package delivery

import (
	"net/http"
	"time"

	"github.com/go-wonk/si/v2"
	"github.com/gorilla/mux"
	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/woong/dto"
	"github.com/w-woong/woong/port"
)

type AppConfigHttpHandler struct {
	timeout time.Duration
	usc     port.AppConfigUsc
}

func NewAppConfigHttpHandler(timeout time.Duration, usc port.AppConfigUsc) *AppConfigHttpHandler {
	return &AppConfigHttpHandler{
		timeout: timeout,
		usc:     usc,
	}
}

func (d *AppConfigHttpHandler) HandleAddAppConfig(w http.ResponseWriter, r *http.Request) {
	var appConfig dto.AppConfig
	reqBody := common.HttpBody{
		Document: &appConfig,
	}

	if err := si.DecodeJson(&reqBody, r.Body); err != nil {
		common.HttpError(w, http.StatusBadRequest)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	rowsAffected, err := d.usc.AddAppConfig(r.Context(), appConfig)
	if err != nil {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	if rowsAffected != 1 {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error("app config was not added", logger.UrlField(r.URL.String()))
		return
	}

	if err := common.HttpBodyOK.EncodeTo(w); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
}

func (d *AppConfigHttpHandler) HandleFindAppConfig(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	appConfig, err := d.usc.FindAppConfig(r.Context(), id)
	if err != nil {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	var resBody common.HttpBody
	resBody.Status = http.StatusOK
	resBody.Count = 1
	resBody.Document = &appConfig

	if err := resBody.EncodeTo(w); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
}
