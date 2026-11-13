// web/src/lib/setup/steps.test.ts
// /setup's three rows for every state of design/specs/2026-11-07-setup.md section 5, on the same
// fixtures the mocks and the parity spec read (tests/e2e/fixtures/home, setup) at the pinned clock.
// The plate's order is Sign in, The companion, The addon (owner 2026-11-10).
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
/** The rows by step. */
function by(rows: StepRow[]): Record<StepRow['id'], StepRow> {
  const find = (id: StepRow['id']): StepRow => {
    const found = rows.find((r) => r.id === id);
    if (found === undefined) throw new Error(`no ${id} row`);
    return found;
  };
  return { signin: find('signin'), companion: find('companion'), addon: find('addon') };
}
const withAddons = (addons: Device['addons']): Device[] => devices.map((d) => ({ ...d, addons }));
const INSTALLED: Device['addons'] = [
  { folder: 'ForeverSixty', version: '2.0.37' },
  { folder: 'ForeverSixtyData', version: '2026.11.09' },
];

describe('setupSteps: signed out', () => {
  it('web-setup: Sign in is the gold step, then the companion, then the addon, both to do', () => {
    const rows = steps({ session: 'out' });
    expect(rows.map((r) => r.title)).toEqual([
      addonCopy.setupSignInTitle,
      addonCopy.setupCompanionTitle,
      addonCopy.setupAddonTitle,
    ]);
    expect(rows.map((r) => [r.id, r.n])).toEqual([
      ['signin', 1],
      ['companion', 2],
      ['addon', 3],
    ]);
    expect(marks(rows)).toEqual(['next', 'todo', 'todo']);
    expect(primaries(rows)).toEqual(['signin']);
    const { signin, companion, addon } = by(rows);
    expect(texts(signin)).toEqual([addonCopy.setupSignInBody]);
    expect(companion.action).toMatchObject({ id: 'download', label: 'Download for Windows' });
    expect(texts(addon)).toEqual([addonCopy.setupAddonNote, addonCopy.setupAddonUnseen]);
    expect(addon.action).toMatchObject({ id: 'install', label: addonCopy.setupInstallAddon, primary: false });
  });

  it('web-setup-mac: the companion step is the Mac download, Windows keeps its own', () => {
    const mac = by(steps({ session: 'out', platform: 'mac' })).companion;
    expect(texts(mac)).toEqual([addonCopy.setupCompanionNote]);
    expect(mac.action).toMatchObject({ id: 'download', label: 'Download for Mac', kind: 'link-button' });
    expect(by(steps({ me: me34, platform: 'mac' })).companion.action).toMatchObject({
      id: 'download',
      label: 'Download for Mac',
    });
  });

  it('web-setup-unknown: a system the page cannot name offers every build under one name', () => {
    const { companion } = by(steps({ session: 'out', platform: 'unknown' }));
    expect(companion.action).toMatchObject({ id: 'download', label: addonCopy.setupDownloadCompanion });
  });

  it('web-setup-mobile: a phone gets the link to copy, never a download', () => {
    const { companion } = by(steps({ session: 'out', platform: 'mobile' }));
    expect(texts(companion)).toEqual([addonCopy.setupCompanionNote, addonCopy.setupCompanionMobile]);
    expect(companion.action).toMatchObject({
      id: 'copylink',
      kind: 'button',
      label: addonCopy.setupCopyLink,
    });
    const signedIn = by(steps({ me: me34, platform: 'mobile' })).companion;
    expect(texts(signedIn)).toEqual([addonCopy.setupCompanionWaiting, addonCopy.setupCompanionMobile]);
    expect(signedIn.action).toMatchObject({ id: 'copylink', primary: true });
    expect(signedIn.mark).toBe('next');
  });

  it('web-setup-export: an export loaded ticks the addon row', () => {
    const rows = steps({ session: 'out', pointer: true });
    expect(marks(rows)).toEqual(['next', 'todo', 'ok']);
    expect(by(rows).addon).toMatchObject({
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
    expect(rows.every((r) => r.action === null && r.lines.every((l) => l.skeleton === true))).toBe(true);
  });

  it('web-setup-loading: sign in and the companion reserve the second line their answer wraps to on a phone', () => {
    // A returning visitor's answer ("Signed in as ... · Updated from Battle.net ...", "JHUNT-PC ·
    // Windows ...") takes two lines on a phone: one reserved line moved the plate 20 px twice
    // (setup-states.spec.ts, layout shift 0.076 against 0.05).
    const rows = steps({ session: 'loading', devices: null });
    expect(rows.map((r) => r.lines.length)).toEqual([2, 2, 1]);
    for (const row of [by(rows).signin, by(rows).companion]) {
      expect(row.lines[0].phoneOnly).toBeUndefined();
      expect(row.lines[1]).toMatchObject({ skeleton: true, phoneOnly: true });
    }
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
    expect(texts(by(rows).companion)).toEqual([addonCopy.setupCompanionWaiting]);
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

  it('web-setup-partial: characters in, the missing count under the signed-in line, step still done', () => {
    const partial: Me = {
      ...me34,
      bnet_import: { state: 'partial', at: '2026-11-06T09:00:00Z', unresolved: 2 },
    };
    const rows = steps({ me: partial });
    expect(rows[0]).toMatchObject({ mark: 'ok', action: null });
    expect(texts(rows[0])).toEqual([
      `Signed in as Hunthrop#1894 · ${characterListCopy.updatedFrom(me34.bnet_imported_at as string, NOW)}`,
      addonCopy.setupImportMissing(2),
    ]);
    expect(rows[0].lines[1]).toMatchObject({ tone: 'warn' });
    const one: Me = { ...me34, bnet_import: { state: 'partial', at: '2026-11-06T09:00:00Z', unresolved: 1 } };
    expect(texts(steps({ me: one })[0])[1]).toBe(
      'One character is still missing. It comes in on a later visit.',
    );
    const done: Me = { ...me34, bnet_import: { state: 'done', at: '2026-11-06T09:00:00Z' } };
    expect(texts(steps({ me: done })[0])).toHaveLength(1);
  });

  it('web-setup-import-failed: no characters, the failed line leads the no-character lines, Refresh stays primary', () => {
    const failed: Me = { ...meNew, bnet_import: { state: 'failed', at: '2026-11-06T09:00:00Z' } };
    const rows = steps({ me: failed });
    expect(rows[0]).toMatchObject({ mark: 'warn' });
    expect(texts(rows[0])).toEqual([
      'Signed in as Hunthrop#1894',
      addonCopy.setupImportFailed,
      characterListCopy.emptyLead,
      characterListCopy.emptyKorea,
    ]);
    expect(rows[0].action).toMatchObject({ id: 'refresh', primary: true });
    // Every character skipped: the count, not the generic line.
    const skipped: Me = {
      ...meNew,
      bnet_import: { state: 'failed', at: '2026-11-06T09:00:00Z', unresolved: 3 },
    };
    expect(texts(steps({ me: skipped })[0])[1]).toBe(addonCopy.setupImportMissing(3));
    // A partial finished under two minutes ago still leads with the wait line.
    const recent: Me = {
      ...meNew,
      bnet_import: { state: 'partial', at: '2026-11-06T09:59:00Z', unresolved: 1 },
    };
    expect(texts(steps({ me: recent })[0])[1]).toBe(addonCopy.setupImportWait);
  });

  it('web-setup-failed: back from Refresh with no import, row 1 failed with the refresh error and Try again', () => {
    const rows = steps({ me: meNew, refreshFailed: true });
    expect(rows[0]).toMatchObject({
      mark: 'miss',
      action: { id: 'refresh', label: characterListCopy.tryAgain, primary: true },
    });
    expect(rows[0].lines[1]).toEqual({ text: characterListCopy.refreshError, tone: 'miss', alert: true });
  });

  it('web-setup-bnet: Battle.net brought the character, the companion is the gold step, the addon a plain to-do', () => {
    const rows = steps({ me: meBnet });
    expect(marks(rows)).toEqual(['ok', 'next', 'todo']);
    expect(texts(rows[0])).toEqual(['Signed in as Hunthrop#1894 · Updated 1 h ago from Battle.net']);
    const { addon } = by(rows);
    // The characters are in from Battle.net: the addon's line never ties them to it, and an addon
    // already installed is said to be seen once the companion runs, never failed.
    expect(texts(addon)).toEqual([addonCopy.setupAddonNote, addonCopy.setupAddonUnseen]);
    expect(addon.lines.every((l) => l.tone === 'muted')).toBe(true);
    expect(addon.action).toMatchObject({ id: 'install', primary: false });
    expect(primaries(rows)).toEqual(['download']);
  });

  it('web-setup-addon: the addon synced, the companion download is the gold step', () => {
    const rows = steps({ me: me34 });
    expect(marks(rows)).toEqual(['ok', 'next', 'ok']);
    expect(texts(by(rows).addon)).toEqual(['Felsnap arrived from the addon 1 day ago']);
    expect(primaries(rows)).toEqual(['download']);
  });

  it('web-setup-pairing: a code out makes the companion step busy with the code, no primary', () => {
    const rows = steps({ me: me34, pairingCode: '4821-9930' });
    expect(by(rows).companion).toMatchObject({
      mark: 'busy',
      action: null,
      lines: [{ text: pairingCopy.waiting, code: '4821-9930' }],
    });
    expect(primaries(rows)).toEqual([]);
  });

  it('web-setup-done: three ticks, the device line, no primary', () => {
    const rows = steps({ me: me34, devices });
    expect(marks(rows)).toEqual(['ok', 'ok', 'ok']);
    expect(texts(by(rows).companion)).toEqual(['Companion signed in on JHUNT-PC · last seen 3 minutes ago']);
    expect(rows.every((r) => r.action === null)).toBe(true);
  });

  it('claims nothing about the companion when the device list did not load', () => {
    const rows = steps({ me: me34, devices: null });
    expect(texts(by(rows).companion)).toEqual([addonCopy.setupCompanionNote]);
  });
});

describe('setupSteps: the addon as the companion sees it', () => {
  it('web-setup-companion: a companion that found no addon makes it the gold to-do, never a fault', () => {
    const rows = steps({ me: meBnet, devices: withAddons([]), addonLatest: '2.0.37' });
    expect(marks(rows)).toEqual(['ok', 'ok', 'next']);
    const { addon } = by(rows);
    expect(texts(addon)).toEqual(['Not in _forever_ on JHUNT-PC yet']);
    expect(addon.lines.every((l) => l.tone === 'muted' && l.alert === undefined)).toBe(true);
    expect(addon.action).toMatchObject({ id: 'install', primary: true });
  });

  it('web-setup-installed: installed on the PC, with its version, done with no action', () => {
    const rows = steps({ me: meBnet, devices: withAddons(INSTALLED), addonLatest: '2.0.37' });
    expect(marks(rows)).toEqual(['ok', 'ok', 'ok']);
    expect(by(rows).addon).toMatchObject({
      action: null,
      lines: [{ text: 'Installed on JHUNT-PC · 2.0.37', tone: 'ok' }],
    });
  });

  it('web-setup-outdated: installed but behind the feed, still done, with what is out and Update buttons', () => {
    const behind = withAddons([{ folder: 'ForeverSixty', version: '2.0.30' }]);
    const { addon } = by(steps({ me: meBnet, devices: behind, addonLatest: '2.0.37' }));
    expect(addon.mark).toBe('ok');
    expect(addon.lines).toEqual([
      { text: 'Installed on JHUNT-PC · 2.0.30', tone: 'ok' },
      { text: '2.0.37 is out', tone: 'text' },
    ]);
    expect(addon.action).toMatchObject({ id: 'update', label: addonCopy.setupUpdateAddon, primary: false });
    // No feed answer: no claim it is behind.
    expect(texts(by(steps({ me: meBnet, devices: behind, addonLatest: null })).addon)).toEqual([
      'Installed on JHUNT-PC · 2.0.30',
    ]);
  });

  it('prefers what the companion found over an older arrival, and keeps the arrival for a paste', () => {
    expect(texts(by(steps({ me: me34, devices: withAddons(INSTALLED) })).addon)).toEqual([
      'Installed on JHUNT-PC · 2.0.37',
    ]);
    // The companion found none, but the addon's data arrived (a paste, another PC): done.
    expect(by(steps({ me: me34, devices: withAddons([]) })).addon).toMatchObject({
      mark: 'ok',
      lines: [{ text: 'Felsnap arrived from the addon 1 day ago' }],
    });
  });

  it('claims nothing from a companion that has never scanned (an older build)', () => {
    const { addon } = by(steps({ me: meBnet, devices }));
    expect(texts(addon)).toEqual([addonCopy.setupAddonNote, addonCopy.setupAddonUnseen]);
    expect(addon.mark).toBe('next');
  });

  it('names the PC that has the addon when another companion does not', () => {
    const laptop: Device = {
      ...devices[0],
      id: 'd2',
      name: 'LAPTOP',
      last_seen_at: '2026-11-06T09:59:00Z',
      addons: [],
    };
    const rows = steps({ me: meBnet, devices: [...withAddons(INSTALLED), laptop] });
    expect(texts(by(rows).addon)).toEqual(['Installed on JHUNT-PC · 2.0.37']);
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
