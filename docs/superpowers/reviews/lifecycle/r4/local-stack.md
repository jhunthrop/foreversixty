# The local stack for walking the guild workflows

The guild master's claim, officer verification, roster management, settings, billing, invites, the loot council and the guild report tabs have only ever run behind Playwright fixtures. This stack puts the real API, a real database and the real site on `localhost`, seeded with one unclaimed guild, so each can be walked in a browser. Everything under `tools/dev-stack/` and `api/cmd/seedguild/lifecycle.go` is local only; nothing here deploys.

## Start it

From a Windows shell in the repository (the scripts run in WSL; there is no Go on Windows):

```
wsl -e bash -lc "tools/dev-stack/seed.sh"     # starts Postgres and the API if they are not up, then seeds
cd web && npm run dev:local-api                # the site on http://localhost:4321, pointed at the API
```

`seed.sh` calls `up.sh` first, so one command brings up everything the site needs. Both are safe to run twice: a running Postgres or API is left alone, and a second seed prints "already seeded" and writes nothing. The first run builds the API and creates a throwaway Postgres cluster in `~/.cache/fs-dev-stack` inside WSL (about a minute); later runs are seconds.

From the Browser pane the same two things are the launch configurations in `.claude/launch.json`: `api-local` (port 8080, runs `up.sh --foreground`, so stopping it stops following the log, not the API) and `web-local-api` (port 4321). Seed once with `seed.sh`; the launch configurations only start the processes.

| What               | Where                                                                   |
| ------------------ | ----------------------------------------------------------------------- |
| Site               | `http://localhost:4321` (`npm run dev:local-api`)                       |
| API                | `http://127.0.0.1:8080`, reached by the site as `http://localhost:8080` |
| Postgres           | `127.0.0.1:5440`, database `forever_dev`, user and password `forever`   |
| API log            | `tools/dev-stack/logs/api.log` (Postgres: `postgres.log`)               |
| Check it           | `wsl -e bash -lc "tools/dev-stack/smoke.sh"`                            |
| Stop it            | `wsl -e bash -lc "tools/dev-stack/down.sh"` (`--wipe` deletes the data) |
| Reset to the start | `seed.sh --remove` then `seed.sh` (new user ids: sign in again)         |

The site reads one variable, `PUBLIC_API_BASE_URL`, and `npm run dev:local-api` sets it to `http://localhost:8080`. It must be `localhost`, not `127.0.0.1`: the session cookie is host-only and `SameSite=Lax`, so the site on `localhost:4321` and the API on `localhost:8080` are the same site and the cookie travels, while `127.0.0.1` for either one would not. The same script sets `FS_LOCAL_GUILD_SHELL=1`, which makes `astro dev` serve the guild page for any `/guild/...` URL (in production the Worker does that; `astro dev` has no Worker). Without it every real guild URL is a 404 in dev.

## Signing in

There is no Battle.net here. The API has a dev sign-in, `POST /v1/dev/login {"battletag": "..."}`, which mints a session for an account that already exists and never creates one. It is gated three ways and must never ship enabled: the route is mounted only when `FS_DEV_LOGIN=1`; the API refuses to start with that flag unless `LISTEN_HOST` is a loopback IP and `PUBLIC_BASE_URL` is not the production origin (`api/internal/config`); and the handler answers 403 to any peer that is not on loopback (`api/internal/auth/devlogin.go`). `up.sh` is the only thing that sets the flag. The tests are `api/internal/config/devlogin_test.go` and `api/internal/auth/devlogin_test.go`.

Sign in from the browser console on a page of the site (`http://localhost:4321/...`), then open the page again. Paste this, changing the battletag:

```js
const csrf = (document.cookie.match(/fs_csrf=([^;]+)/) || [])[1];
const r = await fetch("http://localhost:8080/v1/dev/login", {
  method: "POST",
  credentials: "include",
  headers: {
    "Content-Type": "application/json",
    ...(csrf ? { "X-CSRF-Token": csrf } : {}),
  },
  body: JSON.stringify({ battletag: "hunthrop#1894" }),
});
localStorage.clear(); // the site caches /v1/me and friends per browser; stale entries show the previous user
r.status; // 200, then navigate to the page you want (a reload, not just a hash change)
```

The CSRF header is only needed when switching users with a session already in place. To sign out, `fetch('http://localhost:8080/v1/sessions', {method: 'DELETE', credentials: 'include', headers: {'X-CSRF-Token': csrf}})` and clear local storage the same way.

## The seeded world

Region `us`, ruleset `pvp`, every character level 60 and Alliance. The guild is **Lifecycle Test** (`us/pvp/lifecycle-test`), unclaimed. The seed goes through the same stores a real sync uses, so the guild, memberships and factions are what production code writes.

| Sign in as          | Account holds                             | In Lifecycle Test                                                                 |
| ------------------- | ----------------------------------------- | --------------------------------------------------------------------------------- |
| `hunthrop#1894`     | `us/pvp/bow-jackzon`, `us/pvp/pipsqueak`  | Bow Jackzon (hunter) is the guild master, rank index 0. Pipsqueak is in no guild. |
| `Ashenveil#2201`    | `us/pvp/ashenveil` (paladin)              | An officer, rank index 1. Not yet verified.                                       |
| `Quillfeather#3301` | `quillfeather`, `rootbound`, `cindergale` | Three members, rank indexes 3, 4 and 5. Not yet verified.                         |
| `Wanderer#4401`     | `us/pvp/driftwood` (rogue)                | Nobody: the stranger, and the account that opens an invite link.                  |

What else is in the database:

- **The guild master is verified as a Battle.net import would leave her** (`verified_by = 'bnet'`). Without that the site shows her as a plain member and the claim page refuses her ("you are a member of this guild, so you cannot claim it"), because an export alone verifies nobody. All four accounts carry a placeholder Battle.net subject, since claiming as guild master requires a linked Battle.net account. The officer and members are unverified, so the roster opens with "4 waiting for approval".
- **One synced export per character**, in the addon's FS1 format, with a `guild=Lifecycle Test:<rank>` section, gear (BiS picks for the hunter, rogue and mage), talents and professions. Members share gear with the guild (the default consent is `gear`).
- **The addon roster snapshot** the guild master's outbox would have sent: seven lines, the five members plus two people with no account (Strayhound and Oldmarrow), seen two hours ago. The roster's "In game" column reads from it.
- **One complete report**, `seedguild-lifecycle-report-ony`, "Onyxia" on yesterday's date, owned by the guild master, visibility `guild`: four fights (two whelp packs and two Onyxia kills, the only boss the game data publishes) with a metrics row for each of the five members. Progression shows one named encounter down, 4 pulls, best DPS Bow Jackzon. Its stored events are not in object storage (there is no R2 here), so the report's own detail page has nothing to read; the guild tabs read the database rows and work.
- **Not seeded**: ratings, gear history, a claim, an invite link, billing.

## The workflows, in the order a guild master meets them

Open each page as the account named; sign in first (above). The guild page's tabs are hash routes: `#overview`, `#roster`, `#raids`, `#progression`, `#readiness`, `#loot`, `#settings`. The Playwright specs under `web/tests/e2e` already cover each flow against stubbed API answers; they say what is known to render, and this stack is for what they cannot: the real API in the loop.

| #   | Flow                                                                                                                                                                                                                                                                                                        | As                                                                    | URL                                                    | Covered today by                                                                                                                                             |
| --- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------- | ------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| 0   | Find the guild: the landing lookup, the header doors                                                                                                                                                                                                                                                        | any                                                                   | `/guild`, `/guild/us/pvp/lifecycle-test`               | `guild-landing.spec.ts`, `guild-member-view.spec.ts` (the header doors)                                                                                      |
| 1   | **Before the claim, what a member sees.** Quillfeather is a member waiting on an officer: no roster rows, the line saying what opens it. Wanderer is the stranger: told the roster is not theirs, no claim button. A signed-out visitor sees the three public tabs and the sign-in prompt.                  | `Quillfeather#3301`, `Wanderer#4401`, signed out                      | `/guild/us/pvp/lifecycle-test`, then `/claim`          | `guild-member-view.spec.ts` (the guild page), `guild-centre.spec.ts` (tab strip, public visitor), `guild-claim.spec.ts` (stranger, signed out, plain member) |
| 2   | **Claim.** The guild master claims at once ("claimed by you"); the page says who can act and the claiming rules. One attempt per account every 30 days and one claimed guild per account: reseed (below) to claim again. The officer, once verified, would start a claim a second officer confirms instead. | `hunthrop#1894`                                                       | `/guild/us/pvp/lifecycle-test/claim`                   | `guild-claim.spec.ts` (claim, pending and confirm, contested, release), `guild-centre.spec.ts` (claim, contest)                                              |
| 3   | **Verification of officers and members.** Roster tab: "4 waiting for approval", Approve per row or Approve all, Remove per row. Approve Ashenveil, then sign in as her: she is now a verified officer and sees Guild settings in place of the claim entry. A member never sees Approve or Remove.           | `hunthrop#1894`, then `Ashenveil#2201`, `Quillfeather#3301`           | `/guild/us/pvp/lifecycle-test#roster`                  | `guild-centre.spec.ts` (roster: pin, sort, officer controls; approve; remove), `guild-member-view.spec.ts` (verified officer)                                |
| 4   | **Roster.** Sorting, filters, the In game column from the addon snapshot ("5 of 5 members seen"), the officer-only controls, rank labels, Open in simulator and planner.                                                                                                                                    | `hunthrop#1894`                                                       | `#roster`                                              | `guild-centre.spec.ts` (roster; In game column)                                                                                                              |
| 5   | **Settings.** Default report visibility (public, unlisted, guild), the officer rank threshold, the guild crest, the claim state. After a contest the controls are disabled with the frozen notice.                                                                                                          | `hunthrop#1894`                                                       | `/guild/us/pvp/lifecycle-test/settings` or `#settings` | `guild-settings.spec.ts`, `guild-centre.spec.ts` (crest)                                                                                                     |
| 6   | **Consent levels** are per member, not guild-wide: roster only, gear, or gear and bags. Change one on the account page and watch what the guild master's roster shows for that member.                                                                                                                      | `Quillfeather#3301`, then `hunthrop#1894`                             | `/account`                                             | `guild-account-consent.spec.ts`                                                                                                                              |
| 7   | **Billing.** Settings has a Billing section. Not subscribed: the honest line and a Subscribe link (`/premium/checkout?plan=guild&interval=monthly&guild_id=<id>`), which needs Stripe test mode (below). The plan state itself needs no Stripe: grant it from the CLI and reload Settings.                  | `hunthrop#1894`                                                       | `#settings`                                            | `guild-settings-billing.spec.ts` (not on the plan, billing contact, not the contact, CLI grant)                                                              |
| 8   | **Invites.** Settings, Rotate invite link (with the leak warning). Open the link the API returns, `/guild/invite/<token>`, as a signed-in account with no guild to join as a member; signed out you get a sign-in prompt.                                                                                   | `hunthrop#1894`, then `Wanderer#4401`                                 | `#settings`, then `/guild/invite/<token>`              | `guild-invite.spec.ts`, `guild-settings.spec.ts` (rotate)                                                                                                    |
| 9   | **Loot council.** Loot tab: Onyxia's items ranked by candidate from the roster's gear and attendance. An officer sees Award; a member sees every candidate and no Award. Award an item and the awarded state shows for everyone.                                                                            | `hunthrop#1894`, then `Quillfeather#3301`                             | `#loot`                                                | `guild-centre.spec.ts` (loot, read-only for a member, awarded state)                                                                                         |
| 10  | **Reports, raids, progression, readiness.** Overview's "this week's raid nights" lists the seeded night; Raids lists it with its pulls; Progression shows Onyxia down; Readiness ranks the members worst-first and an officer can copy a Nudge. The `/logs` page grows a tab for the guild.                 | `hunthrop#1894`                                                       | `#raids`, `#progression`, `#readiness`, `/logs`        | `guild-centre.spec.ts` (raids empty state, readiness), `logs-guild-tabs.spec.ts`                                                                             |
| 11  | **Leaving the guild.** Account page, the guild row, Leave. The membership row goes, and a claim the account held is released with it.                                                                                                                                                                       | `Quillfeather#3301` (or the guild master, to watch the claim release) | `/account`                                             | `guild-account-consent.spec.ts`                                                                                                                              |
| 12  | **After the claim, what a member sees.** Reload as Quillfeather once the guild master has approved her: the roster rows, the standing line, no Approve. Compare with step 1.                                                                                                                                | `Quillfeather#3301`                                                   | `/guild/us/pvp/lifecycle-test#roster`                  | `guild-member-view.spec.ts` (verified member), `guild-centre.spec.ts` (member never sees controls)                                                           |

Starting again: `wsl -e bash -lc "tools/dev-stack/seed.sh --remove"` deletes the seeded accounts, characters, guild and report (and with them any claim, invite, award or session), and `seed.sh` writes the world back. User ids change, so sign in again.

## Billing without Stripe

With no Stripe keys the billing routes answer 503 by design, so Subscribe and Manage billing do nothing here. To see a subscribed guild without Stripe, grant the plan from the CLI the way an operator would (the guild id is in `/v1/me` under `guilds`, or printed by `seed.sh`):

```
wsl -e bash -lc 'cd /mnt/c/Users/jhunt/Documents/Repos/fs-guild-local-stack && . tools/dev-stack/lib.sh && \
  DATABASE_URL=$DB_URL API_BASE_URL=http://localhost:8080 PUBLIC_BASE_URL=http://localhost:4321 RESEND_API_KEY=x \
  $FS_DEV_STATE/api grant --guild <id> --plan guild --until 2026-12-31 --note "local walkthrough"'
```

Settings then shows the renewal line with no "Billed by" (no billing contact), as `guild-settings-billing.spec.ts` pins. `api revoke --guild <id> --plan guild --note ...` undoes it. A real checkout needs Stripe test mode: export `STRIPE_SECRET_KEY` (an `sk_test_` key), `STRIPE_WEBHOOK_SECRET` (from `stripe listen --forward-to http://127.0.0.1:8080/v1/billing/webhook`) and `STRIPE_ENVIRONMENT=test` in the WSL shell before `up.sh` (the API inherits them; `PUBLIC_BASE_URL` is `http://localhost:4321`, so the API's own live-key check passes), run `api stripe-setup` once to create the prices, and pay with Stripe's test card. Not exercised here: no Stripe account is attached to this lane.

## Things the walk will run into

- **A stale view after switching accounts.** The site caches `/v1/me` and similar answers in `localStorage` under `fs.q.*` and shows them while it refetches. After signing in as someone else, clear local storage (the snippet does) and navigate; a reload straight after the fetch can still show the previous account.
- **Claiming is once per account per 30 days** and one guild per account. The limit rows go with the account, so `seed.sh --remove` then `seed.sh` is the way to claim again.
- **An unverified officer reads as a member** on the claim page and in `/v1/me`'s guild rank, because the derived membership counts verified characters only. That is the real rule (an officer starts a claim only once verified), and it is why the guild master is seeded as Battle.net-verified.
- **Mail is not wired.** The Resend key is a placeholder; the email sign-in link goes nowhere. Use the dev sign-in.
- **The report's own page** (`/reports/<id>`) has no stored events behind it; judge the guild tabs, not that page.
- **Playwright.** `web/tests/e2e/guild-*.spec.ts` and `logs-guild-tabs.spec.ts` run unchanged against the fixture build and do not touch this stack.
