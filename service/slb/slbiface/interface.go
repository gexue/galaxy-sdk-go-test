// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package slbiface provides an interface to enable mocking the slb service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package slbiface

import (
	"github.com/KscSDK/ksc-sdk-go/service/slb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// SlbAPI provides an interface to enable mocking the
// slb.Slb service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// slb.
//	func myFunc(svc slbiface.SlbAPI) bool {
//	    // Make svc.AddAlbRule request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := slb.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockSlbClient struct {
//	    slbiface.SlbAPI
//	}
//	func (m *mockSlbClient) AddAlbRule(input *map[string]interface{}) (*map[string]interface{}, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockSlbClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type SlbAPI interface {
	AddAlbRule(*map[string]interface{}) (*map[string]interface{}, error)
	AddAlbRuleWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddAlbRuleRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AssociateCertificateWithGroup(*map[string]interface{}) (*map[string]interface{}, error)
	AssociateCertificateWithGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AssociateCertificateWithGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AssociateLoadBalancerAcl(*map[string]interface{}) (*map[string]interface{}, error)
	AssociateLoadBalancerAclWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AssociateLoadBalancerAclRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AssociateMirrorGroup(*map[string]interface{}) (*map[string]interface{}, error)
	AssociateMirrorGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AssociateMirrorGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ConfigureHealthCheck(*map[string]interface{}) (*map[string]interface{}, error)
	ConfigureHealthCheckWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ConfigureHealthCheckRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateAlb(*map[string]interface{}) (*map[string]interface{}, error)
	CreateAlbWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateAlbRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateAlbBackendServerGroup(*map[string]interface{}) (*map[string]interface{}, error)
	CreateAlbBackendServerGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateAlbBackendServerGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateAlbListener(*map[string]interface{}) (*map[string]interface{}, error)
	CreateAlbListenerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateAlbListenerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateAlbListenerCertGroup(*map[string]interface{}) (*map[string]interface{}, error)
	CreateAlbListenerCertGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateAlbListenerCertGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateAlbRuleGroup(*map[string]interface{}) (*map[string]interface{}, error)
	CreateAlbRuleGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateAlbRuleGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateBackendServerGroup(*map[string]interface{}) (*map[string]interface{}, error)
	CreateBackendServerGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateBackendServerGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateHostHeader(*map[string]interface{}) (*map[string]interface{}, error)
	CreateHostHeaderWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateHostHeaderRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateListeners(*map[string]interface{}) (*map[string]interface{}, error)
	CreateListenersWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateListenersRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateLoadBalancer(*map[string]interface{}) (*map[string]interface{}, error)
	CreateLoadBalancerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateLoadBalancerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateLoadBalancerAcl(*map[string]interface{}) (*map[string]interface{}, error)
	CreateLoadBalancerAclWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateLoadBalancerAclRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateLoadBalancerAclEntry(*map[string]interface{}) (*map[string]interface{}, error)
	CreateLoadBalancerAclEntryWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateLoadBalancerAclEntryRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateSlbRule(*map[string]interface{}) (*map[string]interface{}, error)
	CreateSlbRuleWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateSlbRuleRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAlb(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAlbWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAlbRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAlbBackendServerGroup(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAlbBackendServerGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAlbBackendServerGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAlbListener(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAlbListenerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAlbListenerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAlbListenerCertGroup(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAlbListenerCertGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAlbListenerCertGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAlbRule(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAlbRuleWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAlbRuleRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAlbRuleGroup(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAlbRuleGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAlbRuleGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteBackendServerGroup(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteBackendServerGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteBackendServerGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteHealthCheck(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteHealthCheckWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteHealthCheckRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteHostHeader(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteHostHeaderWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteHostHeaderRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteListeners(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteListenersWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteListenersRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteLoadBalancer(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteLoadBalancerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteLoadBalancerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteLoadBalancerAcl(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteLoadBalancerAclWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteLoadBalancerAclRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteLoadBalancerAclEntry(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteLoadBalancerAclEntryWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteLoadBalancerAclEntryRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteRule(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteRuleWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteRuleRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeregisterAlbBackendServer(*map[string]interface{}) (*map[string]interface{}, error)
	DeregisterAlbBackendServerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeregisterAlbBackendServerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeregisterBackendServer(*map[string]interface{}) (*map[string]interface{}, error)
	DeregisterBackendServerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeregisterBackendServerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeregisterBackendServerGroupFromListener(*map[string]interface{}) (*map[string]interface{}, error)
	DeregisterBackendServerGroupFromListenerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeregisterBackendServerGroupFromListenerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeregisterInstancesFromListener(*map[string]interface{}) (*map[string]interface{}, error)
	DeregisterInstancesFromListenerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeregisterInstancesFromListenerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAlbBackendServerGroups(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAlbBackendServerGroupsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAlbBackendServerGroupsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAlbBackendServers(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAlbBackendServersWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAlbBackendServersRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAlbListenerCertGroups(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAlbListenerCertGroupsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAlbListenerCertGroupsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAlbListeners(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAlbListenersWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAlbListenersRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAlbRuleGroups(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAlbRuleGroupsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAlbRuleGroupsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAlbs(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAlbsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAlbsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeBackendServerGroups(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeBackendServerGroupsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeBackendServerGroupsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeBackendServers(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeBackendServersWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeBackendServersRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeHealthChecks(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeHealthChecksWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeHealthChecksRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeHostHeaders(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeHostHeadersWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeHostHeadersRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstancesWithListener(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstancesWithListenerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstancesWithListenerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeListeners(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeListenersWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeListenersRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeLoadBalancerAcls(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeLoadBalancerAclsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeLoadBalancerAclsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeLoadBalancerAttributes(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeLoadBalancerAttributesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeLoadBalancerAttributesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeLoadBalancers(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeLoadBalancersWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeLoadBalancersRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRules(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRulesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRulesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisassociateLoadBalancerAcl(*map[string]interface{}) (*map[string]interface{}, error)
	DisassociateLoadBalancerAclWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisassociateLoadBalancerAclRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisassociateMirrorGroup(*map[string]interface{}) (*map[string]interface{}, error)
	DisassociateMirrorGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisassociateMirrorGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DissociateCertificateWithGroup(*map[string]interface{}) (*map[string]interface{}, error)
	DissociateCertificateWithGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DissociateCertificateWithGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyAlb(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyAlbWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyAlbRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyAlbBackendServer(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyAlbBackendServerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyAlbBackendServerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyAlbBackendServerGroup(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyAlbBackendServerGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyAlbBackendServerGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyAlbListener(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyAlbListenerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyAlbListenerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyAlbRuleGroup(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyAlbRuleGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyAlbRuleGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyBackendServer(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyBackendServerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyBackendServerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyBackendServerGroup(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyBackendServerGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyBackendServerGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyBackendServerGroupHealthCheck(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyBackendServerGroupHealthCheckWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyBackendServerGroupHealthCheckRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyHealthCheck(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyHealthCheckWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyHealthCheckRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyHostHeader(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyHostHeaderWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyHostHeaderRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyInstancesWithListener(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyInstancesWithListenerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyInstancesWithListenerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyListeners(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyListenersWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyListenersRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyLoadBalancer(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyLoadBalancerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyLoadBalancerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyLoadBalancerAcl(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyLoadBalancerAclWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyLoadBalancerAclRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyLoadBalancerAclEntry(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyLoadBalancerAclEntryWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyLoadBalancerAclEntryRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyLoadBalancerAttributes(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyLoadBalancerAttributesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyLoadBalancerAttributesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifySlbRule(*map[string]interface{}) (*map[string]interface{}, error)
	ModifySlbRuleWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifySlbRuleRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RegisterAlbBackendServer(*map[string]interface{}) (*map[string]interface{}, error)
	RegisterAlbBackendServerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RegisterAlbBackendServerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RegisterBackendServer(*map[string]interface{}) (*map[string]interface{}, error)
	RegisterBackendServerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RegisterBackendServerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RegisterBackendServerGroupWithListener(*map[string]interface{}) (*map[string]interface{}, error)
	RegisterBackendServerGroupWithListenerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RegisterBackendServerGroupWithListenerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RegisterInstancesWithListener(*map[string]interface{}) (*map[string]interface{}, error)
	RegisterInstancesWithListenerWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RegisterInstancesWithListenerRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetAlbAccessLog(*map[string]interface{}) (*map[string]interface{}, error)
	SetAlbAccessLogWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetAlbAccessLogRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetAlbDeleteProtection(*map[string]interface{}) (*map[string]interface{}, error)
	SetAlbDeleteProtectionWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetAlbDeleteProtectionRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetAlbModificationProtection(*map[string]interface{}) (*map[string]interface{}, error)
	SetAlbModificationProtectionWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetAlbModificationProtectionRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetAlbName(*map[string]interface{}) (*map[string]interface{}, error)
	SetAlbNameWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetAlbNameRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetAlbStatus(*map[string]interface{}) (*map[string]interface{}, error)
	SetAlbStatusWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetAlbStatusRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetEnableAlbAccessLog(*map[string]interface{}) (*map[string]interface{}, error)
	SetEnableAlbAccessLogWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetEnableAlbAccessLogRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ SlbAPI = (*slb.Slb)(nil)
