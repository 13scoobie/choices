// Copyright 2016 Andrew O'Neill, Nordstrom

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package choices

import (
	"fmt"
	"testing"
)

func TestValueUniform(t *testing.T) {
}

func TestValueWeighted(t *testing.T) {
	w := &Weighted{Choices: []weightedChoice{{"a", 1}, {"b", 2}, {"c", 1}}}
	tests := []struct {
		in   uint64
		want string
	}{
		{
			in:   0,
			want: "a",
		},
		{
			in:   0x8000000000000000,
			want: "b",
		},
		{
			in:   0xffffffffffffffff,
			want: "c",
		},
	}

	for _, test := range tests {
		got, err := w.Choice(test.in)
		if err != nil {
			fmt.Println(err)
		}
		if got != test.want {
			t.Errorf("%v.Value(%v) = %v, want %v", *w, test.in, got, test.want)
		}
	}
}
