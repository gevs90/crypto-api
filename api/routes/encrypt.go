package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gevs90/crypto-api/api/handlers"
	"github.com/gevs90/crypto-api/api/models"
	"github.com/gevs90/crypto-api/db"
	"github.com/gevs90/crypto-api/db/migrations"
	"github.com/gevs90/crypto-api/db/repository"
)

func Encrypt(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestData models.RequestEncrypt

	err := decoder.Decode(&requestData)
	if err != nil {
		panic(err)
	}

	db, err := db.DB()
	if err != nil {
		panic(err)
	}

	migrations.MigrateLogTable(db)
	migrations.MigrateTextTable(db)

	var crypto = handlers.NewCryptoFunctions(repository.NewLogRepository(db), repository.NewTextRepository(db))

	dataEncrypted, err := crypto.EncryptData(requestData.Text)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(dataEncrypted)
}
