# Discord sign-up by character pick: the bot's messages

What a raider and an officer see, in order. Every string is the bot's own (copygen keys in
apps/shared/copy/server.json, rows "Press: not linked" and "Pick" of
design/specs/2026-10-27-discord-bot.md §6.2). Every reply is ephemeral (only the presser sees
it). Names are made up.

Code: api/internal/discordbot/signup_pick.go, autocomplete.go, cmd_signup.go, press.go,
link.go; api/internal/directory/claims.go, handover.go. Web mark:
web/src/components/guild/planner/PlannerRow.svelte.

## 1. A raider with no Forever Sixty account presses In on the raid post

> Run /signup and pick your character under **character**. You are signed up when you send it. Or link Discord to use your Battle.net characters.
>
> [Link Discord]

(The button stays `Link Discord`. Before: "Link your Discord account to Forever Sixty to sign up from here, then press the
button again." with the same button. The link still needs a Battle.net sign-in on the site.)

## 2. They type /signup and focus character

Autocomplete lists the guild's roster from the character directory (every member of the
guild's Blizzard roster, public or not), minus every character an account already holds,
highest level first, at most 25, narrowed as they type:

```
Vela Thornbrook · 41 Hunter
Orrin Ashmantle · 12 Warrior
...
```

The raid option offers nothing until they hold a character (the calendar stays members-only);
with no raid chosen, /signup takes the guild's next raid. Status defaults to In when left out.

## 3. They pick Vela Thornbrook and send

A Discord-only account is made (Discord id and display name only), Vela is claimed for it at
member rank (never officer, whatever the in-game rank), and the sign-up runs as for anyone:

> You picked Vela Thornbrook. Sign in with Battle.net to confirm it is yours.
> Vela Thornbrook is In for Onyxia's Lair as DPS, Thu 10 Dec. 14 of 40 signed up.
>
> [Change spec]

When the character has no known spec, the claim line sits above the usual "Which spec
tonight?" card instead.

## 4. From then on

Presses on the post work like any member's: one character, no select.

> Vela Thornbrook is Out for Onyxia's Lair, Thu 10 Dec.

## 5. Someone else picks the same character

> Vela Thornbrook is already claimed by another account. If it is yours, ask an officer to free it, or sign in with Battle.net.

## 6. A Discord-only account that has not picked anything presses a button

> Pick your character first: run /signup and choose it under **character**.

## 6b. A Discord-only account already holding 3 picks picks another

> You have picked 3 characters already. Ask an officer to free one, or sign in with Battle.net.

## 7. Someone signed in with Battle.net tries a roster pick

> You are signed in with Battle.net. Pick one of your own characters under **character**.

## 8. What officers see

On the raid page roster (web), under the character's row, officers only:

> Picked on Discord by raider#name, not confirmed by Battle.net yet

On the guild page roster the character is listed with a `Picked on Discord` tag (tooltip: the
same sentence), and the officer's Remove frees it: the pick and its sign-ups for coming raids
go, so the real owner can pick it or sign in.

When the owner signs in with Battle.net, the character, its roster row and its
sign-ups move to the owner's account, the mark goes, and nothing is duplicated; if the Discord
user then links the same Discord account from the site, the Discord-only account folds into
the real one.
