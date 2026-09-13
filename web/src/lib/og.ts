import { readFile } from 'node:fs/promises';
import { createRequire } from 'node:module';
import satori from 'satori';
import { Resvg } from '@resvg/resvg-js';

const require = createRequire(import.meta.url);
const cinzel = readFile(require.resolve('@fontsource/cinzel/files/cinzel-latin-700-normal.woff'));
const barlow = readFile(require.resolve('@fontsource/barlow/files/barlow-latin-400-normal.woff'));

export interface OgContent {
  title: string;
  kicker: string;
}

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
          justifyContent: 'flex-end',
          padding: 64,
          background: 'linear-gradient(180deg, #070b12 0%, #0d1522 55%, #1b1410 100%)',
          color: '#f2eee4',
          fontFamily: 'Barlow',
        },
        children: [
          {
            type: 'div',
            props: {
              style: { fontSize: 22, letterSpacing: 3, textTransform: 'uppercase', color: '#e5b955', marginBottom: 18 },
              children: kicker,
            },
          },
          {
            type: 'div',
            props: {
              style: { fontSize: 72, fontFamily: 'Cinzel', fontWeight: 700, lineHeight: 1.05, color: '#e5b955' },
              children: title,
            },
          },
          {
            type: 'div',
            props: {
              style: { fontSize: 24, color: '#9a9484', marginTop: 28 },
              children: 'foreversixty.gg · fan reference, every fact dated and sourced',
            },
          },
        ],
      },
    },
    {
      width: 1200,
      height: 630,
      fonts: [
        { name: 'Cinzel', data: await cinzel, weight: 700, style: 'normal' },
        { name: 'Barlow', data: await barlow, weight: 400, style: 'normal' },
      ],
    }
  );
  const png = new Resvg(svg, { fitTo: { mode: 'width', value: 1200 } }).render().asPng();
  // Buffer is typed as Uint8Array<ArrayBufferLike>, which the DOM lib's BodyInit no
  // longer accepts (it wants Uint8Array<ArrayBuffer>). Uint8Array.from copies into a
  // fresh ArrayBuffer-backed array so `new Response(png)` type-checks at the call site.
  return Uint8Array.from(png);
}
