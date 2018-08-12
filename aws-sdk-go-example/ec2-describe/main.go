package main

import (
    "fmt"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
    /*
    profile file ~/.aws/credentials
    */
    sess, _ := session.NewSessionWithOptions(session.Options{
        Config: aws.Config{
            Region: aws.String("ap-northeast-2"),
        },
        Profile: "goodysuits",
    })

    Describe(*sess)
}

func Describe(sess session.Session) {
    svc := ec2.New(&sess)

    /*
    DescribeInstancesInput struct {
        DryRun *bool `locationName:"dryRun" type:"boolean"`
        Filters []*Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`
        InstanceIds []*string `locationName:"InstanceId" locationNameList:"InstanceId" type:"list"`
        MaxResults *int64 `locationName:"maxResults" type:"integer"`
        NextToken *string `locationName:"nextToken" type:"string"`
    }
    */

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
        fmt.Println("DescribeInstances Fail")
        os.Exit(1)
    } else {
        fmt.Println(result)
    }
}
