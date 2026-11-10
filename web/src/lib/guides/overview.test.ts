// web/src/lib/guides/overview.test.ts
import { describe, expect, it } from 'vitest';
import { splitOverview } from './overview';

const SECTION = `## Overview

### Strengths

- One
- Two

### Weaknesses

- Three

### In a raid

- **Role:** Caster DPS
- **Brings:** Arcane Intellect`;

describe('splitOverview', () => {
  it('fails on an opening "<Name> is the" blurb, and on any text before the first block', () => {
    expect(() =>
      splitOverview('## Overview\n\nFrost is the control spec.\n\n### Strengths\n\n- a', 'x.md'),
    ).toThrow(/x\.md: the Overview opens on a "<Name> is the" blurb/);
    expect(() =>
      splitOverview('## Overview\n\nFrost casts Frostbolt.\n\n### Strengths\n\n- a', 'x.md'),
    ).toThrow(/x\.md: the Overview has text before its first block/);
    expect(() => splitOverview('## Overview\n\nProtection is the shield tank.', 'x.md')).toThrow(/blurb/);
  });

  it('returns each block’s bullets as Markdown', () => {
    const parts = splitOverview(SECTION);
    expect(parts.strengths).toBe('- One\n- Two');
    expect(parts.weaknesses).toBe('- Three');
    expect(parts.raid).toBe('- **Role:** Caster DPS\n- **Brings:** Arcane Intellect');
  });

  it('is empty for a section written as one paragraph, so the page falls back to its prose', () => {
    const parts = splitOverview('## Overview\n\nJust a paragraph.');
    expect([parts.strengths, parts.weaknesses, parts.raid]).toEqual(['', '', '']);
  });

  it('ignores a heading it does not know rather than folding it into a block', () => {
    const parts = splitOverview('## Overview\n\n### Strengths\n\n- a\n\n### Other\n\n- b');
    expect(parts.strengths).toBe('- a');
  });
});
