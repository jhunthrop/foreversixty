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

# Rounds 3 to 5 (2026-11-14): the opt-out page, the officer marks, the bot copy

Captures in r3/, r4/, r5/ (web/tests/captures/character-directory: the opt-out page stubbed and
live, the marks on a fixture guild, the bot's pick messages dumped through its own handlers on a
test database). Surfaces: A /remove-my-character, B the officer's Discord-pick marks, C the bot.

## Round 3: A HOLD (ux, wow, slop, density), B HOLD (all five), C SHIP (all six)

- A: the panel headings did not render (Plate labels are aria only); the lookup answer sat below
  every panel; the Removed state said Removed three times; dangling copy fragments.
- B: the raid roster mark wrapped to four lines; the account table's tag collided with Remove at
  390; the tooltip was the only evidence; the default Roster board showed no mark; a capture
  fixture put the tag on the viewer's own row.
- C SHOULDs taken: a shorter taken line, and "remove" instead of "free" (the roster's Remove).

## What changed

- A: Section panels with visible headings (Your character, Remove with Battle.net, Remove with
  your guild note); the lookup answer under the name, the wait line under Check now; rewritten
  route lines; expiry line cut to the date; Removed drops the pill card for Remove another
  character.
- B: one shared mark, guild/DiscordPickMark.svelte (warn `Picked on Discord` tag plus visible
  "by name", the full sentence as its accessible name and title), on the raid roster (after the
  note), the Roster board and the account table (own line). A member's raid view carries no picker
  names (guilds.hidePickers).
- Privacy (density r4 MUST): the signed-out /code answers an unknown name like a held one (a code,
  never saved), the 404 and the page's not-found line are gone; /status answers only the holder of
  the character's code, so a typed name reveals neither existence nor an opt-out.
- C: taken `{Name} is claimed by another account. Ask an officer to remove it, or sign in with
  Battle.net.`; cap `... Ask an officer to remove one, ...`.

## Round 4: A SHIP (ux, wow), PASS (slop), HOLD (density, the existence probe); B SHIP/PASS (all); C SHIP/PASS (all)

## Round 5: A PASS (density), SHIP (ux); B SHIP (ux, which withdrew its one-line SHOULD: the
narrow cell cut "by br...", so the name wraps under the tag); C stands.

Open for the owner: the privacy wording ("delete what we hold for it and never store it again",
"will not store it again") is unchanged and unapproved; every reviewer flagged it, none rewrote it.
Left SHOULDs: centre the 560 px form column at wide widths (with /connect), a roster count of
unconfirmed picks, a link from the mark to Remove, a house-format timestamp on the spec card,
"Which spec tonight?" for raids days out, pre-existing spec truncation and roster counts.
