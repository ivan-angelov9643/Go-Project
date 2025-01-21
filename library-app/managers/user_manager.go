package managers

import (
	"awesomeProject/library-app/db"
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

func (m *UserManager) GetAll(scopes ...db.DBScope) ([]models.User, error) {
	log.Info("[UserManager.GetAll] Fetching all users")

	var allUsers []models.User
	err := db.ApplyScopes(m.db, scopes).Find(&allUsers).Error
	if err != nil {
		log.Errorf("[UserManager.GetAll] Error fetching all users: %v", err)
		return nil, db.NewDBError(db.InternalError, "[UserManager.GetAll] Error fetching all users: %v", err)
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
		return models.User{}, db.NewDBError(db.InternalError, "[UserManager.Get] Error fetching user with ID %s: %v", idToGet, err)
	}

	log.Infof("[UserManager.Get] Successfully fetched user with ID: %s", idToGet)
	return user, nil
}

func (m *UserManager) Create(newUser models.User) (models.User, error) {
	log.Infof("[UserManager.Create] Creating new user")

	err := newUser.Validate()
	if err != nil {
		return models.User{}, db.NewDBError(db.ValidationError, err.Error())
	}

	err = m.db.Create(&newUser).Error
	if err != nil {
		log.Errorf("[UserManager.Create] Error creating new user with ID %s: %v", newUser.ID, err)
		return models.User{}, db.NewDBError(db.InternalError, "[UserManager.Create] Error creating new user with ID %s: %v", newUser.ID, err)
	}

	log.Infof("[UserManager.Create] Successfully created user with ID: %s", newUser.ID)
	return newUser, nil
}

func (m *UserManager) Update(updatedUser models.User) (models.User, error) {
	log.Infof("[UserManager.Update] Updating user with ID: %s", updatedUser.ID)

	err := updatedUser.Validate()
	if err != nil {
		return models.User{}, db.NewDBError(db.ValidationError, err.Error())
	}

	var user models.User
	err = m.db.First(&user, "id = ?", updatedUser.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[UserManager.Update] User with ID %s does not exist", updatedUser.ID)
			return models.User{}, db.NewDBError(db.NotFoundError, "[UserManager.Update] User with ID %s does not exist", updatedUser.ID)
		}
		log.Errorf("[UserManager.Update] Error fetching user with ID %s: %v", updatedUser.ID, err)
		return models.User{}, db.NewDBError(db.InternalError, "[UserManager.Update] Error fetching user with ID %s: %v", updatedUser.ID, err)
	}

	err = m.db.Model(&user).Updates(updatedUser).Error
	if err != nil {
		log.Errorf("[UserManager.Update] Error updating user with ID %s: %v", updatedUser.ID, err)
		return models.User{}, db.NewDBError(db.InternalError, "[UserManager.Update] Error updating user with ID %s: %v", updatedUser.ID, err)
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
			return models.User{}, db.NewDBError(db.NotFoundError, "[UserManager.Delete] User with ID %s does not exist", idToDelete)
		}
		log.Errorf("[UserManager.Delete] Error fetching user with ID %s: %v", idToDelete, err)
		return models.User{}, db.NewDBError(db.InternalError, "[UserManager.Delete] Error fetching user with ID %s: %v", idToDelete, err)
	}

	err = m.db.Delete(&user).Error
	if err != nil {
		log.Errorf("[UserManager.Delete] Error deleting user with ID %s: %v", idToDelete, err)
		return models.User{}, db.NewDBError(db.InternalError, "[UserManager.Delete] Error deleting user with ID %s: %v", idToDelete, err)
	}

	log.Infof("[UserManager.Delete] Successfully deleted user with ID: %s", idToDelete)
	return user, nil
}

func (m *UserManager) Count(scopes ...db.DBScope) (int64, error) {
	log.Infof("[UserManager.Count] Counting users in the database")

	var count int64
	err := db.ApplyScopes(m.db, scopes).Model(&models.User{}).Count(&count).Error
	if err != nil {
		log.Errorf("[UserManager.Count] Error counting users: %v", err)
		return 0, db.NewDBError(db.InternalError, "[UserManager.Count] Error counting users: %v", err)
	}

	log.Infof("[UserManager.Count] Successfully counted users: %d", count)
	return count, nil
}
