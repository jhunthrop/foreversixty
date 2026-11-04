# Guide accuracy round 1: Shaman

Auditor: shaman (Elemental, Enhancement, Restoration, class index). Audit tree origin/main 64168181f, client 1.60.1.70755, engine pin 2eaae03c6.
Files: `web/src/content/guides/shaman/{index,elemental,enhancement,restoration}.md`.

## Builds decoded (70755 tree order, `data/builds/1.60.1.70755/talents/shaman.json`, tier then column)

All five strings are legal on 70755: 51 points each, no talent over max rank, every tier gate met (5 points per tier spent below), every prerequisite met (Lightning Overload needs Elemental Alacrity 3, Elemental Fury needs Lightning Overload 3, Lightning Mastery needs Call of Thunder, Flurry needs Mental Dexterity 3, Rage of the Farseer needs Mental Quickness 2, Riptide needs Healing Way 3). No removed talent (Lava Burst, Elemental Mastery) is referenced; their slots now hold Call of Thunder (t4c1) and Lightning Mastery (t6c1).

**Elemental `build:`** `5510331300001351/05/053052`
- Elemental 31: Convection 5, Concussion 5, Elemental Devastation 1, Elemental Alacrity 3, Call of Flame 3, Elemental Focus 1, Elemental Warding 3, Call of Thunder 1, Lightning Overload 3, Elemental Fury 5, Lightning Mastery 1
- Enhancement 5: Thundering Strikes 5
- Restoration 15: Tidal Focus 5, Mindfulness 3, Totemic Focus 5, Improved Reincarnation 2

**Elemental `raidBuild:`** `5510331300001351/055/053002`
- Elemental 31: as above
- Enhancement 10: Thundering Strikes 5, Ancestral Knowledge 5
- Restoration 10: Tidal Focus 5, Mindfulness 3, Improved Reincarnation 2

**Enhancement `build:`** `320033/055030310051202051/052`
- Elemental 11: Convection 3, Concussion 2, Elemental Alacrity 3, Call of Flame 3
- Enhancement 33: Thundering Strikes 5, Ancestral Knowledge 5, Mental Dexterity 3, Elemental Weapons 3, Shamanistic Focus 1, Flurry 5, Stormstrike 1, Improved Ghost Wolf 2, Mental Quickness 2, Maelstrom Weapon 5, Rage of the Farseer 1
- Restoration 7: Tidal Focus 5, Mindfulness 2

**Enhancement `raidBuild:`** `0533001/053030310051202051/053`
- Elemental 12: Concussion 5, Elemental Devastation 3, Reverberation 3, Elemental Focus 1
- Enhancement 31: Thundering Strikes 5, Ancestral Knowledge 3, Mental Dexterity 3, Elemental Weapons 3, Shamanistic Focus 1, Flurry 5, Stormstrike 1, Improved Ghost Wolf 2, Mental Quickness 2, Maelstrom Weapon 5, Rage of the Farseer 1
- Restoration 8: Tidal Focus 5, Mindfulness 3

**Restoration `build:`** `0/055/5530003315513151`
- Enhancement 10: Thundering Strikes 5, Ancestral Knowledge 5
- Restoration 41: Improved Healing Wave 5, Tidal Focus 5, Mindfulness 3, Ancestral Healing 3, Healing Focus 3, Water Shield 1, Tidal Mastery 5, Restorative Totems 5, Mana Tide Totem 1, Healing Way 3, Nature's Swiftness 1, Purification 5, Riptide 1

## What checked out (no finding)

Every talent the prose names is in the decoded build with the stated points, and every per-rank value the prose quotes matches the 70755 max-rank text: Convection 10%, Elemental Alacrity 0.5 s, Call of Flame 15% and +9 s (12 to 21 s Flame Shock, 3.5 shock cooldowns), Lightning Overload 30%, Elemental Fury +100% crit bonus, Mindfulness 50%, Improved Reincarnation 20 min off (60 to 40), Mental Dexterity 150%, Elemental Weapons 40%, Flurry 30%, Stormstrike 150% weapon damage, 6 s cooldown (trainables 70753 8000 to 70755 6000), Mental Quickness 30%, Maelstrom 20% a stack, Improved Ghost Wolf 3.0 s (Ghost Wolf 3000 ms, so instant), Healing Focus 70%, Restorative Totems 25%/50%, Healing Way 25%, Ancestral Healing 25%, Purification 10%, Riptide no direct heal, 15 s, 12 s cooldown (70755 change), Lightning Mastery 15 s / 180 s, Rage of the Farseer 25 s / 180 s, Nature's Swiftness 180 s, Mana Tide 300 s.

Spell facts against `spellconst/shaman.json` and `trainables/shaman.json`: Lightning Bolt and Chain Lightning 2500 ms base (Chain Lightning moved 2000 to 2500 in 70755, the guides already say 2.5), Chain Lightning 6 s cooldown, LB rank 10 220 mana / 353 base, rank 9 190 / 295 (16% less, same 0.714 coefficient), Chain Heal rank 1 260 / 251 and rank 3 405 / 490, Mana Spring 10 mana per 2 s for 300 s, Healing Stream 300 s, Searing Totem rank 6 55 s, Magma Totem 20 s, Fire Nova 10 s cooldown, Totemic Call 22 with a 30 s cooldown, Call of the Elements 20, Ghost Wolf 20, Flametongue 10, Windfury Weapon 30, Strength of Earth 10, Searing 10, Chain Lightning 32, Chain Heal 40, Lesser Healing Wave 20, Healing Stream 20, Cure Poison 16, Cure Disease 22, Poison/Disease Cleansing 22/38. Strength of Earth, Grace of Air, Flametongue and resistance totems use effect 65 (raid area aura); Healing Stream, Mana Spring and Mana Tide use effect 35 (party). No guide mentions Lava Burst, Elemental Mastery, Totemic Projection, Earth Shield (class_spell only, not learnable) or any SoD rune.

Racials match `races.json` (Blood Fury AP/RAP/SP 10% 15 s / 2 min, Berserking 10% 15 s / 3 min, Artisan's Precision 1% hit, Stoneform 10% physical, Endurance 5% health, Wind Blessed 2% haste); combos.json has exactly Orc, Dwarf, Tauren, Troll, Windshaper Skyborne as shamans. Windfury Weapon disabling the totem for yourself is the weapon's own text (engine `sim/core/buffs.go:1680-1697`). Class rules hold: no dual wield anywhere, mail at 40, fist weapons named. Item sources check: Frozen and Molten Heart of the Mountain and Totem of Thunder are Enchanting, bind on pickup; Totem of Urgency is Enchanting, binds on equip, "+4% crit on Lesser Healing Wave"; Living Crown and Mooncloth Boots bind on equip; Forest Defender's Axe is a two-handed Blacksmithing craft; Devout set sources match the Restoration BiS file. No em or en dashes in any shaman file.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix | also on r14 |
|---|---|---|---|---|---|---|---|
| SHAMAN-1 | CONFIRMED | consumable | elemental.md:96 | "The flask replaces both elixirs; without one, run Greater Arcane Elixir." | Classic has no flask/elixir exclusivity. The consumables table this same page renders (`web/src/lib/guides/gear-extras.ts:323` reads `sim/request/raid-consumes.json`, shaman-elemental from_level 50) lists `flask_of_supreme_power` AND `greater_arcane_elixir` together; `bis/shaman-elemental.json` presets.raid.consumes the same. `warlock/destruction.md:108` says the opposite ("Flask and elixirs stack as they always have in Classic"). There is also no second caster elixir for a Nature/Fire shaman, so "both elixirs" names nothing. | "Flask of Supreme Power and Greater Arcane Elixir go on together; a flask stacks with your elixirs." | not fixed (elemental.md:95 same text) |
| SHAMAN-2 | LIKELY (NEEDS-OWNER) | rotation | restoration.md:40, restoration.md:83 (and the Chain Heal note in `data/curated/apl/shaman-restoration.json` priority 16) | "Chain Heal cast on a party member only jumps inside your party, so your group placement matters" / "Chain Heal cast on a party member only jumps to other party members" | `design/reviews/2026-11-02-sim-audit-round9/healers.md:21`: Chain Heal is ImplicitTarget 45 (chain-heal ally), 3 targets, 0.5 falloff; cmangos-classic limits jumps to the caster's group when the target is in it, which in a raid is the whole raid. No client row or ruling says party-only. | Weakness line: "Chain Heal jumps to whoever is most hurt near its target, so stand where the damage is." Line 83: "Healing Stream, Mana Spring and Mana Tide only reach your party, so ask to be grouped with whoever takes the damage, usually the tank and the melee." Drop the Chain Heal sentence there and from the APL note. | not fixed (r14 restoration.md:40, :83) |
| SHAMAN-3 | CONFIRMED | gear | elemental.md:106 (also :92) | "Your Leatherworking pieces bind when equipped and can be bought, so your second profession is free" | The raid BiS waist is Skycaller's Leather Waistguard 252476 (tied Skycaller's Mail Belt 252589), Leatherworking; both `bonding: "pickup"` in `data/builds/1.60.1.70755/items/shaman.json`. | Line 106: "Your Leatherworking chest and legs bind when equipped and can be bought, but the Skycaller's belt binds on pickup: take Leatherworking for it, or wear Satyrmane Sash from Lord Vyletongue in Maraudon (`loot.json`) and keep the second profession for Alchemy or gathering." Line 92: "...crafts that bind when equipped, except the Skycaller's belt, which you have to make yourself." | partly (r14 :105 hedges "unless your list carries one that binds on pickup"; still does not name the belt) |
| SHAMAN-4 | CONFIRMED | stats | elemental.md:9, elemental.md:88 | `statPriority: [Intellect, Spell power, Hit, Nature power, Critical strike, Spell haste, Spell penetration]`; prose "Nature power ... still ranks just below plain spell power" | `bis/shaman-elemental.json` band 60 raid (the band the page reads, `class-dps.ts:79` + `presets.ts:36`): scale factors spell_power 2.00, hit 1.37, nature_power 1.29, intellect 1.00, crit 0.74; mp5/spirit/haste not significant. Hit sits between spell power and Nature power; Intellect is fourth, not first. Races prose (:100) "hit is your second stat" agrees with the data, the frontmatter does not. | Frontmatter: `[Spell power, Hit, Nature power, Intellect, Critical strike, Spell haste, Spell penetration]`. Prose: "...but Flame Shock and Searing Totem are fire, so a point of it is worth less than a point of plain spell power." (no rank claim). Re-check at the refresh: the file is carried from 70753. | not checked against r14 (r14 predates 70755) |
| SHAMAN-5 | CONFIRMED | stats | enhancement.md:9 | `statPriority: [Strength, Attack power, Hit, Critical strike, Agility, Intellect, Melee haste]` | `bis/shaman-enhancement.json` band 60 raid: hit 2.51, intellect 1.54, crit 1.34, agility 1.06, strength 1.00, attack_power 0.45, spell_power 0.40 (haste 9.49 per 1%). The prose at :91 already follows this order (hit, Intellect, crit, Agility, Strength); only the frontmatter is backwards. With 70755's Mental Dexterity at 150% (was 100% in the 70753 run) Intellect will only climb at the refresh. | `[Hit, Intellect, Critical strike, Agility, Strength, Attack power, Melee haste]` | n/a |
| SHAMAN-6 | LIKELY | stats | restoration.md:8 | `statPriority: [Healing power, Intellect, Spirit, MP5, Critical strike]` | `bis/shaman-restoration.json` band 60 raid: mp5 1.69, intellect 1.00, healing_power 0.81 (error 0.30, wide), spirit 0.60, crit 0.40. | `[MP5, Intellect, Healing power, Spirit, Critical strike]`; confirm at the refresh given the healing-power error. | n/a |
| SHAMAN-7 | NEEDS-OWNER | leveling | enhancement.md:113 | "Take Stormstrike as soon as the tree opens it." | 70755 moved Stormstrike to tier 3 (t3c2, 15 points, level 24 at the earliest), but spell 17364 carries `spell_level` 40 (`spellconst/shaman.json`), trainables lists it at level 40, and the engine gates it `RequiredLevel: 40` (`sim/shaman/stormstrike.go:41`); the leveling BiS bands take it only from band 40. If the client lets you press it at 24, the line stands; if not, it misleads every leveler for 16 levels. | If gated at 40: "Stormstrike sits in the fourth row but you cannot press it until 40, so put those points into Flurry and Elemental Weapons first and buy Stormstrike at 40." | not fixed |
| SHAMAN-8 | LIKELY | voice | enhancement.md:99 | "Keep Windfury Weapon up all fight, and use the weapon enchant and consumables listed here." | Page-meta ("listed here") and the paragraph names no consumable; the rendered table (raid-consumes.json shaman-enhancement 55) is Elixir of the Mongoose, Juju Power, R.O.I.D.S., Juju Might, Smoked Desert Dumpling, Major Frenzy Potion, Major Mana Potion. | "Windfury Weapon is your weapon's temporary enchant, so there is no room for a sharpening stone or weightstone. Elixir of the Mongoose, Juju Power and Juju Might are your raid buffs; carry a Major Mana Potion for the long fights." (owner may keep the potion line out if the table already says it) | not fixed |
| SHAMAN-9 | LIKELY | voice | restoration.md:97 | "...whichever the list ranks higher for your level." | Page-meta: refers to the BiS list rather than the game. | "Carry a healing staff, or a one-hander with a shield or a held off-hand item; at 60 the staff usually wins." (Glowing Brightwood Staff tops the band-60 file) | not fixed |
| SHAMAN-10 | LIKELY | voice | enhancement.md:41, :66, :95 | "No dual wield" is stated three times on one page (weaknesses, before the pull, best in slot), plus index.md:26. | Repetition reads generated; a top player says it once. | Keep it in Weaknesses and Best in slot; cut the before-the-pull sentence to "Put Windfury Weapon on your weapon." | n/a |
| SHAMAN-11 | NEEDS-SIM | rotation | elemental.md:80 | "Even with raid buffs, a three minute fight of nothing but rank 10 runs you dry before the boss dies." | The APL note (caster tier lane, 2026-10-31) says the raid band was "not mana-limited, out-of-mana well under 1 s of 180 s" before the rank 9/10 pacing line was added. The pacing advice itself is sound; only the absolute claim needs a number. | Question: at the level-60 raid preset, does a Lightning-Bolt-rank-10-only Elemental (same APL minus the rank 9 line) run out of mana before 180 s? If not, soften to "on a long fight, rank 10 alone runs you dry". | n/a |

## Sim findings (engine vs the game, for the next sim round)

- **S1 Stormstrike level gate.** `sim/shaman/stormstrike.go:41` `RequiredLevel: 40`, while 70755 puts Stormstrike in tier 3 (reachable at 24). If the owner rules it castable when talented (SHAMAN-7), the engine gate is wrong and every leveling band 24 to 39 under-values Enhancement.
- **S2 Carried BiS files (refresh list, not a guide fix).** `bis/shaman-*.json` are carried from 70753 (generated 2026-11-01): the band talent strings are in the 70753 order (Enhancement's Improved Ghost Wolf moved from index 5 to 12 in 70755), the weights predate 70755's Mental Dexterity 150%, Stormstrike 150%, Flurry 30% and Lightning Overload 30%, and `bis/shaman-enhancement.json` presets.raid lists `main_hand_imbue:elemental_sharpening_stone` where `sim/request/raid-consumes.json` now has `windfury_weapon`. All four stat findings above (SHAMAN-4/5/6) should be re-read after the refresh.
- **S3 Chain Heal target choice** stays as round 9 left it (any raid member, lowest health share); the guide, not the engine, is the outlier (SHAMAN-2).

## Owner questions

1. **Stormstrike at 24 or 40?** On 70755 the talent is in the fourth row but the spell still says level 40. Can a level 24 to 39 Enhancement shaman press it? (SHAMAN-7, S1)
2. **Chain Heal jumps:** party-only or raid-wide in Forever? Round 9 read cmangos as raid-wide; the Restoration page and APL note say party-only. (SHAMAN-2)
3. **Integrator note for r14:** int/r14's restoration.md:76 says "Cast [Healing Wave] at rank 2 in a raid; a low rank keeps all of your healing bonus". That rests on the open low-rank tooltip ruling (O2 in the refresh-hold memory); Classic's under-level-20 coefficient penalty would make rank 2 a large loss. Do not land that line until the tooltip check is in. r14's shaman pages also predate 70755 and must be rebased onto main's 70755 text, not merged over it.

## Ad-review claim (checked by the integrator)
- AD-1, REJECTED: "Flame Shock ticks for 21 seconds" (elemental.md:59, :72; APL note data/curated/apl/shaman-elemental.json:99). On 70755 every Flame Shock rank lasts 12 s (spellconst/shaman.json 8050 and 29228: duration_ms 12000, period_ms 3000), and Call of Flame reads "increases the duration of your Flame Shock spell by 3 sec" a point (talents/shaman.json, 104770), so the build's 3/3 gives 21 s. Both lines already say "with Call of Flame". The engine agrees (sim/shaman/flame_shock.go: 4 ticks + 1 per Call of Flame point, 3 s each). No change.

## Fix log

Branch gr1/shaman, commit 607ab7de7 (off main ff0f466b0, r14 landed).

- SHAMAN-1: fixed (elemental.md:96), flask and Greater Arcane Elixir go on together.
- SHAMAN-2: softened, not ruled (restoration.md:40, :83; data/curated/apl/shaman-restoration.json Chain Heal note). Party-only claim removed; weakness now "jumps go to whoever is most hurt near its target". Owner question 2 still open.
- SHAMAN-3: fixed (elemental.md:92, :106). Skycaller's belt (252476/252589) binds on pickup; Satyrmane Sash 17755 from Lord Vyletongue (loot.json) named as the alternative.
- SHAMAN-4: fixed (elemental.md:9, :88). statPriority [Spell power, Hit, Nature power, Intellect, Critical strike, Spell haste, Spell penetration] from bis/shaman-elemental.json band 60 raid; prose no longer ranks Nature power.
- SHAMAN-5: fixed by r14 (enhancement.md:9 already [Hit, Intellect, Critical strike, Agility, Strength, Attack power, Spell power, Melee haste], matches band 60 raid scale factors).
- SHAMAN-6: fixed (restoration.md:8), [MP5, Intellect, Healing power, Spirit, Critical strike] verified against bis/shaman-restoration.json band 60 raid (mp5 1.69, int 1.00, hp 0.81, spirit 0.60, crit 0.40); healing power error is wide, re-read at the refresh.
- SHAMAN-7: softened (enhancement.md:113): "opens in the fourth row, but the spell itself is a level 40 spell; if it will not cast for you yet, take it at 40." Owner question 1 still open.
- SHAMAN-8: fixed (enhancement.md:99), consumables named from raid-consumes.json shaman-enhancement 55 (Mongoose, Juju Power, Juju Might, R.O.I.D.S., Smoked Desert Dumpling); no Major Mana Potion, the table no longer carries one.
- SHAMAN-9: fixed (restoration.md:97), "at 60 the staff usually wins" (Glowing Brightwood Staff tops band 60).
- SHAMAN-10: fixed (enhancement.md:66), the before-the-pull dual wield sentence cut; Weaknesses and Best in slot keep it.
- SHAMAN-11: softened (elemental.md:80) to "On a long fight, rank 10 alone runs you dry." No number.
- r14 regression (counted confirmed): Healing Wave "rank 2 in a raid; a low rank keeps all of your healing bonus" (restoration.md:76) softened to "Drop a rank when your mana is behind the fight, and check what the low ranks actually heal for"; the two APL notes stating the full-bonus ruling as fact (shaman-restoration.json NS Healing Wave note and the below-60% Healing Wave note) softened the same way. Actions and ranks untouched.
- AD-1: rejected, left.

SIM findings / build proposals:
- S1 Stormstrike RequiredLevel 40 vs tier-3 talent (sim/shaman/stormstrike.go:41), pending owner ruling.
- S2 bis/shaman-*.json carried from 70753: re-read all three statPriority lists after the refresh; Enhancement presets.raid still lists elemental_sharpening_stone where raid-consumes.json has windfury_weapon.
- S3 Chain Heal target choice unchanged.
- Restoration APL still casts Healing Wave rank 2 (NS line and the below-60% line); once the low-rank tooltip ruling lands, either confirm or move those lines to a higher rank.
- No build-string changes proposed.

## Fix pass 2

MUSTs
1. Restoration NS Healing Wave: fixed (R3+R5). restoration.md:36 Strengths now "Nature's Swiftness makes your biggest Healing Wave instant"; :77 "Nature's Swiftness into your highest rank of Healing Wave ... never spend it on a low rank"; :76 hedge removed; :79 new "Low ranks" paragraph states the geared-60 tooltip check once, as an instruction. Cards: shaman-restoration.json:204 (NS line) "shown at rank 2, but for the save cast your highest rank: Nature's Swiftness makes any rank instant"; :935 (below 60%) "shown at rank 2, the cheapest rank, to save mana in a raid" with the hedge moved to the prose paragraph. Ranks and actions untouched.
2. Chain Heal ranks: fixed (restoration.md:83). spellconst/shaman.json: 1064 rank 1 260 mana / 251, 10623 rank 3 405 / 490, both sp_coefficient 0.714, spell_level 40 and 54 (neither under 20). Base per mana: rank 3 1.21 vs rank 1 0.97, so the reviewer is right on the base heal. Same coefficient means the healing bonus adds equally to both; crossover at about 250 healing, above which rank 1 is slightly better per mana and rank 3 still far bigger. Prose now says exactly that and keeps rank 3 as the default (the card casts rank 3).
3. Enhancement Grace of Air: fixed (R3). shaman-enhancement.json:87 note rewritten: "Grace of Air Totem whenever your Air slot is empty ... Take it when another shaman already has Windfury Totem down near the melee; if none does, Windfury Totem is the stronger drop". Prose (enhancement.md:79) unchanged, already right for raids. Learned 42 / rank 2 56 verified (trainables 8835, 10627).
4. Stormstrike hedge: fixed (enhancement.md:113). spellconst 17364 spell_level 40, trainables level 40, engine RequiredLevel 40; talent tier 3 col 2. Guide's leveling strip (leveling-builds.ts walk of the guide build, no ladder-first for Enhancement) reaches Stormstrike only at band 40 (21 points at 30: Flurry 4). Line: "Stormstrike sits in the fourth row of the tree but is a level 40 spell, so take it at 40 and spend the points before then on Flurry, Improved Ghost Wolf and Mental Quickness."
5. Restoration leveling: fixed (restoration.md:117). Rewritten to the strip's own plan (leveling-builds.ts + ladder-first shaman-restoration): Restoration from the first point, Mindfulness and Water Shield by 20, Mana Tide Totem at 30, Lesser Healing Wave and Healing Stream at 20, Chain Heal at 40, Riptide in by 50 (41 points), the ten Enhancement points from 51.

SHOULDs
- Elemental Warding / Devastation apology: fixed (elemental.md:66, R4). Warding 3/3 = 10% less Fire, Frost and Nature damage (talents 70755); with Devastation's point they are the toll to Elemental Fury (tiers 0 to 4 total 25 only with both).
- Enhancement hit-cap line: left. No spell-cap number added; the split is a rendered footnote (R2).
- Riptide "31st point": fixed with MUST 5 ("in by 50", matches the strip).
- Totem of Urgency padding: cut (restoration.md:113).
- "in Forever" trims: restoration.md:91, enhancement.md:79 (twice), enhancement.md:91.

New SIM findings
- Restoration APL: the Nature's Swiftness Healing Wave (priority 5) casts rank 2 (332, spell_level 6). A save after Nature's Swiftness should be the top rank (Healing Wave rank 10, 25357); rank 2 after NS is almost certainly wrong in real play. Move it to the highest known rank.
- Restoration APL: the below-60% Healing Wave (priority 17) also casts rank 2; the engine applies no low-rank penalty (wowsims-forever sim/shaman, no downrank code), so rank 2 is over-valued if the client tooltip shows a cut. Re-rank after the tooltip ruling; a mid rank (4 to 6) is the usual Classic mana saver.
- Enhancement APL: Air slot casts Grace of Air; in a raid with no other shaman holding Windfury Totem, Windfury Totem for the melee is worth more. Needs a raid-context condition or a preset switch.
