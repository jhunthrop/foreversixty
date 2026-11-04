# Guide accuracy round 16: Paladin

Auditor: paladin (class index, Holy, Protection, Retribution). Audit tree gr16/audit (land/guides-r15 2c3135a2e: fifteen guide rounds on launch-state main). Client 1.60.1.70755, engine pin b2c4a2d05 (sim/enginever/version.go). Read-only: no guide edits, no sims run.

Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, the rendered card notes in `data/curated/apl/paladin-*.json`, `data/builds/1.60.1.70755/addon-builds.json`, `talents/paladin.json`, `trainables/paladin.json`, `spelltips/{paladin,shared}.json`, `talents/hunter.json` / `spelltips/hunter.json` (Heart of the Lion). Before starting I read the round 15 report and fix log, the round 14 report, the ledger through round 15, the race-sims files and the sim round 16 paladin report (P1, P3).

## What changed since round 15's audit (git diff 4600b7d82..2c3135a2e)

- **Guides:** only the two round 15 fixes, both in protection.md (:82 Reckoning, :123 Elixir of Holy Power).
- **Data:** only `addon-data.json` (addon 2.0.25). No paladin change to talents, trainables, spelltips, spellconst, addon-builds.json, rank-choices.json, APL card notes, races, combos or the live bis/paladin-*.json.

## Builds decoded (unchanged, all legal on 70755; tab order checked equal to tier/column order)

- **Retribution** `build` 050250002/502/0523330100130135 (14/7/30): Divine Strength 5, Reverence 2, Improved Seals 5, Holy Conduit 2; Divine Intellect 5, Precision 2; Benediction 5, Improved Judgement 2, Conviction 3, Sanctified Judgement 3, Vindication 3, Seal of Command 1, Sacred Arbiter 1, Two-Handed Weapon Specialization 3, Repentance 1, Vengeance 3, Champion of the Light 5. `raidBuild` 050050002/503/05233301001301351 (12/8/31): Reverence 0, Precision 3, plus Unsealed Wrath 1.
- **Protection** `build` 05005/15302102513501031/5003 (10/33/8): Divine Strength 5, Improved Seals 5; Divine Intellect 1, Redoubt 5, Precision 3, Sacred Duty 2, Templar's Bulwark 1, Shield Specialization 2, Toughness 5, Swift Judgement 1, Improved Righteous Fury 3, Anticipation 5, Holy Shield 1, Iron Creed 3, Shieldsworn 1; Deflection 5, Conviction 3.
- **Holy** `build` 50320215012101251/50322/0503 (31/12/8): Spiritual Focus 5, Healing Light 3, Reverence 2, Purifying Power 2, Divine Favor 1, Illumination 5, Divine Precision 1, Infusion of Light 2, Oath of Light 1, Holy Shock 1, Boundless Light 2, Holy Power 5, Light's Vigil 1; Divine Intellect 5, Precision 3, Guardian's Favor 2, Sacred Duty 2; Benediction 5, Conviction 3.

Leveling orders (addon-builds.json, index 0 = level 10), decoded by row and column:
- Retribution: Benediction 10-14, Improved Judgement 15-16, Conviction 17-19, Sanctified Judgement 20-22, Vindication 23-25, Seal of Command 26, Sacred Arbiter 27, Two-Handed Weapon Specialization 28-30, Repentance 31, Vengeance 32-34, Champion of the Light 35-39 (raid order: Unsealed Wrath 40). retribution.md:74 and :125 match.
- Protection: Divine Intellect 10, Redoubt 11-15, Precision 16-18, Sacred Duty 19, Templar's Bulwark 20, Sacred Duty 21, Shield Specialization 22-23, Toughness 24-28, Swift Judgement 29, Holy Shield 30, ..., Shieldsworn 40. protection.md:133 and :139 match.
- Holy: Spiritual Focus 10-14, Healing Light 15-17, Reverence 18-19, Illumination 20, Divine Favor 21, Illumination 22-24, Oath of Light 25 (16th Holy point, first one row 4 allows), ..., Holy Power 35-39, Light's Vigil 40. holy.md:117 and :123 match.

## Round 15 fixes re-checked

- **protection.md:123 (R15-1):** holds. Every spell it names carries a spell power coefficient on its Holy damage (Holy Shield 0.08, Holy Strike 0.429 on the "additional X as Holy damage" part, Seal of Fury proc 0.1, Judgement of Fury 0.429, Consecration ticks; sim round 16 P3 and the round 15 spellconst reading). Hammer of the Righteous (407632) is "3 times the damage per second of your main hand weapon" with a 0.0 coefficient, so "hits off your weapon, so the elixir does nothing for it" is true. The list leaves out Exorcism, but it no longer claims "every", so nothing is false.
- **protection.md:82 (R15-2):** holds on the measurement. Sim round 16 P1: Reckoning 5 in place of Anticipation 5 gave +2.5% raid threat (a little threat), and on bare 60 the chance of death went from about 27% to about 47-48%. Wording note under SHOULDs. Reckoning's client text ("40% chance after Blocking and a 100% chance after receiving a non-periodic critical strike to gain an extra attack") matches "extra attacks after a block or a crit". The Reckoning against Conviction plus Deflection sim is with sim round 18; not re-filed.

## Checked and correct (no finding)

- **Every talent bullet on all three pages** against the 70755 rank text at the page's rank: Benediction 10% on instants (and the Hammer of Wrath exception, round 13), Improved Judgement 2 s, Conviction 6%, Sanctified Judgement 60%, Sacred Arbiter 20% and refresh, Two-Handed Weapon Specialization 6% Physical and Holy and 20% less threat, Vengeance 3% x5 on non-periodic crits, Champion of the Light 60% of Intellect and 1.0 s off a 1.0 s Hammer of Wrath, Unsealed Wrath 15 s and 3 stacks, Vindication/Unyielding Faith/Seal of Command/Pursuit of Justice on one row, Repentance 6 s 1 min Humanoids; Redoubt 10% chance, 30%, 10 s or 5 blocks; Precision 3%; Sacred Duty 60 s and 10 s (30 s at rank 1); Templar's Bulwark 100% health 8 s; Shield Specialization 30% and 4% once every 3 s; Swift Judgement 1 min, own cooldown; Improved Righteous Fury 9%; Anticipation 20; Holy Shield 30%, 10 s, 4 charges, 20% threat, tier 4 (the 21st point); Iron Creed 30% and 6% for 6 s; Shieldsworn 4 s off a 10 s Holy Strike, uncharged triggers, spell block 6 s; Toughness 10%; Deflection 5%; Spiritual Focus 70%; Healing Light 12%; Reverence 30%; Illumination 50% of base cost; Divine Favor 2 min; Infusion of Light 1.0 s within 15 s; Oath of Light 5 min; Holy Shock 10 s; Boundless Light 30%; Holy Power +50% of the heal crit bonus (175%); Light's Vigil party heal through Holy Shock; Divine Precision 4%/10% at rank 1; Guardian's Favor 2 min and 6 s.
- **index.md:35** milestone rows: every tree has a 1-point talent at tiers 2, 3, 4 and 6 (Divine Favor, Oath of Light, Holy Shock, Light's Vigil; Templar's Bulwark, Swift Judgement, Holy Shield, Shieldsworn; Seal of Command, Sacred Arbiter, Repentance, Unsealed Wrath), which are the 11th, 16th, 21st and 31st points.
- **Trainer levels** (trainables/paladin.json) for every level the four pages name: Devotion Aura, Holy Light, Seal of Righteousness 1; Judgement, Blessing of Might 4; Holy Strike, Holy Light 2, Divine Protection 6; Hammer of Justice, Purify 8; Lay on Hands, Righteous Fury, Seal of Fury 10; Redemption 12; Holy Light 3, Blessing of Wisdom 14; Consecration, Exorcism, Flash of Light, Blessing of Kings, Voice of Truth 20; Concentration Aura 22; Holy Shield 30/40/50/60; Holy Shock ranks 30/40/48/56; Hammer of the Righteous 40; Divine Shield 34 and 50; Hammer of Wrath 44; Holy Wrath 50; Flash of Light rank 5 at 50; Light's Vigil rank 3 at 60. Sacred Shield exists only at level 80, so "No heal over time" (holy.md:40) holds.
- **Spell values (spelltips 70755):** Judgement and Holy Strike 10 s; Judgement 6% base mana, 10 yd; Hammer of the Righteous 6% base mana, 6 s, target plus 3; Hammer of Wrath 6 s, 30 yd, 1.0 s cast; Seal of Fury absorb with a shield and a 3 s taunt; Righteous Fury 30 min, 30% base mana; Seal of the Crusader 40% faster, its Judgement 40 s, refreshed by melee; Blessings and Greater Blessings 60 min; Blessing of Protection 5 min (3 min with Guardian's Favor); Divine Shield rank 2 12 s at half damage; Lay on Hands 20 min, max health, 550 mana at rank 3, drains your mana but not your regeneration; Light's Vigil rank 3 1,340 against ten Flash of Light rank 6 at 140; Holy Shock rank 1 160 against 275/325; Voice of Truth 6 s, 3 min; Concentration Aura 35%; Holy Strike rank 8 20 mana.
- **Racials (spelltips/shared.json):** Ruthless 1% crit, The Human Spirit 10% Spirit, Will to Survive (stuns), Stoneform 10% physical for 8 s on 3 min, Artisan's Precision 1% hit, Big Game Hunter 5% on Beasts, Will of the Forsaken (no immunity text), Cannibalize health and mana, Touch of the Grave drains health. No paid race named anywhere; index.md:39 (Human, Dwarf, Undead; no Skyborne) matches combos.json.
- **Card notes that render** (priorityList notes in all three APL files) agree with the prose and the client values above; the page priority lists match the APL order on every spec.
- **Gear, professions and stat prose:** unchanged since round 15 and still agree with the unchanged live BiS rows and tables. No prose ranks stats or quotes nightly numbers.
- **Source links:** all nine distinct URLs return 200 (curl -L, 2026-11-04) and their labels match.
- **Voice:** no em or en dashes, no site or sim talk.

## Findings

None. No CONFIRMED, LIKELY or NEEDS-OWNER finding this round.

## SHOULDs (wording, not findings)

- protection.md:82: "paying for it with Anticipation's Defense gets you killed on hard pulls". The measured death-chance rise is in fresh 60 gear (bare 60: about 27% to 48%); in raid gear the chance of death stayed at 0 with 5% more damage taken. Tighter: "...but paying for it with Anticipation's Defense makes you take more damage, and in fresh 60 gear it gets you killed far more often".
- Carried: holy.md:81 "Rank 1 costs half as much" (exact against rank 4, 58% of rank 3); "about half as much" covers both.
- Carried: retribution.md:133 "Rank 1 at 135 mana carries as much spell damage as any rank" is exact only to level 37.
- Carried: the Retribution Holy Strike card note's "Twenty mana" against Benediction's 18 (round 3 decision: card notes quote tooltip base cost).

## For the component and data lanes (not guide findings)

- Carried: trainables rows for Holy Light 9 (level 60), Blessing of Might 7 and Blessing of Wisdom 6; the Protection APL "Lay on Hands fires itself below 10% health" note (does not render); the Retribution raid trinket2 Frozen Heart of the Mountain.

## Decisions

- **Zero findings** decided on the 70755 client rows above; every page line was re-read, not only the round 15 diff.
- **R15-2 wording:** kept as a SHOULD, not a finding. "Hard pulls" is a fair player's word for a fight that out-damages your gear, which is exactly where the measurement shows the cost; nothing in it is false.
- **Not re-filed (routed to sim round 18 or settled earlier):** Reckoning against Conviction plus Deflection; Retribution Human against Dwarf in a raid; Seal of Righteousness rank 8's 0.0 coefficient (retribution.md:89, :103, :115 hang on it); Exorcism crit; Oath of Light and low ranks behind the owner's in-game check; Holy Light rank 9 source.
- **Races:** unchanged. Retribution Human (Undead), Holy Human (Undead), Protection Dwarf (Undead). No paid race named.

## Sims requested

None new. The round 15 requests (Reckoning against Conviction 3 plus Deflection 2; Retribution Human against Dwarf raid and bare; Beast target; Protection threat by ability; Crusader opener at bands 20 to 40) are with sim round 18.

## Sim findings

None new. Carried with sim round 18: Seal of Righteousness rank 8's spell power coefficient (client 0.0 on 20293, engine 0.1).

## Sources

- 70755 client data, data/builds/1.60.1.70755: talents/paladin.json, trainables/paladin.json, spelltips/paladin.json (all 144 paladin rows, incl. 407632, 20271, 1310994, 20308, 20293, 53601, 1311595, 10310), spelltips/shared.json (racials), addon-builds.json (paladin-retribution/protection/holy leveling and raid orders), combos.json, races.json.
- data/curated/apl/paladin-{holy,protection,retribution}.json (priorityList notes).
- design/reviews/2026-11-04-sim-audit-round16/paladin.md (P1 Reckoning against Anticipation; P3 Holy Shield, Seal of Fury and Judgement of Fury coefficients); design/reviews/guide-accuracy/round15/paladin.md and round14/paladin.md; ledger.md.
- Links fetched 2026-11-04, all 200: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap ; https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap ; https://www.wowhead.com/forever/news/undead-paladins-coming-to-world-of-warcraft-forever-382823 ; https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/ ; https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/ ; https://foreversixty.gg/planner?class=paladin ; https://foreversixty.gg/sim/weights ; https://foreversixty.gg/data/1.60.1.70755/races.json ; https://foreversixty.gg/data/1.60.1.70755/combos.json

## Counts

| status | count |
|---|---|
| CONFIRMED | 0 |
| LIKELY | 0 |
| NEEDS-SIM | 0 new (round 15 items with sim round 18) |
| NEEDS-OWNER | 0 |

Round 15 fixes: 0 errors. Levels 1 to 19: 0 errors. Race lines: 0 errors. Gear and professions lines: 0 errors. Source links: 0 errors. Card notes: 0 rendered errors. SHOULDs: 1 new, 3 carried.
