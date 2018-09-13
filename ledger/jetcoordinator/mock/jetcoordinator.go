/*
 *    Copyright 2018 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package mock

import (
	"github.com/insolar/insolar/configuration"
	"github.com/insolar/insolar/core"
)

type mockJetCoordinator struct {
	virtualExecutor core.RecordRef
	lightExecutor   core.RecordRef
	heavyExecutor   core.RecordRef

	virtualValidators []core.RecordRef
	lightValidators   []core.RecordRef
}

func NewMockJetCoordinator(conf configuration.JetCoordinator) (core.JetCoordinator, error) {
	virtualExecutor := core.String2Ref(conf.VirtualExecutor)
	lightExecutor := core.String2Ref(conf.LightExecutor)
	heavyExecutor := core.String2Ref(conf.HeavyExecutor)

	virtualValidators := make([]core.RecordRef, len(conf.VirtualValidators))
	for i, vv := range conf.VirtualValidators {
		virtualValidators[i] = core.String2Ref(vv)
	}

	lightValidators := make([]core.RecordRef, len(conf.LightValidators))
	for i, lv := range conf.VirtualValidators {
		lightValidators[i] = core.String2Ref(lv)
	}

	return &mockJetCoordinator{
		virtualExecutor: virtualExecutor,
		lightExecutor:   lightExecutor,
		heavyExecutor:   heavyExecutor,

		virtualValidators: virtualValidators,
		lightValidators:   lightValidators,
	}, nil
}

func (mockJetCoordinator) IsAuthorized(role core.JetRole, obj core.RecordRef, pulse core.PulseNumber, node core.RecordRef) bool {
	panic("implement me")
}

func (mockJetCoordinator) QueryRole(role core.JetRole, obj core.RecordRef, pulse core.PulseNumber) []core.RecordRef {
	panic("implement me")
}
