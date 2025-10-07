package main

import (
	"log"

	"github.com/joho/godotenv"
	"githum.com/erfanyousefi/course-project/internal/db"
	"githum.com/erfanyousefi/course-project/internal/env"
	"githum.com/erfanyousefi/course-project/internal/store"
)

func main() {
	godotenv.Load()
	dbCfg := dbConfig{
		addr:         env.GetString("DB_ADDR", ""),
		maxOpenConns: env.GetInt("DB_MAX_OPEN_CONN", 30),
		maxIdleConns: env.GetInt("DB_MAX_IDLE_CONN", 30),
		maxIdleTime:  env.GetString("DB_MAX_IDLE_Time", "15min "),
	}
	cfg := config{
		addr: env.GetString("ADDR", ":4000"),
		db:   dbCfg,
	}
	db, err := db.New(dbCfg.addr, dbCfg.maxOpenConns, dbCfg.maxIdleConns, dbCfg.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	store := store.NewStorage(db)
	app := &application{
		config: cfg,
		store:  store,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
