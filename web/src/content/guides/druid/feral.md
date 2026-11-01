---
title: Feral Druid in Forever
classSlug: druid
spec: feral
role: dps
build: 'FS1:1.60.1.70707:druid:night-elf:0/50212032123012123001/555023:'
raidBuild: 'FS1:1.60.1.70707:druid:night-elf:05002/50230032023011125001/505003:'
recommendedRaces: [night-elf, tauren]
statPriority: [Strength, Agility, Critical strike, Attack power, Feral attack power, Hit, Melee haste]
description: 'Talents, rotation, stats, and gear for Feral Druid in Forever, covering both Cat Form DPS and Bear Form tanking from the one tree.'
updated: 2026-10-31
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
  - label: 'Forever Sixty rotation data, data/curated/apl/druid-feral.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Feral Combat is one tree for two jobs: Cat Form melee DPS built on Rip and Ferocious Bite, and Bear Form tanking. This guide covers both.

### Strengths

- Cat DPS spends combo points on Rip and Ferocious Bite
- Shifting Power refills 40 energy for mana every 8 seconds
- Omen of Clarity is baseline, so Clearcasting hands out free attacks
- The same tree tanks in Bear Form, so one Druid can fill either role

### Weaknesses

- The cat and bear builds split the 51 points differently, so switching roles means a respec
- A cat gains nothing from Thick Hide's Armor
- A bear has no shield, Parry or Block

### In a raid

- **Role:** Melee DPS in Cat Form, or main tank in Bear Form
- **Brings:** Leader of the Pack, a critical strike aura for the group

## Talents and builds

Client 1.60.1.70707 reworked this tree: Berserk and Improved Shifting Power are gone, Shifting Power and Primal Bite are the two 31-point talents, Lacerate and Predatory Discipline are new, and Furor and Heart of the Wild swapped trees (Furor is now on Feral's top row, Heart of the Wild on Restoration's).

- **Ferocity** — up to 5 less Rage or Energy on Maul, Primal Bite, Swipe, Claw and Rake at rank 5.
- **Shredding Attacks** — up to 18 less energy on Shred at rank 3, the filler you press most.
- **Savage Fury** — up to 10% more damage at rank 2 on Claw, Rake, Shred, Maul and Swipe.
- **Sharpened Claws** and **Natural Instinct** — 6% more critical strike chance in Cat Form and 20% more critical damage on your melee abilities.
- **Predatory Strikes** — up to 150% of your level added to melee attack power in Cat or Bear Form at rank 3.
- **Blood Frenzy** — up to a 100% chance at rank 2 for bonus Rage on a Bear crit, and a 100% chance for a bonus combo point on a Cat crit.
- **Predatory Discipline** (1 point) — new in 1.60.1.70707: Ferocious Bite still deals its extra damage for the energy you have left after its cost, but no longer spends that energy. It is also the prerequisite for Shifting Power.
- **Leader of the Pack** — 3% more critical strike chance for the party and raid, on attacks and spells alike since 1.60.1.70707.
- **Rend and Tear** — up to 10% more damage from your melee abilities on a bleeding target. Your own Rip, Rake, Pounce or Lacerate counts, so once Rip is up the boss stays bleeding; the hits before your first Rip get no bonus. The raid build takes all five; the leveling build takes three, two of them from Feral Instinct, which only helps Swipe (about +1% in our level-60 search on the 1.60.1.70707 trees).
- **Shifting Power** — a 31-point talent behind Predatory Discipline since 1.60.1.70707; see below.

The Restoration points reach **Naturalist**, 2% more damage a rank (10% at 5/5). The leveling build also takes **Heart of the Wild** (10% more Strength in Cat Form) and two ranks of **Natural Shapeshifter**, which make Shifting Power cheaper. The raid build spends its Balance points on **Genesis** (10% more periodic damage, so a bigger Rip) and **Nature's Majesty** (4% critical strike chance) instead of Heart of the Wild. Primal Bite and Lacerate are bear abilities and are left out.

## Rotation and priority

**Cat DPS**: Tiger's Fury is gone: Forever druids cannot learn it. Omen of Clarity is a trainer spell now, so every druid has it from level 20, and a Clearcasting proc makes whichever ability you press next cost nothing. Most of the time that is Shred, the filler, so it needs no button of its own; a free Rip or Ferocious Bite saves less. Leave Rake out: it costs 40 energy for a small bleed and one combo point, and in this site's simulator that energy is worth more spent on Shred and Rip. Rip alone keeps the boss bleeding for Rend and Tear, and keeping Rake up as well still measured about 9% worse in the raid-ready run. At 5 combo points, cast Rip if it isn't already running; once Rip is up, spend 5 combo points on Ferocious Bite instead of letting them cap out uselessly. Shred fills every other global, building combo points the rest of the time.

**Shifting Power**: the tree's energy tool. It is instant, turns a slice of your base mana into 40 energy, and has an 8 second cooldown since client 1.60.1.70707 (Improved Shifting Power, which took 8 seconds off its old 16, is gone and its reduction is now the spell's own). Press it every time it comes off cooldown, but only when your energy is 60 or lower, so the whole 40 fits under the 100 cap; if you are above that, spend a few energy on Shred first and press it on the next global. It sits ahead of every finisher and builder in the order. The mana is otherwise unused in cat form, and Natural Shapeshifter makes it cheaper. Blizzard has not published whether it shares the global cooldown; our simulator treats it as an ordinary global, so if the real spell turns out to be off the global, you will do slightly better than this advice.

**Ferocious Bite with Predatory Discipline**: the bite no longer drains your energy, so there is no reason to hold it back; bite at 5 combo points whenever Rip has time left, as above. Waiting for a full bar before biting measured worse in our raid-ready run.

**Bear tank**: Bear Form tanking has its own guide, [Feral Bear Druid](/guides/druid/feral-bear), with its own simulated rotation, build and stat table.

## Stat priority

Weapon damage per second matters in Forever. In Cat Form your auto attack has your weapon's DPS at one swing a second, and Shred deals 155% of that hit, so a high-DPS two-hander beats a caster staff that carries more Strength and Agility.

## Best in slot

For Bear tanking, look for armor, stamina, and defense skill first.

## Enchants and consumables

## Races

## Professions

Skinning and Leatherworking is the standard Feral pairing in Classic-era play: Skinning is free crafting material from every kill, and Leatherworking turns it into armor a melee Druid can actually wear.

## Leveling

Feral is generally considered the strongest leveling spec for Druid: Cat Form gives questing speed and burst damage, and Bear Form gives the durability to solo tougher pulls or tank early dungeons.
