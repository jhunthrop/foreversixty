// web/src/lib/guild/signup-save.ts
// The one raid sign-up save (component audit 2026-11-10 cluster 5; owner ruling: In, Late, Bench,
// Out everywhere). Every sign-up control goes through here and nowhere else calls the PUT: the
// raid page's strip and officer rows (`sendSignup`, the controller's live dep), the guild band
// and home's raid plaque (`pressSignup`: the band's raid after the answer, a lock or cancel race
// applied), and an officer's quick moves (`trySignup`: a failure the caller always shows).
import { GuildApiError, type GuildRaidView, type NextRaidView } from './api';
import { raidCopy } from './raid-copy';
import { putSignup, type RaidId, type SignupInput, type SignupResult } from './raids-api';

/** Sends one sign-up; rejects with the API's refusal. */
export function sendSignup(
  guildId: number,
  raidId: RaidId,
  key: string,
  body: SignupInput,
): Promise<SignupResult> {
  return putSignup(guildId, raidId, key, body);
}

/** The status line after a refused save (raids spec 4 "Races", 6.1 strip status). */
export function pressFailedLine(thrown: unknown): string {
  if (thrown instanceof GuildApiError && thrown.code === 'raid_locked') return raidCopy.races.locked;
  if (thrown instanceof GuildApiError && thrown.code === 'raid_cancelled') return raidCopy.races.cancelled;
  return raidCopy.stripStatus.failed;
}

/** A race line (lock, cancel) takes focus; the generic failure line is only announced. */
export function isRaceLine(line: string): boolean {
  return line !== raidCopy.stripStatus.failed;
}

/** The next-raid fields after a sign-up, from the PUT's raid: the API's own rule for `own` (the
 *  current character's sign-up, else the account's first, else none). */
export function nextRaidAfter(prev: NextRaidView, view: GuildRaidView, currentKey: string): NextRaidView {
  const signed = view.viewer.characters.filter((c) => c.state !== null);
  const own = signed.find((c) => c.character_key === currentKey) ?? signed[0];
  const { tanks, healers, dps, melee, ranged } = view.counts;
  const note = (key: string): string => view.signups.find((s) => s.character_key === key)?.note ?? '';
  return {
    ...prev,
    starts_at: view.starts_at,
    status: view.status === 'locked' ? 'locked' : 'planned',
    size: view.size,
    counts: { in: view.counts.in, late: view.counts.late, tanks, healers, dps, melee, ranged },
    own:
      own === undefined || own.state === null
        ? null
        : {
            character_key: own.character_key,
            state: own.state,
            role: own.role,
            note: note(own.character_key),
          },
  };
}

/** The next raid after a refused save: a lock locks it, a cancel drops it (`null`); `undefined`
 *  leaves it as it was. */
export function nextRaidAfterRefusal(prev: NextRaidView, thrown: unknown): NextRaidView | null | undefined {
  if (!(thrown instanceof GuildApiError)) return undefined;
  if (thrown.code === 'raid_locked') return { ...prev, status: 'locked' };
  if (thrown.code === 'raid_cancelled') return null;
  return undefined;
}

/** A press on a next-raid control: the raid it leaves, and on a refusal the line to show. */
export type PressOutcome =
  { ok: true; next: NextRaidView } | { ok: false; next: NextRaidView | null | undefined; line: string };

/** Saves one press on the next raid (the guild band, home's plaque). `currentKey` is the header
 *  character the band's `own` prefers. Never rejects. */
export async function pressSignup(
  guildId: number,
  prev: NextRaidView,
  key: string,
  body: SignupInput,
  currentKey: string = key,
  send: typeof sendSignup = sendSignup,
): Promise<PressOutcome> {
  try {
    const { raid } = await send(guildId, prev.id, key, body);
    return { ok: true, next: nextRaidAfter(prev, raid, currentKey) };
  } catch (thrown) {
    return { ok: false, next: nextRaidAfterRefusal(prev, thrown), line: pressFailedLine(thrown) };
  }
}

/** One sign-up for a raid by id (an officer's Move to In): the result, or the line to show.
 *  Never rejects. */
export async function trySignup(
  guildId: number,
  raidId: RaidId,
  key: string,
  body: SignupInput,
  send: typeof sendSignup = sendSignup,
): Promise<{ ok: true; result: SignupResult } | { ok: false; line: string }> {
  try {
    return { ok: true, result: await send(guildId, raidId, key, body) };
  } catch (thrown) {
    return { ok: false, line: pressFailedLine(thrown) };
  }
}
