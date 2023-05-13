package repository

import (
	"golang-store/model/entity"
	"golang-store/utils"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepositoryImpl(db *gorm.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{
		db: db,
	}
}

// Create a new category
func (r *CategoryRepositoryImpl) Create(category entity.Category) (*entity.Category, error) {
	if err := r.db.Create(&category); err != nil {
		return nil, err
	}

	return &category, nil
}

// Update a category
func (r *CategoryRepositoryImpl) Update(category entity.Category) (*entity.Category, error) {
	if err := r.db.Update(&category); err != nil {
		return nil, err
	}
	return &category, nil
}

// find category by id
func (r *CategoryRepositoryImpl) FindByID(id string) (*entity.Category, error) {
	var category entity.Category
	if err := r.db.Where("id =?", id).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// Find all categories
func (r *CategoryRepositoryImpl) FindAll(pagination utils.Pagination) (*utils.Pagination, error) {
	var categories []entity.Category
	if err := r.db.Scope(pagination.Paginate(categories, r.db)).Find(&categories).Error; err != nil {
		return nil, err
	}

	pagination.Rows = categories
	return &pagination, nil
}
