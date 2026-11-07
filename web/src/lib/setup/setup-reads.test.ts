// web/src/lib/setup/setup-reads.test.ts
// /setup's small models: the head answer (answer.ts), the account read (state.ts, B1), the action
// links (links.ts), the visitor's platform (visitor-platform.ts, G4) and the device platform's
// player-facing label (account/platform-label.ts, G5).
import { beforeEach, describe, expect, it, vi } from 'vitest';
import type { Me } from '../account/api';
import { platformLabel } from '../account/platform-label';
import { addonCopy } from '../addon/copy';
import me34Json from '../../../tests/e2e/fixtures/home/me-34.json';
import meNewJson from '../../../tests/e2e/fixtures/home/me-new.json';
import { setupAnswer } from './answer';
import { CURSEFORGE_URL, SETUP_REFRESHED, SETUP_SIGNED_IN, stepHref } from './links';
import { visitorPlatform } from './visitor-platform';

/** What the mocked device list answers, and how often it was asked. */
const companions = { answer: (): Promise<unknown[]> => Promise.resolve([]), calls: 0 };
vi.mock('../account/api', async (importOriginal) => ({
  ...(await importOriginal<typeof import('../account/api')>()),
  listCompanions: () => {
    companions.calls += 1;
    return companions.answer();
  },
}));
const { loadSetupState } = await import('./state');

const NOW = new Date('2026-11-06T10:00:00Z');
const me34 = me34Json as unknown as Me;
const meNew = meNewJson as unknown as Me;

describe('setupAnswer', () => {
  it('is the character: name, class slug, the level line in the client order, imported, two links', () => {
    const answer = setupAnswer(me34, null, NOW);
    expect(answer).toMatchObject({
      name: 'Felsnap',
      classSlug: 'hunter',
      levelLine: 'Level 34 Beast Mastery Hunter',
      imported: 'Updated 1 h ago from Battle.net',
      open: { label: addonCopy.setupOpenCharacter('Felsnap') },
    });
    expect(answer?.open.href).toContain('felsnap');
    expect(answer?.bis?.label).toBe(addonCopy.setupBis);
  });

  it('is nothing signed out or with no character', () => {
    expect(setupAnswer(null, null, NOW)).toBeNull();
    expect(setupAnswer(meNew, null, NOW)).toBeNull();
  });
});

describe('loadSetupState', () => {
  beforeEach(() => {
    companions.calls = 0;
  });

  it('is signed in with the companions', async () => {
    companions.answer = () => Promise.resolve([{ id: 'd1' }]);
    expect(await loadSetupState(() => Promise.resolve(me34))).toEqual({
      session: 'in',
      me: me34,
      devices: [{ id: 'd1' }],
    });
  });

  it('reads a 401 (fetchMe answers null) as signed out', async () => {
    expect(await loadSetupState(() => Promise.resolve(null))).toEqual({
      session: 'out',
      me: null,
      devices: [],
    });
    expect(companions.calls).toBe(0);
  });

  it('reads any other failure (a 500, the network) as an error, never as signed out (B1)', async () => {
    expect(await loadSetupState(() => Promise.reject(new Error('500')))).toEqual({
      session: 'error',
      me: null,
      devices: null,
    });
  });

  it('claims nothing about the companion when the device list fails', async () => {
    companions.answer = () => Promise.reject(new TypeError('device list failed'));
    expect((await loadSetupState(() => Promise.resolve(me34))).devices).toBeNull();
  });
});

describe('stepHref', () => {
  it('sends each link action where it goes, and nothing for a button', () => {
    expect(stepHref('signin', 'x.exe')).toContain(encodeURIComponent(SETUP_SIGNED_IN));
    expect(stepHref('curseforge', 'x.exe')).toBe(CURSEFORGE_URL);
    expect(stepHref('download', 'x.exe')).toMatch(/^\/dl\/.*x\.exe$/);
    expect(stepHref('paste', 'x.exe')).toBe('#paste');
    expect(stepHref('retry', 'x.exe')).toBeNull();
    expect(stepHref('refresh', 'x.exe')).toBeNull();
    expect(SETUP_REFRESHED).toBe('/setup?refreshed=1');
  });
});

describe('visitorPlatform', () => {
  it('reads a Mac from the client hints, the platform or the user agent; everything else is Windows', () => {
    expect(visitorPlatform(undefined)).toBe('windows');
    expect(visitorPlatform({ userAgentData: { platform: 'macOS' } })).toBe('mac');
    expect(visitorPlatform({ platform: 'MacIntel' })).toBe('mac');
    expect(visitorPlatform({ userAgent: 'Mozilla/5.0 (Macintosh; Intel Mac OS X 14_0)' })).toBe('mac');
    expect(visitorPlatform({ platform: 'Win32' })).toBe('windows');
    expect(visitorPlatform({ userAgentData: { platform: 'Windows' }, platform: 'MacIntel' })).toBe('windows');
  });
});

describe('platformLabel', () => {
  it('says the platform as a player reads it (G5)', () => {
    expect(platformLabel('windows/amd64')).toBe('Windows');
    expect(platformLabel('darwin/arm64')).toBe('macOS');
    expect(platformLabel('linux/amd64')).toBe('Linux');
    expect(platformLabel('ios')).toBe('iOS');
    expect(platformLabel('android')).toBe('Android');
    expect(platformLabel('Windows')).toBe('Windows');
    expect(platformLabel('plan9/386')).toBe('plan9/386');
  });
});
