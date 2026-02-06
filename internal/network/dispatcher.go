package network

import (
    "fmt"
    "io"
    "net"
    "github.com/leikyz/lkz-online-services/internal/models"
    "github.com/leikyz/lkz-online-services/internal/registries"
)

type Message interface {
    GetID() uint8
    Serialize() ([]byte, error)
    Deserialize(reader io.Reader) error
    // Ensure the interface matches its intended usage
    Process(c *models.Client, conn net.Conn) (*models.Client, error)
    GetMessageSize() uint16
}

var registry = make(map[uint8]func() Message)

func RegisterMessage(id uint8, factory func() Message) {
    registry[id] = factory
}

func HandleMessage(conn net.Conn) {
    defer conn.Close()

    var sessionClient *models.Client 

    for {
        idBuf := make([]byte, 3)
        if _, err := io.ReadFull(conn, idBuf); err != nil {
            break 
        }
        msgID := idBuf[2]
        fmt.Printf("Message ID : %d\n", msgID)
        factory, ok := registry[msgID]
        if !ok {
            fmt.Printf("Message ID inconnu : %d\n", msgID)
            continue
        }
        msg := factory()

        if err := msg.Deserialize(conn); err != nil {
            fmt.Printf("Erreur désérialisation msg %d: %v\n", msgID, err)
            break
        }

        if msgID == 1 {
            if sessionClient != nil {
                fmt.Println("Tentative de re-création de client ignorée")
                continue 
            }
            
            // Pass nil for the client because it does not exist yet
            client, err := msg.Process(nil, conn) 
            if err != nil {
                fmt.Printf("Erreur création client: %v\n", err)
                break
            }
            sessionClient = client 
            continue
        }

        if sessionClient == nil {
            fmt.Printf("Action refusée pour msg %d : client non identifié\n", msgID)
            continue
        }

        // For other messages, pass the sessionClient
        _, err := msg.Process(sessionClient, conn)
        if err != nil {
            fmt.Printf("Erreur process msg %d: %v\n", msgID, err)
            continue 
        }
    }
    
    if sessionClient != nil {
        fmt.Printf("Déconnexion du client : %s\n", sessionClient.ID)
        registries.Clients.Remove(sessionClient.ID)
    }
}