package models

var dog = "I'm Dog"

func GetDog() string {
	return dog
}

func GetDogAndCat() string {
	return dog + " love cat"
}
