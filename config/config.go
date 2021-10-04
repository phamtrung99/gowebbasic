package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type MysqlEnvConfig struct {
	Host     string
	DB       string
	User     string
	Port     string
	Password string
	URI      string
}

type ServerEnvConfig struct {
	Port string
}

type Pagination struct {
	MovieRowPerPage int
}

func LoadFileEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}

func GetSecretKey() string {
	LoadFileEnv()
	return os.Getenv("SECRET_KEY")
}

func GetMysqlEnvConfig() *MysqlEnvConfig {
	LoadFileEnv()
	var mysql MysqlEnvConfig
	mysql.Host = os.Getenv("MYSQL_HOST")
	mysql.User = os.Getenv("MYSQL_USER")
	mysql.Password = os.Getenv("MYSQL_PASSWORD")
	mysql.DB = os.Getenv("MYSQL_DB")
	mysql.Port = os.Getenv("MYSQL_PORT")

	return &mysql
}

func GetServerConfig() *ServerEnvConfig {
	LoadFileEnv()
	var server ServerEnvConfig
	server.Port = os.Getenv("SERVER_PORT")

	return &server
}

func GetPagination() *Pagination {
	LoadFileEnv()
	var pagi Pagination
	rowPerPage, err := strconv.Atoi(os.Getenv("MOVIE_ROW_PER_PAGE"))
	pagi.MovieRowPerPage = rowPerPage

	if err != nil {
		log.Fatal(err)
	}

	return &pagi
}
