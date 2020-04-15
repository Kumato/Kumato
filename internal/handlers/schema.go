package handlers

const (
	cookieTokenName = "tkn"
)

var (
	uploadPath = ""
	dataPath   = ""
)

type responseMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type IDOnlyRequest struct {
	ID uint32 `json:"id"`
}
