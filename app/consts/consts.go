package consts

// For api handlers
const (
	MethodNotAllowed    = "Method not allowed"
	BadInput            = "Incorrect input"
	EmptyValue          = "The input has an empty value"
	InvalidDate         = "Invalid date. Use MM-YYYY"
	NotExist            = "The subscription does not exist"
	InternalServerError = "Internal server error"
	AlreadyExist        = "The subscription already exists"
)

// For SQL Requests
const (
	InsertQuery         = "insert into subscription (" + ServiceName + ", " + Price + ", " + UserID + ", " + StartDate + ") values ($1, $2, $3, $4)"
	SelectQuery         = "select " + ServiceName + ", " + Price + ", " + UserID + ", " + StartDate + "from subscription where " + ServiceName + " = $1 and " + UserID + " = $2"
	UpdateQuery         = "update subscription set " + Price + " = $1 where " + ServiceName + " = $2 and " + UserID + " = $3 returning " + ServiceName + ", " + Price + ", " + UserID + ", " + StartDate
	DeleteQuery         = "delete from subscription where " + ServiceName + " = $1 and " + UserID + " = &2"
	PriceSelectionQuery = "select " + Price + " from subscription where " + ServiceName + " = $1 and " + UserID + " = $2 and" + StartDate + " = $3"
)

// Requests paths values
const (
	ServiceName = "service_name"
	Price       = "price"
	UserID      = "user_id"
	StartDate   = "start_date"
	APIPathV1   = "/api/v1/subscription"
)

// Other
const (
	Driver     = "postgres"
	TimeFormat = "01-2006"
)
