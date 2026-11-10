// web/src/lib/guides/class-landing.test.ts
import { readdirSync, readFileSync } from 'node:fs';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { proseLabels, specCardLine, splitClassLanding } from './class-landing';

const GUIDES = fileURLToPath(new URL('../../content/guides/', import.meta.url));

function bodyOf(path: string): string {
  return readFileSync(path, 'utf8').split(/^---$/m).slice(2).join('---');
}

const classDirs = readdirSync(GUIDES, { withFileTypes: true })
  .filter((entry) => entry.isDirectory())
  .map((entry) => entry.name);

describe('splitClassLanding', () => {
  const BODY = `## What changed from 1.12

- **Pets** are reworked.

## Races

Skyborne can be hunters.`;

  it('splits the sections', () => {
    const landing = splitClassLanding(BODY, 'x.md');
    expect(landing.sections).toEqual([
      { heading: 'What changed from 1.12', markdown: '- **Pets** are reworked.' },
      { heading: 'Races', markdown: 'Skyborne can be hunters.' },
    ]);
  });

  it('fails on a heading outside the vocabulary', () => {
    expect(() => splitClassLanding('## The three specs\n\n- a', 'x.md')).toThrow(
      /not a class landing heading/,
    );
  });

  it('fails on headings out of order', () => {
    expect(() => splitClassLanding('## Races\n\na\n\n## What changed from 1.12\n\n- b', 'x.md')).toThrow(
      /out of order/,
    );
  });

  it('fails on an intro paragraph before the first heading (no "<Class> is the <noun>" blurb)', () => {
    expect(() => splitClassLanding('Hunter is the class with a pet.\n\n## Races\n\na', 'x.md')).toThrow(
      /no intro paragraph/,
    );
  });
});

describe('every class landing keeps the shape', () => {
  it('covers the nine classes', () => {
    expect(classDirs).toHaveLength(9);
  });

  for (const classSlug of classDirs) {
    it(`${classSlug}: no intro, then What changed from 1.12, then Races`, () => {
      const landing = splitClassLanding(bodyOf(join(GUIDES, classSlug, 'index.md')), `${classSlug}/index.md`);
      expect(landing.sections.slice(0, 2).map((s) => s.heading)).toEqual(['What changed from 1.12', 'Races']);
    });
  }
});

describe('proseLabels', () => {
  it('turns a bold-only paragraph into the Prose glabel, and leaves bold inside a sentence alone', () => {
    expect(proseLabels('<p><strong>Race picks.</strong></p>\n<ul><li>Orc</li></ul>')).toBe(
      '<p class="glabel">Race picks.</p>\n<ul><li>Orc</li></ul>',
    );
    expect(proseLabels('<p><strong>Opener.</strong> Mage Armor first.</p>')).toBe(
      '<p><strong>Opener.</strong> Mage Armor first.</p>',
    );
  });
});

describe('specCardLine', () => {
  it('names the role only when the spec heals or tanks', () => {
    expect(specCardLine('dps')).toBeUndefined();
    expect(specCardLine('healer')).toBe('Healer');
    expect(specCardLine('tank')).toBe('Tank');
  });
});
