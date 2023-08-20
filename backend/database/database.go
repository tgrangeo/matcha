package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/lib/pq"
	"github.com/tgrangeo/matcha/models"
	"github.com/tgrangeo/matcha/utils"
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

func CreateTable(db *sql.DB) error {

	tableDefs := []string{
		`CREATE TABLE IF NOT EXISTS users(
			id SERIAL PRIMARY KEY,
			fname TEXT,
			lname TEXT,
			email TEXT,
			birthdate TEXT,
			pass TEXT,
			bio TEXT,
			imageurl TEXT ARRAY,
			age INTEGER,
			gender INTEGER,
			desiredgender INTEGER ARRAY,
			fame INTEGER,
			tags INTEGER ARRAY,
			pokeball INTEGER ARRAY,
			type INTEGER ARRAY,
			userliked INTEGER ARRAY,
			likedfrom INTEGER ARRAY,
			seenfrom INTEGER ARRAY,
			blocked INTEGER ARRAY,
			convlist INTEGER ARRAY,
			coord JSONB,
			notifs JSONB,
			isactive BOOLEAN,
			temp_token TEXT,
		)`,
		`CREATE TABLE IF NOT EXISTS pokeballs( id SERIAL PRIMARY KEY, name TEXT, title TEXT, description TEXT)`,
		`CREATE TABLE IF NOT EXISTS tags( id SERIAL PRIMARY KEY, name TEXT)`,
		`CREATE TABLE IF NOT EXISTS types( id SERIAL PRIMARY KEY, name TEXT, description TEXT)`,
		`CREATE TABLE IF NOT EXISTS conversations( id SERIAL PRIMARY KEY, user1 INTEGER, user2 INTEGER, messages JSONB)`,
	}
	for _, tableDef := range tableDefs {
		_, err := db.Exec(tableDef)
		if err != nil {
			return err
		}
	}
	fmt.Println("Table created !")
	return nil
}

func InsertUser(db *sql.DB, user models.User) {
	fmt.Println(user)
	//TODO: remettre la secu
	// exists, err := utils.MailExists(db, user)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if exists {
	// 	fmt.Println("email already exist")
	// 	return
	// }

	//crypt pass
	fmt.Println("in db : " + user.Pass)
	crypted, _ := utils.HashPassword(user.Pass)

	//age from birthdate
	age := utils.Age(utils.StringToTime(user.Birthdate), time.Now())

	insertStmt := `INSERT INTO users (fname,lname,email,birthdate,pass,bio,imageurl,age,gender,desiredgender,fame,tags,pokeball,type,userliked,likedfrom,seenfrom,blocked,convlist,coord,notifs,isactive,temp_token) 
	VALUES ($1, $2, $3, $4, $5, $6, $7,$8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)`
	_, err := db.Exec(insertStmt, user.Fname, user.Lname, user.Email, user.Birthdate, crypted, user.Bio, pq.StringArray(user.Imageurl), age, user.Gender, pq.Array(user.Desiredgender), user.Fame, pq.Array(user.Tags),
		pq.Array(user.Pokeball), pq.Array(user.Type), pq.Array(user.Userliked), pq.Array(user.Likedfrom), pq.Array(user.Seenfrom), pq.Array(user.Blocked), pq.Array(user.Convlist), utils.LocToJson(user.Coord), utils.NotifsToJson(user.Notifs), user.Isactive, user.Token)
	if err != nil {
		panic(err)
	}
	fmt.Println("User " + user.Fname + " added !")
}

func UpdateUser(db *sql.DB, user models.User) {
	//crypt pass
	crypted, _ := utils.HashPassword(user.Pass)
	age := utils.Age(utils.StringToTime(user.Birthdate), time.Now())
	insertStmt := `UPDATE users SET fname = $1, lname = $2, email = $3, birthdate = $4, pass = $5, bio = $6, imageurl = $8, age = $9, gender = $10, desiredgender = $11, fame = $12, tags = $13, pokeball = $14, type = $15, userliked = $16, likedfrom = $17, seenfrom = $18, blocked = $19, convlist = $20, coord = $21, notifs = $22, isactive = $23, temp_token = $24 WHERE id = $7`
	_, err := db.Exec(insertStmt, user.Fname, user.Lname, user.Email, user.Birthdate, crypted, user.Bio, user.Id, pq.StringArray(user.Imageurl), age, user.Gender, pq.Array(user.Desiredgender), user.Fame, pq.Array(user.Tags),
		pq.Array(user.Pokeball), pq.Array(user.Type), pq.Array(user.Userliked), pq.Array(user.Likedfrom), pq.Array(user.Seenfrom), pq.Array(user.Blocked), pq.Array(user.Convlist), utils.LocToJson(user.Coord), utils.NotifsToJson(user.Notifs), user.Isactive, user.Token)
	if err != nil {
		panic(err)
	}
	fmt.Println("User " + user.Fname + " updated !")
}

func GetUsers(db *sql.DB) []models.User {
	tab := []models.User{}
	rows, _ := db.Query("SELECT * FROM users")
	for rows.Next() {
		var fname, lname, email, birthdate, pass, bio, token string
		var imageurl []string
		var id, age, gender, fame int64
		var tags, Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, desiredgender []int64
		var coord, notifs []byte
		var isactive bool
		err := rows.Scan(&id, &fname, &lname, &email, &birthdate, &pass, &bio, (*pq.StringArray)(&imageurl), &age, &gender, (*pq.Int64Array)(&desiredgender), &fame, (*pq.Int64Array)(&tags),
			(*pq.Int64Array)(&Type), (*pq.Int64Array)(&pokeball), (*pq.Int64Array)(&userliked), (*pq.Int64Array)(&likedfrom), (*pq.Int64Array)(&seenfrom),
			(*pq.Int64Array)(&blocked), (*pq.Int64Array)(&convlist), &coord, &notifs, &isactive, token)
		usr := models.User{id, fname, lname, email, birthdate, pass, bio, imageurl, age, gender, fame, desiredgender, tags, Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, utils.JsonToLoc(coord), utils.JsonToNotifs(notifs), isactive, token}
		tab = append(tab, usr)
		if err != nil {
			fmt.Println(err)
		}
	}
	return tab
}

func GetUsersById(db *sql.DB, tofind int) models.User {
	row, err := db.Query("SELECT * FROM users WHERE id = $1", tofind)
	if err != nil {
		fmt.Println(err)
	}
	row.Next()
	var fname, lname, email, birthdate, pass, bio, token string
	var imageurl []string
	var id, age, gender, fame int64
	var tags, Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, desiredgender []int64
	var coord, notifs []byte
	var isactive bool
	err = row.Scan(&id, &fname, &lname, &email, &birthdate, &pass, &bio, (*pq.StringArray)(&imageurl), &age, &gender, (*pq.Int64Array)(&desiredgender), &fame, (*pq.Int64Array)(&tags),
		(*pq.Int64Array)(&Type), (*pq.Int64Array)(&pokeball), (*pq.Int64Array)(&userliked), (*pq.Int64Array)(&likedfrom), (*pq.Int64Array)(&seenfrom),
		(*pq.Int64Array)(&blocked), (*pq.Int64Array)(&convlist), &coord, &notifs, &isactive, &token)
	usr := models.User{id, fname, lname, email, birthdate, pass, bio, imageurl, age, gender, fame, desiredgender, tags, Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, utils.JsonToLoc(coord), utils.JsonToNotifs(notifs), isactive, token}
	if err != nil {
		fmt.Println(err)
	}
	return usr
}

func GetUsersByEmail(db *sql.DB, tofind string) models.User {
	row, err := db.Query("SELECT * FROM users WHERE email = $1", tofind)
	if err != nil {
		fmt.Println(err)
	}
	if row.Next() {
		var fname, lname, email, birthdate, pass, bio, token string
		var imageurl []string
		var id, age, gender, fame int64
		var tags, Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, desiredgender []int64
		var coord, notifs []byte
		var isactive bool
		err = row.Scan(&id, &fname, &lname, &email, &birthdate, &pass, &bio, (*pq.StringArray)(&imageurl), &age, &gender, (*pq.Int64Array)(&desiredgender), &fame, (*pq.Int64Array)(&tags),
			(*pq.Int64Array)(&Type), (*pq.Int64Array)(&pokeball), (*pq.Int64Array)(&userliked), (*pq.Int64Array)(&likedfrom), (*pq.Int64Array)(&seenfrom),
			(*pq.Int64Array)(&blocked), (*pq.Int64Array)(&convlist), &coord, &notifs, &isactive, &token)
		usr := models.User{id, fname, lname, email, birthdate, pass, bio, imageurl, age, gender, fame, desiredgender, tags, Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, utils.JsonToLoc(coord), utils.JsonToNotifs(notifs), isactive, token}
		if err != nil {
			fmt.Println(err)
		}
		return usr
	}
	usr := models.User{}
	return usr
}

func GetUsersWhere(db *sql.DB, tofind string, value string) []models.User {
	var rows *sql.Rows
	var err error
	if tofind == "tags" || tofind == "userliked" || tofind == "likedfrom" || tofind == "seenfrom" || tofind == "type" || tofind == "pokeball" || tofind == "convlist" || tofind == "blocked" || tofind == "desiredgender" {
		rows, err = db.Query("SELECT * FROM users WHERE "+tofind+" @> ARRAY[$1]::INTEGER[]", value)
	} else if tofind == "imageurl" {
		rows, err = db.Query("SELECT * FROM users WHERE " + tofind + " @> ARRAY[" + value + "]::TEXT[]")
	} else if tofind == "loc" || tofind == "notifs" {
		return nil
	} else {
		rows, err = db.Query("SELECT * FROM users WHERE "+tofind+" = $1", value)
	}
	if err != nil {
		fmt.Println(err)
	}
	tab := []models.User{}
	for rows.Next() {
		var fname, lname, email, birthdate, pass, bio, token string
		var imageurl []string
		var id, age, gender, fame int64
		var tags, Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, desiredgender []int64
		var coord, notifs []byte
		var isactive bool
		err = rows.Scan(&id, &fname, &lname, &email, &birthdate, &pass, &bio, (*pq.StringArray)(&imageurl), &age, &gender, (*pq.Int64Array)(&desiredgender), &fame, (*pq.Int64Array)(&tags),
			(*pq.Int64Array)(&Type), (*pq.Int64Array)(&pokeball), (*pq.Int64Array)(&userliked), (*pq.Int64Array)(&likedfrom), (*pq.Int64Array)(&seenfrom),
			(*pq.Int64Array)(&blocked), (*pq.Int64Array)(&convlist), &coord, &notifs, &isactive, &token)
		usr := models.User{id, fname, lname, email, birthdate, pass, bio, imageurl, age, gender, fame, desiredgender, tags, Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, utils.JsonToLoc(coord), utils.JsonToNotifs(notifs), isactive, token}

		tab = append(tab, usr)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(tab)
	return tab
}

func DelUserById(db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
}

func DelUsers(db *sql.DB) {
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		panic(err)
	}
	fmt.Println("all users has been deleted !")
}

func DropUsers(db *sql.DB) {
	_, err := db.Exec("DROP TABLE users")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("DROP SEQUENCE IF EXISTS users_id_seq")
	if err != nil {
		panic(err)
	}
	fmt.Println("table users dropped")
}
