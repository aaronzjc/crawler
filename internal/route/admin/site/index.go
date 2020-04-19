package site

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"mu/internal/model"
	"mu/internal/svc/schedule"
	"mu/internal/util/logger"
	"mu/internal/util/req"
)

type InfoForm struct {
	Id int `form:"id" binding:"required"`
}

type ListForm struct {
	Keyword string `form:"keyword"`
}

type UpdateForm model.SiteJson

func Info(c *gin.Context) {
	var r InfoForm
	if err := c.ShouldBindQuery(&r); err != nil {
		req.JSON(c, req.CodeError, "参数异常", nil)
		return
	}

	m := &model.Site{
		ID: r.Id,
	}

	site, err := m.FetchInfo()
	if err != nil || site.ID <= 0 {
		req.JSON(c, req.CodeError, "失败", nil)
		return
	}

	jsonObj, err := site.FormatJson()
	if err != nil {
		req.JSON(c, req.CodeError, err.Error(), nil)
		return
	}
	req.JSON(c, req.CodeSuccess, "成功", jsonObj)
	return
}

func List(c *gin.Context) {
	var r ListForm
	var err error
	if err = c.ShouldBind(&r); err != nil {
		req.JSON(c, req.CodeError, "参数异常", nil)
		return
	}

	m := &model.Site{}
	var sites []model.Site
	if r.Keyword != "" {
		sites, err = m.FetchRows(model.Query{
			Query: "name like ?",
			Args:  []interface{}{"%" + r.Keyword + "%"},
		})
	} else {
		sites, err = m.FetchRows(model.Query{})
	}
	if err != nil {
		req.JSON(c, req.CodeError, err.Error(), nil)
		return
	}

	var result []model.SiteJson
	for _, site := range sites {

		item, err := site.FormatJson()
		if err != nil {
			req.JSON(c, req.CodeError, err.Error(), nil)
			return
		}
		result = append(result, item)
	}

	nodes, _ := (&model.Node{}).FetchRows(model.Query{
		Query: "`enable` = ?",
		Args:  []interface{}{model.Enable},
	})
	nodeJson := make(map[int]model.Node)
	for _, node := range nodes {
		n, _ := node.FormatJson()
		nodeJson[node.ID] = n
	}

	req.JSON(c, req.CodeSuccess, "成功", map[string]interface{}{
		"nodeList": nodeJson,
		"siteList": result,
	})
	return
}

func UpdateSite(c *gin.Context) {
	var err error
	var r UpdateForm
	if err = c.ShouldBindJSON(&r); err != nil {
		req.JSON(c, req.CodeError, "参数异常 "+err.Error(), nil)
		return
	}

	tagBytes, _ := json.Marshal(r.Tags)
	hostsBytes, _ := json.Marshal(r.NodeHosts)
	reqHeaders, _ := json.Marshal(r.ReqHeaders)

	m := model.Site{
		ID:         r.ID,
		Name:       r.Name,
		Key:        r.Key,
		Desc:       r.Desc,
		Cron:       r.Cron,
		Tags:       string(tagBytes),
		Enable:     r.Enable,
		NodeOption: r.NodeOption,
		NodeType:   r.NodeType,
		NodeHosts:  string(hostsBytes),
		ReqHeaders: string(reqHeaders),
	}
	err = m.CheckArgs()
	if err != nil {
		req.JSON(c, req.CodeError, err.Error(), nil)
		return
	}

	data := make(map[string]interface{})
	data["name"] = m.Name
	data["key"] = m.Key
	data["desc"] = m.Desc
	data["cron"] = m.Cron
	data["tags"] = m.Tags
	data["enable"] = m.Enable
	data["node_option"] = m.NodeOption
	data["node_type"] = m.NodeType
	data["node_hosts"] = m.NodeHosts
	data["req_headers"] = m.ReqHeaders

	om, _ := (&model.Site{}).FetchRow(model.Query{
		Query: "`id` = ?",
		Args:  []interface{}{m.ID},
	})

	err = m.Update(data)
	if err != nil {
		req.JSON(c, req.CodeError, "更新站点失败", nil)
		return
	}

	// 检查当前更新状态，操作Job
	s := om.Enable == m.Enable
	if !s {
		if m.Enable == model.Enable {
			// add
			_ = schedule.JobSchedule.AddJob(m)
		} else {
			// delete
			schedule.JobSchedule.RemoveJob(m.Key)
		}
		logger.Info("[%s] cron updated to %s ", m.Key, m.Cron)
	} else {
		if m.Enable == model.Enable && m.Cron != om.Cron {
			// update
			schedule.JobSchedule.RemoveJob(m.Key)
			if err = schedule.JobSchedule.AddJob(m); err == nil {
				logger.Info("[%s] cron updated to %s ", m.Key, m.Cron)
			}
		}
	}

	req.JSON(c, req.CodeSuccess, "成功", nil)
	return
}

func Debug(c *gin.Context) {
	result := schedule.Debug()
	req.JSON(c, req.CodeSuccess, "成功", result)
	return
}
