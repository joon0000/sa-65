package entity

import (
	//"fmt"
	//"time"

	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Role{},
		&Province{},
		&MemberClass{},
		&User{},
	)

	db = database

	password1, err := bcrypt.GenerateFromPassword([]byte("zaq1@wsX"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password3, err := bcrypt.GenerateFromPassword([]byte("1111111111111"), 14)

	//add example data

	//Role
	student := Role{
		Name:       "Student",
		BorrowDay:  3,
		BookRoomHR: 3,
		BookComHR:  4,
	}

	db.Model(&Role{}).Create(&student)

	teacher := Role{
		Name:       "Teacher",
		BorrowDay:  7,
		BookRoomHR: 12,
		BookComHR:  12,
	}
	db.Model(&Role{}).Create(&teacher)

	employee := Role{
		Name:       "Employee",
		BorrowDay:  5,
		BookRoomHR: 6,
		BookComHR:  6,
	}
	db.Model(&Role{}).Create(&employee)

	//province
	korat := Province{
		Name: "Nakhon Ratchasima",
	}
	db.Model(&Province{}).Create(&korat)

	chon := Province{
		Name: "Chonburi",
	}
	db.Model(&Province{}).Create(&chon)

	bangkok := Province{
		Name: "Bangkok",
	}
	db.Model(&Province{}).Create(&bangkok)

	//member
	classic := MemberClass{
		Name:     "classic",
		Discount: 0,
	}
	db.Model(&MemberClass{}).Create(&classic)

	silver := MemberClass{
		Name:     "silver",
		Discount: 5,
	}
	db.Model(&MemberClass{}).Create(&silver)

	gold := MemberClass{
		Name:     "gold",
		Discount: 10,
	}
	db.Model(&MemberClass{}).Create(&gold)

	plat := MemberClass{
		Name:     "platinum",
		Discount: 20,
	}
	db.Model(&MemberClass{}).Create(&plat)

	//user
	db.Model(&User{}).Create(&User{
		Pin:       "B6111111",
		FirstName: "preecha",
		LastName:  "anpa",
		Civ:       "1111111111111",
		Phone:     "0811111111",
		Email:     "preechapat@mail.com",
		Password:  string(password3),
		Address:   "ถนน a อำเภอ v",
		//FK
		Role:        student,
		Province:    korat,
		MemberClass: classic,
	})

	db.Model(&User{}).Create(&User{
		Pin:       "E123456",
		FirstName: "Sirinya",
		LastName:  "kot",
		Civ:       "1234567890123",
		Phone:     "0899999999",
		Email:     "sirinya@mail.com",
		Password:  string(password1),
		Address:   "ถนน c อำเภอ z",
		//FK
		Role:        employee,
		Province:    bangkok,
		MemberClass: plat,
	})

	db.Model(&User{}).Create(&User{
		Pin:       "T654321",
		FirstName: "Wichai",
		LastName:  "Micro",
		Civ:       "3210987654321",
		Phone:     "0823456789",
		Email:     "wichai@mail.com",
		Password:  string(password2),
		Address:   "ถนน c อำเภอ z",
		//FK
		Role:        teacher,
		Province:    bangkok,
		MemberClass: plat,
	})

	/* 	test := GetRoleName(1)
	   	fmt.Printf("User id 1: %s\n", test)
	   	test2 := GetRoleName(2)
	   	fmt.Printf("User id 2: %s\n", test2)
	   	test3 := GetRoleName(3)
	   	fmt.Printf("User id 3: %s\n", test3) */
	GetClass(1)

}
func GetRoleName(id uint) string {
	rn := User{}
	tx := db.Preload("Role").First(&rn, id)
	if tx.Error != nil {
		return "Role not found"
	} else if rn.Role.Name == "Employee" {
		return "admin"
	} else if rn.Role.Name == "Student" || rn.Role.Name == "Teacher" {
		return "user"
	}
	return "err"
}

func GetClass(id uint) {
	class := MemberClass{}
	tx := db.Preload("User").First(&class, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
	fmt.Println(class)
	fmt.Println("ss")
}
