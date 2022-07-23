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
			t.Errorf("expected true but 3 %s ", res.Races[i])
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
	*f.VisibleOnly = true
	var a[]int64
	a = append(a, 3)
    f.MeetingIds = a  

	orderBy := &(racing.ListRacesRequestOrderBy{OrderBy: "advertised_start_time", Order:"ASC"})
    
	var curentADStartTime *timestamppb.Timestamp
	res, _ := r.ListRaces(context.Background(), &racing.ListRacesRequest{Filter:f,OrderBy : orderBy})
	for i:=0; i<len(res.Races); i++{
		log.Println(res.Races[i])
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