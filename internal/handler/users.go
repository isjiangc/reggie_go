package handler

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	v1 "reggie_go/api/v1"
	"reggie_go/internal/service"
	"reggie_go/pkg/helper/rand"
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
	session := sessions.Default(ctx)
	codeSession := session.Get(req.Phone)
	if codeSession != nil && codeSession == req.Code {
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
		return
	}
	v1.HandleError(ctx, http.StatusBadRequest, v1.ErrLoginFail, nil)
	return
}

// SendMsg godoc
// @Summary 发送验证码
// @Schemes
// @Description
// @Tags 客户模块
// @Accept json
// @Produce json
// @Param request body v1.SendMsgRequest true "params"
// @Success 200 {object} v1.Response
// @Router /users/sendMsg [post]
func (h *UsersHandler) SendMsg(ctx *gin.Context) {
	req := v1.SendMsgRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	if req.Phone == "" {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrPhoneNumberIsIllegal, nil)
		return
	}
	code := rand.CreateCaptcha()
	h.logger.Info("验证码为:" + code)
	session := sessions.Default(ctx)
	session.Set(req.Phone, code)
	session.Save()
	v1.HandleSuccess(ctx, "手机验证码短信发送成功")
}
