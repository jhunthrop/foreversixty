// web/src/lib/account/signin-copy.ts
// The one "why sign in" reason line shared by /login (its own paragraph, outside the
// Account island) and /account's signed-out state (SignInPrompt's `line` prop) -- spec
// 2026-09-22 §2.5: the account page reuses the login page's own reason rather than
// inventing a second one.
export const accountSignInCopy = {
  reason:
    'An account is needed to upload logs, pair the companion and claim characters. Reading reports and rankings never needs one.',
} as const;
