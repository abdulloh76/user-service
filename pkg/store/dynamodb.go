package store

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ddbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/abdulloh76/user-service/pkg/types"
	"github.com/abdulloh76/user-service/pkg/utils"
)

type DynamoDBStore struct {
	client    *dynamodb.Client
	tableName string
}

var _ types.UserStore = (*DynamoDBStore)(nil)

func NewDynamoDBStore(ctx context.Context, DYNAMODB_PORT, tableName string) *DynamoDBStore {
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

	return &DynamoDBStore{
		client:    client,
		tableName: tableName,
	}
}

func (d *DynamoDBStore) GetUserDetails(ctx context.Context, id string) (*types.User, error) {
	response, err := d.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &d.tableName,
		Key: map[string]ddbtypes.AttributeValue{
			"id": &ddbtypes.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		return nil, utils.ErrWithDB
	}

	user := types.User{}
	err = attributevalue.UnmarshalMap(response.Item, &user)

	if err != nil {
		return nil, utils.ErrWithDB
	}

	return &user, nil
}

func (d *DynamoDBStore) UpdateUserCredentials(ctx context.Context, id string, credentials types.UpdateCredentialsDto) error {
	_, err := d.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: &d.tableName,
		Key: map[string]ddbtypes.AttributeValue{
			"id": &ddbtypes.AttributeValueMemberS{Value: id},
		},
		ExpressionAttributeValues: map[string]ddbtypes.AttributeValue{
			":fn": &ddbtypes.AttributeValueMemberS{Value: credentials.FirstName},
			":ln": &ddbtypes.AttributeValueMemberS{Value: credentials.LastName},
			":e":  &ddbtypes.AttributeValueMemberS{Value: credentials.Email},
		},
		UpdateExpression: aws.String("set firstName=:fn, lastName=:ln, email=:e"),
		ReturnValues:     "ALL_NEW", // !!! why new?
	})

	if err != nil {
		return utils.ErrWithDB
	}

	return nil
}

func (d *DynamoDBStore) UpdatePassword(ctx context.Context, id string, newPasswordHash string) error {
	_, err := d.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: &d.tableName,
		Key: map[string]ddbtypes.AttributeValue{
			"id": &ddbtypes.AttributeValueMemberS{Value: id},
		},
		ExpressionAttributeValues: map[string]ddbtypes.AttributeValue{
			":pwd": &ddbtypes.AttributeValueMemberS{Value: newPasswordHash},
		},
		UpdateExpression: aws.String("set password=:pwd"),
		ReturnValues:     "ALL_NEW", // !!! why new?
	})

	if err != nil {
		return utils.ErrWithDB
	}

	return nil
}

func (d *DynamoDBStore) UpdateAddress(ctx context.Context, id string, address types.AddressModel) error {
	newAddress, _ := attributevalue.MarshalMap(address)

	_, err := d.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: &d.tableName,
		Key: map[string]ddbtypes.AttributeValue{
			"id": &ddbtypes.AttributeValueMemberS{Value: id},
		},
		ExpressionAttributeValues: map[string]ddbtypes.AttributeValue{
			":add": &ddbtypes.AttributeValueMemberM{Value: newAddress},
		},
		UpdateExpression: aws.String("set address=:add"),
		ReturnValues:     "ALL_NEW", // !!! why new?
	})

	if err != nil {
		return utils.ErrWithDB
	}

	return nil
}

func (d *DynamoDBStore) DeleteUser(ctx context.Context, id string) error {
	_, err := d.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &d.tableName,
		Key: map[string]ddbtypes.AttributeValue{
			"id": &ddbtypes.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		return utils.ErrWithDB
	}

	return nil
}
