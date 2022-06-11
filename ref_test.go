// DejaVu - Data snapshot and sync.
// Copyright (c) 2022-present, b3log.org
//
// DejaVu is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
//
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//
// See the Mulan PSL v2 for more details.

package dejavu

import (
	"testing"

	"github.com/siyuan-note/encryption"
)

func TestTag(t *testing.T) {
	clearTestdata(t)

	aesKey, err := encryption.KDF(testRepoPassword)
	if nil != err {
		return
	}

	repo, err := NewRepo(testDataPath, testRepoPath, aesKey)
	if nil != err {
		t.Fatalf("new repo failed: %s", err)
		return
	}
	index, err := repo.Commit()
	if nil != err {
		t.Fatalf("commit failed: %s", err)
		return
	}
	t.Logf("commit: %s", index.Hash)

	err = repo.AddTag(index.Hash, "v1.0.0")
	if nil != err {
		t.Fatalf("add tag failed: %s", err)
		return
	}

	v100, err := repo.GetTag("v1.0.0")
	if v100 != index.Hash {
		t.Fatalf("get tag failed: %s", err)
		return
	}

	err = repo.AddTag(index.Hash, "v1.0.1")
	if nil != err {
		t.Fatalf("add tag failed: %s", err)
		return
	}

	v101, err := repo.GetTag("v1.0.1")
	if v101 != v100 {
		t.Fatalf("get tag failed: %s", err)
		return
	}
}
