package order

type OrderCreateRequest struct {
	NameProduct []string `json:"nameProduct"`
}

type OrderCreateResponse struct {
	OrderId uint `json:"orderId"`
}

type OrderGetByIdRequest struct {
	OrderId uint `json:"orderId"`
}

type OrderGetByIdResponse struct {
	Order *Order
}

type OrderGetAllResponse struct {
	Orders []*Order
}
