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
updated: 2026-11-02
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Forever Sixty build data, data/builds/1.60.1.70755/talents/paladin.json'
    url: https://foreversixty.gg/data/1.60.1.70755/talents/paladin.json
    kind: datamined
  - label: 'Forever Sixty stat weights, data/curated/specs.json'
    url: https://foreversixty.gg/data/curated/specs.json
    kind: site
  - label: 'Warcraft Tavern, Classic Retribution Paladin rotation'
    url: https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/
    kind: community
  - label: 'Forever Sixty raid phase data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/data/curated/loot/forever-raid-phases.json
    kind: site
  - label: 'Forever Sixty build data, data/builds/1.60.1.70755/races.json'
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

- **Benediction** (5 points): every instant costs 10% less, and nearly everything you press is instant.
- **Improved Judgement** (2 points): Judgement every 8 seconds instead of 10.
- **Conviction** (3 points): 6% crit on everything you do: your swings, Holy Strike, Seal procs, Judgements and Consecration's ticks. Exorcism is the exception: it never crits in Forever.
- **Sanctified Judgement** (3 points): every Judgement hands back 60% of the judged Seal's cost. This is what pays for your Seals, and for the twist at 60.
- **Sacred Arbiter** (1 point): Holy Strike hits 20% harder and refreshes every Judgement on the target.
- **Two-Handed Weapon Specialization** (3 points): 6% more Physical and Holy damage with a two-hander, and 20% less threat. That reaches your Holy Seal procs and Judgements as well as your swings.
- **Vengeance** (3 points): 3% more Physical and Holy damage for 30 seconds after a crit, stacking five times to 15%. With Conviction you reach five stacks early and hold them for the whole fight.
- **Champion of the Light** (5 points): 60% of your Intellect as spell damage, and Hammer of Wrath becomes instant.
- **Unsealed Wrath** (1 point): the raid capstone. Each swing stores a copy of your Seal, up to three, and a swing under a different Seal releases them.

**Seal of Command** is a required point on the way to Vengeance. Below 60 you never cast it: Seal of Righteousness does more with leveling gear. **Vindication** and one point of **Repentance** are filler on the way to Champion of the Light, since nothing else in the tree adds damage there. Repentance still earns its point in a dungeon: a 6 second incapacitate on a Humanoid, on a one minute cooldown.

In Holy, **Divine Strength** is 10% Strength, **Improved Seals** is 15% on every Seal and Judgement, and **Holy Conduit** takes 40% off Consecration, Exorcism and Hammer of Wrath. **Reverence** keeps 30% of your regeneration running in combat, which matters more while leveling than in a raid with Blessing of Wisdom on you. **Divine Intellect** opens Protection and leads to **Precision**. With raid buffs at 60, hit is your best stat per point until you reach the cap.

Leveling order: Benediction and Improved Judgement first, then Conviction and Sanctified Judgement, then down the tree to Champion of the Light by level 40.

## Rotation and priority

**The opener.** Put a Blessing of Might on yourself if nobody has given you one. About a second and a half before the pull, put up your Seal: Seal of Righteousness while leveling, Seal of Command with Unsealed Wrath and raid buffs. Walk in and Judge on your first swing. Judgement never consumes your Seal in Forever, so there is nothing to put back up after it.

**The priority**, top to bottom, every global:

1. **Judgement** on cooldown. It is off the global cooldown, so press it the moment it is ready.
2. **Exorcism** on cooldown, against Undead and Demons only.
3. **Hammer of Wrath** when the target is below 20%, on its 6 second cooldown.
4. Your **Seal** if it is about to run out.
5. **Holy Strike** on cooldown.
6. **Consecration** with whatever mana is spare.

**Seals while leveling.** Stay in Seal of Righteousness. It fires on every swing, it scales with spell damage, and its Judgement hits hard. Do not twist the way 1.12 guides tell you: in Forever a new Seal simply replaces the old one, no swing gets both, and every swap costs a global and a Seal's worth of mana. Spend those globals on Holy Strike instead.

**The raid twist (Unsealed Wrath).** With raid buffs on you, Seal of Command becomes your main Seal. Let three swings land under Seal of Command, so Unsealed Wrath shows three stacks. Switch to Seal of Righteousness for exactly one swing: that swing releases all three stored Commands, each with its normal chance to fire. As soon as it lands (Unsealed Wrath's timer jumps back to 15 seconds), switch back to Seal of Command, and the next swing releases the stored Righteousness hit. Repeat. Two rules keep it from costing you: only twist above 30% mana, and only with raid buffs, because with little attack power Seal of Command does less than Seal of Righteousness and the swaps just burn mana. A swing timer and a stack counter on Unsealed Wrath make this easy.

**Judgements in a raid.** Your own Judgement is for damage. The debuff Judgements are assigned: one paladin keeps Judgement of the Crusader on the boss and another keeps Judgement of Wisdom. If that is you, open with that Seal, Judge it, then go to your damage Seal. The debuff lasts 40 seconds and Holy Strike with Sacred Arbiter refreshes every Judgement you have on the target, so after the first one it costs you nothing. On your own, skip Crusader: one Seal and one Judgement spent on it are worth more as damage.

**Consecration.** It is your most expensive button, so pace it. Cast it only while the mana left after it still pays for a Seal, a Judgement and a Holy Strike, and while your mana percent is at least four fifths of the percent of the fight left (halfway through, 40% mana or more). Fall under that and Consecration waits until you catch up. Spam it and your Seal drops and you go into the execute with nothing for Hammer of Wrath. Hold it too long and you end the fight with a full bar. Before 40, rank 1 does more over a fight than the higher ranks because it costs so little; from 40 move up to the newest rank you can afford. At 60, use rank 4 without raid buffs and rank 5, the top rank, with them. On two or more targets, cast it freely.

**The execute.** Below 20% Hammer of Wrath goes ahead of everything but Judgement and Exorcism. It is instant with Champion of the Light, so you never stop swinging to cast it. Save a Major Mana Potion for this phase if you are low.

**Movement and adds.** Hammer of Wrath works at range and Judgement from 10 yards, so keep pressing them while you close on the boss. On a pack, Consecration first, then the single-target priority on the kill target. Against a pack of Undead or Demons, Holy Wrath hits all of them; on anything else it does nothing, so leave it off your bars outside Scholomance, Stratholme and the like.

## Stat priority

Hit comes first until the cap: every miss is a lost swing and a lost Seal proc, and Precision and Artisan's Precision both help. Intellect is next: Champion of the Light turns 60% of it into spell damage for your Seals, Judgements and Consecration, and it is your mana bar. Then Strength, which Divine Strength raises by 10%, and crit, which feeds Vengeance as well as itself.

## Best in slot

## Enchants and consumables

## Races

**Human** is the Alliance pick: Ruthless is 1% crit on everything, and The Human Spirit helps your regeneration while leveling. **Dwarf** trades that for Artisan's Precision, 1% hit, and Stoneform for a hard-hitting pull. **Undead** is new to the class in Forever and the Horde's only paladin. Touch of the Grave can drain a little health from the target, at most once every 15 seconds. Will of the Forsaken breaks a fear, but no longer makes you immune to the next one, so use it after the fear lands. Cannibalize restores mana as well as health, which cuts your drinking while leveling.

## Professions

Mining and Blacksmithing make your two-handers and plate. Engineering adds explosives for trash packs, and it is the profession that adds damage when your gear is done. Herbalism and Alchemy keep you in mana potions and Elixirs of Holy Power, Forever's 40 Holy damage elixir, which reaches every Seal proc, Judgement and Consecration tick. Pick two that pay for your consumables.

## Leveling

Retribution is the paladin leveling spec, even if you want to heal or tank at 60. Pull with Judgement, Holy Strike on cooldown, and let Seal of Righteousness do the rest. Sanctified Judgement refunds most of each Seal, so on one or two mobs at a time you rarely drink.

From level 20 Consecration is your AoE: two or three mobs on you is the time for it, and rank 1 at 135 mana is the rank to keep on your bar through your 30s. A rank loses its spell damage bonus once you are far above the level you learned it, so from 40 move up to the newest rank you can afford, and keep moving up as you train them. Exorcism (level 20) and Holy Wrath (level 50) are easy kills on Undead and Demons, so quest zones full of them, like the Plaguelands, are paladin country. Repentance takes one Humanoid out of a bad pull, and Hammer of Justice stops a runner or a caster.

Common mistakes:

- Twisting Seals without Unsealed Wrath. It costs mana and globals for nothing.
- Letting the Seal drop. Recast it in its last second or two, not after it is gone.
- Spending the bar on Consecration on a single target and having nothing left for Hammer of Wrath.
- Swinging a one-hander. Two-Handed Weapon Specialization only works with a two-hander, and every Seal proc scales with the weapon.
- Ignoring hit. Below the cap, a point of hit beats a point of Strength.
