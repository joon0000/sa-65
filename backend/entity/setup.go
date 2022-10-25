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

	password1, err := bcrypt.GenerateFromPassword([]byte("1111111111111"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password3, err := bcrypt.GenerateFromPassword([]byte("3333333333333"), 14)

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
		Password:  string(password1),
		Address:   "ถนน a อำเภอ v",
		//FK
		Role:        student,
		Province:    korat,
		MemberClass: classic,
	})

	db.Model(&User{}).Create(&User{
		Pin:       "E123456",
		FirstName: "kawin",
		LastName:  "l.",
		Civ:       "1234567890123",
		Phone:     "0899999999",
		Email:     "kawin@mail.com",
		Password:  string(password2),
		Address:   "หอ b อำเภอ r",
		//FK
		Role:        employee,
		Province:    chon,
		MemberClass: silver,
	})

	db.Model(&User{}).Create(&User{
		Pin:       "T8888",
		FirstName: "sirinya",
		LastName:  "kotpanya",
		Civ:       "3333333333333",
		Phone:     "0823456789",
		Email:     "sirinya@mail.com",
		Password:  string(password3),
		Address:   "บ้าน c อำเภอ q",
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
	// GetClass(3)
	//GetClassStruct(4)
	GetClassStruct2(1)
	GetClassStruct2(2)
	GetClassStruct2(3)

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
	us := User{}
	tx := db.Preload("MemberClass").First(&us, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}

	//fmt.Println(us.MemberClass.ID)
}

/* func GetClassStruct(id uint) {
	mb := []MemberClass{}
	tx := db.Preload("MemberClass").First(&mb, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
	fmt.Println(mb)
} */

func GetClassStruct2(id uint) {
	us := User{}
	mb := []MemberClass{}
	tx := db.Preload("MemberClass").First(&us, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
	er := db.Order("id").First(&mb, "id=?", us.MemberClass.ID)
	if er.Error != nil {
		fmt.Println(tx.Error)
	}
	return
}
