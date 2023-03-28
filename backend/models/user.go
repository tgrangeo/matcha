package models

import (
    _ "github.com/lib/pq"
)

type Loc struct{
    Lat float64 `json:"lat"`
    Long float64 `json:"long"`
} 

type Notif struct{
    From int64 `json:"from"`
    Msg string `json:"msg"`
}

type Msg struct{
    From string `json:"from"`
    Time int64 `json:"time"` 
    Text string `json:"text"` 
}

type Conv struct {
    Id any `json:"id"`
    Usr1 int64 
    Usr2 int64
    Msg []Msg 
}

type Report struct{
    From string `json:"from"`
    To string  `json:"to"`
}

type User struct {
    Id any `json:"id"`
    Fname  string  `json:"fname"`
    Lname  string  `json:"lname"`
    Email string `json:"email"`
    Birthdate string `json:"birthDate"`
    Pass string `json:"pass"`
    Bio string `json:"bio"`
    Imageurl []string `json:"imageurl"` // first == profile picture                 max = 5              no image cannot be like
    Age int64 `json:"age"`
    Gender int64 `json:"gender"`  // 0 homme 1 femme 2 non binaire 
    Fame int64 `json:"fame"`
    Desiredgender []int64 `json:"desiredgender"`// 0 homme 1 femme 2 non bianaires
    Tags []int64 `json:"tags"`
    Type []int64 `json:"type"`
    Pokeball []int64 `json:"pokeball"`
    Userliked []int64 `json:"userliked"`
    Likedfrom []int64 `json:"likedfrom"`
    Seenfrom []int64 `json:"seenfrom"`
    Blocked []int64 `json:"blocked"`
    Convlist []int64 `json:"convlist"`
    Coord Loc `json:"coord"`
    Notifs []Notif `json:"notifs"`
    Isactive bool `json:"isactive"`
}
