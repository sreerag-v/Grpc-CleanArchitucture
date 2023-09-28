package repository

import (
	"Clean-Grpc/pkg/domain"
	interfaces "Clean-Grpc/pkg/repository/interface"
	"context"

	"gorm.io/gorm"
)


type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}
func (c *userDatabase) FindAll(ctx context.Context) ([]domain.Users, error) {
	var users []domain.Users
	err := c.DB.Find(&users).Error

	return users, err
}

func (c *userDatabase) FindByID(ctx context.Context, id uint) (domain.Users, error) {
	var user domain.Users
	err := c.DB.First(&user, id).Error

	return user, err
}

func (c *userDatabase) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	err := c.DB.Save(&user).Error

	return user, err
}

func (c *userDatabase) Delete(ctx context.Context, user domain.Users) error {
	err := c.DB.Delete(&user).Error

	return err
}
