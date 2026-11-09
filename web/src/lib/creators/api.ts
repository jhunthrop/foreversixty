// web/src/lib/creators/api.ts
// The creator program's calls (api/openapi.yaml, /v1/admin/creators and /v1/creators). All go
// through the account module's requestEnvelope (session cookie, CSRF header, the envelope's errors).
import { AccountError, requestEnvelope } from '../account/api';
import { API_BASE_URL } from '../planner/config';
import type {
  ApplicationInput,
  ApplicationStatus,
  Creator,
  CreatorInput,
  CreatorMe,
  CreatorRequest,
  Statement,
  StatementsPage,
  StripeStatus,
} from './model';

const ADMIN = '/v1/admin/creators';
const MINE = '/v1/creators';

export interface CreatorsList {
  creators: Creator[];
  stripe_enabled: boolean;
  payouts_enabled: boolean;
  connect_denied: boolean;
}

export interface MyApplicationRead {
  status: ApplicationStatus | null;
  created_at: string | null;
  creator: boolean;
  /** The day a declined account may apply again, while it is ahead. */
  reapply_on: string | null;
}

/** The data a call must answer; a 2xx with none is an error the page shows. */
async function data<T>(
  path: string,
  init: { method?: string; body?: unknown } = {},
  apiBase = API_BASE_URL,
): Promise<T> {
  const { data: value } = await requestEnvelope<T>(path, apiBase, init);
  if (value === null) throw new AccountError('That didn’t work; try again', 0);
  return value;
}

export const listCreators = (): Promise<CreatorsList> => data<CreatorsList>(ADMIN);
export const addCreator = (input: CreatorInput): Promise<Creator> =>
  data<Creator>(ADMIN, { method: 'POST', body: input });
export const editCreator = (id: number, input: Omit<CreatorInput, 'slug' | 'code'>): Promise<Creator> =>
  data<Creator>(`${ADMIN}/${id}`, { method: 'PATCH', body: input });
export const makeInvite = async (id: number): Promise<string> =>
  (await data<{ url: string }>(`${ADMIN}/${id}/invite`, { method: 'POST' })).url;
export const refreshStripe = async (id: number): Promise<StripeStatus> =>
  (await data<{ stripe_status: StripeStatus }>(`${ADMIN}/${id}/stripe-status`, { method: 'POST' }))
    .stripe_status;
/** A month's statements; with no month, the API's own default, the last closed month. */
export const listStatements = (month = ''): Promise<StatementsPage> =>
  data<StatementsPage>(`${ADMIN}/statements${month === '' ? '' : `?month=${encodeURIComponent(month)}`}`);
export const approveStatement = (id: number, month: string): Promise<Statement> =>
  data<Statement>(`${ADMIN}/${id}/statements/${encodeURIComponent(month)}/approve`, { method: 'POST' });
/** The platform's available USD balance, in cents (the Approve dialog). */
export const payoutBalance = async (): Promise<number> =>
  (await data<{ available_cents: number }>(`${ADMIN}/balance`)).available_cents;
export const listRequests = async (): Promise<CreatorRequest[]> =>
  (await data<{ requests: CreatorRequest[] }>(`${ADMIN}/requests`)).requests;
export const approveRequest = (id: number, input: CreatorInput): Promise<Creator> =>
  data<Creator>(`${ADMIN}/requests/${id}/approve`, { method: 'POST', body: input });
export async function declineRequest(id: number, apiBase: string = API_BASE_URL): Promise<void> {
  await requestEnvelope(`${ADMIN}/requests/${id}/decline`, apiBase, { method: 'POST' });
}

export const invitePreview = async (token: string): Promise<string> =>
  (await data<{ name: string }>(`${MINE}/invite/preview`, { method: 'POST', body: { token } })).name;
export const acceptInvite = async (token: string): Promise<string> =>
  (await data<{ name: string }>(`${MINE}/invite`, { method: 'POST', body: { token } })).name;
export const fetchCreatorMe = (sync = false): Promise<CreatorMe> =>
  data<CreatorMe>(`${MINE}/me${sync ? '?sync=1' : ''}`);
export const onboardingLink = async (): Promise<string> =>
  (await data<{ url: string }>(`${MINE}/me/onboarding`, { method: 'POST' })).url;
export const dashboardLink = async (): Promise<string> =>
  (await data<{ url: string }>(`${MINE}/me/dashboard`, { method: 'POST' })).url;
export const fetchApplication = (): Promise<MyApplicationRead> =>
  data<MyApplicationRead>(`${MINE}/application`);
export const sendApplication = (input: ApplicationInput): Promise<{ status: ApplicationStatus }> =>
  data<{ status: ApplicationStatus }>(`${MINE}/application`, { method: 'POST', body: input });
