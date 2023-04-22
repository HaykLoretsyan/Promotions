package main

type Model struct {
	Id             string  `json:"id" bson:"id"`
	Price          float32 `json:"price" bson:"price"`
	ExpirationDate string  `json:"expiration_date" bson:"expiration_date"`
}
