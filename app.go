package jadepoolsaas

import "errors"

// NewApp creates a new wallet with key and secret.
func NewApp(appKey, appSecret string) *App {
	return NewAppWithAddr(defaultAddr, appKey, appSecret)
}

// NewAppWithAddr creates a new wallet with server addr, key and secret.
func NewAppWithAddr(addr, appKey, appSecret string) *App {
	a := &App{
		Addr:   addr,
		Key:    appKey,
		Secret: appSecret,
	}
	a.session = &session{client: a}
	return a
}

// CreateAddress request new address.
func (a *App) CreateAddress(coinType string) (*Result, error) {
	return a.CreateAddressWithMode(coinType, "")
}

// CreateAddressWithMode request new address for the coin with specified memo.
func (a *App) CreateAddressWithMode(coinType, mode string) (*Result, error) {
	if len(coinType) == 0 {
		return nil, errors.New("coinType is empty")
	}

	return a.session.post("/api/v1/address/"+coinType+"/new", map[string]interface{}{
		"mode": mode,
	})
}

// VerifyAddress verify if an address is valid for specified coin.
func (a *App) VerifyAddress(coinType, address string) (*Result, error) {
	if len(coinType) == 0 || len(address) == 0 {
		return nil, errors.New("coinType or address is empty")
	}

	return a.session.post("/api/v1/address/"+coinType+"/verify", map[string]interface{}{
		"address": address,
	})
}

// GetAssets fetch all assets in the wallet.
func (a *App) GetAssets() (*Result, error) {
	return a.session.get("/api/v1/app/assets")
}

// GetBalance get the balance for specified coin.
func (a *App) GetBalance(coinType string) (*Result, error) {
	if len(coinType) == 0 {
		return nil, errors.New("coinType is empty")
	}

	return a.session.get("/api/v1/app/balance/" + coinType)
}

// GetOrder get order by id.
func (a *App) GetOrder(id string) (*Result, error) {
	if len(id) == 0 {
		return nil, errors.New("id is empty")
	}

	return a.session.get("/api/v1/app/order/" + id)
}

// Withdraw request withdrawal.
func (a *App) Withdraw(id, coinType, to, value string) (*Result, error) {
	return a.WithdrawWithMemo(id, coinType, to, value, "")
}

// WithdrawWithMemo request withdrawal for the coin with specified memo.
func (a *App) WithdrawWithMemo(id, coinType, to, value, memo string) (*Result, error) {
	if len(coinType) == 0 || len(id) == 0 || len(to) == 0 || len(value) == 0 {
		return nil, errors.New("id or coinType or to or value is empty")
	}

	return a.session.post("/api/v1/app/"+coinType+"/withdraw", map[string]interface{}{
		"to":    to,
		"value": value,
		"memo":  memo,
		"id":    id,
	})
}

// Delegate request delegation.
func (a *App) Delegate(id, coinType, value string) (*Result, error) {
	if len(coinType) == 0 || len(id) == 0 || len(value) == 0 {
		return nil, errors.New("id or coinType or value is empty")
	}

	return a.session.post("/api/v1/staking/"+coinType+"/delegate", map[string]interface{}{
		"value": value,
		"id":    id,
	})
}

// UnDelegate request undelegation.
func (a *App) UnDelegate(id, coinType, value string) (*Result, error) {
	if len(coinType) == 0 || len(id) == 0 || len(value) == 0 {
		return nil, errors.New("id or coinType or value is empty")
	}

	return a.session.post("/api/v1/staking/"+coinType+"/undelegate", map[string]interface{}{
		"value": value,
		"id":    id,
	})
}

// GetValidators fetch all validators of specified coin.
func (a *App) GetValidators(coinType string) (*Result, error) {
	if len(coinType) == 0 {
		return nil, errors.New("coinType is empty")
	}

	return a.session.get("/api/v1/staking/" + coinType + "/validators")
}

// AddUrgentStakingFunding add urgent staking funding.
func (a *App) AddUrgentStakingFunding(id, coinType, value string, expiredAt int64) (*Result, error) {
	if len(coinType) == 0 || len(id) == 0 || len(value) == 0 {
		return nil, errors.New("id or coinType or value is empty")
	}

	return a.session.post("/api/v1/staking/"+coinType+"/funding", map[string]interface{}{
		"value":     value,
		"id":        id,
		"expiredAt": expiredAt,
	})
}

// App represent a wallet.
type App struct {
	Addr   string
	Key    string
	Secret string

	session *session
}

func (a *App) getKey() string {
	return a.Key
}

func (a *App) getKeyHeaderName() string {
	return "X-App-Key"
}

func (a *App) getSecret() string {
	return a.Secret
}

func (a *App) getAddr() string {
	return a.Addr
}
