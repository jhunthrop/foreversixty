# Addon UX, persona review round 1

Spec `design/specs/2026-11-08-addon-ux.md` and the round-1 boards (commit 551a6f4c2). Three `wow-player` reviews, one per persona. All three: **HOLD**. Dispositions below; round 2 boards and spec carry the fixes.

## Leveller (hunter through the 20s to 40s): HOLD

Would never look at: Forever Score on a level 24 (panel and tooltip parts), Copy for the site, the player tooltip's live line, the "Beast Mastery estimate" note, the foreign "Right click for Frame Settings" line.

| # | Finding | Disposition |
|---|---|---|
| M1 | Verify the BM rotation (Aimed Shot, Arcane Shot missing, pet) and "Lethal Attacks" | **Verified, kept.** The rotation is `Data.lua` `rotations["hunter-beast-mastery"]` band 22 verbatim; Forever's BM band carries Aimed Shot. Lethal Attacks is talent 105011, BM tier 1, `talents/hunter.json`. Arcane Shot was cut off at four rows: the panel now shows every line of the band under 60 (six at most). Spec §4.3 carries an accuracy note. |
| M2 | Hide the score without logs; units on the gap and parts | **Accepted.** Rule unchanged (no record or no rated fights: no line); Bow Jackzon has a real record in the capture. The live line is cut; `Preparation (flasks, food, enchants)` says what it measures. |
| M3 | Hammer "(thrown)" is wrong; +0.8 vs +0.3; Chausses lacks a level-gated mark on the reward | **Thrown kept (correct):** Cracked Blacksmith Hammer is item subclass 16, inventory type 25, a thrown weapon in Forever's 1.60 data. **+0.8 vs +0.3: consistent**: +0.8 is under the 1.0 floor, so the hammer's own verdict is `Sidegrade`; the states board caption now says so. **Chausses:** no label on the frame (clutter), the hover says `Not for you until level 40` (§2.4). |
| M4 | Truncated talent subtitle | **Fixed.** `Beast Mastery leveling · Change`. |
| S1 | Armor and Stamina on Shift | **Accepted**, §2.3 item 1. |
| S2 | Pet and ammo rows | **Deferred** to a follow-up spec; not in the owner's priority list and adds height. |
| S3 | Required level on panel rows | **Accepted**, `· level 27`. |
| S4 | Chip prefers an actionable state | **Already so** (Take, Equip, raid outrank the idle line); spec says it explicitly. |
| S5 | Drop Copy for the site and "estimate" | **Accepted**, both cut. |
| S6 | Foreign tooltip line in the after mocks | **Accepted**, removed from the after boards (kept in the before board: it is in the real capture). |

## Raider (level 60, raids twice a week): HOLD

Would never look at: Rotation at 60, the finished Talents row, the Score row, the always-on Change build and Sim me.

| # | Finding | Disposition |
|---|---|---|
| M1 | The 60 panel is a leveller's panel | **Accepted.** Rotation collapsed at 60 (expands once after a new rank), finished Talents region not drawn, score moved to the head, footer cut. |
| M2 | In-game In/Out instead of "answer on the app or the site" | **Accepted within the no-network rule.** `In`, `Late`, `Out` write a `raid_answer` entry to the companion outbox (the path Sim me uses); the row reads `sends when the companion runs` until confirmed. The app line is gone. Contract gaps named in §4.3. |
| M3 | Mock the loot roll frame | **Accepted.** `loot-roll-after.png`: arrow, `+5.8 DPS`, `Best in slot`, `Not for you`, no source. |
| M4 | "With Thornhide, Brightwen and 7 more" | **Accepted.** One or two names only (`· with Thornhide`), else nothing; the purpose moves to the hover. |
| M5 | Chip on raid day | **Accepted.** New top state `Tonight 8:00 pm: Onyxia's Lair`, on the chip board. |
| S1 | Explain Preparation | Accepted. |
| S2 | Move Change build and Sim me | Accepted (footer cut). |
| S3 | Drop "Epic" | Accepted. |
| S4 | Raid-loot state | Covered by the loot roll board. |
| S5 | Phase 2 wording | **Out of scope here**: the label and purpose are the API's slot templates (`api/internal/guilds/raid_assign_templates.go`); flagged to the coordinator for the guild lane. |

## Player who hates addon clutter: HOLD

Would never look at: the welcome window, the rotation block, the Tonight paragraph, the Score row and footer buttons, the player tooltip's Shift extras, "check the AH", the chip hover repeating the chip.

| # | Finding | Disposition |
|---|---|---|
| M1 | Bound the pulse | **Accepted.** Five cycles (8 s) after each login, stops for good on the first click or hover, gives up after three logins, never in combat; right-click hides the chip. |
| M2 | Welcome as a small dismissible thing, never again once dismissed | **Accepted in form.** The owner approved a welcome with pointers, so it stays, now a tooltip-style callout docked to the chip (no window, no portrait, nothing centred), two pointers plus the chip itself, dismissed for good by Got it, Escape or a click on the chip. |
| M3 | Rotation, Sim me and Score out of the default panel | **Accepted** at 60 (collapsed rotation, Sim again only as a note link with a sim, score as one number in the head). Under 60 the rotation stays expanded: the leveller persona relies on it and the owner listed it. |
| M4 | Remove "check the AH" | **Accepted.** |
| S1 | Gold is a fourth colour | **Kept.** Gold (`NORMAL_FONT_COLOR`) is the client's own; the raider and leveller both need "best in slot" told apart from "an upgrade". |
| S2 | No arrow glyph in tooltips | **Kept.** One mark language (bag, slot, reward, roll, tooltip); established practice in tooltip addons. |
| S3 | Cut the Shift live line and link hint | Accepted. |
| S4 | Tonight job on one line | Accepted. |
| S5 | A master quiet mode | **Accepted**: `Tooltips only` switch (§3.3). |
| S6 | Real tooltip context instead of "…" | **Partly.** The full-context before/after boards exist (`item-tooltip-after*.png`, `quest-reward-after.png`); the states board keeps "…" to fit twelve cases on one board. |
