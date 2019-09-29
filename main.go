// The main entry into the application
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/urfave/cli"
)

// A container of all the environment configuration for this CLI app
var env CliEnv

// Prepares this CLI application with flags, usage text, etc.
func setup(app *cli.App) {
	app.Name = AppName
	app.Copyright = AppCopyright
	app.HelpName = AppName
	app.Version = AppVersion
	app.Usage = "Displays all outputs in a CloudFormation stack given the names of the outputs, whether they are exported or not."
	app.ArgsUsage = "OUTPUT1 ..."

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "profile, p",
			Usage:       "AWS `PROFILE` to use",
			Destination: &env.AWSProfile,
			EnvVar:      "AWS_PROFILE",
		},
		cli.StringFlag{
			Name:        "region, r",
			Usage:       "AWS `REGION` where the stack resides",
			Destination: &env.AWSRegion,
			EnvVar:      "AWS_DEFAULT_REGION",
		},
		cli.StringFlag{
			Name:        "stack, s",
			Usage:       "`NAME` of the CloudFormation stack",
			Destination: &env.StackName,
		},
	}

	app.Action = start
}

// Main entry into the CLI app.
func main() {
	app := cli.NewApp()
	setup(app)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// Converts the args (which are the names of the outputs in a CFN stack)
// given by the user in the CLI to a map of those same names (since they
// should be unique).  This allows us to easily check if they exist in
// the stack's outputs.
func argsToNameMap(c *cli.Context) map[string]struct{} {
	nameMap := make(map[string]struct{}, c.NArg())
	for _, a := range c.Args() {
		nameMap[a] = struct{}{}
	}

	return nameMap
}

// Returns the stack by its name that was given in the environment
func stack() (cloudformation.Stack, error) {
	config, err := external.LoadDefaultAWSConfig(
		external.WithSharedConfigProfile(env.AWSProfile),
		external.WithRegion(env.AWSRegion),
	)

	client := cloudformation.New(config)

	input := cloudformation.DescribeStacksInput{
		StackName: aws.String(env.StackName),
	}

	req := client.DescribeStacksRequest(&input)
	res, err := req.Send(context.TODO())
	if err != nil {
		return cloudformation.Stack{}, err
	}

	return res.Stacks[0], nil
}

// The meat of the application
func start(c *cli.Context) error {
	if c.NArg() == 0 {
		cli.ShowAppHelpAndExit(c, 1)
	}

	nameMap := argsToNameMap(c)

	if err := env.Validate(); err != nil {
		return err
	}

	stack, err := stack()
	if err != nil {
		return err
	}

	for _, i := range stack.Outputs {
		if _, ok := nameMap[*i.OutputKey]; ok {
			fmt.Println(*i.OutputValue)
		}
	}

	return nil
}
