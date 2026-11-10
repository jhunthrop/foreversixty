# Round 2: wow-player (captures in round-2/)

## Fixed

- The /combatlog lines, "Log out once" said once, and the credit named once.
- The offline bar. The reason it has no Restart is accepted.
- The update bar, the ellipsis and "Fetch my builds from the site".
- Fight counts now agree. The sign-in expiry bar is consistent.
- The Current report row with "Log written".

## Still wrong

1. **`uploading-held`.** The wording is good, but the headline and the tray stay green while uploads are stuck. A held queue should turn the headline and the tray amber.
2. **`device-connected`.**
   - "Log engine 1.0.0" means nothing to a player; call it "Parser version" or drop it.
   - "1 report, 0 fights" contradicts the 3 fights on Status.
   - Confirm "Addon last seen in game" shows when there is data.
3. **`offline-status`.** The stale rows still say "0 so far / The first fight uploads when it ends." Hide that help when stale.
4. **`reports`.** No boss, no kill or wipe. Accepted as a follow-up.
5. **`update-ready`.** "Installs the next time the companion starts" does not say when that is. Acceptable.

## Data doubts

- Confirm the `_classic_beta_` folder name after launch with the integrator.

VERDICT: SHIP WITH FIXES — copy is clean now, but a held upload queue still shows a green "Logging" headline and tray tone, `offline-status` still says "0 so far", and `device-connected` shows "Log engine 1.0.0" and a "0 fights" that contradicts Status.
