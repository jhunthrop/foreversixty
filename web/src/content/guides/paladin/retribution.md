---
title: Retribution Paladin in Forever
classSlug: paladin
spec: retribution
role: dps
build: 'FS1:1.60.1.70938:paladin:human:050250002/502/0523330100130135:'
buildWhy: 'Champion of the Light by 40, and Reverence for questing mana'
raidBuild: 'FS1:1.60.1.70938:paladin:human:050050002/503/05233301001301351:'
raidBuildWhy: 'Trades Reverence for Precision and the Unsealed Wrath twist'
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
    url: https://foreversixty.gg/data/1.60.1.70938/races.json
    kind: datamined
---

## Overview

### Strengths

- The fastest way to level a paladin, with Sanctified Judgement paying for most of your Seals
- Almost every button is instant, so you keep swinging through movement
- Hammer of Wrath is an instant execute on a 6 second cooldown with Champion of the Light
- Unsealed Wrath turns Seal swapping into real damage, as long as your mana keeps pace with the fight

### Weaknesses

- Mana decides how much Consecration and Seal twisting you get in a long fight
- Exorcism and Holy Wrath only work on Undead and Demons
- No baseline damage cooldown: your burst is the execute

### In a raid

- **Role:** Melee damage, and one of the paladins holding a Judgement on the boss
- **Brings:** Greater Blessings (Might or Wisdom, and Kings where no hunter's Heart of the Lion covers the raid), an aura for the melee group, and Judgement of the Crusader or Wisdom on the boss

## Talents and builds

The leveling build is 14 Holy, 7 Protection and 30 Retribution. The raid build moves Reverence's two points into the third rank of Precision and Unsealed Wrath.

- **Benediction** (5 points): your Seals, Judgement, Holy Strike, Consecration and Exorcism cost 10% less. Not Hammer of Wrath: Champion of the Light makes it instant, but Benediction still does not discount it.
- **Improved Judgement** (2 points): Judgement every 8 seconds instead of 10.
- **Conviction** (3 points): 6% crit on everything you do: your swings, Holy Strike, Seal procs, Judgements and Consecration's ticks. Exorcism never crits in Forever.
- **Sanctified Judgement** (3 points): every Judgement hands back 60% of the judged Seal's cost. This is what pays for your Seals, and for the twist at 60.
- **Sacred Arbiter** (1 point): Holy Strike hits 20% harder and refreshes every Judgement on the target.
- **Two-Handed Weapon Specialization** (3 points): 6% more Physical and Holy damage with a two-hander, and 20% less threat. That reaches your Holy Seal procs and Judgements as well as your swings.
- **Vengeance** (3 points): 3% more Physical and Holy damage for 30 seconds after a crit from anything but Consecration's ticks, stacking five times to 15%. With Conviction you reach five stacks early and hold them for the whole fight.
- **Champion of the Light** (5 points): 60% of your Intellect as spell damage, and Hammer of Wrath becomes instant.
- **Unsealed Wrath** (1 point): the raid capstone. Each swing stores a copy of your Seal, up to three, and a swing under a different Seal releases them.

**Seal of Command** is a required point on the way to Vengeance. The leveling build never casts it: without Unsealed Wrath, Seal of Righteousness does more on its own. With Unsealed Wrath, it becomes your main Seal for the twist. **Vindication** is filler on the way to Champion of the Light: once Seal of Command has its point, the rest of that row is Unyielding Faith and Pursuit of Justice, and neither adds damage. **Repentance** is the point Unsealed Wrath requires, and it still earns its place in a dungeon: a 6 second incapacitate on a Humanoid, on a one minute cooldown.

In Holy, **Divine Strength** is 10% Strength, **Improved Seals** is 15% on every Seal and Judgement, and **Holy Conduit** takes 40% off Consecration, Exorcism and Hammer of Wrath. **Reverence** keeps 30% of your mana regeneration running while you cast, which matters more while leveling than in a raid with Blessing of Wisdom on you. **Divine Intellect** opens Protection and leads to **Precision**.

Leveling order: Benediction and Improved Judgement first, then Conviction and Sanctified Judgement, then down the tree to Champion of the Light by level 40.

## Rotation and priority

**Before the pull.** Put a Blessing of Might on yourself if nobody has given you one. If no other paladin has Judgement of the Crusader on the target, put up Seal of the Crusader about a second and a half before the pull and Judge it on your first swing. If the raid already has Crusader on the boss, open straight with your damage Seal and Judge on your first swing.

**Seals while leveling.** Stay in Seal of Righteousness at the top rank you have trained. Train rank 8 at 58: every rank takes the same share of your spell damage, so the newest rank always hits hardest. Do not twist the way 1.12 guides tell you. In Forever a new Seal replaces the old one, no swing gets both, and every swap costs a global and a Seal's worth of mana. Spend those globals on Holy Strike instead.

**The Seal twist.** With Unsealed Wrath, three swings under Seal of Command store three Commands, and one swing under Seal of Righteousness releases them, each with its normal chance to fire. Windfury Totem's extra swings store nothing, so count your own swings, not the totem's. Switch back to Seal of Command the moment that swing lands, and the next swing releases the stored Righteousness hit. Keep pace at four fifths of the fight to come: 80% mana at the pull, 40% halfway and 20% with a quarter left. A swing timer and a stack counter on Unsealed Wrath make this easy.

**Judgements in a raid.** Your own Judgement is for damage. The debuff Judgements are assigned: one paladin keeps Judgement of the Crusader on the boss and another keeps Judgement of Wisdom. If that is you, open with that Seal, Judge it, then go to your damage Seal. The debuff lasts 40 seconds and Holy Strike refreshes it, so after the first one it costs you nothing. Judgement is a ranged attack: it lands on your ranged hit chance, nothing dodges, parries or blocks it, and your weapon's enchant, Windfury Totem and a Hand of Justice never fire from it. Only your swings and Holy Strike can fire those.

**Consecration.** On one target it is filler: running dry for your Seal costs far more than a skipped one. Your regen pays for it almost every time it is ready with Blessing of Wisdom, Judgement of Wisdom and Mana Spring; on your own, hold it for the mana your Seal and Judgements need. On two or more targets, cast it freely.

On one target a cheap Consecration rank you can put down far more often beats the higher ranks:

- **To 52:** rank 1, even after rank 4 arrives at 50.
- **53 to 57:** rank 3. The second point of Holy Conduit at 53 takes 40% off its cost.
- **58 and 59:** rank 4 for a Human or Dwarf. An Undead paladin keeps rank 3.
- **60:** rank 4 without raid buffs, rank 5, the top rank, with them.

**The execute.** Save a Major Mana Potion for it if you are low. Hammer of Wrath at your top rank at every level.

**Movement and adds.** Hammer of Wrath works at range and Judgement from 10 yards, so keep pressing them while you close on the boss. On a pack, Consecration first, then the kill target. Against a pack of Undead or Demons, Holy Wrath hits all of them; on anything else it does nothing, so leave it off your bars outside Scholomance, Stratholme and the like.

## Stat priority

Divine Strength adds 10% to your Strength, which goes into every swing, Holy Strike and Seal of Command proc. Seal of Righteousness scales with your weapon's speed and your spell damage instead, and a two-hander takes about 60% more of that spell damage than a one-hander. Champion of the Light turns 60% of your Intellect into spell damage for your Seals, Judgements, Consecration and Holy Strike. Holy Strike adds its 81 to 105 Holy damage and a share of your spell damage on top of half your weapon's damage, and neither is cut by that half, so Intellect and spell power both add to your damage. A melee crit doubles the hit and feeds Vengeance. Hit counts up to the cap and adds nothing past it; Precision and Artisan's Precision both count toward it.

## Best in slot

## Enchants and consumables

## Races

**Human** for the Alliance. Ruthless adds crit to everything, The Human Spirit keeps your mana up while leveling and in dungeons, and Will to Survive breaks a stun. **Dwarf** matches Human on the raid build in a raid but falls behind solo, with Stoneform for a hard pull. **Undead** is new to the class in Forever and the Horde's only paladin. Will of the Forsaken no longer makes you immune to the next fear, so press it after the fear lands, and Cannibalize restores mana as well as health, which cuts your drinking while leveling.

## Professions

Mining and Blacksmithing make your two-handers and plate. Engineering adds explosives for trash packs, and it is the profession that adds damage when your gear is done. Herbalism and Alchemy keep you in mana potions and Elixirs of Holy Power, Forever's 40 Holy damage elixir, which reaches every Seal proc, Judgement and Consecration tick. Your best gear asks for Enchanting: Frozen Heart of the Mountain binds when picked up, so only an enchanter can wear one. Otherwise pick two that pay for your consumables.

## Leveling

### Levels 1 to 19

Seal of Righteousness and Holy Light come at 1, Judgement and Blessing of Might at 4, Holy Strike at 6, Hammer of Justice at 8 and Lay on Hands at 10. Judgement and Holy Strike both come back every 10 seconds.

At 12, the class quest from your trainer teaches Redemption. Do it before your first dungeon.

Your first five points, 10 to 14, go into Benediction: 10% less mana on every instant, which is every button you fight with. Improved Judgement takes 15 and 16 for a Judgement every 8 seconds, then Conviction from 17 to 19 for 6% crit.

Take quest rewards over vendor gear, and replace your Battleworn Hammer with the first harder-hitting slow two-hander you find: Seal of Righteousness hits harder per swing the slower the weapon. Paladins cannot use wands. You wear mail until plate at 40.

### Levels 20 to 60

Retribution is the paladin leveling spec, even if you want to heal or tank at 60. Sanctified Judgement refunds most of each Seal, so on one or two mobs at a time you rarely drink. The talents after Conviction:

- **23:** one point of Vindication opens its row.
- **24 to 28:** Seal of Command at 24, Sacred Arbiter at 25, Two-Handed Weapon Specialization at 26 to 28.
- **29 and 33:** the last two points of Vindication, which only lowers the attack power of what you hit.
- **30 to 32:** Vengeance.
- **34:** Repentance, to open Champion of the Light.

From level 20 Consecration is your AoE: two or three mobs on you is the time for it. Rank 1, about 120 mana after Benediction, keeps almost all of its spell damage through your 30s for a fraction of the cost, so until your mid 40s it gives you the most damage for your mana. From 40, when a pack is on you and you have the mana, cast the newest rank instead: it hits every mob about twice as hard for the same global. From your mid to late 40s the newest rank also wins on mana on a pack, because rank 1 has lost about half of its spell damage bonus by then.

Against one target, keep Consecration rank 1 all the way to 52. At 53 the second point of Holy Conduit takes 40% off Consecration's cost, and rank 3 becomes the best Consecration you have: it hits far harder than rank 1 and still costs less than your newest rank. A Human or Dwarf moves to rank 4 at 58; an Undead paladin keeps rank 3 to 59.

Exorcism (level 20) and Holy Wrath (level 50) are easy kills on Undead and Demons, so quest zones full of them, like the Plaguelands, are paladin country. From 31, Repentance takes one Humanoid out of a bad pull, and Hammer of Justice stops a runner or a caster.

Common mistakes:

- Twisting Seals without Unsealed Wrath, or while you are behind on mana.
- Letting the Seal drop. Recast it in its last second or two, not after it is gone.
- Spending the bar on Consecration on a single target and having nothing left for Hammer of Wrath.
- Swinging a one-hander. Two-Handed Weapon Specialization only works with a two-hander, and every Seal proc scales with the weapon.
- Getting hit wrong at 60: gear to the cap, then stop.
