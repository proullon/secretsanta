package bucket

import (
	"crypto/rand"
	"fmt"
	"math/big"
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

func randInt(min int, max int) (int, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return 0, err
	}

	return int(r.Int64()), nil
}

func goodRand(name *Person, toAssign []*Person) (int, error) {

	for out := 0; out < 100; out++ {
		r, err := randInt(0, len(toAssign))
		if err != nil {
			return 0, err
		}

		if toAssign[r].InLoveWith != nil && name.Email == toAssign[r].InLoveWith.Email {
			continue
		}

		if toAssign[r].Name != name.Name {
			return r, nil
		}
	}

	return 0, fmt.Errorf("No good match left!")
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

type Gift struct {
	Gifter   *Person
	Receiver *Person
}

func RunChristmasBucket(people []*Person, conf EmailAccount) {

	toAssign := people
	var gifts []Gift

	for i := range people {
		r, err := goodRand(people[i], toAssign)
		if err != nil {
			RunChristmasBucket(people, conf)
			return
		}

		g := Gift{
			Gifter:   people[i],
			Receiver: toAssign[r],
		}
		gifts = append(gifts, g)
		toAssign = removeFromList(r, toAssign)
	}

	for _, g := range gifts {
		fmt.Printf("%s -> %s\n", g.Gifter.Name, g.Receiver.Name)
		body := fmt.Sprintf("<p>Tu offres un cadeau à <b>%s!</b></p>%s", g.Receiver.Name, conf.Body)

		err := sendEmail(conf, conf.Subject, body, g.Gifter.Email)
		if err != nil {
			fmt.Printf("Arg ! %s\n", err)
		}
	}

}
