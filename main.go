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

var env CliEnv

func setup(app *cli.App) {
	app.Name = AppName
	app.Copyright = AppCopyright
	app.HelpName = AppName
	app.Version = AppVersion
	app.Usage = "Displays all outputs in a CloudFormation stack given the names of the outputs, whether they are exported or not."
	app.ArgsUsage = "OUTPUT1 ..."
}

func main() {
	app := cli.NewApp()
	setup(app)

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

	app.Action = func(c *cli.Context) error {
		if c.NArg() == 0 {
			cli.ShowAppHelpAndExit(c, 1)
		}

		if err := env.Validate(); err != nil {
			return err
		}

		config, err := external.LoadDefaultAWSConfig(
			external.WithSharedConfigProfile(env.AWSProfile),
			external.WithRegion(env.AWSRegion),
		)
		if err != nil {
			return err
		}

		client := cloudformation.New(config)

		input := cloudformation.DescribeStacksInput{
			StackName: aws.String(env.StackName),
		}

		req := client.DescribeStacksRequest(&input)
		res, err := req.Send(context.TODO())
		if err != nil {
			return err
		}

		nameMap := make(map[string]struct{}, c.NArg())
		for _, a := range c.Args() {
			nameMap[a] = struct{}{}
		}

		for _, s := range res.Stacks {
			for _, o := range s.Outputs {
				if _, ok := nameMap[*o.OutputKey]; ok {
					fmt.Println(*o.OutputValue)
				}
			}
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
