# F6 Talents, built: wow-player verdict (2026-10-31)

Seen: iOS and Android running-app captures (Obnoxious Yell, Fury 58), the Elemental goldens, the large-text goldens. Checked against data/builds/1.60.1.70658/talents/{warrior,shaman}.json.

**Question the screen answers:** "Is my build the guide build?" The status tag and sentence are the first thing under the title, above the tree. Good.

## What checks out
- Fury tree: every cell's tier, column and rank matches the data (Booming Voice 0/5, Cruelty 5/5, ... Enrage 2/5, Gore Drinker 0/2 locked, Death Wish to Flurry gold arrow, Bloodthirst 1/1). 17 digits of 05153105022011501 sum to 32, the tab says 32. Arms 17 and the 8 differences (Improved Rend +1, Improved Charge +2, Deep Wounds +3, Impale +2) are true.
- Elemental vector: +2 Reverberation, -2 Call of Flame, -2 Elemental Warding, Tidal Focus -1 so Restoration gets a dot; tabs 28/3/17; 3 unspent; "takes 5 points you have not" is true (2+2+1). Matches board: 30/3/18, no dashes.
- Tier gating: Elemental Mastery dim at 28 points, open at 30. Requirement lines use the game's wording ("Requires 30 points in Elemental Talents", "Requires 5 points in Enrage"). Tooltip texts match the data. Arrowheads, gold-when-held lines, cell borders all read like the game.
- No gear, DPS or sim figure on F6. Footer links /planner.

## Blockers
1. **Home's Talents row contradicts F6 for the same character.** home-talents-row shows Obnoxious Yell, Fury 58: "1 of your 51 points differ from the 50 to 59 guide build". F6 for the same character says "8 of your 49 points". A level 58 has 49 points to spend, so 51 is impossible in the game, and a player who taps the row lands on different numbers. The spec says both use one function. Re-capture from one fixture and prove the row equals F6.
2. **The "large" running-app captures are Home, not F6, and Home breaks.** app-talents-differs-large and -large-tree show the Home row at big text: the title splits as "Tal / ent / s", "poin / ts", "diffe / r", "buil / d" down a one-word-wide column beside a huge DIFFERS pill. That is the way into F6 and a player would screenshot it. The golden talents-differs-large (stacked tabs, wrapped sentence) is fine, but no running-app F6 at large text exists. Fix the Home row wrap and capture F6 itself at large text.

## Shoulds
1. **A level 58 can never read Matches.** Everyone who follows the 50 to 59 guide to the letter and spends their extra 8 points anywhere reads "Differs". The "Level 58 is past the 50 to 59 guide build's 41 points, so points past it show as differing" line is honest, but the headline still frets over a build that has every guide point (here yours is at or above the guide in every cell, missing count 0). Owner ruling wanted: points past the guide's length should not count as differing, or the tag should say something like "Has the guide's points".
2. **Fury opens with no dashes while the tag says Differs.** The Arms dot is the only pointer. It works, but the dot is red for every tree; Restoration's difference is "-1" (guide wants it) and Arms-style "over" is the red case. Make the dot take the chip colour (gold when under, red when over, both: red).
3. **Rank-sheet goldens show no tree.** talents-rank-bottom-row and -right-column are a sheet on an empty screen, so I cannot see that the tapped cell is above the sheet and ringed (spec 7.4). The running capture (Gore Drinker) does show it; add the same for a bottom-row cell on both platforms.
4. **Tier 7 hides.** On Differs the capstone is half under the tab bar on both phones (Bloodthirst). Accepted in the spec, but the glass tab bar lets the legend text bleed through behind it on iOS, which looks messy.
5. **Elemental Mastery 0/1 in the 60 guide.** An Elemental shaman with 30 in Elemental and no capstone, 18 in Restoration, looks wrong to a player. Data question for the accuracy reviewer (already noted in spec item 14).

## Nits
- Gore Drinker sheet is half empty at medium height; fine.
- The guide-line box on a talent that is -2 reads "Guide build: rank 3/3" with no "you are 2 short"; the chip does that, acceptable.

## Missing vs the in-game frame
Nothing blocking. Gold for maxed is green here by design.

(First pass ended HOLD; see Re-check below.)

## Re-check (2026-10-31, after the fix round)

Looked at the new captures, not the report.
- Blocker 1 cleared: home-talents-row-ios now reads "8 of your 49 points differ from the 50 to 59 guide build." with Differs, the same sentence as F6. The impossible 51 is gone.
- Blocker 2 cleared: app-talents-differs-large-ios is F6 at big text (tag above a wrapped sentence, the beyond line, stacked segments, stacked tabs with the Arms dot, Fury selected); -large-tree-and shows the tree with icons at fixed size and the legend wrapping. The Home row no longer shatters.
- New talents-restoration-ios: Tidal Focus 4/5 with the gold "-1" chip and dashed ring, Restoration tab dot, Healing Way to Riptide line. Matches the data (Tidal Focus tier 0 column 2, guide 5, yours 4).
- Shoulds 1 (level 58 reads Differs) and 2 (dots all red) stand as owner rulings; not blocking. Should 3 to 5 unchanged.

VERDICT: SHIP
