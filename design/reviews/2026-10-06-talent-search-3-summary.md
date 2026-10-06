# Talent search 3: engine fork models every class talent it used to skip

Branch `talent-search-3`, engine bebda865a (pinned in `sim/go.mod`), client 1.60.1.70009. The
classifier itself (alias-aware static scan + probe precedence, from talent-search-2) is unchanged;
what changed is the engine fork's talent coverage — the previous run's "Engine gaps" lists have
shrunk across every class because the fork now reads talents it used to skip entirely, and its
spell constants were regenerated from 1.60.1.70009.

## Re-run

All 20 written-rotation DPS specs re-run with the tool's defaults (300 screening iterations, 800
confirmation iterations, top 10, seed 7), overwriting `design/reviews/talent-search/*.md`. Each spec
fans out to 14 concurrent sims and saturates the box, so specs ran back-to-back, not as parallel
processes. Total time: 170s for all 20.

## Table 1: per-spec results

"Keep-winner" is the best finalist that drops no talent the final classification still calls
unmodeled (`report.go`'s `bestClean`) — by construction it never removes an engine-gap talent, so
it is always the strongest build the engine's own numbers can back without betting on an
unmodeled talent actually doing something. Gain % and beyond-error are the keep-winner against the
guide. "Own tree max" asks whether the spec's own tree still holds the most points in the
keep-winner build.

| Spec | Guide DPS | Keep-winner DPS | Keep-winner code | Gain % | Beyond error | Points/tree (own tree) | Own tree max | Remaining engine gaps in own tree |
|---|---|---|---|---|---|---|---|---|
| Balance | 170.3 | 181.1 | `FS1:1.60.1.70009:druid:night-elf:522221111540105/0/55333:` | +6.3% | yes | 32/0/19 (Balance) | yes | Improved Wrath, Genesis, Improved Moonfire, Improved Entangling Roots, Overgrowth, Eclipse |
| Feral | 219.7 | 219.7 | `FS1:1.60.1.70009:druid:night-elf:05002/5503202023032210001/5050003:` | +0.0% | no | 0/31/20 (Feral Combat) | yes | Feral Swiftness, Feral Instinct, Brutal Impact, Thick Hide, Feral Charge, Primal Bite, King of the Jungle, Natural Reaction, Rend and Tear |
| Beast Mastery | 233.2 | 253.1 | `FS1:1.60.1.70009:hunter:dwarf:5420001505001251/0053502001/4:` | +8.5% | yes | 31/16/4 (Beast Mastery) | yes | Endurance Training |
| Marksmanship | 227.7 | 246.7 | `FS1:1.60.1.70009:hunter:dwarf:55200004/0050550011500305/5:` | +8.4% | yes | 16/30/5 (Marksmanship) | yes | Careful Aim, Rapid Killing, Improved Arcane Shot, Lone Wolf, Improved Serpent Sting, Rapid Recuperation |
| Survival | 275.3 | 291.2 | `FS1:1.60.1.70009:hunter:dwarf:0/32005500005/500230131050220151:` | +5.8% | yes | 0/20/31 (Survival) | yes | Survivalist, Expose Prey, Survivalist's Discipline |
| Arcane | 456.4 | 498.5 | `FS1:1.60.1.70009:mage:gnome:253225113100011531/032023/005:` | +9.2% | yes | 36/10/5 (Arcane) | yes | Wand Specialization, Improved Channeling, Magic Absorption, Arcane Concentration, Arcane Resilience, Arcane Geometry, Arcane Shielding, Improved Counterspell, Arcane Meditation |
| Fire | 381.4 | 414.8 | `FS1:1.60.1.70009:mage:gnome:2050151/23552100030023051/005:` | +8.7% | yes | 14/32/5 (Fire) | yes | Flame Throwing, Impact, Improved Fire Ward, Hot Streak |
| Frost | 218.3 | 325.4 | `FS1:1.60.1.70009:mage:gnome:2030050001/113023/253511130000030105:` | +49.1% | yes | 11/10/30 (Frost) | yes | Frost Warding, Permafrost, Frostbite, Improved Blizzard, Arctic Reach, Ice Block, Shatter, Fingers of Frost |
| Retribution | 234.5 | 250.3 | `FS1:1.60.1.70009:paladin:human:54003/323/0502533100133032:` | +6.7% | yes | 12/8/31 (Retribution) | yes | Deflection, Improved Judgement, Seal of Command, Twist of Light |
| Shadow | 282.7 | 288.1 | `FS1:1.60.1.70009:priest:gnome:5241110013/0/543110401201300251:` | +1.9% | yes | 18/0/33 (Shadow) | yes | Blackout, Spirit Tap, Shadow Affinity, Shadow Reach, Improved Psychic Scream, Improved Fade, Silence, Early Demise |
| Assassination | 193.7 | 211.8 | `FS1:1.60.1.70009:rogue:night-elf:3250001055050105/325201/51:` | +9.3% | yes | 32/13/6 (Assassination) | yes | Improved Gouge, Remorseless Attacks, Improved Slice and Dice, Improved Expose Armor, Vigor, Improved Kidney Shot, Seal Fate |
| Combat | 234.8 | 234.8 | `FS1:1.60.1.70009:rogue:night-elf:005323101003/32520300001515231/0:` | +0.0% | no | 14/31/6 (Combat) | yes | Lightning Reflexes, Puncturing Wounds, Deflection, Endurance, Riposte, Improved Sprint, Improved Kick |
| Subtlety | 183.1 | 183.1 | `FS1:1.60.1.70009:rogue:night-elf:005323101005/32520300001515211/0:` | +0.0% | no | 5/15/31 (Subtlety) | yes | Camouflage, Master of Deception, Opportunity, Setup, Elusiveness, Dirty Tricks, Improved Ambush, Initiative, Improved Distract, Heightened Senses, Dirty Deeds, Cutthroat, Thousand Cuts |
| Elemental | 152.2 | 222.6 | `FS1:1.60.1.70009:shaman:dwarf:552233130010305/055000001/05002:` | +46.3% | yes | 33/11/7 (Elemental) | yes | Reverberation, Elemental Devastation, Elemental Alacrity, Eye of the Storm |
| Enhancement | 200.6 | 231.4 | `FS1:1.60.1.70009:shaman:dwarf:550333/254130031005002051/0:` | +15.4% | yes | 19/32/0 (Enhancement) | yes | Guardian Totems, Improved Lightning Shield, Anticipation, Toughness, Spirit Weapons, Mental Quickness, Improved Stormstrike, Maelstrom Weapon |
| Affliction | 360.2 | 360.2 | `FS1:1.60.1.70009:warlock:gnome:25550320100201351/0005/0550001:` | +0.0% | no | 32/19/0 (Affliction) | yes | Soul Harvesting, Fel Concentration, Pandemic, Curse of Exhaustion |
| Demonology | 305.7 | 305.7 | `FS1:1.60.1.70009:warlock:gnome:003/235510013002000035/0550005003:` | +0.0% | no | 20/31/0 (Demonology) | yes | Improved Health Funnel, Improved Imp, Demonic Embrace, Demonic Aegis, Improved Voidwalker, Fel Vitality, Demonic Energies, Demonic Sacrifice, Master Summoner, Fel Domination, Demonic Brand, Improved Felhunter, Soul Link, Demonic Pact |
| Destruction | 300.3 | 318.8 | `FS1:1.60.1.70009:warlock:gnome:255323/0/235322510110105:` | +6.2% | yes | 20/0/31 (Destruction) | yes | Destructive Reach, Improved Shadow Bolt, Molten Skin, Cataclysm, Intensity, Pyroclasm |
| Arms | 202.0 | 227.3 | `FS1:1.60.1.70009:warrior:human:05325213032310001/0505/5005:` | +12.5% | yes | 31/10/10 (Arms) | yes | Deflection, Improved Charge, Improved Tactical Mastery, Sweeping Strikes, Improved Slam |
| Fury | 260.9 | 264.5 | `FS1:1.60.1.70009:warrior:human:35311103002/350511005050010051/0:` | +1.4% | yes | 19/32/0 (Fury) | yes | Booming Voice, Blood Craze, Enrage, Improved Berserker Rage |

Every spec's own tree still holds the most points in its keep-winner. The "no remaining engine-gap
talent the winner drops that the guide takes" leg of the adoption rule is satisfied by construction
for all 20: `bestClean` is defined to never drop a talent the final classification calls unmodeled,
so a keep-winner can never drop an engine-gap talent, damage-classified or not.

## Table 2: adoption decisions

ADOPT requires: gain beyond error, gain ≥ 1%, own tree still holds the most points, and no remaining
engine-gap talent in the spec's own tree is a damage talent the guide takes that the winner drops
(always true here, per the note above). Otherwise KEEP.

| Spec | Decision | Gain % | Reason |
|---|---|---|---|
| Balance | ADOPT | +6.3% | beyond error, ≥1%, own tree max |
| Feral | KEEP | +0.0% | guide is already the keep-winner; nothing to swap into beyond error |
| Beast Mastery | ADOPT | +8.5% | beyond error, ≥1%, own tree max |
| Marksmanship | ADOPT | +8.4% | beyond error, ≥1%, own tree max |
| Survival | ADOPT | +5.8% | beyond error, ≥1%, own tree max |
| Arcane | ADOPT | +9.2% | beyond error, ≥1%, own tree max |
| Fire | ADOPT | +8.7% | beyond error, ≥1%, own tree max |
| Frost | ADOPT | +49.1% | beyond error, ≥1%, own tree max |
| Retribution | ADOPT | +6.7% | beyond error, ≥1%, own tree max |
| Shadow | ADOPT | +1.9% | beyond error, ≥1%, own tree max |
| Assassination | ADOPT | +9.3% | beyond error, ≥1%, own tree max |
| Combat | KEEP | +0.0% | guide is already the keep-winner |
| Subtlety | KEEP | +0.0% | guide is already the keep-winner |
| Elemental | ADOPT | +46.3% | beyond error, ≥1%, own tree max |
| Enhancement | ADOPT | +15.4% | beyond error, ≥1%, own tree max |
| Affliction | KEEP | +0.0% | guide is already the keep-winner |
| Demonology | KEEP | +0.0% | guide is already the keep-winner |
| Destruction | ADOPT | +6.2% | beyond error, ≥1%, own tree max |
| Arms | ADOPT | +12.5% | beyond error, ≥1%, own tree max |
| Fury | ADOPT | +1.4% | beyond error, ≥1% (barely), own tree max |

15 ADOPT, 5 KEEP (Feral, Combat, Subtlety, Affliction, Demonology — the same five specs whose guide
was already the clean-build best in talent-search-2).

## Table 3: what changed versus talent-search-2

Decisions flipped KEEP → ADOPT for 7 specs: the gain either didn't exist in talent-search-2 or sat
under the 1% floor, and a newly-modeled talent in the spec's own tree opened it up.

| Spec | v2 gain % | v3 gain % | Newly-modeled talent(s) responsible (probe diff) |
|---|---|---|---|
| Beast Mastery | +0.0% | +8.5% | Deadly Aspects (+14.2 ± 0.8, damage), Focused Fire (+4.6 ± 0.9, damage) |
| Marksmanship | +0.0% | +8.4% | Trueshot Aura (+13.1 ± 1.1, damage), Lethal Attacks (+6.0 ± 1.1, damage) |
| Survival | +0.8% | +5.8% | Resourcefulness (+19.8 ± 1.5, damage), Improved Tracking (+11.6 ± 1.5, damage) |
| Retribution | +0.0% | +6.7% | Sanctified Judgement (+30.4 ± 1.2, damage), Holy Conduit (+10.1 ± 1.2, damage), Champion of the Light (+9.9 ± 1.2, damage) |
| Enhancement | +0.0% | +15.4% | Shamanistic Focus (+22.1 ± 2.0, damage), Rage of the Farseer (+9.6 ± 1.9, damage) |
| Destruction | +0.4% | +6.2% | Agonizing Flames (+21.7 ± 1.4, damage), Shadow and Flame (+15.4 ± 1.3, damage), Fire and Brimstone (+8.7 ± 1.4, damage) |
| Fury | +0.5% | +1.4% | Raging Blows (+4.2 ± 1.7, damage) — just enough to clear the 1% floor |

Decisions held ADOPT in both runs, but the gain magnitude moved because the spec's own tree picked
up newly-modeled damage talents the guide already specs:

| Spec | v2 gain % | v3 gain % | Newly-modeled talent(s) responsible |
|---|---|---|---|
| Balance | +2.6% | +6.3% | Nature's Splendor (+11.0 ± 2.1), Nature's Majesty (+5.0 ± 1.9), Nature's Reach (+2.9 ± 2.1), Nature's Grace (+2.5 ± 2.0), all now `damage` |
| Elemental | +12.0% | +46.3% | Lightning Overload (+5.0 ± 1.6, damage), Call of Thunder (+2.0 ± 1.6, damage); Earthbound/Elemental Reach/Elemental Warding/Improved Fire Nova moved to `modeled, no damage` |
| Arms | +4.4% | +12.5% | Bloodthrill (+37.2 ± 1.5, damage) — the single largest swing of this run |

Decisions held ADOPT with essentially no change in the gain (no own-tree talent crossed into
`damage`): Arcane (+9.2% → +9.2%), Fire (+8.8% → +8.7%, same build and DPS, rounding only), Frost
(+49.1% → +49.1%, identical build and DPS — nothing in Frost's own tree was newly modeled),
Assassination (+9.3% → +9.3%, though Ruthlessness is newly modeled at a slightly negative
−1.3 ± 1.1, so the winner moved that 1 point elsewhere), Shadow (+2.0% → +1.9%, Improved Mind Flay
is now `damage` at +27.5 ± 1.7 but the guide already spent its 2 points there, so guide and winner
DPS both rose together and the gap barely moved; Devouring Contagion is now `modeled, no damage`).

Decisions held KEEP in both runs, but guide DPS itself rose because the guide already specced
talents that are now `damage`: Feral (Blood Frenzy +11.3 ± 0.8, Shredding Attacks +7.9 ± 0.9,
Predatory Instincts +2.3 ± 0.9 — guide 189.4 → 219.7), Combat (Hack and Slash +7.0 ± 1.1, Flawless
Execution +5.5 ± 1.1 — guide 227.9 → 234.8), Affliction (Soul Siphon +21.8 ± 1.3, Improved Drains
+16.5 ± 1.3, Malediction +9.2 ± 1.3, Improved Bane of Agony +4.7 ± 1.3, Malevolence +2.8 ± 1.3 —
guide 351.0 → 360.2), Demonology (Demonic Knowledge +18.1 ± 1.3 — guide 283.0 → 305.7), Subtlety
(Quietus is now `damage` at +1.5 ± 1.0, but the guide was already at its 5/5 max rank, so guide DPS
is unchanged, 183.1 → 183.1).

No gain vanished between the two runs — every spec that was beyond error in talent-search-2 is still
beyond error here, and every KEEP spec is still KEEP. Several engine-gap lists also picked up
talents the classifier hadn't seen under the old names (e.g. Marksmanship's "Improved Serpent
Sting", Retribution's "Improved Judgement", Assassination's "Vigor", Destruction's "Cataclysm") —
these read as talent renames from the 1.60.1.70009 spell-constant regeneration, not new unmodeled
talents the engine skipped before.

## Tests

```
cd sim && go vet ./cmd/talent-search/... && go test ./cmd/talent-search/...
ok  	github.com/jhunthrop/foreversixty/sim/cmd/talent-search	0.443s
```

## Reports

`design/reviews/talent-search/*.md` (all 20 overwritten).
