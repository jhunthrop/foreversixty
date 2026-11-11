# Guild tab: ux-designer verdict, round 4 (2026-11-11)

Reviewed on images: parity/ios-full-guild-tab-leveler.png and parity/android-full-guild-tab-leveler.png at full size, plus round3/ux-designer.md and round3-response.md. The officer, none and none-leveler renders and the device walks were not opened; the officer-side claims (Ping result line, `41 DPS`) are not re-verified here.

## Verdict: SHIP

Must 1 is closed on both full-page renders. The shoulds are routed as follow-ups, and one routing I do not accept as written (should A).

## Must 1: closed

- **ios-full-guild-tab-leveler.png:** under "Raids open at 60 · you're 36", Who is coming reads "Tanks 2", "Healers 5", "Damage 11 · 5 melee, 6 ranged". The green "of 2", "· 3 short" and "19 short" are gone.
- **android-full-guild-tab-leveler.png:** the same three rows; Brightwen's "(Late · 15 min)" stays as a per-person annotation, not a target.
- **Whole page:** both renders run from the guild head through the raid card, Who is coming, Roster, Recent logs, Progression and Pages.
- **Both apps now match the device walks**; one view no longer renders two ways.

## Shoulds: routing check

- **A. Android chevrons (should 4): not fairly routed.** In android-full-guild-tab-leveler.png the Recent logs, Progression and Pages rows (Calendar, Loot history) carry no chevron; iOS has them. These are this lane's rows. Add a chevron to every navigating row on Android, and confirm the officer rows (Plan a raid, Applications, Setup) in android/dark/officer-02-officer-rows.png. The external-link glyph on hand-off rows is still open on both apps.
- **B. Late echo drops the roster (should 2): accept.** Note that the shipped server echo keeps the rest of the roster; check it on the first real device run.
- **C. "Saved." layout shift (should 3): accept the routing to the shared raid flow.** Needs a named owner and a frame after "Saved." clears.
- **D. Calendar leaves the app on Android (should 5): accept.** Needs an owner named in interactions.md.
- **E. Missing raids and Gear gaps above Recently on (guild-officer should 3): accept.**

## New observations (polish, no action required for SHIP)

1. ios-full-guild-tab-leveler.png, Healers row: "Brightwen (Lat…" truncates mid-word. Wrap as Android does, or drop the annotation when it will not fit.
2. android-full-guild-tab-leveler.png: about two thirds of the image is empty background below Pages (capture window). Crop to content next time.
3. The iOS render shows sub-lines smaller than on device; the device walks remain the reference for type size.

## Checklist

- [x] Under 60, Who is coming shows counts alone on iOS and Android full pages.
- [x] "Raids open at 60 · you're 36" in the text ink under the raid card.
- [x] No Before you pull under 60.
- [x] The full-page renders are true whole pages.
- [ ] Android chevrons and hand-off glyphs (should A, follow-up).
- [ ] "Saved." height reserved, with a cleared frame (should C, follow-up).
- [ ] Roster kept under the Late echo (should B, follow-up).
