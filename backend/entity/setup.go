package entity

import (
	//"fmt"
	//"time"

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
		&Employee{},
		&Role{},
		&Province{},
		&MemberClass{},
		&User{},
	)

	db = database

	password1, err := bcrypt.GenerateFromPassword([]byte("zaq1@wsX"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("zxvseta"), 14)
	password3, err := bcrypt.GenerateFromPassword([]byte("1111111111111"), 14)

	//add example data
	//emp

	db.Model(&Employee{}).Create(&Employee{
		Name:     "Sirinya",
		Email:    "sirinya@mail.com",
		Password: string(password1),
	})

	db.Model(&Employee{}).Create(&Employee{
		Name:     "Attawit",
		Email:    "attawit@mail.com",
		Password: string(password2),
	})

	var sirin Employee
	db.Raw("SELECT * FROM employees WHERE email = ?", "sirinya@mail.com").Scan(&sirin)

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
		Employee:    sirin,
		Role:        student,
		Province:    korat,
		MemberClass: classic,
	})
}
