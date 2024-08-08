package db

import "database/sql"

func ConectaComBancoDeDados() *sql.DB {
	conexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
