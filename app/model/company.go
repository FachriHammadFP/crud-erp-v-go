package model

import (
	"erp-go-v2/app/koneksi"
	"fmt"
)

type Company struct {
	Id      int    `json:"Id"`
	Name    string `json:"Name" validate:"required"`
	Address string `json:"Address" validate:"required"`
}

func (c *Company) Create() error {
	db, err := koneksi.ConnectDb()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()
	_, err = db.Exec("insert into companies (name, address) values(?,?)", c.Name, c.Address)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("company added")
	return nil
}
func (c *Company) Read() ([]Company, error) {
	db, err := koneksi.ConnectDb()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("select id,name,address from companies")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var result []Company
	for rows.Next() {
		var each = Company{}
		var err = rows.Scan(&each.Id, &each.Name, &each.Address)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		result = append(result, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return result, nil
}
func (c *Company) ReadOne(id int) (Company, error) {
	var db, err = koneksi.ConnectDb()
	var result = Company{}
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	defer db.Close()

	err = db.QueryRow("select id, name, address from companies where id = ?", id).
		Scan(&result.Id, &result.Name, &result.Address)
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	return result, nil
}
func (c *Company) Update(id int) error {
	db, err := koneksi.ConnectDb()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()
	_, err = db.Exec("update companies set name=? , address=? where id=?", c.Name, c.Address, id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("company updated")
	return nil
}
func (c *Company) Delete(id int) error {
	db, err := koneksi.ConnectDb()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()
	_, err = db.Exec("delete from companies where id=?", id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("company deleted")
	return nil
}
