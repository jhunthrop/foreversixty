package discordbot

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// candidatesTTL is how long an account's candidates are reused. Each read scans the guild's
// roster, and autocomplete asks on every keystroke; a character approved meanwhile shows
// within this window. SignUp checks the character itself, so a stale list never writes.
const candidatesTTL = 30 * time.Second

// candidateSweepAt is the cache size past which a put first drops expired entries.
const candidateSweepAt = 1024

type candidateKey struct{ guildID, userID int64 }

type candidateEntry struct {
	cands []guilds.SignupCandidate
	until time.Time
}

// candidateCache is a per-(guild, account) cache of SignupCandidates. Its zero value is ready.
type candidateCache struct {
	mu sync.Mutex
	by map[candidateKey]candidateEntry
}

func (c *candidateCache) get(k candidateKey, now time.Time) ([]guilds.SignupCandidate, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, ok := c.by[k]
	if !ok || !now.Before(e.until) {
		return nil, false
	}
	return e.cands, true
}

func (c *candidateCache) put(k candidateKey, cands []guilds.SignupCandidate, now time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.by == nil {
		c.by = map[candidateKey]candidateEntry{}
	}
	if len(c.by) >= candidateSweepAt {
		for key, e := range c.by {
			if !now.Before(e.until) {
				delete(c.by, key)
			}
		}
	}
	c.by[k] = candidateEntry{cands: cands, until: now.Add(candidatesTTL)}
}

// signupCandidates is the account's characters in the guild, read through the cache.
func (s *Service) signupCandidates(ctx context.Context, guildID, userID int64) ([]guilds.SignupCandidate, error) {
	k, now := candidateKey{guildID, userID}, s.now()
	if cands, ok := s.candidates.get(k, now); ok {
		return cands, nil
	}
	cands, err := s.d.Guilds.SignupCandidates(ctx, guildID, userID)
	if err != nil {
		return nil, fmt.Errorf("discordbot: signup candidates: %w", err)
	}
	s.candidates.put(k, cands, now)
	return cands, nil
}
