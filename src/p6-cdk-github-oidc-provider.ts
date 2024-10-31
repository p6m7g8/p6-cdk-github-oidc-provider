import * as iam from 'aws-cdk-lib/aws-iam'
import { Construct } from 'constructs'

export class P6CDKGithubOidcProvider extends Construct {
  constructor(scope: Construct, id: string) {
    super(scope, id)

    new iam.OpenIdConnectProvider(this, 'GithubOidc', {
      url: 'https://token.actions.githubusercontent.com',
      clientIds: ['sts.amazonaws.com'],
      thumbprints: ['ffffffffffffffffffffffffffffffffffffffff'],
    })
  }
}
