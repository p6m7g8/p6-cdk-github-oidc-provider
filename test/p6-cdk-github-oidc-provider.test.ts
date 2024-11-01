import * as cdk from 'aws-cdk-lib'
import { Template } from 'aws-cdk-lib/assertions'
import { P6CDKGithubOidcProvider } from '../src'

// Test: Snapshot test for GithubOidcProvider Construct
it('snapshot test for GithubOidcProvider', () => {
  // Set up a CDK app and stack
  const app = new cdk.App()
  const stack = new cdk.Stack(app, 'TestStack')

  // Instantiate the GithubOidcProvider construct
  new P6CDKGithubOidcProvider(stack, 'GithubOidcTest')

  // Generate the CloudFormation template for the stack
  const template = Template.fromStack(stack)

  // Expect the generated template to match the snapshot
  expect(template.toJSON()).toMatchSnapshot()
})
