# Page frame widths, 2026-11-10

Owner: "why don't we have consistent width for every experience?" Measured by
`web/tests/e2e/frame-capture.spec.ts` (support/frame-pages.ts, one page of every type, each loaded
directly). Each cell gives the header bar's width and left edge, then the content box of every
`.page-column` on the page: `body` is the page body, `footer` the footer, `ui-band` the page head
band, `ui-cbar` the control bar and `site-top` the frame around the bar. The raw numbers are in
`before/frame.json` and `after/frame.json`.

## Summary

- Before: 29 page types used the wide frame from 1760 px (1632 px at 1920, 1664 at 2000). Five
  stayed on the 1248 px column at every width because they never opted in with Base's `wide` prop:
  character, guild (every tab), the shared build `/b` (rendered by the API, so no `data-column`),
  `/b-unavailable` and `/app/auth` (not captured; same Base call). On those pages the bar was also
  1248 px wide, below the 1600 px container query that shows the door icons, and a body rule hid
  the "/" hint and the selector's spec line. That is the owner's report: the guild page's bar was
  narrower and had no icons, hint or spec.
- After: one frame on every page. At 1440 px it is 1248 px at x 96, at 1920 px 1632 px at x 144,
  and at 2000 px 1664 px at x 168. The bar, head band, control bar, body and footer share those
  edges on all 34 page types (`web/tests/e2e/page-frame.spec.ts` holds this at 1920).
- Inside the frame the named measures stay: prose at `--measure-reading` (guides, legal,
  changelog, about), the 560 px `form` layout (/connect) and the paperdoll at 640 or 760 px.

## Before

| Page | 1440 px | 1920 px | 2000 px |
|---|---|---|---|
| 404 | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168 |
| home | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168 |
| bis-landing | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| bis | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-cbar 1664 @168 |
| guides | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168 |
| guides-class | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168 |
| guides-spec | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| tiers | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| sim | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168 |
| sim-result | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| sim-gear | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| planner | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| rankings | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| logs | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168 |
| reports | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168 |
| report | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| character | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1248 @336; body 1248 @336; footer 1248 @336; ui-band 1248 @336 | bar 1248 @376; body 1248 @376; footer 1248 @376; ui-band 1248 @376 |
| guild | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1248 @336; body 1248 @336; footer 1248 @336; ui-band 1248 @336 | bar 1248 @376; body 1248 @376; footer 1248 @376; ui-band 1248 @376 |
| guild-roster | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1248 @336; body 1248 @336; footer 1248 @336; ui-band 1248 @336 | bar 1248 @376; body 1248 @376; footer 1248 @376; ui-band 1248 @376 |
| guild-loot | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1248 @336; body 1248 @336; footer 1248 @336; ui-band 1248 @336; ui-cbar 1248 @336 | bar 1248 @376; body 1248 @376; footer 1248 @376; ui-band 1248 @376; ui-cbar 1248 @376 |
| account | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168 |
| setup | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168 |
| connect | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168 |
| login | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168 |
| premium | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| characters | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168 |
| ratings | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168 |
| creators | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168 |
| app | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168 |
| changelog | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| legal | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168 |
| about | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; ui-band 1664 @168 |
| b | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1248 @336; body 1248 @336; footer 1248 @336; ui-band 1248 @336 | bar 1248 @376; body 1248 @376; footer 1248 @376; ui-band 1248 @376 |
| b-unavailable | bar 1248 @96; body 1248 @96; footer 1248 @96; ui-band 1248 @96 | bar 1248 @336; body 1248 @336; footer 1248 @336; ui-band 1248 @336 | bar 1248 @376; body 1248 @376; footer 1248 @376; ui-band 1248 @376 |

## After

| Page | 1440 px | 1920 px | 2000 px |
|---|---|---|---|
| 404 | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| home | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| bis-landing | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| bis | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-cbar 1664 @168 |
| guides | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| guides-class | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| guides-spec | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| tiers | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| sim | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| sim-result | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| sim-gear | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| planner | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| rankings | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| logs | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| reports | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| report | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| character | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| guild | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| guild-roster | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| guild-loot | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| account | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| setup | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| connect | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| login | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| premium | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| characters | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| ratings | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| creators | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| app | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| changelog | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96; ui-cbar 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144; ui-cbar 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168; ui-cbar 1664 @168 |
| legal | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| about | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| b | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
| b-unavailable | bar 1248 @96; body 1248 @96; footer 1248 @96; site-top 1248 @96; ui-band 1248 @96 | bar 1632 @144; body 1632 @144; footer 1632 @144; site-top 1632 @144; ui-band 1632 @144 | bar 1664 @168; body 1664 @168; footer 1664 @168; site-top 1664 @168; ui-band 1664 @168 |
