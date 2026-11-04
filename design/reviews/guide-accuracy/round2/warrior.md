# Guide accuracy round 2: Warrior

Audited on land/guides-r1 609eb471c, client 1.60.1.70755, engine pin d1d9f2775. Files: `web/src/content/guides/warrior/{index,arms,fury,protection}.md`, every line. Round 1 report and ledger read first; nothing fixed, sim-listed or owner-held in round 1 is re-filed.

## Builds decoded (unchanged since round 1; all legal on 70755)

Frontmatter strings match round 1 and `data/builds/1.60.1.70755/addon-builds.json` at level 60:
- Arms `build` `13323213032215201/0505/205`, `raidBuild` `34323213032215201/0505002/2`
- Fury `build` `30323013002/051531025120100231/2` (leveling and raid orders are identical)
- Protection `build` `35300003/0/250533120330001351`

Shipped point orders (index 0 = level 10), the parts the prose leans on:
- Arms leveling: Improved Overpower 22-23, Anger Management 24, Deep Wounds 25-27, Two-Handed Spec 28-30, Impale 31-32, **Bloodthrill 33-34**, Sweeping Strikes 35, Weaponmaster 36-39, Mortal Strike 40, Improved Slam 42-43, then Fury and Protection.
- Fury: Cruelty 10-14, Lingering Rage 15, Unbridled Wrath 16-20, ..., Death Wish 30, Improved Whirlwind 36-37, Bloodthirst 40, Flurry 41, then Arms 42-58, Improved Bloodrage 59-60.
- Protection: Improved Bloodrage 10-11, Shield Specialization 12-16, Anticipation 17-21, ..., Concussion Blow 37, Focused Rage 38-39, Shield Slam 40, **Focused Rage 41**, Bastion 42-46, Improved Heroic Strike 47-49, Deflection 50-54, Improved Rend 55-57, Deep Wounds 58-60.

Rechecked against 70755 and found right: every talent value quoted on all four pages (rank tables in `talents/warrior.json`, including Lingering Rage rank 1 = 2 s, Enrage rank 2 = 4%, Improved Thunder Clap 2 per rank so 3 to 1 is 4 rage, Improved Shield Wall 11 min, Weaponmaster sword 5% extra attack); prerequisites (Anger Management behind Improved Tactical Mastery 3, Mortal Strike behind Sweeping Strikes, Shield Slam behind Concussion Blow, Master of Defense behind Shield Specialization 5); trainer levels and costs in `trainables/warrior.json` (Berserker Stance and Berserker Rage 30, Whirlwind 36 with a 10 s cooldown, Recklessness 50, Execute 15 rage, Death Wish 10 rage, Bloodthirst and Mortal Strike 6 s, Shield Block 16, Defensive Stance 10, Challenging Shout 10 min); `spellconst/warrior.json` (Reckless Swings 1329555: 1 s global, 15 s cooldown, 8 s, 25% base health, 5 rage per swing; Shield Block 2565 off the global, 7 s; Sweeping Strikes 20 s; Demoralizing Shout 11556 -196 plus -1.4 a level over 54 = 204 at 60; Recklessness 15 s); APL orders in `data/curated/apl/warrior-*.json` against every rotation paragraph; racials in `races.json`. Round 1 research sources re-read: Kaivax's 2 Oct beta warrior update does say crit rage went to 100% (was 75%); wow.gg's Forever Protection overview does take Improved Shield Wall 2/2 and says it shortens the cooldown. No em or en dashes, no quoted nightly numbers, no site/sim talk outside the owner-approved planner links.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| WARRIOR-R2-1 | CONFIRMED | rotation | arms.md:83 | "Spend the Charge's rage on Battle Shout as you arrive ... then press Bloodrage as your first swing lands. That is a Rend and a Mortal Strike in the first two globals." | Battle Shout is on the global cooldown (`spellconst/warrior.json` 25289 `gcd_ms 1500`), so when you shout, the shout is your first global and Rend and Mortal Strike are the second and third. The APL's own order (`warrior-arms.json` prepull: Charge, then Battle Shout; priority 3 shout before 4 Rend) plays it that way. | "...then press Bloodrage as your first swing lands. That is a shout, a Rend and a Mortal Strike in your first three globals, or Rend and Mortal Strike in the first two when another warrior's shout is already on you." |
| WARRIOR-R2-2 | CONFIRMED | leveling / spell | arms.md:115, fury.md:118 | "Victory Rush after every kill" | Victory Rush 402927 has a 30 s cooldown (`spellconst/warrior.json` `cooldown_ms 30000`; `trainables/warrior.json` Victory Rush `cooldown_ms 30000`) and must be used within 20 s of the kill, so on quick pulls it is up every other kill at best. Blizzard forums, "Warrior Feedback: Victory Rush Cooldown Conflicts With Its Own Mechanic" describes exactly this. | Both pages: "Charge every pull, and Victory Rush whenever it is up after a kill: it heals 10% and comes back every 30 seconds, so open the next pull with it if the last one died fast." (Fury keeps "and Bloodrage on cooldown" after it.) |
| WARRIOR-R2-3 | CONFIRMED | leveling | protection.md:124 | "Shield Slam is the 31st point, so it arrives at 40; Bastion, then Improved Heroic Strike and Deflection in Arms, fill the levels after it." | `addon-builds.json` warrior-protection: Focused Rage 3 is the point at 41, Bastion 42-46, Improved Heroic Strike 47-49, Deflection 50-54, then Improved Rend and Deep Wounds 55-60. The prose skips the level 41 point. | "Shield Slam is the 31st point, so it arrives at 40. The last point of Focused Rage comes at 41, then Bastion, and the levels after that go to Improved Heroic Strike, Deflection and Deep Wounds in Arms." |
| WARRIOR-R2-4 | LIKELY | leveling | arms.md:115 | "Take Improved Overpower and Bloodthrill as soon as the build reaches them: together they turn every pull into Overpower spam." | Bloodthrill is reachable at 30 (tier 4, 20 Arms points spent by 29), but the shipped order takes Deep Wounds, Two-Handed Weapon Specialization and Impale first and Bloodthrill at 33-34. Read plainly, "as soon as" tells the player to take it at 30, which the strip on the same page does not show. | "Improved Overpower comes in at 22 and Bloodthrill at 33, after Deep Wounds and Impale; once both are in, keep Rend on everything and every pull turns into Overpower spam." |
| WARRIOR-R2-5 | LIKELY | talent / voice | fury.md:76 | "Tactical Mastery itself does little for a warrior who never leaves Berserker Stance; it is the way in." | The talent is Improved Tactical Mastery (`talents/warrior.json` Arms tier 1); Tactical Mastery is the baseline level 14 passive the class page describes (index.md:26). The same page also tells you to Charge freely on trash and to fight in Battle Stance until 36 (fury.md:82, :118), so the warrior does leave Berserker Stance, and those 25 rage carry the Charge rage through the swap. | "Improved Tactical Mastery is the way in. In a raid you never leave Berserker Stance, but on trash it lets a Charge's rage come with you when you swap back." |

## Round 1 research answers re-checked against the client description text

Per the coordinator's round 1 lesson (a description beats an inferred flag), every warrior cost or effect claim was re-read in `spelltips/warrior.json` (the client descriptions):
- Rage from a critical swing, 100%: Kaivax's 2 Oct post says it in words (quoted under Sources). Holds.
- Improved Shield Wall: talent text "Reduces the cooldown of your Shield Wall ability by 11.0 min"; Shield Wall 871 "Reduces the damage taken from all attacks by 60% for 12 sec", cooldown 900000. 4 minutes holds.
- Death Wish while feared: 12328 description "makes you immune to Fear effects" and costs 10 rage, 3 min. It does not say "usable while feared"; that part rests on the dispel-on-immunity attribute plus the Vanilla 1.6.0 change. Nothing in the description contradicts it, and Forever did not touch the spell's fear text, so it holds as a Classic mechanic Forever kept. Not filed.
- Arms opener: Charge 11578 description "generate 15 Rage ... Cannot be used in combat"; Bloodrage 2687 "Generates 10 rage at the cost of health, and then generates an additional 10 rage over 10 sec". Holds.
- Also read for this round's claims: Victory Rush 402927 "Only useable within 20 sec after you kill a non-trivial enemy", cooldown 30000 (WARRIOR-R2-2); Shield Block 2565 "Increases chance to block by 75% for 7 sec, but will only block 2 attacks"; Whirlwind 1680 "up to 4 enemies ... both melee weapons"; Thunder Clap 11581 "20% for 30 sec ... up to 4 targets"; Demoralizing Shout 11556 "by 196 for 45 sec"; Berserker Rage 18499 "immune to Fear and Incapacitate effects and generating extra rage when taking damage. Lasts 10 sec."; Recklessness 1719 "100% increased critical strike chance ... immune to Fear ... 15 sec, but all damage taken is increased by 20%". All agree with the prose.

Data note for the site lane (not a guide finding): `spelltips/warrior.json` gives every Slam rank (1240193, 1464, 8820, 11604, 11605) `cooldown_ms 18000`, while `spellconst/warrior.json` and `trainables/warrior.json` say 15000 and the Wowhead Forever tooltip for 1240193 shows "15 sec cooldown". index.md:29 ("Slam has a 15 second cooldown") is right; any tooltip rendered from spelltips would show 18 s and contradict it.

## Sim findings

- **S5 (NEEDS-SIM).** DPS race order. 70755 racials give Night Elf Elune's Light (10% crit for 15 s, 3 min cooldown) and Dwarf Artisan's Precision (1% hit) (`races.json`), and neither race is named on the Arms or Fury pages, which recommend Human/Orc and Human/Troll. Measure Arms and Fury at 60 raid for Human, Orc, Troll, Night Elf (Elune's Light at 20% with Recklessness) and Dwarf. If Night Elf or Dwarf lands inside the top two, the race lines need a sentence for them. The current lines state nothing false, so this is not a finding yet.
- Round 1 S2 to S4, the Protection Improved Shield Wall build swap and the `expertise` weight removal are still open for the sim round; not re-filed.

## Owner questions

None. Every open item is a data fix or a sim question.

## Sources

- Blizzard forums, Warrior Updates in Today's Beta Build (Kaivax, 2026-10-02), https://us.forums.blizzard.com/en/wow/t/2369360 (re-checked: "Additional Rage generated from landing Critical Strikes increased to 100% increased Rage (was 75%)"; Lingering Rage 2 to 10 s; Berserker Rage at 30; Bloodthirst 45% AP).
- Blizzard forums, Warrior Feedback: Victory Rush Cooldown Conflicts With Its Own Mechanic, https://us.forums.blizzard.com/en/wow/t/2379449/1 (30 s cooldown against a 20 s use window).
- Wowhead Forever tooltip, Slam 1240193, https://nether.wowhead.com/forever/tooltip/spell/1240193 (15 sec cooldown, 1.5 sec cast, 15 rage).
- wow.gg, Protection Warrior Forever overview, https://wow.gg/guides/warrior-protection-forever-overview (re-checked: Improved Shield Wall 2/2, "shortens the cooldown of Shield Wall rather than extending its duration").

## Fix log (round 2)

Findings:
- WARRIOR-R2-1 fixed, arms.md:88 (shout, Rend, Mortal Strike in the first three globals; two when another warrior's shout is on you).
- WARRIOR-R2-2 fixed, arms.md:122 and fury.md:125 (Victory Rush whenever it is up after a kill; heals 10%, 30 s cooldown, per spelltips 402927).
- WARRIOR-R2-3 fixed, protection.md:126 (Focused Rage at 41, then Bastion, then Improved Heroic Strike, Deflection and Deep Wounds).
- WARRIOR-R2-4 fixed, arms.md:122 (Improved Overpower at 22, Bloodthrill at 33 after Deep Wounds and Impale, per addon-builds.json).
- WARRIOR-R2-5 fixed, fury.md:79 (Improved Tactical Mastery named; carries Charge rage through the swap on trash).
- Brief rules: planner links set to the one form as their own paragraph (arms.md:84, fury.md:83, protection.md:90); site source labels reworded with no pipeline nouns on all three spec pages.

Integrator decision, races (S5 closed). Raid builds, band 60, 10,000 iterations, seeds 1|50001, error about ±0.45 a run:
- Arms Alliance: High Order Skyborne 755.3|755.1 leads Night Elf 752.1|751.8 by about 3, well beyond error. Pick High Order Skyborne.
- Arms Horde: Windshaper Skyborne 766.5|765.2 leads Undead 765.1|764.4 by about 1.1 on the means, beyond error. No source backs any Horde race for Arms (Warcraft Tavern names none; one roundup's Orc pick rests on the axe racial Forever removed). Pick Windshaper Skyborne.
- Fury Alliance: High Order Skyborne 832.8|832.5 leads Night Elf 830.7|831.9 by about 1.3, beyond error. Pick High Order Skyborne.
- Fury Horde: Undead 835.6|835.0 leads Windshaper Skyborne 833.8|834.9 by about 0.95 on the means, at the edge of error (the seed-50001 runs overlap). The one Forever source that splits warrior races by spec names Undead for Fury for Touch of the Grave, so consensus breaks the tie. Pick Undead. Warcraft Tavern's Troll for Fury is contradicted beyond error (Troll last, about 4 behind).
- Applied: arms.md:8 `recommendedRaces: [high-order-skyborne, windshaper-skyborne]`; fury.md:7 `[high-order-skyborne, undead]`; race field of arms.md:6-7 and fury.md:6 FS1 strings set to `high-order-skyborne` (talents unchanged). Race prose rewritten, arms.md:112-114 and fury.md:115-117, with every racial value read from spelltips/shared.json racials and races.json (Wind Blessed 2% haste; Touch of the Grave 20% chance, 15 s lockout; Elune's Light 10% crit 15 s / 3 min; Blood Fury 10% AP 15 s / 2 min; Berserking 10% 15 s / 3 min; Eureka! 10% 8 s / 2 min; Artisan's Precision 1% hit). Warcraft Tavern Forever warrior guide added to both pages' sources. Protection untouched (not in scope).
- Consensus research: no Forever source ranks Skyborne for warriors (RankedBoost grades it low but its warrior page still cites removed weapon racials); wowtbc.gg says every race is within 1% for Fury; Warcraft Tavern names Human for Arms (for Spirit) and Troll for Fury; one roundup names Undead for Fury via Touch of the Grave. With no consensus on the Skyborne cases, the beyond-error measurements decide.

New SIM findings: none. Note for the integrator: the race-table component's picks line now names Skyborne and Undead for these pages; check it renders Wind Blessed and Touch of the Grave as the reasons (race-table.ts maps both).

Sources (decision): Warcraft Tavern, Forever warrior guide, https://www.warcrafttavern.com/forever/guides/warrior/ ; wowtbc.gg Forever Fury guide, https://wowtbc.gg/warcraftforever/class-guides/fury-warrior/ ; Warcraft Tavern, Skyborne racials, https://www.warcrafttavern.com/forever/guides/skyborne-racials/ (Wind Blessed 2%) ; MMO-Champion repost of Blizzard's 26 and 27 Oct racial tuning, https://www.mmo-champion.com/content/14645-WoW-Forever-Beta-Development-Notes-Racial-Abilities-and-Traits-Tuning (Wind Blessed 3 to 2%, Touch of the Grave 25 to 15 s) ; RankedBoost warrior race page, https://rankedboost.com/world-of-warcraft/forever/class/warrior/ (stale, axe racial) ; Undead-for-Fury roundup as surfaced by search among https://timesaver.gg/blog/wow-forever-racial-changes-best-race-every-class and https://gamer-guides.com/warcraft-forever/racials/ (page not fetched directly).
