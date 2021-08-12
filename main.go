package main

import (
	"context"
	"encoding/json"
	"fmt"

	"gopkg.in/olivere/elastic.v5"
)

func main() {
	index := "delish-recipes"
	analyzer := "ja_synonym_search_analyzer"
	term := "豚スープ"
	client, _ := elastic.NewSimpleClient(
		elastic.SetURL("http://es:9200"),
	)

	ctx := context.Background()
	resp, _ := client.IndexAnalyze().Analyzer(analyzer).Index(index).Text(term).Do(ctx)
	result, _ := json.Marshal(resp)
	fmt.Println(string(result))
}
