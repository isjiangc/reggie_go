package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "reggie_go/api/v1"
	"reggie_go/internal/service"
	"strconv"
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
