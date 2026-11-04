# Guide accuracy round 11: Warlock

Auditor: warlock lane, round 11. Tree: gr11/audit at d54de48aa (ten guide rounds, sim round 15, leveling-order), client 1.60.1.70755, engine pin 9f7f1f6e4 (`sim/enginever/version.go`; engine source read from gr10/engine at 9f7f1f6e4).
Audited line by line: `web/src/content/guides/warlock/index.md`, `affliction.md`, `demonology.md`, `destruction.md` (frontmatter, every body line, Levels 1 to 19 and Levels 20 to 60), the curated APL notes in `data/curated/apl/warlock-*.json`, and the rendered card lines in `data/builds/1.60.1.70755/addon-data.json` `rotations` (bands 10 to 60) as `web/src/lib/guides/rotation-lists.ts` `clauseOf` cuts them (first sentence, colon or semicolon), plus every source link.
Read first: `round1/warlock.md` to `round10/warlock.md` (every fix log and decision), `round4/warlock-levels.md`, the race-sims files, `ledger.md`.

## What changed since round 10 (84460cec8..d54de48aa)

Only round 10's own fixes: affliction.md :27, :43, :48, :57, :59, :65, :71, :79, :81, :97; demonology.md :84, :128 ("costs no shard", "shard-free"); destruction.md :106 (Ruin wording); APL notes (Affliction CoE, Doom, Amplify, Agony, Wrack, Shadow Bolt; Demonology Soul Fire and Agony; Destruction Agony). No build string, `talents/warlock.json`, `addon-builds.json`, trainables, spellconst or `bis/warlock-*.json` changed. Every round 10 edit was re-checked below.

## Builds (tier, column order; all 51 points, every gate held)

- Affliction leveling `25522300130201351/200522/005`: ILT 2, Suppression 5, IC 5, Grim Reach 2, Soul Harvest 2, Improved Drains 3, Amplify Curse 1, Pandemic 3, Nightfall 2, Siphon Life 1, Soul Siphon 3, Shadow Mastery 5, Wrack 1; IHF 2, Unholy Power 5, Demonic Aegis 2, Improved Voidwalker 2; Bane 5.
- Affliction raid `25520020135001351/00052/005004`: ILT 2, Suppression 5, IC 5, Grim Reach 2, Improved Bane of Agony 2, Amplify Curse 1, Pandemic 3, Malevolence 5, Siphon Life 1, Soul Siphon 3, Shadow Mastery 5, Wrack 1 (35); Unholy Power 5, Demonic Aegis 2; Bane 5, Aftermath 4.
- Demonology leveling `25522/203511310122000135/005` (30 Demonology, no Pact) and raid `05/2035113011220001351/0550005` (Pact as the 31st point; ISB 5, Bane 5, Ruin 5).
- Destruction leveling `25522/0/2053235103101351` and raid `25500000033/0/0050355103101351` (Pandemic 3, Malevolence 3; Aftermath 5, Cataclysm 3).

Shipped orders (`addon-builds.json`, index 0 = level 10) re-derived: Affliction leveling ILT 10-11, Suppression 12-16, IC 17-21, Soul Harvest 24-25, Improved Drains 26-28, Siphon Life 30, Nightfall 34-35, Soul Siphon 36-38, Wrack 40; Demonology IHF 10-11, Demonic Embrace 12-14, Unholy Power 15-19, Sacrifice 26, Decimation 29/31, Soul Link 30, Demonic Knowledge 32-34, Master Demonologist 35-39; Destruction Destructive Reach 10-11, Bane 12-16, Molten Skin 17-19, Shadowburn 30, Conflagrate 34, Incinerate 40, Shadow and Flame 39/41-44, then ILT, Suppression, IC from 45.

## Checked and correct (no finding)

- **Round 10's Affliction rewrite.** :48 (35 Affliction points; Bane and Aftermath for a faster, harder Immolate) matches the decoded raid build. :57 Wrack vs Shadow Bolt re-derived from spellconst: Wrack 1316697 at 60 is 36 + 2.5 x 20 = 86 a second plus 0.286 spell power a tick, 200 mana for 6 s, times 1.36 under Soul Siphon; Shadow Bolt r10 25307 is 268 plus 0.857 spell power, 380 mana, 2.5 s with Bane 5. Per second of casting that is 117 + 0.389 SP against 107 + 0.343 SP, so Wrack is ahead at any spell power before its 10% vulnerability and Pandemic's crit bonus, for about a fifth of the mana per second; "it out-damages Shadow Bolt for a fraction of the mana" holds. :59 trade list, Improved Bane of Agony 2 = 10%, Aftermath 4 = 40% (max 5 = 50%), Malevolence 5% on Shadow spells incl. Wrack, all match `talents/warlock.json`. :65 opener order matches the APL (CoE, Corruption, Doom, Siphon Life, Immolate, Wrack). :71 Soul Siphon at 36% with Corruption, a bane and Siphon Life: the engine counts every Affliction-flagged DoT, Doom included (`sim/warlock/talents.go` soulSiphonMultiplier), and the talent caps at three others. :79, :81, :97 (Eureka! with an Immolate; DoTs and Wrack periodic) hold.
- **Card lines** (as rendered, `addon-data.json`): Affliction Doom "while a minute or more of the fight is left", Amplify "just before the Bane of Agony that covers the last minute" (Amplify gcd_ms 0, spellconst 18288), Agony "whenever no Bane of Doom is on the target" (holds at bands 10 to 60), Wrack "with your DoTs up" (shown from band 40, Wrack at 40 in the order), Shadow Bolt "until Wrack at 40 takes over". Demonology and Destruction Agony lines hold at every band. The Demonology Soul Fire line casting at bands 50 and 60 is the accepted card-action item; its rendered clause is true for the leveling build it shows.
- **Spell facts**: Curse of Recklessness r4 505 armor; Bane of Agony 24 s; Siphon Life 30 s; Curse of the Elements 5 min, 75 and 10%; Bane of Doom 1742 after 1 min (60 s cooldown); Conflagrate 10 s; Shadowburn 15 s and the 8 s shard rule; Soul Fire 6 s cast, 60 s cooldown (4 s with Bane, 2.4 s and 6 s under Decimation); Shadow Trance 10 s (17941); Immolate 2.0 s; Incinerate 2.5 s, 25% with Immolate.
- **Leveling**: every Levels 1 to 19 and 20 to 60 sentence on the three pages matches the orders above and `trainables/warlock.json` (Immolate 1, Corruption 4, Life Tap 6, Bane of Agony 8, Voidwalker 10, Health Funnel 12, Drain Life 14, Succubus 20, Felhunter 30, Soul Fire 48). Wands: Lesser Magic Wand 11287 requires 5, Greater Magic Wand 11288 requires 13 (`items.json`).
- **Races**: no Skyborne warlock (`combos.json`); racial text unchanged. No line contradicts a measured result.
- **Links** (fetched 2026-11-04): Blizzard Deep Dive recap 200, wago.tools SpellEffect 70755 200, But Why Tho 200, foreversixty.gg/sim/weights 200; labels match; source-link rule held.
- **Voice**: no em or en dashes, no site, sim, APL, preset or band talk, no nightly numbers, no stat ranking in a Stat priority section.

## Findings

None. No statement on the four pages, and no rendered card clause, is false on 70755 or contradicts its own page.

## SHOULDs (wording, not findings)

- index.md:19 "Every spec plays out of the same kit: Corruption, a bane, Immolate and Shadow Bolt, with your talents deciding what fills the gaps." Affliction's filler is now Wrack and Destruction's is Incinerate, so Shadow Bolt is the kit only for Demonology at 60. Reads truer as: "Every spec plays out of the same core: Corruption, a bane and Immolate, with your talents deciding the filler: Wrack, Shadow Bolt or Incinerate."
- warlock-affliction.json Agony note, second sentence (not rendered): "Amplify Curse makes it 50% stronger" applies only to the one Agony Amplify precedes; "and Amplify Curse makes the last-minute one 50% stronger" is exact.
- warlock-affliction.json top-level `notes`/`engineering_notes` still describe Wrack as the "#2 damage source behind Shadow Bolt" and the Wrack line as gone (carried from round 10, sim lane; not page text).

## Decisions

| # | decision | reason and sources | guide change |
|---|---|---|---|
| D1 | Round 10's Wrack-filler prose stands. | Re-derived from spellconst 1316697 and 25307 at level 60 (above), Soul Siphon cap and engine count, APL order. | none |
| D2 | Card notes are judged by the rendered clause. | `data/pipeline/addonrotation.py` `_one_line` and `rotation-lists.ts` `clauseOf` show only the first sentence or colon-led phrase; later sentences are SHOULDs at most. | none |
| D3 | Leveling talent orders and race lines stay as shipped. | Brief rules (shipped order; race picks on new builds unmeasured). | none |

## Sims requested

- **Affliction races on the Wrack raid build** (carried from round 10): `25520020135001351/00052/005004`, round 15 rotation, band 60 raid preset, 10,000 iterations, seeds 1 and 50001: Human vs Gnome; Orc vs Undead vs Troll. affliction.md:97 ("Gnome is a step behind", "Troll ... a step behind Orc", "Undead is as good as Orc") was measured on the Shadow Bolt build.
- **PROVISIONAL (BiS re-rank):** all three `bis/warlock-*.json` are still engine d1d9f2775, generated 05:00 to 05:06 UTC, before the Wrack build. After the final re-rank, check affliction.md:9 statPriority and :85 (haste on a channel filler) against the band 60 raid table.

## Sim findings

- None new. Still open: Wrack vulnerability mask 0x402 (Corruption and Agony only) against the tooltip's "other Shadow damage over time effects" (Siphon Life excluded in the engine); Eureka! class masks against the tooltip; Wrack hover text 36 against the talent's 86 (data lane).

## Sources

- Client data, 1.60.1.70755: `talents/warlock.json`, `spellconst/warlock.json` (1316697, 11661, 25307, 603, 18288, 17941), `spelltips/warlock.json` (11717, 18932, 18871, 17924, 11713, 18881, 1311680), `trainables/warlock.json`, `items.json` (11287, 11288), `addon-builds.json`, `addon-data.json` rotations, `bis/warlock-*.json`; `data/curated/apl/warlock-*.json`; `web/src/lib/guides/rotation-lists.ts`; `data/pipeline/addonrotation.py`; engine 9f7f1f6e4 `sim/warlock/talents.go` (Soul Siphon, Pandemic, Malevolence).
- Blizzard, World of Warcraft: Forever Deep Dive Panel Recap, https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200, 2026-11-04).
- wago.tools, SpellEffect at build 1.60.1.70755, https://wago.tools/db2/SpellEffect?build=1.60.1.70755 (200, 2026-11-04).
- But Why Tho, "BLIZZCON 2026: The 'World Of Warcraft: Forever' Developers Talk Balancing Old And New", https://butwhytho.net/2026/09/wow-classic-devs-world-of-warcraft-forever/ (200, 2026-11-04).
- Forever Sixty, stat weights, https://foreversixty.gg/sim/weights (200, 2026-11-04).
- Wowhead Forever tooltip, Eureka! (warlock) 1259821, https://nether.wowhead.com/forever/tooltip/spell/1259821 (round 8 read): non-periodic damaging abilities.

## Count

CONFIRMED 0, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs 3. Decisions 3. Sims requested 2 (one PROVISIONAL). Verdict: zero confirmed errors for warlock.
