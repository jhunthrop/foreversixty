// web/src/components/sim/SimSavePanel.test.ts
// Extracted from SimView.svelte (2026-09-26 layout pass, keeping that file under this
// lane's 800-line ceiling). This pins the one thing an SSR render can prove without a DOM:
// the open button's disabled state follows `result` the same way it always did.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { simCopy } from '../../lib/sim/copy';
import SimSavePanel from './SimSavePanel.svelte';

describe('SimSavePanel', () => {
  it('disables the save button with no result at all', () => {
    const { body } = render(SimSavePanel, {
      props: { result: null, reportTitle: '', onsave: async () => null },
    });
    const match = /<button[^>]*data-testid="sim-save-open"[^>]*>/.exec(body);
    if (match === null) throw new Error('save button not found');
    expect(match[0]).toContain('disabled=""');
    expect(body).not.toContain('data-testid="sim-save-disabled-note"');
  });

  it('disables the save button and states why for an aborted result', () => {
    const { body } = render(SimSavePanel, {
      props: {
        result: { aborted: true } as never,
        reportTitle: 'x',
        onsave: async () => null,
      },
    });
    const match = /<button[^>]*data-testid="sim-save-open"[^>]*>/.exec(body);
    if (match === null) throw new Error('save button not found');
    expect(match[0]).toContain('disabled=""');
    expect(body).toContain('data-testid="sim-save-disabled-note"');
    expect(body).toContain(simCopy.saveAbortedDisabled);
  });

  it('enables the save button for a finished, non-aborted result', () => {
    const { body } = render(SimSavePanel, {
      props: { result: {} as never, reportTitle: 'x', onsave: async () => null },
    });
    const match = /<button[^>]*data-testid="sim-save-open"[^>]*>/.exec(body);
    if (match === null) throw new Error('save button not found');
    expect(match[0]).not.toContain('disabled=""');
  });
});
