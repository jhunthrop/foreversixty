import { describe, expect, it } from 'vitest';
import { FILLED_GLYPHS, GLYPHS, type GlyphName } from './glyphs';

const CARD_HEAD_GLYPHS: readonly GlyphName[] = [
  'scale',
  'bolt',
  'list',
  'pin',
  'sword',
  'tree',
  'book',
  'upload',
  'download',
  'share',
  'layers',
  'target',
  'trophy',
  'search',
  'news',
];

describe('glyphs', () => {
  it.each(CARD_HEAD_GLYPHS)('carries the card-head glyph %s (site alignment S4)', (name) => {
    expect(GLYPHS[name]).toMatch(/^<(path|circle|rect) /);
  });

  it('draws every glyph from SVG primitives only, never a script or a foreign element', () => {
    for (const markup of Object.values(GLYPHS)) {
      const tags = [...markup.matchAll(/<(\w+)/g)].map((m) => m[1]);
      expect(tags.every((tag) => ['path', 'circle', 'rect'].includes(tag))).toBe(true);
    }
  });

  it('marks only solid shapes as filled', () => {
    for (const name of FILLED_GLYPHS) expect(GLYPHS[name]).toContain('stroke="none"');
  });
});
