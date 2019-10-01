package infrastructure

import (
	"github.com/naoki85/my-blog-api-sam/config"
	"testing"
)

func TestGetDynamoDbHandler(t *testing.T) {
	config.InitDbConf("test")
	c := config.GetDbConf()
	_, err := NewDynamoDbHandler(c)

	if err != nil {
		t.Fatalf("Cannot connect to database: %s", err)
	}
}
