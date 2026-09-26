// web/src/lib/guides/sections.test.ts
import { describe, expect, it } from 'vitest';
import { splitSpecSections } from './sections';

const BODY = `## Overview

Some overview text.

## Talents and builds

Pick these talents.

- one
- two

## Stat priority

1. **Attack power** — text.
`;

describe('splitSpecSections', () => {
  it('slices each heading through to the next one, in document order', () => {
    const sections = splitSpecSections(BODY);
    expect([...sections.keys()]).toEqual(['Overview', 'Talents and builds', 'Stat priority']);
    expect(sections.get('Overview')).toBe('## Overview\n\nSome overview text.');
    expect(sections.get('Talents and builds')).toBe(
      '## Talents and builds\n\nPick these talents.\n\n- one\n- two',
    );
    expect(sections.get('Stat priority')).toBe('## Stat priority\n\n1. **Attack power** — text.');
  });

  it('returns an empty map for a body with no recognised heading', () => {
    expect(splitSpecSections('just some text').size).toBe(0);
  });

  it('never returns a key outside SPEC_SECTIONS', () => {
    const sections = splitSpecSections('## Not a real section\n\nx\n\n## Overview\n\ny');
    expect([...sections.keys()]).toEqual(['Overview']);
  });
});
