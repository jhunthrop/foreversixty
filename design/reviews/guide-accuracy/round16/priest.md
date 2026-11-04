# Guide accuracy round 16: Priest

Audited on the round 16 audit tree (land/guides-r15 2c3135a2e: fifteen guide rounds on launch-state main), client 1.60.1.70755, engine pin b2c4a2d05 (`sim/enginever/version.go`).
Files: `web/src/content/guides/priest/index.md`, `discipline.md`, `holy.md`, `shadow.md`; the card notes in `data/curated/apl/priest-*.json`; every frontmatter and body link; the live BiS files `data/builds/1.60.1.70755/bis/priest-{shadow,discipline,holy}.json` against `items/priest.json`. Read first: the round 1 to 15 priest reports and fix logs, round4/priest-levels.md, the race-sims files and the ledger.

## What moved since round 15

`git diff 4600b7d82 2c3135a2e` on the priest side touches one line: shadow.md:107, the round 15 fix of PRIEST-R15-1 (Spirit added to the best-in-slot rule). The APLs, `talents/priest.json`, `spelltips/priest.json`, `trainables/priest.json`, `races.json`, `items/priest.json`, `bis/priest-*` and the priest entries of `addon-builds.json` are unchanged.

## Decoded builds (re-derived, tab order from `talents/priest.json`)

- Shadow leveling `02510503103/0/52542001201300051` (20/0/31): Wand Specialization 2, Twin Disciplines 5, Silent Resolve 1, Mental Strength 5, Mental Agility 3, Inner Focus 1, Meditation 3; Shadow Focus 5, Blackout 2, Spirit Tap 5, Improved Mind Blast 4, Improved Shadow Word: Pain 2, Mind Flay 1, Improved Mind Flay 2, Vampiric Embrace 1, Shadow Weaving 3, Darkness 5, Shadowform 1.
- Shadow raid `00510503103/0/50052201211312251` (18/0/33): as leveling minus Wand Specialization; Shadow Focus 5, Improved Mind Blast 5, Improved Shadow Word: Pain 2, Shadow Reach 2, Mind Flay 1, Improved Mind Flay 2, Dark Inspiration 1, Vampiric Embrace 1, Shadow Weaving 3, Silence 1, Devouring Contagion 2, Early Demise 2, Darkness 5, Shadowform 1. The swap line at shadow.md:73 balances (9 points out, 9 in).
- Discipline `00500503133001143/3350500303/0` (29/22/0): Power Infusion swap at discipline.md:60 is legal (Renewed Hope 5th rank brings the tree to 30, Power Infusion's prerequisite Penance is taken).
- Holy `00500503103/33505003030121431/0` (17/34/0): "17 points in Discipline" and Spiritual Guidance 4 hold.

Leveling orders from `addon-builds.json` (index 0 = level 10): Shadow Wand Specialization 10-11, Spirit Tap 12-16, Shadow Focus 17-21, Improved Mind Blast 22-25, Improved Shadow Word: Pain 26-27, Mind Flay 28, Improved Mind Flay 29-30, Blackout 31-32, Vampiric Embrace 33, Shadow Weaving 34-36, Darkness 37-41, Shadowform 42. Discipline Twin Disciplines 10-14, Mental Strength 15-19, Penance 30. Holy Twilight Focus 10-12, Improved Renew 13-15, Holy Specialization 16-20. Every leveling sentence on the three spec pages matches.

## Findings

None. Counts: CONFIRMED 0, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## SHOULDs (wording, not findings)

| id | file:line | note |
|---|---|---|
| PRIEST-R16-S1 | shadow.md:107 | The round 15 rule "Judge it by that damage number, its Intellect and its Spirit" is true where it is used (every healing piece in the Shadow table that beats a caster piece wins on Intellect or Spirit: Glowing Thresher Cape at 30, Papal Fez, Inquisitor's Shawl, Glowing Eye of Mordresh at 40 and 50, Mooncloth Boots and Torch of Light at 60 bare Alliance; Uther's Strength at 60 raid Alliance beats no caster alternative). But the raid 60 table weights Spirit at 0 on both factions, and the stat paragraph at :103 does not explain Spirit at all, so a raid reader is told to judge by a stat his table ignores. Optional: "...its Intellect and, while you level, its Spirit..." Not false: Spirit is weighed in every bare band, which the table renders. |
| Carried | R15-S1 (discipline.md:64 "ticking"), R14-S1 (shadow.md:103 haste reason), R13-S1, R13-S2, R12-S1/S2, R11-S2, R8-S1/S2, R9-S1 | Unchanged and still optional. |

## Verified clean this round

- **The round 15 fix (shadow.md:107)** against every band and both factions of `bis/priest-shadow.json`: the "third of a healing item's bonus" holds for every healing item picked (Torch of Light 17 healing / 6 damage, Mooncloth Boots 24 / 8, Papal Fez 22 / 8, Inquisitor's Shawl 22 / 7, Glowing Eye of Mordresh 20 / 7, Glowing Thresher Cape 7 / 3 rounding up, Uther's Strength 18 / 6). Spirit carries weight in every bare band (0.17 to 0.93 of spell power) and none in raid, so "a healing piece with good Intellect or Spirit can beat a thin caster piece" is true where it happens. Torch of Light over Brilliant Wand at 60 bare Alliance is now explained.
- **Talent text against `talents/priest.json`**: Shadowform (10%, 50% mana, +100% crit bonus, so a Shadow crit at 200% as shadow.md:103 says), Shadow Weaving, Darkness, Improved Mind Flay, Improved Mind Blast (0.5 s per point), Improved Shadow Word: Pain (6 s), Shadow Focus, Spirit Tap, Blackout (2 points = 4%), Early Demise, Devouring Contagion, Silence, Shadow Reach, Dark Inspiration (1 point = 3%), Silent Resolve (1 point = 10%), Power in Light and Holy Precision (the skip reason), Renewed Hope (4 points = 8%, 4 s), Divine Aegis, Soul Warding (removes the shield's 4 s cooldown, 15% cost), Improved Power Word: Shield, Twilight Focus, Improved Renew, Holy Specialization, Divine Fury, Improved Healing, Litany of Light, Holy Binding, Spiritual Healing, Spiritual Guidance, Spirit of Redemption, Prayer of Mending, Power Infusion, Inner Focus.
- **Spell values against `spelltips/priest.json` and `trainables/priest.json`**: Shadow Word: Death 18 s and 30% backlash; Devouring Plague 24 s cooldown, heals the caster, trainable at 20; Vampiric Embrace 30 s; Fade 10 s on 30 s; Silence 5 s; Lightwell 3 min or 5 charges; Dark Sacrifice 1600 health for 1600 + Spirit mana, 10 min; Divine Grace 10 min, clears Weakened Soul; Penance three heals one second apart, 12 s; Weakened Soul 15 s. Levels 1 to 19 trainer levels (Shadow Word: Pain 4, Power Word: Shield 6, Renew 8, Mind Blast 10, Inner Fire 12, Psychic Scream 14, Heal 16, Dispel Magic 18, Holy Fire 20) and the low-rank lines (Renew rank 7 at 44, Greater Heal rank 2 and 3 at 46 and 52, rank 1 at 40, Heal rank 4 at 34).
- **Racials** against `races.json`: Touch of the Grave (5%, once per 15 s), Elune's Light, Artisan's Precision (1% hit), Berserking, Eureka!, Expansive Mind, Ruthless, The Human Spirit, Engineering Specialization, Will of the Forsaken. `recommendedRaces` matches the prose picks on all three pages; no paid race is recommended; Skyborne line unchanged.
- **Shadow card notes** against the priority prose (shadow.md:81): Shadow Word: Pain, Vampiric Embrace, Devouring Plague, Mind Blast with Inner Focus, Mind Blast, Shadow Word: Death, Mind Flay, Shoot. Consistent.
- **Source links** (fetched 2026-11-04): Blizzard Deep Dive recap 200; Power Up Gaming 200; Talents Forever data.json 200; foreversixty.gg /sim/weights, /planner, /planner?class=priest 200; Icy Veins talent calculator 403 (unknown, not dead, R5-D2).
- No em or en dashes on the four pages or in the priest APL notes; no site or sim words outside the planner and weights links.

## Decisions

- **R16-D1.** The Spirit clause at shadow.md:107 stays a SHOULD: the rule is true in every band the table renders, and the raid table itself shows Spirit at zero.
- **R16-D2.** Shadowfiend stays "not in Forever" (index.md:28): `trainables/priest.json` still carries the rank-0 level-1 placeholder row 401977, which round 2 D1 already settled from the client tables (no SpellLevels row, no teacher). Not re-filed.

## Sims requested

Carried, no prose change waits on them:

- **Discipline race pick with the top-rank shield and the pre-pull Renew**: band 60, published gear, Gnome against Human (Alliance), Undead against Troll (Horde), 10,000 iterations, seeds 1 and 50001.
- **Shadow race lines on the raid build** `50052201211312251`: Human, Night Elf, Dwarf, Gnome, Undead, Troll, same method.
- **BiS re-rank on the shipped builds** (post-launch).

## Sim findings

None new. Still open (routed to sim round 18): S-R2-1 (Power Infusion on self), S-R9-1 (Eureka! periodic mask), X1 (low-rank penalty above 30 and the top-rank exemption), W1, Lightwell Renew's crit source.

## Sources

- Repo data (70755): `talents/priest.json`, `spelltips/priest.json`, `trainables/priest.json`, `races.json`, `addon-builds.json` (priest-shadow, priest-discipline, priest-holy leveling orders), `bis/priest-shadow.json` (every band, weights and slots), `items/priest.json`; `data/curated/apl/priest-shadow.json` notes; `web/src/lib/planner/fs1.ts`
- Race measurements: `design/reviews/guide-accuracy/round3/race-sims.md`
- Blizzard, Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Talents Forever data export: https://talentsforever.com/data.json
- Power Up Gaming, Forever talent changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Icy Veins, Forever talent calculator (403, unknown): https://www.icy-veins.com/wow-forever/news/start-planning-your-wow-forever-character-with-our-new-talent-calculator/
