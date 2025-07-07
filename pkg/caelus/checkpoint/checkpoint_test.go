/*
 * Copyright (c) 2021 Tencent.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 *
 * You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package checkpoint

import (
	"testing"
)

type Data struct {
	A string
	B int
}

func TestCheckpoint(t *testing.T) {
	data := Data{A: "AAA", B: 10}
	InitCheckpointManager("/tmp")
	Save("key1", &data)
	reveiver := &Data{}
	Restore("key1", reveiver)
	if reveiver.A != data.A || reveiver.B != data.B {
		t.Errorf("expect %v, got %v", data, reveiver)
	}
}
