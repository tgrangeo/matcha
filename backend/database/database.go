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
	//TODO: add login
	_, err := db.Exec(`DISCARD ALL`)
	if err != nil {
		panic(err)
	}
	tableDefs := []string{
		`DROP TABLE IF EXISTS users`,
		`CREATE TABLE users (
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
			tags TEXT ARRAY,
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
			login TEXT
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
	fmt.Println("table created")
	return nil
}

func CheckUsers(db *sql.DB) (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", "users")
	var count int
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
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
	crypted, _ := utils.HashPassword(user.Pass)

	//age from birthdate
	age := utils.Age(utils.StringToTime(user.Birthdate), time.Now())

	insertStmt := `INSERT INTO users (fname,lname,email,birthdate,pass,bio,imageurl,age,gender,desiredgender,fame,tags,pokeball,type,userliked,likedfrom,seenfrom,blocked,convlist,coord,notifs,isactive,temp_token,login) 
VALUES ($1, $2, $3, $4, $5, $6, $7,$8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24)`
	_, err := db.Exec(insertStmt, user.Fname, user.Lname, user.Email, user.Birthdate, crypted, user.Bio, pq.StringArray(user.Imageurl), age, user.Gender, pq.Array(user.Desiredgender), user.Fame, pq.Array(user.Tags),
		pq.Array(user.Pokeball), pq.Array(user.Type), pq.Array(user.Userliked), pq.Array(user.Likedfrom), pq.Array(user.Seenfrom), pq.Array(user.Blocked), pq.Array(user.Convlist), utils.LocToJson(user.Coord), utils.NotifsToJson(user.Notifs), user.Isactive, user.Token, user.Login)
	if err != nil {
		panic(err)
	}
	fmt.Println("User " + user.Fname + " added !")
}

func UpdateUser(db *sql.DB, user models.User) {
	crypted, _ := utils.HashPassword(user.Pass)
	age := utils.Age(utils.StringToTime(user.Birthdate), time.Now())
	insertStmt := `UPDATE users SET fname = $1, lname = $2, email = $3, birthdate = $4, pass = $5, bio = $6, imageurl = $8, age = $9, gender = $10, desiredgender = $11, fame = $12, tags = $13, pokeball = $14, type = $15, userliked = $16, likedfrom = $17, seenfrom = $18, blocked = $19, convlist = $20, coord = $21, notifs = $22, isactive = $23, login = $24, temp_token = $25 WHERE id = $7`
	_, err := db.Exec(insertStmt, user.Fname, user.Lname, user.Email, user.Birthdate, crypted, user.Bio, user.Id, pq.StringArray(user.Imageurl), age, user.Gender, pq.Array(user.Desiredgender), user.Fame, pq.Array(user.Tags),
		pq.Array(user.Pokeball), pq.Array(user.Type), pq.Array(user.Userliked), pq.Array(user.Likedfrom), pq.Array(user.Seenfrom), pq.Array(user.Blocked), pq.Array(user.Convlist), utils.LocToJson(user.Coord), utils.NotifsToJson(user.Notifs), user.Isactive, user.Login, user.Token)
	if err != nil {
		panic(err)
	}
	fmt.Println("User " + user.Fname + " updated !")
}

func GetUsers(db *sql.DB) []models.User {
	tab := []models.User{}
	rows, _ := db.Query("SELECT * FROM users")
	for rows.Next() {
		var login, fname, lname, email, birthdate, pass, bio, token string
		var tags, imageurl []string
		var id, age, gender, fame int64
		var Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, desiredgender []int64
		var coord, notifs []byte
		var isactive bool
		err := rows.Scan(&id, &fname, &lname, &email, &birthdate, &pass, &bio, pq.Array(&imageurl), &age, &gender, pq.Array(&desiredgender), &fame, pq.Array(&tags),
			pq.Array(&Type), pq.Array(&pokeball), pq.Array(&userliked), pq.Array(&likedfrom), pq.Array(&seenfrom),
			pq.Array(&blocked), pq.Array(&convlist), &coord, &notifs, &isactive, &token, &login)
		if err != nil {
			fmt.Println(err)
		}
		usr := models.User{id, login, fname, lname, email, birthdate, pass, bio, imageurl, age, gender, fame, desiredgender, tags, Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, utils.JsonToLoc(coord), utils.JsonToNotifs(notifs), isactive, token}
		tab = append(tab, usr)
	}
	return tab
}

func GetUsersById(db *sql.DB, tofind int) models.User {
	row, err := db.Query("SELECT * FROM users WHERE id = $1", tofind)
	if err != nil {
		fmt.Println(err)
	}
	row.Next()
	var login, fname, lname, email, birthdate, pass, bio, token string
	var tags, imageurl []string
	var id, age, gender, fame int64
	var Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, desiredgender []int64
	var coord, notifs []byte
	var isactive bool
	err = row.Scan(&id, &fname, &lname, &email, &birthdate, &pass, &bio, (*pq.StringArray)(&imageurl), &age, &gender, (*pq.Int64Array)(&desiredgender), &fame, (*pq.StringArray)(&tags),
		(*pq.Int64Array)(&Type), (*pq.Int64Array)(&pokeball), (*pq.Int64Array)(&userliked), (*pq.Int64Array)(&likedfrom), (*pq.Int64Array)(&seenfrom),
		(*pq.Int64Array)(&blocked), (*pq.Int64Array)(&convlist), &coord, &notifs, &isactive, &login, &token)
	usr := models.User{id, login, fname, lname, email, birthdate, pass, bio, imageurl, age, gender, fame, desiredgender, tags, Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, utils.JsonToLoc(coord), utils.JsonToNotifs(notifs), isactive, token}
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
		var login, fname, lname, email, birthdate, pass, bio, token string
		var tags, imageurl []string
		var id, age, gender, fame int64
		var Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, desiredgender []int64
		var coord, notifs []byte
		var isactive bool
		err = row.Scan(&id, &fname, &lname, &email, &birthdate, &pass, &bio, (*pq.StringArray)(&imageurl), &age, &gender, (*pq.Int64Array)(&desiredgender), &fame, (*pq.StringArray)(&tags),
			(*pq.Int64Array)(&Type), (*pq.Int64Array)(&pokeball), (*pq.Int64Array)(&userliked), (*pq.Int64Array)(&likedfrom), (*pq.Int64Array)(&seenfrom),
			(*pq.Int64Array)(&blocked), (*pq.Int64Array)(&convlist), &coord, &notifs, &isactive, &login, &token)
		usr := models.User{id, login, fname, lname, email, birthdate, pass, bio, imageurl, age, gender, fame, desiredgender, tags, Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, utils.JsonToLoc(coord), utils.JsonToNotifs(notifs), isactive, token}
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
		var login, fname, lname, email, birthdate, pass, bio, token string
		var tags, imageurl []string
		var id, age, gender, fame int64
		var Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, desiredgender []int64
		var coord, notifs []byte
		var isactive bool
		err = rows.Scan(&id, &fname, &lname, &email, &birthdate, &pass, &bio, (*pq.StringArray)(&imageurl), &age, &gender, (*pq.Int64Array)(&desiredgender), &fame, (*pq.StringArray)(&tags),
			(*pq.Int64Array)(&Type), (*pq.Int64Array)(&pokeball), (*pq.Int64Array)(&userliked), (*pq.Int64Array)(&likedfrom), (*pq.Int64Array)(&seenfrom),
			(*pq.Int64Array)(&blocked), (*pq.Int64Array)(&convlist), &coord, &notifs, &isactive, &token, &login)
		usr := models.User{id, login, fname, lname, email, birthdate, pass, bio, imageurl, age, gender, fame, desiredgender, tags, Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, utils.JsonToLoc(coord), utils.JsonToNotifs(notifs), isactive, token}

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
	_, err = db.Exec(`DISCARD ALL`)
	if err != nil {
		panic(err)
	}
	fmt.Println("table users dropped")
}

// func Initialize(db *sql.DB) {
// 	// Obtenir le nombre total d'utilisateurs dans la table users
// 	var userCount int
// 	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&userCount)
// 	if err != nil {
// 		fmt.Println("Erreur lors de la récupération du nombre d'utilisateurs:", err)
// 		return
// 	}

// 	// Parcourir chaque utilisateur et effectuer des opérations
// 	for i := 1; i <= userCount; i++ {
// 		usr := GetUsersById(db, i)
// 		// Effectuez ici les opérations souhaitées sur l'utilisateur (usr)
// 		fmt.Printf("Utilisateur #%d : %s %s\n", i, usr.Fname, usr.Birthdate)
// 		// Exemple : Mettre à jour l'utilisateur
// 		UpdateUser(db, usr)
// 	}
// }

func Initialize(db *sql.DB) {
	// Retrieve all users from the table users
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Erreur lors de la récupération des utilisateurs:", err)
		return
	}
	defer rows.Close()

	// Initialize a slice to store the users
	var users []models.User

	// Loop through the result set and populate the users slice
	for rows.Next() {
		var login, fname, lname, email, birthdate, pass, bio, token string
		var tags, imageurl []string
		var id, age, gender, fame int64
		var Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, desiredgender []int64
		var coord, notifs []byte
		var isactive bool
		err = rows.Scan(&id, &fname, &lname, &email, &birthdate, &pass, &bio, (*pq.StringArray)(&imageurl), &age, &gender, (*pq.Int64Array)(&desiredgender), &fame, (*pq.StringArray)(&tags),
			(*pq.Int64Array)(&Type), (*pq.Int64Array)(&pokeball), (*pq.Int64Array)(&userliked), (*pq.Int64Array)(&likedfrom), (*pq.Int64Array)(&seenfrom),
			(*pq.Int64Array)(&blocked), (*pq.Int64Array)(&convlist), &coord, &notifs, &isactive, &token, &login)
		usr := models.User{id, login, fname, lname, email, birthdate, pass, bio, imageurl, age, gender, fame, desiredgender, tags, Type, pokeball, userliked, likedfrom, seenfrom, blocked, convlist, utils.JsonToLoc(coord), utils.JsonToNotifs(notifs), isactive, token}
		users = append(users, usr)
	}

	// Process and update each user as needed
	for _, usr := range users {
		//age
		dob, _ := time.Parse("2006-01-02", usr.Birthdate)
		age := utils.Age(dob, time.Now())
		usr.Age = int64(age)
		//login
		usr.Login = fmt.Sprintf("User%d", usr.Id)
		//tags
		usr.Tags = utils.GenerateRandomTags()
		UpdateUser(db, usr)
	}
}
