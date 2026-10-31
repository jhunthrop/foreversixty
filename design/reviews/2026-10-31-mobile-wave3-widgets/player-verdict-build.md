# F14 widgets and live surfaces: wow-player build verdict

Reviewer: wow-player, 2026-10-31. Read: spec 5.6, my design verdict, `apps/shared/widgets/figures.json`, and the captures in `build/` (ios-* and and-*). I looked at the captures one by one. I did not compare against the board PNGs pixel for pixel.

The question: "Am I in tonight, and what is my next upgrade?"

**Verdict: HOLD.** The figures now agree. The words and numbers match across every capture I read: "24 of 40", "in 2 days 2 h", "Thu 8:00 pm", "38 of 40" on the live card, "+5.3 DPS", "8 upgrades · +18.3 DPS". That was my design SHOULD and it is fixed. What holds it is four places where a state loses the answer, or the surface says something false. A raider would mock each of them.

## MUST

1. **The signed-up ring has no unit.** `ios-lock-circular-signed-up.png` shows a ring with a bare "24". My design round asked for "38 of 40" and a ring that "always carries its unit", and the spec says the same. Here it is a number in a circle again. Is it 24 raiders, 24 minutes or 24 DPS? Print "24 of 40", or "24/40" with a small "of 40" under the figure. The countdown ring, `ios-lock-circular-countdown.png`, reads "2d" with nothing around it. It has a unit, so it passes, but the two rings side by side are unlabelled.

2. **The Android notification says "Invites" for a raid that has no invites.** `and-live-raid.png` reads "Invites for Onyxia's Lair in 39 minutes", and the top line of the same card says "until the pull at 8:00 pm". Raids have no invite time yet, so the title contradicts the line above it. The cause is the golden vector "the Live Activity window" in `figures.json`: its `liveTitle` is "Invites for Onyxia's Lair in 39 minutes" while `liveUntil` is "until the pull at 8:00 pm". Only the case "with an invite time" should use "Invites for". With no `invite_at`, use "Onyxia's Lair starts in 39 minutes" and have both platforms read it from the vectors.

3. **The medium widget has no answer in words in two states.**
   - `ios-raid-medium-not-signed-up.png` shows three plain buttons and nothing else. The spec wants a gold "Not signed up" tag. A player glancing at the widget cannot tell "I never answered" from "the widget is broken". The small one, `ios-raid-small-not-signed-up.png`, I did not open, but check that it shows the tag.
   - `ios-raid-medium-locked.png` shows "LOCKED" and drops "You: In · DPS". The sign-ups are locked, and the first thing a raider wants is whether they made the roster. The lock-screen rectangular for the same state does print "You: In · DPS" (golden `you` for locked), so the surfaces disagree.
   - The medium widget never shows role or spec on any state. Only the small widget has "IN · DPS".

4. **The stale small widget hides your answer.** `ios-raid-small-stale.png` and `and-raid2x2-stale.png` replace the "IN · DPS" tag with "Updated 14 h ago". The spec says the normal widget stays and the stamp goes at the foot. A 14-hour-old widget is exactly when you want to see what it last knew. Keep the tag and put the stamp under it, in muted text.

## SHOULD

1. **"39 minutes" beside "38:12".** `and-live-raid.png` shows "in 39 minutes" while the clock runs "38:11". The iOS card shows 38:12. The ceiling rule is in the golden vectors, so it is deliberate, but a player who sees both in one card will say it is off by one. When seconds are on screen, say "in 38 min", or drop the minutes from the title.
2. **The Android notification icon is a blank rounded square.** All four notification captures (`and-live-raid.png`, `and-live-sim-running.png`, `and-live-sim-finished.png`) show a plain square on a blue disc. The spec says the v1 notification small icon. If this is a capture-harness placeholder, say so. If it is the shipped icon, a raider will screenshot it.
3. **"Last: Onyxia's Lair, 3 Dec." predates the raids.** `ios-raid-small-no-raid.png` and `and-raid2x2-no-raid.png`. Raids open on 9 Dec, so there was no Onyxia on 3 Dec. This is a fixture problem, but it is the one string a raider reads and checks. Use a date after 9 Dec, or a real past raid. The platforms also differ: iOS prints "3 Dec." and Android "Thu 3 Dec." and wraps it to a second line. Pick one; the spec says `{d Mon}`, so iOS is right.
4. **The compact island shows "38:12".** `ios-island-compact.png`. The spec and my design round said "38m". The clock in a pill is fine and arguably better. Confirm that is a decision, and keep the same figure in the expanded island.
5. **The expanded island has no signed-up count or guild.** `ios-island-expanded.png` drops "38 of 40" and "OLYMPUS XXVII" that the lock-screen card carries. It is narrower, but "38 of 40" is the number a raid leader checks. It also shows "until the pull at 8:00 pm · You: In · DPS" in one run of dots, which is dense.
6. **The iOS Live surfaces settings capture shows two green blobs.** `ios-settings-live-surfaces.png`. There is no knob and no on/off state. I cannot tell what is on. Android `and-live-settings.png` shows "Server sim progress" off, but the spec says both default on. Re-capture both with the knobs visible and the defaults on.
7. **Not saved sits on the Out button with In still checked.** `ios-raid-medium-not-saved.png`. It is correct (the press was Out, it reverted), but a player would read "Not saved" as being about Out only. Fine, but show it for 10 s as specced and confirm in the capture notes.
8. **Largest text.** `ios-raid-medium-in-ax5.png` fits, but "in 2 days 2 h" runs into the button row. It does not overlap, but there is no gap. The `not-saved-ax5` and `upgrade-medium-ax5` captures I did not open.

## Polish

- The upgrade widgets never say which slot the item is for, except the lock-screen rectangular ("Next upgrade · Waist"). The small widget with only an icon and "Knowledge of the Timbermaw" would help from a slot word. Wowhead and the client always say it.
- The upgrade medium footer "8 upgrades · +18.3 DPS" is correct and uses the unit. Good.
- `ios-live-raid-cancelled.png`: "Cancelled" in red with "You are In · DPS · Lightningbolltz" and "38 of 40" below it reads oddly. Drop the "You are In" line when it is cancelled.
- The small raid widget "BENCH" tag in `ios-raid-small-bench.png` is muted grey, the same weight as "LOCKED". A bench answer is news to the player; use the gold.
- The lock-screen inline and rectangular captures are monochrome with no quality colours, as the spec requires. Good.

## Data I do not believe

Nothing wrong with a figure I could check against the vectors. For the accuracy reviewer: "Knowledge of the Timbermaw" is shown as a purple belt and "Dryad's Wrist Bindings" as purple, but the drop-chance and source fields do not show on a widget, so I cannot say if they are obtainable by band. The 8 upgrades and +18.3 DPS are invented per the brief.

## What agrees

- Count: "24 of 40" on `ios-raid-medium-in.png`, `and-raid4x2-in.png`, the circular ring "24", and the golden `signedUp`.
- Clock: "Thu 8:00 pm" on every home-screen capture; "Thu 8 pm" on the inline capture.
- Countdown: "in 2 days 2 h" on the widgets, the rectangular and "2d" on the ring. All match the vector.
- Live card: 38:12 on the iOS card and island, 38:11 on Android a second later, "38 of 40" on all.
- States have words everywhere I looked: signed out, no guild, no raid, no gear.

## Fixes before SHIP

1. Print "of 40" on the signed-up ring (`ios-lock-circular-signed-up.png`).
2. Replace "Invites for" with a pull-time title when there is no `invite_at`, and change the golden vector.
3. Show the "Not signed up" tag and the "You: In · DPS" answer on the medium widget, including locked.
4. Keep the answer tag on the stale small widgets.
5. Re-capture the settings screens with knobs, and fix the "3 Dec" fixture date.

Re-run and send the updated `ios-lock-circular-signed-up`, `and-live-raid`, `ios-raid-medium-not-signed-up`, `ios-raid-medium-locked`, `ios-raid-small-stale` and `and-raid2x2-stale`.
