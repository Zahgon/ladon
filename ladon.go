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
	"context"
)

// Ladon is an implementation of Warden.
type Ladon struct {
	Manager     Manager
	Matcher     matcher
	AuditLogger AuditLogger
	Metric      Metric
}

func (l *Ladon) matcher() matcher { _ = "STUB: not implemented"; return *new(matcher) }

func (l *Ladon) auditLogger() AuditLogger { _ = "STUB: not implemented"; return *new(AuditLogger) }

func (l *Ladon) metric() Metric { _ = "STUB: not implemented"; return *new(Metric) }

// IsAllowed returns nil if subject s has permission p on resource r with context c or an error otherwise.
func (l *Ladon) IsAllowed(ctx context.Context, r *Request) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Although the manager is responsible of matching the policies, it might decide to just scan for
// subjects, it might return all policies, or it might have a different pattern matching than Golang.
// Thus, we need to make sure that we actually matched the right policies.

// DoPoliciesAllow returns nil if subject s has permission p on resource r with context c for a given policy list or an error otherwise.
// The IsAllowed interface should be preferred since it uses the manager directly. This is a lower level interface for when you don't want to use the ladon manager.
func (l *Ladon) DoPoliciesAllow(ctx context.Context, r *Request, policies []Policy) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Iterate through all policies

// Does the action match with one of the policies?
// This is the first check because usually actions are a superset of get|update|delete|set
// and thus match faster.

// no, continue to next policy

// Does the subject match with one of the policies?
// There are usually less subjects than resources which is why this is checked
// before checking for resources.

// no, continue to next policy

// Does the resource match with one of the policies?

// no, continue to next policy

// Are the policies conditions met?
// This is checked first because it usually has a small complexity.

// no, continue to next policy

// Is the policy's effect `deny`? If yes, this overrides all allow policies -> access denied.

func (l *Ladon) passesConditions(ctx context.Context, p Policy, r *Request) bool {
	_ = "STUB: not implemented"
	return false
}
