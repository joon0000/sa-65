package entity

import (
	"gorm.io/gorm"
	//"time"
)

type EMPLOYEE struct {
	gorm.Model
	NAME     string
	PASSWORD string
	USERS    []USER `gorm:"foreignKey:empID"`
}

type ROLE struct {
	gorm.Model
	NAME        string
	BORROW_DAY  int
	BOOKROOM_HR int
	BOOKCOM_HR  int
	USERS       []USER `gorm:"foreignKey:roleID"`
}

type PROVINCE struct {
	gorm.Model
	NAME  string
	USERS []USER `gorm:"foreignKey:provinceID"`
}

type MemberClass struct {
	gorm.Model
	NAME     string
	DISCOUNT string
	USERS    []USER `gorm:"foreignKey:memberclassID"`
}

type USER struct {
	gorm.Model
	PIN       string `gorm:"uniqueIndex"`
	FirstName string
	LastName  string
	CIV       string `gorm:"uniqueIndex"`
	PHONE     string
	EMAIL     string `gorm:"uniqueIndex"`
	PASSWORD  string
	ADDRESS   string
	//FK
	empID         *uint
	roleID        *uint
	provinceID    *uint
	memberclassID *uint
	//JOIN
	PROVINCE    PROVINCE
	ROLE        ROLE
	MemberClass MemberClass
	EMP         EMPLOYEE
}
