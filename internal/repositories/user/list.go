package user

import "poc-testcontainers/internal/models"

func (u *userRepository) List(filter *models.User, page int) ([]models.User, error) {
	result := []models.User{}
	err := u.db.
		Where(filter).
		Limit(10).
		Offset(10 * page).
		Find(&result).
		Error

	if err != nil {
		return nil, err
	}

	return result, err
}
