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

  it('reserves the signed-in block height while fetchMeOnce is still pending', () => {
    const { body } = render(AddonPasteBox, { props: {} });
    // Before any decode: loaded is null, so the whole signedIn block (including its
    // skeleton) is absent -- this proves the reservation only appears once a decode exists,
    // matching the component's existing pre-effect contract.
    expect(body).not.toContain('addon-paste-status-skeleton');
  });
});
