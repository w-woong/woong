package delivery

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-wonk/si"
	"github.com/gorilla/mux"
	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/woong/dto"
	"github.com/w-woong/woong/port"
)

type HomeHttpHandler struct {
	timeout             time.Duration
	usc                 port.HomeUsc
	homeGroupProductUsc port.HomeGroupProductUsc
}

func NewHomeHttpHandler(timeout time.Duration, usc port.HomeUsc, homeGroupProductUsc port.HomeGroupProductUsc) *HomeHttpHandler {
	return &HomeHttpHandler{
		timeout:             timeout,
		usc:                 usc,
		homeGroupProductUsc: homeGroupProductUsc,
	}
}

func (d *HomeHttpHandler) HandleAddHome(w http.ResponseWriter, r *http.Request) {
	var home dto.Home
	reqBody := common.HttpBody{
		Document: &home,
	}

	if err := si.DecodeJson(&reqBody, r.Body); err != nil {
		common.HttpError(w, http.StatusBadRequest)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
	rowsAffected, err := d.usc.AddHome(r.Context(), home)
	if err != nil {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	if rowsAffected != 1 {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error("home was not added", logger.UrlField(r.URL.String()))
		return
	}

	if err := common.HttpBodyOK.EncodeTo(w); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
}

func (d *HomeHttpHandler) HandleAddHomeGroupProducts(w http.ResponseWriter, r *http.Request) {
	var home dto.HomeGroupProductList
	reqBody := common.HttpBody{
		Documents: &home,
	}

	if err := si.DecodeJson(&reqBody, r.Body); err != nil {
		common.HttpError(w, http.StatusBadRequest)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
	rowsAffected, err := d.homeGroupProductUsc.AddHomeGroupProducts(r.Context(), home)
	if err != nil {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	if rowsAffected != int64(len(home)) {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error("home group products were not added", logger.UrlField(r.URL.String()))
		return
	}

	if err := common.HttpBodyOK.EncodeTo(w); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
}

func (d *HomeHttpHandler) HandleFindByAppConfigID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	home, err := d.usc.FindByAppConfigID(r.Context(), id)
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			if err := common.HttpBodyRecordNotFound.EncodeTo(w); err != nil {
				logger.Error(err.Error(), logger.UrlField(r.URL.String()))
			}
			return
		}
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	resBody := common.HttpBody{
		Status:   http.StatusOK,
		Count:    1,
		Document: &home,
	}

	if err := resBody.EncodeTo(w); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
}
