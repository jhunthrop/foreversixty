---
title: Protection Paladin in Forever
classSlug: paladin
spec: protection
role: tank
build: 'FS1:1.60.1.70790:paladin:dwarf:05005/15302102513501031/5003:'
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
    Expertise,
    Strength,
    Critical strike,
    Armor,
    Attack power,
  ]
description: 'Holding threat with Righteous Fury, Holy Shield and Holy Strike, taunting with Judgement of Fury, and keeping your mana.'
updated: 2026-11-04
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Forever Sixty, paladin talents'
    url: https://foreversixty.gg/planner?class=paladin
    kind: datamined
  - label: "Blizzard, What's Next panel recap"
    url: https://news.blizzard.com/en-gb/article/24303862/world-of-warcraft-forever-whats-next-panel-recap
    kind: blizzard
  - label: 'Forever Sixty, races'
    url: https://foreversixty.gg/data/1.60.1.70790/races.json
    kind: datamined
---

## Overview

Protection tanks with a shield and Holy damage: Righteous Fury on, Holy Shield on cooldown, Holy Strike every 6 seconds, and Judgement of Fury as your taunt.

### Strengths

- Righteous Fury turns your Holy damage into huge threat, and Holy Shield adds more on every block
- Holy Shield and Redoubt keep your block chance high, and Shieldsworn fires Holy Shield's damage on every Holy Strike while it is up and lets you block spells
- Judgement of Fury taunts, and Swift Judgement gives you a second one
- Three defensives (Templar's Bulwark, Divine Shield, Divine Protection) plus Lay on Hands

### Weaknesses

- Mana, not cooldowns, limits what you press: run dry and your threat stops
- Templar's Bulwark, Divine Shield and Divine Protection share Forbearance, so you get one at a time
- Forever has no Blessing of Salvation, so your damage dealers keep all their threat and you have to out-threat them without that help

### In a raid

- **Role:** Main or off tank. You hold the boss, and you swap with Judgement of Fury when the fight calls for it
- **Brings:** Devotion Aura for your group, Greater Blessing of Might (or Kings where no hunter's Heart of the Lion covers the raid), and a Lay on Hands for the raid's worst moment

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

Left out: Reckoning (extra attacks after a block or a crit; it adds a little threat, but paying for it with Anticipation's Defense gets you killed on hard pulls), Stoicism, Guardian's Favor and Improved Hammer of Justice. Pick those up for PvP or utility. There is no Blessing of Sanctuary in Forever, so the tree gives you no blessing of your own.

## Rotation and priority

**Before the pull.** Righteous Fury is a 30 minute buff, not a stance: cast it before the raid (it costs 30% of your base mana, so do it before you drink) and keep it on all night. Without it you make a fraction of the threat. Get Greater Blessing of Kings on yourself unless a hunter's Heart of the Lion already covers you: the two are the same 10% to every stat and do not stack, so with a hunter in the raid take Might instead. Keep Devotion Aura on.

**The opener.** Put up Seal of Fury about 3 seconds before the pull. Pull with Judgement: under Seal of Fury it taunts, so the boss comes straight to you. Press Holy Shield about 2 seconds before the boss reaches you, so all four charges are there for its first swings. Then Swift Judgement, Holy Strike and Hammer of the Righteous in the first few globals. There is no Blessing of Salvation in Forever, so your damage dealers have no threat drop: give them a few seconds on the opener.

**The priority**, top to bottom:

1. **Seal of Fury** in its last 2 seconds, if you still have Holy Shield's mana after it.
2. **Holy Shield** on cooldown. Its cooldown is its duration, so it is always up while it has charges, and it is the first claim on your mana.
3. **Holy Strike** on cooldown: every 6 seconds with Shieldsworn, 20 mana, and your damage reduction from Iron Creed.
4. **Hammer of the Righteous** on cooldown, once Holy Shield's mana is safe. It hits your target and up to three enemies next to it, for Holy damage that Righteous Fury multiplies, so it is where your threat on a pack comes from.
5. **Swift Judgement** whenever it is ready.
6. **Judgement** on cooldown, once Holy Shield's mana is safe.
7. **Exorcism** on Undead and Demons.
8. **Consecration** last, with mana to spare: cast it while the mana you have, plus what you regain less what you spend, still pays for it by the end of the fight, and also once you drop below 35% health, when mana held back may never get used. It makes far less threat per mana than Hammer of the Righteous, so it is how you spend mana you will not need. While you level, rank 1 is the better Consecration at 30 to 33, 40 to 48 and 50 to 57, because it goes down far more often; at the other levels use your newest rank, and rank 1 again at 60 without raid buffs.

**Mana.** You run on blocks: Shield Specialization gives back 4% of your mana on a block, and Holy Strike under Shieldsworn triggers it too. Hold Holy Shield's cost in reserve at all times. If you are going dry, cut Consecration first, then Hammer of the Righteous, and never Holy Shield.

**Taunts and swaps.** Judgement under Seal of Fury makes the target attack you for 3 seconds, from 10 yards. Swift Judgement judges the same Seal, so it is a second taunt on its own minute. On a fight with a tank swap, hold one of them for the swap instead of spending it on threat.

**Adds.** Consecration and Hammer of the Righteous are your multi-target threat; Holy Shield and Shieldsworn's block keep the damage down. Judge the add that is loose, Holy Strike your main target.

**Defensives.** Templar's Bulwark first, at about a third of your health: a full health bar of absorb for 8 seconds. Divine Shield and Divine Protection come after it, once Forbearance is gone. All three put Forbearance on you, so plan one per bad moment, not three in a row. Lay on Hands heals you to full on a 20 minute cooldown, but it takes every point of your mana with it. Your mana keeps regenerating afterwards, but until it comes back you have no Holy Shield and little threat, so save it for the swing that would kill you and have a mana potion ready. As a Dwarf, add Stoneform when you drop under 70% health: 10% less physical damage for 8 seconds, every 3 minutes. It takes a global, but it puts no Forbearance on you.

## Stat priority

Defense takes hits away and lowers the boss's chance to crit you. Stamina is your health bar, and health is what gives every other stat on this list time to work. Dodge and Parry take hits away, Agility is dodge and armor at once, and block value makes every block take more off the hit. More block chance adds nothing you can measure at 60, but keep the shield: your blocks still power Holy Shield's damage and Shield Specialization's mana. Armor takes a slice off every physical hit. Hit matters up to the cap and nothing past it, and Precision already gives you 3% of it. Expertise, which only a few raid sets carry, lowers the chance the boss dodges or parries your swings and Holy Strike, which is threat. The table above has the order.

## Best in slot

## Enchants and consumables

## Races

**Dwarf** is the Alliance tank. Dwarf and Human take hits and hold threat about the same, so Stoneform tips it: it takes physical damage off for a few seconds, a defensive that costs no mana on a hard pull, and Artisan's Precision is 1% hit. **Human** brings Will to Survive, which breaks a stun so you get back to holding the boss, and Ruthless, 1% crit for a little more threat. The Human Spirit does little for a tank. **Undead** is new to paladins in Forever and the only Horde option. Will of the Forsaken breaks a fear, so a feared tank gets back to the boss before it turns on the raid, but it no longer makes you immune, so press it after the fear lands. Touch of the Grave adds a small heal off your attacks, once every 15 seconds at most.

## Professions

Mining and Blacksmithing make plate, shields and the tanking weapons. Engineering gives you explosives for trash threat. With Alchemy, Elixir of Holy Power adds to Holy Shield, Holy Strike, Seal of Fury, your Judgements and Consecration. Hammer of the Righteous hits off your weapon, so the elixir does nothing for it. First Aid is free, and you will use bandages between pulls while leveling.

## Leveling

### Levels 1 to 19

Devotion Aura comes at 1, Judgement at 4, Holy Strike at 6 and Hammer of Justice at 8. At 10 you get Righteous Fury and Seal of Fury: with a shield on, every swing under Seal of Fury gives you an absorb, and judging it hits for Holy damage and taunts for 3 seconds.

At 12, the class quest from your trainer teaches Redemption. Holy Strike rank 2 also arrives at 12: put it on your bar straight away, the extra threat is worth its mana.

Most tanks level as Retribution and pick these up when they respec; the Retribution guide has those points. If you level as Protection, your first point at 10 is Divine Intellect for 2% Intellect. Redoubt takes 11 to 15: hits on you can give 30% more block chance for 10 seconds or 5 blocks. Precision takes 16 to 18 for 3% hit, and the first point of Sacred Duty at 19 takes 30 seconds off Divine Protection.

Take quest rewards over vendor gear. To tank, pair a one-hander with a shield as soon as you have both, since Seal of Fury's absorb needs the shield. Paladins cannot use wands. You wear mail until plate at 40.

### Levels 20 to 60

Level as Retribution and tank dungeons in it: Righteous Fury and Seal of Fury are enough to hold a dungeon group from Ragefire Chasm and the Deadmines on. Switch to Protection when you are tanking more than you are questing. Holy Shield comes at 30, with new ranks at 40, 50 and 60, and Hammer of the Righteous at 40. If you tank in Protection at 20, take Templar's Bulwark before a second point of Sacred Duty: a full health bar of absorb keeps a low-level tank alive far more often than a shorter cooldown on it. Shield Specialization follows at 21 and 22, since your blocks are your mana, then the second point of Sacred Duty and a point of Toughness. Swift Judgement comes at 25, the first level its row opens: a second Judgement and a second taunt are worth more threat than the rest of Toughness, which fills 26 to 29.

Common mistakes:

- Forgetting Righteous Fury after a death or a long break. Check your buffs before every pull.
- Spending mana on Consecration and having none for Holy Shield.
- Pressing Holy Shield while the boss is still running in, so its charges are gone before the first swing.
- Burning both Judgements on threat right before a tank swap.
- Using Divine Shield and Templar's Bulwark back to back: they share Forbearance, so the second one is not there.
