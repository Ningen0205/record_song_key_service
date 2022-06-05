package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"record_song_key_service/src/driver"
	"strings"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Println(err)
	}
}

func main() {
	f := flag.String("mode", "Up", "migration_mode")
	flag.Parse()
	mode := strings.ToLower(*f)

	if mode != "up" && mode != "down" && mode != "show" {
		log.Panicln("mode: up or down only")
	}

	dir := fmt.Sprintf("migration/%s/", mode)

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Panicln(err)
	}

	loadEnv()

	dbDriver := driver.GetInstance()

	db := dbDriver.GetConnection(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_OUT_PORT"),
	)

	for _, file := range files {
		if !file.IsDir() && path.Ext(file.Name()) == ".sql" {
			b, err := ioutil.ReadFile(dir + file.Name())

			if err != nil {
				log.Panicln(err)
			}

			sql := string(b)

			db.Exec(sql)
		}
	}
}
