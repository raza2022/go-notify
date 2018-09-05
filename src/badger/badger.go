package main

import (
	"log"

	"encoding/json"
	"fmt"
	"github.com/dgraph-io/badger"
	"github.com/shirou/gopsutil/host"
)

func main() {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	opts := badger.DefaultOptions
	opts.Dir = "./db"
	opts.ValueDir = "./db"
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	//errr := db.View(func(txn *badger.Txn) error {
	//	fmt.Println("db", txn)
	//	// Your code here…
	//	return nil
	//})
	test1 := db.Update(func(txn *badger.Txn) error {
		host, err := host.Info()
		//message, err := json.Marshal(map[string]string{"foo": "bar"})
		message, err := json.Marshal(host)
		errr := txn.Set([]byte("host"), message)
		fmt.Println("why I have to print this shit", errr)
		//fmt.Println("I got this host", host, message)
		return err
	})
	fmt.Println("I got this shit", test1)

	//query with single names
	test2 := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("host"))
		if err != nil {
			return err
		}
		val, err := item.Value()
		if err != nil {
			return err
		}
		fmt.Printf("The host values is: %s\n", val)
		return nil
	})
	fmt.Printf("just chill it's no effect", test2)
	defer db.Close()
}
