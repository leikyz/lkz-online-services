package models

type Client struct {
	ID        int `json:"id"`
	IPAddress string `json:"ip_address"`
}
