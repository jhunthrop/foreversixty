---
title: Frost Mage in Forever
classSlug: mage
spec: frost
role: dps
build: 'FS1:1.60.1.70009:mage:gnome:203005/113023/253510130100030025:'
raidBuild: 'FS1:1.60.1.70009:mage:gnome:203005/13102/255510032100030025:'
recommendedRaces: [gnome, troll]
statPriority: [Hit, Spell power, Frost power, Intellect, Critical strike, Spell haste, Spell penetration]
description: 'Talents, rotation, stats, and gear for Frost Mage in Forever, and what is confirmed versus projected from the beta.'
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
  - label: 'Forever Sixty rotation data, data/curated/apl/mage-frost.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Frost is Mage's control and survivability spec: Frostbolt's slow keeps enemies at range, Frost Nova and Ice Block cover emergencies, and the tree's damage talents lean on the Shatter interaction — bonus crit chance against a frozen target — rather than a burst cooldown. In a raid it trades some of Fire's sustained output for more personal safety and utility; solo and leveling, that safety is the whole appeal. The beta only reaches level 30, so nothing here about level 60 talents, raiding, or best-in-slot gear has actually been played — it is a projection from the demo trees and 1.12 knowledge, not tested content.

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31. Verified against this build's talent data, the Frost tree's key picks in roughly the order you'd take them are:

- **Improved Frostbolt** — up to 0.5 seconds off Frostbolt's cast time at rank 5, the tree's most direct throughput gain.
- **Ice Shards** — up to 100% more Frost critical strike damage at rank 5, the biggest single damage multiplier in the tree.
- **Piercing Ice** — a flat 6% more Frost damage at rank 3.
- **Shatter** — up to 50% more critical strike chance against a frozen target at rank 3. A raid boss cannot be frozen by Frost Nova, so against a boss this talent only works through Fingers of Frost, below.
- **Fingers of Frost** — each Chill effect (Frostbolt, Frostfire Bolt, Cone of Cold, Improved Blizzard) has a 15% chance to grant a 15 second effect that treats your next 1 spell (2 at rank 2) as if the target were Frozen. That switches on Shatter's crit bonus and Ice Lance's frozen damage bonus against a target that Frost Nova cannot freeze. It needs Ice Lance, so that is a one-point cost.
- **Winter's Chill** — a chance for Frost hits to stack a debuff that raises Ice Lance's and Frostbolt's crit chance against that target, up to 5 stacks at rank 5.

The build above comes from this site's talent search, held to the rule that a replacement must keep every talent the simulator cannot yet credit, stay in the Frost tree, and gain at least 1% beyond the run's measured error. It takes **Wand Specialization** at 2/2 (25% more wand damage), **Improved Channeling** at 3/5 (a 42–60% chance to keep channeling Arcane Missiles through incoming damage — the engine has no code path yet to credit that resistance, so it's an honest unknown rather than a confirmed zero) and **Arcane Concentration** at 5/5 (a 10% chance per damage spell cast to enter Clearcasting, making your very next cast free — by far the largest single swing measured in this build). Fire keeps only the points that unlock Impact and Flame Throwing, which the simulator cannot credit but a Frost mage may still want for their control effects.

The latest change is Fingers of Frost, now modeled: the search's previous build had Shatter at 3/3 but nothing to make a boss count as frozen, so Shatter did nothing. Spending the Arcane Blast, Cold Snap, and Improved Frost Nova points on Ice Lance and Fingers of Frost 2/2 lifts the build by about 10% over the previous spend on the Frostbolt-only rotation, and the Ice Lance line below adds roughly 3% on top. Those three dropped talents are utility the simulator measures as no damage, not damage you lose; Cold Snap and Improved Frost Nova are real control tools, and a player who values them more than that gain can take them back. Ice Barrier stays out for the same reason as before: a pure-damage search does not weigh survivability. This spend is still a projection — the beta cap of 30 has not let anyone test it live, only this site's own simulator. Open the planner at [/planner?class=mage](/planner?class=mage) to build this out.

**Raid build.** For a raid the search found about +3% over the leveling build above in the raid-ready run: it moves two ranks of Improved Fireball, the three Impact points (a stun) and one Frostbite (a freeze chance) into Incineration, Elemental Precision and two more ranks of Frost Channeling, which also takes 30% of the threat off Frost spells. Winter's Chill, Ice Lance, Fingers of Frost, Frost Warding, Flame Throwing and Improved Channeling stay.

## Rotation and priority

Frostbolt is the filler, with the wand coming out once mana runs low. When Fingers of Frost procs, cast **Ice Lance** while the effect is up: the held charge makes the target count as frozen, so Ice Lance deals triple damage and Shatter's crit bonus applies, and each cast spends one charge. In the simulator the Ice Lance line gains roughly 3% over Frostbolt alone, beyond the run's error. Frostbolt's own Chill can roll the next charge, so keep casting while charges are held rather than waiting. Frostfire Bolt, which you can train from level 40, is not a Frost filler: swapping it in for Frostbolt measured roughly a third lower damage, because its direct hit is smaller and the Frost talents above scale Frostbolt, not it. Arcane Blast is no longer part of this rotation or this build: the client prices every cast at 15% of your base mana, and the earlier free-cast line rested on an engine that charged nothing for it. Two modelling choices matter here and are unconfirmed against the beta: the cast that resolves while you hold a charge is the one that benefits (as in TBC and WotLK), and Ice Lance's "300% increased damage to Frozen targets" is read as triple damage rather than quadruple. If it turns out to be quadruple, Ice Lance on a charge is worth more than the figure here. Frost Nova and the other freezes are not part of the boss rotation, since a boss cannot be frozen.

## Stat priority

The table above is this band's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this band's sample size, to tell its true value apart from zero. These numbers come from this site's own level-60 simulator, not from beta play, which only reaches level 30.

## Gear

Look for pieces that lead with spell power, then hit until capped, then crit, following the order above. Specific slot-by-slot picks can't be named honestly yet: the beta caps at level 30, so no one has gear-checked Frost past the early game, and raid loot doesn't exist yet either — nothing raids until 9 December, and even then this client's item table is missing most of the classic raid loot tables it's meant to carry (Onyxia's Lair's own list is currently empty, and Barrow Deeps and Hyjal Summit aren't mapped to items at all). This section will fill in with real picks once raid loot is itemized.

## Enchants and consumables

Target spell power and hit on enchants and consumables, matching the stat priority above. **Enchant Weapon - Spell Power** (verified in this build's enchant data, +30 spell damage) is the standard weapon enchant for a caster. **Elixir of Frost Power** is a verified school-specific damage elixir in this build's data and is the natural pick for Frost over a generic spell-power elixir when the two conflict; **Flask of Supreme Power** covers the general spell-damage slot on a raid night. Beyond naming those, keep it principle-level until raid-tier consumables are confirmed for Forever specifically.

## Races

Mage's playable races are Human, Orc, Undead, Gnome, and Troll on their respective factions, plus Skyborne — but only the Alliance-side Skyborne, since Mage is Alliance-only for that race. Orc is one of Blizzard's six confirmed new race and class pairs, giving Horde a Mage option it didn't have in 1.12.

For Alliance, **Gnome** is the strongest pick: Escape Artist breaks and grants brief immunity to movement impairment on a 2-minute cooldown, which is a direct complement to a spec that already relies on kiting and control, and Expansive Mind's mana bonus supports a rotation that runs on continuous Frostbolt casts. For Horde, **Troll** is the pick: Berserking grants 10% casting and attack speed on a 3-minute cooldown, a straightforward way to squeeze more Frostbolts and Ice Lances into a burst window; it's also the new pairing Blizzard specifically added for this class, though for Warlock rather than Mage — Orc remains Mage's new Horde option, and its Blood Fury spell power cooldown is a reasonable alternative pick if Troll's timing doesn't line up with your raid's needs. Skyborne is available to Alliance Mages as a third option; its Read Ley Line racial restores all health and mana over 15 seconds, useful out of combat, but it doesn't add spell damage or control the way Gnome's kit does, so it's a pick for the race and story rather than raw throughput.

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them. This is general Classic-community convention rather than anything Forever-specific, so treat it as inferred until this site has profession data from the beta.

## Leveling

Frost is the strong leveling choice for Mage generally, and that holds for itself: Frostbolt's slow and Frost Nova's root make solo play safer than either Arcane's burst-dependent kit or Fire's closer-range playstyle. That assessment is inferred from 1.12 knowledge, not tested — the beta only runs to level 30, so nothing about the leveling experience above that has been played in Forever's actual client.
