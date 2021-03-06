// Copyright 2016 Keybase Inc. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

package libkbfs

import (
	"golang.org/x/net/context"

	"github.com/keybase/client/go/protocol/keybase1"
)

func getCurrentUIDAndVerifyingKey(ctx context.Context, cig currentInfoGetter) (
	keybase1.UID, VerifyingKey, error) {
	_, uid, err := cig.GetCurrentUserInfo(ctx)
	if err != nil {
		return keybase1.UID(""), VerifyingKey{}, err
	}

	key, err := cig.GetCurrentVerifyingKey(ctx)
	if err != nil {
		return keybase1.UID(""), VerifyingKey{}, err
	}
	return uid, key, nil
}
