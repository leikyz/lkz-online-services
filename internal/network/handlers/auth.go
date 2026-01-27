package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/leikyz/lkz-online-services/internal/models"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello request received")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Read error :", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if len(body) == 0 {
		fmt.Println("Empty body received")
		return
	}

	messageID := body[0]
	fmt.Printf("Message received ! ID: %d | Size: %d bytes\n", messageID, len(body))

	responseID := byte(1)
	w.Header().Set("Content-Type", "application/octet-stream")

	_, err = w.Write([]byte{responseID})

	if err != nil {
		fmt.Println("Error writing response:", err)
	} else {
		fmt.Printf("Response sent: ID %d\n", responseID)
	}
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	client := models.Client{
		ID:        0,
		IPAddress: r.RemoteAddr,
	}

	fmt.Printf("Client created %v", client)
}
