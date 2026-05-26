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

// Policies is an array of policies.
type Policies []Policy

// Policy represent a policy model.
type Policy interface {
	// GetID returns the policies id.
	GetID() string

	// GetDescription returns the policies description.
	GetDescription() string

	// GetSubjects returns the policies subjects.
	GetSubjects() []string

	// AllowAccess returns true if the policy effect is allow, otherwise false.
	AllowAccess() bool

	// GetEffect returns the policies effect which might be 'allow' or 'deny'.
	GetEffect() string

	// GetResources returns the policies resources.
	GetResources() []string

	// GetActions returns the policies actions.
	GetActions() []string

	// GetConditions returns the policies conditions.
	GetConditions() Conditions

	// GetMeta returns the policies arbitrary metadata set by the user.
	GetMeta() []byte

	// GetStartDelimiter returns the delimiter which identifies the beginning of a regular expression.
	GetStartDelimiter() byte

	// GetEndDelimiter returns the delimiter which identifies the end of a regular expression.
	GetEndDelimiter() byte
}

// DefaultPolicy is the default implementation of the policy interface.
type DefaultPolicy struct {
	ID          string     `json:"id" gorethink:"id"`
	Description string     `json:"description" gorethink:"description"`
	Subjects    []string   `json:"subjects" gorethink:"subjects"`
	Effect      string     `json:"effect" gorethink:"effect"`
	Resources   []string   `json:"resources" gorethink:"resources"`
	Actions     []string   `json:"actions" gorethink:"actions"`
	Conditions  Conditions `json:"conditions" gorethink:"conditions"`
	Meta        []byte     `json:"meta" gorethink:"meta"`
}

// UnmarshalJSON overwrite own policy with values of the given in policy in JSON format
func (p *DefaultPolicy) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalMeta parses the policies []byte encoded metadata and stores the result in the value pointed to by v.
func (p *DefaultPolicy) UnmarshalMeta(v interface{}) error { _ = "STUB: not implemented"; return nil }

// GetID returns the policies id.
func (p *DefaultPolicy) GetID() string {
	_ = "STUB: not implemented"

	// GetDescription returns the policies description.
	return ""
}

func (p *DefaultPolicy) GetDescription() string { _ = "STUB: not implemented"; return "" }

// GetSubjects returns the policies subjects.
func (p *DefaultPolicy) GetSubjects() []string {
	_ = "STUB: not implemented"

	// AllowAccess returns true if the policy effect is allow, otherwise false.
	return nil
}

func (p *DefaultPolicy) AllowAccess() bool { _ = "STUB: not implemented"; return false }

// GetEffect returns the policies effect which might be 'allow' or 'deny'.
func (p *DefaultPolicy) GetEffect() string {
	_ = "STUB: not implemented"

	// GetResources returns the policies resources.
	return ""
}

func (p *DefaultPolicy) GetResources() []string {
	_ = "STUB: not implemented"

	// GetActions returns the policies actions.
	return nil
}

func (p *DefaultPolicy) GetActions() []string {
	_ = "STUB: not implemented"

	// GetConditions returns the policies conditions.
	return nil
}

func (p *DefaultPolicy) GetConditions() Conditions {
	_ = "STUB: not implemented"
	return *

	// GetMeta returns the policies arbitrary metadata set by the user.
	new(Conditions)
}

func (p *DefaultPolicy) GetMeta() []byte {
	_ = "STUB: not implemented"

	// GetEndDelimiter returns the delimiter which identifies the end of a regular expression.
	return nil
}

func (p *DefaultPolicy) GetEndDelimiter() byte {
	_ = "STUB: not implemented"

	// GetStartDelimiter returns the delimiter which identifies the beginning of a regular expression.
	return 0
}

func (p *DefaultPolicy) GetStartDelimiter() byte { _ = "STUB: not implemented"; return 0 }
