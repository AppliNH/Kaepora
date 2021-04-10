package user

import (
	"log"
	"primitivofr/kaepora/kvdb"
)

func ReadAllFromDB() {
	data1, _ := kvdb.ReadAll(dbUsers, "users")
	data2, _ := kvdb.ReadAll(dbKeys, "keys")

	log.Println(data1)
	log.Println(data2)
}
