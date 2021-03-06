// Copyright (c) 2014 - Max Ekman <max@looplab.se>
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

package local

import (
	"testing"

	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/eventbus/testutil"
)

func TestEventBus(t *testing.T) {
	bus := NewEventBus()
	if bus == nil {
		t.Fatal("there should be a bus")
	}

	testutil.EventBusCommonTests(t, bus, bus)
}

func TestEventBusAsync(t *testing.T) {
	bus := NewEventBus()
	if bus == nil {
		t.Fatal("there should be a bus")
	}
	bus.SetHandlingStrategy(eh.AsyncEventHandlingStrategy)

	testutil.EventBusCommonTests(t, bus, bus)
}
