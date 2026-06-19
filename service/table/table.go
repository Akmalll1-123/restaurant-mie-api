package table

type CreateTableRequest struct {
	TableNo  string `json:"table_no"`
	Capacity int    `json:"capacity"`
}

type UpdateTableRequest struct {
	TableNo  string `json:"table_no"`
	Capacity int    `json:"capacity"`
	IsActive bool   `json:"is_active"`
}
