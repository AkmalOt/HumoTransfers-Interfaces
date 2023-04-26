package db

import (
	"awesomeProject2/models"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"io/ioutil"
	"log"

	"gorm.io/gorm"
)

//type Options struct {
//	withoutPrometheus bool
//}
//
//type GormOptions func(o *Options)

//func WithoutPrometheus() GormOptions {
//	return func(o *Options) {
//		o.withoutPrometheus = true
//	}
//}

var Data *gorm.DB

func SetupGorm() *gorm.DB {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	DB := &models.DbData{}
	err = yaml.Unmarshal(yamlFile, DB)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	log.Println(DB)

	conn, err := gorm.Open(postgres.Open(DB.DSN))
	if err != nil {
		log.Println(err, "Не удалось подключиться к базе данных")
		return nil
	}

	//err = conn.AutoMigrate(&UserCards{})
	log.Println("Success connection to", DB)
	Data = conn

	return conn

}
