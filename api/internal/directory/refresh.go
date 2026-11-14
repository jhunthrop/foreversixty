// api/internal/directory/refresh.go
package directory

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/jhunthrop/foreversixty/api/internal/bnetapi"
)

// PassResult is what one pass did, for its log line (counts only).
type PassResult struct {
	Rosters, RosterMembers                     int
	Profiles, Private, Deleted, Misses, Failed int64
	Expired                                    int64
	Calls                                      int64
	RateLimited                                bool
}

// errStop ends a pass early: a 429 (every further call would spend the shared quota for
// another refusal) or the pass's own deadline.
var errStop = errors.New("directory: pass stopped")

// Pass seeds the sweep, reads due rosters, then due characters oldest first, then expires
// what no source confirmed within 30 days. It runs until nothing is due or ctx ends; the
// budget paces every call. A 429 stops it with RateLimited set and no error.
func (s *Service) Pass(ctx context.Context) (PassResult, error) {
	var res PassResult
	startCalls := s.budget().Calls()
	defer func() { res.Calls = s.budget().Calls() - startCalls }()
	if err := s.seed(ctx); err != nil {
		return res, err
	}
	steps := []func(context.Context, *PassResult) error{s.refreshRosters, s.seedNames, s.refreshCharacters}
	for _, step := range steps {
		if err := step(ctx, &res); err != nil {
			if errors.Is(err, errStop) || ctx.Err() != nil {
				break
			}
			return res, err
		}
	}
	expired, err := s.expire(context.WithoutCancel(ctx))
	res.Expired = expired
	return res, err
}

// refreshRosters reads every due guild's roster (and its summary while its ruleset is
// unknown), adding each member to the directory under the guild's ruleset.
// A guild whose read fails is left due for the next pass.
func (s *Service) refreshRosters(ctx context.Context, res *PassResult) error {
	guilds, err := s.dueGuilds(ctx, maxGuildsPerPass)
	if err != nil {
		return err
	}
	for _, g := range guilds {
		n, err := s.refreshRoster(ctx, g)
		if err != nil {
			stop, limited := s.softFail(err, "roster", g.Region+"/"+g.Slug, res)
			res.RateLimited = res.RateLimited || limited
			if stop {
				return errStop
			}
			continue
		}
		res.Rosters++
		res.RosterMembers += n
	}
	return nil
}

// refreshRoster is one guild: its summary when the ruleset is not known yet, then its roster.
func (s *Service) refreshRoster(ctx context.Context, g dirGuild) (int, error) {
	if g.Ruleset == "" {
		if err := s.budget().Spend(ctx); err != nil {
			return 0, errStop
		}
		summary, err := s.Client.GuildSummary(ctx, g.Region, "", g.Slug)
		if isGone(err) {
			return 0, s.guildMissing(ctx, g)
		}
		if err != nil {
			return 0, err
		}
		if err := s.stampGuild(ctx, g, summary); err != nil {
			return 0, err
		}
		g.Ruleset = summary.Ruleset
	}
	if err := s.budget().Spend(ctx); err != nil {
		return 0, errStop
	}
	roster, err := s.Client.GuildRosterBySlug(ctx, g.Region, "", g.Slug)
	if isGone(err) {
		return 0, s.guildMissing(ctx, g)
	}
	if err != nil {
		return 0, err
	}
	return len(roster.Members), s.upsertRoster(ctx, g, roster)
}

// refreshCharacters reads every due character, Workers at a time.
func (s *Service) refreshCharacters(ctx context.Context, res *PassResult) error {
	for {
		due, err := s.dueCharacters(ctx, batchSize)
		if err != nil || len(due) == 0 {
			return err
		}
		if err := s.readAll(ctx, due, res); err != nil {
			return err
		}
	}
}

// readAll reads rows on Workers goroutines; the first stop (429, deadline) ends the batch.
// A row whose read fails is put back for retryAfter, so the pass moves on past it.
func (s *Service) readAll(ctx context.Context, rows []dueRow, res *PassResult) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	work := make(chan dueRow)
	var stopped, limited atomic.Bool
	var wg sync.WaitGroup
	for range s.workers() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for row := range work {
				err := s.refreshOne(ctx, row, res)
				if err == nil {
					continue
				}
				stop, hit429 := s.softFail(err, "character", row.Slug, res)
				if hit429 {
					limited.Store(true)
				}
				if stop {
					stopped.Store(true)
					cancel()
					continue
				}
				if err := s.retryLater(context.WithoutCancel(ctx), row); err != nil {
					s.logger().Warn("directory", "op", "retry_later", "err", err)
				}
			}
		}()
	}
	for _, row := range rows {
		if stopped.Load() || ctx.Err() != nil {
			break
		}
		work <- row
	}
	close(work)
	wg.Wait()
	res.RateLimited = res.RateLimited || limited.Load()
	if stopped.Load() {
		return errStop
	}
	return nil
}

// softFail decides what one failed read does to the pass: a 429 or a spent deadline stops it
// (limited reports the 429); anything else is logged and counted, and the pass goes on.
func (s *Service) softFail(err error, op, what string, res *PassResult) (stop, limited bool) {
	switch {
	case errors.Is(err, bnetapi.ErrRateLimited):
		return true, true
	case errors.Is(err, errStop) || errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded):
		return true, false
	}
	atomic.AddInt64(&res.Failed, 1)
	s.logger().Warn("directory", "op", op, "what", what, "err", err)
	return false, false
}

// refreshOne validates one character through /status, then reads its public profile,
// equipment, specializations, professions and media. A character Blizzard says is gone is
// deleted; one that is not public is kept only while a roster vouches for it.
func (s *Service) refreshOne(ctx context.Context, row dueRow, res *PassResult) error {
	status, err := s.status(ctx, row)
	if err != nil {
		return err
	}
	switch {
	case status.missing:
		return s.notPublic(ctx, row, res)
	case !status.Valid:
		return s.gone(ctx, row, res)
	case row.ID != 0 && status.ID != 0 && status.ID != row.ID:
		// The name now belongs to another character: the old one is gone, the new one is read.
		atomic.AddInt64(&res.Deleted, 1)
		if err := s.deleteCharacter(ctx, row.Region, row.ID); err != nil {
			return err
		}
	}
	read, found, err := s.readProfile(ctx, row)
	if err != nil {
		return err
	}
	if !found {
		return s.notPublic(ctx, row, res)
	}
	if err := s.saveProfile(ctx, row.Region, row.Sources, read); err != nil {
		return err
	}
	atomic.AddInt64(&res.Profiles, 1)
	return nil
}

type statusRead struct {
	bnetapi.CharacterStatus
	missing bool
}

func (s *Service) status(ctx context.Context, row dueRow) (statusRead, error) {
	if err := s.budget().Spend(ctx); err != nil {
		return statusRead{}, errStop
	}
	st, err := s.Client.Status(ctx, row.Region, "", row.Slug)
	if isGone(err) {
		return statusRead{missing: true}, nil
	}
	if err != nil {
		return statusRead{}, fmt.Errorf("status: %w", err)
	}
	return statusRead{CharacterStatus: st}, nil
}

// notPublic is a character whose status or profile answered 404 or 403: a name seed is
// remembered as a miss, a roster member is kept without its profile fields, and a row no
// roster vouches for is deleted.
func (s *Service) notPublic(ctx context.Context, row dueRow, res *PassResult) error {
	if row.ID == 0 {
		atomic.AddInt64(&res.Misses, 1)
		return s.recordMiss(ctx, row.Region, row.Slug)
	}
	if !row.RosterRecent {
		return s.gone(ctx, row, res)
	}
	atomic.AddInt64(&res.Private, 1)
	return s.markPrivate(ctx, row.Region, row.ID)
}

// gone deletes a character Blizzard no longer vouches for, and remembers its name as a miss
// so no seed brings it straight back.
func (s *Service) gone(ctx context.Context, row dueRow, res *PassResult) error {
	if row.ID == 0 {
		atomic.AddInt64(&res.Misses, 1)
		return s.recordMiss(ctx, row.Region, row.Slug)
	}
	atomic.AddInt64(&res.Deleted, 1)
	if err := s.deleteCharacter(ctx, row.Region, row.ID); err != nil {
		return err
	}
	return s.recordMiss(ctx, row.Region, row.Slug)
}

// readProfile reads the summary, then the sub-resources; found is false when the summary
// is not public. A sub-resource that answers 404 or 403 is left empty.
func (s *Service) readProfile(ctx context.Context, row dueRow) (profileRead, bool, error) {
	if err := s.budget().Spend(ctx); err != nil {
		return profileRead{}, false, errStop
	}
	profile, _, err := s.Client.Character(ctx, row.Region, "", row.Slug)
	if isGone(err) {
		return profileRead{}, false, nil
	}
	if err != nil {
		return profileRead{}, false, fmt.Errorf("profile: %w", err)
	}
	read := profileRead{Profile: profile}
	for _, sub := range s.subReads(row, &read) {
		if err := s.budget().Spend(ctx); err != nil {
			return profileRead{}, false, errStop
		}
		if err := sub(ctx); err != nil && !isGone(err) {
			return profileRead{}, false, err
		}
	}
	return read, true, nil
}

// subReads are the four sub-resource reads, each filling its part of read.
func (s *Service) subReads(row dueRow, read *profileRead) []func(context.Context) error {
	region, slug := row.Region, row.Slug
	return []func(context.Context) error{
		func(ctx context.Context) error {
			raw, err := s.Client.Equipment(ctx, region, "", slug)
			if err == nil {
				read.Equipment, err = summarizeEquipment(raw)
			}
			return err
		},
		func(ctx context.Context) error {
			raw, err := s.Client.Specializations(ctx, region, "", slug)
			if err == nil {
				read.Spec, err = activeSpec(raw)
			}
			return err
		},
		func(ctx context.Context) error {
			raw, err := s.Client.Professions(ctx, region, "", slug)
			if err == nil {
				read.Professions, err = summarizeProfessions(raw)
			}
			return err
		},
		func(ctx context.Context) error {
			media, _, err := s.Client.CharacterMedia(ctx, region, "", slug)
			read.AvatarURL = media.AvatarURL
			return err
		},
	}
}

// isGone is a 404 or 403: not found, or not public.
func isGone(err error) bool {
	return errors.Is(err, bnetapi.ErrNotFound) || errors.Is(err, bnetapi.ErrForbidden)
}
