# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### P6CDKGithubOidcProvider <a name="P6CDKGithubOidcProvider" id="p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider"></a>

#### Initializers <a name="Initializers" id="p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.Initializer"></a>

```typescript
import { P6CDKGithubOidcProvider } from 'p6-cdk-github-oidc-provider'

new P6CDKGithubOidcProvider(scope: Construct, id: string)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.Initializer.parameter.id"></a>

- *Type:* string

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### `isConstruct` <a name="isConstruct" id="p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.isConstruct"></a>

```typescript
import { P6CDKGithubOidcProvider } from 'p6-cdk-github-oidc-provider'

P6CDKGithubOidcProvider.isConstruct(x: any)
```

Checks if `x` is a construct.

Use this method instead of `instanceof` to properly detect `Construct`
instances, even when the construct library is symlinked.

Explanation: in JavaScript, multiple copies of the `constructs` library on
disk are seen as independent, completely different libraries. As a
consequence, the class `Construct` in each copy of the `constructs` library
is seen as a different class, and an instance of one class will not test as
`instanceof` the other class. `npm install` will not create installations
like this, but users may manually symlink construct libraries together or
use a monorepo tool: in those cases, multiple copies of the `constructs`
library can be accidentally installed, and `instanceof` will behave
unpredictably. It is safest to avoid using `instanceof`, and using
this type-testing method instead.

###### `x`<sup>Required</sup> <a name="x" id="p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.property.openIdConnectProviderArn">openIdConnectProviderArn</a></code> | <code>aws-cdk-lib.Arn</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `openIdConnectProviderArn`<sup>Required</sup> <a name="openIdConnectProviderArn" id="p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.property.openIdConnectProviderArn"></a>

```typescript
public readonly openIdConnectProviderArn: Arn;
```

- *Type:* aws-cdk-lib.Arn

---





