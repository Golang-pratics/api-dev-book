package model

// Senha representa o formato da requisição de alteração senha
type Senha struct {
	Nova string `json:"nova"`
	Atual string `json:"atual"`
}