package data

import "testing"
import "os"

func getTestVar(varname string, t *testing.T) string {
	val, isSet := os.LookupEnv(varname)

	if !isSet {
		t.Error("Specify TEST_COLLECTION_PATH env var")
		t.Fail()
	}

	return val
}

func TestLoadPapers(t *testing.T) {

	testPath := getTestVar("TEST_COLLECTION_PATH", t)

	testIndex := getTestVar("TEST_COLLECTION_INDEX", t)

	var p = PaperCollection{"Test", testPath, testIndex, nil}

	// load the papers
	err := p.LoadPapers()

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if len(p.Articles) < 1 {
		t.Log("Did not correctly read index file")
		t.Fail()
	}

}

func TestFindPapers(t *testing.T) {

	testPath := getTestVar("TEST_COLLECTION_PATH", t)

	testIndex := getTestVar("TEST_COLLECTION_INDEX", t)

	var p = PaperCollection{"Test", testPath, testIndex, nil}

	// load the papers
	err := p.LoadPapers()

	if err != nil {
		t.Error(err)
	}

	articles := p.AtLeastOnePaperExists()

	/*for i, article := range articles {
		println(i, article.Title, len(p.CheckPapersExist(&article)))

		if i > 4 {
			println("...")
			break
		}
	}*/

	if len(articles) < 1 {
		t.Logf("Did not find any papers that exist :( ")
		t.Fail()
	}

}

func TestPaperCollectionManager(t *testing.T) {

	testPath := getTestVar("TEST_COLLECTION_PATH", t)

	testIndex := getTestVar("TEST_COLLECTION_INDEX", t)

	pm := PaperCollectionManager{Collections: nil}
	pm.AddCollection("Test", testIndex, testPath)

	if pm.Collections[0].Name != "Test" {
		t.Errorf("%s != %s", pm.Collections[0].Name, "Test")
		t.Fail()
	}

	if pm.Collections[0].IndexFile != testIndex {
		t.Errorf("%s != %s", pm.Collections[0].IndexFile, testIndex)
		t.Fail()
	}
}

func TestAggregateArticles(t *testing.T) {

	testPath := getTestVar("TEST_COLLECTION_PATH", t)

	testIndex := getTestVar("TEST_COLLECTION_INDEX", t)

	pm := PaperCollectionManager{Collections: nil}
	pm.AddCollection("Test", testIndex, testPath)
	//pm.AddCollection("Test", testIndex, testPath)

	articles := pm.AggregatePaperArticles()
	println(len(articles))

}
