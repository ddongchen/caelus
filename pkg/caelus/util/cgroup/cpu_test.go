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

package cgroup

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

var (
	tmpCpuCgroupPath = "/sys/fs/cgroup/cpu,cpuacct/test"
	tmpCpuSubPath    = "test"
	cores            = 2
	expectQuota      = "2000000"
)

// TestSetCpuQuota test cpu quota
func TestSetCpuQuota(t *testing.T) {
	err := os.MkdirAll(tmpCpuCgroupPath, 0700)
	if err != nil {
		t.Skipf("mkdir cgroup path(%s) err: %v", tmpCpuCgroupPath, err)
	}
	defer os.RemoveAll(tmpCpuCgroupPath)

	err = SetCpuQuota(tmpCpuSubPath, float64(cores))
	if err != nil {
		t.Fatalf("set cpu quota for %s err: %v", tmpCpuCgroupPath, err)
	}

	quotaBytes, err := ioutil.ReadFile(path.Join(tmpCpuCgroupPath, "cpu.cfs_quota_us"))
	if err != nil {
		t.Fatalf("read cpu quota for %s err: %v", tmpCpuCgroupPath, err)
	}

	if strings.Trim(string(quotaBytes), "\n") != expectQuota {
		t.Fatalf("unexpect quot value: %s, expect: %s", string(quotaBytes), expectQuota)
	}
}
