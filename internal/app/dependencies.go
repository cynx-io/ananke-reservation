package app

import (
	"context"
	"github.com/cynx-io/ananke-reservation/internal/dependencies"
	"github.com/cynx-io/ananke-reservation/internal/dependencies/config"
	"github.com/cynx-io/cynx-core/src/externalapi/email"
	"github.com/cynx-io/cynx-core/src/logger"
	"github.com/cynx-io/cynx-core/src/model/dto"
)

type Dependencies struct {
	DatabaseClient *dependencies.DatabaseClient
}

func NewDependencies(ctx context.Context) *Dependencies {

	logger.Info(ctx, "Connecting to Database")
	databaseClient, err := dependencies.NewDatabaseClient()
	if err != nil {
		logger.Fatal(ctx, "Failed to connect to database: ", err)
	}

	email.Init(ctx, dto.AwsConfig{
		Region:          config.Config.Aws.Region,
		AccessKeyID:     config.Config.Aws.AccessKeyID,
		SecretAccessKey: config.Config.Aws.SecretAccessKey,
	})

	logger.Info(ctx, "Dependencies initialized")
	return &Dependencies{
		DatabaseClient: databaseClient,
	}
}
