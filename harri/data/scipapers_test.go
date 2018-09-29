package data

import "testing"
import "os"

func TestLoadPapers(t *testing.T) {

	var isSet = false
	var testPath, testIndex string

	testPath, isSet = os.LookupEnv("TEST_COLLECTION_PATH")

	if !isSet {
		panic("Specify TEST_COLLECTION_PATH env var")
	}

	testIndex, isSet = os.LookupEnv("TEST_COLLECTION_INDEX")

	if !isSet {
		panic("Specify TEST_COLLECTION_INDEX env var")
	}

	var p = PaperCollection{"Test", testPath, testIndex}

	// load the papers
	p.LoadPapers()
}
