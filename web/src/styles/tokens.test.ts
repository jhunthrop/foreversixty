// The rarity colours are the WoW standard, and three of the six were never chosen to be
// read as small text on a dark panel. This holds every rarity token that reaches text
// through levelColorClass() or rarityClassFor() to WCAG AA for normal text (4.5:1) against
// --color-raised, the panel those labels sit on. The ratios are computed here and the hexes
// are read from tokens.css, so the test cannot drift from the tokens it guards.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { levelColorClass } from '../lib/levels';
import { rarityClassFor } from '../lib/planner/items';

const TOKENS_FILE = path.join(path.dirname(fileURLToPath(import.meta.url)), 'tokens.css');

/** WCAG 2.1 minimum contrast for text below 18.66px bold / 24px regular. */
const AA_NORMAL_TEXT = 4.5;

/** Every `--color-*: #rrggbb;` declaration in tokens.css, keyed without the `--color-` prefix. */
function readColorTokens(): Map<string, string> {
  const source = readFileSync(TOKENS_FILE, 'utf8');
  const tokens = new Map<string, string>();
  for (const [, name, hex] of source.matchAll(/--color-([a-z0-9-]+):\s*(#[0-9a-f]{6})\s*;/gi)) {
    tokens.set(name, hex.toLowerCase());
  }
  return tokens;
}

/** sRGB channel to linear light, per WCAG 2.1 relative luminance. */
function channelLuminance(byte: number): number {
  const srgb = byte / 255;
  return srgb <= 0.03928 ? srgb / 12.92 : ((srgb + 0.055) / 1.055) ** 2.4;
}

function relativeLuminance(hex: string): number {
  const value = Number.parseInt(hex.slice(1), 16);
  const red = channelLuminance((value >> 16) & 0xff);
  const green = channelLuminance((value >> 8) & 0xff);
  const blue = channelLuminance(value & 0xff);
  return 0.2126 * red + 0.7152 * green + 0.0722 * blue;
}

function contrastRatio(foreground: string, background: string): number {
  const [lighter, darker] = [relativeLuminance(foreground), relativeLuminance(background)].sort(
    (a, b) => b - a,
  );
  return (lighter + 0.05) / (darker + 0.05);
}

/**
 * The rarity classes the app actually puts on text, taken from the two functions that emit
 * them rather than from a list here, so a new call site or a remapped quality is covered.
 */
function rarityTextClasses(): string[] {
  const qualities = [0, 1, 2, 3, 4, 5].map((quality) => rarityClassFor(quality));
  const brackets = [10, 30, 60].map((min) => levelColorClass(min));
  return [...new Set([...qualities, ...brackets])].filter((name) => name.startsWith('text-rarity-'));
}

describe('contrastRatio', () => {
  it('matches the WCAG reference values for the extremes', () => {
    expect(contrastRatio('#ffffff', '#000000')).toBeCloseTo(21, 5);
    expect(contrastRatio('#0d111a', '#0d111a')).toBeCloseTo(1, 5);
  });
});

/** One sRGB channel moved by `delta` (0 to 255), as hex. */
function shift(hex: string, delta: number): string {
  const value = Number.parseInt(hex.slice(1), 16);
  const channels = [16, 8, 0].map((bit) => Math.min(255, Math.max(0, ((value >> bit) & 0xff) + delta)));
  return `#${channels.map((c) => c.toString(16).padStart(2, '0')).join('')}`;
}

/** The painted ground's 200 px windows sit within 6.1 of its mean (art.ground_tile, spec round 5);
 *  7 covers every window either way. */
const GROUND_WINDOW = 7;

describe('round 5 ink (design/specs/2026-11-09-directions.md)', () => {
  const tokens = readColorTokens();
  const card = tokens.get('raised') as string;
  const ground = tokens.get('ground') as string;

  it('reads the card and the earth ground out of tokens.css', () => {
    expect(card).toBe('#170e09');
    expect(ground).toBe('#67563a');
  });

  it.each(['text', 'muted-card', 'gold', 'cream', 'strong'])('%s clears AA on the card', (name) => {
    expect(contrastRatio(tokens.get(name) as string, card)).toBeGreaterThanOrEqual(AA_NORMAL_TEXT);
  });

  it.each(['text', 'muted-ground', 'gold', 'cream'])(
    '%s clears AA on the ground at its darkest and lightest window',
    (name) => {
      for (const window of [shift(ground, -GROUND_WINDOW), shift(ground, GROUND_WINDOW)]) {
        expect(contrastRatio(tokens.get(name) as string, window), window).toBeGreaterThanOrEqual(
          AA_NORMAL_TEXT,
        );
      }
    },
  );

  it('cream clears AA on the blue band', () => {
    expect(contrastRatio(tokens.get('cream') as string, tokens.get('band') as string)).toBeGreaterThanOrEqual(
      AA_NORMAL_TEXT,
    );
  });

  it('the client tooltip yellow clears AA on the tooltip ground', () => {
    const yellow = tokens.get('tooltip-yellow') as string;
    expect(contrastRatio(yellow, tokens.get('tooltip') as string)).toBeGreaterThanOrEqual(AA_NORMAL_TEXT);
  });
});

describe('rarity text tokens', () => {
  const tokens = readColorTokens();
  const background = tokens.get('raised');

  it('reads --color-raised out of tokens.css', () => {
    expect(background).toBe('#170e09');
  });

  it('covers all six rarities plus the level brackets', () => {
    expect(rarityTextClasses()).toHaveLength(6);
  });

  it.each(rarityTextClasses())('%s clears AA on --color-raised', (className) => {
    const hex = tokens.get(className.replace(/^text-/, ''));
    expect(hex, `${className} has no matching token in tokens.css`).toBeDefined();
    expect(contrastRatio(hex as string, background as string)).toBeGreaterThanOrEqual(AA_NORMAL_TEXT);
  });
});

describe('the P15 sweep (plan 2026-11-09-redesign, Wave 2)', () => {
  const LEGACY = [
    'color-card-top',
    'color-line-warm',
    'color-line-warm-strong',
    'color-line-card',
    'color-gold-deep',
    'color-gold-light',
    'color-ember',
    'row-hover',
    'color-row-hover',
    'row-band-hover',
    'tile',
    'color-tile',
    'tile-ink',
    'color-tile-ink',
    'radius-panel',
    'radius-control',
    'radius-pill',
  ];

  it('leaves no Direction B legacy token in tokens.css', () => {
    const source = readFileSync(TOKENS_FILE, 'utf8');
    for (const name of LEGACY) expect(source, name).not.toMatch(new RegExp(`--${name}:`));
  });

  it('keeps the warning ink as its own token, AA as text on the card', () => {
    const tokens = readColorTokens();
    expect(tokens.get('warn')).toBe('#d66e28');
    expect(
      contrastRatio(tokens.get('warn') as string, tokens.get('raised') as string),
    ).toBeGreaterThanOrEqual(AA_NORMAL_TEXT);
  });
});
