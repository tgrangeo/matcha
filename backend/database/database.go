package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tgrangeo/matcha/models"
	"os"
)

func ConnectDb() *sql.DB {
	godotenv.Load(".env")
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"db", 5432, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("database connected !")
	return db
}

func CreateTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, fname TEXT NOT NULL, lname TEXT NOT NULL, Bio TEXT, mail TEXT NOT NULL)")
	if err != nil {
		panic(err)
	}
	fmt.Println("table users created")
}


//TODO: make in utils folder
func mailExists(db *sql.DB, user models.User) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE mail = $1"
	var count int
	err := db.QueryRow(query, user.Mail).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func InsertUser(db *sql.DB, user models.User) {
	exists, err := mailExists(db, user)
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

func UpdateUser(db *sql.DB, user models.User, tofind int) {
	insertStmt := `UPDATE users SET fname= $1, lname = $2, Bio = $3, mail = $4 WHERE id = $5`
	_, err := db.Exec(insertStmt, user.First_Name, user.Last_Name, user.Bio, user.Mail, tofind)
	if err != nil {
		panic(err)
	}
	fmt.Println("User " + user.First_Name + " updated !")
}

func GetUsers(db *sql.DB) []models.User {
	tab := []models.User{}
	rows, _ := db.Query("SELECT * FROM users")
	for rows.Next() {
		var id int
		var fname, lname, bio, mail string
		err := rows.Scan(&id, &fname, &lname, &bio, &mail)
		usr := models.User{id, fname, lname, bio, mail}
		fmt.Println(usr)
		tab = append(tab, usr)
		if err != nil {
			fmt.Println(err)
		}
	}
	// fmt.Println(tab)
	return tab
}

func GetUsersById(db *sql.DB, tofind int) models.User {
	var id int
	var fname, lname, bio, mail string
	row, err := db.Query("SELECT * FROM users WHERE id = $1", tofind)
	if err != nil {
		fmt.Println(err)
	}
	row.Next()
	row.Scan(&id, &fname, &lname, &bio, &mail)
	usr := models.User{id, fname, lname, bio, mail}
	// fmt.Println(usr)
	return usr
}

func DelUserById(db *sql.DB, id int){
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
}


func DropUsers(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF EXIST users")
	if err != nil {
		panic(err)
	}
	fmt.Println("table users droped")
}
