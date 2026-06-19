package response

type Success struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Error struct {
	Message string `json:"message"`
}
