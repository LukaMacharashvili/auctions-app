package services

import (
	"context"
	"net/http"

	"github.com/LukaMacharashvili/admin/pkg/helpers"
	"github.com/LukaMacharashvili/admin/pkg/models"
	"github.com/LukaMacharashvili/admin/pkg/pb"
	"github.com/LukaMacharashvili/admin/pkg/repositories"
)

type Server struct {
	AdminRepo *repositories.AdminRepository
}

func (s *Server) Load(req *pb.LoadRequest, stream pb.AdminService_LoadServer) error {
	rows, err := s.AdminRepo.Load()

	if err != nil {
		return err
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Roles, &user.Balance)

		if err != nil {
			return err
		}

		stream.Send(&pb.LoadResponse{User: &pb.User{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
			Role:     user.Roles,
			Balance:  user.Balance,
		}})
	}

	return nil
}

func (s *Server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	var user models.User
	rows, err := s.AdminRepo.Fetch(req.GetUser().GetEmail())

	if err != nil {
		return &pb.CreateResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, err
	}

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Roles, &user.Balance)

		if err == nil {
			if user.Email != "" {
				return &pb.CreateResponse{
					Status: http.StatusConflict,
					Error:  "Email Already Exists",
				}, nil
			}
		}
	}

	user.Username = req.GetUser().GetUsername()
	user.Email = req.GetUser().GetEmail()
	user.Password = helpers.HashPassword(req.GetUser().GetPassword())
	user.Roles = req.GetUser().GetRole()
	user.Balance = req.GetUser().GetBalance()

	created, err := s.AdminRepo.Create(user)

	if !created {
		return nil, err
	}

	return &pb.CreateResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	deleted, err := s.AdminRepo.Delete(req.GetId())

	if !deleted {
		return nil, err
	}

	return &pb.DeleteResponse{
		Deleted: deleted,
	}, nil
}
