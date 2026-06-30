package order

type CreateOrderRequest struct {
	TableID string `json:"table_id"`

	Notes string `json:"notes"`

	Items []CreateOrderItemRequest `json:"items"`
}

type UpdateOrderRequest struct {
	TableID string `json:"table_id"`

	Notes string `json:"notes"`

	Items []CreateOrderItemRequest `json:"items"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status"`
}

type CreateOrderItemRequest struct {
	MenuID uint `json:"menu_id"`
	Qty    int  `json:"qty"`
}
