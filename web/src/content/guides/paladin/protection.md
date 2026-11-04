---
title: Protection Paladin in Forever
classSlug: paladin
spec: protection
role: tank
build: 'FS1:1.60.1.70755:paladin:dwarf:05005/15302102513501031/5003:'
recommendedRaces: [dwarf, undead]
statPriority:
  [
    Defense,
    Block,
    Dodge,
    Parry,
    Agility,
    Block value,
    Stamina,
    Hit,
    Strength,
    Critical strike,
    Armor,
    Attack power,
  ]
description: 'Holding threat with Righteous Fury, Holy Shield and Holy Strike, taunting with Judgement of Fury, and keeping your mana.'
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
  - label: 'Forever Sixty raid phase data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/data/curated/loot/forever-raid-phases.json
    kind: site
  - label: 'Forever Sixty build data, data/builds/1.60.1.70755/races.json'
    url: https://foreversixty.gg/data/1.60.1.70755/races.json
    kind: datamined
---

## Overview

Protection tanks with a shield and Holy damage: Righteous Fury on, Holy Shield on cooldown, Holy Strike every 6 seconds, and Judgement of Fury as your taunt.

### Strengths

- Righteous Fury turns your Holy damage into huge threat, and Holy Shield adds more on every block
- Holy Shield, Redoubt and Shieldsworn keep your block chance high, and Shieldsworn lets you block spells
- Judgement of Fury taunts, and Swift Judgement gives you a second one
- Three defensives (Templar's Bulwark, Divine Shield, Divine Protection) plus Lay on Hands

### Weaknesses

- Mana, not cooldowns, limits what you press: run dry and your threat stops
- Templar's Bulwark, Divine Shield and Divine Protection share Forbearance, so you get one at a time
- No Blessing of Salvation in Forever, so you hold threat against damage dealers at full threat

### In a raid

- **Role:** Main or off tank. You hold the boss, and you swap with Judgement of Fury when the fight calls for it
- **Brings:** Devotion Aura for your group, Greater Blessing of Kings or Might, and a Lay on Hands for the raid's worst moment

## Talents and builds

The build is 10 Holy, 33 Protection and 8 Retribution.

- **Redoubt** (5 points): a hit that damages you has a chance to add 30% block for 10 seconds or 5 blocks.
- **Precision** (3 points): 3% hit, so the Judgement and Holy Strike that carry your threat land.
- **Sacred Duty** (2 points): a minute off Divine Shield, Divine Protection and Templar's Bulwark, and 10 seconds off Forbearance.
- **Templar's Bulwark** (1 point): an absorb worth your whole health bar for 8 seconds. It leads to Swift Judgement.
- **Shield Specialization** (2 points): 30% more absorbed by your shield, and a block gives back 4% of your mana, at most once every 3 seconds. This is where your mana comes from.
- **Swift Judgement** (1 point): a second Judgement of your Seal on its own one minute cooldown, free and off the global cooldown.
- **Improved Righteous Fury** (3 points): 9% less damage taken while Righteous Fury is on. It is a mitigation talent now, not a threat one.
- **Anticipation** (5 points): 20 Defense Skill.
- **Holy Shield** (1 point): 30% more block for 10 seconds or 4 blocks, with Holy damage on each block and 20% extra threat on that damage. It sits on the 21 point row now.
- **Iron Creed** (3 points): 30% more threat on Holy Strike, and while Righteous Fury is on each Holy Strike cuts your damage taken by 6% for 6 seconds.
- **Shieldsworn** (1 point): the capstone. Holy Strike comes back every 6 seconds, which keeps Iron Creed up, and each strike fires Holy Shield's damage and Shield Specialization's mana without spending a charge. For 6 seconds after a strike you can block spells too.

**Toughness** (10% armor from items) and a point of **Divine Intellect** fill the tree. In Retribution, **Deflection** is 5% parry, the only parry talent a paladin has, and **Conviction** is 6% crit on everything you do. In Holy, **Divine Strength** is 10% Strength and **Improved Seals** adds 15% to Seal of Fury and its Judgement.

Left out: Reckoning (extra attacks after a block or a crit, worth less than Conviction), Stoicism, Guardian's Favor and Improved Hammer of Justice. Pick those up for PvP or utility. There is no Blessing of Sanctuary in Forever, so the tree gives you no blessing of your own.

## Rotation and priority

**Before the pull.** Righteous Fury is a 30 minute buff, not a stance: cast it before the raid (it costs 30% of your base mana, so do it before you drink) and keep it on all night. Without it you make a fraction of the threat. Get Greater Blessing of Kings on yourself, and keep Devotion Aura on.

**The opener.** Put up Seal of Fury about 3 seconds before the pull. Pull with Judgement: under Seal of Fury it taunts, so the boss comes straight to you. Press Holy Shield about 2 seconds before the boss reaches you, so all four charges are there for its first swings. Then Swift Judgement, Holy Strike and Hammer of the Righteous in the first few globals. There is no Blessing of Salvation in Forever, so your damage dealers have no threat drop: give them a few seconds on the opener.

**The priority**, top to bottom:

1. **Seal of Fury** in its last 2 seconds, if you still have Holy Shield's mana after it.
2. **Holy Shield** on cooldown. Its cooldown is its duration, so it is always up while it has charges, and it is the first claim on your mana.
3. **Holy Strike** on cooldown: every 6 seconds with Shieldsworn, 20 mana, and your damage reduction from Iron Creed.
4. **Hammer of the Righteous** on cooldown, once Holy Shield's mana is safe. It hits your target and the enemies next to it, and it is your biggest single threat button.
5. **Swift Judgement** whenever it is ready.
6. **Judgement** on cooldown, once Holy Shield's mana is safe.
7. **Exorcism** on Undead and Demons.
8. **Consecration** only above half mana. It makes far less threat per mana than Hammer of the Righteous, so it is how you spend mana you will not need.

**Mana.** You run on blocks: Shield Specialization gives back 4% of your mana on a block, and Holy Strike under Shieldsworn triggers it too. Hold Holy Shield's cost in reserve at all times. If you are going dry, cut Consecration first, then Hammer of the Righteous, and never Holy Shield.

**Taunts and swaps.** Judgement under Seal of Fury makes the target attack you for 3 seconds, from 10 yards. Swift Judgement judges the same Seal, so it is a second taunt on its own minute. On a fight with a tank swap, hold one of them for the swap instead of spending it on threat.

**Adds.** Consecration and Hammer of the Righteous are your multi-target threat; Holy Shield and Shieldsworn's block keep the damage down. Judge the add that is loose, Holy Strike your main target.

**Defensives.** Templar's Bulwark first, at about a third of your health: a full health bar of absorb for 8 seconds. Divine Shield and Divine Protection come after it, once Forbearance is gone. All three put Forbearance on you, so plan one per bad moment, not three in a row. Lay on Hands heals you to full on a 20 minute cooldown, but it takes every point of your mana with it, and with no mana you have no Holy Shield and no threat. Save it for the swing that would kill you, and have a mana potion ready after it.

## Stat priority

Defense comes first: it takes hits away and lowers the boss's chance to crit you. Block, Dodge and Parry come next because they take hits away or blunt them, and Block also feeds Holy Shield's damage, Redoubt and Shield Specialization's mana. Agility is dodge and armor at once, and block value makes every block take more off the hit. Stamina is still health, and armor still takes a slice off every physical hit, so neither ever goes to zero. Hit sits low once Precision's 3% is in.

## Best in slot

## Enchants and consumables

## Races

**Dwarf** is the Alliance tank. Stoneform takes physical damage off for a few seconds, a free defensive on a hard pull, and Artisan's Precision is 1% hit. **Human** brings Ruthless, 1% crit, and The Human Spirit. **Undead** is new to paladins in Forever and the only Horde option. Will of the Forsaken breaks a fear, so a feared tank gets back to the boss before it turns on the raid, but it no longer makes you immune, so press it after the fear lands. Touch of the Grave adds a small heal off your attacks, once every 15 seconds at most.

## Professions

Mining and Blacksmithing make plate, shields and the tanking weapons. Engineering gives you explosives for trash threat. With Alchemy, Elixir of Holy Power adds to every Holy threat button you have. First Aid is free, and you will use bandages between pulls while leveling.

## Leveling

Level as Retribution and tank dungeons in it: you get Righteous Fury and Seal of Fury at level 10, which is enough to hold a dungeon group from Ragefire Chasm and the Deadmines on. Switch to Protection when you are tanking more than you are questing. Holy Shield comes at 30, with new ranks at 40, 50 and 60, and Hammer of the Righteous at 40. If you tank in Protection at 20, take Templar's Bulwark before a second point of Sacred Duty: a full health bar of absorb keeps a low-level tank alive far more often than a shorter cooldown on it.

Common mistakes:

- Forgetting Righteous Fury after a death or a long break. Check your buffs before every pull.
- Spending mana on Consecration and having none for Holy Shield.
- Pressing Holy Shield while the boss is still running in, so its charges are gone before the first swing.
- Burning both Judgements on threat right before a tank swap.
- Using Divine Shield and Templar's Bulwark back to back: they share Forbearance, so the second one is not there.
