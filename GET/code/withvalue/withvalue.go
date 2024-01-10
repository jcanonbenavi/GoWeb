package main

import (
	"context"
	"fmt"
)

type User struct {
	Name  string
	Other string
}

func main() {

	ctx := context.Background() //father context
	user := User{
		Name:  "John Doe",
		Other: "Other data",
	}
	ctx = context.WithValue(ctx, "user", user)
	//ctx is the father context from which we can create a child context, and key value
	//value can be any type (string, struct...)
	funcionOne(ctx)

}

func funcionOne(ctx context.Context) {
	value := ctx.Value("user")
	fmt.Println(value)
}
