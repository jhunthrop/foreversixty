import { describe, expect, it } from 'vitest';
import { META_DESCRIPTION_MAX, factSchema, guideSchema } from './content.config';

const valid = {
  title: 'Hall of Thanes',
  updated: new Date('2026-09-12'),
  confidence: 'single-source',
  sources: [{ label: 'Blizzard panel recap', url: 'https://news.blizzard.com/x', kind: 'blizzard' }],
};

describe('factSchema', () => {
  it('accepts a complete entry', () => {
    expect(factSchema.safeParse(valid).success).toBe(true);
  });
  it('rejects a missing updated date', () => {
    const { updated, ...rest } = valid;
    expect(factSchema.safeParse(rest).success).toBe(false);
  });
  it('rejects an empty sources list', () => {
    expect(factSchema.safeParse({ ...valid, sources: [] }).success).toBe(false);
  });
  it('rejects an unknown source kind', () => {
    expect(
      factSchema.safeParse({ ...valid, sources: [{ ...valid.sources[0], kind: 'rumor' }] }).success,
    ).toBe(false);
  });
});

describe('guideSchema', () => {
  const guide = { ...valid, classSlug: 'mage' };
  const withDescription = (length: number): boolean =>
    guideSchema.safeParse({ ...guide, description: 'x'.repeat(length) }).success;

  it('takes a description that fits a search snippet', () => {
    expect(withDescription(META_DESCRIPTION_MAX)).toBe(true);
  });
  it('rejects a description a search result would cut off', () => {
    expect(withDescription(META_DESCRIPTION_MAX + 1)).toBe(false);
  });
});
