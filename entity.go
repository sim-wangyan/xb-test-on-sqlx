package main

import "time"

type Pet struct {
	Id   uint64 `db:"id"`
	Type string `db:"type"`
}

func (*Pet) TableName() string {
	return "t_pet"
}

type Cat struct {
	Id       uint64    `db:"id"`
	Name     string    `db:"name"`
	PetId    *uint64   `db:"pet_id"`
	Age      *uint     `db:"age"`
	Color    string    `db:"color"`
	Weight   *float64  `db:"weight"`
	IsSold   *bool     `db:"is_sold"`
	Price    *float64  `db:"price"`
	CreateAt time.Time `db:"create_at"`
}

type CatRo struct {
	Name   string   `json:"name, string"`
	IsSold *bool    `json:"isSold, *bool"`
	Price  *float64 `json:"price, *float64"`
	Age    *uint    `json:"age", unit`
}

func (*Cat) TableName() string {
	return "t_cat_go"
}

type Dog struct {
	Id     uint64   `db:"id"`
	Age    *int     `db:"age"`
	PetId  uint64   `db:"pet_id"`
	Weight *float64 `db:"weight"`
	Height *int32   `db:"height"`
}

func (*Dog) TableName() string {
	return "t_dog"
}
