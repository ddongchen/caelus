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

package alarm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/tencent/caelus/pkg/caelus/types"
	"github.com/tencent/caelus/pkg/caelus/util"
	"github.com/tencent/caelus/pkg/util/times"
)

var (
	alarmExecutor = "/tmp/test-alarm.sh"
	alarmMsgFile  = "/tmp/alarms"
	alarmIp       = "127.0.0.1"
)

// AlarmTest shows alarm test stuct
type AlarmTest struct {
	describe     string
	alarmConfig  types.AlarmConfig
	alarmMsg     []string
	expectResult *AlarmBody
}

var alarmCases = []AlarmTest{
	{
		describe: "disable alarm",
		alarmConfig: types.AlarmConfig{
			Enable:       false,
			MessageBatch: 1,
			ChannelName:  types.AlarmTypeLocal,
			AlarmChannel: types.AlarmChannel{
				LocalAlarm: &types.LocalAlarm{
					Executor: alarmExecutor,
				},
			},
			IgnoreAlarmWhenSilence: false,
		},
		alarmMsg: []string{
			"alarm test 1",
		},
		expectResult: nil,
	},
	{
		describe: "message batch is 1",
		alarmConfig: types.AlarmConfig{
			Enable:       true,
			MessageBatch: 1,
			MessageDelay: times.Duration(5 * time.Second),
			ChannelName:  types.AlarmTypeLocal,
			AlarmChannel: types.AlarmChannel{
				LocalAlarm: &types.LocalAlarm{
					Executor: alarmExecutor,
				},
			},
			IgnoreAlarmWhenSilence: false,
		},
		alarmMsg: []string{
			"alarm test 1",
		},
		expectResult: &AlarmBody{
			IP:       alarmIp,
			AlarmMsg: []string{},
		},
	},
	{
		describe: "message batch is 2",
		alarmConfig: types.AlarmConfig{
			Enable:       true,
			MessageBatch: 2,
			MessageDelay: times.Duration(5 * time.Second),
			ChannelName:  types.AlarmTypeLocal,
			AlarmChannel: types.AlarmChannel{
				LocalAlarm: &types.LocalAlarm{
					Executor: alarmExecutor,
				},
			},
			IgnoreAlarmWhenSilence: false,
		},
		alarmMsg: []string{
			"alarm test 1",
			"alarm test 2",
		},
		expectResult: &AlarmBody{
			IP:       alarmIp,
			AlarmMsg: []string{},
		},
	},
	{
		describe: "silence mode",
		alarmConfig: types.AlarmConfig{
			Enable:       true,
			MessageBatch: 1,
			MessageDelay: times.Duration(5 * time.Second),
			ChannelName:  types.AlarmTypeLocal,
			AlarmChannel: types.AlarmChannel{
				LocalAlarm: &types.LocalAlarm{
					Executor: alarmExecutor,
				},
			},
			IgnoreAlarmWhenSilence: true,
		},
		alarmMsg: []string{
			"alarm test 1",
		},
		expectResult: nil,
	},
}

// TestSendAlarm tests send alarm function
func TestSendAlarm(t *testing.T) {
	generatLocalAlarmExecutor()
	defer removeAlarmFile()

	util.SetNodeIP(alarmIp)
	for _, ac := range alarmCases {
		t.Logf("testing alarm: %s", ac.describe)
		clearAlarmMsgFile()

		// just set silence mode equal to alarm ignored value
		util.SilenceMode = ac.alarmConfig.IgnoreAlarmWhenSilence

		am := NewManager(&ac.alarmConfig, nil)
		stopCh := make(chan struct{})
		am.Run(stopCh)

		for _, msg := range ac.alarmMsg {
			if ac.expectResult != nil {
				ac.expectResult.AlarmMsg = append(ac.expectResult.AlarmMsg,
					fmt.Sprintf("%s[%s]", msg, time.Now().Format("15:04:05")))
			}
			SendAlarm(msg)
			time.Sleep(1 * time.Second)
		}

		alarmMsg, err := readAlarmMsgFile()
		if err != nil {
			t.Fatalf("read alarm message from file err: %v", err)
		}

		if ac.expectResult != nil {
			var alarmBody AlarmBody
			if err = json.Unmarshal([]byte(alarmMsg), &alarmBody); err != nil {
				t.Fatalf("unexpected alarm message %s", alarmMsg)
			}
			if ac.expectResult.IP != alarmBody.IP || ac.expectResult.Cluster != alarmBody.Cluster ||
				!stringSliceEqual(ac.expectResult.AlarmMsg, alarmBody.AlarmMsg) {
				t.Fatalf("unexpected alarm message %s expect %v", alarmMsg, *ac.expectResult)
			}
		}
	}
}

func stringSliceEqual(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	sort.Strings(s1)
	sort.Strings(s2)
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func generatLocalAlarmExecutor() {
	executorStr := fmt.Sprintf(`#!/bin/bash
echo $1 >> %s
`, alarmMsgFile)
	ioutil.WriteFile(alarmExecutor, []byte(executorStr), 777)
	os.Chmod(alarmExecutor, 0777)
}

func clearAlarmMsgFile() {
	f, err := os.Create(alarmMsgFile)
	if err == nil {
		f.Close()
	}
}

func readAlarmMsgFile() (string, error) {
	readBytes, err := ioutil.ReadFile(alarmMsgFile)
	if err != nil {
		return "", err
	}

	return strings.Trim(string(readBytes), "\n"), nil
}

func removeAlarmFile() {
	os.Remove(alarmExecutor)
	os.Remove(alarmMsgFile)
}
