import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { describe, expect, it } from 'vitest';
import {
  MODULE_DIR,
  NATIVE_ANDROID_DIR,
  NATIVE_IOS_DIR,
  renderNativeIcons,
  renderUiIcons,
} from '../../../scripts/gen-ui-icons.mjs';
import { iconAttributes, iconMaskUrl, LINE_STROKE_WIDTH } from './icon';
import { ICON_CHEVRON_DOWN, ICON_SEARCH, UI_ICONS } from './ui-icons';

describe('UI_ICONS', () => {
  it('is current with design/assets/icons/ui (node scripts/gen-ui-icons.mjs)', () => {
    const expected = renderUiIcons();
    for (const [file, text] of expected) expect(readFileSync(file, 'utf8')).toBe(text);
    const modules = readdirSync(MODULE_DIR).map((name) => path.join(MODULE_DIR, name));
    expect(modules.filter((file) => !expected.has(file))).toEqual([]);
  });

  it("keeps the apps' tab icons and frame corner current (node scripts/gen-ui-icons.mjs)", () => {
    const expected = renderNativeIcons();
    for (const [file, text] of expected) expect(readFileSync(file, 'utf8')).toBe(text);
    const sets = readdirSync(NATIVE_IOS_DIR).filter((name) => name.endsWith('.imageset'));
    expect(sets.sort()).toEqual(
      ['frame-corner', 'tab-activity', 'tab-gear', 'tab-guild', 'tab-home'].map((n) => `${n}.imageset`),
    );
    const vectors = readdirSync(NATIVE_ANDROID_DIR).filter((name) => /^ic_(tab_|frame_)/.test(name));
    expect(
      vectors.map((name) => path.join(NATIVE_ANDROID_DIR, name)).every((file) => expected.has(file)),
    ).toBe(true);
  });

  it('holds the nav doors and the account menu in the custom set, each one even-odd path', () => {
    for (const name of [
      'bis',
      'guides',
      'tier-list',
      'sim',
      'planner',
      'logs',
      'guild',
      'premium',
      'account',
      'setup',
      'character',
      'paste-export',
      'sign-out',
      'manage',
      'search',
    ] as const) {
      expect(UI_ICONS[name].kind).toBe('fill');
      expect(UI_ICONS[name].body).toMatch(/^<path fill-rule="evenodd" d="M[^"]+"\/>$/);
    }
  });

  it('draws the line family on one 24 px grid from paths and circles only, unpainted', () => {
    const line = Object.values(UI_ICONS).filter((icon) => icon.box === 24);
    expect(line.map((icon) => icon.name)).toEqual(
      expect.arrayContaining(['chevron-down', 'close', 'check', 'menu']),
    );
    for (const icon of line) {
      const tags = [...icon.body.matchAll(/<(\w+)/g)].map((m) => m[1]);
      expect(tags.every((tag) => tag === 'path' || tag === 'circle')).toBe(true);
      expect(icon.body).not.toMatch(/stroke|fill=|style/);
    }
  });
});

describe('iconAttributes', () => {
  it('strokes a line glyph at the family width and fills a custom icon', () => {
    expect(iconAttributes(ICON_CHEVRON_DOWN, 14)).toMatchObject({
      width: '14',
      viewBox: '0 0 24 24',
      fill: 'none',
      stroke: 'currentColor',
      'stroke-width': String(LINE_STROKE_WIDTH),
      'aria-hidden': 'true',
      'data-icon': 'chevron-down',
    });
    const search = iconAttributes(ICON_SEARCH, 20);
    expect(search).toMatchObject({
      fill: 'currentColor',
      viewBox: `0 0 ${ICON_SEARCH.box} ${ICON_SEARCH.box}`,
    });
    expect(search).not.toHaveProperty('stroke');
  });

  it('turns an icon into a CSS mask of the same shape', () => {
    const mask = decodeURIComponent(iconMaskUrl(ICON_CHEVRON_DOWN));
    expect(mask).toContain(`stroke-width='${LINE_STROKE_WIDTH}'`);
    expect(mask).toContain("d='M6 9l6 6 6-6'");
  });
});
