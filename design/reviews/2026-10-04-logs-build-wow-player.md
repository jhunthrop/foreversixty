# /logs landing build — wow-player review (2026-10-04)

Reviewed on the build's own captures: `design/mocks/renders/logs-build/*.png` (37 files on
disk, 33 named in the lane report). Side-by-side: `side-by-side-1440.png`. Mock review:
`design/reviews/2026-10-04-logs-mock-wow-player.md`. Spec: `design/specs/2026-10-04-logs-landing.md`
§10, §12. Lane report: `.worktrees/web-logs-rebuild/design/reviews/2026-10-04-logs-build-report.md`.

## Verdict: FIX

1. **The 360px captures for "has reports" and "zero reports" are swapped** — this is the exact
   failure mode my mock review flagged as the page's worst bug, now back at one specific
   breakpoint. `signed-in-reports-360.png` shows the **sample** hero ("Sanguine Depths, sample
   log"), a signed-out-style Upload panel ("Sign in to upload…"), and no device-status line — even
   though the header still reads "Fixture#4242" signed in. Meanwhile `signed-in-zero-360.png`
   shows the **real** hero ("Molten Core, week 3"), a full device-status line, and the full Upload
   form — the content that belongs in `signed-in-reports-360.png`. Every other width (390, 1024,
   1280, 1440, 1920) has this correctly paired: hero matches the `Your reports` row directly below
   it. Confidence: high that the two 360px files are swapped (confirmed by direct pixel/MD5
   comparison against the correctly-paired 390px pair); medium on whether this is a capture-script
   mislabel or an actual 360px rendering bug — I cannot tell which from static PNGs, but I have no
   valid evidence today that the real "signed-in, has reports" state renders correctly at 360px.
   Re-capture (or fix) and reconfirm before ship.

2. **The upload error state leaks a raw backend error to the player.** `upload-error.png` reads
   `The upload did not finish: part 1 failed three times (The upload did not finish: R2 answered
   500)` — doubled wrapper text, "part 1," and "R2" (Cloudflare's storage product, an
   implementation detail no player should ever see) is exactly the line a raid officer screenshots
   into guild chat to mock. Also on this same panel: `WoWCombatLog.txt · 1 bytes` — wrong
   pluralization, reads amateur next to the client-accurate copy everywhere else on this page.
   Expect: a plain-language line ("The upload didn't finish. Try again.") with no retry-count or
   storage-backend text, and `1 byte` singular. Confidence: high — directly visible in the capture.

3. **The Upload panel's Title field shows a fixed example that coincidentally matches the
   session's own report title, everywhere, including signed out.** `Molten Core, week 3` sits in
   the Title box on `upload-idle.png`, `upload-chosen.png`, `upload-progress.png`,
   `upload-error.png`, and every full-page capture regardless of hero state — it's still there
   verbatim on `signed-in-zero-1440.png` (hero = "Sanguine Depths, sample log") and on
   `signed-out-1440.png` (no session at all). Pixel-sampled, it's dim placeholder-grey, not a
   filled value, so it won't actually submit — but a player skimming the panel has no way to tell
   that at a glance, and the one-time coincidence that it's the literal title of the fixture's own
   real report (used elsewhere on the same page) reads like leftover dev/test text rather than a
   generic example. Expect: a placeholder that's obviously an example (e.g. `e.g. Molten Core,
   week 3` or a zone name that never matches the page's own hero), or just blank with the existing
   `Optional. Without one…` helper line doing the job, as the mock shows. Confidence: medium-high
   on it reading as developer leftovers; this is `Upload.svelte`, kept verbatim by this spec, so
   it may predate this pass — still visible on the rebuilt page today.

4. **One of the two named "companion status" deliverables doesn't prove what it claims to.**
   `companion-status-seen.png` and `companion-status-never-seen.png` are byte-identical (same
   MD5) — both show `MacBook Pro · macOS · last seen 4 minutes ago`. The "paired, not seen yet"
   copy is real and does work (it's visible correctly in `signed-in-zero-1440.png` and
   `signed-in-zero-390.png`'s full-page captures: `Gaming PC · Windows · paired, not seen yet`),
   so I don't think the feature is broken — but the dedicated acceptance crop for that state is a
   duplicate, not a second screenshot. Expect: re-capture
   `companion-status-never-seen.png` for real. Confidence: high on the duplicate (MD5-confirmed);
   low concern on the underlying feature, since it's independently proven elsewhere.

## What reads fixed from my mock review

- **Hero vs `Your reports` mismatch (finding 1):** fixed at every width I could verify (390,
  1024, 1280, 1440, 1920) — `Molten Core, week 3` hero exactly matches the `Your reports` row
  below it (same date, fight count, kill count). Only the 360px pair is in doubt (item 1 above).
- **Retail-ism sample report (finding 2):** not fixed, as expected — `Sanguine Depths, sample
  log` is still the hero/sample everywhere. This is already filed as a blocking data defect by
  the spec (§4.C.3) and the coordinator's own ruling kept it unaltered on purpose pending a real
  fix; not a new defect, just still outstanding before this is a reference example.
- **Phone row text collision (finding 3):** reads fixed — `Molten Core, week 3` stacks above its
  `2026-10-03 · 9 fights · 4 kills · public` line in every phone capture instead of sharing one
  line with a pinned date. I only had a short title to test this against, but the layout itself
  (title line, then stats line below) is visibly different from the mock's inline version and
  matches the described fix.
- **Orphaned dungeon-framing sentence (finding 4):** fixed — "Logs are for group content at any
  level…" sits directly under the spine on every phone and desktop capture, before `Your
  reports`, on both signed-in and signed-out boards.
- **Green "live" status dot (finding 5):** fixed — the companion status line is plain text in
  every capture, no dot, matching the §12 ruling.
- **Hero h1 + Sample pill redundancy (finding 6):** not fixed, but this is a designer ruling, not
  an oversight — the lane report notes the pill was kept deliberately "even though the sample's
  own title already says sample." Still reads slightly redundant to a player, lowest severity,
  called out and intentional.

## Hero line, numbers, companion honesty (the task's direct questions)

- `Last kill, Lucifron: Zulmara · 1,250 DPS · Gear in the planner | Sim` reads right and is
  exactly what I'd click after a pull — fight name, player, DPS, straight to gear/sim. This is
  the site's real differentiator and it's finally on the landing page, not two clicks deep.
- Numbers are consistent between the hero and the row beneath it everywhere except the swapped
  360px pair (item 1): `9 fights · 4 kills` on both the hero facts line and the `Your reports`
  row, same date.
- Companion status is honest in both states where I could verify it (`last seen 4 minutes ago`,
  `paired, not seen yet`), no "Connected" claim, no dot. The pairing success line
  (`Zulmara's PC paired. It starts uploading as soon as you are logging.`) tells you exactly what
  happens next, closing the old "did that work?" gap.
- Signed-out correctly shows the sample labelled `SAMPLE` with a static hook sentence (no fetch),
  and `Recent public reports` shows its own honest empty copy, never the sample a second time.
- Phone at 360 and 390 (aside from the item-1 swap): no clipped controls, guild tab strip and
  every button held its hit area, region order matches §7 (hero → spine → framing line → Your
  reports → Recent public reports → companion → upload).

## Confidence summary

- Item 1 (360px swap): high on observation, medium on root cause — must re-verify before ship.
- Item 2 (raw upload error / "1 bytes"): high.
- Item 3 (Title placeholder reads like dev leftovers): medium-high.
- Item 4 (duplicate companion-status capture): high on the duplicate, low concern on the feature.

---

## Fix round 1 re-check (2026-10-04)

Re-shot files checked: `signed-in-reports-360.png`, `signed-in-zero-360.png`, `upload-error.png`,
`upload-chosen.png`, `upload-idle.png`, `companion-status-never-seen.png`.

1. **360px swap — fixed.** `signed-in-reports-360.png` now shows the real hero ("Molten Core,
   week 3") with the full device-status line and Upload form, matching its own `Your reports` row
   exactly (`9 fights · 4 kills`, same date). `signed-in-zero-360.png` now shows the sample hero
   ("Sanguine Depths, sample log") with the genuinely empty `Your reports` and `Gaming PC ·
   Windows · paired, not seen yet`. Confirmed a capture-isolation artifact as described, not a
   product bug — no longer an open concern.
2. **Upload error copy — fixed.** `upload-error.png` now reads `The upload did not finish. Try
   again; the file is still chosen.` — plain language, no `R2`, no `part 1 failed three times`,
   no doubled wrapper text. `WoWCombatLog.txt · 1 byte` is correctly singular now, on both
   `upload-chosen.png` and `upload-error.png`.
3. **Title placeholder — fixed.** `upload-idle.png`, `upload-chosen.png`, and the 360px full-page
   captures all show a plain dim `Optional` placeholder in the Title field, matching the mock —
   no more `Molten Core, week 3` leftover text anywhere, including on the sample hero.
4. **Companion "never seen" crop — fixed.** `companion-status-never-seen.png` is now a distinct
   file (MD5 differs from `companion-status-seen.png`) and reads `Gaming PC · Windows · paired,
   not seen yet`, genuinely different copy from the "seen" crop's `MacBook Pro · macOS · last
   seen 4 minutes ago`.

All four items resolved. No new issues surfaced in the re-shot files.

## Verdict: SHIP
