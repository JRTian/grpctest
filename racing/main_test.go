package main

import (
	 
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