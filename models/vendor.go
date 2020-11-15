package models

import (
	"math/rand"
	"time"
)

type Vendor struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Enabled bool   `json:"enabled"`
}

type Vendors struct {
	Vendor []Vendor
}

func GetRandom(v []Vendor) *Vendor {
	rand.Seed(time.Now().Unix())
	return &v[rand.Intn(len(v))]
}
