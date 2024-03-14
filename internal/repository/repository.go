package repository

import "gorm.io/gorm"

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func (r *BaseRepository[T]) Insert(entity *T) error {
	return r.DB.Create(&entity).Error
}

func (r *BaseRepository[T]) FindAll() ([]T, error) {
	var entities []T
	err := r.DB.Find(&entities).Error

	return entities, err
}

func (r *BaseRepository[T]) Update(entity *T) error {
	return r.DB.Save(&entity).Error
}

func (r *BaseRepository[T]) Delete(entity *T) error {
	return r.DB.Delete(&entity).Error
}
