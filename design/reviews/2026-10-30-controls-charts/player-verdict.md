# Player verdict: controls and charts (2026-10-30)

Reviewer: veteran Classic raider (Warcraft Logs daily). Read the spec and all six boards.

## Verdict: HOLD

The dropdowns and the interaction model (crosshair, tooltip, brush, drag-zoom, legend Only/Show all) are what raiders already know from Warcraft Logs and Details!. The charts do not yet answer the three post-wipe questions, and several numbers on the boards do not add up.

## Blockers

1. The threat chart cannot say who pulled threat, yet it is titled "Threat on Ragnaros" and shows "100%" / "98%" shares. The lines are cumulative damage and healing built (the note under the chart admits "the base threat model does not [match a taunt]"). Classic threat is not that: Defensive Stance, Righteous Fury and bear form multipliers, Sunder, Heroic Strike and Wrath-of-Ragnaros target swaps all change who holds aggro. A raider reads "Garrosk 100%" as "Garrosk had aggro". Fix: title it for what it is ("Estimated threat" or "Damage and healing built"), drop or relabel the share column ("of top"), and mark the moments the boss changed target if the log has them. The 2:06-2:07 story (Brokk, the top line at 65.4k, dies and Thandril taunts) is exactly what a player wants and the chart only implies it.
2. "Who died and why" is not answered. The death marker's tooltip is `Death · {name}` and §B6.6 says "A marker is never a link". No killing blow, spell, overkill, or jump to the Deaths tab. At 2:07 the board's tooltip shows only the taunt, not Brokk's death one second earlier. Fix: the death event line carries the killer and spell (`Brokk · Wrath of Ragnaros 3,100`), and the marker opens that death on the Deaths tab.
3. "Who was low on damage at which moment" is defeated by the top-8 rule. Lines are chosen by highest total, so the low performers are by definition the ones hidden, and Classic raids are 40 players, not the 10 on the board. Eight noisy 5 s-average lines also cannot be compared by eye. Fix: let the player pick who is drawn (the Source list already exists), always add the Source player and the viewer's own character, and show the raid-median or same-class line to compare against. Keep top 8 only as a default, and say so in the legend ("Top 8 of 40").
4. The threat chart's default lines can hide the tank. Top 8 by threat total will often drop the tank in a 40-man raid, and the taunt marker then refers to a line that is not on screen. Always include players who have a taunt marker in the window, and the Source player.
5. Board numbers contradict each other. Damage board: head says "Raid 4,226 DPS" but the ten chips sum to 4,387 (8 shown = 4,317). Threat board: the legend leads with Brokk 65.4k but the table under it starts at Garrosk 54,771 = 100%, with no Brokk row. The mock's fight is also not Classic-plausible: Ragnaros submerges at 3:00 for about 90 s and cannot be damaged, but the damage lines run at full strength from 3:15 to the kill at 4:12. Fix the generator so the numbers add up and the submerge reads as a gap.

## Shoulds

- The phone board draws "Damage done · Show chart" under the threat chart, which contradicts B1.6 (one chart per tab; the main chart is not drawn on Threat).
- Healers (Seraphine, Oakhorn) are off "by default (outside the top 8)", but on a heal-focused question they are the whole story. Healing tab must default to healers, which the spec implies; say it.
- Cooldown lane only shows raid cooldowns (Death Wish, Innervate). Add Rebirth, Soulstone, Power Infusion, Bloodlust and Mana Tide where the data has them; those are the casts a raider looks for beside a death. Wipe-relevant consumables (Major Healing Potion, Healthstone) would be a bonus.
- Show the death line in the chart for every death, not just in the lane; at 40 players that collapses into "3+ markers in 44 px". Confirm a wipe (dozens of deaths in 5 s) renders as one counted circle with a readable list, and add a wipe board.
- Source list: the "Damage" group name next to "Tanks" and "Healers" is fine. In the Fight style list the top option reads "Patchwerk", good. Group heads and class colours read right.
- Tooltip share column is "% of top value", which on the damage tooltip is not shown; keep it threat-only and label the column.
- The `Log scale` switch is a Retail-WCL idea few Classic players use; fine, but keep it out of the head's first read (put it after the Window select).
- Tooltip "Died 2:06" row is a good touch; keep it, and also dim that player's line after death.
- Phase names: "PHASE 1" and "SUBMERGE" mix a number and a boss mechanic. Use boss-specific names from the encounter data (Rag: "Ragnaros", "Submerge", "Sons of Flame").

## Re-check (fix round 1)

Verdict: SHIP. All five of my blockers are closed, and I confirmed the two UX blockers on the boards and in the spec. No blockers remain.

1. Threat labelling: closed. The title reads "ESTIMATED THREAT ON RAGNAROS", the share column is "Share of estimated threat", and the foot line reads "An estimate: 1 threat per point of damage and 0.5 per point of healing... No stance, form, talent or taunt modifiers yet." with a link to the glossary. The target-change marker ("Ragnaros turned to Thandril") sits beside the taunt at 2:07. The tanks reading at about a third of the top damage dealers is explained in B12.
2. Death tooltip: closed. The tooltip reads "Brokk died · Melee from Ragnaros · 3,862 (1,104 overkill)" at 2:06 and the 2:07 tooltip carries it. B6.6 makes a death marker a link to the Deaths tab, and a cluster links with the window set to its span. The wipe board counts 14 deaths in one circle and its list links to Deaths.
3. Who is drawn: closed. The field reads "Top 8 of 40". The legend holds the raid median (dashed), YOU (Lightningbolltz at 310) and the 8, and the Players multi-select offers a same-class "Shaman median" and every player. The sheet on the phone does the same.
4. Tanks always drawn: closed. Thandril and Brokk are drawn with a pin on the threat board even though they are outside the top 8, and the taunt marker has its line.
5. Numbers: closed. Legend Garrosk 120.6k equals table 120,628, and "34 more rows" plus 6 makes 40. The damage per-player figures (Velyn 157,916 / 340 s = 464) match the chips. The submerge runs 3:00 to 4:30 at zero, named "SUBMERGE", which is right for Ragnaros.
6. Phone fold: closed. The fold stack totals 581 px of the 640, with the brush inside it.
7. Searchable focus: closed. A6.1 gives one rule: the plain list keeps focus on the trigger, and the searchable list moves focus to the filter input and returns it on close.

Non-blocking, for build:
- The Empty board (3:20 to 4:10) shows no legend row although its caption says the legend row stays.
- Submerge is the window where the raid kills Sons of Flame. "No damage in this window." and "Submerge · no damage" should say "no damage to Ragnaros", and ideally offer the Enemy switch to Sons of Flame.
- The threat tooltip covers the plot's left half at 2:07 (the cursor is at 2:07 of a 1:30 to 2:40 window), which is correct under the 45/55 rule but hides the lines you came to read; accept it.
