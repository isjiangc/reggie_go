package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "reggie_go/api/v1"
	"reggie_go/internal/model"
	"reggie_go/internal/service"
	"strconv"
)

type DishHandler struct {
	*Handler
	dishService service.DishService
}

func NewDishHandler(handler *Handler, dishService service.DishService) *DishHandler {
	return &DishHandler{
		Handler:     handler,
		dishService: dishService,
	}
}

// CreateDishWithFlavor godoc
// @Summary 新增菜品
// @Schemes
// @Description
// @Tags 菜品管理
// @Accept json
// @Produce json
// @Param request body v1.CreateDishRequest true "params"
// @Success 200 {object} v1.Response
// @Router /dish [post]
func (h *DishHandler) CreateDishWithFlavor(ctx *gin.Context) {
	req := v1.CreateDishRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	session := sessions.Default(ctx)
	userID := session.Get("employee")
	if userID == nil {
		userID = int64(1)
	}
	categorId, _ := strconv.Atoi(string(req.CategoryId))
	// 设置flavor参数
	var flavor model.DishFlavor
	var flavors []model.DishFlavor

	for i := 0; i < len(req.Flavors); i++ {
		flavor.Name = req.Flavors[i].Name
		flavor.Value = req.Flavors[i].Value
		flavor.CreateUser = userID.(int64)
		flavor.UpdateUser = userID.(int64)
	}
	flavors = append(flavors, flavor)
	dishId, err := h.dishService.SaveDishWithFlavor(ctx, model.Dish{
		Name:        req.Name,
		CategoryId:  int64(categorId),
		Price:       req.Price,
		Code:        req.Code,
		Image:       req.Image,
		Description: req.Description,
		Status:      req.Status,
	}, flavors)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	if dishId < 0 {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, "新增菜品成功")

}
