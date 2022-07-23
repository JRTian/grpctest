Task 1. Add another filter to the existing RPC
implementation:
    (1) added "optional bool visibleOnly = 2";
    (2) added main_test.go as unit test
    (3) use the below to test
        curl -X "POST" "http://localhost:8000/v1/list-races" \
            -H 'Content-Type: application/json' \
            -d $'{
        "filter": {"meeting_ids":["6"], "visible_only":true}
        }'


Task 2. add order by
    (1) added optional order by option, can order by any column name and choose ASC or DESC       
        message ListRacesRequest {
            ListRacesRequestFilter filter = 1;
            optional ListRacesRequestOrderBy order_by = 2;
        }
        //message for doing ordering by
        message ListRacesRequestOrderBy {
            string order_by = 3; // column name in race
            string order = 4;    //asc or desc
        }

    (2) modified unit test
    (3)column name and order seq option can be mixed with lower cases and upper cases letters
    (4) invalid ordering column will be ignored
    (5) use to below to test
    curl -X "POST" "http://localhost:8000/v1/list-races" \
        -H 'Content-Type: application/json' \
        -d $'{
    "filter": {"meeting_ids":["6"], "visible_only":true},
    "order_by":{"order_by":"advertised_start_time", "order":"ASC"}
    }'

Task 3. add status
    (1) wrap the original Race type with RaceWithStatus
    (2) this soluton does not bother the races.go which is the db model
        the thought is status is as extra column which should be seperated from 
        the db model struct
    (3) new return is like
    {"id":"418","meetingId":"418","name":"Massachusetts black cats","number":"1","visible":false,"advertisedStartTime":"2022-07-22T20:17:43Z","status":"CLOSED"},
    {"id":"419","meetingId":"419","name":"Arkansas elves","number":"5","visible":false,"advertisedStartTime":"2022-07-25T05:36:55Z","status":"OPEN"}
