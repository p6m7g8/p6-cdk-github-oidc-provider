import { Stack } from 'aws-cdk-lib'
import * as iam from 'aws-cdk-lib/aws-iam'
import { Construct } from 'constructs'

export interface IP6CDKGithubOidcProviderProps {
  readonly repo: string
}

export class P6CDKGithubOidcProvider extends Construct {
  public readonly roleArn: string

  constructor(scope: Construct, id: string, props: IP6CDKGithubOidcProviderProps) {
    super(scope, id)

    const providerArn = `arn:aws:iam::${Stack.of(this).account}:oidc-provider/token.actions.githubusercontent.com`

    const existingOidcProvider = iam.OpenIdConnectProvider.fromOpenIdConnectProviderArn(this, 'ExistingOidcProvider', providerArn)

    // Create the OIDC Provider for GitHub
    const oidcProvider = existingOidcProvider ?? new iam.OpenIdConnectProvider(this, 'GithubOidc', {
      url: 'https://token.actions.githubusercontent.com',
      clientIds: ['sts.amazonaws.com'],
      thumbprints: ['ffffffffffffffffffffffffffffffffffffffff'],
    })

    // Create the IAM Role for GitHub OIDC Authentication
    const role = new iam.Role(this, 'GithubOidcRole', {
      assumedBy: new iam.FederatedPrincipal(
        oidcProvider.openIdConnectProviderArn,
        {
          StringEquals: {
            'token.actions.githubusercontent.com:aud': 'sts.amazonaws.com',
          },
          StringLike: {
            [`token.actions.githubusercontent.com:sub`]: `repo:${props.repo}:*`,
          },
        },
        'sts:AssumeRoleWithWebIdentity',
      ),
    })

    // Expose the role ARN
    this.roleArn = role.roleArn
  }
}
