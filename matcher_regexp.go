/*
 * Copyright © 2016-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package ladon

import (
	"github.com/dlclark/regexp2"
)

func NewRegexpMatcher(size int) *RegexpMatcher { _ = "STUB: not implemented"; return nil }

// golang-lru only returns an error if the cache's size is 0. This, we can safely ignore this error.

type RegexpMatcher struct {
	*lru.Cache
}

func (m *RegexpMatcher) get(pattern string) *regexp2.Regexp { _ = "STUB: not implemented"; return nil }

func (m *RegexpMatcher) set(pattern string, reg *regexp2.Regexp) { _ = "STUB: not implemented"; return }

// Matches a needle with an array of regular expressions and returns true if a match was found.
func (m *RegexpMatcher) Matches(p Policy, haystack []string, needle string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// This means that the current haystack item does not contain a regular expression

// If we have a simple string match, we've got a match!

// Not string match, but also no regexp, continue with next haystack item

// according to regexp2 documentation: https://github.com/dlclark/regexp2#usage
// The only error that the *Match* methods should return is a Timeout if you set the
// re.MatchTimeout field. Any other error is a bug in the regexp2 package.

// according to regexp2 documentation: https://github.com/dlclark/regexp2#usage
// The only error that the *Match* methods should return is a Timeout if you set the
// re.MatchTimeout field. Any other error is a bug in the regexp2 package.
