package handlers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net/http"

	client "github.com/leikyz/lkz-online-services/internal/clients"
)

type ConnectClientRequest struct {
	MessageID uint8
}

func ConnectClient(w http.ResponseWriter, r *http.Request) {

	// Read data
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Deserialize
	reader := bytes.NewReader(body)
	var req ConnectClientRequest

	// Read data according buffer order
	binary.Read(reader, binary.LittleEndian, &req.MessageID)

	fmt.Printf("ID: %d", req.MessageID)

	// Create client
	client.ClientManager.CreateClient("guest", 1)

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, uint8(1))

	// Send to client
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(buf.Bytes())
}
