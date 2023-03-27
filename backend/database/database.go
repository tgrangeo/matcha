package database

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"github.com/tgrangeo/matcha/models"
	"os"
	
	//utils
	"encoding/json"
	"time"
	"golang.org/x/crypto/bcrypt"
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
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY,
		fname TEXT NOT NULL,
		lname TEXT NOT NULL,
		Bio TEXT,
		mail TEXT NOT NULL,
		type JSONB NOT NULL,
		pokeball JSONB NOT NULL,
		birthdate TEXT NOT NULL,
		age INTEGER,
		pass TEXT NOT NULL,
		gender INTEGER NOT NULL,
		desiredgender INTEGER NOT NULL,
		tags INTEGER ARRAY,
		userliked INTEGER ARRAY,
		likedfrom INTEGER ARRAY,
		seenfrom INTEGER ARRAY
	)`)
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
	layout := "01/02/2006"
	t, _ := time.Parse(layout, s)
	return t
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
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

	//crypt pass
	crypted , _ := HashPassword(user.Pass)

	//create json 
	typeJson := toJson(user.Type)
	pokeJson := toJson(user.Pokeball)

	//age from birthdate
	age := age(stringToTime(user.BirthDate), time.Now())
	fmt.Println(user.Pass)

	insertStmt := `INSERT INTO users (fname,lname,Bio,mail,type,pokeball,birthdate,age,pass,gender,desiredgender,tags,userliked,likedfrom,seenfrom) 
					VALUES ($1, $2, $3, $4, $5, $6, $7,$8, $9, $10, $11, $12, $13, $14, $15)`
	_, err = db.Exec(insertStmt, user.First_Name, user.Last_Name, user.Bio, user.Mail, typeJson, pokeJson, user.BirthDate, age, crypted, user.Gender, user.DesiredGender, pq.Array(user.Tags), pq.Array(user.UserLiked), pq.Array(user.LikedFrom), pq.Array(user.SeenFrom))
	if err != nil {
		panic(err)
	}
	fmt.Println("User " + user.First_Name + " added !")
}

func UpdateUser(db *sql.DB, user models.User, tofind int) {
	//create json 
	typeJson := toJson(user.Type)
	pokeJson := toJson(user.Pokeball)
	//crypt pass
	crypted , _ := HashPassword(user.Pass) 
	age := age(stringToTime(user.BirthDate), time.Now())
	insertStmt := `UPDATE users SET fname = $1, lname = $2, Bio = $3, mail = $4, type = $5, pokeball = $6, birthdate = $8, age = $9, pass = $10, gender = $11, desiredgender = $12, tags = $13, userliked = $14, likedfrom = $15, seenfrom = $16 WHERE id = $7`
	_, err := db.Exec(insertStmt, user.First_Name, user.Last_Name, user.Bio, user.Mail, typeJson, pokeJson, tofind, user.BirthDate, age, crypted, user.Gender, user.DesiredGender, pq.Array(user.Tags), pq.Array(user.UserLiked), pq.Array(user.LikedFrom), pq.Array(user.SeenFrom))
	if err != nil {
		panic(err)
	}
	fmt.Println("User " + user.First_Name + " updated !")
}

func GetUsers(db *sql.DB) []models.User {
	tab := []models.User{}
	rows, _ := db.Query("SELECT * FROM users")
	for rows.Next() {
		var id, age, gender, desired int
		var fname, lname, bio, mail, birthdate, pass string
		var t, p []byte
		var tags, userliked, likedfrom, seenfrom []int64
		err := rows.Scan(&id, &fname, &lname, &bio, &mail, &t, &p, &birthdate, &age, &pass, &gender,&desired,(*pq.Int64Array)(&tags),(*pq.Int64Array)(&userliked),(*pq.Int64Array)(&likedfrom),(*pq.Int64Array)(&seenfrom))
		usr := models.User{id, fname, lname, bio, mail, fromJson(t), fromJson(p), birthdate, age, pass, gender,desired,tags, userliked,likedfrom,seenfrom}
		tab = append(tab, usr)
		if err != nil {
			fmt.Println(err)
		}
	}
	return tab
}

func GetUsersById(db *sql.DB, tofind int) models.User {
	var id,age, gender, desired int
	var fname, lname, bio, mail,birthdate,pass string
	var t, p []byte
	var tags ,userliked, likedfrom,seenfrom[]int64
	row, err := db.Query("SELECT * FROM users WHERE id = $1", tofind)
	if err != nil {
		fmt.Println(err)
	}
	row.Next()
	row.Scan(&id, &fname, &lname, &bio, &mail, &t, &p,&birthdate, &age,&pass, &gender,&desired,(*pq.Int64Array)(&tags),(*pq.Int64Array)(&userliked),(*pq.Int64Array)(&likedfrom),(*pq.Int64Array)(&seenfrom))
	usr := models.User{id, fname, lname, bio, mail, fromJson(t), fromJson(p), birthdate,age, pass,gender,desired,tags,userliked,likedfrom,seenfrom}
	// fmt.Println(usr)
	return usr
}

func GetUsersWhere(db *sql.DB, tofind string, value string) []models.User {
	var rows *sql.Rows
	var err error
	if (tofind == "tags" || tofind == "userliked" || tofind == "likedfrom" || tofind == "seenfrom"){
		rows, err = db.Query("SELECT * FROM users WHERE "+ tofind + " @> ARRAY[" + value + "]::INTEGER[]")
	}else {
		fmt.Println(tofind, value)
		if (tofind == "type"){
			tofind = "type->>'name'" // permet d'aller chercher des elements dans la struct stockée en json (ici le nom du type) 
		}
		if (tofind == "pokeball"){
			tofind = "pokeball->>'name'" // ici le nom de la pokeball 
		}
		rows, err = db.Query("SELECT * FROM users WHERE " + tofind + " = $1", value)
		if err != nil {
			fmt.Println(err)
		}
	}
	tab := []models.User{}
	for rows.Next() {
		var id,age,gender,desired int
		var fname, lname, bio, mail,birthdate,pass string
		var t, p []byte
		var tags, userliked, likedfrom, seenfrom []int64
		err := rows.Scan(&id, &fname, &lname, &bio, &mail, &t, &p, &birthdate,&age,&pass,&gender,&desired,(*pq.Int64Array)(&tags),(*pq.Int64Array)(&userliked),(*pq.Int64Array)(&likedfrom),(*pq.Int64Array)(&seenfrom))
		usr := models.User{id, fname, lname, bio, mail, fromJson(t), fromJson(p),birthdate,age,pass,gender,desired,tags,userliked,likedfrom,seenfrom}
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
