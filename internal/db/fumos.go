package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type FumosRepo struct {
	pool      *pgxpool.Pool
	fumosJSON []FumosJSON
}

type FumosJSON struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

type FumoMediaData struct {
	FumoIDs     []int
	Description *string
	MediaURL    string
	Approved    bool
}

func newFumosRepo(p *pgxpool.Pool) *FumosRepo {
	fumos := readFumosJSON("data/fumos.json")

	return &FumosRepo{
		pool:      p,
		fumosJSON: fumos,
	}
}

func (fr *FumosRepo) uploadMedia(ctx context.Context, data FumoMediaData) error {
	_, err := fr.pool.Exec(
		ctx,
		`INSERT INTO fumo_media (fumo_ids, description, media_url, approved)
		VALUES ($1, $2, $3, $4)`,
		data.FumoIDs,
		data.Description,
		data.MediaURL,
		data.Approved,
	)
	if err != nil {
		return fmt.Errorf(
			"Failed to upload media for fumos \"%v\" (%s): %w",
			data.FumoIDs,
			data.MediaURL,
			err,
		)
	}
	return nil
}

func readFumosJSON(path string) []FumosJSON {
	fileData, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read fumos JSON file: %s\n", err)
	}

	var fumos []FumosJSON
	err = json.Unmarshal(fileData, &fumos)
	if err != nil {
		log.Fatalf("Failed to parse fumos JSON file: %s\n", err)
	}

	return fumos
}
