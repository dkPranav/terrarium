package services

import (
	"context"
	"log"
	"time"

	"github.com/terrariumcloud/terrarium-grpc-gateway/internal/storage"
	terrarium "github.com/terrariumcloud/terrarium-grpc-gateway/pkg/terrarium/module"
	grpc "google.golang.org/grpc"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

const (
	DefaultVersionsTableName      = "terrarium-module-versions"
	DefaultVersionManagerEndpoint = "version_manager:3001"
)

var VersionsTableName string = DefaultVersionsTableName
var VersionManagerEndpoint string = DefaultVersionManagerEndpoint

type VersionManagerService struct {
	UnimplementedVersionManagerServer
	Db     dynamodbiface.DynamoDBAPI
	Table  string
	Schema *dynamodb.CreateTableInput
}

type ModuleVersion struct {
	Name        string `json:"name" bson:"name" dynamodbav:"name"`
	Version     string `json:"version" bson:"version" dynamodbav:"version"`
	CreatedOn   string `json:"created_on" bson:"created_on" dynamodbav:"created_on"`
	PublishedOn string `json:"published_on" bson:"published_on" dynamodbav:"published_on"`
}

func (s *VersionManagerService) RegisterWithServer(grpcServer grpc.ServiceRegistrar) error {
	RegisterVersionManagerServer(grpcServer, s)
	if err := storage.InitializeDynamoDb(s.Table, s.Schema, s.Db); err != nil {
		return err
	}
	return nil
}

// Creates new Module Version with Version Manager service
func (s *VersionManagerService) BeginVersion(ctx context.Context, request *BeginVersionRequest) (*terrarium.BeginVersionResponse, error) {
	log.Println("Creating new version.")

	mv := ModuleVersion{
		Name:      request.GetModule().GetName(),
		Version:   request.GetModule().GetVersion(),
		CreatedOn: time.Now().UTC().String(),
	}

	av, err := dynamodbattribute.MarshalMap(mv)

	if err != nil {
		log.Printf("Failed to marshal: %s", err.Error())
		return nil, err
	}

	in := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(VersionsTableName),
	}

	if _, err = s.Db.PutItem(in); err != nil {
		log.Printf("Failed to put item: %s", err.Error())
		return nil, err
	}

	response := &terrarium.BeginVersionResponse{
		SessionKey: "123",
	}

	log.Println("New version created.")
	return response, nil
}

// Removes Module Version with Version Manager service
func (s *VersionManagerService) AbortVersion(ctx context.Context, request *TerminateVersionRequest) (*terrarium.TransactionStatusResponse, error) {
	log.Println("Aborting module version.")

	in := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(request.Module.GetName()),
			},
			"version": {
				S: aws.String(request.Module.GetVersion()),
			},
		},
		TableName: aws.String(VersionsTableName),
	}

	if _, err := s.Db.DeleteItem(in); err != nil {
		log.Printf("Failed to delete item: %s", err.Error())
		return SessionKeyNotRemoved, err
	}

	log.Println("Module version aborted.")
	return VersionAborted, nil
}

// Updates Module Version to published with Verison Manager service
func (s *VersionManagerService) PublishVersion(ctx context.Context, request *TerminateVersionRequest) (*terrarium.TransactionStatusResponse, error) {
	log.Println("Publishing module version.")

	in := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":published_on": {
				S: aws.String(time.Now().UTC().String()),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"name": {
				S: aws.String(request.Module.GetName()),
			},
			"version": {
				S: aws.String(request.Module.GetVersion()),
			},
		},
		TableName:        aws.String(VersionsTableName),
		UpdateExpression: aws.String("set published_on = :published_on"),
	}

	if _, err := s.Db.UpdateItem(in); err != nil {
		log.Printf("Failed to update item: %s", err.Error())
		return nil, err
	}

	log.Println("Module version published.")
	return VersionPublished, nil
}

// GetModuleVersoinsSchema returns CreateTableInput
// that can be used to create table if it does not exist
func GetModuleVersionsSchema(table string) *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("name"),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},
			{
				AttributeName: aws.String("version"),
				AttributeType: aws.String(dynamodb.ScalarAttributeTypeS),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("name"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("version"),
				KeyType:       aws.String("RANGE"),
			},
		},
		TableName:   aws.String(table),
		BillingMode: aws.String(dynamodb.BillingModeProvisioned),
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	}
}
