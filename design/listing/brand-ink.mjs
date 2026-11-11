// design/listing/brand-ink.mjs
// The token colours the listing plates paint, read from design/tokens/tokens.json (the plates are
// pages with no site CSS, so they cannot read the CSS variables).
import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';

const tokens = JSON.parse(readFileSync(fileURLToPath(new URL('../tokens/tokens.json', import.meta.url)), 'utf8'));
const { base, forever } = tokens.color;

export const INK = Object.freeze({
  bg: base.bg.$value,
  text: base.text.$value,
  line: base.line.$value,
  cream: forever.cream.$value,
  cream2: forever.cream2.$value,
  /** Text set straight on the painted ground. */
  mutedGround: forever.mutedGround.$value,
});
