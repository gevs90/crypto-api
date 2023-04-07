package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gevs90/crypto-api/api/handlers"
	"github.com/gevs90/crypto-api/db"
	"github.com/gevs90/crypto-api/db/migrations"
	"github.com/gevs90/crypto-api/db/repository"
	"github.com/gorilla/mux"
)

func Decrypt(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	record_id, _ := strconv.Atoi(params["record_id"])

	db, err := db.DB()
	if err != nil {
		panic(err)
	}

	migrations.MigrateLogTable(db)
	migrations.MigrateTextTable(db)

	var crypto = handlers.NewCryptoFunctions(repository.NewLogRepository(db), repository.NewTextRepository(db))
	response, err := crypto.DecryptData(record_id)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}
