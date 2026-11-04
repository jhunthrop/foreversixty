package linkedrole

import (
	"bytes"
	"context"
	"strings"
	"sync"
)

// memStore is the Store fake: links by user, profiles by user, guild members by guild.
type memStore struct {
	mu       sync.Mutex
	links    map[int64]Link
	profiles map[int64]Profile
	members  map[int64][]int64
	pushed   map[int64]bool
}

func newMemStore() *memStore {
	return &memStore{links: map[int64]Link{}, profiles: map[int64]Profile{}, members: map[int64][]int64{}, pushed: map[int64]bool{}}
}

func (m *memStore) Get(_ context.Context, userID int64) (Link, bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	l, ok := m.links[userID]
	return l, ok, nil
}

func (m *memStore) Save(_ context.Context, l Link) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	for id, other := range m.links {
		if other.DiscordUserID == l.DiscordUserID && id != l.UserID {
			delete(m.links, id)
		}
	}
	m.links[l.UserID] = l
	return nil
}

func (m *memStore) Rotate(_ context.Context, userID int64, oldBox, newBox []byte) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	l, ok := m.links[userID]
	if !ok || !bytes.Equal(l.RefreshBox, oldBox) {
		return false, nil
	}
	l.RefreshBox = newBox
	m.links[userID] = l
	return true, nil
}

func (m *memStore) DeleteIf(_ context.Context, userID int64, box []byte) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	l, ok := m.links[userID]
	if !ok || !bytes.Equal(l.RefreshBox, box) {
		return false, nil
	}
	delete(m.links, userID)
	return true, nil
}

func (m *memStore) MarkPushed(_ context.Context, userID int64, premium bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.pushed[userID] = premium
	return nil
}

func (m *memStore) Profile(_ context.Context, userID int64) (Profile, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.profiles[userID], nil
}

func (m *memStore) LinkedUsers(_ context.Context, guildID int64) ([]int64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	var ids []int64
	if guildID == 0 {
		for id := range m.links {
			ids = append(ids, id)
		}
		return ids, nil
	}
	for _, id := range m.members[guildID] {
		if _, ok := m.links[id]; ok {
			ids = append(ids, id)
		}
	}
	return ids, nil
}

// plainSealer "seals" by prefixing, so a test can read what was stored.
type plainSealer struct{}

func (plainSealer) Seal(plain string) ([]byte, error) { return []byte("sealed:" + plain), nil }
func (plainSealer) Open(box []byte) (string, error) {
	return strings.TrimPrefix(string(box), "sealed:"), nil
}

// fakeDiscord rotates refresh tokens r0 -> r1 -> ..., refuses tokens in revoked, and records
// pushes by access token.
type fakeDiscord struct {
	mu       sync.Mutex
	revoked  map[string]bool
	exchange Tokens
	userID   string
	pushes   []push
	next     int
}

type push struct {
	Access string
	Conn   Connection
}

func (f *fakeDiscord) AuthorizeURL(state string) string {
	return "https://discord.test/authorize?state=" + state
}

func (f *fakeDiscord) Exchange(context.Context, string) (Tokens, error) { return f.exchange, nil }

func (f *fakeDiscord) Refresh(_ context.Context, refresh string) (Tokens, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.revoked[refresh] {
		return Tokens{}, ErrRevoked
	}
	f.next++
	n := string(rune('0' + f.next))
	return Tokens{Access: "a-" + refresh + "-" + n, Refresh: refresh + "+" + n}, nil
}

func (f *fakeDiscord) UserID(context.Context, string) (string, error) { return f.userID, nil }

func (f *fakeDiscord) Push(_ context.Context, access string, c Connection) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.pushes = append(f.pushes, push{Access: access, Conn: c})
	return nil
}
