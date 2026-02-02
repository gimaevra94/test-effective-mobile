package consts

import "github.com/gimaevra94/test-effective-mobile/app/consts"

// For api handlers
const (
	MethodNotAllowed    = "Method not allowed"
	BadInput            = "Incorrect input"
	EmptyValue          = "The input has an empty value"
	InvalidDate         = "Invalid date. Use MM-YYYY"
	NotExist            = "The row does not exist"
	InternalServerError = "Internal server error"
	AlreadyExist        = "The subscription already exists"
)

// For SQL Requests
const (
	InsertQuery = "insert into subscription (" + ServiceName + ", " + Price + ", " + UserID + ", " + StartDate + ") values ($1, $2, $3, $4)"
	SelectQuery = "select " + ServiceName + ", " + Price + ", " + UserID + ", " + StartDate + "from subscription where " + UserID + " = ? and " + ServiceName + " = ?"
)

// Vars
const (
	ServiceName = "service_name"
	Price       = "price"
	UserID      = "user_id"
	StartDate   = "start_date"
)

// Other
const (
	Driver       = "postgres"
	TimeFormat   = "01-2006"
	APIPathV1    = "/api/v1/subscription"
	User_id      = "user_id"
	Service_name = "service_name"
)
