# Sim audit round 11: auras lane

Scope: consumables, raid buffs, raid debuffs, world buffs (gated off), and the presets (raid, bare and the per-band sets) for every spec. Round 10's auras.md was the starting point; every row below was re-read fresh.

Sources:
- The 1.60.1.70707 client tables with the hotfix push layered in. Every table's sha256 matches `data/builds/1.60.1.70707/hotfix-push.json`. Read with the round-10 lane's `tools/sp.py`, `it.py` and `ranks.py`.
- Wowhead Forever server tooltips (`nether.wowhead.com/forever/tooltip/spell/<id>`), read 2026-11-02. A tooltip and its client row count as one source.
- forever-bugs (github.com/ClassicWoWCommunity/forever-bugs) #20, #64, #109, #215, #270, #278 and #419, read 2026-11-02. AidanZMoon is the responder whose replies round 10 took as Blizzard's (#109).
- warcraft.wiki.gg's Forever pages for the new elixirs (edited 2026-10-27).

Engine branch `lane/r11-auras` (from cfcb76e80). Site branch `wt/r11-auras` (from 8f047a62f).

**Found: 5 CONFIRMED-BUG rows, all fixed, plus one note fix. By origin: old 4, R9 1 (Survival's half of S1). None are R10 or new-data.**
- W1 and W2 move published numbers: the Protection warrior's bare bands gain 2.0 to 3.3% threat.
- M1 moves every mana user at the level-60 raid preset.
- S1 moves only /sim's raid option at levels 35 to 49.

## Raid buffs

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | **The raid's Mana Spring was the plain totem, but a raid's Restoration shaman runs the improved one.**<br>- The preset has a rule: it uses a buff's improved form when the live tree has the talent and a raid would take it. Its notes said no improved form qualified.<br>- Restorative Totems (16187, Restoration tier 3, 5 ranks) reads "Increases the effect of your Mana Spring Totem by 25%". It is in the 70707 tree. forever-bugs #215 asks Blizzard to remove it; that issue is still open.<br>- Our own Restoration shaman raid build (`0000000000000000-055000000000000000-5530003315513151`) puts all 5 points in it.<br>- The preset's reason was "Restorative Totems is a different talent and is not assumed". That breaks its own rule.<br>- Now `mana_spring_totem:improved`. At 60 the raid gives 31.25 mp5 instead of 25. The engine's improved form is the client's ×1.25. | SpellEffect 16187 (aura 108, bp 5 a rank, masks 16384 and 8192); talents/shaman.json Restoration index 10; `core/buffs.go:347`. Test `TestRaidManaSpringIsTheRestorationShamansImprovedTotem` failed first. | CONFIRMED-BUG (client talent plus our own build) | old (S 3f162c29e, the first preset) | S ba04df89a |
| W1 | **The Protection warrior never cast Battle Shout.**<br>- At 60 the raid preset's shout covered for it.<br>- Below 60 the published sets are bare, with no raid warrior, so every band ran with no shout.<br>- Fury and Arms do shout: a prepull cast plus an upkeep line.<br>- The tank's rotation now has the same two lines, the upkeep line just before the Sunder filler.<br>- Threat at bands 20 to 50: +2.0 to +3.3%. At 60 the raid aura is already up, so the line idles and nothing moves. | `data/curated/apl/warrior-protection.json` has had no 11551 since 5114d47f3. Tank fight, horde, 10,000 iterations: band 40 367.4 → 374.8. Test `TestEveryWarriorRotationShouts` failed first (prepull false, upkeep false). The fork golden's NoBuffs rows move because the harness gear barely deals damage, so the shout's 10 rage outweighs its attack power there. | CONFIRMED-BUG | old (S 5114d47f3) | S 2bcc92053, E 7164910d7 |
| W2 | **Battle Shout gave threat. In Forever it gives none.**<br>- forever-bugs #64: a threat parser read 0 threat for Battle Shout, Demoralizing Shout and the paladin blessings.<br>- The reply (2026-09-22): "Many Buffs and Debuffs now cause 0 threat. These ones are all intended to cause 0 threat."<br>- The engine still gave the shout Era's flat threat (the rank's level), on the warrior and on five simulated party hits.<br>- Only threat moves. Fury and Arms DPS are identical. | `sim/warrior/shouts.go` (upstream d30d87d9c). Test `TestBattleShoutGeneratesNoThreat` failed first (270.4 threat on the boss). Goldens: tps rows of TestP1DPSWarrior and TestP1TankWarrior only. | CONFIRMED-BUG (a log plus a reply; Blizzard-attributed single source) | old (upstream wowsims) | E f1c3dbe2b |
| N1 | The warrior tank's options named Commanding Shout "for the health". Commanding Shout (403435) is a Season of Discovery rune with no SkillLineAbility row, so a Forever warrior cannot learn it. The engine reads no shout option at all. The field now says Battle Shout, as Fury's and Arms' do. No number moves. | SpellLevels/SkillLineAbility 403435; grep of the fork for `Options.Shout` (conformance presets only) | note fix | old | S f2b70764e |
| B1 | Every `BuffRanks` table re-checked: spell, learn level and amount for Arcane Intellect, Blessing of Might and Wisdom, Trueshot Aura, Mark of the Wild, Devotion Aura, Power Word: Fortitude, Divine Spirit, Blood Pact, Battle Shout, Thorns, Retribution Aura, Stoneskin and Mana Tide. All match. The script's mismatches are the totems (stated on their aura spells), the Judgements (stated on their triggered spells), Expose Armor (scripted), and two by-design cases: Mark of the Wild's zero stat and resistance rows on ranks that state none, and Heart of the Lion's level 0. Each was checked by hand. | `tools/ranks.py` over `engine/sim/core` | Aligned | - | - |
| B2 | **The level-60 book ranks are real Ahn'Qiraj books**, so leaving them out (`IncludeAQ = false`) is right. The ranks are Battle Shout 25289, Blessing of Might 25291, Blessing of Wisdom 25290, Strength of Earth 25361 and Grace of Air 25359. Their SkillLineAbility rows look exactly like trainer rows (acquire 0), so the client alone cannot tell a book rank apart. But the books are in ItemSparse: 21298 Manual of Battle Shout VII, 21289 Libram: Blessing of Might VII, 21288 Libram: Blessing of Wisdom VI, 21292 Tablet of Strength of Earth Totem V and 21293 Tablet of Grace of Air Totem III. | ItemSparse; SkillLineAbility | Aligned (first time this was checked against the books) | - | - |
| B3 | **A spec's own talent aura is not counted twice with the raid's copy.** Each pair was simmed at bare 50, with and without the raid flag, and came out identical: Moonkin Aura on Balance (191.9 both), Trueshot Aura on Marksmanship (213.4), Strength of Earth on Enhancement (155.4), Mana Spring on Elemental (152.1) and Leader of the Pack on Feral (152.7). In the raid runs each aura appears once. | spec-breakdown, 2,000 iterations | Aligned | - | - |
| B4 | Heart of the Lion is 280 attack power at 60 in the engine and Wowhead's tooltip. The preset note says 276. TrinityCore subtracts SpellLevel (1) for a non-passive spell, which gives 276; the tooltip gives 280. The 4 attack power is inside error. | SpellLevels 409583 (SpellLevel 1, BaseLevel 0); `core/buff_ranks.go` | OPEN (minor; a level-60 character sheet settles it) | - | - |
| B5 | Blessing of Kings with Heart of the Lion: still no source. Both Forever tooltips, warcraft.wiki.gg and forever-bugs say nothing about stacking. | searches 2026-11-02 | OPEN (round 7) | - | - |
| B6 | Battle Shout rank 6 registers under action id 27578, a level-52 duplicate row. Its 185 attack power, 1 a level and zero cost are not used: the grant reads `BattleShoutRankTable` (115 at 60) and `battleShoutRageCost` falls back to 100. Only the displayed id differs. | `warrior/shouts.go`, constants_auto_gen.go:226 | Aligned (cosmetic) | - | - |
| B7 | No other improved form qualifies. The live trees were read for every graded id: Improved Seals is seal damage, not Judgement of the Crusader's debuff; Improved Sunder Armor and Improved Thunder Clap are rage cost; Guardian Totems is Stoneskin, which no preset carries; Improved Expose Armor is a cost. | talents/*.json | Aligned | - | - |

## Raid debuffs

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | Sunder Armor 450 ×5 (major), Faerie Fire 505 and Curse of Recklessness 505 (each additive), Curse of the Elements by rank (no level-60 rank exists), Hunter's Mark 143, Judgement of the Crusader 161: re-read by rank and learn level. All match. | `tools/ranks.py` rows; `core/debuffs.go` | Aligned | - | - |
| D2 | The armor stacking and caster-only debuffs (Shadow Weaving, Fire Vulnerability, Winter's Chill, Stormstrike) stay as round 10 left them. No new evidence. | round 10 D1, D4 | Aligned | - | - |

## Consumables

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Raid rows at levels 35 to 49 left the main hand bare** (found by the regression lane).<br>- Affected: Fury, Arms, the warrior tank, Retribution, the cat, the bear and Survival, plus the rogues' stone beside the poison.<br>- The off hand had a Dense Sharpening Stone from 35, but the main hand waited for the Elemental Sharpening Stone at 50.<br>- Each main-hand list now has the Dense stone (12404, or Dense Weightstone 12643 on a mace, both 8 damage, level 35) after the Elemental one. Under the preset's first-usable-entry rule, 50 and up are unchanged.<br>- A rogue's stone beside its poison now follows the same first entry, so the rogue never gets both stones.<br>- Only /sim's raid option at 35 to 49 moves. Level 60 is identical for every spec. | ItemSparse 12404, 12643, 18262 RequiredLevel 35, 35, 50. Test `TestMeleeMainHandsCarryTheDenseStoneUntilTheElementalStone` failed first (9 rows). | CONFIRMED-BUG | old (warriors, Retribution, druids, rogues); R9 S be3314ab5 (Survival) | S 45594522d |
| C1 | **Forever's new elixirs: still open.**<br>- Grizzly is 25 Strength and 2% crit (1250986), Ferocity 18 Strength and 18 Agility, Owl 25 Intellect and 2% crit, Whale 25 Spirit and 2% crit, Cunning 25 Agility and 25 Intellect, Nature Power 40 nature.<br>- Every one is in the 3 s category 2593 that Mongoose, Greater Arcane and the Classic elixirs share. The client has no Battle or Guardian split.<br>- Neither the tooltips nor warcraft.wiki.gg's Forever pages (2026-10-27) say whether they stack with Mongoose-style elixirs. The wiki's "Battle Elixir" line on Mongoose is retail text. | ItemEffect rows for items 250328 to 250354; Wowhead Forever tooltips | OPEN (a beta character sheet with Mongoose and Grizzly up) | - | - |
| C2 | **The Natural flasks** (274273 to 274276, level 55) give 60 Stamina, plus 4 or 5% crit, hit, haste or dodge and parry reduction "while in Mount Hyjal, Hyjal Summit, and the Barrow Deeps". Those are raid zones that open later (9 December), so the flasks stay out of a Phase 1 preset. | Wowhead Forever tooltips 1293740 to 1293743 | Aligned | - | - |
| C3 | **Mender's Potion** (Major 1251948: +88 healing for 30 s, 2-minute potion timer) is not in the engine. The healers are mana-bound: the Holy paladin's mana lasts 257 of 300 s. Major Mana Potion is about 4,460 mana a fight, about 19,600 healing at its 4.4 healing per mana. Mender's adds about 2,000 (88 healing at about 0.6 coefficient on about 15 casts a window, 2.5 windows). So the mana potion keeps the slot. | `sb -heal` Holy paladin: income table, healed per mana | Aligned (estimate, about 10 to 1; not simmed) | - | - |
| C4 | Brilliant Wizard Oil also gives +36 healing (25113, aura 135), and the engine applies it. Brilliant Mana Oil's 15 mp5 (about 900 mana, about 4,000 healing a fight) beats the wizard oil's extra 6 healing and 1% crit (under 2,000) for a mana-bound healer. The healers keep the mana oil. | SpellEffect 25113, 25116 | Aligned (estimate) | - | - |
| C5 | Elixir of Greater Firepower is still absent: its recipe 21547 is in ItemSparse, but the elixir 21546 is not. Fire keeps Elixir of Fire Power (+10, 7844). | ItemSparse | Aligned | - | - |
| C6 | Every preset consumable maps to the right client item and RequiredLevel. Where two items share a name, the item-level join picks the right one: Elixir of Fortitude resolves to Forever's 250334 (45, 200 health), not Era's.<br>- Level 0: Juju Power, Juju Might, R.O.I.D.S., Scorpok, Cortex, Gizzard Gum, Rumsey Rum, Demonic Rune. Thistle Tea is 5.<br>- Spot amounts: Sage's Tea 44 healing (1249925 → 1249927), Rumsey Rum 15 Stamina, Gizzard Gum 25 Spirit, Brilliant Mana Oil 15 mp5 and 30 healing, Dense stone +8. | `items.json` join script; SpellEffect | Aligned | - | - |
| C7 | Lower potion tiers below 55 in the raid option. The client has Greater (35, 40 attack power) and Superior (45, 56) Frenzy Potions, Greater (35, 23) and Superior (45, 33) Spellblasting Potions, and Greater and Superior Mana Potions. The preset's stand-ins are Frenzy Potion (28) from 25, and no mana potion below 49. Only /sim's raid option below 55 is affected, and the engine has none of the tier potions. | ItemSparse 250935, 250936, 250941, 250942; preset `by_class` | OPEN (whether the Greater and Superior recipes are on the launch vendors: one vendor stock list settles it, then the engine needs the enum values) | - | - |

## Presets: /sim, the ranker and spec-breakdown

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | All 28 specs carry the same set at 60 on every path. spec-breakdown's consumables and buffs equal `sim/request/raid-consumes.json`'s level-60 row for every spec (56 lists, 0 differences). The ranker resolves the same `presets.json` through `Resolve`, `UsableAt` and `Layer`. The page reads the generated `raid-consumes.json` and `bare-preset.json`, which the Go tests hold to the code. | `sb-sets.txt` compare script; `sim/request/raid_consumes.go`; `web/scripts/sync-raid-preset.mjs` | Aligned | - | - |
| P2 | **The bare kits, spec by spec.** Self-buffs at bare 50 were read from the aura tables:<br>- Mage: Mage Armor.<br>- Warlock: Demon Armor.<br>- Hunters: Aspect of the Hawk 14321 and Heart of the Lion, plus Trueshot Aura for Marksmanship.<br>- Enhancement: Strength of Earth 8161 and Grace of Air 8835, the ranks a level-50 shaman knows.<br>- Elemental: Mana Spring 10496.<br>- Shadow: Shadowform.<br>- Retribution: Blessing of Might 19837.<br>- Balance: Moonkin Form. Feral: Cat Form, with Leader of the Pack from its talent.<br>- Healing priests: Inner Fire.<br>- Holy paladin: Concentration Aura.<br>- Protection paladin: Righteous Fury.<br>The one gap was the warrior tank's Battle Shout (W1). | spec-breakdown aura tables; `sim/request/request.go` class options | Aligned after W1 | - | - |
| P3 | The published BiS json files still list Improved Scorch among the raid debuffs (`presets.raid.debuffs`). That is held output from before d1ddb012a removed it. The sims read `presets.json`, so no number uses it. It clears at the owner's refresh. | data/builds/1.60.1.70707/bis/*.json | Aligned (held output) | - | - |

## Round 10's open items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Concentration Aura: the Holy paladin already runs it (`holyPaladinOptions`, `PaladinAura_ConcentrationAura`). The only open question is whether a raid's Holy paladin would pick it over another aura, which is a kit choice. | `sim/request/request.go` | OPEN (owner) | - | - |
| A2 | forever-bugs #419 (shadow wands ignore Shadow Weaving and Shadowform): still open, no reply, last updated 2026-10-21. No published rotation shoots a shadow wand. | gh issue 419 | OPEN | - | - |
| A3 | Rallying Cry and Slip'kik's Forever versions: no new evidence. The Forever tooltips' "No effect while in a raid instance" supports keeping world buffs gated off. | round 10 crit.md "For other lanes" | OPEN (gated off; unreachable) | - | - |

## Before/after

"Before" is `$R/sb-before` from `$R/site-before/sim`. "After" is lane engine f1c3dbe2b with site ba04df89a. 10,000 iterations. Tanks are in threat per second (tank fight), healers in effective HPS, everyone else in DPS.

**Protection warrior, bare bands (W1 and W2):**

| band | horde s1 / s50001 | alliance s1 / s50001 | horde s1 absolute |
|---|---|---|---|
| 20 | +2.9 / +2.9 | +2.9 / +2.8 | 73.3 → 75.4 |
| 30 | +3.1 / +3.3 | +3.3 / +3.2 | 140.8 → 145.1 |
| 40 | +2.0 / +2.0 | +2.0 / +2.0 | 367.4 → 374.8 |
| 50 | +2.3 / +2.2 | +2.3 / +2.3 | 520.4 → 532.6 |
| 60 raid | 0.0 / 0.0 | 0.0 / 0.0 | 1067.5 → 1067.5 |

Fury and Arms at 60: DPS identical (757.8 and 720.8, horde seed 1). W2 moves only threat.

**Every spec at 60, raid preset (M1):**

| spec | horde s1 / s50001 | alliance s1 / s50001 | horde s1 absolute |
|---|---|---|---|
| druid-balance | +0.10 / +0.03 | +0.07 / +0.18 | 571.6 → 572.2 |
| druid-feral-bear | +0.00 / +0.00 | +0.00 / +0.00 | 992.4 → 992.4 |
| druid-feral | +0.03 / +0.04 | +0.03 / +0.04 | 670.0 → 670.2 |
| druid-restoration | +0.45 / +0.45 | +0.45 / +0.40 | 830.0 → 833.7 |
| hunter-beast-mastery | +0.02 / -0.02 | -0.02 / +0.00 | 633.4 → 633.5 |
| hunter-marksmanship | +0.00 / +0.02 | -0.02 / -0.02 | 623.0 → 623.0 |
| hunter-survival | +0.00 / +0.00 | +0.00 / +0.00 | 745.3 → 745.3 |
| mage-arcane | +0.15 / +0.20 | +0.19 / +0.03 | 600.4 → 601.3 |
| mage-fire | +0.21 / +0.17 | +0.25 / +0.25 | 660.1 → 661.5 |
| mage-frost | +0.03 / +0.06 | -0.03 / -0.05 | 623.0 → 623.2 |
| paladin-holy | +0.58 / +0.61 | +0.44 / +0.43 | 641.4 → 645.1 |
| paladin-protection | +0.03 / +0.03 | +0.02 / +0.01 | 1013.5 → 1013.8 |
| paladin-retribution | +0.07 / -0.03 | +0.06 / +0.01 | 685.7 → 686.2 |
| priest-discipline | +0.64 / +0.72 | +0.77 / +0.82 | 625.3 → 629.3 |
| priest-holy | +0.42 / +0.40 | +0.44 / +0.38 | 693.8 → 696.7 |
| priest-shadow | +0.03 / +0.03 | +0.03 / +0.00 | 717.8 → 718.0 |
| rogue-assassination | +0.00 / +0.00 | +0.00 / +0.00 | 619.7 → 619.7 |
| rogue-combat | +0.00 / +0.00 | +0.00 / +0.00 | 617.4 → 617.4 |
| rogue-subtlety | +0.00 / +0.00 | +0.00 / +0.00 | 544.5 → 544.5 |
| shaman-elemental | +0.09 / +0.09 | +0.02 / +0.09 | 466.9 → 467.3 |
| shaman-enhancement | +0.00 / +0.00 | +0.00 / +0.00 | 580.8 → 580.8 |
| shaman-restoration | +0.74 / +0.80 | +0.77 / +0.74 | 834.0 → 840.2 |
| warlock-affliction | +0.19 / +0.16 | +0.16 / +0.19 | 571.0 → 572.1 |
| warlock-demonology | +0.16 / +0.23 | +0.14 / +0.27 | 640.4 → 641.4 |
| warlock-destruction | +0.17 / +0.09 | +0.06 / +0.13 | 636.9 → 638.0 |
| warrior-arms | +0.00 / +0.00 | +0.00 / +0.00 | 720.8 → 720.8 |
| warrior-fury | +0.00 / +0.00 | +0.00 / +0.00 | 757.8 → 757.8 |
| warrior-protection | +0.00 / +0.00 | +0.00 / +0.00 | 1067.5 → 1067.5 |

What moved: the mana users, from the improved Mana Spring (M1).
- Healers: Restoration shaman +0.7 to +0.8, Discipline +0.6 to +0.8, Holy paladin +0.4 to +0.6, Restoration druid and Holy priest about +0.4.
- Warlocks and Fire and Arcane mages: +0.1 to +0.3.
- Every other spec: within 0.1 (seed noise). Mana Spring reaches a non-mana user as nothing.

## For other lanes

- **Platform (web/src/lib/sim):** M1 puts the first graded id, `mana_spring_totem:improved`, into the static raid preset. Three vitest cases fail without a small patch, because they assumed every preset id was plain:
  - `preset-summary.test.ts` (2 cases): `presetSummary` keeps only plain row ids.
  - `settings.test.ts` "reports real totals in the group counters": `selectedIn` names a graded row by its plain id.
  - `raid-buff-names.test.ts` keys its aura map by the plain id.
  - The patch (presetSummary matches the plain or the improved form; the two tests compare plain ids) is at `$R/auras/web-improved-mana-spring.patch`. With it, all 1,147 web/src/lib/sim tests pass.
  - `buffLabel` would show "Mana spring totem: Improved". A nicer label for graded ids is the platform lane's call.
  - The custom panel's three-way grade row already handles the id.
- **Integrator:**
  - TestRotationLadder's warrior-protection golden moves with W1. It was not regenerated: every band's DPS rises (for example 46.9 → 49.1 at 40) and the shout appears in the cast list.
  - The fork goldens moved by W1 and W2 were regenerated in the fork commits: tank NoBuffs rows for W1, and the dps and tank tps rows for W2.
  - Specs outside this lane that move: none outside the presets. M1 moves every mana user at raid 60 (table above). W1 moves warrior-protection's bare bands. W2 moves warrior threat only.
- **Physical lane (warriors):** the warrior APL edit (W1) touches `data/curated/apl/warrior-protection.json` and the fork copy. Merge any physical-lane edit of the same file on top.
- **Healers lane:** forever-bugs #215 may remove Restorative Totems before launch. If it goes, M1 reverts to the plain totem, and the Restoration shaman's own build loses those 5 points.

## Still open

1. C1: whether Forever's new elixirs stack with Mongoose, Greater Arcane and Juju Power.
2. B5: Blessing of Kings with Heart of the Lion (round 7).
3. C7: the Greater and Superior 30-second potions and mana potions as stand-ins below 55 (vendor stock first).
4. B4: Heart of the Lion 276 or 280 at 60.
5. A1: Concentration Aura in the Holy paladin's kit (owner).
6. A2: forever-bugs #419.
7. A3: Rallying Cry and Slip'kik's Forever versions (world buffs gated off).
