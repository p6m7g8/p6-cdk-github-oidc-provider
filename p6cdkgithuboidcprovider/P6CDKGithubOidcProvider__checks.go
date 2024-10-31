//go:build !no_runtime_type_checking

package p6cdkgithuboidcprovider

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateP6CDKGithubOidcProvider_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewP6CDKGithubOidcProviderParameters(scope constructs.Construct, id *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

