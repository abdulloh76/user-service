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

func (d *DynamoDBStore) AllUsers(ctx context.Context) ([]types.User, error) {
	users := []types.User{}

	input := &dynamodb.ScanInput{
		TableName: &d.tableName,
		Limit:     aws.Int32(20),
	}

	result, err := d.client.Scan(ctx, input)

	if err != nil {
		return nil, utils.ErrWithDB
	}

	err = attributevalue.UnmarshalListOfMaps(result.Items, &users)
	if err != nil {
		return users, utils.ErrWithDB
	}

	return users, nil
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

func (d *DynamoDBStore) CreateUser(ctx context.Context, user types.UserBody) error {
	item, err := attributevalue.MarshalMap(&user)
	if err != nil {
		return utils.ErrWithDB
	}

	id, _ := shortid.Generate()
	item["id"] = &ddbtypes.AttributeValueMemberS{Value: id}

	_, err = d.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &d.tableName,
		Item:      item,
	})

	if err != nil {
		return utils.ErrWithDB
	}

	return nil
}

func (d *DynamoDBStore) UpdateUserDetails(ctx context.Context, id string, user types.UserBody) (*types.User, error) {
	address, _ := attributevalue.MarshalMap(user.Address)

	response, err := d.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: &d.tableName,
		Key: map[string]ddbtypes.AttributeValue{
			"id": &ddbtypes.AttributeValueMemberS{Value: id},
		},
		ExpressionAttributeValues: map[string]ddbtypes.AttributeValue{
			":fn":  &ddbtypes.AttributeValueMemberS{Value: user.FirstName},
			":ln":  &ddbtypes.AttributeValueMemberS{Value: user.LastName},
			":e":   &ddbtypes.AttributeValueMemberS{Value: user.Email},
			":add": &ddbtypes.AttributeValueMemberM{Value: address},
		},
		UpdateExpression: aws.String("set firstName=:fn, lastName=:ln, email=:e, address=:add"),
		ReturnValues:     "ALL_NEW",
	})

	if err != nil {
		return nil, utils.ErrWithDB
	}

	updatedUser := &types.User{}
	attributevalue.UnmarshalMap(response.Attributes, &updatedUser)

	return updatedUser, nil
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
