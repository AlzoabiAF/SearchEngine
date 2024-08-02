package search

import (
	"context"
	"encoding/json"
	"os"

	"github.com/olivere/elastic/v7"
)

type Product struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Category     string  `json:"category"`
	Availability bool    `json:"availability"`
}

func SearchPruducts(query string) ([]Product, error) {

	client, err := NewClient()
	if err != nil {
		return nil, elastic.ErrNoClient
	}


	searchResults, err := client.Search().
		Index("products").
		Query(elastic.NewMultiMatchQuery(query, "name", "description")).
		Do(context.Background())

	if err != nil {
		return nil, err
	}

	var products []Product

	for _, hit := range searchResults.Hits.Hits {
		var product Product
		err := json.Unmarshal(hit.Source, &product)
		if err != nil{
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func NewClient() (*elastic.Client, error) {
	return elastic.NewClient(
		elastic.SetURL(os.Getenv("ELASTIC_URL")), 
		elastic.SetBasicAuth(os.Getenv("ELASTIC_USER"), os.Getenv("ELASTIC_PASSWORD")),
	)
}
