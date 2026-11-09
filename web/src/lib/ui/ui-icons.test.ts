import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { OUT, renderUiIcons } from '../../../scripts/gen-ui-icons.mjs';
import { UI_ICONS } from './ui-icons';

describe('UI_ICONS', () => {
  it('is current with design/assets/icons/ui (node scripts/gen-ui-icons.mjs)', () => {
    expect(readFileSync(OUT, 'utf8')).toBe(renderUiIcons());
  });

  it('holds the account menu set, each one path on a square grid', () => {
    for (const name of [
      'account',
      'premium',
      'setup',
      'character',
      'paste-export',
      'sign-out',
      'manage',
    ] as const) {
      expect(UI_ICONS[name].size).toBeGreaterThan(0);
      expect(UI_ICONS[name].d).toMatch(/^M/);
    }
  });
});
