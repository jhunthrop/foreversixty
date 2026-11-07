// web/src/lib/account/signin-copy.ts
// The sign-in words shared by /login (account/AccountLogin.svelte) and /account's signed-out
// state (account/AccountOut.svelte): spec 2026-09-22 §2.5, the account page reuses the login
// page's own reason rather than inventing a second one. The button's label is
// myReportsCopy.signInAction, the one "Sign in with Battle.net" every sign-in button reads.
export const accountSignInCopy = {
  reason:
    'Sign in with Battle.net and your Forever characters come straight in, with your builds, sims and logs saved. Reading reports and rankings never needs an account.',
  /** /login's page title and its plate's name. */
  title: 'Sign in',
  /** Under the button on /login: what Battle.net shares, and what it never does. */
  privacy: 'Battle.net tells the site your BattleTag and your characters. It never sees a password here.',
  /** /login once a session exists: who is signed in, and the way back into the account. */
  signedInAs: (name: string): string => `Signed in as ${name}.`,
  goAccount: 'Your account',
} as const;
