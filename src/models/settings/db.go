package settings

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // O underscore importa o pacote apenas para seus efeitos colaterais (registro do driver)
)

type DBConnectionHanddler struct {
	DB *sql.DB
}

const connStr = "port=5432 user=postgres password=password dbname=postgres sslmode=disable"

// Conecta o DB
func DBConnect() (*DBConnectionHanddler, error) {
	// Settings para conexão do DB
	fmt.Println(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão ao banco de dados: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	return &DBConnectionHanddler{DB: db}, nil
}

// Fecha a conexão
func (handler *DBConnectionHanddler) Close() {
	if err := handler.DB.Close(); err != nil {
		log.Printf("erro ao fechar a conexão com o Banco de Dados: %v", err)
	}
}
