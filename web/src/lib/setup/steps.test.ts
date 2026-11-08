// web/src/lib/setup/steps.test.ts
// /setup's three rows for every state of design/specs/2026-11-07-setup.md section 5, on the same
// fixtures the mocks and the parity spec read (tests/e2e/fixtures/home, setup) at the pinned clock.
import { describe, expect, it } from 'vitest';
import type { Device, Me } from '../account/api';
import { characterListCopy } from '../account/character-list-copy';
import { addonCopy } from '../addon/copy';
import { pairingCopy } from '../reports/copy';
import { yourReportsPageCopy } from '../reports/my-reports-copy';
import me34Json from '../../../tests/e2e/fixtures/home/me-34.json';
import meNewJson from '../../../tests/e2e/fixtures/home/me-new.json';
import devicesJson from '../../../tests/e2e/fixtures/setup/devices.json';
import {
  importPending,
  IMPORT_WAIT_MS,
  sessionOf,
  setupSteps,
  type SetupInputs,
  type StepRow,
} from './steps';

const NOW = new Date('2026-11-06T10:00:00Z');
const me34 = me34Json as unknown as Me;
const meNew = meNewJson as unknown as Me;
const devices = devicesJson as unknown as Device[];
const meBnet: Me = {
  ...me34,
  characters: me34.characters.map((c) => ({ ...c, build: c.build && { ...c.build, source: 'blizzard' } })),
};

function steps(input: Partial<SetupInputs>): StepRow[] {
  return setupSteps({ session: 'in', me: null, devices: [], platform: 'windows', now: NOW, ...input });
}
const marks = (rows: StepRow[]) => rows.map((r) => r.mark);
const primaries = (rows: StepRow[]) =>
  rows.filter((r) => r.action?.primary === true).map((r) => r.action?.id);
const texts = (row: StepRow) => row.lines.map((l) => l.text);

describe('setupSteps: signed out', () => {
  it('web-setup: Sign in is the gold step, the addon and companion are to do', () => {
    const rows = steps({ session: 'out' });
    expect(rows.map((r) => r.title)).toEqual([
      addonCopy.setupStep1Title,
      addonCopy.setupStep2Title,
      addonCopy.setupStep3Title,
    ]);
    expect(marks(rows)).toEqual(['next', 'todo', 'todo']);
    expect(primaries(rows)).toEqual(['signin']);
    expect(texts(rows[0])).toEqual([addonCopy.setupSignInBody]);
    expect(texts(rows[1])).toEqual([addonCopy.setupAddonNote]);
    expect(rows[1].action).toMatchObject({ id: 'curseforge', kind: 'link-button', primary: false });
    expect(rows[2].action).toMatchObject({ id: 'download', label: 'Download for Windows' });
  });

  it('web-setup-mac: row 3 is the Mac download, Windows keeps its own', () => {
    const mac = steps({ session: 'out', platform: 'mac' });
    expect(texts(mac[2])).toEqual([addonCopy.setupCompanionNote]);
    expect(mac[2].action).toMatchObject({ id: 'download', label: 'Download for Mac', kind: 'link-button' });
    expect(steps({ me: me34, platform: 'mac' })[2].action).toMatchObject({
      id: 'download',
      label: 'Download for Mac',
    });
  });

  it('web-setup-unknown: a system the page cannot name offers every build under one name', () => {
    const rows = steps({ session: 'out', platform: 'unknown' });
    expect(rows[2].action).toMatchObject({ id: 'download', label: addonCopy.setupDownloadCompanion });
  });

  it('web-setup-mobile: a phone gets the link to copy, never a download', () => {
    const rows = steps({ session: 'out', platform: 'mobile' });
    expect(texts(rows[2])).toEqual([addonCopy.setupCompanionNote, addonCopy.setupCompanionMobile]);
    expect(rows[2].action).toMatchObject({ id: 'copylink', kind: 'button', label: addonCopy.setupCopyLink });
    const signedIn = steps({ me: me34, platform: 'mobile' });
    expect(texts(signedIn[2])).toEqual([addonCopy.setupCompanionWaiting, addonCopy.setupCompanionMobile]);
    expect(signedIn[2].action).toMatchObject({ id: 'copylink', primary: true });
    expect(signedIn[2].mark).toBe('next');
  });

  it('web-setup-export: an export loaded ticks the addon row', () => {
    const rows = steps({ session: 'out', pointer: true });
    expect(marks(rows)).toEqual(['next', 'ok', 'todo']);
    expect(rows[1]).toMatchObject({
      action: null,
      lines: [{ text: addonCopy.setupAddonDoneLine, tone: 'ok' }],
    });
  });

  it('web-setup-signin-failed: row 1 failed, said as an alert, Sign in again the primary', () => {
    const rows = steps({ session: 'out', signinFailed: true });
    expect(rows[0]).toMatchObject({
      mark: 'miss',
      lines: [{ text: addonCopy.setupSignInFailed, tone: 'miss', alert: true }],
      action: { id: 'signin', primary: true },
    });
    expect(primaries(rows)).toEqual(['signin']);
  });

  it('web-setup-error: the read failed (not a 401), Try again is the primary', () => {
    const rows = steps({ session: 'error', devices: null });
    expect(rows[0]).toMatchObject({
      mark: 'miss',
      lines: [{ text: yourReportsPageCopy.sessionFailed, alert: true }],
      action: { id: 'retry', kind: 'button', primary: true },
    });
    expect(marks(rows)).toEqual(['miss', 'todo', 'todo']);
  });

  it('web-setup-loading: row 1 ok from the hint, every line a reserved skeleton, no action', () => {
    const rows = steps({ session: 'loading', devices: null });
    expect(marks(rows)).toEqual(['ok', 'todo', 'todo']);
    expect(rows.every((r) => r.action === null && r.lines.length === 1 && r.lines[0].skeleton === true)).toBe(
      true,
    );
  });
});

describe('setupSteps: signed in', () => {
  it('web-setup-none: row 1 warns, the lead in gold and the Korea line muted, Refresh the primary (B2)', () => {
    const rows = steps({ me: meNew });
    expect(rows[0].mark).toBe('warn');
    expect(rows[0].lines).toEqual([
      { text: 'Signed in as Hunthrop#1894', tone: 'text' },
      { text: characterListCopy.emptyLead, tone: 'warn' },
      { text: characterListCopy.emptyKorea, tone: 'muted' },
    ]);
    expect(rows[0].action).toMatchObject({ id: 'refresh', kind: 'button', primary: true });
    expect(primaries(rows)).toEqual(['refresh']);
    expect(texts(rows[2])).toEqual([addonCopy.setupCompanionWaiting]);
  });

  it('web-setup-importing: row 1 busy with the wait line, no primary anywhere (rulings Q2)', () => {
    const importing: Me = { ...meNew, bnet_import: { state: 'running', at: '2026-11-06T09:59:30Z' } };
    const rows = steps({ me: importing });
    expect(rows[0]).toMatchObject({ mark: 'busy', action: null });
    expect(texts(rows[0])).toEqual(['Signed in as Hunthrop#1894', addonCopy.setupImportWait]);
    expect(primaries(rows)).toEqual([]);
    expect(rows.some((r) => r.mark === 'next')).toBe(false);
  });

  it('leads the warning with the wait line while an import finished under two minutes ago', () => {
    const recent: Me = { ...meNew, bnet_import: { state: 'done', at: '2026-11-06T09:59:00Z' } };
    expect(texts(steps({ me: recent })[0])).toEqual([
      'Signed in as Hunthrop#1894',
      addonCopy.setupImportWait,
      characterListCopy.emptyLead,
      characterListCopy.emptyKorea,
    ]);
    const old: Me = { ...meNew, bnet_import: { state: 'done', at: '2026-11-06T09:50:00Z' } };
    expect(texts(steps({ me: old })[0])).not.toContain(addonCopy.setupImportWait);
  });

  it('web-setup-failed: back from Refresh with no import, row 1 failed with the refresh error and Try again', () => {
    const rows = steps({ me: meNew, refreshFailed: true });
    expect(rows[0]).toMatchObject({
      mark: 'miss',
      action: { id: 'refresh', label: characterListCopy.tryAgain, primary: true },
    });
    expect(rows[0].lines[1]).toEqual({ text: characterListCopy.refreshError, tone: 'miss', alert: true });
  });

  it('web-setup-bnet: Battle.net brought the character, CurseForge is the gold step (rulings Q1)', () => {
    const rows = steps({ me: meBnet });
    expect(marks(rows)).toEqual(['ok', 'next', 'todo']);
    expect(texts(rows[0])).toEqual(['Signed in as Hunthrop#1894 · Updated 1 h ago from Battle.net']);
    expect(texts(rows[1])).toEqual([addonCopy.setupAddonOptional]);
    expect(primaries(rows)).toEqual(['curseforge']);
  });

  it('web-setup-addon: the addon synced, the companion download is the gold step', () => {
    const rows = steps({ me: me34 });
    expect(marks(rows)).toEqual(['ok', 'ok', 'next']);
    expect(texts(rows[1])).toEqual(['Felsnap arrived from the addon 1 day ago']);
    expect(primaries(rows)).toEqual(['download']);
  });

  it('web-setup-pairing: a code out makes row 3 busy with the code, no primary', () => {
    const rows = steps({ me: me34, pairingCode: '4821-9930' });
    expect(rows[2]).toMatchObject({
      mark: 'busy',
      action: null,
      lines: [{ text: pairingCopy.waiting, code: '4821-9930' }],
    });
    expect(primaries(rows)).toEqual([]);
  });

  it('web-setup-done: three ticks, the device line, no primary', () => {
    const rows = steps({ me: me34, devices });
    expect(marks(rows)).toEqual(['ok', 'ok', 'ok']);
    expect(texts(rows[2])).toEqual(['Companion signed in on JHUNT-PC · last seen 3 minutes ago']);
    expect(rows.every((r) => r.action === null)).toBe(true);
  });

  it('claims nothing about the companion when the device list did not load', () => {
    const rows = steps({ me: me34, devices: null });
    expect(texts(rows[2])).toEqual([addonCopy.setupCompanionNote]);
  });
});

describe('importPending', () => {
  it('is true while running and for two minutes after a finished import', () => {
    expect(
      importPending({ ...meNew, bnet_import: { state: 'running', at: '2026-11-06T09:00:00Z' } }, NOW),
    ).toBe(true);
    const done = (ms: number): Me => ({
      ...meNew,
      bnet_import: { state: 'done', at: new Date(NOW.getTime() - ms).toISOString() },
    });
    expect(importPending(done(IMPORT_WAIT_MS - 1000), NOW)).toBe(true);
    expect(importPending(done(IMPORT_WAIT_MS + 1000), NOW)).toBe(false);
    expect(importPending({ ...meNew, bnet_import: { state: 'failed', at: NOW.toISOString() } }, NOW)).toBe(
      false,
    );
    expect(importPending(meNew, NOW)).toBe(false);
  });
});

describe('sessionOf', () => {
  it('reads a 401 as signed out and any other failure as an error (B1)', () => {
    expect(sessionOf('me')).toBe('in');
    expect(sessionOf('unauthorized')).toBe('out');
    expect(sessionOf('failed')).toBe('error');
  });
});
