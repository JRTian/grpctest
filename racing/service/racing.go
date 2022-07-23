package service

import (
	"time"
	// "log"
	"git.neds.sh/matty/entain/racing/db"
	"git.neds.sh/matty/entain/racing/proto/racing"
	"golang.org/x/net/context"
)

type Racing interface {
	// ListRaces will return a collection of races.
	ListRaces(ctx context.Context, in *racing.ListRacesRequest) (*racing.ListRacesResponse, error)
	GetRace(ctx context.Context, in *racing.GetRaceRequest) (*racing.GetRaceResponse, error)
}

// racingService implements the Racing interface.
type racingService struct {
	racesRepo db.RacesRepo
}

// NewRacingService instantiates and returns a new racingService.
func NewRacingService(racesRepo db.RacesRepo) Racing {
	return &racingService{racesRepo}
}

func (s *racingService) ListRaces(ctx context.Context, in *racing.ListRacesRequest) (*racing.ListRacesResponse, error) {
	races, err := s.racesRepo.List(in.Filter, in.OrderBy)
	if err != nil {
		return nil, err
	}

	// wrap the races to races with status
	racesWithStatus := s.convertRacesToRacesWithStatus(races)
	return &racing.ListRacesResponse{Races: racesWithStatus}, nil
}

func (s *racingService)GetRace(ctx context.Context, in *racing.GetRaceRequest) (*racing.GetRaceResponse, error){
	race, err := s.racesRepo.GetRace(in)
	if err != nil || race == nil{
		return nil, err
	}

	// wrap the races to races with status
	racesWithStatus := s.convertRaceItemToHaveStatus(race)
	return &racing.GetRaceResponse{Race: racesWithStatus}, nil
}

/* to wrap the race with status, private function to ensure the encapsulation */
func (s *racingService) convertRacesToRacesWithStatus(races []*racing.Race) ([]*racing.RaceWithStatus) {
	// copy the input slice
	var racesWithStatus []*racing.RaceWithStatus
    for i:=0; i<len(races); i++{
        race := races[i]
		racesWithStatuItem := s.convertRaceItemToHaveStatus(race)
		racesWithStatus = append(racesWithStatus, racesWithStatuItem)
        // log.Println("racesWithStatuItem ", racesWithStatuItem, nowInSeconds)
    } 

	return racesWithStatus
}

func (s *racingService) convertRaceItemToHaveStatus(race *racing.Race) (*racing.RaceWithStatus) {
	raceWithStatuItem := &(racing.RaceWithStatus{
		Id: race.GetId(),
		MeetingId: race.GetMeetingId(),
		Name:race.GetName(),
		Number:race.GetNumber(),
		Visible:race.GetVisible(),
		AdvertisedStartTime:race.GetAdvertisedStartTime(),
		Status: "",
	})
     
	nowInSeconds := time.Now().Unix()
	if raceWithStatuItem.GetAdvertisedStartTime().GetSeconds() < nowInSeconds{
		raceWithStatuItem.Status = "CLOSED"
	}else{
		raceWithStatuItem.Status = "OPEN"
	}
	return raceWithStatuItem
	// log.Println("racesWithStatuItem ", racesWithStatuItem, nowInSeconds)
}