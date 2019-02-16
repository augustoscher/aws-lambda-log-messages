package model

//MensagemErro representa uma mensagem de erro
type MensagemErro struct {
	ID               string `json:"id"`
	Idintegrador     string `json:"idintegrador"`
	Filial           string `json:"filial"`
	Codintegracao    string `json:"codintegracao"`
	Descricaoerro    string `json:"descricaoerro"`
	Conteudomensagem string `json:"conteudomensagem"`
	Datahora         string `json:"datahora"`
}
