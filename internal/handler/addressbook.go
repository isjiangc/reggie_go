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

type AddressbookHandler struct {
	*Handler
	addressbookService service.AddressbookService
}

func NewAddressbookHandler(handler *Handler, addressbookService service.AddressbookService) *AddressbookHandler {
	return &AddressbookHandler{
		Handler:            handler,
		addressbookService: addressbookService,
	}
}

// GetAddressbook godoc
// @Summary 通过Id获取用户地址信息
// @Schemes
// @Description
// @Tags 地址模块
// @Accept json
// @Produce json
// @Param id path string true "用户Id"
// @Success 200 {object} v1.Response
// @Router /addressBook/list/{id} [get]
func (h *AddressbookHandler) GetAddressbook(ctx *gin.Context) {
	id := ctx.Param("id")
	Id, _ := strconv.Atoi(id)
	getAddressBookData, err := h.addressbookService.GetAddressbook(ctx, &v1.GetAddressBookByUserIdRequest{
		UserId: int64(Id),
	})
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, &getAddressBookData)
}

// UpdateAddressIsDefault godoc
// @Summary 通过UserId和地址Id设置默认地址
// @Schemes
// @Description
// @Tags 地址模块
// @Accept json
// @Produce json
// @Param request body v1.UpdateAddressBookIsDefaultRequest true "params"
// @Success 200 {object} v1.Response
// @Router /addressBook/default [put]
func (h *AddressbookHandler) UpdateAddressIsDefault(ctx *gin.Context) {
	req := v1.UpdateAddressBookIsDefaultRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}
	session := sessions.Default(ctx)
	userID := session.Get("employee")
	// 防止session丢失导致更新失败。默认使用管理员更新
	if userID == nil {
		userID = int64(1)
	}
	err := h.addressbookService.UpdataAddressIsDefault(ctx, &v1.UpdateAddressBookIsDefaultRequest{
		UserId: req.UserId,
		Id:     req.Id,
	}, time.Now(), userID.(int64))
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, "设置默认地址成功")
}
