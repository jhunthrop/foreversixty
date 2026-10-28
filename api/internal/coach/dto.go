package coach

import (
	"slices"
	"time"
)

// Answer statuses beyond the pack's own (§17.2).
const statusWriting = "writing"

// teaserDTO is everything a free owner is sent (§17.2, experience spec §10.2), plus the
// spec's display label the strip and panel title need before any review text.
type teaserDTO struct {
	Status       string          `json:"status"`
	Access       Access          `json:"access"`
	Spec         string          `json:"spec"`
	FindingCount int             `json:"finding_count"`
	FixFirst     *teaserFirstDTO `json:"fix_first"`
}

type teaserFirstDTO struct {
	Title   string `json:"title"`
	Code    Code   `json:"code"`
	SpellID *int64 `json:"spell_id"`
}

// fullDTO is the full and officer answer (§17.2), plus access.
type fullDTO struct {
	ReportID    string               `json:"report_id"`
	FightIndex  int                  `json:"fight_index"`
	PlayerKey   string               `json:"player_key"`
	Spec        string               `json:"spec"`
	Role        string               `json:"role"`
	Size        int                  `json:"size"`
	Kill        bool                 `json:"kill"`
	Status      string               `json:"status"`
	Source      string               `json:"source"`
	Busy        bool                 `json:"busy"`
	Cap         *capDTO              `json:"cap,omitempty"`
	Access      Access               `json:"access"`
	Persona     personaDTO           `json:"persona"`
	ReviewedAt  string               `json:"reviewed_at"`
	Review      reviewDTO            `json:"review"`
	Anchors     map[string]anchorDTO `json:"anchors"`
	BasisLabels []string             `json:"basis_labels"`
}

// capDTO says which of the account's caps holds (§15.4): "day" or "month", and the UTC day
// it lifts, as YYYY-MM-DD.
type capDTO struct {
	Period CapPeriod `json:"period"`
	Until  string    `json:"until"`
}

// capOf is the account cap a busy answer carries; nil when none holds.
func capOf(v *Verdict) *capDTO {
	if v == nil || v.Refusal != RefusalAccountCap {
		return nil
	}
	return &capDTO{Period: v.Period, Until: v.Until.Format(time.DateOnly)}
}

type personaDTO struct {
	Label  string        `json:"label"`
	Status PersonaStatus `json:"status"`
}

type reviewDTO struct {
	Opening      string    `json:"opening"`
	FixFirst     *fixDTO   `json:"fixFirst"`
	ThenFix      []thenDTO `json:"thenFix"`
	DoingWell    []wellDTO `json:"doingWell"`
	CouldntCheck []string  `json:"couldntCheck"`
}

type fixDTO struct {
	Finding   string    `json:"finding"`
	Title     string    `json:"title"`
	Why       string    `json:"why"`
	Drill     string    `json:"drill"`
	DrillText string    `json:"drillText"`
	Anchors   []string  `json:"anchors"`
	Links     []linkDTO `json:"links"`
}

type thenDTO struct {
	Finding string    `json:"finding"`
	Title   string    `json:"title"`
	Why     string    `json:"why"`
	Anchors []string  `json:"anchors"`
	Links   []linkDTO `json:"links"`
}

type wellDTO struct {
	Strength string   `json:"strength"`
	Text     string   `json:"text"`
	Anchors  []string `json:"anchors"`
}

// linkDTO is a link with the pack's sentence, word for word (experience spec §10.4).
type linkDTO struct {
	ID       string  `json:"id"`
	Sentence *string `json:"sentence"`
}

type anchorDTO struct {
	Kind       AnchorKind `json:"kind"`
	SpellID    int64      `json:"spell_id,omitempty"`
	SpellName  string     `json:"spell_name,omitempty"` // the build's name for SpellID, for the chip and the name token
	AtMS       *int64     `json:"at_ms,omitempty"`
	ToMS       *int64     `json:"to_ms,omitempty"`
	Debuff     bool       `json:"debuff,omitempty"`
	TargetGUID string     `json:"target_guid,omitempty"`
}

// served is what one answer carries beside the pack: the review, where it came from and
// when it was written.
type served struct {
	doc        ReviewDoc
	source     string
	status     string
	busy       bool
	cap        *Verdict // the account cap behind busy; nil for the day's spend cap
	persona    Persona
	hasPersona bool
	at         time.Time
}

// teaserOf is the free owner's answer: the count, and the first thing to fix with its
// code and first spell.
func teaserOf(p Pack) teaserDTO {
	t := teaserDTO{Status: string(p.Status), Access: AccessTeaser, Spec: p.Context.SpecLabel, FindingCount: len(p.Findings)}
	for _, f := range p.Findings {
		if f.Rank == 1 {
			t.FixFirst = &teaserFirstDTO{Title: f.Title, Code: f.Code, SpellID: firstSpell(p, f)}
		}
	}
	return t
}

// firstSpell is the finding's first anchor spell, else its first named spell; nil when it
// names none.
func firstSpell(p Pack, f Finding) *int64 {
	for _, id := range f.Anchors {
		for _, a := range p.Anchors {
			if a.ID == id && a.SpellID != 0 {
				spell := a.SpellID
				return &spell
			}
		}
	}
	if len(f.Facts.Spells) > 0 {
		spell := f.Facts.Spells[0].ID
		return &spell
	}
	return nil
}

// fullOf is the full or officer answer for one row.
func fullOf(req fightRequest, l loaded, access Access, sv served) fullDTO {
	c := l.pack.Context
	return fullDTO{
		ReportID: req.reportID, FightIndex: req.fight, PlayerKey: req.player,
		Spec: c.SpecLabel, Role: c.Role, Size: c.Size, Kill: c.Kill,
		Status: sv.status, Source: sv.source, Busy: sv.busy, Cap: capOf(sv.cap), Access: access,
		Persona: personaOf(c.SpecLabel, sv), ReviewedAt: sv.at.UTC().Format(time.RFC3339),
		Review: reviewOf(sv.doc, l.pack), Anchors: anchorsOf(l.pack, l.tables),
		BasisLabels: nonNil(l.pack.BasisLabels),
	}
}

// personaOf is the header: the coach's for a model review, the review's for the rule-only
// rendering, which claims no persona voice (§5.1).
func personaOf(spec string, sv served) personaDTO {
	if sv.source != SourceModel || !sv.hasPersona {
		return personaDTO{Label: Fill(PersonaLabelRules, map[string]string{"Spec": spec}), Status: PersonaFull}
	}
	label := PersonaLabelModel
	if sv.persona.Status == PersonaStarter {
		label = PersonaLabelStarter
	}
	return personaDTO{Label: Fill(label, map[string]string{"Spec": spec}), Status: sv.persona.Status}
}

func reviewOf(doc ReviewDoc, p Pack) reviewDTO {
	out := reviewDTO{Opening: doc.Opening, ThenFix: []thenDTO{}, DoingWell: []wellDTO{},
		CouldntCheck: nonNil(doc.CouldntCheck)}
	if f := doc.FixFirst; f != nil {
		out.FixFirst = &fixDTO{Finding: f.Finding, Title: f.Title, Why: f.Why, Drill: f.Drill,
			DrillText: f.DrillText, Anchors: nonNil(f.Anchors), Links: linksDTO(p, f.Links)}
	}
	for _, t := range doc.ThenFix {
		out.ThenFix = append(out.ThenFix, thenDTO{Finding: t.Finding, Title: t.Title, Why: t.Why,
			Anchors: nonNil(t.Anchors), Links: linksDTO(p, t.Links)})
	}
	for _, w := range doc.DoingWell {
		out.DoingWell = append(out.DoingWell, wellDTO{Strength: w.Strength, Text: w.Text, Anchors: nonNil(w.Anchors)})
	}
	return out
}

func linksDTO(p Pack, ids []string) []linkDTO {
	out := make([]linkDTO, 0, len(ids))
	for _, id := range ids {
		l := linkDTO{ID: id}
		if i := slices.IndexFunc(p.Links, func(pl Link) bool { return pl.ID == id }); i >= 0 {
			sentence := p.Links[i].Sentence
			l.Sentence = &sentence
		}
		out = append(out, l)
	}
	return out
}

// anchorsOf keys the pack's anchors by id, each spell with the build's name for it (the
// class spell data, which the web's own name tables may not hold for a talent). An aura
// anchor on an NPC target (target_guid, which the page's `aura-<target>-<spell>` row is keyed
// by), or whose spell is one of the rule file's DoTs or debuffs, opens the Debuffs tab
// (debuff: true).
func anchorsOf(p Pack, tables Tables) map[string]anchorDTO {
	out := make(map[string]anchorDTO, len(p.Anchors))
	for _, a := range p.Anchors {
		out[a.ID] = anchorDTO{Kind: a.Kind, SpellID: a.SpellID, SpellName: tables.Spells[a.SpellID].Name,
			AtMS: a.AtMS, ToMS: a.ToMS, TargetGUID: a.TargetGUID,
			Debuff: a.Kind == AnchorAura && (a.TargetGUID != "" || debuffAura(tables.Rules, a.SpellID))}
	}
	return out
}

func debuffAura(rules SpecRules, spellID int64) bool {
	for _, a := range rules.Auras {
		if (a.Kind == AuraDot || a.Kind == AuraDebuff) && slices.Contains(a.IDs(), spellID) {
			return true
		}
	}
	return false
}

func nonNil(s []string) []string {
	if s == nil {
		return []string{}
	}
	return s
}
