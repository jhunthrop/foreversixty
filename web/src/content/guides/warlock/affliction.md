---
title: Affliction Warlock in Forever
classSlug: warlock
spec: affliction
role: dps
build: 'FS1:1.60.1.70009:warlock:gnome:25552300120201351/200522/003:'
raidBuild: 'FS1:1.60.1.70009:warlock:gnome:25550300100201351/00052/0550001:'
recommendedRaces: [gnome, troll]
statPriority: [Spell power, Shadow power, Intellect, Hit, Critical strike, Spell haste, Spell penetration]
description: 'Talents, rotation, stats, and gear for Affliction Warlock in Forever, and what is confirmed versus projected from the beta.'
updated: 2026-09-24
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
  - label: 'Forever Sixty rotation data, data/curated/apl/warlock-affliction.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Affliction is Warlock's damage-over-time spec: it layers multiple DoTs onto a target and lets them run, spending Life Tap to keep casting once mana runs low rather than worrying about health, since the drain spells and DoTs keep the Warlock topped up. In a raid it's a low-maintenance, high-uptime caster that keeps producing damage even during movement, once its DoTs are applied. The beta only reaches level 30, so nothing here about level 60 talents, raiding, or best-in-slot gear has actually been played — it is a projection from the demo trees and 1.12 knowledge, not tested content.

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31; separately, Curses and Banes now run as independent categories so a Bane and a Curse can be active on the same target together. Verified against this build's talent data, Affliction's key picks in roughly the order you'd take them are:

- **Suppression** — up to 5% more hit chance and 20% less threat at rank 5, both scarce and valuable for a spec that casts constantly.
- **Improved Corruption** — cuts Corruption's cast time by up to 2 seconds and adds up to 10% more damage at rank 5.
- **Amplify Curse** (1 point) — a required pick, not an optional one, and missing from an earlier draft of this guide's talent section: this build's rotation prepull-casts it before the pull, so it has to be taken for that line to be legal.
- **Pandemic** — up to a 100% critical-strike-damage bonus at rank 3 on Corruption, both Banes, and the drain spells. This build takes two of its three points; the search run measured no damage in the last one, so it went elsewhere.
- **Nightfall** — up to a 4% chance per DoT tick at rank 2 to make your next Shadow Bolt instant, a real mana and time saver.
- **Shadow Mastery** — up to 5% more Shadow damage and life drained at rank 5.
- **Wrack** — the capstone: a DoT that also boosts your other Shadow DoTs on the same target by 10% while it's active.

This build spends 37 points in Affliction to reach Wrack at the bottom, 11 in Demonology and 3 in Destruction. Two things changed from the first draft. Improved Drains and Soul Siphon, three points each, join the Affliction side: Soul Siphon was the single largest unspent gain the talent search found, and Improved Drains was next. Three points of Bane in Destruction shorten the Shadow Bolt cast that fills most globals. They came out of Demonic Embrace, Improved Imp and one point of Pandemic, whose effects (pet stamina, pet imp fire damage, and the extra crit bonus) the simulator measures at zero for a solo caster. Together the swap was worth about +14% in the search run, and about +11% on the level-60 row of the rotation ladder.

What the build deliberately keeps is everything the simulator cannot see: Soul Harvest, Improved Health Funnel, Demonic Aegis and Improved Voidwalker stay in Demonology, since their value is unmeasured rather than measured as zero. Improved Life Tap, Amplify Curse, Siphon Life and Nightfall are untouched, so every ability the rotation casts is still available. A deeper Affliction spec with more Destruction points scored higher in the search, but it drops those unmeasured talents, so it is not the recommendation here. Open the planner at [/planner?class=warlock](/planner?class=warlock) to build this out.

**Raid build.** For a raid the search found about +12% over the leveling build above in the raid-ready run: it moves two points of Pandemic, which the simulator measures at no damage, and the pet and solo-play points the engine cannot see (Soul Harvest, Improved Health Funnel and Improved Voidwalker, which serve a leveling warlock's mana and pet) into five Improved Shadow Bolt, two Bane and Ruin. Suppression (less threat), Amplify Curse, Wrack and Demonic Aegis stay.

## Rotation and priority

Pop a mana potion or gem at the start of the fight if your mana is already below its usual threshold, or save it for a closing burn if the fight is nearly over. From there, put Corruption up first and keep it up throughout the fight — it's the cheapest DoT per point of damage, so upkeep on it should never lapse. Bane of Doom goes up right behind it as your one curse: it does all of its damage in a single hit when its minute is up, so recast it the moment it fires. In our simulator it beat Bane of Agony by a small margin and Curse of the Elements by a wide one, since a lone caster's own damage gains less from Elements than Doom or Agony deal themselves; on a fight shorter than a minute Doom never lands, so use Agony there. Wrack is live in this site's simulator: an instant, no-cooldown DoT that also stacks a 10% Shadow-damage-taken vulnerability on the target, so recast it in its last second to keep the vulnerability from ever dropping. Immolate is the fourth DoT: the search found that a running Immolate out-damages the Shadow Bolt global it costs, so it joins the upkeep order ahead of the filler. Once all four DoTs are running, spend every remaining global on Shadow Bolt. Once mana falls under 50%, wand with Shoot instead: it costs nothing, so mana regenerates while the DoTs keep ticking, and Life Tap is only the backstop below about 10%. Together these changes were worth about +6% in our level-60 search run.

## Stat priority

The table above is this band's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this band's sample size, to tell its true value apart from zero. These numbers come from this site's own level-60 simulator, not from beta play, which only reaches level 30.

## Gear

Look for pieces that lead with spell power, then hit until capped, then crit, following the order above. Specific slot-by-slot picks can't be named honestly yet: the beta caps at level 30, so no one has gear-checked Affliction past the early game, and raid loot doesn't exist yet either — nothing raids until 9 December, and even then this client's item table is missing most of the classic raid loot tables it's meant to carry (Onyxia's Lair's own list is currently empty, and Barrow Deeps and Hyjal Summit aren't mapped to items at all). This section will fill in with real picks once raid loot is itemized.

## Enchants and consumables

Target spell power and hit on enchants and consumables, matching the stat priority above. **Enchant Weapon - Spell Power** (verified in this build's enchant data, +30 spell damage) is the standard weapon enchant for a caster. **Elixir of Shadow Power** is a verified school-specific damage elixir in this build's data and is the natural pick for Affliction over a generic spell-power elixir when the two conflict; **Flask of Supreme Power** covers the general spell-damage slot on a raid night. Beyond naming those, keep it principle-level until raid-tier consumables are confirmed for Forever specifically.

## Races

Warlock's playable races are Human, Orc, Undead, Gnome, and Troll on their respective factions; the class is not open to Skyborne at all. Troll is one of Blizzard's six confirmed new race and class pairs, giving the Horde a Warlock option it didn't have in 1.12.

For Alliance, **Gnome** is the strongest pick: Eureka! cuts the mana cost of your next three abilities by 50% and adds 10% more damage on a 2-minute cooldown, which pairs well with re-applying DoTs after a burst window, and Expansive Mind's mana bonus supports a rotation that keeps casting all fight. For Horde, **Troll** is the pick as the new pairing Blizzard specifically added for this class: Berserking grants 10% casting and attack speed on a 3-minute cooldown, which shortens Shadow Bolt casts between DoT refreshes, and Regeneration keeps some health regeneration running in combat, a fit for a spec that spends health through Life Tap.

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them. This is general Classic-community convention rather than anything Forever-specific, so treat it as inferred until this site has profession data from the beta.

## Leveling

Affliction is generally considered the strong leveling choice for Warlock, since its DoTs keep working while you move or fight multiple targets, and Drain Life covers the healing a solo caster otherwise lacks. That assessment is inferred from 1.12 knowledge, not tested — the beta only runs to level 30, so nothing about the leveling experience above that has been played in Forever's actual client.
