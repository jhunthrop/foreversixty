# Guide accuracy round 17: Hunter

Auditor: hunter. Pages: class index, Beast Mastery, Marksmanship and Survival, plus the card notes in `data/curated/apl/hunter-*.json` and their rendered rows.
- Tree: gr17/audit at 919f7e8a7 (main: sixteen guide rounds plus sim round 18). Engine pin 6ea675594 (`sim/enginever/version.go`), source read from gr17/engine.
- Sim round 19 (not on main): r19/int at 04d7da663, report `design/reviews/2026-11-04-sim-audit-round19/hunter.md`. Its guide diff against 919f7e8a7 touches `survival.md` only (lines 47, 63, 71, 86, 140, 142 and a new paragraph after 144). No hunter APL changed on r19.
- Read first: the round 1 to round 16 hunter reports with their fix logs, ledger.md, the round 18 and round 19 sim reports and round 19's integration.md.

## What moved since round 16

`git diff 2c3135a2e 919f7e8a7` on the hunter inputs:
- **Guide text:** only round 16's four fixes (BM:99, MM:105, the three Best in slot sentences, the three "From 40 you can wear mail" lines, index.md:37) and sim round 18's marksmanship.md:163 (rank 1 again from 52).
- **Card notes:** the Beast Mastery Serpent Sting note (R16-3). The actions are untouched.
- **`rank-choices.json`:** Marksmanship Aimed Shot rank 1 now covers 28-49, 52-59 and bare 60 (sim round 18 M2).
- **The 70755 data job (1904a0e59):** `items/hunter.json` gained class, subclass and inventory-type fields on every row (no name, bind or stat change on any item the pages name); `spellconst/hunter.json` gained one spell, 1328676 "[DNT] Hunter Pet Scaling (Old Testing Backup)", a test aura no page quotes; `hotfix-push.json` counts moved to match. `spelltips/hunter.json`, `talents/hunter.json`, `trainables/hunter.json`, `races.json`, `enchants.json`, `simconsumes.json` and every `bis/hunter-*.json` are byte-identical to round 16's tree.
- **Engine:** one hunter commit between the pins (8c24d6202, Summon Hawk's 18 s hawks). No guide build takes Summon Hawk, so no page number moves.

Every line of the four pages and every card note was read again in full against the 70755 files.

## Builds and leveling order

The frontmatter build strings are unchanged since round 13, so round 16's decode stands (BM 31/20/0 leveling and raid, MM 15/31/5 and 5/33/13, SV 5/11/35; all legal at 51 points, every prerequisite met: Intimidation needs Bestial Swiftness, Bestial Wrath needs Intimidation, Frenzy needs Ferocity 5, Mortal Shots needs Careful Aim 5, Rapid Recuperation needs Rapid Killing 2, Sniper Shot needs Trueshot Aura, Counterattack needs Deterrence, Lacerating Strikes needs Expose Prey 2).

The shipped orders (`addon-builds.json`, index 0 = level 10, entries are tree, tier and column one-based) decoded by name:
- **BM leveling:** Deadly Aspects 10-14, Focused Fire 15-16, Endurance Training 17-20, Bestial Swiftness 21, Unleashed Fury 22-26, Ferocity 27-31, Intimidation 32, Bestial Discipline 33-34, Frenzy 35-39, Bestial Wrath 40, Lethal Attacks 41-45, Improved Stings 46-47, Efficiency 48-51, Careful Aim 52-56, Rapid Killing 57, Trueshot Aura 58, Mortal Shots 59-60. Matches BM:147 and BM:159.
- **MM leveling:** Lethal Attacks 10-14, Efficiency 15-19, Careful Aim 20-24, Rapid Killing 25-26, Trueshot Aura 27, Mortal Shots 28-32, Rapid Recuperation 33-34, Barrage 35, Ranged Weapon Specialization 36-40, Deadly Aspects 41-45, Endurance Training 46-48, Focused Fire 49-50, Unleashed Fury 51-55, Improved Tracking 56-60. Matches MM:153 and MM:159.
- **SV:** Improved Tracking 10-14, Savage Strikes 15-16, Survivalist 17-19, Clever Traps 20-21, Surefooted 22-24, Deterrence 25, Predator's Edge 26-29 and 31, Strider Kick 30, Counterattack 32, Resourcefulness 33-34, Expose Prey 35-36, Survivalist's Discipline 37-38, Lightning Reflexes 39 and 41-44, Lacerating Strikes 40, Deadly Aspects 45-49, Lethal Attacks 50-54, Careful Aim 55-59, Lone Wolf 60. Matches SV:134 and SV:140 ("one of the last points"; it is the last, see the round 19 table). Every tier opens on time (Strider Kick at 30 with 20 Survival points spent, Lightning Reflexes at 39 with 29, Lacerating Strikes at 40 with 30).

The Leveling tree component (`web/src/lib/guides/leveling-builds.ts`) is a line-for-line port of the ladder that writes `addon-builds.json`, so the prose matches what the tree shows.

## Findings

None. No line on the four pages, and no card note, states something false on 70755 or contradicts the page.

## Covered by sim round 19

Lines that would be findings on main 919f7e8a7 but are already rewritten in round 19's guide diff (`git -C r19/int diff 919f7e8a7 HEAD -- web/src/content/guides/hunter`). The r19 wording was checked as new work.

| Main line | Why it would be a finding on main | r19 wording (quoted) | r19 wording checked |
| --- | --- | --- | --- |
| survival.md:140 "Keep a pet out while you level: Lone Wolf is one of the last points this build takes, and until then the pet holds aggro and adds damage. Dismiss it the moment you take Lone Wolf." | Measured (r19 H1, 10000 iterations, seeds 1 and 50001, both factions): out in the world at 60 the pet leads Lone Wolf at every fight length, +0.5 to +2.3%, and takes the hits; Lone Wolf leads only in a raid (+2.4 to +2.7%). "Dismiss it the moment you take Lone Wolf" sends a questing hunter the wrong way. | "Survival levels in melee. Keep a pet out while you level: Lone Wolf is the last point this build takes, at 60, and until then the pet holds aggro and adds damage. Lone Wolf pays off in a raid, where you get the raid buffs and your pet gets none, so your own hits carry more of the damage. Out in the world at 60 a pet still adds a little more damage than Lone Wolf does and takes the hits for you, so keep it out while you quest and farm. The talent does nothing while your pet is out, so dismiss the pet before a raid pull and call it back afterwards." | Right. Lone Wolf is `addon-builds.json` index 50 (level 60), the last point. Lone Wolf's text: "while you do not have an active pet" (415370). The measured direction and size match "a little more". |
| survival.md:47, 63, 71, 86 ("and no pet", "Melee DPS without a pet", "for a melee hunter the talent beats the pet", "Dismiss your pet before the pull") | True on main in a raid, which is the context of each line; r19 narrows them to the raid so they do not contradict the new world advice. | "no pet in a raid"; "Melee DPS, with your pet dismissed"; "so in a raid the talent beats the pet for a melee hunter. Out in the world the pet is worth a little more than the talent and takes the hits, so keep it out there; the point waits for your next raid pull"; "In a raid, dismiss your pet before the pull". | Right. |
| survival.md:142 (the leveling loop had no Immolation Trap) | Not a finding on main (nothing false), noted because r19 adds the line. | "Throw Immolation Trap under anything that will live long enough to burn." | Right. Immolation Trap trains at 16 (`trainables/hunter.json` 13795). |
| survival.md, new paragraph after 144 (Raptor Strike rank 5 at 46-47) | New r19 text, no main counterpart. | "At 46 the new Immolation Trap rank costs more ... Rank 6 of Raptor Strike, trained at 40, is the one to use until then. For 46 and 47, keep rank 5 on your bar instead: most of the hit is your weapon, so it does nearly what rank 6 does for less mana, and the mana it saves goes into traps. Training a new rank upgrades the bar button on its own, so drag rank 5 back from the spellbook. At 48 put rank 7 on the bar." | Right on every value: Immolation Trap rank 4 at 46 costs 190 against rank 3's 135 (14304, 14303); Raptor Strike rank 5 is +35 for 55 mana at 32, rank 6 +40 for 70 at 40, rank 7 +55 for 85 at 48 (14263, 14264, 14265; `trainables/hunter.json`); the row is 14266 rank 5 at 46-47, measured beyond error (r19 S1). |

## Shoulds (wording, not findings)

- **New, r19 survival.md:140:** "where you get the raid buffs and your pet gets none" reads as if the pet gets nothing at all. Your own Heart of the Lion does reach it (settled, sim round 17 H2; the page's own line 71 says "from other players"). "and your pet gets none of theirs" keeps the sentence true to the settled rule.
- **New, SV:124 (Professions):** the Survival set also wears two Leatherworking pieces that bind when equipped (Black Dragonscale Helm 252605, Dawn Armor 252483, both factions, `bis/hunter-survival.json` band 60 raid). The paragraph says nothing false, but Beast Mastery and Marksmanship tell the reader to buy their Leatherworking pieces, and Survival could say the same in one clause.
- **Carried from round 16:** MM:159 "so you have it for most hard pulls"; MM:94 five minutes on the leveling build; SV:65 "smaller than Era's"; the Serpent Sting cards' "throws away a tick"; "the Molten Core tome"; BM:81 naming Deflection; SV:118 Orc and Troll "level" with only Orc tagged; BM:157 "skip ranks 2 to 5 at the trainer"; SV:80 "a third melee special".

## Checked and right

Everything below was re-read this round against the 70755 files, not carried from an earlier report.

- **Round 16's own fixes.**
  - BM:99 (drop Arcane Shot, then Multi-Shot; Serpent Sting last). Costs on the raid build with Efficiency 3: Serpent Sting rank 9 250 (25295), Multi-Shot 13.9% of base mana (2643 `cost_pct`), Aimed Shot rank 6 310 (20904), Arcane Shot rank 8 190 (14287). Serpent Sting 675 Nature over 15 s with no attack power coefficient (effect aura 3, amount 135 every 3000 ms, `ap_coefficient` 0); Multi-Shot has no flat amount (effect 121, amount 0). The damage-per-mana order Serpent Sting, Aimed Shot, Multi-Shot, Arcane Shot holds.
  - MM:105 (Multi-Shot goes before the sting, Aimed Shot and Sniper Shot kept). Same client values with Efficiency 5, Barrage 3 and Ranged Weapon Specialization 5; the sting stays ahead of Multi-Shot before armor by about 13%, further after it. The whole-rotation check is still the carried sim (Sims requested 1).
  - R16-3 card note (BM Serpent Sting, `addon-data.json` rotations row): rendered as written; the action condition (under 1.5 s left and Multi-Shot not about to return) matches the note.
  - R16-4 "mail from 40" lines: leather counts on the band 60 raid rows are BM 4 of 8 both factions, MM 3 Alliance and 5 Horde, SV 5 of 8 both (armor `subclass_id` 2 against 3 on head, shoulders, chest, wrists, hands, waist, legs, feet). "Mail at 40 is the biggest armor upgrade you will get" is an armor statement and stays true.
- **Sim round 18's marksmanship.md:163** against the new row (28-49 rank 1, 52-59 rank 1, bare 60 rank 1): "through 49", "At 50 and 51, rank 4 and rank 1 come out even", "When rank 5 arrives at 52" (20903 at 52), "keep it through 59", "At 60, use rank 1 out in the world, and cast the top rank in a raid" all match. Beast Mastery's row (28-59, bare 60) matches BM:157 "all the way to 60" and "In a raid at 60 ... cast the top rank".
- **Talent text** (`talents/hunter.json`, every rank read):
  - Deadly Aspects: "While Aspect of the Hawk is active, Auto Shot has a 10% chance ... While Aspect of the Beast is active, all melee auto attacks have a 10% chance of increasing melee attack speed by 30% for 12 sec" (BM:77, BM:147, MM:77, SV:77). Falcon's own text (469145) adds the melee proc under Falcon and nothing for Auto Shot, so BM:87, MM:92 and SV:86 hold (round 1 Q1, round 10 D1).
  - Focused Fire 2% "while your pet is active"; Unleashed Fury 15% "pets and hawks"; Ferocity 10%; Frenzy 100% chance at 5 points, 30% for 8 sec; Bestial Wrath 50% for 18 sec, "does not feel ... fear", 120 s (trainables); Endurance Training 3% a point; Bestial Swiftness 30% pet speed; Intimidation 3 s stun on the pet's next hit, high threat; Bestial Discipline 2 points "allows 50% of your Mana regeneration to continue while casting" (BM:81).
  - Careful Aim 100% of Intellect; Lethal Attacks 5%; Efficiency 15% on "Shots, Stings, and melee abilities" (MM:76, MM:153, and SV:82's "comes on top of Resourcefulness's 60%" is right because Efficiency does reach melee abilities); Improved Stings 2 points 13%; Rapid Killing 1 point 1 min and 10%, 2 points 2 min and 20%, on a kill or a death under your Serpent Sting; Rapid Recuperation 2 points "50% of your Mana Regeneration ... for 15 sec, increased to 100% when Rapid Killing is consumed" (MM:84); Mortal Shots "on all ranged abilities", 6% a rank (BM:79 212% and 224%, MM:71 230%, SV:82 "ranged only"); Barrage "Multi-Shot, Aimed Shot, and Volley" 5/10/15%; Ranged Weapon Specialization 5%; Lone Wolf 20% "while you do not have an active pet"; Trueshot Aura rank 1 20, rank 5 100 (20906); Sniper Shot rank 1 +320, rank 3 +590, 4 s cast, 15 s cooldown (trainables 1310786); Hawk Eye range; Improved Concussive Shot a stun chance; Deflection parry (BM:81).
  - Improved Tracking 1% a point on Beasts, Demons, Dragonkin, Elementals, Giants, Humanoids and Undead; Savage Strikes 2% a point "melee abilities"; Survivalist 2% health a point; Clever Traps 30% on Immolation and Explosive Trap; Surefooted 3% hit; Deterrence 25% dodge and parry 10 s, 5 min; Predator's Edge 30% crit damage and 50% off hand (230%, three quarters of a main-hand swing); Counterattack after a parry, cannot be blocked, dodged or parried; Resourcefulness 2 points "60% ... and grants your melee and ranged ability Critical Strikes a 100% chance to allow 50% of your Mana Regeneration to continue while casting for 30 sec" (SV:53, SV:78, SV:102 "half"); Expose Prey 10% on "melee and ranged attacks against enemies afflicted by Hunter's Mark", 10 sec; Survivalist's Discipline 40% off Trap and Deterrence cooldowns; Strider Kick 100% weapon damage; Lightning Reflexes 15%; Lacerating Strikes 40% over 21 sec. Rapid Killing (tier 2, column 0) and Lone Wolf (tier 2, column 3) share a row (SV:82).
- **Trainer levels and ranks** (`trainables/hunter.json`): Heart of the Lion, Raptor Strike, Auto Shot and Track Beasts 1; Raptor Strike ranks 8, 16, 24, 32, 40, 48, 56; Serpent Sting and Monkey 4; Arcane Shot and Hunter's Mark 6; Concussive Shot 8; Hawk, Tame Beast, Call Pet, Revive Pet and Track Humanoids 10; Mend Pet, Wing Clip and Distracting Shot 12; Immolation Trap and Mongoose Bite 16; Multi-Shot and Track Undead 18; Aimed Shot, Cheetah, Freezing Trap and Disengage 20; Aimed Shot rank 2 at 28; Rapid Fire 26; Feign Death, Aspect of the Beast, Lacerate and Intimidation 30; Explosive Trap 34; Sniper Shot ranks 40, 48, 58; Tranquilizing Shot, Falcon, Aspect of the Beast rank 4 and Hawk rank 6 at 58 and rank 7 at 60. Hawk rank 7 (25296) is still `"book": true` in `spellranks.json`, so "Hawk's Ahn'Qiraj rank" stands (round 11 D2).
- **Spell text** (`spelltips/hunter.json`, `spellconst/hunter.json`): Heart of the Lion 10% stats and 40 attack power with 4 a level to max level 60 (409583 auras 99 and 124), 280 at 60 by the client's BaseLevel 0 render and the engine (round 5 D2); Aimed Shot "shares its cooldown with Multi-Shot", 2000 ms cast, category 6000; Multi-Shot "hitting 3 targets"; Rapid Fire "ranged and melee attack speed by 40% for 15 sec", 300 s; Disengage 5 s, "Disables auto-attack"; Feign Death; Tranquilizing Shot "Attempts to remove 1 Frenzy effect"; Distracting Shot "causing threat"; Freezing Trap "Any damage caused will break the ice"; Explosive Trap rank 3 229 + 330 over 20 s within 10 yards for 520; Immolation Trap ranks 50/90/135/190/245; Volley rank 3 112 a second for 6 s, 490 mana; Mongoose Bite "Can only be performed after you dodge"; Lacerate rank 4 406 over 21 s; Wing Clip rank 3 50 damage for 80 (the SV card note); Mend Pet rank 7 480 mana against Revive Pet at 100% of base mana (982 `cost_pct` 100), so BM:101 "cheaper" holds; Aspect of the Beast rank 4 110 melee attack power and "untrackable"; Falcon "the same amount as their highest rank of Aspect of the Hawk" (Hawk rank 6 110 at launch, rank 7 120 from the book); Monkey 8% dodge; Cheetah 30% and the daze; Deterrence 5 min.
- **Rotation prose against the cards** (`data/curated/apl/hunter-*.json`, rendered in `addon-data.json`): the BM order (Hunter's Mark, Bestial Wrath with Rapid Fire, Aimed Shot when the shared cooldown comes up right after an auto shot, Multi-Shot otherwise, Serpent Sting under 1.5 s, Arcane Shot above 25% mana) and the MM order (plus Sniper Shot right after an auto shot) match BM:87-99 and MM:92-105, and the SV order (Raptor Strike, Counterattack, Mongoose Bite, Lacerate, Strider Kick, Immolation Trap, Wing Clip only under Windfury) matches SV:90. Every card note was read: none states a rank, a value or a condition the client contradicts.
- **Stat priority prose** explains and never ranks; each reason is true on 70755 (9% cap, Careful Aim, Mortal Shots and Predator's Edge crit, Lightning Reflexes 15%, Surefooted 3%, Dwarf 1% hit). Frontmatter statPriority still names the stats of the band 60 raid weights in `bis/hunter-*.json`.
- **Best in slot and professions** against the band 60 raid rows, both factions (`bis/hunter-*.json`, `items/hunter.json` after the data job): Frozen Heart 249469 and Molten Heart 249470 bind on pickup, source Enchanting; Legionite Glaive 250619 binds on pickup, Blacksmithing, polearm 3.2 (BM and MM main hand, both factions); Heartseeker 12783 binds when equipped, Blacksmithing, SV off hand both factions; Black Dragonscale Helm 252605 (mail), Dawn Armor 252483 (leather) and Mourning Handguards 279257 (mail) bind when equipped, Leatherworking; Hyper Deluxe Sniper Rifle Mk XVII 279273 binds when equipped, Engineering (BM and MM ranged, both factions). SV wears Teebu's Blazing Longsword 2.9 and Heartseeker 1.9, so "Survival dual-wields" and "pick it for damage and speed" hold; its ranged slot is a world-drop gun or Blackcrow, "a stat stick".
- **Enchants** (`enchants.json`): Sniper Scope and Biznicks 247x128 Accurascope exist on the ranged slot; Agility on gloves (up to Superior, 15), boots (Greater, 7) and cloak (Lesser, 3); Greater Stats on chest; the only Agility bracer enchant is Minor Agility (+1), so "no bracer enchant carries real Agility" holds; weapon Strength 15, Agility 15, Mighty Intellect 22 and 2H Agility 25 exist for SV:110's "all attack power for you".
- **Consumables** (`simconsumes.json`): Major Mana Potion, Major Frenzy Potion, Elixir of the Mongoose, Poached Sunscale Salmon, Ground Scorpok Assay all present.
- **Races** (`races.json` rows for the six races): Elune's Light 10% crit 15 s on 3 min; Ruthless 1% crit; Artisan's Precision 1% hit; Big Game Hunter 5% to Beasts; Blood Fury 10% attack power, ranged attack power and spell power for 15 s on 2 min; Command 5% pet damage; Berserking 10% casting and attack speed 15 s on 3 min; Human Hunter is one of six new pairs. Every race line on the three pages matches, and recommendedRaces matches the prose.
- **Pets:** engine `pet.go` still takes 10% of the hunter's higher attack power and the hunter's crit on every table (index.md:32, BM:52, MM:88, SV:71). Player buffs on the pet are settled and not re-filed.
- **Source links:** all 15 fetched again with a browser user agent. 200: the six foreversixty.gg URLs, the three Blizzard articles, the three Warcraft Tavern pages. 403 (unknown under the brief, labels confirmed in round 9): Wowhead Skyborne First Look, spell=1495, spell=19306.
- **Voice:** no em or en dashes and no site, sim, APL, band, preset or nightly words on any of the four pages.

## Decisions

- **D1. The main pages are clean this round.** Every number, level, rank, talent, item, racial and order was re-derived from the 70755 files and none is wrong. The only line that would have been a finding is survival.md:140, which sim round 19 rewrote on a measured result; its new wording is right.
- **D2. Round 19's survival.md text ships as written, with one SHOULD.** "your pet gets none" is loose against the settled Heart of the Lion rule; "none of theirs" is exact. Not a finding: the same bullet says "from other players" two lines up.
- **D3. Heart of the Lion stays 280.** The spellconst row reads 40 plus 4 a level with spell_level 1; the client renders from BaseLevel 0 (Wowhead Forever's 280 on both 409580 and 409583) and the engine agrees. Round 5 D2 stands; the sim lane's 280/276 owner item is noted, not re-filed.
- **D4. Hawk's Ahn'Qiraj rank stays.** `spellranks.json` still flags 25296 as a book; the trainables row at 60 is the same shape round 11 D2 weighed. The sim lane's "level-60 book ranks" owner item is noted, not re-filed.
- **D5.** All earlier decisions stand.

## Sims requested

1. **Carried from round 16 (HUNTER-R16-1, R16-2; not run in round 19):** BM and MM band 60 raid, both factions, 10000 iterations, seeds 1 and 50001, on fights long enough to run dry (300 s and 420 s). Compare three low-mana orders, each applied below 20% mana: (a) the shipped list; (b) drop Serpent Sting; (c) drop Multi-Shot and keep Serpent Sting. The pages now say (c); confirm it.
2. **Carried (round 16 D2):** the same runs with Aimed Shot rank 1 below 20% mana in place of rank 6.
3. **Carried, on the re-ranked gear (round 19 "Measure at refresh"):** MM race order on the raid build; Survival Crusader against Agility on both weapons; BM and SV race re-runs; BM and MM with the best non-BoP two-hander in Legionite Glaive's place; BM picks and weights with the pet on Heart of the Lion; the 46-47 Raptor Strike row on the re-ranked band 40 gear.

## Sim findings

- **New:** none. The engine's Serpent Sting, Multi-Shot, Aimed Shot, Raptor Strike (now every rank, r19 S2), Rapid Fire melee haste, Lone Wolf, Deadly Aspects under Hawk and Beast, and the pet's 10% attack power and crit copy all match the client rows the pages quote.
- **Carried, open with the sim lane (none claimed in prose):** Ranged Weapon Specialization on Serpent Sting (r19 H3); Lacerating Strikes on refresh and the queued Raptor Strike's off-hand miss (r19 H4); the hawk's level-60 swing (round 18 H4); Heart of the Lion on the pet's buff bar (round 18 H5); the stale top-level APL notes (not rendered: "Aimed Shot is only the pre-pull opener", Falcon "55 melee and 55 ranged").

## Sources

**Forever Sixty client data, 1.60.1.70755** (audit tree)
- `talents/hunter.json` (every rank text), `trainables/hunter.json`, `spellranks.json` (25296 book flag), `addon-builds.json`, `addon-data.json` (rendered hunter rotations).
- `spelltips/hunter.json`: 409580, 13165 to 25296, 13161 to 1299447, 469145, 13163, 5118, 19434, 20904, 2643, 25295, 14287, 1310687 to 1310786, 3045, 781 to 14273, 5384, 19801, 20736 to 15632, 1499 to 14311, 13813 to 14317, 13795 to 14305, 1510 to 14295, 2973 to 14266, 1495 to 14271, 24118 to 1299332, 1317257, 19306 to 20910, 2974 to 14268, 14325, 13544, 982, 19577, 19574, 19263, 20906, 19503, 5116, 1515.
- `spellconst/hunter.json`: 409580, 409583, 982, 13544, 2643, 1495, 19263, 1310496, 25295, 20904, 14287, 1310786, 1328676 (the data job's addition).
- `bis/hunter-{beast-mastery,marksmanship,survival}.json` (bands 40, 50 and 60, both factions), `items/hunter.json`, `enchants.json`, `simconsumes.json`, `races.json`, `hotfix-push.json`.

**Repository**
- `data/curated/apl/hunter-{beast-mastery,marksmanship,survival}.json`; `data/curated/rank-choices.json` (hunter rows).
- `web/src/lib/guides/leveling-builds.ts` (the Leveling tree's ladder port).
- Engine 6ea675594 (gr17/engine): `sim/hunter/pet.go:204-241`, `sim/hunter/aspect_of_the_falcon.go`, `sim/hunter/heart_of_the_lion.go`, `sim/core/buffs.go:1437-1460`; `git log b2c4a2d05..6ea675594 -- sim/hunter` (8c24d6202 only).
- design/reviews/2026-11-04-sim-audit-round18/hunter.md (H2 to H5, M2, S3); r19/int design/reviews/2026-11-04-sim-audit-round19/hunter.md (H1, H5, H6, S1, S2) and integration.md; `git -C r19/int diff 919f7e8a7 HEAD -- web/src/content/guides data/curated/apl`.
- design/reviews/guide-accuracy/round1, round5 (D2), round10 (D1), round11 (D2), round15 and round16 hunter reports; ledger.md.

**Web** (fetched 2026-11-04, browser user agent)
- The 15 frontmatter URLs, with the status codes above.

## Counts

- CONFIRMED: 0
- LIKELY: 0
- NEEDS-SIM: 0 (two carried sim requests, neither claimed as a finding)
- NEEDS-OWNER: 0
- Covered by sim round 19: 1 line that would have been CONFIRMED on main (survival.md:140), r19 wording right; 3 r19 rewordings of lines true both before and after.
- SHOULDs: 2 new, 9 carried.

## Fix log (round 17)

- No finding on main. survival.md:140 (Lone Wolf dismiss line) is covered by sim round 19's landing; its "your pet gets none" SHOULD is routed to that lane.
