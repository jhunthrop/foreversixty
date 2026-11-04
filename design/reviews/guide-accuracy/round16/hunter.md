# Guide accuracy round 16: Hunter

Auditor: hunter. Pages: class index, Beast Mastery, Marksmanship and Survival, plus the card notes in `data/curated/apl/hunter-*.json`.
- Tree: gr16/audit at 2c3135a2e (land/guides-r15 on launch-state main).
- Client: 1.60.1.70755. Engine pin: b2c4a2d05 (sim/enginever/version.go); engine source read from gr15/engine, which is checked out at b2c4a2d05.
- Read first: the round 1 to round 15 hunter reports with their fix logs, round4/hunter-levels.md, the race-sims files, ledger.md, and design/reviews/2026-11-04-sim-audit-round17/hunter.md.

## What moved since round 15

`git diff 4600b7d82..2c3135a2e` touches only the round 15 fixes:
- beast-mastery.md:60 and index.md:32 (your own Heart of the Lion still reaches the pet).
- beast-mastery.md:157 and marksmanship.md:163 (the Aimed Shot rank 1 reason).

No hunter APL, BiS file, talent file, trainables row, addon-builds row or rank-choices row changed. Every line was read again in full this round.

## Builds and leveling order

The decode is unchanged (tab order = tier, then column, against `talents/hunter.json`):
- **BM leveling** 31/20/0: Deadly Aspects 5, Endurance Training 4, Focused Fire 2, Bestial Swiftness 1, Unleashed Fury 5, Ferocity 5, Intimidation 1, Bestial Discipline 2, Frenzy 5, Bestial Wrath 1; Lethal Attacks 5, Improved Stings 2, Efficiency 4, Careful Aim 5, Rapid Killing 1, Trueshot Aura 1, Mortal Shots 2.
- **BM raid** 31/20/0: the same, with Efficiency 3, no Rapid Killing, Mortal Shots 4.
- **MM leveling** 15/31/5: Deadly Aspects 5, Endurance Training 3, Focused Fire 2, Unleashed Fury 5; Lethal Attacks 5, Efficiency 5, Careful Aim 5, Rapid Killing 2, Trueshot Aura 1, Mortal Shots 5, Rapid Recuperation 2, Barrage 1, Ranged Weapon Specialization 5; Improved Tracking 5.
- **MM raid** 5/33/13: Deadly Aspects 5; Lethal Attacks 5, Improved Stings 2, Efficiency 5, Careful Aim 5, Lone Wolf 1, Trueshot Aura 1, Mortal Shots 5, Barrage 3, Ranged Weapon Specialization 5, Sniper Shot 1; Improved Tracking 5, Survivalist 5, Surefooted 3.
- **SV** 5/11/35: Deadly Aspects 5; Lethal Attacks 5, Careful Aim 5, Lone Wolf 1; Improved Tracking 5, Savage Strikes 2, Survivalist 3, Clever Traps 2, Surefooted 3, Deterrence 1, Predator's Edge 5, Counterattack 1, Resourcefulness 2, Expose Prey 2, Survivalist's Discipline 2, Strider Kick 1, Lightning Reflexes 5, Lacerating Strikes 1.

All five are legal at 51 points (prerequisites: Intimidation needs Bestial Swiftness, Bestial Wrath needs Intimidation, Frenzy needs Ferocity, Mortal Shots needs Careful Aim, Rapid Recuperation needs Rapid Killing, Sniper Shot needs Trueshot Aura, Counterattack needs Deterrence, Lacerating Strikes needs Expose Prey).

The shipped orders (`addon-builds.json`, index 0 = level 10) still match the prose:
- BM: Deadly Aspects 10-14, Focused Fire 15-16, Endurance Training 17-20, Bestial Wrath 40 (BM:147, BM:159).
- MM: Lethal Attacks 10-14, Efficiency 15-19, Rapid Killing 25-26, Rapid Recuperation 33-34 (MM:153, MM:159).
- SV: Improved Tracking 10-14, Savage Strikes 15-16, Survivalist 17-19, Lone Wolf at 60 (SV:134, SV:140).

## Findings

### HUNTER-R16-1 (CONFIRMED, rotation): Beast Mastery's low-mana order drops Serpent Sting, its most mana-efficient shot, and keeps Multi-Shot, its least efficient one after Arcane Shot
- **File:** web/src/content/guides/hunter/beast-mastery.md:99
- **Quote:** "When mana gets low, drop Arcane Shot first, then Serpent Sting, and keep Aimed Shot and Multi-Shot going."
- **Evidence:**
  - **Costs** (`spellconst/hunter.json`, engine `multi_shot.go:33`, hunter base mana 1720 in `core/base_stats_auto_gen.go:42`):
    - Serpent Sting rank 9 (25295): 250 mana.
    - Multi-Shot (2643): 13.9% of base mana, 239.
    - Aimed Shot rank 6 (20904): 310.
    - Arcane Shot rank 8 (14287): 190.
    - The raid build's Efficiency 3 takes 9% off all four.
  - **Damage:**
    - Serpent Sting: 135 Nature every 3 s for 15 s = 675. It has no attack power coefficient (`serpent_sting.go` BonusCoefficient 0), and Improved Stings 2 makes it 763. It is Nature, so armor does not touch it.
    - Multi-Shot: normalized weapon damage + RAP x 2.8 / 14 + ammo, with no flat amount (spellconst effect amount 0; `multi_shot.go` ApplyEffects). That is about 536 at band 60 raid. It is derived from round 15's measured rank 1 to rank 6 ratio of 65%: (X+40)/(X+350) = 0.65. It is physical, so it loses some of that to armor.
    - Aimed Shot rank 6 is the same shot plus 350.
    - Arcane Shot is a flat 280.
  - **Damage per mana point at 60, raid build** (crit, hit and Focused Fire scale every shot alike):
    - Serpent Sting: 763 / 227.5 = 3.35.
    - Aimed Shot rank 6: 886 / 282 = 3.14, before armor.
    - Multi-Shot: 536 / 217.5 = 2.46, before armor.
    - Arcane Shot: 280 / 173 = 1.62.
  - **Serpent Sting's crits:** its ticks crit on the ranged table and take Mortal Shots (`serpent_sting.go` CritDamageBonus; `core/spell_outcome.go:38-57`, DefenseTypeRanged), the same as Multi-Shot's.
  - **How robust:** for Multi-Shot to pass Serpent Sting per mana, the shot before its flat amount would need to be above about 730. That takes roughly 3000 ranged attack power, far past launch gear.
  - **The page against itself:** the Marksmanship card note on the same shot says Serpent Sting "is the best damage for its mana you have between the big shots, so keep it rolling for the whole fight".
- **Fix** (replace that paragraph):
  > When mana gets low, drop Arcane Shot first, then Multi-Shot, and let Aimed Shot take the shared cooldown only when its cast fits after an auto shot. Serpent Sting is the last shot you give up: it does more damage for its mana than anything else you fire. Drink a Major Mana Potion before you need it, not at 5%.

### HUNTER-R16-2 (LIKELY, rotation): Marksmanship's low-mana order has the same inversion between Serpent Sting and Multi-Shot
- **File:** web/src/content/guides/hunter/marksmanship.md:105
- **Quote:** "When mana gets low, stop Arcane Shot first, then Serpent Sting, and keep Aimed Shot, Multi-Shot and Sniper Shot going."
- **Evidence:**
  - The same client values as R16-1, on the raid build: Efficiency 5 (15% off), Improved Stings 2, Barrage 3 (15% on Multi-Shot and Aimed Shot), and Ranged Weapon Specialization (5% on weapon shots). Lone Wolf and Improved Tracking scale every shot alike.
  - **Damage per mana point at 60:**
    - Aimed Shot rank 6: 886 x 1.21 / 263.5 = 4.06.
    - Sniper Shot rank 3: 1126 x 1.05 / 310 = 3.81.
    - Serpent Sting: 763 / 212.5 = 3.59.
    - Multi-Shot: 536 x 1.21 / 203 = 3.18, before armor.
    - Arcane Shot: 1.73.
  - **Margin:** Serpent Sting is about 13% ahead of Multi-Shot before armor, and further ahead after it. Multi-Shot would pass it only with the shot above about 605 before its flat amount, which is roughly 2400 ranged attack power on the target.
  - Marked LIKELY rather than CONFIRMED because this margin is narrower than Beast Mastery's. See Sims requested 1.
  - **The page against itself:** the page's own Serpent Sting card says it "is the best damage for its mana you have between the big shots, so keep it rolling for the whole fight". The prose stops it second.
- **Fix:**
  > When mana gets low, stop Arcane Shot first, then Multi-Shot, and keep Aimed Shot, Sniper Shot and Serpent Sting going. Drink a Major Mana Potion before you need it.

### HUNTER-R16-3 (LIKELY, card note): the Beast Mastery Serpent Sting card says Multi-Shot is "always the better global", but on one target the sting outdamages it
- **File:** data/curated/apl/hunter-beast-mastery.json, rotation.priorityList[6].notes
- **Quote:** "Refreshing any earlier throws away a tick, and Multi-Shot is always the better global."
- **Evidence:**
  - **One global each on a single boss:** Serpent Sting rank 9 with Improved Stings 2 is 763 Nature with no armor. Multi-Shot is about 536 physical before armor.
  - **Why Multi-Shot still goes first:** the order on the card is right, for a different reason.
    - A late Multi-Shot pushes back the shared six second cooldown, which loses about a quarter of a shot for each 1.5 s.
    - A late sting loses only 1.5 s of a 15 s DoT, about a tenth of it.
  - "Always" fails on one target. Multi-Shot is the better global only on three targets.
- **Fix** (replace that sentence):
  > Refreshing any earlier throws away a tick. When both are ready, Multi-Shot goes first: a late Multi-Shot pushes back the shared cooldown, while a late sting costs only a moment of its damage.

### HUNTER-R16-4 (LIKELY, gear): "mail from 40" reads as a switch, but every best in slot list from 40 to 60 keeps leather pieces
- **Files:**
  - beast-mastery.md:117, marksmanship.md:123 and survival.md:106: "Hunters wear leather until 40 and mail from 40."
  - Also index.md:37, and the "Wear leather until mail at 40" lines at BM:149, MM:155 and SV:136.
- **Evidence:** `bis/hunter-*.json` against `items.json` (armor subclass 2 = leather, 3 = mail), counting head, shoulders, chest, wrists, hands, waist, legs and feet:
  - BM band 60 raid, both factions: 4 leather of 8, among them Dawn Armor (chest), Forest Stalker's Bracers and Shadowcraft Boots.
  - MM band 60 raid: 3 leather on Alliance, 5 on Horde.
  - SV band 60 raid: 5 leather of 8 on both factions.
  - The bands at 40 and 50 are 2 to 5 leather of 8.
  - The owner rule (mail at 40) is a proficiency rule. As a statement of what you wear, "mail from 40" contradicts the lists the page links to and renders.
  - Round 14 passed this line as "no row breaks the owner rule". That is true of the rule, but not of the sentence as a reader takes it.
- **Fix** (the Best in slot sentence on all three spec pages):
  > Hunters wear leather until 40 and can wear mail from 40, but mail is not a must: Agility decides each slot, and a strong leather piece often beats the mail you can get, so your list mixes both at 60.
- index.md:37 can stay as a proficiency line ("Mail from 40") if its second sentence reads: "Hunters wear leather until 40, and good leather stays in the mix after it."

### Re-checked this round and still true

- **Talent text** (70755 `talents/hunter.json`, every rank):
  - Mortal Shots 2 and 4 points give 212% and 224%; 5 points give 230%. Predator's Edge uses the same aura 108 / misc 15 bonus modifier (spellconst 19485 and 1310627).
  - Rapid Killing 1 point: 1 min and 10%. 2 points: 2 min and 20%.
  - Barrage 5/10/15%; Improved Stings 2 points 13%; Efficiency 3% a point; Survivalist 2% a point; Savage Strikes 2% a point.
  - Expose Prey 10%; Lacerating Strikes 40% over 21 s; Resourcefulness 60%.
  - Lone Wolf 20%; Trueshot Aura rank 5 = 100 at 60; Focused Fire 2%.
- **Trainer levels** (`trainables/hunter.json`):
  - Serpent Sting and Monkey at 4; Arcane Shot and Hunter's Mark at 6; Concussive Shot at 8; Hawk at 10; Wing Clip and Mend Pet at 12.
  - Raptor Strike ranks at 1, 8 and 16; Mongoose Bite at 16; Multi-Shot at 18.
  - Aimed Shot and Cheetah at 20; Rapid Fire at 26; Feign Death, Beast and Lacerate at 30.
  - Sniper Shot ranks at 40, 48 and 58; Falcon and Tranquilizing Shot at 60.
  - Track Humanoids at 10 and Undead at 18.
- **Spell text and constants** (`spelltips/hunter.json`, `spellconst/hunter.json`):
  - Sniper Shot rank 3: +590, 4 s cast, 15 s cooldown.
  - Aimed Shot: 2 s cast, 6 s shared cooldown.
  - Disengage 5 s and "Disables auto-attack".
  - Heart of the Lion 40 + 4 a level = 280 at 60.
  - Rapid Fire hastes melee and ranged.
  - Beast rank 4 = 110 = Falcon at launch.
- **Races** (`races.json`, `combos.json`):
  - Racial values: Elune's Light, Ruthless, Artisan's Precision, Big Game Hunter, Command, Blood Fury (2 min) and Berserking (3 min).
  - Human Hunter is one of six new pairs; Skyborne can be Hunters on both factions.
  - recommendedRaces matches the prose on all three pages.
- **Professions:**
  - Frozen and Molten Heart of the Mountain are BoP Enchanting; Legionite Glaive is a BoP Blacksmithing polearm.
  - Hyper Deluxe Sniper Rifle, Black Dragonscale Helm, Dawn Armor, Mourning Handguards, Scalegut Treaders and Heartseeker bind on equip (`items/hunter.json`).
- **Ammo, quivers, enchants and raid totems:**
  - Ammo and quivers (`items.json`): Jagged Arrow 40, Ice Threaded 51, Thorium Headed Arrow and Thorium Shells 52, Quickdraw Quiver 40, Harpy Hide Quiver and Gnoll Skin Bandolier 55.
  - Scopes are in `enchants.json`. The only Agility bracer enchant is Minor Agility (+1).
  - Strength of Earth Totem's top launch rank is 53 Strength and Curse of Recklessness's is 505 armor, both below Era's.
- **Source links:** all 15 were fetched again with a browser user agent.
  - 200: the six foreversixty.gg URLs, the three Blizzard articles and the three Warcraft Tavern pages.
  - 403 (unknown under the brief): Wowhead Skyborne First Look, spell=1495 and spell=19306.
- **Round 15's fixes:**
  - BM "well over half": rank 1 is 62 to 65% of the newest rank at 52 to 60, and about 80% at 28 to 35.
  - MM "most of what the newer ranks do": about 72 to 80% through 49.
  - The Heart of the Lion clause on index.md:32 and BM:60 matches engine H2.
- **Voice:** no em or en dashes, and no site, sim, APL, band, preset or nightly words in the prose.

### SHOULDs (wording, not findings)

- **Carried:**
  - MM:159 "so you have it for most hard pulls".
  - MM:94 five minutes on the leveling build.
  - SV:65 "smaller than Era's".
  - The Serpent Sting cards' "throws away a tick".
  - "The Molten Core tome".
  - BM:81 naming Deflection.
  - SV:118 Orc and Troll "level" with only Orc tagged.
  - BM:157 "skip ranks 2 to 5 at the trainer".
- **New, SV:80:** "Strider Kick is a third melee special". Raptor Strike, Mongoose Bite, Counterattack and Lacerate are already on the page, so "another melee special" reads truer.

## Decisions

- **D1. Under a mana squeeze, the order you drop shots follows damage per mana, and on 70755 Serpent Sting leads it.** Forever's 70753 push raised Serpent Sting to 675, and Multi-Shot lost its flat damage. So the Era habit of dropping the sting first is now backwards. BM is CONFIRMED because no reachable attack power flips it. MM is LIKELY because its Multi-Shot carries Barrage and Ranged Weapon Specialization, which narrows the margin to about 13%.
- **D2. The rank 1 Aimed Shot advice for the raid stays as written.** The page casts the top rank in a raid with Judgement of Wisdom and a potion. Rank 1's damage per mana (about 6.7) would also beat Serpent Sting under a squeeze. That is a new raid tactic, so it goes to Sims requested 2, not into this round's fixes.
- **D3. "Mail from 40" is a proficiency rule, and the prose should say so.** Round 14's pass held for the rule, not for the sentence.
- **D4.** All earlier decisions stand.

## Sims requested

1. **New (R16-1, R16-2):** BM and MM band 60 raid, both factions, 10000 iterations, seeds 1|50001, on a fight long enough to run dry (300 s and 420 s). Compare three low-mana orders, each applied below 20% mana:
   - (a) the shipped list;
   - (b) drop Serpent Sting below 20% mana;
   - (c) drop Multi-Shot below 20% mana, with Serpent Sting kept.

   The page should give whichever of (b) and (c) wins.
2. **New (D2):** the same runs with Aimed Shot rank 1 below 20% mana in place of rank 6.
3. **Carried:**
   - MM bare 53, 56 and 59, Aimed Shot rank 1 against the newest rank (HUNTER-R15-3).
   - SV bare 60, Lone Wolf against the pet, 10000 iterations.
   - MM race order on the raid build.
   - Survival Crusader against Agility on both weapons.
   - BM and SV race re-runs on the re-ranked gear.
   - BM and MM with the best non-BoP two-hander in Legionite Glaive's place.
   - BM picks and weights with the pet on Heart of the Lion.

## Sim findings

- **Carried, LIKELY:**
  - Falcon's melee proc does not need the talent in the engine.
  - Summon Hawk is unmodeled (owner check).
  - The stale top-level APL notes (not rendered): "Aimed Shot is only the pre-pull opener", and Falcon "55 melee and 55 ranged".
- **New:** none. The engine's Serpent Sting (no coefficient, ranged-table tick crits with Mortal Shots) and Multi-Shot (no flat amount, 13.9% of base mana) both match the client.

## Sources

**Forever Sixty client data, 1.60.1.70755**
- `spellconst/hunter.json`: 25295 Serpent Sting, 2643 Multi-Shot, 20904 Aimed Shot, 14287 Arcane Shot, 1310786 Sniper Shot, 19485 Mortal Shots, 1310627 Predator's Edge, 409580 and 409583 Heart of the Lion.
- `spelltips/hunter.json`: the hunter kit, including 469145 Falcon, 3045 Rapid Fire and 781 Disengage. `spelltips/shaman.json` and `spelltips/warlock.json`: Strength of Earth Totem, Windfury Totem and Curse of Recklessness.
- `talents/hunter.json`, `trainables/hunter.json`, `addon-builds.json`.
- `bis/hunter-{beast-mastery,marksmanship,survival}.json` (bands 40 to 60, both factions), `items.json`, `items/hunter.json`, `enchants.json`, `races.json`, `combos.json`, `simconsumes.json`.

**Repository**
- Engine b2c4a2d05 (gr15/engine):
  - `sim/hunter/serpent_sting.go`, `sim/hunter/multi_shot.go:33-100`, `sim/hunter/pet.go:205-230`.
  - `sim/core/spell_outcome.go:38-102`, `sim/core/dot.go:396-440`, `sim/core/base_stats_auto_gen.go:42`.
- `data/curated/apl/hunter-{beast-mastery,marksmanship,survival}.json` (card notes); `sim/request/raid-consumes.json`.
- design/reviews/guide-accuracy/round15/hunter.md (the band 60 raid rank 1 to rank 6 ratio).

**Web** (fetched 2026-11-04, browser user agent)
- The 15 frontmatter URLs, with the status codes above.

## Fix log (round 16)

- HUNTER-R16-1: fixed, beast-mastery.md:99, report text applied (Arcane Shot, then Multi-Shot; Serpent Sting last). Costs checked in 70755 spelltips (Serpent Sting r9 250, Multi-Shot 13.9% of base).
- HUNTER-R16-2: fixed, marksmanship.md:105, worded to hold whatever the requested sim shows: "On one target Serpent Sting does more damage for its mana than Multi-Shot, so let Multi-Shot go before the sting, and keep Aimed Shot and Sniper Shot going." The sim stays requested.
- HUNTER-R16-3: fixed, `data/curated/apl/hunter-beast-mastery.json` Serpent Sting notes: reason only, report text applied. Action untouched.
- HUNTER-R16-4: fixed. Best in slot sentence on beast-mastery.md:117, marksmanship.md:123, survival.md:106: mail from 40 kept, leather pieces stay in the best gear. Leveling lines BM:149, MM:155, SV:136: "From 40 you can wear mail, but keep a leather piece whenever it is the better one." index.md:37: "and good leather stays in the mix after it." Leather counts re-checked on the band 60 raid rows (BM 4/8, MM 3 and 5, SV 5).
