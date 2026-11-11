# brand-consistency-reviewer: Discord bot (2026-11-10)

Scope: the reviewer read the desktop captures in `2026-11-09-f-discord/discord`, `2026-11-10-unify-server/after/discord-*` and `captures/identity-sheet.png`, and grepped `api/internal/discordbot`, `card` and `sheet`. It did not open the phone captures or the payload JSON.

## DRIFT

**BLOCKER:** none. Embed and card colours all come from `brand.Color*` (`tokens_gen.go`), and the bot copy has no em dashes.

**SHOULD**
1. **Two competing Discord identities** (`captures/identity-sheet.png`).
   - Server icon: `design/community/discord/assets/icon.png` is the black-disc LX (the iOS AppIcon family).
   - Bot avatar: the painted seal.
   - Web favicon: the flat teal seal.
   - `assets/banner.png` uses forced tracked caps ("FOREVER SIXTY", "PLAY YOUR CLASS BETTER").
   - `design/assets/art/discord-banner-2026-11-09.jpg` (tavern, painted seal, Marcellus sentence case) matches the web.
   - Fix: make the 11-09 banner the one banner and use the painted seal for the server icon. Retire `banner.png` and the black-disc `icon.png`.
2. **Raid post title date** (`unify-server/after/discord-post-full-desktop.png`): "Onyxia's Lair · Thursday 10 December" via `textfmt.LongDay` (`cmd_report.go:137`). The web raid row reads "Wed 13 Jan". Fix: use `textfmt.RaidDay`, or confirm that the long form is a ruling for titles.
3. **"When" and "Guild time" are two fields.** The web has one cell: time with zone, plus the relative time. Keep the Discord local-time stamp (platform idiom) and drop the duplicate label.
4. **Item thumbnail has no quality ring** (`char-card-own-desktop.png`, `item_embed.go:60`). Discord cannot border a thumbnail, so bake the ring in through `sheet/palette.go` `qualityRing`.
5. **Character thumbnail is unringed hood art** (`char-card-own-desktop.png`, `bis-card-4-ready-desktop.png`). Serve the ringed HD crest. `design/community/discord/assets/crest-*.png` exists for 9 classes.

**NICE**
- `card/card.go:322` truncates with an ASCII `"..."` where the rest of the product uses "…".
- `cmd_report.go:181`, `:206` and `:208` hand-build plurals ("%d %ss").
- `card.go:331` has its own `parseHexColor` with a cream fallback, a second colour path.
- `copy_guide.go:71` and `:77` and `copy_lookup.go:65` hold the bot's own strings. They match the web today ("Best in slot", "BiS") but are not generated from shared copy.

## Shared-source consumption

| Area | Generated | Hardcoded | Cite |
|---|---|---|---|
| Embed colours | `brand.RGB(brand.Color*)`: accent, kill, class, quality | 0 colour ints | `item_embed.go:36-37`, `:48`, `:75` |
| Card palette | `brand.Color*` | `colorBlack` | `sheet/palette.go:23` |
| Fonts | `brand.Marcellus400` and Open Sans 400/600/700 (the web families) | pixel sizes 26-72, off the type scale | `card.go:110-116` |
| Icons, emoji | app emoji sync, brand/art seal and ground | crest PNGs committed by hand | `emoji_sync.go` |
| Copy | none | all strings in `copy_*.go` | `copy_lookup.go:65` |
| Formats | `textfmt` (`formats.json` replayed in `textfmt_test.go`), `upgrades.SignedFigure` | `%d` counts | `cmd_report.go:206` |

## Components

| Component | Web | Discord | Status |
|---|---|---|---|
| Item line | Quality-bordered icon | "Slot · Quality · item level N" | Same grammar; ring missing (SHOULD 4) |
| Crest | Ringed | Unringed (SHOULD 5) | Drift |
| Raid row | Bar, chip and relative time | Text "40 of 40" with role columns | Role wording matches; allowed adaptation |
| Sign-up | Own controls | In / Late / Bench / Out / "Raid page ↗" as Discord buttons | Allowed |
| Figures | "+9.2 DPS", "item level 61" | Same | Matches |

## Stale captures

- `2026-11-09-f-discord/discord/*` predates `5b8e8a02c` ("one house format for Discord") and `cb89f2a47`. Re-render them.
- `2026-11-10-unify-server/after/discord-*` is current.

VERDICT: DRIFT
