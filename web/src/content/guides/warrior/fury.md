---
title: Fury Warrior in Forever
classSlug: warrior
spec: fury
role: dps
description: 'Talents, rotation, stat priority, and race picks for Fury Warrior in Forever, with beta-versus-projection called out.'
updated: 2026-09-24
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Talents Forever (demo transcription)'
    url: https://talentsforever.com/data.json
    kind: community
  - label: 'Warcraft Tavern: Classic Fury Warrior DPS rotation'
    url: https://www.warcrafttavern.com/wow-classic/guides/dps-warrior-rotation/
    kind: community
  - label: 'Icy Veins, the racial rework'
    url: https://www.icy-veins.com/wow-forever/news/the-racial-rework-gives-every-race-a-reason-to-matter-in-warcraft-forever/
    kind: community
  - label: 'Forever Sixty stat weights, data/curated/specs.json'
    url: https://foreversixty.gg/sim/weights
    kind: site
  - label: 'Forever Sixty rotation data, data/curated/apl/warrior-fury.json'
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

Fury is Forever's dual-wield DPS spec, built around Bloodthirst's frequent big hit and Whirlwind's multi-target filler rather than Arms's single hard-hitting finisher. In a raid or group it plays a similar single-target damage role to Arms but leans on faster, more consistent swings from two weapons instead of one heavy two-handed weapon, and it picks up more incidental area damage from Whirlwind hitting everything nearby. The beta only reaches level 30, so nothing below about level 60 play, raiding, or best-in-slot gear has actually been tested — it's a projection from the demo trees and 1.12 knowledge, not confirmed play.

## Talents and builds

Verified against this build's talent data, in roughly the order you'd take them:

- **Cruelty** — up to 5% flat critical strike chance with melee attacks at max rank, a simple multiplier taken early.
- **Dual Wield Specialization** — up to 25% more off-hand weapon damage, 100% more off-hand Rage generation, and 10% more off-hand hit chance at max rank; the core talent for a two-weapon build.
- **Enrage** — up to a 30% chance to deal 10% bonus Physical damage for 12 seconds after taking any damaging hit, which comes up often on a spec that's usually in melee range.
- **Flurry** (requires 5 points in Enrage) — up to 25% more melee attack speed for your next 3 swings after a melee crit, compounding with Cruelty's crit chance.
- **Death Wish** — the spec's burst cooldown: 20% more Physical damage and Fear immunity, at the cost of 5% more damage taken, for 30 seconds.
- **Bloodthirst** (requires 1 point in Death Wish) — the capstone: an instant attack for damage equal to 35% of Attack Power plus 30, and a 10% movement speed bonus for 10 seconds, the highest damage-per-Rage button in the kit.

Also worth a point if you lean into Whirlwind uptime: **Raging Blows**, a new Forever talent that makes Whirlwind also strike with your off-hand weapon and reduces Cleave's Rage cost. It's a single-point situational pick rather than a core damage multiplier, so it isn't in the priority list above, but it's real in this build's data.

A build reaching Bloodthirst spends roughly 31 points in Fury, with the remainder commonly split into a few Arms points for Deep Wounds and Impale — Fury's high crit rate keeps both the bleed and the crit-damage talent relevant — rather than Protection, which has little to offer a dual-wielding damage build. Open the planner at [/planner?class=warrior](/planner?class=warrior) to build this out.

## Rotation and priority

This site's own Fury priority list opens the same way Arms does: top up Rage with Bloodrage below 80, since Rage limits the rest of the rotation. Death Wish is used on cooldown as the spec's burst window, and Battle Shout is kept active for the group buff. Bloodthirst is the priority hit whenever its six-second cooldown is up, since it's the highest damage-per-Rage button in the kit; Whirlwind fills the gap whenever Bloodthirst isn't ready, hitting everything nearby rather than sitting idle. Below 20% target health, Execute replaces the rest of the priority. Leftover Rage goes into Heroic Strike, but only above 40 Rage, so it never starves Bloodthirst or Execute of what they need.

## Stat priority

In this site's own simulator weighting, in priority order:

1. **Attack power** — the reference stat; every weapon swing and special attack scales off it directly.
2. **Strength** — the primary source of attack power, split between two weapons instead of one.
3. **Agility** — a smaller attack power contribution plus some crit and dodge.
4. **Critical strike** — feeds Flurry's attack-speed proc and Enrage's damage proc, both of which trigger off hits landing, so it does double duty for Fury specifically.
5. **Hit** — keeps Bloodthirst and Whirlwind landing reliably; matters more before you're near the hit cap.
6. **Melee haste** — more swings mean more chances at Flurry and Enrage procs, though this site's own weighting still ranks it last of the six.

## Gear

Look for attack power and strength first, then crit, then hit, following the order above. Because Fury dual-wields, both main-hand and off-hand weapon damage matter, unlike Arms's single two-handed weapon. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30, and this build's raid loot tables are incomplete until the first raid tier opens on 9 December 2026 and gets curated. This section fills in once raid loot is itemized.

## Enchants and consumables

Target strength and attack power on weapon and gear enchants, matching the stat priority above; because Fury carries two weapons, both need an enchant to get full value. **Enchant Weapon - Strength** and **Enchant Weapon - Crusader** (a proc-based strength and self-healing enchant) are both verified in this build's enchant data. For consumables, **Mighty Rage Potion** and **Elixir of Ogre Strength** are verified options in this build's consumable list. Beyond naming those, specific best-in-slot consumable stacking isn't something this site can confirm yet at level 30.

## Races

Fury can be played by every race the Warrior class allows: Human, Dwarf, Night Elf, and Gnome on the Alliance; Orc, Undead, Tauren, and Troll on the Horde; and, new in Forever, the Skyborne on either faction, since Skyborne is a new race rather than one of Blizzard's six new class pairings. For Alliance, **Human** is a strong pick if you're dual-wielding swords: Sword Specialization adds 2% ability critical strike chance while wielding a sword, stacking with Cruelty and giving Flurry more chances to trigger. For Horde, **Troll** pairs well through Berserking, a 10% attack speed cooldown on a 3-minute timer — a direct swing-speed boost for a dual-wield spec that already leans on attack speed for Flurry uptime. Skyborne is open to Warrior on both factions with no restriction; its racials (a downward glide, a short full health-and-mana restore, 1% haste) read as utility rather than damage, so it's a valid but not obviously stronger pick than Human or Troll for Fury specifically.

## Professions

Blacksmithing and Mining remain the standard Classic-era combination for a melee DPS Warrior: Blacksmithing's crafted weapons, armor, and socket additions have direct combat value, and Mining removes the need to buy ore. This is general Classic-community convention, not confirmed for Forever specifically.

## Leveling

Fury is generally considered the strongest leveling spec for Warrior of the three, since dual-wielding two fast weapons together with Bloodthirst clears trash quickly. That comparison is carried over from 1.12 knowledge rather than tested — the beta only reaches level 30, so nothing about leveling above that has actually been played in Forever's client.
