package order

import (
	"net/http"
	"project/configs"
	"project/pkg/middleware"
	"project/pkg/req"
	"project/pkg/res"
	"strconv"
)

type IsAuthRepo interface {
	GetPhone(phoneNum string) (uint, error)
}

type OrderHandlerDeps struct {
	OrderRepository *OrderRepository
	AuthRepository  IsAuthRepo
	Config          *configs.Config
}

type OrderHandler struct {
	OrderRepository *OrderRepository
	AuthRepository  IsAuthRepo
}

func NewOrderHandler(router *http.ServeMux, deps OrderHandlerDeps) {
	handler := &OrderHandler{
		OrderRepository: deps.OrderRepository,
		AuthRepository:  deps.AuthRepository,
	}
	router.Handle("POST /order", middleware.IsAuth(handler.NewOrder(), deps.Config))
	router.Handle("GET /order/{id}", middleware.IsAuth(handler.GetOrderById(), deps.Config))
	router.Handle("GET /my-orders", middleware.IsAuth(handler.GetOrders(), deps.Config))
}

func (handler *OrderHandler) NewOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phoneNum, ok := r.Context().Value(middleware.ContextPhoneKey).(string)
		if !ok {
			return
		}
		phoneId, err := handler.AuthRepository.GetPhone(phoneNum)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		body, err := req.HandleBody[OrderCreateRequest](&w, r)
		if err != nil {
			return
		}
		orderId, err := handler.OrderRepository.CreateOrder(phoneId, body.NameProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp := &OrderCreateResponse{
			OrderId: orderId,
		}
		res.Json(w, resp, http.StatusCreated)
	}
}

func (handler *OrderHandler) GetOrderById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		order, err := handler.OrderRepository.GetOrderById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp := &OrderGetByIdResponse{
			Order: order,
		}
		res.Json(w, resp, http.StatusCreated)
	}
}

func (handler *OrderHandler) GetOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phoneNum, ok := r.Context().Value(middleware.ContextPhoneKey).(string)
		if !ok {
			return
		}
		phoneId, err := handler.AuthRepository.GetPhone(phoneNum)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		orders, err := handler.OrderRepository.GetOrders(phoneId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := &OrderGetAllResponse{
			Orders: orders,
		}
		res.Json(w, resp, http.StatusCreated)
	}
}
