package main

import (
	"io/ioutil"
	"log"

	pb "github.com/denis-zakharov/goprotobuf/tutorialpb"
	"google.golang.org/protobuf/proto"
)

func main() {
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	book := &pb.AddressBook{}
	book.People = append(book.People, &p)

	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode an address book:", err)
	}
	if err := ioutil.WriteFile("book.pb", out, 0644); err != nil {
		log.Fatalln("Failed to write an address book:", err)
	}

	rbook := &pb.AddressBook{}
	protoBytes, err := ioutil.ReadFile("book.pb")
	if err != nil {
		log.Fatalln("Failed to read an address book:", err)
	}
	if err := proto.Unmarshal(protoBytes, rbook); err != nil {
		log.Fatalln("Failed to decode an address book:", err)
	}

	log.Default().Printf("%+v\n", rbook)
}
