package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	_ "github.com/lib/pq"
	"github.com/tgrangeo/matcha/models"
)

func ConnectDb() *sql.DB {
	godotenv.Load(".env")
    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        "db", 5432, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
    db, err := sql.Open("postgres",connStr)
    if err != nil {
        panic(err)
    }
	fmt.Println("database connected !")
	return db
}


func CreateTable(db *sql.DB){
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, fname TEXT NOT NULL, lname TEXT NOT NULL, Bio TEXT, mail TEXT NOT NULL)")
    if err != nil {
        panic(err)
    }
    fmt.Println("table users created")
}


func mailExists(db *sql.DB, user models.User) (bool,error) {
		query := "SELECT COUNT(*) FROM users WHERE mail = $1"
		var count int
		err := db.QueryRow(query, user.Mail).Scan(&count)
		if err != nil {
			return false, err
		}
		return count > 0, nil
}

func  InsertUser(db *sql.DB, user models.User){
	exists, err := mailExists(db,user)
    if err != nil {
        fmt.Println(err)
    }
    if exists {
		fmt.Println("mail already exist")
        return
	}

    insertStmt := `INSERT INTO users (fname,lname,Bio,mail) VALUES ($1, $2, $3, $4)`
    _, err = db.Exec(insertStmt, user.First_Name, user.Last_Name, user.Bio, user.Mail)
    if err != nil {
        panic(err)
    }
    fmt.Println("User " + user.First_Name + " added !")
}

func GetUsers(db *sql.DB) []models.User {
	tab := []models.User{}
	rows, _ := db.Query("SELECT * FROM users")
	for rows.Next() {
		var id int
		var fname, lname, bio, mail string
		err := rows.Scan(&id, &fname, &lname, &bio, &mail)
		if err != nil {
			fmt.Println(err)
		}
	}
	return tab
}

func DropUsers(db *sql.DB){
	_, err := db.Exec("CREATE TABLE IF EXIST users")
    if err != nil {
        panic(err)
    }
    fmt.Println("table users droped")
}