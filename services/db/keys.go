package db

import (
	"encoding/json"
	"log"
	"primitivofr/kaepora/kvdb"
	"primitivofr/kaepora/models"
)

var keysDbName = "kaepora-keys-db"
var keysBucketName = "keys"

func NewKeysDb() (*KvDb, error) {
	return NewKvDb(keysDbName, keysBucketName)
}

func ReadAllKeys() ([]models.UserKeys, error) {

	db, err := NewKeysDb()

	if err != nil {
		return nil, err
	}

	db.Connect()
	defer db.Close()

	data, err := kvdb.ReadAll(db.boltDriver, db.BucketName)
	log.Println(data)
	if err != nil {
		return nil, err
	}

	var userKeys []models.UserKeys

	for _, item := range data {
		var userKeysItem models.UserKeys
		json.Unmarshal([]byte(item), &userKeysItem)
		userKeys = append(userKeys, userKeysItem)
	}

	return userKeys, nil
}
