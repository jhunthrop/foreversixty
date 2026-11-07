// web/src/lib/setup/steps.ts
// /setup's three steps as data (design/specs/2026-11-07-setup.md section 5): each row's mark, its
// lines and its one action, for every state the page can be in. The page's island draws these rows
// and the mock generator (design/mocks/setup_facts.mjs) renders the same rows, so the mock and the
// build share one truth. Pure: the reads (`/v1/me`, `/v1/devices`, the hints) come in as inputs.
import type { Device, Me } from '../account/api';
import { characterListCopy } from '../account/character-list-copy';
import { pairingCopy } from '../reports/copy';
import { yourReportsPageCopy } from '../reports/my-reports-copy';
import { addonCopy } from '../addon/copy';
import {
  addonLine,
  companionLine,
  companionPlatformsLine,
  newestAddonCharacter,
  signInLine,
} from './progress';
import type { VisitorPlatform } from './visitor-platform';

/** `ok` done, `next` the first step not done (the page's one primary), `todo` optional or later,
 *  `warn` signed in with nothing to show, `miss` failed, `busy` waiting on something running. */
export type StepMark = 'ok' | 'next' | 'todo' | 'warn' | 'miss' | 'busy';
export type StepTone = 'muted' | 'text' | 'ok' | 'warn' | 'miss';

export interface StepLineOut {
  text: string;
  tone: StepTone;
  /** An error sentence: drawn with `role="alert"`. */
  alert?: boolean;
  /** Literal text a player types or matches, drawn as `Code` after the text. */
  code?: string;
  /** A reserved line while the reads are in flight. */
  skeleton?: boolean;
}

export type StepActionId = 'signin' | 'retry' | 'refresh' | 'curseforge' | 'download' | 'paste';

export interface StepAction {
  id: StepActionId;
  label: string;
  /** `button` and `link-button` are the system Button; `link` is the gold text link with `›`. */
  kind: 'button' | 'link-button' | 'link';
  primary: boolean;
}

export interface StepRow {
  n: 1 | 2 | 3;
  title: string;
  mark: StepMark;
  lines: StepLineOut[];
  action: StepAction | null;
}

/** How the visitor's account read went: signed out (401 or no session), failed, still in flight, or signed in. */
export type SessionState = 'out' | 'error' | 'loading' | 'in';

export interface SetupInputs {
  session: SessionState;
  me: Me | null;
  /** Null when the device list did not load (no claim is made about the companion). */
  devices: readonly Device[] | null;
  platform: VisitorPlatform;
  /** Back from a Battle.net callback that failed (`?signin=failed`). */
  signinFailed?: boolean;
  /** Back from Refresh with no import (`compareSnapshot` said `error`). */
  refreshFailed?: boolean;
  /** Signed out with an export loaded or a pointer stored. */
  pointer?: boolean;
  /** A fallback pairing code is out and unused. */
  pairingCode?: string | null;
  now?: Date;
}

/** Two minutes: the API's `importContinueBudget`, the longest an import can still be landing. */
export const IMPORT_WAIT_MS = 2 * 60 * 1000;

/** The import is running, or finished so recently that characters may still be landing (rulings Q2). */
export function importPending(me: Me, now: Date): boolean {
  const run = me.bnet_import;
  if (run === undefined) return false;
  if (run.state === 'running') return true;
  return run.state === 'done' && now.getTime() - new Date(run.at).getTime() < IMPORT_WAIT_MS;
}

const line = (text: string, tone: StepTone = 'muted', extra: Partial<StepLineOut> = {}): StepLineOut => ({
  text,
  tone,
  ...extra,
});
const SKELETON: StepLineOut = { text: '', tone: 'muted', skeleton: true };

function act(id: StepActionId, label: string, kind: StepAction['kind'] = 'link-button'): StepAction {
  return { id, label, kind, primary: false };
}

const SIGN_IN = act('signin', addonCopy.setupSignInAction);
const CURSEFORGE = act('curseforge', addonCopy.installCurseForge);

function downloadLabel(platform: string): string {
  return addonCopy.setupCompanionDownload(platform);
}

/** Row 3 not done: the download on Windows, the export route on a Mac (U10). */
function companionTodo(
  platform: VisitorPlatform,
  text: string,
  downloadPlatform: string,
  platforms: readonly string[],
): StepRow {
  if (platform === 'mac') {
    const macLine = companionPlatformsLine(platforms) ?? text;
    return row(3, 'todo', [line(macLine)], act('paste', addonCopy.setupTryExport, 'link'));
  }
  return row(3, 'todo', [line(text)], act('download', downloadLabel(downloadPlatform)));
}

function row(n: 1 | 2 | 3, mark: StepMark, lines: StepLineOut[], action: StepAction | null): StepRow {
  const titles = [addonCopy.setupStep1Title, addonCopy.setupStep2Title, addonCopy.setupStep3Title];
  return { n, title: titles[n - 1], mark, lines, action };
}

/** The first `todo` row with an action becomes `next` and carries the page's one primary (W.10). */
function promoteFirst(rows: StepRow[]): StepRow[] {
  // A primary already placed (sign in, retry, refresh), or a step still running: no second gold.
  if (rows.some((r) => r.action?.primary === true || r.mark === 'busy')) return rows;
  const first = rows.findIndex((r) => r.mark === 'todo' && r.action !== null);
  return rows.map((r, i) =>
    i === first && r.action !== null
      ? { ...r, mark: 'next', action: { ...r.action, primary: r.action.kind !== 'link' } }
      : r,
  );
}

function primary(action: StepAction): StepAction {
  return { ...action, primary: true };
}

/** Signed out, failed or loading: no account to read. */
function anonymousRows(input: SetupInputs, platforms: readonly string[]): StepRow[] {
  const download = platforms[0] ?? 'Windows';
  const addon =
    input.pointer === true
      ? row(2, 'ok', [line(addonCopy.setupAddonDoneLine, 'ok')], null)
      : row(2, 'todo', [line(addonCopy.setupAddonNote)], CURSEFORGE);
  const companion = companionTodo(input.platform, addonCopy.setupCompanionNote, download, platforms);
  if (input.session === 'loading') {
    return [
      row(1, 'ok', [SKELETON], null),
      row(2, 'todo', [SKELETON], null),
      row(3, 'todo', [SKELETON], null),
    ];
  }
  if (input.session === 'error') {
    const retry = primary(act('retry', characterListCopy.tryAgain, 'button'));
    return [
      row(1, 'miss', [line(yourReportsPageCopy.sessionFailed, 'miss', { alert: true })], retry),
      addon,
      companion,
    ];
  }
  const signIn =
    input.signinFailed === true
      ? row(1, 'miss', [line(addonCopy.setupSignInFailed, 'miss', { alert: true })], primary(SIGN_IN))
      : row(1, 'next', [line(addonCopy.setupSignInBody)], primary(SIGN_IN));
  return [signIn, addon, companion];
}

/** Row 1 signed in: who, and whether the characters are here, coming, or missing. */
function signInRow(me: Me, input: SetupInputs, now: Date): StepRow {
  const signed = signInLine(me)?.text ?? '';
  if (me.characters.length > 0) {
    const imported =
      me.bnet_imported_at === undefined
        ? ''
        : ` · ${characterListCopy.updatedFrom(me.bnet_imported_at, now)}`;
    return row(1, 'ok', [line(`${signed}${imported}`, 'ok')], null);
  }
  if (me.bnet_import?.state === 'running') {
    return row(1, 'busy', [line(signed, 'text'), line(addonCopy.setupImportWait)], null);
  }
  if (input.refreshFailed === true) {
    const retry = primary(act('refresh', characterListCopy.tryAgain, 'button'));
    return row(
      1,
      'miss',
      [line(signed, 'text'), line(characterListCopy.refreshError, 'miss', { alert: true })],
      retry,
    );
  }
  const wait = importPending(me, now) ? [line(addonCopy.setupImportWait, 'warn')] : [];
  return row(
    1,
    'warn',
    [
      line(signed, 'text'),
      ...wait,
      line(characterListCopy.emptyLead, 'warn'),
      line(characterListCopy.emptyKorea),
    ],
    primary(act('refresh', characterListCopy.refreshFromBattlenet, 'button')),
  );
}

function companionRow(me: Me, input: SetupInputs, platforms: readonly string[], now: Date): StepRow {
  const download = platforms[0] ?? 'Windows';
  const status = companionLine(me, input.devices, now);
  if (status?.done === true) return row(3, 'ok', [line(status.text, 'ok')], null);
  if (input.pairingCode != null && input.pairingCode !== '') {
    return row(3, 'busy', [line(pairingCopy.waiting, 'muted', { code: input.pairingCode })], null);
  }
  return companionTodo(input.platform, status?.text ?? addonCopy.setupCompanionNote, download, platforms);
}

/** Every row of the plate for one state. `platforms` is companion.json's download platforms. */
export function setupSteps(input: SetupInputs, platforms: readonly string[]): StepRow[] {
  const now = input.now ?? new Date();
  const me = input.me;
  if (input.session !== 'in' || me === null) return promoteFirst(anonymousRows(input, platforms));
  const first = signInRow(me, input, now);
  const addonDone = newestAddonCharacter(me.characters) !== null;
  const addonText = addonLine(me, now)?.text ?? addonCopy.setupAddonNote;
  const addon = addonDone
    ? row(2, 'ok', [line(addonText, 'ok')], null)
    : row(2, 'todo', [line(addonText)], CURSEFORGE);
  return promoteFirst([first, addon, companionRow(me, input, platforms, now)]);
}

/** The session the page reads from `/v1/me`'s answer: a 401 (or no session) is signed out, anything else that fails is an error (B1). */
export function sessionOf(result: 'me' | 'unauthorized' | 'failed'): SessionState {
  if (result === 'me') return 'in';
  return result === 'unauthorized' ? 'out' : 'error';
}
