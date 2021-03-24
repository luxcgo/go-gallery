package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/luxcgo/go-gallery/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "secret"
	dbname   = "luxcgo_gallery"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;uniqueIndex"`
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, port, user, password, dbname)
	us, err := models.NewUserService(dsn)
	if err != nil {
		panic(err)
	}

	// Update the call to ByID to instead be ByEmail
	foundUser, err := us.ByID(3)
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)

	us.Update(&models.User{Model: gorm.Model{ID: 3}, Name: "sm", Email: "m@s.com"})
	foundUser, err = us.ByID(3)
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)
}

func main2() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, port, user, password, dbname)
	us, err := models.NewUserService(dsn)
	if err != nil {
		panic(err)
	}

	// use models packge methods
	// query a user
	user, err := us.ByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	// drop the table and recreate it
	us.DestructiveReset()

	// query a user
	user, err = us.ByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	// create a user
	u := models.User{
		Name:  "Michael Scott",
		Email: "michael@dundermifflin.com",
	}
	if err := us.Create(&u); err != nil {
		panic(err)
	}

	// query a user
	foundUser, err := us.ByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)
}

func main1() {
	// connect to db
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, port, user, password, dbname)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// create table
	db.AutoMigrate(&User{})

	// create a user
	name, email := getInfo()
	u := &User{
		Name:  name,
		Email: email,
	}
	if err = db.Create(u).Error; err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", u)

	// query a user
	var v User
	db.First(&v)
	if db.Error != nil {
		panic(db.Error)
	}
	fmt.Println(v)

}

func getInfo() (name, email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("What is your email?")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)
	return name, email
}
