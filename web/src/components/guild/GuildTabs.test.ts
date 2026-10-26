// web/src/components/guild/GuildTabs.test.ts
import { describe, expect, it } from 'vitest';
import { tabAfterHashChange, tabsForRole } from './GuildTabs.svelte';

const OFFICER = tabsForRole('officer');

describe('tabAfterHashChange', () => {
  it('follows a hash that names a tab', () => {
    expect(tabAfterHashChange('#raids', OFFICER, 'settings')).toBe('raids');
  });

  it('keeps the open tab for an in-page anchor that names no tab (Settings · #billing)', () => {
    expect(tabAfterHashChange('#billing', OFFICER, 'settings')).toBe('settings');
    expect(tabAfterHashChange('#discord', OFFICER, 'settings')).toBe('settings');
  });

  it('falls back to the first allowed tab for a tab the role may not open, or no hash', () => {
    expect(tabAfterHashChange('#settings', tabsForRole('member'), 'raids')).toBe('overview');
    expect(tabAfterHashChange('', OFFICER, 'settings')).toBe('overview');
  });
});
