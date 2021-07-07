// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package billunioniface provides an interface to enable mocking the bill-union service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package billunioniface

import (
	"github.com/KscSDK/ksc-sdk-go/service/billunion"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// BillUnionAPI provides an interface to enable mocking the
// billunion.BillUnion service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // bill-union.
//    func myFunc(svc billunioniface.BillUnionAPI) bool {
//        // Make svc.DescribeBillSummaryByPayMode request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := billunion.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockBillUnionClient struct {
//        billunioniface.BillUnionAPI
//    }
//    func (m *mockBillUnionClient) DescribeBillSummaryByPayMode(input *map[string]interface{}) (*map[string]interface{}, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockBillUnionClient{}
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
type BillUnionAPI interface {
	DescribeBillSummaryByPayMode(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeBillSummaryByPayModeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeBillSummaryByPayModeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeBillSummaryByProduct(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeBillSummaryByProductWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeBillSummaryByProductRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeBillSummaryByProject(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeBillSummaryByProjectWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeBillSummaryByProjectRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstanceSummaryBills(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstanceSummaryBillsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstanceSummaryBillsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeProductCode(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeProductCodeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeProductCodeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ BillUnionAPI = (*billunion.BillUnion)(nil)
