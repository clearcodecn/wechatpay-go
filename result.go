// Copyright The Wechat Pay Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package wechatpay

import "encoding/json"

// Result is a result after call client.Do
type Result struct {
	Body      []byte
	Timestamp int64
	Nonce     string
	Signature string
	SerialNo  string
	Err       error
}

// Scan data from the response into the dest object.
func (r *Result) Scan(dest interface{}) error {
	if r.Error() != nil {
		return r.Err
	}

	if len(r.Body) == 0 {
		return nil
	}

	if err := json.Unmarshal(r.Body, dest); err != nil {
		return err
	}

	return nil
}

// Error return the error.
func (r *Result) Error() error {
	return r.Err
}
