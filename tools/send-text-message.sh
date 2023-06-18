PORT=$1
MESSAGE=$2
curl -X POST http://localhost:$PORT/v1/text -d "{\"text\" : \" ${MESSAGE}  \"}"
