// api/internal/guilds/item_needs.go
//
// GET /v1/guilds/{id}/items/{item}/needs: who in the guild needs an item and who has it, for
// the apps' item sheet (spec 2026-11-08-app-flows F5). A verified character needs it when the
// BiS band of its spec at its own level band (upgrades.LevelBand, the raid preset at 60, as the
// site's BiS page reads) names the item as a slot's pick or a listed alternative and it does
// not wear it; it has it when it wears it. Gear is known only at gear consent (the roster's
// rule): a character with no gear on file is neither, only counted in unknown_gear.
package guilds

import (
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/upgrades"
)

// ItemNeeds is the body of GET …/items/{item}/needs. NextRaid is the guild's next raid in the
// zone that drops the item (item_next_raid.go), null when none is planned or no table raid
// drops it.
type ItemNeeds struct {
	ItemID      int           `json:"item_id"`
	Needs       []ItemNeed    `json:"needs"`
	Has         []ItemHolder  `json:"has"`
	UnknownGear int           `json:"unknown_gear"`
	NextRaid    *ItemNextRaid `json:"next_raid"`
}

// ItemNeed is one character whose band names the item. Pick is true when it is the slot's pick,
// false for a listed alternative; Slot is the band's slot that names it.
type ItemNeed struct {
	ItemHolder
	Level int    `json:"level"`
	Band  int    `json:"band"`
	Slot  string `json:"slot"`
	Pick  bool   `json:"pick"`
}

// ItemHolder is a roster character, as the sheet names it.
type ItemHolder struct {
	CharacterKey string `json:"character_key"`
	Name         string `json:"name"`
	Class        string `json:"class"`
	Spec         string `json:"spec"`
}

func (s *Service) itemNeeds(w http.ResponseWriter, r *http.Request) {
	guildID, gok := guildIDFrom(r)
	itemID, err := strconv.Atoi(r.PathValue("item"))
	if !gok || err != nil || itemID <= 0 {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such item", nil)
		return
	}
	_, verified, err := s.Accounts.GuildRank(r.Context(), guildID, auth.ActorFrom(r.Context()).UserID)
	if err != nil {
		s.fail(w, r, "item_needs", err, "could not load who needs that just now")
		return
	}
	if !verified {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	lr, err := s.Store.lootRoster(r.Context(), guildID)
	if err != nil {
		s.fail(w, r, "item_needs", err, "could not load who needs that just now")
		return
	}
	out := s.Store.guildItemNeeds(lr.roster, itemID)
	if out.NextRaid, err = s.Store.itemNextRaid(r.Context(), guildID, itemID, time.Now()); err != nil {
		s.fail(w, r, "item_needs", err, "could not load who needs that just now")
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, out)
}

// guildItemNeeds sorts the verified roster into who needs itemID, who has it, and how many have
// no gear on file. Needs read picks first, then by name; has by name.
func (s *Store) guildItemNeeds(roster []RosterRow, itemID int) ItemNeeds {
	out := ItemNeeds{ItemID: itemID, Needs: []ItemNeed{}, Has: []ItemHolder{}}
	for _, row := range roster {
		if !row.Verified {
			continue
		}
		if len(row.gear) == 0 {
			out.UnknownGear++
			continue
		}
		if wears(row, itemID) {
			out.Has = append(out.Has, holderOf(row))
			continue
		}
		if need, ok := s.needOf(row, itemID); ok {
			out.Needs = append(out.Needs, need)
		}
	}
	slices.SortStableFunc(out.Needs, func(a, b ItemNeed) int {
		if a.Pick != b.Pick {
			if a.Pick {
				return -1
			}
			return 1
		}
		return strings.Compare(a.Name, b.Name)
	})
	slices.SortStableFunc(out.Has, func(a, b ItemHolder) int { return strings.Compare(a.Name, b.Name) })
	return out
}

// wears reports whether row wears itemID in any slot.
func wears(row RosterRow, itemID int) bool {
	for _, id := range row.gear {
		if id == itemID {
			return true
		}
	}
	return false
}

// needOf is row's need for itemID: its band at its level names the item in some slot.
func (s *Store) needOf(row RosterRow, itemID int) (ItemNeed, bool) {
	level, ok := rosterLevel(row)
	if !ok {
		return ItemNeed{}, false
	}
	band, ok := upgrades.LevelBand(&level)
	if !ok {
		return ItemNeed{}, false
	}
	b, ok := s.loadBandAt(row.className, row.specStr, row.faction, band)
	if !ok {
		return ItemNeed{}, false
	}
	for _, slot := range b.Slots {
		if slot.ItemID > 0 && namesItem(slot, itemID) {
			return ItemNeed{ItemHolder: holderOf(row), Level: level, Band: band, Slot: slot.Slot, Pick: slot.ItemID == itemID}, true
		}
	}
	return ItemNeed{}, false
}

// holderOf is row as the sheet names it.
func holderOf(row RosterRow) ItemHolder {
	h := ItemHolder{CharacterKey: awardKeyOf(row), Name: row.Name}
	if row.Class != nil {
		h.Class = *row.Class
	}
	if row.Spec != nil {
		h.Spec = *row.Spec
	}
	return h
}
