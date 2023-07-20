package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


var(
	//StringConexaoBanco é a string de conexão com o banco de dados
	StringConexaoBanco = ""
	//porta onde a api está rodando
	Porta = 0
	// chave usada para assinar o token
	SecretKey []byte
	
)
// carregar vai inicializar as variaves de ambiente
func Carregar(){
	var erro error

	if erro = godotenv.Load(); erro != nil{
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORTA"))
	if erro != nil{
		Porta = 4000
	}

	StringConexaoBanco = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}

// var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
//     "password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)