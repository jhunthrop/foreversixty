// web/src/lib/billing/api.ts
// Every browser call to the billing half of the API (spec section 4). Mirrors guild/api.ts's
// pattern exactly: the shared transport (account/api.ts's requestEnvelope), a module-local
// error class carrying the HTTP status so a caller can branch on 503 (billing not configured
// yet) or 409 (a guild already has the plan) without parsing a message string.
import { AccountError, requestEnvelope, type EnvelopeResult } from '../account/api';
import { API_BASE_URL } from '../planner/config';
import type { Interval, PlanKey } from './plans';

export const BILLING_API_FAILED = 'That did not work; try again';

export class BillingApiError extends Error {
  constructor(
    message: string,
    readonly status: number,
  ) {
    super(message);
    this.name = 'BillingApiError';
  }
}

async function call<T>(
  path: string,
  apiBase: string,
  init: { method?: string; body?: unknown } = {},
): Promise<T> {
  let result: EnvelopeResult<T>;
  try {
    result = await requestEnvelope<T>(path, apiBase, { ...init, failureMessage: BILLING_API_FAILED });
  } catch (error) {
    if (error instanceof AccountError) throw new BillingApiError(error.message, error.status);
    throw new BillingApiError(BILLING_API_FAILED, 0);
  }
  if (result.data === null) throw new BillingApiError(result.message ?? BILLING_API_FAILED, result.status);
  return result.data;
}

export interface CheckoutRequest {
  plan: PlanKey;
  interval: Interval;
  /** Required by the API when plan is 'guild'; omitted for a personal premium purchase. */
  guildId?: number;
  /** Only ever sent from the guild settings page's "Take over billing" button. */
  intent?: 'transfer';
}

export interface CheckoutResult {
  checkout_url: string;
}

export interface PortalResult {
  portal_url: string;
}

export function startCheckout(
  input: CheckoutRequest,
  apiBase: string = API_BASE_URL,
): Promise<CheckoutResult> {
  const body: Record<string, unknown> = { plan: input.plan, interval: input.interval };
  if (input.guildId !== undefined) body.guild_id = input.guildId;
  if (input.intent !== undefined) body.intent = input.intent;
  return call<CheckoutResult>('/v1/billing/checkout', apiBase, { method: 'POST', body });
}

/** guildId omitted asks for the caller's own personal billing portal. */
export function openPortal(
  guildId: number | undefined,
  apiBase: string = API_BASE_URL,
): Promise<PortalResult> {
  const body = guildId === undefined ? {} : { guild_id: guildId };
  return call<PortalResult>('/v1/billing/portal', apiBase, { method: 'POST', body });
}
