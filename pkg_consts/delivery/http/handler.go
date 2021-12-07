package http

import (
	"net/http"

	"github.com/DarkSoul94/helpdesk2/pkg_consts"
	"github.com/gin-gonic/gin"
)

type ConstsHandler struct {
	uc pkg_consts.IConstsUsecase
}

func NewConstsHandler(uc pkg_consts.IConstsUsecase) *ConstsHandler {
	return &ConstsHandler{
		uc: uc,
	}
}

func (h *ConstsHandler) SetConst(c *gin.Context) {
	data := make(map[string]interface{})
	key := c.Param("const")

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	err := h.uc.SetConst(key, data)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func (h *ConstsHandler) GetConst(c *gin.Context) {
	key := c.Param("const")

	data, err := h.uc.GetConst(key)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *ConstsHandler) GetSettings(c *gin.Context) {
	data, err := h.uc.GetConst(pkg_consts.KeyConfig)
	if err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *ConstsHandler) SetSettings(c *gin.Context) {
	var data = make(map[string]interface{})

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}

	if err := h.uc.SetConst(pkg_consts.KeyConfig, data); err != nil {
		c.JSON(err.Code(), map[string]interface{}{"status": "error", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})

}
