# Guide accuracy round 13: Warlock

Auditor: warlock lane, round 13. Tree: gr13/audit at 604b78a80 (twelve guide rounds, sim round 15, guide-header-build), client 1.60.1.70755, engine pin 9f7f1f6e4 (`sim/enginever/version.go`).
Audited line by line: `web/src/content/guides/warlock/index.md`, `affliction.md`, `demonology.md`, `destruction.md` (frontmatter, every body line, Levels 1 to 19 and Levels 20 to 60), the curated APLs `data/curated/apl/warlock-*.json` (prepull, priority, notes), the rendered card clauses in `data/builds/1.60.1.70755/addon-data.json` `rotations`, and every source link.
Read first: `round1/warlock.md` to `round12/warlock.md` (fix logs and decisions), `round4/warlock-levels.md`, the race-sims files, `ledger.md` (through round 12).

## What changed since round 12 (a8b147a6f..604b78a80)

Nothing for warlock. The round 12 landing touched druid, hunter, mage, paladin and rogue guides, the druid-restoration and rogue-assassination APLs, and rogue Shank card lines in `addon-data.json`. No warlock guide, APL, `talents/warlock.json`, `addon-builds.json`, trainables, spelltips, spellconst or `bis/warlock-*.json` changed. The audit below is a fresh full pass anyway, not a diff check.

## Builds (tier, column order; all 51 points, every gate held)

- Affliction leveling `25522300130201351/200522/005`; raid `25520020135001351/00052/005004` (ILT 2, Suppression 5, IC 5, Grim Reach 2, Improved Bane of Agony 2, Amplify Curse 1, Pandemic 3, Malevolence 5, Siphon Life 1, Soul Siphon 3, Shadow Mastery 5, Wrack 1; Unholy Power 5, Demonic Aegis 2; Bane 5, Aftermath 4).
- Demonology leveling `25522/203511310122000135/005`; raid `05/2035113011220001351/0550005` (Demonic Energies out, Improved Sayaad 1 and Demonic Pact 1 in; ISB 5, Bane 5, Ruin 5).
- Destruction leveling `25522/0/2053235103101351` (35 Destruction); raid `25500000033/0/0050355103101351` (33 Destruction: Destructive Reach and Molten Skin out, Cataclysm 3, Aftermath 5; Pandemic 3, Malevolence 3).
- Prerequisites re-read from `talents/warlock.json` `prereq_talent_id`: Wrack needs Siphon Life, Soul Link needs Demonic Sacrifice, Demonic Pact needs Soul Link, Fire and Brimstone needs Conflagrate, Incinerate needs Bane of Havoc (destruction.md:63 "Bane of Havoc is the path to Incinerate" holds).

Shipped orders re-derived from `addon-builds.json` (index 0 = level 10, [tree,row,col] 1-based):
- Affliction leveling: ILT 10-11, Suppression 12-16, IC 17-21, Grim Reach 22-23, Soul Harvest 24-25, Improved Drains 26-28, Amplify 29, Siphon Life 30, Pandemic 31-33, Nightfall 34-35, Soul Siphon 36-38, Shadow Mastery 39/41-44, Wrack 40, then IHF, Unholy Power, Aegis, Improved Voidwalker, Bane 56-60.
- Demonology leveling: IHF 10-11, Demonic Embrace 12-14, Unholy Power 15-19, Aegis 20, IVW 21, Fel Vitality 22-24, Demonic Energies 25, Sacrifice 26, Master Summoner 27-28, Decimation 29/31, Soul Link 30, Demonic Knowledge 32-34, Master Demonologist 35-39, then ILT, Suppression, IC, Grim Reach, Soul Harvest, Bane.
- Destruction leveling: Destructive Reach 10-11, Bane 12-16, Molten Skin 17-19, Cataclysm 20-21, Aftermath 22-24, Ruin 25-29, Shadowburn 30, Agonizing Flames 31-33, Conflagrate 34, Havoc 35, Fire and Brimstone 36-38, Shadow and Flame 39/41-44, Incinerate 40, then ILT 45-46, Suppression 47-51, IC 52-56, Grim Reach, Soul Harvest.

Every Levels 1 to 19 and Levels 20 to 60 sentence on the three pages matches these orders (Affliction "10 and 11 ... ILT", "Suppression from 12 to 16", "IC starts at 17 ... at 21 it is instant", Siphon Life from 30, Wrack from 40, Soul Siphon's 36%; Demonology IHF 10-11, Demonic Embrace 12-14, Unholy Power 15-19, Soul Link from 30, Master Demonologist from 35 and full at 39; Destruction Destructive Reach 10-11, Bane 12-16, Molten Skin 17-19, Shadowburn from 30, Shadow and Flame full at 44 with Incinerate at 40, then ILT, Suppression, IC). No Levels 1 to 19 paragraph contradicts the Levels 20 to 60 text below it.

## Checked and correct (no finding)

- **Talent text** (`talents/warlock.json`, at the ranks taken): Suppression 5% hit / 20% threat; IC 2 s off and 10%; Amplify Curse (Weakness or Agony 50%, Exhaustion 20%, 30 s; `spellconst` 18288 gcd 0, 180 s); Pandemic +100% bonus on Corruption, both banes, drains, Siphon Life, Wrack; Nightfall 4% on Corruption, Drain Soul, Drain Life, Wrack; Shadow Trance 10 s (17941); Shadow Mastery 10%; Soul Siphon 12% per other effect, 36% cap; Wrack 86 a second and 10% for 6 s; Improved Bane of Agony 2 = 10%; Malevolence 5 = 5% (3 = 3%); IHF 2 (40%, 30%, no threat, any demon health); Demonic Embrace 3 = 9%; Aegis 1 = 15%; Improved Sayaad 1 = 10%; Fel Vitality 15%; Demonic Sacrifice Imp Shadow 15%, Succubus/Incubus Fire 15%; Master Summoner 4 s and 40%; Decimation 90% cooldown cut, 6% below 35%, 10 s of 40% faster shard-free Soul Fire; Soul Link 30% and 3%; Demonic Knowledge 100% of level; Master Demonologist Succubus Shadow 10%, Voidwalker physical 10%; Demonic Pact; ISB 20% for 12 s, your attacks only; Bane 0.5 s and 2 s; Molten Skin 3 = 6%; Aftermath 10% a point; Ruin +100% bonus on Destruction spells; Agonizing Flames 10%; Fire and Brimstone 25%; Shadow and Flame 20%/20% for 20 s, keep and refund at 5; Bane of Havoc 15%; Incinerate 25% with Immolate.
- **Skill lines** (`trainables/warlock.json`): Shadow Bolt, Immolate, Soul Fire, Searing Pain, Incinerate, Conflagrate, Shadowburn are Destruction, so Ruin, Destructive Reach and Agonizing Flames reach them as the pages say; Corruption, banes, curses, drains, Siphon Life and Wrack are Affliction (Grim Reach).
- **Spell facts** (`spelltips`, `spellconst`): Corruption r1 to r7 all 2.0 s casts (r7 438 over 18 s); Immolate 2.0 s at every rank; Shadow Bolt r10 3.0 s; Incinerate 2.5 s; Soul Fire 6 s and 60 s (4 s with Bane; 2.4 s and 6 s under Decimation); Conflagrate 10 s; Shadowburn 15 s and the 8 s shard rule; Bane of Agony 552 over 24 s (12 ticks at 0.133 each); Bane of Doom 1742 after 1 min, 60 s, coefficient 4.0, so Doom's 1742 + 4.0 SP a minute beats Agony's 1380 + 4.0 SP a minute for one global against two and a half (demonology.md:80 holds at any spell power); Curse of the Elements r4 misc mask 126 (Holy through Arcane), 75 and 10%, 5 min (index.md:24 holds); Curse of Recklessness r4 505; Life Tap 420 plus Spirit; Drain Soul shard chance on damage, always on a kill; Create Spellstone (2% haste, Shadow up to 21) and Create Firestone (2% crit, Fire up to 21) both weapon enchants; summons 10 s, Succubus 100% base mana; Portal of Summoning 60, one hour.
- **Trainer levels**: Immolate 1, Corruption 4, Life Tap 6, Bane of Agony 8, Voidwalker 10, Health Funnel 12, Drain Life 14, Succubus and Incubus 20, Felhunter and Siphon Life 30, Wrack and Incinerate 40, Death Coil 42, Soul Fire 48, Bane of Doom and Portal 60. Unstable Affliction and Haunt are `class_spell` rows only (index.md:31 holds).
- **Rotation sections against the raid builds** (the header card and lists follow `raidBuild:`): openers match the APLs (Affliction CoE, Corruption, Doom, Siphon Life, Immolate, Wrack; Demonology prepull Imp, Sacrifice, Succubus, Soul Link at -2 s, then CoE, Corruption, Doom, Immolate, Shadow Bolt; Destruction Immolate at -1.5 s, CoE, Corruption, Doom, Conflagrate, Shadowburn, Incinerate). Immolate at 3 s left with a 1.5 s Bane cast gives up only the last tick, as all three pages say. Movement lines hold for each raid build (Demonology raid has no IC, Corruption 2 s; Destruction raid IC 5).
- **Rendered card clauses** (all bands): every warlock clause re-read; each is true for the build it shows. The Demonology Soul Fire clause names its Demonic Pact gate (component item, round 12 D2).
- **Races** (`races.json`): Ruthless 1% crit; The Human Spirit 10% Spirit; Eureka! 10% and 10% for 8 s on 2 min; Expansive Mind 5%; Blood Fury 10% spell power for 15 s on 2 min; Command 5% to warlock pets; Berserking 10% for 15 s on 3 min; Touch of the Grave 5% of max health, once every 15 s. No race line contradicts a measured result or the client.
- **Links** (fetched 2026-11-04): Blizzard Deep Dive recap 200 (names the warlock changes); wago.tools SpellEffect 70755 200; But Why Tho 200 (names Ana Resendez and Clayton Stone, as the label says); foreversixty.gg/sim/weights 200. Source-link rule holds.
- **Voice**: no em or en dashes, no site, sim, APL, preset or band talk, no nightly numbers, no stat ranking in any Stat priority section.

## Findings

None. No statement on the four pages, and no rendered card clause, is false on 70755, contradicts its own page, or contradicts the header build card.

## SHOULDs (wording, not findings)

- index.md:19 (carried from rounds 11 and 12): "Corruption, a bane, Immolate and Shadow Bolt" reads truer as "Every spec plays out of the same core: Corruption, a bane and Immolate, with your talents deciding the filler: Wrack, Shadow Bolt or Incinerate." Neither Affliction nor Destruction casts Shadow Bolt at 60.
- destruction.md:128 "Drain Life and Death Coil are your only heals": Death Coil arrives at 42, so for most of 20 to 60 Drain Life is the only one. "Drain Life is your heal, and Death Coil joins it at 42" is exact.
- `warlock-affliction.json` top-level `notes`/`engineering_notes` still call Wrack the "#2 damage source behind Shadow Bolt" (sim lane, not page text; carried).

## Decisions

| # | decision | reason and sources | guide change |
|---|---|---|---|
| D1 | Warlock prose stays as is: zero confirmed errors, third round running. | Full fresh pass above against `talents`, `spelltips`, `spellconst`, `trainables`, `addon-builds.json`, APLs, `races.json` and the rendered cards. | none |
| D2 | Bane of Doom over Bane of Agony on a boss holds at every spell power. | `spellconst` 603 (1742, SP 4.0 per minute) against 11713 (46 a tick, 0.133 a tick, 12 ticks over 24 s = 1380 and 4.0 SP a minute). | none |
| D3 | Leveling talent orders and race lines stay as shipped. | Brief rules (shipped order; race picks on new builds unmeasured; race re-runs with the sim lane). | none |

## Sims requested

- **Affliction races on the Wrack raid build** (carried, already with the sim lane): `25520020135001351/00052/005004`, round 15 rotation, band 60 raid preset, 10,000 iterations, seeds 1 and 50001: Human vs Gnome; Orc vs Undead vs Troll. affliction.md:97 was measured on the Shadow Bolt build.
- **PROVISIONAL (BiS re-rank on the shipped raid builds, with the sim lane):** all three `bis/warlock-*.json` are still engine d1d9f2775, generated 2026-11-04 05:00 to 05:06 UTC. After the re-rank lands, check affliction.md:9 statPriority and :85 (haste on a channel filler), demonology.md:10 and destruction.md:10 against the band 60 raid table.

## Sim findings

- None new. Still open: Wrack vulnerability mask 0x402 (Corruption and Agony only) against the tooltip's "other Shadow damage over time effects"; Eureka! class masks against the tooltip; Wrack hover text 36 against the talent's 86 (data lane); Decimation's 0% client effect (known engine gap).

## Sources

- Client data, 1.60.1.70755: `talents/warlock.json`, `spelltips/warlock.json`, `spellconst/warlock.json` (603, 11713, 1311680, 18288, 11689, 11675, 17941, 25307, 1316697, 17924, 1293813), `trainables/warlock.json`, `races.json`, `addon-builds.json`, `addon-data.json` rotations, `bis/warlock-*.json`; `data/curated/apl/warlock-*.json`.
- Blizzard, World of Warcraft: Forever Deep Dive Panel Recap, https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200, 2026-11-04).
- wago.tools, SpellEffect at build 1.60.1.70755, https://wago.tools/db2/SpellEffect?build=1.60.1.70755 (200, 2026-11-04).
- But Why Tho, "BLIZZCON 2026: The 'World Of Warcraft: Forever' Developers Talk Balancing Old And New", https://butwhytho.net/2026/09/wow-classic-devs-world-of-warcraft-forever/ (200, 2026-11-04).
- Forever Sixty, stat weights, https://foreversixty.gg/sim/weights (200, 2026-11-04).
- Wowhead Forever tooltips read in earlier rounds: Portal of Summoning 437169, Create Firestone 17953, Shadowburn 18871 (Soul Shard reagents), Eureka! 1259821 (non-periodic damaging or healing abilities).

## Count

CONFIRMED 0, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs 3. Decisions 3. Sims requested 2 (one PROVISIONAL, both already with the sim lane). Verdict: zero confirmed errors for warlock, third round running.
