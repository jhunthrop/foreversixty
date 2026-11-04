# Guide accuracy round 12: Priest

Audited on the round 12 audit tree (main a8b147a6f: guide rounds 1 to 11, sim round 15, leveling-order, guide-header-build), client 1.60.1.70755, engine pin 9f7f1f6e4.
Files: `web/src/content/guides/priest/index.md`, `discipline.md`, `holy.md`, `shadow.md`; the card notes and opener cards in `data/curated/apl/priest-*.json`; every frontmatter and body link. Read first: the round 1 to round 11 priest reports and fix logs, round4/priest-levels.md, the race-sims files and the ledger through round 11.

## What moved since round 11

- Guide text: only the round 11 fixes (shadow.md:73 and :79, holy.md:30 and :80). The builds, the point orders and the APL files are unchanged (`git diff d54de48aa HEAD`).
- guide-header-build (3bd24d3a2): the rail's build card, the key-talent chips and the rotation card now use `raidBuild` where a guide has one. That is Shadow only. Discipline and Holy have no `raidBuild`, so their header still shows the level 60 build. I re-read every Shadow line against the raid build the header now loads.

## Decoded builds (70755 `talents/priest.json` array order; `addon-builds.json` index 0 = level 10)

- **Shadow leveling** `02510503103/0/52542001201300051` (20/0/31). Order: Wand Specialization 10-11, Spirit Tap 12-16, Shadow Focus 17-21, Improved Mind Blast 22-25, Improved Shadow Word: Pain 26-27, Mind Flay 28, Improved Mind Flay 29-30, Blackout 31-32, Vampiric Embrace 33, Shadow Weaving 34-36, Darkness 37-41, Shadowform 42, Twin Disciplines 43-47, Silent Resolve 48, Mental Strength 49-53, Mental Agility 54-56, Inner Focus 57, Meditation 58-60.
- **Shadow raid** (shown in the header now) `00510503103/0/50052201211312251` (18/0/33): Shadow Focus 5, Improved Mind Blast 5, Improved Shadow Word: Pain 2, Shadow Reach 2, Mind Flay 1, Improved Mind Flay 2, Dark Inspiration 1, Vampiric Embrace 1, Shadow Weaving 3, Silence 1, Devouring Contagion 2, Early Demise 2, Darkness 5, Shadowform 1 (prereq 105820 = Vampiric Embrace, taken; 32 points before tier 6). Discipline: Twin Disciplines 5, Silent Resolve 1, Mental Strength 5, Mental Agility 3, Inner Focus 1, Meditation 3. The diff is 9 points each way, which matches shadow.md:73.
- **Discipline** `00500503133001143/3350500303/0` (29/22/0): Twin Disciplines 10-14, Mental Strength 15-19, Meditation 20, Mental Agility 21-23, Inner Focus 24, Improved Power Word: Shield 25-27, Meditation 28-29, Penance 30, Soul Warding 31, Renewed Hope 32-35, Divine Aegis 36-38, Twilight Focus 39-41, Improved Renew 42-44, Holy Specialization 45-49, Divine Fury 50-54, Inspiration 55-57, Improved Healing 58-60.
- **Holy** `00500503103/33505003030121431/0` (17/34/0): Twilight Focus 10-12, Improved Renew 13-15, Holy Specialization 16-20, Divine Fury 21-25, Inspiration 26-28, Improved Healing 29-31, Holy Binding 32, Litany of Light 33-34, Spiritual Healing 35-37, Spirit of Redemption 38, Spiritual Guidance 39, Prayer of Mending 40 (prereq 105854 = Spirit of Redemption), Spiritual Guidance 41-43, Twin Disciplines 44-48, Mental Strength 49-53, Mental Agility 54-56, Inner Focus 57, Meditation 58-60.

## Findings

None. Counts: CONFIRMED 0, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## SHOULDs (wording, not findings)

| id | file:line | text | suggested wording |
|---|---|---|---|
| PRIEST-R12-S1 | `priest-discipline.json` priorityList[17] (Discipline card note) | "Lesser Heal when Greater Heal rank 1 would overheal: a small top-up for little mana." | Row [17] has the same health and mana-pace conditions as [16] (Greater Heal rank 1, below 65%) and sits under it, so the card casts Lesser Heal only when it cannot cast Greater Heal rank 1. It never checks for overheal. Lesser Heal rank 3 (learned at 10) also gets none of your bonus at 60, and discipline.md:74 says to downrank only within the ranks from your forties. The card action is the accepted low-rank item. A note that does not argue with the prose: "Lesser Heal only when Greater Heal rank 1 is out of reach: a small top-up for little mana, though at 60 it gets almost nothing from your gear." |
| PRIEST-R12-S2 | `priest-holy.json` priorityList[20] (Holy card note) | "Heal rank 4 when Greater Heal would overheal or your mana runs short" | Same pattern: [20] has the same conditions as [10] (Greater Heal rank 4 on the tank below 80%), so it fires only when Greater Heal cannot be cast. "or your mana runs short" is the true half. Suggested: "Heal rank 4 when your mana will not cover Greater Heal: cheaper, though it keeps only part of your bonus at 60." |
| Carried | shadow.md:69 (R11-S2); holy.md:59, :51 (R8-S1, R8-S2); shadow.md:65 (R9-S1) | | Still optional. shadow.md:69 "your instants hit harder and cost less" is true to the client text: both Twin Disciplines and Mental Agility name "instant cast spells". The R11-S2 mask gap is engine detail, not a false line. |

## Verified clean this round

- **Round 11 fixes, re-checked.** shadow.md:79 opener: the pre-cast Mind Blast (APL prepull 10947 r9 at -1.5 s) lands at 0 and comes back at 5.5 s (cooldown 8000 ms, minus 2.5 s from the raid build's Improved Mind Blast 5/5). Shadow Word: Pain, Vampiric Embrace and Devouring Plague take the globals at 0, 1.5 and 3.0, and Shadow Word: Death (row [7], no condition) takes 4.5. Inner Focus (off the global, cooldown 180000) plus Mind Blast at 6.0, then Mind Flay. Shadow Weaving reaches 5 stacks within the string (Mind Blast, Pain, Plague, Death, Mind Blast). shadow.md:73 names the 9 traded points correctly. holy.md:30 "five heals" agrees with the Holy card [2] and Blizzard's deep dive. holy.md:80 "your most expensive heal": Prayer of Healing rank 4 at 1030 mana beats Greater Heal rank 4 at 655, or 557 after Improved Healing.
- **Header build card (guide-header-build).** Shadow's header now loads and labels the raid build. No Shadow line calls the header or "the build above" the leveling build. Leveling claims say "leveling" (:41, :61, :62, :73, Leveling section). Raid-only talents carry "raid build" (:64 to :68). The bullets the key-talent chips read give the right ranks for both builds: Improved Mind Blast "4 to 5", Spirit Tap and Blackout "leveling, none in the raid build", Early Demise 2 (raid 2/2, level 60 0/2, and its line says "the raid build's execute"). The Rotation section, now titled "For the raid build", names only talents the raid build takes: Early Demise at :89, Improved Mind Flay, Inner Focus, Vampiric Embrace, and the Silent Resolve to Mental Strength move at :95. That move keeps tier 2's 10-point gate (Twin Disciplines 5, Silent Resolve 2, Mental Strength 4 = 11). The Discipline and Holy headers are unchanged.
- **Leveling prose against `addon-builds.json`**, level by level: shadow.md:131 and :139 (Wand Specialization 10-11, Spirit Tap 12-16, Shadow Focus 17-21, Improved Mind Blast to 25, Improved Shadow Word: Pain 26-27, Mind Flay 28, Improved Mind Flay 29-30, Blackout 31-32, Shadowform 42, "cost you those two levels of the form"); discipline.md:118 and :124 (Twin Disciplines 10-14, Mental Strength 15-19, Penance 30); holy.md:116 (Twilight Focus 10-12, Improved Renew 13-15, Holy Specialization from 16). Trainer levels (`trainables/priest.json`): Smite 1, Shadow Word: Pain 4, Power Word: Shield 6, Renew 8, Mind Blast 10, Inner Fire 12, Psychic Scream 14 (rank 1, 2 enemies), Heal 16, Dispel Magic 18, Holy Fire 20, Devouring Plague 20, Flash Heal 20, Heal rank 3 at 28, Shadow Word: Death 32, Lightwell 40, Fear Ward 20, Divine Spirit 30; Penance ranks 30/40/50/60 and Prayer of Mending ranks 40/50/60 (index.md:27).
- **Low-rank heals**, using the 17-level grace and then 5% a level: Renew rank 7 (44) keeps the full bonus and ranks 5 (32) and 6 (38) keep part. Heal rank 4 (34) keeps part, Heal ranks 1 to 3 and Lesser Heal (1, 4, 10) next to none. Greater Heal ranks 2 and 3 (46, 52) keep the full bonus, rank 1 (40) keeps 85%. The pages say the same: discipline.md:70, :72, :74; holy.md:68, :70, :72; index.md:29; Disc card [16] "keeps most"; Holy card [20] "only part". Greater Heal rank 4 (10965) is the top trainer rank. Ranks 25314, 25315 and 25316 are Ahn'Qiraj books in `book-ranks.json`.
- **Talent text** (70755 rank text): Shadowform, Shadow Weaving, Darkness, Improved Mind Flay, Improved Mind Blast, Improved Shadow Word: Pain (6 s on 18 s), Shadow Focus, Spirit Tap 5/5, Blackout 2/5 (4%, 3 s), Vampiric Embrace (20%, 30 s, cooldown 30 s), Early Demise, Devouring Contagion, Silence (5 s and 3 s), Shadow Reach, Dark Inspiration 1/5 (3%, 10 s), Silent Resolve 1/3 (10%), Penance, Renewed Hope 4/5 (8%, 4 s), Divine Aegis, Soul Warding (the shield's 4000 ms cooldown removed, 15%), Improved Power Word: Shield, Improved Healing, Divine Fury, Holy Specialization, Improved Renew, Twilight Focus 70%, Inspiration (armor only), Litany of Light, Holy Binding, Spirit of Redemption, Spiritual Guidance 4/5, Spiritual Healing, Prayer of Mending, Power Infusion (20%, 15 s, 180000). The Power Infusion swap is legal: Discipline reaches 30 with Renewed Hope 5, and Holy keeps 17 points ahead of Improved Healing's tier.
- **Spells, races and consumables.** Dark Sacrifice rank 5 trades 1600 health for 1600 + Spirit mana over 15 s, cooldown 10 min. Divine Grace: below 50%, removes Weakened Soul, not on self, 10 min. Contingency Plan triggers under 35%. Lightwell lasts 3 min or 5 charges and being attacked cancels a click. Fade lasts 10 s, cooldown 30 s. Prayer of Healing heals "the target and their party". The `races.json` racials hold: Ruthless, The Human Spirit, Artisan's Precision, Stoneform, Elune's Light, Berserking, Touch of the Grave, Will of the Forsaken, Eureka!, Expansive Mind, Engineering Specialization. The race lines agree with round3/race-sims.md; per the round 11 note, the new builds are still unmeasured. Skyborne cannot be priests (`combos.json`). Every consumable is in `simconsumes.json`. Wands: Lesser Magic Wand 5, Greater Magic Wand 13.
- **Card notes**: Shadow [1] to [9] and Discipline [1], [2], [7], [14], [15], [16]; Holy [1], [2], [4], [5], [9], [10], [15], [25] match the spell and rank each row casts and the 70755 values. The two exceptions are the wording in S1 and S2.
- **Source links** (fetched 2026-11-04): Blizzard Deep Dive recap 200; Talents Forever data.json 200 (labelled beta client data); Power Up Gaming 200; foreversixty.gg /sim/weights (Shadow only), /planner and /planner?class=priest 200; Icy Veins talent calculator 403 (unknown, not dead, R5-D2). Every entry follows the round 4 source-link rule.
- No em or en dashes. No site or sim words outside the planner links. The stat prose explains without ranking stats. No paid race is recommended.

## Decisions (no owner questions)

- **R12-D1 The Twin Disciplines / Mental Agility line on Shadow stays as written.** The client text for both talents says "instant cast spells", and Shadow Word: Pain, Devouring Plague, Shadow Word: Death and Vampiric Embrace are all instants. A narrower engine mask is a sim detail, so "your instants hit harder and cost less" is not false on 70755.
- **R12-D2 The healer card notes that say "would overheal" are SHOULDs, not findings.** They give a real player's reason to downrank. The rows' conditions fire them only when the costlier heal cannot be cast, so the suggested wording names that instead.

## Sims requested

Carried from round 11, still open. No priest prose changes until they run:

- **Discipline race pick with the top-rank shield**: spec-breakdown band 60, published gear, Gnome against Human (Alliance) and Undead against Troll (Horde), raid entry with Power Word: Shield rank 10, 10,000 iterations, seeds 1 and 50001. Confirms discipline.md:100-106 and index.md:35.
- **Shadow race lines on the raid build** `50052201211312251` with the pre-cast Mind Blast: Human, Night Elf, Dwarf and Gnome (Alliance), Undead and Troll (Horde), same method. Confirms shadow.md:115-119 and index.md:37.
- **PROVISIONAL, BiS re-rank**: no priest prose names a best-in-slot piece, so there is no row to compare.

## Sim findings (engine vs game)

None new. Still open, not re-filed:

- Power Infusion on yourself (S-R2-1).
- The Eureka! masks against the tooltip (S-R9-1).
- Healer haste weights (X1).
- Discipline pre-pull Renew (D2).
- The low-rank curve above level 30 (W1).
- The SW:P cutoff (APL `remainingTime >= 10` against the page's 6 seconds).

## Sources

- Repo data (70755): `data/builds/1.60.1.70755/{talents,trainables,spelltips}/priest.json`, `talents.json` (prereq ids), `addon-builds.json`, `races.json`, `combos.json`, `simconsumes.json`; `data/curated/apl/priest-*.json`; `data/curated/book-ranks.json`
- guide-header-build: commit 3bd24d3a2 (`web/src/lib/guides/key-talents.ts`, `guide-rotation.ts`, `GuideActionRail.astro`)
- Blizzard, "World of Warcraft: Forever Class Deep Dives, Priest and Warrior" (30 Sep 2026): https://worldofwarcraft.blizzard.com/news/24301514
- Blizzard, Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Talents Forever data export: https://talentsforever.com/data.json; Power Up Gaming: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/; Icy Veins talent calculator news (403)
