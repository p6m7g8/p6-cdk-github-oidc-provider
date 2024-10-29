package p6cdkgithuboidcprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/p6m7g8/p6-cdk-github-oidc-provider/p6cdkgithuboidcprovider/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/p6m7g8/p6-cdk-github-oidc-provider/p6cdkgithuboidcprovider/internal"
)

type P6CDKGithubOidcProvider interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	RoleArn() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for P6CDKGithubOidcProvider
type jsiiProxy_P6CDKGithubOidcProvider struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_P6CDKGithubOidcProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_P6CDKGithubOidcProvider) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}


func NewP6CDKGithubOidcProvider(scope constructs.Construct, id *string, props IP6CDKGithubOidcProviderProps) P6CDKGithubOidcProvider {
	_init_.Initialize()

	if err := validateNewP6CDKGithubOidcProviderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_P6CDKGithubOidcProvider{}

	_jsii_.Create(
		"p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewP6CDKGithubOidcProvider_Override(p P6CDKGithubOidcProvider, scope constructs.Construct, id *string, props IP6CDKGithubOidcProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider",
		[]interface{}{scope, id, props},
		p,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func P6CDKGithubOidcProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateP6CDKGithubOidcProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_P6CDKGithubOidcProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

