package nosql

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type DefaultResult struct {
	UID       string
	IsDefault bool
}

var (
	someQueryString string = "some query string"
	dg              *DormDB
)

func init() {
	dg = NewDGraphClient(&DgraphCfg{Url: "127.0.0.1:8090"})
}

func TestGraphQuery(t *testing.T) {

	resp, err := dg.Query(someQueryString)
	if err != nil {
		log.Fatal(err)
	}
	var result DefaultResult
	_ = json.Unmarshal(resp.GetJson(), &result)
	fmt.Printf("%s", "hello")
	fmt.Printf("%+v", result)
}

func TestDGraphReadonlyQuery(t *testing.T) {

	resp, err := dg.QueryReadOnly(someQueryString)
	if err != nil {
		log.Fatal(err)
	}

	var result DefaultResult
	_ = json.Unmarshal(resp.GetJson(), &result)
	fmt.Printf("%s", "hello")
	fmt.Printf("%+v", result)
}

func TestColumnExistQueryString(t *testing.T) {

	resp, err := dg.QueryReadOnly(someQueryString)
	if err != nil {
		log.Fatal(err)
	}
	var result DefaultResult
	_ = json.Unmarshal(resp.GetJson(), &result)

	fmt.Printf("%s\n", "hello")
	fmt.Printf("%+v", result)
}
