package models

import (
	_ "github.com/lib/pq"
)

type Loc struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type Notif struct {
	From int64  `json:"from"`
	Msg  string `json:"msg"`
}

type Msg struct {
	From string `json:"from"`
	Time int64  `json:"time"`
	Text string `json:"text"`
}

type Conv struct {
	Id   any `json:"id"`
	Usr1 int64
	Usr2 int64
	Msg  []Msg
}

type Report struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type User struct {
	Id            any      `json:"id"`
	Login         string   `json:"login"`
	Fname         string   `json:"fname"`
	Lname         string   `json:"lname"`
	Email         string   `json:"email"`
	Birthdate     string   `json:"birthDate"`
	Pass          string   `json:"pass"`
	Bio           string   `json:"bio"`
	Imageurl      []string `json:"imageurl"` // first == profile picture                 max = 5              no image cannot be like
	Age           int64    `json:"age"`
	Gender        int64    `json:"gender"` // 0 homme 1 femme 2 non binaire
	Fame          int64    `json:"fame"`
	Desiredgender []int64  `json:"desiredgender"` // 0 homme 1 femme 2 non bianaires
	Tags          []int64  `json:"tags"`
	Type          []int64  `json:"type"`
	Pokeball      []int64  `json:"pokeball"`
	Userliked     []int64  `json:"userliked"`
	Likedfrom     []int64  `json:"likedfrom"`
	Seenfrom      []int64  `json:"seenfrom"`
	Blocked       []int64  `json:"blocked"`
	Convlist      []int64  `json:"convlist"`
	Coord         Loc      `json:"coord"`
	Notifs        []Notif  `json:"notifs"`
	Isactive      bool     `json:"isactive"`
	Token         string   `json:"token"`
}

type UserPublic struct {
	Login         string   `json:"login"`
	Fname         string   `json:"fname"`
	Lname         string   `json:"lname"`
	Birthdate     string   `json:"birthDate"`
	Bio           string   `json:"bio"`
	Imageurl      []string `json:"imageurl"` // first == profile picture                 max = 5              no image cannot be like
	Age           int64    `json:"age"`
	Gender        int64    `json:"gender"` // 0 homme 1 femme 2 non binaire
	Fame          int64    `json:"fame"`
	Desiredgender []int64  `json:"desiredgender"` // 0 homme 1 femme 2 non bianaires
	Tags          []int64  `json:"tags"`
	Type          []int64  `json:"type"`
	Pokeball      []int64  `json:"pokeball"`
	Coord         Loc      `json:"coord"`
	LikedByMe     bool     `json:"likedbyme"`
}

func NewUserPublic(usr User) *UserPublic {
	return &UserPublic{
		Login:         usr.Login,
		Fname:         usr.Fname,
		Lname:         usr.Lname,
		Birthdate:     usr.Birthdate,
		Bio:           usr.Bio,
		Imageurl:      usr.Imageurl,
		Age:           usr.Age,
		Gender:        usr.Gender,
		Fame:          usr.Fame,
		Desiredgender: usr.Desiredgender,
		Tags:          usr.Tags,
		Type:          usr.Type,
		Pokeball:      usr.Pokeball,
		Coord:         usr.Coord,
		LikedByMe:     false,
	}
}

//constructor for a new user
func NewSubUser(id interface{}, login, fname, lname, email, bd, pass, tok string, age, gender, poketype, pokeball int64) *User {
	return &User{
		Login:     login,
		Id:        id,
		Fname:     fname,
		Lname:     lname,
		Gender:    gender,
		Type:      []int64{poketype},
		Pokeball:  []int64{pokeball},
		Birthdate: bd,
		Email:     email,
		Pass:      pass,
		Token:     tok,
		// Initialize all other fields
		Age:           0,
		Desiredgender: []int64{},
		Bio:           "",
		Imageurl:      []string{"https://cdn.pixabay.com/photo/2015/04/23/22/00/tree-736885__480.jpg"},
		Fame:          0,
		Tags:          []int64{},
		Userliked:     nil,
		Likedfrom:     nil,
		Seenfrom:      nil,
		Blocked:       nil,
		Convlist:      nil,
		Coord:         Loc{},
		Notifs:        nil,
		Isactive:      false,
	}
}

type NewUserInput struct {
	Login     string `json:"login"`
	Birthdate string `json:"birthdate"`
	Type      int64  `json:"type"`
	Pokeball  int64  `json:"pokeball"`
	Gender    int64 `json:"gender"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Email     string `json:"email"`
	Pass      string `json:"pass"`
}