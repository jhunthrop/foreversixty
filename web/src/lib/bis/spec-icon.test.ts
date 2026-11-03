// web/src/lib/bis/spec-icon.test.ts
import { existsSync, readdirSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import activeBuild from '../../data/active-build.json';
import { publishedIcon } from './spec-icon';

const ICONS = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  '../../../public/data',
  activeBuild.build,
  'icons',
);

describe('publishedIcon', () => {
  it('passes nothing through as nothing', () => {
    expect(publishedIcon(activeBuild.build, undefined)).toBeUndefined();
  });

  it('drops an icon this build does not publish', () => {
    expect(publishedIcon(activeBuild.build, 'no_such_icon_anywhere')).toBeUndefined();
  });

  it.skipIf(!existsSync(ICONS))('keeps an icon this build publishes', () => {
    const stem = readdirSync(ICONS)
      .find((file) => file.endsWith('.webp'))
      ?.replace(/\.webp$/, '');
    if (stem === undefined) return;
    expect(publishedIcon(activeBuild.build, stem)).toBe(stem);
  });
});
