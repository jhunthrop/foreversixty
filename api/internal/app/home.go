package app

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/fs1"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
	"github.com/jhunthrop/foreversixty/api/internal/upgrades"
)

// home answers GET /v1/characters/{region}/{ruleset}/{name}/home. Who may see it is the
// character page's rule (its found=false is a 404) and it is cached as that page is.
func (s *Service) home(w http.ResponseWriter, r *http.Request) {
	h, found, err := s.HomeOf(r.Context(), r.PathValue("region"), r.PathValue("ruleset"), r.PathValue("name"))
	if err != nil {
		s.logger().Error("app", "id", httpx.RequestIDFrom(r.Context()), "op", "home", "err", err)
		httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not read that character just now", nil)
		return
	}
	if !found {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character", nil)
		return
	}
	rankings.CacheLive(w)
	httpx.WriteOK(w, r, http.StatusOK, h)
}

// HomeOf is the home answer for a character; found is false exactly when its public
// character page is.
func (s *Service) HomeOf(ctx context.Context, region, ruleset, name string) (Home, bool, error) {
	card, found, err := s.Characters.CharacterCard(ctx, region, ruleset, name)
	if err != nil || !found {
		return Home{}, false, err
	}
	in, err := s.inputOf(ctx, card)
	if err != nil {
		return Home{}, false, err
	}
	return answer(in), true, nil
}

// inputOf reads what the answer needs beside the card: the export, the band and the items.
func (s *Service) inputOf(ctx context.Context, card rankings.Character) (homeInput, error) {
	build, _ := s.Data.ActiveBuild()
	class := strings.ToLower(card.Character.Class)
	in := homeInput{card: card, level: card.Character.Level, spec: strOf(card.GearSpec), build: build,
		items: func(id int) (trees.Item, bool) { return s.Data.BuildItem(build, class, id) }}
	exp, has, err := s.Exports.Export(ctx, card.Character.Key)
	if err != nil {
		return homeInput{}, fmt.Errorf("app: home of %s: %w", card.Character.Key, err)
	}
	var decoded fs1.Decoded
	if has {
		in.synced, in.source = &exp.CapturedAt, exp.Source
		if decoded, in.hasExport = fs1.Decode(exp.Code); in.hasExport {
			in.talents = decoded.Talents.Trees
			if in.level == nil && decoded.HasLevel {
				in.level = &decoded.Level
			}
		}
	}
	in.band, in.found = s.bandOf(in.spec, in.level, s.factionOf(card, decoded))
	return in, nil
}

// factionOf is the character's race's faction, else the export's race's.
func (s *Service) factionOf(card rankings.Character, decoded fs1.Decoded) string {
	if f := s.Data.RaceFaction(strOf(card.Character.Race)); f != "" {
		return f
	}
	return sims.RaceFaction(decoded.DataBuild, decoded.RaceSlug)
}

// bandOf is the spec's band for the level and faction under the preset the band reads at
// (upgrades.PresetFor, the /char panel's choice); found is false when any part is unknown or
// nothing is published. A file that cannot be read is logged and reads as not published.
func (s *Service) bandOf(spec string, level *int, faction string) (bis.Band, bool) {
	band, ok := upgrades.LevelBand(level)
	if !ok || spec == "" || faction == "" {
		return bis.Band{}, false
	}
	b, found, err := bis.LoadPresetBand(s.Data.BisDataDir(), s.Data.BisBuild(), spec, band, faction, upgrades.PresetFor(band))
	if err != nil {
		s.logger().Warn("app", "op", "home_band", "spec", spec, "err", err)
	}
	return b, found && err == nil
}
