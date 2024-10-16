package kv

import (
	"errors"
	"go.etcd.io/bbolt"
)

type Key int

const (
	JWT Key = iota
)

func (k Key) String() string {
	return [...]string{"jwt"}[k]
}
func (k Key) Bytes() []byte {
	return []byte(k.String())
}

func GetDB() (*bbolt.DB, error) {
  db, err := bbolt.Open("bbolt.db", 0600, nil)
	if err != nil {
	  return nil, err
	}
	return db, nil 
}

func GetBucket(tx *bbolt.Tx) *bbolt.Bucket {
	return tx.Bucket([]byte("client_bucket"))
}
func GetBucketUnsafe(tx *bbolt.Tx) (*bbolt.Bucket, error) {
	bucket := tx.Bucket([]byte("client_bucket"))
	if bucket == nil {
		_, err := tx.CreateBucket([]byte("client_bucket"))
		if err != nil {
			return nil, err
		}
		return tx.Bucket([]byte("client_bucket")), nil
	}
	return bucket, nil
}

func Put(db *bbolt.DB, key Key, value[] byte) error {
	err := db.Update(func(tx *bbolt.Tx) error {
		bucket, err := GetBucketUnsafe(tx)
		if err != nil {
			return err
		}

    return bucket.Put(key.Bytes(), value)
	})
	return err
}
func Get(db *bbolt.DB, key Key) ([]byte, error) {
	var value []byte
	err := db.View(func(tx *bbolt.Tx) error {
		bucket := GetBucket(tx)
		if bucket == nil {
			return errors.New("bucket not found")
		}
		value = bucket.Get(key.Bytes())
		return nil
	})
	if err != nil {
		return nil, err
	}
	return value, nil
}
