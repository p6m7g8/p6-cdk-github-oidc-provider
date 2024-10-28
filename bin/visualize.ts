import type { Construct } from 'constructs'
import process from 'node:process'
import * as cdk from 'aws-cdk-lib'
import { P6CDKGithubOidcProvider } from '../src'

class VisualizeStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props: cdk.StackProps) {
    super(scope, id, props)

    new P6CDKGithubOidcProvider(this, 'MyWebsite', {
      repo: 'p6m7g8/p6m7g8.org',
    })
  }
}

const theEnv = {
  account: process.env.CDK_DEPLOY_ACCOUNT || process.env.CDK_DEFAULT_ACCOUNT,
  region: process.env.CDK_DEPLOY_REGION || process.env.CDK_DEFAULT_REGION,
}

const app = new cdk.App()
new VisualizeStack(app, 'VisualizeStack', { env: theEnv })
app.synth()
