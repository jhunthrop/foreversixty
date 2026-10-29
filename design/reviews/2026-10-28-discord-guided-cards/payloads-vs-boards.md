# Real payloads against the boards (Task 15, offline half)

Source: `api/internal/discordbot/payload_dump_test.go` (`TestDumpGuidedPayloads`, skipped unless `GUIDED_PAYLOADS_DIR` is set) writes what the handlers and the drainer answer to `design/mocks/2026-10-28-discord-guided-cards/payloads/<frame>.json`. `gen_guided_cards.mjs` draws each as `<frame>-real-{desktop,phone}.png` next to the board `<frame>-{desktop,phone}.png`. A field-by-field JSON diff (board message against payload, sentinel marks removed) produced the table below.

Differences that are the test world, not the handlers, and are not repeated per frame: base URLs (`site.test`, `api.test` against `foreversixty.gg`), item names and icons, raid ids and `raid:{id}:*` ids, share tokens, option `value`s (the boards use readable stand-ins), the locked select's single `none` option, and Discord wire fields the boards omit (`nonce`, `enforce_nonce`, `message_reference`, `allowed_mentions`, `default:false`). Every `default: true` in a payload matches the option the board shows chosen; every button style, label and row order matched.

## Verdicts

| Frame | Verdict | Differences |
|---|---|---|
| bis-card-1-class | match | none beyond the test world |
| bis-card-2-spec | match | none |
| bis-card-3-range | match | none |
| bis-card-4-ready | match | none (same payload as card 3; the boards differ only in which list is drawn open) |
| bis-card-nudge | match | none |
| bis-card-main | match | none |
| bis-card-main-changed | match | `Use Zulmara` present, picks and lead identical |
| bis-card-stale | defect | content is `This card is out of date. Run //bis again.` (double slash); the board and spec §G.7 say `Run /bis again.` See D1 |
| bis-result | cosmetic | the payload closes with the footer `Beast Mastery Hunter · level 60 Phase 1 raid · Forever Sixty` and a timestamp on its last embed; the board has neither (it reused the old `bis-top` payload). The payload is right, the board is stale. One embed instead of five is the fixture (one BiS slot). Row is Share in channel, Change picks, link, as drawn |
| post-scheduled | cosmetic | `Signed up` is not inline in the payload (inline on the board). Title, description, colour, footer and the five buttons match. The board draws no roster here, so nothing else is missing |
| post-partial | defect (gated) | not dumped: it differs from post-scheduled only by the roster block, which is not built (R1) |
| post-tomorrow | defect (gated) | title `Tomorrow: Onyxia's Lair · Thursday 10 December`, description, colour and buttons match. `Signed up` reads `**33 of 40** · 33 DPS`; the board has the Tanks, Healers, Melee, Ranged, Late, Bench and Out fields with crest emoji. Roster block not built (R1) |
| post-full | defect (gated) | not dumped, roster only (R1) |
| post-reminder | match | reply text, `message_reference`, and the four sign-up buttons plus `Raid page` match. The post above it carries the roster difference of post-tomorrow (R1) |
| post-reminder-webhook | defect (gated) | the content line and the masked link match the board. The webhook sends content only; the board also shows the Tomorrow embed with the roster. Whether a webhook reminder carries the embed is a spec question to settle with the roster (R1) |
| spec-ask-card | defect (gated) | the `Which spec tonight?` card is not built: no `gspc` route exists in `api/internal/discordbot` (R1) |
| spec-ask-done | defect (gated) | same |

Counts: 17 boards. match 8, cosmetic 2, defect 7 (1 real handler defect, 6 gated on the roster work).

Frames with no board, drawn as `<frame>-real`: char-card-unlinked, char-card-own, char-card-unknown, signup-card, signup-card-in, attendance-select, attendance-pick, readiness-select, readiness-pick, report-select, report-pick. Their modals (`char-card-modal`, `attendance-modal`) are JSON only. Each was read against spec §G.4 and §G.7: labels, placeholders, styles (Success only on the forward action), `default` choices, row order and the Share in channel, Someone else, Character page order all match. Two fixture notes, not defects: the `/signup` option descriptions drop `{Guild time} · ` because the test guild has no zone (Task 6 report), and the `/signup` default character is Duskstrider because the test characters carry no main flag.

## Defects in the handlers

- D1 (real). Every stale card reads `Run //bis again.` The copy `CardStale` is `This card is out of date. Run /{command} again.` (`api/internal/discordbot/copy_guide.go:112`), `staleCard` substitutes the command (`api/internal/discordbot/card.go:59`), and every caller passes the command with its slash already on: `bis_press.go:68`, `char_card.go:241`, `officer_cards.go:88`, `officer_scope_cards.go:50`, `officer_scope_cards.go:133`, `signup_card.go:290`. The test at `bis_press_test.go:185` builds its expectation with the same function, so it cannot see it. Fix: pass `bis` not `/bis` at the six call sites (or drop the slash from the copy), then assert the literal `Run /bis again.`

## Gated differences (R1)

R1. The roster block of spec §P.9 and the `Which spec tonight?` card of §P.10 (amendment b4245917c) are the Task 10b build, held at the owner's gate (progress log). Until it lands, the post, the Tomorrow edit, the partial and full posts, the webhook reminder embed and both spec-ask frames differ from the boards exactly as listed above. Nothing else on those posts differs.
