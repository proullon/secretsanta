package main

import (
	"github.com/proullon/christmas/bucket"
)

func main() {
	people := []*bucket.Person{
		&bucket.Person{
			Email: "foo@bar.com",
			Name:  "Foo",
		},
		&bucket.Person{
			Email: "bar@bar.com",
			Name:  "Bar",
		},
		&bucket.Person{
			Email: "alice@example.com",
			Name:  "Alice",
		},
		&bucket.Person{
			Email: "bob@example.com",
			Name:  "Bob",
		},
	}

	conf := bucket.EmailAccount{
		Port:     bucket.GmailPort,
		Server:   bucket.GmailServer,
		Email:    "me@gmail.com",
		Password: "p4$$w0rd",
		Subject:  "The christmas gift repartition !",
		Body:     "Optional body to say hello all, see you soon !",
	}

	// Disable possibility for Alice and Bob
	// to offer a gift to each other
	bucket.LoveLove("Alice", "Bob", people)

	// Go !
	bucket.RunChristmasBucket(people, conf)
}
