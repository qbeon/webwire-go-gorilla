package webwire

import (
	"github.com/qbeon/webwire-go/message"
)

// handleSessionClosure handles session destruction requests
// and returns an error if the ongoing connection cannot be proceeded
func (srv *server) handleSessionClosure(
	con *connection,
	msg *message.Message,
) {
	finalize := func() {
		srv.deregisterHandler(con)

		// Release message buffer
		msg.Close()
	}

	if !srv.sessionsEnabled {
		srv.failMsg(con, msg, ErrSessionsDisabled{})
		finalize()
		return
	}

	if !con.HasSession() {
		// Send confirmation even though no session was closed
		srv.fulfillMsg(con, msg, Payload{})
		finalize()
		return
	}

	// Deregister session from active sessions registry destroying it if it's
	// the last connection left
	srv.sessionRegistry.deregister(con, true)

	// Reset the session on the connection
	con.setSession(nil)

	// Send confirmation
	srv.fulfillMsg(con, msg, Payload{})
	finalize()
}
