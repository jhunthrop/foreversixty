# Guild raids and sign-ups: wow-player review (Lane G1)

Reviewer: wow-player persona (Classic raid officer since 2019, Raid-Helper and loot-council veteran).
Seen rendered: rows 1-16 and 18 in this folder, full pages and crops at 360/390/1440/2000.
Fixture names (Barrow Deeps, Hyjal Summit, scene-to-scene roster differences) and the separate
fix round (` · ` spacing, note bullets, native checkboxes, the two row heights) are not judged here.

## 1. The question each surface answers

- **Raid page:** "Am I in for Thursday, at what time for me, and who else is coming?" Yes, first
  thing: when, countdown, count, then YOUR SIGN-UP with In/Late/Bench/Out on every own character
  (05 at 360/390/1440). The two-character strip is exactly what an alt-heavy guild needs.
- **Guild band line:** "Is there a raid and am I on it?" Yes: `Thursday 8 pm, you are In on Zulmara
  as DPS, 14 of 40 signed up.` plus segments and `Onyxia's Lair →` (14-band-in). One tap really is
  one tap on a phone: the band segments are full width and 44 px at 390.
- **Raids tab:** "What is planned and am I signed?" Yes: date block, time, countdown, count, chip (01, 02).
- **Form:** "Post Thursday's Ony in 30 seconds from my phone." Yes: zone, date (defaults a week out),
  8 pm and 3 h prefilled, `Post the raid`. Four taps if the defaults are right (13-form-default).
- **Copy list:** "Paste the roster into #raid-signups." Nearly (see 2.3).

## 2. What a Classic player would find wrong, most damaging first

1. **A cancelled raid still counts down and still asks you to sign up.** 08-page-cancelled-390: the
   `CANCELLED` pill sits above a gold `in 2 days 7 h` and the gold progress bar, as if it is still on.
   02-tab-member-390: the cancelled Hyjal row shows a gold `NOT SIGNED UP` chip beside `CANCELLED`
   (and it truncates the title to `Hyjal S...`), with a gold `in 5 days 7 h`. That is the one thing
   that gets screenshotted: "the site wants me to sign up for the raid we cancelled." Fix: on a
   cancelled raid drop the countdown and the bar on the page and the row, draw the count muted, and
   show no viewer chip at all (`CANCELLED` alone). `Planned 3` should count only live raids, or the
   cancelled row moves to the bottom with the count reading `Planned 2`.
2. **The count and the role split do not add up.** Every surface reads `14 of 40 signed up · 1 tank
   · 3 healers · 9 DPS` (05, 09, 01 row `14/40`). An officer adds 1 + 3 + 9 = 13 and stops trusting
   the number. In + Late in the total is right (Raid-Helper counts Late too); just say so. Fix, page
   band and tab row: `14 of 40 signed up · 1 late · 1 tank · 3 healers · 9 DPS` (drop `1 late` when
   zero). Copy list line 2 (16-copy-failed): `14 of 40: 1 tank, 3 healers, 9 DPS, 1 late`.
3. **Which segment I chose is too faint at a glance.** 05-strip-360: `In` is marked only by a
   slightly lighter cell and a thin green underline; `Bench` on Duskstrider by a grey one. On a
   phone in daylight I cannot tell In from Out without squinting, and that is the page's one job.
   Fix: the checked segment gets a filled tint (`--kill` at ~15% for In/Late, `--line` grey for
   Bench/Out) plus the existing 2 px line, as Raid-Helper's lit button does.
4. **Locked segments look live.** 07-page-locked/lock-race: after `Sign-ups are locked. An officer
   can still change yours.` the four segments look exactly like the enabled ones. Fix: draw them at
   the disabled opacity so nobody taps three times wondering why nothing happens.
5. **The needs line is red.** 14-band-in-390 and 14-band-unsigned-1440: `What the guild needs from
   you before Thursday: 3 gear upgrades waiting (23.8 DPS)...` is in the wipe red, under a sign-up
   line. Red is for failures (spec §5); this reads like a scolding from the GM. Fix: `--muted` text
   with the counts in `--strong`. (The line itself predates G1; the `before Thursday` hookup is new.)
6. **Band line says "signed up" twice when unsigned.** 14-band-unsigned: `Thursday 8 pm, you have
   not signed up, 13 of 40 signed up.` Replace with `Thursday 8 pm, you have not signed up yet.
   13 of 40 are in.` (keep `{n} signed up` for unknown size: `13 are in.`).
7. **Cancel confirm is safe but flat.** 09-confirm-390: a second tap is enough to stop a fat finger,
   good. But `Cancel the raid` and `Keep it` are twin grey buttons. Fix: `Cancel the raid` in the
   `--wipe` outline, `Keep it` first in focus order. And `Post the cancellation to Discord` should
   start ticked when the raid was posted to Discord; an unticked default means 40 people still
   show up to Ony.
8. **Phone titles truncate before the chip.** 02-390 `Barrow Deeps: first c...`. Let the title wrap
   to two lines and keep the chip on line 1 right; a raider needs the whole title.

## 3. Missing against what officers already use

- **Discord-local time in the copied list.** Line 1 reads `Onyxia's Lair · Thu 10 Dec, 8:00 pm CST`.
  In Discord, half the guild converts CST in their head and gets it wrong. Raid-Helper prints a
  `<t:unix:F>` stamp that Discord shows in each reader's own time. Fix: append ` · <t:{unix}:F>` to
  line 1 so the paste renders locally in Discord and still reads as guild time in plain text.
- **Guild time on the raid page for a raider in another zone is not captured.** The band line has
  ` (9 pm guild time)` in the spec, the form has `For you: Thu 17 Dec, 9:00 pm.` (13-for-you, clear
  and correct), but every raid page capture runs in Chicago so I never saw the page's two-clock line.
  Add one capture of 05 under New York before ship; I cannot approve a line I have not seen.
- **Class counts for the officer** (follow-up, not a blocker): Classic comps are built on classes,
  not roles (how many warlocks for curses, shamans for totems, mages for decurse). A
  `4 Warriors · 3 Hunters · 2 Rogues...` line under `Sign-ups 14/40` in officer view is what
  Raid-Helper shows and what I would check first.

## 4. Data I do not believe

Nothing load-bearing. Roles and specs read right (Protection tank, Holy/Restoration healers, Arcane
mage marked `DPS` on the Late row). Fixture only: Barrow Deeps and Hyjal Summit are not Classic
raids, and `Hyjal Summit · 0 bosses · 7 wipes · 109 min` under Raid nights would be mocked if it
ever reached production data. The officer roster at 1440 (09-roster) leaves a large empty block
right of Late/Bench/Out; already noted in the README, worth fixing with the three-column layout.

## 5. Words

`In · Late · Bench · Out` is exactly the Raid-Helper vocabulary a Classic raider expects; Bench as a
self-chosen "I can come if needed" reads right. `Lock sign-ups`, `Copy the sign-up list`, `Keep it`,
`Sign-ups are locked. An officer can still change yours.`, `This raid was locked while you were
signing up.`, `Only 2 more fit: a guild plans at most 12 raids ahead.` are all things an officer
would actually say. No marketing copy anywhere on these surfaces. Nothing to remove.

## Verdict: SHIP WITH FIXES

Blocking:
1. Cancelled raid (08-cancelled, 01/02 tab rows): no countdown, no bar, no viewer chip; `CANCELLED` only.
2. Count line (05, 09, 01, 02, 16): `14 of 40 signed up · 1 late · 1 tank · 3 healers · 9 DPS`;
   copy line 2 `14 of 40: 1 tank, 3 healers, 9 DPS, 1 late`.
3. Checked segment (05-strip, 14-band): filled tint on the chosen segment.
4. One New York capture of the raid page band showing the guild-time line.

Should fix in this lane:
5. Locked segments drawn disabled (07).
6. Unsigned band line: `Thursday 8 pm, you have not signed up yet. 13 of 40 are in.` (14-unsigned).
7. Cancel confirm: `Cancel the raid` in wipe outline; Discord box ticked when the raid was posted (09-confirm).
8. Copy list line 1: `Onyxia's Lair · Thu 10 Dec, 8:00 pm CST · <t:{unix}:F>` (16).
9. Needs line out of red (14-band-*).
10. Phone titles wrap instead of truncating (02-390).
