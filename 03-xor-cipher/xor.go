package main

import (
	"bytes"
	"encoding/gob"
)

//func main (){}

func gobEncode(value interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(value)
	return buf.Bytes(), err
}

func gobDecode(data []byte, value interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(value)
}

func xorCipher(text string, key string) ([]byte, error) {
	textBuf, err := gobEncode(text)
	if err != nil {
		return nil, err
	}

	keyBuf, err := gobEncode(key)
	if err != nil {
		return nil, err
	}

	xoredBuf := make([]byte, len(textBuf))
	for i, c := range textBuf {
		xoredBuf[i] = c ^ keyBuf[i%len(keyBuf)]
	}

	return xoredBuf, nil
}

func xorDecipher(cipherText []byte, key string) (string, error) {
	keyBuf, err := gobEncode(key)
	if err != nil {
		return "", err
	}

	decipheredBuf := make([]byte, len(cipherText))
	for i, c := range cipherText {
		decipheredBuf[i] = c ^ keyBuf[i%len(keyBuf)]
	}

	var decipheredText string
	err = gobDecode(decipheredBuf, &decipheredText)
	if err != nil {
		return "", err
	}

	return decipheredText, nil
}

