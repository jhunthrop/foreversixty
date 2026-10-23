---
title: Arms Warrior in Forever
classSlug: warrior
spec: arms
role: dps
build: 'FS1:1.60.1.70532:warrior:human:03325213032515001/0505/005:'
raidBuild: 'FS1:1.60.1.70532:warrior:human:02305213032515001/55050000001/2:'
recommendedRaces: [human, orc]
statPriority: [Strength, Attack power, Critical strike, Hit, Melee haste, Agility]
description: 'Talents, rotation, stat priority, and race picks for Arms Warrior in Forever, with beta-versus-projection called out.'
updated: 2026-09-24
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Talents Forever (demo transcription)'
    url: https://talentsforever.com/data.json
    kind: community
  - label: 'Warcraft Tavern: WoW Classic PvE DPS Warrior guide, Arms section'
    url: https://www.warcrafttavern.com/wow-classic/guides/pve-dps-warrior/
    kind: community
  - label: 'Icy Veins, the racial rework'
    url: https://www.icy-veins.com/wow-forever/news/the-racial-rework-gives-every-race-a-reason-to-matter-in-warcraft-forever/
    kind: community
  - label: 'Forever Sixty stat weights, data/curated/specs.json'
    url: https://foreversixty.gg/sim/weights
    kind: site
  - label: 'Forever Sixty rotation data, data/curated/apl/warrior-arms.json'
    url: https://foreversixty.gg/sim/specs
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'This site, character planner'
    url: https://foreversixty.gg/planner
    kind: site
---

## Overview

Arms is Forever's two-handed warrior spec, built around Mortal Strike, a hard-hitting attack that also cuts the target's incoming healing.

### Strengths

- Mortal Strike is the highest damage-per-Rage button in the kit
- Steady single-target damage
- Rend's bleed is amplified by Improved Rend and Deep Wounds
- Overpower is a free follow-up after the target dodges

### Weaknesses

- Rage limits the whole rotation and needs Bloodrage to top up
- Little area damage next to Fury's Whirlwind
- Its main hit depends on a six-second cooldown

### In a raid

- **Role:** Two-handed melee damage
- **Brings:** Mortal Strike's healing debuff on the target and Battle Shout for the group

## Talents and builds

- **Improved Rend** — up to 35% more Bleed damage from Rend at max rank; take it early since it costs nothing to maintain once the bleed is applied.
- **Deep Wounds** (requires 3 points in Improved Rend) — your own critical strikes add a second bleed dealing 60% of your weapon's average damage over 12 seconds, stacking with Rend's.
- **Two-Handed Weapon Specialization** — up to 3% more damage with the two-handed weapon Arms is built around.
- **Impale** — up to 20% more critical strike damage on your abilities, multiplying whatever crit chance you already have.
- **Improved Overpower** — up to 50% more critical strike chance on Overpower, the free attack that becomes available after the target dodges you.
- **Sweeping Strikes**, then **Mortal Strike** — Sweeping Strikes is a one-point prerequisite for Mortal Strike at the bottom of the tree in this build's data, and it also cleaves your next 5 swings onto a second target, so it isn't a wasted point even outside its role as a stepping stone. Mortal Strike itself deals weapon damage plus 85 and reduces healing the target receives by 50% for 10 seconds.
- **Bloodthrill**, taken to 5/5 — while Rend is ticking on the target, each Main Hand hit has a 20% chance (at max rank) to let you use Overpower for the next 6 seconds, outside its usual dodge-only window. Overpower already comes free on every dodge in this rotation; Bloodthrill is what makes it come up on plain hits too, and it was the single largest gain this build's level-60 search found for any point in the Arms tree.
- **Weaponmaster**, taken to 5/5 — a weapon-dependent bonus; with the sword in this build's gear set it gives successful melee attacks a 5% chance to trigger an extra attack, and the sim measured it as the next-largest Arms gain after Bloodthrill.

A build reaching Mortal Strike spends 36 points in Arms, with the rest as Cruelty and Unbridled Wrath in Fury (crit and extra Rage) and the five points of Iron Will in Protection, the defensive filler that replaced Toughness in the live tree. Exact splits vary by preference and aren't fixed by anything confirmed for Forever. The extra Bloodthrill and Weaponmaster ranks come from points the previous build spent where the engine measures no damage at all: two of the five Deflection points and all five of Shield Specialization, a defensive talent with no effect on a damage rotation. Nothing the rotation casts was given up: Mortal Strike, Overpower, Slam and Whirlwind all still fire at level 60, and every talent the sim does not model (Deflection, Improved Charge, Improved Tactical Mastery, Sweeping Strikes, Iron Will) is still taken. This build measures about +10.5% over the previous spend in our level-60 search run. Open the planner at [/planner?class=warrior](/planner?class=warrior) to build this out.

**Raid build.** For a raid the search found about +1% over the leveling build above in the raid-ready run: it moves a Deflection point, the two Improved Charge points and the five Iron Will points (resistance to stuns and fears, which matters in PvP, not on a raid boss) into five Booming Voice, Improved Execute and two Improved Bloodrage. Mortal Strike stays.

## Rotation and priority

This site's own Arms priority list opens by topping up Rage with Bloodrage whenever it drops below 80, since Rage is the resource limiting everything else in the rotation. Keep Battle Shout active for the group buff, and keep Rend applied to the target — its bleed is what Improved Rend and Deep Wounds both amplify. Mortal Strike is the priority hit whenever its six-second cooldown is up, since it's the highest damage-per-Rage button in the kit. Overpower is a free follow-up whenever the target dodges you, so it's cast on sight rather than held for later.

## Stat priority

The table above is this level range's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this level range's sample size, to tell its true value apart from zero.

## Best in slot

Look for pieces that lead with attack power and strength, then critical strike, then hit, following the order above — a two-handed weapon with a high top-end damage roll matters more for Arms than it does for the dual-wielding specs.

## Enchants and consumables

Target strength and attack power on weapon and gear enchants, matching the stat priority above.

## Races

Arms can be played by every race the Warrior class allows: Human, Dwarf, Night Elf, and Gnome on the Alliance; Orc, Undead, Tauren, and Troll on the Horde; and, new in Forever, the Skyborne on either faction, since Skyborne is a new race rather than one of Blizzard's six new class pairings. For Alliance, **Human** is a solid pick for a two-handed build: Sword Specialization adds 2% ability critical strike chance while wielding a sword (Warrior has no spells, so only the ability half of that racial applies), and Will to Survive removes all active Stun effects on a 3-minute cooldown, which keeps Mortal Strike's cooldown moving through fights with heavy crowd control. For Horde, **Orc** is the strongest pick: Blood Fury grants 10% attack power for 15 seconds on a 2-minute cooldown, a straightforward damage cooldown that lines up with Recklessness, and Axe Specialization adds 1% crit chance if you're wielding an axe. Skyborne is open to Warrior on both factions with no restriction; its racials (a downward glide, a short full health-and-mana restore, 1% haste) read as utility rather than damage, so it's a valid but not obviously stronger pick than Human or Orc for Arms specifically.

## Professions

Blacksmithing paired with Mining is the standard Classic-era convention for a plate melee DPS class: Blacksmithing crafts weapons and armor and adds extra sockets to gear, and Mining supplies the ore to use it without buying materials.

## Leveling

Fury is generally the stronger leveling choice for Warrior, since dual-wielding and Bloodthirst clear trash faster than Arms's slower two-handed swing timer. Arms levels fine on its own, particularly once Mortal Strike is trained.
