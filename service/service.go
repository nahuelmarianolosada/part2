package service

import (
	"fmt"
	"strconv"
	"log"
	"part2/repository"
)

type NumberService interface {
	GetAll() []string
	GetByID(id int) *string
	Insert(id int) error
}

type NumberServiceImpl struct {
	NumberCollection repository.Repository
}

func (ns *NumberServiceImpl) GetAll() []string {
	values, err := ns.NumberCollection.GetAll()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	for idx, value := range values {
		intValue, _ := strconv.Atoi(value)
		values[idx] = CheckValue(intValue)
	}
	return values
}

func (ns *NumberServiceImpl) GetByID(id int) *string {
	valueType := CheckValue(id)
	return &valueType
}

func (ns *NumberServiceImpl) Insert(id int) error {
	return ns.NumberCollection.Insert(id)
}

var maps = map[string]func(i int) bool{
	"Type 3": is3And5Multiple,
	"Type 1": is3Multiple,
	"Type 2": is5Multiple,
}

func CheckValue(i int) string {
	for typeFound, fun := range maps {
		if fun(i) {
			return typeFound
		}
	}
	return fmt.Sprintf("%d", i)
}

func is3Multiple(value int) bool {
	return value%3 == 0
}

func is5Multiple(value int) bool {
	return value%5 == 0
}

func is3And5Multiple(value int) bool {
	return is3Multiple(value) && is5Multiple(value)
}
