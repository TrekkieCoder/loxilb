curl -X 'POST' \
  'http://172.17.0.3:11111/netlox/v1/config/cistate' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "instance": "default",
  "state": "BACKUP",
  "vip": "0.0.0.0"
}'

curl -X 'POST' \
  'http://172.17.0.4:11111/netlox/v1/config/cistate' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "instance": "default",
  "state": "BACKUP",
  "vip": "0.0.0.0"
}'

curl -X 'POST' \
  'http://172.17.0.5:11111/netlox/v1/config/cistate' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "instance": "default",
  "state": "MASTER",
  "vip": "0.0.0.0"
}'
