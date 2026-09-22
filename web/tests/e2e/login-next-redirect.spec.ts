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

test('an unsafe protocol-relative ?next= value is rejected and falls back to /logs', async ({ page }) => {
  await page.goto('/login?next=//evil.example/x');

  const battlenet = page.getByTestId('battlenet');
  await expect(battlenet).toHaveAttribute('href', /next=%2Flogs$/);
  await expect(battlenet).not.toHaveAttribute('href', /evil\.example/);
});
