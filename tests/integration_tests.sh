#!/bin/bash
echo 'GET http://localhost:8080/v1/get-remaining-requests' | vegeta attack -rate 5 -duration 20s | vegeta encode > result.log
ok=`cat result.log | jq .code | grep 200 | wc -l`
too_many_reqs=`cat result.log | jq .code | grep 429 | wc -l`

if [ "$ok" -eq 60 ] && [ "$too_many_reqs" -eq 40 ]; then
    echo "Success"
else
    echo "Failed\tStatus Ok: $ok, Status Too Many Requests: $too_many_reqs"
fi
