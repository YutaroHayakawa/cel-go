// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package interpreter

import (
	"github.com/google/cel-go/common/types/ref"
)

// EvalState tracks the values associated with expression ids during execution.
type EvalState interface {
	// Value returns the observed value of the given expression id if found, and a nil false
	// result if not.
	Value(int64) (ref.Value, bool)

	// SetValue sets the observed value of the expression id.
	SetValue(int64, ref.Value)

	// Reset clears the previously recorded expression values.
	Reset()
}

// evalState permits the mutation of evaluation state for a given expression id.
type evalState struct {
	values map[int64]ref.Value
}

// NewEvalState returns an EvalState instanced used to observe the intermediate
// evaluations of an expression.
func NewEvalState() EvalState {
	return &evalState{
		values: make(map[int64]ref.Value),
	}
}

// Value is an implementation of the EvalState interface method.
func (s *evalState) Value(exprID int64) (ref.Value, bool) {
	val, found := s.values[exprID]
	return val, found
}

// SetValue is an implementation of the EvalState interface method.
func (s *evalState) SetValue(exprID int64, val ref.Value) {
	s.values[exprID] = val
}

func (s *evalState) Reset() {
	s.values = map[int64]ref.Value{}
}
