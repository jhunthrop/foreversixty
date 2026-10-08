// web/tests/e2e/planner-guide-link.spec.ts
// Live defect 2026-10-05: a guide's "Load this build" (`/planner?code=` for one class) opened
// by a signed-in visitor whose main is another class loaded the planner but not the talents.
// The signed-in main-class default switched class while the talent file was still loading,
// so the code no longer matched and was dropped. A link that names a build owns the planner.
import { expect, test } from '@playwright/test';
import { meAddonFixture } from '../../src/fixtures/me-addon';

// The fixture data build ships a trimmed warrior and hunter talent file, so this is the suite's
// known-good fixture Fury Warrior code (current-character.spec.ts), not a live guide's own
// code; the signed-in fixture's main is a Hunter, which reproduces the class mismatch.
const CODE = 'FS1:1.60.1.70291:warrior:orc:0/5530515/0:';

for (const signedIn of [false, true]) {
  test(`a guide build code loads its own class and talents ${signedIn ? 'signed in with a Hunter main' : 'signed out'}`, async ({
    page,
  }) => {
    await page.route('**/v1/me', (route) =>
      route.fulfill(
        signedIn
          ? {
              status: 200,
              contentType: 'application/json',
              body: JSON.stringify({ ok: true, data: meAddonFixture, error: null, request_id: 'e2e' }),
            }
          : {
              status: 401,
              contentType: 'application/json',
              body: JSON.stringify({ ok: false, data: null, error: null, request_id: 'e2e' }),
            },
      ),
    );
    await page.goto(`/planner?code=${encodeURIComponent(CODE)}`);
    await expect(page.getByLabel('Class')).toHaveValue('warrior');
    await expect(page.getByTestId('tree-points-164')).toHaveText('24');
    await expect(page.getByText('Talents loaded from a character', { exact: false })).toBeVisible();
  });
}
