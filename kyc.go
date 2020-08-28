package jadepoolsaas

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

// FileGet get file.
func (k *KYC) FileGet(fileID, filePath string) (*Result, error) {
	return k.session.getFile("/api/v1/file/" + fileID, filePath)
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
