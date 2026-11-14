# /admin/activity review, 2026-11-13

Captures: `admin-activity-1440.png`, `admin-activity-2000.png`, `admin-activity-390.png`,
`admin-activity-not-owner-1440.png` (web/tests/captures/admin-activity, sample figures only).
Internal owner page: ux-designer only.

## Round 1

- ux-designer: FIX. The 390 page overflowed and the grids clipped; the figures sat far from their
  labels at 1440 and 2000; the running week read as final in retention and the headline; the
  headline dividers dangled on a phone.

## Round 2 (all of the above applied)

- ux-designer: SHIP WITH FIXES. Must fix: a blank cohort row read as broken, and nothing said the
  phone grids scroll. Not blocking: phone row labels wrap ("9 / Nov"), the empty right half of
  the panels at 1440.

## Round 3

- Applied both must-fixes: a muted dash for weeks not reached yet (named in the retention meta),
  and a right-edge fade where a phone grid scrolls. The captures in this folder are this round's.
