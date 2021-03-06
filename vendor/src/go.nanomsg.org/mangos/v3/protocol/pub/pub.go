// Copyright 2018 The Mangos Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package pub implements the PUB protocol.  This protocol publishes messages
// to subscribers (SUB peers).  The subscribers will filter incoming messages
// from the publisher based on their subscription.
package pub

import (
	"go.nanomsg.org/mangos/v3/protocol"
	"go.nanomsg.org/mangos/v3/protocol/xpub"
)

type socket struct {
	protocol.Protocol
}

// Protocol identity information.
const (
	Self     = protocol.ProtoPub
	Peer     = protocol.ProtoSub
	SelfName = "pub"
	PeerName = "sub"
)

func (s *socket) GetOption(name string) (interface{}, error) {
	switch name {
	case protocol.OptionRaw:
		return false, nil
	}
	return s.Protocol.GetOption(name)
}

// NewProtocol returns a new protocol implementation.
func NewProtocol() protocol.Protocol {
	s := &socket{
		Protocol: xpub.NewProtocol(),
	}
	return s
}

// NewSocket allocates a raw Socket using the PAIR protocol.
func NewSocket() (protocol.Socket, error) {
	return protocol.MakeSocket(NewProtocol()), nil
}
