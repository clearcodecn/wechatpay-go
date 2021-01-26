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

package sign

import (
	"testing"
)

func TestAes256Gcm(t *testing.T) {
	cases := []struct {
		key  []byte
		noce []byte
		data []byte
		text string
	}{
		{
			[]byte("AES256Key-32Characters1234567890"),
			[]byte("eabb3e044577"),
			[]byte("certificate"),
			"exampleplaintext",
		},
	}

	for _, c := range cases {
		secret, err := EncryptByAes256Gcm(c.key, c.noce, c.data, c.text)
		if err != nil {
			t.Fatal(err)
		}

		plain, err := DecryptByAes256Gcm(c.key, c.noce, c.data, secret)
		if err != nil {
			t.Fatal(err)
		}

		plainTxt := string(plain)
		if plainTxt != c.text {
			t.Fatal("invalid aes-256-gcm")
		}
	}
}
