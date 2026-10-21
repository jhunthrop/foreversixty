// web/src/lib/account/signin-copy.ts
// The one "why sign in" reason line shared by /login (its own paragraph, outside the
// Account island) and /account's signed-out state (SignInPrompt's `line` prop) -- spec
// 2026-09-22 §2.5: the account page reuses the login page's own reason rather than
// inventing a second one.
export const accountSignInCopy = {
  reason: 'Sign in to save your builds, sims and logs. Reading reports and rankings never needs an account.',
} as const;

/** "or use an email link": the alternative to the Battle.net button, worded once for every
 *  sign-in surface (EmailLinkAlternative.svelte). */
export const emailLinkCopy = {
  lead: 'or',
  link: 'use an email link',
  href: '/login',
} as const;
