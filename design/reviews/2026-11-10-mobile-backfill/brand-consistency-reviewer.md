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

