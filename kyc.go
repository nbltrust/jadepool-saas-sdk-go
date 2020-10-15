package jadepoolsaas

import (
	"bytes"
	"github.com/imroc/req"
)

// NewKYCWithAddr creates a new kyc instance with server addr, key and secret.
func NewKYCWithAddr(addr, appKey, appSecret string) *KYC {
	a := &KYC{
		Addr:   addr,
		Key:    appKey,
		Secret: appSecret,
	}
	a.session = &session{client: a}
	return a
}

// FileUpload upload file.
func (k *KYC) FileUpload(filePath string) (*Result, error) {
	return k.session.postFile("/api/v1/file", filePath)
}

// FileUpload2 upload file.
func (k *KYC) FileUpload2(fileName string, file *bytes.Reader) (*Result, error) {
	return k.session.postFile2("/api/v1/file", fileName, file)
}

// FileGet get file.
func (k *KYC) FileGet(fileID, filePath string) (*Result, error) {
	return k.session.getFile("/api/v1/file/" + fileID, filePath)
}

// FileGet2 get file.
func (k *KYC) FileGet2(fileID string) (*req.Resp, error) {
	return k.session.getFile2("/api/v1/file/" + fileID)
}

// ApplicationCreate create an application.
func (k *KYC) ApplicationCreate(mType, identifier string) (*Result, error) {
	return k.session.post("/api/v1/application", map[string]interface{}{
		"type": mType,
		"identifier": identifier,
	})
}

// ApplicationUpdate update the application.
func (k *KYC) ApplicationUpdate(applicationID, key, value string) (*Result, error) {
	return k.session.patch("/api/v1/application/" + applicationID, map[string]interface{}{
		key: value,
	})
}

// ApplicationUpdate2 update the application.
func (k *KYC) ApplicationUpdate2(applicationID string, content map[string]interface{}) (*Result, error) {
	return k.session.patch("/api/v1/application/" + applicationID, content)
}

// ApplicationGet get the application.
func (k *KYC) ApplicationGet(applicationID string) (*Result, error) {
	return k.session.get("/api/v1/application/" + applicationID)
}

// ApplicationSubmit submit the application.
func (k *KYC) ApplicationSubmit(applicationID string) (*Result, error) {
	ret, err := k.session.get("/api/v1/application/" + applicationID)
	if err != nil {
		return nil, err
	}

	return k.session.put("/api/v1/application/" + applicationID, ret.Data["detail"].(map[string]interface{}))
}

// KYC represent a kyc instance.
type KYC struct {
	Addr   string
	Key    string
	Secret string

	session *session
}

func (k *KYC) getKey() string {
	return k.Key
}

func (k *KYC) getKeyHeaderName() string {
	return "X-API-Key"
}

func (k *KYC) getSecret() string {
	return k.Secret
}

func (k *KYC) getAddr() string {
	return k.Addr
}
