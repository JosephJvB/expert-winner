package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var client *s3.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}
	
	client = s3.NewFromConfig(cfg)
}
func main() {
	var token *string
	keys := []string{}
	x := 1
	for ok := true; ok; ok = (token != nil) {
		fmt.Println("loopNum", x)
		x++

		r, err := getNext(token)
		if err != nil {
			panic(err)
		}

		fmt.Println("got", len(r.Contents), "keys")
		fmt.Println("totalKeys:", len(keys))

		for i := 0; i < len(r.Contents); i++ {
			keys = append(keys, *r.Contents[i].Key)
		}
		token = r.NextContinuationToken
	}
	fmt.Println("done. Got", len(keys), "keys in total")
}

func getNext(key *string) (*s3.ListObjectsV2Output, error) {
	if key != nil {
		fmt.Println("  > key ==", *key)
	}
	params := &s3.ListObjectsV2Input{
		Bucket: aws.String("milkbooks-design"),
		Prefix: aws.String("library/1951600/2.2022/thumbnail"),
	}
	if key != nil {
		params.ContinuationToken = key
	}
	return client.ListObjectsV2(context.TODO(), params)
}
