// web/src/lib/planner/tree-source.test.ts
import { describe, expect, it } from 'vitest';
import { PREBETA_BUILD, treeSourceNotice } from './tree-source';

describe('treeSourceNotice', () => {
  it('names the client build the trees were read from', () => {
    expect(treeSourceNotice('1.60.1.69893')).toBe(
      'Talent trees read from the game client, build 1.60.1.69893.',
    );
  });

  it('says so plainly while the trees are still the pre-beta snapshot', () => {
    expect(treeSourceNotice(PREBETA_BUILD)).toBe(
      'Talent trees from a pre-beta Wowhead snapshot; the client’s own trees replace them at the beta.',
    );
  });

  it('never claims the trees are Classic Era', () => {
    expect(treeSourceNotice('1.15.9.69722')).not.toContain('Classic Era');
  });
});
