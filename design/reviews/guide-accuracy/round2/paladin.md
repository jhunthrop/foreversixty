# Guide accuracy round 2: Paladin

Auditor: paladin (class index, Holy, Protection, Retribution). Audit tree land/guides-r1 609eb471c, client 1.60.1.70755, engine pin d1d9f2775. Read-only: no guide edits, no sims run.

Files: `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, plus the rotation card notes the pages render (`data/curated/apl/paladin-*.json`).

## Builds decoded (talents[] order in talents/paladin.json; all legal on 70755: 51 points, tier gates and prerequisites met)

The builds have not changed since round 1, and I decoded them again from scratch.

- **Retribution `build`** `050250002/502/0523330100130135` (14/7/30). Holy: Divine Strength 5, Reverence 2, Improved Seals 5, Holy Conduit 2. Protection: Divine Intellect 5, Precision 2. Retribution: Benediction 5, Improved Judgement 2, Conviction 3, Sanctified Judgement 3, Vindication 3, Seal of Command 1, Sacred Arbiter 1, Two-Handed Weapon Specialization 3, Repentance 1, Vengeance 3, Champion of the Light 5.
- **Retribution `raidBuild`** `050050002/503/05233301001301351` (12/8/31). The same, minus Reverence 2, plus Precision 3 and Unsealed Wrath 1.
- **Protection `build`** `05005/15302102513501031/5003` (10/33/8). Holy: Divine Strength 5, Improved Seals 5. Protection: Divine Intellect 1, Redoubt 5, Precision 3, Sacred Duty 2, Templar's Bulwark 1, Shield Specialization 2, Toughness 5, Swift Judgement 1, Improved Righteous Fury 3, Anticipation 5, Holy Shield 1, Iron Creed 3, Shieldsworn 1. Retribution: Deflection 5, Conviction 3.
- **Holy `build`** `50320215012101251/50322/0503` (31/12/8). Holy: Spiritual Focus 5, Healing Light 3, Reverence 2, Purifying Power 2, Divine Favor 1, Illumination 5, Divine Precision 1, Infusion of Light 2, Oath of Light 1, Holy Shock 1, Boundless Light 2, Holy Power 5, Light's Vigil 1. Protection: Divine Intellect 5, Precision 3, Guardian's Favor 2, Sacred Duty 2. Retribution: Benediction 5, Conviction 3.

Prerequisites on 70755: Illumination needs Reverence 2. Boundless Light and Light's Vigil need Holy Shock. Shield Specialization needs Redoubt 5. Swift Judgement needs Templar's Bulwark. Shieldsworn needs Holy Shield. Vengeance needs Seal of Command, which bears out retribution.md:70. Unsealed Wrath needs Repentance.

## Leveling prose against the shipped point order (addon-builds.json, index 0 = level 10)

- **Retribution:** Benediction 10-14, Improved Judgement 15-16, Conviction 17-19, Sanctified Judgement 20-22, Vindication, Seal of Command, Sacred Arbiter, Two-Handed Weapon Specialization, Repentance, Vengeance, then Champion of the Light 35-39, then Holy. retribution.md:74 ("... to Champion of the Light by level 40") matches.
- **Protection:** Templar's Bulwark at 20 and the second Sacred Duty point at 21, Holy Shield at 30, Shieldsworn at 40. protection.md:129 matches.
- **Holy:** Oath of Light at 25 as the 16th point (ladder-first.json paladin-holy row), Divine Precision at 29, Infusion of Light at 30-31. holy.md:112 matches. Round 1's S1 has landed.

## Checked and correct (no finding)

I read every talent line against the 70755 descriptions in talents/paladin.json. Spell levels, ranks, costs, cooldowns and ranges were checked against trainables/paladin.json and spelltips/paladin.json:

- Holy Strike 6 (10 s cooldown, 20 mana at rank 8).
- Judgement 4 (10 s, 10 yd, off the GCD).
- Swift Judgement (60 s, free, off the GCD).
- Seal of Fury and Righteous Fury at 10. Righteous Fury lasts 30 min and costs 30% of base mana. The Seal of Fury Judgement taunts for 3 s.
- Consecration 20 (rank 1 at 135 mana, 8 s cooldown), Exorcism 20 (Undead or Demon only, 15 s), Holy Wrath 50 (Undead and Demons within 20 yd).
- Hammer of the Righteous 40 (6 s, hits up to 3 extra targets), Hammer of Wrath 44 (1.0 s cast, 6 s cooldown, 30 yd).
- Holy Shield ranks at 30, 40, 50 and 60 (10 s cooldown, 10 s duration, 4 charges), Holy Shock (10 s cooldown, ranks at 30, 40, 48 and 56).
- Flash of Light: rank 1 at 20, rank 4 at 42, 1.5 s cast. Holy Light rank 9 is Libram: Holy Light IX (items.json, book-ranks.json).
- Light's Vigil: 1340 mana at 60, 1.5 s cast.
- Blessing of Protection: 5 min, 3 min with Guardian's Favor.
- Divine Shield, Divine Protection and Templar's Bulwark all apply Forbearance.
- Blessing of Kings and Greater Blessing of Kings: 60 min. Heart of the Lion 409580: "increasing total stats by 10% and Attack Power by 40 for all nearby allies".

The rest:

- **Stat priorities:** all three match bis band 60 raid order.
  - Ret: Hit 2.16 > Int 1.07 > Str 1.00 > Crit 0.99 > Agi 0.77 > SP 0.57 > AP 0.41.
  - Prot: Defense > Block > Dodge > Parry > Agility > Block value > Stamina > Hit > Str > Crit > Armor > AP. Expertise is correctly left out because no item carries it.
  - Holy: MP5 > Int > Healing > Crit > Spirit.
- **Racials:** checked against the client descriptions in spelltips/shared.json: Ruthless, The Human Spirit, Will to Survive, Stoneform (10% physical, 8 s), Artisan's Precision, Will of the Forsaken (no immunity), Cannibalize (health and mana), Touch of the Grave (drain).
- **Classes:** combos.json allows Human, Dwarf and Undead only.
- **Elixir of Holy Power:** item 21546 exists in 70755, required level 40, spell 1310077.
- **Rotation prose:** matches the APLs step for step. That covers Ret's 30% mana and 1200 AP twist gates and the Unsealed Wrath timer cue, Prot's 35% Bulwark threshold and the half-mana Consecration gate, and Holy's 40/70/90/80 thresholds with the -5 point pacing.
- **Style:** no em or en dashes and no site, sim or planner talk in the markdown.

**Lay on Hands:** protection.md:109 and holy.md:92 are known and fixed on a follow-up branch, so they are not re-filed. Nothing else on the pages repeats the "costs no mana" claim. holy.md:90 ("your only instant heal outside Lay on Hands"), protection.md:53 and :64, holy.md:50 and the Prot card note at paladin-protection.json:61 say nothing about its cost.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R2-1 | CONFIRMED | rotation (card note) | data/curated/apl/paladin-holy.json:584 (rendered on holy page) | "Oath of Light adds the same flat amount to every rank, so the 35 mana rank 1 heals the tank for well over half of what the top rank does at a quarter of the cost." | spellconst 19750: base 46, +1/level, capped at level 25, so 51. 19943: 303 + 2.6/level, so about 308 at 60. Oath of Light 1328548 adds 1.92/level, about 115 at 60. Under Forever's level-gap rule (round 1 L1) rank 1, learned at 20, keeps none of your healing power at 60. Even with zero healing power, rank 1 lands for 166 against 423, which is 39%, not "well over half". The note also contradicts holy.md:81 and :116 ("Rank 1 is a trap there"). | Note: "On the tank below 90%: your staple. Oath of Light adds the same flat amount to every rank, and rank 4 is the cheapest rank that still keeps nearly all of your healing power at 60." Land it together with the queued action change 19750 r1 to 19941 r4 (round 1 L1(b)) so the card casts what it says. |
| PALADIN-R2-2 | CONFIRMED | rotation (card note) | data/curated/apl/paladin-holy.json:255 (rendered on holy page) | "Rank 1 heals about two thirds as much as the top rank for half its mana" | spelltips 1311606 rank 1: 114 healing for 160 mana. 20930 rank 4: 320 for 325. That is 36% of the heal before healing power, and rank 1 (learned at 30) keeps only about 35% of your healing power at 60 while rank 4 keeps all of it. The note contradicts holy.md:84 and :96 ("Use Holy Shock rank 3 or 4 ... at 60 it keeps only part of your healing power"). | Note: "On anyone below 70%, with the same mana to spare. Use rank 3 or 4: rank 1 costs half as much but has lost most of your healing power by 60. On anyone but your Oath target, Boundless Light adds 30% crit, which Illumination turns back into mana." Land it with the queued action change 1311606 r1 to 20929 r3 at lines 246, 306, 366, 425, 484 and 543 (round 1 L1(b)). |
| PALADIN-R2-3 | CONFIRMED | rotation (card note) | data/curated/apl/paladin-retribution.json:148 (rendered on retribution page) | "every Seal costs a global and over 200 mana" | trainables: Seal of Command rank 5 costs 210 and Seal of Righteousness rank 8 costs 200. Benediction 5/5 (in both Ret builds) is "Reduces the Mana cost of all instant cast spells and abilities by 10%", and both Seals are instant, so they cost 189 and 180. Neither is over 200. | "... every Seal costs a global and about 190 mana, and with less attack power Seal of Command does less than Seal of Righteousness." |
| PALADIN-R2-4 | LIKELY | voice / healing | holy.md:36 | "Oath of Light adds a flat amount to every heal on the tank, so a middle rank Flash of Light heals them almost like the top rank for two thirds of the mana" | At 60 with Oath, rank 4 lands for about 275 plus 95% of your healing power, against 423 plus all of it for rank 6. That is roughly 65% (no gear) to 75% (raid gear) of the top rank's heal, for 90 against 140 mana (64%). It is the more efficient rank, but "almost like the top rank" overstates the heal. | "Oath of Light adds a flat amount to every heal on the tank, so a middle rank Flash of Light on them heals for most of what the top rank does and gets more healing out of every point of mana" |

## Decisions

- **Retribution race: Human, for leveling and raid builds.**
  - **Reason, raid 60:** the racial sims (raid preset, band 60, published gear, 10,000 iterations, seeds 1 and 50001, design/reviews/guide-accuracy/round2/race-sims.md) are a tie, Human 761.98 ±0.51 against Dwarf 761.75 ±0.50.
  - **Reason, bare 60:** Human wins beyond error, 264.17 against 261.03 (+1.2%, errors 0.15). That is The Human Spirit's 10% Spirit paying for more Consecration. The bare 60 weights also show 1% hit and 1% crit at near parity (40.4 against 39.6 per 1%), so Artisan's Precision buys no edge there.
  - **Reason, the rest:** Big Game Hunter only works on Beasts, and Will to Survive is a stun break Dwarf has no answer to. No Forever source names a different winner. Blizzard says the races are tuned for similar offensive power, and wowtbc.gg's Forever guide says every race is within 1% for a spec. With raids tied, the race that wins solo, in dungeons and while leveling is the pick.
  - **Sources:** race-sims.md; races.json and spelltips/shared.json (Ruthless, The Human Spirit, Will to Survive, Artisan's Precision, Stoneform, Big Game Hunter); Blizzard Deep Dive panel recap https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap ; wowtbc.gg, Retribution Paladin PvE Guide, WoW Forever https://wowtbc.gg/warcraftforever/class-guides/retribution-paladin/ (read 2026-11-04: "each race for a spec are within 1% DPS of each other"; no winner named); wow.gg, WoW Forever Retribution Paladin (beta) https://wow.gg/guides/wow-forever-retribution-paladin-beta (no race section).
  - **Guide change:**
    - Frontmatter stays as it is: `build` and `raidBuild` race `human`, `recommendedRaces: [human, undead]`.
    - retribution.md:111: replace the first two sentences ("**Human** is the Alliance pick: ... a hard-hitting pull.") with: "**Human** is the Alliance pick. Ruthless is 1% crit on everything, The Human Spirit's 10% Spirit keeps your mana up while leveling and in dungeons, and Will to Survive breaks a stun. **Dwarf** comes out even in a raid, where Artisan's Precision, 1% hit, is worth as much as Ruthless, and Stoneform takes 10% off physical damage for 8 seconds on a hard pull, but it falls behind everywhere else." Leave the Undead sentences as they are.
  - **Closes:** PALADIN-12 / S2.

## Round 1 research answers re-checked against the client text

- **O1 (Exorcism never crits):** stands. The 70755 descriptions (879 to 10314) say nothing about crit either way, and the beta-log evidence is unchanged.
- **O2 (Lay on Hands costs no mana):** WRONG. The 70755 descriptions of 633, 2800 and 10310 read "Drains all of the Paladin's remaining Mana when used, but does not interrupt Mana regeneration." This is already corrected on the follow-up branch, so it is not re-filed. Its "needs sim round" item (drop `SpendMana` in lay_on_hands.go) must be withdrawn: the engine's drain is right.
- **O3 (Kings and Heart of the Lion do not stack):** the forum source checks out. The thread discusses the SoD rune under its earlier name, Aspect of the Lion: Kings replaces it on the target, and only the hunter's own extra 10% stays. The 70755 Heart of the Lion description (409580) is the same 10% total stats plus 40 attack power, so the guides' "same 10% to every stat, do not stack" holds.
- **S1 (Oath of Light at 25):** confirmed. Oath of Light's per-level points (6.67 / 1.92 / 3.83) are in spellconst 1328548, and the shipped order and ladder-first.json now take it at 25.
- **L1 (level-gap rule):** the guide's rank advice matches it. Flash of Light rank 4 (learned 42) keeps 95% of your healing power at 60, rank 3 (34) about 55%, and Holy Shock rank 3 (48) keeps all of it.

## Sim findings

- **Withdraw round 1's O2 engine item.** Keep `paladin.SpendMana(sim, paladin.CurrentMana(), ...)` in sim/paladin/lay_on_hands.go. The 70755 description confirms the drain.
- **PALADIN-R2-1 and R2-2** depend on round 1's queued Holy card rank changes (Flash of Light on the tank 19750 to 19941, Holy Shock 1311606 to 20929) and the engine's level-gap curve. Land the notes with the actions.

## Owner questions

None. The Retribution race is decided above. No paladin item is a true taste call.

## Sources

- 70755 client data: talents/paladin.json, trainables/paladin.json, spelltips/paladin.json, spelltips/shared.json (racials), spelltips/hunter.json (Heart of the Lion 409580), spellconst/paladin.json, races.json, combos.json, items.json, simconsumes.json, addon-builds.json, bis/paladin-*.json; data/curated/ladder-first.json, rank-choices.json, book-ranks.json, apl/paladin-*.json; engine d1d9f2775 proto/common.proto:491-500.
- design/reviews/guide-accuracy/round2/race-sims.md.
- Blizzard, World of Warcraft Forever Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- wowtbc.gg, Retribution Paladin PvE Guide, WoW Forever: https://wowtbc.gg/warcraftforever/class-guides/retribution-paladin/
- wow.gg, WoW Forever Retribution Paladin (beta): https://wow.gg/guides/wow-forever-retribution-paladin-beta
- Blizzard forums (SoD), Blessing of Kings doesn't stack with Aspect of Lion BUT: https://us.forums.blizzard.com/en/wow/t/blessing-of-kings-doesnt-stack-with-aspect-of-lion-but/1738967

## Counts

CONFIRMED 3 (all rendered card notes), LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0. Guide markdown proper: 0 confirmed errors outside the Lay on Hands lines already being fixed.

## Fix log (round 2)

- **PALADIN-R2-1** fixed, data/curated/apl/paladin-holy.json:584. The card still casts Flash of Light rank 1 on the tank (action unchanged until the sim round), so the note now states only what is true of any rank: "On the tank below 90%: your staple. Oath of Light adds the same flat amount to every rank you cast on the tank." The false "well over half of the top rank" claim is gone. The auditor's rank 4 wording lands with the queued action change 19750 r1 to 19941 r4.
- **PALADIN-R2-2** fixed, data/curated/apl/paladin-holy.json:255. Dropped "Rank 1 heals about two thirds as much as the top rank for half its mana"; the note keeps the threshold and the Boundless Light / Illumination line. The "use rank 3 or 4" wording lands with the queued action change 1311606 r1 to 20929 r3.
- **PALADIN-R2-3** fixed, data/curated/apl/paladin-retribution.json:148: "over 200 mana" to "about 190 mana" (Benediction 5/5 on instant Seals: 189 and 180).
- **PALADIN-R2-4** (LIKELY) fixed, web/src/content/guides/paladin/holy.md:36, verified from the auditor's spellconst figures (rank 4 with Oath about 65 to 75% of rank 6 for 64% of the mana): now "heals for most of what the top rank does and gets more healing out of every point of mana".
- **Decision, Retribution race Human** applied, web/src/content/guides/paladin/retribution.md:111, the first two sentences replaced with the decided text; Undead sentences kept. Frontmatter (`human` in both FS1 strings, `recommendedRaces: [human, undead]`) already matched, unchanged.
- **Source labels**: every "Forever Sixty" label in index.md:17, holy.md:16-25, protection.md:30-39, retribution.md:17-29 rewritten to player words with no paths ("Forever Sixty, paladin talents", "..., what each stat is worth", "..., raids by phase", "..., races", "..., which races can be paladins"). URLs unchanged.
- **Planner links**: none on the paladin pages; nothing to change.
- **Lay on Hands**: already fixed on main (ab42101f7); not touched.
- **Left**: the file-level `notes` description in paladin-retribution.json:4 still says "200+ mana at 60 (Seal of Righteousness rank 8 is 200, Seal of Command rank 5 is 210)". That is the base cost and is a developer description, not a rendered card note; left for the sim round's note cleanup.

### New SIM findings

None beyond the auditor's: land the Holy card rank changes (19750 r1 to 19941 r4 on the tank; 1311606 r1 to 20929 r3 at the six Holy Shock lines) and then restore the rank-specific note wording from PALADIN-R2-1 and R2-2.
