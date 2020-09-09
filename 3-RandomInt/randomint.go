package main

import(
	"fmt"
	"math/rand"
	"time"
)

func Seed(min, max int ) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func main(){
	randint := Seed(1,100)
	fmt.Println("Your random number is:", randint)
}


