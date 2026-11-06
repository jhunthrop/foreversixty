// web/tests/e2e/support/sim-race.ts
// A logged fight records no race and the sim never guesses one (lib/sim/character.ts
// PENDING_RACE), so a sim opened from a fight first asks for it. Compare mode starts its run
// unpressed once the question is answered (SimView.svelte enterCompare).
import { expect, type Page } from '@playwright/test';
import { choose } from './select';

/** Asserts the race question is up and nothing has run, then answers it with the first race offered. */
export async function answerRaceQuestion(page: Page): Promise<void> {
  const picker = page.getByTestId('sim-race-picker');
  await expect(picker).toBeAttached();
  await expect(page.getByTestId('compare-headline')).toHaveCount(0);
  const race = await picker.evaluate(
    (select: HTMLSelectElement) => [...select.options].find((option) => option.value !== '')?.value ?? '',
  );
  expect(race).not.toBe('');
  await choose(picker, race);
}
