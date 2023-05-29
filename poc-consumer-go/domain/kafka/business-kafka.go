package domain

type BusinessKafka struct {
	CodigoCliente string `json:"CD_CLIENTE"`
	NomeCliente   string `json:"NM_CLIENTE"`
	CpfCliente    string `json:"CD_CPF_CLIENTE"`
	Telefone      string `json:"TEL_CLIENTE"`
	DataCadastro  string `json:"DT_CADASTRO"`
	Filial        string `json:"CD_FILIAL"`
	Regiao        string `json:"CD_REGIAO"`
	Status        string `json:"STATUS"`
}
