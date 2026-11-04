# Guide accuracy round 16: Warlock

Auditor: warlock lane, round 16. Tree: gr16/audit at 2c3135a2e (land/guides-r15: fifteen guide rounds on launch-state main), client 1.60.1.70755, engine pin b2c4a2d05 (`sim/enginever/version.go`).
Audited line by line: `web/src/content/guides/warlock/index.md`, `affliction.md`, `demonology.md`, `destruction.md` (frontmatter, every body line, Levels 1 to 19 and Levels 20 to 60), the rendered warlock card clauses in `addon-data.json` `rotations`, `addon-builds.json`, and every source link.
Read first: `round1/warlock.md` to `round15/warlock.md` (with fix logs), `round4/warlock-levels.md`, the race-sims files, `ledger.md`, and sim round 17 `hunter.md` H2 (the pet-buff evidence round 15 relied on).

## What changed since round 15 (4600b7d82..2c3135a2e)

- Round 15's own fix (WARLOCK-R15-1): index.md:29 and demonology.md:64 rewritten (the demon and group buffs). Checked as new work below.
- `addon-data.json` changed only in rogue Backstab clauses (addon 2.0.25). Every warlock path is identical (diffed key by key).
- No change to `talents/warlock.json`, trainables, spelltips, `races.json`, `stat-weights.json`, `addon-builds.json`, `bis/warlock-*.json` or the warlock APLs.

## Builds (tier, column order; decoded again)

- Affliction leveling `25522300132001351/200522/005`: ILT 2, Suppression 5, IC 5, Grim Reach 2, Soul Harvest 2, Improved Drains 3, Amplify Curse 1, Pandemic 3, Malevolence 2, Siphon Life 1, Soul Siphon 3, Shadow Mastery 5, Wrack 1 (35); IHF 2, Unholy Power 5, Demonic Aegis 2, Improved Voidwalker 2 (11); Bane 5.
- Affliction raid `25520020135001351/00052/005004`: ILT 2, Suppression 5, IC 5, Grim Reach 2, Improved Bane of Agony 2, Amplify 1, Pandemic 3, Malevolence 5, Siphon Life 1, Soul Siphon 3, Shadow Mastery 5, Wrack 1 (35); Unholy Power 5, Aegis 2 (7); Bane 5, Aftermath 4 (9).
- Demonology leveling `25522/203511310122000135/005` and raid `05/2035103021220001351/0550005`: as round 14 recorded (raid drops ILT, IC, Grim Reach, Soul Harvest, Demonic Energies, Improved Voidwalker = 13; adds Improved Sayaad 2, Pact 1, ISB 5, Ruin 5 = 13).
- Destruction leveling `25522/0/2053235103101351` and raid `25500000033/0/0050355103101351`: raid drops Destructive Reach 2 and Molten Skin 3, adds Aftermath +2 and Cataclysm +1; Affliction drops Grim Reach 2 and Soul Harvest 2 for Pandemic 3 and Malevolence 3. destruction.md:50 holds.

Shipped orders (`addon-builds.json`, index 0 = level 10), re-derived:
- Affliction: ILT 10-11, Suppression 12-16, IC 17-21, Grim Reach 22-23, Soul Harvest 24-25, Improved Drains 26-28, Amplify 29, Siphon Life 30, Pandemic 31-33, Malevolence 34-35, Soul Siphon 36-38, Shadow Mastery 39 and 41-44, Wrack 40, then Demonology and Bane.
- Demonology: IHF 10-11, Demonic Embrace 12-14, Unholy Power 15-19, Aegis 20, IVW 21, Fel Vitality 22-24, Demonic Energies 25, Sacrifice 26, Master Summoner 27-28, Decimation 29 and 31, Soul Link 30, Demonic Knowledge 32-34, Master Demonologist 35-39, then ILT, Suppression, IC, Grim Reach, Soul Harvest, Bane.
- Destruction: Destructive Reach 10-11, Bane 12-16, Molten Skin 17-19, Cataclysm 20-21, Aftermath 22-24, Ruin 25-29, Shadowburn 30, Agonizing Flames 31-33, Conflagrate 34, Havoc 35, Fire and Brimstone 36-38, Shadow and Flame 39 and 41-44, Incinerate 40, then ILT, Suppression, IC, Grim Reach, Soul Harvest.

Every level named in the Levels 1 to 19 and Levels 20 to 60 paragraphs on all three pages matches these orders and the trainer rows (Immolate 1, Corruption 4, Life Tap 6, Bane of Agony 8, Voidwalker 10, Health Funnel 12, Drain Life 14, Succubus 20, Felhunter 30, Death Coil 42, Soul Fire 48). The two leveling blocks read in order and do not contradict each other.

## Round 15's fix, checked as new work

Client check (wago SpellMisc and SpellEffect 70755, local copy `scratchpad/r17/warlock/wago/*.layered.csv`), Attributes_3 0x100 ("can only target players"), every player rank:
- Flagged (skip the demon): Prayer of Fortitude 21562/21564, Gift of the Wild 21849/21850, Arcane Brilliance 23028, Prayer of Spirit 27681, Greater Blessing of Might 25782/25916, Kings 25898, Wisdom 25894/25918, Battle Shout every player rank (5242 to 25289), Strength of Earth (8076 to 25362), Grace of Air (8836, 10626, 25360), Windfury Totem (8515, 10609, 10612), Devotion Aura every rank (465 to 10293), the paladin resistance auras (19891/19899/19900 and the Frost and Shadow lines), Leader of the Pack 24932, Moonkin Aura 24907. index.md:29's list holds.
- Not flagged (reach the demon): Heart of the Lion 409583, Blood Pact (6307 to 11767), Fire/Frost/Nature Resistance Totem auras (8185/10534/10535, 8182/10476/10477, 10596/10598/10599), Mana Spring, Tranquil Air, **Retribution Aura every rank (7294, 10298, 10299, 10300, 10301)**, **Concentration Aura 19746**, Sanctity Aura 20218. Retribution and Concentration Aura are trainer spells for every paladin (`trainables/paladin.json`: Retribution 16 to 56, Concentration 22).
- Sim round 17 `hunter.md` H2: the beta timelines show Retribution Aura and Blood Pact landing on pets.
- Trueshot Aura in Forever is 1299346, ranged attack power only (aura 124), so it adds nothing to a demon. "The one exception for damage is Heart of the Lion" holds.

index.md:29 ("Blood Pact and a shaman's resistance totems still reach it too, but a paladin's Devotion and resistance auras do not") is true. demonology.md:64 says something wider, and that is false. See R16-1.

## Findings

### WARLOCK-R16-1 (CONFIRMED, kind: stale-70755 / round 15 edit)
- **Where:** demonology.md:64.
- **Quote:** "A shaman's resistance totems still reach it, so park your Succubus in a Fire Resistance Totem on Fire fights; a paladin's auras do not."
- **Evidence:** Retribution Aura (all five ranks, 7294 to 10301) and Concentration Aura 19746 have Attributes_3 0 in SpellMisc 70755. They are area party auras (effect 35, aura 15 and 149), the same shape as Blood Pact, and the beta logs show Retribution Aura on pets (sim round 17 hunter H2). Only Devotion Aura and the three resistance auras carry the 0x100 bit. So a paladin's Retribution and Concentration Aura do reach your demon. The page also contradicts index.md:29, which correctly limits the rule to "Devotion and resistance auras". This is round 15's own fix: round 15's evidence listed Retribution Aura 10301 and Concentration Aura 19746 as unflagged, and the demonology rewrite then dropped that distinction.
- **Fix** (demonology.md:64, replace that sentence): "A shaman's resistance totems still reach it, so park your Succubus in a Fire Resistance Totem on Fire fights. A paladin's Devotion and resistance auras do not, but Retribution Aura does."

No other statement on the four pages, and no rendered warlock card clause, is false on 70755. None contradicts its own page, the header build card, the live BiS rows, the race table or the stat table.

## Checked and correct (no finding), fresh pass

- **Talent text at the ranks taken** (`talents/warlock.json`): every rank effect the pages state holds. Suppression 5%/20%; IC 0.4 s a point, 10%; Amplify Curse 50% or 20%, 30 s; Pandemic +100% bonus on Corruption, both banes, the drains, Siphon Life and Wrack; Malevolence 2 = 2%, 3 = 3%, 5 = 5%; Soul Siphon 12% per effect, 36% cap; Shadow Mastery 10%; Wrack 6 s, every second, 10% on your other Shadow DoTs; Improved Bane of Agony 2 = 10%. IHF 2 (40%, 30% cheaper, no threat, any demon health); Demonic Embrace 3 = 9%; Unholy Power 5 = 10%; Aegis 1 = 15%; Improved Sayaad 2 = 20%; Fel Vitality 15%; Sacrifice (Imp Shadow 15%, Succubus Fire 15%); Master Summoner 4 s and 40%; Decimation (90%, 6%, 10 s, 40%, no shard); Soul Link 30% and 3%; Demonic Knowledge 100% of level; Master Demonologist (Succubus 10% Shadow, Voidwalker 10% physical); Pact. ISB 20% for 12 s, your attacks; Bane 0.5 s and 2 s; Destructive Reach 20%; Molten Skin 3 = 6%; Aftermath 30/40/50%; Ruin +100%; Agonizing Flames 10%; Fire and Brimstone 25%; Shadow and Flame 20%/20 s with keep and refund at 5; Bane of Havoc 15%; Incinerate +25% with Immolate.
- **Spell facts** (`spelltips`, `trainables`): unchanged since round 14, where each was verified. Cast times give Bane's 1.5 / 2.0 / 2.5 / 4 s and Decimation's 2.4 s on a 6 s cooldown. Unstable Affliction and Haunt are `class_spell` rows only, so index.md:31 holds.
- **Rotation sections and card clauses:** they match the APLs. The Demonology Soul Fire card line at bands 48 to 60 is the accepted card-action item (rounds 11 and 12).
- **Races:** `races.json` and the measured results are unchanged, so every race line holds as round 14 recorded. Each page's "Human and Orc are the picks" matches `recommendedRaces`.
- **Gear, professions, consumables, stats:** BiS, weights and race data unchanged. Every line still matches its table. The stat paragraphs explain and never rank.
- **Links** (curl, 2026-11-04, all 200): the Blizzard Deep Dive recap, wago.tools SpellEffect 70755, But Why Tho, foreversixty.gg/sim/weights. Each label matches its page, as round 15 recorded.
- **Voice:** no em or en dashes; no site, sim, APL, preset or band talk outside source URLs; no nightly numbers.

## SHOULDs (wording, not findings)

- index.md:29: for parity with demonology.md once R16-1 lands, it could read "but a paladin's Devotion and resistance auras do not; Retribution Aura does." It is true as written.
- index.md:29: "Heart of the Lion ... is the exception", followed by "Blood Pact and a shaman's resistance totems still reach it too". This is true, because the exception is among damage and stat buffs, but "the one damage buff that reaches it" would read cleaner.
- Carried: index.md:19 filler wording; destruction.md:128 "Death Coil joins it at 42"; demonology.md:62 one clause on Improved Sayaad 2 of 3; affliction.md:54 leveling comparison inside the key-talent bullet.

## Decisions

| # | decision | reason and sources | guide change |
|---|---|---|---|
| D1 | Round 15's group-buff rule stands, with Retribution and Concentration Aura on the "reaches the demon" side. | SpellMisc Attributes_3 70755 per rank (listed above); sim round 17 hunter H2 beta-log evidence. | R16-1 |
| D2 | Trueshot Aura does not break "the one exception for damage". | Forever Trueshot 1299346 is aura 124 (ranged attack power) only. | none |
| D3 | Leveling orders stay as shipped; every prose level matches. | `addon-builds.json`. | none |

## Sims requested

- None new. The round 15 race items (Affliction on the Wrack raid build with Heart of the Lion on the demon; Demonology and Destruction re-runs on b2c4a2d05) are routed to sim round 18 and are not re-filed.

## Sim findings

- None new. Still open, all carried: Eureka! on Incinerate (client mask against the tooltip); demon crit at 60, Soul Siphon's count, Wrack and Doom at 60 (all waiting for a launch log); Decimation's 0% client effect (known engine gap). The engine gives the demon no Retribution Aura, Blood Pact or resistance totem. None of these moves raid damage, so there is no sim finding.

## Sources

- Client data, 1.60.1.70755: `talents/warlock.json`, `trainables/warlock.json`, `trainables/paladin.json`, `talents/hunter.json` (Trueshot Aura 1299346), `addon-builds.json`, `addon-data.json`. Also wago SpellMisc, SpellEffect and SpellName at 70755 (`scratchpad/r17/warlock/wago/*.layered.csv`), for the spell ids listed above.
- Sim-accuracy round 17 `design/reviews/2026-11-04-sim-audit-round17/hunter.md` H2 (beta timelines: Retribution Aura and Blood Pact on pets).
- Blizzard, World of Warcraft: Forever Deep Dive Panel Recap, https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200, 2026-11-04).
- wago.tools, SpellEffect at build 1.60.1.70755, https://wago.tools/db2/SpellEffect?build=1.60.1.70755 (200, 2026-11-04).
- But Why Tho, "BLIZZCON 2026: The 'World Of Warcraft: Forever' Developers Talk Balancing Old And New", https://butwhytho.net/2026/09/wow-classic-devs-world-of-warcraft-forever/ (200, 2026-11-04).
- Forever Sixty, stat weights, https://foreversixty.gg/sim/weights (200, 2026-11-04).

## Count

CONFIRMED 1, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs 2 new plus 4 carried. Decisions 3. Sims requested 0 new (round 15's two are with sim round 18).

## Fix log (round 16)

- WARLOCK-R16-1: fixed, demonology.md:64: "A paladin's Devotion and resistance auras do not, but Retribution Aura does." Matches index.md:29.
- SHOULDs: left (not in this round's scope).
