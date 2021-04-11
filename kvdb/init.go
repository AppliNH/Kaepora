package kvdb

import (
	"os"
	"path/filepath"

	"github.com/boltdb/bolt"
)

// InitDB init a boltdb object + db file + bucket if needed
// dbName is the name of the db (folder + .db file)
// bucketName is the name of the bucket that will contain data
func InitDB(dbName string, bucketName string) (*bolt.DB, error) {

	ex, _ := os.Executable()

	exPath := filepath.Dir(ex)
	//fmt.Println(exPath)

	if _, err := os.Stat(exPath + "/kaepora-db/" + dbName); os.IsNotExist(err) {
		err := os.MkdirAll(exPath+"/kaepora-db/"+dbName, 0700)
		if err != nil {
			panic(err)
		}
	}

	db, err := bolt.Open(exPath+"/kaepora-db/"+dbName+"/"+dbName+".db", 0700, nil)
	if err != nil {
		panic(err)
	}

	if bucketName != "" {

		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
			if err != nil {
				return err
			}
			return nil
		})
	}

	return db, err

}
