package order

import (
	"project/db"
	"project/internal/product"
)

type OrderRepository struct {
	DataBase *db.Db
}

func NewOrderRepository(db *db.Db) *OrderRepository {
	return &OrderRepository{
		DataBase: db,
	}
}

func (repo *OrderRepository) CreateOrder(phoneId uint, products []string) (uint, error) {
	var productsSlice []product.Product
	var getProduct product.Product
	for _, product := range products {
		result := repo.DataBase.DB.First(&getProduct, "name = ?", product)
		if result.Error != nil {
			return 0, result.Error
		}
		productsSlice = append(productsSlice, getProduct)
	}
	order := &Order{
		PhoneId:  phoneId,
		Products: productsSlice,
	}
	result := repo.DataBase.DB.Create(order)
	if result.Error != nil {
		return 0, result.Error
	}
	return order.ID, nil
}

func (repo *OrderRepository) GetOrderById(orderId uint) (*Order, error) {
	var order Order
	result := repo.DataBase.DB.Preload("Products").First(&order, orderId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (repo *OrderRepository) GetOrders(phoneId uint) ([]*Order, error) {
	var orders []*Order
	result := repo.DataBase.DB.Preload("Products").Where(&Order{PhoneId: phoneId}).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}
