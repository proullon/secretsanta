package bucket

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Email      string
	Name       string
	InLoveWith *Person
}

func LoveLove(name1 string, name2 string, people []*Person) {
	for _, one := range people {
		for _, two := range people {
			if one.Name == name1 && two.Name == name2 {
				one.InLoveWith = two
				two.InLoveWith = one
			}
		}
	}
}

func randInt(min int, max int) int {
	time.Sleep(235 * time.Microsecond)
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func goodRand(name *Person, toAssign []*Person) int {

	for {
		r := randInt(0, len(toAssign))

		if toAssign[r].InLoveWith != nil && name.Email == toAssign[r].InLoveWith.Email {
			continue
		}

		if toAssign[r].Name != name.Name {
			return r
		}
	}

}

func removeFromList(r int, toAssign []*Person) []*Person {
	var newToAssign []*Person

	for i, name := range toAssign {
		if i != r {
			newToAssign = append(newToAssign, name)
		}
	}

	return newToAssign
}

func RunChristmasBucket(people []*Person, conf EmailAccount) {

	toAssign := people

	for i := range people {
		fmt.Printf("Sending %d/%d\n", i+1, len(people))
		r := goodRand(people[i], toAssign)

		body := fmt.Sprintf("You offer a gift to %s\n%s\n", toAssign[r].Name, conf.Body)
		err := sendEmail(conf, conf.Subject, body, people[i].Email)
		if err != nil {
			fmt.Printf("Arg ! %s\n", err)
		}

		toAssign = removeFromList(r, toAssign)
	}
}
