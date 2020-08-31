package jadepoolsaas

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/imroc/req"
)

type params req.Param

// Result request result.
type Result struct {
	Code    int
	Data    map[string]interface{}
	Message string
	Sign    string
}

type session struct {
	client     client
	nonceCount int
}

func (session *session) get(path string) (*Result, error) {
	return session.getWithParams(path, map[string]interface{}{})
}

func (session *session) getWithParams(path string, params params) (*Result, error) {
	url := session.getURL(path)
	err := session.prepareParams(params)
	if err != nil {
		return nil, err
	}

	r, err := req.Get(url, session.commonHeaders(), req.Param(params))
	if err != nil {
		return nil, err
	}
	if r.Response().StatusCode != 200 {
		return nil, fmt.Errorf("http error code:%d", r.Response().StatusCode)
	}

	var result Result
	err = r.ToJSON(&result)
	if err != nil {
		return nil, fmt.Errorf("parse body to json failed: %v", err)
	}

	if err = result.error(session.client.getSecret()); err != nil {
		return nil, err
	}

	return &result, err
}

func (session *session) getFile(path string, filePath string) (*Result, error) {
	params := params{}

	url := session.getURL(path)
	err := session.prepareParams(params)
	if err != nil {
		return nil, err
	}

	r, err := req.Get(url, session.commonHeaders(), req.Param(params))
	if err != nil {
		return nil, err
	}
	if r.Response().StatusCode != 200 {
		return nil, fmt.Errorf("http error code:%d", r.Response().StatusCode)
	}

	err = r.ToFile(filePath)
	if err != nil {
		return nil, err
	}

	var result Result
	result.Data = map[string]interface{}{
		"filePath": filePath,
	}
	return &result, err
}

func (session *session) post(path string, params params) (*Result, error) {
	url := session.getURL(path)
	err := session.prepareParams(params)
	if err != nil {
		return nil, err
	}

	r, err := req.Post(url, session.commonHeaders(), req.BodyJSON(&params))
	if err != nil {
		return nil, err
	}
	if r.Response().StatusCode != 200 {
		return nil, fmt.Errorf("http error code:%d", r.Response().StatusCode)
	}

	var result Result
	err = r.ToJSON(&result)
	if err != nil {
		return nil, fmt.Errorf("parse body to json failed: %v", err)
	}

	if err = result.error(session.client.getSecret()); err != nil {
		return nil, err
	}

	return &result, err
}

func (session *session) patch(path string, params params) (*Result, error) {
	url := session.getURL(path)
	err := session.prepareParams(params)
	if err != nil {
		return nil, err
	}

	r, err := req.Patch(url, session.commonHeaders(), req.BodyJSON(&params))
	if err != nil {
		return nil, err
	}
	if r.Response().StatusCode != 200 {
		return nil, fmt.Errorf("http error code:%d", r.Response().StatusCode)
	}

	var result Result
	err = r.ToJSON(&result)
	if err != nil {
		return nil, fmt.Errorf("parse body to json failed: %v", err)
	}

	if err = result.error(session.client.getSecret()); err != nil {
		return nil, err
	}

	return &result, err
}

func (session *session) postFile(path string, filePath string) (*Result, error) {
	params := params{}

	url := session.getURL(path)
	err := session.prepareParams(params)
	if err != nil {
		return nil, err
	}

	r, err := req.Post(url, session.commonHeaders(), req.File(filePath), req.Param(params))
	if err != nil {
		return nil, err
	}
	if r.Response().StatusCode != 200 {
		return nil, fmt.Errorf("http error code:%d", r.Response().StatusCode)
	}

	var result Result
	err = r.ToJSON(&result)
	if err != nil {
		return nil, fmt.Errorf("parse body to json failed: %v", err)
	}

	if err = result.error(session.client.getSecret()); err != nil {
		return nil, err
	}

	return &result, err
}

func (session *session) put(path string, params params) (*Result, error) {
	url := session.getURL(path)
	err := session.prepareParams(params)
	if err != nil {
		return nil, err
	}

	r, err := req.Put(url, session.commonHeaders(), req.BodyJSON(&params))
	if err != nil {
		return nil, err
	}
	if r.Response().StatusCode != 200 {
		return nil, fmt.Errorf("http error code:%d", r.Response().StatusCode)
	}

	var result Result
	err = r.ToJSON(&result)
	if err != nil {
		return nil, fmt.Errorf("parse body to json failed: %v", err)
	}

	if err = result.error(session.client.getSecret()); err != nil {
		return nil, err
	}

	return &result, err
}

func (session *session) delete(path string) (*Result, error) {
	return session.deleteWithParams(path, map[string]interface{}{})
}

func (session *session) deleteWithParams(path string, params params) (*Result, error) {
	url := session.getURL(path)
	err := session.prepareParams(params)
	if err != nil {
		return nil, err
	}

	r, err := req.Delete(url, session.commonHeaders(), req.QueryParam(params))
	if err != nil {
		return nil, err
	}
	if r.Response().StatusCode != 200 {
		return nil, fmt.Errorf("http error code:%d", r.Response().StatusCode)
	}

	var result Result
	err = r.ToJSON(&result)
	if err != nil {
		return nil, fmt.Errorf("parse body to json failed: %v", err)
	}

	if err = result.error(session.client.getSecret()); err != nil {
		return nil, err
	}

	return &result, err
}

func (session *session) getURL(path string) string {
	return fmt.Sprintf("%s%s", session.client.getAddr(), path)
}

func (session *session) prepareParams(params params) error {
	timestamp := time.Now().Unix()
	params["timestamp"] = timestamp
	params["nonce"] = session.genNonce(timestamp)
	return params.sign(session.client.getSecret())
}

func (session *session) commonHeaders() req.Header {
	keyName := session.client.getKeyHeaderName()
	return req.Header{
		"Content-Type": "application/json",
		keyName:        session.client.getKey(),
	}
}

func (session *session) genNonce(timestamp int64) string {
	session.nonceCount++
	return fmt.Sprintf("%d%d%d", session.nonceCount, timestamp, rand.Int63n(timestamp))
}

func (params *params) sign(secret string) error {
	sign, err := signHMACSHA256(params, secret)
	if err != nil {
		return err
	}
	(*params)["sign"] = sign
	return nil
}

func (result *Result) error(secret string) error {
	if !result.success() {
		return errors.New(result.Message)
	}

	if !result.checkSign(secret) {
		return errors.New("check sign failed")
	}
	return nil
}

func (result *Result) success() bool {
	return result.Code == 0
}

func (result *Result) checkSign(secret string) bool {
	mySign, err := signHMACSHA256(result.Data, secret)
	if err != nil {
		return false
	}

	return result.Sign == mySign
}
