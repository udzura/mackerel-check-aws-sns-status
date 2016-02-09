package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	ARN               string `short:"a" long:"arn" required:"true" description:"Platform application ARN to check"`
	WarnThreshold     int    `short:"w" long:"warn" required:"false" default:"30" description:"A threshold to warn cert expiration (in days)"`
	CriticalThreshold int    `short:"c" long:"critical" required:"false" default:"14" description:"A threshold to judge critical for cert expiration (in days)"`
}

func main() {
	run()
}

func run() {
	_, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		os.Exit(127)
	}

	svc := sns.New(session.New())
	params := &sns.GetPlatformApplicationAttributesInput{
		PlatformApplicationArn: aws.String(opts.ARN),
	}

	resp, err := svc.GetPlatformApplicationAttributes(params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(3)
	}

	attr := resp.Attributes

	if val, ok := attr["AppleCertificateExpirationDate"]; ok {
		expireAt, err := time.Parse("2006-01-02T15:04:05Z", *val)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
			os.Exit(3)
		}

		switch {
		case time.Now().After(expireAt):
			fmt.Fprintf(os.Stderr, "Cert is now expired: %s\n", expireAt)
			os.Exit(3)
		case time.Now().AddDate(0, 0, opts.CriticalThreshold).After(expireAt):
			duration := time.Now().Sub(expireAt) / (24 * time.Hour) * -1
			fmt.Fprintf(os.Stderr, "Cert is going to expire in %d days: %s\n", duration, expireAt)
			os.Exit(2)
		case time.Now().AddDate(0, 0, opts.WarnThreshold).After(expireAt):
			duration := time.Now().Sub(expireAt) / (24 * time.Hour) * -1
			fmt.Fprintf(os.Stderr, "Cert is going to expire in %d days: %s\n", duration, expireAt)
			os.Exit(1)
		default:
			fmt.Fprintf(os.Stderr, "Cert is OK: %s\n", expireAt)
			os.Exit(0)
		}
	}

	if val, ok := attr["Enabled"]; ok {
		enabled := *val
		if enabled != "true" {
			fmt.Fprintf(os.Stderr, "Endpoint is disabled!!\n")
			os.Exit(3)
		}
	}

	fmt.Fprintf(os.Stderr, "Endpoint is enabled\n")
	os.Exit(0)
}
