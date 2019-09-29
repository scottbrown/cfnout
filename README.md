# cfnout

A utility tool when working with AWS CloudFormation stacks that allows you
to quickly return the value of a named stack output without having to
worry about JMESPATH or JSON vs TEXT output.

## Usage

```bash
$ cfnout --profile PROFILE --region REGION --stack STACK OUTPUT_NAME
```

where:
* PROFILE: the name of the AWS profile to use from `.aws/config`
* REGION: the region to target in AWS
* STACK: the name of the stack
* OUTPUT_NAME: The name of the output variable to return

## License

[MIT](LICENSE)
