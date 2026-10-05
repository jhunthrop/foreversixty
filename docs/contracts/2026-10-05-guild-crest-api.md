# Guild crest (profile image): API contract

A guild may carry an uploaded crest. It appears wherever the guild's identity mark appears,
starting with the guild page header ring (design/specs/2026-10-04-guild-page.md §12.2); the
faction logo is the default whenever no crest is set. Officers and the leader of a claimed
guild manage it; a moderator may remove it.

## Storage

- Column `guilds.crest_key text null`, `guilds.crest_updated_at timestamptz null` (migration 0032).
- Object in R2 under `guilds/<guild_id>/crest/<sha256-prefix>.webp`, 256x256, WebP quality 85,
  alpha kept, produced server-side from the upload. The original is never stored.
- Served at `GET /v1/guilds/{id}/crest.webp` with `Cache-Control: public, max-age=31536000,
  immutable` and an `ETag`; the URL the API hands out carries the key's hash as `?v=` so a new
  upload is a new URL. 404 when no crest is set.

## Responses

`guild` objects on the home response and the public guild page gain:
```json
"crest_url": "https://api.foreversixty.gg/v1/guilds/2/crest.webp?v=3f9a1c" | null
```

## Endpoints (officer or leader of a claimed guild, CSRF and the usual rate limits)

- `PUT /v1/guilds/{id}/crest` multipart field `image`: PNG, JPEG or WebP, at most 2 MiB,
  at least 128x128, any aspect (the server centre-crops to square, then resizes to 256).
  Returns `{ "crest_url": "..." }`. Errors with code `invalid` and a plain message for
  type, size and dimensions ("That file is 3.1 MB; the limit is 2 MB.").
- `DELETE /v1/guilds/{id}/crest` clears the crest (and deletes the object). 204.

## Web

- Settings tab, officer: a "Guild crest" block: current crest or the faction logo with the
  caption "Default: your faction's logo", a file picker (accept image/png,image/jpeg,image/webp),
  a 64px ring preview of the chosen file before saving, Save and Remove buttons with busy and
  error states, the size and type rules stated in one line under the picker.
- Header ring: crest when `crest_url` is set, faction logo otherwise, neutral disc when neither.
  The ring colour stays the faction colour. Same at 44px on phone.
- Every place that later shows a guild mark reads the same `crest_url ?? faction logo` rule from
  one helper (`guildMarkSrc`), never two copies.
