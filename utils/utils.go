/*
Package utils ...

Copyright © 2019 hajime-terasawa <terako.studio@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package utils

import (
	"math/rand"
)

// Contains assert an array includes an element or not.
func Contains(s []interface{}, e interface{}) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}

	return false
}

// Shuffle extracts elements from slice by random indexing.
func Shuffle(s []interface{}) interface{} {
	//nolint:gosec
	idx := rand.Intn(len(s))
	return s[idx]
}
