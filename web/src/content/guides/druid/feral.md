---
title: Feral Druid in Forever
classSlug: druid
spec: feral
role: dps
build: 'FS1:1.60.1.70532:druid:night-elf:01/55232032121032012001/50532:'
raidBuild: 'FS1:1.60.1.70532:druid:night-elf:05002/45211031021032212001/5053:'
recommendedRaces: [night-elf, tauren]
statPriority: [Strength, Agility, Critical strike, Attack power, Feral attack power, Hit, Melee haste]
description: 'Talents, rotation, stats, and gear for Feral Druid in Forever, covering both Cat Form DPS and Bear Form tanking from the one tree.'
updated: 2026-10-07
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
- Shifting Power refills 40 energy for mana every 16 seconds
- Tiger's Fury refunds energy on a 30 second cooldown
- The same tree tanks in Bear Form, so one Druid can fill either role

### Weaknesses

- The cat and bear builds split the 51 points differently, so switching roles means a respec
- A cat gains nothing from Thick Hide's Armor
- A bear has no shield, Parry or Block

### In a raid

- **Role:** Melee DPS in Cat Form, or main tank in Bear Form
- **Brings:** Leader of the Pack, a critical strike aura for the group

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31. Two structural changes affect this tree specifically: Feral Charge splits into a Cat version, a gap-closing leap that dazes on landing, sitting next to the older Bear version built for interrupting; and Furor's cat-form energy refund now scales with how long you'd been out of form beforehand, rather than handing back a flat amount every time.

- **Ferocity** — up to 5 less Rage or Energy on Maul, Mangle, Swipe, Claw, and Rake at rank 5.
- **Thick Hide** — extra base armor while shapeshifted, scaling with defense skill, essential for a Bear tank and useful padding for a Cat.
- **Savage Fury** — up to 10% more damage at rank 2 on Claw, Rake, Shred, Maul, and Swipe.
- **Predatory Strikes** — up to 150% of your level added to melee attack power in Cat or Bear Form at rank 3.
- **Blood Frenzy** — up to a 100% chance at rank 2 for bonus Rage on a Bear crit, and a 100% chance for a bonus combo point on a Cat crit, supporting both forms from one talent. (Named "Primal Fury" in an earlier draft of this guide; this build's own data calls it Blood Frenzy, so the talent section and `build:` string below use that name.)
- **Berserk** — the capstone: Mangle hits up to 3 targets with no cooldown, and critical strike chance on combo-point generators rises 100%, for 15 seconds.

The point split differs by role. This build (a cat-DPS spread) spends 35 points in Feral and reaches Berserk through the offensive column: Savage Fury, Sharpened Claws, Predatory Strikes, Blood Frenzy and Leader of the Pack, then Shredding Attacks at 3 of 3 to open Shifting Power, Predatory Instincts at 2 of 2 for the bigger crits, and Heart of the Wild at 5 of 5. The tier-gate fillers Feral Swiftness, Feral Instinct, Brutal Impact and Feral Charge come along on the way, and Thick Hide is held at a single rank, since a Cat gains nothing from its armor in the simulator. One point goes to Genesis in Balance, and the other 15 go to Restoration: Naturalist maxed for the flat damage bonus, Nature's Focus and Subtlety kept as the cheapest way to meet the tier gates, and 2 ranks of Natural Shapeshifter, which the two points freed from Thick Hide pay for. Furor is not in this build. In the simulator a Cat that powershifts out of and back into Cat Form gains nothing from Furor, because the refund is a share of the energy you already had, and Shifting Power is a better refill than any powershift. In the search run this spread beats the old Furor and Natural Shapeshifter spread by about 21%, almost all of it from Shredding Attacks, Shifting Power, Predatory Instincts and Heart of the Wild; moving the two Thick Hide ranks into Natural Shapeshifter adds about another 2% on top, because the simulator credits Natural Shapeshifter with a small gain (the probe measures it as a cheaper Shifting Power). A bear-tank build stays in Feral for more of its points, weighting the defensive column instead (Feral Instinct, Thick Hide, Natural Reaction) before reaching Berserk, and typically only dips into Restoration for Furor rather than Naturalist. Open the planner at [/planner?class=druid](/planner?class=druid) to build this out.

**Raid build.** For a raid the search found about +10% over the leveling build above in the raid-ready run, most of it from four more points of Genesis, two of Nature's Majesty and two of Improved Shifting Power in place of one rank each of Ferocity and Savage Fury, two of Natural Shapeshifter and four utility points: two of Feral Instinct (stealth detection and a Swipe bonus the cat rotation never casts), Brutal Impact (a stun) and Feral Charge. Leader of the Pack (the raid's crit aura), Shifting Power, Nature's Focus and Subtlety stay.

## Rotation and priority

**Cat DPS**: open with Tiger's Fury, since it costs nothing and refunds energy on its own 30-second cooldown, and keep it on cooldown throughout. Leave Rake out: it costs 40 energy for a small bleed and one combo point, and in this site's simulator that energy is worth more spent on Shred and Rip, which measured about 5% better on the bare character and about 9% better in the raid-ready run. At 5 combo points, cast Rip if it isn't already running; once Rip is up, spend 5 combo points on Ferocious Bite instead of letting them cap out uselessly.

**Shifting Power**: this is the new energy tool in the tree, replacing King of the Jungle. It is instant, turns a slice of your base mana into 40 energy, and has a 16 second cooldown (Improved Shifting Power takes 4 or 8 seconds off that). Press it every time it comes off cooldown, but only when your energy is 60 or lower, so the whole 40 fits under the 100 cap; if you are above that, spend a few energy on Shred first and press it on the next global.

**Bear tank**: Bear Form tanking has its own guide, [Feral Bear Druid](/guides/druid/feral-bear), with its own simulated rotation, build and stat table.

## Stat priority

The table above is this level range's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this level range's sample size, to tell its true value apart from zero.

## Best in slot

For Cat DPS, look for pieces that lead with attack power and agility, then crit and hit, following the priority above. For Bear tanking, look for armor, stamina, and defense skill first.

## Enchants and consumables

For Cat DPS, target attack power and agility on enchants and consumables.

## Races

Druid's playable races are Night Elf and Tauren on their respective factions, plus Skyborne on both sides — Druid is one of the few classes open to both the Alliance's High Order Skyborne and the Horde's Windshaper Skyborne. None of Blizzard's six confirmed new race and class pairs touches Druid, so this list is unchanged from 1.12.

For Alliance, **Night Elf** is the strongest pick for either role: Shadowmeld lets you stealth even in combat on a 2-minute cooldown, which meshes with Feral's own Prowl mechanic for repositioning or a surprise opener, and Quickness's passive dodge helps a Bear tank's avoidance and a Cat's survivability alike. For Horde, **Tauren** is the pick: War Stomp stuns up to 5 enemies within 8 yards on a 2-minute cooldown, a strong tool for a Bear tank picking up multiple adds, and Endurance's 5% health and 1% hit help both the tanking and DPS role. Skyborne is available to Druid on both factions, but its racials — Read Ley Line's health and mana regen, Elemental Insight's damage bonus against Elementals — lean toward caster support rather than melee combat, so it's a less targeted pick for Feral than Night Elf or Tauren.

## Professions

Skinning and Leatherworking is the standard Feral pairing in Classic-era play: Skinning is free crafting material from every kill, and Leatherworking turns it into armor a melee Druid can actually wear.

## Leveling

Feral is generally considered the strongest leveling spec for Druid: Cat Form gives questing speed and burst damage, and Bear Form gives the durability to solo tougher pulls or tank early dungeons.
