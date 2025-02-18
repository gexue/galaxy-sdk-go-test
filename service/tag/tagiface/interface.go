// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package tagiface provides an interface to enable mocking the tag service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package tagiface

import (
	"github.com/gexue/galaxy-sdk-go-test/service/tag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// TagAPI provides an interface to enable mocking the
// tag.Tag service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // tag.
//    func myFunc(svc tagiface.TagAPI) bool {
//        // Make svc.CreateTags request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := tag.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockTagClient struct {
//        tagiface.TagAPI
//    }
//    func (m *mockTagClient) CreateTags(input *map[string]interface{}) (*map[string]interface{}, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockTagClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type TagAPI interface {
	CreateTags(*map[string]interface{}) (*map[string]interface{}, error)
	CreateTagsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateTagsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteTags(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteTagsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteTagsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeTagKeys(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeTagKeysWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeTagKeysRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeTagValues(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeTagValuesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeTagValuesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeTags(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeTagsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeTagsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ TagAPI = (*tag.Tag)(nil)
