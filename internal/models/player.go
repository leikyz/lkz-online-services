package models

type Player struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	MMR      int    `json:"mmr"`
}
