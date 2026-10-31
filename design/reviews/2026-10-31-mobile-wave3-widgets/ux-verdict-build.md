# F14 widgets and live surfaces: build review (UX)

Reviewed: 40 ios-* and 32 and-* renders in `build/` against spec §5.6, §4, §6, §7, §9, §13 and the boards `f14-surfaces-ios.png` and `f14-surfaces-and.png`. Pixels only, no diff.

**Verdict: HOLD.** Six MUSTs. The structure, copy, figures and state coverage match the boards; the holds are one truncated failure label, one clipped footer, one Android tag that disagrees with iOS and the spec, one false notification title, a missing crest and an unverified tint.

## MUST
1. **iOS `Not saved` truncates at the largest size** (`ios-raid-medium-not-saved-ax5.png`: `Not sav...`). It is the only feedback that a tap failed, and §9 forbids cutting an answer. Fix: let the failed button wrap to two lines (`Not` / `saved`, lineLimit 2, 44 pt height kept) or widen it by sizing In/Late/Out to content for the 10 s. Re-capture at default and largest.
2. **Android 4x2 upgrades footer is clipped at the cap** (`and-upgrade4x2-fs2.png`: `8 upgrades · +18.3 DPS` is cut by the card edge). This is the exact defect fixed on the board in the design pass, and §9 says the footer never collapses. Fix: at fontScale above 1.0 drop row height or icon to 28 (as the design pass did), or give the card the height. Re-capture.
3. **Android Late and Out tags disagree with the spec and with iOS** (`and-raid2x2-late.png` shows `LATE · DPS` in kill green; `and-raid2x2-out.png` shows `OUT · DPS`; iOS `ios-raid-small-late.png` / `-out.png` show `LATE` / `OUT` in muted). §5.6: `Late`, `Bench`, `Out` stand alone; kill is for `In · DPS`. `WidgetDesign.kt` maps Late to kill and Android builds its own tag word. One WidgetFigures must produce both; add Late, Bench and Out cases to `figures.json` so the goldens catch this.
4. **The raid notification title contradicts its own countdown** (`and-live-raid.png`: title `Invites for Onyxia's Lair in 39 minutes`, header `until the pull at 8:00 pm`). With no `invite_at` the clock counts to the pull, not to invites, and iOS says `until the pull` (`ios-live-raid-38-min.png`). Fix: when `invite_at` is null the title is `Raid starts in {n} minutes`; add the copy key to the spec and the golden `liveTitle` for the no-invite case.
5. **Android upgrade widgets have no crest** (`and-upgrade2x2.png`, `and-upgrade4x2.png`; `and-upgrade2x2-no-gear.png`, `-no-home.png`). §5.6 and the board put a 16 dp crest before `Next upgrade` and `Upgrades · {name}`, and the owner's one-crest rule applies. iOS has it (`ios-upgrade-small.png`). The Android raid widgets have it, so this is an omission.
6. **Lock-screen tint is not verified** (design verdict item 3, §5.6 acceptance: "each lock-screen widget with the system tint"). `ios-lock-*.png` are one default rendering on flat brown; no vibrant or accented rendering is shown, and no `widgetAccentable` is set in the extension. Add a capture of the five accessory widgets under a tinted lock screen (device or simulator with a tint), and mark the ring and the number as the accented group. Also the Android notification small icon renders as an empty rounded square in `and-live-raid.png`, `-sim-running.png`, `-sim-finished.png`: capture the real v1 small icon on device, or this is a missing-icon defect.

## SHOULD
1. **Inline can lose the answer.** `ios-lock-inline.png` carries the full zone, `Onyxia's Lair · Thu 8 pm · In`; the spec uses `Onyxia`. A long zone (`Temple of Ahn'Qiraj`) will ellipsise the `In` at the end, which is the point of the widget. Put the answer where it survives (`In · Thu 8 pm · {zone}`) or use a short zone name, and capture the longest zone.
2. **iOS 18 Tinted home-screen mode is not captured or handled.** Item quality colours and the kill checked state vanish there; nothing sets `widgetAccentedRenderingMode` on item icons. Capture `ios-upgrade-medium` in Tinted mode and keep the icon borders legible.
3. **Medium raid widgets say nothing for `Not signed up` or `Choose a spec`** (`ios-raid-medium-not-signed-up.png`, `-choose-spec.png`, `and-raid4x2-not-signed-up.png`, `-choose-spec.png`). §5.6: the tag shows and a spec press opens the ask. Today the buttons look identical to a normal unanswered row, so `Choose a spec` buttons that open the app instead of answering are a surprise. Put the gold tag in the empty bottom-left of the medium.
4. **At the largest size the countdown touches the In button** (`ios-raid-medium-in-ax5.png`, `ios-raid-medium-not-saved-ax5.png`: `in 2 days 2 h` ends about 12 pt from the button and sits on its top edge). Add 8 pt of clear space above the button row.
5. **`Last:` date format differs** (`and-raid2x2-no-raid.png`: `Thu 3 Dec.` wraps with an orphan `Dec.`; `ios-raid-small-no-raid.png`: `3 Dec.`). The spec is `{d Mon}`. Use the iOS form on both.
6. **Acceptance captures missing from the set:** light and dark wallpaper for F14.1; widget gallery entries on both platforms; the quick-settings tile (active and `No raid planned`); Android cancelled and failed live captures; iOS upgrade no-home state; the Android 4x2 in the `no-gear` state; iOS Live Activity at 5 min exists, Android at 5 min does not.
7. **iOS Live Activity toggles show no thumb** (`ios-settings-live-surfaces.png`, `-ax5.png`): the state cannot be read. Re-render, or note it as a snapshot artifact with a device screenshot.
8. **Android sim progress bar is the system blue** (`and-live-sim-running.png`); the board and the iOS Live Activity use gold. Set the notification colour to gold.

## Polish
- `ios-island-expanded.png` reads `until the pull at 8:00 pm · You: In · DPS`; the lock-screen activity reads `You are In · DPS · Lightningbolltz`. Same snapshot, two phrasings. Use `live.you` in both.
- `and-live-raid.png` drops the guild (`Olympus XXVII`) the board shows in the header.
- Android micro labels (`NEXT RAID`) have no letter-spacing; iOS and the board do (`and-raid2x2-in.png` vs `ios-raid-small-in.png`).
- iOS expanded Island has no seal.

## Rulings on the five decisions
- **(a) Text caps at xxxLarge on iOS and 1.3x on Android: ACCEPT.** WidgetKit and Apple's own widgets behave this way, and the 4x2 must not drop its zone or countdown (§9). Conditions: the cap must not clip anything (MUST 2 and MUST 1 are exactly that), and §4 and §9 must be amended to say widgets cap and the app they open scales fully.
- **(b) Island compact trailing shows a ticking `38:12`: ACCEPT.** `38:12` is more useful than `38m`, it fits in the 44 pt frame (`ios-island-compact.png`), and it ends at 60:00 at most. Amend the spec line and the board. Check on device that it never widens the compact region.
- **(c) Upgrade row opens Home, not the item sheet: ACCEPT as interim.** Without an item deep link the sheet cannot open. Do not style rows as separate targets, and log the item deep link as an open item; the spec text in §5.6 must say the row opens Home until then.
- **(d) No Android picker preview: ACCEPT for this lane, not for launch (SHOULD).** The picker is where a player decides to add the widget. A static preview PNG from the existing Roborazzi renders (`and-raid4x2-in.png`, `and-upgrade4x2.png`) is nearly free. The acceptance list asks for the gallery entry on both platforms.
- **(e) `snapshot.upgrades` null shows `Open the app`: ACCEPT as interim, polish.** It is words, not an empty box (`and-upgrade2x2-no-home.png`), but it says nothing. Prefer `Open the app to see your upgrades.` once the spec adds the line; capture the iOS state too.

## Checklist (§13)
- [x] Every small state has words; no spinner, no empty box (all small states, both platforms).
- [x] In, Late, Out checked state: kill ground and 2 pt underline (`ios-raid-medium-in/late/out.png`, `and-raid4x2-*.png`).
- [x] Locked and cancelled drop the buttons; cancelled drops the countdown.
- [x] Stale: `Updated 14 h ago` at the foot, buttons kept.
- [x] Copy verbatim for signed out, no guild, no raid, locked, cancelled, choose a spec, no gear, live sim running, finished, failed.
- [x] Medium buttons 44 pt iOS, 48 dp Android.
- [x] Figures identical across surfaces for the countdown, time, signed up (one WidgetFigures; `figures.json` read).
- [ ] Failure line legible at the largest size (MUST 1).
- [ ] Nothing clipped at the cap on the medium widgets (MUST 2).
- [ ] Same tag on both platforms (MUST 3).
- [ ] Live surfaces say true things (MUST 4).
- [ ] Crest on every widget header (MUST 5).
- [ ] Lock-screen tint verified; notification icon shown (MUST 6).
- [ ] Acceptance captures complete (SHOULD 6).

**HOLD.** Re-submit with MUSTs 1 to 6 re-captured; the SHOULDs can follow in the same pass or be logged.
