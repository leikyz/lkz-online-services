package network


import (
	"fmt"
	"net/http"
	"github.com/leikyz/lkz-online-services/internal/network/handlers"
	"github.com/leikyz/lkz-online-services/internal/models"
)

func RegisterHandlers() {
	http.HandleFunc("/hello", handlers.SayHello)
}

func CreatePlayer() {
	player := models.Player{
		ID:       "1",
		Username: "PlayerOne",
		MMR:      1500,
	}

	fmt.Println("Created player:", player)
}

