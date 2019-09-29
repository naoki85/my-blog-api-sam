package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/naoki85/my-blog-api-sam/config"
)

type DynamoDbHandler struct {
	Conn *dynamodb.DynamoDB
}

func NewDynamoDbHandler(c *config.Config) (*DynamoDbHandler, error) {
	sess := session.Must(session.NewSession())
	config := aws.NewConfig().WithRegion("ap-north-east-1")
	config = config.WithEndpoint(c.DynamoDbEndpoint)

	DynamoDbHandler := new(DynamoDbHandler)
	DynamoDbHandler.Conn = dynamodb.New(sess, config)
	return DynamoDbHandler, nil
}

//func (handler *DynamoDbHandler) Execute(statement string, args ...interface{}) (DynamoDbResult, error) {
//	res := DynamoDbResult{}
//	result, err := handler.Conn.GetItem(&dynamodb.GetItemInput{
//		TableName: aws.String("RecommendedBooks"),
//		AttributesToGet: []*string{
//			aws.String("Id"),
//			aws.String("Link"),
//			aws.String("ImageUrl"),
//			aws.String("ButtonUrl"),
//		}})
//	if err != nil {
//		log.Printf("%s", err.Error())
//		return res, err
//	}
//	res.Result = result
//	return res, err
//}
//
//func (handler *DynamoDbHandler) Query(statement string, args ...interface{}) (DynamoDbRow, error) {
//	rows, err := handler.Conn.Query(statement, args...)
//	if err != nil {
//		log.Printf("%s", err.Error())
//		return new(SqlRow), err
//	}
//	row := new(SqlRow)
//	row.Rows = rows
//	return row, nil
//}
//
//type DynamoDbResult struct {}
//
//func (r DynamoDbResult) LastInsertId() (int64, error) {
//	return 0, nil
//}
//
//func (r DynamoDbResult) RowsAffected() (int64, error) {
//	return 0, nil
//}
//
//type DynamoDbRow struct {
//	Result
//}
//
//func (r DynamoDbRow) Scan(dest ...interface{}) error {
//	return r.Rows.Scan(dest...)
//}
//
//func (r DynamoDbRow) Next() bool {
//	return r.Rows.Next()
//}
//
//func (r DynamoDbRow) Close() error {
//	return r.Rows.Close()
//}
