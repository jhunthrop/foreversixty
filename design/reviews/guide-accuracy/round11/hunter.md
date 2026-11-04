# Guide accuracy round 11: Hunter

Auditor: hunter. Pages: class index, Beast Mastery, Marksmanship and Survival, plus the card notes in `data/curated/apl/hunter-*.json`.
- Tree: gr11/audit at d54de48aa (main after guide round 10, leveling-order and rotation-gates merged).
- Client: 1.60.1.70755.
- Engine pin: 9f7f1f6e4 (sim/enginever/version.go). Engine source read from gr10/engine (HEAD 9f7f1f6e4).
- Read first: the round 1 to round 10 hunter reports with every fix log, round4/hunter-levels.md, the round 2 to round 5 race-sims files and ledger.md (through the round 10 section).

## What moved since round 10

`git log 84460cec8..HEAD` on the hunter files shows one hunter commit, eabba8f84 (the round 10 fixes): marksmanship.md and the Sniper Shot note in `hunter-marksmanship.json`. Nothing else hunter-specific moved:
- `addon-builds.json`, `talents/hunter.json`, `trainables/hunter.json` and the hunter BiS files are unchanged.
- The BM and SV pages and the index are unchanged.
- The other commits in range are component work (ec60e6abf strip/planner/home on the ladder order, b53e82298 rotation card gates, b7ff3cb00 leveling-build reservation). Those are the component lane's, so I filed nothing on them.

## Builds and leveling order

Decoded again against `talents/hunter.json`, sorted by tier and then column. All five strings are legal at 51 points, and every prerequisite is met: Intimidation needs Bestial Swiftness, Frenzy needs Ferocity 5, Bestial Wrath needs Intimidation, Mortal Shots needs Careful Aim 5, Rapid Recuperation needs Rapid Killing 2, Sniper Shot needs Trueshot Aura, Counterattack needs Deterrence, and Lacerating Strikes needs Expose Prey 2.

| Build | Split | Points |
|---|---|---|
| BM leveling | 31/20/0 | Deadly Aspects 5, Endurance Training 4, Focused Fire 2, Bestial Swiftness 1, Unleashed Fury 5, Ferocity 5, Intimidation 1, Bestial Discipline 2, Frenzy 5, Bestial Wrath 1; Lethal Attacks 5, Improved Stings 2, Efficiency 4, Careful Aim 5, Rapid Killing 1, Trueshot Aura 1, Mortal Shots 2 |
| BM raid | 31/20/0 | As leveling, except Efficiency 3, Rapid Killing 0, Mortal Shots 4 |
| MM leveling | 15/31/5 | Deadly Aspects 5, Endurance Training 3, Focused Fire 2, Unleashed Fury 5; Lethal Attacks 5, Efficiency 5, Careful Aim 5, Rapid Killing 2, Trueshot Aura 1, Mortal Shots 5, Rapid Recuperation 2, Barrage 1, Ranged Weapon Specialization 5; Improved Tracking 5 |
| MM raid | 5/33/13 | Deadly Aspects 5; Lethal Attacks 5, Improved Stings 2, Efficiency 5, Careful Aim 5, Lone Wolf 1, Trueshot Aura 1, Mortal Shots 5, Barrage 3, Ranged Weapon Specialization 5, Sniper Shot 1; Improved Tracking 5, Survivalist 5, Surefooted 3 |
| SV | 5/11/35 | Deadly Aspects 5; Lethal Attacks 5, Careful Aim 5, Lone Wolf 1; Improved Tracking 5, Savage Strikes 2, Survivalist 3, Clever Traps 2, Surefooted 3, Deterrence 1, Predator's Edge 5, Counterattack 1, Resourcefulness 2, Expose Prey 2, Survivalist's Discipline 2, Strider Kick 1, Lightning Reflexes 5, Lacerating Strikes 1 |

Shipped orders from `addon-builds.json` (index 0 = level 10). I simulated every order point by point for row gates and prerequisites, and every point is legal:
- **BM:** Deadly Aspects 10-14, Focused Fire 15-16, Endurance Training 17-20, Bestial Swiftness 21, Unleashed Fury 22-26, Ferocity 27-31, Intimidation 32, Bestial Discipline 33-34, Frenzy 35-39, Bestial Wrath 40, Lethal Attacks 41-45, Improved Stings 46-47, Efficiency 48-51, Careful Aim 52-56, Rapid Killing 57, Trueshot Aura 58, Mortal Shots 59-60.
- **MM leveling:** Lethal Attacks 10-14, Efficiency 15-19, Careful Aim 20-24, Rapid Killing 25-26, Trueshot Aura 27, Mortal Shots 28-32, Rapid Recuperation 33-34, Barrage 35, Ranged Weapon Specialization 36-40, Deadly Aspects 41-45, Endurance Training 46-48, Focused Fire 49-50, Unleashed Fury 51-55, Improved Tracking 56-60.
- **MM raid:** Sniper Shot at 40, after exactly 30 Marksmanship points, so tier 7 is open.
- **SV:** Improved Tracking 10-14, Savage Strikes 15-16, Survivalist 17-19, ... Lacerating Strikes 40, ... Careful Aim 55-59, Lone Wolf 60.

Every leveling sentence matches these orders:
- BM: "first five points go into Deadly Aspects", "Focused Fire at 15 and 16 ... Endurance Training starts at 17", and "Bestial Wrath comes the same level" as mail at 40.
- MM: "first five points go into Lethal Attacks", "Efficiency from 15", "Rapid Killing at 25 and 26" (Rapid Fire trained at 26) and "Rapid Recuperation at 33 and 34".
- SV: "Improved Tracking" first, "Savage Strikes at 15 and 16", "Survivalist from 17", and "Lone Wolf is one of the last points" (it is the level 60 point).

## Findings

None. No statement on the four hunter pages or in the hunter card notes is false on 70755 or contradicts the page.

### Round 10's fixes, re-checked

All hold against the 70755 rank text, the spelltips and the decoded strings:
- **MM:69, the build paragraph.** 31/15/5 leveling. The raid trade, with the count checked: 31 - 4 (Rapid Killing 2, Rapid Recuperation 2) + Lone Wolf 1 + Sniper Shot 1 + Barrage 2 + Improved Stings 2 = 33.
- **MM:72 and :80.**
  - Barrage is 5% a rank (19461 rank 3 is 15%).
  - Improved Stings rank 2 is 13% (1310661).
- **MM:82, Sniper Shot.** Spelltip 1310786: "ranged damage plus 590", 4000 ms cast, 15000 ms cooldown, 365 mana. Rank 3 is trained at 58 (trainables: 40, 48, 58).
- **MM:82, "the hardest hit you have".**
  - The engine computes both shots from normalized weapon damage (`sim/hunter/aimed_shot.go:63`, `sniper_shot.go`).
  - Aimed Shot is (N + 350) x 1.15 with Barrage. Sniper Shot is N + 590, with no Barrage multiplier (Barrage's text names only Multi-Shot, Aimed Shot and Volley).
  - Sniper Shot is the larger hit while N is under 1250. A raid bow with about 2500 ranged attack power gives N of about 800, so the claim holds.
- **MM:83, Rapid Killing rank 2 (415405).** "2 min" off Rapid Fire, and "when you kill a non-trivial enemy or it dies while afflicted by your Serpent Sting ... next Shot ... by 20%".
- **MM:84, Rapid Recuperation rank 2 (1223987).** "50% ... increased to 100% when Rapid Killing is consumed", 15 sec. Rank 1 is 25%/50%, so "half ... all of it" is the 2-point value.
- **MM:88.** Sniper Shot's 4 s cast is twice Aimed Shot's 2 s. In the engine, both casts scale with ranged speed, so the ratio holds at any haste. Survivalist (tier 2) is the 10th Survival point that opens Surefooted's tier.
- **MM:96 and the Sniper Shot card note.**
  - The APL order is Aimed Shot (next auto over 1.2 s), Multi-Shot (no gate), then Sniper Shot (next auto over 1.2 s). So Sniper Shot fires only while the shared cooldown is not ready, which is what both texts say.
  - The 15 s cooldown is its own timer (`NewTimer()` in `sniper_shot.go`).
- **MM:52, :59 and :94.** The raid build has no Rapid Killing, so Rapid Fire's five minutes and "once in a normal boss fight" hold. "The shots on your shared cooldown" is true for Barrage, Mortal Shots and Ranged Weapon Specialization.
- **MM:105.** Sniper Shot costs a flat 365, and keeping it through low mana is consistent with the priority.
- **MM:159.** At 25, Rapid Killing rank 1 takes 1 minute off; at 26 (two points) it takes 2. Rapid Fire is trained at 26. Rapid Recuperation's tier-5 gate needs 20 Marksmanship points, and the build has 23 at level 33.

### What was checked again and holds

- **Every talent value on all three pages** against the 70755 rank text:
  - Mortal Shots 30% bonus (230%, and 212%/224% on BM). Predator's Edge 30% bonus and 50% off hand.
  - Unleashed Fury 15%, Ferocity 10%, Frenzy 100% / 30% / 8 sec, Focused Fire 2%, Bestial Wrath 50% / 18 sec, Bestial Discipline 50%.
  - Deadly Aspects 10% / 30% / 12 sec (Hawk: Auto Shot; Beast: melee autos).
  - Lethal Attacks 5%, Efficiency 15%, Careful Aim 100% of Intellect, Lone Wolf 20%, Ranged Weapon Specialization 5%.
  - Improved Tracking 5% for the seven listed types.
  - Savage Strikes 4%, Survivalist 10%, Surefooted 3%.
  - Clever Traps 30% Immolation/Explosive damage, Survivalist's Discipline 40% trap cooldown.
  - Resourcefulness 60% and "100% chance to allow 50% of your Mana Regeneration ... 30 sec".
  - Expose Prey 10% / 10 sec, Lightning Reflexes 15%, Lacerating Strikes 40% over 21 sec.
- **Trainer levels.** `trainables/hunter.json` and `spellranks.json` hold for every level the pages name:
  - Serpent Sting and Monkey at 4; Arcane Shot and Hunter's Mark at 6; Concussive Shot at 8; Hawk and Tame Beast at 10.
  - Wing Clip and Mend Pet at 12; Mongoose Bite at 16; Multi-Shot and Track Undead at 18.
  - Aimed Shot, Cheetah, Disengage and Freezing Trap at 20; Rapid Fire at 26.
  - Feign Death, Beast and Lacerate at 30; Volley at 40; Falcon and Tranquilizing Shot at 60.
  - No Kill Command (so "no execute phase" holds), and no page names a class_spell rune (Wyvern Strike, Widow Bite, Hydra Shot).
- **SV:86, the Falcon line.**
  - Hawk rank 7 (25296, 120) is `"book": true` in `spellranks.json`, so "10 more once you own Hawk's Ahn'Qiraj rank" holds: Falcon 110 = Beast rank 4 110 at 60 until the book.
  - Tranquilizing Shot (19801) has no book flag, so the "bought from your trainer at 60" lines hold.
- **Card notes.**
  - The BM, MM and SV notes all agree with their APL lines and gates.
  - SV's priority sentence matches the APL order: Raptor Strike, Counterattack, Mongoose Bite (gated on 5302), Lacerate, Strider Kick, Immolation Trap, then Wing Clip under Windfury.
- **Races, consumables, enchants, ammo and quivers.** Unchanged since round 9 and still true. Race orders on the new MM raid build are unmeasured, and no line contradicts a measured result or the client.
- **Voice.**
  - No em or en dashes on any page or in any card note.
  - No site, sim, APL, preset, band or pipeline words in the prose (only `kind: site` in frontmatter).
  - No quoted nightly numbers, and no ranking words in Stat priority.
- **Source links.** All 15 URLs were fetched with a browser user agent.
  - 200: the three Blizzard articles, the three Warcraft Tavern pages and the six foreversixty.gg URLs.
  - 403: Wowhead's Skyborne First Look, spell=1495 and spell=19306. This counts as unknown, not dead; round 9 confirmed both spell labels through the nether tooltip endpoint.
  - Body links go to /bis/hunter/<spec>. Every entry follows the source-link rule.

### SHOULDs (wording, not findings)

- **MM:159.** "So you have it for most hard pulls" oversells a three minute cooldown a little. "So it is back for the next hard pull" would be safer.
- **MM:94.** "Rapid Fire's five minute cooldown" is right for the raid build. A player who raids on the leveling build (MM:107 allows it) has three minutes from Rapid Killing.
- **Carried from round 9:**
  - survival.md:65, "most of the melee buffs are smaller than Era's".
  - The Serpent Sting card's "throws away a tick".
  - "The Molten Core tome".
  - BM:81 naming Deflection.

## Decisions

- **D1. "The hardest hit you have" stands** (MM:82 and the Sniper Shot note). See the arithmetic under the re-checked fixes. The claim fails only above a normalized shot base of 1250, which no level 60 raid bow reaches.
- **D2. "Hawk's Ahn'Qiraj rank" stands.** The 70755 `spellranks.json` flags 25296 as a book. `trainables/hunter.json` lists it as a skill_line_ability row at 60, but the book flag is the specific signal, and the APL notes agree that rank 6 is the trainer's top rank.
- **D3.** Round 10's D1 to D4 and round 9's D1 to D3 stand.

## Sims requested

1. **MM best in slot on the new builds (PROVISIONAL).**
   - Band 60 raid in `bis/hunter-marksmanship.json` (last written de044125c, 05:46 UTC) still carries the old raid string `0053550011503050` (Improved Stings 3, no Sniper Shot). The leveling bands carry the old leveling string.
   - Re-rank on `5/0052550011503051/50005003` and `532000050/005055200152105/5`, then re-read the MM Stat priority table, the enchant line (MM:127) and the BiS prose.
2. **MM race order on the new raid build.**
   - Night Elf against Human against Dwarf, and Troll against Orc, at seeds 1|50001, band 60 raid.
   - MM:135-137 and `recommendedRaces: [night-elf, troll]` rest on round 2's sims of the old build. Sniper Shot's 4 s cast changes how much Berserking's haste is worth against Blood Fury's attack power.
3. **Carried:**
   - Survival weapon enchant: Crusader on both against Agility on both, `5/005005001/500230231051222151`, Night Elf, seeds 1|50001.
   - Hunter race re-runs after the final re-rank: SV Orc against Troll, BM Night Elf against Human, BM Orc against Troll.

## Sim findings

- **LIKELY, carried from round 10: Aspect of the Falcon's melee proc does not need the talent in the engine (b86fa7d39).** The client text ties the melee proc to "the Improved Aspect of the Hawk talent". No page states a chance.
- **Note for the sim lane (not a guide finding): Sniper Shot's cast time scales with ranged haste in the engine.** `sniper_shot.go` uses CastTime = 4000 ms / RangedSwingSpeed. Aimed Shot follows the same Classic rule, but the 70755 client does not state it for Sniper Shot. Confirm in game that a quiver and Rapid Fire shorten the cast.
- **Carried:**
  - Summon Hawk is unmodeled.
  - Pet non-stat buffs are blocked by RefusePlayerBuffs.
  - The APL top-level notes (not rendered) still say "Aimed Shot is only the pre-pull opener" and call Falcon 55/55.

## Sources

**Forever Sixty client data, 1.60.1.70755**
- `talents/hunter.json`: rank text and prerequisites for all three trees, including 415405, 1223987, 19461, 1310661, 1310687 and 19485.
- `spelltips/hunter.json`:
  - Sniper Shot 1310687, 1310785 and 1310786; Aimed Shot 19434 to 20904; Multi-Shot 2643; Rapid Fire 3045.
  - Hawk 13165 to 25296; Beast 13161 and 1299445 to 1299447; Falcon 469145.
  - Trueshot Aura 1299346 to 20906; Lacerate; Strider Kick 1317257; Disengage; Volley; Tranquilizing Shot 19801; Heart of the Lion 409580.
- `trainables/hunter.json` and `spellranks.json`: the book flags on 25296 and 25295; Tranquilizing Shot unflagged.
- `addon-builds.json`: the BM, MM and SV leveling and raid orders.
- `bis/hunter-*.json`: band talent strings.

**Repository**
- `data/curated/apl/hunter-{beast-mastery,marksmanship,survival}.json`: card notes and gates.
- Commit eabba8f84: the round 10 fixes.
- `design/reviews/guide-accuracy/ledger.md`: the round 10 section.

**Engine 9f7f1f6e4**
- `sim/hunter/sniper_shot.go`: normalized weapon damage, its own timer, cast time over RangedSwingSpeed.
- `sim/hunter/aimed_shot.go:57,63`: Barrage multiplier and normalized weapon damage.

**Web** (fetched 2026-11-04)
- Blizzard, World of Warcraft: Forever Class Deep Dives: Hunter and Druid, https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid
- Blizzard, Deep Dive panel recap, https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Blizzard, What's Next panel recap, https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap
- Warcraft Tavern, Hunter Class Guide (Forever), https://warcrafttavern.com/forever/guides/hunter/
- Warcraft Tavern, Hunter Leveling Guide (Forever), https://www.warcrafttavern.com/forever/guides/hunter-leveling/
- Warcraft Tavern, Hunter and Druid Deep Dive, https://www.warcrafttavern.com/forever/news/powershifting-is-back-sort-of-hunter-druid-deep-dive-for-wow-forever/
- Wowhead Forever, spell=1495, spell=19306 and Skyborne First Look: 403 to automated fetches, so unknown.
