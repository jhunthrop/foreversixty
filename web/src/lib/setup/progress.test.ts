// web/src/lib/setup/progress.test.ts
// The lines under /setup's steps for a signed-in account (lib/setup/progress.ts).
import { describe, expect, it } from 'vitest';
import type { Device, Me, MeCharacter } from '../account/api';
import { addonCopy } from '../addon/copy';
import { addonLine, companionAddon, companionLine, newestAddonCharacter, signInLine } from './progress';

describe('the setup done lines (progress.ts)', () => {
  const NOW = new Date('2026-10-20T12:00:00Z');
  const user = { id: 1, battletag: 'hunthrop#1894', email: null, role: 'user', anonymize: false };
  const character = (
    name: string,
    capturedAt: string | null,
    source: 'addon' | 'blizzard' = 'addon',
  ): MeCharacter => ({
    key: `us/hardcore/${name.toLowerCase()}`,
    region: 'us',
    ruleset: 'hardcore',
    name,
    ...(capturedAt === null ? {} : { build: { source, captured_at: capturedAt } }),
  });
  const me = (characters: MeCharacter[], battletag: string | null = 'hunthrop#1894'): Me => ({
    user: { ...user, battletag },
    characters,
    guilds: [],
  });
  const device = (name: string, lastSeen: string | null, created = '2026-10-01T00:00:00Z'): Device => ({
    id: name,
    name,
    platform: 'windows',
    created_at: created,
    last_seen_at: lastSeen,
  });

  it('says nothing signed out', () => {
    expect(signInLine(null)).toBeNull();
    expect(addonLine(null, [], null, NOW)).toBeNull();
    expect(companionLine(null, [], NOW)).toBeNull();
  });

  it('names the signed-in account: the battletag, else the email', () => {
    expect(signInLine(me([]))).toEqual({ done: true, text: 'Signed in as hunthrop#1894' });
    const emailOnly = { ...me([], null), user: { ...user, battletag: null, email: 'a@b.gg' } };
    expect(signInLine(emailOnly)?.text).toBe('Signed in as a@b.gg');
  });

  it('names the newest character the addon synced and when, ignoring Battle.net-only builds', () => {
    const account = me([
      character('Old Hunter', '2026-10-10T00:00:00Z'),
      character('Bow Jackzon', '2026-10-20T11:00:00Z'),
      character('Blizzard Only', '2026-10-20T11:59:00Z', 'blizzard'),
      character('No Build', null),
    ]);
    expect(newestAddonCharacter(account.characters)?.character.name).toBe('Bow Jackzon');
    expect(addonLine(account, [], null, NOW)).toMatchObject({
      done: true,
      text: 'Bow Jackzon arrived from the addon 1 hour ago',
    });
  });

  it('never ties Battle.net characters to the addon: the plain note until something shows it', () => {
    const bnet = me([character('Blizzard Only', '2026-10-20T11:59:00Z', 'blizzard')]);
    const unseen = { done: false, text: addonCopy.setupAddonNote, scanned: false };
    expect(addonLine(bnet, [], null, NOW)).toEqual(unseen);
    expect(addonLine(me([]), [], null, NOW)).toEqual(unseen);
  });

  it('reads the addon off the newest companion that scanned, and nothing from one that never did', () => {
    const scanned = (name: string, seen: string, version: string | null): Device => ({
      ...device(name, seen),
      addons: version === null ? [] : [{ folder: 'ForeverSixty', version }],
    });
    expect(companionAddon([device('OLD', '2026-10-20T11:00:00Z')])).toBeNull();
    expect(companionAddon(null)).toBeNull();
    const none = companionAddon([scanned('JH-DESKTOP', '2026-10-20T11:00:00Z', null)]);
    expect(none).toMatchObject({ device: { name: 'JH-DESKTOP' }, version: null });
    const found = companionAddon([
      scanned('OLD', '2026-10-19T00:00:00Z', '2.0.30'),
      scanned('JH-DESKTOP', '2026-10-20T11:00:00Z', '2.0.37'),
      scanned('BARE', '2026-10-20T11:59:00Z', null),
    ]);
    expect(found).toMatchObject({ device: { name: 'JH-DESKTOP' }, version: '2.0.37' });
    const bnet = me([character('Blizzard Only', '2026-10-20T11:59:00Z', 'blizzard')]);
    const at = (version: string) => [scanned('JH-DESKTOP', '2026-10-20T11:00:00Z', version)];
    expect(addonLine(bnet, at('2.0.36'), '2.0.37', NOW)).toEqual({
      done: true,
      text: 'Installed on JH-DESKTOP · 2.0.36',
      behind: '2.0.37 is out',
      scanned: true,
    });
    expect(addonLine(bnet, at(''), '2.0.37', NOW)).toMatchObject({
      done: true,
      text: 'Installed on JH-DESKTOP',
    });
  });

  it('names the most recently seen companion device and when', () => {
    const devices = [
      device('OLD-LAPTOP', '2026-10-02T00:00:00Z'),
      device('JH-DESKTOP', '2026-10-20T11:57:00Z'),
      device('NEVER-SEEN', null, '2026-10-03T00:00:00Z'),
    ];
    expect(companionLine(me([]), devices, NOW)).toEqual({
      done: true,
      text: 'Companion signed in on JH-DESKTOP · last seen 3 minutes ago',
    });
  });

  it('falls back to the creation time for a device never heard from', () => {
    const line = companionLine(me([]), [device('JH-DESKTOP', null, '2026-10-20T11:58:00Z')], NOW);
    expect(line?.text).toBe('Companion signed in on JH-DESKTOP · last seen 2 minutes ago');
  });

  it('says no companion has signed in, and claims nothing when the device list did not load', () => {
    expect(companionLine(me([]), [], NOW)).toEqual({ done: false, text: 'No companion signed in yet' });
    expect(companionLine(me([]), null, NOW)).toBeNull();
  });
});
