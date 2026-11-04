# Guide accuracy round 11: Priest

Audited on the round 11 audit tree (main d54de48aa, round 10 landed), client 1.60.1.70755, engine pin 9f7f1f6e4 (`sim/enginever/version.go`; engine source read at scratchpad/gr10/engine).
Files: `web/src/content/guides/priest/index.md`, `discipline.md`, `holy.md`, `shadow.md`; the card notes and opener cards from `data/curated/apl/priest-*.json`; every frontmatter and body link. Read first: round 1 to round 10 priest reports and fix logs, round4/priest-levels.md, the race-sims files and the ledger through round 10.

## What moved since round 10

Only the round 10 fixes (shadow.md Leveling, talent lines, opener and pre-pull; discipline.md shield lines; holy.md filler; the Discipline and Holy card notes). Builds, point orders and card actions are unchanged since round 10. Every round 10 edit was re-checked against the 70755 data; one is wrong (R11-1).

## Decoded builds (70755, `talents/priest.json` array order; `addon-builds.json` index 0 = level 10)

- **Shadow leveling** `02510503103/0/52542001201300051` (20/0/31). Order: Wand Specialization 10-11, Spirit Tap 12-16, Shadow Focus 17-21, Improved Mind Blast 22-25, Improved Shadow Word: Pain 26-27, Mind Flay 28, Improved Mind Flay 29-30, Blackout 31-32, Vampiric Embrace 33, Shadow Weaving 34-36, Darkness 37-41, Shadowform 42, Twin Disciplines 43-47, Silent Resolve 48, Mental Strength 49-53, Mental Agility 54-56, Inner Focus 57, Meditation 58-60. Shadowform's tier 6 gate is met at 41 (30 Shadow points).
- **Shadow raid** `00510503103/0/50052201211312251` (18/0/33): Shadow Focus 5, Improved Mind Blast 5, ISWP 2, Shadow Reach 2, Mind Flay 1, Improved Mind Flay 2, Dark Inspiration 1, Vampiric Embrace 1, Shadow Weaving 3, Silence 1, Devouring Contagion 2, Early Demise 2, Darkness 5, Shadowform 1; Discipline as leveling minus Wand Specialization. The diff to leveling is 9 points each way, exactly as shadow.md:73 lists.
- **Discipline** `00500503133001143/3350500303/0` (29/22/0) and **Holy** `00500503103/33505003030121431/0` (17/34/0): unchanged; orders as in round 10 (Penance at 30; Holy Twilight Focus 10-12, Improved Renew 13-15, Holy Specialization 16-20, Prayer of Mending 40). The Power Infusion swap (Improved Renew 3 to 1, Renewed Hope 5, Power Infusion) stays legal: Disc reaches 30 before tier 6, Penance is the prerequisite, Holy keeps 17 points ahead of Improved Healing.

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| PRIEST-R11-1 | CONFIRMED | rotation (round 10 edit) | shadow.md:79 | "**Opener.** Mind Blast lands as the boss is engaged, then Shadow Word: Pain, Vampiric Embrace, Devouring Plague, Mind Blast again (with Inner Focus, so it is free and more likely to crit), Shadow Word: Death, and Mind Flay." | The pre-cast Mind Blast (APL prepull 10947 r9 at -1.5 s, 1.5 s cast) lands at 0 and starts its cooldown there: 8 s base (`spelltips/priest.json` 10947 cooldown_ms 8000) minus 2.5 s from Improved Mind Blast 5/5 is 5.5 s, so it is back at 5.5 s (6 s with the leveling build's 4 points). Shadow Word: Pain, Vampiric Embrace and Devouring Plague take three 1.5 s globals (0, 1.5, 3.0), so at 4.5 s Mind Blast is still on cooldown and the card's next ready row is Shadow Word: Death (priest-shadow.json priorityList [5] strictSequence Inner Focus + Mind Blast and [6] Mind Blast are not ready; [7] Shadow Word: Death has no condition). Mind Blast with Inner Focus follows at 6.0 s. The round 10 rewrite put the two in the wrong order; casting it as written means a second of standing still. | "**Opener.** Mind Blast lands as the boss is engaged, then Shadow Word: Pain, Vampiric Embrace, Devouring Plague and Shadow Word: Death. By then Mind Blast is back: Inner Focus first, so it is free and more likely to crit, then Mind Flay. By the end of that string Shadow Weaving is at five stacks." |

Counts: CONFIRMED 1, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## SHOULDs (wording, not findings)

| id | file:line | text | suggested wording |
|---|---|---|---|
| PRIEST-R11-S1 | shadow.md:73 | "The leveling build opens with two points in Wand Specialization ..., and its last points go to Discipline. The raid build trades them, Blackout and Spirit Tap for ..." | "them" reads as the Discipline points, which the raid build keeps. "The raid build trades Wand Specialization, Blackout and Spirit Tap for ..." |
| PRIEST-R11-S2 | shadow.md:69 | "your instants hit harder and cost less" | Mental Agility's client mask (engine `PriestSpellMaskMentalAgility`, 14520) leaves out Shadow Word: Death and Mind Flay, and Twin Disciplines covers Mind Flay. Truer: "Shadow Word: Pain, Devouring Plague, Shadow Word: Death and Mind Flay hit harder, your DoTs cost less, and Inner Focus is a free Mind Blast every 3 minutes." |
| PRIEST-R11-S3 | holy.md:80 | "Prayer of Healing, your most expensive spell" | Prayer of Fortitude (3400 base) costs more. "your most expensive heal". |
| PRIEST-R11-S4 | holy.md:30 | "five jumps for one cast" | The Holy card note says "five heals for one cast", and Blizzard's deep dive says it "continues jumping until it heals a total of 5 times". The client text reads "Jumps up to 5 times", so the line is not false, but "five heals for one cast" agrees with the card and the designers. |
| Carried | holy.md:59, :51, :30; shadow.md:65 | R8-S1, R8-S2, R8-S4, R9-S1 | Still optional. |

## Verified clean this round (no finding)

- **Round 10 fixes re-checked.** Shadow leveling text (shadow.md:34, :41, :131, :137, :139) matches `addon-builds.json` level by level, including "Shadowform is your point at 42" and "The two early points in Wand Specialization cost you those two levels of the form". Spirit Tap 5/5 rank text ("100% chance ... For the duration, 50% of your Mana regeneration will continue while casting") backs :61 and :131. Blackout 2/5 4%, 3 s. New talent lines :64 to :68 match the 70755 rank text (Devouring Contagion 50% and 2 enemies within 10 yds; Silence 5 s silence, 3 s interrupt; Shadow Reach 20%; Dark Inspiration 1/5 3% for 10 s). Before the pull (:77) matches the opener card (Shadowform -20 s, cost_pct 40, so "the form costs mana" holds; Mind Blast -1.5 s). Discipline :26, :38, :68 and the Discipline card note [7]: Blizzard's Priest and Warrior deep dive, fetched this round, says "Rage gain also ignores all damage absorption effects, so Power Word: Shield will no longer decrease Rage gained", and the card casts rank 10. Holy :26, :70: Greater Heal rank 2, 3, 4 learned at 46, 52, 58 keep the full bonus at 60; rank 1 (40) is 3 levels past the 17-level grace, 85%. Heal rank 4 and Greater Heal are both 3.0 s base, 2.5 s with Divine Fury. Card notes Disc [16], [17], Holy [10], [15], [20], [25] name the spell and rank each row casts.
- **Litany of Light** "even at different ranks, earn nothing" (holy.md:56): the deep dive says the refund needs "the name of the spell" to differ, and the engine keys it on SpellCode, shared by every rank.
- **Every other talent line** on the three pages against its 70755 rank text: Shadowform (crit bonus +100% doubles the 50% spell crit bonus, so "a Shadow crit hits for twice a normal hit" at :103 holds), Shadow Weaving 2% from you, 5 stacks, 15 s; Darkness, Improved Mind Flay 20% per instance, Improved Mind Blast 0.5 s a point off 8 s, ISWP 6 s, Shadow Focus 5%, Early Demise 30%, Silent Resolve 10% a point, Twin Disciplines, Mental Agility, Inner Focus (3 min, 25%, non-periodic), Meditation 50%, Mental Strength 10%; Penance (three heals 1 s apart, 12 s), Renewed Hope 4/5 (8%, 4 s), Divine Aegis 15%, Soul Warding (4 s cooldown removed, 15%), Improved Power Word: Shield 20%, Holy Specialization, Divine Fury, Improved Healing, Improved Renew, Twilight Focus 70%, Inspiration (armor only), Holy Binding 50%, Spirit of Redemption 15 s, Spiritual Guidance, Spiritual Healing 10%, Prayer of Mending (prerequisite Spirit of Redemption), Power Infusion 20% for 15 s, 3 min.
- **Spell facts** (`spelltips/priest.json`, `trainables/priest.json`): Shadow Word: Death 18 s, 30% max health backlash; Devouring Plague 24 s cooldown and duration, heals the caster; Vampiric Embrace 20%, 30 s, 30 s cooldown; SW:P r8 18 s in 3 s ticks; Mind Flay 3 ticks, 50% slow; Fade 10 s, 30 s cooldown; Lightwell 3 min or 5 charges, attacked cancels; Prayer of Healing heals the target and their party; Divine Grace below 50%, removes Weakened Soul, 10 min; Contingency Plan under 35%; Dark Sacrifice 1600 health for 1600 + Spirit mana, 10 min; Weakened Soul 15 s. Learn levels in every Levels 1 to 19 paragraph (Smite 1, SW:P 4, PW:S 6, Renew 8, Mind Blast 10, Inner Fire 12, Psychic Scream 14, Heal 16, Dispel Magic 18, Holy Fire 20) and in the leveling prose (Devouring Plague 20, Heal rank 3 at 28). Prayer of Healing rank 4 as "the top rank a trainer teaches" (Disc card [14]) agrees with `data/curated/book-ranks.json` (25316 is an Ahn'Qiraj book rank).
- **Wands and consumables**: Lesser Magic Wand required level 5, Greater Magic Wand 13 (`items.json` 11287, 11288); Elven Wand and Cinder Wand exist. Flask of Supreme Power, Elixir of Shadow Power, Greater Arcane Elixir, Brilliant Wizard and Mana Oil, Major Spellblasting and Major Mender's Potion, Flask of Distilled Wisdom, Greater Mageblood, Greater Cleric's, Elixir of Sages, Sage's Tea, Major Mana Potion all in `simconsumes.json`.
- **Racials** against `races.json`: Ruthless 1% crit, The Human Spirit 10%, Artisan's Precision 1% hit, Stoneform, Elune's Light 10% / 15 s / 3 min, Berserking 10% casting speed / 15 s / 3 min, Touch of the Grave (5% max health, once every 15 s), Will of the Forsaken (charm, fear, sleep), Eureka! 10% / 8 s / 2 min, Expansive Mind 5%, Engineering Specialization. Index Eureka! line follows the round 9 rule.
- **Index page**: every line holds.
- **Source links** (fetched 2026-11-04): Blizzard Deep Dive recap 200; Talents Forever data.json 200 (labelled beta client data); Power Up Gaming 200; foreversixty.gg /sim/weights (Shadow only), /planner and /planner?class=priest 200; Icy Veins talent calculator 403 (unknown, not dead, R5-D2). Entries follow the round 4 source-link rule.
- No em or en dashes; no site or sim words outside the planner links; stat prose explains without ranking.

## Decisions (no owner questions)

- **R11-D1 Opener order follows the cooldown arithmetic.** With the pre-cast Mind Blast, the second Mind Blast comes after Shadow Word: Death, not before it. The card already plays it that way; only the prose moves.
- **R11-D2 Improved Mind Flay keeps its client text.** Blizzard's 30 Sep deep dive says Improved Mind Flay "increases the damage and range of Mind Flay but reduces its movement speed slowing effect"; the 70755 rank text names only the per-tick ramp, and Mind Flay rank 6 still reads "slowing their movement speed by 50%". The client description wins (round 2 lesson), so shadow.md:91 "Mind Flay slows by 50%" stands. If a log shows a smaller slow with the talent, drop the number.

## Sims requested

Carried from round 10, still open (no priest prose changes until they run):

- **Discipline race pick with the top-rank shield**: spec-breakdown band 60, published gear, Gnome against Human (Alliance) and Undead against Troll (Horde), raid entry with Power Word: Shield rank 10, 10,000 iterations, seeds 1 and 50001. Confirms discipline.md:100-102 and index.md:35.
- **Shadow race lines on the raid build** `50052201211312251` with the pre-cast Mind Blast: Human, Night Elf, Dwarf, Gnome (Alliance) and Undead, Troll (Horde), same method. Confirms shadow.md:115-119 and index.md:37.
- **PROVISIONAL, BiS re-rank**: no priest prose names a best-in-slot piece, so there is no row to compare.

## Sim findings (engine vs game)

None new. Still open, not re-filed: Power Infusion on yourself (S-R2-1); the Eureka! masks against the tooltip (S-R9-1); healer haste weights (X1); Discipline pre-pull Renew (D2); the low-rank curve above level 30 (W1); the SW:P cutoff (APL `remainingTime >= 10` against the page's 6 seconds).

## Sources

- Repo data (70755): `data/builds/1.60.1.70755/{talents,trainables,spelltips}/priest.json`, `addon-builds.json`, `items.json`, `simconsumes.json`, `races.json`; `data/curated/apl/priest-*.json`; `data/curated/book-ranks.json`
- Engine 9f7f1f6e4: `sim/priest/healing_talents.go` (Litany of Light keyed on SpellCode), `sim/priest/priest.go` (Twin Disciplines and Mental Agility masks), `sim/priest/mind_flay.go`
- Blizzard, "World of Warcraft: Forever Class Deep Dives, Priest and Warrior" (30 Sep 2026): https://worldofwarcraft.blizzard.com/news/24301514 (Power Word: Shield and rage; Litany of Light "name of the spell"; Prayer of Mending "heals a total of 5 times"; Improved Mind Flay)
- Blizzard, Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Talents Forever data export: https://talentsforever.com/data.json; Power Up Gaming: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/; Icy Veins talent calculator news (403)

## Fix log (round 11)

- PRIEST-R11-1: fixed, shadow.md:79 with the report's text. Checked: Mind Blast 10947 cooldown 8000 ms, Improved Mind Blast 5/5 -2.5 s, so back at 5.5 s; three globals end at 4.5 s, Shadow Word: Death (priorityList[7]) fills it, Inner Focus + Mind Blast at 6.0 s.
- PRIEST-R11-S1: applied, shadow.md:73 "trades Wand Specialization, Blackout and Spirit Tap".
- PRIEST-R11-S3: applied, holy.md:80 "your most expensive heal" (Prayer of Fortitude 3400 is dearer than Prayer of Healing).
- PRIEST-R11-S4: applied, holy.md:30 "five heals for one cast".
- PRIEST-R11-S2: left, needs the engine mask read, not one line.
