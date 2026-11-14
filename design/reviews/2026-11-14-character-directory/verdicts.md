# Verdicts

## Round 1 (2026-11-14), on bot-messages.md as first written

## guild-officer-reviewer: HOLD

MUSTs: (1) an officer had no way to undo a wrong claim; removal left the character held;
(2) one Discord user could claim any number of characters; (3) the officer mark did not say
who picked. SHOULDs: claim line read as ownership; say who to ask on "taken"; shorter
"pick first".

## ux-designer: HOLD until MUST 1

MUSTs: (1) the not-linked press reply did not say the sign-up happens on send, and repeated
the link route the button already carries; (2) "pick first" lost the status the raider
pressed; (3) the claim line explained officer tooling to the raider. SHOULDs: tighter taken
and Battle.net refusals; officer tag tooltip in the same words as the row line.

## What changed

- An officer's Remove on a picked character frees it: its characters row and its sign-ups for
  coming raids go (guilds/roster.go releaseDiscordPick), and picked characters now appear on
  the guild page roster with a `Picked on Discord` tag, so Remove is reachable.
- At most 3 picks per Discord-only account (directory.MaxDiscordPicks), with its own refusal.
- The marks carry the picker's Discord name (`discord_picked_by` on the raid view and roster).
- Copy: every string the reviewers rewrote, as now in bot-messages.md.

# Round 2 verdicts (2026-11-14)

## guild-officer-reviewer: SHIP

Round 1 MUSTs closed in code (release on Remove, the 3-pick cap, the picker's name).
SHOULDs taken: a comment on the cap; the roster's picker name only reaches officers
(home.go). Left: naming the officer on "taken".

## ux-designer: HOLD, two copy MUSTs, both applied as worded

- The not-linked reply now names the Link Discord route the button carries:
  `Run /signup and pick your character under **character**. You are signed up when you send it. Or link Discord to use your Battle.net characters.`
- "Pick first" drops "set your status" (status defaults to In):
  `Pick your character first: run /signup and choose it under **character**.`

With both MUSTs applied word for word, the copy is as the designer asked; SHOULDs (a shorter
"taken" line, a touch fallback for the roster tooltip) are left for the next pass. The picker
name is the display name stored at the first pick (discord_links.display_name).
