# Player verdict: mobile sign-in, account, switcher (2026-10-31)

Verdict: HOLD

Seen: signedout, signin-failed, switcher, you, devices, delete, no-characters (iOS), switcher/delete-confirm/devices-revoke-confirm (Android). Not seen: a Premium delete screen, the Battle.net consent sheet, a character with a spec-less or non-Orc row.

## Blockers
1. The switcher never shows a realm or a ruleset. Rows read "Orc Fury Warrior · Level 58" over "OLYMPUS XXVII" (which a player reads as a realm but is the guild) or "No guild". Two same-named characters on different realms, or a Forever and an Era-ruleset character, cannot be told apart. Add realm (and region where it differs) and the ruleset to each row; keep the guild on its own line.
2. Delete account never tells a paying player what they lose. The list covers data only. The confirm dialog is just "Delete your Forever Sixty account? This cannot be undone." Nothing on screen says a Premium plan is cancelled with no further renewal, the founder price (renews 27 Oct 2027, to 6 Jan 2027 window) is gone for good, or that a guild plan or seats you gave out drop to free. The cancel line exists only in the report, and no Premium delete capture shows it. Put the plan line on the screen and in the dialog.
3. Unsettled promise on guild reports. The account page says reports are never rewritten; the privacy page says they go with the account. The delete list says "the reports parsed from them" without saying whether raiders' reports vanish for the guild. Owner decision needed before this ships.

## Shoulds
- Hunter row shows "Orc Hunter · Level 31" with no spec, and the selector chip shows "Fury · 58" only. Fine for Fury; show a spec (or "No spec yet") consistently.
- Devices: "Companion on Windows" uses a phone icon; Android shows a telephone handset for "This phone" and "Pixel 9". Use a monitor for the companion and a smartphone glyph.
- Sign-in trust: the screen never says Battle.net is read-only, or that Forever never sees your password. Add one line under the button.
- "Pick your class opens best in slot, talents and sims for any spec, no account needed." Reads clumsy; say "Pick your class to see BiS, talents and sims for any spec. No account needed."
- No-characters: "Battle.net lists a character once it has logged in on Forever" is wrong; Battle.net does not list it, we do. Say "Forever lists a character once it has logged in with the addon."
- Switcher footer "after it logs in on Forever, or after the addon syncs it" repeats the same idea; one sentence.
- Revoke dialog: "It is signed out of Forever Sixty at once." name the device kind and say it can sign in again.
- Delete: "The rest is gone within thirty days" is vague; say "Everything listed is deleted within thirty days."
- Live sign-in cannot finish until the Apple Team ID and Android fingerprint land (report concern 6); do not ship a build with sign_in on before then.
