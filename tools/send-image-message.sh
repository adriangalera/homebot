PORT=$1
FILENAME=$2
curl -X POST http://localhost:"$PORT"/v1/image --data "{\"filename\":\"$FILENAME\"}"
