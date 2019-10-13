# my-blog-api-sam

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)

## Setup process

Start Mysql container.

```shell
docker network create lambda-local
docker-compose up dynamodb
```

### Building

```shell
GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world
or
make build
```

### Local development

```bash
sam local start-api --env-vars env.json --docker-network lambda-local
```
Local test user is,

```
email: hoge@example.com
password: hogehoge
```

## Packaging and deployment

First and foremost, we need a `S3 bucket` where we can upload our Lambda functions packaged as ZIP before we deploy anything - If you don't have a S3 bucket to store code artifacts then this is a good time to create one:

```bash
aws s3 mb s3://BUCKET_NAME
```
And run `deploy.sh` . This script executes package and deploy.

```bash
deploy.sh
```

After deployment is complete you can run the following command to retrieve the API Gateway Endpoint URL:

```bash
aws cloudformation describe-stacks \
    --stack-name my-blog-api-sam \
    --query 'Stacks[].Outputs'
``` 

### Testing

First, check status of docker db.

```shell
docker-compose up -d dynamodb
```
And, execute all tests.

```shell
export TEST_FLAG=1
go test ./infrastructure -v -cover
go test ./repository -v -cover
go test ./usecase -v -cover
go test ./controller -v -cover
go test ./handler -v -cover
```

## DynamoDB Tips

### CLI

```bash
# Create table
$ aws dynamodb create-table --cli-input-json file://docker/dynamodb/sample/table.json --endpoint-url http://127.0.0.1:3307 --region ap-northeast-1
# Show table
$ aws dynamodb describe-table --table-name TableName --endpoint-url http://127.0.0.1:3307 --region ap-northeast-1
# Insert test data
$ aws dynamodb batch-write-item --request-items file://docker/dynamodb/sample/data.json --endpoint-url http://127.0.0.1:3307 --region ap-northeast-1
# Scan table
$ aws dynamodb scan --table-name TableName --endpoint-url http://127.0.0.1:3307 --region ap-northeast-1
# Delete table
$ aws dynamodb delete-table --table-name TableName --endpoint-url http://127.0.0.1:3307 --region ap-northeast-1
```
