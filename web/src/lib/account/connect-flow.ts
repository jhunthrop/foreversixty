// web/src/lib/account/connect-flow.ts
// The pure half of /connect: reading the code out of the query, showing it the way the
// companion does, and turning a failed call into the state the page lands in.
import { AccountError } from './api';

export type ConnectState =
  | 'loading'
  | 'signedOut'
  | 'pending'
  | 'success'
  | 'cancelled'
  | 'expired'
  | 'used'
  | 'failed';

/** Upper-case, spaces and hyphens stripped: the API compares codes the same way. */
export function normalizeCode(raw: string | null | undefined): string {
  return (raw ?? '').replace(/[\s-]/g, '').toUpperCase();
}

/** `KQ4-7TM`: a hyphen after the third symbol. Anything that is not six symbols is shown as typed. */
export function displayCode(raw: string | null | undefined): string {
  const code = normalizeCode(raw);
  return code.length === 6 ? `${code.slice(0, 3)}-${code.slice(3)}` : code;
}

/** Where a failed lookup/approve/deny leaves the page. */
export function stateForError(error: unknown): ConnectState {
  if (!(error instanceof AccountError)) return 'failed';
  if (error.status === 404 || error.status === 410) return 'expired';
  if (error.status === 409) return 'used';
  if (error.status === 401 || error.status === 403) return 'signedOut';
  return 'failed';
}
