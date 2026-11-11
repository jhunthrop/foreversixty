# Guild tab (iOS + Android), guild-officer review, round 3, 2026-11-11

**Job of the page:** on the phone, an officer sees who is coming and what the raid still needs. They act on it (Ping on Discord, applications, setup) and see who is missing or behind. A member sees when the raid is, their answer, and their own job.

**Evidence limit.** The `parity/*-full-*.png` files are not full-page renders. The iOS officer image is 1179x2556 and ends in the Assignments heading. The Android one is 1081x2401 and ends under the tab bar. So I could not see Loot, Gear gaps, Your attendance, Missing raids, or the gold winner name rendered. The iOS device walk `ios/dark/officer-06-scrolled.png` shows the Roster section down to Close to 60, with the item-level rows and no gear gaps in frame. I checked those features in code only:
- `apps/ios/App/Guild/Logic/GuildTabPresenter.swift`: `gearGaps` (the three largest `addon.gearGap` values), `yourAttendance` (from `attendance30`), and the gold own-award mark.
- `apps/ios/App/Guild/Views/GuildTabViews.swift`: the group labels and order.

The data wiring looks right. The finding below asks for the full-page captures.

## Round-2 items

| Round 2 | Status |
|---|---|
| 1. Loot has no "you won" mark | Resolved in code: the winner shows in gold when it is one of your characters. Not seen rendered. Loot fairness (awards per raider) stays "needs data": the app reads `limit=1`, and the web loot page's per-raider counts are not in the app contract. |
| 2. Gear readiness only as item level on recent names | Resolved: Gear gaps shows the top three by `roster[].addon.gear_gap`, with spec and item level. It is a gap rank, not BiS progress. That is fine for a phone. |
| 3. Member attendance absent | Resolved: `7 of 8 nights` comes from `attendance_30`. Next upgrade is still absent. It needs an app-readable BiS field, so it is "needs data". Gear owns it. |
| 4. Ping on Discord posted-state | Open, see SHOULD 2. The label is now the web planner's wording, and the screenshots show the officer card and button. I did not verify the posted state afterwards. |
| 5. Per-raider parse and deaths | "Needs data", unchanged. |
| 6. Per-boss kills and wipes with dates | "Needs data", unchanged. Progression shows `1 of 1 down · 6 pulls`, which is enough for now. |

## What works at a glance (officer screenshots)
- **Raid card:** raid name, `8:00 pm CST · pull in 2 h 48 min`, invites time, `As Thandrel, Damage`, and the In / Late / Bench / Out control with the current state filled.
- **Needs you:** `5 have not answered, 2 of them healers`, `Healers short by 3`, and a Ping on Discord button, all in one card. This is the main officer win. It beats a spreadsheet plus Discord.
- **Officer rows:** `Plan a raid · Onyxia Wed, Barrow Deeps Sat`, `Applications · 1 waiting` with the applicant named, and `Setup · 4 of 8 done`. Each has a chevron and each is one tap.
- **Who is coming:** Tanks 2 of 2 (green), Healers 5 of 8 · 3 short (red), Damage 11 of 30 with the melee/ranged split, a Bench/Out/not-answered line, and missing buffs with icons (Fort, WF 1 of 2). That is the comp-gap check, with Late and its minutes shown inline.
- **Roster:** `30 of 38 at 60`, recently on with item level, close to 60 by level.

## Findings

### SHOULD
1. **Capture evidence is cropped.** The "full" renders stop at about one and a half screens, so the new Loot, Gear gaps, Your attendance, Missing raids and gold-name marks have no image anywhere. Re-render real full-page captures of the officer state so the round-3 additions can be judged visually. This is a capture fix, not a feature gap.
2. **Ping on Discord needs a visible after-state.** After the tap, the button should show posted, with the time or who was pinged, and not be tappable again. Confirm this in `interactions.md` or a capture. The API write is the one the web planner already uses. If the state is not shown, the officer cannot tell whether the ping went out.
3. **Roster is long for officers.** Gear gaps, Missing raids and Recently on are stacked in one panel under Roster. An officer wanting "who needs a talk" has to scroll past five recent names. Consider putting Missing raids and Gear gaps above Recently on for officers. The data is already there, so this is only ordering.

### NICE
- The assignments footer reads "Assigning is on the site." That is honest, and is the one planned hand-off.
- Recent-name rows are not tappable to a raider detail. That would need the roster page, which is already linked.

### Needs data (not built, not blocking)
- Loot fairness: awards per raider and attendance-weighted counts.
- Per-raider deaths, avoidable damage, and parse versus spec median.
- Boss kills and wipes with dates.
- A member's next upgrade and where it drops: needs an app-readable BiS field.
- An account-wide list of your applications.

## Verdict
Every round-2 item that the existing API can answer is now answered. The officer card does the main job: who has not signed, which role is short, one tap to Discord, and applications and setup one tap away. The members' views (answer, job, attendance, loot win mark) cover what a member checks before a raid. What remains is the missing capture evidence and the Discord after-state. Neither needs new data. As an away-from-PC companion, an officer would open this before raid instead of the spreadsheet, though the sheet still wins for loot fairness and performance until the "needs data" items exist.

VERDICT: WORTH PAYING FOR

Conditions: SHOULD 1 and 2 are confirmed before merge.
