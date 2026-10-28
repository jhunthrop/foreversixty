// web/src/lib/billing/claim-flow.ts
// The decisions behind the checkout card's in-place claim (spec 2026-10-27-subscribe-and-join
// §5.3, §9.1): which steps open a confirm, the words of each, what a call's outcome is, and
// which §9.1 refusal line an API refusal becomes. No DOM; the two API calls are injected.
//
// The claim routes answer by status (api/internal/guilds/handler.go), not by Err* value:
// 409 `conflict` = ErrAlreadyClaimed | ErrClaimPending | ErrAlreadyClaimsAnotherGuild;
// 429 `rate_limited` = ErrClaimRateLimited; 403 `forbidden` = ErrNotEligible |
// ErrNoBattleNetIdentity (claim) or ErrSameAccount | ErrNotEligible (confirm); confirm's 404
// `not_found` = ErrNoPendingClaim. A 409 is told apart by the step the re-read /v1/me reports.
import { checkoutGuildCopy as copy, claimErrorLine } from './checkout-guild-copy';
import type { SubscribeStep } from './contract';
import type { StepGuild } from './subscribe-step';

export type ClaimKind = 'claim' | 'start_claim' | 'confirm_claim';

const CLAIM_KINDS: readonly SubscribeStep[] = ['claim', 'start_claim', 'confirm_claim'];

export function claimKindOf(step: SubscribeStep | null): ClaimKind | null {
  return step !== null && CLAIM_KINDS.includes(step) ? (step as ClaimKind) : null;
}

export type ClaimGuild = Pick<StepGuild, 'id' | 'name' | 'claim'>;

function claimantName(guild: ClaimGuild): string {
  return guild.claim?.holder_name ?? copy.rank.officer;
}

export function openerLabel(kind: ClaimKind, guildName: string): string {
  if (kind === 'claim') return copy.claim.button(guildName);
  return kind === 'start_claim' ? copy.startClaim.button : copy.confirmClaim.button;
}

export interface ClaimConfirmWords {
  sentence: string;
  commit: string;
  busy: string;
}

export function confirmWords(kind: ClaimKind, guild: ClaimGuild): ClaimConfirmWords {
  switch (kind) {
    case 'claim':
      return {
        sentence: copy.confirm.claim(guild.name),
        commit: copy.confirm.yesClaim,
        busy: copy.busy.claiming,
      };
    case 'start_claim':
      return {
        sentence: copy.confirm.startClaim(guild.name),
        commit: copy.confirm.yesStartClaim,
        busy: copy.busy.starting,
      };
    case 'confirm_claim':
      return {
        sentence: copy.confirm.confirmClaim(claimantName(guild), guild.name),
        commit: copy.confirm.yesConfirmClaim,
        busy: copy.busy.confirming,
      };
  }
}

/** The two calls, injected: `claim` is `claimGuild` (master claim or officer start), `confirm` is `confirmClaim`. */
export interface ClaimIo {
  claim: (guildId: number) => Promise<{ status: 'confirmed' | 'pending' }>;
  confirm: (guildId: number) => Promise<unknown>;
}

export type ClaimOutcome = { ok: true; line: string } | { ok: false; code: string; rereads: boolean };

const STALE_CODES: readonly string[] = ['conflict', 'not_found'];

function codeOf(error: unknown): string {
  if (typeof error === 'object' && error !== null && 'code' in error) return String(error.code);
  return '';
}

/** Runs a confirm block's commit. A refusal carries the API's code; `conflict` and `not_found` mean /v1/me is stale. */
export async function performClaim(kind: ClaimKind, guild: ClaimGuild, io: ClaimIo): Promise<ClaimOutcome> {
  try {
    if (kind === 'confirm_claim') {
      await io.confirm(guild.id);
      return { ok: true, line: copy.done.confirmClaim(claimantName(guild), guild.name) };
    }
    const { status } = await io.claim(guild.id);
    return { ok: true, line: status === 'pending' ? copy.done.startClaim : copy.done.claim(guild.name) };
  } catch (error) {
    const code = codeOf(error);
    return { ok: false, code, rereads: STALE_CODES.includes(code) };
  }
}

const PENDING_STEPS: readonly SubscribeStep[] = ['claim_pending_yours', 'confirm_claim', 'claim_pending_ask'];

/** The refusal key of an API refusal; `stepAfter` is the step the re-read /v1/me gave the card. */
export function refusalKey(kind: ClaimKind, code: string, stepAfter: SubscribeStep | null): string {
  if (code === 'rate_limited') return 'rate_limited';
  if (code === 'forbidden') return kind === 'confirm_claim' ? 'same_account' : 'not_eligible';
  if (code === 'not_found') return 'already_claimed';
  if (code !== 'conflict') return '';
  if (stepAfter === 'claim_blocked') return 'holds_other';
  return stepAfter !== null && PENDING_STEPS.includes(stepAfter) ? 'claim_pending' : 'already_claimed';
}

export function refusalLine(
  kind: ClaimKind,
  code: string,
  stepAfter: SubscribeStep | null,
  guildName: string,
): string {
  return claimErrorLine(refusalKey(kind, code, stepAfter), guildName);
}

/** The in-place claim's state, owned by the shell: one confirm open across all cards. */
export interface ClaimControls {
  open: boolean;
  /** Another card's confirm is open: this opener is disabled. */
  locked: boolean;
  error: string | null;
  io: ClaimIo;
  onOpen: () => void;
  onCancel: () => void;
  onDone: (line: string) => void;
  onRefused: (code: string, rereads: boolean) => void;
}

/** The outcome line a claim left on the card (`status`) or a refusal that closed the block (`alert`). */
export interface CardNotice {
  text: string;
  role: 'status' | 'alert';
}
