# Guide accuracy round 8: Warrior

Audited on the round 8 audit tree (land/guides-r7 head 09b98d9e4), client 1.60.1.70755, engine pin d1d9f2775 (`sim/enginever/version.go`). Files: `web/src/content/guides/warrior/{index,arms,fury,protection}.md`. I read every line, including frontmatter sources, Levels 1 to 19, Levels 20 to 60, stat prose and race prose, and every card note in `data/curated/apl/warrior-{arms,fury,protection}.json`. I read the rounds 1 to 7 warrior reports with their fix logs, `round4/warrior-levels.md`, race-sims for rounds 2 to 5 and the ledger first. Anything already fixed, listed for the sim round or handed to the data or component lanes is not filed again.

## Builds decoded (unchanged; all legal on 70755)

- Arms `build` `13323213032215201/0505/205`, `raidBuild` `34323213032215201/0505002/2` (night-elf)
- Fury `build` `30323013002/051531025120100231/2` (night-elf)
- Protection `build` `35300003/0/250533120330001351` (night-elf)

`addon-builds.json` and `talents/warrior.json` have not changed since round 7 (last commit bce8c200b). The shipped orders are the ones round 7 lists:
- Arms: IHS 10, Deflection 11-13, Improved Rend 14-16, Improved Charge 17-18, ITM 19-21, Improved Overpower 22-23, AM 24, Deep Wounds 25-27, 2H Spec 28-30, Impale 31-32, Bloodthrill 33-34, Sweeping Strikes 35, Weaponmaster 36-39, MS 40, Weaponmaster 41, Improved Slam 42-43, Cruelty 44-48, Unbridled Wrath 49-53, Improved Bloodrage 54-55, Iron Will 56-60.
- Fury: Cruelty 10-14, Lingering Rage 15, Unbridled Wrath 16-20, Furious Precision 21-23, Piercing Howl 24, Enrage 25-26, DW Spec 27-29, Death Wish 30, DW Spec 31-32, Reckless Swings 33, Improved Execute 34-35, Improved Whirlwind 36-37, Flurry 38-39, Bloodthirst 40, Flurry 41, Arms 42-58, Improved Bloodrage 59-60.
- Protection: Improved Bloodrage 10-11, Shield Spec 12-16, Anticipation 17-21, Improved Revenge 22-24, Improved Thunder Clap 25-27, Last Stand 28, Master of Defense 29-30, Defiance 31-33, Improved Sunder Armor 34-36, Concussion Blow 37, Focused Rage 38-39, Shield Slam 40, Focused Rage 41, Bastion 42-46, IHS 47-49, Deflection 50-54, Improved Rend 55-57, Deep Wounds 58-60.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| WARRIOR-R8-1 | CONFIRMED | race | arms.md:111; fury.md:114 | "Gnome's Eureka!, 10% more damage from your abilities for 8 seconds every 2 minutes" | The warrior row's client text (Wowhead Forever tooltip 1259813; `spelltips/shared.json` Eureka! has the same text) reads: "Your **non-periodic** damaging abilities have their Rage cost reduced by 10% and deal 10% more damage for 8 sec." The engine applies it per class row with a separate periodic modifier (`sim/core/eureka.go` at d1d9f2775). Rend ticks and Deep Wounds get nothing from it, and Arms keeps Rend up all fight. The rogue page (combat.md:112, a round 7 MUST) and the warlock pages already state the limit; the warrior pages do not. | Both pages: "Gnome's Eureka!, 10% more damage and 10% less rage on your non-periodic abilities for 8 seconds every 2 minutes," (the rest of the sentence is unchanged). |
| WARRIOR-R8-2 | CONFIRMED | card note | `warrior-protection.json` priorityList[8] (Thunder Clap) | "Keep it on the boss: its attack speed slow is your biggest cut to damage taken after Shield Block." | Thunder Clap 11581 is "increasing the time between their attacks by 20%", so it cuts melee damage by 1 - 1/1.2 = 16.7%. Demoralizing Shout at 60 removes 204 of a level 63 boss's roughly 270 attack power. A boss swing scales as weapon damage x (1 + 0.15873% x AP), so the shout cuts damage by 1 - 1.105/1.429 = 22.6%. The same file's notes record the engine's measurements on the tank harness: Thunder Clap upkeep "another 16 percent", Demoralizing Shout "damage taken about 22 percent lower". In Forever, Demoralizing Shout is the bigger cut. (In Classic, Thunder Clap's 10% slow made the note true.) | "Keep it on the boss: its attack speed slow is a straight cut to the damage you take. Refresh it in its last 3 seconds." |
| WARRIOR-R8-3 | LIKELY | stale-70755 | arms.md:60; fury.md:59; `warrior-arms.json` priorityList[3], `warrior-fury.json` priorityList[4] (Battle Shout) | "**Brings:** ... Battle Shout for your group" / card "It is 10 rage for the whole group's attack power." | The Forever shout covers the raid: 11551 and 25289 read "increasing the melee attack power of all **party or raid** members within 20 yards ... Lasts 3 min" (`spelltips/warrior.json`; Wowhead Forever tooltips 11551 and 25289). The Classic 11551 reads "all party members within 20 yards ... Lasts 2 min" (Wowhead Classic tooltip). In Classic, "for your group" meant one shout per party, which is how raids placed their warriors. In Forever one shout reaches every raid member in range. The same pages and the Protection note already treat it as one raid-wide aura ("If a raid warrior's shout is already on you, skip yours"). Filed LIKELY only because "group" can be read loosely. | arms.md:60: "**Brings:** Mortal Strike's 50% healing cut for the bosses that heal, and Battle Shout, which now reaches every party and raid member within 20 yards". fury.md:59: "**Brings:** Battle Shout, which now reaches every party and raid member within 20 yards". Both card notes: "Whenever it drops: 10 rage for the attack power of everyone within 20 yards, raid included." |

## Shoulds (wording, not findings)

- fury.md:64 "It is the button you never delay." The Execute paragraph (fury.md:90) then puts Execute ahead of Bloodthirst. The card note's wording, "the button you lose the most by delaying", says the same thing without the clash.
- index.md "What changed from 1.12" has no Battle Shout bullet. Party-only and 2 minutes became party or raid and 3 minutes, which is a real change: "**Battle Shout** reaches every party and raid member within 20 yards and lasts 3 minutes."
- arms.md:111 "sit just behind it": "it" points back to Night Elf in the previous paragraph. "sit just behind Night Elf" reads cleaner.

## Checked and right

- **Levels 1 to 19, all three pages**: every trainer level matches `trainables/warrior.json` (Charge and Rend 4, Thunder Clap 6, Hamstring 8, Bloodrage 10, Overpower and Shield Bash 12, Revenge, Demoralizing Shout and Tactical Mastery 14, Shield Block 16). Every point by level matches the orders above. The rank values hold: IHS 1, Deflection 1% a rank, Improved Rend 35%, Improved Charge +3 a rank, Cruelty 5%, Lingering Rage 2 s at rank 1, Unbridled Wrath 60% at 5, Improved Bloodrage 50% at 2, and Shield Spec's 100% chance at rank 5 only. The 1 to 19 and 20 to 60 blocks read in order, and neither contradicts the strip.
- **Talent text** re-read in the 70755 rank tables: Weaponmaster (5% crit, 15% armor, 5% extra attack), Bloodthrill main hand 20%, Blood Craze (crit or a hit over 20% of max health), Anger Management 1 every 3 s, Improved Slam, Impale (220%, round 7's fix holds), Furious Precision 10% off-hand hit, Reckless Swings, Improved Whirlwind, Flurry, Death Wish 20% / 5%, Enrage 4% at 2, Shield Spec, Master of Defense, Anticipation, Defiance, Bastion, Focused Rage, Improved Thunder Clap, Improved Shield Wall, Vanguard and Improved Shield Bash. The Improved Shield Wall swap is still legal: 25 points sit below tier 4 after it, and 31 below Bastion.
- **Client descriptions** (`spelltips/warrior.json`): Execute spends every extra point of rage; Overpower 5 s after a dodge; Whirlwind 4 targets with both weapons, Berserker Stance; Sweeping Strikes 5 attacks, Battle Stance; Recklessness 100% / 15 s / 20% / 30 min; Berserker Rage 10 s; Shield Block 75% / 7 s / 2 attacks / 5 s; Shield Wall 60% / 12 s / 15 min; Last Stand 30% / 20 s; Challenging Shout 6 s / 10 min; Victory Rush 10% / 20 s window / 30 s; Taunt "no effect if the target is already attacking you"; Defensive Stance +30% threat. The `class_spell` rows in trainables (Devastate 20243, Howling Blade, Recycle) come from SkillLineAbility-less data and are not learnable, so leaving Devastate out of the Protection page is right.
- **Demoralizing Shout 204 at 60 (protection.md:92) holds.** Wowhead Forever's tooltip shows 210, but that figure is computed at the rank's MaxLevel 64. SpellLevels for 11556 has BaseLevel 54, so a level 60 tank gets 196 + 6 x 1.4 = 204 (the same check gives 116 for 11551 at its MaxLevel 61). Round 4's decision stands.
- **Racials** (`spelltips/shared.json`): Elune's Light, Ruthless, Artisan's Precision, Blood Fury, Berserking 10%, Touch of the Grave (20%, up to 5% health), Wind Blessed 2%, Quickness, Endurance and Stoneform all match the prose. Race order matches `round2/race-sims.md` (guide builds) and `round3/race-sims.md` (Protection). Skyborne is named only as the paid option.
- **Consumables**: Elixir of the Mongoose (25 Agility, 2% crit), R.O.I.D.S. (25 Strength), Winterfall Firewater (35 AP), Smoked Desert Dumplings (20 Strength, Forever spell 1248401) and Mighty Rage Potion all exist with those effects.
- **Card notes** other than R8-2 and R8-3: each matches its action, the APL thresholds and the client text.
- **Source links**, all fetched this round: Blizzard recap, Wowhead Skyborne First Look, Power Up Gaming, Talents Forever data.json, Warcraft Tavern Classic PvE DPS warrior, Warcraft Tavern Forever warrior and Fury guides, and foreversixty.gg /sim/weights, /sim/drops, /planner and /bis/warrior/protection all return 200. Icy Veins racial rework returns 403, so its status is unknown, not dead. The source-link rule holds.
- No em or en dashes and no quoted nightly numbers. Stat prose explains and never ranks. No gear is ranked by a nightly value.

## Decisions

- **Eureka! names its limit (R8-1).** The client text says non-periodic, and the rogue and warlock pages already say so.
- **Thunder Clap is not the biggest cut after Shield Block (R8-2).** Forever doubled the slow to 20%, but Demoralizing Shout's 204 against a level 63 boss's roughly 270 AP is still bigger, both by Classic's boss-swing formula and by the engine's own recorded measurement. The prose order (Thunder Clap, then Demoralizing Shout) is the APL's priority, not a size claim, so it stays.
- **Battle Shout is raid-wide in Forever (R8-3).** Settled by the 70755 client text against the Classic text.

## Sims requested

None new. Still open from earlier rounds: Dwarf vs Human breakdown (Arms and Fury), Elune's Light and Blood Fury held-timing comparisons, the Stoneform rotation line, S2 Improved Berserker Rage, S3 Spearing Strike, S4 third Improved Heroic Strike point, the Protection Improved Shield Wall build and Shield Wall line, and the `expertise` weight removal.

## Sim findings

None new. The engine's boss swing scales with creature attack power as a share, with 268 at level 63. That matches the Classic formula the Blizzard forum analysis gives, so R8-2 is a card-note error only.

## Sources

- Blizzard forums, Boss Attacks and You: Beware Math Inside, https://us.forums.blizzard.com/en/wow/t/boss-attacks-and-you-beware-math-inside/459866 (boss auto attack = weapon damage x (1 + 0.15873% x AP); a level 63 boss has about 270 AP).
- Wowhead Forever tooltips, https://nether.wowhead.com/forever/tooltip/spell/<id>: Eureka! (Warrior) 1259813, Battle Shout 11551 and 25289, Demoralizing Shout 11556, Thunder Clap 11581, Whirlwind 1680, Sweeping Strikes 12292, Execute 20662.
- Wowhead Classic tooltips: Battle Shout 11551, https://nether.wowhead.com/classic/tooltip/spell/11551 ("all party members ... Lasts 2 min"); Shield Block 2565, https://nether.wowhead.com/classic/tooltip/spell/2565.
- Client data 1.60.1.70755: `spelltips/warrior.json`, `spelltips/shared.json` (racials, consumables), `spellconst/warrior.json` (11551, 11556), `talents/warrior.json`, `trainables/warrior.json`, `simconsumes.json`, `addon-builds.json`; SpellLevels (11556 BaseLevel 54, MaxLevel 64).
- Engine d1d9f2775: `sim/core/eureka.go`, `sim/warrior/demoralizing_shout.go`.

## Fix log (round 8)

- WARRIOR-R8-1 fixed, arms.md:111 and fury.md:114: "10% more damage and 10% less rage on your non-periodic abilities". Checked against the engine's warrior class mask (eureka_auto_gen.go 1259813): no Rend or Deep Wounds, no periodic modifier.
- WARRIOR-R8-2 fixed, warrior-protection.json Thunder Clap notes only: "its attack speed slow is a straight cut to the damage you take. Refresh it in its last 3 seconds." Action untouched.
- WARRIOR-R8-3 fixed after the client check (spelltips/warrior.json 6673 to 25289: "all party or raid members within 20 yards ... Lasts 3 min"): arms.md:60, fury.md:59, and the Battle Shout notes in warrior-arms.json and warrior-fury.json ("Whenever it drops: 10 rage for the attack power of everyone within 20 yards, raid included.").
- SHOULDs (fury "never delay", index Battle Shout bullet, "behind it") left: wording only, outside the assigned findings.
- Prose check: warrior clean. No new SIM findings.
