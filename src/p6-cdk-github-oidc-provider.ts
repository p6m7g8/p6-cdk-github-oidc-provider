import type * as cdk from 'aws-cdk-lib'
import * as iam from 'aws-cdk-lib/aws-iam'
import { Construct } from 'constructs'

export class P6CDKGithubOidcProvider extends Construct {
  public readonly openIdConnectProviderArn: cdk.Arn

  constructor(scope: Construct, id: string) {
    super(scope, id)

    const oidcProvider = new iam.OpenIdConnectProvider(this, 'GithubOidc', {
      url: 'https://token.actions.githubusercontent.com',
      clientIds: ['sts.amazonaws.com'],
      thumbprints: ['ffffffffffffffffffffffffffffffffffffffff'],
    })

    this.openIdConnectProviderArn = oidcProvider.openIdConnectProviderArn as cdk.Arn
  }
}
