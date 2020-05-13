package model

import (
	"erp-go-v2/app/koneksi"
	"fmt"
)

type Vendor struct {
	Id      int    `json:"Id"`
	Number  string `json:"Number" validate:"required"`
	Name    string `json:"Name" validate:"required"`
	Address string `json:"Address" validate:"required"`
	// Contact_person string `json:"Contact_person"`
}

func (v *Vendor) Create() error {
	db, err := koneksi.ConnectDb()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()
	// _, err = db.Exec("insert into vendors (number, name, address, contact_person) values(?,?,?,?)", v.Number, v.Name, v.Address, v.Contact_person)
	_, err = db.Exec("insert into vendors (number, name, address) values(?,?,?)", v.Number, v.Name, v.Address)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Vendor added")
	return nil
}
func (v *Vendor) Read() ([]Vendor, error) {
	db, err := koneksi.ConnectDb()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer db.Close()
	// rows, err := db.Query("select id, number, name, address, contact_person from vendors")
	rows, err := db.Query("select id, number, name, address from vendors")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var result []Vendor
	for rows.Next() {
		var each = Vendor{}
		// var err = rows.Scan(&each.Id, &each.Number, &each.Name, &each.Address, &each.Contact_person)
		var err = rows.Scan(&each.Id, &each.Number, &each.Name, &each.Address)
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
func (v *Vendor) ReadOne(id int) (Vendor, error) {
	var db, err = koneksi.ConnectDb()
	var result = Vendor{}
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	defer db.Close()

	// err = db.QueryRow("select id, number, name, address, contact_person from vendors where id = ?", id).
	// 	Scan(&result.Id, &result.Number, &result.Name, &result.Address, &result.Contact_person)
	err = db.QueryRow("select id, number, name, address from vendors where id = ?", id).
		Scan(&result.Id, &result.Number, &result.Name, &result.Address)
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	return result, nil
}
func (v *Vendor) Update(id int) error {
	db, err := koneksi.ConnectDb()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()
	_, err = db.Exec("update vendors set number=?, name=?, address=? where id=?", v.Number, v.Name, v.Address, id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Vendor updated")
	return nil
}
func (v *Vendor) Delete(id int) error {
	db, err := koneksi.ConnectDb()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()
	_, err = db.Exec("delete from vendors where id=?", id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Vendor deleted")
	return nil
}
