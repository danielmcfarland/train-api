HASH=$(git rev-parse --short HEAD)
GOOS=linux GOARCH=arm64 go build -v -o train-api-arm64 .

docker build -t train-api .
docker tag train-api registry.r-01.an0thr.com/mcc/train-api:latest
docker push registry.r-01.an0thr.com/f/train-api:latest
