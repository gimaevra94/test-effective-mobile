package structs

import "time"

type Subscription struct {
	ServiceName       string `json:"service_name" gorm:"column:service_name;uniqueIndex:user_service"`
	Price             int    `json:"price" gorm:"column:price"`
	UserID            string `json:"user_id" gorm:"column:user_id;uniqueIndex:user_service"`
	StartDate         string `json:"start_date" gorm:"column:start_date;type:date"`
	FormatedStartDate time.Time
}

type Responce struct {
	Msg string `json:"msg"`
}
