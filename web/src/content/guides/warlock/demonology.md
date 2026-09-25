---
title: Demonology Warlock in Forever
classSlug: warlock
spec: demonology
role: dps
description: 'Talents, rotation, stats, and gear for Demonology Warlock in Forever, and what is confirmed versus projected from the beta.'
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
  - label: 'Forever Sixty rotation data, data/curated/apl/warlock-demonology.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Demonology is Warlock's pet-focused spec: its talents strengthen and empower a summoned demon, and its signature choice — Demonic Sacrifice, giving up the pet for a standing buff instead — decides whether you play with a demon out or without one. Casting-wise it's closer to Affliction than a pet-and-forget spec: Corruption and Shadow Bolt still carry most of the direct damage, with Soul Fire as an occasional cooldown. Demonology also picked up Incubus, a new demon summon alongside the Succubus with several talents referencing it, in the BlizzCon demo. The beta only reaches level 30, so nothing here about level 60 talents, raiding, or best-in-slot gear has actually been played — it is a projection from the demo trees and 1.12 knowledge, not tested content.

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31. Verified against this build's talent data, Demonology's key picks in roughly the order you'd take them are:

- **Fel Vitality** — up to 15% more maximum health, mana, and pet stats at rank 3, an early quality-of-life pick.
- **Demonic Sacrifice** — sacrifices your active demon for a long-lasting buff instead, the talent that defines whether you're playing pet-out or pet-sacrificed Demonology.
- **Master Summoner** — up to 4 seconds off summon cast time and 40% off its mana cost at rank 2, useful whenever you need to re-summon mid-fight.
- **Demonic Knowledge** — up to 100% of your level added to your and your pet's spell damage while a demon is active, one of the tree's largest damage talents.
- **Master Demonologist** — a scaling buff to both Warlock and demon that differs by pet, up to 10% at rank 5 for whichever pet is out.
- **Demonic Pact** — the capstone: your Demonic Sacrifice buff is no longer cancelled by summoning a different pet, only by re-summoning the sacrificed one.

A typical Demonology build spends roughly 31 points in this tree, with the remaining 20 points usually going into Affliction for Corruption and Suppression — a common 1.12 hybrid pattern that likely still applies, though it isn't confirmed for Forever specifically. Open the planner at [/planner?class=warlock](/planner?class=warlock) to build this out.

## Rotation and priority

Pop a mana potion at the start of the fight if your mana is already below its usual threshold. Keep Corruption running throughout — it's cheap, sustained damage that doesn't compete with your Shadow Bolt casts for a global cooldown. Soul Fire goes out on cooldown whenever it's up; in a full build it's meant to synergize with Decimation, which drops its cooldown to near nothing after a Shadow Bolt or Searing Pain crit, but this site's simulator hasn't implemented that shortened cooldown yet, so today Soul Fire simply fires about once a minute on its own timer. Shadow Bolt is the default filler for every other global, and it's also what feeds the Decimation proc once that talent is live. Life Tap whenever mana drops below about 10% to keep the rotation running.

## Stat priority

In this site's own simulator weighting, in priority order:

1. **Spell power** — the reference stat; every Corruption tick and Shadow Bolt scales off it directly.
2. **Intellect** — a larger mana pool, and Fel Vitality adds further to both the Warlock's and the pet's pools.
3. **Critical strike** — feeds Soul Fire and Shadow Bolt crits, and eventually the Decimation proc once implemented.
4. **Hit** — needed to stop missing casts against raid-level bosses.
5. **Spell haste** — shortens Shadow Bolt's cast time, the spec's main filler.
6. **Spell penetration** — only matters against targets with meaningful shadow or fire resistance.
7. **Shadow damage and Fire damage (school power)** — the narrowest stats, appearing on very few items, relevant depending on which pet talents you've taken.

## Gear

Look for pieces that lead with spell power, then hit until capped, then crit, following the order above; stamina from Demonic Embrace-adjacent gear is a secondary consideration if you're sacrificing your pet and losing its damage-soak utility. Specific slot-by-slot picks can't be named honestly yet: the beta caps at level 30, so no one has gear-checked Demonology past the early game, and raid loot doesn't exist yet either — nothing raids until 9 December, and even then this client's item table is missing most of the classic raid loot tables it's meant to carry (Onyxia's Lair's own list is currently empty, and Barrow Deeps and Hyjal Summit aren't mapped to items at all). This section will fill in with real picks once raid loot is itemized.

## Enchants and consumables

Target spell power and hit on enchants and consumables, matching the stat priority above. **Enchant Weapon - Spell Power** (verified in this build's enchant data, +30 spell damage) is the standard weapon enchant for a caster. Demonology splits its school damage between Shadow and Fire depending on pet talents taken, so a generic spell-power consumable such as **Flask of Supreme Power** is the safer pick over a single school-specific elixir like **Elixir of Shadow Power** or **Elixir of Fire Power**, both verified in this build's data. Beyond naming those, keep it principle-level until raid-tier consumables are confirmed for Forever specifically.

## Races

Warlock's playable races are Human, Orc, Undead, Gnome, and Troll on their respective factions; the class is not open to Skyborne at all. Troll is one of Blizzard's six confirmed new race and class pairs, giving the Horde a Warlock option it didn't have in 1.12.

For Alliance, **Gnome** is the strongest pick: Expansive Mind's mana bonus stretches a pool that also has to sustain your pet's mana through Fel Vitality, and Eureka!'s mana-cost reduction is a direct complement to re-summoning demons or firing off Soul Fire. For Horde, **Troll** is the pick as the new pairing Blizzard specifically added for this class: Berserking's 10% casting and attack speed on a 3-minute cooldown shortens Shadow Bolt casts, and Regeneration keeps some health regeneration running in combat, useful for a spec that spends health through Life Tap and, if pet-sacrificed, loses the Voidwalker's mana-return option.

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them. This is general Classic-community convention rather than anything Forever-specific, so treat it as inferred until this site has profession data from the beta.

## Leveling

Affliction is generally the stronger leveling choice for Warlock; Demonology can level well too, particularly with a Voidwalker tanking damage while you cast, but it's usually considered slower solo than Affliction's DoT-and-move playstyle. That assessment is inferred from 1.12 knowledge, not tested — the beta only runs to level 30, so nothing about the leveling experience above that has been played in Forever's actual client.
