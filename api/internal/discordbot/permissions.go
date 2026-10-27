package discordbot

import (
	"encoding/json"
	"sort"
	"strconv"
)

// Perm is a Discord permission bit set. Discord sends it as a decimal string.
type Perm uint64

// Permission bits the bot reasons about.
const (
	PermAdministrator      Perm = 1 << 3
	PermViewChannel        Perm = 1 << 10
	PermSendMessages       Perm = 1 << 11
	PermEmbedLinks         Perm = 1 << 14
	PermReadMessageHistory Perm = 1 << 16
)

// UnmarshalJSON reads Discord's decimal string.
func (p *Perm) UnmarshalJSON(raw []byte) error {
	var s string
	if err := json.Unmarshal(raw, &s); err != nil {
		return err
	}
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	*p = Perm(v)
	return nil
}

// Channel types and overwrite kinds from Discord.
const (
	channelText     = 0
	channelCategory = 4

	// OverwriteRole and OverwriteMember are the overwrite target kinds.
	OverwriteRole   = 0
	OverwriteMember = 1
)

// Overwrite is one permission overwrite on a channel.
type Overwrite struct {
	ID    string `json:"id"`
	Type  int    `json:"type"`
	Allow Perm   `json:"allow"`
	Deny  Perm   `json:"deny"`
}

// RawChannel is a channel as Discord lists it.
type RawChannel struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Type       int         `json:"type"`
	Position   int         `json:"position"`
	ParentID   *string     `json:"parent_id"`
	Overwrites []Overwrite `json:"permission_overwrites"`
}

// Role is a server role.
type Role struct {
	ID          string `json:"id"`
	Permissions Perm   `json:"permissions"`
	Position    int    `json:"position"`
}

// ServerLayout is what the bot learns about a server to find its channels.
// BotUserID is the bot's own member id, for member overwrites.
type ServerLayout struct {
	Channels   []RawChannel
	Roles      []Role
	BotRoleIDs []string
	BotUserID  string
}

// Channel is a channel a member can pick for posts.
type Channel struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Category *string `json:"category"`
}

// PostableChannels lists the text channels where the bot may post an embed (postPermissions,
// the same set channel_unusable names):
// uncategorised first, then by category position, each by position, ties by id.
func PostableChannels(serverID string, l ServerLayout) []Channel {
	base := basePermissions(serverID, l)
	categories := map[string]RawChannel{}
	for _, c := range l.Channels {
		if c.Type == channelCategory {
			categories[c.ID] = c
		}
	}
	var picked []RawChannel
	for _, c := range l.Channels {
		if c.Type == channelText && channelPermissions(base, serverID, l, c)&postPermissions == postPermissions {
			picked = append(picked, c)
		}
	}
	sort.SliceStable(picked, func(i, j int) bool { return before(picked[i], picked[j], categories) })
	out := make([]Channel, 0, len(picked))
	for _, c := range picked {
		out = append(out, Channel{ID: c.ID, Name: c.Name, Category: categoryName(c, categories)})
	}
	return out
}

// basePermissions is @everyone's permissions OR the bot's roles'.
func basePermissions(serverID string, l ServerLayout) Perm {
	bot := map[string]bool{}
	for _, id := range l.BotRoleIDs {
		bot[id] = true
	}
	var perms Perm
	for _, r := range l.Roles {
		if r.ID == serverID || bot[r.ID] {
			perms |= r.Permissions
		}
	}
	return perms
}

// channelPermissions applies the channel's overwrites in Discord's order:
// @everyone, then role overwrites (deny all, then allow all), then the member.
func channelPermissions(base Perm, serverID string, l ServerLayout, c RawChannel) Perm {
	if base&PermAdministrator != 0 {
		return ^Perm(0)
	}
	perms := base
	bot := map[string]bool{}
	for _, id := range l.BotRoleIDs {
		bot[id] = true
	}
	var roleDeny, roleAllow Perm
	var member *Overwrite
	for i, o := range c.Overwrites {
		switch {
		case o.Type == OverwriteRole && o.ID == serverID:
			perms = perms&^o.Deny | o.Allow
		case o.Type == OverwriteRole && bot[o.ID]:
			roleDeny |= o.Deny
			roleAllow |= o.Allow
		case o.Type == OverwriteMember && o.ID == l.BotUserID:
			member = &c.Overwrites[i]
		}
	}
	perms = perms&^roleDeny | roleAllow
	if member != nil {
		perms = perms&^member.Deny | member.Allow
	}
	return perms
}

func categoryName(c RawChannel, categories map[string]RawChannel) *string {
	if c.ParentID == nil {
		return nil
	}
	parent, ok := categories[*c.ParentID]
	if !ok {
		return nil
	}
	name := parent.Name
	return &name
}

// before orders uncategorised channels first, then by category, then position.
func before(a, b RawChannel, categories map[string]RawChannel) bool {
	ca, cb := categoryOf(a, categories), categoryOf(b, categories)
	switch {
	case (ca == nil) != (cb == nil):
		return ca == nil
	case ca != nil && ca.ID != cb.ID:
		if ca.Position != cb.Position {
			return ca.Position < cb.Position
		}
		return ca.ID < cb.ID
	case a.Position != b.Position:
		return a.Position < b.Position
	}
	return a.ID < b.ID
}

func categoryOf(c RawChannel, categories map[string]RawChannel) *RawChannel {
	if c.ParentID == nil {
		return nil
	}
	if parent, ok := categories[*c.ParentID]; ok {
		return &parent
	}
	return nil
}
