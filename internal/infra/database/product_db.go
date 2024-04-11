package database

import (
	"github.com/wagaodev/Go-Expert-API/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error

	if sort != "asc" && sort != "desc" {
		sort = "asc" // Isso garante um fallback seguro
	}

	query := p.DB.Order("created_at " + sort).Limit(limit).Offset((page - 1) * limit)
	err = query.Find(&products).Error

	return products, err
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	if err := p.DB.First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}
