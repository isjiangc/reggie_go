package handler

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "reggie_go/api/v1"
	"reggie_go/internal/service"
)

type EmployeeHandler struct {
	*Handler
	employeeService service.EmployeeService
}

func NewEmployeeHandler(handler *Handler, employeeService service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		Handler:         handler,
		employeeService: employeeService,
	}
}

// Login godoc
// @Summary 员工登录
// @Schemes
// @Description
// @Tags 员工模块
// @Accept json
// @Produce json
// @Param request body v1.EmployeeLoginRequest true "params"
// @Success 200 {object} v1.EmployeeLoginResponse
// @Router /employee/login [post]
func (h *EmployeeHandler) Login(ctx *gin.Context) {
	req := v1.EmployeeLoginRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	res, err := h.employeeService.Login(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrInternalServerError, nil)
		return
	}
	session := sessions.Default(ctx)
	session.Set("employee", res.Id)
	session.Save()
	v1.HandleSuccess(ctx, &res)
}

// Logout godoc
// @Summary 员工退出
// @Schemes
// @Description
// @Tags 员工模块
// @Accept json
// @Produce json
// @Success 200 {object} v1.Response
// @Router /employee/logout [post]
func (h *EmployeeHandler) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("employee")
	v1.HandleSuccess(ctx, "退出成功")
}

// Save godoc
// @Summary 新增员工
// @Schemes
// @Description
// @Tags 员工模块
// @Accept json
// @Produce json
// @Param request body v1.CreateEmployeeRequest true "params"
// @Success 200 {object} v1.Response
// @Router /employee [post]
func (h *EmployeeHandler) Save(ctx *gin.Context) {
	req := v1.CreateEmployeeRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	session := sessions.Default(ctx)
	userID := session.Get("employee")
	fmt.Println(userID)
	err := h.employeeService.CreateEmployee(ctx, &service.CreateEmployeeData{
		Name:       req.Name,
		Username:   req.Username,
		Phone:      req.Phone,
		Sex:        req.Sex,
		IdNumber:   req.IdNumber,
		CreateUser: userID.(int64),
		UpdateUser: userID.(int64),
	})
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, "新增员工成功")
}
