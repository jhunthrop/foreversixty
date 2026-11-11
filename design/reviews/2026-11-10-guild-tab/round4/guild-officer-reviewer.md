# Round 4 verdict: Guild tab in the native apps (officer view)

**Page job:** the officer opens the Guild tab on a phone before raid, sees what needs them (unanswered signups, applicants), who is coming against the comp, who is slipping, and what dropped, and acts from there (Mark In/Bench, Ping on Discord, open the planner).

## Condition 1: full-page captures. MET.

Read `parity/ios-full-guild-tab-officer.png` (786x7899) and `parity/android-full-guild-tab-officer.png` (721x8050); the `-guild-tab.png` files are whole pages too. Same section order on both platforms:

- **Raid card and Needs you:** Onyxia's Lair at 8:00 pm CST, pull in 2 h 48 min, In / Late / Bench / Out; "5 have not answered, 2 of them healers" with Ping on Discord, plus "1 applicant".
- **Officer and Assignments:** Plan a raid, Applications `1 waiting`, Setup `4 of 8 done`; the assignment slots (`1 of 1`, `3 of 3`, `0 of 2`) and the planner link.
- **Who is coming:** Tanks 2 of 2, Healers 5 of 8 · 3 short, Damage 11 of 30 · 19 short, bench count, not synced.
- **Before you pull:** attunement, consumable and enchant checks, with an Auction house link on the missing flask.
- **Roster:** Recently on, Missing raids (`2 of 8 nights`), Gear gaps (`41 DPS`, `37 DPS`, `22 DPS`), Your attendance (`7 of 8 nights`).
- **Loot and logs:** Loot with Head of Onyxia "Not awarded yet", Recent logs, Progression, Pages.

My round-3 asks (Missing raids, Gear gaps, attendance) are all present, with gaps in DPS and ranked. The loot award in gold is not visible at the size I viewed (the item names are rarity purple); not a blocker.

Weaker spots, not blocking: Healers/Damage name lists truncate ("Aenwyn, Brightwen (Lat...") on iOS more than Android.

## Condition 2: Ping on Discord after-state. PARTLY MET.

**Code and tests (met):** `MentionWords.line` gives `Mentioned 2 healers in Discord.`, the left-out note, and `Nobody to mention: none of them has linked Discord.`; `NotAnsweredModel.mention` stores `mentionNote` on success and routes a refusal through `SignupRefusal.sentence`; a `mentioning` flag blocks a double tap. `NotAnsweredTest` pins the strings; iOS `RaidFlowTests` covers `mentionStatus` and the no-bot case. The officer gets a visible, accurate, persistent result in words.

**The device capture (not met):** `officer-07-ping-list-settled.png` shows "Not answered 0", "Everyone has answered." The fixture raid has no `not_signed_raiders`, so Mention could not be tapped; the result line has not been seen rendered. The Needs you card (5) and the list (0) disagree for the same fixture. Accepted for the merge because the logic is unit-tested on both platforms and the strings are the ones the officer push uses.

## Findings

- **SHOULD:** Seed `not_signed_raiders` in the fixture and record one real Mention tap per platform (result line and left-out line), right after merge; it also fixes the 5 vs 0 contradiction.
- **SHOULD:** Crop-check the gold award on the Loot rows.
- **NICE:** Stop the Healers/Damage name lists truncating on iOS, or tap to expand.
- **NICE, needs data:** per-raider performance vs spec median and deaths.

## Value test

An officer would open this before raid instead of the sheet: countdown, comp gaps, who has not answered with one-tap Ping, who has not synced, who is missing nights, who is under-geared with a DPS cost. A member gets Your job, Your attendance and the Before-you-pull checks. Worth paying for on a phone; the remaining gaps are evidence gaps, not feature gaps.

Condition 1 met in full. Condition 2 met in code and tests, the on-device capture open as a SHOULD.

VERDICT: WORTH PAYING FOR
