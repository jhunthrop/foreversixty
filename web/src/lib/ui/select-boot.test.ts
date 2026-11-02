// @vitest-environment jsdom
// web/src/lib/ui/select-boot.test.ts
// Select.astro boots its static fields after first paint (src/lib/after-paint.ts), so a value
// chosen on the native <select> before the boot lands must still reach the face, and every
// later change must follow it.
import { afterEach, describe, expect, it } from 'vitest';
import { bootStaticFields } from './select-boot';

function staticField(): HTMLElement {
  document.body.innerHTML = `
    <div data-listbox data-static>
      <span class="sel-face"><span class="sel-label">Arms</span></span>
      <select>
        <option value="arms" selected>Arms</option>
        <option value="fury">Fury</option>
        <option value="prot">Protection</option>
      </select>
    </div>`;
  const field = document.querySelector<HTMLElement>('[data-listbox]');
  if (field === null) throw new Error('no field');
  return field;
}

const label = (field: HTMLElement): string | null | undefined =>
  field.querySelector('.sel-label')?.textContent;

describe('bootStaticFields', () => {
  afterEach(() => {
    document.body.innerHTML = '';
  });

  it('paints the face from a value picked before the boot', () => {
    const field = staticField();
    const select = field.querySelector('select');
    if (select === null) throw new Error('no select');
    select.value = 'fury';
    bootStaticFields();
    expect(label(field)).toBe('Fury');
  });

  it('follows every change after the boot', () => {
    const field = staticField();
    bootStaticFields();
    const select = field.querySelector('select');
    if (select === null) throw new Error('no select');
    select.value = 'prot';
    select.dispatchEvent(new Event('change', { bubbles: true }));
    expect(label(field)).toBe('Protection');
  });
});
