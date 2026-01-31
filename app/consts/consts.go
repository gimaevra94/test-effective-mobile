package consts

// For api handlers
const (
	MethodNotAllowed    = "Method not allowed"
	BadInput            = "Bad input"
	EmptyValue          = "Empty value"
	InvalidDate         = "Invalid date. Use MM-YYYY"
	InternalServerError = "Internal server error"
)

// SQL requests
const (
	InsertQuery = "insert into subscriptions (service_name, price, user_id, start_date) values ($1, $2, $3, $4)"
	SelectQuery = "select * where user_id = ? and service_name = ? inline 1"
)
