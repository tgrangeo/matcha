package utils 

import (
	"database/sql"
	"github.com/tgrangeo/matcha/models"
	"encoding/json"
	"time"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
	"regexp"
)

func MailExists(db *sql.DB, user models.User) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE email = $1"
	var count int
	err := db.QueryRow(query, user.Email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func LocToJson(t models.Loc) []byte {
	ret, _ := json.Marshal(t)
	return ret
}

func JsonToLoc(t []byte) models.Loc {
	var data models.Loc
	json.Unmarshal(t, &data)
	return data
}

func NotifsToJson(t []models.Notif) []byte {
	ret, _ := json.Marshal(t)
	return ret
}

func JsonToNotifs(t []byte) []models.Notif {
	var data []models.Notif
	json.Unmarshal(t, &data)
	return data
}

func Age(birthdate, today time.Time) int {
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

func StringToTime(s string) time.Time {
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

func SpecialChar(str string) bool{
	ret, _ := regexp.MatchString("/^[a-z\\d\\-_\\s]+$/i", str)
	return ret
}

func CheckUser(user models.User) (int,string){
	//mail 
	_, err := mail.ParseAddress(user.Email)
    if (err != nil){
		return 1 , "invalid mail format"
	}
	//age > 18
	a := Age(StringToTime(user.Birthdate), time.Now())
	if a < 18 {
		return 1 , "this guy is very young"
	}
	//bio max 250 char +  no special char 
	if SpecialChar(user.Bio){
		return 1 , "special char in bio"
	}
	if len(user.Bio) > 250{
		return 1 , "bio too long"
	}
	//pass regex: Minimum eight characters, at least one uppercase letter, one lowercase letter and one number
	ret, _ := regexp.MatchString("^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)[a-zA-Z\\d]{8,}$", user.Pass)
	if ret{
		return 1 , "this pass is not enough secure"
	}
	//lname fname +  no special char
	if SpecialChar(user.Lname){
		return 1 , "special char in last name"
	}
	if SpecialChar(user.Fname){
		return 1 , "special char in first name"
	}

	return 0, ""
}