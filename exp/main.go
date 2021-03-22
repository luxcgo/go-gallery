package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&User{})

	// name, email := getInfo()
	// u := &User{
	// 	Name:  name,
	// 	Email: email,
	// }
	// if err = db.Create(u).Error; err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", u)

	var u User
	db.First(&u)
	if db.Error != nil {
		panic(db.Error)
	}
	fmt.Println(u)

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
