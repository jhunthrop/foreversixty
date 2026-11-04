# Guide accuracy round 2: Priest

Audited on land/guides-r1 head 609eb471c, client 1.60.1.70755, engine pin d1d9f2775.
Files: `web/src/content/guides/priest/index.md`, `discipline.md`, `holy.md`, `shadow.md`, plus the card notes in `data/curated/apl/priest-*.json`.

## Decoded builds (FS1 digit order = array order of `data/builds/1.60.1.70755/talents/priest.json`; all legal on 70755, 51 points, every tier gate and prerequisite met)

**Discipline `build:`** `00500503133001143/3350500303/0` (29/22/0)
- Discipline: Twin Disciplines 5, Mental Strength 5, Mental Agility 3, Inner Focus 1, Improved Power Word: Shield 3, Meditation 3, Soul Warding 1 (needs IPWS 3: met), Penance 1, Renewed Hope 4 (needs Soul Warding: met), Divine Aegis 3
- Holy: Twilight Focus 3, Improved Renew 3, Holy Specialization 5, Divine Fury 5, Inspiration 3, Improved Healing 3
- Leveling order (`addon-builds.json` priest-discipline): Twin Disciplines 10-14, Mental Strength 15-19, Meditation 20, Mental Agility 21-23, Inner Focus 24, IPWS 25-27, Meditation 28-29, Penance 30, Soul Warding 31, Renewed Hope 32-35, Divine Aegis 36-38, then Holy. Prose "Penance comes at level 30" matches.
- Power Infusion swap (Improved Renew -2, Renewed Hope +1, Power Infusion +1) is legal: Disc reaches 30 before tier 6, Penance is PI's prerequisite, Holy keeps 17 points before Improved Healing.

**Holy `build:`** `00500503103/33505003030121431/0` (17/34/0)
- Discipline: Twin Disciplines 5, Mental Strength 5, Mental Agility 3, Inner Focus 1, Meditation 3
- Holy: Twilight Focus 3, Improved Renew 3, Holy Specialization 5, Divine Fury 5, Inspiration 3, Improved Healing 3, Holy Binding 1, Litany of Light 2, Spirit of Redemption 1, Spiritual Guidance 4, Spiritual Healing 3, Prayer of Mending 1 (needs Spirit of Redemption: met)
- Leveling order: Holy Binding 32, Litany 33-34, Spiritual Healing 35-37, Spirit of Redemption 38, Prayer of Mending 40, Discipline from 44.

**Shadow `build:` (leveling)** `02510503103/0/54342001201300051` (20/0/31)
- Discipline: Wand Specialization 2, Twin Disciplines 5, Silent Resolve 1, Mental Strength 5, Mental Agility 3, Inner Focus 1, Meditation 3
- Shadow: Shadow Focus 5, Blackout 4, Spirit Tap 3, Improved Mind Blast 4, Improved Shadow Word: Pain 2, Mind Flay 1, Improved Mind Flay 2, Vampiric Embrace 1, Shadow Weaving 3, Darkness 5, Shadowform 1
- Order: Shadow Focus 10-14, Blackout 15-18, Spirit Tap 19-21, IMB 22-25, ISWP 26-27, Mind Flay 28, IMF 29-30, VE 31, Shadow Weaving 32-34, Darkness 35-39, Shadowform 40, Wand Specialization 41-42, Discipline 43-60. shadow.md:34, :41, :125 and :69 match it.

**Shadow `raidBuild:`** `00510503103/0/52352101201300251` (18/0/33)
- Discipline: Twin Disciplines 5, Silent Resolve 1, Mental Strength 5, Mental Agility 3, Inner Focus 1, Meditation 3
- Shadow: Shadow Focus 5, Blackout 2, Spirit Tap 3, Improved Mind Blast 5, ISWP 2, Shadow Reach 1, Mind Flay 1, Improved Mind Flay 2, VE 1, Shadow Weaving 3, Early Demise 2, Darkness 5, Shadowform 1
- shadow.md:69 "trades [Wand Specialization 2] and two points of Blackout for Early Demise, Shadow Reach and the last point of Improved Mind Blast" is exactly the diff (-2 WS, -2 Blackout, +2 ED, +1 SR, +1 IMB).

## Verified clean this round (no finding)

- Every talent line on the three spec pages against the 70755 talent text at the build's rank: Shadowform, Shadow Weaving ("from you"), Darkness, Improved Mind Flay (20% per tick), Improved Mind Blast (0.5 s a point), ISWP (+6 s), Shadow Focus, Spirit Tap 3/5 (60%, 50% regen while casting), Blackout 4/2 (8%/4%, 3 s), VE (20%, 30 s), Early Demise (30%), Silent Resolve (10% a point), Penance (three heals, 1 s apart, 12 s cd), Renewed Hope 4/5 (8%, 4 s), Divine Aegis (15%), Soul Warding (removes PW:S's 4 s cooldown, 15% cost), IPWS (20%), Improved Healing (Lesser Heal, Heal, Greater Heal, Penance, Prayer of Mending), Divine Fury (Heal, Greater Heal to 2.5 s), Twilight Focus (70%), Improved Renew (15%), Holy Binding (50%, others only), Spiritual Healing (10%), Prayer of Mending (next hit or direct heal, jumps up to 5 times, 10 s cd).
- Spell descriptions (`spelltips/priest.json`, per the round 2 lesson): PW:S Weakened Soul 15 s; Lightwell 3 min or 5 charges, being attacked cancels, 10 min cd; Prayer of Healing "heals the target and their party"; Divine Grace below 50%, removes Weakened Soul, cannot be cast on self, 10 min; Contingency Plan under 35%, 30 s ward; Confounding Flash up to 5 enemies, 8 yd, 3 s; Devouring Plague 24 s, heals the caster; SW:D 30% backlash, 18 s; Fade 30 s cd; Inner Focus 3 min, non-periodic crit only (so "do not let it fall on a Renew" is right); Shadowform 40% base mana (cost_pct 40, so "the form costs mana" is right); Mind Blast "causes a high amount of threat"; Dark Sacrifice 1600 health for 1600 + Spirit mana, 10 min; Desperate Prayer free.
- Low-rank arithmetic on the shipped prose (index.md:29, discipline.md:77, holy.md:74) against the 70755 learn levels and the ForeverChanges curve: Renew r1-r3 and Heal r1-r2 and Lesser Heal keep nothing at 60, Renew r4 15%, r5 45%, r6 75%, Heal r4 55%, Greater Heal r1 85%, r2 and Renew r7 full. "Renew rank 5 and 6 and Heal rank 4 keep only part of it" and "Greater Heal rank 1 or 2" hold.
- Racials against `races.json`: The Human Spirit, Ruthless, Artisan's Precision (spell hit too), Stoneform, Elune's Light 10% / 15 s / 3 min, Berserking 10% casting speed, Eureka! 10% / 8 s / 2 min, Expansive Mind 5%, Engineering Specialization, Touch of the Grave, Will of the Forsaken (removes fear).
- Consumables: every name exists in 70755 `items.json`, including Forever's Elixir of Sages (250338, req 55, 25 Spirit and 2% crit per the engine's client read) beside Era's Elixir of the Sages; Major Mender's (88 healing) and Major Spellblasting (47 spell damage) are 30 s potions in the potion category, so "shares the potion cooldown with your mana potion" is right.
- Stat prose and frontmatter against band 60 raid scale factors (`bis/priest-*.json`): Disc mp5 2.52 > Spirit 1.30 > healing 1.17 > Int 1.00 > crit 0.67; Holy mp5 2.73 > healing 1.69 > Spirit 1.45 > Int > crit; Shadow spell power 5.75 > hit 4.61 > haste 4.44 > crit 3.83 > Int. No quoted numbers.
- Round 1 research answers re-checked against their sources and the client text: ForeverChanges' downranking page states exactly the rule round 1 quotes ("For 17 levels after you learn a rank, it gets all of its bonus", "loses 5% of its full bonus at each level", "37 levels after you learn it, only the rank's own heal or damage is left", measured on beta 1.60.1.70532 up to 30). SW:P cutoff: 70755 SW:P r8 description "762 Shadow damage over 18 sec" is 127 a tick every 3 s; Mind Flay r6 "390 Shadow damage over 3 sec" is 130 a tick, so the two-tick (6 s) answer stands. Spirit Tap 5/5: rank 5 description "100% chance". Devouring Contagion: "Reduces the mana cost of your Devouring Plague by 50%" and spreads on death, as round 1 said. Mind Flay legal at 20: tier 2 needs 10 points, true on 70755.
- No em/en dashes, no site/sim words in the prose outside the planner links (see Decisions).

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| PRIEST-R2-1 | CONFIRMED | stale / voice | index.md:27; shadow.md:41 | "**Mind Flay** is a Shadow talent instead of a trainer spell" (under "What changed from 1.12"); "Mind Flay is a talent in Forever" | Mind Flay has always been a Shadow talent: Era client `data/builds/1.15.9.69722/talents/priest.json` has Mind Flay in the Shadow tree at tier 2, exactly where 70755 has it. Nothing changed, so it does not belong in a "what changed" list, and the Shadow line sells an old fact as a Forever one. | index.md:27: "**Penance** and **Prayer of Mending** are the deep Discipline and Holy talents, and the trainer teaches their higher ranks once you have them." shadow.md:41: "Mind Flay is a talent and the leveling build reaches it at 28, so until then you wand between casts" |
| PRIEST-R2-2 | CONFIRMED | talent | discipline.md:63 | "**Power Infusion** is not in the build because it adds nothing to your own healing." | 70755 description (`spelltips/priest.json`, 10060): "Infuses the target with power, increasing their spell damage and healing done by 20% for 15 sec." No "cannot be cast on self" clause, while Divine Grace's description carries one; Forever Discipline macros self-cast it (Rankedboost, Forever Discipline macros: "hold Ctrl to keep it for yourself"). On yourself it is 20% more healing for 15 s every 3 min, which is not nothing. | "**Power Infusion** is not in the build. On yourself it is 15 seconds of 20% more healing every 3 minutes for a fifth of your base mana; on a mage or warlock's burst it is worth far more to the raid. When you have a caster worth buffing, take it: move two points out of Improved Renew, one into Renewed Hope's fifth rank and one into Power Infusion, and give it to them on their burst." (Whether self-cast PI beats the two Improved Renew points is S-R2-1.) |
| PRIEST-R2-3 | LIKELY | rotation | holy.md:59 | "Litany of Light (2 points): mana back whenever your heal is a different spell from the last one. Weaving Heal, Flash Heal and Prayer of Mending earns it all fight." | Talent text: "gain Mana equal to 10% of the base cost of the spell if your previous heal was a different spell". The same page makes downranked Greater Heal the raid filler (holy.md:29, :72) and Flash Heal emergencies only (holy.md:76), so the spells a Holy priest actually weaves are Greater Heal, Renew and Prayer of Mending; two Greater Heals of different ranks are still the same spell. | "**Litany of Light** (2 points): mana back whenever your heal is a different spell from the last one. Put a Renew or a Prayer of Mending between your Greater Heals and it pays all fight; two Greater Heals in a row, even at different ranks, earn nothing." |
| PRIEST-R2-4 | LIKELY | voice | discipline.md:87, holy.md:40, shadow.md:93 | "Shadowfiend cannot be learned at launch." / "you have no Shadowfiend at launch" | A Classic priest never had Shadowfiend (it is a Burning Crusade spell and a Season of Discovery rune), so saying it is missing on three spec pages reads as a data note, not a player's advice. The class page already says it once (index.md:28). Decision D1 below settles the fact. | discipline.md:87 and shadow.md:93: delete the sentence. holy.md:40: "Mana is the limit on every long fight" (drop the Shadowfiend clause). Keep index.md:28 as D1 words it. |

Counts: CONFIRMED 2, LIKELY 2, NEEDS-SIM 1 (S-R2-1), NEEDS-OWNER 0.

## Decisions (owner: nothing goes back as a taste call)

- **D1 Shadowfiend is not learnable at launch. Keep the line on the class page only.** Reason: on the 70755 client, SkillLineAbility row 49886 (SkillLine 78 Shadow Magic, Spell 401977, AcquireMethod 0) exists, but SpellLevels has no row for 401977, no talent in `talents/priest.json` teaches it, the trainer export gives it only a rank-0 level-1 placeholder, and the 70755 hotfix push does not mention it. Its only learning path in this client is the Season of Discovery rune system (Engrave Pants - Shadowfiend 402861 on the Runes skill line 2851), which Forever does not have. Penance and Prayer of Mending, which also have no SpellLevels row, are taught by talents; Shadowfiend has no such source. No Forever source (Blizzard dev notes, the Priest Deep Dive, Warcraft Tavern, Mobalytics, wow.gg, Icy Veins) lists it. Guide change: index.md:28 becomes "**Shadowfiend**, the Season of Discovery rune, is not in Forever." and the three spec mentions go (PRIEST-R2-4). The APL notes already state it and do not cast it.
- **D2 Planner links stay, in one form, for every class.** Reason: the owner's one-job-per-page rule says "the planner owns talents ... A page links to the owner of a fact instead of restating it", so a guide pointing to the planner for build changes is what the owner asked for; the guide-voice rule bans site and sim talk inside the advice (how a number was found, "our search", "this site"), not navigation, and the prose test deliberately skips link targets so planner links pass. One form avoids three phrasings: a link that names the action, "[Open the planner](/planner?class=<class>) to change the build." (the form the rogue pages use). Priest changes: discipline.md:63 last sentence "[Open the planner](/planner?class=priest) to make the swap."; holy.md:64 "Spirit of Redemption is only the path to Prayer of Mending. [Open the planner](/planner?class=priest) to change the build."; shadow.md:69 last sentence "[Open the planner](/planner?class=priest) to change either build." Integrator: apply the same link text to mage (arcane, fire, frost) and warrior (arms, fury, protection) where the link sits mid-sentence on the word "planner".

## Sim findings (next sim round)

- **S-R2-1 Power Infusion on yourself.** Question for a Discipline heal-builds run at 60 raid: does the PI swap (Improved Renew 3 to 1, Renewed Hope 4 to 5, Power Infusion 1, cast on self on cooldown) beat the shipped build beyond error? If it does, the build takes PI and discipline.md:63 says so.
- Still open from round 1, not re-filed: engine low-rank healing curve and the healer card ranks (Disc and Holy Renew r1, Disc Lesser Heal r1, Holy Heal r1 filler; their card notes describe those casts); Shadow leveling Mind Flay at 20 and Spirit Tap 5/5; Shadow raid Blackout 2 + Spirit Tap 3 move; SW:P cutoff `remainingTime >= 10` to `>= 6`; healer haste weights negative.

## Owner questions

None. Both questions round 1 left for the owner are decided above (D1, D2).

## Sources

- 70755 client tables on wago.tools: SkillLineAbility https://wago.tools/db2/SkillLineAbility/csv?build=1.60.1.70755 and SpellLevels https://wago.tools/db2/SpellLevels/csv?build=1.60.1.70755 (Shadowfiend 401977, row 49886; no SpellLevels row)
- ForeverChanges, "WoW Forever downranking: the level penalty and which ranks to keep": https://foreverchanges.pro/downranking
- Rankedboost, WoW Forever Discipline Priest macros (Power Infusion self-cast): https://rankedboost.com/world-of-warcraft/forever/macros/priest/discipline/
- Blizzard, World of Warcraft: Forever Class Deep Dives, Priest and Warrior: https://news.blizzard.com/en-us/article/24301514/world-of-warcraft-forever-class-deep-dives-priest-and-warrior
- Warcraft Tavern, Forever Priest Guide: https://www.warcrafttavern.com/forever/guides/priest/ ; Mobalytics, Forever Shadow Priest guide: https://mobalytics.gg/wow-forever/classes/shadow-priest-guide ; Icy Veins, Forever Shadow Priest guide: https://www.icy-veins.com/wow-forever/shadow-priest-ranged-dps-pve-guide (none mention Shadowfiend)
- Owner rules: memory feedback-one-job-per-page.md, feedback-guide-voice.md; `web/src/content/guides/_prose.test.ts` (link targets skipped)
- Repo data: `data/builds/1.60.1.70755/{talents,trainables,spelltips}/priest.json`, `races.json`, `items.json`, `simconsumes.json`, `addon-builds.json`, `bis/priest-*.json`; `data/builds/1.15.9.69722/talents/priest.json`; engine `sim/core/consumes.go` at d1d9f2775

## Fix log (round 2)

- PRIEST-R2-1 fixed: index.md:27 now lists only Penance and Prayer of Mending as the deep talents (Mind Flay dropped from "What changed"); shadow.md:41 "Mind Flay is a talent and the leveling build reaches it at 28".
- PRIEST-R2-2 fixed (true part only, per integrator): discipline.md:63 states PI works on yourself for 15 s of 20% more healing every 3 min, keeps the caster swap advice; no claim on which build is better (S-R2-1 open).
- PRIEST-R2-3 fixed: holy.md:59 Litany of Light now says weave a Renew or Prayer of Mending between Greater Heals; same spell at another rank earns nothing (talent text "previous heal was a different spell").
- PRIEST-R2-4 / D1 fixed: index.md:28 "**Shadowfiend**, the Season of Discovery rune, is not in Forever."; Shadowfiend sentences removed from discipline.md:87, shadow.md:93, holy.md:40. APL notes already true (not cast, not learnable); unchanged.
- D2 fixed: planner links in the one form at discipline.md:63 ("to make the swap"), holy.md:64 ("to change the build"), shadow.md:69 ("to change either build"), each its own sentence.
- Source labels: discipline/holy/shadow frontmatter now "Forever Sixty, stat weights" and "Forever Sixty, character planner" (hunter's form).
- No build, race, card action or leveling-order changes. No new SIM findings; S-R2-1 and the round 1 carry-overs stand.
