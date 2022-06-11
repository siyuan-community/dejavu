// DejaVu - Data snapshot and sync.
// Copyright (c) 2022-present, b3log.org
//
// DejaVu is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package dejavu

import (
	"bytes"
	"testing"

	"github.com/siyuan-note/encryption"
)

func TestPutGet(t *testing.T) {
	clearTestdata(t)

	aesKey, err := encryption.KDF(testRepoPassword)
	if nil != err {
		return
	}

	store := NewStore(testRepoPath+"/objects/", aesKey)

	data := []byte("Hello!")
	chunk := &Chunk{Hash: Hash(data), Data: data}
	err = store.PutChunk(chunk)
	if nil != err {
		t.Fatalf("put failed: %s", err)
		return
	}

	chunk, err = store.GetChunk(chunk.Hash)
	if nil != err {
		t.Fatalf("get failed: %s", err)
		return
	}
	if 0 != bytes.Compare(chunk.Data, data) {
		t.Fatalf("data not match")
		return
	}

	err = store.Remove(chunk.Hash)
	if nil != err {
		t.Fatalf("remove failed: %s", err)
		return
	}

	chunk, err = store.GetChunk(chunk.Hash)
	if nil != chunk {
		t.Fatalf("get should be failed")
		return
	}
}
