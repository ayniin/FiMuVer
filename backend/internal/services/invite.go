package services

import (
	"crypto/rand"
	"errors"
	"fimuver/internal/db"
	"fimuver/internal/models"
	"math/big"
	"time"

	"gorm.io/gorm"
)

// ValidatedInvite wraps a validated InviteCode for use in handlers
type ValidatedInvite struct {
	Invite *models.InviteCode
}

const codeAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const codeLength = 8

type InviteService struct {
	db *db.Database
}

func NewInviteService(db *db.Database) *InviteService {
	return &InviteService{db: db}
}

func GenerateCode() (string, error) {
	result := make([]byte, codeLength)
	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(codeAlphabet))))
		if err != nil {
			return "", err
		}
		result[i] = codeAlphabet[n.Int64()]
	}
	return string(result), nil
}

func (s *InviteService) CreateInviteCode(userID uint, maxUses int, expiresAt *time.Time) (*models.InviteCode, error) {
	code, err := GenerateCode()
	if err != nil {
		return nil, err
	}

	invite := &models.InviteCode{
		Code:            code,
		CreatedByUserID: userID,
		MaxUses:         maxUses,
		CurrentUses:     0,
		ExpiresAt:       expiresAt,
	}

	if err := s.db.DB.Create(invite).Error; err != nil {
		return nil, err
	}
	return invite, nil
}

func (s *InviteService) GetInviteCode(code string) (*models.InviteCode, error) {
	var invite models.InviteCode
	if err := s.db.DB.Where("code = ?", code).First(&invite).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invite code nicht gefunden")
		}
		return nil, err
	}
	return &invite, nil
}

func (s *InviteService) ValidateInviteCode(code string) (*models.InviteCode, error) {
	invite, err := s.GetInviteCode(code)
	if err != nil {
		return nil, err
	}

	if invite.ExpiresAt != nil && time.Now().After(*invite.ExpiresAt) {
		return nil, errors.New("invite code ist abgelaufen")
	}

	if invite.MaxUses > 0 && invite.CurrentUses >= invite.MaxUses {
		return nil, errors.New("invite code hat maximale Verwendungen erreicht")
	}

	return invite, nil
}

func (s *InviteService) UseInviteCode(invite *models.InviteCode, usedByUserID uint) error {
	now := time.Now()
	return s.db.DB.Model(invite).Updates(map[string]interface{}{
		"current_uses":     invite.CurrentUses + 1,
		"used_by_user_id":  usedByUserID,
		"used_at":          now,
	}).Error
}

func (s *InviteService) GetInviteCodesByUser(userID uint) ([]models.InviteCode, error) {
	var codes []models.InviteCode
	if err := s.db.DB.Where("created_by_user_id = ?", userID).Order("created_at desc").Find(&codes).Error; err != nil {
		return nil, err
	}
	return codes, nil
}

func (s *InviteService) DeleteInviteCode(id uint, userID uint) error {
	result := s.db.DB.Where("id = ? AND created_by_user_id = ?", id, userID).Delete(&models.InviteCode{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("invite code nicht gefunden oder keine Berechtigung")
	}
	return nil
}
