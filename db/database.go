package db

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type Database interface {
	CreateUser(user *User) error
	GetUser(email string) (*User, error)

	CreateLink(link *Link) error
	GetLink(alias string) (*Link, error)
	GetLinks(ownerID string) ([]*Link, error)
}

type MySql struct {
	*gorm.DB
}

type User struct {
	UserID   uuid.UUID `gorm:"primaryKey;not null"`
	Username string    `gorm:"unique;not null"`
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
}

type Link struct {
	gorm.Model
	URL     string    `gorm:"not null"`
	Alias   string    `gorm:"unique;not null"`
	OwnerID uuid.UUID `gorm:"not null"`
}

func NewMYSQLdb(username, password, host, port, dbName string) Database {
	var err error
	var count int64

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, host, port)
	initDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connection,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	initDB.Raw("SELECT count(*) FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = ?", dbName).Scan(&count)
	if count <= 0 {
		if err := initDB.Exec("CREATE DATABASE IF NOT EXISTS " + dbName).Error; err != nil {
			panic("Error creating database: " + err.Error())
		}
	}

	connection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connection,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	err = DB.AutoMigrate(&User{})
	if err != nil {
		panic(err.Error())
		return nil
	}
	err = DB.AutoMigrate(&Link{})
	if err != nil {
		panic(err.Error())
		return nil
	}
	return &MySql{DB}
}

func (db *MySql) CreateUser(user *User) error {
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *MySql) GetUser(email string) (*User, error) {
	var user User
	err := db.DB.Table("users").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *MySql) CreateLink(link *Link) error {
	err := db.Create(link).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *MySql) GetLink(alias string) (*Link, error) {
	var link Link
	err := db.DB.Table("links").Where("alias = ?", alias).First(&link).Error
	if err != nil {
		return nil, err
	}
	return &link, nil
}

func (db *MySql) GetLinks(ownerID string) ([]*Link, error) {
	var links []*Link
	err := db.DB.Table("links").Where("owner_id = ?", ownerID).Find(&links).Error
	if err != nil {
		return nil, err
	}
	return links, nil
}
