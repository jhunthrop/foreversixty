# Guide accuracy round 12: Warlock

Auditor: warlock lane, round 12. Tree: gr12/audit at a8b147a6f (eleven guide rounds, sim round 15, leveling-order, guide-header-build), client 1.60.1.70755, engine pin 9f7f1f6e4 (`sim/enginever/version.go`).
Audited line by line: `web/src/content/guides/warlock/index.md`, `affliction.md`, `demonology.md`, `destruction.md` (frontmatter, every body line, Levels 1 to 19 and Levels 20 to 60), the curated APL notes in `data/curated/apl/warlock-*.json`, the rendered card lines in `data/builds/1.60.1.70755/addon-data.json` `rotations` (bands 10 to 60), and every source link.
Read first: `round1/warlock.md` to `round11/warlock.md` (every fix log and decision), `round4/warlock-levels.md`, the race-sims files, `ledger.md`.

## What changed since round 11 (d54de48aa..a8b147a6f)

- No warlock guide file, curated APL, `talents/warlock.json`, `addon-builds.json`, trainables, spelltips or `bis/warlock-*.json` changed. The `addon-data.json` diff touches only rogue and shaman card lines.
- guide-header-build (3bd24d3a2) landed. For a guide with a `raidBuild:` (all three warlock specs), the page now:
  - shows the raid build in the header rail, titled "Raid build" (`GuideActionRail.astro` `isRaidBuild`), with the rail's four rotation lines taken from the Priority list (`rotationListsData.priority.slice(0, 4)`);
  - prints "For the raid build." under "Rotation and priority" and builds the Opener, Priority and Cooldowns lists from the raid build's talents (`guide-rotation.ts` `describedBuildTrees`), including talent cast-time cuts (`talentCastTimeCuts`, Bane's "Reduces the casting time of your Shadow Bolt, Immolate, and Incinerate spells by 0.5 sec");
  - gives a key-talent chip the raid build's rank, and both ranks ("Level 60 n/m", "Raid n/m") when the two builds differ (`key-talents.ts` `refFor`).

Every page's prose was re-read against that framing: everything under Rotation and priority, Stat priority and the Overview must now hold for the raid build, and every key-talent bullet must hold against the chips it sits beside.

## Builds (tier, column order; all 51 points, every gate held)

- Affliction leveling `25522300130201351/200522/005`: ILT 2, Suppression 5, IC 5, Grim Reach 2, Soul Harvest 2, Improved Drains 3, Amplify Curse 1, Pandemic 3, Nightfall 2, Siphon Life 1, Soul Siphon 3, Shadow Mastery 5, Wrack 1; IHF 2, Unholy Power 5, Demonic Aegis 2, Improved Voidwalker 2; Bane 5.
- Affliction raid `25520020135001351/00052/005004`: ILT 2, Suppression 5, IC 5, Grim Reach 2, Improved Bane of Agony 2, Amplify Curse 1, Pandemic 3, Malevolence 5, Siphon Life 1, Soul Siphon 3, Shadow Mastery 5, Wrack 1; Unholy Power 5, Demonic Aegis 2; Bane 5, Aftermath 4.
- Demonology leveling `25522/203511310122000135/005`: ILT 2, Suppression 5, IC 5, Grim Reach 2, Soul Harvest 2; IHF 2, Demonic Embrace 3, Unholy Power 5, Demonic Aegis 1, Improved Voidwalker 1, Fel Vitality 3, Demonic Energies 1, Demonic Sacrifice 1, Master Summoner 2, Decimation 2, Soul Link 1, Demonic Knowledge 3, Master Demonologist 5; Bane 5.
- Demonology raid `05/2035113011220001351/0550005`: Suppression 5; the same Demonology core minus Demonic Energies, plus Improved Sayaad 1 and Demonic Pact 1 (31); Improved Shadow Bolt 5, Bane 5, Ruin 5.
- Destruction leveling `25522/0/2053235103101351`: ILT 2, Suppression 5, IC 5, Grim Reach 2, Soul Harvest 2; Destructive Reach 2, Bane 5, Molten Skin 3, Cataclysm 2, Aftermath 3, Ruin 5, Shadowburn 1, Agonizing Flames 3, Conflagrate 1, Bane of Havoc 1, Fire and Brimstone 3, Shadow and Flame 5, Incinerate 1.
- Destruction raid `25500000033/0/0050355103101351`: ILT 2, Suppression 5, IC 5, Pandemic 3, Malevolence 3; Bane 5, Cataclysm 3, Aftermath 5, Ruin 5, Shadowburn 1, Agonizing Flames 3, Conflagrate 1, Bane of Havoc 1, Fire and Brimstone 3, Shadow and Flame 5, Incinerate 1.

Shipped leveling orders (`addon-builds.json`, index 0 = level 10) are unchanged from round 11 and still match every Levels 1 to 19 and 20 to 60 sentence.

## Checked and correct (no finding)

- **Rotation sections against the raid build** (now labelled "For the raid build."):
  - Affliction :63 Unholy Power (raid 5), :67 Improved Bane of Agony and Pandemic (raid 2 and 3), :69 Immolate started at about 3 s (raid Immolate 2.0 s base, `spelltips` 11668, minus Bane 5 = 1.5 s, lands about 1.5 s early and gives up the last of its five 3 s ticks, as the line says), :71 Wrack and Soul Siphon, :73 Improved Life Tap 20% (raid ILT 2). Nothing in the section is leveling-only.
  - Demonology :76 6 s summons (Master Summoner 2 in the raid build), :82 Immolate timing, :84 Soul Fire 4 s outside Decimation (6000 ms, `spelltips` 17924, minus Bane 2 s), 2.4 s and 6 s under Decimation 2 (40% cast, 90% off 60 s), :94 "Corruption needs you to stand still in this build" (raid has no Improved Corruption; Corruption 11672 is 2000 ms).
  - Destruction :67 "You have no Demonology talents" (raid Demonology 0), :69 Immolate 1.5 s before the pull (Bane 5), :73 Corruption 18 s (11672 "over 18 sec"), :86 Improved Life Tap 20% (raid ILT 2), :92 Corruption instant (raid IC 5).
- **Key-talent bullets against the new chips.** Every bullet whose ranks differ between the builds says so in its own words: Affliction Nightfall "(2 points) ... The raid build drops it"; Destruction Aftermath "(3 points, 5 in the raid build)" and Pandemic and Malevolence "(raid build, 6 points)". Every other bullet's count matches both builds (Suppression 5 in all six; Demonology's seven bullets identical in both builds).
- **Overview and header.** Each Overview, Strengths, Weaknesses and In a raid line holds for the raid build now shown in the rail (Affliction Wrack filler and four or five opening globals; Demonology Pact setup, Master Demonologist 10%, Decimation 6%; Destruction Shadow and Flame 20%/20% for 20 s on 10 s and 15 s cooldowns, Ruin +100% bonus). No prose refers to the build card as the leveling build.
- **Demonology Soul Fire card line under "For the raid build."** Its rendered clause is "Every 6 seconds while Decimation is up, without Demonic Pact": the APL gate `not auraIsKnown 425464` resolves to "Demonic Pact" through the talent table (`rotation-lists.ts` `talentSpells`, `withTalentGates`), so the card agrees with demonology.md:84 ("Skip it anyway"). See the component note below.
- **Talent text** re-read at the ranks taken (`talents/warlock.json`): Improved Health Funnel 2 (40%, 30% cheaper, no threat, any demon health), Demonic Embrace 3 (9%), Demonic Aegis 1 (15%), Improved Sayaad 1 (10% Lash of Pain), Unholy Power 5 (10% pet damage), Master Demonologist 5 (Voidwalker 10% physical), Molten Skin 3 (6%), Destructive Reach 2 (20%), Ruin 5 (+100% bonus), Fire and Brimstone 3 (25%), Shadow and Flame 5 (100% keep and refund), Pandemic 3 (+100% bonus), Soul Siphon 3 (12% per other effect, 36% cap), Wrack 86 a second for 6 s, Nightfall 2 (4%).
- **Spell facts** (`spelltips/warlock.json`): Shadow Bolt 3.0 s, Incinerate 2.5 s, Soul Fire 6 s and 60 s, Conflagrate 10 s, Shadowburn 15 s and 8 s shard rule, Bane of Doom 1742 after 1 min, Bane of Agony 24 s, Siphon Life 30 s, Amplify Curse 3 min and 30 s.
- **Races**: `races.json` notes confirm Berserking 10% for 15 s on 3 min, Touch of the Grave up to 5% of max health once every 15 s, Command 5% for warlock pets. No line contradicts a measured result.
- **Links** (fetched 2026-11-04): Blizzard Deep Dive recap 200, wago.tools SpellEffect 70755 200, But Why Tho 200, foreversixty.gg/sim/weights 200; labels match; the source-link rule holds.
- **Voice**: no em or en dashes, no site, sim, APL, preset or band talk, no nightly numbers, no stat ranking in a Stat priority section.

## Findings

None. No statement on the four pages, and no rendered card clause, is false on 70755, contradicts its own page, or contradicts the header build card.

## SHOULDs (wording, not findings)

- index.md:19 (carried from round 11): "Corruption, a bane, Immolate and Shadow Bolt" reads truer as "Every spec plays out of the same core: Corruption, a bane and Immolate, with your talents deciding the filler: Wrack, Shadow Bolt or Incinerate."
- Component lane (not prose): the Demonology header rail is titled "Raid build" and its first of four rotation lines is Soul Fire "... without Demonic Pact", a line that build never casts. The clause is honest, but the rail would read cleaner if lines whose talent gate the described build fails were left out of the rail and the Priority list.
- `warlock-affliction.json` top-level `notes`/`engineering_notes` still call Wrack the "#2 damage source behind Shadow Bolt" (sim lane, not page text; carried).

## Decisions

| # | decision | reason and sources | guide change |
|---|---|---|---|
| D1 | The warlock prose already fits guide-header-build's framing (rail, rotation lists and chips follow the raid build). | Every Rotation, Stat priority and Overview line re-checked against the decoded raid builds above; bullets name both builds' ranks where they differ. | none |
| D2 | A gated card line that the described build never fires is a component item, not a guide finding, while its rendered gate words are true. | `apl-gates.ts` `gateWords`, `rotation-lists.ts` `withTalentGates`; brief (component-rendered parts). | none |
| D3 | Leveling talent orders and race lines stay as shipped. | Brief rules (shipped order; race picks on new builds unmeasured). | none |

## Sims requested

- **Affliction races on the Wrack raid build** (carried from rounds 10 and 11): `25520020135001351/00052/005004`, round 15 rotation, band 60 raid preset, 10,000 iterations, seeds 1 and 50001: Human vs Gnome; Orc vs Undead vs Troll. affliction.md:97 ("Gnome is a step behind", "Troll ... a step behind Orc", "Undead is as good as Orc") was measured on the Shadow Bolt build.
- **PROVISIONAL (BiS re-rank):** all three `bis/warlock-*.json` are still engine d1d9f2775, generated 2026-11-04 05:00 to 05:06 UTC; the 13:00 UTC re-rank has not reached this tree. After it lands, check affliction.md:9 statPriority and :85 (haste on a channel filler) against the band 60 raid table, and demonology.md:10 and destruction.md:10 likewise.

## Sim findings

- None new. Still open: Wrack vulnerability mask 0x402 (Corruption and Agony only) against the tooltip's "other Shadow damage over time effects"; Eureka! class masks against the tooltip (Incinerate in neither, a periodic 10% on the drains and Wrack); Wrack hover text 36 against the talent's 86 (data lane).

## Sources

- Client data, 1.60.1.70755: `talents/warlock.json`, `spelltips/warlock.json` (11672, 17924, 11668, 1293813, 11661, 18871, 18932, 603, 11713, 18881, 1316697, 18288), `trainables/warlock.json`, `races.json`, `addon-builds.json`, `addon-data.json` rotations, `bis/warlock-*.json`; `data/curated/apl/warlock-*.json`.
- Site code at a8b147a6f: `web/src/pages/guides/[class]/[spec].astro`, `web/src/lib/guides/guide-rotation.ts`, `rotation-lists.ts`, `key-talents.ts`, `web/src/lib/bis/apl-gates.ts`, `web/src/components/guides/GuideActionRail.astro`, `web/src/lib/guides/copy.ts`.
- Blizzard, World of Warcraft: Forever Deep Dive Panel Recap, https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200, 2026-11-04).
- wago.tools, SpellEffect at build 1.60.1.70755, https://wago.tools/db2/SpellEffect?build=1.60.1.70755 (200, 2026-11-04).
- But Why Tho, "BLIZZCON 2026: The 'World Of Warcraft: Forever' Developers Talk Balancing Old And New", https://butwhytho.net/2026/09/wow-classic-devs-world-of-warcraft-forever/ (200, 2026-11-04).
- Forever Sixty, stat weights, https://foreversixty.gg/sim/weights (200, 2026-11-04).
- Wowhead Forever tooltip, Eureka! (warlock) 1259821, https://nether.wowhead.com/forever/tooltip/spell/1259821 (round 8 read): non-periodic damaging abilities.

## Count

CONFIRMED 0, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs 3 (one for the component lane). Decisions 3. Sims requested 2 (one PROVISIONAL). Verdict: zero confirmed errors for warlock, second round running.
