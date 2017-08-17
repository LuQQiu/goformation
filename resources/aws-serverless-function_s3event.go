package resources


// AWSServerlessFunction_S3Event AWS CloudFormation Resource (AWS::Serverless::Function.S3Event)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#s3
type AWSServerlessFunction_S3Event struct {
    
    // Bucket AWS CloudFormation Property
    // Required: true
    // See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#s3
    Bucket string `json:"Bucket,omitempty"`
    
    // Events AWS CloudFormation Property
    // Required: true
    // See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#s3
    Events *AWSServerlessFunction_StringOrListOfString `json:"Events,omitempty"`
    
    // Filter AWS CloudFormation Property
    // Required: false
    // See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#s3
    Filter *AWSServerlessFunction_S3NotificationFilter `json:"Filter,omitempty"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessFunction_S3Event) AWSCloudFormationType() string {
    return "AWS::Serverless::Function.S3Event"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSServerlessFunction_S3Event) AWSCloudFormationSpecificationVersion() string {
    return "2016-10-31"
}