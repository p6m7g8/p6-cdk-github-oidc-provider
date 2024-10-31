// A AWS CDK Github OIDC Provider
package p6cdkgithuboidcprovider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider",
		reflect.TypeOf((*P6CDKGithubOidcProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_P6CDKGithubOidcProvider{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
}
