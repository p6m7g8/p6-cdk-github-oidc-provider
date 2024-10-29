// A AWS CDK Github OIDC Provider
package p6cdkgithuboidcprovider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"p6-cdk-github-oidc-provider.IP6CDKGithubOidcProviderProps",
		reflect.TypeOf((*IP6CDKGithubOidcProviderProps)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "repo", GoGetter: "Repo"},
		},
		func() interface{} {
			return &jsiiProxy_IP6CDKGithubOidcProviderProps{}
		},
	)
	_jsii_.RegisterClass(
		"p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider",
		reflect.TypeOf((*P6CDKGithubOidcProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_P6CDKGithubOidcProvider{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
}
