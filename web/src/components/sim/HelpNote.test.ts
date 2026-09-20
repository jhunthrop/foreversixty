// web/src/components/sim/HelpNote.test.ts
// Task 7. This project's vitest config hands `mount()` Svelte's server entry (lib/report/
// tree-sizes.ts's own note), so there is no client-side click to simulate here -- the same
// constraint SubstitutionChips.test.ts already works around with a static render. HelpNote
// forwards `open` to Disclosure's own bindable for exactly this: rendering both states by
// prop proves the ARIA wiring is correct in each one, and Playwright (sim-help.spec.ts, at
// 390x844) covers the actual tap, the second tap that hides it again, and Escape.
import { createRawSnippet } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { simCopy } from '../../lib/sim/copy';
import HelpNote from './HelpNote.svelte';

const BODY_TEXT = 'Every iteration draws its own fight length inside this band, the way real pulls do.';

function bodySnippet(text: string): ReturnType<typeof createRawSnippet> {
  return createRawSnippet(() => ({
    render: () => `<p data-testid="help-note-body">${text}</p>`,
  }));
}

function renderNote(open: boolean, disabled = false): string {
  const { body } = render(HelpNote, {
    props: {
      label: 'Fight length',
      id: 'sim-duration',
      open,
      disabled,
      children: bodySnippet(BODY_TEXT),
    },
  });
  return body;
}

describe('HelpNote', () => {
  it('names the trigger after the control it explains -- "What Fight length means" -- rather than a word shared by every note', () => {
    const body = renderNote(false);
    expect(body).toContain(simCopy.helpTrigger('Fight length'));
    expect(simCopy.helpTrigger('Fight length')).toBe('What Fight length means');
  });

  it('the trigger is a real button, never a title=', () => {
    const body = renderNote(false);
    expect(body).toMatch(/<button[^>]*type="button"[^>]*data-testid="sim-duration-help-trigger"/);
    expect(body).not.toContain('title=');
  });

  it('carries Disclosure’s own 44px floor rather than a smaller bespoke one', () => {
    const body = renderNote(false);
    expect(body).toMatch(/<button[^>]*class="[^"]*\bmin-h-11\b[^"]*\bmin-w-11\b[^"]*"/);
  });

  it('closed by default: aria-expanded is false, aria-controls names the panel, and the panel is absent from the page', () => {
    const body = renderNote(false);
    expect(body).toContain('aria-expanded="false"');
    expect(body).toContain('aria-controls="sim-duration-help-panel"');
    expect(body).not.toContain('help-note-body');
    expect(body).not.toContain(BODY_TEXT);
  });

  it('open: aria-expanded is true and the panel renders the note body at the id aria-controls named', () => {
    const body = renderNote(true);
    expect(body).toContain('aria-expanded="true"');
    expect(body).toContain('id="sim-duration-help-panel"');
    expect(body).toContain(BODY_TEXT);
  });

  it('forwards disabled to the trigger', () => {
    const enabled = renderNote(false, false);
    const disabled = renderNote(false, true);
    expect(enabled).not.toMatch(/<button[^>]*\sdisabled/);
    expect(disabled).toMatch(/<button[^>]*\sdisabled/);
  });

  it('accepts structured content, not only a plain sentence -- Fight style’s own dl of nine styles', () => {
    const dlSnippet = createRawSnippet(() => ({
      render: () => '<dl><div><dt>Patchwerk</dt><dd>One target, standing still.</dd></div></dl>',
    }));
    const { body } = render(HelpNote, {
      props: { label: 'Fight style', id: 'sim-style', open: true, children: dlSnippet },
    });
    expect(body).toContain('<dl>');
    expect(body).toContain('Patchwerk');
  });
});
