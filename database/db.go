package database

import (
	"fmt"
	"log"
	"os"

	"github.com/tgrangeo/go_serv/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDb() {
	var err error

	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect to databse. \n", err)
		os.Exit(1)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Todolist{})
}

func DropTable(){
	db.Migrator().DropTable(&models.Todolist{})
}

func InsertDB() {
	Create("task","tttttttttt")
	Create("effiency","eewqerqwerqweerqwerqw")
	Create("ff","se rendre")
	//printAll()
}

func checkName(name string) int{
	todos := FindAll()
	for value := range todos{
		if (name == todos[value].Name){
			return 1
		}
	}   
	return 0
}

func Create(name string, desc string) int{
	fmt.Println(name)
	//check if name already exist
	if (checkName(name) == 1){
		return 1
	}
	//create in db
	todo := models.Todolist{Name: name, Description: desc}
	db.Create(&todo)
	fmt.Println("new todo added:",name, desc)
	return 0
}

func Update(old string, newName string, newDescription string){
	db.Model(&models.Todolist{}).Where("name = ?", old).Update("Name", newName)
	db.Model(&models.Todolist{}).Where("name = ?", old).Update("Description", newDescription)
}

func printAll() {
	var todos []models.Todolist
	db.Find(&todos)
	for x := range todos {
		fmt.Println(todos[x].Name)
	}
}

func FindAll() []models.Todolist {
	var todos []models.Todolist
	db.Find(&todos)
	return todos
}

func DeleteByName(name string) {
	db.Where("name = ?", name).Delete(&models.Todolist{})
}

func DeleteAll(){
 	db.Exec("TRUNCATE TABLE todolists;")
	printAll()
}

func FindByName(name string) *models.Todolist {
	var ret models.Todolist
	db.Where("name = ?", name).First(ret)
	return &ret
}
