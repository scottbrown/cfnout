package main

// Contains the configuration for the CLI app
type CliEnv struct {
	AWSProfile string
	AWSRegion  string
	StackName  string
}

// Returns nil if the environment configuration is valid; otherwise error.
func (e CliEnv) Validate() error {
	if e.AWSProfile == "" {
		return ErrMissingAWSProfile
	}

	if e.AWSRegion == "" {
		return ErrMissingAWSRegion
	}

	if e.StackName == "" {
		return ErrMissingStackName
	}

	return nil
}
