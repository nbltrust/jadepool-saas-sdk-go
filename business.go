package jadepoolsaas

import (
	"encoding/base64"
	"os"
)

// NewBusinessWithAddr creates a new business instance with server addr, business key, the private key pem file of your service and the public key pem file of xpert.
func NewBusinessWithAddr(addr, businessKey, pemFilePath, pubPemFilePath string) (*Business, error) {
	data, err := os.ReadFile(pemFilePath)
	if err != nil {
		return nil, err
	}

	var pubData []byte
	if len(pubPemFilePath) > 0 {
		pubData, err = os.ReadFile(pubPemFilePath)
		if err != nil {
			return nil, err
		}
	}

	a := &Business{
		Addr:   addr,
		Key:    businessKey,
		Secret: base64.StdEncoding.EncodeToString(data),
	}
	if len(pubData) > 0 {
		a.PubKey = base64.StdEncoding.EncodeToString(pubData)
	}
	a.session = &session{client: a}
	return a, nil
}

// ClientGet fetch client info.
func (b *Business) ClientGet(userID uint) (*BusinessResult, error) {
	return b.session.businessGetWithParams("/api/v1/business/client", map[string]interface{}{
		"userID": userID,
	})
}

// AssetsGet fetch all assets in the wallet.
func (b *Business) AssetsGet() (*BusinessResult, error) {
	return b.session.businessGet("/api/v1/business/assets")
}

// WalletBalancesGet fetch the asset balance in the wallet.
func (b *Business) WalletBalancesGet(userID, assetID uint) (*BusinessResult, error) {
	return b.session.businessGetWithParams("/api/v1/business/wallet/balances", map[string]interface{}{
		"userID":  userID,
		"assetID": assetID,
	})
}

// Business represents a business instance.
type Business struct {
	Addr   string
	Key    string
	Secret string
	PubKey string

	session *session
}

func (b *Business) getKey() string {
	return b.Key
}

func (b *Business) getKeyHeaderName() string {
	return "X-BUSINESS-KEY"
}

func (b *Business) getSecret() string {
	return b.Secret
}

func (b *Business) getPubKey() string {
	return b.PubKey
}

func (b *Business) getAddr() string {
	return b.Addr
}
