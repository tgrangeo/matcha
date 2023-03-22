package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/tgrangeo/matcha/models"
	"os"
	"encoding/json"
	"time"
)

func ConnectDb() *sql.DB {
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
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, fname TEXT NOT NULL, lname TEXT NOT NULL, Bio TEXT, mail TEXT NOT NULL, type JSONB NOT NULL, pokeball JSONB NOT NULL, birthdate TEXT NOT NULL, age INTEGER)")
	if err != nil {
		panic(err)
	}
	fmt.Println("table users created")
}


//////////////////////////////////////////////////////////////////////////////TODO: make an utils package 
func mailExists(db *sql.DB, user models.User) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE mail = $1"
	var count int
	err := db.QueryRow(query, user.Mail).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func toJson(t models.Type) []byte {
	ret , _  := json.Marshal(t)
	return ret
}

func fromJson(t []byte) models.Type {
	var data models.Type
	json.Unmarshal(t, &data)
	//ret := models.Type{data[0].name, data[0].desc}
	return data
}

func age(birthdate, today time.Time) int {
	today = today.In(birthdate.Location())
	ty, tm, td := today.Date()
	today = time.Date(ty, tm, td, 0, 0, 0, 0, time.UTC)
	by, bm, bd := birthdate.Date()
	birthdate = time.Date(by, bm, bd, 0, 0, 0, 0, time.UTC)
	if today.Before(birthdate) {
		return 0
	}
	age := ty - by
	anniversary := birthdate.AddDate(age, 0, 0)
	if anniversary.After(today) {
		age--
	}
	return age
}

func stringToTime(s string) time.Time{
	layout := "00/00/0000"
	t, _ := time.Parse(layout, s)
	return t
}
///////////////////////////////////////////////////////////////////////////////////////////////////////

func InsertUser(db *sql.DB, user models.User) {
	exists, err := mailExists(db, user)
	if err != nil {
		fmt.Println(err)
	}
	if exists {
		fmt.Println("mail already exist")
		return
	}


	//create json 
	typeJson := toJson(user.Type)
	pokeJson := toJson(user.Pokeball)

	age := age(stringToTime(user.BirthDate), time.Now())
	fmt.Println(age)

	insertStmt := `INSERT INTO users (fname,lname,Bio,mail,type,pokeball,birthdate,age) VALUES ($1, $2, $3, $4, $5, $6, $7,$8)`
	_, err = db.Exec(insertStmt, user.First_Name, user.Last_Name, user.Bio, user.Mail, typeJson, pokeJson, user.BirthDate, age)
	if err != nil {
		panic(err)
	}
	fmt.Println("User " + user.First_Name + " added !")
}

func UpdateUser(db *sql.DB, user models.User, tofind int) {
	//create json 
	typeJson := toJson(user.Type)
	pokeJson := toJson(user.Pokeball)
	age := age(stringToTime(user.BirthDate), time.Now())
	insertStmt := `UPDATE users SET fname = $1, lname = $2, Bio = $3, mail = $4, type = $5, pokeball = $6, birthdate = $8, age = $9 WHERE id = $7`
	_, err := db.Exec(insertStmt, user.First_Name, user.Last_Name, user.Bio, user.Mail, typeJson, pokeJson, tofind, user.BirthDate, age)
	if err != nil {
		panic(err)
	}
	fmt.Println("User " + user.First_Name + " updated !")
}

func GetUsers(db *sql.DB) []models.User {
	tab := []models.User{}
	rows, _ := db.Query("SELECT * FROM users")
	for rows.Next() {
		var id, age int
		var fname, lname, bio, mail, birthdate string
		var t, p []byte
		err := rows.Scan(&id, &fname, &lname, &bio, &mail, &t, &p, &birthdate, &age)
		usr := models.User{id, fname, lname, bio, mail, fromJson(t), fromJson(p), birthdate, age}
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
	var id,age int
	var fname, lname, bio, mail,birthdate string
	var t, p []byte
	row, err := db.Query("SELECT * FROM users WHERE id = $1", tofind)
	if err != nil {
		fmt.Println(err)
	}
	row.Next()
	row.Scan(&id, &fname, &lname, &bio, &mail, &t, &p,&birthdate, &age)
	usr := models.User{id, fname, lname, bio, mail, fromJson(t), fromJson(p), birthdate,age}
	// fmt.Println(usr)
	return usr
}

func GetUsersWhere(db *sql.DB, tofind string, value string) []models.User {
	fmt.Println(tofind, value)
	if (tofind == "type"){
		tofind = "type->>'name'" // permet d'aller chercher des elements dans la struct stockée en json (ici le nom du type) 
	}
	if (tofind == "pokeball"){
		tofind = "pokeball->>'name'" // ici le nom de la pokeball 
	}
	rows, err := db.Query("SELECT * FROM users WHERE " + tofind + " = $1", value)
	if err != nil {
		fmt.Println(err)
	}
	tab := []models.User{}
	for rows.Next() {
		var id,age int
		var fname, lname, bio, mail,birthdate string
		var t, p []byte
		err := rows.Scan(&id, &fname, &lname, &bio, &mail, &t, &p, &birthdate,&age)
		usr := models.User{id, fname, lname, bio, mail, fromJson(t), fromJson(p),birthdate,age}
		fmt.Println(usr)
		tab = append(tab, usr)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(tab)
	return tab;
}

func DelUserById(db *sql.DB, id int){
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
}

func DelUser(db *sql.DB){
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		panic(err)
	}
	fmt.Println("all users has been deleted !")
}


func DropUsers(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF EXIST users")
	if err != nil {
		panic(err)
	}
	fmt.Println("table users droped")
}
