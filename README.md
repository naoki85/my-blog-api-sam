# my-blog-api-sam

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)

## Setup process

Start Mysql container.

```shell
docker network create lambda-local
docker-compose up db
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

If the previous command ran successfully you should now be able to hit the following local endpoint to invoke your function `http://localhost:3000/hello`

## Packaging and deployment

First and foremost, we need a `S3 bucket` where we can upload our Lambda functions packaged as ZIP before we deploy anything - If you don't have a S3 bucket to store code artifacts then this is a good time to create one:

```bash
aws s3 mb s3://BUCKET_NAME
```

After create S3 bucket, Set following environment variables.

```bash
export USERNAME='username'
export PASSWORD='password'
export HOST='host'
export PORT='port'
export DBNAME='dbname'
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

We use `testing` package that is built-in in Golang and you can simply run the following command to run our tests:

```shell
go test -v ./handler
```
