package repository

import (
	"context"

	"gorm.io/gorm"
)

type Repository[E any] interface {
	GetAll(context.Context) ([]E, error)
	GetByID(context.Context, int64) (E, error)
	Save(context.Context, *E) error
	Update(context.Context, *E) error
	Delete(context.Context, *E) error
}

type GormRepository[E any] struct {
	DB *gorm.DB
}

func (r *GormRepository[E]) GetAll(ctx context.Context) ([]E, error) {
	var entities []E

	if err := r.DB.WithContext(ctx).Find(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}

func (r *GormRepository[E]) GetByID(ctx context.Context, id int64) (E, error) {
	var entity E

	if err := r.DB.WithContext(ctx).First(&entity, "id = ?", id).Error; err != nil {
		return *new(E), err
	}

	return entity, nil
}

func (r *GormRepository[E]) Save(ctx context.Context, entity *E) error {
	if err := r.DB.WithContext(ctx).Create(&entity).Error; err != nil {
		return err
	}

	return nil
}

func (r *GormRepository[E]) Update(ctx context.Context, entity *E) error {
	if err := r.DB.WithContext(ctx).Updates(&entity).Error; err != nil {
		return err
	}

	return nil
}

func (r *GormRepository[E]) Delete(ctx context.Context, entity *E) error {
	if err := r.DB.WithContext(ctx).Delete(&entity).Error; err != nil {
		return err
	}

	return nil
}
