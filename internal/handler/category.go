package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "reggie_go/api/v1"
	"reggie_go/internal/service"
	"strconv"
	"time"
)

type CategoryHandler struct {
	*Handler
	categoryService service.CategoryService
}

func NewCategoryHandler(handler *Handler, categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		Handler:         handler,
		categoryService: categoryService,
	}
}

// CreateCategory godoc
// @Summary 新增分类
// @Schemes
// @Description
// @Tags 分类模块
// @Accept json
// @Produce json
// @Param request body v1.CreateCategoryRequest true "params"
// @Success 200 {object} v1.Response
// @Router /category [post]
func (h *CategoryHandler) CreateCategory(ctx *gin.Context) {
	req := v1.CreateCategoryRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	cateType, _ := strconv.Atoi(req.Type)
	cateSort, _ := strconv.Atoi(req.Sort)
	session := sessions.Default(ctx)
	userID := session.Get("employee")
	if userID == nil {
		userID = int64(1)
	}
	err := h.categoryService.CreateCategory(ctx, &service.CreateCategoryData{
		Type:       cateType,
		Name:       req.Name,
		Sort:       cateSort,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		CreateUser: userID.(int64),
		UpdateUser: userID.(int64),
	})
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, "新增分类成功")
}
