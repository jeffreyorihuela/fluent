package db

import (
	"context"
	"fluent/ent"
	"log"

	_ "github.com/lib/pq"
)

func GetClient() (*ent.Client, error) {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5434 user=postgres dbname=postgres password=123456 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}
