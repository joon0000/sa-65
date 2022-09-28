package entity

import (
	"gorm.io/gorm"
)

type EMPLOYEE struct {
	gorm.Model
	NAME     string
	PASSWORD string
	USERS    []USER `gorm:"foreignKey:EMP_ID"`
}

type ROLE struct {
	gorm.Model
	NAME        string
	BORROW_DAY  int
	BOOKROOM_HR int
	BOOKCOM_HR  int
	USERS       []USER `gorm:"foreignKey:ROLE_ID"`
}

type PROVINCE struct {
	gorm.Model
	NAME  string
	USERS []USER `gorm:"foreignKey:PROVINCE_ID"`
}

type MemberClass struct {
	gorm.Model
	NAME     string
	DISCOUNT string
	USERS    []USER `gorm:"foreignKey:MemberClass_ID"`
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
	EMP_ID         *uint
	ROLE_ID        *uint
	PROVINCE_ID    *uint
	MemberClass_ID *uint
	//JOIN
	PROVINCE    PROVINCE
	ROLE        ROLE
	MemberClass MemberClass
	EMP         EMPLOYEE
}
