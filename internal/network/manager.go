package network

import (
	"net/http"

	"github.com/leikyz/lkz-online-services/internal/network/handlers"
)

func RegisterHandlers(mux *http.ServeMux) {

	mux.HandleFunc("/hello", handlers.SayHello)
}
