package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func load(filename string, target interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, target)
}

func TestJSONStruct(t *testing.T) {
	invoices := make([]*Invoice, 1)
	plays := make(map[string]*Play)

	testData := []struct {
		filename string
		model    interface{}
	}{
		{"testdata/invoices.json", &invoices},
		{"testdata/plays.json", &plays},
	}

	for _, td := range testData {
		err := load(td.filename, td.model)
		if err != nil {
			t.Fatalf("can't load %s, %s", td.filename, err)
		}
	}

	if len(invoices) != 1 {
		t.Fatalf("invoices is empty, want 1")
	}

	t.Logf("Invoice.Customer: %s", invoices[0].Customer)
	for i, p := range invoices[0].Performances {
		t.Logf("Invoice.Performances[%d]: %+v", i, p)
	}

	if len(plays) != 3 {
		t.Fatalf("plays is invalid got: %d, want: 3", len(plays))
	}
	for k, p := range plays {
		t.Logf("Plays[%s]: %+v", k, p)
	}
}
