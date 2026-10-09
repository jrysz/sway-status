package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	FullText string `json:"full_text"`
	Markup   string `json:"markup,omitempty"`
}

func print() Block {
	return Block{
		FullText: time.Now().Format("15:04") + " ",
	}
}

func main() {
	fmt.Println(`{"version":1}`)
	fmt.Println("[")

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	first := true

	for range ticker.C {
		blocks := []Block{
			print(),
		}

		data, err := json.Marshal(blocks)
		if err != nil {
			continue
		}

		if !first {
			fmt.Println(",")
		}
		first = false

		fmt.Println(string(data))
	}
}
