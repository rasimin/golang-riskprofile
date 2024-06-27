package service

import (
	"context"
	"fmt"
	"project-riskprofile/entity"
)

// IUserService mendefinisikan interface untuk layanan pengguna
type IsubmissionService interface {
	CreateSubmi(ctx context.Context, user *entity.Submission) (entity.Submission, error)
	GetUSubmiByID(ctx context.Context, id int) (entity.Submission, error)
	// 	GetUserByEmail(ctx context.Context, id string) (entity.User, error)
	// UpdateSubmi(ctx context.Context, id int, user entity.Submission) (entity.Submission, error)
	DeleteSubmi(ctx context.Context, id int) error
	GetAllSubmi(ctx context.Context) ([]entity.Submission, error)
}

// IUserRepository mendefinisikan interface untuk repository pengguna
type IsubmissionRepository interface {
	CreateSubmi(ctx context.Context, user *entity.Submission) (entity.Submission, error)
	GetUSubmiByID(ctx context.Context, id int) (entity.Submission, error)
	// GetUserByEmail(ctx context.Context, id string) (entity.User, error)
	// UpdateSubmi(ctx context.Context, id int, user entity.Submission) (entity.Submission, error)
	DeleteSubmi(ctx context.Context, id int) error
	GetAllSubmi(ctx context.Context) ([]entity.Submission, error)
}

// userService adalah implementasi dari IUserService yang menggunakan IUserRepository
type submissionService struct {
	submissionRepo IsubmissionRepository
}

// NewUserService membuat instance baru dari userService
func NewSubmissionService(submissionRepo IsubmissionRepository) IsubmissionService {
	return &submissionService{submissionRepo: submissionRepo}
}

// CreateUser membuat pengguna baru
func (s *submissionService) CreateSubmi(ctx context.Context, subm *entity.Submission) (entity.Submission, error) {
	// Memanggil CreateUser dari repository untuk membuat pengguna baru
	createdUser, err := s.submissionRepo.CreateSubmi(ctx, subm)
	if err != nil {
		return entity.Submission{}, fmt.Errorf("gagal membuat Submission: %v", err)
	}
	return createdUser, nil
}

// GetUserByID mendapatkan pengguna berdasarkan ID
func (s *submissionService) GetUSubmiByID(ctx context.Context, id int) (entity.Submission, error) {

	// Memanggil GetUserByID dari repository untuk mendapatkan pengguna berdasarkan ID
	user, err := s.submissionRepo.GetUSubmiByID(ctx, id)
	if err != nil {
		return entity.Submission{}, fmt.Errorf("gagal mendapatkan pengguna berdasarkan ID: %v", err)
	}
	return user, nil
}

// // GetUserByID mendapatkan pengguna berdasarkan ID
// func (s *userService) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
// 	// Memanggil GetUserByID dari repository untuk mendapatkan pengguna berdasarkan ID
// 	user, err := s.userRepo.GetUserByEmail(ctx, email)
// 	if err != nil {
// 		return entity.User{}, fmt.Errorf("gagal mendapatkan pengguna berdasarkan ID: %v", err)
// 	}
// 	return user, nil
// }

// UpdateUser memperbarui data pengguna
// func (s *submissionService) UpdateSubmi(ctx context.Context, id int, user entity.Submission) (entity.Submission, error) {
// 	// Memanggil UpdateUser dari repository untuk memperbarui data pengguna
// 	updatedUser, err := s.submissionRepo.UpdateSubmi(ctx, id, user)
// 	if err != nil {
// 		return entity.Submission{}, fmt.Errorf("gagal memperbarui pengguna: %v", err)
// 	}
// 	return updatedUser, nil
// }

// DeleteUser menghapus pengguna berdasarkan ID
func (s *submissionService) DeleteSubmi(ctx context.Context, id int) error {
	// Memanggil DeleteUser dari repository untuk menghapus pengguna berdasarkan ID
	err := s.submissionRepo.DeleteSubmi(ctx, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus pengguna: %v", err)
	}
	return nil
}

// GetAllUsers mendapatkan semua pengguna
func (s *submissionService) GetAllSubmi(ctx context.Context) ([]entity.Submission, error) {
	// Memanggil GetAllUsers dari repository untuk mendapatkan semua pengguna
	users, err := s.submissionRepo.GetAllSubmi(ctx)
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan semua pengguna: %v", err)
	}
	return users, nil
}
