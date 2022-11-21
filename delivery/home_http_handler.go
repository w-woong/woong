package delivery

import (
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/woong/port"
)

type HomeHttpHandler struct {
	timeout time.Duration
	usc     port.HomeUsc
}

func NewHomeHttpHandler(timeout time.Duration, usc port.HomeUsc) *HomeHttpHandler {
	return &HomeHttpHandler{
		timeout: timeout,
		usc:     usc,
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
