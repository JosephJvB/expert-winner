package clients

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var ddbClient *dynamodb.Client

func init() {
	log.SetPrefix("clients.ddb: ")
	log.SetFlags(0)
	log.Println("init()")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatal(err)
	}
	
	ddbClient = dynamodb.NewFromConfig(cfg)
}

type UserDoc struct {
	EmailAddress string
	UserData string
}
type UserDetails struct {
	FirstName string
	LastName string
}
type UserData struct {
	UserDetails UserDetails
}

func RunDdb() {
	log.Println("RunDdb()")
	r, err := getItem("joevanbo+6@gmail.com")
	if err != nil {
		panic(err)
	}
	log.Println("Got userData raw", r.Item)
	user := UserDoc{}
	err = attributevalue.UnmarshalMap(r.Item, &user)
	if err != nil {
		panic(err)
	}
	log.Println("user data as struct", user)
}

func getItem(email string) (*dynamodb.GetItemOutput, error) {
	log.Println("getItem()", email)
	params := &dynamodb.GetItemInput{
		TableName: aws.String("DevUserData"),
		Key: map[string]types.AttributeValue{
			"EmailAddress": &types.AttributeValueMemberS{Value: email},
		},
	}
	return ddbClient.GetItem(context.TODO(), params)
}