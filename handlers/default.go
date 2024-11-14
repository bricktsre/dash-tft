package handlers

import (
	"context"
	"net/http"

	"github.com/bricktsre/tftcalculator/session"
)

type XService interface {

}

func New(log *slog.Logger, xs XService) *Default Handler {
	return &DefaultHandler{
		Log: 	  log,
		XService: xs,
	}
}

type DefaultHandler struct {
	log		 *slog.Logger
	XService XService
}

