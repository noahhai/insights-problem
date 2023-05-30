package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	body := []byte(`{
		"title": "testing",
		"body": "data"
	}`)

   resp, err := http.Post(os.Getenv("POST_ENDPOINT"), "application/text", bytes.NewBuffer(body))
   if err != nil {
      log.Fatalln(err)
   }
   
   body, err = ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }

   sb := string(body)
   log.Printf(sb)
}