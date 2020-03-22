#!/bin/bash

aws dynamodb create-table --cli-input-json file://docker/dynamodb/sample/Categories.json --endpoint-url http://127.0.0.1:3307 --region ap-northeast-1
aws dynamodb create-table --cli-input-json file://docker/dynamodb/sample/IdCounter.json --endpoint-url http://127.0.0.1:3307 --region ap-northeast-1
aws dynamodb create-table --cli-input-json file://docker/dynamodb/sample/Posts.json --endpoint-url http://127.0.0.1:3307 --region ap-northeast-1
aws dynamodb create-table --cli-input-json file://docker/dynamodb/sample/RecommendedBooks.json --endpoint-url http://127.0.0.1:3307 --region ap-northeast-1
aws dynamodb create-table --cli-input-json file://docker/dynamodb/sample/Users.json --endpoint-url http://127.0.0.1:3307 --region ap-northeast-1

aws dynamodb batch-write-item --request-items file://docker/dynamodb/sample/CategoriesData.json --endpoint-url http://127.0.0.1:3307 --region ap-northeast-1
aws dynamodb batch-write-item --request-items file://docker/dynamodb/sample/IdCounterData.json --endpoint-url http://127.0.0.1:3307 --region ap-northeast-1
aws dynamodb batch-write-item --request-items file://docker/dynamodb/sample/PostsData.json --endpoint-url http://127.0.0.1:3307 --region ap-northeast-1
aws dynamodb batch-write-item --request-items file://docker/dynamodb/sample/RecommendedBooksData.json --endpoint-url http://127.0.0.1:3307 --region ap-northeast-1
aws dynamodb batch-write-item --request-items file://docker/dynamodb/sample/UsersData.json --endpoint-url http://127.0.0.1:3307 --region ap-northeast-1
