# State

## FanToken
Definition of data structure of Fungible Token

```go
type FanToken struct {
	Denom		string
	Name		string
	MaxSupply	sdk.Int
	Mintable	bool
	MetadataUri string
	Owner		string
}
```

## Params
Params is a module-wide configuration structure that stores system parameters and defines overall functioning of the fan token module.

```go
type Params struct {
	IssuePrice	sdk.Coin
}
```