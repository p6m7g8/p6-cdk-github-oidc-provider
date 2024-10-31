package p6cdkgithuboidcprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

type IP6CDKGithubOidcProviderProps interface {
	Policies() *[]awsiam.IManagedPolicy
	Repo() *string
}

// The jsii proxy for IP6CDKGithubOidcProviderProps
type jsiiProxy_IP6CDKGithubOidcProviderProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IP6CDKGithubOidcProviderProps) Policies() *[]awsiam.IManagedPolicy {
	var returns *[]awsiam.IManagedPolicy
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IP6CDKGithubOidcProviderProps) Repo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repo",
		&returns,
	)
	return returns
}

