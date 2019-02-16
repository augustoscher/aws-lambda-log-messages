package main

import (
	"fmt"

	"bitbucket.org/augustoscher/lambda-log-erros/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	uuid "github.com/satori/go.uuid"
)

const tableName = "mensagemerro"

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-2"))

// GetMensagemErro retrieves one MensagemErro from the DB based on its ID
func GetMensagemErro(uuid string) (model.MensagemErro, error) {
	// Prepares the input to retrieve the item with the given ID
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(uuid),
			},
		},
	}

	// Retrieves the item
	result, err := db.GetItem(input)
	if err != nil {
		return model.MensagemErro{}, err
	}
	if result.Item == nil {
		return model.MensagemErro{}, nil
	}

	// Unmarshals the object retrieved into a domain struct
	var msg model.MensagemErro
	err = dynamodbattribute.UnmarshalMap(result.Item, &msg)
	if err != nil {
		return model.MensagemErro{}, err
	}

	return msg, nil
}

// GetMensagensErro retrieves all the users from the DB
func GetMensagensErro() ([]model.MensagemErro, error) {

	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	result, err := db.Scan(input)
	if err != nil {
		return []model.MensagemErro{}, err
	}
	if len(result.Items) == 0 {
		return []model.MensagemErro{}, nil
	}

	var msgs []model.MensagemErro
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &msgs)
	if err != nil {
		return []model.MensagemErro{}, err
	}

	return msgs, nil
}

// CreateMensagemErro inserts a new MensagemErro item to the table.
func CreateMensagemErro(msg model.MensagemErro) error {

	// Generates a new random ID
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	// Creates the item that's going to be inserted
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(fmt.Sprintf("%v", uuid)),
			},
			"codintegracao": {
				S: aws.String(msg.Codintegracao),
			},
			"conteudomensagem": {
				S: aws.String(msg.Conteudomensagem),
			},
			"descricaoerro": {
				S: aws.String(msg.Descricaoerro),
			},
		},
	}

	_, err = db.PutItem(input)
	return err
}
