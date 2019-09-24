package jadepoolsaas

import "errors"

// NewCompany creates a new company with key and secret.
func NewCompany(key, secret string) *Company {
	return NewCompanyWithAddr(defaultAddr, key, secret)
}

// NewCompanyWithAddr creates a new company with server addr, key and secret.
func NewCompanyWithAddr(addr, key, secret string) *Company {
	a := &Company{
		Addr:   addr,
		Key:    key,
		Secret: secret,
	}
	a.session = &session{client: a}
	return a
}

// GetFundingWallets get all wallets in the company.
func (c *Company) GetFundingWallets() (*Result, error) {
	return c.session.get("/api/v1/funding/balances")
}

// FundingTransfer transfer the funding between wallets.
func (c *Company) FundingTransfer(from, to, coinType, value string) (*Result, error) {
	return c.FundingTransferWithMemo(from, to, coinType, value, "")
}

// FundingTransferWithMemo transfer the funding between wallets for the coin with specified memo.
func (c *Company) FundingTransferWithMemo(from, to, coinType, value, memo string) (*Result, error) {
	if len(coinType) == 0 || len(from) == 0 || len(to) == 0 || len(value) == 0 {
		return nil, errors.New("from or coinType or to or value is empty")
	}

	return c.session.post("/api/v1/funding/transfer", map[string]interface{}{
		"from":      from,
		"to":        to,
		"value":     value,
		"assetName": coinType,
		"memo":      memo,
	})
}

// GetFundingRecords get funding records.
func (c *Company) GetFundingRecords(page, amount int) (*Result, error) {
	return c.FilterFundingRecords(page, amount, "DESC", "", "", "", "", "created_at")
}

// FilterFundingRecords get funding records with filters.
func (c *Company) FilterFundingRecords(page, amount int, sort, coins, froms, toes, coinType, orderBy string) (*Result, error) {
	if len(sort) == 0 || len(orderBy) == 0 {
		return nil, errors.New("sort or orderBy is empty")
	}

	if page <= 0 {
		page = 1
	}
	if amount <= 0 {
		amount = 10
	}

	return c.session.getWithParams("/api/v1/funding/records", map[string]interface{}{
		"page":    page,
		"amount":  amount,
		"sort":    sort,
		"coins":   coins,
		"froms":   froms,
		"toes":    toes,
		"type":    coinType,
		"orderBy": orderBy,
	})
}

// Company represent a company.
type Company struct {
	Addr   string
	Key    string
	Secret string

	session *session
}

func (c *Company) getKey() string {
	return c.Key
}

func (c *Company) getKeyHeaderName() string {
	return "X-Company-Key"
}

func (c *Company) getSecret() string {
	return c.Secret
}

func (c *Company) getAddr() string {
	return c.Addr
}
