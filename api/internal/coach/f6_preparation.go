package coach

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	// prepAnchorRef is the outcome's one local anchor, the Buffs tab.
	prepAnchorRef = "#0"
	// prepMajorSize is the raid size at which a missing flask, food or weapon enchant is
	// major; at 10 and 20 it is minor (§6.2 F6).
	prepMajorSize = 40
	// potionCooldownMS is the shared cooldown of combat potions: a fight holds one use at
	// the pull and one more per cooldown, so a short fight is not faulted for the rest.
	potionCooldownMS = 120_000
	mainHandSlot     = 15

	prepFactUses     = "combat potions used"
	prepFactAllowed  = "combat potions the fight could hold"
	prepFactUnused   = "combat potions unused"
	prepFactEnchants = "items without a permanent enchant"
	prepIncomplete   = "{Category}: the coach's {list} list isn't complete yet."

	// The words each missing thing goes by in DetailAlsoMissing's {items}.
	prepMissingFlask    = "a flask"
	prepMissingFood     = "food"
	prepMissingOil      = "a weapon oil or stone"
	prepMissingEnchant  = "a {Slot} enchant"
	prepMissingPotions  = "combat potions"
	prepMissingBlessing = "a blessing"
)

// prepKind is what an issue is about; issues of one kind group into one title.
type prepKind int

const (
	kindFlask prepKind = iota
	kindWeaponOil
	kindFood
	kindEnchant
	kindRaidBuff
	kindPotions
)

// Order within a severity tier: the most severe, then the earliest, names the finding.
const (
	orderFlask = iota
	orderPermanentWeapon
	orderWeaponOil
	orderFood
	orderGearEnchant
	orderPotions
	orderRaidBuff // last: a raid-buff title means raid buffs alone (§6.1, never rank 1)
)

// enchantSlotNames labels the gear indexes in rating.DisciplineEnchantSlots, the one
// list of slots both the score and this check read.
var enchantSlotNames = map[int]string{4: "Chest", 7: "Feet", 14: "Back", mainHandSlot: "Main Hand"}

// prepIssue is one thing missing at the pull.
type prepIssue struct {
	order    int
	severity Severity
	kind     prepKind
	name     string       // the slot (enchants) or the buff as the copy names it
	class    string       // the raid buff's provider class, lower case
	missing  string       // the words for DetailAlsoMissing
	item     string       // a name the finding quotes, for its facts; "" when none
	potions  *potionCount // set for the potion issue
}

// potionCount is the combat potions used against what the fight could hold.
type potionCount struct{ uses, allowed int64 }

// prepCategory is a persistent-buff category the sheet can show and its wording.
type prepCategory struct {
	category   string
	order      int
	kind       prepKind
	missing    string
	incomplete string
}

func prepCategories() []prepCategory {
	return []prepCategory{
		{consumables.CategoryFlask, orderFlask, kindFlask, prepMissingFlask, Fill(prepIncomplete, map[string]string{"Category": "Flask", "list": "flask"})},
		{consumables.CategoryWeaponEnchant, orderWeaponOil, kindWeaponOil, prepMissingOil, Drafts["cannot_weapon_enchant"]},
		{consumables.CategoryFood, orderFood, kindFood, prepMissingFood, FoodIncomplete},
	}
}

// preparation is F6 (§6.2): what the player brought into the pull, read from the
// character sheet's seeded auras and gear. No sheet is "couldn't check", never a claim.
func preparation(in Input, _ Counted) outcome {
	pull := consumables.AtPull(in.Fight, in.Row.GUID, in.Tables.Consumables)
	if !pull.Known {
		return outcome{CouldNotCheck: []string{NoSheet}}
	}
	issues, have, couldNot := categoryIssues(in, pull)
	issues = append(issues, gearIssues(in)...)
	issues = append(issues, raidBuffIssues(in)...)
	potion, potionNote := potionIssue(in, pull)
	issues = append(issues, potion...)
	couldNot = append(couldNot, potionNote...)
	facts := slices.Concat(have, entryNames(pull.Present[consumables.CategoryWorldBuffs]))
	out := outcome{CouldNotCheck: couldNot, Anchors: []Anchor{{Kind: AnchorBuffs}}}
	if len(issues) == 0 {
		out.Strengths = preparedStrengths(have, facts)
		return out
	}
	out.Finding, out.RaidBuffOnly, out.MissingFlask = prepFinding(in, issues, have, facts), allPrepKind(issues, kindRaidBuff), anyPrepKind(issues, kindFlask)
	return out
}

func entryNames(entries []consumables.Entry) []string {
	out := make([]string, len(entries))
	for i, e := range entries {
		out[i] = e.Name
	}
	return out
}

func severityRank(s Severity) int {
	if s == SeverityMajor {
		return 0
	}
	return 1
}

func sizeSeverity(size int) Severity {
	if size >= prepMajorSize {
		return SeverityMajor
	}
	return SeverityMinor
}

// categoryIssues checks flask, weapon oil and food. A category absent from the sheet is
// missing only when the catalogue's list for it is complete; else it is couldn't-check.
// have is the names present.
func categoryIssues(in Input, pull consumables.Pull) (issues []prepIssue, have, couldNot []string) {
	for _, c := range prepCategories() {
		present := pull.Present[c.category]
		switch {
		case len(present) > 0:
			have = append(have, entryNames(present)...)
		case !in.Tables.Consumables.IsComplete(c.category):
			couldNot = append(couldNot, c.incomplete)
		case in.Tables.Consumables.MissedAtPull(pull, c.category):
			issues = append(issues, prepIssue{order: c.order, severity: sizeSeverity(in.Size), kind: c.kind, missing: c.missing})
		}
	}
	return issues, have, couldNot
}

// gearIssues is the permanent enchants missing from the sheet's gear, one per slot
// rating.Enchanted checks. The weapon's is major at every size.
func gearIssues(in Input) []prepIssue {
	var out []prepIssue
	for _, c := range in.Fight.Combatants {
		if c.GUID != in.Row.GUID {
			continue
		}
		for _, slot := range rating.DisciplineEnchantSlots {
			label := enchantSlotNames[slot]
			if slot < len(c.Gear) && c.Gear[slot].ID != 0 && !rating.Enchanted(onlySlot(c, slot)) {
				out = append(out, gearIssue(in.Size, slot, label))
			}
		}
	}
	slices.SortFunc(out, func(a, b prepIssue) int { return strings.Compare(a.name, b.name) })
	return out
}

// onlySlot is c with every gear slot but one emptied, so rating.Enchanted judges that
// slot alone.
func onlySlot(c summary.CombatantRow, slot int) summary.CombatantRow {
	gear := make([]event.Item, len(c.Gear))
	gear[slot] = c.Gear[slot]
	c.Gear = gear
	return c
}

func gearIssue(size, slot int, label string) prepIssue {
	issue := prepIssue{order: orderGearEnchant, severity: sizeSeverity(size), kind: kindEnchant, name: label,
		missing: Fill(prepMissingEnchant, map[string]string{"Slot": label})}
	if slot == mainHandSlot {
		issue.order, issue.severity = orderPermanentWeapon, SeverityMajor
	}
	return issue
}

// potionIssue is the combat potions used against what the fight could hold, a seeded
// pre-pull potion counting as one use. A list the catalogue calls incomplete is
// couldn't-check.
func potionIssue(in Input, pull consumables.Pull) ([]prepIssue, []string) {
	rc := in.Tables.Consumables
	if !rc.IsComplete(consumables.CategoryCombatPotion) {
		return nil, []string{Fill(prepIncomplete, map[string]string{"Category": "Combat potions", "list": "potion"})}
	}
	uses := pull.PotionUses
	if pull.Has(consumables.CategoryCombatPotion) {
		uses++
	}
	allowed := min(rc.CombatPotion.MaxUses, 1+in.Fight.DurationMS/potionCooldownMS)
	if uses >= allowed {
		return nil, nil
	}
	return []prepIssue{{order: orderPotions, severity: SeverityMinor, kind: kindPotions, missing: prepMissingPotions,
		potions: &potionCount{uses, allowed}}}, nil
}

func prepFinding(in Input, issues []prepIssue, have, facts []string) *Finding {
	slices.SortStableFunc(issues, func(a, b prepIssue) int {
		if a.severity != b.severity {
			return severityRank(a.severity) - severityRank(b.severity)
		}
		return a.order - b.order
	})
	for _, i := range issues {
		if i.item != "" {
			facts = append(facts, i.item)
		}
	}
	lead, rest := leadGroup(issues)
	detail := prepDetail(lead, have)
	if len(rest) > 0 {
		missing := make([]string, len(rest))
		for i, r := range rest {
			missing[i] = r.missing
		}
		detail += " " + Fill(DetailAlsoMissing, map[string]string{"items": joinAnd(missing)})
	}
	return &Finding{
		Code:     CodePreparation,
		Severity: issues[0].severity,
		Title:    prepTitle(lead),
		Detail:   detail,
		Basis:    Basis{Label: prepBasis(in, lead[0].kind), Kind: BasisOwnLog},
		Anchors:  []string{prepAnchorRef},
		Facts:    Facts{Items: slices.Compact(slices.Sorted(slices.Values(facts))), Numbers: prepNumbers(issues, lead)},
	}
}

// leadGroup splits the sorted issues into the first one's kind (every enchant or every
// raid buff missing, which one title counts) and the rest, in order.
func leadGroup(issues []prepIssue) (lead, rest []prepIssue) {
	for _, i := range issues {
		if i.kind == issues[0].kind {
			lead = append(lead, i)
		} else {
			rest = append(rest, i)
		}
	}
	return lead, rest
}

func prepTitle(lead []prepIssue) string {
	n, first := strconv.Itoa(len(lead)), lead[0]
	switch first.kind {
	case kindFlask:
		return TitleNoFlask
	case kindFood:
		return TitleNoFood
	case kindWeaponOil:
		return Drafts["title_no_weapon_oil"]
	case kindEnchant:
		if len(lead) > 1 {
			return Fill(TitleNoEnchants, map[string]string{"n": n})
		}
		return Fill(TitleNoWeaponEnchant, map[string]string{"Slot": first.name})
	case kindRaidBuff:
		if len(lead) > 1 {
			return Fill(TitleRaidBuffs, map[string]string{"n": n})
		}
		return Fill(TitleRaidBuff, map[string]string{"Buff": first.name})
	}
	return Fill(TitlePotions, map[string]string{"n": strconv.FormatInt(first.potions.unused(), 10)})
}

func prepDetail(lead []prepIssue, have []string) string {
	first := lead[0]
	switch first.kind {
	case kindFlask:
		return withItems(DetailNoFlask, DetailNoFlaskBare, have)
	case kindFood:
		return withItems(DetailNoFood, DetailNoFoodBare, have)
	case kindWeaponOil:
		return Drafts["detail_no_weapon_oil"]
	case kindEnchant:
		if len(lead) > 1 {
			return Fill(DetailNoEnchants, map[string]string{"slots": joinAnd(names(lead))})
		}
		return Fill(DetailNoWeaponEnchant, map[string]string{"Slot": first.name})
	case kindRaidBuff:
		if len(lead) > 1 {
			return Fill(DetailRaidBuffs, map[string]string{"buffs": joinAnd(missingWords(lead))})
		}
		return upperFirst(Fill(DetailRaidBuff, map[string]string{"Buff": first.missing, "class": first.class}))
	}
	return Fill(DetailPotions, map[string]string{
		"used": strconv.FormatInt(first.potions.uses, 10), "max": strconv.FormatInt(first.potions.allowed, 10),
	})
}

// withItems is format with what the player pulled with, or bare when nothing was.
func withItems(format, bare string, have []string) string {
	if len(have) == 0 {
		return bare
	}
	return Fill(format, map[string]string{"items": joinAnd(have)})
}

func names(issues []prepIssue) []string {
	out := make([]string, len(issues))
	for i, is := range issues {
		out[i] = is.name
	}
	return out
}

func missingWords(issues []prepIssue) []string {
	out := make([]string, len(issues))
	for i, is := range issues {
		out[i] = is.missing
	}
	return out
}

func upperFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func prepBasis(in Input, kind prepKind) string {
	switch kind {
	case kindEnchant:
		return BasisGear
	case kindRaidBuff:
		return BasisRaidBuff
	}
	return Fill(BasisSheet, map[string]string{"roles": roleLabel(in.Row.Role)})
}

func (p potionCount) unused() int64 { return p.allowed - p.uses }

// prepNumbers are the potion counts when potions are short, and the count of items
// without an enchant when the title states it.
func prepNumbers(issues, lead []prepIssue) []NumberFact {
	var out []NumberFact
	if lead[0].kind == kindEnchant && len(lead) > 1 {
		out = append(out, NumberFact{V: float64(len(lead)), Text: strconv.Itoa(len(lead)), Unit: "items", Of: prepFactEnchants})
	}
	for _, i := range issues {
		if p := i.potions; p != nil {
			out = append(out,
				NumberFact{V: float64(p.uses), Text: strconv.FormatInt(p.uses, 10), Unit: "uses", Of: prepFactUses},
				NumberFact{V: float64(p.allowed), Text: strconv.FormatInt(p.allowed, 10), Unit: "uses", Of: prepFactAllowed},
				NumberFact{V: float64(p.unused()), Text: strconv.FormatInt(p.unused(), 10), Unit: "uses", Of: prepFactUnused})
		}
	}
	return out
}

func roleLabel(role string) string {
	switch role {
	case RoleHealer:
		return "healers"
	case RoleTank:
		return "tanks"
	}
	return "damage dealers"
}

// preparedStrengths is the one strength when nothing is missing: the flask named when
// there is one, else the draft's general sentence. Nothing present is no strength.
func preparedStrengths(have, facts []string) []Strength {
	if len(have) == 0 {
		return nil
	}
	sentence := Drafts["strength_consumables"]
	if flask := firstFlask(have); flask != "" {
		sentence = Fill(StrengthAtPull, map[string]string{"Item": flask})
	}
	return []Strength{{Sentence: sentence, Facts: Facts{Items: slices.Clone(facts)}, Anchors: []string{prepAnchorRef}}}
}

// firstFlask is the first of have the catalogue lists as a flask, "" if none.
func firstFlask(have []string) string {
	flasks := entryNames(consumables.Load().AllRoles().Flask)
	for _, h := range have {
		if slices.Contains(flasks, h) {
			return h
		}
	}
	return ""
}

// The raid buffs: flagged only when a provider's class is in the roster, the buff does
// something for this player, and (blessings) the paladins there were could have given it.

const blessingClass = "Paladin"

// blessingNames are the paladin blessings, one given by each paladin; blessingWord names
// a shortfall when the log cannot say which was owed.
var blessingNames = []string{"Blessing of Kings", "Blessing of Might", "Blessing of Wisdom"}

const blessingWord = "blessing"

func raidBuffIssues(in Input) []prepIssue {
	var out []prepIssue
	var blessings []consumables.RaidBuffState
	for _, s := range consumables.RaidBuffsAtPull(in.Fight, in.Row.GUID, in.Tables.RaidBuffs) {
		switch {
		case !s.Known || !s.ProviderInRoster || !buffMatters(in.Row, s.Buff.Name):
		case slices.Contains(blessingNames, s.Buff.Name):
			blessings = append(blessings, s)
		case !s.Present:
			out = append(out, buffIssue(s.Buff.Name, s.Buff.Class))
		}
	}
	return append(out, blessingIssue(in, blessings)...)
}

func buffIssue(name, class string) prepIssue {
	return prepIssue{order: orderRaidBuff, severity: SeverityMinor, kind: kindRaidBuff, item: name,
		name: name, missing: name, class: strings.ToLower(class)}
}

// blessingIssue is a shortfall in the blessings the player holds against the paladins in
// the roster: each gives one. The log does not say which one was owed, so a blessing is
// named only when it is the single one the player lacks. relevant holds the blessings
// that matter to this player.
func blessingIssue(in Input, relevant []consumables.RaidBuffState) []prepIssue {
	held, absent := 0, []string{}
	for _, s := range relevant {
		if s.Present {
			held++
		} else {
			absent = append(absent, s.Buff.Name)
		}
	}
	if held >= min(paladinCount(in), len(relevant)) {
		return nil
	}
	if len(absent) == 1 {
		return []prepIssue{buffIssue(absent[0], blessingClass)}
	}
	issue := buffIssue(blessingWord, blessingClass)
	issue.missing = prepMissingBlessing
	return []prepIssue{issue}
}

func paladinCount(in Input) int {
	seen := map[string]bool{}
	for _, r := range in.Fight.Roster {
		if r.Class == blessingClass {
			seen[r.GUID] = true
		}
	}
	return len(seen)
}

// buffMatters is whether the buff does anything for this player: Wisdom and Intellect
// are for mana users, Might and Battle Shout for attack-power users. A buff not listed
// here is never claimed.
func buffMatters(row summary.RosterRow, buff string) bool {
	switch buff {
	case "Power Word: Fortitude", "Mark of the Wild", "Blessing of Kings":
		return true
	case "Arcane Intellect", "Blessing of Wisdom":
		return usesMana(row)
	case "Blessing of Might", "Battle Shout":
		return usesAttackPower(row)
	}
	return false
}

func usesMana(row summary.RosterRow) bool {
	switch row.Class {
	case "Mage", "Priest", "Warlock", "Paladin", "Shaman", "Hunter":
		return true
	case "Druid":
		return !strings.HasPrefix(row.Spec, "Feral")
	}
	return false
}

func usesAttackPower(row summary.RosterRow) bool {
	switch row.Class {
	case "Warrior", "Rogue", "Hunter":
		return true
	case "Paladin":
		return row.Spec != "Holy"
	case "Druid":
		return strings.HasPrefix(row.Spec, "Feral")
	case "Shaman":
		return row.Spec == "Enhancement"
	}
	return false
}
