package main

import (
	"bytes"
	"goat/pkg/ai"
	goat "goat/pkg/http"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {
	buf, err := ioutil.ReadFile("goat-dump-v1")
	if err != nil {
		panic(err)
	}

	generator := ai.NewGenerator(bytes.NewReader(buf))
	app := goat.NewApp(generator)
	log.Fatal(http.ListenAndServe(":8080", app))
}


