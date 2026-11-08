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
  it('sends each link action where it goes, and nothing for a button or the download', () => {
    expect(stepHref('signin')).toContain(encodeURIComponent(SETUP_SIGNED_IN));
    expect(stepHref('curseforge')).toBe(CURSEFORGE_URL);
    expect(stepHref('paste')).toBe('#paste');
    expect(stepHref('download')).toBeNull();
    expect(stepHref('copylink')).toBeNull();
    expect(stepHref('retry')).toBeNull();
    expect(stepHref('refresh')).toBeNull();
    expect(SETUP_REFRESHED).toBe('/setup?refreshed=1');
  });
});

const UA = {
  windows:
    'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0 Safari/537.36',
  mac: 'Mozilla/5.0 (Macintosh; Intel Mac OS X 14_0) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 Safari/605.1.15',
  iphone:
    'Mozilla/5.0 (iPhone; CPU iPhone OS 18_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148',
  android:
    'Mozilla/5.0 (Linux; Android 15; Pixel 7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0 Mobile Safari/537.36',
  linux: 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0 Safari/537.36',
};

describe('visitorPlatform', () => {
  it('reads Windows and a Mac from the client hints, the platform or the user agent', () => {
    expect(visitorPlatform({ userAgentData: { platform: 'macOS', mobile: false } })).toBe('mac');
    expect(visitorPlatform({ platform: 'MacIntel', maxTouchPoints: 0 })).toBe('mac');
    expect(visitorPlatform({ userAgent: UA.mac })).toBe('mac');
    expect(visitorPlatform({ userAgentData: { platform: 'Windows', mobile: false } })).toBe('windows');
    expect(visitorPlatform({ platform: 'Win32', userAgent: UA.windows })).toBe('windows');
    expect(visitorPlatform({ userAgentData: { platform: 'Windows' }, platform: 'MacIntel' })).toBe('windows');
  });

  it('reads a phone or tablet, an iPad asking for the desktop site included', () => {
    expect(visitorPlatform({ userAgentData: { platform: 'Android', mobile: true } })).toBe('mobile');
    expect(visitorPlatform({ userAgentData: { platform: 'Android', mobile: false } })).toBe('mobile');
    expect(visitorPlatform({ platform: 'iPhone', userAgent: UA.iphone })).toBe('mobile');
    expect(visitorPlatform({ platform: 'Linux armv81', userAgent: UA.android })).toBe('mobile');
    expect(visitorPlatform({ platform: 'MacIntel', userAgent: UA.mac, maxTouchPoints: 5 })).toBe('mobile');
  });

  it('calls everything else unknown, the server render included', () => {
    expect(visitorPlatform(undefined)).toBe('unknown');
    expect(visitorPlatform({ platform: 'Linux x86_64', userAgent: UA.linux })).toBe('unknown');
    expect(visitorPlatform({ userAgentData: { platform: 'Chrome OS', mobile: false } })).toBe('unknown');
    expect(visitorPlatform({})).toBe('unknown');
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
