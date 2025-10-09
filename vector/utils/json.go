// Copyright (c) 2025 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ParseJsonUseNumber2(input []byte, target interface{}) error {
	var d *json.Decoder
	var err error
	d = json.NewDecoder(bytes.NewBuffer(input))
	if d == nil {
		return fmt.Errorf("ParseJsonUseNumber init NewDecoder failed")
	}
	d.UseNumber()
	err = d.Decode(&target)
	if err != nil {
		return fmt.Errorf("ParseJsonUseNumber Decode failed %v", err)
	}
	return nil
}

func SerilizeToJsonBytesUseNumber(source interface{}) ([]byte, error) {
	//  buffer := make([]byte, 0)
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(source)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
