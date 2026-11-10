// web/src/lib/setup/downloads.test.ts
import { describe, expect, it } from 'vitest';
import { COMPANION_DOWNLOADS, downloadFor, otherDownloads, type CompanionDownload } from './downloads';

const WINDOWS: CompanionDownload = {
  os: 'windows',
  platform: 'Windows',
  asset: 'foreversixty-companion-setup.exe',
  firstRun: 'w',
};
const MAC: CompanionDownload = {
  os: 'mac',
  platform: 'Mac',
  asset: 'foreversixty-companion_darwin_universal.app.zip',
  firstRun: 'm',
};

describe('downloadFor', () => {
  it('gives Windows and a Mac their own build, whatever the list order', () => {
    expect(downloadFor('windows', [MAC, WINDOWS])).toBe(WINDOWS);
    expect(downloadFor('mac', [WINDOWS, MAC])).toBe(MAC);
  });

  it('gives a phone and an unknown system no own build', () => {
    expect(downloadFor('mobile', [WINDOWS, MAC])).toBeNull();
    expect(downloadFor('unknown', [WINDOWS, MAC])).toBeNull();
  });

  it('reads companion.json by default, which lists Windows and the Mac', () => {
    expect(COMPANION_DOWNLOADS.map((d) => d.os)).toEqual(['windows', 'mac']);
    expect(downloadFor('mac')?.asset).toBe('foreversixty-companion_darwin_universal.app.zip');
  });
});

describe('otherDownloads', () => {
  it('is every build but the own one, and every build when there is none', () => {
    expect(otherDownloads('windows', [WINDOWS, MAC])).toEqual([MAC]);
    expect(otherDownloads('mac', [WINDOWS, MAC])).toEqual([WINDOWS]);
    expect(otherDownloads('unknown', [WINDOWS, MAC])).toEqual([WINDOWS, MAC]);
  });
});
