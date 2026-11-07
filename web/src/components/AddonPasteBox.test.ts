// web/src/components/AddonPasteBox.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import AddonPasteBox from './AddonPasteBox.svelte';
import { addonCopy } from '../lib/addon/copy';

describe('AddonPasteBox', () => {
  it('is a FormRow: the labelled export field and Load, with no decode result yet', () => {
    const { body } = render(AddonPasteBox, { props: { signedIn: null } });
    expect(body).toContain('data-testid="addon-paste-box"');
    expect(body).toContain(`aria-label="${addonCopy.setupTryExport}"`);
    expect(body).toContain(addonCopy.pasteFieldLabel);
    expect(body).toContain(`placeholder="${addonCopy.pastePlaceholder}"`);
    expect(body).toMatch(/<button[^>]*type="submit"[^>]*disabled/);
    expect(body).not.toContain('data-testid="addon-paste-save"');
    expect(body).not.toContain('data-testid="addon-paste-signin-hint"');
  });

  it('draws no card surface (the flat system)', () => {
    const { body } = render(AddonPasteBox, { props: { signedIn: true } });
    expect(body).not.toMatch(/bg-raised|rounded-panel|border-line-warm/);
    // Before any decode the status block (and its skeleton) is absent.
    expect(body).not.toContain('addon-paste-status-skeleton');
  });
});
