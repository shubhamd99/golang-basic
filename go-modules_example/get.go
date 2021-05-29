package main

import (
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"net/http"
	quoteV3 "rsc.io/quote/v3"
)

func GetRouter(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logger, _ := zap.NewProduction()
	logger.Info("success")
	logger.Info(quoteV3.HelloV3())
}
