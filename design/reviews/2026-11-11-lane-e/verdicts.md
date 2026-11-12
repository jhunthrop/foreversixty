# Lane E persona gate, 11 Nov 2026

Captures in this folder (1440 and 390): `guild-empty`, `error-account`, `filter-popover-rankings`, `filter-popover-bis` (390 shows the phone control stack open), `planner-share` and `planner-share-copied`, `boss-picker` and `boss-picker-crop`. Made by `web/tests/e2e/lane-e-captures.spec.ts` (`LANE_E_CAPTURE=1`).

| Reviewer | Round 1 | Round 2 | Final |
|---|---|---|---|
| ux-designer | FIX | FIX (phone chips 32 px; EmptyState action white underline) | SHIP |
| wow-player | FIX | SHIP | SHIP |
| ai-slop-reviewer | FIX | PASS | PASS |
| info-density-reviewer | FIX | FIX (Metric twice on phone, judged pre-existing) | PASS |
| guild-officer-reviewer | FIX | FIX (dense roster busy, Nudge not announced, stale comment) | PASS |
| brand-consistency-reviewer (spec) | FIX | FIX (Lua key names, `selected` vs primary) | PASS |

## Fixed in the rounds
- Account load error says what failed: "Could not reach Forever Sixty. Check your connection."
- Last raid empty leads to the companion ("Get the companion ›"); EmptyState's line action is the gold link everywhere.
- Rankings' Clear filters is inert at defaults.
- Roster Approve, Decline and Remove (both rosters) show a spinner plus "Approving…", "Declining…", "Removing…" on the pressed button only.
- Nudge reads "Copy failed" on a refusal and is announced; disabled with no text.
- Chips are 44 px below 768 at any pointer.
- "Copied. Paste it in game with /fs"; the duplicate `copiedLink` is gone; the BiS Sources footnote is shorter.
- The addon/companion spec was corrected twice against the code.

## Hand-offs (raised, outside lane E)
- Guild overview: "Raids open now" icons right-aligned at 1440; the "Your first raid" explainer line; the "Your guild" tag on a guild's own feed.
- Rankings: the Filters group's Metric select repeats the bar's Damage/Healing tabs (logs spec K2): label it "More metrics" or drop the two.
- Planner share panel layout (button grouping, Show DPS toggle contrast, preview card).
- Lane A: the Protection guide-build icon is an empty frame.
- Lane B: "1830" in the history readout against "1,830" in the score panel.
- Lane D: `guild/dense/WednesdayPanel.svelte` still writes "Try again" (named in the guard until D lands); `GuildSeatRow` is the last `BUSY_CLASS` user (delete `lib/ui/busy.ts` after).
- `report/ReportLoot.svelte`'s empty line keeps a link mid-sentence, which EmptyState's one line cannot carry.
