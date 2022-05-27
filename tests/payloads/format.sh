# Uses JQ to format the payloads

for payload in */**.json
do
    if jq . $payload > $payload.tmp
    then
        cat $payload.tmp > $payload
        rm $payload.tmp
    fi
done