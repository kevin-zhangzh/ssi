package component

import "ssi/schema"

type Wallet struct {
	DidMap map[string]schema.DIDDoc
	VcMap  map[string]schema.Credential
}

func (w *Wallet) GetDID(did string) (ok bool, doc schema.DIDDoc) {
	return
}

func (w *Wallet) SaveDID(doc schema.DIDDoc) {

}

func (w *Wallet) GetCredential(id string) (ok bool, vc schema.Credential) {
	return
}

func (w *Wallet) SaveCredential(vc schema.Credential) {

}
