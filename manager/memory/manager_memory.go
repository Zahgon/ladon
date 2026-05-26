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

package memory

import (
	"context"
	"sync"

	. "github.com/ory/ladon"
)

// MemoryManager is an in-memory (non-persistent) implementation of Manager.
type MemoryManager struct {
	Policies map[string]Policy
	sync.RWMutex
}

// NewMemoryManager constructs and initializes new MemoryManager with no policies.
func NewMemoryManager() *MemoryManager { _ = "STUB: not implemented"; return nil }

// Update updates an existing policy.
func (m *MemoryManager) Update(ctx context.Context, policy Policy) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAll returns all policies.
func (m *MemoryManager) GetAll(ctx context.Context, limit, offset int64) (Policies, error) {
	_ = "STUB: not implemented"
	return *new(Policies), nil
}

// Create a new pollicy to MemoryManager.
func (m *MemoryManager) Create(ctx context.Context, policy Policy) error {
	_ = "STUB: not implemented"
	return nil
}

// Get retrieves a policy.
func (m *MemoryManager) Get(ctx context.Context, id string) (Policy, error) {
	_ = "STUB: not implemented"
	return *new(Policy), nil
}

// Delete removes a policy.
func (m *MemoryManager) Delete(ctx context.Context, id string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MemoryManager) findAllPolicies() (Policies, error) {
	_ = "STUB: not implemented"
	return *new(Policies), nil
}

// FindRequestCandidates returns candidates that could match the request object. It either returns
// a set that exactly matches the request, or a superset of it. If an error occurs, it returns nil and
// the error.
func (m *MemoryManager) FindRequestCandidates(ctx context.Context, r *Request) (Policies, error) {
	_ = "STUB: not implemented"
	return *new(Policies), nil
}

// FindPoliciesForSubject returns policies that could match the subject. It either returns
// a set of policies that applies to the subject, or a superset of it.
// If an error occurs, it returns nil and the error.
func (m *MemoryManager) FindPoliciesForSubject(ctx context.Context, subject string) (Policies, error) {
	_ = "STUB: not implemented"
	return *new(Policies), nil
}

// FindPoliciesForResource returns policies that could match the resource. It either returns
// a set of policies that apply to the resource, or a superset of it.
// If an error occurs, it returns nil and the error.
func (m *MemoryManager) FindPoliciesForResource(ctx context.Context, resource string) (Policies, error) {
	_ = "STUB: not implemented"
	return *new(Policies), nil
}
