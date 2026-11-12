# guild-officer-reviewer, round 1 (guild-tab-undecided.png and onboarding-4-guild.png)

**Page job.** The guild tab and its tour page tell a raider what the guild is doing and how they fit in; for an officer the roster is a quick read of who is raid-ready.

**Empty cell versus a guess.** Empty beats a wrong guess: spec drives the comp count, the gear-gap unit and the "who needs what" read, and a 2/2 paladin shown as Holy would put a wrong healer count into those. The rule is sound (a single point never decides, 8/4 asks, 10/3 and 6/1/1 are sure). The cost: the empty cells carry no explanation, and a level 13 and a 24 look like a member who never synced; their Gear and Build cells are empty too.

## Findings

1. SHOULD: the empty cell should say why on hover. For others, a dim hover with the split ("Spec not set: 12/9 points, no clear tree"), the cell itself blank; never a leaning label. For the own row, "Pick your path" as a click target that opens the tiles.
2. SHOULD: the row loses its other data (gear gap and build depend on the spec); level and score from the record should still show.
3. NICE: the board's caption is cut off ("so hers shows"); re-render.
4. NICE: add an undecided level 60 row, the case officers care about.

**The guild page's lead.** Confirmed in the addon: the lead string, the Tonight text, RaidAssign reading the next raid, sign-ups queued through the panel's Tonight buttons. Not checked here: the API's assignment templates and the Discord bot's raid post. The wording is honest and short and makes only conditional promises; "coming" on the in-game and phone line avoids over-promising.

**Ranked wants for the empty cell:** a hover with the split (the addon's own data for its own row; others' splits need data the hello does not carry); an officer's click to ask a player to set their spec (needs a whisper or Discord path, none today); never "undecided" as cell text on a 40-row roster.

**VERDICT: CLOSE.** Keep the change; add the hover or prompt so an empty cell is not mistaken for a missing sync; re-render the caption and add a 60 row.
