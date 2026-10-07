---
title: Arms Warrior in Forever
classSlug: warrior
spec: arms
role: dps
build: 'FS1:1.60.1.70009:warrior:human:03325213032515001/0505/005:'
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

Arms is Forever's two-handed weapon DPS spec, built around Mortal Strike: a hard-hitting finisher that also cuts the target's incoming healing for several seconds. In a raid or group it plays a steady single-target damage role rather than Fury's swing-speed-driven output or Protection's threat generation, and its one raid-wide contribution, the Mortal Strike healing debuff, matters most against anything that heals itself or gets healed through the fight. The beta only reaches level 30, so nothing below about level 60 play, raiding, or best-in-slot gear has actually been tested — it's a projection from the demo trees and 1.12 knowledge, not confirmed play.

## Talents and builds

Verified against this build's talent data, in roughly the order you'd take them:

- **Improved Rend** — up to 35% more Bleed damage from Rend at max rank; take it early since it costs nothing to maintain once the bleed is applied.
- **Deep Wounds** (requires 3 points in Improved Rend) — your own critical strikes add a second bleed dealing 60% of your weapon's average damage over 12 seconds, stacking with Rend's.
- **Two-Handed Weapon Specialization** — up to 3% more damage with the two-handed weapon Arms is built around.
- **Impale** — up to 20% more critical strike damage on your abilities, multiplying whatever crit chance you already have.
- **Improved Overpower** — up to 50% more critical strike chance on Overpower, the free attack that becomes available after the target dodges you.
- **Sweeping Strikes**, then **Mortal Strike** — Sweeping Strikes is a one-point prerequisite for Mortal Strike at the bottom of the tree in this build's data, and it also cleaves your next 5 swings onto a second target, so it isn't a wasted point even outside its role as a stepping stone. Mortal Strike itself deals weapon damage plus 85 and reduces healing the target receives by 50% for 10 seconds.
- **Bloodthrill**, taken to 5/5 — while Rend is ticking on the target, each Main Hand hit has a 20% chance (at max rank) to let you use Overpower for the next 6 seconds, outside its usual dodge-only window. Overpower already comes free on every dodge in this rotation; Bloodthrill is what makes it come up on plain hits too, and it was the single largest gain this build's level-60 search found for any point in the Arms tree.
- **Weaponmaster**, taken to 5/5 — a weapon-dependent bonus; with the sword in this build's gear set it gives successful melee attacks a 5% chance to trigger an extra attack, and the sim measured it as the next-largest Arms gain after Bloodthrill.

A build reaching Mortal Strike spends 36 points in Arms, with the rest as Cruelty and Unbridled Wrath in Fury (crit and extra Rage) and the five points of Iron Will in Protection, the defensive filler that replaced Toughness in the live tree. Exact splits vary by preference and aren't fixed by anything confirmed for Forever. The extra Bloodthrill and Weaponmaster ranks come from points the previous build spent where the engine measures no damage at all: two of the five Deflection points and all five of Shield Specialization, a defensive talent with no effect on a damage rotation. Nothing the rotation casts was given up: Mortal Strike, Overpower, Slam and Whirlwind all still fire at level 60, and every talent the sim does not model (Deflection, Improved Charge, Improved Tactical Mastery, Sweeping Strikes, Iron Will) is still taken. This build measures about +10.5% over the previous spend in our level-60 search run. Open the planner at [/planner?class=warrior](/planner?class=warrior) to build this out.

## Rotation and priority

This site's own Arms priority list opens by topping up Rage with Bloodrage whenever it drops below 80, since Rage is the resource limiting everything else in the rotation. Keep Battle Shout active for the group buff, and keep Rend applied to the target — its bleed is what Improved Rend and Deep Wounds both amplify. Mortal Strike is the priority hit whenever its six-second cooldown is up, since it's the highest damage-per-Rage button in the kit. Overpower is a free follow-up whenever the target dodges you, so it's cast on sight rather than held for later. Below 20% target health, Execute replaces the rest of the priority. There is no Heroic Strike rage dump any more: the search found that every point of Rage spent on it was Rage Mortal Strike, Overpower and Execute would have used better, and dropping it was worth about +1% in our level-60 search run. Arms has no burst cooldown in this rotation: Recklessness is Berserker-Stance-only, and Arms correctly stays in Battle Stance the whole fight for Overpower, so it can never be cast without a stance dance this priority list doesn't do; the spec also has no Death Wish-style talent of its own to fill that role.

## Stat priority

The table above is this band's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this band's sample size, to tell its true value apart from zero. These numbers come from this site's own level-60 simulator, not from beta play, which only reaches level 30.

## Gear

Look for pieces that lead with attack power and strength, then critical strike, then hit, following the order above — a two-handed weapon with a high top-end damage roll matters more for Arms than it does for the dual-wielding specs. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30, so nobody has gear-checked Arms past the early game, and this build's raid loot tables are themselves incomplete — nothing raids in-game until the first tier opens on 9 December 2026, and even then several of this client's raid item tables are missing most of what earlier sourcing expects. This section fills in once raid loot is itemized and the spec is played past 30.

## Enchants and consumables

Target strength and attack power on weapon and gear enchants, matching the stat priority above. **Enchant Weapon - Strength** and **Enchant Bracer - Superior Strength** are both verified in this build's enchant data. For consumables, **Mighty Rage Potion** and **Elixir of Ogre Strength** are verified options in this build's consumable list. Beyond naming those, specific best-in-slot consumable stacking isn't something this site can confirm yet at level 30.

## Races

Arms can be played by every race the Warrior class allows: Human, Dwarf, Night Elf, and Gnome on the Alliance; Orc, Undead, Tauren, and Troll on the Horde; and, new in Forever, the Skyborne on either faction, since Skyborne is a new race rather than one of Blizzard's six new class pairings. For Alliance, **Human** is a solid pick for a two-handed build: Sword Specialization adds 2% ability critical strike chance while wielding a sword (Warrior has no spells, so only the ability half of that racial applies), and Will to Survive removes all active Stun effects on a 3-minute cooldown, which keeps Mortal Strike's cooldown moving through fights with heavy crowd control. For Horde, **Orc** is the strongest pick: Blood Fury grants 10% attack power for 15 seconds on a 2-minute cooldown, a straightforward damage cooldown that lines up with Recklessness, and Axe Specialization adds 1% crit chance if you're wielding an axe. Skyborne is open to Warrior on both factions with no restriction; its racials (a downward glide, a short full health-and-mana restore, 1% haste) read as utility rather than damage, so it's a valid but not obviously stronger pick than Human or Orc for Arms specifically.

## Professions

Blacksmithing paired with Mining is the standard Classic-era convention for a plate melee DPS class: Blacksmithing crafts weapons and armor and adds extra sockets to gear, and Mining supplies the ore to use it without buying materials. This is general Classic-community practice rather than anything confirmed for Forever specifically.

## Leveling

Fury is generally the stronger leveling choice for Warrior, since dual-wielding and Bloodthirst clear trash faster than Arms's slower two-handed swing timer. Arms levels fine on its own, particularly once Mortal Strike is trained, but that comparison is carried over from 1.12 leveling knowledge — the beta only runs to level 30, so nothing about the leveling experience above that has actually been played in Forever's client.
