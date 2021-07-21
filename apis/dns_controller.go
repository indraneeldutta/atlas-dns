package apis

import (
	"errors"
	"net/http"

	"github.com/atlas-dns/common"
	"github.com/atlas-dns/models"
	dnsservice "github.com/atlas-dns/services/dns_service"
	"github.com/gin-gonic/gin"
)

const (
	GETLOC = "/drone/loc"
)

// DndController ..
type DnsController struct {
	service dnsservice.IDnsService
}

// NewDnsController initialises all the routes for Dns Services
func NewDnsController(router *gin.RouterGroup, service dnsservice.IDnsService) *DnsController {
	controller := DnsController{
		service: service,
	}

	router.POST(GETLOC, controller.getDroneLoc)

	return &controller
}

func (c *DnsController) getDroneLoc(ginCtx *gin.Context) {
	context := common.CreateLoggableContextFromRequest(ginCtx.Request, common.Log)
	context.Logger.Infof("request received for POST: %v", GETLOC)

	var request models.DroneRequest
	if errors.Is(ginCtx.ShouldBindJSON(&request), nil) {
		response, err := c.service.GetDroneLoc(context, request)
		if !errors.Is(err, nil) {
			ginCtx.SecureJSON(http.StatusInternalServerError, gin.H{})
			context.Logger.Error(err)
			return
		}
		ginCtx.SecureJSON(http.StatusOK, gin.H{
			"data": response,
		})
		return
	}
	ginCtx.SecureJSON(http.StatusBadRequest, gin.H{})
}
