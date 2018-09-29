package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

/*PaperCollection represents a collection of newspaper articles linked to scientific papers*/
type PaperCollection struct {
	Name       string
	PathPrefix string
	IndexFile  string
	Articles   []PaperArticle
}

/*PaperArticle is a data structure containing a news article and links to one or more scientific papers*/
type PaperArticle struct {
	Title   string            `json:"webTitle"`
	URL     string            `json:"webUrl"`
	Authors []string          `json:"authors"`
	Fields  map[string]string `json:"fields"`
	Papers  []string          `json:"papers"`
}

/*LoadPapers reads the json index file and then loads newspaper article information*/
func (p *PaperCollection) LoadPapers() error {

	// parse index file
	var f, err = os.Open(p.IndexFile)

	if err != nil {
		return err
	}

	fileData, err := ioutil.ReadAll(f)

	if err != nil {
		return err
	}

	err = json.Unmarshal(fileData, &p.Articles)

	if err != nil {
		return err
	}

	defer f.Close()

	return nil
}

/*AtLeastOnePaperExists returns articles with one or more linked scientific paper  that exists on filesystem*/
func (p *PaperCollection) AtLeastOnePaperExists() []PaperArticle {

	var result []PaperArticle

	for _, article := range p.Articles {

		if len(p.CheckPapersExist(&article)) > 0 {
			result = append(result, article)
		}

	}

	return result
}

/*CheckPapersExist returns a list of paper files known to exist for a given article
*
 */
func (p *PaperCollection) CheckPapersExist(article *PaperArticle) []string {

	var fullPaths []string

	for _, paper := range article.Papers {

		if strings.HasSuffix(paper, "pdf") {
			paper = paper[0:len(paper)-3] + "pdfx.xml"

		}

		//prepend prefix
		paper = p.PathPrefix + "/" + paper

		inList := false
		//check to see if path already in fullPaths
		for _, path := range fullPaths {
			if path == paper {
				inList = true
				break
			}
		}

		if _, err := os.Stat(paper); err == nil && !inList {
			fullPaths = append(fullPaths, paper)
		}

	}

	return fullPaths
}
