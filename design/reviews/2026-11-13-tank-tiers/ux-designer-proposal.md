# Tank tier list: ux-designer verdict on the proposed fix

Reviewed: /tiers/tank at 1440 and 390 (live captures), tier-row-view.ts, tier-copy.ts, tier-list.ts.
Reference: Wowhead tank tier tables and the Details! damage-taken window: one sorted number, bars that read as share of the top, green only where better.

Verdict: SHIP with changes.

## Item by item

1. Real maximum health. SHIP. It removes the double count: Health now says the bear has the most (9,171) and Taken / s says the mitigation is why it still loses. Highest is green. Head stays "Health", title "Maximum health with this gear and raid buffs", unit HP, phone label "Health" (not "Effective health"). Placeholder: no "n/a" cells. Until every row in the band carries metrics.health, drop the Health column on desktop and the Health line on phone for that band. Never fall back to the old effective-health number under the new label. Whole numbers, grouped.
2. vs #1. SHIP with changes. Whole percent: Taken / s is whole numbers, so 52.9 against 53.2 is noise on a 1 per second gap. Cells read "+53% taken" in the default text colour, never green (a plus that is bad must not look like a win). Leader keeps "best". Title (desktop) and aria-label: "Takes 53% more damage a second than Protection Warrior", leader name from data. A title does not exist on touch, so the visible text must stand alone: it does.
3. Bars. Pick the first option: bar = best / taken, no floor, tank only. Warrior 100%, bear and paladin 65%. A longer bar means better on every tab; the shortest-is-best alternative breaks that and the green leader. The floor at 90% of the lowest row is what turned 0.65 into 0.16.
4. Rail note. SHIP with the copy below.

## Copy, verbatim

Rail "What the number is":
"Taken / s is the damage the tank takes a second from the boss, lower is better, and the list is sorted by it. Health is maximum health with this gear and raid buffs. Threat / s is threat a second while tanking. vs #1 is how much more damage a second than the top spec. The bar is the top spec's damage taken as a share of this spec's. The best of each column is green: lowest taken, highest health, highest threat."

Phone gap label: "More damage taken than #1".

## Other things that mislead on this tab

- MUST: on 390 the vs #1 figure is missing for every tank row. tierFigures returns phoneGap only on the fallback path; the full path returns [sorted, health, threat]. Phone shows no gap at all.
- SHOULD: paladin is ranked 3rd on 352 against 351. That is a tie inside sim error; show both as rank 2 (or a tie mark) as healers do with healerTie.
- SHOULD: the sorted column has no sort mark. Add the same one the other tabs use on the sorted head.
- SHOULD: Damage and Healers keep the 90% floor, so their small gaps look large and the tank bars would now follow a different rule. Out of scope here; log it so the three tabs share one bar rule.
- SHOULD: phone figures carry no unit; the unit only rides in the head. Add HP and /s after the values or the phone label.

## Acceptance captures

1440 and 390, Alliance and Horde: all rows with health; a band without metrics.health (column hidden, no gap, no shift); a tie row. Check bar lengths 100/65/65 by measuring the rendered track.
