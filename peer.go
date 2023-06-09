package ssi

import "ssi/component"

// peer represent ssi system roles: issuer, holder and verifier

type Peer struct {
	Agent  component.Agent
	Wallet component.Wallet
}
