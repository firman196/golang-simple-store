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
	err := r.db.Create(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

// Update a category
func (r *CategoryRepositoryImpl) Update(category entity.Category) (*entity.Category, error) {
	err := r.db.Model(&category).Where("id = ?", category.Id).Updates(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// find category by id
func (r *CategoryRepositoryImpl) FindByID(id int) (*entity.Category, error) {
	var category entity.Category
	err := r.db.Model(&category).Where("id =?", id).First(&category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

// Find all categories
func (r *CategoryRepositoryImpl) FindAll(pagination utils.Pagination) (*utils.Pagination, error) {
	var categories []entity.Category
	if err := r.db.Scopes(pagination.Paginate(categories, r.db)).Find(&categories).Error; err != nil {
		return nil, err
	}

	pagination.Rows = categories
	return &pagination, nil
}
