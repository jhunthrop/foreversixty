# Round 14 wow-player review: physical, tanks, hybrids, warlock (2026-11-04)

Reviewer: veteran Classic raider. Forever client rows beat Era memory where the reports cite them.

## Verdicts per row

| row | verdict | why a player cares |
|---|---|---|
| tanks O1 Templar's Bulwark at paladin band 20 | SHIP | A level 20 protection paladin has no avoidance to speak of; a whole-health-bar absorb every 5 minutes is the thing that keeps him alive. Owner exception stands. FIX (optional): add to the guide's Leveling section "Take Templar's Bulwark as soon as it opens, ahead of a second point of Sacred Duty." Otherwise the guide's "Sacred Duty (2 points)" line disagrees with what the addon sets at 20. |
| tanks O2 boss AP 268, swing x (1 + AP/630) | SHIP | Matches how Classic tanks remember demo shout: a double-digit cut on boss melee (the old "about 16%" with the improved 146). 204 giving 22.7% is bigger than Era only because Forever's rank is bigger. Just do not publish 22.7% in prose (round 6 rule); the guide says "a real slice", which is fine. |
| tanks O3 raid debuff options at the client's 204 | SHIP | No published number moves. |
| tanks O4 Demoralizing Shout / Roar in the rotation | SHIP, owner confirms | See the answer below. |
| tanks B1 Frenzied Regeneration 1% a rage | SHIP | Era's 0.3% was the old value; Forever's own tooltip says 1%. Guide text "heal you for your whole health bar given the rage" is already consistent. |
| tanks B2 Roar threat | SHIP | Small. |
| tanks B5 shout/roar can miss | SHIP as OPEN | Classic players do see "resisted" on Demo Shout against bosses; keep the miss roll. |
| physical R1 Premeditation without Stealth | SHIP, one check | Shapeshift mask 0 plus Wowhead Forever having no "Requires Stealth" line is two independent sources, and Forever changing it is plausible. Check: the guide says the points "last 20 seconds". Era's tooltip says 10 seconds; confirm the combo-point expiry from the Forever aura row before that number stays in subtlety.md (lines 69 and the APL note). If it is 10, change both to "10 seconds" and the "three combo points or fewer" rule is unchanged. Also say in the guide "if it is off cooldown while Hemorrhage or Rupture is mid-run, press it anyway" only if the sim does so; otherwise leave. |
| regmelee H1 BM Arcane Shot downrank rows removed | SHIP | Nobody wants to cast rank 1 Arcane Shot at 50 when rank 8 hits for 280. Players will not miss it. |
| regmelee S1 Aspect of the Falcon | SHIP | Survival.md wording "at 60 it gives the same melee attack power as Beast (10 more once you have Hawk's Ahn'Qiraj rank)" is correct against rank 4 Beast 110 and rank 7 Hawk 120. It is a long sentence; split at "but Deadly Aspects". |
| hybrids W1 to W5 cat weapon DPS | SHIP | Forever's druid deep dive beats the cmangos Era rule. feral.md lines 42, 92 and 100 now agree with the engine. A "high-DPS two-hander beats a caster staff" line is what every cat will test in the first hour; if it is wrong the forums find it fast, so keep the in-game check on the open list. |
| core C1 hits inside the same 10 ms batch | SHIP | Invisible to players; fixes a Discipline dip at -0.01% haste. |
| core C5 Searing Totem full 55 s | SHIP | Matches a shaman dropping a new one at expiry. Mention in release notes. |
| warlock A1 Immolate cast speed | SHIP | Berserking and haste not speeding Immolate would have been mocked. |
| warlock C1 Orc over Troll for Affliction and Destruction | SHIP with FIX | Every Classic warlock expects Orc (Blood Fury, Command). But demonology.md still says "Gnome and Troll are the picks" while Orc beats Troll there by +0.8% and Undead by +0.4 to 0.6%. Players will ask why Orc is right for two specs and not the third. Owner call below. Do not write "Orc over Troll" sim talk into the guides; the race paragraphs already avoid it. |
| warlock D4 Unholy Power 5 over Improved Imp 3 | SHIP | Leveling warlocks run Succubus or Voidwalker; Improved Imp buffs a pet nobody has out. demonology.md line 53 says exactly that. Raid build already takes Improved Imp 0. |
| data R1 Gnome Expansive Mind rage/energy | SHIP | A gnome warrior or rogue at 105 is the client's number; the effect is small and rogue-guide race paragraphs need no change. If any guide says "Expansive Mind adds mana only", fix it. |
| data K1 to K3 rank rows | SHIP | Dropping rows that tie or lose is right; a player reads "cast the top rank" as the default. |

## Would a Classic tank keep Demoralizing Shout / Roar up?

Yes, with two caveats.
- In dungeons and on trash, a tank keeps Demo Shout up without thinking about it. On a raid boss a warrior tank in Era usually let a druid or off-tank hold it, because rage (10 per cast) competes with Shield Slam, Revenge and Sunder and the threat gain per rage is poor. -2.5% threat on a warrior and -5.4% on a bear for roughly a fifth less damage taken is the trade a healer would take every time, and it matches the guides' own wording. The bear is the stronger case: 98% to 10% chance of death is not a rotation tweak, it is the difference between a bear that can tank and one that cannot.
- The 5.4% bear threat cost is noticeable to a DPS who is close to pulling. Keep the guide line "unless a druid or another warrior already holds it: two do not stack" and put the same sentence in the APL line notes. They already carry it.
- Keep it ranked below Shield Block/Revenge/Shield Slam and above Sunder, as adopted. Do not move the roar above Primal Bite (7.5% death but -6.8% threat); correctly not adopted.

## Owner calls I recommend

1. Confirm O4 (adopt). I would adopt.
2. Race for Demonology: Orc across all three warlock specs for one message, even though Demonology is under 1%. Recommended: yes, list [gnome, orc], because a split recommendation reads as a data bug to a player.
3. Premeditation combo-point lifetime (10 vs 20 s): verify from the aura row before shipping the number.
4. Add the one Templar's Bulwark sentence to paladin/protection.md Leveling (see O1).

## Release-note lines

- Protection warriors and bears now keep Demoralizing Shout and Roar on the boss; damage taken is about a fifth lower at 60.
- Frenzied Regeneration now heals 1% of your health per rage, as the Forever client says.
- Bosses' attack power now follows Classic: Demoralizing effects matter again.
- Protection paladins take Templar's Bulwark at 20.
- Subtlety rogues press Premeditation again in combat, every 2 minutes; it needs no Stealth in Forever.
- Cat Form attacks use your weapon's DPS; weapon choice matters again for cats.
- Beast Mastery casts its top Arcane Shot rank at every level.
- Gnome warriors and rogues get Expansive Mind's +5% rage or energy.
- Warlocks: Immolate now speeds up with haste and Berserking; Orc is the pick for Affliction and Destruction; leveling Demonology takes Unholy Power.
- Searing Totem now stands its full 55 seconds before it is dropped again.
