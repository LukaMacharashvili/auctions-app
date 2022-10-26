package services

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/LukaMacharashvili/item/pkg/models"
	"github.com/LukaMacharashvili/item/pkg/pb"
	"github.com/LukaMacharashvili/item/pkg/repositories"
)

type Server struct {
	ItemRepo *repositories.ItemRepository
	UserRepo *repositories.UserRepository
}

func (s *Server) Load(req *pb.LoadRequest, stream pb.ItemService_LoadServer) error {
	rows, err := s.ItemRepo.Load()

	if err != nil {
		return err
	}

	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.StartingPrice, &item.Image, &item.Status, &item.HighestBid, &item.OwnerId, &item.HighestBidder)

		if err != nil {
			return err
		}

		stream.Send(&pb.LoadResponse{Item: &pb.Item{
			Id:            item.Id,
			Name:          item.Name,
			Description:   item.Description,
			StartingPrice: item.StartingPrice,
			Image:         item.Image,
			Status:        item.Status,
			HighestBid:    item.HighestBid,
			OwnerId:       item.OwnerId,
			HighestBidder: item.HighestBidder,
		}})
	}

	return nil
}

func (s *Server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	var item models.Item

	item.Name = req.GetItem().GetName()
	item.Description = req.GetItem().GetDescription()
	item.StartingPrice = req.GetItem().GetStartingPrice()
	item.Image = req.GetItem().GetImage()
	item.Status = req.GetItem().GetStatus()
	item.HighestBid = req.GetItem().GetHighestBid()
	item.OwnerId = req.GetItem().GetOwnerId()
	item.HighestBidder = req.GetItem().GetHighestBidder()

	created, err := s.ItemRepo.Create(item)

	if !created {
		return nil, err
	}

	return &pb.CreateResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	deleted, err := s.ItemRepo.Delete(req.GetId())

	if !deleted {
		return nil, err
	}

	return &pb.DeleteResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) Fetch(ctx context.Context, req *pb.FetchRequest) (*pb.FetchResponse, error) {
	var item models.Item
	rows, err := s.ItemRepo.Fetch(req.GetId())

	if err != nil {
		return &pb.FetchResponse{
			Error: err.Error(),
		}, err
	}

	for rows.Next() {
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.StartingPrice, &item.Image, &item.Status, &item.HighestBid, &item.OwnerId, &item.HighestBidder)
		if err != nil {
			return &pb.FetchResponse{
				Error: err.Error(),
			}, err
		}
	}

	return &pb.FetchResponse{
		Item: &pb.Item{
			Id:            item.Id,
			Name:          item.Name,
			Description:   item.Description,
			StartingPrice: item.StartingPrice,
			Image:         item.Image,
			Status:        item.Status,
			HighestBid:    item.HighestBid,
			OwnerId:       item.OwnerId,
			HighestBidder: item.HighestBidder,
		},
	}, nil
}

func (s *Server) Bid(ctx context.Context, req *pb.BidRequest) (*pb.BidResponse, error) {
	var item models.Item
	var user models.User

	fetchChannel := make(chan error)
	var rows1 *sql.Rows
	var rows2 *sql.Rows
	var err error

	go func() {
		rows1, err = s.ItemRepo.Fetch(req.GetId())

		fetchChannel <- err
	}()
	go func() {
		rows2, err = s.UserRepo.Fetch(req.GetUserId())

		fetchChannel <- err
	}()

	for i := 0; i < 2; i++ {
		fetchErr := <-fetchChannel
		if fetchErr != nil {
			return &pb.BidResponse{
				Error: err.Error(),
			}, err
		}
	}

	for rows1.Next() {
		err := rows1.Scan(&item.Id, &item.Name, &item.Description, &item.StartingPrice, &item.Image, &item.Status, &item.HighestBid, &item.OwnerId, &item.HighestBidder)
		if err != nil {
			return &pb.BidResponse{
				Error: err.Error(),
			}, err
		}
	}

	for rows2.Next() {
		err := rows2.Scan(&user.Id, &user.Balance)

		if err != nil {
			return &pb.BidResponse{
				Error: err.Error(),
			}, err
		}
	}

	if req.GetAmount() > item.StartingPrice && req.GetAmount() > item.HighestBid && user.Balance >= req.GetAmount() {
		item.HighestBid = req.GetAmount()
		item.HighestBidder = req.GetUserId()

		_, err := s.ItemRepo.UpdateHighestBid(req.GetId(), user.Id, req.GetAmount())

		if err != nil {
			return &pb.BidResponse{
				Error: err.Error(),
			}, err
		} else {
			return &pb.BidResponse{
				Status: http.StatusOK,
			}, nil
		}
	} else {
		return &pb.BidResponse{
			Error: "amount not enaugh",
		}, errors.New("amount not enaugh")
	}

}

func (s *Server) FinishAuction(ctx context.Context, req *pb.FinishAuctionRequest) (*pb.FinishAuctionResponse, error) {
	var item models.Item
	var user models.User

	fetchChannel := make(chan error)
	var rows1 *sql.Rows
	var rows2 *sql.Rows
	var err error

	go func() {
		rows1, err = s.ItemRepo.Fetch(req.GetId())

		fetchChannel <- err
	}()
	go func() {
		rows2, err = s.UserRepo.Fetch(req.GetOwnerId())

		fetchChannel <- err
	}()

	for i := 0; i < 2; i++ {
		fetchErr := <-fetchChannel
		if fetchErr != nil {
			return &pb.FinishAuctionResponse{
				Error: err.Error(),
			}, err
		}
	}

	for rows1.Next() {
		err := rows1.Scan(&item.Id, &item.Name, &item.Description, &item.StartingPrice, &item.Image, &item.Status, &item.HighestBid, &item.OwnerId, &item.HighestBidder)
		if err != nil {
			return &pb.FinishAuctionResponse{
				Error: err.Error(),
			}, err
		}
	}

	for rows2.Next() {
		err := rows2.Scan(&user.Id, &user.Balance)

		if err != nil {
			return &pb.FinishAuctionResponse{
				Error: err.Error(),
			}, err
		}
	}

	updateChannel := make(chan error)
	if user.Balance >= item.HighestBid {
		go func() {
			_, err := s.ItemRepo.Delete(req.GetId())

			updateChannel <- err
		}()
		go func() {
			_, err := s.UserRepo.DecreaseBalance(item.HighestBidder, item.HighestBid)

			updateChannel <- err
		}()

		for i := 0; i < 2; i++ {
			updateError := <-updateChannel
			if updateError != nil {
				return &pb.FinishAuctionResponse{
					Error: updateError.Error(),
				}, updateError
			}
		}
	}

	return &pb.FinishAuctionResponse{
		Status: http.StatusOK,
	}, nil
}
