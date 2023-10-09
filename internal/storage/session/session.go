// internal/storage/session/session.go

package session

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"
)

type Session struct {
	Username     string
	LastActivity time.Time
}

const sessionTimeout = 1 * time.Hour

var sessionStore = make(map[string]*Session)
var sessionMutex = &sync.Mutex{}

func createSessionID() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func CreateSessionForUser(username string) (string, error) {
	sessionID, err := createSessionID()
	if err != nil {
		return "", err
	}

	newSession := &Session{
		Username:     username,
		LastActivity: time.Now(),
	}

	// Lock the session store to prevent race conditions
	sessionMutex.Lock()
	defer sessionMutex.Unlock()

	sessionStore[sessionID] = newSession
	return sessionID, nil
}

func GetSession(sessionID string) *Session {
	// Lock the session store to prevent race conditions
	sessionMutex.Lock()
	defer sessionMutex.Unlock()

	return sessionStore[sessionID]
}

func IsValidSession(sessionID string) bool {
	sessionMutex.Lock()
	defer sessionMutex.Unlock()

	session, exists := sessionStore[sessionID]
	if !exists {
		return false
	}

	// Check if the session has expired
	if time.Since(session.LastActivity) > sessionTimeout {
		delete(sessionStore, sessionID) // Remove expired session
		return false
	}

	// Update the LastActivity timestamp to extend the session's life
	//session.LastActivity = time.Now()

	return true
}
