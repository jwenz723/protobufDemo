package main

import (
	"github.com/jwenz723/protobufDemo/protogen"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	//createEncodedFile("addressbook4.dat")

	files, err := filepath.Glob("./*.dat")
	if err != nil {
		log.Fatalln("Failed to get files: ", err)
	}

	for _, f := range files {
		// Read the existing address book.
		in, err := ioutil.ReadFile(f)
		if err != nil {
			log.Fatalln("Error reading file:", err)
		}
		inBook := &contacts.AddressBook{}
		if err := proto.Unmarshal(in, inBook); err != nil {
			log.Fatalln("Failed to parse address book:", err)
		}

		log.Printf("\n%s: %v\n", f, inBook)

		for _, p := range inBook.People {
			log.Printf("%s - %s, %s, %v\n", f, p.Name, p.Email, p.Address)
		}
	}
}

// Creates a protobuf encoded file
func createEncodedFile(fileName string) {
	p := contacts.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*contacts.Person_PhoneNumber{
			{Number: "555-4321", Type: contacts.Person_HOME},
		},
		Address: &contacts.Person_Address{
			Address: "123 blueberry",
			City: "Salt Lake",
			State: "Utah",
		},
	}
	p2 := contacts.Person{
		Id:    1234,
		Name:  "Jane Doe",
		Email: "jadoe@example.com",
		Phones: []*contacts.Person_PhoneNumber{
			{Number: "555-2222", Type: contacts.Person_MOBILE},
		},
		Address: &contacts.Person_Address{
			Address: "123 strawberry",
			City: "Denver",
			State: "Colorado",
		},
	}

	book := &contacts.AddressBook{
		People: []*contacts.Person{&p, &p2},
	}

	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fileName, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	log.Println("Created file: ", fileName)
}
