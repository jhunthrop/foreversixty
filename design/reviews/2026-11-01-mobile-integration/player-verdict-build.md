# Player verdict: F15 system integration, build review

## Round 1

Reviewer: veteran Classic raider. Judged on the captures in `build/ios/` and `build/and/` against spec §5.7 F15 (lines ~630-718), §6 Strip and Search-button rows, and the fixture story (Lightningbolltz, Elemental Shaman, Olympus XXVII, Onyxia's Lair Thursday 8 pm, 25 of 40). The planner follow-ups (F11.4 / F13.5) are not in these captures and are not judged here.

What is right: the Siri copy matches the spec word for word (confirm question, `You are In for Onyxia's Lair, Thursday at 8 pm, as DPS on Lightningbolltz. 25 of 40 signed up.`, the refusal sentence, the `3 tanks, 6 healers, 14 DPS, 1 late` line). The three strips carry the exact spec strings, gold `Retry` / `Update`, and survive largest text on both platforms. The Android toast reads `You are In for Onyxia's Lair.` and the shortcut never opens the app. Share texts match the table (`Lightningbolltz, Elemental Shaman, level 60`, `Elemental Shaman guide`, `Healer tier list`, `Onyxia's Lair · 14 Dec`). Offline data copy, the remove question and the footer match the spec; no build ids or hashes.

### MUSTs (block ship)

1. **The raid count disagrees between two answers a player hears back to back.** `f15-siri-signedup-ios.png` says `You are In ... 25 of 40 signed up.` `f15-siri-nextraid-ios.png` (and -large) says `You are In. 24 of 40 signed up.` and `f15-siri-whoiscoming-ios.png` says `24 of 40 signed up: 3 tanks, 6 healers, 14 DPS, 1 late.` The fixture story is 25 of 40 with the player In. In a raid leader's ears one of these is a lie. Spec line: intent.signedUp and the Next raid / Who is coming rows use the same count. Make all three read the same number (25 if the player is In, and the who-is-coming role split must add up to it).
2. **A character name with a space.** `f15-spotlight-entries-ios.png` (and -large) lists `Obnoxious Yell`, `Fury Warrior · 58 · OLYMPUS XXVII`. Character names in this game cannot contain spaces, and the same character is `Obnoxiousyell`, `Warrior · Level 60 · US Normal` in `f15-search-watching-and.png`. A player would screenshot this. Same character, same name, same level on every surface (spec §5.7 Spotlight: the app's characters). Fix the fixture or the index.
3. **The same guide has two dates.** Spotlight shows `Elemental Shaman guide`, `Guide · Updated 24 Sep 2026`; the guide header in `f15-strip-newdata-ios.png` and the Guides root in `f15-guides-search-and.png` say `Updated 1 Nov 2026`. One guide, one date.
4. **iOS Guides root has no search button.** `f15-guides-searchbutton-ios.png` shows the large title `GUIDES`, the segments and the class grid, but no magnifier and no character selector in the bar. Android (`f15-guides-search-and.png`) has the magnifier top right. The follow-up (§6 Search button row: `Search`, opens F12) is the point of this capture; as shown it is not built or not visible on iOS. Both apps must tell the same story.
5. **Largest text breaks words in the middle (iOS).** `f15-siri-upgrade-large-ios.png`: title truncated to `What's my next...`, `Leatherworki` / `ng`, `Lightningbo` / `lltz.`. `f15-search-watching-large-ios.png`: `Lightningboll` / `tz`, placeholder cut to `Search...`, chips clipped at `Characte`. `f15-spotlight-entries-large-ios.png`: the row icons are pushed to the left edge, clipped, and overlap the titles (book icon sits on `Shaman guide`). A player at AX5 sees a broken line. Names must wrap at the word or shrink, never split mid-word; the Siri snippet must put the icon above the text at this size as the Offline data rows already do.
6. **Acceptance captures missing.** Spec acceptance (line ~717) asks for: the share sheet for a character, a guide and a report on both platforms (Android has none); F15.3 launcher search on Android (only the long-press menu is shown); the Siri `Which raid?` and spec question (`Which spec is Lightningbolltz on Olympus XXVII?`); a Shortcuts-app run without confirmation; the Control Center control `In for the next raid`; the Assistant `OPEN_APP_FEATURE` entry points. The iOS share captures also show only the title and link, not a sheet; if that is the accepted stand-in say so, but the Android ones are simply absent.

### SHOULDs

1. **Source wording differs across surfaces.** Siri snippet (`f15-siri-upgrade-ios.png`): `Chest · item level 58 · Leatherworking`. The upgrade list (`f15-strip-stale-ios.png`): `Chest · Crafted: Leatherworking`. A player reads a bare `Leatherworking` as a vendor or a zone. Say `Crafted: Leatherworking` in the snippet too (spec: `{source}`).
2. **Offline `Update now` still looks live.** `f15-offlinedata-offline-ios.png` prints `Update now needs a connection.` under a full-gold button. The spec says writes stay disabled with their reason; the button should look disabled.
3. **Switches have no knob in the drawn iOS captures** (`f15-offlinedata-data-ios.png`, `-empty-ios.png`): empty gold pills, so on/off cannot be judged. The real-app capture (`f15-offlinedata-app-ios.png`) has them; re-draw the snapshots.
4. **Android Offline data row `Item icons` ends at `41 MB`** (`f15-offline-data-and.png`) while iOS says `41 MB · Updated today`. Spec row format is `{what} · {size} · Updated {when}`. Also `UPDATE NOW` / `REMOVE DOWNLOADED DATA` are all caps on Android, and on iOS only `UPDATE NOW`; the spec label is `Remove downloaded data`.
5. **Guides footer and Offline data disagree.** iOS Guides root: `Saved on this phone: Warrior, Shaman · 32 KB.`; Offline data (real app): `Warrior · 23 KB`. Android Guides root says `Nothing is saved on this phone yet.` next to Offline data showing 48 MB with Guides saved. Same phone, one number.
6. **Android launcher menu** (`f15-launcher-shortcuts-emulator-and.png`): `Next raid`, `Sign up In`, `Voltaris`, `Stormwrath`. The spec's static `My upgrades` and `Search` are not shown, the current character (Lightningbolltz) is not shown, and `Stormwrath` is not one of the two watched characters in Search (`Voltaris`, `Obnoxiousyell`). iOS Search watches `Morvane` and `Kraggor` (US PvP) instead. The fixtures should tell one story across platforms.
7. **Spotlight set is thin.** Own characters should lead (up to 9); Lightningbolltz is absent, and `Yellfury`, `Hunter · 31` has no realm while the other row does. Raid rows are right (`Onyxia's Lair · Thu 8 pm`).
8. **Search rows:** `Elemental Shaman guide` carries an external-link arrow in both apps. If it opens the native guide (follow-up) it should show the chevron like `Best in slot` and `Rankings`.
9. **Odd empty space above the strip.** `f15-strip-newdata-ios.png` and the Offline data iOS captures leave about 270 px of blank above the strip or title. Looks like a missing bar in the drawn view.
10. **Offline strip time.** `f15-strip-offline-large-ios.png` reads `at Thu 9:41 pm` while the stale strip reads `at 9:41 pm` (v1 text, so tolerable), but the Android offline strip says plain `9:41 pm`; pick one.

### Data a player will doubt (for the accuracy reviewer, not blocking here)
- `Legionite Glaive` (Main hand, Crafted: Blacksmithing, `+21.2 DPS`) and `Hyper Deluxe Sniper Rifle Mk XVII` (Crafted: Engineering): neither name rings from Classic; confirm they are Forever-new crafts.
- `Ironfeather Breastplate` (leather chest, ilvl 58) as an Elemental Shaman's next upgrade replacing `Vest of Elements`: legal as leather below mail, but check the source and the +3.8 DPS.
- Guide stat priority in `f15-strip-newdata-ios.png` lists `Nature power` third; unusual, outside F15 scope.
- Report dated `14 Dec` is a Monday, while the raids are Thursdays (10 and 17 Dec); fixture, but a raider will notice.

(Round 1 verdict: HOLD)

## Round 2

Reviewer: same veteran raider, on the new captures. Rulings in the brief are not re-raised (confirm order, alt name Obnoxious Yell, 25 of 40 with 3/6/15 + 1 late, system-UI gaps, Android caps).

### Round 1 check
- M1 raid count: FIXED. Sign-up, next raid and who-is-coming all read 25 of 40 (`f15-siri-signedup-ios.png`, `-nextraid-ios.png`, `-whoiscoming-ios.png`: `25 of 40 signed up: 3 tanks, 6 healers, 15 DPS, 1 late.`); it adds up.
- M2 name with a space: closed by ruling for the name. Spotlight now says `Obnoxiousyell`. The level still disagrees, see SHOULD 1.
- M3 guide date: FIXED. `Updated 1 Nov 2026` in Spotlight (`f15-spotlight-entries-ios.png`), Guides root on both apps.
- M4 iOS Guides search button: FIXED. Magnifier top right in `f15-guides-searchbutton-ios.png`.
- M5 mid-word breaks at AX5: FIXED for the Siri snippet (`f15-siri-upgrade-large-ios.png`: icon above, words wrap whole) and Search (`f15-search-watching-large-ios.png`, names wrap at the word). Spotlight leftovers in SHOULD 2.
- M6 acceptance captures: mostly FIXED. Android share sheets for character, guide, report; launcher shortcuts static and signed in; launcher search; toast; iOS system share and Shortcuts run all present. `f15-launcher-search-emulator-and.png` shows only the app icon (documented gap, on-device pass). Siri `Which raid?`, spec question and Control are present as our own views.
- S1 `Crafted: Leatherworking` in the snippet: FIXED.
- S2 Update now disabled offline: FIXED (`f15-offlinedata-offline-ios.png` dimmed, with the reason line).
- S3 switch knobs: FIXED (white knob in the drawn captures).
- S4 Android `Item icons` row and iOS: both still end at the size (`41 MB`, `3.3 MB`) with no `Updated`. Spec row is `{what} · {size} · Updated {when}`. Now the same on both apps; left as SHOULD 4.
- S5 Guides footer vs Offline data: Android `Saved on this phone: ... 2.3 MB` against Offline data `2.4 MB`; iOS `23 KB` against `31 KB`. Closer, not equal (SHOULD 3).
- S6 launcher menu: Voltaris and Stormwrath now match Search's Watching on both apps. Current character still not in the menu (platform slot rule, ruled).
- S8 chevron on the guide row: FIXED. S9 blank band above the strip: not re-checked, not blocking.

### MUSTs (block ship)

1. **Gear gap for the same character is 36.0 on iOS and 3.0 on Android, and the worn chest differs.** `f15-char-open-planner-ios.png` shows Lightningbolltz in `Vest of Elements`, `Gear gap 36.0 points`, gold dots on five slots. `f15-char-open-planner-and.png` shows `Ironfeather Breastplate`, `Gear gap 3.0 points`, and one dot. Siri's next upgrade (`f15-siri-upgrade-ios.png`) says Ironfeather Breastplate replaces Vest of Elements, so Android's character already wears the upgrade. A raider comparing phones would call one of them wrong. Spec: one character, one story (fixtures follow the spec story). Make both the same worn set and the same gap.
2. **The Compare-with-mine screen tells two different stories.** `f15-shared-build-compare-ios.png`: `3 of the shared build's points still to take.`, Elemental 30 / Enhancement 3 / Restoration 15, no differing talent marked in the tree. `f15-shared-build-compare-and.png`: `2 of your 48 points differ from the shared build.` plus `The shared build takes 5 points you have not.`, Elemental 28 / Resto 17, with +2, -2, -2 markers. Same fixture character, same shared build (Orc Shaman 30/3/18), different counts; and iOS says DIFFERS with no dashed talent anywhere on screen to explain it. Spec F13.5 / plan.cmp.* copy: one wording, one answer. Align the fixture and the sentence on both apps. Builds seen on the character page says `30/0/21`, a third build; fine as a different row, but name it so it is not read as the same build.

### SHOULDs

1. **Level disagrees.** iOS Guides root selector `Obnoxious Yell` / `Fury · 58` (`f15-guides-searchbutton-ios.png`); Spotlight `Obnoxiousyell`, `Fury Warrior · 60`; Android Search `Level 60`. Name stays by ruling; make the level 60 (or 58) everywhere.
2. **iOS Spotlight at AX5** (`f15-spotlight-entries-large-ios.png`): book and calendar icons are drawn oversized and touch the left edge, with the title hard against them; the character crests are small beside them. Cap the icon size like the crests.
3. **Offline data vs Guides footer sizes** (see check S5): same phone, one number (`2.3 MB` vs `2.4 MB`, `23 KB` vs `31 KB`).
4. **`Item icons` row has no `Updated {when}`** on either app (`f15-offline-data-and.png`, `f15-offlinedata-app-ios.png`).
5. **iOS Search at AX5**: placeholder cut to `Search...` and the chip row clipped at `Characte` (`f15-search-watching-large-ios.png`). The chips scroll, so tolerable; shorten the placeholder to `Search` at this size.
6. **Open in planner** (`f15-char-open-planner-ios.png` / `-and.png`): iOS is a separate card with a gold label and chevron, Android sits inside the Builds seen card; the iOS row also lacks the build icon the Android row lacks. Pick one layout. Item icons in the iOS gear grid are blank squares in this drawn capture (Android shows real icons); redraw with icons.
7. **Android character share sheet** (`f15-share-character-emulator-and.png`) sits over the signed-out landing screen, so it does not show the app sharing from the character page. Re-capture from the page.

### Data a player will doubt (for the accuracy reviewer)
Unchanged from round 1 (Legionite Glaive, Ironfeather set, Nature power third in the guide, report dated `14 Dec` against Thursday raids). New: `Vest of Elements` shown as Chest `ilvl 63` while `Hand of Edward the Odd` is `ilvl 62` main hand: confirm item levels.

Verdict: HOLD

