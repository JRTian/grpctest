package service

import (
	// "log"
	"git.neds.sh/matty/entain/sporting/db"
	"git.neds.sh/matty/entain/sporting/proto/sporting"
	"golang.org/x/net/context"
)

type Sporting interface {
	// ListSports will return a collection of races.
	ListEvents(ctx context.Context, in *sporting.ListEventsRequest) (*sporting.ListEventsResponse, error)
}

// sportingService implements the Sporting interface.
type sportingService struct {
	sportsRepo db.SportsRepo
}

// NewRacingService instantiates and returns a new racingService.
func NewSportingService(sportsRepo db.SportsRepo) Sporting {
	return &sportingService{sportsRepo}
}

func (s *sportingService) ListEvents(ctx context.Context, in *sporting.ListEventsRequest) (*sporting.ListEventsResponse, error) {
	sports, err := s.sportsRepo.List(in.Filter)
	if err != nil {
		return nil, err
	}

	return &sporting.ListEventsResponse{Sports: sports}, nil
}

 

 