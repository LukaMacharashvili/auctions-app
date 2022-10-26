package services

import (
	"context"
	"net/http"

	"github.com/LukaMacharashvili/auth/pkg/helpers"
	"github.com/LukaMacharashvili/auth/pkg/models"
	"github.com/LukaMacharashvili/auth/pkg/pb"
	"github.com/LukaMacharashvili/auth/pkg/repositories"
)

type Server struct {
	AuthRepo *repositories.AuthRepository
	Jwt      *helpers.JwtWrapper
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User
	rows, err := s.AuthRepo.Fetch(req.Email)

	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, err
	}

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Roles, &user.Balance)

		if err == nil {
			if user.Email != "" {
				return &pb.RegisterResponse{
					Status: http.StatusConflict,
					Error:  "Email Already Exists",
				}, nil
			}
		}
	}

	user.Username = req.Username
	user.Email = req.Email
	user.Password = helpers.HashPassword(req.Password)

	created, err := s.AuthRepo.Create(user)

	if !created {
		return nil, err
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	rows, err := s.AuthRepo.Fetch(req.Email)

	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, err
	}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Roles, &user.Balance)
	}

	match := helpers.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var user models.User

	rows, err := s.AuthRepo.Fetch(claims.Email)

	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Roles, &user.Balance)
	}

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
		Roles:  user.Roles,
	}, nil
}
