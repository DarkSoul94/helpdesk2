package http

import (
	"net/http"
	"strconv"

	"github.com/DarkSoul94/helpdesk2/dto"
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
