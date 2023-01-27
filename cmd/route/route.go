package route

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/w-woong/common"
	"github.com/w-woong/common/middlewares"
	"github.com/w-woong/woong/delivery"
	"github.com/w-woong/woong/port"
)

func HomeRoute(router *mux.Router, conf common.ConfigHttp, usc port.HomeUsc) *delivery.HomeHttpHandler {

	handler := delivery.NewHomeHttpHandler(time.Duration(conf.Timeout)*time.Second, usc)

	router.HandleFunc("/v1/woong/home",
		middlewares.AuthBearerHandler(handler.HandleAddHome, conf.BearerToken)).Methods(http.MethodPost)

	router.HandleFunc("/v1/woong/home/appconfig/{id}",
		middlewares.AuthBearerHandler(handler.HandleFindByAppConfigID, conf.BearerToken),
	).Methods(http.MethodGet)

	return handler
}

func AppConfigRoute(router *mux.Router, conf common.ConfigHttp, usc port.AppConfigUsc) *delivery.AppConfigHttpHandler {

	handler := delivery.NewAppConfigHttpHandler(time.Duration(conf.Timeout)*time.Second, usc)
	router.HandleFunc("/v1/woong/appconfig",
		middlewares.AuthBearerHandler(handler.HandleAddAppConfig, conf.BearerToken),
	).Methods(http.MethodPost)

	router.HandleFunc("/v1/woong/appconfig/{id}",
		middlewares.AuthBearerHandler(handler.HandleFindAppConfig, conf.BearerToken),
	).Methods(http.MethodGet)

	return handler
}

// func HomeRoute(router *mux.Router, conf common.ConfigHttp,
// 	validator commonport.IDTokenValidators, usc port.CartUsc, userSvc commonport.UserSvc) *delivery.CartHttpHandler {

// 	handler := delivery.NewCartHttpHandler(time.Duration(conf.Timeout)*time.Second, usc)

// 	router.HandleFunc("/v1/order/cart", middlewares.AuthIDTokenUserAccountHandler(
// 		handler.HandleFindByUserID, validator, userSvc,
// 	)).Methods(http.MethodGet)
// 	router.HandleFunc("/v1/order/cart/_find-or-create", middlewares.AuthIDTokenUserAccountHandler(
// 		handler.HandleFindOrCreateByUserID, validator, userSvc,
// 	)).Methods(http.MethodGet)

// 	router.HandleFunc("/test/order/cart", middlewares.AuthIDTokenUserAccountHandler(
// 		handler.HandleTestRefreshError, validator, userSvc,
// 	)).Methods(http.MethodGet)

// 	router.HandleFunc("/v1/order/cart/product", middlewares.AuthIDTokenUserAccountHandler(
// 		handler.HandleAddCartProduct, validator, userSvc,
// 	)).Methods(http.MethodPost)

// 	return handler
// }
