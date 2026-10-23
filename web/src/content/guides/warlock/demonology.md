---
title: Demonology Warlock in Forever
classSlug: warlock
spec: demonology
role: dps
build: 'FS1:1.60.1.70532:warlock:gnome:25522/233211310122000135/005:'
raidBuild: 'FS1:1.60.1.70532:warlock:gnome:055/203511311112000135/0050051:'
recommendedRaces: [gnome, troll]
statPriority:
  [Spell power, Hit, Intellect, Critical strike, Fire power, Spell haste, Spell penetration, Shadow power]
description: 'Talents, rotation, stats, and gear for Demonology Warlock in Forever, and what is confirmed versus projected from the beta.'
updated: 2026-10-08
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

Demonology is Warlock's pet spec: its talents empower a summoned demon, and Demonic Sacrifice trades the pet for a standing buff.

### Strengths

- Talents strengthen and empower the summoned demon
- Demonic Sacrifice offers a play style without a pet
- Incubus joins Succubus as a new summon with several talents behind it
- Soul Fire adds an occasional large hit

### Weaknesses

- Corruption and Shadow Bolt still carry most of the direct damage
- The mana pool also sustains the pet through Fel Vitality
- A dead pet must be re-summoned, costing mana and time

### In a raid

- **Role:** Ranged shadow damage with a demon
- **Brings:** A curse on the target (Bane of Doom in its rotation), plus Healthstones and a Soulstone

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31.

- **Fel Vitality** — up to 15% more maximum health, mana, and pet stats at rank 3, an early quality-of-life pick.
- **Demonic Sacrifice** — sacrifices your active demon for a long-lasting buff instead, the talent that defines whether you're playing pet-out or pet-sacrificed Demonology.
- **Master Summoner** — up to 4 seconds off summon cast time and 40% off its mana cost at rank 2, useful whenever you need to re-summon mid-fight.
- **Demonic Knowledge** — up to 100% of your level added to your and your pet's spell damage while a demon is active, one of the tree's largest damage talents.
- **Master Demonologist** — a scaling buff to both Warlock and demon that differs by pet, up to 10% at rank 5 for whichever pet is out.
- **Decimation** — the search run's best use of the points freed from Demonic Embrace; the simulator credits it with a measurable damage gain.

This build spends 30 points in Demonology, with 16 in Affliction (Improved Corruption and Suppression maxed, then Improved Life Tap, Soul Harvest and Grim Reach, the two-rank range talent the current client put in Malediction's slot) and 5 in Destruction on Bane. Against the first draft it drops three Improved Drains points, two Demonic Embrace points and the Demonic Pact capstone, all of which the simulator measures at zero for a solo caster, and spends them on two Decimation points and four Bane points, which shorten the Shadow Bolt cast. That was worth about +10% in the search run, and about +8% on the level-60 row of the rotation ladder. Demonic Sacrifice and Soul Link stay, and so do the talents the simulator cannot yet see (Soul Harvest, Improved Health Funnel, Demonic Aegis, Improved Voidwalker). A deep Affliction build scored higher still, but it abandons the Demonology tree and those talents, so it is not the recommendation for this spec. Open the planner at [/planner?class=warlock](/planner?class=warlock) to build this out.

**Raid build.** For a raid the search found about +6% over the previous raid build in the raid-ready run, now that Searing Pain is the filler and Soul Fire is in the loop. It drops the three Malediction points and a Master Summoner point and spends them on two more Aftermath points, a Ruin point and an Improved Sayaad point. With Searing Pain as the filler Improved Shadow Bolt no longer earns a point, which is why the raid build carries none of it. Suppression (less threat), Soul Link, Demonic Aegis and Decimation stay.

## Rotation and priority

Pop a mana potion at the start of the fight if your mana is already below its usual threshold. Keep Bane of Doom and Corruption running throughout: Doom deals its whole damage when its minute is up, so recast it the moment it fires (it edged out Bane of Agony in our simulator; on a fight shorter than a minute use Agony instead), and Corruption is cheap, sustained damage that doesn't compete with your Shadow Bolt casts for a global. Together these changes were worth about +13% in our level-60 search run.

## Stat priority

The table above is this level range's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this level range's sample size, to tell its true value apart from zero.

## Best in slot

Look for pieces that lead with spell power, then hit until capped, then crit, following the order above; stamina from Demonic Embrace-adjacent gear is a secondary consideration if you're sacrificing your pet and losing its damage-soak utility.

## Enchants and consumables

Target spell power and hit on enchants and consumables, matching the stat priority above.

## Races

Warlock's playable races are Human, Orc, Undead, Gnome, and Troll on their respective factions; the class is not open to Skyborne at all. Troll is one of Blizzard's six confirmed new race and class pairs, giving the Horde a Warlock option it didn't have in 1.12.

For Alliance, **Gnome** is the strongest pick: Expansive Mind's mana bonus stretches a pool that also has to sustain your pet's mana through Fel Vitality, and Eureka!'s mana-cost reduction is a direct complement to re-summoning demons or firing off Soul Fire. For Horde, **Troll** is the pick as the new pairing Blizzard specifically added for this class: Berserking's 10% casting and attack speed on a 3-minute cooldown shortens Shadow Bolt casts, and Regeneration keeps some health regeneration running in combat, useful for a spec that spends health through Life Tap and, if pet-sacrificed, loses the Voidwalker's mana-return option.

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them.

## Leveling

Affliction is generally the stronger leveling choice for Warlock; Demonology can level well too, particularly with a Voidwalker tanking damage while you cast, but it's usually considered slower solo than Affliction's DoT-and-move playstyle.
