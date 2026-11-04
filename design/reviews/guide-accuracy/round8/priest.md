# Guide accuracy round 8: Priest

Audited on the round 8 audit tree (land/guides-r7, HEAD 09b98d9e4), client 1.60.1.70755, engine pin d1d9f2775 (`sim/enginever/version.go`).
Files: `web/src/content/guides/priest/index.md`, `discipline.md`, `holy.md`, `shadow.md`, the card notes in `data/curated/apl/priest-*.json`, and every frontmatter and body link. Read first: round 1 to round 7 priest reports (every fix log), round4/priest-levels.md, round2 to round7 race-sims.md and the ledger through round 7.

What moved since the round 7 audit tree (394562f20): `git diff --stat 394562f20..HEAD` over the priest guides, `data/curated/apl/priest-*`, `addon-builds.json`, `talents/`, `trainables/`, `spelltips/priest.json` and `bis/` shows one line: index.md:5 `updated: 2026-11-03` (the PRIEST-R7-1 fix). The Skyborne merge (715359838, 858394236) touches no priest page or priest data. Every FS1 string is unchanged since round 2.

## Decoded builds and point orders (re-derived this round; all legal on 70755)

FS1 digits decoded in `talents/priest.json` array order; `addon-builds.json` entries [tree, tier+1, column+1], index 0 = level 10.

- Discipline `00500503133001143/3350500303/0` (gnome), leveling = raid: Twin Disciplines 10-14, Mental Strength 15-19, Meditation 20, Mental Agility 21-23, Inner Focus 24, IPWS 25-27, Meditation 28-29, Penance 30, Soul Warding 31, Renewed Hope 32-35, Divine Aegis 36-38, Twilight Focus 39-41, Improved Renew 42-44, Holy Specialization 45-49, Divine Fury 50-54, Inspiration 55-57, Improved Healing 58-60. Tier gates: Penance (tier 4) at 30 with 20 spent, Divine Aegis (tier 5) at 36 with 26.
- Holy `00500503103/33505003030121431/0` (human), leveling = raid: Twilight Focus 10-12, Improved Renew 13-15, Holy Specialization 16-20, Divine Fury 21-25, Inspiration 26-28, Improved Healing 29-31, Holy Binding 32, Litany of Light 33-34, Spiritual Healing 35-37, Spirit of Redemption 38, Spiritual Guidance 39, Prayer of Mending 40 (30 Holy spent), Spiritual Guidance 41-43, Twin Disciplines 44-48, Mental Strength 49-53, Mental Agility 54-56, Inner Focus 57, Meditation 58-60.
- Shadow leveling `02510503103/0/54342001201300051` (human): Shadow Focus 10-14, Blackout 15-18, Spirit Tap 19-21, IMB 22-25, ISWP 26-27, Mind Flay 28, IMF 29-30, VE 31, Shadow Weaving 32-34, Darkness 35-39, Shadowform 40, Wand Specialization 41-42, Twin Disciplines 43-47, Silent Resolve 48, Mental Strength 49-53, Mental Agility 54-56, Inner Focus 57, Meditation 58-60.
- Shadow raid `00510503103/0/52352101201300251`: Shadow Focus 10-14, Blackout 15-16, Spirit Tap 17-19, IMB 20-24, ISWP 25-26, Shadow Reach 27, Mind Flay 28, IMF 29-30, VE 31, Shadow Weaving 32-34, Early Demise 35-36, Darkness 37-39, Shadowform 40, Darkness 41-42, then Discipline as leveling minus Wand Specialization. The trade at shadow.md:69 (-2 WS, -2 Blackout, +2 ED, +1 SR, +1 IMB) balances.
- Power Infusion swap (discipline.md:60): Discipline 29 + Renewed Hope 5th = 30 for tier 6; Holy keeps TF 3, IR 1, HS 5, DF 5, Insp 3 = 17 ahead of Improved Healing's tier 3 (15). Legal.

## Verified clean this round (no finding)

**Leveling, page order (1 to 19, then 20 to 60, strip under).** Trainer levels re-read from `trainables/priest.json`: Smite 1, Lesser Heal 1/4/10, Shadow Word: Pain 4, Power Word: Shield 6, Renew 8, Fade 8, Mind Blast 10, Inner Fire 12, Psychic Scream 14, Heal 16/22/28/34, Dispel Magic 18, Holy Fire, Flash Heal, Devouring Plague, Fear Ward at 20, Mind Flay talent ranks from 20. Every talent sentence matches the orders above: Shadow "first five ... Shadow Focus ... next four into Blackout", then "From 19 ... Spirit Tap to 21, Improved Mind Blast to 25 and Improved Shadow Word: Pain at 26 and 27, then Mind Flay at 28", "Improved Mind Flay, Vampiric Embrace, Shadow Weaving and Darkness follow, and Shadowform is your point at 40" (exact order); Discipline TD 10-14, MS 15-19, Penance 30; Holy TF 10-12, IR 13-15, HS from 16. Heal rank 2 (22) and rank 3 (28) are both inside the 17-level full-bonus window at 30 (holy.md:122). Wands and quests: Lesser Magic Wand 11287 required 5, Greater Magic Wand 11288 required 13 (`items.json`); Elven Wand 5604 alliance-only from Bashal'Aran (quest 957), Cinder Wand 5242 horde-only from Arugal's Folly (quest 99) (`loot.json` quests map). No contradiction between the two blocks or with the strips.

**Talent lines, rank by rank** against `talents/priest.json` 70755 text: Shadowform (10%, half cost on Shadow spells, +100% crit bonus, no healing spells), Shadow Weaving 3/3 100% "from you", 2%, 15 s, 5 stacks, Darkness 10%, IMF 20% per instance, IMB 0.5 s a point, ISWP 3 s a point, Shadow Focus 5%, Spirit Tap 3/5 60% with 50% regen while casting and VE deaths, Blackout 4/5 8% and 2/5 4% for 3 s, VE 20% for 30 s, Early Demise 2/2 30%, Silent Resolve 10% a point, Twin Disciplines "instant cast spells" 5%, Mental Agility "Smite, Holy Fire, and instant cast spells" 10%, Inner Focus (free, 25% on non-periodic), Meditation 50%, Mental Strength 10%, Penance (three heals, 1 s apart; 12 s on the trainer rows), Renewed Hope 4/5 8% and 4 s on Flash Heal, Lesser Heal, Heal, Greater Heal, Penance, Divine Aegis 3/3 15% of non-periodic crits, Soul Warding (removes the 4 s cooldown, 15%), IPWS 20%, Twilight Focus 70%, Improved Renew 15%, Holy Specialization 5% on Holy spells, Divine Fury 0.5 s on Heal and Greater Heal, Inspiration armor only, Improved Healing 3/3 15% on Lesser Heal, Heal, Greater Heal, Penance, Prayer of Mending, Holy Binding 50% on others, Litany of Light 10% "previous heal was a different spell", Spirit of Redemption 15 s of free heals, Spiritual Guidance 4/5 20% of Spirit, Spiritual Healing 3/3 10%, Prayer of Mending "Jumps up to 5 times", 10 s.

**Spell text** (`spelltips/priest.json`): Shadow Word: Death "If your target is not killed ... 30% of your maximum health", 18 s; Fade 10 s, 30 s cooldown; Devouring Plague 24 s cooldown and duration, heals the caster; Vampiric Embrace 30 s, cooldown 30 s; Mind Blast "a high amount of threat", 8 s; Mind Flay r6 3 s channel, 50% slow; SW:P r8 18 s; Lightwell "3 min or 5 charges", "Being attacked cancels the effect", 10 min; Prayer of Healing "heals the target and their party"; Power Word: Shield "spellcasting will not be interrupted by damage", Weakened Soul 15 s; Divine Grace below 50%, removes Weakened Soul, not on self, 10 min; Desperate Prayer free self heal; Dark Sacrifice 1600 health over 15 s for 1600 + Spirit mana, 10 min; Contingency Plan below 35%, 30 s ward; Confounding Flash 5 enemies, 8 yd, 3 s; Power Infusion 20% for 15 s, 3 min, no self-cast restriction; Heal and Greater Heal 3.0 s, so 2.5 s with Divine Fury; Holy Fire 3.5 s. Index:26 baseline list (Fear Ward, Devouring Plague, Shadow Word: Death, Divine Spirit, Lightwell) are all `skill_line_ability` rows; Holy Nova stays a Holy talent, which index.md does not claim otherwise.

**Races** against `races.json` 70755 text and round3/race-sims.md: Ruthless 1% spell and attack crit, The Human Spirit 10%, Elune's Light 10% crit 15 s / 3 min, Artisan's Precision 1% hit, Stoneform, Berserking 10% casting speed 15 s / 3 min, Touch of the Grave up to 5% max health once per 15 s, Will of the Forsaken removes Charm, Fear and Sleep, Eureka! 10% / 10% / 8 s / 2 min, Expansive Mind 5% mana, Engineering Specialization 20%. Race spells per index.md:42-46 match the race-locked rows. Every comparative ("a hair behind", "just as good", "close", "clearly less", "well ahead", "of all") holds on the round 3 numbers. Skyborne (`combos.json` races 95 and 96): warrior, hunter, rogue, mage, druid and warrior, hunter, rogue, shaman, druid, so index.md:48's list is the union and no priest row exists.

**Low ranks** (index.md:29, discipline.md:74, holy.md:72) against learn levels and the ForeverChanges curve: Lesser Heal r3 (10), Renew r1-r3, Heal r1-r2 keep nothing at 60; Heal r3 25%, Heal r4 55%, Renew r5 45%, r6 75%, r7 (44) full; Greater Heal r1 (40) 85%. "Next to nothing", "only part of it" and "from your forties on" hold.

**Stat prose** (all three) explains and never ranks; every reason is the 70755 text (Mental Strength 10%, Divine Aegis on non-periodic crits, Meditation 50%, Spiritual Guidance, Inspiration armor only, 16% spell hit cap with Shadow Focus 5%, DoT and Mind Flay crits, Shadowform crit bonus, Intellect spell crit). Healing-item data model (`items/priest.json`): "damage and healing" items carry both `spell_power` and `healing`, so "Spell power with no healing on it does nothing for a heal" is right; phase 1 healing items carry a third as spell damage (Hide of the Wild 42/14, Cauterizing Band 46/15, Devout set). No ranking word outside the tables except measured race lines and non-nightly rules.

**Consumables**: every name exists in `items.json`; flask plus three elixirs stack in the band 55 kits; Major Mender's and Major Spellblasting share the potion cooldown.

**Card notes**: no new false note. The healer rank-1 notes (Discipline "Lesser Heal rank 1 ... for very little mana", Holy "Heal is the filler: rank 1 in a raid") and the top-level "learned below level 20" summaries stay with the rank move (known and accepted). Shadow notes match the page.

**Round 7 fix re-checked**: index.md:5 `updated: 2026-11-03` is the date of the last content change (77bd7c12b); the round 7 commit changed only the stamp. Holds.

**Source links** (fetched 2026-11-04): Blizzard Deep Dive recap 200; Power Up Gaming 200; Talents Forever data.json 200 (labelled as beta client data); foreversixty.gg/sim/weights 200 (Shadow only, a damage spec); /planner and /planner?class=priest 200; Icy Veins talent calculator news 403 to automated fetch (unknown, not dead, R5-D2). Every entry fits the round 4 source-link rule. No em or en dashes; no site or sim words outside frontmatter `kind` and the planner links.

## Findings

None.

Counts: CONFIRMED 0, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## SHOULDs (wording, not findings)

| id | file:line | text | why | suggested wording |
|---|---|---|---|---|
| PRIEST-R8-S1 | holy.md:59 | "**Twin Disciplines** (5 points) and **Mental Agility** (3 points): Renew heals 5% more, and Renew and Power Word: Shield cost 10% less." | True, but it leaves out Holy's capstone. Prayer of Mending is an instant: Twin Disciplines reaches its heals (client heal-spell mask 0x200000 in eff0, engine `sim/priest/priest.go` PriestSpellMaskTwinDisciplines) and Mental Agility's client cost mask names it (PriestSpellMaskMentalAgility). A Holy player would say what the 8 points do for the spell cast every 10 seconds. | "**Twin Disciplines** (5 points) and **Mental Agility** (3 points): Renew and Prayer of Mending heal 5% more, and they cost 10% less, along with Power Word: Shield." |
| PRIEST-R8-S2 | holy.md:51 | "**Spiritual Healing** (3 points): 10% more healing on everything." | Lightwell's renew and Holy Binding are not in its mask (`priest.go` PriestSpellMaskSpiritualHealing). "Everything" is loose; every spell you cast is the true scope. | "**Spiritual Healing** (3 points): 10% more healing on every heal you cast." |
| PRIEST-R8-S3 | index.md:38, shadow.md:115 | "Eureka! (8 seconds of 10% cheaper and 10% stronger spells ...)" / "Eureka! makes your spells 10% cheaper and 10% stronger" | The client row covers non-periodic damaging and healing abilities (Wowhead 1259823), not Shadow Word: Pain, Devouring Plague, Renew or the shield. shadow.md:95 already says it skips the DoTs, so the page does not contradict itself, but the rogue page was tightened the same way in round 7. | index.md:38: "Eureka! (8 seconds of 10% cheaper and 10% stronger direct damage and heals, every 2 minutes)". shadow.md:115: "Eureka! makes Mind Blast, Shadow Word: Death and Mind Flay 10% cheaper and 10% stronger for 8 seconds every 2 minutes". |
| PRIEST-R8-S4 | holy.md:30 | "five jumps for one cast" | The client says "Jumps up to 5 times"; the engine and the card note count five heals. The prose matches the tooltip, so it stays unless the class wants one phrasing; "five heals for one cast" is the version a player feels. | "five heals for one cast" |

## Decisions (no owner questions)

- **R8-D1 No finding this round.** Every statement on the four priest pages is true on 70755 and none contradicts its page. The four SHOULDs are wording; the integrator may apply them without a re-audit (each is narrower than the text it replaces).
- **R8-D2 Shadow and Discipline frontmatter `statPriority`** stay as R6/R7 decided (the rendered order comes from band 60 scale factors).
- **R8-D3 Card notes**: as R4-D2 to R7-D3. No change.
- **R8-D4 Discipline 1 to 19 Twin Disciplines line** (discipline.md:118, "which at this level means Shadow Word: Pain and Renew"): a Human's Divine Grace and a Dwarf's Desperate Prayer (both level 10, both in the Twin Disciplines mask) would also count, but the page's races are Gnome and Undead, for whom the sentence is exact. Not filed.

## Sims requested

None. Race re-runs on the re-ranked gear stay on the ledger's open list (round 7), and nothing on the priest pages depends on them beyond "close" wording that the round 3 numbers support.

## Sim findings (engine vs game)

None new. Still open, not re-filed: Starshards cooldown and a Shadow line (S-R3-1/2); Power Infusion on yourself (S-R2-1); the Forever low-rank healing curve and the healer card ranks with their notes; Shadow leveling Mind Flay at 20 and Spirit Tap 5/5; the Shadow raid Blackout 2 + Spirit Tap 3 move; the SW:P cutoff `remainingTime >= 10` to `>= 6`; healer haste weights (no priest item carries haste).

## Sources

- Repo data (70755): `data/builds/1.60.1.70755/{talents,trainables,spelltips}/priest.json`, `items/priest.json`, `items.json`, `loot.json`, `races.json`, `combos.json`, `addon-builds.json`; `data/curated/apl/priest-*.json`; round3/race-sims.md
- Engine d1d9f2775: `sim/priest/priest.go` (Twin Disciplines, Mental Agility and Spiritual Healing masks), `sim/priest/prayer_of_mending.go` (five charges)
- Wowhead Forever tooltip, Eureka! (priest) 1259823: https://nether.wowhead.com/forever/tooltip/spell/1259823
- ForeverChanges, WoW Forever downranking: https://foreverchanges.pro/downranking
- Blizzard, World of Warcraft: Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200)
- Power Up Gaming, New WoW Forever talents: changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/ (200)
- Talents Forever data export: https://talentsforever.com/data.json (200)
- Forever Sixty, stat weights: https://foreversixty.gg/sim/weights (200); character planner: https://foreversixty.gg/planner and https://foreversixty.gg/planner?class=priest (200)
- Icy Veins, Forever talent calculator news: https://www.icy-veins.com/wow-forever/news/start-planning-your-wow-forever-character-with-our-new-talent-calculator/ (403, unknown)

## Fix log (round 8)

- PRIEST-R8-S3 applied. index.md:38: "10% cheaper and 10% stronger heals and damage spells every 2 minutes, never a DoT, Renew or a shield" (instead of "direct damage and heals", since the client mask also covers Mind Flay). shadow.md:115: "makes Mind Blast, Shadow Word: Death and Mind Flay 10% cheaper and 10% stronger". Checked against the engine's priest row (eureka_auto_gen.go 1259823): Mind Blast, Shadow Word: Death and Mind Flay in the damage mask; Shadow Word: Pain, Devouring Plague, Renew and Power Word: Shield in none.
- S1, S2, S4 left: not in this pass's scope. Prose check: priest clean. No new SIM findings.
