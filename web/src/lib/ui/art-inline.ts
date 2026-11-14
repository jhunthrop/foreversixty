// web/src/lib/ui/art-inline.ts
// A page's own band art inside the page (2026-11-13, band LCP). The band is the largest paint on
// most pages, and as a CSS background its file was found only after the stylesheet and then
// queued behind the document and the faces: real visitors saw it seconds after the title. Base.astro
// inlines the page's phone file (lib/ui/assets.ts artCssVars) so it paints with the first frame; the
// desktop file is drawn over it (ui/band.css), so a wide screen has the art at once and the sharper
// copy when it lands. Server only: it reads web/public at build time (the build and the tests run
// from web/).

import { readFileSync } from 'node:fs';
import path from 'node:path';

const PUBLIC_DIR = path.join(process.cwd(), 'public');
const MIME: Readonly<Record<string, string>> = { '.webp': 'image/webp' };
const cache = new Map<string, string>();

function read(file: string): string {
  const mime = MIME[path.extname(file)];
  if (mime === undefined) throw new Error(`art-inline: no data URL type for ${file}`);
  return `data:${mime};base64,${readFileSync(path.join(PUBLIC_DIR, file)).toString('base64')}`;
}

/** A /public art file (an assets.ts path) as a data URL, read once per build. */
export function inlineArt(file: string): string {
  const known = cache.get(file);
  if (known !== undefined) return known;
  const dataUrl = read(file);
  cache.set(file, dataUrl);
  return dataUrl;
}
