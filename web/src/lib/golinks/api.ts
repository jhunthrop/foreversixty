// web/src/lib/golinks/api.ts
// The owner's /admin/links calls: every link with its funnel, and a new link. Both go through
// the account module's requestEnvelope (session cookie, CSRF header, the envelope's errors).
import { AccountError, requestEnvelope } from '../account/api';
import { API_BASE_URL } from '../planner/config';

export interface GoLinkInput {
  slug: string;
  target: string;
  label: string;
}

export interface GoLink extends GoLinkInput {
  created_at: string;
}

/** One link's funnel: people's clicks (bots apart), then the accounts it brought. */
export interface GoLinkFunnel extends GoLink {
  clicks: number;
  bot_clicks: number;
  sign_ins: number;
  linked: number;
  premium: number;
}

export const ADMIN_LINKS_PATH = '/v1/admin/golinks';

export async function listGoLinks(apiBase: string = API_BASE_URL): Promise<GoLinkFunnel[]> {
  const { data } = await requestEnvelope<{ links: GoLinkFunnel[] }>(ADMIN_LINKS_PATH, apiBase);
  return data?.links ?? [];
}

export async function createGoLink(input: GoLinkInput, apiBase: string = API_BASE_URL): Promise<GoLink> {
  const { data } = await requestEnvelope<GoLink>(ADMIN_LINKS_PATH, apiBase, { method: 'POST', body: input });
  if (data === null) throw new AccountError('The link was not saved; try again', 0);
  return data;
}
