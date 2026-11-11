# Lane C (unify-server) UX review, 2026-11-10

**SHIP**

Judged on the before/after captures. Every after reads right against the owner rulings and shows no regression. No must-fix findings.

## Findings

1. **nit** `after/discord-post-full-desktop.png`, `discord-raid-answer-phone.png`: "When" reads "8:00 PM (in 4 hours)" and "Guild time" reads "8 pm CST" side by side. The first is Discord's own `<t:unix:t>` render and cannot be changed; the second is our ruled site clock. Correct per the rulings. No action.
2. **nit** `after/app-confirm-390.png` vs `before/`: the wordmark and title now use Marcellus with the flat seal, and the body and button use Open Sans. This matches the companion and site headers, which is the intent. The apostrophe in "Didn't" is typographic. Line breaks are better ("it up / from there" no longer orphans "up"). Phone and desktop are both fine.
3. **nit** `after/app-confirm-1440.png`: the content column is about 1000 px, centred, with the wordmark left-aligned to the card edge. Fine for a bridge page. The button is visible at once. No action.
4. **should** `after/companion-reports-en-US.png`: the Zul'Gurub row ("not on the site yet") has no Copy link, as before. Not caused by this lane, so no action here.
5. **nit** `after/b-*.jpg`: byte-different from `before/`, but layout, copy and spacing are identical. The difference is font rasterisation (different capture machine). Acceptable as "no visual change", but the pair is not a pixel proof. Recapture both on one machine if a pixel proof is wanted.

## Checked and correct

- Discord, `discord-post-*`, `discord-raid-answer-*`, `discord-webhook-raid-moved-*` at desktop and phone: guild time "8 pm CST" matches the site clock (formats.json: "8 pm", no ":00"). When and Was stay Discord-rendered. Layout unchanged, nothing clipped at 390.
- Companion en-US: "Today 6:40 pm", "Yesterday 4:40 pm", "Sun 18 Oct, 2:40 pm", "Last played 6:40 pm". Month and day order match `weekday_day_month`.
- Companion en-GB: "Today 18:40", "Sun 18 Oct, 14:40", "Last played 18:40". 24-hour, matching the "20:00" ruling.
- Ages are the compact form everywhere: "just now", "1 min ago", "4 min ago", "8 min ago". The old "8 minutes ago" on the Device page (`before/companion-device-connected-en-GB.png`) is gone.
- Grouped numbers and spoken durations appear only on text-only surfaces (feed, push), so I could not check them on screenshots; the formats.json cases cover them.
