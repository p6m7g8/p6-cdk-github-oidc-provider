# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### P6CDKGithubOidcProvider <a name="P6CDKGithubOidcProvider" id="p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider"></a>

#### Initializers <a name="Initializers" id="p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.Initializer"></a>

```typescript
import { P6CDKGithubOidcProvider } from 'p6-cdk-github-oidc-provider'

new P6CDKGithubOidcProvider(scope: Construct, id: string, props: IP6CDKGithubOidcProviderProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.Initializer.parameter.props">props</a></code> | <code><a href="#p6-cdk-github-oidc-provider.IP6CDKGithubOidcProviderProps">IP6CDKGithubOidcProviderProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.Initializer.parameter.props"></a>

- *Type:* <a href="#p6-cdk-github-oidc-provider.IP6CDKGithubOidcProviderProps">IP6CDKGithubOidcProviderProps</a>

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
| <code><a href="#p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.property.roleArn">roleArn</a></code> | <code>string</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `roleArn`<sup>Required</sup> <a name="roleArn" id="p6-cdk-github-oidc-provider.P6CDKGithubOidcProvider.property.roleArn"></a>

```typescript
public readonly roleArn: string;
```

- *Type:* string

---




## Protocols <a name="Protocols" id="Protocols"></a>

### IP6CDKGithubOidcProviderProps <a name="IP6CDKGithubOidcProviderProps" id="p6-cdk-github-oidc-provider.IP6CDKGithubOidcProviderProps"></a>

- *Implemented By:* <a href="#p6-cdk-github-oidc-provider.IP6CDKGithubOidcProviderProps">IP6CDKGithubOidcProviderProps</a>


#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#p6-cdk-github-oidc-provider.IP6CDKGithubOidcProviderProps.property.repo">repo</a></code> | <code>string</code> | *No description.* |
| <code><a href="#p6-cdk-github-oidc-provider.IP6CDKGithubOidcProviderProps.property.policies">policies</a></code> | <code>aws-cdk-lib.aws_iam.IManagedPolicy[]</code> | *No description.* |

---

##### `repo`<sup>Required</sup> <a name="repo" id="p6-cdk-github-oidc-provider.IP6CDKGithubOidcProviderProps.property.repo"></a>

```typescript
public readonly repo: string;
```

- *Type:* string

---

##### `policies`<sup>Optional</sup> <a name="policies" id="p6-cdk-github-oidc-provider.IP6CDKGithubOidcProviderProps.property.policies"></a>

```typescript
public readonly policies: IManagedPolicy[];
```

- *Type:* aws-cdk-lib.aws_iam.IManagedPolicy[]

---

