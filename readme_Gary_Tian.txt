1. Add another filter to the existing RPC
implementation:
    (1) added "optional bool visibleOnly = 2";
    (2) added main_test.go as unit test
    (3) use the below to test
        curl -X "POST" "http://localhost:8000/v1/list-races" \
            -H 'Content-Type: application/json' \
            -d $'{
        "filter": {"meeting_ids":["6"], "visibleOnly":true}
        }'