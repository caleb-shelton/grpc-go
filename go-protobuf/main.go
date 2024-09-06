package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	elliot := &Person{
		Name: "Elliot",
		Age: 24,
		SocialFollowers: &SocialFollowers{
			Youtube: 2500,
			X: 1400,
		},
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}

	// printing out our raw protobuf object
	fmt.Println(data)

	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// print out `newElliot` object
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot.SocialFollowers.GetX())
	fmt.Println(newElliot.SocialFollowers.GetYoutube())
}