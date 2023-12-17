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
	for _, v := range req.Flavors {
		flavor.Name = v.Name
		flavor.Value = v.Value
		flavors = append(flavors, flavor)
	}
	dishId, err := h.dishService.SaveDishWithFlavor(ctx, model.Dish{
		Name:        req.Name,
		CategoryId:  int64(categorId),
		Price:       req.Price,
		Code:        req.Code,
		Image:       req.Image,
		Description: req.Description,
		Status:      req.Status,
		CreateUser:  userID.(int64),
		UpdateUser:  userID.(int64),
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

// GetDishList godoc
// @Summary 分页查询
// @Schemes
// @Description
// @Tags 菜品管理
// @Accept json
// @Produce json
// @Param page query string false "页数"
// @Param size query string false "每页数"
// @Param name query string false "菜品名称"
// @Success 200 {object} v1.GetDishByPageResponse
// @Router /dish/page [get]
func (h *DishHandler) GetDishList(ctx *gin.Context) {
	pageNum := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("size", "10")
	name := ctx.Query("name")
	page, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	dishByPage, err := h.dishService.GetDishByPage(ctx, &v1.GetDishByPageRequest{
		PageNum:  page,
		PageSize: size,
		Name:     name,
	})
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrInternalServerError, nil)
		return
	}
	v1.HandleSuccess(ctx, &dishByPage)
}

// GetDishById godoc
// @Summary 通过id获取菜品信息
// @Schemes
// @Description
// @Tags 菜品管理
// @Accept json
// @Produce json
// @Param id path string true "菜品Id"
// @Success 200 {object} v1.GetDishByIdResponse
// @Router /dish/{id} [get]
func (h *DishHandler) GetDishById(ctx *gin.Context) {
	id := ctx.Param("id")
	disId, _ := strconv.Atoi(id)
	dishByIdData, err := h.dishService.GetDishById(ctx, &v1.GetDishByIdRequest{
		Id: int64(disId),
	})
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, &dishByIdData)

}
