// web/src/lib/guides/overview.test.ts
import { describe, expect, it } from 'vitest';
import { splitOverview } from './overview';

const SECTION = `## Overview

Frost is the **Mage** control spec, see [the planner](/planner).

### Strengths

- One
- Two

### Weaknesses

- Three

### In a raid

- **Role:** Caster DPS
- **Brings:** Arcane Intellect`;

describe('splitOverview', () => {
  it('takes the sentence before the first block as a plain-text summary', () => {
    expect(splitOverview(SECTION).summary).toBe('Frost is the Mage control spec, see the planner.');
  });

  it('returns each block’s bullets as Markdown', () => {
    const parts = splitOverview(SECTION);
    expect(parts.strengths).toBe('- One\n- Two');
    expect(parts.weaknesses).toBe('- Three');
    expect(parts.raid).toBe('- **Role:** Caster DPS\n- **Brings:** Arcane Intellect');
  });

  it('is empty for a section written as one paragraph, so the page falls back to its prose', () => {
    const parts = splitOverview('## Overview\n\nJust a paragraph.');
    expect(parts.summary).toBe('Just a paragraph.');
    expect([parts.strengths, parts.weaknesses, parts.raid]).toEqual(['', '', '']);
  });

  it('ignores a heading it does not know rather than folding it into a block', () => {
    const parts = splitOverview('## Overview\n\nS.\n\n### Strengths\n\n- a\n\n### Other\n\n- b');
    expect(parts.strengths).toBe('- a');
  });
});
