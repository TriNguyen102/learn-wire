package handler

import (
	"learn-wire/service"
	"log"
	"net/http"
)

func NewMyHandler(svc *service.MyService) *MyHandler {
	return &MyHandler{svc: svc}
}

type MyHandler struct {
	svc *service.MyService
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	isExisted, errorConnect := h.svc.SomeMethod()

	if !isExisted {
		log.Println("Database is already Existed!")
	}

	if errorConnect != nil {
		log.Println("Connect to Database fail!")
	}
}
