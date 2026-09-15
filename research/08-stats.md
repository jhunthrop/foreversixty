# WoW: Forever — Stats and Combat Formulas (for the simulator engine)

Research date: **2026-09-14**. Beta client ships **2026-09-17**; everything below is pre-beta.
Scope: every fact that becomes code in an attack table, a stat-dependency graph, or a rating
conversion. Companion to [07-simulator.md](07-simulator.md) §5 and [01-official-facts.md](01-official-facts.md) §13.

## 0. How to read this

**Source kinds** are marked on every fact: `blizzard` (Blizzard-published text), `datamined`
(pulled from a live data endpoint), `community` (an outlet, a fan site, a player), `site` (our own
verification). `SINGLE-SOURCE` marks a fact only one outlet carries.

**Reliability tiers**, in descending order, because this document mixes them:

| Tier | What it is | How much weight |
|---|---|---|
| A | Blizzard's own published prose | Build on it |
| B | A panel slide photographed and transcribed with numbers | Build on it, re-verify at beta |
| C | A live data endpoint (Wowhead's Forever talent data) | Wording reliable, per-rank scaling not |
| D | A fan dataset read off stream footage | Corroboration only |
| E | A player's inference from a screenshot | Note it, never code it |

**Two things this document does not do:** it does not fill gaps with vanilla 1.12 knowledge, and it
does not turn an absence of evidence into a finding. Where Forever has said nothing, it says so, and
§11 gives the exact beta test that settles it.

**Attribution required by licence:** the fan dataset used in §2, §5 and §6 is
`https://talentsforever.com/data.json`, CC BY 4.0, credit string *"Data from talentsforever.com
(https://talentsforever.com)"*. Any page on foreversixty.gg that republishes it needs that link.

### Verifications I ran today (site)

| Check | Result | When |
|---|---|---|
| `nether.wowhead.com/forever/data/gear-planner?dv=100` | 200, 5,058,446 bytes, 9,107 items, **all `versionNum: 11300`** (patch 1.13.0) | 2026-09-14 |
| Item stat keys in that payload | still the **split Era keys** `mlehitpct` / `rgdhitpct` / `splhitpct`, `mlecritstrkpct` / `rgdcritstrkpct` / `splcritstrkpct`; **no unified hit key, no expertise key** | 2026-09-14 |
| `baseStats.stats`, `critPhysical`, `critSpell`, `randPropPoints`, `reforgeStats` in that payload | **all empty or all zero** — no conversion tables published | 2026-09-14 |
| `X'caliboar` (item 10758) tooltip in the Forever env | still the Era item: +20 Strength, +8 Stamina, no creature-type effect | 2026-09-14 |
| **`nether.wowhead.com/forever/data/talents-classic?dv=100`** | **200, 228,707 bytes — real Forever talent data**, 470 talents, 27 trees, new spell IDs (105726, 110876, …) | 2026-09-14 |
| `wago.tools/builds` | **no Forever product branch**. `wow_classic_titan` is at 3.80.2.69815 | 2026-09-14 |
| `wow_classic_titan` 3.80.2.69815 `ChrRaces` + `SpellName` | 21 races, **no Skyborne, no Seal of Fury, no Drain Hope** — re-confirmed **not** Forever | 2026-09-14 |

So the answer to "did Wowhead's Forever environment change since Sept 14?" is **the item environment
has not** — it is still an Era scaffold — **but a Forever talent endpoint that nobody had found is
live**, and it is the best data source that exists before the beta. It is the backbone of §2 and §5.

---

## 1. The three load-bearing claims

### 1.1 Can damage-over-time effects critically strike? — **Probably yes, and the engine should be built for it, but no source states it**

**Verdict: REFUTED as stated, REPLACED by something stronger.**

The claim entered our research via [06-since-announcement.md](06-since-announcement.md) §4, which
lists a Warlock talent as *"Malevolence (DoTs can crit)"*, single-sourced to Power Up Gaming. **That
is wrong.** Two independent datasets give Malevolence's actual text, and it is an ordinary crit-chance
talent:

> **Malevolence** — "Increases the critical effect chance of your Shadow spells by 1%." (5 ranks)

— Wowhead Forever talent data, spell **110876**, `nether.wowhead.com/forever/data/talents-classic?dv=100` (datamined, verified 2026-09-14); identical text in `talentsforever.com/data.json` (community).

`06-since-announcement.md` §4 should be corrected.

**But the real evidence is a different talent, and it is better evidence.**

> **Pandemic** — "Increases the critical strike damage bonus of your Corruption, Bane of Agony, Bane of
> Doom, Drain Soul, Drain Life, Siphon Life, and Drain Hope spells by 33%." (3 ranks: 33 / 66 / 100%)

— Wowhead Forever talent data, spell **105917** (datamined, verified 2026-09-14); identical in
talentsforever (community); also carried by classicwowforever.com (community).

Every spell in that list deals **only** periodic damage. Corruption, Bane of Agony and Bane of Doom
have no direct-damage component at all; Drain Life, Drain Soul and Siphon Life tick; Drain Hope is new
and is explicitly a tick ("dealing 52 Shadow damage every 1 sec", spell 105909). A talent that raises
the *critical strike damage bonus* of a list consisting exclusively of periodic effects is inert unless
those periodic effects can crit. In vanilla 1.12 this talent could not exist.

**This is an inference from tooltip text, not a statement.** Label it that way on the site. One
community site makes the same reading and attaches a caution worth repeating:

> "This is not a critical strike chance bonus, and should not be read as one."
> — classicwowforever.com (community)

The caution is correct and does not weaken the inference: a *damage* bonus on a periodic-only spell
list still presupposes that periodic ticks produce crits.

**Scope — what is and is not established:**

- **All periodic damage, or some?** Unknown. Pandemic is Warlock Affliction and names seven specific
  spells. Nothing indicates a global rule. Two generic periodic-damage talents exist and are silent on
  crit: Warlock **Malediction** (105923) "Increases all periodic damage done by your Warlock spells by
  1%" and Druid **Genesis** (104924) "Increases the periodic damage and healing done by your spells and
  abilities by 1%".
- **Do periodic heals crit?** No evidence either way. Druid **Improved Regrowth** (104913) raises "the
  critical effect chance of your Regrowth spell" — Regrowth is direct + HoT, so this does not separate
  the two. No talent anywhere references a critical periodic *heal*.
- **What multiplier?** Not stated. Pandemic's phrasing ("critical strike damage **bonus** … by 33%")
  matches the vanilla convention where the bonus is the part above 1.0, so a Forever DoT crit most
  likely starts from the 1.5× spell baseline (a 50% bonus) and Pandemic at rank 3 (+100%) takes those
  seven spells to 2.5×. **That arithmetic is inference, not a source.**
- **Interaction with existing DoT talents?** Not stated anywhere.

**Counter-evidence worth holding in view.** Forever's talent text uses the qualifier
**"non-periodic"** in nine places — Mage **Hot Streak** (105786) "Your *non-periodic* critical strikes
with Fireball, Frostfire Bolt, Fire Blast, and Scorch", Mage **Combustion** (105781) "until you have
caused 4 *non-periodic* critical strikes with Fire spells", Priest **Inspiration** (105860) "Your
*non-periodic* critical heals", Druid **Nature's Grace** (104934) "All *non-periodic* spell criticals",
Druid **Primal Fury** (104947), Paladin **Reckoning** (105627) "*non-periodic* critical strike", Priest
**Prayer of Mending** (105851) "*non-periodic* healing". A designer only writes "non-periodic critical
strike" into seven talents if periodic critical strikes are a thing that would otherwise trigger them.
**This is the strongest corroboration in the document**, and it is also inference.

Note that most of these qualifiers already exist in TBC/WotLK-era text for the same talents, so their
presence is partly inherited wording rather than a fresh Forever decision. Treat the two inferences
(Pandemic, and the "non-periodic" qualifier density) as mutually reinforcing but neither as proof.

**Blizzard has said nothing.** A full grep of all eight Blizzard Forever articles returns **zero**
occurrences of "periodic", "tick", or "damage over time" (blizzard, verified across
news.blizzard.com Forever articles). No blue post addresses it.

**Engine call:** build the DoT tick path so periodic crit is a per-spell flag with its own multiplier,
defaulting to off, and flip it on for the Warlock Affliction set on day one of beta. Do not hardcode
vanilla's "DoTs never crit".

### 1.2 Unified hit — **CONFIRMED, and it is one flat percentage, not a rating**

**Verdict: CONFIRMED.** Blizzard, verbatim:

> "*World of Warcraft: Forever* keeps the familiar stats from Classic, but reworks several of them and
> adds new ones so gearing choices can vary more based on role, activity, and playstyle. **Spell,
> melee, and ranged hit chance are combined, as are spell, melee, and ranged critical chance.** That
> helps hybrid specializations such as Retribution Paladins and Enhancement Shaman, while also letting
> Warriors, Hunters, and Rogues benefit when improving the reliability of tools such as Taunts,
> Poisons, and Traps."

— Blizzard, "World of Warcraft: Forever Deep Dive Panel Recap",
https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (blizzard)

**One rating or one percentage?** **One flat percentage.** The panel showed item cards, and Output Lag
photographed and transcribed them:

> "Lionheart Helm is a level 56 plate helm with +18 Strength and **2.0 percent to both hit and critical
> strike chance**."

— Output Lag, "World of Warcraft Forever unifies hit and crit stats and gives healing gear bonus
damage", Steven Mills, 2026-09-13,
https://outputlag.com/news/world-of-warcraft-forever-unifies-hit-and-crit-stats-and-gives-healing-gear-bonus-damage/ (community, Tier B — transcribed from the slide)

This is the **single most important engine fact in the document** and it is covered in §3.

**Is it literally one stat?** Yes, and the talent trees confirm it from the other direction. Forever
rewrote the class talents that used to grant split hit or resist-reduction into one unified hit stat.
All from Wowhead's Forever talent data (datamined, verified 2026-09-14), rank 1 quoted:

| Class / tree | Talent (spell id) | Forever text | Classic text |
|---|---|---|---|
| Paladin Protection | Precision (105638) | "Increases your chance to hit with **all spells and attacks** by 1%." | "…with melee weapons by 1%." |
| Warrior Fury | Precision (105929) | "Increases your chance to hit with **all abilities and attacks** by 1%." | *(new)* |
| Rogue Combat | Precision (105737) | "Increases your chance to hit with **all attacks and Poisons** by 1%." | "…with melee weapons by 1%." |
| Shaman Restoration | Tidal Focus (104735) | "Reduces the Mana cost of your healing spells by 1% and increases your chance to hit with **all spells and attacks** by 1%." | mana cost only |
| Warlock Affliction | Suppression (105925) | "Increases your chance to hit with **all spells and attacks** by 1% and reduces all threat you generate by 4%." | "Reduces the chance for enemies to resist your Affliction spells by 2%." |
| Druid Balance | Nature's Reach (104929) | "…increases the chance for **all your spells and attacks** to hit by 2%." | range only |
| Hunter Survival | Surefooted (104987) | "Increases your **hit chance** by 1%…" | "Increases hit chance by 1%…" |
| Hunter Survival | Survival Tactics (104988) | "Increases your chance to hit with your **Trap and Feign Death** abilities by 5%." | *(new)* |
| Mage Arcane | Arcane Focus (105814) | "…hit with your **Arcane spells** by 1%." | "Reduces the chance that the opponent can resist your Arcane spells by 2%." |
| Mage Frost | Elemental Precision (105778) | "…hit with **Frost and Fire spells** by 1%." | resist reduction |
| Priest Shadow | Shadow Focus (110851) | "…hit with your **Shadow spells** by 1%." | resist reduction |
| Paladin Holy | Divine Precision (105324) | "…hit with **Holy spells** by 6%." | *(new)* |
| Priest Discipline | Holy Precision (110852) | "…hit with **Holy spells** by 6%." | *(new)* |

Two structural readings fall out, and both are engine-relevant:

1. **The vanilla "resist chance" talents have been converted into hit talents.** Arcane Focus,
   Elemental Precision, Shadow Focus and Suppression all used to reduce the target's chance to *resist*.
   In Forever they grant *hit*. That is a real mechanical change: binary spell resists and spell miss
   were separate systems in vanilla, and Forever appears to have folded the caster-side lever into hit.
2. **School-scoped hit still exists alongside global hit.** Arcane Focus is Arcane-only, Shadow Focus is
   Shadow-only, Divine/Holy Precision are Holy-only, Survival Tactics is Trap/Feign-Death-only, while
   Precision / Suppression / Tidal Focus / Nature's Reach are global. So the *stat* is unified; the
   *modifiers* are not. The engine needs a global hit stat plus per-school and per-spell hit adders.

**Hit caps.** **Not stated by anyone.** Blizzard gave no conversion numbers and no caps; Output Lag
says so explicitly: "The panel gave no conversion numbers for the unified stats, no per-item cap for
weapon skill". Whether the vanilla 16% spell-hit cap against a +3 raid boss survives is unknown. One
suggestive number: **Divine Precision and Holy Precision each give 6% per rank over 3 ranks = 18%**,
which would be oddly generous against a 16% cap and may indicate a different cap, a different base
spell miss, or simply that Wowhead's rank scaling is wrong (see §0 tier C). Do not plan around it.

**Dual-wield miss penalty — strong evidence it survives.**

> **Dual Wield Specialization** (Warrior Fury, spell 105933) — "Increases your off-hand weapon damage by
> 5%, off-hand Rage generation by 20%, and **chance to hit with off-hand attacks by 2%**." (5 ranks, to
> +10% off-hand hit)

— Wowhead Forever talent data (datamined, verified 2026-09-14); identical in talentsforever (community)

A talent granting hit **specifically and only to off-hand attacks** is meaningless unless off-hand
attacks carry their own miss penalty. In vanilla that penalty is a flat +19% miss on white swings.
Forever giving a Fury Warrior up to 10% of it back is a strong signal the penalty exists and has
roughly the same shape. **Inference from tooltip, not a statement.**

**Interaction with weapon skill — this is the gap, and it is still open.**

This was flagged in [07-simulator.md](07-simulator.md) §5.3 as "completely unspecified", and after a
full pass it remains **the largest unresolved mechanic in the game**. What is now known:

- Blizzard: "**Weapon skill still works as it always has**, but items with weapon skill offer less of it
  per item so they are not always the obvious best choice." (blizzard, Deep Dive recap). "Works as it
  always has" is the only statement about its mechanics, and it is doing enormous work: in vanilla,
  weapon skill reduces miss, reduces dodge, reduces the glancing-blow penalty, reduces crit
  suppression, and gates the level-based miss table. If all of that is literally unchanged, then weapon
  skill and the new unified hit stat operate on the same miss roll through two different paths.
- Weapon skill has been **stripped out of talents and racials entirely.** Across all 470 Forever
  talents the strings "weapon skill" and "Weapon Skill" appear **zero** times (datamined, verified
  2026-09-14). The Rogue talent that granted it now grants dodge/parry reduction instead (§1.3). And:
  > "Racials that used to grant Weapon Skill now grant critical strike chance with the matching weapon
  > type" — Icy Veins (community), corroborated by every racial in the talentsforever set: Human Sword
  > Specialization is now "+2% critical strike chance with all spells and attacks, auto attacks
  > included, while you have a sword or two-handed sword equipped"; Dwarf Mace Specialization and Orc
  > Axe Specialization are "+1% critical strike chance".
- Weapon skill **survives only on items**, and the per-item amount is now tiny. The panel card:
  > "Edgemaster's Handguards, level 44 mail, reduce the chance to be dodged or parried by 1.0 percent
  > and add **+1 Weapon Damage plus +1 to Axes, Daggers, and Swords**."
  — Output Lag (community, Tier B). The vanilla Edgemaster's Handguards give **+7** to each of those
  three skills. Forever gives **+1**. That is the "less of it per item" quantified, and it is a 7×
  reduction.

**What is still unknown:** whether unified hit and weapon skill both feed the same miss number and in
what order; whether weapon skill still reduces dodge (and therefore double-dips with the new expertise
stat, see §1.3); whether it still governs glancing-blow damage and crit suppression. §11.2 gives the
test.

### 1.3 Expertise — **CONFIRMED as a mechanic, name only half-confirmed, and it STACKS with weapon skill**

**Verdict: CONFIRMED that it reduces both dodge and parry. No rating conversion exists because there is
no rating.**

Blizzard's own prose describes the mechanic without naming it:

> "New stats broaden those choices further. **Some items can reduce the chance for attacks to be parried
> or dodged**, while spellcaster-oriented weapons now grant spell damage and healing…"

— Blizzard Deep Dive recap (blizzard)

**Both dodge and parry: confirmed**, by Blizzard's own wording above and independently by the talent
data:

> **Weapon Expertise** (Rogue Combat, spell 105726) — "Reduces the chance for your attacks to be
> **Dodged or Parried** by 1%." (2 ranks → 2%). Classic text: "Increases your skill with Sword, Fist and
> Dagger weapons by 3."

— Wowhead Forever talent data (datamined, verified 2026-09-14); identical in talentsforever (community)

That single talent is the cleanest evidence in the document: a talent that granted **weapon skill** in
vanilla now grants **dodge/parry reduction** in Forever, expressed as a flat percentage.

**The name.** Wowhead's liveblog states it flatly:
> "Weapon skill was a little too easy to get and a little too strong in Classic. It's a little harder to
> get now. **Expertise has also been added.**"
— https://www.wowhead.com/forever/news/world-of-warcraft-forever-deep-dive-liveblog-382862 (community)

Output Lag reports the panel slide listed five stats — "**Hit Chance, Crit Chance, Weapon Skill, Bonus
Healing, and Expertise**" — but adds a caveat worth carrying:
> "The panel … **did not confirm Expertise as the in-game name for the dodge and parry reduction**."
(community)

Blizzard's published *articles* never use the word: a grep of all eight returns **zero** occurrences of
"Expertise" (blizzard). So: the word is on the slide and in the liveblog; the binding of that word to
this specific mechanic is not stated. Wowhead's own Forever talent data calls the Rogue talent "Weapon
Expertise" and describes the effect without using "expertise" as a noun. Use "expertise" as our term,
note the caveat once, move on.

**Rating conversion: there is none, because Forever does not use ratings.** The later-expansion figure
of 3.94 rating per expertise point at level 80, and 4 expertise = 1% dodge/parry reduction, **does not
apply**. Forever puts the reduction on the item as a flat percentage:

> "Edgemaster's Handguards, level 44 mail, **reduce the chance to be dodged or parried by 1.0 percent**"
— Output Lag (community, Tier B)

So the unit on the item is a percentage, and the Rogue talent's unit is a percentage. The engine stores
expertise as a percentage and subtracts it from dodge and parry directly. See §3.

**Does it replace or stack with weapon skill's dodge/parry reduction? — IT STACKS, and the evidence is
one item.** Edgemaster's Handguards carries **both**: 1.0% dodge/parry reduction **and** +1 to Axes,
Daggers and Swords, on the same card. If expertise had replaced weapon skill's dodge reduction,
Blizzard would not have shipped both on one glove and then told us weapon skill "still works as it
always has". They coexist.

**The consequence for the engine is a real risk.** If weapon skill still reduces dodge *and* expertise
reduces dodge, a melee character has two independent dodge-reduction paths and the ordering matters for
the single-roll table. Nobody has said how they compose. §11.3 gives the test.

---

## 2. Ratings or flat percentages? — **FLAT PERCENTAGES. This is the major engine decision and it is settled.**

Vanilla 1.12 had no rating system for most stats; Classic added a few; TBC introduced ratings wholesale.
Forever's model is **vanilla's**: item stats are flat percentages at level 60.

**Evidence, strongest first:**

1. **The panel item cards.** Lionheart Helm: "+18 Strength and **2.0 percent** to both hit and critical
   strike chance". Edgemaster's Handguards: "reduce the chance to be dodged or parried by **1.0
   percent**". — Output Lag (community, Tier B, transcribed from photographed slides). Percent signs on
   the item, not rating integers.
2. **Every talent in the game speaks in percentages.** Across 470 Forever talents there is not a single
   reference to a rating. Hit talents give "1%", parry talents give "1%", dodge talents give "1%",
   Weapon Expertise gives "1%". The strings "rating" (as a stat), "resilience", "armor penetration" and
   "spell penetration" appear **zero** times. (datamined, verified 2026-09-14)
3. **Racials speak in percentages.** Tauren Endurance "+5% total Health, **+1% Hit Chance**"; Human Sword
   Specialization "**+2% critical strike chance**"; Skyborne Wind Blessed "Passive. **1% increased**
   melee, ranged, and spellcasting Haste". (talentsforever, community)
4. **Blizzard said nothing to the contrary** — no Blizzard article mentions ratings, conversions or caps
   (blizzard, verified across all eight Forever articles).

**One qualification.** No source states the *level-60 base* values that percentages sit on top of: base
crit per class, base miss, base dodge, base parry. Those are conversions from primary stats and from the
level table, not item stats, and they are unpublished. See §7 and §11.

**Engine consequence:** `sim/core/base_stats_auto_gen.go` in `wowsims/classic` already encodes exactly
this model —

```go
// Crit/Hit/Haste ratings are straight percentage values in classic
const HasteRatingPerHastePercent = 1
const CritRatingPerCritChance = 1
const MeleeHitRatingPerHitChance = 1
const SpellHitRatingPerHitChance = 1
const ExpertiseRatingPerExpertiseChance = 1
```

— so the fork inherits the right model for free. This removes what would have been the single largest
piece of engine work. See §12.

---

## 3. Primary stats

Blizzard published **nothing** on Strength, Agility, Stamina, Intellect or Spirit as itemization
concepts. The word "Strength" appears twice in the Deep Dive recap, both times meaning "class strengths
and weaknesses" (blizzard). There is no statement that any primary-stat conversion changed.

What exists is indirect and thin:

| Stat | Forever evidence | Kind |
|---|---|---|
| Strength | Appears on item cards: Lionheart Helm +18 Strength; X'caliboar +20 Strength | community (Tier B) |
| Agility | No Forever statement of any kind. Output Lag's summary asserts "Agility still buys melee crit and Intellect still buys spell crit at their own per-class rates, and base crit is untouched" — but that sentence is the **writer's framing**, not a quoted designer line, and should be treated as tier E | community, INFERENCE |
| Stamina | Appears on item cards (X'caliboar +7 Stamina, Hide of the Wild +8 Stamina). No conversion statement | community |
| Intellect | Multiple talents convert it — see §7. Appears on cards (Hide of the Wild +10 Intellect) | datamined |
| Spirit | Human racial "The Human Spirit: +5% Spirit"; Druid Living Spirit "Increases your Spirit by 5%"; Priest Spiritual Guidance converts it — see §7 | datamined / community |

**Per-class differences: nothing published.** No source gives strength-to-attack-power per class,
agility-to-crit per class, stamina-to-health, or the intellect-to-spell-crit table. Wowhead's Forever
gear planner would normally carry these in its `baseStats`, `critPhysical` and `critSpell` sections —
**I checked today and all three are empty or all-zero** (site, verified 2026-09-14). §11.5.

**One primary-stat mechanic did change**, and it is a defensive one:
> Stoneform's "defensive bonus now **reduces Physical damage taken instead of increasing Armor**, making
> it useful to more classes." — Blizzard Deep Dive recap (blizzard)

---

## 4. Secondary and combat stats

Consolidated table. "Forever status" is what the evidence supports; blank cells are genuine unknowns.

| Stat | In Forever? | Unit | Evidence |
|---|---|---|---|
| **Hit** | Yes, **unified** across spell/melee/ranged, and covers Taunts, Poisons, Traps | flat % | blizzard (Deep Dive) + item card 2.0% (community Tier B) + 13 talents (datamined) |
| **Crit** | Yes, **unified** across spell/melee/ranged | flat % | blizzard + item card 2.0% + racials (community) |
| **Expertise** (dodge/parry reduction) | Yes. Both dodge and parry | flat % | blizzard (mechanic) + slide name (community) + item card 1.0% + Rogue talent 1%/rank (datamined) |
| **Weapon skill** | Yes, **items only**, ~7× less per item (+1 where vanilla gave +7) | integer skill points | blizzard + item card (community Tier B) + zero talent/racial occurrences (datamined) |
| **Haste** | Yes, exists; still described as three things ("melee, ranged, and spellcasting Haste") | flat % | Skyborne racial +1% (community); Troll Berserking +10% "spellcasting and attack speed"; Shaman Rage of the Farseer +30% melee attack and spell casting speed (datamined). **No Blizzard statement; the word "haste" appears 0 times in Blizzard's articles** |
| **Attack power** | Yes | flat points | item card Graverobber's Shovel +12 AP (community); Trueshot Aura +30 RAP (datamined); Blizzard names a campfire "Sharpening Wheel that increases Attack Power" (blizzard) |
| **Ranged attack power** | Yes, distinct from melee AP | flat points | Trueshot Aura "Increases the **Ranged Attack Power** of party members within 45 yards by 30" (datamined) |
| **Spell damage** | Yes | flat points | Hide of the Wild "damage by up to 14" (community Tier B) |
| **Spell power** | **Not a Forever term.** Blizzard says "spell damage and healing" throughout; "spell power" appears 0 times in Blizzard articles. But Orc Blood Fury is "Increases Attack Power and **Spell Power** by 10%" (community) | — | mixed; treat "spell damage" as canonical |
| **Healing power / bonus healing** | Yes, and **carries ⅓ as much bonus damage** | flat points | blizzard + two numeric confirmations, §6 |
| **Mana per 5** | Yes | flat points | Blessing of Wisdom "restoring 24 mana every 5 seconds" (community). No item-side statement |
| **Spirit regeneration** | Yes, and the mana-while-casting talent family is intact (Meditation, Arcane Meditation, Mindfulness, Reflection all "Allows 17% of your Mana regeneration to continue while casting") | % of regen | datamined |
| **Armor penetration** | **Not as a stat.** But **percentage armor ignore exists as an effect** on three talents | % ignore | datamined — see below |
| **Resilience** | **NO. Explicitly excluded.** | — | Zierhut, verbatim: "although I'll say **we're not adding resilience**." (community Tier B, Output Lag) |
| **Defense** | Yes, still a **skill** with base = 5 × level | integer skill points | Warrior/Paladin **Anticipation** (105975 / 105636) "Increases your **Defense Skill** by 4"; Druid Thick Hide (104942) "…for each point of defense skill beyond **five times your level**" (datamined) |
| **Dodge** | Yes | flat % | Druid Natural Reaction +1%/rank, Rogue Lightning Reflexes +1%/rank, Shaman Anticipation +2%/rank, Night Elf Quickness +1% (datamined/community) |
| **Parry** | Yes | flat % | Deflection exists in four trees at 1–2%/rank (datamined) |
| **Block** | Yes | flat % | Warrior Shield Specialization "+1% chance to Block"; Paladin Holy Shield "+20% block for 10 sec" (datamined) |
| **Block value** | Yes | flat points | Warrior Shield Slam "causing 421 to 439 damage, **increased by your Block Value**" (datamined). Present in the Era item payload as `blockamount` |
| **Resistances** | Yes, all five schools, flat points | flat points | Paladin resistance auras "30 additional Fire resistance"; Mage Arcane Subtlety "Increases all your resistances by 5"; Mark of the Wild "+7 all resistances" (datamined) |
| **Spell penetration** | Yes, and it is **zone-conditional on at least one item** | flat points | Rune of Perfection, level 20 trinket, "gives spells **2 magical resistance piercing**, plus 3 more in forest and woodland areas" (community Tier B) |

**Percentage armor ignore** is a new-in-Forever effect and needs engine support. Three talents carry it,
all datamined, verified 2026-09-14:

- Warrior Arms **Weaponmaster** (105944, 5 ranks) — a *merged* weapon-specialisation talent: "Axe/Polearm:
  Increases your critical strike chance by 1%. **Mace/Staff: Your attacks ignore 3% of your target's
  armor.** Sword: Your successful melee attacks have a 1% chance to trigger an extra attack." At rank 5,
  5% crit / **15% armor ignore** / 5% extra-attack.
- Rogue Combat **Hack and Slash** (105727, 5 ranks) — same merged shape: "Axe/Sword: … extra attack.
  Dagger/Fist: … critical strike chance by 1%. **Mace: Your attacks ignore 3% of your target's armor.**"
- Rogue Subtlety **Serrated Blades** (105752, 3 ranks) — "Causes your attacks to **ignore 3% of your
  target's Armor** and increases the damage dealt by your Rupture ability by 10%." (→ 9% / 30%)

Two things here matter beyond armor. First, vanilla's separate per-weapon specialisation talents have
been **consolidated into one talent per class** whose effect switches on equipped weapon type — that is
a talent-modelling change, not just a stat change. Second, vanilla's Mace Specialization was a **stun
proc**; in Forever it is **armor ignore**. The engine needs a weapon-subclass-conditional modifier and a
percentage armor-ignore multiplier that composes with the target's armor before mitigation.

---

## 5. The attack table

**Nothing has been said about it. Not one sentence, by anyone.**

Verified absences (blizzard, all eight Forever articles; datamined, all 470 Forever talents):

| Concept | Blizzard articles | Forever talent text |
|---|---|---|
| attack table | 0 | — |
| glancing blow | 0 | **0** |
| crushing blow | 0 | **0** |
| dual-wield miss penalty | 0 | 0 (but see the Dual Wield Specialization inference, §1.2) |
| level-difference miss table | 0 | 0 |

What can be established indirectly:

- **The outcome vocabulary is intact.** Hunter **Counterattack** (104989): "Counterattack **cannot be
  blocked, dodged, or parried**." Rogue **Riposte** (105735) and Counterattack both trigger "after
  parrying an opponent's attack". Warrior **Master of Defense** (105971) and Shaman **Improved
  Stormstrike** (104742) trigger "when you **Dodge or Parry**". Paladin **Reckoning** (105627) triggers
  "after **Blocking** a melee attack". So miss / dodge / parry / block / crit all still exist as discrete
  outcomes. (datamined)
- **Single roll vs two roll: unknown.** No evidence either way.
- **Block still subtracts a value rather than a percentage** — implied by Shield Slam being "increased by
  your Block Value" and by the Era payload's `blockamount` key, not proven.
- **Crit suppression / the +3-level boss table: unknown**, and it is entangled with weapon skill (§1.2).

**This is the single biggest hole in the document and it is unavoidable.** Vanilla's attack table is
derived from weapon skill vs defense skill; Forever keeps both, says weapon skill "works as it always
has", and then adds a second dodge/parry lever. The composition is untestable until Sept 17. §11.2,
§11.3, §11.4.

---

## 6. Spell mechanics

**Spell coefficients.** No statement, and no data source will supply them. `SpellScaling` does not exist
in the Classic-lineage DB2 schema (`404`, verified in [07-simulator.md](07-simulator.md) §5.2) and
`SpellEffect.EffectBonusCoefficient` is routinely 0 or wrong for Classic-lineage spells. Coefficients in
Forever will be a convention (cast_time/3.5 direct, duration/15 periodic, halved for hybrids) plus
per-spell overrides, exactly as in vanilla, and **Forever's new spells have coefficients nobody knows**.
Drain Hope (105909), Holy Strike, Seal of Fury, Light's Vigil, Mutilate, Venom, Cutthroat, Restless
Blades, Berserk, Genesis and Twist of Light all need one. §11.6.

**Partial resists.** No statement. Indirect evidence that the *caster-side* resist lever changed: four
vanilla talents that reduced target resist chance (Arcane Focus, Elemental Precision, Shadow Focus,
Suppression) now grant **hit** instead (§1.2). Whether the partial-resist damage-reduction system itself
survives is unknown. **Spell penetration survives as an item stat** (Rune of Perfection, §4), which
implies resistance still mitigates, which implies partial resists still exist. Inference.

**Spell crit multiplier.** No direct statement, but the talent arithmetic is unambiguous and consistent
with vanilla's 1.5× baseline expressed as a "+50% bonus":

- Mage **Ice Shards** (105777): +20% crit damage bonus per rank, 5 ranks → **+100%**
- Shaman **Elemental Fury** (104766): +20%/rank, 5 ranks → **+100%**
- Warlock **Ruin** (105883): +20%/rank → **+100%**
- Druid **Vengeance** (104932): +20%/rank → **+100%**
- Mage **Arcane Mind** (105800): "+2% Intellect and +20% crit damage bonus of Arcane spells" → **+100%**
- Priest **Shadowform** (105817): "increasing the critical strike damage bonus of your Shadow spells by
  **100%**"

Five separate talents landing on exactly +100% is the vanilla pattern of "double the 50% bonus to reach
a 2.0× crit". So **spell crit is 1.5× baseline and these talents take it to 2.0×** — an inference from
convergent tooltip arithmetic, not a statement. Physical crit follows the same shape: Warrior **Impale**
(105947) +10%/rank → +20%, Hunter **Mortal Shots** (105002) +6%/rank → +30%, Rogue **Lethality**
(105716) +6%/rank → +30%, all identical to their vanilla values, implying a **2.0× physical baseline**
unchanged. (all datamined, verified 2026-09-14)

**"Critical effect chance" vs "critical strike chance".** Forever's text uses "critical **effect**
chance" for healing (Priest Holy Specialization, Shaman Tidal Mastery, Druid Improved Regrowth, Paladin
Divine Favor and Illumination) and for Warlock Malevolence, and "critical **strike** chance" elsewhere.
Priest **Inner Focus** (105844) keeps the vanilla qualifier: "increases its critical effect chance by
25% **if it is capable of a critical effect**" — which is direct evidence that **not everything can
crit** in Forever. Which things cannot is exactly the §1.1 question.

**Spell haste at 60.** Exists but is thin. No talent grants flat passive spell haste. The only passive
source found is the Skyborne racial **Wind Blessed**, "1% increased melee, ranged, and spellcasting
Haste". Actives exist: Troll **Berserking** (+10% spellcasting and attack speed), Shaman **Rage of the
Farseer** (+30% melee attack and spell casting speed), Druid **Nature's Grace** (+10% spellcasting speed
and −10% GCD for 3 sec). Warlock **Create Spellstone** grants "spell haste by 1%". No item has been shown
with haste on it. (datamined / community)

Note that **haste is still described as three separate things** ("melee, ranged, and spellcasting Haste")
in the one racial that grants all three — so Forever unified hit and crit but appears **not** to have
unified haste. Worth confirming.

**GCD reduction is a separate lever from haste.** Druid Gift of the Earthmother (104918) "Reduces the
global cooldown by 0.5 seconds on your Rejuvenation, Swiftmend, and Wild Growth spells"; Warrior Improved
Slam (110858) "Reduces the global cooldown and cast time of your Slam ability by 0.25 sec. In addition,
**Slam no longer interrupts your melee swing time**." The engine needs per-spell GCD overrides and a
swing-timer-preservation flag.

**Downranking.** No statement from anyone. Rank structures clearly still exist (the Forever spellbooks
carry Rank 1..N for every spell). Whether downranking remains efficient depends on unpublished
coefficients and mana costs. §11.6.

---

## 7. Stat conversions and dependencies

Every conversion below comes from Forever talent text (datamined via
`nether.wowhead.com/forever/data/talents-classic?dv=100`, verified 2026-09-14, corroborated by
talentsforever). These are **talent-granted** conversions. The **baseline** conversions (agility→crit,
agility→dodge, agility→armor, strength→AP per class, stamina→health, intellect→mana, intellect→spell
crit, spirit→regen) are **entirely unpublished** — see §3 and §11.5.

| Conversion | Talent (spell id) | Text | Max |
|---|---|---|---|
| Intellect → spell damage **and** healing | Paladin **Champion of the Light** (110882) | "Increases your spell damage and healing by up to 33% of your Intellect." | 100% at 3 ranks |
| Intellect → spell damage and healing | Shaman **Mental Quickness** (104744) | "…by up to 15% of your Intellect." | 30% at 2 ranks |
| Intellect → attack power | Hunter **Careful Aim** (105008) | "Increases your Attack Power by 20% of your Intellect." | 100% at 5 ranks |
| Intellect → attack power | Shaman **Mental Dexterity** (104755) | "Increases your Attack Power by an amount equal to 33% of your Intellect." | 100% at 3 ranks |
| Intellect → armor | Mage **Arcane Resilience** (105809) | "Increases your Armor by an amount equal to 25% of your Intellect." | 50% at 2 ranks |
| Spirit → spell healing **and** spell damage | Priest **Spiritual Guidance** (105853) | "Increases your spell healing by up to **5%** of your total Spirit and your spell damage by up to **1%** of your total Spirit." → rank 5: "**25%** … and **8%**" | 25% / 8% |
| Level → melee attack power | Druid **Predatory Strikes** (104952) | "Increases your melee Attack Power by 50% of your level." | 150% at 3 ranks |
| Level → spell damage and healing | Warlock **Demonic Knowledge** (105893) | "…by up to 33% of your level while you have a summoned Demon pet active." | 100% at 3 ranks |
| Defense skill → armor (forms) | Druid **Thick Hide** (104942) | "…you gain 1 additional base Armor per level and another 0.67 base Armor for each point of defense skill beyond five times your level." | 3 / 2.0 at 3 ranks |
| Attack power → damage | Warrior **Bloodthirst** (105930) | "damage equal to **35% of your Attack Power** plus 30" | — |

**Spiritual Guidance independently confirms the one-third rule.** Blizzard says bonus healing "now also
includes one third as much bonus damage". Spiritual Guidance at rank 5 gives **25% of Spirit to healing
and 8% to spell damage** — and 25 ÷ 3 = 8.33, which rounds to 8. At rank 1 it is 5% and 1% (5 ÷ 3 = 1.67
→ 1). **The one-third relationship holds inside talents as well as on items**, which strongly suggests it
is a global rule applied at the point where a healing bonus is granted, not an itemisation convention.
This matters for the engine: implement it as a derived stat (`SpellDamage += HealingPower / 3`), not as
an item-parse special case.

---

## 8. Diminishing returns

**No source mentions diminishing returns on anything** — not on dodge, parry, block, resilience (which
does not exist), avoidance, or crowd control. Zero occurrences across Blizzard's articles and across all
470 Forever talents (blizzard / datamined, verified 2026-09-14). Vanilla had none on avoidance. Assume
none; verify at beta (§11.7).

---

## 9. Item budget and itemization

Blizzard's own summary, verbatim:

> "These itemization updates reach across the entire game. Every dungeon drop has been re-examined,
> memorable items remain memorable, less exciting items have been adjusted, and hundreds of new drops
> have been added. **Unique dungeon boss items are now blue quality and set bonuses have been improved to
> feel more useful and exciting.**"
>
> "Quest rewards have also been updated or expanded to support more classes and roles. **Mor'ladim's quest
> in Duskwood, for example, still offers a two-handed sword, but now also grants the Ladimore Heirloom
> Ring for spellcasters. Hemet Nesingwary's quests in Stranglethorn also include new options such as the
> Master Hunter's Spellsword.** World-drop epics have been improved, hundreds of new rare-creature drops
> have been added…"

— Blizzard Deep Dive recap (blizzard)

**Correction for the site:** Blizzard's word is "**blue quality**", not "rare quality".
[01-official-facts.md](01-official-facts.md) §13 and
[06-since-announcement.md](06-since-announcement.md) §5 both say "Rare quality" and should be reworded to
quote Blizzard.

**No Blizzard article gives a single number for a single item.** Every number below comes from a
photographed panel slide (community, Tier B).

### Items with concrete numbers

| Item | Level | Stats as shown | Source |
|---|---|---|---|
| **Lionheart Helm** | 56, plate | +18 Strength, **2.0% hit**, **2.0% crit** | Output Lag (community, Tier B) |
| **Edgemaster's Handguards** | 44, mail | **1.0% reduced chance to be dodged or parried**, +1 Weapon Damage, **+1 to Axes, Daggers, Swords** (vanilla: +7 each) | Output Lag |
| **Hide of the Wild** | 57, cloak | +10 Intellect, +8 Stamina, **healing up to 42 and damage up to 14** for all spells | Output Lag |
| **Staff of Westfall** | — | up to **48 healing and 16 spell damage** and **2% movement speed**, but only **in Elwynn Forest, Westfall, Redridge Mountains, and the Deadmines** | Output Lag |
| **Rune of Perfection** | 20, trinket | spells gain **2 magical resistance piercing, plus 3 more in forest and woodland areas** | Output Lag |
| **Worgenbane Talisman** | 21, trinket | **stuns a target Worgen for 4 seconds, 5 minute cooldown** | Output Lag |
| **X'caliboar** | 37, two-hand sword | +20 Strength, +7 Stamina, **chance on hit to deal 33 to 53 Holy damage, tripled against Undead and Swine** | Output Lag |
| **Graverobber's Shovel** | req 23, two-hand mace | 68–102 damage, 3.70 speed, +10 Stamina, **+12 Attack Power, +12 Attack Power vs. Undead**; Use: digs a grave at Raven Hill Cemetery and Scarlet Monastery Graveyard, 2-min cooldown, 20 charges | Warcraft Tavern / Output Lag (community), **What's Next slide, SINGLE-SOURCE per outlet** |
| **Tidal Charm** | — | extended effect against specific enemy archetypes such as the Naga | Output Lag, SINGLE-SOURCE |
| **Whitemane's Chapeau** | — | "class-themed for healers" | Output Lag, SINGLE-SOURCE |
| **Sharpened Cutlery** | — | named only, no effect reported | Output Lag, SINGLE-SOURCE |

### What these numbers actually establish

1. **Two independent confirmations of the one-third rule.** Hide of the Wild: 42 healing → **14** damage
   (42 ÷ 3 = 14 exactly). Staff of Westfall: 48 healing → **16** damage (48 ÷ 3 = 16 exactly). Combined
   with Spiritual Guidance (§7), the rule is confirmed three ways and is exact, not approximate.
2. **Weapon skill per item fell 7× .** Edgemaster's went from +7/+7/+7 to +1/+1/+1.
3. **Conditional effects come in at least four shapes**, and each needs different engine support:
   - **zone list** (Staff of Westfall: four named zones including a dungeon)
   - **biome class** (Rune of Perfection: "forest and woodland areas"; the slide also named Mountains,
     Deserts, Cities and Caverns)
   - **creature type** (X'caliboar: Undead and Swine; Graverobber's Shovel: +12 AP vs Undead; Worgenbane:
     Worgen only; Tidal Charm: Naga)
   - **conditional magnitude, not on/off** (Rune of Perfection gives 2 base **plus 3 more** in forest —
     an additive bonus, not a gate)
4. **A creature-type multiplier can be 3×** (X'caliboar's proc against Undead and Swine), which is far
   outside the range of vanilla's creature-type items and is a real sim input for encounter-specific
   gear sets.
5. **"Swine" is a creature type nobody has seen.** Zierhut on the panel: tripled against Undead and Swine,
   "whatever that is". The engine's creature-type enum needs to be extensible.

**Item budget as such: no statement.** Blizzard published nothing about stat budget, item level or the
random-property point table, and Wowhead's Forever `randPropPoints` is all zeros (site, verified
2026-09-14).

---

## 10. What Blizzard has *not* said

This section exists so nobody re-runs these searches. Verified by keyword grep across all eight Blizzard
Forever articles plus the WoW: Forever Blizzard forums (blizzard, 2026-09-14):

- **Zero Blizzard mentions of:** haste, expertise (the word), resilience, defense as a stat, armor
  penetration, mana per 5, block value, resistances, spell power, ratings vs percentages, attack table,
  glancing blows, crushing blows, dual-wield miss, level-difference miss, spell coefficients, partial
  resists, spell crit multiplier, downranking, periodic/tick/damage-over-time, diminishing returns, item
  budget or item level.
- **No developer blue post on stats exists.** There are exactly four Blizzard staff topics in the Forever
  forums, all by Kaivax, and the only substantive one corrects Krol'dok Stronghold's level range to
  40–45. Player threads asking about weapon-skill tooltips, Intellect→spell-damage conversion and
  itemization sit **unanswered**.
- **Blizzard's total published surface on combat math is five sentences**: hit/crit combination, weapon
  skill per-item reduction, bonus healing → ⅓ bonus damage, the dodge/parry-reduction stat, and caster
  weapons granting spell damage and healing.

---

## 11. Settled only by the beta

Sept 17. Each entry is the exact test to run on day one.

**11.1 — Can DoTs crit, and which ones?**
Roll a Warlock to a level where Corruption is available. Log 500+ Corruption ticks with the combat log
and count `SPELL_PERIODIC_DAMAGE` events carrying the critical flag. Repeat for Bane of Agony, Immolate
(Destruction, to test a non-Affliction DoT), Rend (physical DoT), Serpent Sting (ranged DoT), Moonfire
(Balance), Rejuvenation (periodic *heal*), and Renew. Record the observed crit rate against the character
sheet crit to determine whether periodic crit uses the same crit chance. Then take 3/3 Pandemic and
measure the crit damage multiplier on Corruption ticks specifically — that gives the periodic crit
multiplier and validates whether 1.5× is the baseline. **This single test settles §1.1 completely.**

**11.2 — How do unified hit and weapon skill compose?**
Park a character in front of a +3-level target with a weapon type they have 0 bonus skill in. Record
2,000 white swings: miss / dodge / parry / glance / crit / hit counts. Equip a +1 weapon-skill item of
that type and repeat. Then swap to +1% hit instead and repeat. Three miss rates from the same baseline
tell you whether weapon skill and hit feed the same number, whether weapon skill still moves dodge, and
whether it still moves the glance rate and glance damage. Log glance damage distributions in all three
runs.

**11.3 — Does expertise stack with weapon skill on dodge?**
Same rig, target *behind* (no parry). Baseline dodge rate; then +1 weapon skill; then +1.0% expertise;
then both. If both-together ≈ sum of the singles, they stack additively and the engine subtracts both.
If both-together ≈ the larger single, one supersedes the other.

**11.4 — Is the attack table still a single roll, and does it still include glancing and crushing?**
From 11.2's 2,000-swing logs: check whether the outcome frequencies sum to 1.0 under a single-roll
ordering (miss → dodge → parry → glance → block → crit → hit). Look for any crushing blows against a
+3 target. Separately, confirm the dual-wield penalty by comparing white-swing miss rate two-handed vs
dual-wielding at identical hit, and confirm specials bypass it.

**11.5 — The base conversion tables.**
Read the character sheet at level 60 for every class: base crit at 0 agility-equivalent, crit per point
of agility, dodge per point of agility, armor per point of agility, AP per point of strength, health per
point of stamina, mana per point of intellect, spell crit per point of intellect, mana regen per point of
spirit. Then re-pull `nether.wowhead.com/forever/data/gear-planner?dv=100` — its `baseStats`,
`critPhysical` and `critSpell` sections should populate within days of the client shipping, and they are
exactly the shape `wowsims`' `wowhead_db.go` already parses.

**11.6 — Spell coefficients for every new and changed ability.**
With 0 spell damage, cast each spell 50 times and record average damage; add a known flat spell-damage
value and repeat; the slope is the coefficient. Priority list: Drain Hope, Holy Strike, Seal of Fury,
Light's Vigil, Consecration (its first-four-targets split), Mutilate, Venom, Cutthroat, Restless Blades,
Genesis, Berserk, Twist of Light, and every Warlock Bane. Also measure downranking efficiency for the
big healers once coefficients are known.

**11.7 — Diminishing returns.**
Stack dodge from 5% to 30% in steps and check linearity. Same for parry and block.

**11.8 — Is haste unified?**
Take the Skyborne Wind Blessed racial and check whether the character sheet shows one haste line or
three. Measure swing timer, ranged shot timer and cast time change from the same 1%.

**11.9 — Hit caps.**
Stack hit against a +3-level boss dummy and find the value at which spell misses stop, and separately the
value at which white-swing misses stop. Compare to vanilla's 16% / 8%(+ DW 19%).

**11.10 — The item stat vocabulary.**
Pull `nether.wowhead.com/forever/data/gear-planner?dv=100` and diff the item stat keys against today's
Era set (`mlehitpct`, `rgdhitpct`, `splhitpct`, `mlecritstrkpct`, `rgdcritstrkpct`, `splcritstrkpct`).
A unified `hitpct` / `critstrkpct` key and a new expertise key appearing is the confirmation that the
data environment has flipped to Forever content, and it is the trigger to run the whole gear pipeline.

**11.11 — Conditional item effects.**
Establish where the biome/zone and creature-type conditions live. Check whether `ItemEffect` gains a
`PlayerConditionID` that resolves to a zone list, and whether creature-type conditions are readable or
script-side. This decides whether our item database can carry conditions as data or needs a hand-written
table.

---

## 12. What this changes in the engine

Targets are `wowsims/classic` paths, which is the fork base recommended in
[07-simulator.md](07-simulator.md) §6.1. Line references verified against a local checkout today.

### 12.1 `sim/core/stats/stats.go` — the Stat enum

The enum already carries `Expertise`, `Resilience`, `ArmorPenetration`, `SpellPenetration`, `Defense`,
`Block`, `BlockValue`, `Dodge`, `Parry`, `HealingPower`, `SpellDamage`, `MP5`, all five resistances, and a
separate `WeaponSkills [WeaponSkillLen]float64` array keyed by weapon subclass. Changes:

1. **Merge `SpellHit` + `MeleeHit` → one `Hit`; merge `SpellCrit` + `MeleeCrit` → one `Crit`.** This is
   the enum renumbering already flagged in the Phase 2 notes. It touches every spec's
   `ApplyTalents`, every item parse, every stat weight, and the protobuf. Keep the old names as
   deprecated aliases only for the duration of the port, then delete — no backwards-compatibility
   shims in the shipped enum.
2. **Delete `Resilience`.** Zierhut: "we're not adding resilience." Removing it is free and prevents the
   UI from offering a stat weight for a stat that cannot exist.
3. **Keep `WeaponSkills` exactly as is.** Weapon skill survives on items at ~1/7 the magnitude; the array
   and `GetWeaponSkill()` need no structural change, only new item data.
4. **Add a percentage armor-ignore concept.** `ArmorPenetration` in the enum is a flat value; Forever's
   three talents ignore a *percentage* of target armor. Either add `ArmorPenetrationPercent` or handle it
   as a `PseudoStats` multiplier applied before mitigation. Prefer `PseudoStats` — it is per-attacker and
   conditional on weapon subclass, which a flat stat cannot express.
5. **Do not add a Haste merge.** Forever appears to keep melee/ranged/spell haste separate (§6); leave
   `MeleeHaste` and `SpellHaste` alone until 11.8 says otherwise.

### 12.2 `sim/core/stats/deps.go` — the dependency graph

`safeDepsOrder` is the topological order for stat dependencies. Changes:

1. **Add `HealingPower → SpellDamage` at a ratio of 1/3**, globally, as a stat dependency. This is the
   cleanest expression of Blizzard's rule and it is confirmed three ways (§7, §9). **Correction
   (2026-09-15, from reading `sim/core/stats/deps.go` in the fork):** `safeDepsOrder` lists
   `SpellDamage` *before* `HealingPower`, the reverse of what this section first claimed, so the
   dependency needs `HealingPower` moved ahead of `SpellDamage` in that order as well as the
   registration. The engine plan's Task 6 carries the reorder and a check that the order is safe.
2. **Collapse the `SpellCrit` and `MeleeCrit` entries into one `Crit`** to match 12.1.
3. **New per-class talent dependencies** from §7: Intellect→SpellDamage/HealingPower (Paladin Champion of
   the Light, Shaman Mental Quickness), Intellect→AttackPower (Hunter Careful Aim, Shaman Mental
   Dexterity), Intellect→Armor (Mage Arcane Resilience), Spirit→HealingPower and Spirit→SpellDamage at
   *different* ratios (Priest Spiritual Guidance, 25% and 8%), Defense→Armor (Druid Thick Hide). Note
   Spiritual Guidance is the one place where the ⅓ rule is baked into a talent rather than derived — make
   sure it does not double-apply on top of the global dependency from item 1.
4. **The baseline conversions are unknown** (§3, §11.5). Leave the Era values in place, tag them, and
   regenerate from the beta client.

### 12.3 `sim/core/spell_outcome.go` and `sim/core/spell_result.go` — the attack table and crit

The attack table lives in `outcomeMeleeWhite` (single roll, ordering miss → dodge → parry → glance →
block → crit → hit) and the `applyAttackTable*` helpers at lines ~650–780. Assessment:

1. **The ordering needs no change until 11.4 says so.** Keep it.
2. **Expertise already works the way Forever needs.** `applyAttackTableDodge` and `applyAttackTableParry`
   already read `attackTable.Attacker.stats[stats.Expertise] / 100` and subtract it from both — the SoD
   lineage already made expertise a percentage. This is a **zero-change** subsystem, which is a genuine
   windfall. Only `ExpertisePerQuarterPercentReduction = 2.500000` in `base_stats_auto_gen.go` is stale
   and should go.
3. **The dual-wield penalty is hardcoded `missChance += 0.19` in `applyAttackTableMiss`**, with
   `applyAttackTableMissNoDWPenalty` for specials. Forever's Dual Wield Specialization grants off-hand
   hit specifically (§1.2), so the penalty structure is right but **the 0.19 constant must become a tuned
   config value** measured by test 11.4, not a literal.
4. **DoT crit is the real work.** The machinery already exists — `Dot.OutcomeTick` (never crits),
   `Dot.OutcomeTickPhysicalCrit`, `Dot.OutcomeSnapshotCrit`, `Dot.OutcomeMagicHitAndSnapshotCrit` — so
   enabling periodic crit is a matter of **assigning the right outcome function per spell**, plus one new
   variant for a magic DoT that rolls crit **per tick** rather than snapshotting. Build that variant now.
   Then: default every Forever DoT to `OutcomeTick`, and drive the exceptions from a per-spell config
   flag that test 11.1 populates. Do not branch on school.
5. **`CritMultiplier` needs a periodic-specific path.** §6 establishes 1.5× spell / 2.0× physical as the
   probable baselines with talents adding to the *bonus*; the periodic multiplier is unknown and may
   differ. Give `Dot` its own `CritMultiplier` rather than borrowing the parent spell's.
6. **Add a weapon-subclass-conditional modifier hook** for Weaponmaster / Hack and Slash (§4). These
   switch effect on equipped weapon type within a single talent, which the current one-talent-one-effect
   `ApplyTalents` shape does not express.

### 12.4 `sim/core/base_stats_auto_gen.go` — rating conversions

**Largely already correct, which is the headline.** The file states:

```go
// Crit/Hit/Haste ratings are straight percentage values in classic
const CritRatingPerCritChance = 1
const MeleeHitRatingPerHitChance = 1
const SpellHitRatingPerHitChance = 1
```

Forever uses flat percentages (§2), so these 1:1 constants are right. Changes:

1. **Delete `ResilienceRatingPerCritReductionChance = 28.750002`** — resilience does not exist.
2. **Set `ExpertiseRatingPerExpertiseChance` to 1 and delete
   `ExpertisePerQuarterPercentReduction = 2.500000`** — the item unit is a percentage (Edgemaster's 1.0%),
   not expertise points, so the quarter-percent quantisation from later expansions is wrong.
3. **Everything under `// TODO: Update Defense/Dodge/Parry rates` still needs the beta** (11.5).
4. **Regenerate the base-stat table for 10 races** (nine plus Skyborne's two faction variants) and
   whatever class list Forever ships. Today's Wowhead `baseStats.raceOffsets` has 8 races and is Era data
   (site, verified 2026-09-14).

### 12.5 `sim/core/target.go` — the weapon-skill attack table

`NewAttackTable` (lines ~300–385) derives `BaseMissChance`, `BaseParryChance`, `BaseDodgeChance`,
`BaseGlanceChance`, `GlanceMultiplierMin/Max`, `HitSuppression` and `MeleeCritSuppression` from
`weaponSkill` vs `targetDefense`, citing the magey/classic-warrior wiki. **This is the subsystem most at
risk.** Forever keeps weapon skill and says it "works as it always has", but also adds a second
dodge-reduction lever and cut per-item weapon skill 7×. Do not touch the formulas yet; instead:

1. **Extract all nine magic constants into a named, versioned config struct** so tests 11.2–11.4 can
   fit them without a code change.
2. **Add a regression fixture** that pins today's Era-derived numbers, so any Forever divergence shows up
   as a failing test rather than silent drift.

### 12.6 Per-spell code and the item database

1. **Every spec's hit and crit talents need rewriting**, not just renaming — thirteen talents changed
   meaning (§1.2), four of them converting from resist-reduction to hit. Grep each spec for `SpellHit` /
   `MeleeHit` and rebuild against the table in §1.2.
2. **Conditional item effects need a new data shape** — zone list, biome class, creature type, and
   additive-bonus-not-gate (§9). This is not expressible in the current `item_effects.go` pattern and
   should be designed before beta so the item pipeline has somewhere to put the data.
3. **The creature-type enum must be extensible** — "Swine" is not a vanilla creature type.
4. **Talent modelling needs a weapon-subclass switch** (12.3 item 6) and per-spell GCD overrides plus a
   swing-timer-preservation flag (§6, Improved Slam).
5. **Trigger the gear pipeline on test 11.10**, not on a date — the Wowhead environment is still Era-shaped
   and the unified stat keys appearing is the real signal.

### 12.7 Ordered work list

| # | Task | Blocked on beta? | Touches |
|---|---|---|---|
| 1 | Merge hit and crit in the Stat enum; delete Resilience | No | `stats.go`, proto, every spec |
| 2 | `HealingPower → SpellDamage` at ⅓ as a stat dependency | No | `deps.go` |
| 3 | Fix the expertise constants; delete the quarter-percent quantisation | No | `base_stats_auto_gen.go` |
| 4 | Add a per-tick magic-DoT crit outcome variant + `Dot.CritMultiplier` | No | `spell_outcome.go` |
| 5 | Extract the weapon-skill attack-table constants into config + pin a regression fixture | No | `target.go` |
| 6 | Add `PseudoStats` percentage armor ignore + weapon-subclass conditional modifiers | No | `spell_result.go`, `character.go` |
| 7 | Design the conditional-item-effect data shape (zone / biome / creature type) | No | `item_effects.go`, item pipeline |
| 8 | Rewrite the 13 changed hit/crit talents | No | per-spec `talents.go` |
| 9 | Turn on DoT crit per spell | **Yes** (11.1) | per-spec spell code |
| 10 | Fit miss / dodge / parry / glance constants; set the DW penalty | **Yes** (11.2–11.4) | `target.go` |
| 11 | Regenerate base stats and conversion tables | **Yes** (11.5, 11.10) | `base_stats_auto_gen.go` |
| 12 | Coefficients for every new ability | **Yes** (11.6) | per-spec spell code |

Items 1–8 are ~60% of the stat work and none of them wait for Sept 17.

---

## 13. Corrections to earlier research

1. **`06-since-announcement.md` §4, Warlock:** "Malevolence (DoTs can crit)" is **wrong**. Malevolence is
   "Increases the critical effect chance of your Shadow spells by 1%." The DoT-crit evidence is
   **Pandemic**, and it is an inference, not a statement. (§1.1)
2. **`01-official-facts.md` §13 and `06-since-announcement.md` §5:** Blizzard's word for the new dungeon
   boss item quality is "**blue quality**", not "rare quality". Quote Blizzard. (§9)
3. **`07-simulator.md` §5.2** says the Wowhead Forever environment serves only Era-shaped data. That is
   still true of the **item** endpoint, but `nether.wowhead.com/forever/data/talents-classic?dv=100`
   serves **real Forever talent data** with new spell IDs, and it is the best pre-beta source that
   exists. Add it to the data ledger. (§0)
4. **`07-simulator.md` §5.3** lists the unified-hit/weapon-skill interaction as "completely unspecified".
   Still true, but it can now be narrowed: weapon skill is gone from all talents and racials and survives
   only on items at ~1/7 the per-item magnitude, and expertise demonstrably coexists with it on the same
   item. (§1.2, §1.3)
