// web/tests/e2e/login-next-redirect.spec.ts
// Covers the /login?next=... -> "Sign in with Battle.net" link round-trip. login.astro and
// Account.svelte's mode === 'login' branch resolve a real ?next= query param client-side (a
// $effect reading window.location.search, since this site is a fully static build with no
// per-request Astro frontmatter -- see Account.svelte's own comment above `resolvedNext`).
// tests/e2e/auth.spec.ts (not owned by this lane) only covers the no-next-param default case;
// this file is the only coverage that a real ?next= value actually survives into the link's
// href, and that safeNextPath's protection against an unsafe value holds end-to-end through
// the real component, not just in safe-next.ts's own unit tests.
import { expect, test } from '@playwright/test';

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});

test.beforeEach(async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: null, request_id: 'r' }, 401)),
  );
});

test('a safe ?next= value survives into the Battle.net link, post-hydration', async ({ page }) => {
  await page.goto('/login?next=%2Fpremium%2Fcheckout%3Fplan%3Dpremium%26interval%3Dmonthly');

  // toHaveAttribute auto-retries, so this asserts the post-hydration value the $effect
  // writes, not the build-time-frozen default baked into the static HTML.
  await expect(page.getByTestId('battlenet')).toHaveAttribute(
    'href',
    /next=%2Fpremium%2Fcheckout%3Fplan%3Dpremium%26interval%3Dmonthly$/,
  );
});

test('an unsafe protocol-relative ?next= value is rejected and falls back to the hub', async ({ page }) => {
  await page.goto('/login?next=//evil.example/x');

  // Every sign-in link's default landed on the hub, not /logs, once this lane's "sign-in
  // defaults" change shipped (login.astro's own `next="/account?signed_in=1"` fallback) --
  // safeNextPath's rejection falls back to whatever `next` the caller passed in, so this
  // now asserts that new default rather than the old /logs one.
  const battlenet = page.getByTestId('battlenet');
  await expect(battlenet).toHaveAttribute('href', /next=%2Faccount%3Fsigned_in%3D1$/);
  await expect(battlenet).not.toHaveAttribute('href', /evil\.example/);
});
