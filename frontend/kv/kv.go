package kv

import (
  "go.etcd.io/bbolt"
)

func GetDB() (*bbolt.DB, error) {
  db, err := bbolt.Open("bolt.db", 0600, nil)
	if err != nil {
	  return nil, err
	}
	defer db.Close()
	return db, nil 
}

func StoreJWT(db *bbolt.DB, token[] byte) error {
	return nil // TODO: 
	// bucket, err := db.CreateBucketIfNotExists([]byte("JWTBucket"))
	// 	if err != nil {
	// 		return err
	// 	}
	//   _, err := bucket.Put([]byte("jwt"))
	// return err
}
