package postgres_gorm

import (
	"context"
	"errors"

	// "errors"
	"log"
	"project-riskprofile/entity"
	"project-riskprofile/service"

	"gorm.io/gorm"
)

// GormDBIface defines an interface for GORM DB methods used in the repository
type GormDBIface2 interface {
	WithContext(ctx context.Context) *gorm.DB
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}

type submissionRepository struct {
	db GormDBIface2
}

// NewUserRepository membuat instance baru dari userRepository
func NewsubmissionRepository(db GormDBIface) service.IsubmissionRepository {
	return &submissionRepository{db: db}
}

// CreateUser membuat pengguna baru dalam basis data
func (r *submissionRepository) CreateSubmi(ctx context.Context, subm *entity.Submission) (entity.Submission, error) {
	if err := r.db.WithContext(ctx).Create(subm).Error; err != nil {
		log.Printf("Error creating user: %v\n", err)
		return entity.Submission{}, err
	}
	return *subm, nil
}

// GetUserByID mengambil pengguna berdasarkan ID
func (r *submissionRepository) GetUSubmiByID(ctx context.Context, id int) (entity.Submission, error) {
	var user entity.Submission
	if err := r.db.WithContext(ctx).Select("id", "user_id", "answers", "risk_score", "risk_category", "created_at", "updated_at").Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Submission{}, nil
		}

		log.Printf("Error getting user by ID: %v\n", err)
		return entity.Submission{}, err
	}
	return user, nil
}

func (r *submissionRepository) GetUSubmiByUserID(ctx context.Context, id int) (entity.Submission, error) {
	var user entity.Submission
	if err := r.db.WithContext(ctx).Select("id", "user_id", "answers", "risk_score", "risk_category", "created_at", "updated_at").Where("user_id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Submission{}, nil
		}

		log.Printf("Error getting user by ID: %v\n", err)
		return entity.Submission{}, err
	}
	return user, nil
}

// // GetUserByEmail mengambil pengguna berdasarkan email
// func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
// 	var user entity.User
// 	if err := r.db.WithContext(ctx).Select("id", "name", "email", "password", "created_at", "updated_at").Where("email = ?", email).First(&user).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return entity.User{}, nil
// 		}
// 		log.Printf("Error getting user by ID: %v\n", err)
// 		return entity.User{}, err
// 	}
// 	return user, nil
// }

// UpdateUser memperbarui informasi pengguna dalam basis data
// func (r *submissionRepository) UpdateSubmi(ctx context.Context, id int, user entity.Submission) (entity.Submission, error) {
// 	// Menemukan pengguna yang akan diperbarui
// 	var existingUser entity.Submission
// 	if err := r.db.WithContext(ctx).Select("id", "user_id", "answers", "risk_score", "risk_category", "created_at", "updated_at").First(&existingUser, id).Error; err != nil {
// 		log.Printf("Error finding user to update: %v\n", err)
// 		return entity.Submission{}, err
// 	}

// 	// Memperbarui informasi pengguna
// 	existingUser.UserID = user.UserID
// 	existingUser.Answers = user.Answers
// 	existingUser.RiskScore = user.RiskScore
// 	existingUser.RiskCategory = user.RiskCategory

// 	if err := r.db.WithContext(ctx).Save(&existingUser).Error; err != nil {
// 		log.Printf("Error updating user: %v\n", err)
// 		return entity.Submission{}, err
// 	}
// 	return existingUser, nil
// }

// 	// Memperbarui informasi pengguna
// 	existingUser.Name = user.Name
// 	existingUser.Email = user.Email
// 	if err := r.db.WithContext(ctx).Save(&existingUser).Error; err != nil {
// 		log.Printf("Error updating user: %v\n", err)
// 		return entity.User{}, err
// 	}
// 	return existingUser, nil
// }

// DeleteUser menghapus pengguna berdasarkan ID
func (r *submissionRepository) DeleteSubmi(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Submission{}, id).Error; err != nil {
		log.Printf("Error deleting user: %v\n", err)
		return err
	}
	return nil
}

// // GetAllUsers mengambil semua pengguna dari basis data
func (r *submissionRepository) GetAllSubmi(ctx context.Context) ([]entity.Submission, error) {
	var users []entity.Submission
	if err := r.db.WithContext(ctx).Select("id", "user_id", "answers", "risk_score", "risk_category", "created_at", "updated_at").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users, nil
		}
		log.Printf("Error getting all users: %v\n", err)
		return nil, err
	}
	return users, nil
}
