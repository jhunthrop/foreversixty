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
import { addonLine, companionLine, signInLine } from './progress';
import { downloadFor } from './downloads';
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
  /** Drawn on a phone only (below md): the second line an answer wraps to there. */
  phoneOnly?: boolean;
}

/** `install` and `update` are the addon's store buttons as one group (CurseForge, then Wago). */
export type StepActionId =
  'signin' | 'retry' | 'refresh' | 'install' | 'update' | 'download' | 'paste' | 'copylink';

export interface StepAction {
  id: StepActionId;
  label: string;
  /** `button` and `link-button` are the system Button; `link` is the gold text link with `›`. */
  kind: 'button' | 'link-button' | 'link';
  primary: boolean;
}

/** The three steps, in the plate's order: the companion before the addon, because the companion
 *  is what moves the addon's data and it sees the addon in the game folder (owner 2026-11-10). */
export const STEP_ORDER = ['signin', 'companion', 'addon'] as const;
export type StepId = (typeof STEP_ORDER)[number];

export interface StepRow {
  id: StepId;
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
  /** The addon feed's newest version; null when the feed did not answer (no "behind" claim). */
  addonLatest?: string | null;
  now?: Date;
}

/** Two minutes: the API's `importContinueBudget`, the longest an import can still be landing. */
export const IMPORT_WAIT_MS = 2 * 60 * 1000;

/** The import is running, or finished so recently that characters may still be landing (rulings Q2). */
export function importPending(me: Me, now: Date): boolean {
  const run = me.bnet_import;
  if (run === undefined) return false;
  if (run.state === 'running') return true;
  const finished = run.state === 'done' || run.state === 'partial';
  return finished && now.getTime() - new Date(run.at).getTime() < IMPORT_WAIT_MS;
}

/** The line that says what the last import left out, when it did: the missing count (partial, or
 *  failed with every character skipped), else, with no character in, that a failed run did not
 *  finish. A failed state beside characters says nothing: the visit refresh brought them in and
 *  only the next sign-in rewrites the state. Empty otherwise. */
function importShortfall(me: Me): StepLineOut[] {
  const run = me.bnet_import;
  if (run === undefined || run.state === 'running' || run.state === 'done') return [];
  const missing = run.unresolved ?? 0;
  if (missing > 0) return [line(addonCopy.setupImportMissing(missing), 'warn')];
  const failedEmpty = run.state === 'failed' && me.characters.length === 0;
  return failedEmpty ? [line(addonCopy.setupImportFailed, 'warn')] : [];
}

const line = (text: string, tone: StepTone = 'muted', extra: Partial<StepLineOut> = {}): StepLineOut => ({
  text,
  tone,
  ...extra,
});
const SKELETON: StepLineOut = { text: '', tone: 'muted', skeleton: true };
/** The second line a signed-in answer wraps to on a phone (rows 1 and 3: who and when, the PC and
 *  when it was seen), reserved there so the plate does not move when the reads answer. */
const SKELETON_PHONE: StepLineOut = { ...SKELETON, phoneOnly: true };

function act(id: StepActionId, label: string, kind: StepAction['kind'] = 'link-button'): StepAction {
  return { id, label, kind, primary: false };
}

const SIGN_IN = act('signin', addonCopy.setupSignInAction);
const INSTALL = act('install', addonCopy.setupInstallAddon);
const UPDATE = act('update', addonCopy.setupUpdateAddon);

/** The companion step not done. Windows and Mac: their own build. A system the page cannot name: both builds
 *  (the action's label is the group's name; ui draws one button per build). A phone or tablet: no
 *  download, the link to open on the computer instead. */
function companionTodo(platform: VisitorPlatform, text: string): StepRow {
  if (platform === 'mobile') {
    const lines = [line(text), line(addonCopy.setupCompanionMobile)];
    return row('companion', 'todo', lines, act('copylink', addonCopy.setupCopyLink, 'button'));
  }
  const own = downloadFor(platform);
  const label =
    own === null ? addonCopy.setupDownloadCompanion : addonCopy.setupCompanionDownload(own.platform);
  return row('companion', 'todo', [line(text)], act('download', label));
}

const TITLES: Record<StepId, string> = {
  signin: addonCopy.setupSignInTitle,
  companion: addonCopy.setupCompanionTitle,
  addon: addonCopy.setupAddonTitle,
};

function row(id: StepId, mark: StepMark, lines: StepLineOut[], action: StepAction | null): StepRow {
  const n = (STEP_ORDER.indexOf(id) + 1) as StepRow['n'];
  return { id, n, title: TITLES[id], mark, lines, action };
}

/** The addon step not done, with the install buttons. Before any companion has looked: what it does
 *  and that an installed one will be seen. After a look that found none: on which PC it is missing. */
function addonTodo(notOn: string | null): StepRow {
  const lines =
    notOn === null ? [line(addonCopy.setupAddonNote), line(addonCopy.setupAddonUnseen)] : [line(notOn)];
  return row('addon', 'todo', lines, INSTALL);
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
function anonymousRows(input: SetupInputs): StepRow[] {
  const addon =
    input.pointer === true
      ? row('addon', 'ok', [line(addonCopy.setupAddonDoneLine, 'ok')], null)
      : addonTodo(null);
  const companion = companionTodo(input.platform, addonCopy.setupCompanionNote);
  if (input.session === 'loading') {
    return [
      row('signin', 'ok', [SKELETON, SKELETON_PHONE], null),
      row('companion', 'todo', [SKELETON, SKELETON_PHONE], null),
      row('addon', 'todo', [SKELETON], null),
    ];
  }
  if (input.session === 'error') {
    const retry = primary(act('retry', characterListCopy.tryAgain, 'button'));
    return [
      row('signin', 'miss', [line(yourReportsPageCopy.sessionFailed, 'miss', { alert: true })], retry),
      companion,
      addon,
    ];
  }
  const signIn =
    input.signinFailed === true
      ? row('signin', 'miss', [line(addonCopy.setupSignInFailed, 'miss', { alert: true })], primary(SIGN_IN))
      : row('signin', 'next', [line(addonCopy.setupSignInBody)], primary(SIGN_IN));
  return [signIn, companion, addon];
}

/** Row 1 signed in: who, and whether the characters are here, coming, or missing. */
function signInRow(me: Me, input: SetupInputs, now: Date): StepRow {
  const signed = signInLine(me)?.text ?? '';
  if (me.characters.length > 0) {
    const imported =
      me.bnet_imported_at === undefined
        ? ''
        : ` · ${characterListCopy.updatedFrom(me.bnet_imported_at, now)}`;
    return row('signin', 'ok', [line(`${signed}${imported}`, 'ok'), ...importShortfall(me)], null);
  }
  if (me.bnet_import?.state === 'running') {
    return row('signin', 'busy', [line(signed, 'text'), line(addonCopy.setupImportWait)], null);
  }
  if (input.refreshFailed === true) {
    const retry = primary(act('refresh', characterListCopy.tryAgain, 'button'));
    return row(
      'signin',
      'miss',
      [line(signed, 'text'), line(characterListCopy.refreshError, 'miss', { alert: true })],
      retry,
    );
  }
  // A run that keyed nothing cannot say the account has no characters (ux-designer and
  // wow-player 2026-11-12): the shortfall line stands alone, the Korea line stays muted.
  const shortfall = importPending(me, now) ? [] : importShortfall(me);
  const wait = importPending(me, now) ? [line(addonCopy.setupImportWait, 'warn')] : [];
  const empty = shortfall.length > 0 ? [] : [line(characterListCopy.emptyLead, 'warn')];
  return row(
    'signin',
    'warn',
    [line(signed, 'text'), ...wait, ...shortfall, ...empty, line(characterListCopy.emptyKorea)],
    primary(act('refresh', characterListCopy.refreshFromBattlenet, 'button')),
  );
}

function companionRow(me: Me, input: SetupInputs, now: Date): StepRow {
  const status = companionLine(me, input.devices, now);
  if (status?.done === true) return row('companion', 'ok', [line(status.text, 'ok')], null);
  if (input.pairingCode != null && input.pairingCode !== '') {
    return row('companion', 'busy', [line(pairingCopy.waiting, 'muted', { code: input.pairingCode })], null);
  }
  return companionTodo(input.platform, status?.text ?? addonCopy.setupCompanionNote);
}

/** Every row of the plate for one state. */
export function setupSteps(input: SetupInputs): StepRow[] {
  const now = input.now ?? new Date();
  const me = input.me;
  if (input.session !== 'in' || me === null) return promoteFirst(anonymousRows(input));
  const first = signInRow(me, input, now);
  return promoteFirst([first, companionRow(me, input, now), addonRow(me, input, now)]);
}

/** The addon step for an account: done when a companion found it installed or its data arrived. */
function addonRow(me: Me, input: SetupInputs, now: Date): StepRow {
  const status = addonLine(me, input.devices, input.addonLatest ?? null, now);
  if (status === null || !status.done) return addonTodo(status?.scanned === true ? status.text : null);
  if (status.behind === undefined) return row('addon', 'ok', [line(status.text, 'ok')], null);
  // Behind the feed: still done (it is installed), with what is out and the stores to update from.
  return row('addon', 'ok', [line(status.text, 'ok'), line(status.behind, 'text')], UPDATE);
}

/** The session the page reads from `/v1/me`'s answer: a 401 (or no session) is signed out, anything else that fails is an error (B1). */
export function sessionOf(result: 'me' | 'unauthorized' | 'failed'): SessionState {
  if (result === 'me') return 'in';
  return result === 'unauthorized' ? 'out' : 'error';
}
