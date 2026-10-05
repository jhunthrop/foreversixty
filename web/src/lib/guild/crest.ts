// web/src/lib/guild/crest.ts
// Client-side rules for the Settings tab's crest picker (docs/contracts/2026-10-05-guild-
// crest-api.md): type and size, checked before a request is ever made, in the same plain
// wording the API itself answers with for the same two failures -- so a client-caught
// error and a server-caught one read identically. Dimensions (at least 128x128) are left
// to the server: a browser would need to decode the image just to ask it its own size,
// and the contract's only client-side pre-check is "type and 2 MiB".
export const CREST_MAX_BYTES = 2 * 1024 * 1024;
export const CREST_ACCEPTED_TYPES = ['image/png', 'image/jpeg', 'image/webp'] as const;

/** The picker's own `accept` attribute, one place so the input element and this module's
 *  own check can never drift apart. */
export const CREST_ACCEPT_ATTR = CREST_ACCEPTED_TYPES.join(',');

export const CREST_RULES_LINE = 'PNG, JPEG or WebP, up to 2 MB, at least 128 by 128; we crop to a square.';

/**
 * `null` when the file passes both checks; otherwise the plain message to show in place,
 * matching the contract's own server-side wording for the size case ("That file is 3.1 MB;
 * the limit is 2 MB.") exactly, so a client-caught oversize file and a server-caught one
 * never read differently.
 */
export function precheckCrestFile(file: File): string | null {
  if (!CREST_ACCEPTED_TYPES.includes(file.type as (typeof CREST_ACCEPTED_TYPES)[number])) {
    return 'That file needs to be a PNG, JPEG or WebP image.';
  }
  if (file.size > CREST_MAX_BYTES) {
    const sizeMb = (file.size / (1024 * 1024)).toFixed(1);
    return `That file is ${sizeMb} MB; the limit is 2 MB.`;
  }
  return null;
}
