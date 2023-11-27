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

// GetCategoryList godoc
// @Summary 分页查询
// @Schemes
// @Description
// @Tags 分类模块
// @Accept json
// @Produce json
// @Param page query string false "页数"
// @Param size query string false "每页数"
// @Success 200 {object} v1.GetCategoryPageResponse
// @Router /category/page [get]
func (h *CategoryHandler) GetCategoryList(ctx *gin.Context) {
	pageNum := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("size", "10")
	page, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	categoryPage, err := h.categoryService.GetCategoryPage(ctx, &v1.GetCategoryPageRequest{
		PageNum:  page,
		PageSize: size,
	})
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrInternalServerError, nil)
		return
	}
	v1.HandleSuccess(ctx, &categoryPage)
}

// DeleteCategory godoc
// @Summary 删除分类
// @Schemes
// @Description
// @Tags 分类模块
// @Accept json
// @Produce json
// @Param id query string false "分类ID"
// @Success 200 {object} v1.Response
// @Router /category [delete]
func (h *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	Id := ctx.Query("id")
	if Id == "" {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrDeleteCategoryIdIsNotEmpty, nil)
		return
	}
	id, _ := strconv.Atoi(Id)
	err := h.categoryService.DeleteCategory(ctx, &v1.DeleteCategoryRequest{
		Id: int64(id),
	})
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, "删除分类成功")
}
