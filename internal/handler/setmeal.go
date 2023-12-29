package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "reggie_go/api/v1"
	"reggie_go/internal/service"
	"strconv"
)

type SetmealHandler struct {
	*Handler
	setmealService service.SetmealService
}

func NewSetmealHandler(handler *Handler, setmealService service.SetmealService) *SetmealHandler {
	return &SetmealHandler{
		Handler:        handler,
		setmealService: setmealService,
	}
}

// GetSetmealList godoc
// @Summary 分页查询
// @Schemes
// @Description
// @Tags 套餐模块
// @Accept json
// @Produce json
// @Param page query string false "页数"
// @Param size query string false "每页数"
// @Param name query string false "套餐名称"
// @Success 200 {object} v1.GetSetmealByPageResponse
// @Router /setmeal/page [get]
func (s *SetmealHandler) GetSetmealList(ctx *gin.Context) {
	pageNum := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("size", "10")
	name := ctx.Query("name")
	page, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	setmealByPage, err := s.setmealService.GetSetmealByPage(ctx, &v1.GetSetmealByPageRequest{
		PageNum:  page,
		PageSize: size,
		Name:     name,
	})
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrInternalServerError, nil)
		return
	}
	v1.HandleSuccess(ctx, &setmealByPage)
}
