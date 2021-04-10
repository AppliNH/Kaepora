package db

import (
	"github.com/boltdb/bolt"
)

var dbKeys *bolt.DB
var dbUsers *bolt.DB

var usersDbName = "kaepora-users-db"
var usersBucketName = "users"

func NewUsersDb() (*KvDb, error) {
	return NewKvDb(usersDbName, usersBucketName)
}
