# Battle.net-only sign-in and the signed-in, no-character home: review verdicts

Captures were taken 2026-11-03 from a fixture build with a mocked `/v1/me`, covering three states: anonymous, signed in with no characters, and signed in with one character. Each state was captured on home, /sim, /account and the header menu, at widths 1440 and 390.

- ux-designer: round 1 NO-SHIP. The Get set up card was missing for the no-character state, and the /sim intro offered sign-in. Round 2: SHIP WITH FIXES, with no must-fix items left.
- wow-player: round 1 NO-SHIP. The /sim intro offered sign-in, and the Last raid line showed in an empty Characters panel. The guides icons were a fixture-build artifact. Round 2: SHIP.

Open shoulds, not blocking:
- /sim with no characters: the Your characters card is stranded in the left column at 1440, the damage-specs note appears twice, and "Load a character" shows with nothing to load.
- Header menu with no characters: there is no Get the addon row, and Your account is linked three times.
- /account with no characters: Paste an export should be a secondary button. The Korea line and the "pairing code fallback" wording need a plainer rewrite.
- /account with characters: "Officer Verified" has no separator, and the gear-sync line is cramped.
