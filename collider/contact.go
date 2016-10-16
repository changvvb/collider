// Copyright (c) 2014 The WebRTC project authors. All Rights Reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package collider

import (
//"errors"
//"fmt"
//"io"
//"log"
//"net/http"
//"time"
)

type contact struct {
	owner     *client
	clientsID []string
}

func newContact(clientID string) *contact {
	//在这里查找数据库;
	return &contact{
		clientsID: []string{"001", "002", "003"},
	}
}

// remove closes the client connection and removes the client specified by the |clientID|.
func (ct *contact) remove(clientID string) {
}
