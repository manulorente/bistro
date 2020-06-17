package models

import (
	//"errors"
	"html"
	"strings"
	"time"
	//"github.com/jinzhu/gorm"
)

type Product struct {
	//ID   			int       	`json:"id" gorm:"primary_key;auto_increment"`
	Category    	string   	`json:"cat,omitempty" gorm:"size:100;not null;unique"`
	Name        	string    	`json:"name" gorm:"size:100;not null;unique"`
	Description 	string    	`json:"description,omitempty" gorm:"size:255;not null"`
	Size        	string    	`json:"size" gorm:"size:100;not null"`
	Price       	string    	`json:"price" gorm:"size:8;not null"`
	User 			User		`json:"user"`
	UserID 			int			`json:"user_id" gorm:"not null"`
	CreatedAt   	time.Time   `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   	time.Time   `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (p *Product) Prepare() {
	//p.ID = 0
	p.Category = html.EscapeString(strings.TrimSpace(p.Category))
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	p.Size = html.EscapeString(strings.TrimSpace(p.Size))
	p.Price = html.EscapeString(strings.TrimSpace(p.Price))
	p.User = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

/*
func (p *Product) Validate() error {

	if p.Category == "" {
		return errors.New("Required Category")
	}
	if p.Name == "" {
		return errors.New("Required Name")
	}
	if p.Description == "" {
		return errors.New("Required Description")
	}
	if p.Size == "" {
		return errors.New("Required Size")
	}
	if p.Price == "" {
		return errors.New("Required Price")
	}
	if p.UserID < 1 {
		return errors.New("Required User")
	}
	return nil
}

func (p *Product) SaveProduct(db *gorm.DB) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Product{}, err
		}
	}
	return p, nil
}

func (p *Product) FindAllProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Limit(100).Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}
	if len(products) > 0 {
		for i, _ := range products {
			err := db.Debug().Model(&Product{}).Where("id = ?", products[i].UserID).Take(&products[i].User).Error
			if err != nil {
				return &[]Product{}, err
			}
		}
	}
	return &products, nil
}

func (p *Product) FindProductByID(db *gorm.DB, pid uint64) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&Product{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Product{}, err
		}
	}
	return p, nil
}

func (p *Product) UpdateProduct(db *gorm.DB) (*Product, error) {

	var err error

	err = db.Debug().Model(&Product{}).Where("id = ?", p.ID).Updates(Product{Category: p.Category, Name: p.Name, Description: p.Description, Size: p.Size, Price: p.Price, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Product{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Product{}, err
		}
	}
	return p, nil
}

func (p *Product) DeleteProduct(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Product{}).Where("id = ? and user_id = ?", pid, uid).Take(&Product{}).Delete(&Product{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Post not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

/*
// Return a list of all the articles
func GetAllProducts() []Product {
	return ProductsList
}

func CreateNewProduct(cat, name, description, size, price string) (*Product, error) {
    p := Product{
		ID: len(ProductsList) + 1, 
		Cat : cat,
		Name: name,
		Description: description,
		Size: size,
		Price: price + "€",
		CreatedAt: time.Now(),
	}

    ProductsList = append(ProductsList, p)

    return &p, nil
}*/
