package main

import (
	"errors"
)

// Error used when the --profile flag is missing
var ErrMissingAWSProfile error = errors.New("Missing AWS Profile")

// Error used when the --region flag is missing
var ErrMissingAWSRegion error = errors.New("Missing AWS Region")

// Error used when the --stack flag is missing
var ErrMissingStackName error = errors.New("Missing CloudFormation stack name")
