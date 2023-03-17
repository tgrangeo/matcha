package models

import (
    _ "github.com/lib/pq"
)


type User struct {
    First_Name  string  `json:"First_Name"`
    Last_Name  string  `json:"Last_Name"`
    Bio string `json:"Bio"`
    Mail string `json:"Mail"`
}