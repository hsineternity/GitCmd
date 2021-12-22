package main

import (
	"fmt"

	"github.com/hsineternity/GitCmd/models"
)

func main() {
	fmt.Println(models.GetAuthor())
	fmt.Println(models.GetDog())
}
