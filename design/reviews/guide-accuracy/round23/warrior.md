# Guide accuracy round 23: Warrior

Audited on the round 23 audit tree (24ee457d0: main after guide round 22; no sim round 21 commit in the log, so there is no sim round 21 class report; engine pin d5c708834 in `sim/enginever/version.go`, engine source at `gr23/engine`) against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/warrior/{index,arms,fury,protection}.md`, every note in `data/curated/apl/warrior-{arms,fury,protection}.json`, every rendered row of `addon-data.json` (`rotations/warrior-arms` 20 bands, `warrior-fury` 19, `warrior-protection` 24), both shipped orders per spec in `addon-builds.json` (re-decoded against tier and column), `talents/warrior.json`, `spelltips/warrior.json`, `spellconst/warrior.json`, `trainables/warrior.json`, `races.json`, `items/warrior.json`, `loot.json`, `bis/warrior-*.json`, `data/curated/content-phases.json` and the engine's warrior code. I read the round 22 warrior report, the ledger's Round 19 to 22 sections and the briefs.

Rendering as rounds 20 to 22 documented: the addon row carries the note up to the first ". " at every band; the guide card prints up to the first ": ", ". " or "; "; the sim drawer shows the whole note with no level, so every sentence counts at every band.

## Changes since round 22's tree (987d08a7c)

`git diff 987d08a7c HEAD` for warrior: the Protection Shield Slam note (`warrior-protection.json`, the `sim/request` copy) and the regenerated `addon-data.json` / `Data.lua` text. No warrior guide page, talent, trainable, race or order file moved.

**Round 22's fix, re-checked as new work:**

- R22-1 Shield Slam note "On cooldown: a hard hit with a very high amount of threat, and it grows with your block value.": right at every band 40 to 60. "very high amount of threat" is the client's own wording on all four ranks (23922 to 23925); block value is added on all four; "a hard hit" compares with nothing, so Revenge's attack power share cannot break it on either side. "On cooldown" matches the card (no condition; 6 s, `category_cooldown_ms` 6000). The text is in `addon-data.json` (11 bands), `Data.lua` (11) and the `sim/request` copy; no row still says "your biggest hit". Side effects: the Revenge note's "most threat per rage" is unaffected; no page line depended on the old superlative (protection.md:69 "Your biggest button" stays a Should).

## Orders (index 0 = level 10)

Unchanged from round 22, re-decoded: Arms leveling Improved Heroic Strike 10, Deflection 11-13, Improved Rend 14-16, Improved Charge 17-18, Improved Tactical Mastery 19-21, Improved Overpower 22-23, Anger Management 24, Deep Wounds 25-27, Two-Handed Weapon Specialization 28-30, Impale 31-32, Bloodthrill 33-34, Sweeping Strikes 35, Weaponmaster 36-39 and 41, Mortal Strike 40, Improved Slam 42-43, Cruelty 44-48, Unbridled Wrath 49-53, Improved Bloodrage 54-55, Iron Will 56-60. Arms raid Improved Heroic Strike 10-12, Improved Rend 13-15, Improved Tactical Mastery 16-18, Improved Overpower 19-20, Anger Management 21, Deep Wounds 22-24, Spearing Strike 25, Two-Handed 26-28, Impale 29-30, Bloodthrill 31-32, Sweeping Strikes 33, Weaponmaster 34-38, Improved Slam 39 and 41, Mortal Strike 40, Cruelty 42-46, Unbridled Wrath 47-51, Enrage 52-56, Reckless Swings 57, Improved Execute 58-59, Improved Bloodrage 60. Fury (both identical) Cruelty 10-14, Unbridled Wrath 15-19, Furious Precision 20-22, Enrage 23-24, Dual Wield Specialization 25-29, Death Wish 30, Reckless Swings 31, Improved Execute 32-33, Improved Berserker Rage 34-35, Improved Whirlwind 36-37, Flurry 38-39 and 41, Bloodthirst 40, Improved Heroic Strike 42-44, Improved Rend 45-47, Improved Charge 48-49, Improved Tactical Mastery 50-52, Anger Management 53, Deep Wounds 54-56, Impale 57-58, Improved Bloodrage 59-60. Protection (both identical) Improved Bloodrage 10-11, Shield Specialization 12-16, Anticipation 17-21, Improved Revenge 22-24, Improved Thunder Clap 25-27, Last Stand 28, Master of Defense 29-30, Defiance 31-33, Improved Sunder Armor 34-36, Concussion Blow 37, Focused Rage 38-39 and 41, Shield Slam 40, Bastion 42-46, Improved Heroic Strike 47-49, Deflection 50-54, Improved Rend 55-57, Deep Wounds 58-60.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| WARRIOR-R23-1 | CONFIRMED | page prose ("every" claim that holds for only some races) | fury.md:114 | "Whatever you pick, press its active racial on cooldown and line it up with Death Wish when the timers allow." | Every race has two actives (`races.json`) and on five of the races this paragraph and the one above it discuss, neither adds damage. The Undead the page recommends: Will of the Forsaken (removes Charm, Fear and Sleep, 2 min, no immunity) and Cannibalize, and its damage racial Touch of the Grave is passive. Pressing Will of the Forsaken on cooldown throws away the fear break, the very thing fury.md:92 tells the player to plan around on Magmadar, Onyxia and Nefarian. Human: Will to Survive (stun break, 3 min), Perception. Dwarf: Stoneform, which protection.md:115 says "only helps when you press it, so save it", Find Treasure. Tauren: War Stomp (stun), Cultivation. Skyborne: Walk on Air, Read Leyline or Skysight. Only Elune's Light, Blood Fury, Berserking and Eureka! are damage actives. Every earlier round read this sentence as the damage races it was written for; "whatever you pick" says otherwise. | "If your race has a damage racial, Elune's Light, Blood Fury, Berserking or Eureka!, press it on cooldown and line it up with Death Wish when the timers allow." |
| WARRIOR-R23-2 | LIKELY | page prose (stated reason gives the cooldown where the rule is the kill window; the condition points the wrong way) | arms.md:131; fury.md:134 | "Victory Rush whenever it is up after a kill: it heals 10% and comes back every 30 seconds, so open the next pull with it if the last one died fast." | Client 402927: "Only useable within 20 sec after you kill a non-trivial enemy", `cooldown_ms` 30000. The reason to press it on the next pull is that a dead mob leaves no target and the window lasts 20 seconds; the 30 second cooldown does not lead to that, and the window is never mentioned. "If the last one died fast" is often the opposite of when it is available: press it at the start of a pull, kill that mob at 12 seconds, and it is back at 30 with the window shutting at 32, so 2 seconds to use it, while a mob that took 30 seconds leaves the whole 20. A fast kill only helps when the last pull did not spend it. The same sentence's "open the next pull with it" also sits beside "Charge every pull": the Charge comes first, since Victory Rush needs melee range. Not CONFIRMED because the sentence is gated by "whenever it is up" and following it costs nothing. | Both pages: "Charge every pull, and Victory Rush whenever it is up: it heals 10%, works for 20 seconds after a kill and comes back every 30 seconds, so when the next mob is close, press it right after the Charge." |

No card note needs a change, so no `addon-data.json` or `Data.lua` regeneration.

Counts: CONFIRMED 1, LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0.

## Shoulds (wording, not findings)

- **New, arms.md:53** "Rage starved without a raid around you": on Forever's normalized rage (`sim/core/rage.go`; index.md:28) attack power and raid damage buffs add no swing rage, so a raid helps only through crit buffs (a crit swing pays double), a Windfury Totem's extra swings and raid damage taken, and a solo Arms warrior earns rage from being hit that a raid DPS does not. Arms is short of rage in a raid too (arms.md:71). The line is not false, since those raid sources are real. Suggested: "Short of rage everywhere, so every point has a job". See Sims requested.
- **New, protection.md:89** "Bloodrage before you pull so your first Sunder Armor or Shield Slam is paid for, and Battle Shout with it": at 60 with a self-shout the bar at the pull is about 8.75 (15 now, 1.5 a second from 2.5 s before, less the 10 shout), a hair under Sunder's 9 and far under Shield Slam's 17. Without the shout it is 18.75 and both are paid. True in the case the sentence leads with; only the shout branch misses.
- **New, arms.md:85** "a shout, a Rend and a Mortal Strike in your first three globals": on the raid build without Improved Charge, with no crit and no potion, the bar reaches about 29 at the third global against Mortal Strike's 30, so it can slip about half a second. The page's own Mighty Rage Potion on the pull covers it, and the level 60 build's 21 rage Charge does too.
- **New, arms.md:91** "that first Execute is your biggest hit of the fight": holds with the bar arms.md:48 and the note ask for, comparing like with like. A non-crit Execute at about 60 rage is below a Mortal Strike crit at 220%; an Execute crit at that rage is above it. Round 22 read this as right against a full bar; left unfiled for the same reason.
- **New, fury.md:86 and the Fury Reckless Swings note** "the healers' to cover", "which the healers cover": solo from 31 there are no healers. Not a false statement about the group play the card describes.
- **New, arms.md:109 and fury.md:112** "Recklessness, which already makes every hit crit": glancing white swings do not crit. The reason (a crit racial adds nothing inside it) still holds.
- **New, protection.md:103** "only a few late sets carry" Expertise: also true of rings and a trinket that are not sets (Band of the Better Half 284716 is a set piece, but the sixteen Signets of Vigilance and Adaptive Combat Assistant 272437 are not). All are later content (`rep:the-watchers:*` and vendor 227853 `opens: later`, content-phases.json), so "late" stands. Dwarven Tree Chopper and Servomechanic Sledgehammer carry it early as two-handers a tank does not use. "a few late pieces" is exact.
- **Carried from round 22:** arms.md:131 Sweeping Strikes and Whirlwind undated ("once you have both"); fury.md:134 Bloodthirst undated ("once ... are in"); protection.md:69 "Your biggest button"; protection.md:75 "the three buttons you press most after Shield Slam"; Arms Reckless Swings note "white swings for the next 8 seconds" (next 4 within 8 s); Protection Shield Block note's two-attack cap; Fury Execute "First"; Protection Defensive Stance note at 10 to 27 naming Revenge, Shield Block and Shield Wall; Arms Hamstring "never from 25% down" before Execute; fury.md:112 "edges Night Elf by a hair"; protection.md:113 "no race's health swings less"; arms.md:125 skipping the level 19 point; fury.md:74 "Fury never casts Rend"; arms.md:115 "other pair"; fury.md:54 "debuff"; fury.md:100 Edgemaster wording; "From 36" lead on the level 60 card; Arms Heroic Strike note "above 90" against the card's 90 or more.

## Checked and right this round

- **Reason each line gives (round 23's lens), all four pages and every note:**
  - Rage: rage per swing set by speed with crit double (`rage.go` 3.46 / 4.50 / 0.5 / 2.0) supports index.md:28, fury.md:100 and :104 ("how often you land"), arms.md:99 ("swings are rage"). Off-hand hit is rage holds because a landed swing is what pays.
  - Improved Execute "so the rage you hold goes into damage" (5 more rage converts); Improved Bloodrage over Charge (30 against 21; 25 at 59 against 21); Improved Tactical Mastery "lets a Charge's rage come with you" (25 against 15 / 21; Fury holds it from 52, the line is a talent note on the finished build); Deep Wounds "needs the Improved Rend points" (prerequisite Improved Rend 3); Anger Management behind Improved Tactical Mastery 3; Mortal Strike behind Sweeping Strikes; Shield Slam behind Concussion Blow; Improved Charge "part of the points the rows above Deep Wounds ask for" (tier 2 needs 10 points; without those two the rows hold 9).
  - Overpower ahead of Mortal Strike: its window closes, the cooldown does not. Whirlwind and Sweeping Strikes rage: 25 kept pays Whirlwind's 25 exactly. Spearing Strike stops at the swap (Battle Stance only, engine mask). Overpower and Rend stop in Berserker Stance (engine masks), Mortal Strike and Execute do not.
  - Thunder Clap "a straight cut to the damage you take" (measured 16%); Demoralizing Shout "a real slice" (204 at 60: 196 + 6 x 1.4 per level, `spellconst` 11556; 268 attack power boss, about 23% of a swing, owner ruling in the card's notes); Night Elf, Tauren and Dwarf reasons against `races.json`; Blacksmithing "plate and shields" (Evergreen Shield 279262, Jagged Obsidian Shield 22198 in `crafted:blacksmithing*`); two-hander skill only on the weapon (of every warrior item, only Obsidian Edged Blade carries Two-handed Swords; Edgemaster's is Axes, Daggers, Swords).
- **"Whenever it is ready" against the card's own condition, every row:** Shield Block "whenever it is down" = `not auraIsActive`; Revenge and Shield Slam unconditional; Thunder Clap and Demoralizing Shout "last 3 seconds" = under 3 s remaining, durations at every rank well above it (Thunder Clap 10 to 30 s, 6 s cooldown; shout 45 s); Rend "keep it ticking" = `not dotIsActive`; Battle Shout "whenever it drops" = `not auraIsActive`; Bloodrage "under 80" / "under 60" match; Fury Bloodthirst, Whirlwind, Reckless Swings "on cooldown" unconditional; Berserker Rage level 36 and rage under 80; Heroic Strike 90, 60, 40 (Should on "above"); Death Wish remaining 210 s = a second one fits; Recklessness execute or last 15 s; Arms swap rage 25 or less after the first Execute, with Recklessness known and ready; Slam cast under 1.5 s = from the first Improved Slam point (L 42, R 39), and rank 1's text already removes the swing delay.
- **Both sides of every comparison across all bands:** Revenge per rage (5, 2 at 41) against Sunder 15 / 12 / 9, Shield Slam 20 / 17, Thunder Clap 20 / 14 / 11, Heroic Strike 15 / 12 / 9; Bloodrage with Improved Bloodrage against Charge at 59 and 60; Bloodthirst against Whirlwind per cooldown second at low and high attack power (with Improved Whirlwind's 8 s and 30%); Execute cost 15 / 10 (Fury 32, Arms raid 58); Thunder Clap "4 more" with Improved Shield Wall (one point left of three: 2 against 6); Shield Wall 15 to 4 minutes. Round 22's Revenge against Shield Slam no longer appears as a comparison.
- **Every superlative and quoted cost, duration and cooldown:** as round 22's list, re-read against `spelltips` and `spellconst`; nothing moved. Reckless Swings 1 s global (`gcd_ms` 1000), 25% base health (`cost_pct`); Shield Block `gcd_ms` 0; Bloodrage `gcd_ms` 0; Enrage 4% at 2 points, 30% at every rank; Furious Precision 10% at 3; Improved Slam 1 s and 10 s at 2; Sweeping Strikes five charges.
- **Spells and talents named in leveling paragraphs against their levels:** arms.md:121 to :131, fury.md:124 to :134, protection.md:125 to :135 against `trainables` and the orders. Clean apart from the Shoulds carried.
- **index.md** every line. **Gear lines** against BiS and bonding: Frozen Heart 249469 and Molten Heart 249470 (bind on pickup, Enchanting; Fury Alliance raid row carries both); Arcanite Champion 12790 and Fury's Blacksmithing weapons bind on equip. **Race orders** against round 19's measured runs, unchanged. No page names 70755; no em or en dashes. Frontmatter sources unchanged since they last passed.

## Decisions

- **D1. R23-1 is CONFIRMED on `races.json` and the page's own line 92.** The sentence's subject is explicit ("Whatever you pick") and following it on the recommended Horde race loses a fear break.
- **D2. R23-2 is LIKELY, not CONFIRMED.** The numbers are hard; what is judgment is how far "whenever it is up" and "if the last one died fast" mislead. Take it: the fix also gives the 20 second window neither page states.
- **D3. "Rage starved without a raid" stays a Should.** Raid crit buffs and Windfury do add rage, so without a measurement the sentence is not false.
- **D4. Expertise "only a few late sets" stays a Should.** Round 7's check missed the non-set pieces, but every non-set piece a tank would wear is later content, so the line misleads no one at launch.
- **D5. No web research was needed.** Every call this round is decided by client text or the page itself.
- **D6. Rows that render before their talent or spell exists stay with the component lane,** as in rounds 15 to 22.

## Sims requested

- Only if the arms.md:53 wording is kept: Arms level 60 build, rage gained per second by source, raid preset against bare, 10,000 iterations, seeds 1 and 50001.
- Carried, only if the wording is kept: Protection raid preset at 60, average landed Revenge (11601) against Shield Slam (23925); Fury Alliance Skyborne against Night Elf at 30,000 iterations (fury.md:112); Protection race TMI on one shared raid preset (protection.md:113).

## Engine questions

- **New, W23-1 (extends W22-1):** the trainer rows at 60 list more than Revenge: Battle Shout rank 7 (25289) and Heroic Strike rank 9 (25286) sit in `trainables/warrior.json` at 60 with `source: skill_line_ability`, beside Revenge rank 6 (25288), while the cards cast 11551, 11567 and 11601. Settle all three together with the book-rank question. Not a guide finding.
- Carried: W22-1, W21-1 Classic weapon speeds, X1, P5 threat constants, AC Arcanite Champion PPM, C1 Eureka! with Execute (#439), C2 Deep Wounds' double attacker multiplier, whether an outright-missed Heroic Strike spends a Reckless Swings charge.

## Sim findings

- None new.

## Sources

- Client 1.60.1.70790: `spelltips/warrior.json` (402927 Victory Rush description and every spell quoted above), `spellconst/warrior.json` (gcd, duration, `points_per_level`), `talents/warrior.json`, `trainables/warrior.json`, `races.json`, `items/warrior.json`, `loot.json`, `addon-builds.json`, `addon-data.json`, `bis/warrior-*.json`.
- Curated: `data/curated/apl/warrior-{arms,fury,protection}.json` (rows, conditions, notes), `data/curated/content-phases.json` (The Watchers).
- Engine d5c708834: `sim/core/rage.go`, `sim/warrior/{reckless_swings,execute,sweeping_strikes,stances,talents}.go`, the `RegisterSpell` stance masks across `sim/warrior`.
- Reviews: `design/reviews/guide-accuracy/round22/warrior.md`, `ledger.md` (Rounds 19 to 22).

## Note sweep

One line per row note: spec, row, bands it renders, each named talent / spell / rank / value with the level the player has it (L = leveling order, R = raid order), verdict.

Arms (20 bands, 10 to 60)

- Berserker Stance (30 to 60): "From 50" (Recklessness 50; card asks `spellIsKnown` 1719); Improved Tactical Mastery 3 (L 21, R 18) 25 rage; Mortal Strike (40) and Execute (24) work there, Overpower and Rend do not (masks). OK.
- Bloodrage (10 to 60): 10 + 10 over 10 s, 60 s (2687); a quarter per point (12301); L 54 / 55, R 60. OK.
- Battle Shout (10 to 60, ranks 6673 to 11551): 10 rage every rank, 20 yards, raid. OK.
- Rend (10 to 60): bleed ignores armor; Overpower and Bloodthrill named only "once Bloodthrill is in" (L 33, R 31); main-hand hits on your Rend (1289682). OK.
- Overpower (12 to 60): Mortal Strike "from 40"; dodge opens it; Bloodthrill "once the talent is in"; Improved Overpower 2 (L 23, R 20) "most of the time" (50% plus base crit). OK.
- Mortal Strike (40 to 60, ranks 12294 to 21553): 6 s every rank; healing cut 50%, 10 s. OK.
- Reckless Swings (10 to 60, component row; talent R 57): raid build only, after Mortal Strike (40); 8 s; 5 rage; quarter of base health. OK (Should on the four-attack cap).
- Execute (24 to 60): below 20%; first spends the bar; Mortal Strike "from 40". OK.
- Slam (20 to 60): Improved Slam L 42, R 39 (rank 1 already frees the swing); Mortal Strikes "from 40"; swing timer kept. OK.
- Spearing Strike (10 to 60, component row; talent R 25): raid build only, Battle Stance (mask), 75%, twice on dragons and giants. OK.
- Hamstring (10 to 60, ranks 1715 to 7373): Windfury Totem gate; 10 rage every rank; never from 25% down (Should); rage other attacks need. OK.
- Heroic Strike (10 to 60): at 90 or more (Should); Overpower 12, Execute 24, Mortal Strike 40. OK.

Fury (19 bands; leveling and raid orders identical)

- Bloodrage (10 to 60): 10 + 10, 60 s; Improved Bloodrage rank 1 at 59, rank 2 at 60. OK.
- Death Wish (10 to 60; talent 30, rows 10 to 28 component): pull if a second fits (210 s), execute phase, last 30 s. OK.
- Recklessness (50 to 60): execute phase or last 15 s; once a fight (1800000 ms). OK.
- Battle Shout (10 to 60): as Arms. OK.
- Execute (24 to 60): first below 20%; every point of rage; Bloodthirst "from 40". OK.
- Bloodthirst (40 to 60): "the button you lose the most by delaying". OK.
- Whirlwind (36 to 60): four targets, both weapons (1680; Dual Wield 20); Bloodthirst "from 40". OK.
- Reckless Swings (10 to 60; talent 31): behind Whirlwind "once you have it" (36); four white swings, 5 rage, quarter of base health; "healers cover" (Should, solo). OK.
- Berserker Rage (36 to 60): "From 36"; Improved Berserker Rage 2 at 35, 10 rage; 30 s (18499); Magmadar, Onyxia, Nefarian. OK.
- Overpower (12 to 32, levels 12 to 35): Battle Stance before 36; Whirlwind 36. OK.
- Heroic Strike (10 to 60): above 60; 25% stop once Execute (24) is known; no global. OK.

Protection (24 bands; leveling and raid orders identical)

- Defensive Stance (10 to 60): threat bonus; Revenge 14, Shield Block 16, Shield Wall 28 "all need it" (masks). OK (Should carried).
- Last Stand (10 to 60; talent 28, rows 10 to 26 component): below 35%; 30% for 20 s, lost after (12975). OK.
- Shield Wall (28 to 60): below 40%; 60% for 12 s; 15 min (871; no Improved Shield Wall in the build). OK.
- Bloodrage (10 to 60): before the pull, under 60; 10 + 10; Improved Bloodrage rank 1 at 10, rank 2 from 11. OK.
- Shield Block (16 to 60): 75% for 7 s, `gcd_ms` 0; every block lights Revenge (14). OK (Should carried).
- Revenge (14 to 60, ranks 6572 to 11601): dodge, parry or block; most threat per rage (5, 2 with Focused Rage 38, 39, 41). OK.
- Shield Slam (40 to 60, ranks 23922 to 23925): "a hard hit", "very high amount of threat" (client text), grows with block value. OK (round 22's fix holds).
- Thunder Clap (10 to 60, 6343 to 11581): 20% attack speed slow every rank; refresh under 3 s. OK.
- Demoralizing Shout (14 to 60, 1160 to 11556): attack power stripped at every rank; does not stack with Demoralizing Roar; refresh under 3 s. OK.
- Battle Shout (10 to 60): only when nobody else's is up; same aura at 60. OK.
- Sunder Armor (10 to 60, 7386 to 11597): five stacks, then threat. OK.
- Heroic Strike (10 to 60): above 40, queued, no global. OK.
