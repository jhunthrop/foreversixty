// web/src/components/sim/tools/SaveSimForm.test.ts
// Task 5 (newcomer MAJOR, review.md:360-363): the Save button's disabled explanation used to
// live only in a `title`, invisible on a phone. A static render proves the fix without a
// browser: the note is either in the markup or it is not, and producing it never involves a
// hover -- the same static-render technique SubstitutionChips.test.ts already uses for the
// same class of question.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { simCopy } from '../../../lib/sim/copy';
import SaveSimForm from './SaveSimForm.svelte';

function renderForm(canSave: boolean): string {
  const { body } = render(SaveSimForm, {
    props: { onsave: async () => null, titleFor: 'My build', canSave },
  });
  return body;
}

describe('SaveSimForm', () => {
  it('explains a disabled Save button as visible text, not a hover title', () => {
    const body = renderForm(false);
    expect(body).toContain(simCopy.saveAbortedDisabled);
    expect(body).not.toContain(`title="${simCopy.saveAbortedDisabled}"`);
  });

  it('renders no disabled note when the run can be saved', () => {
    const body = renderForm(true);
    expect(body).not.toContain(simCopy.saveAbortedDisabled);
  });
});
