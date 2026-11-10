# Round 1: ai-slop-reviewer (captures in round-1/)

## First impression

This reads as made by someone who plays WoW. It has the site's painted band and the seal, real zone and boss names, a monospace sign-in code and a real `/combatlog` instruction. It does not read as generated. The problems left are one nested card and a lot of grey explainer copy.

## Tells, most damaging first

### Tell 8: a card inside a card

`settings.png`, Game folders: the folder row is a bordered, rounded box inside the Settings panel, which is itself a rounded panel.

Fix: make the folder rows plain rows split by a hairline.

### Tell 22: qualifiers the context already makes clear

- `device-connected.png`: "A private store in your user profile." sits under "Windows Credential Manager".
- `logging-credited.png`, Addon row: "up to date" appears twice under "Last synced 1 minute ago".
- `logging-credited.png`, headline: the sub-line restates "Logging".

Fix: show "Windows Credential Manager" alone. Show "Addon 2.0.16 · Data 2026.11.09" with no state words unless one is stale.

### Tells 7 and 21: grey explainer sentences

- `settings.png`: "Anyone can find it in rankings and search.", the CurseForge/Wago sentence, and "Automatic reads who is logging from the combat log itself. Pick a character only if that is wrong."
- `reports-empty.png`: "It switches off if you type it again, so leave it if it is already on."
- `sign-in-waiting.png`: "Check that your browser shows this code:".

Fix: keep the password note under Sign in, because it is a real trust point. Cut the rest to one clause each or drop them.

### Tells 11 and 13: the empty state

`reports-empty.png` is a centred stack: a teal glyph, a serif heading and a grey paragraph in a big empty panel. That is the stock empty-state shape.

Fix: left-align it, drop the glyph, and leave the `/combatlog` line as the only content.

### Tell 9: an icon on every action

Fix: keep icons on the nav and the spinner states. Make the links text-only everywhere else.

### Repeated buttons

`reports.png` has eight identical blue "Copy link" buttons with icons. They are louder than the report names.

Fix: use a plain text link, or an icon-only control.

### Button roles

Blue-filled secondaries (Open my reports, Open the browser again, Copy link) compete with the teal primary.

Fix: make it clear which is primary and which is secondary.

### No instance found

- Em dashes: none in copy.js, index.html or app.js.
- Tells 1, 2, 3, 4, 5, 12, 14 and 15: none. The text-shadow halo on headings is minor and defensible on the painted ground.

## Keep

- The band and the seal.
- The ornate frame on the lead panel only.
- The deliberately uneven hierarchy.
- The mono sign-in code and its countdown.
- The real report rows with state dots.
- Layouts that are not centred.
- The real Save state.

VERDICT: FAIL — The nested Classic beta folder box inside the Settings panel (`settings.png`) is a tell-8 card-in-card that blocks a pass, and the grey explainer lines and repeated "up to date" qualifiers need cutting too.
