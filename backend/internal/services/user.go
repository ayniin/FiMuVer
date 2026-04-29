package services

import (
	"errors"
	"fimuver/internal/auth"
	"fimuver/internal/db"
	"fimuver/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *db.Database
}

func NewUserService(db *db.Database) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if err := s.db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) GetUserByName(name string) (models.User, error) {
	var user models.User
	if err := s.db.DB.Where("username = ?", name).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) GetUserByID(id int) (models.User, error) {
	var user models.User
	if err := s.db.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	_, err := s.GetUserByEmail(user.Email)

	if err == nil {
		return models.User{}, errors.New("email already exists")
	}

	// Prüfen ob Fehler "not found" ist (GORM-spezifisch)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// echter Fehler (z.B. DB down)
		return models.User{}, err
	}

	_, err = s.GetUserByName(user.Username)
	if err == nil {
		return models.User{}, errors.New("username already exists")
	}

	// Prüfen ob Fehler "not found" ist (GORM-spezifisch)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// echter Fehler (z.B. DB down)
		return models.User{}, err
	}

	// Hash password before storing
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	user.Password = string(hashed)

	// User existiert nicht → erstellen
	if err := s.db.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) LoginUser(email string, password string) (models.User, string, error) {
	// Try to find user by email
	existingUser, err := s.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, "", errors.New("combination not valid")
		}
		return models.User{}, "", err
	}

	// check password hash
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password)); err != nil {
		return models.User{}, "", errors.New("combination not valid")
	}

	// generate JWT token
	token, err := auth.GenerateToken(existingUser.ID)
	if err != nil {
		return models.User{}, "", err
	}

	return existingUser, token, nil
}
