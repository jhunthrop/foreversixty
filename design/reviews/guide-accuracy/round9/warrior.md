# Guide accuracy round 9: Warrior

Audited on the round 9 audit tree (land/guides-r8 head 645caa983), client 1.60.1.70755, engine pin d1d9f2775 (`sim/enginever/version.go`). Files: `web/src/content/guides/warrior/{index,arms,fury,protection}.md`. I read every line, including frontmatter sources, Levels 1 to 19, Levels 20 to 60, stat, race and profession prose, and every card note in `data/curated/apl/warrior-{arms,fury,protection}.json`. I read the rounds 1 to 8 warrior reports and fix logs, `round4/warrior-levels.md`, race-sims for rounds 2 to 5 and the ledger first. Since round 8 the only warrior changes are the three round 8 fixes: Eureka! wording, the Thunder Clap note and Battle Shout raid-wide (`git diff 09b98d9e4 HEAD`). I re-checked all three against the client text.

## Builds decoded (unchanged; all legal on 70755)

- Arms `build` `13323213032215201/0505/205`, `raidBuild` `34323213032215201/0505002/2` (night-elf). 51 points each.
- Fury `build` `30323013002/051531025120100231/2` (night-elf). 51 points.
- Protection `build` `35300003/0/250533120330001351` (night-elf). 51 points.

I re-derived the shipped orders from `addon-builds.json` against `talents/warrior.json` tier and column (index 0 = level 10). They match rounds 5 to 8 exactly:
- Arms: IHS 10, Deflection 11-13, Improved Rend 14-16, Improved Charge 17-18, ITM 19-21, Improved Overpower 22-23, AM 24, Deep Wounds 25-27, 2H Spec 28-30, Impale 31-32, Bloodthrill 33-34, Sweeping Strikes 35, Weaponmaster 36-39, MS 40, Weaponmaster 41, Improved Slam 42-43, Cruelty 44-48, Unbridled Wrath 49-53, Improved Bloodrage 54-55, Iron Will 56-60.
- Fury: Cruelty 10-14, Lingering Rage 15, Unbridled Wrath 16-20, Furious Precision 21-23, Piercing Howl 24, Enrage 25-26, DW Spec 27-29, Death Wish 30, DW Spec 31-32, Reckless Swings 33, Improved Execute 34-35, Improved Whirlwind 36-37, Flurry 38-39, Bloodthirst 40, Flurry 41, IHS 42-44, Improved Rend 45-47, Improved Charge 48-49, ITM 50-52, AM 53, Deep Wounds 54-56, Impale 57-58, Improved Bloodrage 59-60.
- Protection: Improved Bloodrage 10-11, Shield Spec 12-16, Anticipation 17-21, Improved Revenge 22-24, Improved Thunder Clap 25-27, Last Stand 28, Master of Defense 29-30, Defiance 31-33, Improved Sunder Armor 34-36, Concussion Blow 37, Focused Rage 38-39, Shield Slam 40, Focused Rage 41, Bastion 42-46, IHS 47-49, Deflection 50-54, Improved Rend 55-57, Deep Wounds 58-60.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| WARRIOR-R9-1 | CONFIRMED | race (round 8 edit incomplete) | arms.md:111; fury.md:114 | "Gnome's Eureka!, 10% more damage and 10% less rage on your non-periodic abilities for 8 seconds every 2 minutes" | The warrior row's client text (1259813, cited in round 8) says "Your non-periodic **damaging** abilities have their Rage cost reduced by 10% and deal 10% more damage for 8 sec." Round 8 left out "damaging", so the page now gives a rage cut to every non-periodic ability. The engine's warrior Cost list (`sim/core/eureka_auto_gen.go` at d1d9f2775, Class_ClassWarrior) holds only damaging spells: Heroic Strike, Cleave, Mortal Strike, Bloodthirst, Whirlwind, Execute, Overpower, Slam, Hamstring, Revenge, Shield Slam, Thunder Clap, Intercept, Pummel, Shield Bash and Concussion Blow. Battle Shout, Sunder Armor, Death Wish, Sweeping Strikes and Demoralizing Shout are non-periodic, they cost rage, and none of them is on that list. The round 9 brief asks for every Eureka! line to be checked against "non-periodic damaging abilities". | Both pages: "Gnome's Eureka!, 10% more damage and 10% less rage on your non-periodic damaging abilities for 8 seconds every 2 minutes," (the rest of the sentence is unchanged). |

## Shoulds (wording, not findings)

- `warrior-arms.json` priorityList[5] (Overpower) note: "with Improved Overpower it nearly always crits." Overpower's crit is sheet crit plus 50%, so roughly three in four crit on raid gear. The page says "crits most of the time" (arms.md:68). Use the page's wording: "...and with Improved Overpower it crits most of the time."
- fury.md:92 "Berserker Rage breaks fear". The 70755 text (18499) only says "becoming immune to Fear and Incapacitate effects ... Lasts 10 sec". I found no source that settles whether it can be pressed while already feared in Forever (Wowpedia's Berserker Rage page says "break fear"). The sentence already tells you to hold it for the known fears, so nothing on the page depends on this. Not filed.
- fury.md:78 "a boss you stand behind never gives you a parry": the parry is yours, against the boss's swings. "A boss you stand behind never swings at you, so Deflection's parry does nothing" says it plainly.
- Carried from round 8, still open: fury.md:64 "the button you never delay" against Execute first at 20%; the index Battle Shout bullet; arms.md:111 "behind it" pointing back to Night Elf.

## Checked and right

- **Round 8 fixes re-checked**: Battle Shout lines (arms.md:60, fury.md:59 and both card notes) match 6673 to 25289: "all party or raid members within 20 yards ... Lasts 3 min". The Thunder Clap note ("a straight cut to the damage you take") matches 11581's 20% slow. Eureka! is right apart from R9-1.
- **Crit-damage lesson**: Impale (arms.md:68, fury.md:74), "220% of a normal hit instead of double", matches 16493 rank 2, "+20% to the critical strike damage bonus" (1 + 1.0 x 1.2 = 2.2). No other line on the warrior pages states crit damage. "A critical swing brings twice the rage" is about rage, not damage.
- **Levels 1 to 19, all three pages**, against `trainables/warrior.json`: Charge and Rend 4, Thunder Clap 6, Hamstring 8, Bloodrage, Defensive Stance, Sunder Armor and Taunt 10, Overpower and Shield Bash 12, Revenge, Demoralizing Shout and Tactical Mastery (1310185, "retain up to 10 Rage") 14, Shield Block 16. Every point by level matches the orders above. Rank values hold: IHS 1, Deflection 1% a rank, Improved Rend 35%, Improved Charge +6 at rank 2, Cruelty 5%, Lingering Rage 2 s at rank 1, Unbridled Wrath 60%, Improved Bloodrage 50%, and Shield Spec's 100% chance at rank 5 only. The 1 to 19 and 20 to 60 blocks read in order and do not contradict the strip. Victory Rush 20, Berserker Stance 30, Whirlwind 36, Execute 24 and Mortal Strike, Bloodthirst and Shield Slam 40 match too.
- **Talent text** (70755 rank tables): every value on the pages matches, including Weaponmaster 5% / 15% / 5%, Bloodthrill main hand 20%, Improved Slam -5 s and -0.5 s with no swing delay, Reckless Swings and Enrage. Reckless Swings is 1 s GCD, 15 s, 8 s and 25% of base health (spellconst 1329555), and the engine confirms its 4 charges are white swings only (`sim/warrior/reckless_swings.go` ProcMaskMeleeWhiteHit). Enrage is 2% a rank, so 4% at 2. All prerequisites are met, and the Protection Improved Shield Wall swap stays legal (25 points below tier 4, 31 below Bastion, 36 below Shield Slam).
- **Client spell text** (`spelltips/warrior.json`, `spellconst/warrior.json`): Execute 15 rage, 10 with Improved Execute; Slam 1.5 s cast and 15 s cooldown; Mortal Strike and Bloodthirst 6 s; Sweeping Strikes 5 attacks; Whirlwind 4 targets with both weapons; Recklessness 100% / 15 s / 30 min; Death Wish 10 rage / 30 s / 3 min / 5%; Shield Block 75% / 7 s / 2 attacks / 5 s with no GCD; Shield Wall 60% / 12 s / 15 min; Last Stand 30% / 20 s; Challenging Shout 10 yd / 6 s / 10 min; Thunder Clap 20 rage, 20% slow, 30 s, 4 targets; Demoralizing Shout 45 s (204 at 60 holds per round 8); Victory Rush 10% heal, 20 s window, 30 s; Bloodrage 10 + 10 over 10 s, 60 s; Charge rank 3 15 rage, out of combat only.
- **Death Wish while feared**: the 70755 SpellMisc row for 12328 is identical to Era's (`scratchpad/wago/SpellMisc.csv` against `SpellMisc_era.csv`), so the Classic behaviour of breaking fear stands.
- **Racials** (`spelltips/shared.json`; engine `sim/core/racials.go`): Elune's Light 10% / 15 s / 3 min, Blood Fury 10% AP / 15 s / 2 min, Berserking 10% / 15 s / 3 min, Touch of the Grave 20% with a 15 s internal cooldown, Wind Blessed 2%, Ruthless, Artisan's Precision, Quickness, Endurance and Stoneform (10% physical, 8 s, 3 min) all match. Tauren's racials (War Stomp, Cultivation, Endurance, Plainsrunning) carry no damage, so it comes last. Race order matches the race sims, and Skyborne is named only as the paid option.
- **Card notes**, apart from the Overpower SHOULD: every note matches its action, the APL thresholds (Arms HS 90, Bloodrage under 80, Recklessness swap at 25 rage or less; Fury HS 60 with none from 25%, Death Wish 210 s / execute / 30 s; Protection Bloodrage under 60, Last Stand 35%, Shield Wall 40%, HS 40) and the client text.
- **Stat prose** on all three pages explains and never ranks. Each reason holds on 70755. Nothing else on the pages ranks stats or gear by a nightly value.
- **Source links**, fetched this round: the Blizzard recap, Wowhead Skyborne First Look, Power Up Gaming, Talents Forever data.json, both Warcraft Tavern Classic and Forever warrior pages, the Warcraft Tavern Fury guide, and foreversixty.gg /sim/weights, /sim/drops, /planner and /bis/warrior/protection all return 200. Icy Veins racial rework returns 403, so its status is unknown, not dead. The source-link rule holds.
- No em or en dashes, no quoted nightly numbers, and no site or sim talk outside the planner links.

## Decisions

- **Eureka! names "damaging" (R9-1).** Settled by the client text and the engine's warrior cost list. Shouts, Sunder Armor, Death Wish and Sweeping Strikes get no rage cut.
- **Berserker Rage "breaks fear" stays.** The client text cannot settle it and the line does not depend on it.

## Sims requested

None new. Still open from earlier rounds, none of them claimed in prose:
- Dwarf against Human breakdown (Arms and Fury)
- Elune's Light and Blood Fury held-timing comparisons
- the Stoneform rotation line
- S2 Improved Berserker Rage
- S3 Spearing Strike
- S4 a third Improved Heroic Strike point
- the Protection Improved Shield Wall build and Shield Wall line
- removing the `expertise` weight

## Sim findings

None new. The engine's warrior Eureka! lists match the client text: a damage and cost list of damaging abilities only, and an empty periodic list.

## Sources

- Client data 1.60.1.70755: `talents/warrior.json`, `trainables/warrior.json`, `spelltips/warrior.json`, `spelltips/shared.json`, `spellconst/warrior.json` (1329555, 12292, 2687, 29131, 1719, 12328, 871, 2565, 402927), `addon-builds.json`, `bis/warrior-{arms,fury,protection}.json`.
- wago.tools SpellMisc at 1.60.1.70753 and Era (scratchpad copies `wago/SpellMisc.csv`, `wago/SpellMisc_era.csv`), row 12328.
- Engine d1d9f2775: `sim/core/eureka_auto_gen.go` (Class_ClassWarrior, 1259813), `sim/warrior/reckless_swings.go`, `sim/core/racials.go`.
- Wowhead Forever tooltip, Eureka! (Warrior) 1259813, https://nether.wowhead.com/forever/tooltip/spell/1259813.
- Wowpedia, Berserker Rage, https://wowpedia.fandom.com/wiki/Berserker_Rage.

## Fix log (round 9)

Applied by the integrator; see the round 9 ledger entry and the commit for the exact lines.
