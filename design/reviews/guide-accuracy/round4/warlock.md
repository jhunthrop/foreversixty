# Guide accuracy round 4: Warlock

Auditor: warlock lane, round 4. Tree: land/guides-r4 752c02e43, client 1.60.1.70755, engine pin d1d9f2775.
Audited line by line: `web/src/content/guides/warlock/index.md`, `affliction.md`, `demonology.md`, `destruction.md`, the rendered card notes in `data/curated/apl/warlock-*.json`, every frontmatter source link, and the new `### Levels 1 to 19` subsections (notes: `round4/warlock-levels.md`).
Read first: `round1/warlock.md`, `round2/warlock.md`, `round3/warlock.md`, `round2/race-sims.md`, `round3/race-sims.md`, `ledger.md`. Nothing below re-files a fixed, sim-round or data-lane item unless the page now states it falsely. Since round 3 the only warlock change is the three Levels 1 to 19 blocks (752c02e43); round 3's fixes (c2cece8b7, f085c45f9) were re-checked and hold.

## Builds and shipped orders (decoded again on 70755)

All six frontmatter strings still decode to the `addon-builds.json` point counts (51 each, every gate and prerequisite holds). Orders for levels 10 to 19 (index 0 = level 10), against which the new subsections were checked:

- **Affliction:** Improved Life Tap 10-11, Suppression 12-16, Improved Corruption 17-21.
- **Demonology:** Improved Health Funnel 10-11, Demonic Embrace 12-14, Unholy Power 15-19.
- **Destruction:** Destructive Reach 10-11, Bane 12-16, Molten Skin 17-19.

Every talent order and rank effect in the three Levels 1 to 19 blocks matches (Improved Life Tap 20%, Suppression 5% / 20%, Improved Health Funnel 40% / 30% / no threat / any demon health at 2, Demonic Embrace 9% at 3, Unholy Power 10% at 5, Destructive Reach 20%, Bane 0.5 s at 5, Molten Skin 6% at 3). None of the blocks contradicts the Levels 20 to 60 text below it or the strip.

## Checked and correct (no finding)

- Trainer levels in the new blocks (`trainables/warlock.json`): Immolate 1, Corruption 4 (rank 1 a 2.0 s cast, spelltips 172), Life Tap 6, Bane of Agony 8, Summon Voidwalker 10, Health Funnel 12, Drain Life 14. Voidwalker by a race-specific class quest at 10 (Warcraft Tavern Forever leveling guide).
- Wands (`items.json`): Lesser Magic Wand 11287 required level 5, Greater Magic Wand 11288 required level 13, both Enchanting crafts (also Warcraft Tavern).
- Health Funnel "wakes up at 12": Health Funnel rank 1 is level 12; Improved Health Funnel's text needs nothing else.
- The rest of every page against the 70755 talent and spell text, again: Pandemic list, Malevolence, Nightfall, Soul Siphon "other Affliction effects", Wrack, Siphon Life 30 s, Decimation (Shadow Bolt or Searing Pain, 6%, 10 s, 40%, no shard), Master Summoner, Soul Link, Demonic Knowledge, Master Demonologist, Demonic Pact, Improved Sayaad 10% at 1, Aftermath 30% / 50%, Fire and Brimstone 25%, Shadow and Flame 20% / 20 s and its rank 5 guarantees, Shadowburn's 8 s shard rule, Bane of Havoc 15%, Incinerate +25%.
- Card notes (rendered): nothing false beyond the accepted Agony and Soul Fire actions.
- Links: all four warlock URLs answer 200. The Deep Dive recap says "Troll Warlock" and is titled as labelled. The wago and Forever Sixty links do not show what their labels promise (R4-3, R4-5).
- No em or en dashes, no site, sim, preset or band talk in the prose, and no quoted nightly numbers.

## Findings

| id | status | kind | file:line | wrong claim (quoted) | evidence | exact fix |
|---|---|---|---|---|---|---|
| WARLOCK-R4-1 | CONFIRMED | rotation | demonology.md:83 | "Doom's damage scales hard with your spell damage, which is why it beats Agony here." | `spellconst/warlock.json`: Bane of Doom 1742 base, coefficient 4.0 per 60 s. Bane of Agony r6: 12 ticks × 0.133 = 1.6 per 24 s, which is 4.0 per minute. Both scale the same per minute. Doom wins on base damage (1742 against 1380 a minute) and globals (one cast against two and a half). | "Doom does more damage per minute than Agony for one global instead of two and a half, which is why it beats Agony here." |
| WARLOCK-R4-2 | CONFIRMED | leveling | affliction.md:111 | "Improved Corruption starts at 17 and makes Corruption instant at 21. Until then Corruption has a 2 second cast" | Talent text: 0.4 s off Corruption's cast per point. So the cast is 1.6 s at 17, 1.2 s at 18, 0.8 s at 19 and 0.4 s at 20. Only before 17 is it 2 seconds. | "Improved Corruption starts at 17: each point takes 0.4 seconds off Corruption's 2 second cast, and at 21 it is instant. Until then, open with it while the mob is still running at you." |
| WARLOCK-R4-3 | CONFIRMED | source link | affliction.md:18-19, demonology.md:18-19, destruction.md:18-19, index.md:11-12 | label "Forever client spell and talent data", url `https://wago.tools/db2/SpellEffect` | The page loads with `currentVersion` 12.1.7.70668 (retail), not a Forever build. It is also one spell table, not the talent data. 1.60.1.70755 is in its build list, and `?build=1.60.1.70755` answers 200. | url `https://wago.tools/db2/SpellEffect?build=1.60.1.70755`, label 'Forever client spell data (wago.tools, build 70755)'. Alternatively, to cover talents too, add 'Forever Sixty, talent trees' at `https://foreversixty.gg/data/1.60.1.70755/talents/warlock.json` (200), the form the hunter guides use. |
| WARLOCK-R4-4 | CONFIRMED | source link | index.md:14 | "But Why Tho, interview with Anna Resendez and Clayton Stone" | The article spells her name "Ana Resendez" every time (Lead Software Engineer). It never mentions warlocks, but the index does not depend on it for a class claim. | label 'But Why Tho, interview with Ana Resendez and Clayton Stone'. The site /sources page has the same misspelling (site lane). |
| WARLOCK-R4-5 | CONFIRMED | source link | affliction.md:23-25, demonology.md:24-26, destruction.md:24-26 | label "Forever Sixty, rotation", url `https://foreversixty.gg/sources` | That URL is "How sources are classified" (pill kinds, panel recaps, interviews) and never uses the word rotation. No public page shows the rotation outside the guide itself. Round 2 set this label (384717fd4), so this is a round 2 edit that is wrong. | Delete the entry in all three guides, since the rotation is on the page itself. If the site kind must stay, relabel it 'Forever Sixty, how sources are classified' (same url). Mage and druid guides carry the same label (cross-class, site lane). |
| WARLOCK-R4-6 | LIKELY | rotation | destruction.md:97 | "on two targets, put Bane of Havoc on the one that must die and spend your casts on the other: 15% of that damage also lands on the Havoc target." | Havoc text: "15% of all damage done by the Warlock to other targets to also be dealt to the cursed target". As written, the target that must die gets only the 15% echo and the other one gets your full casts. The standard use is the reverse: hard-cast the priority target and Havoc the second one. | "On two targets, spend your casts on the one that must die and put Bane of Havoc on the other: 15% of everything you deal also lands on it. Havoc takes your bane slot, so no Doom on that target." |
| WARLOCK-R4-7 | LIKELY | leveling | destruction.md:124 | "Destructive Reach for 20% more range on Shadow Bolt and Immolate, so the mob takes an extra cast before it reaches you." | 30 to 36 yards (spelltips `range_yd` 30). That is 6 yards, which a running mob covers in under a second. Shadow Bolt is a 3 second cast at those levels, so this is not an extra cast. | "Your points at 10 and 11 go into Destructive Reach: 20% more range on Shadow Bolt and Immolate, so you start the pull from farther out and the first bolt lands before the mob is moving." |
| WARLOCK-R4-8 | LIKELY | talent | affliction.md:53 | "Suppression (5 points): ... Every warlock build starts here." | Every build takes Suppression 5, but Demonology takes it at 42-46 and Destruction at 47-51 (`addon-builds.json`), and their Levels 1 to 19 blocks open with Improved Health Funnel and Destructive Reach. "Starts here" reads as a point order the other specs do not follow. | "Every warlock build takes all five." |
| WARLOCK-R4-9 | LIKELY | voice | affliction.md:112, demonology.md:125, destruction.md:125 | The wand and gear bullet is word for word the same in all three guides. | This is the template-copy pattern round 1 fixed (WARLOCK-18). | Keep the facts in one spec and give the others a line of their own. Destruction: "Wands still do most of the killing before Bane cuts your casts; a Lesser Magic Wand at 5 and a Greater Magic Wand at 13 from an enchanter carry you to 20." Demonology: "Let the Voidwalker hold the mob and wand it down; a Lesser Magic Wand at 5 and a Greater Magic Wand at 13 do more than your bolts." Keep Affliction as written. |
| WARLOCK-R4-10 | NEEDS-SIM | consumable/race | affliction.md:8, :100; demonology.md:8, :113; destruction.md:8, :113 | Orc as the Horde pick; Demonology "Orc fits Demonology best of any Horde race"; Troll "a step behind Orc". | The Horde race was never measured for warlock: round 2 sims are Alliance only, and round 3 has no warlock rows. Undead beat Orc for every other Horde caster in round 3 (mage Arcane 623.3 against 619.3, Fire 676.0 against 673.3, Shadow priest 695.0 against Troll 687.8), and the owner rule is that race picks are measured. Warcraft Tavern names Orc for Demonology, so the prose is sourced, not contradicted. | Sims requested S1. If Undead leads beyond error, the pick becomes `[human, undead]` with an Undead-first Horde sentence; otherwise no change. |

## Decisions

| # | decision | reason and sources | guide change |
|---|---|---|---|
| D1 | The reason Doom beats Agony is base damage and globals, not scaling. | Client spellconst rows 603 and 11713 (same 4.0 spell power per minute). | WARLOCK-R4-1 |
| D2 | Bane of Havoc goes on the secondary target. | Havoc's 70755 text copies your damage on *other* targets to the Havoc target, so the target that matters should take the full casts. | WARLOCK-R4-6 |
| D3 | Frontmatter links name their build and their real content. | wago default is retail 12.1.7; /sources has no rotation; But Why Tho spells "Ana". | WARLOCK-R4-3, R4-4, R4-5 |
| D4 | The Levels 1 to 19 talent orders stay as shipped (Improved Life Tap first for Affliction). | Brief rule: prose follows `addon-builds.json`. Warcraft Tavern's Improved Corruption first would move no checkpoint question here, and there is no measured case to reorder. | none |

## Sims requested

- **S1 Horde race, all three specs:** raid preset, band 60, published gear and each guide's raid build, 10,000 iterations, seeds 1 and 50001: Orc vs Undead vs Troll for `warlock-affliction`, `warlock-demonology` and `warlock-destruction` (WARLOCK-R4-10).
- Carried from round 3, still open: Destruction weapon enchant (Brilliant Wizard Oil vs Major Firestone vs Major Spellstone), and Destruction Gnome vs Human once the Eureka! mask covers Incinerate.

## Sim findings

- None new. Still open from earlier rounds and not re-filed: Curse of the Elements before Corruption in all three APLs, the Affliction card casting Agony and prepull Amplify instead of Doom, the Demonology card's Soul Fire line, Wrack's aura mask (Siphon Life), the Wrack build on 70755, the Eureka! mask (Incinerate in, channel ticks out), and the carried bis and dps-per-point weights.

## Sources

- Warcraft Tavern, Warlock Leveling Guide (WoW Forever), https://www.warcrafttavern.com/forever/guides/warlock-leveling/ (read 2026-11-04): Voidwalker from a race-specific class quest at 10; trainer levels 1 to 20; Lesser (5) and Greater (13) Magic Wand from Enchanting; the beta build puts Improved Corruption first.
- But Why Tho, "BLIZZCON 2026: The 'World Of Warcraft: Forever' Developers Talk Balancing Old And New", https://butwhytho.net/2026/09/wow-classic-devs-world-of-warcraft-forever/ : interviewees Ana Resendez and Clayton Stone; no warlock content.
- Blizzard, World of Warcraft: Forever Deep Dive Panel Recap, https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap : "Orc Mage, Troll Warlock, and Undead Paladin".
- wago.tools, SpellEffect, https://wago.tools/db2/SpellEffect (default `currentVersion` 12.1.7.70668) and https://wago.tools/db2/SpellEffect?build=1.60.1.70755 (200).
- Forever Sixty, https://foreversixty.gg/sources ("How sources are classified"; no rotation content).
- Client data, 1.60.1.70755: `talents/warlock.json`, `trainables/warlock.json`, `spellconst/warlock.json` (603, 11713, 172), `spelltips/warlock.json` (172, 348, 689, 697, 755, 1120), `items.json` (11287, 11288), `addon-builds.json`, `races.json`; `round2/race-sims.md`, `round3/race-sims.md`.

## Count

CONFIRMED 5 (R4-1 to R4-5), LIKELY 4 (R4-6 to R4-9), NEEDS-SIM 1 (R4-10), NEEDS-OWNER 0. Decisions 4. Verdict: not zero. R4-3 to R4-5 are source links. The only confirmed play errors are one wrong reason (Doom) and one cast-time slip in the new Levels 1 to 19 text.

## Fix log (round 4)

- WARLOCK-R4-1 / D1 fixed: demonology.md:83, Doom beats Agony on damage per minute and globals, not scaling.
- WARLOCK-R4-2 fixed: affliction.md:111, Improved Corruption takes 0.4 s per point off the 2 s cast, instant at 21.
- WARLOCK-R4-3 / D3 fixed: index.md:11-12, affliction.md:17-18, demonology.md:18-19, destruction.md:18-19, url `https://wago.tools/db2/SpellEffect?build=1.60.1.70755`, label 'Forever client spell data (wago.tools, build 70755)'.
- WARLOCK-R4-4 / D3 fixed: index.md:14, "Ana Resendez".
- WARLOCK-R4-5 / D3 fixed per the round 4 brief (which overrides the audit's "delete" for damage specs): affliction.md:23-24, demonology.md:24-25, destruction.md:24-25, relabelled "Forever Sixty, how <Spec> plays" at https://foreversixty.gg/sim/specs. No guide links /sources now. "Forever Sixty, stat weights" stays at /sim/weights (all three are damage specs).
- WARLOCK-R4-6 / D2 fixed: destruction.md:97, hard-cast the priority target, Havoc the other (Havoc's 70755 text: damage to other targets echoes onto the cursed one).
- WARLOCK-R4-7 fixed: destruction.md:124, Destructive Reach now "open from farther out and get a little more casting done", no "extra cast" claim.
- WARLOCK-R4-8 fixed: affliction.md:53, "Every warlock build takes all five."
- WARLOCK-R4-9 fixed: demonology.md:125 and destruction.md:125 rewritten in their own words (same wand facts: Lesser 5, Greater 13, Enchanting); affliction.md:112 kept.
- WARLOCK-R4-10 resolved by round4/race-sims.md: Orc stays the Horde pick for all three; `recommendedRaces` stays [human, orc], FS1 race field stays human, no build change. Race prose now names Undead with its client text (Touch of the Grave: chance to drain up to 5% of maximum health, at most once every 15 seconds): affliction.md:100 "Undead is as good as Orc for Affliction" (582.5 against 582.1, a tie); demonology.md:113 and destruction.md:113 "Undead is close behind Orc". Troll stays "a step behind Orc" (last in all three).
- D4: no change, as decided.
- Card notes: no rendered note states anything false; none edited. Index has no Levels 1 to 19 block; the three spec blocks were re-read after the edits and still match the shipped orders.
- New SIM findings: none.
