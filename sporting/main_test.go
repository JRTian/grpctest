package main

import (
	// "log"
	 
	"git.neds.sh/matty/entain/sporting/proto/sporting" 
	"google.golang.org/grpc"
 
	"context"
	"testing"
)


func TestNameFilter(t *testing.T) {
	var cConn *grpc.ClientConn
	cConn, _ = grpc.Dial(":10000", grpc.WithInsecure())
	 

	r := sporting.NewSportingClient(cConn)
	name := "Connecticut cats"
	f := &(sporting.ListSportsRequestFilter{})
	var a[]string
	a = append(a, name)
    f.Names = a 
	 

	res, err := r.ListEvents(context.Background(), &sporting.ListEventsRequest{Filter:f})
	if err != nil{
		t.Errorf("expected nil %s ", err)
		return
	}
	for i:=0; i<len(res.Sports); i++{
		mid := res.Sports[i].GetName()
		if mid != name{
			t.Errorf("expected 3 but %s ", res.Sports[i])
		}
	}
}

