// web/src/lib/guides/class-landing.test.ts
import { readdirSync, readFileSync } from 'node:fs';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { cardLineFor, proseLabels, splitClassLanding } from './class-landing';

const GUIDES = fileURLToPath(new URL('../../content/guides/', import.meta.url));

function bodyOf(path: string): string {
  return readFileSync(path, 'utf8').split(/^---$/m).slice(2).join('---');
}

const classDirs = readdirSync(GUIDES, { withFileTypes: true })
  .filter((entry) => entry.isDirectory())
  .map((entry) => entry.name);

describe('cardLineFor', () => {
  it('fails on a spec guide with no Overview', () => {
    expect(() => cardLineFor('## Talents and builds\n\nx', 'x.md')).toThrow(/no ## Overview/);
  });

  it('fails on a line too long for a phone card', () => {
    expect(() => cardLineFor(`## Overview\n\n${'Long words. '.repeat(20)}`, 'x.md')).toThrow(/over 200/);
  });
});

describe('splitClassLanding', () => {
  const BODY = `Hunter is the class with a pet.

## What changed from 1.12

- **Pets** are reworked.

## Races

Skyborne can be hunters.`;

  it('splits the identity from the sections', () => {
    const landing = splitClassLanding(BODY, 'x.md');
    expect(landing.identity).toBe('Hunter is the class with a pet.');
    expect(landing.sections).toEqual([
      { heading: 'What changed from 1.12', markdown: '- **Pets** are reworked.' },
      { heading: 'Races', markdown: 'Skyborne can be hunters.' },
    ]);
  });

  it('fails on a heading outside the vocabulary', () => {
    expect(() => splitClassLanding('Id.\n\n## The three specs\n\n- a', 'x.md')).toThrow(
      /not a class landing heading/,
    );
  });

  it('fails on headings out of order', () => {
    expect(() =>
      splitClassLanding('Id.\n\n## Races\n\na\n\n## What changed from 1.12\n\n- b', 'x.md'),
    ).toThrow(/out of order/);
  });

  it('fails on an identity with Markdown in it', () => {
    expect(() => splitClassLanding('The **hunter** shoots.\n\n## Races\n\na', 'x.md')).toThrow(/plain text/);
  });

  it('fails on an identity of two paragraphs', () => {
    expect(() => splitClassLanding('One.\n\nTwo.\n\n## Races\n\na', 'x.md')).toThrow(
      /more than one paragraph/,
    );
  });
});

describe('every class landing keeps the shape', () => {
  it('covers the nine classes', () => {
    expect(classDirs).toHaveLength(9);
  });

  for (const classSlug of classDirs) {
    it(`${classSlug}: identity, then What changed from 1.12, then Races`, () => {
      const landing = splitClassLanding(bodyOf(join(GUIDES, classSlug, 'index.md')), `${classSlug}/index.md`);
      expect(landing.sections.slice(0, 2).map((s) => s.heading)).toEqual(['What changed from 1.12', 'Races']);
    });

    it(`${classSlug}: every spec card line says how the spec plays, in one card-sized line`, () => {
      const specFiles = readdirSync(join(GUIDES, classSlug)).filter(
        (f) => f.endsWith('.md') && f !== 'index.md',
      );
      for (const file of specFiles) {
        expect(() => cardLineFor(bodyOf(join(GUIDES, classSlug, file)), file)).not.toThrow();
      }
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
