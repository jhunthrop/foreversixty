---
title: Retribution Paladin in Forever
classSlug: paladin
spec: retribution
role: dps
build: 'FS1:1.60.1.70755:paladin:human:050250002/502/0523330100130135:'
raidBuild: 'FS1:1.60.1.70755:paladin:human:050050002/503/05233301001301351:'
recommendedRaces: [human, undead]
statPriority: [Hit, Intellect, Strength, Critical strike, Agility, Spell power, Attack power, Melee haste]
description: 'Leveling and raiding with Seals, Judgement and an instant Hammer of Wrath, including the raid Unsealed Wrath swap.'
updated: 2026-11-04
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Forever Sixty, paladin talents'
    url: https://foreversixty.gg/planner?class=paladin
    kind: datamined
  - label: 'Forever Sixty, what each stat is worth'
    url: https://foreversixty.gg/sim/weights
    kind: site
  - label: 'Warcraft Tavern, Classic Retribution Paladin rotation'
    url: https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/
    kind: community
  - label: "Blizzard, What's Next panel recap"
    url: https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap
    kind: blizzard
  - label: 'Forever Sixty, races'
    url: https://foreversixty.gg/data/1.60.1.70755/races.json
    kind: datamined
---

## Overview

Retribution is a two-handed Seal fighter: keep a Seal up, Judge every 8 seconds, Holy Strike on cooldown, and finish with an instant Hammer of Wrath.

### Strengths

- The fastest way to level a paladin, with Sanctified Judgement paying for most of your Seals
- Almost every button is instant, so you keep swinging through movement
- Hammer of Wrath is an instant execute on a 6 second cooldown with Champion of the Light
- Unsealed Wrath turns Seal swapping into real damage once you have raid buffs

### Weaknesses

- Mana decides how much Consecration you get in a long fight
- Exorcism and Holy Wrath only work on Undead and Demons
- No baseline damage cooldown: your burst is the execute

### In a raid

- **Role:** Melee damage, and one of the paladins holding a Judgement on the boss
- **Brings:** Greater Blessings (Might or Wisdom, and Kings where no hunter's Heart of the Lion covers the raid), an aura for the melee group, and Judgement of the Crusader or Wisdom on the boss

## Talents and builds

The leveling build is 14 Holy, 7 Protection and 30 Retribution. The raid build moves Reverence's two points into the third rank of Precision and Unsealed Wrath.

- **Benediction** (5 points): your Seals, Judgement, Holy Strike, Consecration and Exorcism cost 10% less. Hammer of Wrath is the exception: Champion of the Light makes it instant, but it still costs full price.
- **Improved Judgement** (2 points): Judgement every 8 seconds instead of 10.
- **Conviction** (3 points): 6% crit on everything you do: your swings, Holy Strike, Seal procs, Judgements and Consecration's ticks. Exorcism is the exception: it never crits in Forever.
- **Sanctified Judgement** (3 points): every Judgement hands back 60% of the judged Seal's cost. This is what pays for your Seals, and for the twist at 60.
- **Sacred Arbiter** (1 point): Holy Strike hits 20% harder and refreshes every Judgement on the target.
- **Two-Handed Weapon Specialization** (3 points): 6% more Physical and Holy damage with a two-hander, and 20% less threat. That reaches your Holy Seal procs and Judgements as well as your swings.
- **Vengeance** (3 points): 3% more Physical and Holy damage for 30 seconds after a crit from anything but Consecration's ticks, stacking five times to 15%. With Conviction you reach five stacks early and hold them for the whole fight.
- **Champion of the Light** (5 points): 60% of your Intellect as spell damage, and Hammer of Wrath becomes instant.
- **Unsealed Wrath** (1 point): the raid capstone. Each swing stores a copy of your Seal, up to three, and a swing under a different Seal releases them.

**Seal of Command** is a required point on the way to Vengeance. Below 60 you never cast it: Seal of Righteousness does more with leveling gear. **Vindication** is filler on the way to Champion of the Light: once Seal of Command has its point, the rest of that row is Unyielding Faith and Pursuit of Justice, and neither adds damage. **Repentance** is the point Unsealed Wrath requires, and it still earns its place in a dungeon: a 6 second incapacitate on a Humanoid, on a one minute cooldown.

In Holy, **Divine Strength** is 10% Strength, **Improved Seals** is 15% on every Seal and Judgement, and **Holy Conduit** takes 40% off Consecration, Exorcism and Hammer of Wrath. **Reverence** keeps 30% of your mana regeneration running while you cast, which matters more while leveling than in a raid with Blessing of Wisdom on you. **Divine Intellect** opens Protection and leads to **Precision**. At 60, hit is worth taking up to the cap and adds nothing past it.

Leveling order: Benediction and Improved Judgement first, then Conviction and Sanctified Judgement, then down the tree to Champion of the Light by level 40.

## Rotation and priority

**The opener.** Put a Blessing of Might on yourself if nobody has given you one. If no other paladin has Judgement of the Crusader on the target, put up Seal of the Crusader about a second and a half before the pull, Judge it on your first swing, then put up your damage Seal: Seal of Righteousness while leveling, Seal of Command with Unsealed Wrath and raid buffs. Your swings and Holy Strike keep the Crusader debuff up, so that one Judgement adds Holy damage to every Seal proc, Judgement, Holy Strike and Consecration tick for the rest of the fight. If the raid already has Crusader on the boss, open straight with your damage Seal and Judge on your first swing. Judgement never consumes your Seal in Forever, so there is nothing to put back up after it.

**The priority**, top to bottom, every global:

1. **Judgement** on cooldown. It is off the global cooldown, so press it the moment it is ready.
2. **Exorcism** on cooldown, against Undead and Demons only.
3. **Hammer of Wrath** when the target is below 20%, on its 6 second cooldown.
4. Your **Seal** if it is about to run out.
5. **Holy Strike** on cooldown.
6. **Consecration** with whatever mana is spare.

**Seals while leveling.** After the opening Crusader Judgement, stay in Seal of Righteousness. It fires on every swing, it scales with spell damage, and its Judgement hits hard. Do not twist the way 1.12 guides tell you: in Forever a new Seal simply replaces the old one, no swing gets both, and every swap costs a global and a Seal's worth of mana. Spend those globals on Holy Strike instead.

**The raid twist (Unsealed Wrath).** With raid buffs on you, Seal of Command becomes your main Seal. Let three swings land under Seal of Command, so Unsealed Wrath shows three stacks. Switch to Seal of Righteousness for exactly one swing: that swing releases all three stored Commands, each with its normal chance to fire. As soon as it lands (Unsealed Wrath's timer jumps back to 15 seconds), switch back to Seal of Command, and the next swing releases the stored Righteousness hit. Repeat. Two rules keep it from costing you: only twist above 30% mana, and only with raid buffs, because with little attack power Seal of Command does less than Seal of Righteousness and the swaps just burn mana. A swing timer and a stack counter on Unsealed Wrath make this easy.

**Judgements in a raid.** Your own Judgement is for damage. The debuff Judgements are assigned: one paladin keeps Judgement of the Crusader on the boss and another keeps Judgement of Wisdom. If that is you, open with that Seal, Judge it, then go to your damage Seal. The debuff lasts 40 seconds and Holy Strike with Sacred Arbiter refreshes every Judgement you have on the target, so after the first one it costs you nothing. On your own, know what each Seal gives you. Seal of the Crusader is attack power and faster, lighter swings, and its Judgement makes the target take more Holy damage for 40 seconds, with your melee keeping it up. Seal of Righteousness adds Holy damage to every swing, and its Judgement is a Holy hit. Judgement no longer consumes your Seal, so judging Crusader once and then putting up Righteousness costs you one extra Seal and one Judgement, not a Seal every time. On your own, always open with Crusader as above: one Seal and one Judgement cover the whole fight.

**Consecration.** It is your most expensive button, so pace it. Cast it only while the mana left after it still pays for a Seal, a Judgement and a Holy Strike, and while your mana percent is at least four fifths of the percent of the fight left (halfway through, 40% mana or more). Fall under that and Consecration waits until you catch up. Spam it and your Seal drops and you go into the execute with nothing for Hammer of Wrath. Hold it too long and you end the fight with a full bar. Below 60, rank 1 does the most over a fight: it costs 135 mana, and the mana it saves goes into Seals, Judgements and Holy Strike. At 60, use rank 4 without raid buffs and rank 5, the top rank, with them. On two or more targets, cast it freely.

**The execute.** Below 20% Hammer of Wrath goes ahead of everything but Judgement and Exorcism. It is instant with Champion of the Light, so you never stop swinging to cast it. Save a Major Mana Potion for this phase if you are low.

**Movement and adds.** Hammer of Wrath works at range and Judgement from 10 yards, so keep pressing them while you close on the boss. On a pack, Consecration first, then the single-target priority on the kill target. Against a pack of Undead or Demons, Holy Wrath hits all of them; on anything else it does nothing, so leave it off your bars outside Scholomance, Stratholme and the like.

## Stat priority

Strength is attack power on every swing, every Holy Strike and every Seal of Command proc, and Divine Strength raises it by 10%. Seal of Righteousness works differently: it scales with your weapon's speed and your spell damage, not attack power. Crit counts twice: a melee crit doubles the hit, and it feeds Vengeance. Champion of the Light turns 60% of your Intellect into spell damage for your Seals, Judgements and Consecration, so spell power and Intellect both add to your damage, and Intellect is also your mana bar. Agility is crit and armor at once. Hit matters up to the cap and nothing past it. Below the cap, in leveling or fresh 60 gear, every miss is a lost swing and a lost Seal proc, so take hit until you get there; Precision and Artisan's Precision both count toward it. The table above has the order for your gear.

## Best in slot

## Enchants and consumables

## Races

**Human** is the Alliance pick. Ruthless is 1% crit on everything, The Human Spirit's 10% Spirit keeps your mana up while leveling and in dungeons, and Will to Survive breaks a stun. **Dwarf** matches Human on the raid build in a raid, and Human is ahead solo, which is why Human stays the pick. Dwarf's Big Game Hunter adds 5% to the damage you deal to Beasts. Stoneform takes 10% off physical damage for 8 seconds on a hard pull. **Undead** is new to the class in Forever and the Horde's only paladin. Touch of the Grave can drain a little health from the target, at most once every 15 seconds. Will of the Forsaken breaks a fear, but no longer makes you immune to the next one, so use it after the fear lands. Cannibalize restores mana as well as health, which cuts your drinking while leveling.

## Professions

Mining and Blacksmithing make your two-handers and plate. Engineering adds explosives for trash packs, and it is the profession that adds damage when your gear is done. Herbalism and Alchemy keep you in mana potions and Elixirs of Holy Power, Forever's 40 Holy damage elixir, which reaches every Seal proc, Judgement and Consecration tick. Pick two that pay for your consumables.

## Leveling

### Levels 1 to 19

Seal of Righteousness and Holy Light come at 1, Judgement and Blessing of Might at 4, Holy Strike at 6, Hammer of Justice at 8 and Lay on Hands at 10. Judgement does not use up your Seal, and it and Holy Strike both come back every 10 seconds.

At 12, the class quest from your trainer teaches Redemption. Do it before your first dungeon.

Your first five points, 10 to 14, go into Benediction: 10% less mana on every instant, which is every button you fight with. Improved Judgement takes 15 and 16 for a Judgement every 8 seconds, then Conviction from 17 to 19 for 6% crit.

Take quest rewards over vendor gear, and replace your Battleworn Hammer with the first harder-hitting slow two-hander you find: Seal of Righteousness hits harder per swing the slower the weapon. Paladins cannot use wands. You wear mail until plate at 40.

### Levels 20 to 60

Retribution is the paladin leveling spec, even if you want to heal or tank at 60. Pull with Seal of the Crusader and a Judgement on your first swing, then Seal of Righteousness and Holy Strike on cooldown. Sanctified Judgement refunds most of each Seal, so on one or two mobs at a time you rarely drink.

From level 20 Consecration is your AoE: two or three mobs on you is the time for it, and rank 1 at 135 mana is the rank to keep on your bar until 60: a rank loses its spell damage bonus once you are far above the level you learned it, but rank 1 has so little to lose that its low cost wins. Exorcism (level 20) and Holy Wrath (level 50) are easy kills on Undead and Demons, so quest zones full of them, like the Plaguelands, are paladin country. Repentance takes one Humanoid out of a bad pull, and Hammer of Justice stops a runner or a caster.

Common mistakes:

- Twisting Seals without Unsealed Wrath. It costs mana and globals for nothing.
- Letting the Seal drop. Recast it in its last second or two, not after it is gone.
- Spending the bar on Consecration on a single target and having nothing left for Hammer of Wrath.
- Swinging a one-hander. Two-Handed Weapon Specialization only works with a two-hander, and every Seal proc scales with the weapon.
- Getting hit wrong at 60. Below the cap every miss costs a swing and a Seal proc, so gear for it; once you reach the cap, stop taking more, because past it hit adds nothing.
