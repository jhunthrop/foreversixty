# Guide accuracy round 6: Warlock

Auditor: warlock lane, round 6. Tree: gr6/audit at 367cbadb7 (land/guides-r5, all five rounds applied), client 1.60.1.70755, engine pin d1d9f2775 (`sim/enginever/version.go`).
Audited line by line: `web/src/content/guides/warlock/index.md`, `affliction.md`, `demonology.md`, `destruction.md` (frontmatter, every body line, Levels 1 to 19 and Levels 20 to 60), the rendered card notes in `data/curated/apl/warlock-*.json`, and every source link.
Read first: `round1/warlock.md` to `round5/warlock.md` (with every fix log), `round4/warlock-levels.md`, `round2/race-sims.md` to `round5/race-sims.md`, `ledger.md`. Nothing below re-files a fixed, sim-round, accepted or data-lane item. Round 5's fixes and its review pass (stat prose, BiS prose, traded-talent lists) were re-checked against the current text and data.

## Builds and shipped orders (decoded again on 70755)

All six frontmatter strings decode on `talents/warlock.json` (tier, column order) to exactly the `addon-builds.json` point counts, 51 each, every gate and prerequisite held.

- **Affliction leveling** `25522300130201351/200522/005`: Improved Life Tap 10-11, Suppression 12-16, Improved Corruption 17-21, Grim Reach 22-23, Soul Harvest 24-25, Improved Drains 26-28, Amplify Curse 29, Siphon Life 30, Pandemic 31-33, Nightfall 34-35, Soul Siphon 36-38, Shadow Mastery 39 and 41-44, Wrack 40, Improved Health Funnel 45-46, Unholy Power 47-51, Demonic Aegis 52-53, Improved Voidwalker 54-55, Bane 56-60.
- **Affliction raid** `2552001013500105/00052/0550004` (30 Affliction: Improved Life Tap 2, Suppression 5, Improved Corruption 5, Grim Reach 2, Improved Bane of Agony 1, Amplify Curse 1, Pandemic 3, Malevolence 5, Siphon Life 1, Shadow Mastery 5; Unholy Power 5, Demonic Aegis 2; Improved Shadow Bolt 5, Bane 5, Ruin 4).
- **Demonology leveling** `25522/203511310122000135/005`: Improved Health Funnel 10-11, Demonic Embrace 12-14, Unholy Power 15-19, Demonic Aegis 20, Improved Voidwalker 21, Fel Vitality 22-24, Demonic Energies 25, Demonic Sacrifice 26, Master Summoner 27-28, Decimation 29 and 31, Soul Link 30, Demonic Knowledge 32-34, Master Demonologist 35-39, Improved Life Tap 40-41, Suppression 42-46, Improved Corruption 47-51, Grim Reach 52-53, Soul Harvest 54-55, Bane 56-60.
- **Demonology raid** `05/2035113011220001351/0550005` (Pact at 40 on 30 Demonology points).
- **Destruction leveling** `25522/0/2053235103101351`: Destructive Reach 10-11, Bane 12-16, Molten Skin 17-19, Cataclysm 20-21, Aftermath 22-24, Ruin 25-29, Shadowburn 30, Agonizing Flames 31-33, Conflagrate 34, Bane of Havoc 35, Fire and Brimstone 36-38, Shadow and Flame 39 and 41-44, Incinerate 40, Improved Life Tap 45-46, Suppression 47-51, Improved Corruption 52-56, Grim Reach 57-58, Soul Harvest 59-60.
- **Destruction raid** `25500000033/0/0050355103101351` (Affliction 18 with Pandemic before Malevolence at 15 points; Destruction 33).

## Checked and correct (no finding)

- **Stat prose against the band 60 raid tables** (`bis/warlock-<spec>.json`, generated 2026-11-04T05:0x, rows ordered as `lib/bis/scale-view.ts` sorts them, not-significant rows last):
  - Affliction: hit 4.15, spell power 3.94, Shadow power 3.48, crit 2.67, intellect 1.00, spell penetration not significant. Prose: hit to the cap, spell power, Shadow power close behind, crit, intellect, Spirit only through Life Tap. Matches.
  - Demonology: hit 2.87, spell power 2.63, Shadow power 2.13, crit 1.67, intellect 1.00, Fire power 0.50; Spirit is flagged insignificant, so it renders last. Prose order matches.
  - Destruction: hit 4.09, spell power 3.31, Fire power 2.10, crit 1.99, Shadow power 1.21, intellect 1.00, Spirit no gain. Prose matches, reasons stable, no numbers.
- **BiS prose:** "The pre-raid set leans on Tailoring and Enchanting crafts" holds on all three raid tables (Tailoring head or body pieces, Enchanting trinket and wand). The Affliction professions line names no piece.
- **Leveling:** both blocks read in order on all three pages; every talent, level and rank effect in Levels 1 to 19 matches the shipped orders above and the trainer rows (Immolate 1, Corruption 4 at 2.0 s, Life Tap 6, Bane of Agony 8, Voidwalker 10, Health Funnel 12, Drain Life 14); the 20 to 60 text matches (Siphon Life 30, Soul Siphon by 38 with three effects, Wrack 40, Succubus 20, Felhunter 30, Soul Fire 48 after Decimation 29/31, Shadowburn talent 30, Shadow and Flame full at 44, Incinerate 40, then Improved Life Tap, Suppression, Improved Corruption). Wands: Lesser Magic Wand req 5, Greater Magic Wand req 13, Enchanting.
- **Talent text at the decoded ranks:** every number on the four pages again (Improved Corruption 0.4 s a point, Pandemic list, Malevolence 1% a point, Nightfall 4%, Soul Siphon 12%/36%, Shadow Mastery 10%, Wrack text, Improved Bane of Agony 5% at 1, Ruin 80% at 4 and 100% at 5, Demonic Embrace 9% at 3, Demonic Aegis 15% at 1, Improved Voidwalker 10% at 1, Improved Sayaad 10% at 1, Unholy Power 10% at 5, Fel Vitality 15%, Master Summoner 4 s / 40%, Demonic Knowledge 100% of level, Master Demonologist, Soul Link 30%/3%, Decimation 90%/6%/40%/10 s, Demonic Pact, Bane, Molten Skin 6% at 3, Destructive Reach 20%, Cataclysm, Aftermath 30%/50%, Agonizing Flames 10%, Fire and Brimstone 25%, Shadow and Flame 20%/20 s and its rank 5 guarantees, Bane of Havoc 15%, Incinerate +25%).
- **Spell text** (`spelltips/warlock.json`): Shadow Bolt r10 3.0 s, Immolate 2.0 s, Incinerate 2.5 s, Soul Fire 6 s / 60 s, so Bane's 2.5 / 1.5 / 2.0 / 4 s and Decimation's 2.4 s hold; Curse of the Elements r4 75 and 10%, 5 min; Curse of Recklessness r4 505; Doom 1742 after 1 min, 60 s cooldown; Agony 24 s build-up; Amplify Curse 3 min, 30 s; Shadowburn 8 s shard rule; Conflagrate 10 s; Death Coil heals 100%; Life Tap base plus Spirit; Hellfire burns the caster; Spellstone 2% haste plus Shadow, Firestone 2% crit plus Fire.
- **Round 5's traded-talent lists** (affliction.md:59, demonology.md:60, destruction.md:50) decode exactly against both builds.
- **Races:** picks and sentences match `round2/race-sims.md` (Human) and `round4/race-sims.md` (Orc; Undead a tie for Affliction, close behind for the others; Troll last). No paid race is named for warlock (no Skyborne warlock in `combos.json`).
- **Card notes:** nothing false beyond the accepted Agony / prepull Amplify (Affliction) and Soul Fire (Demonology) actions, except the "cheap" wording in R6-3.
- **Links** (fetched 2026-11-04): Blizzard Deep Dive recap (200, "Troll Warlock" among the six new pairs); But Why Tho (200, "Ana Resendez" and Clayton Stone); `https://foreversixty.gg/sim/weights` (200, the stat weights page); `https://wago.tools/db2/SpellEffect?build=1.60.1.70755` answered 403 to the automated fetch this time (unknown, not dead; round 5 read it at 200 with the build on the page).
- **Voice:** no em or en dashes, no site, sim, search, APL, preset or band talk, no quoted nightly numbers.

## Findings

| id | status | kind | file:line | wrong claim (quoted) | evidence | exact fix |
|---|---|---|---|---|---|---|
| WARLOCK-R6-1 | CONFIRMED | gear | demonology.md:114 | "Your best head, Spellpower Goggles Xtreme Plus, is an engineer's craft; if you cannot get a pair, the tailored Felcloth Hood is nearly as good, since its 30 Shadow power suits your Shadow Bolts." | (1) The page's own band 60 raid table (`bis/warlock-demonology.json` / `.md:167`) puts Dreamweave Circlet (Tailoring) in the head slot; neither the goggles nor Felcloth Hood is among its alternatives. Round 5's review pass dropped the same "Engineering head" claim from the BiS paragraph for this reason and missed this one. (2) Wowhead Forever tooltip for 15999: "Requires Engineering (270)". A warlock who follows this paragraph (Tailoring and Enchanting) cannot wear the goggles at all, so "if you cannot get a pair" is wrong advice; `items.json` has required_level 0 for 15999, consistent with a skill gate. | "Tailoring and Enchanting, same as any warlock: tailored cloth and enchanted caster pieces fill your pre-raid set, your head included. Engineering goggles need Engineering to wear, so they are only yours if you train it in place of Enchanting. Alchemy is the swap if your guild wants a flask from you every raid night." |
| WARLOCK-R6-2 | CONFIRMED | gear | destruction.md:114 | "Your best head, Spellpower Goggles Xtreme Plus, is an engineer's craft; if you cannot get a pair, the tailored Dreamweave Circlet is the next best." | Same Engineering (270) requirement (Wowhead Forever 15999): a Tailoring and Enchanting warlock cannot equip them, so the "get a pair" fallback is false. "Dreamweave Circlet is the next best" is not on the page's band 60 raid table (`bis/warlock-destruction.md`, head alternatives: two PvP pieces, then Red Mageweave Headband), and which head leads moves with the nightly re-rank (it already moved for Demonology). | "Tailoring and Enchanting. Tailored cloth carries your pre-raid set, Enchanting adds a few strong caster pieces, and you enchant your own gear. Engineering goggles need Engineering to wear, so they are only yours if you train it in place of Enchanting. Alchemy is the swap when your guild expects a flask every raid night, and it keeps your Fire Power elixirs cheap." |
| WARLOCK-R6-3 | LIKELY | rotation, voice | destruction.md:73; card notes warlock-destruction.json priorityList[0], warlock-affliction.json priorityList[2] | "Corruption whenever it has fallen off. It is instant and cheap." / "Instant with Improved Corruption, and cheap." | Corruption r7 (25311) costs 340 mana: more than Incinerate r3 (325, 293 after the raid build's Cataclysm 3), Conflagrate r6 (255), Bane of Doom (300). It is worth casting for its damage per global, not because it is cheap. | destruction.md:73 "Corruption whenever it has fallen off. It is instant and ticks for 18 seconds off one global." Destruction note: "Corruption whenever it has fallen off. Instant with Improved Corruption, and 18 seconds of damage from one global." Affliction note: "Corruption whenever it has fallen off. Instant with Improved Corruption, and never clipped." |
| WARLOCK-R6-4 | LIKELY | leveling | demonology.md:128 | "Use the Voidwalker when pulls hurt, since Master Demonologist then cuts the physical damage you take by 10%." | The paragraph covers 20 to 60; the shipped order takes Master Demonologist at 35-39, 2% a point (talent text), so there is no reduction before 35 and 10% only from 39. | "Use the Voidwalker when pulls hurt; from 35 Master Demonologist also cuts the physical damage you take while it is out, by 10% once it is full at 39." |

## Decisions

| # | decision | reason and sources | guide change |
|---|---|---|---|
| D1 | Professions prose names the goggles only for the stable rule (Engineering to wear), never as "your best head". | Round 6 brief: gear prose names a piece only for a reason that does not move nightly; Wowhead Forever 15999 "Requires Engineering (270)". | R6-1, R6-2 |
| D2 | Affliction keeps "close to it with Soul Siphon" (affliction.md:57). | Round 1 research: Warcraft Tavern's Forever Affliction guide runs Wrack as filler behind 3/3 Soul Siphon and calls the Shadow Bolt build "only slightly behind"; "close" is the conservative reading of that and no 70755 measurement contradicts it. | none |
| D3 | Levels 1 to 19 talent orders stay as shipped. | Brief rule (prose follows `addon-builds.json`); rounds 4 and 5 D3/D4. | none |

## Sims requested

None new. Still unrun from earlier rounds: Destruction weapon enchant (Brilliant Wizard Oil vs Major Firestone vs Major Spellstone, band 60 raid, published gear and raid build, Human), and Destruction Gnome vs Human once the Eureka! mask covers Incinerate.

## Sim findings

None new. Still open, not re-filed: Curse of the Elements before Corruption in all three APLs; the Affliction card casting Agony and a prepull Amplify instead of Doom; the Demonology card's Soul Fire line; Wrack's aura mask (Siphon Life); the Wrack build on 70755; the Eureka! mask; the Destruction band 60 raid BiS rows still ranked on the pre-round-2 raid build (`2053345...` with Molten Skin) and the Gnome BiS captions (data lane); Wrack hover text at 36 a second.

## Sources

- Wowhead Forever tooltip, Spellpower Goggles Xtreme Plus 15999, https://nether.wowhead.com/forever/tooltip/item/15999 (read 2026-11-04): "Requires Engineering (270)", 27 spell damage and healing, binds when equipped.
- Blizzard, World of Warcraft: Forever Deep Dive Panel Recap, https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (read 2026-11-04).
- But Why Tho, "BLIZZCON 2026: The 'World Of Warcraft: Forever' Developers Talk Balancing Old And New", https://butwhytho.net/2026/09/wow-classic-devs-world-of-warcraft-forever/ (read 2026-11-04).
- Forever Sixty, stat weights, https://foreversixty.gg/sim/weights (read 2026-11-04).
- wago.tools, SpellEffect at build 1.60.1.70755, https://wago.tools/db2/SpellEffect?build=1.60.1.70755 (403 to the automated fetch 2026-11-04; unknown, not dead).
- Warcraft Tavern, Affliction Warlock Guide (WoW Forever), https://www.warcrafttavern.com/forever/guides/affliction-warlock/ (as cited in round 1, for D2).
- Client data, 1.60.1.70755: `talents/warlock.json`, `spelltips/warlock.json` (25311, 1293813, 18932, 603, 17728, 17953), `items.json` and `items/warlock.json` (15999, 10041, 14111, 10033), `addon-builds.json`, `bis/warlock-*.json` and `.md`, `combos.json`; `web/src/lib/bis/scale-view.ts`, `weight-reading.ts`.

## Count

CONFIRMED 2 (R6-1, R6-2), LIKELY 2 (R6-3, R6-4), NEEDS-SIM 0, NEEDS-OWNER 0. Decisions 3. Verdict: not zero. Both confirmed errors are the same Professions sentence on two pages. It names a nightly-moving best head that the Demonology table no longer shows, and it offers as a fallback goggles that need Engineering 270 to wear, which the reader of a Tailoring and Enchanting paragraph cannot equip.

## Fix log (round 6)

- WARLOCK-R6-1 fixed, demonology.md:114. The goggles are named only with the rule that they need Engineering to wear, never as a best head; Felcloth Hood line dropped. "your head included" left out of the proposed text, since the head slot moves nightly.
- WARLOCK-R6-2 fixed, destruction.md:114. Same rule; the Dreamweave Circlet fallback dropped.
- WARLOCK-R6-3 fixed after checking spelltips: Corruption r7 (25311) costs 340 mana, more than Incinerate r3 (325) and Conflagrate r6 (255). destruction.md:73 now "It is instant and ticks for 18 seconds off one global."; warlock-destruction.json:53 note now "...Instant with Improved Corruption, and 18 seconds of damage from one global."; warlock-affliction.json:72 note drops "cheap" ("...Instant with Improved Corruption, and never clipped."). Card actions untouched.
- WARLOCK-R6-4 fixed, demonology.md:128. The talent text confirms Master Demonologist (23785) at 2% a point, and the shipped order takes it at 35 to 39, so the line now says from 35, 10% once full at 39.
- Stat priority rule (every spec page): affliction.md:85, demonology.md:98, destruction.md:98 rewritten without ranking words. Each stat keeps its lasting reason (hit matters up to the cap and nothing past it, Suppression's 5%, school coverage, Ruin, Improved Shadow Bolt, Pandemic, Life Tap) and points to the table for the order. Demonology now names Decimation's Soul Fires as Fire, so it no longer says Fire power reaches only Immolate.
- Same rule elsewhere: destruction.md:106 drops "a Spellstone's 2% haste is worth less to you"; the Firestone and Spellstone reasons stay (Ruin, school coverage). The weapon oil pick stays, because its reason (a stone costs a shard) does not move nightly.
- D1 applied (R6-1, R6-2). D2 and D3: no change, as decided.
- `updated:` set to 2026-11-04 on affliction.md, demonology.md and destruction.md. index.md is unchanged, so its date stays.
- Prose check `-t warlock`: passes.
- New SIM findings: none.
