// web/src/components/AddonPasteBox.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import AddonPasteBox from './AddonPasteBox.svelte';
import { addonCopy } from '../lib/addon/copy';

describe('AddonPasteBox', () => {
  it('renders the paste box at the #paste anchor, with no decode result yet', () => {
    const { body } = render(AddonPasteBox);
    expect(body).toContain('id="paste"');
    expect(body).toContain('data-testid="addon-paste-box"');
    expect(body).toContain(addonCopy.pasteTitle);
    expect(body).not.toContain('data-testid="addon-paste-save"');
    expect(body).not.toContain('data-testid="addon-paste-signin-hint"');
  });
});
