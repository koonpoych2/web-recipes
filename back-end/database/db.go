package database

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getParameters(names []string) (map[string]string, error) {

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	// Create SSM client
	client := ssm.NewFromConfig(cfg)

	// Fetch parameters
	resp, err := client.GetParameters(context.TODO(), &ssm.GetParametersInput{
		Names:          names,
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get parameters: %w", err)
	}

	// Store parameters in a map
	params := make(map[string]string)
	for _, param := range resp.Parameters {
		params[*param.Name] = *param.Value
	}

	return params, nil
}

func ConnectDB() {

	parameterNames := []string{
		"/mycook/db-host",
		"/mycook/db-name",
		"/mycook/db-password",
		"/mycook/db-username",
	}

	params, err := getParameters(parameterNames)
	if err != nil {
		log.Fatal("Error retrieving parameters:", err)
	}

	fmt.Println("Retrieved parameters:", params)
	// var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		params["/mycook/db-username"], params["/mycook/db-password"],
		params["/mycook/db-host"], "3306", params["/mycook/db-name"],
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connected successfully")
}
