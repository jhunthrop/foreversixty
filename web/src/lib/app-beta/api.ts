// web/src/lib/app-beta/api.ts
// The two calls behind the app beta (api/internal/apptesters): the public signup on /app and the
// owner's list on /admin/testers. Both go through requestEnvelope. The signup sends no cookie
// ('omit'): it is not tied to who asks, and a request with no session is outside the API's
// double-submit CSRF check by construction (there is no ambient credential to ride on).
import { requestEnvelope } from '../account/api';
import { API_BASE_URL } from '../planner/config';
import type { Platform } from './platform';

export const APP_TESTERS_PATH = '/v1/app-testers';
export const ADMIN_APP_TESTERS_PATH = '/v1/admin/app-testers';

export interface TesterSignup {
  email: string;
  platform: Platform;
}

export interface AppTester {
  email: string;
  platform: Platform;
  created_at: string;
}

/** Signs the address up; throws AccountError (422 with fields, 429, 5xx, offline). The API
 *  answers a new address and one already on the list alike, so there is nothing to return. */
export async function signUpTester(signup: TesterSignup, apiBase: string = API_BASE_URL): Promise<void> {
  await requestEnvelope<{ status: 'received' }>(APP_TESTERS_PATH, apiBase, {
    method: 'POST',
    body: signup,
    credentials: 'omit',
  });
}

/** Every signup, newest first (the owner only: 401 and 403 throw). */
export async function listAppTesters(apiBase: string = API_BASE_URL): Promise<AppTester[]> {
  const { data } = await requestEnvelope<{ testers: AppTester[] }>(ADMIN_APP_TESTERS_PATH, apiBase);
  return data?.testers ?? [];
}
