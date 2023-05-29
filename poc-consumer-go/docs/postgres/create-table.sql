create table ClientesVendas(
	CODIGO_CLIENTE INT not null,
	NOME_CLIENTE VARCHAR(44) not null,
	TELEFONE VARCHAR(22) not null,
	DATA_CADASTRO DATE not null,
	FILIAL INT not null,
	REGIAO INT not null,
	STATUS VARCHAR(1))