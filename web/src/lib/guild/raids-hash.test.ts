// web/src/lib/guild/raids-hash.test.ts
// The night deep link: built with the report id escaped, read back, and owned by the Raids tab.
import { describe, expect, it } from 'vitest';
import { tabAfterHashChange, tabFromHash } from '../../components/guild/GuildTabs.svelte';
import { nightFromHash, nightHref, tabIdOfHash } from './raids-hash';

describe('raids-hash', () => {
  it('builds #raids-{report_id} and reads it back', () => {
    expect(nightHref('/guild/us/pvp/olympus-xxvii', 'abc123def')).toBe(
      '/guild/us/pvp/olympus-xxvii#raids-abc123def',
    );
    expect(nightFromHash('#raids-abc123def')).toBe('abc123def');
    expect(nightFromHash(`#${nightHref('', 'a b').split('#')[1]}`)).toBe('a b');
  });

  it('reads no night from the tab hash, an empty id or a broken escape', () => {
    expect(nightFromHash('#raids')).toBeNull();
    expect(nightFromHash('#raids-')).toBeNull();
    expect(nightFromHash('#raids-%E0%A4%A')).toBeNull();
    expect(nightFromHash('#roster')).toBeNull();
  });

  it('keeps case in the id but lower-cases a tab hash', () => {
    expect(nightFromHash('#raids-AbC')).toBe('AbC');
    expect(tabIdOfHash('#Roster')).toBe('roster');
    expect(tabIdOfHash('#raids-AbC')).toBe('raids');
  });

  it('routes a night hash to the Raids tab, on load and on a hash change', () => {
    const tabs = ['overview', 'raids', 'progression'] as const;
    expect(tabFromHash('#raids-abc123def', tabs)).toBe('raids');
    expect(tabAfterHashChange('#raids-abc123def', tabs, 'overview')).toBe('raids');
  });
});
