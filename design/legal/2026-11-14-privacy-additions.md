# Privacy page additions: character directory (draft)

Status: DRAFT for the owner's approval. Not published. Until it is approved and merged into
`web/src/content/pages/privacy.md`, the character directory may fill, serve the Discord
character pick and help the Battle.net import, but no public page shows a character of a player
who has no account.

Basis: docs/superpowers/specs/2026-11-14-character-directory-discovery.md §5 and §6; Blizzard
Developer API Terms of Use 2(m), 2(p), 2(s), 11(c). Each block below names where it goes in
the current page. The `updated:` date moves to the day it is published.

## 1. Under "What is collected", a new paragraph after the first

> We also read public character information from Blizzard's Battle.net API for characters of
> players who have no account here: the character's name, level, class, race, faction,
> guild and guild rank, realm type, item level, talent specialization, professions,
> equipped items, portrait and last login. It comes from public guild rosters and public
> character profiles only. We never read a BattleTag, an account or anything that needs a
> player to sign in, and we never link this information to a person until that player
> signs in and claims the character.
>
> If you sign up for a raid from our Discord bot without an account here, we create a
> Discord-only account that holds your Discord user id, your Discord display name and the
> character you picked. It has no email, no BattleTag and no password.

## 2. "How Battle.net data is used", replacing its one paragraph

> Character information comes from Blizzard Entertainment's Battle.net API. Blizzard does not
> endorse this site. We use it to know who is in a guild, to fill guild rosters, to let a
> guild's members sign up for raids from Discord, and to work out a character's realm type
> when it is imported. Anything sourced from Blizzard's Developer API is checked against
> Blizzard and refreshed at least every thirty days, and deleted when Blizzard says the
> character no longer exists or when it has not been refreshed within thirty days, per
> Blizzard's own Developer API Terms of Use. It is never sold or placed behind a paywall.

## 3. A new section after "How Battle.net data is used": "Remove your character"

> To have a character removed from what we read from Blizzard, and kept out from then on,
> use [Remove my character](/remove-my-character). Sign in with the Battle.net account the
> character is on and it is removed at once; or, without an account, put the code that page
> gives you in the character's public guild note, and it is removed the next time a
> guildmate who runs our addon logs out. Once removed, we delete what we hold for the
> character and never store it again. You can also email
> [help@foreversixty.gg](mailto:help@foreversixty.gg). A character Blizzard stops sharing,
> because it was deleted or its data was withdrawn through Blizzard, is deleted here at our
> next check, within thirty days.

## 4. Under "Data retention", a new paragraph after the first

> A Discord-only account is kept until its player asks us to delete it or signs in with
> Battle.net. When that player signs in with the Battle.net account that owns the picked
> character, the character and its raid sign-ups move to that account; linking the same
> Discord account from the site merges the Discord-only account into it and deletes it.

## 5. The last sentence of "Requests and contact", replacing it

> Those rights are yours wherever you live. The site runs no advertising and sells no data.
> It keeps public game information per character, as described above, not profiles of
> people, and joins it to a person only when that player claims the character.

## Notes for the owner

- Wording 1 lists exactly what the directory stores today (`character_directory` columns);
  statistics, reputations, quests and collections are readable without sign-in but not kept.
- The guild-note code reaches us only through the addon's guild roster snapshot, which sends
  the code alone, never the rest of a note (docs/contracts/addon-v2.md, `optout_codes`).
- Wording 5 is the reading of "builds no profiles" the discovery doc asks for; the old
  sentence can stay if the owner prefers, but then the directory needs this explanation
  elsewhere on the page.
