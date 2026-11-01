# Sim audit round 2: hybrids (2026-11-01)

This audit covers Druid (Balance, Feral cat, Feral bear, Restoration), Shaman (Elemental, Enhancement, Restoration) and Paladin (Retribution, Protection, Holy).

**Setup:**
- Client 1.60.1.70707.
- Engine: pin 17d31ecfa, plus fork branch `lane/r2-hybrids` (head 0c8df7db2).
- Site: branch `wt/r2-hybrids`.
- Round 1 is `design/reviews/2026-10-31-theorycraft-audit/hybrids.md`.

**How claims were checked:**
- Client tables: the raw 70707 wago CSVs (`SpellEffect`, `SpellLevels`, `SpellAuraOptions`, `Spell`, `SpellName`).
- Wowhead Forever tooltips: `nether.wowhead.com/forever/tooltip/spell/<id>`, fetched 2026-10-31.
- Sim runs: `sim/cmd/spec-breakdown` and `sim/cmd/talent-search`, raid preset on the level-60 alliance band, 10000 iterations at seeds 1 and 50001.
- Real logs: our public API (see Real logs below).

**Classes:**
- **CONFIRMED-BUG:** evidence attached, fixed end to end.
- **FOREVER-CHANGE:** we are right.
- **OPEN:** each row says what would settle it.
- **Aligned:** checked and correct.

Fix-sha column: "E" is a fork commit on `lane/r2-hybrids`, "S" is a site commit on `wt/r2-hybrids`.

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| C1 | **The crit rule's side effects are all as ruled.** Consecration is a non-channelled ground dot and crits on the spell table. Hurricane is channelled and unflagged, so it does not crit. Moonfire, Insect Swarm and Flame Shock ticks crit with the spell multiplier. Rip, Rake, Lacerate and Pounce roll the melee table with the physical multiplier. Holy Shock and Exorcism are direct and unchanged. | **Engine:**<br>- `sim/core/dot.go:398` `periodicDamageCanCrit`.<br>- Consecration is a Dot with `IsAOE` and no channel flag (`sim/paladin/consecration.go:111`).<br>- Hurricane carries `SpellFlagChanneled` (`sim/druid/hurricane.go:34`).<br>- `Dot.outcomeCritTick` picks the melee table for physical dots (`sim/core/spell_outcome.go:46`).<br>**Sims:** Consecration ticks crit 38.4%, the paladin's unified spell crit (Forever merges melee and spell crit into one stat, `sim/paladin/item_sets_pve.go:26`). Rip ticks crit 38.5%, the cat's melee crit. | Aligned | - |
| C2 | **No proc double-triggers off a tick.** | **Hooks:** Omen of Clarity, Vengeance (paladin and druid), Elemental Focus, the Seals and Unsealed Wrath all hook `OnSpellHitDealt` or `OnCastComplete`. None of them hooks `OnPeriodicDamageDealt` (no druid, shaman or paladin file uses it).<br>**Client:** Omen's proc mask is 0x14014 (`SpellAuraOptions` 129438): melee auto, melee special, harmful and helpful spells. It has no periodic bit. Vengeance's text says "non-periodic critical strike". | Aligned | - |
| C3 | **Blessing of Kings and Heart of the Lion: no Forever evidence either way.** Kings stays out of the raid preset. | **No Forever source:** none of the Wowhead Forever tooltips for 409583, 409580, 20217 or 25898 carries exclusivity text, and no Forever patch note or report covers the question.<br>**Client:** both are aura 137 at +10. Kings carries `SpellLabel` 26 and Heart of the Lion none (round 1).<br>**Season of Discovery (Heart of the Lion's origin), the only evidence:** Kings replaced the area aura and left the hunter's self 10% (Blizzard forums, 2023-12-11, https://us.forums.blizzard.com/en/wow/t/blessing-of-kings-doesnt-stack-with-aspect-of-lion-but/1738967).<br>**Stakes:** if they stack, Kings is worth +4.9% to Retribution (measured). | OPEN: a Forever character sheet with both buffs settles it | - |
| C4 | **Heart of the Lion gave 276 attack power at 60; Forever gives 280.** The engine counted levels from `SpellLevel` 1 instead of `BaseLevel` 0. | **Client:** `SpellLevels` row 139819 for 409583 has BaseLevel 0, SpellLevel 1 and MaxLevel 60; base 40 plus 4 a level.<br>**Wowhead Forever:** the buff reads "Attack Power by 280" (`ppl0:60:40:400`).<br>**Fix:** `sim/core/buff_ranks.go:140`. | CONFIRMED-BUG (+4 AP for every melee and ranged character under the raid preset) | E 580225579 |
| C5 | **Judgement of the Crusader used Era's 140 at every rank. Forever's ranks are 23, 35, 58, 92, 127 and 161.** | **Client:** `SpellEffect` aura 14 base points for 21183, 20188, 20300, 20301, 20302 and 20303.<br>**Wowhead Forever:** 20303 reads "up to 161".<br>**Fix:** `JudgementOfTheCrusaderAura` now takes the rank's amount (`sim/core/debuffs.go:470`, `sim/paladin/sotc.go`).<br>No published number moves: the raid preset does not carry the debuff. | CONFIRMED-BUG | E 0c8df7db2 |
| C6 | **The raid preset leaves out Judgement of the Crusader**, while it carries Judgement of Wisdom and Judgement of Light. | **Learnable:** Seal of the Crusader is trainable (`SkillLineAbility` 11547ff).<br>**Measured:** adding the debuff (`-buffs judgement_of_the_crusader`) moves Retribution about +12%. It moves every Holy-damage spec, which includes priests.<br>**Settle:** a preset-policy call. How many paladins does the Phase 1 raid assume, and does one judge Crusader? | OPEN (cross-group if adopted) | - |
| C7 | **Omen of Clarity's spell proc rate is still unconfirmed, but the engine's effective rate matches the only number found.** | **Client:** 16864 has ProcChance 100, a 10 s `ProcCategoryRecovery`, and `SpellProcsPerMinuteID` 0. Wowhead Forever shows "Aura proc chance: 100%" and no PPM.<br>**Only number found:** a Reddit claim (unsourced, 2026-09-30) of "4% chance on spells and 10 sec internal cd" (https://www.reddit.com/r/classicwow/comments/1wu7hcu/), which Moonkin Form doubles to 8% (the Forever deep dive confirms the doubling).<br>**Engine** (`sim/druid/talents.go:331`): 8.8 Clearcasting procs a fight from about 95 Balance casts, about 9% a cast. That is in line with the claim.<br>Melee has no Forever number; it stays at 2 PPM. | OPEN: a caster-druid log with Clearcasting events settles it | - |
| C8 | Cross-check of client "the damage is spell X" references across every trainable druid, shaman and paladin spell. This search found C5 and Shaman E1. | **Found correct:** Hurricane (1278965/1278968/1278759), Consecration (1280345–1280349), Holy Shock (1311604ff) and Seal of Fury (1311647ff) all use the Forever ids. | Aligned | - |

## Druid: Balance

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| B1 | **Raid talent search with `-keep` on the current engine:** the guide build is best (520.4 ± 1.1). | `sim/cmd/talent-search -preset raid -keep "Moonkin Form,Insect Swarm,Eclipse,Nature's Focus,Subtlety"`. | Aligned | - |
| B2 | Moonfire and Insect Swarm ticks crit (31% and 27% in the run). | The engine follows C1.<br>**Forever log:** our public report `z2opt2x2gs5u` (2026-09-17, starting zone) shows Moonfire rank 1 ticks critting 22 of 459 times, at ×1.5. | Aligned | - |
| B3 | `sim/druid/insect_swarm.go:106` zeroes `SnapshotCritChance` outside Moonkin Form. The line is dead: `Dot.OutcomeTick` rolls the live spell crit chance and never reads the snapshot. | Code read. No DPS effect. | Aligned (dead line; cleanup only) | - |
| B4 | Omen of Clarity rate. | See C7. | OPEN | - |

## Druid: Feral (cat)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | **The cat's paw ignored the equipped weapon.** It was pinned at Era's 54.8 dps (43.84–65.76 at 1.0 s, `sim/druid/forms.go:31` before). Forever's form attack carries the weapon's DPS. | **Blizzard druid deep dive (2026-09-30):** "the Druid's melee auto attack DPS in Bear, Cat or Dire Bear Form is now the same as the DPS of the Druid's equipped weapon … speed … 1.0 or 2.5 seconds … Abilities that deal weapon damage likewise use these damage values" (https://us.forums.blizzard.com/en/wow/t/2367361; also Warcraft Tavern's write-up).<br>**Icy Veins (2026-10-26):** 14 AP = 1 DPS on auto attacks.<br>**Shred:** 155% of that hit plus 80 (`SpellEffect` 9830, Wowhead Forever).<br>**Fix:** `formPawWeapon` keeps the weapon's DPS and spread at the form's speed (`sim/druid/forms.go:50`). An unarmed druid keeps Era's paw. | CONFIRMED-BUG | E 4cb26aa16 |
| F2 | **The published band's main hand is Elemental Mage Staff (944).** The BiS ranker never valued weapon DPS for a cat. | `bis/druid-feral.json`, band 60 alliance raid.<br>**With the fix:** the published set moves −3.7%. Grand Marshal's Demolisher (234568, already an alternative) moves +12.8% over today's published number (Before/after below).<br>The nightly BiS regeneration picks the new weapon. Nothing nightly-owned was committed. | Follows from F1 | (nightly) |
| F3 | **The Ravage opener dealt no damage.** It was a prepull action at −1 s, and the engine counts no damage dealt before the pull (`sim/core/spell_result.go:425`). | **Run:** prepull Ravage, avg hit 0. Ravage in the priority list, avg hit 1125.<br>Now first in the priority list; it fires once from Prowl.<br>The other APLs' prepull casts are buffs or hardcasts that land after 0 (Elemental's Lightning Bolt at −2.5 s measured the same at −2 s and −3 s). | CONFIRMED-BUG (+0.5%) | S cc087dbd5, E e0188df98 |
| F4 | Raid talent search with `-keep`: the guide build is best (683.9 ± 0.9). | Talent-search run on the current engine. Rerun after F1: weapon DPS does not change talent values. | Aligned | - |
| F5 | Rip, Rake and Pounce bleeds: melee table, physical multiplier (Rip 38.5% crit). Omen of Clarity procs 7.6 times a fight from melee and Shred. Rend and Tear applies to cats. Tiger's Fury is absent. | Round-1 regressions re-run on the 70707 engine. | Aligned | - |
| F6 | The guide now says why weapon DPS matters. | `web/src/content/guides/druid/feral.md` Stat priority. | - | S cc087dbd5 |

## Druid: Feral (bear)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| G1 | The bear paw also ignored the weapon (54.8 dps at 2.5 s). It now has the weapon's DPS at 2.5 s. | Same deep dive as F1. Fork tank goldens regenerated (+0.4% DPS; the test weapon is close to 54.8 dps). | CONFIRMED-BUG | E 4cb26aa16 |
| G2 | Threat moves −2.2% on today's band weapon. The nightly BiS will re-pick it. | Before/after below. The guide (`feral-bear.md`) says a low-DPS staff costs threat. | Follows from G1 | S cc087dbd5 |
| G3 | Lacerate crits on the melee table. | `sim/druid/tank/lacerate_crit_test.go` passes. | Aligned | - |

## Druid: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Wild Growth is skipped. | **Wowhead Forever (408120, 1238214, 1238215):** 3 s cast, 6 s cooldown, heals the target's party over 7 s.<br>**Measured:** heal-builds put it about 5% less healing for any point it takes. The APL's 70707 note, dated 2026-10-31, records that run. | Aligned under our heal profile (OPEN for a party-damage profile, as in round 1) | - |
| R2 | Bloom (1328621) is modeled and taken. | `sim/druid/bloom.go`. | Aligned | - |
| R3 | Omen of Clarity on heals: 8.97 procs a fight. | C7. | OPEN (rate) | - |

## Shaman: Elemental

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| E1 | **Fire Nova dealt the vanilla totem's damage, about 1.8 times too much.** | **Client:** Forever's Fire Nova (408341–408345) names its damage `$408423s1` … `$408428s1`. Rank 5 (408428) is 239 base points plus 1.3 a level to 57, rolled 0.1098 wide: 232–259 at 60.<br>**Wowhead Forever:** "232 to 259".<br>**Engine:** it dealt 11307's 419 plus 3.4 a level (412–460), and rank 1 used the totem's 0.10 coefficient instead of 0.143.<br>**Fix:** `sim/shaman/fire_nova.go:25`. Shaman conformance golden regenerated. | CONFIRMED-BUG | E 5148265ce, E 1e8d916a6 |
| E2 | **Fire Nova on Clearcasting** (kami-labs, single source). | Measured after E1: +1.0% raid and +2% bare, with less time out of mana. Fire Nova on cooldown loses about 5%.<br>The APL line is gated on `auraIsKnown(16246)` so a build without Elemental Focus never reaches it. | CONFIRMED (rotation gain) | S c46ae9d4f, E fddf8ebe7 |
| E3 | **Improved Reincarnation hold.** The search's best build drops it for +3.2% (477.2 vs 462.4, ±1.3). | **Client:** Forever's Improved Reincarnation adds "+4% maximum health" (talents/shaman.json 104737). That effect is unmodeled, and the adoption rule keeps unmodeled guide talents.<br>**Without the drop:** the clean-variant search says the guide is best. | Held (owner call: is +4% health raid-irrelevant?) | - |
| E4 | Flame Shock ticks crit with Elemental Fury's multiplier, inherited from the parent spell. | C1. | Aligned | - |

## Shaman: Enhancement

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| N1 | **Windfury Weapon against Flametongue Weapon:** Windfury is right. | **Measured:** Flametongue −3.4% (597.8 vs 618.7), with Windfury Totem's procs counted for the Flametongue run.<br>**Wowhead Forever:** 16362 is "20% chance of granting you 2 extra attacks with 433 extra attack power". Flametongue 16339 is "(1728/77 − 1) to (1728/25)" per weapon speed. | Aligned | - |
| N2 | Raid talent search with `-keep`: the best variant that keeps the unmodeled guide talents is +0.45% (Elemental Focus to Reverberation). | Below the 1% adoption floor. | Aligned | - |
| N3 | Lightning Bolt at 5 Maelstrom stacks: 10.2 casts a fight. | Round-1 regression re-run. | Aligned | - |

## Shaman: Restoration

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | Round-1 build (Thundering Strikes 5, Ancestral Knowledge 5) and rotation (Riptide, Chain Heal, Healing Wave) are intact; no OOM anomalies. | Heal-profile run, 612 effective HPS both seeds. | Aligned | - |

## Paladin: Retribution

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **Round-1 X3 is settled: the 70707 client carries the reworked trees.** Unsealed Wrath (1310735) replaced Twist of Light, and Shieldsworn and Oath of Light are present too. | `data/builds/1.60.1.70707/talents/paladin.json`. The engine models Unsealed Wrath (`sim/paladin/unsealed_wrath.go:78`). | FOREVER-CHANGE (followed) | - |
| P2 | **Unsealed Wrath twist.** Store three Seal of Righteousness copies, swing once under Seal of Command, which releases three extra Righteousness hits, then switch straight back. This beats no twist by +1.3% on a legal build. | **Build:** the Reverence point goes to Unsealed Wrath (raidBuild `050050002/503/05233301001301351`).<br>**Result:** 554.2 vs 547.2, both seeds, error ±0.4, no time out of mana.<br>**Variants:** an even three-and-three swap +0.9%; Seal of Command as the only Seal −2.7%; ungated swaps on the bare preset −16%, hence the 30% mana gate.<br>**Rejected point sources:** Benediction 4 measured +0.3% more, but leaves Retribution's tier 1 short of 5 points (`builds.test.ts`).<br>**Gate:** both APL lines open with `auraIsKnown`. An unresolved `auraNumStacks` drops out of an `and`, which made a talentless build twist (−17%) until guarded.<br>Unsealed Wrath joins the raid keeps. | CONFIRMED (build and rotation gain) | S 652253d40, E e0188df98 |
| P3 | Seal of Command as the primary Seal. | −2.7% (see P2). | Aligned (Righteousness stays primary) | - |
| P4 | Judgement of the Crusader amounts and preset. | C5, C6. | Bug fixed; preset OPEN | E 0c8df7db2 |
| P5 | Raid talent search with `-keep` on the pre-twist build: the guide was best. | Talent-search run. | Aligned | - |

## Paladin: Protection

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| T1 | Consecration ticks crit under the ruling (C1). Prot's threat moves +0.1% from Heart of the Lion (C4). | Tank run. | Aligned | E 580225579 |
| T2 | The fork's Protection goldens moved with Judgement of the Crusader's 161 (FullBuffs enables the debuff). | C5. | Follows from C5 | E 0c8df7db2 |

## Paladin: Holy

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| L1 | Holy Shock, Flash of Light, Holy Light and Illumination runs show no anomaly. Holy Shock's heal resolves through 1311605. | Heal-profile run. | Aligned | - |

## Real logs

- Warcraft Logs (fresh, classic, vanilla, forever subdomains) refuse non-browser fetches (403), and no Forever report is indexed by search.
- Our public feed (`GET https://api.foreversixty.gg/v1/reports/recent`, `api/internal/reports/recent.go`) lists three public reports:
  - Two have no fights.
  - The third, `z2opt2x2gs5u`, is from 2026-09-17: 42 starting-zone fights, and every class spell in it is a druid's.
- **What that log settles:** Moonfire periodic ticks crit at ×1.5 (22 of 459 ticks). This supports C1 and the spell crit multiplier.
- **What it cannot settle:** it has no Clearcasting events, cat white hits, Seals, imbues or level-60 DPS. No sim band can be compared against it.

## Before/after (relative %)

**Method:**
- Raid preset, level-60 alliance band, 10000 iterations, seeds 1 and 50001. The two seeds agreed within ±0.1%.
- **Before:** pin 17d31ecfa with origin/main rotations.
- **After:** lane engine 0c8df7db2 with the `wt/r2-hybrids` rotations.
- Healers are scored on effective HPS on the heal profile, tanks on threat per second.

| Spec | Change | Before → after |
|---|---|---|
| Balance | none | 0.0% |
| Feral (published band, caster staff) | F1 weapon DPS, F3 Ravage, C4 | −3.7% |
| Feral (Grand Marshal's Demolisher 234568 in the main hand) | same changes | +9.4% on that weapon; +12.8% over today's published number |
| Feral bear (threat, published band) | G1 | −2.2% |
| Restoration Druid | none | 0.0% |
| Elemental | E1 and E2 | +1.0% |
| Enhancement | C4 | +0.1% |
| Restoration Shaman | none | 0.0% |
| Retribution (published band build) | C4 | +0.1% |
| Retribution (new raidBuild, with the twist) | P2 and C4 | +1.4% (+1.3% from the twist alone) |
| Protection Paladin (threat) | C4 | +0.1% |
| Holy Paladin | none | 0.0% |

The bands and BiS sets are nightly-owned. Feral and bear gear, the Retribution band talents and every weight follow at the next nightly run, with the new engine pin. Nothing nightly-owned was committed.

## Still open

1. **C3, Kings and Heart of the Lion stacking.** No Forever evidence; Season of Discovery says they do not stack. A Forever character sheet with both buffs settles it. If they stack, every spec moves (Retribution +4.9%).
2. **C6, Judgement of the Crusader in the raid preset.** This is a preset-policy call. Adopting it moves Retribution about +12% and every Holy-damage spec, priests included.
3. **C7, Omen of Clarity's spell and melee rates.** A caster or cat log with Clearcasting (16870) events settles them.
4. **E3, the Improved Reincarnation hold** (+3.2% for Elemental). This is an owner call on the adoption rule.
5. **R1, Wild Growth under a party-damage heal profile.** Carried over from round 1.
6. **Unsealed Wrath's unconfirmed reading** (`sim/paladin/unsealed_wrath.go:60`): do special attacks store or consume copies, and does each copy keep its own 15 s? A Retribution log settles both.
7. **Warcraft Logs Forever data** needs a browser session.

## Sources

- Blizzard druid and hunter deep dive (2026-09-30): https://us.forums.blizzard.com/en/wow/t/2367361 ; Warcraft Tavern's write-up: https://www.warcrafttavern.com/forever/news/powershifting-is-back-sort-of-hunter-druid-deep-dive-for-wow-forever/
- Blizzard Beta Dev Notes, druid changes (2026-10-15): https://us.forums.blizzard.com/en/wow/t/druid-changes/2384207
- Wowhead Forever tooltips (fetched 2026-10-31): spells 409583, 16864, 9830, 20303, 21183, 20188, 408345, 11307, 16362, 16339, 20375, 20920, 21084.
- Season of Discovery, Kings and Heart of the Lion (2023-12-11): https://us.forums.blizzard.com/en/wow/t/blessing-of-kings-doesnt-stack-with-aspect-of-lion-but/1738967
- Reddit, Omen of Clarity rate claim (2026-09-30): https://www.reddit.com/r/classicwow/comments/1wu7hcu/
- Icy Veins, Countdown to Forever 8 (2026-10-26): https://www.icy-veins.com/wow-forever/news/countdown-to-forever-8-how-druids-are-shaping-up/
- Client 1.60.1.70707 wago tables: `SpellEffect`, `SpellLevels` (rows 139819, 139677–139681), `SpellAuraOptions` (129438), `Spell`, `SpellName`.
