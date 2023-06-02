package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/VikeLabs/uvic-api-go/banner"
)

func main() {
	term := "202305"
	cx, err := banner.New(term)
	if err != nil {
		log.Fatal(err)
	}

	offset := 1
	var buf []banner.Datum

	for {
		log.Println("offset:", offset)
		response, err := cx.GetData(offset, banner.PageMaxSize)
		if err != nil {
			if errors.Is(err, banner.ErrEmptyOffset) {
				log.Println("done")
				break
			}
			log.Fatal(err)
		}

		buf = append(buf, response.Data...)
		if len(buf) >= 1000 {
			break
		}

		offset++
	}

	jsonData, _ := json.MarshalIndent(buf, "", "  ")
	err = os.WriteFile("uvic-202305-example.json", jsonData, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
