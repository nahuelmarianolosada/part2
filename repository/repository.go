package repository

import (
	"fmt"
	"strconv"
)

type Repository interface {
	Insert(value int) error
	GetAll() ([]string, error)
	GetByID(id int) (string, error)
}

type NumberCollection struct {
	db []string
}

func (nm *NumberCollection) Insert(value int) error {
	if nm.db == nil {
		nm.db = []string{strconv.Itoa(value)}
		return nil
	}
	nm.db = append(nm.db, strconv.Itoa(value))
	fmt.Println(fmt.Sprintf("=============== db %+v", nm.db))
	return nil
}

func (nm *NumberCollection) GetAll() ([]string, error) {
	fmt.Println(fmt.Sprintf("=============== db in getAll() %+v", nm.db))
	return nm.db, nil
}

func (nm *NumberCollection) GetByID(id int) (string, error) {
	return nm.db[id], nil
}
