// web/tests/e2e/planner-guide-link.spec.ts
// Live defect 2026-10-05: a guide's "Load this build" (`/planner?code=` for one class) opened
// by a signed-in visitor whose main is another class loaded the planner but not the talents.
// The signed-in main-class default switched class while the talent file was still loading,
// so the code no longer matched and was dropped. A link that names a build owns the planner.
// Bug B1 (spec 2026-11-08 4.3): a code that names no character is a link, so its note says so
// (`plannerCopy.linkedNote`) and never "Talents loaded from a character".
import { expect, test } from './support/test';
import { meAddonFixture } from '../../src/fixtures/me-addon';
import { plannerCopy } from '../../src/lib/planner/copy';
import { ACTIVE_BUILD } from './support/active-build';
import { expectClass } from './support/planner';

// The fixture data build ships a trimmed warrior and hunter talent file, so this is the suite's
// known-good fixture Fury Warrior code (current-character.spec.ts), not a live guide's own
// code; the signed-in fixture's main is a Hunter, which reproduces the class mismatch. Stamped
// with the active build, as a guide's own link is: a code from an older client build reads the
// older-build note instead (and is reconstructed, not linked).
const CODE = `FS1:${ACTIVE_BUILD}:warrior:orc:0/5530515/0:`;

for (const signedIn of [false, true]) {
  test(`a guide build code loads its own class and talents ${signedIn ? 'signed in with a Hunter main' : 'signed out'}`, async ({
    page,
    isMobile,
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
    await expectClass(page, 'warrior');
    await expect(page.getByTestId('tree-points-164')).toHaveText('24');
    // The long sentence from 768, the short one on a phone (Notice `short`).
    const note = page.getByTestId('planner-code-note');
    await expect(note).toContainText(plannerCopy.linkedNote);
    await expect(note).toContainText(plannerCopy.linkedNoteShort);
    await expect(
      note.getByText(isMobile ? plannerCopy.linkedNoteShort : plannerCopy.linkedNote, { exact: true }),
    ).toBeVisible();
    await expect(page.getByText('Talents loaded from a character')).toHaveCount(0);
  });
}
