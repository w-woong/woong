package delivery_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/w-woong/woong/delivery"
	"github.com/w-woong/woong/dto"
	"github.com/w-woong/woong/port/mocks"
)

var (
	dnt, _    = time.Parse("20060102150405", "20221120121212")
	dnt2, _   = time.ParseInLocation("20060102150405", "20221120121212", time.Local)
	appConfig = dto.AppConfig{
		ID:        "88ca4470-7690-4073-a0c8-5cc3b84bc2a6",
		CreatedAt: &dnt2,
		UpdatedAt: &dnt2,
		Name:      "Test-Woong-App2",
	}
)

func Test_AppConfigHttpHandler_HandleAddAppConfig(t *testing.T) {
	urlPath := "/v1/woong/appconfig/{id}"

	ctrl := gomock.NewController(t)
	usc := mocks.NewMockAppConfigUsc(ctrl)
	usc.EXPECT().FindAppConfig(gomock.Any(), appConfig.ID).Return(appConfig, nil).AnyTimes()
	handler := delivery.NewAppConfigHttpHandler(3*time.Second, usc)

	router := mux.NewRouter()
	router.HandleFunc(urlPath, handler.HandleFindAppConfig).Methods(http.MethodGet)

	req, err := http.NewRequest(http.MethodGet, "/v1/woong/appconfig/"+appConfig.ID, nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	resBody, err := io.ReadAll(rr.Body)
	assert.Nil(t, err)
	fmt.Println(string(resBody))

}
