package main

import (
	"context"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/boom"
	_ "go.mercari.io/datastore/clouddatastore"
	"log"
)

const (
	ProjectID = "morning-tide"
)

type TestData struct {
	ID int64 `json:"id" datastore:"-" boom:"id"`
	Name string `json:"name"`
}

func main() {
	ctx := context.Background()
	opts := datastore.WithProjectID(ProjectID)
	ds, err := datastore.FromContext(ctx, opts)
	if err != nil {
		log.Fatalf("Failed to create datastore client: %v", err)
	}

	b := boom.FromClient(ctx, ds)
	data := &TestData{Name: "hello-ds-name-example"}
	key, err := b.Put(data)
	if err != nil {
		log.Fatalf("Failed to put datastore: %v", err)
	}

	log.Printf("inserted new data (key: %s)", key)
}