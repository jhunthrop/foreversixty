// web/src/lib/subscribe.ts
// Maps POST /v1/subscribe responses to what the homepage says. Kept out of the component
// so the wording is unit-tested without a browser, the way lib/planner/share.ts is.

export const SUBSCRIBED = 'Check your email and confirm the address.';
export const RATE_LIMITED = 'Too many attempts from this connection; try again shortly.';
export const SUBSCRIBE_FAILED = 'That did not go through; try again.';
export const MAILTO_PROMPT = 'The subscribe service is not answering. Email the address below instead.';

export interface SubscribeMessage {
  kind: 'done' | 'error' | 'offline';
  message: string;
}

/**
 * `status` is the HTTP status, or 0 when the request never completed.
 * `apiMessage` is `error.message` from the envelope, when there is one.
 */
export function subscribeMessageFor(status: number, apiMessage: string | null): SubscribeMessage {
  if (status === 202) return { kind: 'done', message: SUBSCRIBED };
  if (status === 429) return { kind: 'error', message: RATE_LIMITED };
  if (status === 400) return { kind: 'error', message: apiMessage ?? SUBSCRIBE_FAILED };
  if (status === 0 || status >= 500) return { kind: 'offline', message: MAILTO_PROMPT };
  return { kind: 'error', message: SUBSCRIBE_FAILED };
}
