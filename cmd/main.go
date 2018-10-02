package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/ravenscroftj/goharri/harri/data"
)

type collection struct {
	Name       string
	IndexFile  string
	PrefixPath string
}

type PaperCollections struct {
	Collections []collection
}

func main() {
	println("Hello!")

	var collections PaperCollections
	//var config configStruct

	if _, err := toml.DecodeFile("config.toml", &collections); err != nil {
		println(err.Error())
		return
	}

	pm := data.PaperCollectionManager{Collections: nil}

	//print(config.Title)
	for _, col := range collections.Collections {
		pm.AddCollection(col.Name, col.IndexFile, col.PrefixPath)
	}

	fmt.Printf("Found %d collections and %d scientific papers", len(collections.Collections), len(pm.AggregatePaperArticles()))

}
