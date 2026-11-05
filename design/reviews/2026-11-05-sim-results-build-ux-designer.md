# Saved sim result, run kind: build review (ux-designer)

Spec: `design/specs/2026-11-04-sim-results.md` (§15 rulings binding). Mocks: `design/mocks/2026-11-04-sim-results/*.png`.
Reviewed: the integrated captures in `shots/int/`, the lane state captures in `shots/lane-{a,b,c}/`, and the integrated `web/dist` served locally (`/sim/simfullfixt1`, `/sim/simfixtureab`) at 390, 1024, 1440 and 2000 to measure what the captures could not show.

## Verdict: SHIP WITH FIXES

The composition matches the mock region by region at 1440, 1024, 2000 and 390. The band, the nav, Worth knowing, Damage by ability, the spread, the timeline, the paperdoll, talents, resource, buffs, rotation and settings are all in the spec's order with the spec's components and copy. The hierarchy holds: one 64px gold figure, one filled gold button. Every §15 ruling renders: R1 (Frozen Heart of the Mountain stands in for Chthonic Power, no Thrash, no empty boxes in the integrated build), R2 (Faerie Fire present, Curse of the Elements under "No effect on your damage"), R3 (attack power 1,775 and haste "0%", no dashes), R4 (armor and armor after debuffs), R6 (enchant lines at 1440 and 390), R7 (Crit header tooltip, verbatim).

One blocker stops it as it stands: the lazy cards reserve the fixture page's own heights on every result, so any run whose cards are shorter than the Fury fixture's shows black holes up to 460px between cards. Fix that and the four must-fixes below and it ships.

Comparison with the competitor (§2), as the review must state it:
1. **Hierarchy:** one 64px figure. Every other fact is a quiet 15px value. Holds.
2. **Answer before diagnosis:** Damage by ability sits above the timeline. Holds.
3. **Every ability and item is real:** holds, except the rage "Refund" row (deviation 4).
4. **Honest about old runs:** the fallback copy all renders. The holes make old runs look broken (deviation 1).
5. **Phone first:** every lane is kept, names sit inside the plot, the readout is pinned under it, nothing truncates. Holds.

## Deviations

### 1. Lazy card reservations: fixed fixture heights stay after mount. BLOCKER

- **Element:** the `section` wrappers of Timeline, Gear, Talents, Resource, Buffs and Rotation in `SavedSim`.
- **Spec:** §4.2 and §8 say every skeleton is its ready height, computed from data the main chunk already has. §3.7 gives the timeline's formula. "Nothing moves", and nothing is left empty.
- **Shipped:** every section carries a hard-coded min-height measured on the fixture page: Talents `min-h-[687px] md:min-h-[582px]`, Resource `min-h-[824px] md:min-h-[561px]`, Buffs `min-h-[1231px] md:min-h-[775px]`, Gear `min-h-[1193px] md:min-h-[1086px]`, Rotation `min-h-[2005px] md:min-h-[1395px]`. The timeline gets an inline `min-height: 1528px`. These stay after the card mounts. I measured them live at 1440:
  - **Fixture run:** the timeline card is 1,367px inside a 1,528px box, which leaves a 161px hole under the cast list (visible in `int/full-1440.png` between the timeline and Gear). Rotation leaves 42px.
  - **Older run (`simfixtureab`, A15):** Talents fallback 129px in 582 (453px hole). Resource 246 in 561 (315). Buffs 315 in 775 (460). Gear 933 in 1,086 (153). `int/older-1440.png` shows three black gaps.
  - **Lazy chunk failure (A16):** the timeline LoadError card is about 156px inside the 1,528px box, which leaves about 1,370px of black.
  - **Any other class or spec:** a run with fewer rotation lines, fewer auras or no pet still gets the Fury fixture's heights.
- **Fix:**
  - Drop the hard-coded `min-h-[…]` classes.
  - Compute each card's reserve in the main chunk from the result. Use the §3.7 formula for the timeline. For the other cards, use the row counts the result already gives: rotation lines, aura rows, resource source rows, gear present or not, and `talentsFit` (with the fallback it is the 129px head-and-line card).
  - Remove the reserve (set `min-height` to none) once the lazy card has mounted **or** has failed. Below-the-fold growth is not counted as CLS, so release is safe.
  - Fix the timeline formula: it reserves 161px more than the ready card at 1440 with Swings and ticks off. Check that it counts 31 lanes (12 Casts, 10 On you, 9 On the target) and 3 group headings, and that it does not add the Swings and ticks group.
  - Add a Playwright assertion on both fixtures at 390 and 1440: for every lazy section, the section's height equals its card's height within 2px after load.

### 2. Rotation notes still use three Execute thresholds. MUST FIX

- **Element:** Rotation card notes (lines 3, 4, 6, 12) on the fixture run and the older run.
- **Spec:** §3.11 and §14 finding 5 call the window "the Execute phase", always with a capital E. The curated Fury notes are rewritten and ship with the guide text.
- **Shipped:**
  - Line 3 Death Wish: "otherwise in the execute phase".
  - Line 4 Recklessness: "In the execute phase".
  - Line 6 Execute: "First, below 20%: it takes every point of rage you hold".
  - Line 12 Heroic Strike: "once you have Execute never from 25% down".
  - Settings says "Last 25% of the fight, standing in for the boss under 20% health". The page therefore still states three thresholds, which is the wow-player round 1 finding.
- **Fix:** apply the §3.11 rewrites to the curated Fury APL notes in `rotations.json`, with the same guide text:
  - "First, in the Execute phase: it takes ...";
  - "... once you have Execute never in the Execute phase: hold that rage for it.";
  - "the execute phase" becomes "the Execute phase" twice.

  Grep every APL with an execute line for "below 20%", "from 25% down" and lowercase "execute phase". Get a wow-player SHIP on the warrior text under the sim-changes-ship-their-guide-text rule.

### 3. The Deep Wounds target bar starts at the pull. MUST FIX (lane E data)

- **Element:** timeline, On the target, the Deep Wounds lane (`A7-1440-close.png`, `A7-390-*`, integrated full pages).
- **Spec:** §10 check 8 says a DEBUFF track's uptime must agree with its periodic row. §3.7 says target lanes are ordered by first application.
- **Shipped:** the Deep Wounds bar runs from 0:00. In Close up, the Deep Wounds ticks lane directly above starts at about 0:06. The page contradicts itself in two adjacent lanes. The 98.9% uptime is consistent with a bar that is only 2 s too long, so the segment start is the error.
- **Fix:** lane E starts the `sample_auras` Deep Wounds segment at its first application in the sample (the first crit that applies it), not at the pull. Extend the §10 check 8 fixture test to compare that first segment start with the first `spell:12867` tick minus one tick period.

### 4. Rage source "Refund": engine term and generic glyph. MUST FIX

- **Element:** Rage card, "Gained from", the row "Refund 78" (1440 and 390).
- **Spec:** §2 point 3, §3.10 and R1 say every row is a real, named game source with its own icon. No engine vocabulary reaches the page.
- **Shipped:** the label is "Refund" next to a bare circular-arrows glyph. No player knows what it refunds.
- **Fix:**
  - Label: "Refunded on a miss, dodge or parry". This is the 80% rage return on a failed special.
  - Icon: a 24px `IconTile` holding that glyph, so the row reads as deliberate and not as a missing icon.
  - Lane E confirms that the engine source is that refund and nothing else. If it is a different mechanic, use that mechanic's in-game name.

### 5. First screen at 1440 x 900: the third Damage row is cut. POLISH

- **Element:** the first viewport (`A3-first-screen-1440.png`, live).
- **Spec:** §1 asks for at least the first three rows of Damage by ability at 1440 x 900.
- **Shipped:** Execute and Off hand are whole. Main hand starts at about 885px and is cut at 900. The global header grew after the mock.
- **Fix:** recover 48px:
  - band top padding 28 to 20px;
  - facts row top gap 24 to 16px;
  - the gap from nav to first card 32 to 24px;
  - Worth knowing note row padding 12 to 8px vertical, and CardHead bottom padding 4px less on the warnings card only.

  Re-take A3 at 1440 x 900.

### 6. Timeline loading placeholder head differs from the ready head. POLISH

- **Element:** timeline CardHead while the chunk loads (live, chunk held).
- **Spec:** §3.7 says the card renders its CardHead and a Skeleton until the chunk arrives. The swap is one `.reveal`.
- **Shipped:** the placeholder head uses the chart glyph and no meta. The ready head uses the clock glyph and "The median run · 2:53 long". The glyph swaps and the meta pops in on load.
- **Fix:** render the real CardHead (clock glyph, label, title, meta from `sample_duration_ms`) in the main chunk, and lazy-load only the body.

### 7. Talents meta spacing. POLISH

- **Element:** Talents CardHead meta.
- **Spec:** §3.9 gives "17 / 32 / 2" (numeric). The band descriptor keeps "17/32/2 talents".
- **Shipped:** "17/32/2". The older run shows "0/24/0".
- **Fix:** format the meta with thin-spaced slashes, "17 / 32 / 2", from the same `talentSplit()` value.

### 8. Timeline On the target: nine identical bars. POLISH (spec gap, my call)

- **Element:** On the target group in the timeline.
- **Spec:** §3.7 orders target lanes by first application. It is silent on debuffs that are up all fight and do nothing for this spec.
- **Shipped:** nine full-length gold bars. Five of them are Judgement of Wisdom, Judgement of Light, Judgement of the Crusader, Curse of the Elements and Hunter's Mark, which the Buffs card itself marks "No effect on your damage". Deep Wounds, the one bar that varies, sits last.
- **Fix:** apply the Buffs card's "No effect on your damage" rule (R2) to timeline target lanes: damage-relevant lanes first, then the no-effect lanes at 50% opacity under the same 13px muted sub-label. Lane order inside each part stays by first application. Optional for this landing; it lands in the spec as a §3.7 amendment if taken.

### 9. Acceptance captures incomplete or not showing their state. MUST FIX (evidence, not product)

- `int/full-390.png` and `int/sbs-390.png` lose everything below about 8,290 css px: Rotation is cut after its head and Settings is black. This is Chromium's tall full-page capture limit. Live at 390, Rotation (2,005px) and Settings (1,444px) render correctly. Re-take 390 as stitched segments (scroll and clip in 4,000px slices).
- `A14-lazy-skeleton-1440.png` shows no lazy card skeleton. I verified it live with the timeline chunk held: head plus skeleton. Re-take with a chunk held.
- `A16-lazy-failed-{1440,390}.png` show the Settings card, not a failed card. I verified it live with the chunk aborted: "This part of the page didn't load." plus Try again. Re-take after deviation 1 is fixed, so the hole is gone.
- `A5`: the capture shows W1, W2 and W4. The spec asked for W1, W2 and W5. With the three-note cap and the priority order, W4 must be removed from the variant for W5 to show. Re-take with Recklessness' casts removed.
- Not attached: A11b (an imported character with enchants on a real import, the Extra attacks row at 390), A11c (a suppressed-debuff chip variant), A12 (Settings Show all expanded on phone), A18 (Lighthouse for `/sim/simfixtureab.html`). The island size line is in `build-int.log`: `dist/sim-island.js` 78,384 bytes gzipped, under the 90 KB budget.

### Observation for the accuracy loop (not a page deviation)

The Fury rotation note for Berserker Rage says "Improved Berserker Rage makes it 10 rage", and the run casts it 6 times. "Gained from" has no Berserker Rage row. Either the published talents skip Improved Berserker Rage, in which case the note overclaims for this build, or the engine does not credit the rage. This is passed on, not raised.

## Checklist (spec §9, §10, §15)

**Acceptance captures (§9)**
- [x] A1 default full page at 1024, 1440, 1920 and 2000. Column centred at 1344 with full-bleed band at 1920 and 2000.
- [ ] A1 at 390: capture truncated (deviation 9). The live render is complete.
- [x] A2 side by side at 1440 and 390.
- [ ] A3 first screen: 390 passes (crest, spec, figure, error, Open in the sim). 1440 x 900 misses the third row (deviation 5).
- [x] A4 nav stuck, compact figure "904.4 DPS", current link Timeline. Verified live: focus moves to the Timeline `h2` and `aria-current` is set.
- [ ] A5 warnings: shows W1, W2 and W4, not W5 (deviation 9). The rows themselves match §3.5.
- [x] A6 row hover, focus ring on the Crit header, Crit ascending ▴, Execute spell tooltip, tap tooltip at 390, R7 tooltip text verbatim.
- [x] A7 Whole fight, Close up with brush window, Swings and ticks, crosshair readout, pinned readout with close, tick tooltip with outcome line, keyboard focus, at 1440 and 390.
- [x] A8 phone scrolled to the end: sticky icon column and axis, in-plot lane names.
- [x] A9 histogram tooltip "910 to 919 DPS: 795 runs".
- [x] A10 item tooltips on the left and right columns, enchant lines in the client's text ("+7 Strength", "Crusader", "Minor Speed Increase"), active set foot, at 1440 and 390.
- [x] A11 three trees, phone tabs, talent tooltip, §4.7 fallback line verbatim.
- [ ] A11b imported character with enchants, Extra attacks row at 390: not attached.
- [ ] A11c Target armor shown (yes, integrated); suppressed-debuff chip variant not attached.
- [ ] A12 rotation with dimmed lines: yes. Settings Show all expanded on phone: not attached.
- [x] A13 Copied and Copy failed.
- [ ] A14 shell skeleton: yes. Lazy card skeleton: capture does not show it (verified live).
- [ ] A15 older result: every fallback line renders. Holes between cards (deviation 1).
- [ ] A16 load failed: yes. Lazy failure: wrong capture, and a 1,370px hole live (deviation 1).
- [x] A17 stale: W1 note, "Older engine" tag.
- [ ] A18 Lighthouse report: not attached. Island size 78.4 KB gzipped: passes.

**Accuracy (§10)**
- [x] 1. Execute 190.8, 21.1%, 24, 1,493, 52.2%. Main hand 63 swings, 506, 42.6%, Miss 19.0%, Dodge 6.3%, Glance 39.7%. Extra attacks 21 and 557. Deep Wounds 59 ticks, 141, 37.3%.
- [x] 2. Actions per minute 40.4.
- [x] 3. Split 17/32/2 in the band equals the Talents meta (spacing aside, deviation 7).
- [ ] 4. Every key named and iconed: the rage "Refund" row has neither a game name nor a game icon (deviation 4).
- [x] 5. W4 alone on the fixture. W2 fires on the variant.
- [x] 6. The histogram reads 10,000 runs.
- [x] 7. Rage sources sum to 2,127 and 2,116 (to 1).
- [ ] 8. Debuff track agrees with ticks: the Deep Wounds bar starts before its first tick (deviation 3).

**§15 rulings**
- [x] R1 icons: no empty box in the integrated build. Chthonic Power shown as Frozen Heart of the Mountain. Thrash absent.
- [x] R2 buffs: same rule and list as Settings. Faerie Fire present. Curse of the Elements under "No effect on your damage".
- [x] R3 stats: attack power 1,775 and haste 0%, no dashes.
- [x] R4 target: "Level 63 boss · 3,731 armor, 471 after debuffs".
- [x] R5 rage: Bloodrage is a gain, never a spend.
- [x] R6 enchants at 1440 and 390.
- [x] R7 crit convention and header tooltip text verbatim.

**Layout and states (§3 to §8)**
- [x] Section order and nav labels. Nav drops Timeline without a sample.
- [x] Band hierarchy: one 64px gold figure, one filled gold button. Phone buttons are stacked full width.
- [x] Damage table columns, sub-lines with nowrap pairs, white rows with weapon names in quality colour, Extra attacks row with "From Windfury Totem 21", Biggest hit hidden, Uptime only on Deep Wounds.
- [x] Spread: real bins, 16th to 84th percentile band, caption "Two runs in three land between 861 and 958 DPS.", six tiles. At 1024, histogram 7fr plus tiles 5fr.
- [x] Timeline: groups, Execute phase wash, resource lane, brush, cast list foot copy, Whole fight landings only, phone in-plot names.
- [x] Gear: enchant line, stats block, set foot, Best in slot link.
- [x] Talents: trees with art and ranks. Phone tabs with the most-points tab selected.
- [x] Buffs: bars, Up all fight, No effect group, phone bars.
- [ ] Rotation: structure, dimming and counts correct. Notes not rewritten (deviation 2).
- [x] Settings: rows, chips, damage-first grouping, phone Show all 14 and Show all 8, actions grid.
- [ ] No layout shift and no empty space after hydration (deviation 1).
- [x] Errors: "This sim did not load" card and LoadError copy verbatim.

## Exact fixes, in landing order

1. Deviation 1: data-computed reserves, released on mount or failure, plus the height-equality e2e on both fixtures.
2. Deviation 2: curated APL note rewrites with wow-player SHIP on the warrior text.
3. Deviation 3: lane E Deep Wounds segment start, plus fixture test.
4. Deviation 4: Refund row label and IconTile.
5. Re-take A1-390 (segmented), A3-1440, A5 (W5), A14 lazy, A16 lazy, and attach A11b, A11c, A12 phone Show all and A18.
6. Polish 5 to 8 when convenient.

## Round 2 (land/simview at 2a93bf7b6)

### Verdict: SHIP WITH FIXES

The blocker is fixed. Every round 1 must-fix is fixed or answered. Two small must-fixes remain, both one-line data or text changes. Once they are in, ship without another review round; the integrator checks them on the captures.

Reviewed: `int2/sbs-{1440,390}` and their small versions, the stitched full pages for both fixtures at 390, 1024, 1440, 1920 and 2000, and `fix-layout/` and `fix-page/`.

### Round 1 items

1. **Blocker, lazy reservations: fixed.**
   - The gap under the timeline is gone at 1440.
   - The older run has no holes between cards at 1440 or 390.
   - A failed timeline is a 158px error card with Gear directly under it (`lazy-failed-{390,1440}`).
   - The loading head is the real head: clock glyph and "The median run · 2:53 long".
   - The new `sim-saved-layout.spec.ts` keeps it fixed.
   - CLS on the older fixture is 0.0186 at 1440. That is inside the 0.05 budget.
2. **Must fix, Execute wording: fixed.**
   - Every note says "in the Execute phase".
   - Heroic Strike reads "never in the Execute phase: hold that rage for Execute".
   - Overpower now reads "Never at 60: Overpower needs Battle Stance, and Fury fights in Berserker Stance."
   - Settings and the notes now state one window.
3. **Must fix, Deep Wounds bar: withdrawn.** An off-hand crit at 0:00.01 applies it, and a fixture test pins that. My reading of the Close up capture was wrong.
4. **Must fix, Refund row: fixed.** It reads "Refunded on a miss, dodge or parry" in a 24px IconTile.
5. **Polish, first screen at 1440 x 900: still open.** Main hand still starts at about 885px. It stays polish.
6. **Polish, timeline placeholder head: fixed.**
7. **Polish, Talents meta: fixed.** It reads "17 / 32 / 2".
8. **Polish, target lanes: taken.** Raid auras that are up all fight are left out of the timeline lanes, and On the target now shows Deep Wounds alone. This departs from §7's "any lane of the timeline never collapses". I accept it because those auras are listed in the Buffs card with their uptimes. The spec owner amends §3.7 and §7 to say so.
9. **Captures:** the stitched full pages and the lazy captures resolve every item, except the two named below.

### Coordinator changes accepted, spec copy to follow

- **R7 Crit tooltip, reworded:** "Crits out of the hits that landed, glancing blows included. Misses, dodges and parries are left out, so a row can read above or below your crit chance on the boss." It is accurate and clearer than the §15 text. §6 and §15 R7 take the new wording.
- **The Hit cell shows the run's measured misses** ("Missed in this run: 0% of specials · 13.9% of white swings") instead of the 220px bar to 27% and the reference line. Both figures agree with the table (white rows 19.0%, 8.5% and 19.0%; no special shows a miss). §3.8 and §6 take the new line.
- **The IN PROGRESS fidelity note on the saved Rotation card** reads as one quiet line with a muted Tag. It is in the right place. §3.11 takes it.

### Remaining must fix

**R2-1. Blood Fury has no icon (older run, `int2/full-simfixtureab-*`).** MUST FIX
- **Spec:** R1 says no icon exemptions at merge, and an action key with no icon never draws an empty box.
- **Shipped:** Buffs, On you, "Blood Fury 16.7% 2×" has no icon at all; the name sits where the icon should be. The cast record has the same gap. The full fixture is clean, but the older fixture is the A15 acceptance page.
- **Fix:**
  - Add the orc racial Blood Fury to `data/curated/extra-spell-icons.json`: the cast id and the aura id the old result carries. Take the icon from the client's SpellIcon table.
  - Extend the §11.4 name-and-icon test to iterate `web/src/fixtures/sim/result.json` as well as `result-full.json`.

**R2-2. The item tooltip prints the enchant as "Strength +7".** MUST FIX
- **Spec:** §3.8 and R6 say the enchant text is the client's own `SpellItemEnchantment` name, "+7 Strength". The paperdoll line under the item already prints "+7 Strength".
- **Shipped:** in `tooltip-hands-panel-1440`, the green line inside the Raider Gauntlets tooltip reads "Strength +7". The card and its own tooltip disagree, and the tooltip is not what the client prints. "Crusader" on Persuader is right.
- **Fix:** the tooltip's enchant line reads `enchants.json` `display`, the same field as the paperdoll line, for every enchant. Add a unit test that the tooltip and the paperdoll line render the same string for 1887 (+7 Strength) and Crusader.

### A11c and A18

- **A11c, the suppressed-debuff chip: does not block.** The raid preset has no exclusive pair, so no shipped page can show the state today. It is a fixture variant owed with the first preset that carries one. Add it to the follow-up list.
- **A18, Lighthouse: blocks the push, not this verdict.** §8's budgets on `/sim/simfixtureab.html` at 360px (performance ≥ 0.90, accessibility ≥ 0.95, SEO ≥ 0.95, TBT ≤ 200 ms, CLS ≤ 0.05) are a gate the integrator runs before pushing, per the shared-layout rule. The island is 78.4 KB gzipped against the 90 KB budget, and CLS is measured inside budget, so I expect it green. If it is red, that is a must-fix for the lane that owns the failing metric.

### Checklist delta from round 1

- [x] A1 at 390 (stitched).
- [x] A14 lazy skeleton.
- [x] A15 older result: no holes; fallbacks verbatim. Blood Fury icon is open (R2-1).
- [x] A16 lazy failure.
- [x] §10.4 for the full fixture. Open for the older fixture (R2-1).
- [x] §10.8 Deep Wounds (withdrawn).
- [x] Rotation wording.
- [x] No empty space after hydration.
- [ ] A3 at 1440 x 900: third row cut (polish).
- [ ] A5 with W5 shown: not re-taken. Polish; the cap behaviour is right.
- [ ] A11c: deferred, see above.
- [ ] A18: integrator gate.
- [ ] R6 tooltip text (R2-2).
