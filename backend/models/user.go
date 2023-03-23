package models

import (
    _ "github.com/lib/pq"
)


type User struct {
    Id any `json:"Id"`
    First_Name  string  `json:"First_Name"`
    Last_Name  string  `json:"Last_Name"`
    Bio string `json:"Bio"`
    Mail string `json:"Mail"`
    Type Type `json:"Type"`
    Pokeball Type `json:"Pokeball"`
    BirthDate string `json:"BirthDate"`
    Age int `json:"age"`
    Pass string `json:"pass"`



    //v2
    //image 
    //localisation
    //want
    //gender 
    //array of tags
    //likedfrom
    //userliked
    //user who see my profile
}