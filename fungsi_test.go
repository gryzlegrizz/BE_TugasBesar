package _tes

import (
	"fmt"
	"testing"

	// "github.com/gryzlegrizz/BE_TugasBesar/model"
	"github.com/gryzlegrizz/BE_TugasBesar/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
) 

func TestInsertParfume(t* testing.T){
	nama := 		"Van Persie"
	jenis := 		"Eau de Parfum"
	merk := 		"Chanel"
	deskripsi := 	"Parfum yang sangat harum dan tahan lama"
	harga :=		1500000
	thn := 			2019
	stok := 		100
	ukuran := 		"50ml"
	insertedID, err := module.InsertParfume(module.MongoConn, "parfume", nama, jenis, merk, deskripsi, harga, thn, stok, ukuran)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan ID: %v\n", insertedID.Hex())
}

func TestGetParfumeFromID(t *testing.T) {
    id := "66756b2bb0584a360d9709d8"
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        t.Fatalf("Error converting id to ObjectID: %v", err)
    }
    parfume, err := module.GetParfumeFromID(objectID, module.MongoConn, "parfume")
    if err != nil {
        t.Fatalf("Error calling GetParfumeFromID: %v", err)
    }
	fmt.Println(parfume)
}

func TestGetAllParfume(t *testing.T) {
	data := module.GetAllParfume(module.MongoConn, "parfume")
	fmt.Println(data)
}

func TestUpdateParfume(t *testing.T) {
	id := "66756b2bb0584a360d9709d8"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("Error converting id to ObjectID: %v", err)
	}
	nama := "Morris Garages"
	jenis := "Eau de Parfum"
	merk := "Morris"
	deskripsi := "Parfum yang harum dan murah"
	harga := 150000
	thn := 2018
	stok := 100
	ukuran := "50ml"
	err = module.UpdateParfume(objectID, module.MongoConn, "parfume", nama, jenis, merk, deskripsi, harga, thn, stok, ukuran)
	if err != nil {
		t.Fatalf("Error calling UpdateParfume: %v", err)
	}
	fmt.Println("Data berhasil diupdate")
}

func TestDeleteParfumeByID(t *testing.T) {
	id := "66756b2bb0584a360d9709d8"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("Error converting id to ObjectID: %v", err)
	}
	err = module.DeleteParfumeByID(objectID, module.MongoConn, "parfume")
	if err != nil {
		t.Fatalf("Error calling DeleteParfume: %v", err)
	}
	_, err = module.GetParfumeFromID(objectID, module.MongoConn, "parfume")
	if err == nil {
		t.Fatalf("Data masih ada")
	}
}

func TestInsertUser(t *testing.T){
	username := "user"
	password := "user"
	email := "user"
	phone := "08123456789"
	address := "Jl. Jalan"
	insertedID, err := module.InsertUser(module.MongoConn, "user", username, password, email, phone, address)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan ID: %v\n", insertedID.Hex())
}

func TestGetUserFromID(t *testing.T) {
	id := "667e684f71fbe2689d38ad23"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("Error converting id to ObjectID: %v", err)
	}
	user, err := module.GetUserFromID(objectID, module.MongoConn, "user")
	if err != nil {
		t.Fatalf("Error calling GetUserFromID: %v", err)
	}
	fmt.Println(user)
}

func TestUpdateUser(t *testing.T) {
	id := "667e684f71fbe2689d38ad23"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("Error converting id to ObjectID: %v", err)
	}
	username := "admin"
	password := "user"
	email := "user"
	phone := "000000000000"
	address := "Jl. Jalan"
	err = module.UpdateUser(objectID, module.MongoConn, "user", username, password, email, phone, address)
	if err != nil {
		t.Fatalf("Error calling UpdateUser: %v", err)
	}
	fmt.Println("Data berhasil diupdate")
}

func TestDeleteUserByID(t *testing.T) {
	id := "667e684f71fbe2689d38ad23"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("Error converting id to ObjectID: %v", err)
	}
	err = module.DeleteUserByID(objectID, module.MongoConn, "user")
	if err != nil {
		t.Fatalf("Error calling DeleteUser: %v", err)
	}
	_, err = module.GetUserFromID(objectID, module.MongoConn, "user")
	if err == nil {
		t.Fatalf("Data masih ada")
	}
}

func TestGetAllUser(t *testing.T) {
	data := module.GetAllUser(module.MongoConn, "user")
	fmt.Println(data)
}