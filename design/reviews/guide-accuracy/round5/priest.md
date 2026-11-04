# Guide accuracy round 5: Priest

Audited on the round 5 audit tree (land/guides-r4, HEAD 42f10dda9 on top of 752c02e43), client 1.60.1.70755, engine pin d1d9f2775.
Files: `web/src/content/guides/priest/index.md`, `discipline.md`, `holy.md`, `shadow.md`, the card notes in `data/curated/apl/priest-*.json`, and every frontmatter source link. Read first: round1 through round4 priest reports (round 4 includes its fix log), round4/priest-levels.md, round2/3/4 race-sims.md and the ledger, including the round 4 source-link rule.

## Decoded builds and point orders (re-derived this round; unchanged since round 2, all legal on 70755)

The FS1 strings are byte-identical to round 4: Discipline `00500503133001143/3350500303/0` (gnome), Holy `00500503103/33505003030121431/0` (human), Shadow leveling `02510503103/0/54342001201300051` and raid `00510503103/0/52352101201300251` (human). `addon-builds.json`, decoded again against `talents/priest.json` ([tree, tier+1, column+1], index 0 = level 10):

- Discipline (leveling = raid): Twin Disciplines 10-14, Mental Strength 15-19, Meditation 20, Mental Agility 21-23, Inner Focus 24, IPWS 25-27, Meditation 28-29, Penance 30, Soul Warding 31, Renewed Hope 32-35, Divine Aegis 36-38, Twilight Focus 39-41, Improved Renew 42-44, Holy Specialization 45-49, Divine Fury 50-54, Inspiration 55-57, Improved Healing 58-60.
- Holy (leveling = raid): Twilight Focus 10-12, Improved Renew 13-15, Holy Specialization 16-20, Divine Fury 21-25, Inspiration 26-28, Improved Healing 29-31, Holy Binding 32, Litany of Light 33-34, Spiritual Healing 35-37, Spirit of Redemption 38, Spiritual Guidance 39, Prayer of Mending 40, Spiritual Guidance 41-43, Twin Disciplines 44-48, Mental Strength 49-53, Mental Agility 54-56, Inner Focus 57, Meditation 58-60.
- Shadow leveling: Shadow Focus 10-14, Blackout 15-18, Spirit Tap 19-21, IMB 22-25, ISWP 26-27, Mind Flay 28, IMF 29-30, VE 31, Shadow Weaving 32-34, Darkness 35-39, Shadowform 40, Wand Specialization 41-42, Twin Disciplines 43-47, Silent Resolve 48, Mental Strength 49-53, Mental Agility 54-56, Inner Focus 57, Meditation 58-60.
- Shadow raid: Shadow Focus 10-14, Blackout 15-16, Spirit Tap 17-19, IMB 20-24, ISWP 25-26, Shadow Reach 27, Mind Flay 28, IMF 29-30, VE 31, Shadow Weaving 32-34, Early Demise 35-36, Darkness 37-39, Shadowform 40, Darkness 41-42, then Discipline as leveling minus Wand Specialization.

## Verified clean this round (no finding)

**Levels 1 to 19 (all three, now paragraphs).** Trainer levels from `trainables/priest.json`: Smite 1, Shadow Word: Pain 4, Power Word: Shield 6, Renew 8, Mind Blast 10, Inner Fire 12, Psychic Scream 14, Heal 16, Dispel Magic 18, Holy Fire 20. Talent paragraphs match the orders above exactly (Shadow Focus 5 then Blackout 4 to 18, and the Levels 20 to 60 paragraph picks up "From 19 ... Spirit Tap to 21"; Twin Disciplines 10-14 and Mental Strength 15-19; Twilight Focus 10-12, Improved Renew 13-15, Holy Specialization from 16). Rank text at those levels: Shadow Focus 5%, Blackout "a small chance" (8% at 4), Twin Disciplines "instant cast spells" (the round 4 fix to Shadow Word: Pain and Renew holds; Power Word: Shield's mask 0x1 is in neither effect), Mental Strength 10%, Twilight Focus 70% ("most hits"), Improved Renew 15%. Items (`items.json`): Lesser Magic Wand 11287 required level 5, Greater Magic Wand 11288 required level 13, Elven Wand 5604 and Cinder Wand 5242 required level 0. Quests (`loot.json` quests map): 5604 from Bashal'Aran (quest 957, alliance, level 13), 5242 from Arugal's Folly (quest 99, horde, level 15). Worn Mace (36) is the starting mace. No contradiction with the Levels 20 to 60 text (Shadow wands until Mind Flay at 28; Holy's Heal from 16 and Heal rank 2 vs rank 3 at 30, Heal ranks at 16/22/28/34; Discipline's Penance at 30) or with the strips.

**Round 4 fixes re-checked, all true.** discipline.md:118 (Twin Disciplines names Shadow Word: Pain and Renew only); healer pages carry no stat-weights source; "Talents Forever, talent data from the beta client" on all four pages; holy.md:116 "most hits"; holy.md:62 Spirit of Redemption (70755 text: 15 sec, "can cast any healing spell free of cost").

**Everything else, line by line** against `talents/priest.json` rank text and `spelltips/priest.json`: Shadowform ("you may not cast healing spells", so Power Word: Shield still works, shadow.md:137 and :43 hold; 40% base mana); Shadow Weaving "from you", 15 s, 5 stacks; Darkness 10%; Improved Mind Flay 20% per instance; IMB 0.5 s a point off Mind Blast's 8 s; ISWP 6 s; Spirit Tap 3/5 60%, 50% regen while casting, VE deaths; Blackout 8%/4%, 3 s; VE 20%, 30 s, cooldown 30 s; Early Demise 30%; Silent Resolve 10% a point; Shadow Word: Death 18 s cooldown, 30% max health backlash; Devouring Plague 24 s cooldown and duration, heals the caster; Fade 30 s cooldown; Mind Flay 50% slow; SW:P r8 18 s (3 s ticks, the 6 second cutoff holds); Penance three heals 1 s apart, 12 s; Renewed Hope 4/5 8% and 4 s, on Flash Heal, Lesser Heal, Heal, Greater Heal, Penance; Divine Aegis 15% of non-periodic crits; Soul Warding 4 s and 15%; IPWS 20%; Inner Focus 3 min, 25% non-periodic crit; Power Infusion 20% for 15 s, 3 min cooldown, no self-cast restriction; Improved Healing names Lesser Heal, Heal, Greater Heal, Penance, Prayer of Mending; Divine Fury 0.5 s on Heal and Greater Heal; Holy Binding 50% on others; Litany of Light 10% "previous heal was a different spell"; Spiritual Guidance 4/5 20% of Spirit; Spiritual Healing 10%; Prayer of Mending "Jumps up to 5 times", 10 s; Lightwell 3 min or 5 charges, "Being attacked cancels the effect", 10 min; Prayer of Healing "the target and their party"; Divine Grace below 50%, removes Weakened Soul, 10 min, not on self; Contingency Plan under 35%, 30 s ward, 10 min; Dark Sacrifice 1600 health for 1600 + Spirit mana, 10 min. Learn levels for index.md:26-27 (Fear Ward 20, Devouring Plague 20, Divine Spirit 30, Shadow Word: Death 32, Lightwell 40; Penance ranks 40/50/60, Prayer of Mending 50/60). Low-rank arithmetic (index.md:29, discipline.md:74, holy.md:72) against the 70755 learn levels and the ForeverChanges curve: Renew r5 (32) 45%, r6 (38) 75%, r7 (44) full; Heal r4 (34) 55%; Greater Heal r1 (40) 85%. Racials against `races.json`. Race prose against round3 race-sims (re-measured and held in round 4). Stat prose and frontmatter against band 60 raid scale factors, re-read: Disc mp5 2.52 > Spirit 1.30 > healing 1.17 > Int 1.00 > crit 0.67; Holy mp5 2.73 > healing 1.69 > Spirit 1.45 > Int > crit 0.56; Shadow spell power = Shadow power 5.75 > hit 4.61 > haste 4.44 > crit 3.83 > Int 1; hit_to_cap spell 16, Shadow 5, to cap 11 (shadow.md:99). Skyborne classes (index.md:48). No em or en dashes; no site or sim words outside the planner links (round 2 D2).

**Card notes.** No new false note. The healer rank-1 notes and the "learned below level 20" top-level notes stay with the rank move (R3-D1, R4-D2). Shadow notes match the page (VE cost, Devouring Plague, Inner Focus, Mind Blast, SW:D backlash, Mind Flay ticks).

**Source links (ledger round 4 rule).** Blizzard Deep Dive panel recap (24303313): HTTP 200; names the Gnome priest, baseline Divine Spirit, Cannibalize mana for priests. Power Up Gaming, Forever talent changes vs Classic: 200. Talents Forever data.json: 200, labelled as beta client data per the rule. foreversixty.gg/sim/weights: 200, Shadow only (damage spec, per the rule); the healer pages no longer carry it. foreversixty.gg/planner and /planner?class=priest: 200, labelled "character planner" (talent sources open the planner). No /sources links, no "how <spec> plays" entries, no wago.tools links. Icy Veins talent calculator news (index.md): 403 to automated fetch, unknown, not dead (the shaman index carries the same entry).

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| PRIEST-R5-1 | LIKELY | race / rotation (round 3 rendered-review edit) | shadow.md:95 | "Gnome Eureka! on a Mind Blast and a Shadow Word: Death (it skips every periodic tick, Mind Flay's included, so never spend it on a refresh)." | The priest Eureka! row (1259823) as generated from the 70755 client effect masks (`sim/core/eureka_auto_gen.go` at d1d9f2775): Mind Flay 18807 and every other Mind Flay rank sit in the damage set (SPELLMOD_DAMAGE) AND in the periodic set (SPELLMOD_DOT), beside Starshards; the periodic set holds no Shadow Word: Pain or Devouring Plague id. The cost set reaches Mind Blast and Mind Flay too. So the client rows boost Mind Flay's ticks and skip only the true DoTs. The tooltip (Wowhead Forever 1259823: "non-periodic damaging and healing abilities") reads the same once a channel counts as a cast ability rather than a DoT: Blizzard put Mind Flay in the DoT modifier on purpose, and put no DoT there. Round 1 read the masks this way; round 3's rendered review changed the line by analogy with the warlock tooltip, without a priest mask check. Not CONFIRMED only because no Forever combat log has been read. | "Gnome Eureka! on a Mind Blast and a Shadow Word: Death, and fill the rest of it with Mind Flay. It skips Shadow Word: Pain and Devouring Plague, so never spend it on a refresh." |
| PRIEST-R5-2 | LIKELY | talent / voice | discipline.md:58 | "**Improved Renew** (3 points) and **Twilight Focus** (3 points): a better Renew on the tank, and casts that raid damage rarely pushes back." | Twilight Focus 3/3: "Gives your spells a 70% chance to not lose casting time when you take damage" (`talents/priest.json`). Three hits in ten still push you back, which is not rare. Holy's own lines say it straight ("70% of the hits you take", holy.md:58; "most hits", holy.md:116), and round 4 fixed the same overstatement on the Holy page (PRIEST-R4-4). | "**Improved Renew** (3 points) and **Twilight Focus** (3 points): a better Renew on the tank, and most raid hits no longer push back your casts." |

Counts: CONFIRMED 0, LIKELY 2, NEEDS-SIM 0, NEEDS-OWNER 0.

## Decisions (no owner questions)

- **R5-D1 Eureka! on Mind Flay follows the client masks.** The rows the server runs name Mind Flay in the periodic modifier and leave every DoT out; the tooltip is the same template sentence on every class's row. The page should match the rows (PRIEST-R5-1). Engine side, nothing to do for the priest: the engine already applies the bonus to Mind Flay ticks. The warlock decision (WARLOCK-R3-2 / D2, Incinerate from the tooltip) is not touched by this: that row has no Incinerate id, a gap the tooltip fills; this row has an explicit Mind Flay id, which the tooltip does not contradict.
- **R5-D2 The index Icy Veins entry stays.** It is a community news source, not a talent-tree link, and it is not dead (403 only to automated fetches). Not a finding.
- **R5-D3 Card notes**: as R4-D2. No change.

## Sims requested

None. Both findings are text against client data.

## Sim findings (engine vs game)

None new. Still open from earlier rounds, not re-filed: Starshards cooldown and a Shadow line (S-R3-1/2); Power Infusion on yourself (S-R2-1); the Forever low-rank healing curve and the healer card ranks with their notes; Shadow leveling Mind Flay at 20 and Spirit Tap 5/5; the Shadow raid Blackout 2 + Spirit Tap 3 move; the SW:P cutoff `remainingTime >= 10` to `>= 6`; negative healer haste weights. On the Eureka! mask item carried in round 4's list: for the priest the engine already matches the client rows (R5-D1), so the item is warlock-only.

## Sources

- Repo data (70755): `data/builds/1.60.1.70755/{talents,trainables,spelltips}/priest.json`, `addon-builds.json`, `items.json`, `loot.json` (quests map), `races.json`, `bis/priest-*.json`; `data/curated/apl/priest-*.json`, `data/curated/rank-choices.json`; round3/race-sims.md, round4/race-sims.md
- Engine d1d9f2775: `sim/core/eureka_auto_gen.go` (priest row 1259823: damage, periodic and cost sets), `sim/core/eureka.go`
- Wowhead Forever tooltip, Eureka! (priest) 1259823: https://nether.wowhead.com/forever/tooltip/spell/1259823 ("non-periodic damaging and healing abilities ... 10% ... 8 sec", 2 min)
- Blizzard, World of Warcraft: Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200)
- Power Up Gaming, New WoW Forever talents: changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/ (200)
- Talents Forever data export: https://talentsforever.com/data.json (200)
- Forever Sixty, stat weights: https://foreversixty.gg/sim/weights (200); character planner: https://foreversixty.gg/planner (200)
- Icy Veins, Forever talent calculator news: https://www.icy-veins.com/wow-forever/news/start-planning-your-wow-forever-character-with-our-new-talent-calculator/ (403, unknown)
- ForeverChanges, WoW Forever downranking: https://foreverchanges.pro/downranking (curve used for the low-rank arithmetic, as in round 2)

## Fix log (round 5)

- **PRIEST-R5-1** (LIKELY, adopted by the integrator per R5-D1): fixed, shadow.md:95. Now "Gnome Eureka! on a Mind Blast and a Shadow Word: Death, and fill the rest of it with Mind Flay. It skips Shadow Word: Pain and Devouring Plague, so never spend it on a refresh." Swept every priest page and `priest-*.json` card note for any other line saying Eureka! skips Mind Flay or periodic ticks: none (the other Eureka! lines, discipline.md:100, holy.md:84 and :104, shadow.md:115, index.md:35 and :38, say only 10% cheaper and stronger for 8 seconds every 2 minutes, which holds).
- **PRIEST-R5-2** (LIKELY, verified now): fixed, discipline.md:58. Twilight Focus 3/3 is "a 70% chance to not lose casting time when you take damage" (`talents/priest.json`), so "rarely pushes back" overstated it. Now "most raid hits no longer push back your casts", matching holy.md:58 and :116.
- **R5-D1**: applied (PRIEST-R5-1). No engine change for the priest.
- **R5-D2**: applied, the index Icy Veins entry stays.
- **R5-D3**: applied, card notes unchanged.
- New SIM findings: none.

## Fix log (round 5, review pass)

- **MUST 1** (discipline.md:88): applied as the reviewer wrote it. Band 60 raid scale factors (`bis/priest-discipline.json`, both factions): mp5, healing power, Intellect, crit, Spirit last. Haste is not on items for this spec and stays out of the prose.
- **MUST 2** (holy.md:88): applied as the reviewer wrote it. Band 60 raid: Alliance Human has Spirit just ahead of healing power, Horde Undead has healing power ahead of Spirit, so "gear decides which leads" holds for both. mp5 first, Intellect, crit last unchanged.
- **MUST 3** (`priest-shadow.json`, first Mind Blast line, notes only): now "Mind Blast goes on cooldown, ahead of Mind Flay. When Inner Focus is up, use it first: the Mind Blast is free and has 25% more crit chance." Action untouched; file still parses as JSON.
- **SHOULD** (shadow.md:99 hit line): applied. "the other 11% comes from gear" became "your gear finds the rest", which no longer disagrees with the component's 1% built in plus 10% gear split. The component line itself is not mine.
- Eureka! SHOULD: rejected by the integrator, R5-D1 stands. Lesser Heal icon and BiS race captions: data lane, untouched.
