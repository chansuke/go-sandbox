package main

import "fmt"

// Duck walks like duck
type Duck interface {
	Duckwalk() string
}

// Dog barks like dog
type Dog interface {
	Bark() string
}

type Godzilla struct {
	soundOfWalking string
	soundOfBarking string
}

func (g Godzilla) Duckwalk() string {
	return g.soundOfWalking
}

func (g Godzilla) Bark() string {
	return g.soundOfBarking
}

func LetDuckWalk(duck Duck) string {
	return duck.Duckwalk()
}

func LetDogBark(dog Dog) string {
	return dog.Bark()
}

func main() {
	godzilla := &Godzilla{
		"ぴょん",
		"わん",
	}
	fmt.Printf("ゴジラが来た => %+v\n", godzilla)
	fmt.Println(
		"アヒルなら歩いてみろよ",
		LetDuckWalk(godzilla),
		"犬なら吠えてみろよ",
		LetDogBark(godzilla),
	)
}
