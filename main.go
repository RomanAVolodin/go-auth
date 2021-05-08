package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string `json:"first_name"`
}

func main()  {
	p1 := person{
		First: "Roman",
	}

	p2 := person{
		First: "Julia",
	}

	xp := []person{p1, p2}

	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Print JSON", string(bs))

	xp2 := []person{}

	err = json.Unmarshal(bs, &xp2)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("back to go datastructure", xp2)
}