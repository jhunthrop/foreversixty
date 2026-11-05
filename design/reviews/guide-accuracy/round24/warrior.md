# Guide accuracy round 24: Warrior

Audited on the round 24 audit tree (ea9adea1c: main after guide round 23 landed, including the sim round 21 merge 617d4e9c4; engine pin 170a59c39 in `sim/enginever/version.go`, engine source at `gr24/engine`). I checked against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`).

Files read:
- Guide pages: `web/src/content/guides/warrior/{index,arms,fury,protection}.md`.
- Cards: every note in `data/curated/apl/warrior-{arms,fury,protection}.json`.
- Addon rows: every rendered row of `addon-data.json`, `rotations/warrior-{arms,fury,protection}`.
- Clause rendering: `web/src/lib/guides/rotation-lists.ts` (`clauseOf`).
- Client data: `talents/warrior.json`, `spelltips/{warrior,shared}.json`, `trainables/warrior.json`, `races.json`, `addon-builds.json`.

Reviews read first:
- Sim round 21's `design/reviews/2026-11-05-sim-audit-round21/warrior.md` and `integration.md`.
- Round 23's `warrior.md`.
- The ledger's Round 19 to 23 sections.

How notes render (as rounds 20 to 23 documented):
- The addon row shows the note up to the first ". ".
- `clauseOf` (the guide card and the level 60 lists) shows it up to the first ". ", ": ", "; " or " -- ".
- The sim drawer shows the whole note.

## New work since round 22's tree (24ee457d0)

I ran `git diff 24ee457d0 HEAD` on the warrior pages, the warrior cards, `addon-builds.json`, warrior BiS, `races.json` and `spelltips/warrior.json`:
- Only the three guide pages moved.
- The `rotations/warrior-*` blocks of `addon-data.json` are byte-identical to 24ee457d0.
- No card, order, rank row or talent file changed.

**Sim round 21 (audited first):**
- **protection.md:52, new Strengths bullet.** "Defensive Stance raises all your threat by 80%, and Defiance stacks another 15% on top with a shield on." The 80% comes from Blizzard's 3 Nov Forever notes and Wowhead's live tooltip for 71 and 7376 (sim S1). Defiance's 15% is the 70790 talent text (12792 rank 3, "an additional 15% while a shield is equipped"). Wowhead's live tooltip for 12792 also says 15%; I re-read it on 5 Nov ("Increases all threat generated in Defensive stance by an additional 15% while a shield is equipped"). "Stacks ... on top" matches the client formula `($7376s3+100)*($12792s1/100+1)-100`. The stance is trained at 10, and the line describes the finished build. OK.
- **protection.md:53, merged bullet.** "Shield Block is off the global cooldown and up most of the fight, and every block, dodge and parry pays rage back through Shield Specialization and Master of Defense." Shield Block is 7 s on a 5 s cooldown with `gcd_ms` 0 (2565). At 5/5, Shield Specialization gives "a 100% chance to generate 5 Rage when you Block". At 2/2, Master of Defense gives "a 100% chance to generate 5 Rage when you Dodge or Parry while a shield is equipped" (1310316). Both are complete by 30 on the shipped order (Shield Specialization 12 to 16, Master of Defense 29 to 30). "Up most of the fight" stays a carried Should (two-attack cap). OK.
- **protection.md:71, Defiance talent line.** "on top of the stance's own 80%, so the two together come to a little over double": 1.8 x 1.15 = 2.07. OK.
- **Left behind by the 80% change?**
  - No line on any warrior page states 30% or reasons from the old size.
  - Bastion's "10% more threat from everything that hits" (:70) is multiplicative, so it is unchanged.
  - Revenge "most threat per rage" and Shield Slam "a very high amount of threat" are relative, and the stance scales them equally.
  - The race lines (:113, :115) rest on damage taken, which the change does not touch (sim: the stance's -10% is untouched).
  - The hover tooltip on Defensive Stance still renders the cached 30% (`spelltips` 71). That is component and data lane, already in sim round 21's "For other lanes", so it is not a guide finding.
- **Sim round 21's "For the guide lane" item was not applied by round 23's merge:** fury.md:112 "edges Night Elf by a hair". Filed below as **WARRIOR-R24-1**.

**Round 23's fixes (re-checked as new work):**
- **R23-1, fury.md:114.** "If your race has a damage racial, Elune's Light, Blood Fury, Berserking or Eureka!, press it on cooldown and line it up with Death Wish when the timers allow."
  - The list is complete against `races.json` and `spelltips/shared.json` racials. Every other race's actives are utility:
    - Will to Survive and Perception (Human).
    - Stoneform and Find Treasure (Dwarf).
    - Shadowmeld (Night Elf).
    - Will of the Forsaken and Cannibalize (Undead).
    - War Stomp and Cultivation (Tauren).
    - Escape Artist (Gnome).
    - Rapid Regeneration (Troll).
    - Walk on Air, Read Leyline and Skysight (Skyborne).
  - The cooldowns let them line up with Death Wish's 3 minutes: Elune's Light and Berserking are 3 min, Blood Fury and Eureka! are 2 min.
  - "On cooldown" against "never inside Recklessness" for Elune's Light: the same paragraph's last sentence ("a crit racial does not") and :92 carry the exception. That is a Should, not a finding (below). OK.
- **R23-2, arms.md:131 and fury.md:134.** "Victory Rush whenever it is up: it heals 10%, works for 20 seconds after a kill and comes back every 30 seconds, so when the next mob is close, press it right after the Charge."
  - The client text for 402927 matches: "healing you for 10% of your maximum health. Only useable within 20 sec after you kill a non-trivial enemy", `cooldown_ms` 30000.
  - It is trained at 20 (`trainables` 402927), and both lines sit under "Levels 20 to 60".
  - Charge (out of combat only, 100 / 6178 / 11578) and then Victory Rush is the order melee range needs.
  - On Fury after 36 the page's own :84 covers Charge or Intercept. OK.

## Orders (index 0 = level 10)

`addon-builds.json` is unchanged since round 22. The decode is as in round 23's report, and I re-checked it for every line below. Protection: Improved Bloodrage 10-11, Shield Specialization 12-16, Anticipation 17-21, Improved Revenge 22-24, Improved Thunder Clap 25-27, Last Stand 28, Master of Defense 29-30, Defiance 31-33, Improved Sunder Armor 34-36, Concussion Blow 37, Focused Rage 38-39 and 41, Shield Slam 40, Bastion 42-46, Improved Heroic Strike 47-49, Deflection 50-54, Improved Rend 55-57, Deep Wounds 58-60.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| WARRIOR-R24-1 | CONFIRMED | page prose (a race lead the measurement does not show; routed request now resolved) | fury.md:112 | "On the Alliance, Skyborne's Wind Blessed, 2% haste that never turns off, edges Night Elf by a hair, but it needs the paid Skyborne Heroic Pack" | The measurement rounds 19 to 23 asked for has been run. Sim round 21, Fury F2: raid 60, frozen BiS, guide raid build, 10,000 iterations. Night Elf scored 918.5 / 919.8 and both Skyborne races 918.7 / 920.1 at seeds 1 / 50001. That is +0.02% and +0.03%, inside the ±0.5% error. The sim lane's call: "'level', as arms.md:109 says, is the honest word" (For the guide lane). Round 18's runs gave the same 0.03 / 0.04%. The sentence claims a lead the data cannot resolve, and that lead is for the paid race the owner rule says is never the pick. The same page calls the Horde case "level" on equal evidence, and arms.md:109 calls the Alliance Arms gap "level". This was carried as a Should while it waited on the sim. The sim has now decided it, so it is a finding. | "On the Alliance, Skyborne's Wind Blessed, 2% haste that never turns off, is level with Night Elf, but it needs the paid Skyborne Heroic Pack, so take it only if you already own it." |

No card note needs a change, so `addon-data.json` and `Data.lua` do not need regenerating.

Counts: CONFIRMED 1, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## Shoulds (wording, not findings)

- **New, fury.md:114** "press it on cooldown". Elune's Light is in the list, and :112 and :92 say never inside Recklessness. The paragraph's own last sentence ("a crit racial does not") carries the exception, so the line is not false. Suggested: "...press it on cooldown, a crit racial outside Recklessness, and line it up with Death Wish...".
- **New, fury.md:63** Flurry "comes out at the faster speed without using one up". That is true of a Heroic Strike that lands, is dodged or is parried. One that misses outright spends a charge (`sim/warrior/talents.go` `spendsFlurryCharge`, from Tawl's 70450 log). At the special hit cap :100 describes, raid Heroic Strikes almost never miss.
- **New, the sweep's "first sentence at 60" lens.** The Fury Berserker Rage addon line and clause read "From 36, under 80 rage". The Magmadar, Onyxia and Nefarian hold is in the second sentence. The first sentence is exactly the row's condition, the card has no encounter gate, and :92 teaches the hold. Not false.
- **Carried from round 23 (unchanged lines):**
  - arms.md:53 "Rage starved without a raid".
  - protection.md:89 Bloodrage with the shout branch.
  - arms.md:85 "first three globals".
  - arms.md:91 "biggest hit of the fight".
  - fury.md:86 and the Fury Reckless Swings note's "healers".
  - arms.md:109 and fury.md:112 "Recklessness, which already makes every hit crit" (glancing).
  - protection.md:103 "only a few late sets".
  - arms.md:131 Sweeping Strikes and Whirlwind undated.
  - fury.md:134 Bloodthirst undated.
  - protection.md:69 "Your biggest button".
  - protection.md:75 "three buttons you press most".
  - The Arms Reckless Swings note "white swings for the next 8 seconds" (four within 8 s).
  - The Protection Shield Block note's two-attack cap (also protection.md:53 "up most of the fight").
  - Fury Execute "First".
  - Protection Defensive Stance note at 10 to 27.
  - Arms Hamstring "never from 25% down" before Execute.
  - protection.md:113 "no race's health swings less".
  - arms.md:125 skipping the level 19 point.
  - fury.md:74 "Fury never casts Rend".
  - arms.md:115 "other pair".
  - fury.md:54 "debuff".
  - fury.md:100 Edgemaster wording.
  - "From 36" lead.
  - The Arms and Protection Heroic Strike notes' "above 90" and "above 40" against the cards' 90-or-more and 40-or-more.

## Checked and right this round (round 24 lenses)

- **Class index teaching a rotation the cards no longer use.**
  - index.md has no rotation. Its only play line (:22, Charge or Bloodrage to open) matches the Arms pull (Charge) and the Fury raid pull (Bloodrage).
  - :26 to :35 were checked against the client:
    - Tactical Mastery 1310185 is trained at 14 and keeps 10 rage, plus 15 from Improved Tactical Mastery.
    - Victory Rush at 20.
    - Slam 15 s, and 10 s with a 1 s cast under Improved Slam 2 ("no longer interrupts or delays your melee swing").
    - Shield Block 7 s / 5 s.
    - Shield Wall 60%, 12 s, 15 min, with Improved Shield Wall -11 min.
    - Death Wish 5% damage taken.
- **Reasons against the numbers.**
  - Reckless Swings (fury.md:71): "20 rage every 15 seconds is a Heroic Strike". 4 x 5 = 20 is at least the 15 cost at 31 to 41 and the 12 cost from 44.
  - Arms Reckless Swings "two or three swings with a slow two-hander": 8 s at a 3.4 to 3.8 s speed, plus haste.
  - The Improved Bloodrage pull: 30 against 21 at 60.
  - Improved Shield Wall swap: 6 to 2 = "4 more rage", and 15 to 4 minutes.
  - Focused Rage on Revenge: 5 to 2.
  - Impale: 220%.
  - Enrage: 2 points 4% (Fury), 5 points 10% (Arms raid).
  - Anger Management: 10 rage a half minute.
  - Improved Charge: +3 at 17 and +6 at 18.
  - Demoralizing Shout: 204 at 60.
  - Defiance with the stance: 2.07.
  - Challenging Shout: 10 yd, 6 s, 10 min (1161).
- **"Only" and "every" claims against the same page.**
  - grep of "only", "never", "every", "always", "best", "most" and "biggest" on all four pages.
  - Spearing Strike "only works in Battle Stance": engine mask, and the client "Deals ... ".
  - Frozen Heart "only an enchanter can make": bind on pickup, Enchanting.
  - "Arms only needs one weapon".
  - "take it only if you already own it".
  - Thunder Clap "only costs 4 more".
  - Taunt "only to take a target back".
  - None is contradicted by another spell on its page.
- **Level 60 rule in the first sentence.** I read every warrior note's `clauseOf` output and addon sentence at band 60. Each first sentence carries its row's own condition:
  - The Arms swap "From 50, below 20% and once your first Execute has emptied the bar".
  - Spearing Strike and Reckless Swings "raid build only".
  - Hamstring's totem and 25% rules.
  - Heroic Strike's thresholds.
  - Death Wish and Recklessness timing.
  - Fury Heroic Strike's execute hold.
  - Protection Battle Shout "Whenever it drops": another warrior's shout is the same aura 11551, so it has not dropped.
  - The only rule left to a later sentence is the Berserker Rage fear hold (Should above).
- **Race lines** against round 19's measured order and sim round 21's F2 run: arms.md:109 and :111, fury.md:112 (R24-1) and :114, protection.md:113 and :115.
- **Leveling paragraphs** against `trainables` and the orders (arms.md:121 to :131, fury.md:124 to :134, protection.md:125 to :135). Every spell named is trained by the level stated, and every talent level matches the order.

## Decisions

- **D1. R24-1 is CONFIRMED.** The comparison was routed to the sim with the condition "only if the wording is kept". The sim ran it and found no lead, so the page now states a lead that is not there, for the paid race.
- **D2. Defiance stays 15% on the page.** Blizzard's 3 Nov note says 7/13/20%, but the 70790 talent text, the trait curve and Wowhead's live tooltip for 12792 (re-read 5 Nov) all say 15%. Wowhead's 7376 figure of 107% = 1.8 x 1.15 shows the stance's 80% is live and Defiance's 20% is not. Not a finding. It stays with sim round 21's open S2.
- **D3. The Defensive Stance hover tooltip still reads 30%, beside the page's 80%.** This is component and data lane (hover tooltip text is excluded; sim round 21 "For other lanes"). Not filed here.
- **D4. The fury.md:114 Elune's Light "on cooldown" is a Should.** The sentence after it states the exception.
- **D5. Rows that render before their talent or spell exists stay with the component lane,** as in rounds 15 to 23.

## Sims requested

- None new. Carried, only if the wording is kept:
  - Protection raid preset at 60, average landed Revenge (11601) against Shield Slam (23925) (protection.md:69).
  - Protection race TMI on one shared raid preset (protection.md:113).
  - Arms level 60 build, rage per second by source, raid against bare (arms.md:53).

## Engine questions

- Carried, none new:
  - S2: Defiance 7/13/20 against the client's 5/10/15.
  - W23-1 and W22-1: trainer rows at 60 for Battle Shout rank 7, Heroic Strike rank 9 and Revenge rank 6 against the cards' ranks.
  - W21-1: Classic weapon speeds.
  - X1.
  - P5 / S3: Revenge, Shield Slam and Thunder Clap threat constants.
  - AC: Arcanite Champion PPM.
  - C2 / S4: Deep Wounds' double attacker multiplier.
  - Whether an outright-missed Heroic Strike spends a Reckless Swings charge.

## Sim findings

- None new. Sim round 21's S1 (Defensive Stance 80%) is reflected correctly on protection.md.

## Sources

- Client 1.60.1.70790:
  - `spelltips/warrior.json` (71, 2565, 402927, 1161, 871, 12328, 1329555, 1310222, 1464 to 11605, 100 / 6178 / 11578).
  - `spelltips/shared.json` racials.
  - `talents/warrior.json` (Defiance 12792, Shield Specialization 12298, Master of Defense 1310316, Improved Slam, Impale, Enrage, Improved Execute, Improved Berserker Rage, Flurry).
  - `trainables/warrior.json`, `races.json`, `addon-builds.json`, `addon-data.json`.
- Curated: `data/curated/apl/warrior-{arms,fury,protection}.json`. `web/src/lib/guides/rotation-lists.ts`.
- Engine 170a59c39: `sim/warrior/talents.go` (`spendsFlurryCharge`, `countsAsWhiteSwing`), `white_swing_procs_test.go`.
- Reviews:
  - `design/reviews/2026-11-05-sim-audit-round21/warrior.md` (S1, S2, F2, P2, For the guide lane) and `integration.md`.
  - `design/reviews/guide-accuracy/round23/warrior.md`.
  - `ledger.md` Rounds 19 to 23.
- Web:
  - Wowhead Forever tooltip, Defiance, https://nether.wowhead.com/forever/tooltip/spell/12792 (read 5 Nov: "an additional 15% while a shield is equipped").
  - Blizzard, WoW Forever patch notes thread, https://us.forums.blizzard.com/en/wow/t/2360696 (3 Nov: Defensive Stance 80%, Defiance 7/13/20), as cited by sim round 21.

## Note sweep

One line per row note:
- spec, row, and the bands it renders;
- each named talent, spell, rank or value, with the level the player has it (L = leveling order, R = raid order);
- the verdict.

No note or row changed since round 23. I re-read every line at every band under this round's lenses: reason against the numbers, "only", first sentence at 60.

Arms (20 bands, 10 to 60)

- **Berserker Stance swap** (30 to 60). OK.
  - "From 50": Recklessness at 50, and the card asks `spellIsKnown` 1719.
  - Improved Tactical Mastery 3 (L 21, R 18) gives 25 rage.
  - Mortal Strike (40) and Execute (24) work in Berserker Stance; Overpower and Rend do not (masks).
  - The first sentence carries the rule.
- **Bloodrage** (10 to 60). OK.
  - 10 + 10 over 10 s, 60 s.
  - A quarter per point (12301): L 54 / 55, R 60.
- **Battle Shout** (10 to 60, ranks 6673 to 11551): 10 rage at every rank, 20 yards, raid. OK.
- **Rend** (10 to 60). OK.
  - The bleed ignores armor.
  - Bloodthrill is named only "once Bloodthrill is in" (L 33, R 31).
  - Main-hand hits on your Rend.
- **Overpower** (12 to 60). OK.
  - Mortal Strike "from 40"; a dodge opens it.
  - Bloodthrill "once the talent is in".
  - Improved Overpower 2 (L 23, R 20) "most of the time": 50% plus base crit.
- **Mortal Strike** (40 to 60): 6 s at every rank; healing cut 50%. OK.
- **Reckless Swings** (10 to 60, component row; talent R 57). OK (Should on the four-attack cap).
  - Raid build only, after Mortal Strike (40).
  - 8 s, 5 rage, a quarter of base health.
- **Execute** (24 to 60): below 20%; the first one spends the bar; Mortal Strike "from 40". OK.
- **Slam** (20 to 60). OK.
  - Improved Slam L 42, R 39. Its talent text: "no longer interrupts or delays your melee swing", so without it Slam does.
  - Mortal Strikes "from 40".
- **Spearing Strike** (10 to 60, component row; talent R 25). OK.
  - Raid build only, Battle Stance.
  - 75% weapon damage, twice that on dragons and giants (client: Giants, Dragonkin, mounted).
- **Hamstring** (10 to 60, ranks 1715 to 7373). OK (Should on 25% before Execute).
  - Windfury Totem gate.
  - 10 rage at every rank; never from 25% down.
- **Heroic Strike** (10 to 60): 90 (Should on "above"); Overpower 12, Execute 24, Mortal Strike 40. OK.

Fury (19 bands; leveling and raid orders identical)

- **Bloodrage** (10 to 60): Improved Bloodrage rank 1 at 59, rank 2 at 60. OK.
- **Death Wish** (10 to 60; talent 30, rows 10 to 28 component): second one fits at 210 s; execute phase; last 30 s. OK.
- **Recklessness** (50 to 60): execute phase or last 15 s; once a fight (1800000 ms). OK.
- **Battle Shout** (10 to 60): as Arms. OK.
- **Execute** (24 to 60): first below 20%; every point of rage; Bloodthirst "from 40". OK.
- **Bloodthirst** (40 to 60): "the button you lose the most by delaying". OK.
- **Whirlwind** (36 to 60): four targets, both weapons (1680); Bloodthirst "from 40". OK.
- **Reckless Swings** (10 to 60; talent 31). OK.
  - Behind Whirlwind "once you have it" (36).
  - Four swings, 5 rage, a quarter of base health.
  - "healers cover" is a Should.
- **Berserker Rage** (36 to 60). OK (Should on the first sentence at 60).
  - "From 36"; Improved Berserker Rage 2 at 35 makes it 10 rage; 30 s (18499).
  - Hold for the fear in the second sentence.
- **Overpower** (12 to 32, levels 12 to 35): Battle Stance before 36. OK.
- **Heroic Strike** (10 to 60): above 60; the 25% stop applies once Execute (24) is known; no global. OK.

Protection (24 bands; leveling and raid orders identical)

- **Defensive Stance** (10 to 60). OK.
  - "the threat bonus" carries no number, so it is unaffected by 30 to 80%.
  - Revenge 14, Shield Block 16, Shield Wall 28 (Should carried).
- **Last Stand** (10 to 60; talent 28): below 35%; 30% for 20 s, lost after (12975). OK.
- **Shield Wall** (28 to 60): below 40%; 60%, 12 s, 15 min (871). OK.
- **Bloodrage** (10 to 60): before the pull, then under 60; Improved Bloodrage rank 1 at 10, rank 2 from 11. OK.
- **Shield Block** (16 to 60): 75% for 7 s, off the global; every block lights Revenge (14). OK (Should carried).
- **Revenge** (14 to 60, ranks 6572 to 11601): dodge, parry or block; most threat per rage (5, then 2 from 41 with Focused Rage). OK.
- **Shield Slam** (40 to 60, ranks 23922 to 23925): "a hard hit", "very high amount of threat" (client text), grows with block value. OK.
- **Thunder Clap** (10 to 60, 6343 to 11581): 20% slow at every rank; refresh under 3 s (durations 10 to 30 s). OK.
- **Demoralizing Shout** (14 to 60, 1160 to 11556): attack power stripped; does not stack with Demoralizing Roar; refresh under 3 s. OK.
- **Battle Shout** (10 to 60): "Whenever it drops" is true at 60 with the same aura from another warrior. OK.
- **Sunder Armor** (10 to 60, 7386 to 11597): five stacks, then threat. OK.
- **Heroic Strike** (10 to 60): above 40 (Should), queued, no global. OK.
