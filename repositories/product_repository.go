package repositories

import (
	"go-assignment/models"

	"gorm.io/gorm"
)

type ProductRepositoryQ interface {
	GetProducts(products *[]models.Product)
	GetProduct(product *models.Product, id int)
}

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (productRepository *ProductRepository) GetProducts(products *[]models.Product, filters map[string]interface{}) {
	query := productRepository.DB
	for key, value := range filters {
		if value == "nil" {
			continue
		}
		if value != "" {
			switch key {
			case "category", "name", "sku":
				query = query.Where(key+" LIKE ?", "%"+value.(string)+"%")
			case "price__gte":
				query = query.Where("price >= ?", value)
			case "price__lte":
				if value.(float64) != 0 {
					query = query.Where("price <= ?", value)
				}
			}
		}
	}
	order := filters["order_by"]
	if order != "" {
		query = query.Order(order)
	}
	query.Find(products)
}

func (productRepository *ProductRepository) GetProduct(product *models.Product, id int) {
	productRepository.DB.Where("id = ? ", id).First(product)
}
