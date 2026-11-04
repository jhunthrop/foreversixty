# Guide accuracy round 14: Warlock

Auditor: warlock lane, round 14 (final launch state). Tree: gr14/audit at d7fc51fe7 (thirteen guide rounds, sim round 16), client 1.60.1.70755, engine pin f637ec32c (`sim/enginever/version.go`).
Audited line by line: `web/src/content/guides/warlock/index.md`, `affliction.md`, `demonology.md`, `destruction.md` (frontmatter, every body line, Levels 1 to 19 and Levels 20 to 60), the curated APLs `data/curated/apl/warlock-*.json`, the rendered card clauses in `data/builds/1.60.1.70755/addon-data.json` `rotations`, the live `bis/warlock-*.json` (both factions, every band 60 row), `races.json`, `combos.json`, `items/warlock.json`, and every source link.
Read first: `round1/warlock.md` to `round13/warlock.md`, `round4/warlock-levels.md`, the race-sims files (rounds 2 to 5), `ledger.md` (through round 13), and the sim-accuracy round 16 warlock report (`design/reviews/2026-11-04-sim-audit-round16/warlock.md`).

## What changed since round 13 (604b78a80..d7fc51fe7)

- **Demonology raid build** (sim round 16 D1, commits 0c9f72749 and cad97d206): `05/2035113011220001351/0550005` became `05/2035103021220001351/0550005`. The point in Improved Voidwalker moved to a second point in Improved Sayaad. demonology.md:7, :60 and :62 were rewritten to match. This is the only warlock prose change.
- `addon-builds.json`: the Demonology raid order follows the new build. The leveling orders did not change.
- `addon-data.json` warlock rotations: per-level rank rows were added (sim round 16 W1, Shadow Bolt ranks by level). No warlock card clause text was added or removed. I diffed every `condition` string between 604b78a80 and HEAD, and the sets are identical.
- No change to the warlock APLs, `talents/warlock.json`, trainables, spelltips, `races.json` or `bis/warlock-*.json`. The BiS files are still the de044125c picks (engine d1d9f2775, 2026-11-04 05:00 to 05:06 UTC), and they are what ships.

## Builds (tier, column order; all 51 points, every gate held)

- Affliction leveling `25522300130201351/200522/005`: ILT 2, Suppression 5, IC 5, Grim Reach 2, Soul Harvest 2, Improved Drains 3, Amplify Curse 1, Pandemic 3, Nightfall 2, Siphon Life 1, Soul Siphon 3, Shadow Mastery 5, Wrack 1 (35); IHF 2, Unholy Power 5, Demonic Aegis 2, Improved Voidwalker 2 (11); Bane 5.
- Affliction raid `25520020135001351/00052/005004`: ILT 2, Suppression 5, IC 5, Grim Reach 2, Improved Bane of Agony 2, Amplify Curse 1, Pandemic 3, Malevolence 5, Siphon Life 1, Soul Siphon 3, Shadow Mastery 5, Wrack 1 (35); Unholy Power 5, Demonic Aegis 2 (7); Bane 5, Aftermath 4 (9).
- Demonology leveling `25522/203511310122000135/005`: ILT 2, Suppression 5, IC 5, Grim Reach 2, Soul Harvest 2 (16); IHF 2, Demonic Embrace 3, Unholy Power 5, Demonic Aegis 1, Improved Voidwalker 1, Fel Vitality 3, Demonic Energies 1, Demonic Sacrifice 1, Master Summoner 2, Decimation 2, Soul Link 1, Demonic Knowledge 3, Master Demonologist 5 (30); Bane 5.
- **Demonology raid (new)** `05/2035103021220001351/0550005`: Suppression 5; IHF 2, Demonic Embrace 3, Unholy Power 5, Demonic Aegis 1, Fel Vitality 3, Improved Sayaad 2, Demonic Sacrifice 1, Master Summoner 2, Decimation 2, Soul Link 1, Demonic Knowledge 3, Master Demonologist 5, Demonic Pact 1 (31); Improved Shadow Bolt 5, Bane 5, Ruin 5 (15). The gates hold. Tier 3 needs 10 points below it and has 14. Decimation needs 15 and has 19. Soul Link needs 20 and has 21. Demonic Pact needs 30 and has 30. The prerequisites hold too: Soul Link needs Demonic Sacrifice, and Pact needs Soul Link.
- Destruction leveling `25522/0/2053235103101351`: ILT 2, Suppression 5, IC 5, Grim Reach 2, Soul Harvest 2 (16); Destructive Reach 2, Bane 5, Molten Skin 3, Cataclysm 2, Aftermath 3, Ruin 5, Shadowburn 1, Agonizing Flames 3, Conflagrate 1, Bane of Havoc 1, Fire and Brimstone 3, Shadow and Flame 5, Incinerate 1 (35).
- Destruction raid `25500000033/0/0050355103101351`: ILT 2, Suppression 5, IC 5, Pandemic 3, Malevolence 3 (18; Pandemic's tier needs 10 and has 12, Malevolence's needs 15 and has 15); Bane 5, Cataclysm 3, Aftermath 5, Ruin 5, Shadowburn 1, Agonizing Flames 3, Conflagrate 1, Bane of Havoc 1, Fire and Brimstone 3, Shadow and Flame 5, Incinerate 1 (33).

Shipped orders, re-derived from `addon-builds.json` (index 0 = level 10):
- Affliction leveling: ILT 10-11, Suppression 12-16, IC 17-21, Grim Reach 22-23, Soul Harvest 24-25, Improved Drains 26-28, Amplify 29, Siphon Life 30, Pandemic 31-33, Nightfall 34-35, Soul Siphon 36-38, Shadow Mastery 39 and 41-44, Wrack 40, IHF 45-46, Unholy Power 47-51, Aegis 52-53, Improved Voidwalker 54-55, Bane 56-60.
- Demonology leveling: IHF 10-11, Demonic Embrace 12-14, Unholy Power 15-19, Aegis 20, Improved Voidwalker 21, Fel Vitality 22-24, Demonic Energies 25, Sacrifice 26, Master Summoner 27-28, Decimation 29 and 31, Soul Link 30, Demonic Knowledge 32-34, Master Demonologist 35-39, then ILT, Suppression, IC, Grim Reach, Soul Harvest and Bane.
- Demonology raid (new): as leveling to 20, then Fel Vitality 21-23, Improved Sayaad 24-25, Sacrifice 26, Master Summoner 27-28, Decimation 29 and 31, Soul Link 30, Demonic Knowledge 32-34, Master Demonologist 35-39, Pact 40, Suppression 41-45, Improved Shadow Bolt 46-50, Bane 51-55, Ruin 56-60.
- Destruction leveling: Destructive Reach 10-11, Bane 12-16, Molten Skin 17-19, Cataclysm 20-21, Aftermath 22-24, Ruin 25-29, Shadowburn 30, Agonizing Flames 31-33, Conflagrate 34, Havoc 35, Fire and Brimstone 36-38, Shadow and Flame 39 and 41-44, Incinerate 40, ILT 45-46, Suppression 47-51, IC 52-56, Grim Reach 57-58, Soul Harvest 59-60.

Every sentence in Levels 1 to 19 and Levels 20 to 60 on all three pages matches these orders. On Affliction: ILT at 10-11, Suppression 12-16, IC from 17 and instant at 21, Siphon Life from 30, Nightfall procs, Soul Siphon's 36%. On Demonology: IHF 10-11 waking up at 12 with Health Funnel, Demonic Embrace 12-14 for 9%, Unholy Power 15-19, Soul Link from 30, Master Demonologist from 35 and full at 39, and Soul Fire at 48 with Decimation already taken at 29 and 31. On Destruction: Destructive Reach 10-11, Bane 12-16, Molten Skin 17-19 for 6%, Shadowburn from 30, Shadow and Flame full at 44 with Incinerate at 40, then ILT, Suppression and IC. Neither Levels 1 to 19 block contradicts the Levels 20 to 60 text below it.

## The round 16 Demonology rewrite (re-checked as a fresh edit)

- demonology.md:7 `raidBuild:` decodes to the build above. The tree totals are 5/31/15.
- :60 "To pay for it, the raid build gives up Improved Corruption, Improved Life Tap, Grim Reach, Soul Harvest, Demonic Energies and the leveling build's point in Improved Voidwalker." This is exact. The leveling build has ILT 2, IC 5, Grim Reach 2, Soul Harvest 2, Demonic Energies 1 and Improved Voidwalker 1, and the raid build has none of them. What the raid build adds (Pact, the second Improved Sayaad point, Improved Shadow Bolt 5, Ruin 5) is named in the same paragraph and the next.
- :62 "The two points in Improved Sayaad make Lash of Pain 20% stronger." `talents/warlock.json` Improved Sayaad gives 10%, 20% and 30% at ranks 1, 2 and 3, so two points are 20%. "Improved Voidwalker and Improved Imp stay empty": both are 0 in the raid build. "the Voidwalker never comes out in a raid and the Imp is sacrificed before the pull": this matches the setup list at :70-:74 and the APL prepull. The survival clauses also hold at the ranks taken: IHF 2 (40%, 30% cheaper), Demonic Embrace 3 (9%), Demonic Aegis 1 (15%).
- :50 "The raid build goes to Pact and spends its Destruction points on Shadow Bolt": Improved Shadow Bolt, Bane and Ruin. This still holds.
- The key-talent bullets (:52 to :58) give the same ranks in both builds: Sacrifice 1, Master Demonologist 5, Soul Link 1, Demonic Knowledge 3, Decimation 2, Master Summoner 2, Fel Vitality 3. So no bullet contradicts a raid chip.
- No other warlock line names Improved Voidwalker or Improved Sayaad in a raid context. Affliction :59 ("along with Improved Health Funnel and Improved Voidwalker") is about the Affliction builds and is still true there.

## Round 14 focus: gear, professions, races and stats against the live tables

**Best in slot and Professions prose against `bis/warlock-*.json` (band 60, raid and bare, both factions).** Binding is read from `items/warlock.json` `bonding`.
- "The pre-raid set leans on Tailoring and Enchanting crafts" (all three pages) holds. In the raid tables, Tailoring fills 6 slots for Affliction and Demonology and 4 for Destruction, on both factions. Enchanting fills 2 on every table. Robe of the Void (Tailoring) is bind on pickup, and so are Frozen Heart of the Mountain and Brilliant Wand (both Enchanting). The rest of the tailored picks (Felcloth Hood, Mantle of the Timbermaw, Fel Cape, Gloves of Spell Mastery, Earthen Silk Slippers, Dreamweave Circlet) bind on equip. So the advice that the crafter keeps the pieces, especially the bind-on-pickup ones, is right.
- affliction.md:101 "a couple of enchanted caster pieces": the raid tables have exactly two (Frozen Heart of the Mountain, plus Greater Eternal Wand on the Alliance side or Brilliant Wand on the Horde side).
- destruction.md:114 and demonology.md:114 "Engineering goggles need Engineering to wear, so they are only yours if you train it in place of Enchanting". Spellpower Goggles Xtreme Plus is the Destruction raid head on both factions and a bare band 60 pick for all three specs. The line explains why the recommended pair does not wear them. It neither contradicts the table nor ranks gear.
- No warlock page names a specific item in the Best in slot section. The only items named anywhere are the leveling wands: Lesser Magic Wand (11287, required level 5) and Greater Magic Wand (11288, required level 13), as the pages say.
- Consumables prose (Elixir of Shadow Power, Elixir of Fire Power, Brilliant Wizard Oil, Major Mana Potion, Demonic Rune) matches the raid preset's consumables on all three specs. Spellstone (2% haste, Shadow up to 21) and Firestone (2% crit, Fire up to 21) match `spelltips` 17728 and 17953.

**Race prose against `races.json`, `combos.json`, the race table and the measured results**
- The race table tags come from `recommendedRaces: [human, orc]`, and each page's prose says "Human and Orc are the picks". These agree. The BiS captions (Gnome and Orc for Affliction and Destruction; Gnome and Troll for Demonology) differ, but the brief excludes those captions.
- `combos.json`: warlock races are Human, Orc, Undead, Gnome and Troll, with Troll `new_in_forever`. So index.md:35 holds ("Troll warlocks are new for the Horde ... Alliance warlocks are Human or Gnome. Skyborne cannot be warlocks").
- Each racial matches `races.json`: Ruthless 1% crit; The Human Spirit 10% Spirit; Eureka! 10% damage and 10% lower cost for 8 s on 2 min; Expansive Mind 5% mana; Blood Fury 10% spell power for 15 s on 2 min; Command 5% to warlock pets; Berserking 10% cast speed for 15 s on 3 min; Touch of the Grave drains up to 5% of max health, at most once every 15 s.
- Measured results: Affliction on the shipped Wrack raid build (sim round 16 A1, seeds 1 and 50001): Human ahead of Gnome by 0.2%, Orc ahead of Undead by 0.4%, Troll 1.5% behind Orc. affliction.md:97 ("Gnome is a step behind", "Troll ... a step behind Orc", "Undead is as good as Orc") holds on the build it describes. Demonology and Destruction still rest on round 4's measurement: Orc ahead of Undead beyond error, Troll last. Their "Undead is close behind Orc" and Troll "a step behind Orc" lines match it. Re-runs on the current builds are with the sim lane, and no line contradicts a measured result.

**Stat prose against the band 60 raid weights (the page tables)**
- Each Stat priority paragraph explains and never ranks. The reasons are stable and true on 70755. Suppression gives 5% hit. Shadow power covers the Shadow spells each page names. Fire power covers Immolate, Rain of Fire, Incinerate and Conflagrate. Crit pays through Pandemic, Ruin and Improved Shadow Bolt on the page that says so. Intellect gives about 1% crit per 60 at 60 (`playerexpectedstat.csv`, 0.0165% a point). Spirit only matters through Life Tap (420 + Spirit, `spelltips` 11689).
- Spirit is `insignificant: true` in all three raid tables (-0.019, 0.626 and -0.264, with errors of 0.18 to 0.27), and no `statPriority` lists Spirit. So "let it decide between two otherwise equal pieces" contradicts nothing shown.
- No prose names a stat as better than another, and no prose quotes a weight.

## Checked and correct (no finding), fresh pass

- **Talent text at the ranks taken** (`talents/warlock.json`): Suppression 5% and 20%; IC 2 s and 10%; Amplify Curse 50% or 20% for 30 s, with no global cooldown and a 3 min cooldown (`spelltips` 18288); Pandemic +100% bonus on Corruption, both banes, the drains, Siphon Life and Wrack; Nightfall 4%; Malevolence 5 = 5% and 3 = 3%; Improved Bane of Agony 2 = 10%; Soul Siphon 12% per effect, 36% cap; Shadow Mastery 10%; Wrack 6 s, 10%; IHF 2; Demonic Embrace 3 = 9%; Aegis 1 = 15%; Fel Vitality 15%; Improved Sayaad 2 = 20%; Demonic Sacrifice (Imp Shadow 15%, Succubus/Incubus Fire 15%, any summon cancels it); Master Summoner 4 s and 40%; Decimation (90%, 6%, 10 s, 40% faster, no shard, Shadow Bolt or Searing Pain); Soul Link 30% and 3%; Demonic Knowledge 100% of level; Master Demonologist (Succubus 10% Shadow, Voidwalker 10% physical); Pact (only resummoning the sacrificed pet cancels); Improved Shadow Bolt 20% for 12 s, your attacks only; Bane 0.5 s and 2 s; Destructive Reach 20%; Molten Skin 3 = 6%; Cataclysm; Aftermath 30% at 3, 40% at 4, 50% at 5; Ruin +100% bonus; Agonizing Flames 10%; Fire and Brimstone 25%; Shadow and Flame 20%/20% for 20 s, with keep and refund at 5; Bane of Havoc 15%; Incinerate 25% with Immolate.
- **Spell facts** (`spelltips`): Corruption 2.0 s base and 18 s; Immolate 2.0 s base, with a 15 s burn; Shadow Bolt 3.0 s; Incinerate 2.5 s; Soul Fire 6 s and 60 s; Conflagrate 10 s; Shadowburn 15 s and the 8 s shard rule; Bane of Agony 24 s, "dealt slowly at first, and builds up" (affliction.md:69, :79 hold); Bane of Doom 1742 after 1 min; Siphon Life 30 s; Curse of the Elements 75 and 10%, 5 min; Curse of Recklessness 505 armor; Life Tap 420 + Spirit; Health Funnel costs health (IHF "30% cheaper" holds).
- **Trainer levels** (`trainables/warlock.json`): Immolate 1, Corruption 4, Life Tap 6, Bane of Agony 8, Voidwalker 10, Health Funnel 12, Drain Life 14, Succubus and Incubus 20, Felhunter 30, Subjugate Demon 30 (index.md:29), Siphon Life talent 30, Wrack and Incinerate 40, Death Coil 42, Soul Fire 48, Bane of Doom and Portal of Summoning 60. Unstable Affliction and Haunt are `class_spell` rows only (index.md:31 holds).
- **Rotation sections against the raid builds** (the header card and the lists follow `raidBuild:`): the openers match the APL priority order. Immolate started at 3 s with a 1.5 s cast gives up only the last tick. The movement lines hold for each raid build (Demonology raid has no IC, Destruction raid has IC 5).
- **Rendered card clauses**: unchanged text, each true for its build (see "What changed").
- **Links** (fetched 2026-11-04 with curl): Blizzard Deep Dive recap 200; wago.tools SpellEffect 70755 200; But Why Tho 200; foreversixty.gg/sim/weights 200. The labels match their pages, and the source-link rule holds.
- **Voice**: no em or en dashes; no site, sim, APL, preset or band talk; no nightly numbers; no stat ranking.

## Findings

None. On the four pages, no statement and no rendered card clause is false on 70755. None contradicts its own page, the header build card, the live BiS rows (either faction), the race table or the stat table. The round 16 Demonology rewrite is correct.

## SHOULDs (wording, not findings)

- index.md:19 (carried from rounds 11 to 13): "Corruption, a bane, Immolate and Shadow Bolt" reads truer as "Every spec plays out of the same core: Corruption, a bane and Immolate, with your talents deciding the filler: Wrack, Shadow Bolt or Incinerate."
- destruction.md:128 (carried): "Drain Life and Death Coil are your only heals". Death Coil arrives at 42, so "Drain Life is your heal, and Death Coil joins it at 42" is exact.
- demonology.md:62 (round 16 wow-player SHOULD, optional): one clause on why Improved Sayaad stops at 2 of 3. For example: "The third point would come from Improved Health Funnel or Demonic Aegis, which the raid keeps for survival."
- `warlock-affliction.json` top-level `notes`/`engineering_notes` still call Wrack the "#2 damage source behind Shadow Bolt". This is for the sim lane and is not rendered (carried).

## Decisions

| # | decision | reason and sources | guide change |
|---|---|---|---|
| D1 | The round 16 Demonology raid-build rewrite (demonology.md:7, :60, :62) stands as written. | Decoded build; `talents/warlock.json` Improved Sayaad ranks (10/20/30%); `addon-builds.json` raid order; sim round 16 D1 measurement (+0.07% at both seeds, both factions). | none |
| D2 | The gear and professions prose agrees with the live BiS rows on both factions; the Tailoring and Enchanting pair is right because the bind-on-pickup crafts in the tables are Tailoring (Robe of the Void) and Enchanting (Frozen Heart of the Mountain, Brilliant Wand). | `bis/warlock-*.json` band 60 raid and bare; `items/warlock.json` `bonding`. | none |
| D3 | Human and Orc stay the stated picks. Affliction's race lines are now measured on the shipped Wrack build and hold. Demonology and Destruction rest on round 4 until the sim lane re-runs them. | Sim round 16 A1; round4/race-sims.md; brief (BiS race captions excluded). | none |
| D4 | The leveling talent orders stay as shipped; every prose level matches them. | `addon-builds.json`. | none |

## Sims requested

- **Demonology races on the new raid build** (with the sim lane under "race re-runs on the new builds"; sim round 16 lists Demonology Horde as an open owner question): `05/2035103021220001351/0550005`, band 60 raid preset, 10,000 iterations, seeds 1 and 50001. Compare Orc, Undead and Troll, then Human and Gnome. demonology.md:110 ("Orc fits Demonology best of any Horde race", "Undead is close behind Orc", Troll "a step behind Orc") rests on round 4's run on an older build. If Troll or Undead comes out ahead of Orc beyond error, the line changes.
- **Destruction races on the shipped raid build** (same lane item): `25500000033/0/0050355103101351`, the same preset and seeds. Compare Orc, Undead and Troll, then Human and Gnome. destruction.md:110 rests on round 4.
- **Closed:** Affliction races on the Wrack raid build, measured in sim round 16 A1. The page holds.

## Sim findings

- None new. Still open, all carried: Eureka!'s client mask against its tooltip (Wrack gets a periodic 10% in the engine, and the page follows the tooltip); Wrack hover text 36 against the talent's 86 (data lane); Decimation's 0% client effect (known engine gap); demon crit at 60, Soul Siphon's count and Doom at 60, all waiting for a launch log.

## Sources

- Client data, 1.60.1.70755: `talents/warlock.json`, `spelltips/warlock.json` (11713, 603, 18881, 1316697, 11672, 11668, 17924, 18871, 1293813, 17962, 11661, 1311680, 11717, 11708, 18288, 11695, 11689, 25307, 25311, 25309, 17953, 17728), `trainables/warlock.json`, `races.json`, `combos.json`, `items.json`, `items/warlock.json`, `addon-builds.json`, `addon-data.json` rotations, `bis/warlock-*.json`, `gametables/playerexpectedstat.csv`; `data/curated/apl/warlock-*.json`.
- Sim-accuracy round 16 warlock report, `design/reviews/2026-11-04-sim-audit-round16/warlock.md` (D1, A1, W1), commits 0c9f72749 and cad97d206.
- Blizzard, World of Warcraft: Forever Deep Dive Panel Recap, https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200, 2026-11-04).
- wago.tools, SpellEffect at build 1.60.1.70755, https://wago.tools/db2/SpellEffect?build=1.60.1.70755 (200, 2026-11-04).
- But Why Tho, "BLIZZCON 2026: The 'World Of Warcraft: Forever' Developers Talk Balancing Old And New", https://butwhytho.net/2026/09/wow-classic-devs-world-of-warcraft-forever/ (200, 2026-11-04).
- Forever Sixty, stat weights, https://foreversixty.gg/sim/weights (200, 2026-11-04).

## Count

CONFIRMED 0, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs 4. Decisions 4. Sims requested 2 (both already with the sim lane), 1 closed. Verdict: zero confirmed errors for warlock, fourth round running.
