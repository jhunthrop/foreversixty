// web/src/lib/chart/theme.test.ts
// @vitest-environment jsdom
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { CHART_ACCENT, CHART_INK, chartFont, inkCanvas, inkCss } from './theme';

const TOKENS = readFileSync(
  path.join(path.dirname(fileURLToPath(import.meta.url)), '../../styles/tokens.css'),
  'utf8',
);

describe('the chart theme', () => {
  it('names only tokens tokens.css defines', () => {
    for (const [role, ink] of Object.entries(CHART_INK)) {
      expect(TOKENS, `${role}: ${ink.token}`).toContain(`${ink.token}:`);
    }
  });

  it('draws the primary series in the accent and gridlines in faint white', () => {
    expect(CHART_ACCENT).toBe('var(--color-accent)');
    expect(inkCss(CHART_INK.grid)).toBe('color-mix(in srgb, var(--color-strong) 8%, transparent)');
  });

  it('resolves an ink against the computed tokens for a canvas', () => {
    const el = document.createElement('div');
    el.style.setProperty('--color-accent', '#0093b2');
    el.style.setProperty('--font-body', "'Open Sans', Arial, sans-serif");
    document.body.append(el);
    expect(inkCanvas(el, CHART_INK.accent)).toBe('#0093b2');
    expect(inkCanvas(el, CHART_INK.selection)).toBe('rgba(0, 147, 178, 0.14)');
    expect(chartFont(el)).toBe("400 12px 'Open Sans', Arial, sans-serif");
    el.remove();
  });
});
