package p6cdkgithuboidcprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IP6CDKGithubOidcProviderProps interface {
	Repo() *string
}

// The jsii proxy for IP6CDKGithubOidcProviderProps
type jsiiProxy_IP6CDKGithubOidcProviderProps struct {
	_ byte // padding
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

