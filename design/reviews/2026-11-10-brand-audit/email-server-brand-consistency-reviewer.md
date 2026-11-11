# brand-consistency-reviewer: email, server pages and push copy (2026-11-10)

This review read the code and the captures (`captures/email-inbox.png`, `2026-11-10-h-api-html/after`, `2026-11-10-unify-server/after`). It did not render any pages.

## Email

### BLOCKER

**E1. Every mail is plain text, with no seal, wordmark or footer.**
- `api/internal/mail/resend.go:33` sends only `text`. The bodies are in `billingmail.go` (`compose`) and `auth/handler.go:268-272`.
- Money mail (renewal, seat, comeback) reads as a bare system notice.
- There is no legal footer (Commish LLC, help@) and no button, only raw URLs.

**E2. The sender and the contact address disagree.**
- `api/internal/config/config.go:16` sends from `hello@foreversixty.gg`.
- The legal pages and `shared.account.delete.email` say `help@foreversixty.gg` (20 uses, against 5 for `hello@`, all in api config, README and tests).
- Fix: use one address, or set `Reply-To: help@`. Update `api/README.md:33` and `:824`.

### SHOULD

**E3. The mail sentences are hand-built in Go, not in `apps/shared/copy`.**
- Mail: "renews on 2 December 2026 for $150, your founder price; …" (`billingmail.go:89`).
- Account page and apps: "Premium · renews {date}" and "{date} · {price}" (`en.json:783`, `:789`).
- `EmailDate` is `textfmt.LongDate` in Pacific time (`billingmail/price.go:60`). That long date form has no `formats.json` case.
- Fix: move the bodies to `mail.*` keys, and add a `long_date_year` case to `formats.json`.

**E4. Each body is a single 40-word sentence.**
- Lead with a figure block (plan, date, price), the way the account page does.

**E5. Mail prices come from `billingmail.FormatPrice`, not the web's price rules.**
- Pin both with one fixture.

### NICE

**E6. No brand prefix in subjects, and no preheader.**

**E7. The sign-in mail is behind `EMAIL_SIGN_IN`, which is off by default.**
- It still needs the same shell when it is turned on.

### Minimal branded mail

1. Add `HTML` to `mail.Message`. Resend already accepts an `html` field.
2. Write one `api/internal/mail/render.go` (or `billingmail/render.go`) that reads `api/internal/brand/tokens_gen.go`:
   - ground `ColorForeverGround`, panel `ColorBaseRaised`, text `ColorBaseText`, button `ColorForeverTeal`, link `ColorBaseGold`;
   - the `FontTitle` wordmark with a serif fallback.
3. Lay it out as a 560 px column with tables and inline CSS.
4. Seal: the hosted flat seal PNG at 40 px. Do not use an SVG data URL, because Gmail drops it.
5. Body: the sentence, one teal button with the existing label, and the raw URL below it.
6. Footer: the web footer line ("Fan-made player tool…"), Commish LLC, and help@.
7. All six mails and the sign-in mail use the one shell.

## Server pages

### SHOULD

**S1. The `app-confirm` captures are stale.**
- `h-api-html/after/app-confirm-1440.png` still shows the Battle.net sentence that `app_confirm.go:48-51` dropped.
- The capture shows no seal, although `:45` renders one at 40 px.
- Fix: recapture, and confirm the data-URL seal paints.

**S2. The flat seal (`brand.NavSealSVG`) is correct at 40 px.**
- Confirm it reads as the same ring as the web header's seal.

### NICE

**S3. Untokened shadow.**
- `app_confirm.go:145` hardcodes `text-shadow rgba(0,0,0,.45)`.
- Route it through a token if it matches the web band title; otherwise drop it.

**S4. Panel edge.**
- `:147` uses `#000`. Use a token.

**S5. Band mask.**
- `:141` uses `#000` and `rgba(0,0,0,.5)` for the mask alpha. Low risk.

**S6. Confirm copy is a local struct.**
- `:50` `appConfirmCopy` holds "Continue to Battle.net" and "Sign in to the app". Move them to shared keys.

**S7. Column width is hardcoded in two places.**
- `site/page.go:130-137` (/b) is token- and Tailwind-driven, but sets `max-width:560px`.
- `app_confirm.go` also has `columnWidth` 560.
- Use one token.

**S8. The /b mock is the stale side.**
- In `side-b-unavailable-1440.png` the mock nav shows icons the web lacks, so the build matches the web and the mock does not.

The API serves no other `text/html` handlers.

## Push copy

### SHOULD

**P1. Push strings live in a Go struct.**
- `notify/events.go:79-90` (`pushCopy`) holds "Your %s sim finished" and "%s DPS".
- These duplicate `home.next.simFinished` and `web/src/lib/sim/server-runs-copy.ts:40`.

**P2. The same words appear again in `notify/embed.go:75-86`.**

### NICE

**P3. Separators match the web.**
- The "·" separators and "Sign-ups are open: Wed 4 Nov, 8 pm CST." follow `formats.json`.

## Shared-source consumption

| Surface | Tokens | Seal | Copy | Formats |
|---|---|---|---|---|
| Email | none | none | Go literals | `textfmt.LongDate` (no `formats.json` case) |
| app-confirm | `brand` tokens, plus some `#000`/`rgba` | flat seal | local struct | n/a |
| /b | web CSS variables and Tailwind | web chrome | web | web |
| Push | n/a | n/a | Go struct, written twice | `textfmt` |

VERDICT: DRIFT
