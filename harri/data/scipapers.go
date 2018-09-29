package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type PaperCollection struct {
	Name       string
	PathPrefix string
	IndexFile  string
	articles   *[]PaperArticle
}

/* PaperArticle is a data structure containing a news article and links to one or more scientific papers*/
type PaperArticle struct {
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}

/* Load papers from index */
func (p PaperCollection) LoadPapers() {

	// parse index file
	var f, err = os.Open(p.IndexFile)

	if err != nil {
		panic(err)
	}

	var fileData, e = ioutil.ReadAll(f)

	if e != nil {
		panic(err)
	}

	var articles []PaperArticle

	json.Unmarshal(fileData, &articles)

	p.articles = &articles

	defer f.Close()
}
