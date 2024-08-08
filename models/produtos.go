package models

import "go-web-app/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscarTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos;")
	if err != nil {
		panic("Select error:" + err.Error())
	}
	produto := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

		produtos = append(produtos, produto)
	}

	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	insereDadosNoBanco, err := db.Prepare("INSERT INTO PRODUTOS (NOME, DESCRICAO, PRECO, QUANTIDADE) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
}

func DeletarProduto(id string) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	deletarProduto, err := db.Prepare("DELETE FROM PRODUTOS WHERE ID = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)
}

func EditaProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	ProdutoDoBanco, err := db.Query("SELECT * FROM PRODUTOS WHERE ID=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for ProdutoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = ProdutoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}

	return produtoParaAtualizar
}

func AtualizaProduto(id, quantidade int, nome, descricao string, preco float64) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	AtualizaProduto, err := db.Prepare("UPDATE PRODUTOS SET NOME=$1, DESCRICAO=$2, PRECO=$3, QUANTIDADE=$4 WHERE ID=$5")
	if err != nil {
		panic(err.Error())
	}

	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
}
