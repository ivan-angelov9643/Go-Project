package managers

import (
	"awesomeProject/library-app/global/db_error"
	"awesomeProject/library-app/models"
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserManager struct {
	db *gorm.DB
}

func NewUserManager(db *gorm.DB) *UserManager {
	log.Info("[NewUserManager] Initializing UserManager")
	return &UserManager{db}
}

func (m *UserManager) GetAll() ([]models.User, error) {
	log.Info("[UserManager.GetAll] Fetching all users")

	var allUsers []models.User
	err := m.db.Find(&allUsers).Error
	if err != nil {
		log.Errorf("[UserManager.GetAll] Error fetching all users: %v", err)
		return nil, db_error.NewDBError(db_error.InternalError, "[UserManager.GetAll] Error fetching all users: %v", err)
	}

	log.Infof("[UserManager.GetAll] Successfully fetched all users")
	return allUsers, nil
}

func (m *UserManager) Get(idToGet uuid.UUID) (models.User, error) {
	log.Infof("[UserManager.Get] Fetching user with ID: %s", idToGet)

	var user models.User
	err := m.db.First(&user, "id = ?", idToGet).Error
	if err != nil {
		log.Errorf("[UserManager.Get] Error fetching user with ID %s: %v", idToGet, err)
		return models.User{}, db_error.NewDBError(db_error.InternalError, "[UserManager.Get] Error fetching user with ID %s: %v", idToGet, err)
	}

	log.Infof("[UserManager.Get] Successfully fetched user with ID: %s", idToGet)
	return user, nil
}

func (m *UserManager) Create(newUser models.User) (models.User, error) {
	log.Infof("[UserManager.Create] Creating new user")

	err := newUser.Validate()
	if err != nil {
		return models.User{}, db_error.NewDBError(db_error.ValidationError, err.Error())
	}

	err = m.db.Create(&newUser).Error
	if err != nil {
		log.Errorf("[UserManager.Create] Error creating new user with ID %s: %v", newUser.ID, err)
		return models.User{}, db_error.NewDBError(db_error.InternalError, "[UserManager.Create] Error creating new user with ID %s: %v", newUser.ID, err)
	}

	log.Infof("[UserManager.Create] Successfully created user with ID: %s", newUser.ID)
	return newUser, nil
}

func (m *UserManager) Update(updatedUser models.User) (models.User, error) {
	log.Infof("[UserManager.Update] Updating user with ID: %s", updatedUser.ID)

	err := updatedUser.Validate()
	if err != nil {
		return models.User{}, db_error.NewDBError(db_error.ValidationError, err.Error())
	}

	var user models.User
	err = m.db.First(&user, "id = ?", updatedUser.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[UserManager.Update] User with ID %s does not exist", updatedUser.ID)
			return models.User{}, db_error.NewDBError(db_error.NotFoundError, "[UserManager.Update] User with ID %s does not exist", updatedUser.ID)
		}
		log.Errorf("[UserManager.Update] Error fetching user with ID %s: %v", updatedUser.ID, err)
		return models.User{}, db_error.NewDBError(db_error.InternalError, "[UserManager.Update] Error fetching user with ID %s: %v", updatedUser.ID, err)
	}

	err = m.db.Model(&user).Updates(updatedUser).Error
	if err != nil {
		log.Errorf("[UserManager.Update] Error updating user with ID %s: %v", updatedUser.ID, err)
		return models.User{}, db_error.NewDBError(db_error.InternalError, "[UserManager.Update] Error updating user with ID %s: %v", updatedUser.ID, err)
	}

	log.Infof("[UserManager.Update] Successfully updated user with ID: %s", updatedUser.ID)
	return updatedUser, nil
}

func (m *UserManager) Delete(idToDelete uuid.UUID) (models.User, error) {
	log.Infof("[UserManager.Delete] Deleting user with ID: %s", idToDelete)

	var user models.User
	err := m.db.First(&user, "id = ?", idToDelete).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[UserManager.Delete] User with ID %s does not exist", idToDelete)
			return models.User{}, db_error.NewDBError(db_error.NotFoundError, "[UserManager.Delete] User with ID %s does not exist", idToDelete)
		}
		log.Errorf("[UserManager.Delete] Error fetching user with ID %s: %v", idToDelete, err)
		return models.User{}, db_error.NewDBError(db_error.InternalError, "[UserManager.Delete] Error fetching user with ID %s: %v", idToDelete, err)
	}

	err = m.db.Delete(&user).Error
	if err != nil {
		log.Errorf("[UserManager.Delete] Error deleting user with ID %s: %v", idToDelete, err)
		return models.User{}, db_error.NewDBError(db_error.InternalError, "[UserManager.Delete] Error deleting user with ID %s: %v", idToDelete, err)
	}

	log.Infof("[UserManager.Delete] Successfully deleted user with ID: %s", idToDelete)
	return user, nil
}
