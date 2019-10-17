package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordID      int
	customerID int
}

type employee struct {
	name    string
	ID      int
	address string
	salary  int
	country string
}

// func createQuery(o order) string {
// 	i := fmt.Sprintf("INSERT INTO order VALUES(%d, %d)\n", o.ordID, o.customerID)
// 	return i
// }
// func createQuery(q interface{}) string {
// 	// t := reflect.TypeOf(q)
// 	// v := reflect.ValueOf(q)
// 	// k := t.Kind()

// 	// fmt.Println("Type: ", t)
// 	// fmt.Println("Value: ", v)
// 	// fmt.Println("Kind: ", k)

// 	if reflect.ValueOf(q).Kind() == reflect.Struct {
// 		v := reflect.ValueOf(q)
// 		fmt.Println("Nb of fields", v.NumField())
// 		for i := 0; i < v.NumField(); i++ {
// 			fmt.Printf("Field %d type %T value %v\n", i, v.Field(i), v.Field(i))
// 		}
// 	}

// 	return "pouet"
// }
func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")
}

func main() {
	// intro
	i := 10
	fmt.Printf("%d %T\n", i, i)

	//
	o := order{
		ordID:      1234,
		customerID: 567,
	}
	createQuery(o)

	e := employee{
		name:    "Pouet",
		ID:      565,
		address: "Somewhere DTC",
		salary:  123456,
		country: "Pouet pouet",
	}
	createQuery(e)

	ii := 90
	createQuery(ii)
}
