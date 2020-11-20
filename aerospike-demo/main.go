package main

import (
	"fmt"

	as "github.com/aerospike/aerospike-client-go"
)

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	client, err := as.NewClient("localhost", 3000)
	if err != nil {
		panic(err)
	}

	// ns.set
	key, _ := as.NewKey("test", "foo",
		"key value goes here and can be any supported primitive")

	bin1 := as.NewBin("bin1", "value1")
	bin2 := as.NewBin("bin2", "value2")

	// Write a record
	err = client.PutBins(nil, key, bin1, bin2)
	if err != nil {
		panic(err)
	}

	// Read a record
	record, err := client.Get(nil, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(record)

	client.Close()
}
