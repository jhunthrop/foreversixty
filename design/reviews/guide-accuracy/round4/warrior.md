# Guide accuracy round 4: Warrior

Audited on land/guides-r4 752c02e43, client 1.60.1.70755, engine pin d1d9f2775. Files: `web/src/content/guides/warrior/{index,arms,fury,protection}.md`, every line, plus the card notes in `data/curated/apl/warrior-{arms,fury,protection}.json`. Read first: round 1, 2 and 3 warrior reports, `round2/race-sims.md`, `round3/race-sims.md`, the ledger and `round4/warrior-levels.md`. Nothing fixed, sim-listed or handed to the data lane is re-filed unless the page now states it falsely.

## Builds decoded (unchanged; all legal on 70755)

- Arms `build` `13323213032215201/0505/205`, `raidBuild` `34323213032215201/0505002/2` (race night-elf)
- Fury `build` `30323013002/051531025120100231/2` (race night-elf; leveling and raid orders identical)
- Protection `build` `35300003/0/250533120330001351` (race night-elf)

Shipped point orders (`addon-builds.json`, index 0 = level 10), re-derived this round against `talents/warrior.json` tier/column:
- Arms leveling: Improved Heroic Strike 10, Deflection 11-13, Improved Rend 14-16, Improved Charge 17-18, Improved Tactical Mastery 19-21, Improved Overpower 22-23, Anger Management 24, Deep Wounds 25-27, Two-Handed Spec 28-30, Impale 31-32, Bloodthrill 33-34, Sweeping Strikes 35, Weaponmaster 36-39, Mortal Strike 40, Weaponmaster 41, Improved Slam 42-43, Cruelty 44-48, Unbridled Wrath 49-53, Improved Bloodrage 54-55, Iron Will 56-60.
- Fury: Cruelty 10-14, Lingering Rage 15, Unbridled Wrath 16-20, Furious Precision 21-23, Piercing Howl 24, Enrage 25-26, Dual Wield Spec 27-29, Death Wish 30, Dual Wield Spec 31-32, Reckless Swings 33, Improved Execute 34-35, Improved Whirlwind 36-37, Flurry 38-39, Bloodthirst 40, Flurry 41, Arms 42-58, Improved Bloodrage 59-60.
- Protection: Improved Bloodrage 10-11, Shield Specialization 12-16, Anticipation 17-21, Improved Revenge 22-24, Improved Thunder Clap 25-27, Last Stand 28, Master of Defense 29-30, Defiance 31-33, Improved Sunder Armor 34-36, Concussion Blow 37, Focused Rage 38-39, Shield Slam 40, Focused Rage 41, Bastion 42-46, Improved Heroic Strike 47-49, Deflection 50-54, Improved Rend 55-57, Deep Wounds 58-60.

## Levels 1 to 19 (new this round): checked

- Trainer levels against `trainables/warrior.json`: Battle Shout 1, Charge 4, Rend 4, Thunder Clap 6, Hamstring 8, Bloodrage / Defensive Stance / Sunder Armor / Taunt 10, Overpower and Shield Bash 12, Revenge / Demoralizing Shout / Tactical Mastery 14, Shield Block 16, Victory Rush / Slam / Cleave / Retaliation 20. All right.
- Tactical Mastery 1310185: Wowhead Forever tooltip "You retain up to 10 Rage when you change Stances". Right.
- Charge rank 1 is 9 rage (`spelltips/warrior.json`), so "Improved Charge at 17 and 18 for 6 more Rage on every opener" is right at every rank.
- Talent points in all three Levels 1 to 19 lists match the shipped orders above exactly (Arms 10/11-13/14-16/17-18; Fury 10-14/15/16-20; Protection 10-11/12-16/17). Rank values right: Improved Heroic Strike rank 1 = 1 rage, Cruelty 5 = 5%, Lingering Rage 1 = 2 s, Unbridled Wrath 5 = 60%, Improved Bloodrage 2 = 50%, Shield Specialization rank 5 = 100% chance at 5 rage (ranks 1-4 are a chance, so "from the fifth point every block" is right), Anticipation = Defense.
- Level 10 class quest teaching Defensive Stance, Taunt and Sunder Armor: Warcraft Tavern's Forever leveling guide (class quests table) and Method's Forever leveling guide both say so; Method describes the follow-up quest that rewards "a powerful one hand weapon with several weapon type options". Dual Wield at 20 from the trainer: Warcraft Tavern Forever. Shield Block "two attacks": client text "will only block 2 attacks". Mail until plate at 40: owner class rule.
- No contradiction with the Levels 20 to 60 text or the strip (Fury fights in Battle Stance until 36 in both; Protection Shield Slam at 40 and Focused Rage at 41 in both).

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| WARRIOR-R4-1 | CONFIRMED | spell (round 3 edit wrong) | protection.md:98 | "**Demoralizing Shout.** 196 less melee attack power at 60" | 196 is rank 5's value at its own learn level, 54. `spellconst/warrior.json` 11556: `amount -196`, `spell_level 54`, `points_per_level -1.4`, `max_level 64`, so at 60 it is 196 + 6 x 1.4 = 204. The engine models it the same way (`sim/warrior/demoralizing_shout.go`: "a further 1.4 per caster level up to the rank's max level"). The page also contradicts itself: the Protection card note for the same action (`warrior-protection.json` priority 9) says "204 less melee attack power at 60". Round 3's rendered-review MUST changed the prose from 204 to 196 by reading the base tooltip. | "**Demoralizing Shout.** 204 less melee attack power at 60 takes a real slice off every swing, so keep it on the boss and refresh it in its last 3 seconds, unless a druid or another warrior already holds it: two do not stack." |
| WARRIOR-R4-2 | CONFIRMED | source link | arms.md:35-36, fury.md:34-35, protection.md:42-43 | label 'Forever Sixty, where the raid gear drops', url https://foreversixty.gg/sources | The URL resolves (200) to "How sources are classified": the page that explains the Blizzard / Datamined / Community / This site pills and what "single-source" means. Nothing on it says where raid gear drops. | Arms and Fury: url https://foreversixty.gg/sim/drops, label 'Forever Sixty, which boss drops are upgrades' (the Droptimizer page: "Every item those bosses drop is simulated ... the upgrades are listed by boss"). Protection: that page is damage specs only, so url https://foreversixty.gg/bis/warrior/protection, label 'Forever Sixty, Protection best in slot'. |
| WARRIOR-R4-3 | CONFIRMED | source link | arms.md:32-33, fury.md:31-32, protection.md:39-40 | label 'Forever Sixty, how Arms plays' / 'how Fury plays' / 'how Protection plays', url https://foreversixty.gg/sim/specs | The URL resolves to "Spec support": which specs the simulator models and how close each is to real parses (`web/src/pages/sim/specs.astro`: "Which World of Warcraft: Forever specs the simulator models, how close each one is to real parses"). It says nothing about how any spec plays, and for Protection it says the opposite of the label: "Healing and tanking specs are not simulated yet." | Arms and Fury: relabel 'Forever Sixty, which specs the simulator models' (or drop the entry). Protection: drop the entry. |
| WARRIOR-R4-4 | CONFIRMED | stale / class page | index.md:32 | "**Sweeping Strikes** lasts 20 seconds." (under "What changed from 1.12") | Not a change. Classic Era client data shows the same 20 s buff: Wowhead Classic tooltip 12292 buff "Your next 5 melee attacks strike an additional nearby opponent. 20 seconds remaining" (and Requires Battle Stance, 30 Rage, 30 s cooldown, all identical to Forever's 12292). The "lasts 10 sec" wiki text round 3 weighed is post-1.12: Warcraft Wiki's own history line "Sweeping Strikes now lasts 10 seconds and affects your next 10 swings" is patch 3.0.2. Forever's 12292 is unchanged from 1.12 in every field, so the bullet lists something that did not change. | Delete the bullet. |
| WARRIOR-R4-5 | LIKELY | source label | index.md:17, arms.md:17, fury.md:16, protection.md:30 | label 'Talents Forever (demo transcription)' | The URL resolves; its own `_readme` says the talents, racials and spellbooks now "come from the WoW Forever beta client's own data files (the site began as frame-by-frame reads of the BlizzCon 2026 demo)". The label describes what the export used to be. | Relabel 'Talents Forever, talent calculator from the beta client'. |
| WARRIOR-R4-6 | LIKELY | race | arms.md:114 | "Troll's Berserking only covers 15 seconds in every 3 minutes, so it trails the races that are always on." | `round2/race-sims.md` Arms Horde: Tauren 753.85 is last, behind Troll 757.86. Tauren's Endurance is always on, and Troll beats it. The sentence reads as Troll last on the Horde. Only the single-seed table has Tauren, so LIKELY. | "Troll's Berserking only covers 15 seconds in every 3 minutes, so it trails Undead and Orc; Tauren brings nothing for damage and comes last." |

## Checked and right

- Every talent value on all four pages against the 70755 rank tables (Bloodthrill 20%, Improved Overpower 50%, Weaponmaster 5 ranks, Improved Rend 35%, Deep Wounds 60%, Impale 20%, Two-Handed Spec 6%, Improved Tactical Mastery +15, Improved Slam -5 s / -0.5 s, Flurry 30%, Bloodthirst 45% AP, Death Wish 20% / 5% / 30 s, Dual Wield Spec 25% / 50%, Furious Precision 10%, Improved Execute -5, Improved Whirlwind -2 s / 30%, Reckless Swings 4 swings / 5 rage / 8 s, Enrage 2 = 4%, Piercing Howl 10 yd / 6 s, Bastion 10%, Defiance 15%, Master of Defense 2 = 100%, Anticipation 20, Improved Revenge 60%, Improved Thunder Clap 6, Improved Sunder Armor 3, Focused Rage 3, Last Stand 30% / 20 s, Concussion Blow 5 s, Vanguard, Improved Shield Bash, Improved Shield Wall 11 min).
- Client descriptions: Mortal Strike 50% healing cut; Shield Block 75% / 7 s / 2 attacks, 5 s cooldown; Shield Wall 60% / 12 s / 15 min; Recklessness 100% crit / 15 s / 20% damage taken; Death Wish fear immunity; racials in `spelltips/shared.json` (Elune's Light, Touch of the Grave, Wind Blessed, Blood Fury, Berserking, Eureka!, Ruthless, Artisan's Precision, Stoneform, Endurance, Quickness 1% dodge and 2% speed).
- Rotation paragraphs and every card note against the APL orders and thresholds (Arms Charge then shout prepull, Bloodrage under 80, Heroic Strike at 90, Hamstring only under Windfury and not from 25%, Recklessness swap at 25 rage or less; Fury Death Wish 210 s / execute / last 30 s, Recklessness execute / last 15 s, Heroic Strike 60 and none from 25%; Protection Last Stand 35%, Shield Wall 40%, Thunder Clap and Demoralizing Shout refresh at 3 s, Heroic Strike above 40). The only card-note/prose mismatch is R4-1.
- Round 3's three fixes (Elune's Light outside Recklessness, Skyborne vs Undead on the Horde, Fury first ten points) are right. Race prose against `round2/race-sims.md` and `round3/race-sims.md` (Protection: Night Elf least damage taken, Tauren lowest TMI).
- Source links: Blizzard Deep Dive recap, Power Up Gaming, Warcraft Tavern Forever warrior and Fury guides, Warcraft Tavern Classic PvE DPS warrior (has an Arms talents and Arms rotation section), foreversixty.gg weights and planner all resolve and match their labels. Icy Veins racial rework and Wowhead Skyborne First Look return 403 to automated fetches: unknown, not dead (Wowhead's title does match the label).
- No em or en dashes, no quoted nightly numbers, no site or sim talk outside the one-form planner links.

## Decisions

- **Demoralizing Shout is 204 at 60** (R4-1): the client's points-per-level field settles it, and the engine already uses it. The tooltip shows the base 196 because it is the rank's level 54 value; the prose speaks to a level 60 tank.
- **Sweeping Strikes bullet goes** (R4-4): Classic Era client data, not wiki text, is the 1.12 reference; it matches Forever field for field.
- **Dwarf "1% hit matters least" line (arms.md:114, fury.md:117): not re-filed.** Round 3 decided it and sim request 3 there is still open.
- **Class quest and Dual Wield claims stand** on Warcraft Tavern Forever and Method Forever leveling guides (both cited below).

## Sims requested

None new. Still open from round 3: Elune's Light timing confirmation, Dwarf vs Human breakdown, Stoneform rotation line for the tank fight; from earlier rounds: S2 Improved Berserker Rage, S3 Spearing Strike, S4 third Improved Heroic Strike point, Protection Improved Shield Wall build and Shield Wall APL line, `expertise` weight removal.

## Sim findings

None new. The engine's Demoralizing Shout scaling (1.4 a level to the rank's max level) matches the client; the guide prose is what is wrong.

## Sources

- Wowhead Classic tooltip, Sweeping Strikes 12292, https://nether.wowhead.com/classic/tooltip/spell/12292 (buff "20 seconds remaining"; Battle Stance, 30 Rage, 30 s cooldown).
- Wowhead Forever tooltip, Sweeping Strikes 12292, https://nether.wowhead.com/forever/tooltip/spell/12292 (same text, requires level 30).
- Warcraft Wiki, Sweeping Strikes, https://warcraft.wiki.gg/wiki/Sweeping_Strikes (3.0.2: "now lasts 10 seconds and affects your next 10 swings").
- Wowhead Forever tooltip, Tactical Mastery 1310185, https://nether.wowhead.com/forever/tooltip/spell/1310185 (retain up to 10 Rage, 25 with Improved Tactical Mastery; level 14).
- Warcraft Tavern, Forever Warrior Leveling Guide, https://www.warcrafttavern.com/forever/guides/warrior-leveling/ (level 10 class quest: Defensive Stance, Taunt, Sunder Armor; Dual Wield at 20 from the trainer; trainer levels 1 to 20).
- Method, WoW Forever Warrior Leveling Guide and Talents, https://www.method.gg/wow-forever/wow-forever-warrior-leveling-guide-and-talents (follow-up quest rewards a one-hand weapon with several weapon types).
- Forever Sixty, How sources are classified, https://foreversixty.gg/sources (R4-2: page content).
- Forever Sixty, Spec support, https://foreversixty.gg/sim/specs (R4-3: "Healing and tanking specs are not simulated yet").
- Talents Forever data export, https://talentsforever.com/data.json (R4-5: `_readme`).

## Fix log (round 4)

- WARRIOR-R4-1 fixed: protection.md:92 Demoralizing Shout back to 204 less melee attack power at 60 (rank 5 is 196 at its level 54 plus 1.4 a level); now matches the card note.
- WARRIOR-R4-2 fixed per brief: "Forever Sixty, where the raid gear drops" now links https://foreversixty.gg/sim/drops on arms.md:36, fury.md:35, protection.md:37 (the brief's raid loot rule, kept on Protection since that page lists drops by boss).
- WARRIOR-R4-3 per brief: Arms and Fury keep "Forever Sixty, how <Spec> plays" on /sim/specs (brief label); Protection entry deleted. Protection's "what each stat is worth" (/sim/weights) entry deleted too (tank page).
- WARRIOR-R4-4 fixed: index.md "Sweeping Strikes lasts 20 seconds" bullet deleted.
- WARRIOR-R4-5 fixed: "Talents Forever, talent data from the beta client" on index.md:17, arms.md:17, fury.md:16, protection.md:30.
- WARRIOR-R4-6 fixed: arms.md:114 Troll now trails Undead and Orc, Tauren last (round2/race-sims.md Arms Horde). Fury's matching line left: Tauren was not run for Fury and Troll is last of the Horde races that were.
- Decisions applied: Demoralizing Shout 204, Sweeping Strikes bullet gone; Dwarf line and class quest / Dual Wield claims stand as decided. No race, build, action or Levels 1 to 19 changes; no wago links on these pages.
- New SIM findings: none.
