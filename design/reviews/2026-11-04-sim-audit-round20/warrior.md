# Round 20, lane warrior (Arms, Fury, Protection)

Site `wt/r20-warrior` on int/r20 (25f3e06d9). Engine `lane/r20-warrior` on fork int/r20 (8a915c0fb). Client 1.60.1.70755 with the late hotfix re-layer. Measured against a lane base built from fork 8a915c0fb (`sb-base`), which equals `$R/before.tsv` to the decimal for all twelve warrior values.

**Result: 2 confirmed warrior bugs, both `old`. One is an engine fix (Flurry after a missed Heroic Strike, single source). The other is a card/APL fix (Fury Berserker Rage and Overpower level gates, `guides-handoff`).** The Protection Bloodrage note is a text fix (WARRIOR-R19-2). Nothing outside `sim/warrior`, the warrior APLs and fury.md is touched. **No spec outside the class moves.**

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **A Heroic Strike that misses outright spends a Flurry charge.** Round 19 kept Heroic Strike and Cleave off the Flurry charge, because none of 228 landed Heroic Strikes took one. It left nine charges spent right after a missed Heroic Strike unmodelled. I split those failed Heroic Strikes by outcome in Tawl's build 70450 log (one-hander and shield, so no off-hand swing shares the timestamp). With Flurry up, **all 9 outright misses removed a charge within 1 ms, and none of the 10 dodges and parries did.** Over the whole log: 11 misses (9 under Flurry), 22 dodges and parries, and Unbridled Wrath never procs after any of them. White swings spend a charge whether they land or miss (299 of 380 hits, 70 of 73 misses). Fix: `spendsFlurryCharge` covers any white swing, plus a next-swing ability whose outcome is a miss. Test `TestAMissedHeroicStrikeSpendsAFlurryCharge` failed first (undefined helper) and passes now. **Single source** (one player's log). The DPS effect is within noise (see Before/after): an outright miss on a yellow attack is rare, and Flurry is refreshed so often that one extra spent charge seldom costs a swing. | magey/forever-warrior discussion #30 log (`r19/warrior/logs-dl/tawl_flurry_1h_70450.txt`), `$R/warrior/flurry2.py`; SpellAuraOptions 12966 ProcTypeMask 4 | CONFIRMED (single source, flagged) | old (round 19 saw the 9 and left them) | E 105a9db7a |
| P5 | Revenge, Shield Slam and Thunder Clap threat. The forever-bugs tracker was re-read 2026-11-05 01:55 UTC. Every threat issue (#63, #64, #66, #121, #278, #324) is unchanged since round 19, and nothing new has been filed since 31 October except #445-#448 (not warrior). No blue post and no threat capture exist. | `gh search issues --repo ClassicWoWCommunity/forever-bugs threat`, issue list updated since 2026-10-31 | OPEN (carried) | - | - |
| AC | Arcanite Champion's proc rate: still 1 PPM in `common/item_effects.go:227` (shared). There is no client rate row and no Forever source; it is a level-60 item and no level-60 log exists. | round 19 sources | OPEN (carried) | - | - |
| C1 | Eureka! with Execute (#439). There is a new comment, oNodrak 2026-10-30: "The interaction went from applying the 10% damage bonus to applying neither effect" (screenshot only). It is still one reporter, with no Blizzard answer, on a race no warrior guide picks. The engine applies both (`core/eureka_auto_gen.go`, shared). Not changed. | https://github.com/ClassicWoWCommunity/forever-bugs/issues/439 | OPEN (shared, flagged for platform) | - | - |
| C2 | Deep Wounds' attacker multiplier is applied twice. `procDeepWounds` multiplies the new bleed by the auto attack's `AttackerDamageMultiplier` (Death Wish, Two-Handed Specialization, Enrage, Dual Wield Specialization on the off hand). Each tick then applies the Deep Wounds spell's own multiplier again (`deep_wounds.go` OnTick). No beta log carries a Deep Wounds event (0 in all six community logs), and #234 says only "scales with your weapon damage". Estimated size if it is wrong: one extra multiplier on 5-7% of damage, about 0.3-0.5% DPS. A level-40+ log with Deep Wounds ticks next to a known weapon would settle it. | code; logs grep | OPEN (carried) | - | - |
| DW | Fury Death Wish timing (platform S dd169626d): the APL (`remainingTime >= 210s` or execute phase or `remainingTime <= 30s`), the card note ("At the pull when a second one still fits, otherwise in the execute phase, or in the last 30 seconds when that phase is shorter than the buff") and fury.md:92 all say the same thing. Not redone. | `data/curated/apl/warrior-fury.json`, fury.md:92 | Aligned | - | - |
| R19 | Fresh re-audit of round 19's three engine fixes. (1) `countsAsWhiteSwing`: Unbridled Wrath needs a landed white swing, Reckless Swings a landed white swing, Flurry any white swing (and now the S1 miss case). `white_swing_procs_test.go` passes. (2) Heroic Strike threat per rank: `TestHeroicStrikeThreatIsPerRank` passes, ranks 1-9 = 20/39/59/78/98/118/137/145/175. (3) Deep Wounds: rollover, the running tick timer and no attack power all pass their tests (`deep_wounds_formula_test.go`, `dps_warrior/deep_wounds_refresh_test.go`), and ticks crit by the bleed rule without re-proccing (`deep_wounds_crit_test.go`). The off-hand half matches the server's off-hand damage field, which carries the 50% penalty, with Dual Wield Specialization applied through the auto's multiplier (one emulator lineage). The fork goldens `TestP1DPSWarrior` and `TestP1TankWarrior` are green at the lane base (`-tags with_db`). No regression. | tests in `sim/warrior` | Aligned | - | - |

## Guide-accuracy round 19 leads (sim/APL items)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| G1 | **Fury Berserker Rage line dead on the 30-35 bands (WARRIOR-R19-1).** The action had no gate. Below 36 a Fury warrior fights in Battle Stance (`stances.go` defaultStance: Berserker Stance only once Whirlwind, 36, is known), and the engine refuses a Berserker Stance button there. So the row rendered on the addon's 30 and 32 bands, told the player to press it, and never fired (0 casts at 30 and 34). The action now carries `currentLevel >= 36`, and the note takes the report's wording: "From 36, under 80 rage: Improved Berserker Rage makes it 10 rage, back every 30 seconds. Hold it for the fear on Magmadar, Onyxia and Nefarian." **Why currentLevel and not `spellIsKnown(Whirlwind)`:** the engine registers Whirlwind at every level, because `RequiredLevel: 36` only blocks the cast. So `spellIsKnown(1680)` is true at 30, and my first attempt (`not spellIsKnown(1680)` on Overpower) dropped Overpower at 30 (-8.8% at level 30, caught before commit). `currentLevel` is read the same way by the engine, the guide's level gate (`apl-level-gate.ts`) and the addon generator (`addonrotation.py _level_truth`). | trainables (Whirlwind 1680 at 36, Berserker Rage 18499 at 30); engine casts | CONFIRMED | old, guides-handoff | S c779ea2de, E 8524027de (fork copy) |
| G2 | **Fury Overpower lingered from 36 to 60.** Its note says "Only before level 36"; it is a Battle Stance button and never fires from 36. It now carries `currentLevel < 36`. On the BiS page's band-30 panel (which shows the band's top entry, level 38) it now drops off, as it should. | same | (part of G1) | old, guides-handoff | S c779ea2de |
| G3 | No cast moves with G1/G2. Fury horde, 3,000 iterations, seed 1: bare 30 60.3 / 60.3, 34 69.0 / 69.0, 36 71.4 / 71.4, 40 141.4 / 141.4, and raid 60 904.2 / 904.2 before / after. Berserker Rage casts at 36 and up are unchanged (6.41/iter at 36). Guard test `test_fury_berserker_rage_starts_at_36_and_overpower_ends_there` (red on the old APL, green now). | `$R/warrior/furygate.sh` | - | - | S fc74b5ba4 |
| G4 | **Arms Hamstring has no Execute gate.** Not taken. The step needs a Windfury Totem (`auraIsKnown(10610)`), and the totem's first rank (8512) trains at **32** (`trainables/shaman.json`), after Execute at **24** (5308). Below 24 there is no totem to be on, so the extra `or not spellIsKnown(Execute)` branch could never change a cast. The requested run (bands 10-22 with a Windfury Totem) confirms this: at level 22 with `-buffs windfury_totem`, Hamstring casts 0 times and no Windfury aura appears (`request/preset.go` windfuryTotemFirstLevel 32). The card text stays. | trainables; spec-breakdown | Aligned | - | - |
| G5 | **Protection Bloodrage note at band 10 (WARRIOR-R19-2).** The note credited the second point of Improved Bloodrage at 10. It now reads "... a quarter more with your first point of Improved Bloodrage at 10 and half again from 11", matching protection.md:129. Text only. | talents (Improved Bloodrage 25/50%); Protection order 10, 11 | CONFIRMED (text) | old, guides-handoff | S c779ea2de |

## Arms

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Builds decoded against `talents/warrior.json`. Raid: Improved Heroic Strike 3, Improved Rend 3, Improved Tactical Mastery 3, Improved Overpower 2, Anger Management 1, Deep Wounds 3, Spearing Strike 1, Two-Handed Spec 3, Impale 2, Bloodthrill 2, Sweeping Strikes 1, Weaponmaster 5, Improved Slam 2, Mortal Strike 1 / Cruelty 5, Unbridled Wrath 5, Enrage 5, Reckless Swings 1, Improved Execute 2 / Improved Bloodrage 1. Every point is cast, scales something cast, or is a prerequisite (Sweeping Strikes for Mortal Strike, Improved Rend 3 for Deep Wounds). The leveling build's Deflection 3 and Iron Will 5 stay leveling choices (round 17 A5). | talents | Aligned | - | - |
| A2 | Rotation lines cast and gated as written: Recklessness swap under 25 rage below 20%, Rend kept for Bloodthrill, Overpower first, Mortal Strike, Reckless Swings, Execute, Slam only under 1.5 s, Spearing Strike, Hamstring only under a Windfury Totem (G4), Heroic Strike at 90. Unchanged by S1 (Arms has no Flurry: 0.00% at raid 60). | APL; measure | Aligned | - | - |

## Fury

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | G1, G2, S1 above. | | | | |
| F2 | Build: Cruelty 5, Unbridled Wrath 5, Furious Precision 3, Enrage 2, Dual Wield Spec 5, Reckless Swings 1, Improved Execute 2, Improved Berserker Rage 2, Death Wish 1, Improved Whirlwind 2, Flurry 3, Bloodthirst 1, Improved Bloodrage 2 / Arms 20 incl. tier fillers (Improved Charge 2, Improved Tactical Mastery 3; round 17 F5). Improved Berserker Rage is cast from 36 (G1). | talents | Aligned | - | - |
| F3 | Improved Berserker Rage lands at 34 and 35 in the shipped order, so "From 36 ... Improved Berserker Rage makes it 10 rage" is true wherever the row now shows. | addon-builds order | Aligned | - | - |

## Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Defensive Stance passive 7376 in both 70755 and 70790: -10% damage taken (aura 87), -10% damage done (aura 79), +30% threat (aura 10). The engine has all three (`stances.go` defensiveStance*). Defiance multiplies on top (separate MOD_THREAT auras). | `$R/warrior/client/se-*-7376.csv` | Aligned | - | - |
| P2 | Threat/s unchanged by every commit here (1088.6 / 1079.1 alliance / horde, seed 1). | measure | Aligned | - | - |
| P3 | G5 (Bloodrage note). P5 threat constants stay OPEN. | | | | |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | **Engine `spellIsKnown` is true for a spell below its RequiredLevel** (Whirlwind at 30). Nothing ships on it now (G1 uses currentLevel), but any future `spellIsKnown` gate on a level-gated spell registered at every level will disagree between the engine and the guide's level gate, which reads trainables. Flag for platform: register such spells only from their level, or make `spellIsKnown` honour `RequiredLevel`. | `core/apl_values_spell.go` GetBool (`spell != nil`); `whirlwind.go` RequiredLevel | For platform | - | - |

## Before/after

Raid 60, 10,000 iterations, guide builds via `rt`, frozen BiS (`ALLRT=1 BISDIR=$R/bis-frozen measure.sh`). Lane base `sb-base` (fork 8a915c0fb) = `$R/before.tsv`.

| Spec | seed 1 / 50001, alliance | seed 1 / 50001, horde | change |
|---|---|---|---|
| Arms | 853.0 / 852.0 → 853.0 / 852.0 | 855.5 / 854.4 → 855.5 / 854.4 | 0.00% |
| Fury | 917.6 / 918.3 → 917.6 / 918.3 | 909.8 / 910.5 → 909.8 / 910.5 | 0.00% |
| Protection (threat/s) | 1088.6 / 1087.9 → same | 1079.1 / 1079.6 → same | 0.00% |

Raid 60 is identical to the decimal: at raid hit an outright Heroic Strike miss is rare, and one extra spent charge seldom costs a swing. Bare Fury, horde, 10,000 iterations (base → after): 40: 141.3 / 141.2 → 141.3 / 141.2; 50: 219.7 / 220.1 → 219.8 / 220.0. All within ±0.05%, which is noise. The APL changes move no cast (G3). The fork golden TestP1DPSWarrior moved ±0.3% on its all-items rows (regenerated in E 105a9db7a).

## Files to regenerate

- `data/builds/<active>/addon-data.json` and `addon/ForeverSixty/Data.lua`: the Fury rows (Berserker Rage from 36 with the new note, Overpower only to 35), the Protection Bloodrage note, and the Fury Death Wish note (platform's dd169626d, still pending). A preview splice was built and captured for the review and then reverted; nothing generated is committed.
- No ladder golden or dps-per-point change is expected (no cast moves, and the Flurry change is noise), but the integrator's regen covers them anyway.
- Fork golden regenerated in this lane: `sim/warrior/dps_warrior/TestP1DPSWarrior.results` (E 105a9db7a).

## Measure at refresh

- None new. Carried: warrior BiS after the re-rank (round 19).

## For other lanes

- **Platform:** X1 (`spellIsKnown` below RequiredLevel); C1 Eureka! with Execute (#439, shared racial code, now said to apply neither effect, single reporter); AC Arcanite Champion PPM (shared item code).
- Fork goldens in this package need `-tags with_db` (without it every test panics "No item with id").

## For the guide lane

- Fury Bloodrage card note (the guide-accuracy round 19 Should): the level-proof wording "Whenever rage is under 80: 10 rage at once and 10 more over 10 seconds, back every minute, half again with Improved Bloodrage at 59 and 60" is still open. Not taken here (a Should, and outside the requested items).
- The guide card at 60 now reads "From 36, under 80 rage". The wow-player called the "From 36" noise on a level-60 card. It is a nit, and the addon bands are where it matters.

## Guide text

- Done in this lane (S c779ea2de). fury.md leveling (Levels 20 to 60) now reads "Berserker Stance and Berserker Rage train at 30, but Berserker Rage only works in that stance, so leave it on the bar until 36: from there Whirlwind makes Berserker Stance your home, Berserker Rage joins the rotation, and pulls of two or three go quick." Also the Fury Berserker Rage and Protection Bloodrage card notes. `_prose.test.ts` passes.
- wow-player: **SHIP with one SHOULD** on the rendered Fury card (1440 and 390), the full Fury page and the Protection card. Captures are in `$R/warrior/shots/`; the addon rows were rebuilt for the preview only. The SHOULD said Berserker Rage might train at 32, as in vanilla. Answered from client data: `trainables/warrior.json` and `addon-data.json ability_levels` put 18499 at 30, as do Blizzard's 1 October notes (`berserker_rage.go:9`; guide-accuracy round 19 checked the same and filed nothing). The sentence stays.

## Still open

- P5 threat constants (threat-meter capture or blue post); AC Arcanite Champion's PPM; C1 Eureka! with Execute; C2 Deep Wounds' double multiplier (a 40+ log with Deep Wounds ticks).
- Whether an outright-missed Heroic Strike also spends a Reckless Swings charge. There is no log: Tawl's log has no Reckless Swings, and the engine keeps it landed-white-swings-only.
- WARRIOR-13 (Improved Shield Wall) stays with the owner.
- **Logs:** last looked 2026-11-05 02:00 UTC. foreverlogs.gg now answers "Report not found" for every id probed (1, 2, 10, 2700, 2814, 2815, 2820, 2850, 2900, 3000, 3200), including the beta reports earlier rounds read, so the beta archive looks gone after launch. The listing needs a token. `api.foreversixty.gg/v1/reports/recent` still holds only the 2026-10-20 report. Warcraft Logs is blocked (gcloud); I did not work around it.

## Sources

- forever-bugs: https://github.com/ClassicWoWCommunity/forever-bugs/issues/439 (comment 2026-10-30), threat search and issues updated since 2026-10-31 (read 2026-11-05 01:55 UTC).
- Log: https://github.com/magey/forever-warrior/discussions/30 (Tawl, build 70450, 2026-10-20), local copy from round 19.
- Client: `data/builds/1.60.1.70755/trainables/{warrior,shaman}.json`, `talents/warrior.json`; SpellEffect 7376 (70755 and 70790).
- Guide-accuracy round 19 warrior report (WARRIOR-R19-1, R19-2, Sim findings, Sims requested).
