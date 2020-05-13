package model

import (
	"erp-go-v2/app/koneksi"
	"fmt"
)

type Product struct {
	Id            int    `json:"Id"`
	Name          string `json:"Name" validate:"required"`
	Cover         string `json:"Cover" validate:"required"`
	Cost          int64  `json:"Cost" validate:"required"`
	Price         int64  `json:"Price" validate:"required"`
	Average_price int64  `json:"Average_price" validate:"required"`
}

func (p *Product) Create() error {
	db, err := koneksi.ConnectDb()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()
	_, err = db.Exec("insert into products (name, cover, cost, price, average_price) values(?,?,?,?,?)", p.Name, p.Cover, p.Cost, p.Price, p.Average_price)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Product added")
	return nil
}
func (p *Product) Read() ([]Product, error) {
	db, err := koneksi.ConnectDb()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("select id, name, cover, cost, price, average_price from products")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var result []Product
	for rows.Next() {
		var each = Product{}
		var err = rows.Scan(&each.Id, &each.Name, &each.Cover, &each.Cost, &each.Price, &each.Average_price)
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
func (p *Product) ReadOne(id int) (Product, error) {
	var db, err = koneksi.ConnectDb()
	var result = Product{}
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	defer db.Close()

	err = db.QueryRow("select id, name, cover, cost, price, average_price from products where id = ?", id).
		Scan(&result.Id, &result.Name, &result.Cover, &result.Cost, &result.Price, &result.Average_price)
	if err != nil {
		fmt.Println(err.Error())
		return result, err
	}
	return result, nil
}
func (p *Product) Update(id int) error {
	db, err := koneksi.ConnectDb()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()
	_, err = db.Exec("update products set name=?, cover=?, cost=?, price=?, average_price=? where id=?", p.Name, p.Cover, p.Cost, p.Price, p.Average_price, id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Product updated")
	return nil
}
func (p *Product) Delete(id int) error {
	db, err := koneksi.ConnectDb()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer db.Close()
	_, err = db.Exec("delete from products where id=?", id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Product deleted")
	return nil
}
