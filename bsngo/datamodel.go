package bsngo

import (
	"errors"
	"strings"
)

// Data
type BaseModel struct {
	BaseKey   string
	BaseValue string
}

func (b BaseModel) Check() (err error) {
	if strings.TrimSpace(b.BaseKey) == "" {
		err = errors.New("base key can't be empty")
	}
	return
}

func GetMainKey(baseKey string) string {
	return preKey + baseKey
}

// History Info
type DTOHistoryModel struct {
	TxId      string `json:"txId"`
	Value     string `json:"dataInfo"`
	Timestamp string `json:"txTime"`
	IsDelete  bool   `json:"isDelete"`
}
