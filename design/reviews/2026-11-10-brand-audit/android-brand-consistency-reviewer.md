# brand-consistency-reviewer: Android app (2026-11-10)

Scope:
- Read: the m0, m2a, m3a and m4a-m5a side-by-sides, the m2i and m3i comparisons, `unify-format/after/home-390.jpg`, the Kotlin sources, and the generated token, copy and icon files.
- Not opened: the widget captures. They are judged stale by date alone.

## DRIFT

**BLOCKER**
1. **Retired styles are still live on about half the app.** These are the styles `design/specs/2026-11-07-mobile-redesign.md:165` retired. Counts across `app/src/main` and `core/*/src/main`:

   | Retired style | Count |
   |---|---|
   | `quietCard` | 65 |
   | `primaryCard` | 22 |
   | `surfaceBand` | 12 |
   | `MicroLabel` | 62 |
   | `ScaledType.title/h2/stat/sec/cardHead` | 41 |
   | `.uppercase(` | 57 |

   - Worst `.uppercase(` sites: `raids/ui/RaidPage.kt:167,223,300,350`, `raids/ui/Roster.kt:120`, `raids/ui/RaidNight.kt:67`, `core/designsystem/.../Surfaces.kt:49,66`.
   - The m0 to m5 captures cover only Home, Raid, Gear, Activity and Menu. Guild, Rankings, Reports, Planner and Guides still show the pre-redesign language. Files still using these styles: reports 6, raids/ui 6, guides 6, search 3, rankings 3, planner 3, account 3.
   - Fix: migrate those screens to the spec 3.4 primitives, delete the retired helpers, and make a surface-guard test fail on them.
   - `home/GearScreen.kt:34` has no non-test caller. It is dead code: delete it.
2. **`m0/android-home-34.png` is broken or stale.** Item and activity icons render as empty quality-bordered squares, and the header is a bare person glyph with no crest. Re-shoot, and check the loading and error state in `items/ItemIcon.kt`.
3. **The Raid screen lacks the guild-links block and the search action** that iOS and the mock carry (`m2a/side-by-side-raid.png`). The bar title duplicates the page title (N3). iOS leaves the bar empty. Pick one rule for both apps.

**SHOULD**
4. **30 hardcoded colours.**
   - `GameColors.kt:45` (`0xFFF08CC0`), `paperdoll/Paperdoll.kt:165` (`0xFF14171F`), `guild/platform/AndroidCalendar.kt:116` (`0xFFE5B955`).
   - `Color.Black` used for hairlines at `Surfaces.kt:129,176`, `ItemIcon.kt:143` and `PaperdollHeader.kt:133`. Use `FsColor.line` instead.
   - The XML colours are not generated: `res/values/notification.xml:5,8` (`#FF2C8A9E`) and `values/themes.xml:10,12` (`#FF07090D`, `#FF408076`).
   - The parse-tier map in `reports/ui/ReportsChrome.kt:153` is local.
5. **Text shadows** at `home/flows/HomeHero.kt:72` and `PaperdollHeader.kt:156`. `SignupViews.kt:294` uses `FsRadius.legacyControl`.
6. **Type.**
   - System monospace at `ScaledType.kt:36-40`, `BisScreen.kt:203` and `TalentTree.kt:276`.
   - Widgets default to `FontFamily.SansSerif` (`widgets/glance/WidgetDesign.kt:56`).
   - 61 `.sp` literals and 1130 `.dp` literals sit outside Tokens.
7. **Icons.**
   - 228 Material icon references, for example Refresh at `HomeScreen.kt:204`, Lock, List, Check and KeyboardArrowDown at `Roster.kt:21-25`, and Check at `Surfaces.kt:65`.
   - About 40 hand-maintained `ic_*.xml` drawables (`ic_source_*`, `ic_shield`, `ic_trophy`, `ic_talents`) are not generated. `gen-ui-icons.mjs` only emits `ic_tab_*` and `ic_frame_corner`.
   - A Material back arrow on the native bar is allowed (N3).
8. **The live-raid notification diverges from iOS** (`m4a-m5a/side-by-side-lock.png`). The Android card is light grey with a black timer, while iOS is dark with a gold timer. The small icon is an olive "LX" disc (`ic_notification_seal.xml`), not the flat seal. The accent is `#2C8A9E` in place of the token accent.
9. **Item tooltip** (`m3a/side-by-side-item.png`). It drops "drop chance" and "Wed with Olympus XXVII", and uses a straight apostrophe where the web uses a curly ’.
10. **The crest has a blue halo** in the item sheet capture. The web has no glow.

## Shared-source consumption

| Shared source | Generated | Hardcoded |
|---|---|---|
| Tokens | `Tokens.kt` | 30 colours, 1130 `.dp`, 61 `.sp`, and the XML colours (not generated) |
| Icons | 5 icons (`ic_tab_*`, `ic_frame_corner`) | 228 Material references and about 40 hand-written vectors |
| Copy | `strings.xml` (1803 strings, from `gen-mobile-copy.mjs`) | 57 `.uppercase(` transforms |
| Formats | `Formats` in 3 files (`RaidCardWords.kt`, `HomePresenter.kt`, `Applicant.kt`) | 89 local `.format(`, `SimpleDateFormat` and `DateTimeFormatter` calls |

## iOS vs Android

- **Sign-up control:** iOS uses four filled buttons, Android a segmented pill. The mock uses the segmented pill on both.
- **Bar title:** Android shows one, iOS leaves the bar empty.
- **Guild-links block:** iOS has it, Android does not.
- **Item tooltip:** iOS shows "8% drop chance", Android shows "8%".
- **Lock-screen timer:** gold on iOS, black on Android.

## Stale captures

- `m0/android-home-34.png`.
- All of `2026-10-31-mobile-wave3-widgets/build/and-*.png`.
- `m0/android-design-system-sheet.png` and `m0/android-activity.png`.
- Guild, Rankings, Reports, Planner and Guides have never been captured.

## Notifications

Channel names come from `strings.xml`. Push title and body text come from the server (`PushNotifier.kt:29-30`, `RaidActionReceiver.kt:76-77`). The live card copy is in the web's voice.

VERDICT: DRIFT
