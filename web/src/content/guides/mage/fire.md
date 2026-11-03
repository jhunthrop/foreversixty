---
title: Fire Mage in Forever
classSlug: mage
spec: fire
role: dps
build: 'FS1:1.60.1.70753:mage:gnome:2520150000002/51521001230133031/003:'
raidBuild: 'FS1:1.60.1.70753:mage:gnome:2530050000002/50520201230133031/003:'
recommendedRaces: [gnome, orc]
statPriority: [Hit, Intellect, Spell power, Fire power, Critical strike, Spell haste, Spell penetration]
description: 'Fire Mage in Forever: the pre-pull Pyroblast, Scorch spam, Heating Up and Combustion, Ignite and threat, mana over a boss fight, talents, stats and leveling.'
updated: 2026-11-02
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Talents Forever (demo transcription)'
    url: https://talentsforever.com/data.json
    kind: community
  - label: 'Forever Sixty stat weights, data/curated/stat-weights.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty rotation data, data/curated/apl/mage-fire.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Fire spams Scorch, keeps its own Improved Scorch stacks up, turns crits into fast Pyroblasts with Heating Up, and cashes Combustion into a run of Ignites.

### Strengths

- The cheapest filler a mage has: Scorch is 150 mana for a 1.5 second cast
- Every crit pays twice, once through Ignite and once through Heating Up
- Fire Power rewards you for every burn you keep on the boss
- Fire Blast and a 1.5 second Pyroblast give you damage on the move

### Weaknesses

- Runs out of mana before a long boss dies without gems, potions and runes
- No Presence of Mind and no Ice Block: they sit in the Arcane and Frost trees, which this build only dips into
- Fire-immune and Fire-resistant bosses, Ragnaros first among them, shut the spec down

### In a raid

- **Role:** Sustained ranged caster DPS with a Combustion window every 3 minutes
- **Brings:** Arcane Brilliance, Remove Lesser Curse, Polymorph and Counterspell. Improved Scorch now only boosts your own Fire damage, so every Fire mage stacks their own.

## Talents and builds

- **Improved Scorch** (3 points): every Scorch that lands adds 3% Fire damage taken from you, up to 5 stacks for 30 seconds. A missed Scorch adds no stack.
- **Heating Up** (1 point): Scorch, Fire Blast and Fireball crits cut your next Pyroblast's cast by 25%, up to 3 stacks, and every Fire Blast hit adds a stack on top. Three stacks make Pyroblast a 1.5 second cast.
- **Pyroblast** (1 point): the payoff for Heating Up, and the spell you cast before the pull.
- **Ignite** (5 points): every Fire crit burns the boss for another 40% of the hit over 4 seconds.
- **Fire Power** (3 points): 3% more Fire damage for each of your Fire effects on the target, up to 12%. Improved Scorch, Ignite and Pyroblast's burn are three of them.
- **Incineration** and **Critical Mass** (8 points): 10% more crit on Scorch and Fire Blast and 6% more on every Fire spell, which is what keeps Ignite and Heating Up rolling.
- **Master of Elements** (3 points): every Fire crit refunds 30% of the spell's base cost. It is the steady trickle that makes the gems and potions last.
- **Improved Fire Blast** (2 points): Fire Blast every 6 seconds instead of 8, and every Fire Blast is a Heating Up stack.
- **Combustion** (1 point): each Fire hit adds 10% Fire crit until three non-periodic crits land, on a 3 minute cooldown.

Fire takes exactly 31 points to Combustion and skips Blast Wave and Improved Flamestrike, which are AoE only. Accelerant only speeds up Fireball, Frostfire Bolt and Arcane Wrath, none of which you cast at 60, so it gets at most a spare point.

The other 20 points buy mana. Seventeen in Arcane reach **Arcane Concentration** (5 points, a free spell 10% of the time) and two ranks of **Arcane Meditation** (a third of your mana regeneration keeps running while you cast). With Mage Armor on top, most of your regeneration runs through the whole fight. Three go into **Precision** in Frost for 6% hit.

The 60 build keeps a point in **Impact** for the odd stun on mobs while you farm, and its spare Accelerant point. The raid build moves Accelerant and Impact, plus the point of Magic Absorption, into two ranks of **Burning Soul** and a third rank of Arcane Resilience. Burning Soul is less threat and pushback on your Fire spells, and raid bosses cannot be stunned, so Impact does nothing there. Build it in the [talent planner](/planner?class=mage).

## Rotation and priority

**Opener.** Mage Armor, Arcane Brilliance and your food buff go up before the pull. Conjure a Mana Ruby and a Mana Citrine and drink back to full. Start Pyroblast 6 seconds before the pull so it lands just after the tank's first hit; it costs no Heating Up, and its burn starts your Fire Power count. Then Scorch to 5 stacks. Fire Blast whenever it is ready.

**The loop.** Keep 5 Improved Scorch stacks rolling: every filler Scorch refreshes them, so you only think about the refresh after a movement break or a target swap. Fire Blast every time it is off cooldown. Pyroblast only on 3 stacks of Heating Up. Everything else is Scorch. A Pyroblast at the full 6 seconds, or on 1 or 2 stacks, loses to Scorch, and Fireball as a filler loses to Scorch even while your mana is full.

**Combustion.** Press it once the boss carries 5 stacks, together with your on-use trinkets and your racial: Eureka! on a Gnome, Blood Fury on an Orc, Berserking on a Troll. A fresh tank needs a few seconds before you do. Combustion lasts until three crits land, and every crit in it is an Ignite and a Heating Up stack, so the window usually ends in a 1.5 second Pyroblast.

**Ignite and threat.** Ignite's damage is your damage and your threat. Fire mages pull more than anyone else in the first twenty seconds, mostly off a Combustion window that crits three times in a row. You have no Ice Block and no Invisibility, so the only fix is not to pull: watch your threat, hold Combustion on a fresh tank, and stop casting for a global if you are close. Burning Soul takes some off the top.

**Mana in a 3 minute fight.** Fire runs dry first. In order: a Mana Ruby when you are 1,200 mana down, a Major Mana Potion, a Demonic Rune, then the Mana Citrine two minutes after the Ruby. Gems, runes and the Robe of the Archmage share a 1 minute cooldown with each other, so spread them out. Evocation is an 8 minute cooldown and comes once a fight: channel it near empty with more than 30 seconds left, never in the last 30. Keep Fire Blast on cooldown even when you are short: skipping it saves less damage than it costs. Never downrank Scorch; the top rank is already the cheapest damage you have.

**AoE.** Flamestrike is a 3 second cast that burns the ground for 8 seconds, so cast it where the tank is going to stand, then follow with Arcane Explosion in the pack or a Blizzard from range. Flamestrike costs 990 mana, so two of them and you are on the wand. Let the tank get a few seconds of threat first, and Frost Nova and Blink out if the pack turns on you.

**Execute and movement.** Fire has no execute. On the move, Fire Blast is instant and a 3-stack Pyroblast is short enough to cast between steps. Scorch's stacks last 30 seconds, so a short movement phase does not drop them.

Frostfire Bolt is not a Fire filler. It costs more than Scorch, takes twice as long and its burn mixes Frost and Fire, which Fire Power does not count.

## Stat priority

Hit and spell power lead for Fire until the hit cap: every missed Scorch is a missed stack, a missed Ignite and a missed Heating Up chance. Intellect is your mana bar, and Fire is the spec that runs out. Crit feeds Ignite, Heating Up and Master of Elements at once.

## Best in slot

## Enchants and consumables

Fire casts nothing but Fire, so the Fire-only elixir and enchants are full value for you. Bring Major Mana Potions and Demonic Runes to every boss; you will use both.

## Races

**Gnome** is the best Alliance Fire mage: Eureka! is 10% more damage and 10% cheaper spells for 8 seconds on a 2 minute cooldown, and Expansive Mind adds 5% mana. Press it with Combustion. **Human** brings 1% crit, which Fire turns into Ignite. On the Horde, **Orc** Blood Fury adds 10% spell power for 15 seconds every 2 minutes and **Troll** Berserking adds 10% cast speed for 15 seconds every 3 minutes. Both line up with Combustion. **Undead** gets a small Shadow drain from Touch of the Grave and Will of the Forsaken for fear fights.

## Professions

Tailoring and Enchanting are the standard pair: Tailoring makes the early cloth sets and Enchanting puts the spell power enchants on your own gear. Alchemy pays for itself once you raid every week, because Fire drinks a Major Mana Potion on every boss. Gnomes who want the explosives take Engineering.

## Leveling

Fireball is your nuke until Scorch at 22, then Scorch and Fire Blast take over on single targets. Fire levels best from 30, once Improved Scorch and Pyroblast are in, and Impact's stuns stop casters and runners for you. Open on a mob with Pyroblast from range, Scorch or Fireball as it closes, and finish it with Fire Blast. Keep Frost Nova and Blink for anything that reaches you, and Frostbolt for kiting: Fire has no slow of its own. Flamestrike into Arcane Explosion clears packs once you have the Stamina for it. Train Mage Armor at 34 and wear it for the mana; switch to Ice Armor when you are taking melee hits in the open world.
