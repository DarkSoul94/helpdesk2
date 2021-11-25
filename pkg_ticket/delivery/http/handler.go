package http

import (
	"net/http"
	"strconv"

	"github.com/DarkSoul94/helpdesk2/dto"
	"github.com/DarkSoul94/helpdesk2/global_const"
	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket"
	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	uc pkg_ticket.ITicketUsecase
}

func NewTicketHandler(uc pkg_ticket.ITicketUsecase) *TicketHandler {
	return &TicketHandler{
		uc: uc,
	}
}

func (h *TicketHandler) CreateCategory(c *gin.Context) {
	var cat dto.InpCategory

	if err := c.BindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	id, err := h.uc.CreateCategory(dto.ToModelCategory(cat))
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok", "category_id": id})
}

func (h *TicketHandler) UpdateCategory(c *gin.Context) {
	var cat dto.InpCategory

	if err := c.BindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	err := h.uc.UpdateCategory(dto.ToModelCategory(cat))
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func (h *TicketHandler) CreateCategorySection(c *gin.Context) {
	var inpSection dto.OutCategorySection

	if err := c.BindJSON(&inpSection); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	id, err := h.uc.CreateCategorySection(dto.ToModelCategorySection(inpSection))
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok", "category_section_id": id})
}

func (h *TicketHandler) UpdateCategorySection(c *gin.Context) {
	var inpSection dto.OutCategorySection

	if err := c.BindJSON(&inpSection); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	err := h.uc.UpdateCategorySection(dto.ToModelCategorySection(inpSection))
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func (h *TicketHandler) GetCategorySection(c *gin.Context) {
	forSearch, err := strconv.ParseBool(c.Request.URL.Query().Get("for_search"))
	if err != nil {
		forSearch = false
	}

	var outSectionList []dto.OutSectionWithCategory
	if sectionList, err := h.uc.GetCategorySection(forSearch); err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	} else {
		for _, sec := range sectionList {
			outSectionList = append(outSectionList, dto.ToOutSectionWithCategory(sec))
		}
	}

	c.JSON(http.StatusOK, outSectionList)
}

func (h *TicketHandler) GetCategorySectionList(c *gin.Context) {
	catWithSec, err := h.uc.GetCategorySectionList()
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	var outList []dto.OutCategoryWithSections

	for _, list := range catWithSec {
		outList = append(outList, dto.ToOutCategoryWithSections(list))
	}

	c.JSON(http.StatusOK, outList)
}

func (h *TicketHandler) CreateRegion(c *gin.Context) {
	var reg dto.InpRegion

	if err := c.BindJSON(&reg); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	id, err := h.uc.CreateRegion(dto.ToModelRegion(reg))
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok", "region_id": id})
}

func (h *TicketHandler) UpdateRegion(c *gin.Context) {
	var reg dto.InpRegion

	if err := c.BindJSON(&reg); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	err := h.uc.UpdateRegion(dto.ToModelRegion(reg))
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func (h *TicketHandler) DeleteRegion(c *gin.Context) {
	regionID, err := strconv.ParseUint(c.Request.URL.Query().Get("region_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	if err := h.uc.DeleteRegion(regionID); err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func (h *TicketHandler) CreateFilial(c *gin.Context) {
	var fil dto.InpFilial

	if err := c.BindJSON(&fil); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	id, err := h.uc.CreateFilial(dto.ToModelFilial(fil))
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok", "filial_id": id})
}

func (h *TicketHandler) UpdateFilial(c *gin.Context) {
	var fil dto.InpFilial

	if err := c.BindJSON(&fil); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	err := h.uc.UpdateFilial(dto.ToModelFilial(fil))
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func (h *TicketHandler) DeleteFilial(c *gin.Context) {
	id, err := strconv.ParseUint(c.Request.URL.Query().Get("filial_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	if err := h.uc.DeleteFilial(id); err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func (h *TicketHandler) GetFilialList(c *gin.Context) {
	list, err := h.uc.GetRegionsWithFilials()
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	var outList []dto.OutRegionWithFilials
	for _, reg := range list {
		outList = append(outList, dto.ToOutRegionWithFilials(reg))
	}

	c.JSON(http.StatusOK, outList)
}

func (h *TicketHandler) GetTicketStatuses(c *gin.Context) {
	user, _ := c.Get(global_const.CtxUserKey)

	list, err := h.uc.GetTicketStatuses(user.(*models.User).Group.ID, false)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	var outList []dto.OutTicketStatus
	for _, stat := range list {
		outList = append(outList, dto.ToOutTicketStatus(stat))
	}

	c.JSON(http.StatusOK, outList)
}

func (h *TicketHandler) GetAllTicketStatuses(c *gin.Context) {
	user, _ := c.Get(global_const.CtxUserKey)

	list, err := h.uc.GetTicketStatuses(user.(*models.User).Group.ID, true)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	var outList []dto.OutTicketStatus
	for _, stat := range list {
		outList = append(outList, dto.ToOutTicketStatus(stat))
	}

	c.JSON(http.StatusOK, outList)
}

func (h *TicketHandler) GetTicketStatusHistory(c *gin.Context) {
	ticketID, _ := strconv.ParseUint(c.Request.URL.Query().Get("ticket_id"), 10, 64)
	user, _ := c.Get(global_const.CtxUserKey)

	historyList, err := h.uc.GetAllTicketStatusHistory(ticketID, user.(*models.User))
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	outList := make([]dto.OutTicketStatusHistory, 0)
	for _, history := range historyList {
		outList = append(outList, dto.ToOutTicketStatusHistory(history))
	}

	c.JSON(http.StatusOK, outList)
}

func (h *TicketHandler) CreateTicket(c *gin.Context) {
	var ticket dto.NewTicket

	if err := c.BindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	user, _ := c.Get(global_const.CtxUserKey)
	ticket.Author = dto.OutUser{
		ID:    user.(*models.User).ID,
		Name:  user.(*models.User).Name,
		Email: user.(*models.User).Email,
	}
	ticket.Ip = c.GetHeader("X-Forwarded-For")

	id, err := h.uc.CreateTicket(dto.NewTicketToModelTicket(ticket))
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok", "ticket_id": id})
}

func (h *TicketHandler) GetTicketsList(c *gin.Context) {
	count, _ := strconv.Atoi(c.Request.URL.Query().Get("count"))
	offset, _ := strconv.Atoi(c.Request.URL.Query().Get("offset"))
	user, _ := c.Get(global_const.CtxUserKey)

	list, tags, priority, err := h.uc.GetTicketList(user.(*models.User), count, offset)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	outList := make([]dto.OutTicketForList, 0)
	for _, tick := range list {
		outList = append(outList, dto.ToOutTicketForList(tick, priority))
	}

	c.JSON(http.StatusOK, map[string]interface{}{"fields": tags, "tickets": outList})
}

func (h *TicketHandler) GetTicket(c *gin.Context) {
	ticketID, _ := strconv.ParseUint(c.Request.URL.Query().Get("ticket_id"), 10, 64)
	user, _ := c.Get(global_const.CtxUserKey)

	ticket, err := h.uc.GetTicket(ticketID, user.(*models.User))
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToOutTicket(ticket))
}

func (h *TicketHandler) CreateComment(c *gin.Context) {
	var comment dto.InpComment

	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "error": err.Error()})
		return
	}

	user, _ := c.Get(global_const.CtxUserKey)

	mComment := dto.ToModelComment(comment)
	mComment.Author = user.(*models.User)

	id, err := h.uc.CreateComment(mComment)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok", "comment_id": id})
}

func (h *TicketHandler) CheckNeedApprovalTicketExist(c *gin.Context) {
	user, _ := c.Get(global_const.CtxUserKey)

	exist := h.uc.CheckNeedApprovalTicketExist(user.(*models.User).Group.ID)

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok", "exist": exist})
}

func (h *TicketHandler) GetApprovalTicketList(c *gin.Context) {
	count, _ := strconv.Atoi(c.Request.URL.Query().Get("count"))
	offset, _ := strconv.Atoi(c.Request.URL.Query().Get("offset"))
	user, _ := c.Get(global_const.CtxUserKey)

	list, tags, err := h.uc.GetApprovalTicketList(user.(*models.User).Group.ID, count, offset)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	outList := make([]dto.OutTicketForList, 0)
	for _, tick := range list {
		outList = append(outList, dto.ToOutTicketForList(tick, nil))
	}

	c.JSON(http.StatusOK, map[string]interface{}{"fields": tags, "tickets": outList})
}
