package entity

import (
	"gorm.io/gorm"
	//"time"
)

type Employee struct {
	gorm.Model
	NAME     string
	Email    string `gorm:"uniqueIndex"`
	PASSWORD string
	USERS    []User `gorm:"foreignKey:EmployeeID"`
}

type Role struct {
	gorm.Model
	NAME        string
	BORROW_DAY  int
	BOOKROOM_HR int
	BOOKCOM_HR  int
	USERS       []User `gorm:"foreignKey:RoleID"`
}

type Province struct {
	gorm.Model
	NAME  string
	USERS []User `gorm:"foreignKey:ProvinceID"`
}

type MemberClass struct {
	gorm.Model
	NAME     string
	DISCOUNT int
	USERS    []User `gorm:"foreignKey:MemberClassID"`
}

type User struct {
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
	EmployeeID    *uint
	RoleID        *uint
	ProvinceID    *uint
	MemberClassID *uint
	//JOIN
	Province    Province
	Role        Role
	MemberClass MemberClass
	Employee    Employee
}
