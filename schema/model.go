package schema

type DIDDoc struct {
	DID string `json:"did"`
}

type Credential struct {
	ID         string `json:"id"`
	Properties []Tags `json:"properties"`
}

type Tags struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
