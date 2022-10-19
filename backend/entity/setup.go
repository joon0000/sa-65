package entity

import (
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
	//add example data
	//emp
	/* 	sirin := Employee{
	   		NAME:     "Sirinya",
	   		PASSWORD: "zaq1@wsX",

	   	}
	   	db.Model(&Employee{}).Create(&sirin)

	   	atta := Employee{
	   		NAME:     "Attawit",
	   		PASSWORD: "zxvseta",
	   	}
	   	db.Model(&Employee{}).Create(&atta) */

	db.Model(&Employee{}).Create(&Employee{
		NAME:     "Sirinya",
		Email:    "sirinya@mail.com",
		PASSWORD: "zaq1@wsX",
	})

	db.Model(&Employee{}).Create(&Employee{
		NAME:     "Attawit",
		Email:    "attawit@mail.com",
		PASSWORD: "zxvseta",
	})

	var sirin Employee
	db.Raw("SELECT * FROM employees WHERE email = ?", "sirinya@mail.com").Scan(&sirin)

	//Role
	student := Role{
		NAME:        "Student",
		BORROW_DAY:  3,
		BOOKROOM_HR: 3,
		BOOKCOM_HR:  4,
	}
	db.Model(&Role{}).Create(&student)

	teacher := Role{
		NAME:        "Teacher",
		BORROW_DAY:  7,
		BOOKROOM_HR: 12,
		BOOKCOM_HR:  12,
	}
	db.Model(&Role{}).Create(&teacher)

	//province
	korat := Province{
		NAME: "Nakhon Ratchasima",
	}
	db.Model(&Province{}).Create(&korat)

	chon := Province{
		NAME: "Chonburi",
	}
	db.Model(&Province{}).Create(&chon)

	bangkok := Province{
		NAME: "Bangkok",
	}
	db.Model(&Province{}).Create(&bangkok)

	//member
	classic := MemberClass{
		NAME:     "classic",
		DISCOUNT: 0,
	}
	db.Model(&MemberClass{}).Create(&classic)

	silver := MemberClass{
		NAME:     "silver",
		DISCOUNT: 5,
	}
	db.Model(&MemberClass{}).Create(&silver)

	gold := MemberClass{
		NAME:     "gold",
		DISCOUNT: 10,
	}
	db.Model(&MemberClass{}).Create(&gold)

	plat := MemberClass{
		NAME:     "platinum",
		DISCOUNT: 20,
	}
	db.Model(&MemberClass{}).Create(&plat)

	//user
	db.Model(&User{}).Create(&User{
		PIN:       "B6111111",
		FirstName: "preecha",
		LastName:  "anpa",
		CIV:       "1111111111111",
		PHONE:     "0811111111",
		EMAIL:     "preechapat@mail.com",
		PASSWORD:  "1111111111111",
		ADDRESS:   "ถนน a อำเภอ v",
		//FK
		Employee:    sirin,
		Role:        student,
		Province:    korat,
		MemberClass: classic,
	})
}
