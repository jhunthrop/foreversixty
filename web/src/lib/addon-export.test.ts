import { describe, expect, it, vi } from 'vitest';

const fetchSimInput = vi.fn();
vi.mock('./sim/api', () => ({ fetchSimInput: (...args: unknown[]) => fetchSimInput(...args) }));

const { lookupAddonExport } = await import('./addon-export');

const PATH = { region: 'us' as const, ruleset: 'normal' as const, slug: 'thrallgar' };

describe('lookupAddonExport', () => {
  it('carries the FS1 code when the newest source is an addon export', async () => {
    fetchSimInput.mockResolvedValueOnce({
      spec: 'warrior-fury',
      gear: 'FS1:1.60.1.69893:warrior:human:0/0/0:',
      talents: '',
      buffs: [],
      captured_at: '2026-09-20T00:00:00Z',
      source: 'addon',
    });

    const result = await lookupAddonExport(PATH);
    expect(result.path).toBe(PATH);
    expect(result.code).toBe('FS1:1.60.1.69893:warrior:human:0/0/0:');
  });

  it('carries the FS1 code when the newest source is a Battle.net import', async () => {
    fetchSimInput.mockResolvedValueOnce({
      spec: 'warrior-fury',
      gear: 'FS1:1.60.1.69893:warrior:human:0/0/0:',
      talents: '',
      buffs: [],
      captured_at: '2026-09-20T00:00:00Z',
      source: 'blizzard',
    });

    const result = await lookupAddonExport(PATH);
    expect(result.code).toBe('FS1:1.60.1.69893:warrior:human:0/0/0:');
  });

  it('reads no export when the newest source is a logged fight', async () => {
    fetchSimInput.mockResolvedValueOnce({
      spec: 'mage-fire',
      gear: { trinkets: [] },
      talents: '31/0/20',
      buffs: [],
      captured_at: '2026-09-20T00:00:00Z',
      source: 'fight',
    });

    const result = await lookupAddonExport(PATH);
    expect(result.code).toBeNull();
  });

  it('reads no export, not an error, when the API call fails', async () => {
    fetchSimInput.mockRejectedValueOnce(new Error('offline'));

    const result = await lookupAddonExport(PATH);
    expect(result.code).toBeNull();
  });
});
