---
title: Shadow Priest in Forever
classSlug: priest
spec: shadow
role: dps
build: 'FS1:1.60.1.70009:priest:gnome:5240110313/0/543120301201300051:'
raidBuild: 'FS1:1.60.1.70009:priest:gnome:3250010313/0/523120501201300251:'
recommendedRaces: [gnome, undead]
statPriority: [Intellect, Hit, Spell power, Shadow power, Critical strike, Spell haste, Spell penetration]
description: 'Shadow Priest overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
updated: 2026-09-24
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Talents Forever (demo transcription)'
    url: https://talentsforever.com/data.json
    kind: community
  - label: 'This site, stat weights'
    url: https://foreversixty.gg/sim/weights
    kind: site
  - label: 'This site, character planner'
    url: https://foreversixty.gg/planner
    kind: site
---

## Overview

Shadow is Priest's damage-over-time and mind-magic DPS spec, built around maintaining Shadow Word: Pain and layering Mind Blast and Mind Flay on top of it, with a small amount of self-sufficiency from Vampiric Embrace-style healing. It's also the class's strongest solo and leveling spec, since its damage kit doubles as sustain. The beta caps at level 30, so how Shadow's damage output compares to other casters at level 60, or performs in a raid, is a projection from the demo talent trees and 1.12 knowledge, not confirmed play.

## Talents and builds

In rough priority order: **Mind Flay** (1 point) is a required pick, not an optional one — Forever moved it out of the baseline trainer-taught kit and into this talent, so without it there is no filler spell to fill the GCDs between cooldowns at all. **Improved Mind Flay**, ranked further down the tree, is its scaling once the 1-point prerequisite is in. **Shadow Weaving** stacks a Shadow damage buff from spell crits, a compounding multiplier for a spec casting Shadow spells constantly. **Shadow Focus**, now taken at its full 5/5, improves your chance to hit with Shadow spells by 5% — a talent search found this a stronger pick than the fifth point in Improved Mind Blast below. **Darkness** adds a flat percentage to all Shadow damage, a broad multiplier late in the tree. **Shadowform**, the tree's capstone, increases Shadow damage by 10% and reduces the mana cost of Shadow spells, at the cost of being unable to cast non-Shadow spells while it's active.

**Improved Mind Blast** still cuts Mind Blast's cooldown and now holds 3 of its 5 points, down from 4. The engine measures each of those points at no damage in this build, so the point moves to **Mental Agility**, taken at 3/3 (a mana-cost cut on Smite, Holy Fire and instant casts, measured by the engine as a real damage gain). **Silent Resolve**'s point (a threat talent, zero damage) and both points in **Early Demise** go to the same place, and the freed Shadow points go into a second rank of **Improved Shadow Word: Pain** (6 extra seconds of Shadow Word: Pain duration), which the probe credits as a clear damage gain. Early Demise, which raised Shadow Word: Death's critical strike chance below 20% health, is no longer taken: the engine measures it at zero damage, and Shadow Word: Death is still cast in the execute window. **Power in Light**, **Holy Precision**, **Improved Power Word: Shield** and **Blackout** hold their ranks: the engine has no code path for them, so this site cannot measure what they are worth and the build leaves them alone. This build measures about +6.9% over the previous spend in our level-60 search run.

Point allocation runs deep into Shadow to reach Shadowform at the bottom of the tree, 33 points, with the remaining 18 in Discipline: **Inner Focus** (1 point) is a required pick alongside the Meditation talent's mana regeneration while casting — it's the free, empowered Devouring Plague cast this build's Rotation section calls out below, so it has to be taken for that line to be legal, not just Meditation. This spend is still a projection — the beta cap of 30 has not let anyone test it live, only this site's own simulator. Open the planner at [/planner?class=priest](/planner?class=priest) to build this out.

**Raid build.** For a raid the search found about +2% over the leveling build above in the raid-ready run: it moves the Discipline points that only help before Shadowform (Power in Light and Holy Precision, which strengthen Holy spells) and Blackout (a stun chance) into Twin Disciplines, Early Demise and two ranks of Improved Mind Blast. Shadow Affinity (less threat), Improved Power Word: Shield, Inner Focus, Vampiric Embrace and Shadowform stay.

## Rotation and priority

The core loop this site's simulator plays: keep Shadow Word: Pain up on the target unless the fight is close to ending, since refreshing it late would waste most of the remaining damage-over-time value. Below 20% target health, Shadow Word: Death is the priority (it is the execute spell, though this build no longer takes Early Demise's bonus critical strike chance on it), falling back to Mind Blast on cooldown and Mind Flay otherwise on the rare global where Shadow Word: Death is still on its own 15-second cooldown inside that window. Outside the execute window, Inner Focus into Devouring Plague is used as a free, empowered cast of the spec's biggest nuke, at its full top rank. After that, Mind Flay on every free global comes before Mind Blast: swapping them was the largest gain the search found (about +9% in our level-60 search run), so Mind Blast outside the execute window is only a fallback for when Mind Flay cannot be cast. Mind Flay needs its talent taken, since Forever gates it behind that pick rather than granting it as a baseline trainer spell (see Talents and builds below for when to take it while leveling). Wand shots fill any global where nothing else is ready.

## Stat priority

The table above is this band's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this band's sample size, to tell its true value apart from zero. These numbers come from this site's own level-60 simulator, not from beta play, which only reaches level 30.

## Gear

Prioritize spell power and intellect first, then crit and hit, then spell haste, spell penetration, and Shadow-specific power. Because bonus healing gear in Forever now also carries a fraction of bonus damage, some cloth pieces itemized for healers may carry incidental value for Shadow that they wouldn't have in 1.12 — worth checking case by case rather than assuming. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30 and this build's raid loot tables are still missing most items the community's sourcing expects, with nothing raiding in-game until the first tier opens on 9 December 2026.

## Enchants and consumables

Weapon enchants should target spell power; **Enchant Weapon - Spell Power** is a verified entry in this build's enchant data. For consumables, **Elixir of Shadow Power** and **Greater Arcane Elixir** are both verified options tied historically to Shadow and spell damage respectively. A fuller consumable stack isn't something this site is naming until the class is played past level 30.

## Races

Shadow can be played by Human, Dwarf, Night Elf, and Gnome on the Alliance, and Undead and Troll on the Horde. Priest is not available to the Skyborne on either faction. For Alliance, **Gnome** is a strong pick for a caster DPS spec: Eureka! makes the next three abilities cost 50% less mana and deal 10% more damage on a 2-minute cooldown, functioning as a burst-and-sustain cooldown that lines up naturally with Mind Blast and Devouring Plague casts. For Horde, **Undead** is the strongest pick: Dark Sacrifice's health-to-mana conversion gives Shadow a way to keep casting through extended fights or pulls without downtime, and Touch of the Grave adds passive chance-on-hit damage that stacks with the spec's DoT-and-nuke kit.

## Professions

Community convention for Shadow Priests favors **Tailoring** for caster cloth stats alongside **Enchanting** for weapon and gear enchants, the same combination general caster practice recommends across classes. This is Classic-era community convention, not confirmed for Forever.

## Leveling

Shadow is the strongest leveling spec for Priest by a wide margin: its damage-over-time kit provides both offense and self-healing that Discipline and Holy, built around supporting other players, don't have solo. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested past that point in Forever.

Take **Mind Flay**'s single point the moment it's available at level 20, ahead of the level-60 build's own point order above. That order maxes Shadow Focus, Blackout and Spirit Tap first for a level-60 raider's hit and threat needs, but a leveling character walking those same points in sequence doesn't reach Mind Flay's row until several levels later — and until then there is no ranged filler at all beyond the wand, since Forever moved Mind Flay out of the baseline kit. Grab the 1-point prerequisite first, then fill Shadow Focus/Blackout/Spirit Tap and the rest of the tree behind it.
