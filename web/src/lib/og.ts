// web/src/lib/og.ts
// The link preview (round 5, design/specs/2026-11-09-directions.md "Link preview"): the owner's
// road-and-glow art under the night veil, our painted seal and "Forever Sixty" in Marcellus, the
// page's title, its kicker, and the fan-made line, so a shared link can never read as an official
// promo. Satori cannot read CSS custom properties, so the few colours it paints are the tokens'
// values, named after them (design/tokens/tokens.json, color.forever and color.base).
import { readFile } from 'node:fs/promises';
import { createRequire } from 'node:module';
import path from 'node:path';
import satori from 'satori';
import { Resvg } from '@resvg/resvg-js';
import { FAVICON, OG_ART } from './ui/assets';

const require = createRequire(import.meta.url);
const font = (file: string) => readFile(require.resolve(file));
const marcellus = font('@fontsource/marcellus/files/marcellus-latin-400-normal.woff');
const openSans = font('@fontsource/open-sans/files/open-sans-latin-400-normal.woff');

/** The public file the site serves at `url` (/art/og-art.jpg), as a data URI. */
async function publicDataUri(url: string, type: string): Promise<string> {
  const bytes = await readFile(path.resolve(process.cwd(), 'public', url.replace(/^\//, '')));
  return `data:${type};base64,${bytes.toString('base64')}`;
}

const art = publicDataUri(OG_ART, 'image/jpeg');
/** The painted seal at 180 px (the apple-touch icon is the painted seal). */
const seal = publicDataUri(FAVICON.appleTouch, 'image/png');

/** The token values the card paints (satori reads no CSS). */
const INK = {
  cream: '#f3eee2', // --color-cream
  cream2: '#fcedcf', // --color-cream-2
  text: '#dcd8d2', // --color-text
  veil: 'rgba(7, 29, 46, 0.42)', // the preview's night veil (art_crops.OG_VEIL)
} as const;

/** Copy on every card: the fan-made line (owner rule: never mistakable for an official page). */
export const OG_FAN_LINE = 'Fan-made player tool for WoW: Forever';

export interface OgContent {
  title: string;
  kicker: string;
}

const text = (fontSize: number, color: string, children: string, extra: Record<string, unknown> = {}) => ({
  type: 'div',
  props: { style: { fontSize, color, ...extra }, children },
});

export async function renderOg({ title, kicker }: OgContent): Promise<Uint8Array<ArrayBuffer>> {
  const svg = await satori(
    {
      type: 'div',
      props: {
        style: {
          width: 1200,
          height: 630,
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'space-between',
          padding: 64,
          backgroundImage: `linear-gradient(${INK.veil}, ${INK.veil}), url(${await art})`,
          backgroundSize: '1200px 630px',
          fontFamily: 'Open Sans',
        },
        children: [
          {
            type: 'div',
            props: {
              style: { display: 'flex', alignItems: 'center', gap: 22 },
              children: [
                { type: 'img', props: { src: await seal, width: 88, height: 88 } },
                text(40, INK.cream2, 'Forever Sixty', { fontFamily: 'Marcellus' }),
              ],
            },
          },
          {
            type: 'div',
            props: {
              style: { display: 'flex', flexDirection: 'column' },
              children: [
                text(76, INK.cream, title, { fontFamily: 'Marcellus', lineHeight: 1.05, letterSpacing: -1 }),
                // No eyebrow over the title: the page's kind rides the fan-made line (ai-slop, round 5).
                text(24, INK.text, `${kicker} · ${OG_FAN_LINE}`, { marginTop: 22 }),
              ],
            },
          },
        ],
      },
    },
    {
      width: 1200,
      height: 630,
      fonts: [
        { name: 'Marcellus', data: await marcellus, weight: 400, style: 'normal' },
        { name: 'Open Sans', data: await openSans, weight: 400, style: 'normal' },
      ],
    },
  );
  const png = new Resvg(svg, { fitTo: { mode: 'width', value: 1200 } }).render().asPng();
  // Buffer is typed as Uint8Array<ArrayBufferLike>, which the DOM lib's BodyInit no
  // longer accepts (it wants Uint8Array<ArrayBuffer>). Uint8Array.from copies into a
  // fresh ArrayBuffer-backed array so `new Response(png)` type-checks at the call site.
  return Uint8Array.from(png);
}
