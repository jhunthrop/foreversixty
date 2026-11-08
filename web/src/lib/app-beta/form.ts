// web/src/lib/app-beta/form.ts
// What one press of "Sign up to test" comes to: the line under the field (invalid), the line
// under the button (failed), or the line that replaces the form (done). Pure over its sender so
// every outcome is tested without a DOM; ui TesterForm.svelte only draws it.
import { AccountError } from '../account/api';
import { signUpTester, type SignupStatus, type TesterSignup } from './api';
import { appBetaCopy as copy } from './copy';
import { isEmailShaped, type Platform } from './platform';

export type SignupOutcome =
  | { kind: 'done'; message: string }
  | { kind: 'invalid'; message: string }
  | { kind: 'failed'; message: string };

type Send = (signup: TesterSignup) => Promise<SignupStatus>;

const TOO_MANY = 429;
const UNPROCESSABLE = 422;

function failure(error: unknown): SignupOutcome {
  if (error instanceof AccountError && error.status === UNPROCESSABLE && error.fields?.email !== undefined) {
    return { kind: 'invalid', message: copy.invalidEmail };
  }
  if (error instanceof AccountError && error.status === TOO_MANY)
    return { kind: 'failed', message: copy.tooMany };
  return { kind: 'failed', message: copy.failed };
}

export async function submitSignup(
  rawEmail: string,
  platform: Platform,
  send: Send = signUpTester,
): Promise<SignupOutcome> {
  const email = rawEmail.trim();
  if (!isEmailShaped(email)) return { kind: 'invalid', message: copy.invalidEmail };
  try {
    const status = await send({ email, platform });
    return { kind: 'done', message: status === 'created' ? copy.created(platform) : copy.exists(platform) };
  } catch (error) {
    return failure(error);
  }
}
