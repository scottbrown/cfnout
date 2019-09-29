package main

import (
	"errors"
)

var ErrMissingAWSProfile error = errors.New("Missing AWS Profile")

var ErrMissingAWSRegion error = errors.New("Missing AWS Region")

var ErrMissingStackName error = errors.New("Missing CloudFormation stack name")
