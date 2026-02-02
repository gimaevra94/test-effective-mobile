package consts

// For api handlers
const (
	MethodNotAllowed    = "Method not allowed"
	BadInput            = "Incorrect input"
	EmptyValue          = "The input has an empty value"
	InvalidDate         = "Invalid date. Use MM-YYYY"
	NotExist            = "The row does not exist"
	InternalServerError = "Internal server error"
)

// SQL requests
const (
	InsertQuery = "insert into subscriptions (service_name, price, user_id, start_date) values ($1, $2, $3, $4)"
	SelectQuery = "select service_name, price, user_id, start_date from subscriptions where user_id = ? and service_name = ?"
)
