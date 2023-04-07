package handlers

import (
	"fmt"

	"github.com/gevs90/crypto-api/api/models"
	"github.com/gevs90/crypto-api/api/utils"
	dbmodel "github.com/gevs90/crypto-api/db/models"
	"github.com/gevs90/crypto-api/db/repository"
)

type cryptoFunctions struct {
	lr *repository.LogRepository
	tr *repository.TextRepository
}

func NewCryptoFunctions(lr *repository.LogRepository, tr *repository.TextRepository) *cryptoFunctions {
	return &cryptoFunctions{
		lr: lr,
		tr: tr,
	}
}

var kc = utils.NewKeyCache()

func (nc *cryptoFunctions) EncryptData(text string) (models.ResponseEncryptedText, error) {

	nc.lr.Create(&dbmodel.Log{
		Url:    "/v1/encrypt",
		Method: "POST",
	})

	ciphertext, key, err := kc.EncryptString(text)
	if err != nil {
		fmt.Println("encryption error:", err)
		return models.ResponseEncryptedText{}, err
	}
	var textModel = &dbmodel.Text{
		EncryptedText: ciphertext,
		EncryptionKey: key,
	}
	result := nc.tr.Create(textModel)
	if result != nil {
		fmt.Println("error creating record [Text]:", err)
		return models.ResponseEncryptedText{}, err
	}

	return models.ResponseEncryptedText{
		ID:            textModel.Model.ID,
		EncryptedText: ciphertext,
		EncryptionKey: key,
		Text:          text,
	}, nil
}

func (nc *cryptoFunctions) DecryptData(record_id int) (models.ResponseEncryptedText, error) {
	nc.lr.Create(&dbmodel.Log{
		Url:    "/v1/encrypt/" + fmt.Sprint(record_id),
		Method: "GET",
	})

	text_record, err := nc.tr.Find(record_id)
	if err != nil {
		fmt.Println("encryption error:", err)
		return models.ResponseEncryptedText{}, err
	}

	text, key, err := kc.DecryptString(text_record.EncryptedText, text_record.EncryptionKey)
	if err != nil {
		fmt.Println("encryption error:", err)
		return models.ResponseEncryptedText{}, err
	}

	return models.ResponseEncryptedText{
		ID:            text_record.ID,
		EncryptedText: text_record.EncryptedText,
		EncryptionKey: key,
		Text:          text,
	}, nil
}
