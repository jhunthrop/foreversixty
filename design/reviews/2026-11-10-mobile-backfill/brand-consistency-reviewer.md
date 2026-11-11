# brand-consistency-reviewer: mobile backfill (both platforms vs the web)

## Round 1 (2026-11-10, base ac2f37dec)

# Mobile backfill, round 1: brand-consistency verdict (iOS and Android vs the web)

I judged the captures against the shared sources. I did not run either app. The web counterparts are the live `/`, `/bis`, `/character` and the home, raid and account-menu mocks in `web/`.

The sources that decide the contested calls:
- Web sign-up control: `web/src/components/guild/SignupSegments.svelte`, which uses the `ui/Segmented` primitive (decision D11).
- Spec: row 143 and row 198 of `design/specs/2026-11-07-mobile-redesign.md`.
- Mocks: `.seg` in `design/mocks/app/ios-home-raid-day.html` and `design/mocks/app/android-home-raid-day.html`.

## BLOCKER

**B1. The iOS sign-up answer is four separate filled pills; it should be the Segmented control.**
- **What is wrong:** `RaidAnswerButtons` in `apps/ios/App/Raids/Views/NextRaidBlock.swift` draws four `FSButtonStyle` buttons. The chosen one is `.primary` teal and the other three `.secondary` blue, with gaps between them. This is on Home (`home-01-top.png`), the Raid tab (`raid-01-top.png`) and the Home and Raid sign-up steps.
- **What the web does:** one connected control with hairline dividers and one selected segment.
- **What is ruled:** spec row 143 says `Segmented` is a 1 pt `line` outline, 4 radius, hairline dividers, 14/16 600, selected in gold on the `selected` ground with a 2 pt gold underline. Row 198 says the raid page's In/Late/Bench/Out is a Segmented with a haptic.
- **Why a player feels it:** the selected answer is the only teal block, so it reads as a primary action, not a state.
- **Where it should match:** the mock's `.seg` in `design/mocks/app/ios-home-raid-day.html` (outline, 4 px radius, gold selected).
- **Which platform is nearer:** Android is closer, since it is one connected control with a selected segment.
- **Fix:** add a `Segmented` primitive to `apps/ios/Sources/DesignSystem`. There is none today (grep finds no `Segmented` in `apps/ios/App` or `Sources`), so it also fails the spec's section 3.4 primitive list. Draw the answers with it, keep the accessibility-size two-by-two fallback, and fire the R14 haptic on change. The `LateMinutesChips` and `LateMinutesRow` chips can stay.

**B2. Android's Segmented has the wrong shape and colour.**
- **What is wrong:** `apps/android/app/src/main/kotlin/gg/foreversixty/app/flows/FlowSegmented.kt` draws a full pill (`RoundedCornerShape(percent = 50)`). The selected segment is solid `FsColor.tealEdge`, with text colour `FsColor.strong` throughout.
- **What the web and spec say:** 4 radius, with the selected segment in gold text on the `selected` ground and a gold underline (spec row 143, mock `.seg`).
- **Platform idiom:** the plan says "Android's … 48 dp segmented answers". That ruling covers the 48 dp height and the M3 segmented idiom. It does not cover the teal fill and pill.
- **Why it matters:** B1 and B2 together mean the two apps draw the one sign-up three ways: web (Segmented), iOS (buttons) and Android (a pill).
- **Fix:** change `FlowSegmented` to the tokenised outline and gold selection, ideally as a `core/designsystem` primitive shared by both platforms, and keep the 48 dp height.

**B3. Character art is the Battle.net portrait on both apps, not the circular crest.**
- **What is wrong:**
  - Android: the profile button, the drawer header and the switcher rows show race faces in a blue or red faction ring (`android-menu.png`, `android-home-raid-day.png`). `ProfileButton.kt` line 39 and `CharacterPortrait.kt` draw `avatarUrl` first and the class crest only as a fallback.
  - Android Gear: the spec header (the "Fury Warrior" band, `android-gear.png`) uses the portrait.
  - iOS: the same code path (`apps/ios/App/Shell/CharacterPortrait.swift`) is used in the drawer and the switcher.
  - Both platforms' Gear headers should follow the web character page, which shows the LX seal.
- **What the web does:** the account-menu direction and the character page show the class crest in a gold ring.
- **Why it matters:** the owner's "One crest language" ruling bans Blizzard avatars everywhere.
- **Cross-platform drift:** the iOS Gear header uses the LX seal and Android uses the portrait.
- **Fix:** make `CharacterPortrait` on both platforms draw the class crest, ringed in the class colour as on the web. Drop `avatarUrl` from the player-facing look.
- **Related (the iOS profile button):** the glyph in place of the crest is already being fixed, but the fixed button must be the class crest and not a portrait. The faction emblem badge and red ring on the iOS profile button and switcher rows are a platform-only decoration the web does not use. The Android drawer shows the same badge.

## SHOULD

**S1. The raid push says "Thu 5 Nov, 20:00 CST", but the apps say "8:00 pm CST".**
- **What is wrong:** `api/internal/guilds/raids_events.go` sets `raidWhenLayout = "Mon 2 Jan, 15:04 MST"`. It reaches the iOS notification (`notif-02-actions.png`) and presumably Android.
- **What the rest of the product uses:** `3:04 pm MST` in `api/internal/notify/raid_embed.go`, `api/internal/discordbot/signup_card.go`, `api/internal/guilds/raids_confirmation.go` and `calendar_ics.go`, and `h:mm a zzz` in both `RaidClock.swift` and `RaidClock.kt`.
- **Fix:** set `raidWhenLayout` to `"Mon 2 Jan, 3:04 pm MST"`.

**S2. The Android tab roots show the page title in the top bar at rest.**
- **What is wrong:** in `android-raid.png`, `android-gear.png` and `android-activity.png` the title sits beside the avatar in the bar.
- **What is ruled:** N3 says the title is Marcellus once, in the PageHeader, and the bar has no title at rest; it fades in on scroll.
- **Evidence:** iOS follows this (large title under the band). I could not confirm whether the live app differs from the Roborazzi frames.
- **Fix:** render the PageHeader in the content and fade `BAR_TITLE` in on scroll, per N3.

**S3. Raid and Activity feed times use bucketed formats the web does not use, and they differ on the same fixture.**
- **What is wrong:** both apps use the same bucketed rule (today `8 h ago`, yesterday `6:02 pm`, this week `Tue`, earlier `3 Nov`). The web home feed shows `10 h ago` on every row (`web-home-dark.png`).
- **Evidence:** the rule is in `apps/ios/App/Activity/ActivityModel.swift:107` and `ActivityScreen.kt:198`. The iOS capture shows `Sat`, `Sat` because of fixture dates.
- **Fix:** decide this once. If bucketing stays, record it as a ruling. Both platforms agree with each other, so the drift is against the web only.

**S4. Home's type roles are one-off sizes instead of tokens.**
- **What is wrong:** each of these is a copy of a token with a local override.
  - Android: `HomeHero.kt` (`fontSize = 30.sp`, `24.sp`, `14.sp`), `HomeRaidCard.kt` (`lineHeight = 38.sp`, `12.sp`), `GearFlowScreen.kt` (`14.sp`), `ProfileMenu.kt` (`26.sp`, `16.sp`, `12.sp`), and `RaidTabParts.kt` (`13.sp`).
  - iOS: `NextRaidBlock.swift` uses an inline `TypeRole(face: FSFont.title, size: 24, lineHeight: 30, scalesWith: .title2)`, and `ActivityScreen.swift:165` uses `.font(.custom(FSFont.bodySemibold, size: 14, ...))`.
- **Why it matters:** these happen to match today, but a scale change in `design/tokens/build.mjs` will not reach them. This is the hardcoded drift the review brief asks about.
- **Fix:** add the named roles (hero name, card title, fact line, unit) to `design/tokens/tokens.json` and regenerate `DesignTokens.swift` and `Tokens.kt`.

**S5. A hardcoded colour in the Android Gear code.**
- **What is wrong:** `apps/android/app/src/main/kotlin/gg/foreversixty/app/paperdoll/Paperdoll.kt:165` has `Color(0xFF14171F)`. `core/designsystem/GameColors.kt:45` has `parsePinkText = Color(0xFFF08CC0)`.
- **Fix:** move both into the token pipeline (`tokens.json` into `Tokens.kt`).

**S6. The Raid tab shows counts differently between the platforms.**
- **What is wrong:** iOS says `Tanks 3`, `Healers 6`, `Damage 15`. Android says `Tanks 2 of 2` and `Healers 5 of 8`, tone-coloured green or red, with `· 3 short`.
- **What is ruled:** web counts are `3 of 4` and `3 tanks` (`web-home-mock`).
- **Fix:** put one tally model in the shared copy. A quota-less count may read plainly, but the needed and short counts and the tone must match.

**S7. "Equipped" appears on Android Gear rows and not on iOS.**
- **What is wrong:** `android-gear.png` marks `Neck · Equipped`, while iOS shows only the slot or the gain.
- **What the web does:** the web character paperdoll shows the slot only.
- **Fix:** choose one row line. Slot plus gain is the web's; drop `Equipped`, or add it on the web.

## NICE

- **N-1:** the iOS drawer line reads `Fury Warrior · 58 · OLYMPUS XXVII`. The drawer takes the guild name straight from the fixture, which hardcodes `OLYMPUS XXVII` in `apps/ios/AppTests/System/SystemTestSupport.swift`. There is no `.uppercased()` in the drawer code, so this is a fixture artifact, not an N1 breach. Use `Olympus XXVII` in the fixture so goldens do not teach the wrong case.
- **N-2:** drawer row anatomy. The web account menu has icon tiles per row, a `Premium` pill, a plan value at the right and a Sign out row. The Android drawer has two-line rows (`Notifications / Raids, gear and sims, logs on`) and puts Sign out under Settings. The iOS drawer shows a value only on Plan. Align the sublines between the platforms (they are the same destinations). The web icon tiles are ALLOWED to differ under rule 8.
- **N-3:** the Android notification buttons are fully round pills. That is the M3 notification idiom (rule 8) and passes.
- **N-4:** the Search button is visible in the iOS tab headers but absent from the Android Roborazzi frames. Spec R1 says Search stays a toolbar button. `ForeverAppScaffold` has a `RootSearchBar`, so this may be a capture gap. Confirm it in the emulator.
- **N-5:** the Android panel header on Home and Raid uses a portrait where the web uses the crest. This is covered by B3.

## ALLOWED differences (not drift)

- The floating glass iOS tab bar and the Material 3 `NavigationBar`: spec 4.1 and rule 8.
- iOS large title versus Android top bar: spec N3 and rule 8, subject to S2.
- Pull to refresh, haptics, safe areas, `-x2_0` and `x1_5` font scaling (rule 8).
- Sentence-case labels (N1). I found no forced caps in the four tabs' code. The `uppercased()` uses sit in the raid page, Roster and BiS screens, which are out of scope, and they are retired styles for the other lane.
- Sheets for tooltips and the Late chips: spec row 155.
- Raid-day order on Home (raid card first on raid day, upgrades first otherwise): spec R0, R2 and the "Any other day" table.
- Light mode: not applicable, since the app and the site are dark only.

## Do not re-report (fixes in progress, taken as read)

The iOS `Raids start at 60` text and `Every slot matches` text, and the iOS person glyph in the profile button. The Android list in the BRIEF, none of which I re-reported.

VERDICT: DRIFT

## Round 2 (2026-11-10, base 9bbf15e37)

# Round 2 verdict: brand consistency, both apps vs the web

I judged the current captures, the Android parity goldens, the token sources and the code. I did not run either app. Android `notif-02-expanded.png` is a black loading frame, so I judged the Android push from `notif-01-shade.png`. The `light-*` captures in `ios/` and `android-r2/` are identical to the dark ones, which confirms the app is dark only.

## Round-1 findings

| # | Status | Notes |
|---|---|---|
| B1 iOS answers as four buttons | FIXED | `FSSegmented` is one connected group with a 4 pt radius, hairline dividers and a teal selected segment. It is used on Home and Raid, stacks at accessibility sizes (`ax5-raid.png`), and exposes a Picker to VoiceOver. |
| B2 Android Segmented shape and colour | KEPT for the pill and the teal-chosen idea; the fill token is a NEW finding (D1) | I accept the round-5 mock ruling (`.r5raid .seg`) and the M3 pill under rule 8. The web's cyan underline versus teal fill is a mock-ruled translation. The fill colour is a separate problem. |
| B3 Portraits | KEPT | The owner exception in app-flows lines 142-146 and tenet 2 answers it. The captures follow it: the button and drawer head show the person glyph in the faction ring with the badge, Home's head is the crest, and the Switch rows show the crest. Gear is the seal on both apps. |
| S1 Push time format | FIXED | iOS reads `Sign-ups are open: Thu 5 Nov, 8:00 pm CST. 24 of 40 signed.` The Android shade shows the same text, truncated. |
| S2 Android bar title at rest | KEPT | The round-5 mocks and M3 top app bar decide it (rule 8). |
| S3 Feed time buckets | KEPT | The ruling is accepted, but Android breaks it on Home (see D3). |
| S4 One-off type sizes | STILL OPEN, partly | The named roles exist and are used on the main paths. The residue is in D4. |
| S5 Hardcoded colours | FIXED | `Paperdoll.kt` and `GameColors` no longer carry hex. The only hex left in the non-token Kotlin is `AndroidCalendar.kt:116`, the system calendar colour, and iOS has the same in `EventKitCalendarAdapter.swift:36`. That is a NICE token candidate. |
| S6 Tallies differ | FIXED on Raid, STILL OPEN on iOS Home | The Raid tab tallies read the same on both apps. iOS Home is new drift (D2). |
| S7 `Equipped` | FIXED in code | `gear.worn` is a shared key, used by both platforms and shown in the Android golden. The iOS tab capture has no worn pick, so I could not see it on screen. |
| N-1 `OLYMPUS XXVII` | STILL OPEN, NICE | It still shows in the Activity feed body (`Loot awarded` row, both apps). The server text or the fixture should read `Olympus XXVII`. |
| N-2 Drawer sublines | FIXED | Both drawers carry Plan `Free` and Devices `3 signed in · newest <device>`. |
| N-4 Search button | FIXED | Android has the search glyph in the bar. |

## New findings

### SHOULD (noticeable to a regular)

**D1. The selected answer is two different teals across the apps, and on Android two inside one card.**
- **Where:** `FlowSegmented.kt` fills the chosen segment with `FsColor.tealEdge` (#408076, pale). iOS `Segmented.swift` fills with `FsColor.teal` (#0f6064) and uses `tealEdge` only as the hairline. The web primary uses the dark `teal` fill with the `tealEdge` border.
- **What a player sees:** on Android the `In` chip on Raid is dark teal and the segment under it is pale teal. iOS shows one teal.
- **Fix:** `ChoiceColors.ENABLED.selected` becomes `FsColor.teal`, with the selected segment's edge in `FsColor.tealEdge`. The pill shape stays under the ruling.

**D2. iOS Home's raid card drops the role tally when the raid has no target. Android shows it.**
- **Where:** Android Home reads `Tanks 3 · Healers 6 · Damage 14`. iOS Home (`NextRaidBlock.swift` line 141, `HomeFlowPresenter.tally`) shows nothing.
- **Ruling broken:** the S6 decision says `Healers 6` without a target, on both platforms. iOS's Raid tab does this correctly.
- **Fix:** have `HomeFlowPresenter.tally(_:)` fall back to plain counts, the same path as `RaidRootView`.

**D3. Android Home reads `· 3 d ago` where Activity and iOS Home read `· Sat`.**
- **Ruling broken:** the round-2 decision says Home's `New for you` rows use the same buckets "so one row never reads two ways". Android Home's Molten Core row says `3 d ago`.
- **Fix:** route Android Home's `New for you` age through the same bucket formatter as `ActivityScreen.kt:198`.

**D4. Raid counts and the Covered line disagree between the apps on the same fixture.**
- **Damage split:** Android's Raid tab says `Damage 15 · 6 melee, 8 ranged`, which adds to 14, and Android Home says `Damage 14`. iOS says `Damage 15`, `7 melee, 8 ranged`. One of them miscounts, probably a Late or Bench boundary, and the Android split does not add up to its own total.
- **Covered line:**
  - iOS reads `Missing Blessing of Kings, Blessing of Might, Blessing of Wisdom and 2 more` in muted text.
  - Android lists six `Missing …` items in wipe red, and the last one is a bare `Missing Blessings`.
  - The decision for G4 is one line, `Missing Fortitude · Windfury 1 of 2`.
- **Fix:** one shared tally and Covered builder, fed by the same `counts` and `composition` fields, with one tone.

**D5. Android has no stacked fallback for the answer control at large font.**
- **Where:** `android-r2/fs2-raid.png` shows `Bench` pressed against both dividers at scale 2. iOS stacks the segments at accessibility sizes.
- **Fix:** in `FlowSegmented`, stack the segments or use a 2×2 grid when the font scale is at least 1.5. Follow the existing `isLargeText()` helper.

**D6. Type residue the S4 fix missed, on the four tabs and the drawer.**
- **iOS:**
  - `NextRaidBlock.swift:82` uses `TypeRole(... size: 36, lineHeight: 38 ...)`. Android's countdown is `FlowCountdown.kt:138`, a local `CLOCK` of 40/42. Same job, different size, neither tokenised.
  - `NextRaidBlock.swift:257` and `HomeFlowView.swift:198` use `.custom(... size: 14 / 15)`.
  - `ProfileMenu.swift:221, 248, 340, 363, 372, 389` use inline `TypeRole`. The `drawerName` and `drawerSection` roles exist and Android uses them.
- **Android:**
  - `ProfileButton.kt:96` has `BAR_TITLE = 22.sp/28.sp`.
  - `Paperdoll.kt:167-168` has 13/12.sp.
  - `FlowRows.kt:251` has a 13.sp unit where `FsType.figureUnit` exists.
  - `FlowRows.kt:356, 367` have 13.sp.
- **Fix:** add roles for `countdown` (the token exists, so use it on both), `barTitle` and `paperdollName/Detail`, and replace every inline size above with its role.

### NICE

- **D7. The synced stamp sits in different places on Home.**
  - iOS puts `Synced 12 d ago` at the top of the upgrades card. Android puts it at the bottom.
  - The decision says the stamp goes under the section meta, so Android's bottom position is wrong.
  - Move both to just under the `13 upgrades · +157.6 DPS` meta line.
- **D8. `Bring flasks` sits inside the raid card on iOS and outside it on Android.** Pick one placement. The mocks decide which.
- **D9. The feed body and the push write the same raid time two ways.** The Home row says `Thursday 8:00 PM.` and the push says `Thu 5 Nov, 8:00 pm CST`. Use the shared layout (`Mon 2 Jan, 3:04 pm MST`) in the feed text too.
- **D10. Activity's `Show older` is on Android and not in the iOS capture.** The BRIEF lists feed paging as a fixture limit, so I only mention it here. Check it on the next iOS capture.

## ALLOWED differences (cited)

- The floating glass iOS tab bar against the M3 `NavigationBar` and its pill indicator: spec 4.1 and rule 8.
- iOS large title under the band against the Android bar title at rest: N3, the round-2 KEEP, rule 8.
- Android's Segmented as one 48 dp pill against iOS's 4 pt radius, 44 pt group: the round-5 mock ruling and rule 8.
- Android push buttons are `In`, `Late`, `Bench` (three), because M3 notifications cap actions at three: rule 8.
- Round pill action buttons and the system notification chrome: rule 8.
- The Android teal LX small icon against the iOS app-icon square: rule 8.
- Pull to refresh, haptics, safe areas, `ax5-*` and `fs2-*` scaling: rule 8.
- The person glyph in the faction ring on the profile button and drawer head, and the crest on Home and the Switch rows: the owner exception, KEPT.

VERDICT: DRIFT

## Round 3 (2026-11-10, base 7911bf87e)

# Round 3 verdict: brand consistency, both apps vs the web

I judged the iOS step captures (including `light-*` and `ax5-*`), the Android `android-r3/` captures (including `fs2-*`), the parity goldens, the round-3 decisions and the code. I did not run either app. The iOS and Android `light-*` captures are identical to the dark ones, so the app is dark only, as the site is. I judged the Android push from `notif-01-shade.png`. The iOS Gear counter (`0 of 17` vs `3 of 17`) and the Gear `Equipped` cells are known and not re-reported. I did not open the round-3 mocks or the web captures, so web comparisons rest on the earlier rounds and the decisions file.

## Round-2 findings (D1-D10)

| # | Status | Evidence |
|---|---|---|
| D1 Two teals on the selected answer | FIXED | `Segmented.swift:83-84` and `FlowSegmented.kt` both fill `FsColor.teal` inside a `tealEdge` edge. Home and Raid read the same teal on both apps. A new teal split appears on the Activity chips (N1). |
| D2 iOS Home drops the tally | FIXED | iOS Home reads `Tanks 3 · Healers 6 · Damage 15`, the same as Android. |
| D3 Android Home `3 d ago` | FIXED | Both Home feeds read `Sat` on its own meta line, the same as Activity. |
| D4 Counts and Covered line | FIXED for counts, PARTLY OPEN for the Covered line | `Damage 15` equals `7 melee, 8 ranged` on Home and Raid, on both apps. `Missing` appears once, at most three names and `and 2 more`, with wipe red on the names only. What is left is N3. |
| D5 Android no stacked answer at large font | FIXED | `fs2-raid.png` stacks to 2×2. iOS stacks to one column at `ax5` (N5, NICE). |
| D6 Type residue | PARTLY OPEN | iOS Home, the drawer and the countdown are clean: `NextRaidBlock.swift` and `ProfileMenu.swift` now use roles. Residue is listed in N4. |
| D7 Synced stamp position | FIXED for position, tone differs | The stamp sits under the meta line on both apps. The tone differs (N2). |
| D8 `Bring flasks` placement | FIXED | It sits inside the raid card on both Raid tabs. |
| D9 Feed time layout vs push | FIXED | The feed and the push both read `Thu 12 Nov, 8:00 pm CST`. The push body leads with the count, `24 of 40 signed.`, on both platforms (decision s3). |
| D10 iOS `Show older` | FIXED | It is on iOS Activity. Alignment differs (N8). |
| Carried N-1 `OLYMPUS XXVII` | STILL OPEN, NICE | The `Loot awarded` body reads `by OLYMPUS XXVII` on both apps. Fix the server text or the fixture to `Olympus XXVII`. |

The KEEPs hold, and I do not contest them:
- The person glyph in the faction ring (owner exception).
- The Android bar title at rest (the round-5 mocks and rule 8).
- Liquid Glass on the iOS tab bar and the M3 `NavigationBar` (spec 4.1, rule 8).
- The pill `FlowSegmented` against the 4 pt `FSSegmented` (mock ruling, rule 8).
- Three notification actions on Android (rule 8).

## New findings, ranked

### SHOULD

**N1. The selected filter chip is a different teal on each app (DRIFT).**
- **Where:** iOS `ActivityScreen.swift:179` fills `FSColor.tealEdge` (pale `#408076`) for the selected chip and while pressed. Android `FlowChip` fills dark teal. `activity-01-top.png` shows both.
- **Why a player feels it:** this is the D1 split again, one control over from the Segmented. The web primary is `teal` fill with a `tealEdge` border.
- **Fix:** in `ChipStyle`, use `FSColor.teal` as the selected fill with a `tealEdge` hairline. Keep the pressed state as the `tealEdge.opacity(0.25)` overlay that `Segmented.swift:94` uses.

**N2. The stale `Synced 12 d ago` stamp is muted on iOS Home and warning-orange on Android Home.**
- **Where:** iOS `Panels.swift:165` always draws the stamp in `FSColor.muted`. Android passes `stampWarn = block.syncedStale` (`HomeFlowScreen.kt:115`). The iOS drawer already shows the same age in warning orange (`menu-01-open.png`).
- **What is ruled:** round-3 change 12 sets the warning tone past 7 days. The fixture is 12 days stale, so Home should show it.
- **Fix:** give `PanelSectionHead` a `stampStale` flag. Fed from the Presenter, it draws the stamp in the drawer's warning token. The Presenter already has the sync age.

**N3. The Covered line has broken or missing spell icons, and its punctuation differs.**
- **Where:** `raid-01-top.png` on iOS draws empty bordered boxes beside `Blessing of Might` and `Blessing of Wisdom`. Android leaves a blank gap in the same places. Only Kings has an icon. iOS adds trailing commas, and Android does not.
- **Also:** decision 4 asks for the data's short `label` (`Kings`, `Might`...) where one exists, and the line still shows the long names. That is only a drift if the fixture carries labels, so check that.
- **Fix:**
  - Render no icon frame when the stem is missing, on both apps. Or ship the three icons as faithful HD reproductions per the owner's icon ruling.
  - Use the shared Covered builder's output for the separators on both apps.

**N4. Type residue remains on the four tabs.**
- **iOS:** `GearListView.swift:97` and `:107` use inline `TypeRole(... size: 13 ...)` and `20` for the Gear rows. These are the slot and gain lines on a tab root.
- **Android:**
  - `FlowFigures.kt:45` and `:55` have 13.sp and 12.sp.
  - `FlowLayout.kt:304` and `:307` have `ACTION` 14.sp and `PANEL_TITLE` 21.sp.
  - `FlowRows.kt:338` and `:345` carry local `lineHeight` copies.
  - `ForeverTheme.kt:52` has `TOP_BAR_TITLE = 22.sp`.
  - `PaperdollHeader.kt:301-304` has 19, 17 and 11 sp.
  - `ScaledType.kt` carries literal sizes (`cardHead`, `lead`, `segment`).
- **Why it matters:** the Android Raid card title is visibly smaller than the iOS one on the same card, because the iOS Home title wraps to three lines where Android's fits one. A scale change in `design/tokens/build.mjs` will not reach any of these.
- **Fix:**
  - Add roles to `design/tokens/tokens.json` for `cardTitle`, `figureUnit`, `paperdollName`, `paperdollDetail`, `barTitle` and `panelTitle`.
  - Regenerate `DesignTokens.swift` and `Tokens.kt`.
  - Point every site above at its role, and make the iOS and Android Home raid titles use the same `cardTitle`.

### NICE

**N5. The answer control stacks differently at large type.** iOS `ax5-raid.png` goes to one column of four. Android `fs2-raid.png` goes to 2×2. Round-3 change a2 said "as iOS does", so pick one. I would take the 2×2 on both. The pill outline around a 2×2 on Android has a rounded grid corner that the web never draws, so set the stacked form's radius to 4.

**N6. The Damage row is laid out two ways.**
- **iOS:** `Damage 15`, then the names, then `7 melee, 8 ranged` on its own line.
- **Android:** `Damage 15 · 7 melee, 8 ranged`, then the names.
- **Fix:** take the Android form. The role label, count and split read as one line, and the names stay the lead.

**N7. The Home tally touches the Segmented on iOS.** `Tanks 3 · Healers 6 · Damage 15` sits flush on the control's top edge in `home-01-top.png`. Android leaves a gap. Add the same 8 pt spacing in `NextRaidBlock.swift`.

**N8. `Show older` is centred on iOS and left-aligned on Android.** The web's pager link is left-aligned under its list. Left-align it on iOS, with a 44 pt target on both.

## ALLOWED differences (cited)

- The floating glass iOS tab bar against the M3 `NavigationBar` and its pill indicator: spec 4.1, rule 8.
- iOS large title under the band against the Android bar title at rest: N3, rule 8.
- The M3 pill answer control and the three notification actions: rule 8, the round-5 mock ruling.
- The system notification chrome (the Android teal LX small icon, the iOS LX tile): rule 8.
- The person glyph in the faction ring: the owner exception, KEPT.
- Safe areas, haptics and `ax5-*` and `fs2-*` scaling: rule 8.

VERDICT: DRIFT

## Round 4 (2026-11-11, base 693e7a952)

Round 4 verdict on the four tab roots, the profile drawer and the raid notification, judged against the web. I read the iOS captures, the Android r4 captures, the code and the web tokens. I did not open the Android parity goldens or the round-4 mocks.

Two Android captures are not usable: `android-r4/gear-01-top.png` and `android-r4/activity-01-top.png` show the system notification shade, not the app. I judged Android Gear and Activity from `light-gear.png` and `light-activity.png`. The `light-*` sets are still the dark app, which matches the brief's "dark only". Redo the two bad frames.

## Round-3 findings

| # | Status | Evidence |
|---|---|---|
| N1 Activity chip teal | FIXED | iOS `activity-01-top.png` fills the selected `All` chip with `teal` inside a `tealEdge` hairline. It matches Android and the Segmented. |
| N2 Stale tone | FIXED | Home `Synced 12 d ago` and Gear `synced 12 d ago` are warning orange on both apps, and the drawer matches. |
| N3 Covered line | FIXED, one NICE residue | iOS `light-raid.png` reads `Missing [Kings icon] Kings · Might · Wisdom and 2 more` with no empty frames and one separator. Might and Wisdom lose their art, which is the known fixture limit. The Kings label sits a few pixels lower than Might and Wisdom, so the baseline shifts when an icon is present. Align the icon on the text baseline. |
| N4 Type residue | FIXED on the files I could check | The named Android sites (`FlowFigures.kt`, `FlowLayout.kt`, `ForeverTheme.kt`) no longer have literal `.sp` sizes. One remains: `FlowRows.kt:348` has `FsType.small.copy(lineHeight = 18.sp)`. That is a role with a line-height override. Fold the 18 into a token role (NICE). I could not find `GearListView.swift` at the cited path and did not re-verify it. |
| N5 Large-text stack | FIXED in code, not seen | The mock rulings and `fs2-raid.png` use the 2×2 form at radius 4. I did not open `ax5-raid.png` on iOS. |
| N6 Damage split | FIXED | `Damage 15 · 7 melee, 8 ranged` is one line on both apps, with the names under it. |
| N7 iOS tally spacing | FIXED | There is a gap above the Segmented on iOS Home, as on Android. |
| N8 Show older | FIXED | It is left-aligned under the list on both apps. |
| Carried: `OLYMPUS XXVII` | KEPT | Both apps still read `by OLYMPUS XXVII.` Round-4 decisions KEEP it as the guild's own typing, and I accept that. |

## New findings

**SHOULD**

**R4-1. Preview mention is a blue button on iOS and a neutral grey outline on Android (DRIFT).**
- **iOS:** `NextRaidBlock.swift:144` uses `FSButtonStyle`, whose `.secondary` kind is `blue` with a `blueEdge` hairline. That is the web's secondary button, per `web/src/styles/tokens.css:119`.
- **Android:** `FlowRows.kt:269-276` draws every non-primary `FlowButton` as `raised` with a `line` border. `HomeFlowScreen.kt:132` uses it for the same action. It is the same job and the same screen, and `raid-officer-01-top.png` shows the two differ.
- **Fix:** add a secondary kind to `FlowButton` that reads the generated `FsColor.blue` and `FsColor.blueEdge` tokens. Make `FlowButton`'s non-primary default that kind.

**R4-2. Android's primary button is a gold gradient (DRIFT).**
- **Where:** `FlowRows.kt:269` fills primary buttons with `Brush.verticalGradient(goldHover, gold)` and `bg` text. `ApplicantScreen.kt:85` (Accept) and `NotAnsweredScreen.kt:108` (the role Mention button) use it.
- **Web and iOS:** the web primary is `teal` with a `tealEdge` border. iOS `FSButtonStyle.primary` uses the same. The web draws no gradients.
- **Scope:** the screens that use it are secondary, but the notification actions land there. This is a shared component, so it will spread.
- **Fix:** primary becomes `FsColor.teal` fill, `FsColor.tealEdge` border and `strong` text, with no brush.

**R4-3. The iOS Activity tab badge does not show at launch (DRIFT).**
- **iOS:** `RootView.swift:195` sets `shell.badges[.activity]` only from `ActivityScreen`'s `onUnread`. iOS Home, Raid and Gear show no badge on Activity, although Home's own link reads `All activity (3)`.
- **Android:** the badge `3` shows on every tab.
- **Fix:** load the unread count at shell start, the way line 117 loads the Raid badge. `RaidBadgeCount.load` is the pattern.

**NICE**

**R4-4. The Raid badge reads 6 on Android's officer view and 1 on iOS's.** I cannot say whether these are different states. `ios/r4-officer-raid-01-top.png` shows 1 and `android-r4/raid-officer-01-top.png` shows 6, both on the officer state. Check that both build the count from the same rule.

**R4-5. The two apps use different clocks in the captures.**
- **Android:** `Today / Molten Core / just now`, then `This week`.
- **iOS:** one `This week` group, with `Thu`.
- **Fix:** pin one fixture clock so the grouping and stamps compare.

**R4-6. The stamp meta sits alone on its own line on both apps in the Activity rows.** It reads `…CST.` then `Thu`. The web keeps the meta on the same line when it fits. This is the same on both apps, so it is consistent with the brief's ruling and I do not list it as drift.

## ALLOWED

- The floating glass iOS tab bar against the M3 `NavigationBar` and its pill indicator: spec 4.1, rule 8.
- The iOS large title against the Android bar title at rest: rule 8.
- The M3 pill answer control and the three notification actions: rule 8, the round-5 mock ruling.
- The person glyph in the faction ring: the owner exception.
- The `OLYMPUS XXVII` loot body: the round-4 KEEP.
- The `Panels` on list blocks: the round-4 KEEP, and the web home and guild pages draw the same.

The two button findings are real component drift. A regular who opens the app after the site will feel both. Only CONSISTENT passes the gate.

VERDICT: DRIFT

## Round 5

Round 5 brand-consistency verdict on the four tab roots, the profile drawer and the raid notification: I find no BLOCKER or SHOULD drift. I checked two parts of the screen set directly and the rest only through the decisions files and the code, so the verdict rests on that scope.

**What I looked at.** I opened the iOS and Android Home captures and the officer Raid captures on both apps. I read `round-4-decisions.md` and `round-5-decisions.md`, my round-4 findings, and the cited code. I did not open the Gear, Activity or drawer captures, the Android parity goldens or the web captures one by one.

**Round-4 findings**
- **R4-1 (Preview mention, secondary button): FIXED.** In `android-r5/raid-officer-01-top.png` and `ios/r4-officer-raid-01-top.png` both apps draw Preview mention as the same blue fill with a blue edge and white label. `FlowRows.kt:260-272` now maps non-primary to `BrandButtonKind.Secondary`.
- **R4-2 (Android gold gradient): FIXED.** I found no `Brush` or gradient in `FlowRows.kt`; primary is `BrandButtonKind.Primary`, the shared token kind.
- **R4-3 (iOS Activity badge at launch): FIXED.** `RootView.swift:121-123` loads `ActivityBadgeCount` at shell start. The iOS Home capture shows the Activity badge `3` and the Raid badge `1`.
- **R4-4 (Raid badge, 6 on Android and 1 on iOS): FIXED on the officer state.** Both officer Raid frames read `6`. The Home frames read `1` on both apps for the member state.

**Home, side by side.** Home matches between the two apps and the web's roles:
- Serif section heads, a gold `Gear ›` link and a gold `All activity (3) ›` link.
- Quality-coloured item names with plate icons, and `+2.8 DPS` as a figure plus a muted unit.
- The `In` segment filled in teal.
- The circular ringed crest, and the same tab glyphs on both apps.
- Whole-number figures where the house format asks for them. `+18.0 DPS` stays one decimal because it is under 100.

**Colour.** My grep for hardcoded colours in the app sources found literals only in widget snapshot tests and the calendar adapter. `EventKitCalendarAdapter.swift:36` sets a literal `CGColor` for the calendar colour, which is a system calendar and not app UI.

**ALLOWED differences**
- The iOS floating tab bar and large title, against the Android `NavigationBar` and bar title (rule 8).
- The pill segmented control on Android, against the rounded rectangle on iOS (rule 8).
- Panels on the list blocks, and `OLYMPUS XXVII` as the guild typed it (both KEEP rulings).
- The age on its own meta line under feed rows (KEEP, round 3 decision 3).
- Android's three notification actions (platform limit).

**NIT**
1. **Feed clocks differ between the apps.** The Home feed row on Android reads `just now` where iOS reads `Thu`. This comes from the capture clocks, which my round-4 R4-5 flagged. It is a fixture difference, not app behaviour. Pin one fixture clock so a future comparison lines up.
2. **The Home raid title wraps on iOS and not on Android.** `Onyxia's Lair · Sat 19 Dec, 7:00 pm CST` breaks across two lines on iOS and fits on one line on Android. The brief already lists this as a dated known drift, so I am not counting it.
3. **App raid clocks still read `7:00 pm`**, while server text follows the house format and drops `:00`. The brief assigns that to the format lane, so I am not counting it.

VERDICT: CONSISTENT

## Round 6

Round 6 holds up on every item it changed, and I found one carry-over figure in the stale Gear frame that breaks the brief's own house format. I checked each change against `apps/shared/copy/en.json`, the shared copy keys, the app code and the captures. I did not open the `web/` 390 px captures. The Gear comparison below is made against the house-format ruling in `BRIEF.md`.

**What round 6 changed**
- **Drawer Notifications value:** FIXED. It uses the shared keys `menu.notifications.allOn`, `.someOff` and `.off`.
  - `NotificationsLine.kt` builds the value from those keys. The first group name stays capitalised, the rest are lowercased, and the old `menu.notifications.on` key is gone.
  - The iOS drawer reads `Gear and sims, guild ladder, tiers off` and the Android drawer reads `All on`. The two differ only because the fixtures differ. The wording rule is the same on both.
  - Type, hairline rows, the serif section heads and the gold `Refresh` link match between the two apps.
  - Both drawers read `Updated 9 h ago from Battle.net`, so the stale-APK problem from round 5 is gone.
- **Filtered roster plate:** FIXED. Both apps show one population, `4 of 38` with `on the roster`, and the guild-wide `in game, no account` and `waiting` figures are gone. The gold filter chip with its close glyph, the crest tiles, the quality-coloured names and the attendance pips are identical across the two apps. The back and search controls are platform idiom (rule 8).
- **Push failure notifications:** FIXED. Both apps print `That did not save. Try again.` (`common.failed`) with `Open the application.` (`push.applicant.failed`). That is the shared copy, with no rewritten strings and no em dashes. The LX seal is used as the notification icon.
- **Stale Gear frame:** FIXED for the stale state. Both apps show a warning-coloured `synced 9 d ago` on the meta line. The same warning colour shows on iOS Home as `Synced 37 d ago`, a stale age the brief already dates. Item names carry their quality colours and the icons have quality borders.

**BLOCKER**
- None.

**SHOULD**
1. **`182.0 DPS` in the Gear header on both apps** (`android-r6/gear-01-stale.png`, `ios/r6-gear-01-stale.png`). The brief says figures are whole from 100, and Activity now reads `204 DPS`. Gear is a three-digit figure and still prints one decimal.
   - Fix: route the Gear header through the same house number formatter that Activity uses (`apps/shared/formats.json`, `numbers`).
   - Round 6 did not touch this code. It only shows up because the stale Gear frame is new.

**NIT**
1. The iOS Gear `Main hand · Equipped` cell wraps at phone width. This is a KEEP ruling and I am not counting it.
2. The Android notification shade shows the emulator's `Serial console enabled` line. It belongs to the emulator, not the app.
3. The app clocks still read `7:00 pm`. The brief assigns that to the format lane.

VERDICT: CONSISTENT

## Round 7

**Round 7 brand-consistency: Gear, Home, Where line, push failures, drawer Notifications**

I viewed these captures: android-r7 gear-01-top, gear-02-sheet-where, push-raid-short-mention-failed and drawer-01-notifications-some-off, plus ios r7-gear-02-sheet-where-settled and r7-home-01-top. I did not open the other r7 captures (the accept-failed push, drawer all-on and all-off, iOS push and menu frames). I checked copy keys in `/Users/jh/code/forever/.claude/worktrees/agent-aee50bfb43330b1e0/apps/shared/copy/en.json` and `apps/shared/formats.json`. I did not compare against the web 390 px captures or `design/tokens/tokens.json`, and I did not grep the app sources for hardcoded colours or strings. This verdict rests on the screenshots and the copy and format files only.

**BLOCKER**
- None.

**SHOULD**
- None. Round 6 S1 is fixed. Gear now reads `182 DPS` on both platforms, and the whole-from-100 rule matches `formats.json`. Home and Gear agree on `+18.0 DPS` and `13 upgrades`.

**NIT**
1. **Where line differs between platforms.** On Android the sheet row has the source as its title (`Crafted: Blacksmithing`) with `Where to get it` as a grey sub-line (`gear.sheet.where`). On iOS the row is just `World drop`, with no sub-line.
   - Fix: pick one anatomy and apply it on both. The label as a sub-line reads like a heading demoted to caption, so I'd drop it on Android to match iOS.
   - These were different items, so confirm the iOS crafted case before you change anything.
2. **Where line uses plain source strings, not the gear.where keys.** The `gear.where.*` keys (`drop`, `boss`, `ask`, `rep`) have the shape `{items} · {boss}, {chance}`. The sheet shows bare `World drop` and `Crafted: Blacksmithing`. Check that the sheet reuses `bis.source.crafted` and the shared source strings, not locally built ones.
3. **Zero gain shows as `+0.0 DPS` on Shoulder.** This is a drift risk if the web prints a neutral `0` or nothing for a zero gain. Check the web BiS and gear row.
4. **Header sub-line casing.** `synced 10 h ago` is lowercase mid-line (`gear.header.synced`), while Home reads `Synced 10 h ago` as a standalone line. That is the same fact in two casings, and each follows its key's position. It is acceptable.
5. **Home raid line.** It still shows `7:00 pm CST` and the duplicated `Thu`. The brief lists the app clock as the format lane's work and the duplicate day as a round 6 KEEP, so I did not count either.
6. **Push failure copy.** I checked only the Mention failure in the capture. It matches the `push.short.failed` and `push.short.retry` keys exactly, uses no em dash, and the title reads `Could not mention the healers in Discord`. Ruled OK.
7. **Drawer Notifications value.** `Gear and sims, guild ladder, tiers off` is sentence case and states what is off. Ruled OK.

All round 7 changes use shared keys and house formats, and the cross-platform difference in nit 1 is small and fixable in one line.

VERDICT: CONSISTENT

## Round 8

Round 8 holds up against the web. I found no BLOCKER and no SHOULD. I did not read the ios `r8-gear-sheet-upgrade` or the decline-failure frames.

**What I checked**
- iOS `r8-gear-sheet-crafted-boss-runner-up-settled` and `r8-push-short-fail-03-result`.
- Android `gear-02-sheet-crafted-skill` and `gear-04-sheet-upgrade`.
- Code: `apps/ios/App/Items/ItemVocabulary.swift`, `apps/ios/App/Items/ItemSheet.swift`, `apps/android/.../gear/flows/GearComposer.kt`, `GearItemSheet.kt`, and `apps/shared/copy/en.json`.

**Source row (ALLOWED, round-7 nit closed)**
- The row is now one anatomy on both apps: the source alone, with its figure and no caption or item icon, so it no longer reads as an item row.
  - The crafted pick reads `Crafted: Blacksmithing (245)`. That is the shared key `bis.source.craftedSkill`, not a rewritten string.
  - The vendor upgrade reads `Vendor: Lady Palanseer`, from `bis.source.vendor`.
  - The boss-drop shape (`Onyxia, Onyxia's Lair, 8%`) is built in code on both apps. I confirmed it in `ItemSheet.swift:142`. I could not confirm it from a capture, because the brief says none shows a boss drop.
- Both apps now carry the same row order and the same runner-up line. That answers the round-7 note on the Where row.

**Drop chance (ALLOWED)**
- iOS `ItemNumbers.percent` and Android `chanceWord` both match the web's `dropPercentLabel` in `web/src/lib/items/drop-chance.ts`.
  - Both give a whole percent, half up, and `<1%` under one.
  - Both cite the web function in their comments.
- iOS uses `.toNearestOrAwayFromZero` and Android uses `roundToInt`. For positive chances both agree with the web's `Math.round`.

**`No gain` (ALLOWED, round-7 nit closed)**
- The text comes from `gear.cell.noGain` on both apps, and Gear's Shoulder cell reads `Shoulder · No gain`.
- A real negative gain still prints with a true minus (`−1.3 DPS`), as the server writes it. That is consistent with `signedTenths`.

**Mention failure (ALLOWED)**
- The title `Could not mention the healers` fits one iOS banner line, from the changed `push.short.failed` key. The body keeps the raid name and the retry.

**NIT**
- No capture shows a boss-drop row with its chance. This is declared in the brief. The shape is covered by code and parity data, so it does not block.
- Android's `Runner-up · PvP rank 11 · Blood Guard · Horde · Same DPS as the pick` wraps to a second line, and iOS wraps it too. That is the platform's width, not drift.

VERDICT: CONSISTENT

## Round 9

Round 9 is consistent with the web. I found no BLOCKER and no SHOULD, only two NITs.

I judged from the iOS `r9-gear-sheet-crafted-pvp-runner-up-settled.png` and `android-r9/gear-02-sheet-crafted-skill.png`. I did not open the other two sheets, the Android goldens `tooltip-weapon` and `tooltip-large`, or the 390 px web captures. For the web side I read `web/src/lib/bis/copy.ts`, `web/src/lib/bis/doll-view.ts`, `web/src/lib/bis/frame-view.ts` and `web/src/lib/ui/paperdoll.ts`.

**Matches the web**
- **Copy key:** `Same DPS` comes from the shared key `items.sameDps` ("Same {unit}"). The unit word is the same one the signed deltas use. `GearItemSheet.kt` and `ItemTooltipModel.kt` read `R.string.items_sameDps`, which comes from the shared copy. iOS reads the generated `Copy.swift`, line 659. Neither app hardcodes the string.
- **Wording:** The web already says "same DPS" for a tie. It does so in `doll-view.ts` and in the paperdoll tooltip "(same DPS)". "Same DPS" as a standalone figure-like tail is the same word, capitalised for a standalone tail.
- **Row anatomy:** In both captures the runner-up row has a quality-bordered icon, a quality-coloured name, and a muted tail of `Runner-up · source · figure`. This matches the web's tooltip "next best" row, with `·` separators and no em dashes.
- **Source row:** The pick's source row reads `Crafted: Blacksmithing (245)`, which is `bis.source.craftedSkill` as on the web. The runner-up does not repeat its source when it matches the pick's. A differing source is named in the shared `bis.source.*` words.
- **Type, colour and shape:** Both captures are dark. Rows are hairline-separated rather than stacked cards. The blue name is the rare quality colour and the pick is green, as on the web. I saw no new shadows or glows, and no hardcoded strings or colours.
- **Both platforms agree:** The two platforms show the same row.

**NIT**
1. **PvP source drops the faction word.** The app reads `PvP rank 11 · Blood Guard`. The web's `pvpSourceLabel` in `web/src/lib/bis/copy.ts` still reads `PvP rank 11 · Knight-Lieutenant · Alliance`. This is an owned decision (round-8 ruling 3, from the ai-slop review), and the faction is redundant on a list already filtered to one faction. But a player who moves between the web and the app will see the same source worded two ways. Either carry the ruling into the web's `pvpSourceLabel`, or record it in the decisions as a deliberate native-only difference. If you carry it to the web, check `source-cell.test.ts` and `copy.test.ts`, which assert the faction.
2. **Tie handling differs from the web's Gear rows.** The web's `rowFigure` in `frame-view.ts` hides a tie on frame rows ("Same DPS on most rows was noise"). The app shows `Same DPS` on a runner-up. That is the correct place for it, because the web also says "same DPS" in the runner-up tooltip. I'm recording it only so nobody copies the frame-row suppression into the sheet.

VERDICT: CONSISTENT

