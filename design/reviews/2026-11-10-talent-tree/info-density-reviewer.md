Verdict: DENSE

Scope: the trees and tooltips only. Judged from after/*.png and before/*.png, GameTooltip.svelte, lib/tooltips/game-tip.ts (GameTipModel: icon, subtitle, lines, description, nextRank, requirement, note) and inventory.md.

The unification is a clear gain in what is shown. Every read-only tree now has prerequisite arrows, client art, spec icons in the heads and a tooltip with icon, Rank n/m, gold rank text and next rank. /sim before a run and /sim/talents had no tooltip before. Both now open one (sim-talents-hover.png). The sim phone view went from a hidden grid to three tabs with the points on each tab (sim-result-390.png). Nothing on these surfaces fails. The findings below are ranked additions.

Findings

1. Guide At a glance and Leveling trees: no hover tooltip in the captures. guide-glance-hover.png and guide-leveling-hover.png show no tooltip open, although the inventory says the guide host serves `talent:<id>:<rank>` keys. Either the capture missed it or the guide path still opens nothing. Confirm before ship. If it is the capture, re-shoot with the tooltip open. If it is real, this is the one MUST, because the owner's complaint was exactly "some don't work on hover".

2. Tooltip placement covers the neighbouring data. On /sim/talents the tooltip lands on the Gains and Drops list (sim-talents-hover.png). In the planner it covers the Protection tree and the cells below the hovered one (planner-hover.png). Flip or shift the tooltip so it keeps clear of the Gains/Drops column (placement in tooltip-position.ts). This is a minor fix.

3. Leveling lane tree does not mark the point taken at the chosen level. The slider reads "Level 60: Improved Bloodrage, rank 2" (guide-leveling-hover.png), but the tree shows only the end state with no highlight on that talent. Ring the cell the slider names, using the same cell data. A raider scrubbing levels wants to see where the new point went. This is a SHOULD.

4. Talent tooltips carry no spell facts for talents that grant a spell. GameTipModel has `lines` (cost, range, cast, cooldown) and SpellTipRow has `cost`, `cast_time_ms`, `cooldown_ms` and `range_yd`. Active talents (Bloodthirst, Death Wish, Whirlwind-type picks) show only name, rank and description (Booming Voice shows a passive description). Fill `lines` from the spell's SpellTipRow when the talent teaches a spell. This is a SHOULD, and the data exists.

5. Compact trees at 18 to 22 px are crowded. At guide-glance-1440 and guide-leveling the rank badges overlap the neighbouring icons. On phone (guide-glance-390) three side-by-side trees stay at that size. Either hide the badge below a size threshold and keep the ring and tooltip, or use the sim's tab pattern on phone. A reader tapping a 20 px cell is the weak point. This is a SHOULD.

6. Sim talent builds: no per-cell diff on the trees. The Gains and Drops list sits beside the trees but the trees do not mark which cells differ. Add a `+n`/`-n` mark on changed cells, from the same loadout delta already computed for the list. Hovering a list row could highlight its cell. This is a COULD.

7. Spell icons and names in the tree heads are good, but spec points have no context. "17 / 32 / 2" is shown with no tree-point requirement to the next tier row. The planner already shows requirement refusals in the tooltip. Showing the row-unlock threshold (5 points per row) on the tree's tier edge would give a veteran what the client shows. This is a COULD.

8. The planner's tooltip shows the band line well (see "Levels 20 to 29 build: 0 of 5"). Read-only trees do not show how the cell compares with the guide build. Add the guide's rank to the sim /talents tooltip for the current-build row ("Guide build: 5/5"). This is a COULD, and it links to the planner rather than restating it.

9. The sim result tree is not tied to the run. A talent's value (DPS gained per point) would be the most useful tooltip line on a result, but it is the sim's own number and does not exist per talent in the repo data I checked.

Phone check: planner keeps tabs; sim gets tabs with points (good); guides keep three small trees (finding 5). Empty and low-data states: the "Your current build 0/20/0" row on /sim/talents shows an honest tree with the counts. Nothing blank.

Needs data
- Per-talent DPS value or marginal gain per point, for the sim result tooltip (finding 9).
- Where a talent is trained or learned (a tier unlock level) if the client talent table carries it; not confirmed.
