# Guide accuracy round 4: Hunter

Auditor: hunter (class index, Beast Mastery, Marksmanship, Survival). Tree: land/guides-r4 752c02e43, client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go).
Files: `web/src/content/guides/hunter/{index,beast-mastery,marksmanship,survival}.md`, card notes in `data/curated/apl/hunter-*.json`. Read first: round1/2/3 hunter reports, round2/race-sims.md, round3/race-sims.md (no hunter rows), ledger.md, round4/hunter-levels.md.

## Builds and leveling order

Build strings and statPriority are the same as round 3 (`git diff 6595e62d7 752c02e43` touches only the Leveling sections of the three spec pages). The decodes in round3/hunter.md still hold: all legal, 51 points.

Leveling order (`addon-builds.json`, index 0 = level 10), levels 10 to 21:
- BM: Deadly Aspects 10-14, Focused Fire 15-16, Endurance Training 17-20, Bestial Swiftness 21.
- MM: Lethal Attacks 10-14, Improved Stings 15, Efficiency 16-20, Careful Aim 21.
- SV: Improved Tracking 10-14, Savage Strikes 15-16, Survivalist 17-19, Clever Traps 20-21.

## New "Levels 1 to 19" subsections: what was checked and holds

- Every level against `trainables/hunter.json`: Raptor Strike 1 (ranks 8, 16), Auto Shot 1, Track Beasts 1, Serpent Sting 4, Aspect of the Monkey 4, Arcane Shot 6, Hunter's Mark 6, Concussive Shot 8, Aspect of the Hawk 10, Tame Beast 10, Track Humanoids 10, Mend Pet 12, Wing Clip 12, Mongoose Bite 16, Multi-Shot 18, Track Undead 18. All right.
- Talent points against the shipped order and the 70755 rank text. BM: Deadly Aspects 2% per rank, Auto Shot under Hawk, 30% for 12 sec. Focused Fire 1%/2% for you and the pet; its row opens at 5 points, so 15 is legal. Endurance Training from 17. MM: Lethal Attacks 1% crit per rank. Improved Stings rank 1 is 6% Serpent Sting. Efficiency 3% per rank (shots, stings and melee abilities). SV: Improved Tracking 1% per rank against the seven trackable types. Savage Strikes 2% melee ability crit per rank. Survivalist 2% health per rank. All match.
- Mongoose Bite "Can only be performed after you dodge" (spelltips 1495 and the Wowhead Forever tooltip). Raptor Strike "melee weapon damage plus N" (2973 to 14266), so the slow two-hander advice before Dual Wield at 20 holds. Feed Pet "Feeding your pet increases happiness" (6991), so "feed it to keep it happy" holds.
- Taming quests at 10 that end by teaching Tame Beast. Warcraft Tavern's Forever pet guide says "Once you have completed all three quests, you will learn Tame Beast", starting from the trainer at 10. Holds.
- No contradiction with "Levels 20 to 60" or the strip. The BM "Tame your first pet at 10" and Deadly Aspects/Focused Fire lines moved up and are not repeated. Survival's Monkey-until-30 line below agrees with Monkey at 4 above. Mail at 40 and leather until 40 agree everywhere.
- Voice: no em or en dashes, no site/sim words, no quoted nightly numbers.

## Rest of the pages: re-checked and holds

Every talent value in the prose against the 70755 rank text: Unleashed Fury "pets and hawks by 15%", Ferocity 10%, Frenzy 100%/30%/8 sec, Bestial Discipline (pet focus plus 50% mana regen while casting), Bestial Swiftness 30% pet speed, Rapid Killing (1 min and 10% at rank 1), Mortal Shots "ranged abilities" 30%, Lone Wolf 20%, Resourcefulness 60% plus the crit regen, and Expose Prey on "melee and ranged attacks against enemies afflicted by Hunter's Mark". Disengage "Disables auto-attack." (781). Counterattack is a talent at 30 (19306). Strider Kick is 100% weapon damage, 5.8% base mana, 8 sec. Heart of the Lion is level 1 and 280 attack power at 60 (Wowhead Forever tooltip 409580: "Requires level 1", `ppl0:60:40:400` gives 280). The round 3 fixes (Skyborne paid wording, Disengage, the Expose Prey card note, Monkey until 30) are correct. Race prose still matches the guide-build reruns in round2/race-sims.md.

## Source links

Resolved (gr4/links.txt plus direct fetches): all foreversixty.gg data and site links 200. Blizzard Deep Dive panel recap 200. Blizzard What's Next recap 200; it says Skyborne "can also play Warriors, Hunters, Rogues, and Druids", which supports index.md. Warcraft Tavern Forever Hunter Guide 200. Warcraft Tavern Hunter and Druid Deep Dive 200. Wowhead Skyborne First Look 200. Wowhead spell=1495 returns 202 to fetchers (a bot check, not dead); the nether tooltip endpoint returns Mongoose Bite.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| HUNTER-R4-1 | CONFIRMED | card note | data/curated/apl/hunter-survival.json, Wing Clip entry notes | "Without the totem it is 80 mana for 50 damage, so leave it alone." | Wing Clip rank 3 (14268) costs 80 base mana, but this is the Survival card, and the Survival build has Resourcefulness 2/2: "Reduces the Mana cost of your Trap and melee abilities by 60%" (talents/hunter.json). Wing Clip is a melee ability (5 yd), so it costs 32. The same card's Raptor Strike note already says "Resourcefulness takes 60% off its cost". The note contradicts the card. | "Wing Clip on any free global under a Windfury Totem, as a cheap melee special that can proc Windfury. It is a main hand hit that also rolls your weapon procs. Without the totem it is 50 damage for a global, so leave it alone." |
| HUNTER-R4-2 | CONFIRMED | source link | survival.md:19-20 | label 'Wowhead Forever, Mongoose Bite and Counterattack tooltips', url https://www.wowhead.com/forever/spell=1495 | The link is the Mongoose Bite page only (tooltip: "Counterattack the enemy for melee weapon damage plus 15. Can only be performed after you dodge."). The word "Counterattack" appears only as Mongoose Bite's verb. Counterattack the talent is spell 19306, which has its own page. The label promises a tooltip the link does not give. | Relabel to 'Wowhead Forever, Mongoose Bite' and add a second source: label 'Wowhead Forever, Counterattack', url https://www.wowhead.com/forever/spell=19306, kind community. |
| HUNTER-R4-3 | LIKELY | leveling | beast-mastery.md:137, marksmanship.md:140, survival.md:124 | The Levels 1 to 19 bullets never name Heart of the Lion. The lane notes (round4/hunter-levels.md) left it out as "SoD rune spell ids, not trainer spells". | trainables/hunter.json has Heart of the Lion 409580 as `skill_line_ability`, active, level 1. The Wowhead Forever tooltip says "Requires level 1". index.md:31 already says "a hunter spell from level 1". The notes' reason is wrong. 10% more stats for the group plus another 10% for you is the biggest buff a level 1 hunter owns, and a day-1 player needs to hear it first. | Add to the first bullet of each spec, before "Train": "Cast Heart of the Lion at level 1 and keep it up all the way to 60: 10% more stats and some attack power for everyone near you, and another 10% stats for you." |
| HUNTER-R4-4 | LIKELY | pet | index.md:29, beast-mastery.md:54,95, marksmanship.md:79, survival.md:65 | "it gets no buffs from your raid" / "The pet gets no raid buffs in Forever, so only your own gear makes it stronger" / "the only way to make it stronger is your own gear" | Blizzard's Forever Hunter and Druid class deep dive says: "Pets can no longer receive player buffs that increase their stats." That rules out stat buffs (Mark of the Wild, Kings, Fortitude), not every buff. Boss debuffs (Sunder Armor, Faerie Fire, Curse of Recklessness) still raise what the pet hits for, so "only your own gear" overstates it. The engine refuses all player buffs (sim/core/pet.go RefusePlayerBuffs), and no Forever source confirms that breadth. | index.md:29: "...and it takes no stat buffs from other players." BM:54: "The pet takes no stat buffs from other players in Forever, so your gear is what makes it stronger." BM:95: "Other players' stat buffs do not reach your pet in Forever, so your own gear is what makes it stronger." MM:79 and SV:65: "...and take no stat buffs from other players...". If the sim round shows that non-stat buffs (Windfury, Trueshot Aura, Leader of the Pack) do reach pets in Forever, add them to the engine (see Sim findings). |
| HUNTER-R4-5 | LIKELY | source link | beast-mastery.md:14-16, marksmanship.md:14-16, survival.md:13-15 | The only Blizzard source on the spec pages is 'Blizzard, Deep Dive panel recap' | That recap names Hunters only in passing: Human Hunter and level-up weapons. It says nothing about pets, Survival melee, Aimed Shot, Lone Wolf or aspects. Blizzard's own hunter article is "World of Warcraft Forever Class Deep Dives: Hunter and Druid" (https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid, 200). It is the primary source for Aimed Shot at 20 sharing Multi-Shot's cooldown, Lone Wolf, Careful Aim, Predator's Edge, Resourcefulness, Aspect of the Beast's melee AP and the pet stat-buff rule. The druid pages already cite it. The link resolves and matches its label, so this is not a dead link. | Add to all three spec pages (and index.md): label 'Blizzard, Class Deep Dives: Hunter and Druid', url https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid, kind blizzard. Keep the panel recap on index.md for the Human Hunter line. |

## Decisions

- **D1. The round 4 lane notes on Heart of the Lion are wrong** (409580 is a level 1 skill-line spell in Forever, not an SoD rune row). Only the notes say otherwise; the guide never states it, so it is HUNTER-R4-3, an omission, not a false-claim finding.
- **D2. Aspect of the Monkey dodge.** Blizzard's deep dive says Monkey "now grants twice as much Dodge", but the 70755 client text (13163) still says 8%. The guides quote no number ("its dodges open Mongoose Bite"), so nothing changes. The client description wins on any future number.
- **D3. Savage Strikes.** Blizzard's deep dive says "all melee damage by 2/4%", but the 70755 rank text says melee ability crit 2%/4%. The guides follow the client and stay as they are.
- **D4. Tame Beast by quest.** Kept. Warcraft Tavern's Forever pet guide confirms the three-quest chain from the trainer at 10.

## Sims requested

1. **Carried, still unrun: Sniper Shot** (round 3 request 1). MM raidBuild `5/005355001150305/50005003` against `5/0053550011503051/50005002` (Sniper Shot 1 for Surefooted 1), Night Elf, band 60 raid preset, seeds 1|50001, with Sniper Shot cast when the shared cooldown is down and the next auto shot is at least 4 s away. Adopt only beyond error.
2. **Carried, still unrun: Summon Hawk** (round 3 request 2). BM raidBuild with Summon Hawk 1 for the fourth Endurance Training point (`5320001505101251/00523500014/0`), cast on cooldown in Arcane Shot's place. Adopt only beyond error.

## Sim findings

- **Pet buffs (HUNTER-R4-4).** sim/core/pet.go `RefusePlayerBuffs` blocks every player buff on the hunter pet. Blizzard's wording is "player buffs that increase their stats". If Forever lets non-stat buffs (Windfury Totem, Leader of the Pack, Trueshot Aura, Battle Shout if it counts as non-stat) reach pets, the engine undervalues BM's pet in a raid. Check in game or on beta logs before changing anything.
- Carried unchanged: stale spelltips hover text (data lane), bis files pre-push (refresh), APL top-level Falcon notes (notes only).

## Sources

- Blizzard, World of Warcraft Forever Class Deep Dives: Hunter and Druid, https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid ("Pets can no longer receive player buffs that increase their stats"; Aimed Shot baseline at 20; Aspect of the Monkey; Savage Strikes)
- Blizzard, Deep Dive panel recap, https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (Hunters only in passing; Human Hunter)
- Blizzard, What's Next panel recap, https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap (Skyborne can play Hunters)
- Warcraft Tavern, Hunter Pets Guide (Forever), https://www.warcrafttavern.com/forever/guides/hunter-pets/ (three taming quests from the trainer at 10 teach Tame Beast; happiness and loyalty remain)
- Wowhead Forever tooltips: Mongoose Bite https://nether.wowhead.com/forever/tooltip/spell/1495, Counterattack https://nether.wowhead.com/forever/tooltip/spell/19306, Tame Beast https://nether.wowhead.com/forever/tooltip/spell/1515, Heart of the Lion https://nether.wowhead.com/forever/tooltip/spell/409580 ("Requires level 1")
- Forever Sixty client data 1.60.1.70755: trainables/hunter.json, talents/hunter.json, spelltips/hunter.json (13163, 6991, 2973-14266, 14268, 1317257, 781, 409580), spellconst/hunter.json, addon-builds.json
- Engine d1d9f2775: sim/hunter/pet.go:103, sim/core/pet.go:139 (RefusePlayerBuffs)

### Integrator: carried sims run (raid band 60, Night Elf, seeds 1|50001)
- Summon Hawk 1 for Endurance Training 4 (BM 5320001505101251/00523500014/0): 674.81|674.91, identical to the published build. The engine does not model Summon Hawk, so the sim cannot decide; the guide keeps the build and must not claim a measured verdict.
- Sniper Shot 1 for one Surefooted point (MM 5/0053550011503051/50005002): 738.23|738.24 against 745.93|745.94. The Marksmanship rotation never casts Sniper Shot, so this only measures the lost hit; not a verdict on the talent. Keep the build; sim round: add a Sniper Shot line before re-measuring.

## Fix log (round 4)

- HUNTER-R4-1 fixed: data/curated/apl/hunter-survival.json Wing Clip note (line 167) now "Without the totem it is 50 damage for a global, so leave it alone." (Resourcefulness makes it 32 mana, so the 80 mana claim is gone). Action unchanged.
- HUNTER-R4-2 fixed: survival.md:22-27 relabelled 'Wowhead Forever, Mongoose Bite' (spell=1495) and added 'Wowhead Forever, Counterattack' (spell=19306, community).
- HUNTER-R4-3 fixed (verified: trainables 409580 skill_line_ability level 1, spelltips 409580 text, ppl gives 280 at 60): first Levels 1 to 19 bullet in beast-mastery.md:140, marksmanship.md:143, survival.md:130 now opens "Train Heart of the Lion at level 1 and keep it up all the way to 60: 10% more stats and extra attack power (280 at 60) for everyone near you, and another 10% stats for you."
- HUNTER-R4-4 fixed (Blizzard deep dive: "player buffs that increase their stats"): index.md:32, beast-mastery.md:57 and :98, marksmanship.md:82, survival.md:71 now say the pet takes no stat buffs from other players; "only your own gear" softened to "your gear is what makes it stronger".
- HUNTER-R4-5 fixed: 'Blizzard, Class Deep Dives: Hunter and Druid' (kind blizzard) added first in sources on index.md, beast-mastery.md, marksmanship.md, survival.md; panel recap kept.
- D1-D4: no change needed (as decided).
- Integrator note, Summon Hawk / Sniper Shot: builds unchanged. No prose claims a measured verdict on either. BM never names Summon Hawk. marksmanship.md:82 keeps the play reason for skipping Sniper Shot (a point from elsewhere, a 4 second cast that holds the auto shot) and the "longer than any other shot you have" overreach (Volley channels 6 s) is now "twice as long as Aimed Shot" (spelltips 1310687 4000 ms vs Aimed Shot 2 s).
- Round 4 source rules: stat weights already point at /sim/weights on the three damage pages; no rotation, tank fight, /sources, wago or Talents Forever entries on hunter pages.
- New SIM findings: none beyond those above (pet non-stat buffs; Sniper Shot line needed in the MM rotation before re-measuring; Summon Hawk unmodelled in the engine).
