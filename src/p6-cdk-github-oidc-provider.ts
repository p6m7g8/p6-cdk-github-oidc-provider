import * as iam from 'aws-cdk-lib/aws-iam'
import { Construct } from 'constructs'

export interface IP6CDKGithubOidcProviderProps {
  readonly repo: string
  readonly policies?: iam.IManagedPolicy[]
}

export class P6CDKGithubOidcProvider extends Construct {
  public readonly roleArn: string

  constructor(scope: Construct, id: string, props: IP6CDKGithubOidcProviderProps) {
    super(scope, id)

    let oidcProvider = this.node.tryFindChild('GithubOidc') as iam.OpenIdConnectProvider
    if (!oidcProvider) {
      oidcProvider = new iam.OpenIdConnectProvider(this, 'GithubOidc', {
        url: 'https://token.actions.githubusercontent.com',
        clientIds: ['sts.amazonaws.com'],
        thumbprints: ['ffffffffffffffffffffffffffffffffffffffff'],
      })
    }

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

    // Attach provided policies to the role
    if (props.policies) {
      props.policies.forEach(policy => role.addManagedPolicy(policy))
    }

    // Expose the role ARN
    this.roleArn = role.roleArn
  }
}
