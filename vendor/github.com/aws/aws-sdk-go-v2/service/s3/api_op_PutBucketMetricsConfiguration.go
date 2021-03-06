// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3/internal/arn"
)

type PutBucketMetricsConfigurationInput struct {
	_ struct{} `type:"structure" payload:"MetricsConfiguration"`

	// The name of the bucket for which the metrics configuration is set.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The ID used to identify the metrics configuration.
	//
	// Id is a required field
	Id *string `location:"querystring" locationName:"id" type:"string" required:"true"`

	// Specifies the metrics configuration.
	//
	// MetricsConfiguration is a required field
	MetricsConfiguration *MetricsConfiguration `locationName:"MetricsConfiguration" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutBucketMetricsConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketMetricsConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketMetricsConfigurationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if s.MetricsConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("MetricsConfiguration"))
	}
	if s.MetricsConfiguration != nil {
		if err := s.MetricsConfiguration.Validate(); err != nil {
			invalidParams.AddNested("MetricsConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketMetricsConfigurationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketMetricsConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.MetricsConfiguration != nil {
		v := s.MetricsConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "MetricsConfiguration", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "id", protocol.StringValue(v), metadata)
	}
	return nil
}

func (s *PutBucketMetricsConfigurationInput) getEndpointARN() (arn.Resource, error) {
	if s.Bucket == nil {
		return nil, fmt.Errorf("member Bucket is nil")
	}
	return parseEndpointARN(*s.Bucket)
}

func (s *PutBucketMetricsConfigurationInput) hasEndpointARN() bool {
	if s.Bucket == nil {
		return false
	}
	return arn.IsARN(*s.Bucket)
}

type PutBucketMetricsConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketMetricsConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketMetricsConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutBucketMetricsConfiguration = "PutBucketMetricsConfiguration"

// PutBucketMetricsConfigurationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Sets a metrics configuration (specified by the metrics configuration ID)
// for the bucket. You can have up to 1,000 metrics configurations per bucket.
// If you're updating an existing metrics configuration, note that this is a
// full replacement of the existing metrics configuration. If you don't include
// the elements you want to keep, they are erased.
//
// To use this operation, you must have permissions to perform the s3:PutMetricsConfiguration
// action. The bucket owner has this permission by default. The bucket owner
// can grant this permission to others. For more information about permissions,
// see Permissions Related to Bucket Subresource Operations (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html).
//
// For information about CloudWatch request metrics for Amazon S3, see Monitoring
// Metrics with Amazon CloudWatch (https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html).
//
// The following operations are related to PutBucketMetricsConfiguration:
//
//    * DeleteBucketMetricsConfiguration
//
//    * PutBucketMetricsConfiguration
//
//    * ListBucketMetricsConfigurations
//
// GetBucketLifecycle has the following special error:
//
//    * Error code: TooManyConfigurations Description: You are attempting to
//    create a new configuration but have already reached the 1,000-configuration
//    limit. HTTP Status Code: HTTP 400 Bad Request
//
//    // Example sending a request using PutBucketMetricsConfigurationRequest.
//    req := client.PutBucketMetricsConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketMetricsConfiguration
func (c *Client) PutBucketMetricsConfigurationRequest(input *PutBucketMetricsConfigurationInput) PutBucketMetricsConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutBucketMetricsConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?metrics",
	}

	if input == nil {
		input = &PutBucketMetricsConfigurationInput{}
	}

	req := c.newRequest(op, input, &PutBucketMetricsConfigurationOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutBucketMetricsConfigurationRequest{Request: req, Input: input, Copy: c.PutBucketMetricsConfigurationRequest}
}

// PutBucketMetricsConfigurationRequest is the request type for the
// PutBucketMetricsConfiguration API operation.
type PutBucketMetricsConfigurationRequest struct {
	*aws.Request
	Input *PutBucketMetricsConfigurationInput
	Copy  func(*PutBucketMetricsConfigurationInput) PutBucketMetricsConfigurationRequest
}

// Send marshals and sends the PutBucketMetricsConfiguration API request.
func (r PutBucketMetricsConfigurationRequest) Send(ctx context.Context) (*PutBucketMetricsConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketMetricsConfigurationResponse{
		PutBucketMetricsConfigurationOutput: r.Request.Data.(*PutBucketMetricsConfigurationOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketMetricsConfigurationResponse is the response type for the
// PutBucketMetricsConfiguration API operation.
type PutBucketMetricsConfigurationResponse struct {
	*PutBucketMetricsConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketMetricsConfiguration request.
func (r *PutBucketMetricsConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
