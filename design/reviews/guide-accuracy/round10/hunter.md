# Guide accuracy round 10: Hunter

Auditor: hunter. Pages: class index, Beast Mastery, Marksmanship and Survival, plus the card notes in `data/curated/apl/hunter-*.json`.
- Tree: gr10/audit at 84460cec8 (origin/main: guide rounds 1 to 9 plus sim round 15).
- Client: 1.60.1.70755.
- Engine pin: 9f7f1f6e4 (sim/enginever/version.go). Engine source read from gr10/engine.
- Read first: the round 1 to round 9 hunter reports with every fix log, round4/hunter-levels.md, the round 2 to round 5 race-sims files and ledger.md.

## What moved since round 9

`git diff 645caa983 HEAD` on the hunter files shows that only Marksmanship moved. Two sim round 15 commits changed it, and neither touched the prose:
- **1b923c4f0, "Marksmanship raids with Sniper Shot".** It changed the raid build (Sniper Shot replaces the third Improved Stings point). It also added a Sniper Shot line to the rotation (1310786 rank 3, gated on next auto shot over 1.2 s) with an empty note, "left for the guide lane".
- **1134376ac, "Marksmanship levels with Rapid Killing and Rapid Recuperation".** It changed the leveling build. Ferocity 1, Improved Stings 1 and two Barrage ranks went into Rapid Killing 2 and Rapid Recuperation 2.
- `addon-builds.json` was regenerated to match both strings.

Other files:
- The BM and SV pages, the index, the BM and SV card notes, talents, trainables, spelltips and races did not move.
- The only hunter engine change is b86fa7d39 (Aspect of the Falcon's melee proc). See Sim findings.

## Builds and leveling order

Decoded against `talents/hunter.json`, in tier then column order. All builds are 51 points and legal:
- Rapid Recuperation's prerequisite is Rapid Killing 2/2.
- Sniper Shot's prerequisite is Trueshot Aura 1/1, with 32 Marksmanship points above tier 7.

| Build | Split | Points |
|---|---|---|
| BM leveling / raid | 31/20/0 | Unchanged since round 9 |
| **MM leveling** `532000050/005055200152105/5` | **15/31/5** | Deadly Aspects 5, Endurance Training 3, Focused Fire 2, Unleashed Fury 5; Lethal Attacks 5, Efficiency 5, Careful Aim 5, **Rapid Killing 2**, Trueshot Aura 1, Mortal Shots 5, **Rapid Recuperation 2**, **Barrage 1**, Ranged Weapon Specialization 5; Improved Tracking 5. **No Improved Stings, no Ferocity.** |
| **MM raid** `5/0052550011503051/50005003` | 5/33/13 | Deadly Aspects 5; Lethal Attacks 5, **Improved Stings 2**, Efficiency 5, Careful Aim 5, Lone Wolf 1, Trueshot Aura 1, Mortal Shots 5, Barrage 3, Ranged Weapon Specialization 5, **Sniper Shot 1**; Improved Tracking 5, Survivalist 5, Surefooted 3 |
| SV | 5/11/35 | Unchanged since round 9 |

Shipped orders from `addon-builds.json` (index 0 = level 10):
- **MM leveling:** Lethal Attacks 10-14, Efficiency 15-19, Careful Aim 20-24, Rapid Killing 25-26, Trueshot Aura 27, Mortal Shots 28-32, Rapid Recuperation 33-34, Barrage 35, Ranged Weapon Specialization 36-40, Deadly Aspects 41-45, Endurance Training 46-48, Focused Fire 49-50, Unleashed Fury 51-55, Improved Tracking 56-60.
- **MM raid:** Lethal Attacks 10-14, Improved Stings 15-16, Efficiency 17-21, Careful Aim 22-26, Lone Wolf 27, Trueshot Aura 28, Mortal Shots 29-33, Barrage 34-36, Ranged Weapon Specialization 37-39, Sniper Shot 40, Ranged Weapon Specialization 41-42, Deadly Aspects 43-47, Improved Tracking 48-52, Survivalist 53-57, Surefooted 58-60.
- **BM and SV:** identical to round 9. Every BM and SV leveling sentence still matches.

Client values for the new talents (70755 rank text and spelltips):
- **Rapid Killing (415405), rank 2:** "Reduces the cooldown on your Rapid Fire ability by 2 min. ... increasing the damage of your next Shot ability within 20 sec by 20%."
- **Rapid Recuperation (1223987), rank 2:** "Afflicting a target with Serpent Sting allows 50% of your Mana Regeneration to continue while casting for 15 sec, increased to 100% when Rapid Killing is consumed."
- **Barrage (19461):** 5% a rank. Rank 1 is 5%, rank 3 is 15%.
- **Improved Stings (1310661):** 6%, 13% and 20% Serpent Sting damage.
- **Sniper Shot:** a tier 7 talent with ranks at 40, 48 and 58 (1310687, 1310785, 1310786). Rank 3: "A long-range shot that deals ranged damage plus 590 and increases the range of your next 3 Shots by 10 yards for 10 sec." 4000 ms cast, 15000 ms cooldown, no category cooldown (its own timer, not the Aimed/Multi one), 365 mana. Effect 121 is normalized weapon damage (engine `sim/hunter/sniper_shot.go:82`). Aimed Shot rank 6 (20904) is "increases ranged damage by 350" on a 2000 ms cast.

## Findings

All six findings are on marksmanship.md. All six come from the two sim round 15 build commits, which did not update the prose.

### HUNTER-R10-1. CONFIRMED. talent. marksmanship.md:85
**Quote:** "Sniper Shot, the capstone, costs a point you would take from elsewhere in either build, and its 4 second cast holds your auto shot twice as long as Aimed Shot, so neither build takes it."

The same paragraph says the raid build "spends the points on Improved Stings and Surefooted's hit".

**Evidence:**
- The raid string's 16th Marksmanship digit is 1: Sniper Shot 1/1.
- The addon raid order has `[2,7,2]` at index 30, which is level 40.
- `hunter-marksmanship.json` casts 1310786 rank 3.
- The leveling build does skip it.

**Fix (replace from "So the raid build drops" to the end of the paragraph):**
> So the raid build drops the pet talents, takes Lone Wolf, and spends the points on Sniper Shot, the last ranks of Barrage, Improved Stings and Surefooted's hit; Survivalist is the way to Surefooted. Sniper Shot, the capstone, is your biggest shot: weapon damage plus 590 at rank 3, on its own 15 second cooldown. Its 4 second cast holds your auto shot twice as long as Aimed Shot, so it goes out right after an auto shot leaves, the same as Aimed Shot. While leveling, keep the pet: it holds what you pull, the leveling build keeps its talents, and its last Marksmanship points go into Rapid Killing and Rapid Recuperation instead of Sniper Shot.

### HUNTER-R10-2. CONFIRMED. talent. marksmanship.md:69
**Quote:** "thirty points in Marksmanship take the talents that raise your big shots and your auto shot, sixteen in Beast Mastery take the pet talents and Deadly Aspects ... thirty-three points in Marksmanship add Lone Wolf and the rest of Improved Stings"

**Evidence:**
- The leveling split is 15/31/5, not 16/30/5.
- The leveling build has no Improved Stings, so "the rest of" is false.
- Against the leveling build, the raid build adds Lone Wolf, Sniper Shot, two Barrage ranks and Improved Stings 2/3, and drops Rapid Killing and Rapid Recuperation.

**Fix:**
> There are two builds. You level with the pet: thirty-one points in Marksmanship take the talents that raise your big shots and your auto shot, plus Rapid Killing for a shorter Rapid Fire and Rapid Recuperation to keep your mana coming back while you shoot; fifteen in Beast Mastery take Deadly Aspects and the pet talents, and five in Survival take Improved Tracking. At 60, for raids, the pet goes: thirty-three points in Marksmanship trade Rapid Killing and Rapid Recuperation for Lone Wolf, Sniper Shot, the last two ranks of Barrage and two points of Improved Stings, five in Beast Mastery keep Deadly Aspects, and thirteen in Survival take Improved Tracking, Survivalist and Surefooted.

### HUNTER-R10-3. CONFIRMED. talent. marksmanship.md:80
**Quote:** "**Improved Stings** (3 points in the raid build): 6% more Serpent Sting damage with the 1 point you level with, 20% with all 3 in the raid build"

**Evidence:** The raid build has Improved Stings 2/3, which is 13% (rank text 1310661). The leveling build has 0.

**Fix:**
> - **Improved Stings** (2 points, raid build): 13% more Serpent Sting damage

### HUNTER-R10-4. CONFIRMED. talent. marksmanship.md:72
**Quote:** "**Barrage** (3 points): 15% more damage on Multi-Shot, Aimed Shot and Volley"

**Evidence:** The leveling build has Barrage 1/3, which is 5%. The list states every other talent with the points both builds share.

**Fix:**
> - **Barrage** (3 points in the raid build): 15% more damage on Multi-Shot, Aimed Shot and Volley, 5% with the 1 point you level with

Add the two leveling talents and the raid capstone to the list, with their 70755 values:
> - **Rapid Killing** (2 points, leveling build): two minutes off Rapid Fire, and your next shot after a kill hits 20% harder
> - **Rapid Recuperation** (2 points, leveling build): after your Serpent Sting lands, half your mana regeneration keeps running while you cast for 15 seconds, all of it once Rapid Killing is spent
> - **Sniper Shot** (1 point, raid build): weapon damage plus 590 at rank 3 on a 4 second cast and its own 15 second cooldown, the hardest hit you have

### HUNTER-R10-5. CONFIRMED. leveling. marksmanship.md:150
**Quote:** "Improved Stings at 15 adds 6% to Serpent Sting, and Efficiency from 16 cuts the mana cost of your shots and stings by 3% a point"

**Evidence:** The shipped leveling order has Efficiency 15-19 and no Improved Stings point at any level.

**Fix:**
> Your first five points go into Lethal Attacks, 1% crit on every attack per point. Efficiency from 15 cuts the mana cost of your shots and stings by 3% a point, which is fewer drinking breaks.

**SHOULD:** the Levels 20 to 60 block could add "Rapid Killing at 25 and 26 takes two minutes off Rapid Fire, so you have it for most hard pulls" (two minutes off five is three). It is not required: nothing there is false.

### HUNTER-R10-6. CONFIRMED. rotation. marksmanship.md:48, :93, :102 and the Sniper Shot card
**Quotes:**
- :48: "Marksmanship's damage is your own shots: Aimed Shot and Multi-Shot on their shared cooldown, Serpent Sting rolling, Arcane Shot in the gaps."
- :102: "stop Arcane Shot first, then Serpent Sting, and keep Aimed Shot and Multi-Shot going."

**Evidence:**
- The page's own raid build takes Sniper Shot (R10-1).
- The rotation casts it right after Multi-Shot and before Serpent Sting, gated on the next auto shot being more than 1.2 s away.
- Neither the overview nor the priority nor the mana line names it, so a player running the page's raid build never learns when to press its capstone.
- The card note is empty (commit 1b923c4f0 left it "for the guide lane").

**Fixes:**
- :48:
  > Marksmanship's damage is your own shots: Aimed Shot and Multi-Shot on their shared cooldown, Sniper Shot in a raid, Serpent Sting rolling, Arcane Shot in the gaps. In a raid, no pet: Lone Wolf.
- :93, after the Multi-Shot sentence:
  > With the raid build, Sniper Shot comes next: when it is ready and the shared cooldown is not, cast it right after an auto shot fires, never partway to the next one.
- :102:
  > When mana gets low, stop Arcane Shot first, then Serpent Sting, and keep Aimed Shot, Multi-Shot and Sniper Shot going.
- Card note for the 1310786 line:
  > Sniper Shot when the shared cooldown is down and an auto shot has just left. Its 4 second cast holds your next auto, so start it right after one fires, never partway to the next. It is your hardest hit, on its own 15 second cooldown.

## What was checked and holds

**Marksmanship, apart from the findings:**
- Strengths and weaknesses hold. Rapid Fire's five minute cooldown in the raid context is right: the raid build has no Rapid Killing.
- The 4 second against 2 second cast comparison is right.
- Trueshot Aura 100 at 60 (20906) holds.
- The Hawk-over-Falcon reasoning holds (see D1).
- The pet, threat, trash and utility lines hold, and so do Levels 20 to 60.

**Beast Mastery, Survival and the index:**
- No text changed, and every claim round 9 verified still holds on 70755 data.
- I re-checked the spots most likely to move with sim round 15. These all hold:
  - BM:81, the Rapid Killing rank 1 values.
  - SV:82, Rapid Killing's row and the ten points.
  - SV:86, Falcon (see D1).
  - SV:144, Dual Wield at 20 (round 5 D3).
  - SV:78, Resourcefulness "50% of your Mana Regeneration" for 30 sec.
  - Index:33, the Aimed/Multi shared cooldown.

**Prior checks still holding:**
- Round 9's crit-damage checks (Mortal Shots 212/224/230%, Predator's Edge 230%) stand.
- No Eureka! line exists on any hunter page.
- Skyborne is named only as the paid option.

**Source links:** I fetched all 15 URLs with a browser user agent.
- 200: the three Blizzard articles, the three Warcraft Tavern pages and the six foreversixty.gg URLs.
- 403: Wowhead's Skyborne First Look, spell=1495 and spell=19306. These count as unknown, not dead; round 9 confirmed both spell labels through the nether tooltip endpoint.
- Every entry still follows the source-link rule.

**Voice:** No em or en dashes, no site, sim or pipeline words, no quoted nightly numbers, no ranking in Stat priority, and no removed spells or runes.

### SHOULDs (wording, not findings)
- **MM Levels 20 to 60:** name Rapid Killing at 25 (see R10-5).
- **Carried from round 9:**
  - survival.md:65, "most of the melee buffs are smaller than Era's".
  - The Serpent Sting card's "throws away a tick".
  - "the Molten Core tome".
  - BM:81 naming Deflection.

## Decisions

- **D1. The Falcon lines stand** (BM:87, MM:89, SV:86). The 70755 Falcon description (469145) says: "The hunter takes on the aspects of a Falcon, increasing melee and ranged Attack Power by the same amount as their highest rank of Aspect of the Hawk. While Aspect of the Falcon is active, the Improved Aspect of the Hawk talent can now proc from and benefit melee attacks." Deadly Aspects' rank text ties the Auto Shot proc to "While Aspect of the Hawk is active". Two readings agree with the pages:
  - Falcon's proc aura is melee only (ProcTypeMask 20).
  - "Never from your auto shots" and "procs from your swings under either one" both hold.
  - No page states a Falcon proc chance.
  - "Falcon is not a clear upgrade" agrees with b86fa7d39's measurement (Beast ahead).
- **D2. Sniper Shot is no longer engine-blocked.** It is modeled (`sim/hunter/sniper_shot.go`) and cast by the rotation, so round 9's "blocked on the engine: Sniper Shot" item is closed. The guide now has to teach it (R10-1, R10-6).
- **D3. The leveling build's Rapid Killing and Rapid Recuperation are taken as shipped.** Commit 1134376ac decided them under the owner's tie rule from Forever leveling guides. The guide only has to describe them truthfully (R10-2, R10-4).
- **D4.** Round 9's D1 to D3 stand.

## Sims requested

1. **MM BiS on the new raid build (PROVISIONAL).** Band 60 raid in `bis/hunter-marksmanship.json` still carries the old talent string `0053550011503050` (Improved Stings 3, no Sniper Shot). Bands 50 to 58 carry the old leveling string. The 13:00 UTC re-rank should run on `5/0052550011503051/50005003` and `532000050/005055200152105/5`, and the Stat priority table and prose should be re-read after it. Any prose against BiS row mismatch is provisional until then.
2. **MM race order on the new raid build.** Night Elf against Human against Dwarf, and Troll against Orc, seeds 1|50001, band 60 raid. MM:132-134 and `recommendedRaces: [night-elf, troll]` rest on round 2's sims of the old build, and Sniper Shot's 4 s cast shifts the value of Berserking's haste against Blood Fury's attack power. This folds into the carried hunter race re-runs below.
3. **Carried:**
   - The Survival weapon enchant: Crusader on both weapons against Agility on both, `5/005005001/500230231051222151`, Night Elf, seeds 1|50001.
   - Hunter race re-runs after the final re-rank: SV Orc against Troll, BM Night Elf against Human, BM Orc against Troll.
   - Summon Hawk is unmodeled.

## Sim findings

- **LIKELY. Aspect of the Falcon's melee proc does not need the talent in the engine (b86fa7d39).** The engine rolls Quick Strikes (469144) at a flat 5% under Falcon with no Deadly Aspects points. The 70755 description says it is "the Improved Aspect of the Hawk talent" that "can now proc from and benefit melee attacks", which ties the proc to the talent. The ProcChance 5 in SpellAuraOptions 239200 may be the per-rank base that Deadly Aspects scales. No guide states the chance, and Beast stays the Survival aspect either way. The sim lane should check whether a 0-point Deadly Aspects character procs Quick Strikes under Falcon in the client.
- **Carried:**
  - Summon Hawk is unmodeled.
  - Pet non-stat buffs are blocked by RefusePlayerBuffs.
  - The APL top-level notes (not rendered) still say "Aimed Shot is only the pre-pull opener" and call Falcon 55/55.

## Sources

**Forever Sixty client data, 1.60.1.70755**
- `talents/hunter.json`: Rapid Killing 415405, Rapid Recuperation 1223987 (prerequisite 105005), Barrage 19461, Improved Stings 1310661, Sniper Shot (prerequisite 105004), Endurance Training, Ferocity, Resourcefulness.
- `spelltips/hunter.json` and `spellconst/hunter.json`: 1310687, 1310785 and 1310786 Sniper Shot; 20904 Aimed Shot; 2643 Multi-Shot; 20906 Trueshot Aura; 469145 Aspect of the Falcon.
- `trainables/hunter.json` and `spellranks.json`: Sniper Shot ranks at 40, 48 and 58; Aimed Shot; Rapid Fire at 26.
- `addon-builds.json`: hunter-marksmanship leveling and raid orders.
- `bis/hunter-*.json`: band talent strings.

**Repository**
- Commits 1b923c4f0 and 1134376ac: the MM raid and leveling builds and the Sniper Shot rotation line.
- `data/curated/apl/hunter-marksmanship.json`: the Sniper Shot line and its empty note.

**Engine 9f7f1f6e4**
- `sim/hunter/sniper_shot.go:11-24,82`: normalized weapon damage.
- `sim/hunter/constants_auto_gen.go:4463-4472`: Sniper Shot ranks.
- Commit b86fa7d39, `sim/hunter/aspect_of_the_falcon.go`: the Falcon proc.

**Web** (fetched 2026-11-04)
- Blizzard, World of Warcraft: Forever Class Deep Dives: Hunter and Druid, https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid
- Blizzard, Deep Dive panel recap, https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Blizzard, What's Next panel recap, https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap
- Warcraft Tavern, Hunter Class Guide (Forever), https://warcrafttavern.com/forever/guides/hunter/
- Warcraft Tavern, Hunter Leveling Guide (Forever), https://www.warcrafttavern.com/forever/guides/hunter-leveling/
- Warcraft Tavern, Hunter and Druid Deep Dive, https://www.warcrafttavern.com/forever/news/powershifting-is-back-sort-of-hunter-druid-deep-dive-for-wow-forever/
- Wowhead Forever, spell=1495, spell=19306 and Skyborne First Look: 403 to automated fetches, so unknown.

## Fix log (round 10)

Files changed: web/src/content/guides/hunter/marksmanship.md and the Sniper Shot line's `notes` in data/curated/apl/hunter-marksmanship.json (no action, build string or BiS file touched). `updated:` on marksmanship.md was already 2026-11-04. Prose test (`-t hunter`) passes.

- **HUNTER-R10-1** fixed, marksmanship.md:88. The auditor's replacement, checked against spelltips 1310786 (590, 4 s cast, 15 s cooldown) and the raid string (Sniper Shot 1/1, Barrage 3, Improved Stings 2).
- **HUNTER-R10-2** fixed, marksmanship.md:69. 31/15/5 leveling and the raid trade (Rapid Killing and Rapid Recuperation out; Lone Wolf, Sniper Shot, two Barrage ranks and Improved Stings 2 in), decoded from both strings.
- **HUNTER-R10-3** fixed, marksmanship.md:80. Improved Stings 2 points, 13%.
- **HUNTER-R10-4** fixed, marksmanship.md:72 (Barrage 15% raid, 5% leveling) and new bullets :82-84 for Sniper Shot, Rapid Killing (rank 2 text: 2 min off Rapid Fire, 20% next Shot) and Rapid Recuperation (rank 2: 50%, 100% once Rapid Killing is consumed).
- **HUNTER-R10-5** fixed, marksmanship.md:153. Improved Stings sentence removed; Efficiency from 15. The SHOULD is applied at :159: Rapid Killing at 25-26 (Rapid Fire trained at 26) and Rapid Recuperation at 33-34, per the shipped order.
- **HUNTER-R10-6** fixed: :48 overview names Sniper Shot; :96 priority sentence after Multi-Shot (matches the APL: after Multi-Shot, next ranged auto over 1.2 s); :105 mana line keeps Sniper Shot; card note written for the 1310786 line (auditor's text, true to the gate and the 15 s own cooldown).
- **Extra, marksmanship.md:52.** The Strengths line called Aimed Shot and Multi-Shot "your biggest shots", false now that Sniper Shot is the hardest hit; it now says "the shots on your shared cooldown" (Barrage does not touch Sniper Shot, so the line stays true).
- **Stat priority, every spec page.** BM, MM and SV Stat priority prose re-read: each gives reasons and the cap rule, points to the table, and has no ranking words; the prose test agrees. No change needed, so BM, SV and the index were not edited.
- **D1 to D4** applied as written: Falcon lines stand; Sniper Shot is taught (R10-1, R10-6); the leveling talents are described as shipped.
- **SHOULDs carried from round 9** (survival.md:65, the Serpent Sting card's "throws away a tick", "the Molten Core tome", BM:81 Deflection): left, wording only and nothing false.
- **New SIM findings:** none. The requested sims (MM BiS and race order on the new raid build, the carried hunter runs) and the Falcon proc LIKELY stand as written above.
