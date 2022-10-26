package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/LukaMacharashvili/wallet/pkg/models"
	"github.com/LukaMacharashvili/wallet/pkg/pb"
	"github.com/LukaMacharashvili/wallet/pkg/repositories"
)

type Server struct {
	RaiseReqRepo    *repositories.RaiseReqRepository
	TransactionRepo *repositories.TransactionRepository
	UserRepo        *repositories.UserRepository
}

func (s *Server) LoadTransactions(req *pb.LoadTransactionsRequest, stream pb.WalletService_LoadTransactionsServer) error {
	rows, err := s.TransactionRepo.LoadTransaction(req.GetOwnerId())

	if err != nil {
		return err
	}

	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.Id, &transaction.OwnerId, &transaction.Amount)

		if err != nil {
			return err
		}

		stream.Send(&pb.LoadTransactionsResponse{Transaction: &pb.Transaction{
			Id:      transaction.Id,
			OwnerId: transaction.OwnerId,
			Amount:  transaction.Amount,
		}})
	}

	return nil
}

func (s *Server) LoadRaiseReqs(req *pb.LoadRaiseReqsRequest, stream pb.WalletService_LoadRaiseReqsServer) error {
	rows, err := s.RaiseReqRepo.LoadRaiseReq(req.GetOwnerId())

	if err != nil {
		return err
	}

	for rows.Next() {
		var raiseReq models.RaiseReq
		err := rows.Scan(&raiseReq.Id, &raiseReq.OwnerId, &raiseReq.Amount, &raiseReq.Status)

		if err != nil {
			return err
		}

		stream.Send(&pb.LoadRaiseReqsResponse{RaiseReq: &pb.RaiseReq{
			Id:      raiseReq.Id,
			OwnerId: raiseReq.OwnerId,
			Amount:  raiseReq.Amount,
			Status:  raiseReq.Status,
		}})
	}

	return nil
}

func (s *Server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	var raiseReq models.RaiseReq

	raiseReq.OwnerId = req.GetRaiseReq().GetOwnerId()
	raiseReq.Amount = req.GetRaiseReq().GetAmount()
	raiseReq.Status = req.GetRaiseReq().GetStatus()

	created, err := s.RaiseReqRepo.Create(raiseReq)

	if !created {
		return nil, err
	}

	return &pb.CreateResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) ChangeStatus(ctx context.Context, req *pb.ChangeStatusRequest) (*pb.ChangeStatusResponse, error) {
	var raiseReq models.RaiseReq
	rows, err := s.RaiseReqRepo.Fetch(req.GetId())

	if err != nil {
		return &pb.ChangeStatusResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, err
	}

	for rows.Next() {
		err := rows.Scan(&raiseReq.Id, &raiseReq.OwnerId, &raiseReq.Amount, &raiseReq.Status)

		if err != nil {
			return &pb.ChangeStatusResponse{
				Status: http.StatusConflict,
				Error:  err.Error(),
			}, err
		}
	}
	repoChannel := make(chan bool)

	var transaction models.Transaction
	if raiseReq.Status != "Approved" && req.GetStatus() == "Approved" {
		transaction.Amount = raiseReq.Amount
		transaction.OwnerId = raiseReq.OwnerId

		go func() {
			resp, _ := s.UserRepo.IncreaseBalance(transaction.OwnerId, transaction.Amount)
			repoChannel <- resp
		}()
		go func() {
			resp, _ := s.RaiseReqRepo.UpdateStatus(req.GetId(), req.GetStatus())
			repoChannel <- resp
		}()
		go func() {
			resp, _ := s.TransactionRepo.Create(transaction)
			repoChannel <- resp
		}()
		for i := 0; i < 3; i++ {
			respBool := <-repoChannel
			if !respBool {
				return &pb.ChangeStatusResponse{
					Status: http.StatusInternalServerError,
				}, errors.New("something went wrong")
			}
		}
	} else if raiseReq.Status == "Approved" && req.GetStatus() == "Rejected" {
		transaction.Amount = -raiseReq.Amount
		transaction.OwnerId = raiseReq.OwnerId

		go func() {
			resp, _ := s.UserRepo.DecreaseBalance(transaction.OwnerId, -transaction.Amount)
			repoChannel <- resp
		}()
		go func() {
			resp, _ := s.RaiseReqRepo.UpdateStatus(req.GetId(), req.GetStatus())
			repoChannel <- resp
		}()
		go func() {
			resp, _ := s.TransactionRepo.Create(transaction)
			repoChannel <- resp
		}()
		for i := 0; i < 3; i++ {
			respBool := <-repoChannel
			if !respBool {
				return &pb.ChangeStatusResponse{
					Status: http.StatusInternalServerError,
				}, errors.New("something went wrong")
			}
		}
	} else {
		_, err = s.RaiseReqRepo.UpdateStatus(req.GetId(), req.GetStatus())
	}

	if err != nil {
		return nil, err
	}

	return &pb.ChangeStatusResponse{
		Status: http.StatusOK,
	}, nil
}
