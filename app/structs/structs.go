package structs

type Subscription struct {
	ServiceName string `json:"service_name" gorm:"column:service_name"`
	Price       int    `json:"price" gorm:"column:price"`
	UserID      string `json:"user_id" gorm:"column:user_id"`
	StartDate   string `json:"start_date" gorm:"column:start_date"`
}

type Responce struct {
	Msg string `json:"msg"`
}
