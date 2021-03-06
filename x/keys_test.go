/*
 * Copyright 2016-2018 Dgraph Labs, Inc. and Contributors
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
 */

package x

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDataKey(t *testing.T) {
	var uid uint64
	for uid = 0; uid < 1001; uid++ {
		sattr := fmt.Sprintf("attr:%d", uid)
		key := DataKey(sattr, uid)
		pk := Parse(key)

		require.True(t, pk.IsData())
		require.Equal(t, sattr, pk.Attr)
		require.Equal(t, uid, pk.Uid)
		require.Equal(t, uint64(0), pk.StartUid)
	}

	keys := make([]string, 0, 1024)
	for uid = 1024; uid >= 1; uid-- {
		key := DataKey("testing.key", uid)
		keys = append(keys, string(key))
	}
	// Test that sorting is as expected.
	sort.Strings(keys)
	require.True(t, sort.StringsAreSorted(keys))
	for i, key := range keys {
		exp := DataKey("testing.key", uint64(i+1))
		require.Equal(t, string(exp), key)
	}
}

func TestParseKeysWithStartUid(t *testing.T) {
	var uid uint64
	startUid := uint64(1024)
	for uid = 0; uid < 1001; uid++ {
		sattr := fmt.Sprintf("attr:%d", uid)
		key := DataKey(sattr, uid)
		key = GetSplitKey(key, startUid)
		pk := Parse(key)

		require.True(t, pk.IsData())
		require.Equal(t, sattr, pk.Attr)
		require.Equal(t, uid, pk.Uid)
		require.Equal(t, startUid, pk.StartUid)
	}
}

func TestIndexKey(t *testing.T) {
	var uid uint64
	for uid = 0; uid < 1001; uid++ {
		sattr := fmt.Sprintf("attr:%d", uid)
		sterm := fmt.Sprintf("term:%d", uid)

		key := IndexKey(sattr, sterm)
		pk := Parse(key)

		require.True(t, pk.IsIndex())
		require.Equal(t, sattr, pk.Attr)
		require.Equal(t, sterm, pk.Term)
	}
}

func TestReverseKey(t *testing.T) {
	var uid uint64
	for uid = 0; uid < 1001; uid++ {
		sattr := fmt.Sprintf("attr:%d", uid)

		key := ReverseKey(sattr, uid)
		pk := Parse(key)

		require.True(t, pk.IsReverse())
		require.Equal(t, sattr, pk.Attr)
		require.Equal(t, uid, pk.Uid)
	}
}

func TestSchemaKey(t *testing.T) {
	var uid uint64
	for uid = 0; uid < 1001; uid++ {
		sattr := fmt.Sprintf("attr:%d", uid)

		key := SchemaKey(sattr)
		pk := Parse(key)

		require.True(t, pk.IsSchema())
		require.Equal(t, sattr, pk.Attr)
	}
}
