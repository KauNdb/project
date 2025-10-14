package product

import "project/db"

type ProductRepository struct {
	DataBase *db.Db
}

func NewProductRepository(db *db.Db) *ProductRepository {
	return &ProductRepository{
		DataBase: db,
	}
}

func (repo *ProductRepository) Create(product *Product) (*Product, error) {
	result := repo.DataBase.DB.Create(product)
	if result != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *ProductRepository) Update(product *Product) (*Product, error) {
	result := repo.DataBase.DB.Updates(product)
	if result != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *ProductRepository) Delete(id uint) error {
	result := repo.DataBase.DB.Delete(&Product{}, id)
	if result != nil {
		return result.Error
	}

	return nil
}

func (repo *ProductRepository) GetById(id uint) (*Product, error) {
	var product Product
	result := repo.DataBase.DB.First(&product, id)
	if result != nil {
		return nil, result.Error
	}

	return &product, nil
}
