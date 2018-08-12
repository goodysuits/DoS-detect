package main

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
    sess, _ := session.NewSessionWithOptions(session.Options{
        Config: aws.Config{
            Region: aws.String("ap-northeast-2"),
        },
        Profile: "goodysuits-ec2-user",
    })

    DescribeInstances(*sess)
}

func DescribeInstances(sess session.Session) {
    svc := ec2.New(&sess)

    name := "*test*"
    params := &ec2.DescribeInstancesInput{
        Filters: []*ec2.Filter{
            {
                Name:   aws.String("tag:Name"),
                Values: []*string{aws.String(name)},
            },
        },
    }

    result, err := svc.DescribeInstances(params)

    if err != nil {
        if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
            default:
                fmt.Println(aerr.Error())
            }
        } else {
            fmt.Println(err.Error())
        }

        return
    }

    fmt.Println(result)
}
