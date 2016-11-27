// Copyright (c) 2014 The WebRTC project authors. All Rights Reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

package collider

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"time"
)

const maxQueuedMsgCount = 1024

const (
	OFFLINE = "OFFLINE"
	ONLINE  = "ONLINE"
	BUSY    = "BUSY"
	LEAVE   = "LEAVE"
)

type client struct {
	id string
	// rwc is the interface to access the websocket connection.
	// It is set after the client registers with the server.
	rwc io.ReadWriteCloser
	// msgs is the queued messages sent from this client.
	msgs []string
	// timer is used to remove this client if unregistered after a timeout.
	timer    *time.Timer
	contact_ *contact
	//状态
	state string
}

var registered_clients map[string]*client

func newClient(id string, t *time.Timer) *client {
	c := client{id: id, timer: t}
	c.contact_ = newContact(id)
	return &c
}

func (c *client) setTimer(t *time.Timer) {
	if c.timer != nil {
		c.timer.Stop()
	}
	c.timer = t
}

// register binds the ReadWriteCloser to the client if it's not done yet.
func (c *client) register(rwc io.ReadWriteCloser) error {
	if c.rwc != nil {
		log.Printf("Not registering because the client %s already has a connection", c.id)
		return errors.New("Duplicated registration")
	}

	registered_clients[c.id] = c
	c.setTimer(nil)
	c.rwc = rwc

	//set state
	c.state = ONLINE
	c.getContactState()
	c.informState()
	c.getOfflineMessage()

	return nil
}

// deregister closes the ReadWriteCloser if it exists.
func (c *client) deregister() {
	c.state = OFFLINE
	c.informState()

	if c.rwc != nil {
		c.rwc.Close()
		c.rwc = nil
	}
	delete(registered_clients, c.id)
}

// registered returns true if the client has registered.
func (c *client) registered() bool {
	return c.rwc != nil
}

// enqueue adds a message to the client's message queue.
func (c *client) enqueue(msg string) error {
	if len(c.msgs) >= maxQueuedMsgCount {
		return errors.New("Too many messages queued for the client")
	}
	c.msgs = append(c.msgs, msg)
	return nil
}

// sendQueued the queued messages to the other client.
func (c *client) sendQueued(other *client) error {
	if c.id == other.id || other.rwc == nil {
		return errors.New("Invalid client")
	}
	for _, m := range c.msgs {
		sendServerMsg(other.rwc, "", m)
	}
	c.msgs = nil
	log.Printf("Sent queued messages from %s to %s", c.id, other.id)
	return nil
}

// send sends the message to the other client if the other client has registered,
// or queues the message otherwise.
func (c *client) send(other *client, cmd string, msg string) error {
	if c.id == other.id {
		return errors.New("Invalid client")
		log.Printf("Invalid client")
	}
	if other.rwc != nil {
		log.Printf("sending %s to %s from %s, cmd is %s", msg, other.id, c.id, cmd)
		return sendServerMsg(other.rwc, cmd, msg)
	}
	return c.enqueue(msg)
}

//通过ClientID发送信息
func (c *client) sendByID(OtherClientID string, cmd string, msg string) error {
	if other := registered_clients[OtherClientID]; other != nil {
		if other.rwc != nil {
			log.Printf("sending %s to %s from %s, cmd is %s", msg, other.id, c.id, cmd)
			m := wsServerMsg{
				Msg:  msg,
				Cmd:  cmd,
				From: c.id,
			}
			return send(other.rwc, m)
		}
	} else {
		log.Println("The receiver is offline now")

		db, err := sql.Open("mysql", MYSQL_CONNECT_STRING)
		if err != nil {
			return err
		}

		stmt, err := db.Prepare("INSERT INTO offlineMessage(cmd,fromid,toid,msg,created) VALUES (?,?,?,?,?)")
		if err != nil {
			return err
		}
		res, err := stmt.Exec(cmd, c.id, OtherClientID, msg, time.Now())
		if err != nil {
			log.Println("Error,exec", err)
			return nil
		}
		affect, err := res.RowsAffected()
		if err != nil {
			log.Println("Affect,exec")
			return err
		}
		if affect != 0 {
			log.Printf("insert offlineMessage successfully")
		}
	}
	return nil
}

func (c *client) informState() {
	m := wsServerMsg{
		Msg:  c.state,
		Cmd:  "inform",
		From: c.id,
	}
	for _, contact_ := range c.contact_.clientsID {
		if client_ := registered_clients[contact_]; client_ != nil {
			send(client_.rwc, m)
		}
	}

}

func (c *client) getContactState() {
	for _, contact_ := range c.contact_.clientsID {
		state, client_ := c.getOneStateByID(contact_)
		if client_ != nil {
			m := wsServerMsg{
				Cmd:  "contact_state",
				From: contact_,
				Msg:  state,
			}
			log.Printf("m.Msg:%s", m.Msg)
			send(c.rwc, m)

		}
	}
}

func (c *client) getOneStateByID(ClientID string) (string, *client) {
	if client_ := registered_clients[ClientID]; client_ != nil {
		return client_.state, client_
	} else {
		return "OFFLINE", nil
	}
}

const MYSQL_CONNECT_STRING = "root:root@tcp(localhost:3306)/im?parseTime=true"

func (c *client) getOfflineMessage() {
	db, err := sql.Open("mysql", MYSQL_CONNECT_STRING)
	checkErr(err)

	rows, err := db.Query("SELECT * FROM offlineMessage WHERE toid=" + c.id)
	checkErr(err)
	for rows.Next() {
		var cmd string
		var origin string
		var message string
		var to string
		var msgTime time.Time
		var id int
		err = rows.Scan(&id, &cmd, &to, &origin, &msgTime, &message)
		checkErr(err)
		m := wsServerMsg{
			From: origin,
			Msg:  message,
			Cmd:  "offlinemessage",
			Time: msgTime.Local(),
		}
		log.Printf("%+v\n", m)
		send(c.rwc, m)
	}
	stmt, err := db.Prepare("DELETE FROM offlineMessage WHERE toid=?")
	checkErr(err)
	_, err = stmt.Exec(c.id)
	checkErr(err)
	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
