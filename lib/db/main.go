package db

import (
	"log"

	"github.com/applinh/kaepora/kvdb"

	"github.com/boltdb/bolt"
)

type KvDb struct {
	DbName     string
	BucketName string
	boltDriver *bolt.DB
}

type KvDbInteract interface {
	Connect() error
	Read(key string) (string, error)
	Write(key string, value string) error
	Close() error
}

func (k *KvDb) Connect() error {
	var err error
	k.boltDriver, err = kvdb.InitDB(k.DbName, k.BucketName)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
func (k *KvDb) Close() error {
	return k.boltDriver.Close()
}
func (k *KvDb) Read(key string) (string, error) {
	k.Connect()
	defer k.Close()

	return kvdb.ReadData(k.boltDriver, k.BucketName, key)
}

func (k *KvDb) Write(key string, value string) error {
	k.Connect()
	defer k.Close()

	return kvdb.WriteData(k.boltDriver, k.BucketName, key, value)
}

func NewKvDb(dbName string, bucketName string) (*KvDb, error) {
	return &KvDb{dbName, bucketName, nil}, nil
}
