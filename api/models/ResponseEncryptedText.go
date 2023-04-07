package models

type ResponseEncryptedText struct {
	ID            uint
	Text          string
	EncryptedText string
	EncryptionKey string
}
