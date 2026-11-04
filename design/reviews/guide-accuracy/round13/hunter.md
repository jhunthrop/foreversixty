# Guide accuracy round 13: Hunter

Auditor: hunter. Pages: class index, Beast Mastery, Marksmanship and Survival, plus the card notes in `data/curated/apl/hunter-*.json`.
- Tree: gr13/audit at 604b78a80 (main after guide round 12).
- Client: 1.60.1.70755. Engine pin: 9f7f1f6e4 (sim/enginever/version.go); engine source read from gr10/engine.
- Read first: the round 1 to round 12 hunter reports with every fix log, round4/hunter-levels.md, the race-sims files (rounds 2 to 5) and ledger.md through the round 12 section.

## What moved since round 12

`git diff a8b147a6f..604b78a80` touches one hunter line: beast-mastery.md:81, the round 12 SHOULD on Improved Stings. No hunter APL, `addon-builds.json`, `talents/hunter.json`, `trainables/hunter.json` or hunter BiS file changed. The hunter BiS files are still de044125c (05:46 UTC), so the final re-rank has not landed in this tree.

### Round 12's edit, re-checked

BM:81 now reads "Improved Stings (2 points) adds Serpent Sting damage and, in the raid build, is the way down to Trueshot Aura and Mortal Shots."
- Improved Stings rank 2 is "Increases the damage of your Serpent Sting ability by 13%" (talents/hunter.json). True.
- Raid order (addon-builds.json hunter-beast-mastery raid): Lethal Attacks 5, Improved Stings 2, Efficiency 3, Careful Aim 5 is 15 Marksmanship points before Trueshot Aura (tier 3) at 56 and Mortal Shots (tier 3, needs Careful Aim 5) at 57 to 60. Without Improved Stings the raid build has 13. True.
- Level 60 order: Lethal Attacks 5, Improved Stings 2, Efficiency 4, Careful Aim 5, Rapid Killing 1 is 17, so it does not need Improved Stings to get there, which is why the sentence names the raid build. Consistent.

## Builds and leveling order

Decoded again against `talents/hunter.json` (tier, then column). All five strings are legal at 51 points and every prerequisite is met (Intimidation needs Bestial Swiftness 1; Bestial Wrath needs Intimidation; Frenzy needs Ferocity 5; Mortal Shots needs Careful Aim 5; Rapid Recuperation needs Rapid Killing 2; Sniper Shot needs Trueshot Aura; Counterattack needs Deterrence; Lacerating Strikes needs Expose Prey 2). The decode is the same as round 12:

| Build | Split | Points |
|---|---|---|
| BM leveling | 31/20/0 | Deadly Aspects 5, Endurance Training 4, Focused Fire 2, Bestial Swiftness 1, Unleashed Fury 5, Ferocity 5, Intimidation 1, Bestial Discipline 2, Frenzy 5, Bestial Wrath 1; Lethal Attacks 5, Improved Stings 2, Efficiency 4, Careful Aim 5, Rapid Killing 1, Trueshot Aura 1, Mortal Shots 2 |
| BM raid | 31/20/0 | As leveling, except Efficiency 3, Rapid Killing 0, Mortal Shots 4 |
| MM leveling | 15/31/5 | Deadly Aspects 5, Endurance Training 3, Focused Fire 2, Unleashed Fury 5; Lethal Attacks 5, Efficiency 5, Careful Aim 5, Rapid Killing 2, Trueshot Aura 1, Mortal Shots 5, Rapid Recuperation 2, Barrage 1, Ranged Weapon Specialization 5; Improved Tracking 5 |
| MM raid | 5/33/13 | Deadly Aspects 5; Lethal Attacks 5, Improved Stings 2, Efficiency 5, Careful Aim 5, Lone Wolf 1, Trueshot Aura 1, Mortal Shots 5, Barrage 3, Ranged Weapon Specialization 5, Sniper Shot 1; Improved Tracking 5, Survivalist 5, Surefooted 3 |
| SV | 5/11/35 | Deadly Aspects 5; Lethal Attacks 5, Careful Aim 5, Lone Wolf 1; Improved Tracking 5, Savage Strikes 2, Survivalist 3, Clever Traps 2, Surefooted 3, Deterrence 1, Predator's Edge 5, Counterattack 1, Resourcefulness 2, Expose Prey 2, Survivalist's Discipline 2, Strider Kick 1, Lightning Reflexes 5, Lacerating Strikes 1 |

Shipped orders (`addon-builds.json`, index 0 = level 10) against every leveling sentence:
- BM: Deadly Aspects 10-14; Focused Fire 15-16, the first level tier 1 is open; Endurance Training 17-20; Bestial Wrath 40, the level mail arrives. BM:147 and BM:157 match.
- MM: Lethal Attacks 10-14; Efficiency 15-19; Rapid Killing 25-26 (Rapid Fire trains at 26); Rapid Recuperation 33-34. MM:153 and MM:159 match.
- SV: Improved Tracking 10-14; Savage Strikes 15-16; Survivalist 17-19; Strider Kick 30; Lacerating Strikes 40; Lone Wolf 60, the last point. SV:134 and SV:140 match.

## Findings

None. Nothing on the four hunter pages or in the hunter card notes is false on 70755, contradicts the page, or contradicts the header build card.

### Re-checked line by line this round and still true

- **Talent rank text (70755):** Deadly Aspects 2% a point (Hawk on Auto Shot, Beast on melee autos, 30% for 12 sec); Endurance Training 3% a point; Focused Fire 1% a point; Bestial Swiftness has no outdoor limit in Forever ("movement speed of your pets by 30%"), so "gets the cat to the target sooner after every pull" holds indoors too; Unleashed Fury 15%; Ferocity 10%; Frenzy 100% for 30% over 8 sec; Bestial Discipline 25% a point; Bestial Wrath 50% for 18 sec, "does not feel ... fear"; Intimidation 3 sec stun; Lethal Attacks 1% a point; Improved Stings 13% at 2; Efficiency 3% a point on Shots, Stings and melee; Careful Aim 100% of Intellect; Rapid Killing 1 min/10% and 2 min/20%; Lone Wolf 20%; Mortal Shots 6% of the crit bonus a point (212, 224, 230%); Rapid Recuperation 50%, 100% once Rapid Killing is consumed, 15 sec; Barrage 5% a point; Ranged Weapon Specialization 5%; Improved Tracking 5% on the seven types listed; Savage Strikes 2% crit a point; Survivalist 2% a point; Surefooted 3% hit; Predator's Edge 30% crit bonus (aura 108 / misc 15, round 8 D2) and 50% off hand; Resourcefulness 60%, and on crits 50% regen for 30 sec; Expose Prey 10% for 10 sec; Survivalist's Discipline 40% off trap and Deterrence cooldowns; Strider Kick 100% weapon damage; Lightning Reflexes 15% Agility; Lacerating Strikes 40% over 21 sec.
- **Spell text (spelltips/hunter.json):** Aimed Shot 2 s cast, 6 s, shares with Multi-Shot; Multi-Shot instant (0 ms), 3 targets, so "on the run" (MM:103) holds; Arcane Shot instant; Rapid Fire 40% ranged and melee for 15 sec on 300 s; Sniper Shot ranks 320/450/590 at 40/48/58, 4000 ms, 15000 ms; Aspect of the Hawk 110 at rank 6, 120 at rank 7 (book); Beast 50/70/90/110; Falcon "the same amount as their highest rank of Aspect of the Hawk"; Monkey 8% dodge; Cheetah dazes for 4 sec; Heart of the Lion 10% stats, Attack Power 40 (+4 a level, 280 at 60), another 10% for the hunter; Trueshot Aura rank 5 100 to "party or raid members"; Mongoose Bite "Can only be performed after you dodge"; Counterattack after a parry, cannot be dodged, blocked or parried; Wing Clip rank 3 50 damage; Freezing Trap "Any damage caused will break the ice"; Volley a 6 sec channel; Feign Death; Tranquilizing Shot removes 1 Frenzy.
- **Trainer levels (trainables/hunter.json):** every level on the pages (Heart of the Lion, Raptor Strike 1/8/16, Serpent Sting and Monkey 4, Arcane Shot and Hunter's Mark 6, Concussive Shot 8, Hawk, Tame Beast and Track Humanoids 10, Wing Clip and Mend Pet 12, Mongoose Bite 16, Multi-Shot and Track Undead 18, Aimed Shot, Cheetah, Disengage and Freezing Trap 20, Rapid Fire 26, Feign Death, Beast and Lacerate 30, Explosive Trap 34, Volley 40, Falcon and Tranquilizing Shot 60).
- **Engine (9f7f1f6e4):** Raptor Strike is queued on the next swing (`raptor_strike.go` TryRaptorStrike, the "Raptor Strike Queued" aura), so "replaces your next white hit instead of costing a global" holds. The cat runs Bite and Claw (`pet.go` PetConfigs Cat).
- **Racials (races.json):** Elune's Light 10% crit, 15 s, 3 min; Ruthless 1% crit; Artisan's Precision 1% hit; Big Game Hunter 5% to Beasts; Blood Fury 10% AP, 15 s, 2 min; Command 5% pet damage; Berserking 10% speed, 15 s, 3 min; Wind Blessed 2% haste. combos.json: Human Hunter is new in Forever; both Skyborne races can be Hunters.
- **Raid buffs (SV:65):** Strength of Earth rank 5 is 53 (Era 77); Curse of Recklessness rank 4 is 505 armor (Era 640).
- **Stat prose against band 60 raid weights:** every reason is stable and true. Haste: `haste_on_items` is false and the haste scale factor per 1% is large on all three specs, so "rare on gear but strong when you find it" holds. Survival's Strength factor is 1.21 attack power a point (Heart of the Lion's two 10%), so "Strength is attack power directly" holds. No paragraph ranks stats.
- **Card notes:** all fourteen hunter notes read against the cards and client values; no false statement.
- **Voice:** no em or en dashes; no site, sim, APL, preset, band or pipeline words in the prose; no quoted nightly numbers.
- **Source links:** all 15 URLs fetched again with a browser user agent: 200 for the three Blizzard articles, the three Warcraft Tavern pages and the six foreversixty.gg URLs; 403 for Wowhead Skyborne First Look, spell=1495 and spell=19306 (unknown under the brief, their labels confirmed in round 9). Every entry follows the round 4 source-link rule. The Blizzard hunter deep dive was re-read in full this round (see Decisions D2).

### SHOULDs (wording, not findings, carried)

- MM:159, "so you have it for most hard pulls".
- MM:94, five minutes for a player who raids on the leveling build.
- survival.md:65, "smaller than Era's" wording.
- The Serpent Sting cards' "Refreshing any earlier throws away a tick".
- "The Molten Core tome".
- BM:81 naming Deflection, a Survival talent, in a build that puts no points in Survival.

## Decisions

- **D1. Round 12's Improved Stings edit stands** (see above).
- **D2. The Blizzard deep dive's Savage Strikes line does not override the client.** The deep dive says Savage Strikes "Now increases all melee damage by 2/4%". The 70755 rank text says "Increases the critical strike chance of all your melee abilities by 4%". Following the round 1 lesson, the client text wins, and the pages (SV:75, SV:102, SV:134) follow it. The same article confirms the pages elsewhere: pets "can no longer receive player buffs that increase their stats", Disengage's threat reduction doubled, the 8 yard dead zone, Intimidation needs Bestial Swiftness, Sniper Shot is a 4 second cast on a 15 second cooldown.
- **D3.** Round 12's D1 to D3 and the earlier rounds' decisions stand.

## Sims requested

1. **MM best in slot on the current builds (PROVISIONAL, carried).** `bis/hunter-marksmanship.json` (de044125c) still carries `0053550011503050`. Re-rank on `5/0052550011503051/50005003` and `532000050/005055200152105/5`, then re-read the MM stat table, the enchant table and the hit sentence at MM:119.
2. **MM race order on the new raid build (carried).** Night Elf against Human against Dwarf, and Troll against Orc, at seeds 1|50001, band 60 raid.
3. **Carried from earlier rounds:**
   - Survival weapon enchant: Crusader on both weapons against Agility on both, `5/005005001/500230231051222151`, Night Elf, seeds 1|50001.
   - Race re-runs after the final re-rank: SV Orc against Troll, BM Night Elf against Human, BM Orc against Troll.

## Sim findings

- **Carried, LIKELY:**
  - Aspect of the Falcon's melee proc does not need the talent in the engine (b86fa7d39).
  - Sniper Shot's cast time scales with ranged haste in the engine (`sniper_shot.go`), which the 70755 text does not state.
  - Summon Hawk is unmodeled.
  - Pet non-stat buffs are blocked by RefusePlayerBuffs.
  - The APL top-level notes (not rendered) still say "Aimed Shot is only the pre-pull opener" and call Falcon 55/55.

## Sources

**Forever Sixty client data, 1.60.1.70755**
- `talents/hunter.json`: rank text, tier, column and prerequisites for all three trees.
- `trainables/hunter.json`: every trainer level above.
- `spelltips/hunter.json`: 19434 to 20904, 2643, 3044 to 14287, 3045, 1310687 to 1310786, 13161 and 1299445 to 1299447, 13165 to 25296, 469145, 13163, 5118, 409580, 20906, 1495 to 14271, 19306 to 20910, 14268, 1499 to 14311, 1510 to 14295, 5384, 19801.
- `spelltips/shaman.json` 25361 and `spelltips/warlock.json` 11717.
- `races.json` and `combos.json`.
- `addon-builds.json`: the shipped orders.
- `bis/hunter-*.json`: band 60 raid weights, haste factor and haste_on_items.

**Repository**
- `data/curated/apl/hunter-{beast-mastery,marksmanship,survival}.json`: card notes.
- Engine 9f7f1f6e4: `sim/hunter/raptor_strike.go`, `sim/hunter/pet.go`, `sim/hunter/talents.go`.

**Web** (fetched 2026-11-04, browser user agent)
- Blizzard, World of Warcraft: Forever Class Deep Dives: Hunter and Druid, https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid (200, read in full)
- Blizzard, Deep Dive panel recap, https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200)
- Blizzard, What's Next panel recap, https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap (200)
- Warcraft Tavern, Hunter Class Guide (Forever), https://warcrafttavern.com/forever/guides/hunter/ (200)
- Warcraft Tavern, Hunter Leveling Guide (Forever), https://www.warcrafttavern.com/forever/guides/hunter-leveling/ (200)
- Warcraft Tavern, Hunter and Druid Deep Dive, https://www.warcrafttavern.com/forever/news/powershifting-is-back-sort-of-hunter-druid-deep-dive-for-wow-forever/ (200)
- Wowhead Forever: Skyborne First Look, spell=1495 and spell=19306 (403 to automated fetches, so unknown)
