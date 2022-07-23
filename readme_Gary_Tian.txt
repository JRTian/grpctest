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