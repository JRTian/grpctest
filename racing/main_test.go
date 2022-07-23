package main

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	 
	"git.neds.sh/matty/entain/racing/proto/racing" 
	"google.golang.org/grpc"
 
	"context"
	"testing"
)


func TestMeetingIDFilter(t *testing.T) {
	var cConn *grpc.ClientConn
	cConn, _ = grpc.Dial(":9000", grpc.WithInsecure())
	 

	r := racing.NewRacingClient(cConn)
	f := &(racing.ListRacesRequestFilter{})
	var a[]int64
	a = append(a, 3)
    f.MeetingIds = a  

	res, err := r.ListRaces(context.Background(), &racing.ListRacesRequest{Filter:f,})
	if err != nil{
		t.Errorf("expected nil %s ", err)
		return
	}
	for i:=0; i<len(res.Races); i++{
		mid := res.Races[i].GetMeetingId()
		if mid != 3{
			t.Errorf("expected 3 but %s ", res.Races[i])
			return
		}
	}
}

func TestVisibilityFilter(t *testing.T) {
	var cConn *grpc.ClientConn
	cConn, _ = grpc.Dial(":9000", grpc.WithInsecure())
	 

	r := racing.NewRacingClient(cConn)
	f := &(racing.ListRacesRequestFilter{})
    f.VisibleOnly = new(bool)
	*f.VisibleOnly = true

	res, _ := r.ListRaces(context.Background(), &racing.ListRacesRequest{Filter:f,})
	for i:=0; i<len(res.Races); i++{
		visible := res.Races[i].GetVisible()
		if !visible{
			t.Errorf("expected true but fale %s ", res.Races[i])
		}
	}
	 
}

func TestOrderBy(t *testing.T) {
	var cConn *grpc.ClientConn
	cConn, _ = grpc.Dial(":9000", grpc.WithInsecure())
	 

	r := racing.NewRacingClient(cConn)
	f := &(racing.ListRacesRequestFilter{})
    f.VisibleOnly = new(bool)
	*f.VisibleOnly = false
	
	var a[]int64
	a = append(a, 3)
    f.MeetingIds = a  
    
	orderBy := &(racing.ListRacesRequestOrderBy{OrderBy: "advertised_start_time", Order:"ASC"})
    
	var curentADStartTime *timestamppb.Timestamp
	res, _ := r.ListRaces(context.Background(), &racing.ListRacesRequest{Filter:f,OrderBy : orderBy})
	for i:=0; i<len(res.Races); i++{
		// log.Println(res.Races[i])
		advertised_start_time := res.Races[i].GetAdvertisedStartTime()

		if i == 0{
			curentADStartTime = advertised_start_time
			continue
		}
		
		if advertised_start_time.GetSeconds() < curentADStartTime.GetSeconds(){
			t.Errorf("expected true but fale %s ", res.Races[i])
			return
		}
		curentADStartTime = advertised_start_time
	}
	
}

func TestGetRaceById(t *testing.T) {
	var cConn *grpc.ClientConn
	cConn, _ = grpc.Dial(":9000", grpc.WithInsecure())
	 

	r := racing.NewRacingClient(cConn)
	f := &(racing.GetRaceRequest{})
    f.Id = 1

	res, _ := r.GetRace(context.Background(), f)
	
	id := res.Race.GetId()
	if id != 1{
		t.Errorf("expected 1 but %s ", res)
	}
	log.Println("TestGetRaceById", res)
}