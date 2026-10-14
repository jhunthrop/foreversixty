# The weekly capture loop

Every week of the addon v2 build ends with the same loop: shoot every surface in every state on the beta client, hand the
folder to the two personas, and merge the week's lane PRs only when both return SHIP. The spec's section 8 states the rule
("the two personas review screenshots of every surface in every state ... before the week closes. SHIP from both is the
gate"); the plan's Task L9.1 is the step list. This file is the checklist; `addon/tools/capture_checklist.md` is the list of
captures.

## Layout

```
docs/superpowers/reviews/
  README.md            this file
  week1/               the Oct 11 to 14 session: 371 jpgs and chip-captures.md (the manifest), web-addon/, web-score/
  week2/  week3/  week4/
    <captures>.jpg     one file per capture, named as below
    captures.md        the manifest: one row per file, "what it shows", the build it was shot at
    ux-designer.md     the designer persona's report
    wow-player.md      the player persona's report
```

A week's folder holds that week's captures and nothing else. A re-shoot after a fix round is a new file with the next
number, not an overwrite, so the manifest keeps what each round showed.

## Naming

Files are `<surface>-<N>-<state>.jpg`, as in `marks5-strip.jpg`, `tip7-chest.jpg`, `dialog6-2-row-preview.jpg`,
`hunter-24-flyout-screen.jpg`. `<surface>` is the surface (or the character, `hunter-`, when one round covered several), `N`
is the round counter within the week (it goes up with every fix round, so a re-shoot never reuses a name), and `<state>`
is a few lowercase words. A whole-screen game shot and its crop share a stem: `-screen.jpg` and `-crop.jpg`.

Week 1 kept those names. From week 2 the checklist gives each capture a stable name,
`<surface>-<level>-<state>.jpg` with the level token `l1`, `l20`, `l60` or `combat` (`chip-l20-point-to-spend.jpg`); a
round's counter goes after the surface when a re-shoot is needed (`chip2-l20-point-to-spend.jpg`). The manifest row is the
join: it says which checklist line a file answers.

## Taking a capture

1. Log in on the beta (`C:\Program Files (x86)\World of Warcraft\_classic_beta_\`) at the build the lane PR is on, `/reload`
   if the files changed on disk (the first reload after a commit sometimes does not take: a second one does).
2. Reach the state with the row or section the checklist line names (`addon/README.md` rows 1 to 76, spec section 4). Fixtures
   below, if the line asks for them.
3. Take a game screenshot, not an OS one, so no overlay is in it:
   `/run C_Timer.After(3, function() Screenshot() end)`, then put the mouse where the state needs it (hover, pressed) and
   wait. Week 1 used a 2 s delay in places; 3 s leaves time to move the pointer. For a held key (Shift for the score
   components) hold it past the delay.
4. The file lands in `C:\Program Files (x86)\World of Warcraft\_classic_beta_\Screenshots\` as a 3840 by 2160 jpg. Copy it
   into the week's folder under the checklist name.
5. Crop with PIL to the surface, keeping a margin, and save a `-crop.jpg` beside the whole-screen shot for anything small
   (the chip, the marks, a 12 px arrow): `Image.open(f).crop((l, t, r, b)).save(out, quality=92)`; for a 12 px mark add a 4x
   resize (`Image.NEAREST`) so the persona sees the pixels. Keep the whole-screen shot too, because the personas judge
   placement against the client's own frames.
6. Add the manifest row the moment the file exists: the file, what it shows in the words on screen (the text the player
   reads), the build and character. Say what you did not capture and why, in the manifest, not in your head.

Capture rules carried over from week 1:

- Read the shot by eye before it goes in the manifest. A timed shot can catch a tooltip fading or the wrong hover; say so and
  re-shoot instead of captioning what you hoped to see.
- Never claim a state a file does not show. A seeded or faked state (`seedLive`, a `/run` override of `IsShiftKeyDown`) is
  named as such in its row.
- Restore every override (a switch turned off, a build changed) before the next shot.
- Hover and pressed looks cannot always be shown with the capture tool's pointer; a row that cannot be shown is marked owed
  for a human hover, not captured.
- The glow, the chip and the tooltip are proven on a real client screenshot, never by a test.

## Fixtures

Two stand-in addons put known data in front of the surfaces. They never ship. Both are swapped in at the character screen
(`/logout` first), used for the captures, and restored afterwards.

| Fixture | Folder in the repository | Swapped over | Gives |
|---|---|---|---|
| Data | `addon/fixtures/ForeverSixtyData-l44` | the client's `Interface\AddOns\ForeverSixtyData` | Bow Jackzon 1240 (rare), Smoke Gouda 1610 (epic), Facep Alm (no logs), the guild roster scores. Dated: move `generated` and `as_of` forward after three days or the client says "too old". |
| Inbox | `addon/fixtures/ForeverSixtyInbox-l53` | the client's `Interface\AddOns\ForeverSixtyInbox` | a `weights`, a `sim_result`, a `last_raid` dated 2026-10-12 (told once per login, only within 14 days: move `date` forward after 2026-10-26) and a `loot_values`, for Bow Jackzon. |

The commands are in `addon/README.md` ("The L5.3 inbox fixture", "The L4.4 data fixture"). In short, from the AddOns folder
(`C:\Program Files (x86)\World of Warcraft\_classic_beta_\Interface\AddOns`):

```powershell
# swap in the data fixture (removes only the link, never the repository folder)
cmd /c rmdir ForeverSixtyData
cmd /c mklink /D ForeverSixtyData C:\Users\jhunt\Documents\Repos\foreversixty\addon\fixtures\ForeverSixtyData-l44
# swap in the inbox fixture
(Get-Item ForeverSixtyInbox).Delete()
New-Item -ItemType Junction -Path ForeverSixtyInbox -Target C:\Users\jhunt\Documents\Repos\foreversixty\addon\fixtures\ForeverSixtyInbox-l53
```

Restore both at the end of the session, at the character screen, with the real folders as targets:

```powershell
cmd /c rmdir ForeverSixtyData
cmd /c mklink /D ForeverSixtyData C:\Users\jhunt\Documents\Repos\foreversixty\addon\ForeverSixtyData
(Get-Item ForeverSixtyInbox).Delete()
New-Item -ItemType SymbolicLink -Path ForeverSixtyInbox -Target C:\Users\jhunt\Documents\Repos\foreversixty\addon\ForeverSixtyInbox
```

Write "fixtures restored" in the manifest. A fixture left in place shadows the real companion inbox and the real data addon,
and the next person's `/fs diag` will say things that are not true. `/fs diag` and the Settings foot show which data addon is
in force ("Player ratings ...: current" with the data fixture).

## The persona review

1. Finish the checklist as far as the characters allow. What cannot be shot stays `owed:` in the manifest with its blocker
   (a level-60 character, the user's test group, spare gear, a hostile mob).
2. Give `ux-designer` and `wow-player` the week's folder and the manifest, the spec's section 4 for the lanes in the week,
   and the lanes' PR numbers. Give them the screenshots, not the diffs: they review what a player sees. Run them in parallel;
   neither sees the other's report until both are filed.
3. Each returns a report: SHIP or FIX, with must-fix items numbered and each pinned to a file. The week's lanes need SHIP from
   both before their PRs merge. A FIX on either sends the lane back to its implementer; the re-shoot goes in the same folder
   under the next round number, and both personas review the changed surfaces again.
4. Items a persona wants shot that the characters cannot show yet are recorded as owed in the checklist and carried to the
   next week; they block a SHIP only when the persona says the surface cannot be judged without them.

## Filing

Commit the week's folder with the manifest and the two persona reports (`ux-designer.md`, `wow-player.md`) once both say SHIP,
or at the end of the week with the FIX items listed if they do not. The commit is `docs(addon): week N captures and reviews`.
Update the status column of `addon/tools/capture_checklist.md` in the same commit: replace `owed:` with `captured:` and the
file names, only for lines whose files are now in a manifest. Week 1's folder (`week1/`, with `chip-captures.md` as its
manifest) is the starting status of that file.

The L9.2 release (`addon-v2.0.0`) waits for week 4's SHIP from both.
