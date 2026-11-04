# Guide accuracy round 13: Priest

Audited on the round 13 audit tree (main 604b78a80: twelve guide rounds, sim round 15, guide-header-build), client 1.60.1.70755, engine pin 9f7f1f6e4 (`sim/enginever/version.go`; engine source read at scratchpad/gr10/engine).
Files: `web/src/content/guides/priest/index.md`, `discipline.md`, `holy.md`, `shadow.md`; every card note and opener card in `data/curated/apl/priest-*.json`; every frontmatter and body link. Read first: the round 1 to 12 priest reports and fix logs, round4/priest-levels.md, the round 2 to 5 race-sims files and the ledger through round 12.

## What moved since round 12

Nothing on the priest side. `git diff a8b147a6f 604b78a80` over the four guides, `data/curated/apl/priest-*.json`, `addon-builds.json` and `sim/enginever` is empty. The round 12 SHOULDs (S1, S2) were not applied, so the two card notes they named are unchanged. Every line was still re-read from scratch against the 70755 data.

## Decoded builds (70755 `talents/priest.json`, `addon-builds.json` decoded by tree, tier and column; index 0 = level 10)

- **Shadow leveling** `02510503103/0/52542001201300051` (20/0/31): Wand Specialization 10-11, Spirit Tap 12-16, Shadow Focus 17-21, Improved Mind Blast 22-25, Improved Shadow Word: Pain 26-27, Mind Flay 28, Improved Mind Flay 29-30, Blackout 31-32, Vampiric Embrace 33, Shadow Weaving 34-36, Darkness 37-41, Shadowform 42, Twin Disciplines 43-47, Silent Resolve 48, Mental Strength 49-53, Mental Agility 54-56, Inner Focus 57, Meditation 58-60. Tier gates hold (Mind Flay with 16 Shadow points, Vampiric Embrace with 21, Darkness with 25, Shadowform with 30).
- **Shadow raid** (the header build) `00510503103/0/50052201211312251` (18/0/33): Shadow Focus 10-14, Improved Mind Blast 15-18, Improved Shadow Word: Pain 19-20, Mind Flay 21, Improved Mind Flay 22-23, Improved Mind Blast 24, Shadow Reach 25-26, Dark Inspiration 27, Vampiric Embrace 28, Shadow Weaving 29-31, Silence 32, Devouring Contagion 33-34, Early Demise 35-36 (25 Shadow points before tier 5), Darkness 37-41, Shadowform 42, Discipline as leveling minus Wand Specialization.
- **Discipline** `00500503133001143/3350500303/0` (29/22/0): Twin Disciplines 10-14, Mental Strength 15-19, Meditation 20, Mental Agility 21-23, Inner Focus 24, Improved Power Word: Shield 25-27, Meditation 28-29, Penance 30, Soul Warding 31, Renewed Hope 32-35, Divine Aegis 36-38, Twilight Focus 39-41, Improved Renew 42-44, Holy Specialization 45-49, Divine Fury 50-54, Inspiration 55-57, Improved Healing 58-60.
- **Holy** `00500503103/33505003030121431/0` (17/34/0): Twilight Focus 10-12, Improved Renew 13-15, Holy Specialization 16-20, Divine Fury 21-25, Inspiration 26-28, Improved Healing 29-31, Holy Binding 32, Litany of Light 33-34, Spiritual Healing 35-37 (25 Holy points before tier 5), Spirit of Redemption 38, Spiritual Guidance 39, Prayer of Mending 40 (30 points, prerequisite Spirit of Redemption), Spiritual Guidance 41-43, Twin Disciplines 44-48, Mental Strength 49-53, Mental Agility 54-56, Inner Focus 57, Meditation 58-60.
- **Power Infusion swap** (discipline.md:60), re-derived by tier: Improved Renew 3 to 1 leaves Holy tier 0 at 9 points, so Divine Fury (tier 1, needs 5) holds, Inspiration (tier 2, needs 10) has 14 and Improved Healing (tier 3, needs 15) has 17. Renewed Hope 5 brings Discipline to 30 before Power Infusion's tier 6, and Penance is its prerequisite. Legal.

## Findings

None. Counts: CONFIRMED 0, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## SHOULDs (wording, not findings)

| id | file:line | text | suggested wording |
|---|---|---|---|
| PRIEST-R13-S1 | index.md:48 | "Skyborne ... cannot be priests: they play warrior, hunter, rogue, shaman, mage or druid." | True as a union, but `combos.json` gives High Order Skyborne (Alliance) warrior, hunter, rogue, mage, druid and Windshaper Skyborne (Horde) warrior, hunter, rogue, shaman, druid. A Skyborne cannot be both a shaman and a mage on one side. Optional: "they play warrior, hunter, rogue or druid, plus mage on the Alliance side or shaman on the Horde side." |
| PRIEST-R13-S2 | shadow.md:91 | "Put Shadow Word: Pain on adds that will live 15 seconds or more" | The same page's Execute line (:89) says two ticks, 6 seconds, already out-damage the Mind Flay the global would buy. 15 seconds is a safe mana rule, not a false one, but the two numbers read as if they disagree. Optional: "on adds that will live long enough for a few ticks, and keep your rotation on the kill target." |
| Carried | `priest-discipline.json` [17], `priest-holy.json` [20] (R12-S1, R12-S2) | "Lesser Heal when Greater Heal rank 1 would overheal"; "Heal rank 4 when Greater Heal would overheal or your mana runs short" | Unchanged and still optional. Both rows carry the same conditions as the row above them, so they fire only when the dearer heal cannot be cast. |
| Carried | shadow.md:69 (R11-S2); holy.md:59, :51 (R8-S1, R8-S2); shadow.md:65 (R9-S1) | | Still optional. |

## Verified clean this round

- **Talent lines, every one, against the 70755 rank text** (`talents/priest.json`): Shadowform (10%, half cost on Shadow spells, +100% crit bonus so a Shadow crit is 200% of a hit, no healing spells; Power Word: Shield is an absorb, so shadow.md:141 holds), Shadow Weaving 3/3 (100%, 2% "from you", 5 stacks, 15 s), Darkness 10%, Improved Mind Flay 20% per instance, Improved Mind Blast 0.5 s a point off 8 s, Improved Shadow Word: Pain 6 s, Shadow Focus 5%, Spirit Tap 5/5 (100%, Spirit doubled 15 s, 50% regeneration while casting), Blackout 2/5 (4%, 3 s), Vampiric Embrace 20% for 30 s, Early Demise 2/2 30% at or below 20%, Devouring Contagion 2/2 (50%, 2 enemies within 10 yds), Silence (5 s, 3 s interrupt), Shadow Reach 20%, Dark Inspiration 1/5 3% for 10 s, Silent Resolve 10% a point, Twin Disciplines and Mental Agility ("instant cast spells"; Penance and Renew in Twin Disciplines' mask per round 4), Inner Focus (free, +25% on non-periodic), Meditation 50%, Mental Strength 10%, Penance (three heals 1 s apart, 12 s), Renewed Hope 4/5 (8%, 4 s, the five spells named), Divine Aegis 15% of non-periodic crits, Soul Warding (removes the 4 s cooldown, 15%), Improved Power Word: Shield 20%, Holy Specialization 5% on Holy spells, Divine Fury 0.5 s on Heal and Greater Heal, Improved Healing 15% (Lesser Heal, Heal, Greater Heal, Penance, Prayer of Mending: each page names the ones its build has), Improved Renew 15%, Twilight Focus 70%, Inspiration armor only, Holy Binding 50% on Flash Heal on others, Litany of Light "previous heal was a different spell", Spirit of Redemption 15 s of free heals, Spiritual Guidance, Spiritual Healing 10%, Prayer of Mending ("Jumps up to 5 times", 10 s cooldown), Power Infusion 20% for 15 s, 3 min.
- **Spells** (`spelltips/priest.json`, `trainables/priest.json`): Shadow Word: Death 18 s, 30% maximum health backlash; Devouring Plague 24 s duration and cooldown, heals the caster; Mind Blast "a high amount of threat"; Mind Flay 3 s channel, 50% slow; Fade 10 s on 30 s; Lightwell 3 min or 5 charges, "Being attacked cancels the effect"; Prayer of Healing "the target and their party", rank 4 1030 mana against Greater Heal rank 4 at 655, so "your most expensive heal" holds; Flash Heal rank 7 is the dearest heal per point healed (825 for 380, against Greater Heal rank 4 at 1782 for 557 and Heal rank 4 at 651 for 259 after Improved Healing); Penance rank 4 at 1275 for 327 is the best heal per mana on the Discipline card; Prayer of Mending rank 3 at 390 for five heals; Divine Grace (free, below 50%, removes Weakened Soul, not on self, 10 min); Desperate Prayer free, 10 min; Dark Sacrifice 1600 health over 15 s for 1600 + Spirit mana, 10 min; Contingency Plan under 35%; Confounding Flash confuses up to 5 enemies within 8 yds; Fear Ward, Divine Spirit, Lightwell, Shadow Word: Death and Devouring Plague are trainer spells (`source: skill_line_ability`). Shadowfiend's level 1 row stays the round 2 D1 decision.
- **Leveling**: both Levels 1 to 19 and Levels 20 to 60 blocks on all three pages match the point orders above level by level and the trainer levels (Smite 1, Shadow Word: Pain 4, Power Word: Shield 6, Renew 8, Mind Blast 10, Inner Fire 12, Psychic Scream 14 at 2 enemies, Heal 16, Dispel Magic 18, Holy Fire and Devouring Plague 20, Heal rank 3 at 28). The two blocks read right in order and do not contradict each other. Wands and quest wands as in rounds 4 and 5 (`items.json` 11287 req 5, 11288 req 13, 5604, 5242).
- **Low-rank heals** against the engine's 17-level grace and 5% a level (`sim/core/spell.go:785-790`): Greater Heal ranks 2 to 4 keep the whole bonus at 60, rank 1 85%; Renew rank 7 whole, ranks 5 and 6 45% and 75%, ranks 1 to 3 none, rank 4 15%; Heal rank 4 55%, rank 3 25%, ranks 1 and 2 none; Lesser Heal none. index.md:29, discipline.md:70 to :74, holy.md:68 to :72 and the card notes say the same.
- **Card notes**: every rendered note on the three cards against its row's spell, rank and condition (Disc Penance tank 95% / party 70%, shield on Weakened Soul off and tank under 90%, Flash Heal under 35% / 30%, Prayer of Healing three of four under 85%, Greater Heal rank 1 under 65%; Holy Flash Heal 35% / 30%, Greater Heal and Heal waits 50% / 35%; Shadow Vampiric Embrace skipped under 10 s, Dark Sacrifice under a fifth). Only R12-S1 and R12-S2 remain, as SHOULDs.
- **Opener** (shadow.md:77, :79) against the prepull (Shadowform -20 s, Mind Blast rank 9 -1.5 s) and the round 11 cooldown arithmetic: still right.
- **Races**: race prose against round3/race-sims.md (Shadow Human 690.7, Night Elf 690.4, Dwarf 689.6, Gnome 687.8, Undead 695.0, Troll 687.8; Discipline Gnome 600.6, Human 598.0, Undead 585.2; Holy Undead 837.2, Human 827.5, Gnome 821.6, the other three 804 to 809). Racials against `races.json` (Ruthless, The Human Spirit, Artisan's Precision, Stoneform, Elune's Light, Berserking, Touch of the Grave, Will of the Forsaken, Eureka! per the ledger round 9 rule, Expansive Mind, Engineering Specialization). No paid race recommended; Skyborne cannot be priests (`combos.json`).
- **Consumables**: every name is in `items.json` (Elixir of Sages is Forever's 250338, not Era's Elixir of the Sages; Demonic Rune 12662; Major Mender's 250949; Major Spellblasting 250937).
- **Header build card**: Shadow loads the raid build; every Shadow talent line gives the right count for both builds and tags the raid-only ones. Discipline and Holy have no `raidBuild`.
- **Source links** (fetched 2026-11-04): Blizzard Deep Dive recap 200; Talents Forever data.json 200; Power Up Gaming 200; foreversixty.gg /sim/weights, /planner, /planner?class=priest 200; Icy Veins talent calculator 403 (unknown, not dead, R5-D2).
- No em or en dashes in the pages or the card notes; no site or sim words outside the planner links; the stat paragraphs explain without ranking.

## Decisions (no owner questions)

- **R13-D1 Round 12's decisions stand.** R12-D1 (Twin Disciplines and Mental Agility "your instants" follows the client text) and R12-D2 (the "would overheal" healer notes are SHOULDs) were re-checked against the client text and the row conditions; neither puts a false statement on the page.
- **R13-D2 The Skyborne line is not a finding.** The class list is the true union of the two Skyborne races; the per-faction split is wording (S1).

## Sims requested

Carried, still open (no priest prose changes until they run):

- **Discipline race pick with the top-rank shield**: spec-breakdown band 60, published gear, Gnome against Human (Alliance) and Undead against Troll (Horde), raid entry with Power Word: Shield rank 10, 10,000 iterations, seeds 1 and 50001. Confirms discipline.md:100-106 and index.md:35.
- **Shadow race lines on the raid build** `50052201211312251` with the pre-cast Mind Blast: Human, Night Elf, Dwarf, Gnome (Alliance), Undead and Troll (Horde), same method. Confirms shadow.md:115-119 and index.md:37.
- **PROVISIONAL, BiS re-rank**: the priest BiS files are still the 04:01 to 04:24 UTC run. No priest prose names a best-in-slot piece, so there is no row to compare.

## Sim findings (engine vs game)

None new. Still open, not re-filed: Power Infusion on yourself (S-R2-1); the Eureka! masks against the tooltip (S-R9-1); healer haste weights (X1); Discipline pre-pull Renew (D2); the low-rank curve above level 30 (W1); the SW:P cutoff (APL `remainingTime >= 10` against the page's 6 seconds).

## Sources

- Repo data (70755): `data/builds/1.60.1.70755/{talents,trainables,spelltips}/priest.json`, `addon-builds.json`, `races.json`, `combos.json`, `classes.json`, `items.json`, `simconsumes.json`, `simbuffs.json`, `bis/priest-*.json`; `data/curated/apl/priest-*.json`
- Engine 9f7f1f6e4: `sim/core/spell.go:785-800` (level penalty grace 17, 5% a level)
- Race measurements: `design/reviews/guide-accuracy/round3/race-sims.md`
- Blizzard, "World of Warcraft: Forever Class Deep Dives, Priest and Warrior" (30 Sep 2026): https://worldofwarcraft.blizzard.com/news/24301514
- Blizzard, Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Talents Forever data export: https://talentsforever.com/data.json; Power Up Gaming: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/; Icy Veins talent calculator news (403, unknown)
