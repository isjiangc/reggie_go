package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "reggie_go/api/v1"
	"reggie_go/internal/service"
)

type UsersHandler struct {
	*Handler
	usersService service.UsersService
}

func NewUsersHandler(handler *Handler, usersService service.UsersService) *UsersHandler {
	return &UsersHandler{
		Handler:      handler,
		usersService: usersService,
	}
}

// UsersLogin godoc
// @Summary 用户登录
// @Schemes
// @Description
// @Tags 客户模块
// @Accept json
// @Produce json
// @Param request body v1.UserLoginRequest true "params"
// @Success 200 {object} v1.Response
// @Router /users/login [post]
func (h *UsersHandler) UsersLogin(ctx *gin.Context) {
	req := v1.UserLoginRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	users, err := h.usersService.UserLogin(ctx, &v1.UserLoginRequest{
		Phone: req.Phone,
		Code:  req.Code,
	})
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrInternalServerError, nil)
		return
	}
	session := sessions.Default(ctx)
	session.Set("users", users.Phone)
	session.Save()
	v1.HandleSuccess(ctx, &users)
}
