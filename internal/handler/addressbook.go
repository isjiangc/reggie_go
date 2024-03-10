package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	v1 "reggie_go/api/v1"
	"reggie_go/internal/service"
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

// GetAddressbookByUserId godoc
// @Summary 通过用户Id获取用户地址信息
// @Schemes
// @Description
// @Tags 地址模块
// @Accept json
// @Produce json
// @Param userid path string true "用户Id"
// @Success 200 {object} v1.Response
// @Router /addressBook/list/{userid} [get]
func (h *AddressbookHandler) GetAddressbookByUserId(ctx *gin.Context) {
	id := ctx.Param("userid")
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

// GetAddressBookById godoc
// @Summary 通过Id获取用户地址信息
// @Schemes
// @Description
// @Tags 地址模块
// @Accept json
// @Produce json
// @Param id path string true "Id"
// @Success 200 {object} v1.Response
// @Router /addressBook/{id} [get]
func (h *AddressbookHandler) GetAddressBookById(ctx *gin.Context) {
	id := ctx.Param("id")
	Id, _ := strconv.Atoi(id)
	addressbook, err := h.addressbookService.GetAddressById(ctx, &v1.GetAddressBookByIdRequest{
		Id: int64(Id),
	})
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, addressbook)
}

// SaveAddressBook godoc
// @Summary 新增地址
// @Schemes
// @Description
// @Tags 地址模块
// @Accept json
// @Produce json
// @Param request body v1.SaveAddressBookRequest true "params"
// @Success 200 {object} v1.Response
// @Router /addressBook [post]
func (h *AddressbookHandler) SaveAddressBook(ctx *gin.Context) {
	req := v1.SaveAddressBookRequest{}
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
	err := h.addressbookService.SaveAddressBook(ctx, &v1.SaveAddressBookRequest{
		AddressBook: v1.AddressBook{
			UserId:     req.UserId,
			Consignee:  req.Consignee,
			Sex:        req.Sex,
			Phone:      req.Phone,
			Detail:     req.Detail,
			Label:      req.Label,
			CreateUser: userID.(int64),
			UpdateUser: userID.(int64),
		},
	})
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, "新增地址成功")
}

// GetDefaultAddressBook godoc
// @Summary 通过用户Id获取用户默认地址
// @Schemes
// @Description
// @Tags 地址模块
// @Accept json
// @Produce json
// @Param userid path string true "用户Id"
// @Success 200 {object} v1.Response
// @Router /addressBook/default/{userid} [get]
func (h *AddressbookHandler) GetDefaultAddressBook(ctx *gin.Context) {
	id := ctx.Param("userid")
	Id, _ := strconv.Atoi(id)
	getAddressBookData, err := h.addressbookService.GetDefaultAddressBook(ctx, &v1.GetDefaultAddressBookRequest{
		UserId: int64(Id),
	})
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, &getAddressBookData)
}
