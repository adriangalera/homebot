sudo apt-get install jq -y
curl https://api.telegram.org/bot$1/getUpdates | jq -r '.'