# BA3 brand-consistency verdict

Reference: live web (unify-format/after guild-raids-1440, footer), footer-copy.ts, tokens.css.

## Checked, consistent
- E1 fixed: every letter has the painted seal, Marcellus wordmark, earth ground (#67563a = --color-ground), raised dark panel with warm hairline, teal button, gold link, and the site footer lines (shared.footer.*) plus a contact line. Plain-text parts carry the same blocks in order.
- Colours and font stacks come from brand/tokens_gen.go; no hardcoded hex in notify, discordbot, mail, billingmail, auth, subscribe.
- E2 fixed: help@ sender and reply-to; no hello@ left in api/internal.
- Discord SHOULD 2 fixed: "Onyxia's Lair · Thu 10 Dec" matches the web raid row. SHOULD 3 fixed: one When cell. Kill times m:ss match the web.
- Push: "Your ... report is ready" and "sim finished" from server.json; "report" matches the site term. No em dashes.
- Mail dates "12 Dec 2026" and "$150 a year" match the account plan.
- ALLOWED: Discord timestamp chips vs web "8 pm CST / in 4 days"; push app-icon placeholder is a renderer stand-in.

## BLOCKER
None.

## SHOULD
None.

## NICE
1. Discord "Signed up 25 of 40" vs web "25/40" with bar; share a key if one exists.
2. Raw URL under the button wraps mid-word on phone (captures/mail-seat-removed-phone.png); consider a break rule at / and &.
3. Out of scope, listed: unringed hood crest thumbnail (lane D), server icon/banner (BA2), app_confirm.go #000/text-shadow token.

VERDICT: CONSISTENT
