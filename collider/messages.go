// Copyright (c) 2014 The WebRTC project authors. All Rights Reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package collider

import (
	"encoding/json"
	"io"
)

// WebSocket message from the client.
type wsClientMsg struct {
	Cmd      string `json:"cmd"`
	RoomID   string `json:"roomid"`
	To       string `json:"to"`
	ClientID string `json:"clientid"`
	Msg      string `json:"msg"`
}

// wsServerMsg is a message sent to a client on behalf of another client.
type wsServerMsg struct {
	Cmd   string `json:"cmd"`
	From  string `json:"from"`
	Msg   string `json:"msg"`
	Error string `json:"error"`
}

// sendServerMsg sends a wsServerMsg composed from |msg| to the connection.
func sendServerMsg(w io.Writer, cmd string, msg string) error {
	m := wsServerMsg{
		Msg: msg,
		Cmd: cmd,
	}
	return send(w, m)
}

// sendServerErr sends a wsServerMsg composed from |errMsg| to the connection.
func sendServerErr(w io.Writer, errMsg string) error {
	m := wsServerMsg{
		Error: errMsg,
	}
	return send(w, m)
}

// send writes a generic object as JSON to the writer.
func send(w io.Writer, data interface{}) error {
	enc := json.NewEncoder(w)
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}
