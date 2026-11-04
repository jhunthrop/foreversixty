# Guide accuracy round 7: Priest

Audited on the round 7 audit tree (land/guides-r6, HEAD 394562f20), client 1.60.1.70755, engine pin d1d9f2775 (`sim/enginever/version.go`).
Files: `web/src/content/guides/priest/index.md`, `discipline.md`, `holy.md`, `shadow.md`, the card notes in `data/curated/apl/priest-*.json`, and every frontmatter and body link. Read first: round 1 to round 6 priest reports (every fix log), round4/priest-levels.md, round2 to round5 race-sims.md and the ledger through round 6.

What moved since the round 6 audit tree (367cbadb7): only the round 6 commit 394562f20 on the priest pages (the three `## Stat priority` paragraphs rewritten to explain and not rank, holy.md:122 and discipline.md:124 leveling openers, `updated: 2026-11-04`). The merges 2b869f01d and f6503fd74 and the data fix 707bfd345 touch BiS components, `items/priest.json` and spelltip icons, no talent, trainer, spellconst or `addon-builds.json` row. Every FS1 string is unchanged since round 2.

## Decoded builds and point orders (unchanged, all legal on 70755)

- Discipline `00500503133001143/3350500303/0` (gnome): Twin Disciplines 5, Mental Strength 5, Mental Agility 3, Inner Focus 1, IPWS 3, Meditation 3, Soul Warding 1, Penance 1, Renewed Hope 4, Divine Aegis 3 (29); Twilight Focus 3, Improved Renew 3, Holy Specialization 5, Divine Fury 5, Inspiration 3, Improved Healing 3 (22). Order TD 10-14, MS 15-19, Med 20, MA 21-23, IF 24, IPWS 25-27, Med 28-29, Penance 30, SW 31, RH 32-35, DA 36-38, then Holy 39-60.
- Holy `00500503103/33505003030121431/0` (human): Discipline 17, Holy 34 (TF 3, IR 3, HS 5, DF 5, Insp 3, IH 3, Holy Binding 1, Litany 2, SoR 1, Spiritual Guidance 4, Spiritual Healing 3, PoM 1). Order TF 10-12, IR 13-15, HS 16-20, DF 21-25, ... PoM 40, Discipline 44-60.
- Shadow leveling `02510503103/0/54342001201300051`, raid `00510503103/0/52352101201300251` (human). Leveling order SF 10-14, Blackout 15-18, Spirit Tap 19-21, IMB 22-25, ISWP 26-27, Mind Flay 28, IMF 29-30, VE 31, Shadow Weaving 32-34, Darkness 35-39, Shadowform 40, Wand Specialization 41-42, Discipline 43-60. Raid trade (-2 WS, -2 Blackout, +2 Early Demise, +1 Shadow Reach, +1 IMB) balances; shadow.md:69 holds.
- Power Infusion swap (discipline.md:60): Disc 29 + Renewed Hope 5th = 30 before PI's tier, Penance is its prerequisite, Holy keeps 17 (TF 3, IR 1, HS 5, DF 5, Insp 3) ahead of Improved Healing's 15-point tier. Legal.

## Verified clean this round (no finding)

**Round 6 stat prose, reason by reason, against 70755 text.**
- Discipline (discipline.md:88): mp5 is mana while casting (true by definition); "Healing power is the size of every shield, Penance and heal": Power Word: Shield r10 (10901) carries `sp_coefficient` 0.1 in `spellconst/priest.json`, so the shield does scale with healing; Mental Strength "Increases your total Intellect by 10%"; Divine Aegis "Your non-periodic critical heals create a protective shield ... 15%", so "every direct heal that crits" is exact; Meditation "50% of your Mana regeneration to continue while casting". No ranking words.
- Holy (holy.md:88): Spiritual Guidance "spell healing by up to 25% of your total Spirit" at 5/5 (20% at the build's 4/5); Meditation 3/3 in the build; Inspiration "non-periodic critical heals increase your target's Armor by 25%"; no Divine Aegis or Renewed Hope in the Holy build, so a crit is a bigger heal plus the armor, as written. No ranking words.
- Shadow (shadow.md:99): 16% spell hit cap (`hit_to_cap` spell 16), Shadow Focus "5%"; DoT and Mind Flay crits (owner ruling forever-dots-crit, periodic crit flag on Mind Flay); Shadowform "increasing the critical strike damage bonus of your Shadow spells by 100%"; Intellect gives priest spell crit (round 3, gametable SpellCritPerIntellect); spell penetration "Not significant" in the table. No ranking words.
- `haste_on_items` false on every priest band; no prose mentions haste.

**Ranking by a nightly value elsewhere on the pages.** None. The only comparatives outside Stat priority are race lines from measured sims (round3/race-sims.md, held in round 4), "on a long fight the mana is worth more" (one potion cooldown, a rule), "a healer's cloth piece still loses to a real caster piece" (the third-as-damage rule), and holy.md:58 Twilight Focus "worth more than any throughput point" (a talent judgement, no stat or gear, nothing nightly). The Shadow frontmatter `statPriority` still lists Spell power before Hit and Discipline's lists Spirit second; it is only the stat SET once band 60 data loads (`railStatRows` sorts by scale factor; `[spec].astro:225`, `guide-file.ts:339` fall back to it only when no rows match), and the round 6 integrator left it on purpose. Not re-filed.

**Leveling, both blocks in page order.** Trainer levels re-read from `trainables/priest.json`: Smite 1, Lesser Heal 1/4/10, Shadow Word: Pain 4, Power Word: Shield 6, Renew 8, Fade 8, Mind Blast 10, Inner Fire 12, Psychic Scream 14, Heal 16/22/28/34, Dispel Magic 18, Holy Fire, Flash Heal, Devouring Plague, Fear Ward, Mind Flay rank 1 at 20, Penance 30/40/50/60, Prayer of Healing 30, Divine Spirit 30, Shadow Word: Death 32, Greater Heal 40/46/52/58/60, Lightwell 40, Prayer of Mending 40/50/60. Talent sentences against the orders above: Shadow "first five ... Shadow Focus ... next four into Blackout" then "From 19 the build takes Spirit Tap to 21" (no gap, no repeat); Discipline TD 10-14, MS 15-19, Penance 30; Holy TF 10-12, IR 13-15, HS from 16. The round 6 openers (holy.md:122, discipline.md:124) now read as a choice made at 20, after the 1 to 19 points, and match the strips. Holy "a level 30 priest who casts Heal rank 2 instead of rank 3": ranks 22 and 28, both inside the 17-level full-bonus window at 30. Wands and quests as round 5 verified (11287 req 5, 11288 req 13, Elven Wand from Bashal'Aran 957, Cinder Wand from Arugal's Folly 99).

**Everything else, line by line**, against `talents/priest.json` rank text, `spelltips/priest.json`, `spellconst/priest.json`, `races.json`, `trainables/priest.json` and `sim/request/raid-consumes.json`:
- Discipline tree: Power in Light "Smite and Penance" and Holy Precision "Holy spells" (shadow.md:52); Silent Resolve "threat generated by your spells by 30%" at 3/3, 10% a point, all spells, so it does cut Shadow threat (shadow.md:67, :91); Inner Focus 3 min, 25% on non-periodic spells.
- Shadow tree: Shadow Weaving 100%, 2%, 15 s, 5 stacks, "from you"; Improved Mind Flay 20% per instance; Spirit Tap 3/5 60%; Blackout 4/5 8%, 2/5 4%; IMB 0.5 s a point; ISWP 6 s; VE 20%, 30 s; Early Demise 30%; Darkness 10%; Shadowform 10%, half cost, +100% crit bonus.
- Spells: Shadow Word: Death "backlash damage equal to 30% of your maximum health", 18 s; Fade "for 10 sec", 30 s cooldown; Devouring Plague 24 s, heals the caster; Mind Blast "a high amount of threat"; Lightwell "3 min or 5 charges", "Being attacked cancels the effect"; Prayer of Healing "the target and their party"; Prayer of Mending "next time they take damage or receive non-periodic healing", "Jumps up to 5 times", 10 s; Power Word: Shield "spellcasting will not be interrupted by damage" (so "Power Word: Shield keeps you casting", holy.md:122, is the client text); Divine Grace free, below 50%, removes Weakened Soul, not on self, 10 min; Desperate Prayer free; Heal and Greater Heal 3.0 s casts, so 2.5 s with Divine Fury; Holy Fire 3.5 s.
- Racials and race spells: Escape Artist, Eureka!, Expansive Mind and Engineering Specialization for the Gnome (`races.json` text); Confounding Flash "up to 5 enemies within 8 yds for 3 sec"; Contingency Plan below 35%; Elune's Grace and Starshards for the Night Elf; Dark Sacrifice 5 ranks to 60.
- Consumables: every name in `items.json` (Greater Mageblood Elixir 250341, Greater Cleric's Elixir 250333, Elixir of Sages 250338, Sage's Tea 249870, Major Mender's 250949, Major Spellblasting 250937), and the band 55 healer raid kit carries flask, Greater Mageblood, Greater Cleric's and Elixir of Sages together, so the stacked list holds.
- Race prose against round3/race-sims.md: Shadow Human 690.7, Night Elf 690.4, Dwarf 689.6, Gnome 687.8 (±0.27), Undead 695.0, Troll 687.8; Discipline Gnome 600.6, Human 598.0, Night Elf 585.1, Dwarf 581.9, Undead 585.2, Troll 581.0; Holy Undead 837.2, Human 827.5, Gnome 821.6, the rest 804 to 808. Every "a hair behind", "close", "clearly less" and "of all" holds.
- No em or en dashes; no site or sim words outside the planner links (round 2 D2).

**Card notes.** No new false note. The healer rank-1 notes (Discipline "Lesser Heal rank 1 ... for very little mana", Holy "Heal is the filler: rank 1 in a raid") and the top-level "learned below level 20" summaries stay with the rank move (R3-D1, R4-D2). Shadow notes match the page.

**Source links.** Blizzard Deep Dive recap 200; Power Up Gaming 200; Talents Forever data.json 200 (labelled as beta client data); foreversixty.gg/sim/weights 200 (Shadow only, a damage spec); /planner and /planner?class=priest 200; Icy Veins talent calculator news 403 to automated fetch (unknown, not dead, R5-D2). Every entry fits the round 4 source-link rule.

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| PRIEST-R7-1 | LIKELY | meta (date) | index.md:5 | `updated: 2026-11-02` | `git log -- web/src/content/guides/priest/index.md`: content changed on 2026-11-03 in 384717fd4 (Shadowfiend line, round 2), c2cece8b7 (race picks split by spec from the measured sims, round 3) and 77bd7c12b (source label, round 4), and the stamp was never moved. The class page hides its header (`[class]/index.astro` passes `hideHeader`), so the date shows only in `article:modified_time` and the JSON-LD (`layouts/Content.astro:74, :83`), which is why it is not CONFIRMED as a reader-facing error. Every other priest page carries 2026-11-04. The same stale 2026-11-02 sits on the druid, mage, paladin, rogue, shaman, warrior and warlock class pages, and each of those files has a later commit (last commit 2026-11-03 or 2026-11-04 by `git log -1`; the integrator should confirm the content changed and not only a merge). | `updated: 2026-11-04` the next time index.md is touched (or now, since the race lines changed on 11-03). Integrator: check each class `index.md` against its `git log` the same way. |

Counts: CONFIRMED 0, LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0.

## Decisions (no owner questions)

- **R7-D1 Holy crit sentence stays** (holy.md:88, "A critical heal is a bigger heal, and for Holy that is all it is, since Inspiration only adds armor to the target"). It is true on 70755: the Holy build has no Divine Aegis, Renewed Hope or other crit-to-healing talent, and Inspiration's only effect is armor. It explains and does not rank. Not a finding.
- **R7-D2 Shadow and Discipline frontmatter `statPriority` order stays** (as the round 6 integrator decided): it is the set of stats the table shows, and the rendered order comes from the band 60 scale factors.
- **R7-D3 Card notes**: as R4-D2, R5-D3 and R6-D3. No change.

## Sims requested

None. The one finding is page metadata against git history.

## Sim findings (engine vs game)

None new. Still open from earlier rounds, not re-filed: Starshards cooldown and a Shadow line (S-R3-1/2); Power Infusion on yourself (S-R2-1); the Forever low-rank healing curve and the healer card ranks with their notes; Shadow leveling Mind Flay at 20 and Spirit Tap 5/5; the Shadow raid Blackout 2 + Spirit Tap 3 move; the SW:P cutoff `remainingTime >= 10` to `>= 6`; the healer haste weights (no priest item carries haste, so no prose depends on them).

## Sources

- Repo data (70755): `data/builds/1.60.1.70755/{talents,trainables,spelltips,spellconst}/priest.json`, `races.json`, `items.json`, `addon-builds.json`, `bis/priest-*.json`; `sim/request/raid-consumes.json`; `data/curated/apl/priest-*.json`; round3/race-sims.md
- Site code: `web/src/pages/guides/[class]/index.astro`, `web/src/pages/guides/[class]/[spec].astro:222-225`, `web/src/layouts/Content.astro`, `web/src/lib/guides/rail-stats.ts`, `web/src/lib/app-data/guide-file.ts:339`
- Commits: 394562f20 (round 6 fixes), 384717fd4, c2cece8b7, 77bd7c12b (index.md edits on 2026-11-03)
- Blizzard, World of Warcraft: Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200)
- Power Up Gaming, New WoW Forever talents: changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/ (200)
- Talents Forever data export: https://talentsforever.com/data.json (200)
- Forever Sixty, stat weights: https://foreversixty.gg/sim/weights (200); character planner: https://foreversixty.gg/planner and https://foreversixty.gg/planner?class=priest (200)
- Icy Veins, Forever talent calculator news: https://www.icy-veins.com/wow-forever/news/start-planning-your-wow-forever-character-with-our-new-talent-calculator/ (403, unknown)

## Fix log (round 7)

- PRIEST-R7-1: fixed, index.md:5 `updated: 2026-11-03`, the date of the last content change (`git log -1 --format=%cs` = 2026-11-03, 77bd7c12b, round 4 source label). Set to the content date per the integrator's instruction rather than the brief's blanket 2026-11-04, since this edit changes only the stamp.
- R7-D1, R7-D2, R7-D3: applied as written, no change.
- No other files touched. SIM findings: none.
