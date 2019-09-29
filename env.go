package main

type CliEnv struct {
	AWSProfile string
	AWSRegion  string
	StackName  string
}

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
