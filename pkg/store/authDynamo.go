package store

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ddbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/teris-io/shortid"

	"github.com/abdulloh76/user-service/pkg/types"
	"github.com/abdulloh76/user-service/pkg/utils"
)

type AuthDynamoStore struct {
	client    *dynamodb.Client
	tableName string
}

var _ types.AuthStore = (*AuthDynamoStore)(nil)

func NewAuthDynamoStore(ctx context.Context, DYNAMODB_PORT, tableName string) *AuthDynamoStore {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("localhost"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:" + DYNAMODB_PORT}, nil
			})))

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)

	return &AuthDynamoStore{
		client:    client,
		tableName: tableName,
	}
}

func (d *AuthDynamoStore) CreateUser(ctx context.Context, userCredentials types.SignInCredentials) (string, error) {
	credentials, err := attributevalue.MarshalMap(&userCredentials)
	if err != nil {
		return "", utils.ErrJsonUnmarshal
	}

	id, _ := shortid.Generate()
	credentials["id"] = &ddbtypes.AttributeValueMemberS{Value: id}

	_, err = d.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &d.tableName,
		Item:      credentials,
	})

	if err != nil {
		return "", utils.ErrWithDB
	}

	return id, nil
}

func (d *AuthDynamoStore) GetUser(ctx context.Context, email, passwordHash string) (*types.User, error) {
	response, err := d.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &d.tableName,
		Key: map[string]ddbtypes.AttributeValue{
			"email": &ddbtypes.AttributeValueMemberS{Value: email},
		},
	})
	if err != nil {
		return nil, utils.ErrWithDB
	}
	if len(response.Item) == 0 {
		return nil, utils.ErrUserNotExists
	}

	user := types.User{}
	attributevalue.UnmarshalMap(response.Item, &user)

	if user.Password != passwordHash {
		return nil, utils.ErrWrongPassword
	}

	return &user, nil
}
