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
    Age int `json:"Age"`
    Pass string `json:"Pass"`
    // Fame int `json:Pass`

    Gender int `json:"Gender"`  // 0 homme 1 femme 2 autres
    DesiredGender int `json:"DesiredGender"`// 0 homme 1 femme 2 autres (3h+f ???? )


    Tags []string `json:"Tags"`
    // UserLiked []string `json:"UserLiked"`
    // LikedFrom []string `json:"LikedFrom"`
    // SeenFrom []string `json:"SeenFrom"`


    //v2
    //image 
    //localisation

    //INT
    //want
    //gender

    //ARRAY
    //array of tags
    //likedfrom
    //userliked
    //user who see my profile
}