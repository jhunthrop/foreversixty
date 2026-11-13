// web/src/lib/og-tier.ts
// The tier list's link preview (owner 2026-11-13: a shared /tiers link shows the list itself, not
// the generic page card). The site's dark panel, not the road art: the page's title and role with
// the seal and the name across the top, the page's meta line, then the list as the page draws it, each tier a strip led by its letter
// in its item quality colour, each spec its ringed crest and its full name in its class colour.
// The meta line and the Updated stamp are the page's own. Colours are read from the generated
// tokens (src/styles/tokens.css), since satori reads no CSS.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import sharp from 'sharp';
import type { TierCardGroup, TierCardModel, TierCardSpec } from './tiers/tier-card';
import type { TierLetter } from './tiers/tier-rules';
import { classCrestSrc } from './class-crest';
import { isClassSlug } from './ui/primitives';
import { INK, OG_BRAND, OG_SIZE, box, renderCardPng, sealImage, text } from './og';

const TOKENS_CSS = readFileSync(path.resolve(process.cwd(), 'src/styles/tokens.css'), 'utf8');

/** A generated token's value (`--color-class-mage` -> `#3fc7eb`); a missing one fails the build. */
export function tokenValue(name: string): string {
  const match = new RegExp(`${name}:\\s*([^;]+);`).exec(TOKENS_CSS);
  if (match === null) throw new Error(`og-tier: no token ${name} in src/styles/tokens.css`);
  return match[1]!.trim();
}

/** Each letter's colour, the page's (TIER_LETTER_TONES): S legendary down to D poor. */
const LETTER_TOKEN: Readonly<Record<TierLetter, string>> = {
  S: '--color-rarity-legendary',
  A: '--color-rarity-epic-text',
  B: '--color-rarity-rare-text',
  C: '--color-rarity-uncommon',
  D: '--color-rarity-poor',
};

const PANEL = {
  bg: tokenValue('--color-bg'),
  raised: tokenValue('--color-raised'),
  line: tokenValue('--color-line-soft'),
  muted: tokenValue('--color-muted'),
  text: tokenValue('--color-text'),
} as const;

/** The card's grid: the head (title, role, seal and name) and the meta line over the list. */
const PAD = 44;
const LETTER_WIDTH = 64;
const CELL_GAP = 10;
const CREST_PX = 38;
const NAME_PX = 22;
/** A ranked list (tanks, healers: a handful of rows, no letters) draws one spec a row, larger. */
const RANKED = { crest: 52, name: 30, rankWidth: 64 } as const;
/** Specs per strip line: three keep the longest name ("Beast Mastery Hunter") on one line at a
 *  size a phone preview can read. */
const PER_LINE = 3;
const LIST_WIDTH = OG_SIZE.width - PAD * 2;
/** The list's inner padding beside the cells. */
const CELLS_PAD = 12;
const CELL_WIDTH = Math.floor((LIST_WIDTH - LETTER_WIDTH - 2 * CELLS_PAD) / PER_LINE);

function classColour(slug: string): string {
  return tokenValue(`--color-class-${slug}`);
}

/** The class's HD crest as a PNG data URI: the site serves WebP, which the card's renderer cannot read. */
async function crestUri(slug: string): Promise<string> {
  if (!isClassSlug(slug)) throw new Error(`og-tier: no crest for class ${slug}`);
  const file = path.resolve(process.cwd(), 'public', classCrestSrc(slug).replace(/^\//, ''));
  const png = await sharp(file)
    .resize(CREST_PX * 2, CREST_PX * 2)
    .png()
    .toBuffer();
  return `data:image/png;base64,${png.toString('base64')}`;
}

function crestImage(slug: string, uri: string, px: number): unknown {
  return {
    type: 'img',
    props: {
      src: uri,
      width: px,
      height: px,
      style: { borderRadius: 999, border: `2px solid ${classColour(slug)}`, backgroundColor: PANEL.bg },
    },
  };
}

async function specCell(spec: TierCardSpec): Promise<unknown> {
  return box({ alignItems: 'center', gap: CELL_GAP, width: CELL_WIDTH, height: CREST_PX + 10 }, [
    crestImage(spec.classSlug, await crestUri(spec.classSlug), CREST_PX),
    text(NAME_PX, classColour(spec.classSlug), spec.name, { whiteSpace: 'nowrap' }),
  ]);
}

function rule(last: boolean): string {
  return last ? 'none' : `1px solid ${PANEL.line}`;
}

/** One letter's strip on the damage list: the letter, then its specs four to a line. */
async function letterStrip(
  letter: TierLetter,
  specs: readonly TierCardSpec[],
  last: boolean,
): Promise<unknown> {
  const cells = await Promise.all(specs.map(specCell));
  return box({ alignItems: 'stretch', borderBottom: rule(last), flexGrow: specs.length }, [
    box({ width: LETTER_WIDTH, alignItems: 'center', justifyContent: 'center', flexShrink: 0 }, [
      text(40, tokenValue(LETTER_TOKEN[letter]), letter, { fontFamily: 'Marcellus' }),
    ]),
    box({ flexWrap: 'wrap', alignContent: 'center', padding: `2px ${CELLS_PAD}px`, flexGrow: 1 }, cells),
  ]);
}

/** One spec's row on a ranked list: its rank as the page prints it, its crest, its name and the
 *  figure the list is sorted on. */
async function rankedRow(spec: TierCardSpec, last: boolean): Promise<unknown> {
  return box(
    { alignItems: 'center', gap: 18, borderBottom: rule(last), flexGrow: 1, padding: `0 ${CELLS_PAD}px` },
    [
      text(RANKED.name - 4, PANEL.muted, spec.rank, {
        width: RANKED.rankWidth,
        justifyContent: 'center',
      }),
      crestImage(spec.classSlug, await crestUri(spec.classSlug), RANKED.crest),
      text(RANKED.name, classColour(spec.classSlug), spec.name, { flexGrow: 1 }),
      text(RANKED.name - 4, PANEL.text, spec.figure ?? '', { paddingRight: CELLS_PAD }),
    ],
  );
}

/** The list's rows: a strip per letter, or a row per spec on a list without letters. */
async function listRows(groups: readonly TierCardGroup[]): Promise<unknown[]> {
  const nested = await Promise.all(
    groups.map(async (group, index) => {
      const lastGroup = index === groups.length - 1;
      if (group.letter !== null) return [await letterStrip(group.letter, group.specs, lastGroup)];
      return Promise.all(
        group.specs.map((spec, at) => rankedRow(spec, lastGroup && at === group.specs.length - 1)),
      );
    }),
  );
  return nested.flat();
}

async function head(model: TierCardModel): Promise<unknown> {
  return box({ flexDirection: 'column', gap: 10 }, [
    box({ alignItems: 'center', justifyContent: 'space-between' }, [
      box({ alignItems: 'baseline', gap: 20 }, [
        text(56, INK.cream, model.title, { fontFamily: 'Marcellus', letterSpacing: -1 }),
        text(32, INK.cream2, model.role, { fontFamily: 'Marcellus' }),
      ]),
      box({ alignItems: 'center', gap: 14 }, [
        await sealImage(56),
        text(28, INK.cream2, OG_BRAND, { fontFamily: 'Marcellus' }),
      ]),
    ]),
    text(22, PANEL.text, model.meta),
  ]);
}

export async function renderTierOg(model: TierCardModel): Promise<Uint8Array<ArrayBuffer>> {
  const strips = await listRows(model.groups);
  return renderCardPng({
    type: 'div',
    props: {
      style: {
        ...OG_SIZE,
        display: 'flex',
        flexDirection: 'column',
        gap: 22,
        padding: `36px ${PAD}px ${PAD}px`,
        backgroundColor: PANEL.bg,
        fontFamily: 'Open Sans',
      },
      children: [
        await head(model),
        box(
          {
            flexDirection: 'column',
            flexGrow: 1,
            backgroundColor: PANEL.raised,
            border: `1px solid ${PANEL.line}`,
            borderRadius: 10,
            padding: '4px 0',
          },
          strips,
        ),
      ],
    },
  });
}
