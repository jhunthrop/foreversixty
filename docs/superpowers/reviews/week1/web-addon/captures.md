# L8.3 captures (fs-web-v2 lane/web-v2, fixture build, system Chrome via Playwright)

Desktop = 1280 wide, phone = 360 wide. Roster shots use the guild fixture (viewer Zulmara, member), clock pinned to 2026-12-22 21:46Z.

## /setup (each as `setup-desktop-*` and `setup-phone-*`)
- `setup-<vp>-full.png`: whole page, images forced eager.
- `setup-<vp>-surface-4.1.png` ... `4.11.png`: one per surface of the "What it does in the game" section (4.4 and 4.9 text-only for quest/roll and messaging).
- `setup-<vp>-commands.png`: the Commands list.
- `setup-<vp>-handshake.png`: the companion handshake.
- Horizontal overflow checked at both widths: none.

## Guild Roster tab
- `roster-desktop-snapshot.png`, `roster-phone-snapshot.png`: snapshot line, In game column (Zulmara, Kraggor with level/spec/gap/build/seen; Duskstrider consent-withheld shows "seen 23:12" only; unmatched rows "—").
- `roster-desktop-no-snapshot.png`, `roster-phone-no-snapshot.png`: "No addon snapshot yet ..." line, every row "—".

## Lighthouse (/setup, mobile defaults, 3 runs identical)
performance 96, accessibility 100, best-practices 93, seo 100; LCP 2.6 s, CLS 0, TBT 0.
