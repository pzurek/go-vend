package main

import (
	"encoding/json"
	"fmt"
	"github.com/pzurek/vendengo/vend"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// to run:
// go run vend.go <store name> <username> <password>
func main() {
	args := os.Args
	storeName := args[1]
	username := args[2]
	password := args[3]

	baseUrl := fmt.Sprintf("https://%s.vendhq.com/api/", storeName)

	client := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl+"config", nil)
	req.SetBasicAuth(username, password)
	res, err := client.Do(req)
	defer res.Body.Close()

	if err != nil {
		log.Fatal("Error: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatal("Response status: ", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal("Error: %s", err)
	}

	cr := new(vend.ConfigResponse)
	c := new(vend.Config)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &cr)

	if err != nil {
		log.Fatal(err)
	}

	c = cr.Config

	fmt.Printf("Retailer name: %+v\n", *c.RetailerName)
	fmt.Printf("Retailer id: %+v\n", *c.RetailerId)
	fmt.Printf("Account state: %+v\n", *c.AccountState)
	fmt.Printf("User name: %+v\n", *c.UserName)
	fmt.Printf("Currency name:  %+v\n", *c.Currency)
}
