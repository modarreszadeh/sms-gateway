#!/bin/bash

url="http://localhost:5000/v1/api/sms/send"

data='{
    "sender": "09393639116",
    "receptor": "09121111111",
    "message": "test"
}'

concurrency=1000

total_requests=1000000

tmpfile=$(mktemp)

echo "$data" > $tmpfile

ab -n $total_requests -c $concurrency -H 'userId: 65b909e462cafc46de1202c9' -T "application/json" -p $tmpfile -k $url

rm $tmpfile
