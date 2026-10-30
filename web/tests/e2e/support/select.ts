// web/tests/e2e/support/select.ts
// Driving the shared Select (design/specs/2026-10-30-controls-charts.md §A3). Its native <select>
// stays the source of truth for value and change at every pointer, but on a fine pointer the
// listbox hides it under a trigger once it loads after first paint. These helpers keep a spec
// about something else (a filter, a setting) independent of which control is showing; the
// dropdown's own behaviour is driven through the real UI in select.spec.ts.
import type { Locator, Page } from '@playwright/test';

type Choice = Parameters<Locator['selectOption']>[0];

/** Picks an option on the field's native select, whether or not the listbox has covered it. */
export async function choose(select: Locator, choice: Choice): Promise<void> {
  await select.selectOption(choice, { force: true });
}

/** The control a person sees and focuses for the field with this test id: the listbox trigger
 *  once enhanced, else the native select. */
export function fieldControl(scope: Page | Locator, testid: string): Locator {
  return scope.locator(`[data-testid="${testid}-trigger"], select[data-testid="${testid}"]:not([hidden])`);
}

/** The field's visible box (the wrapper the face is drawn in). */
export function fieldBox(scope: Page | Locator, testid: string): Locator {
  return scope.locator(`.sel:has(> select[data-testid="${testid}"])`);
}

/** Waits until the listbox has enhanced the field (fine pointer only). */
export async function enhanced(scope: Page | Locator, testid: string): Promise<Locator> {
  const trigger = scope.locator(`[data-testid="${testid}-trigger"]`);
  await trigger.waitFor();
  return trigger;
}

/** The field's native select by its label. The listbox trigger is labelled by the same label,
 *  so a bare getByLabel matches both once the field is enhanced. */
export function nativeSelect(scope: Page | Locator, label: string): Locator {
  return scope.getByLabel(label).and(scope.locator('select'));
}

/** The open list of the field whose select carries this test id. */
export function listOf(scope: Page | Locator, testid: string): Locator {
  return scope.locator(`.sel:has(> select[data-testid="${testid}"]) > .sel-pop`);
}
