
## Round 1

Verdict: FIX (two must-fix, rest optional). The left column is dense and well-sourced; the right column leaves visible dead ground and drops fields the sim message already carries.

Page job: how am I doing, and what is the one thing to do about it. Reachable: ratings pack record (`score`, `tier`, five parts, `as_of`, `fights`, `stale`), `Score.guild(name)` plus `data.characters` (guild members' scores), inbox `last_raid` (per boss `percentile`, `biggest_loss`, `loss_seconds`, `deaths`, `consumables_missed`, `adherence`), `loot_values` (`raid`, `items[].planned`), `sim_result` (`dps`, `error`, `generated_at`, `slots[].slot/item_id/delta`), `upgrade` (`item_name`, `source`, `slot`), `ack`.

### Must-fix

1. Dead space in the right column. me-raid-day.png and me-sim-queued.png leave about 120 px of empty ground under Sim again; me-no-companion.png leaves about 270 px under two Right now rows (the left column runs to the footer). Fill from existing data, no new chrome: in Your sim, put the slot name under each gain icon (`Hands`, from `sim_result.slots[].slot`; for the upgrade fallback `upgrade.slot`), and add the item name in the gain's hover. The gain row is three bare icons with numbers; a raider cannot tell which slot is the +9.0 without hovering. This also lets the sim block sit at a fixed height so Right now and Your sim stop floating.
2. No-companion state wastes the column. The Right now rows there are only Gear and Builds, and Your sim is dropped (spec 8, correct). Use the pack data that is present: under Right now show the guild line the header already has as a row only when it has a figure. Concretely, put the player's place in the guild beside the tier as `5th of 31 in Thornwake`, from `Score.guild(name).members` against `data.characters[*].score` (the pack is local, no invented data). It is context on the one number the tab leads with (the score has tier and gap-to-next but no comparison), and it fills the no-companion rail. Link it to Guild, do not copy the Guild tab's top five. If the pack lacks the guild record, show nothing.

### Optional, ranked

1. Last raid: add `Followed 82%` per boss from `last_raid.bosses[].adherence` (already in the contract, already used by the Rotation lead's `74% followed`; MeTab.bossRow drops it). It fits if Biggest loss narrows; Rotation keeps the detail, Me shows the figure only.
2. Last raid: make each boss row a hover/click into Rotation for that boss (the Rotation tab owns the per-boss detail), so a row has a drill-down.
3. Last raid footer: add the planned drops of that raid as a count (`3 of your plan drop`, from `loot_values.items[].planned`) linking Gear. MeTab reads `loot_values.raid` only; the list is otherwise unused on this tab.
4. Score block: the part bars show only Preparation's reason. Show the weakest part's gap in the same slot as the reason when it is not Preparation (`Parses 228 of 600`'s lowest-share part) only if Score.parts already ranks quality; otherwise skip.
5. Sim date `Mon 7 Dec` could read `Mon 7 Dec, 2 days old` coloured warning past 7 days (`sim_result.generated_at`), so stale numbers are not mistaken for current ones.

### Empty and low-data states

me-no-companion.png, me-no-pack.png, me-stale-no-guild.png, me-new-character.png are honest and name what fills them (app address, pack source). me-new-character.png is two short rows over a mostly empty window; acceptable at level 8, since a section with nothing to say is left out per spec 8. Stale pack states the age in the footer in the warning colour: good.

### Phone

Not applicable (fixed 740 px window); all suggestions above stay inside the existing two columns and add no scrolling.

### Needs data (not a fail)

- Raid-ready list DPS for scale beside the sim (Data.lua has no list DPS; dropped by the lane, correct).
- Previous sim DPS, to show a delta against the last run (the inbox keeps one `sim_result`).
- Previous Forever Score, for a trend or delta (the pack carries one `as_of` record; the spec already names score history in game as out of scope).
- Per-boss spec median parse (named out of scope in spec section 9).

## Round 2

Verdict: PASS.

1. Must-fix 1 resolved: the gains are three rows of icon, name in quality colour and DPS, and the right column now runs nearly to the footer in me-raid-day.png; "Simmed Mon 7 Dec" labels the date. Rotation's row also gained "Tue 8 Dec · parse 58".
2. Must-fix 2 withdrawn: guild place belongs to the Guild tab (spec 4.5 guildYouAre); one job per page says link, not restate. The no-companion rail ending in ground is spec 8's state plus the tallest-tab height rule (spec 3.1), not thin data.

Optional items from Round 1 stand (per-boss `adherence` column, planned-drop count from `loot_values`, boss row drill-down into Rotation, sim age in the warning colour when stale). None blocks.
