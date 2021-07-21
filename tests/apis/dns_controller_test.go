package apis

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/atlas-dns/apis"
	"github.com/atlas-dns/common"
	"github.com/atlas-dns/models"
	mock_dnsservice "github.com/atlas-dns/tests/service/mock_service"
	"github.com/atlas-dns/tests/test_utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type DnsControllerTestSuite struct {
	suite.Suite
	ctx        *common.Context
	ctrl       *gomock.Controller
	controller *apis.DnsController
	dnsService *mock_dnsservice.MockIDnsService
	gin        *gin.Engine
}

func (suite *DnsControllerTestSuite) BeforeTest(suiteName, testName string) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(suite.T())
	suite.ctx = test_utils.TestContext
	suite.ctrl = ctrl
	suite.dnsService = mock_dnsservice.NewMockIDnsService(ctrl)
	suite.gin = gin.Default()
	suite.controller = apis.NewDnsController(suite.gin.Group("/v1"), suite.dnsService)
}

func (suite *DnsControllerTestSuite) AfterTest(suiteName, testName string) {
	suite.ctrl.Finish()
}

func TestDnsControllerTestSuite(t *testing.T) {
	suite.Run(t, new(DnsControllerTestSuite))
}

func (suite *DnsControllerTestSuite) TestGetDroneLoc() {
	request := models.DroneRequest{
		ID:  "id",
		X:   "1",
		Y:   "2",
		Z:   "3",
		Vel: "4",
	}
	data, _ := json.Marshal(request)

	response := models.DroneResponse{
		Loc: 123.1,
	}

	suite.dnsService.EXPECT().GetDroneLoc(test_utils.EqContextMatcher(suite.ctx), request).Return(response, nil)
	r, _ := http.NewRequest(http.MethodPost, "/v1/drone/loc", bytes.NewBuffer(data))
	w := httptest.NewRecorder()
	suite.gin.ServeHTTP(w, r)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *DnsControllerTestSuite) TestGetDroneLoc_Error() {
	request := models.DroneRequest{
		ID: "id",
		X:  "1",
		Y:  "2",
		Z:  "3",
	}
	data, _ := json.Marshal(request)

	response := models.DroneResponse{
		Loc: 123.1,
	}

	suite.dnsService.EXPECT().GetDroneLoc(test_utils.EqContextMatcher(suite.ctx), request).Return(response, errors.New("error"))
	r, _ := http.NewRequest(http.MethodPost, "/v1/drone/loc", bytes.NewBuffer(data))
	w := httptest.NewRecorder()
	suite.gin.ServeHTTP(w, r)

	suite.Equal(http.StatusInternalServerError, w.Code)
}

func (suite *DnsControllerTestSuite) TestGetDroneLoc_BadRequest() {
	data := `{id:"1", x:1}`

	r, _ := http.NewRequest(http.MethodPost, "/v1/drone/loc", bytes.NewBuffer([]byte(data)))
	w := httptest.NewRecorder()
	suite.gin.ServeHTTP(w, r)

	suite.Equal(http.StatusBadRequest, w.Code)
}
