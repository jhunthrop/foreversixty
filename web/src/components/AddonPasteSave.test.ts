// web/src/components/AddonPasteSave.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import AddonPasteSave from './AddonPasteSave.svelte';
import { addonCopy } from '../lib/addon/copy';

const CODE = 'FS1:1.60.1.69893:priest:human:0/0/0:';

describe('AddonPasteSave', () => {
  it('shows the name/region/ruleset fields and the save action when signed in', () => {
    const { body } = render(AddonPasteSave, { props: { signedIn: true, code: CODE } });
    expect(body).toContain('data-testid="addon-paste-save"');
    expect(body).toContain('data-testid="addon-paste-name"');
    expect(body).toContain('data-testid="addon-paste-region"');
    expect(body).toContain('data-testid="addon-paste-ruleset"');
    expect(body).toContain(addonCopy.pasteSaveAction);
    expect(body).not.toContain('data-testid="addon-paste-signin-hint"');
  });

  it('shows only the sign-in hint when signed out, no fields at all', () => {
    const { body } = render(AddonPasteSave, { props: { signedIn: false, code: CODE } });
    expect(body).toContain('data-testid="addon-paste-signin-hint"');
    expect(body).toContain(addonCopy.pasteSignInHint);
    expect(body).not.toContain('data-testid="addon-paste-save"');
  });

  it('starts from the name and the realm’s ruleset when the export names its character', () => {
    const { body } = render(AddonPasteSave, {
      props: { signedIn: true, code: CODE, character: { name: 'Bow Jackzon', realm: 'Classic Beta PvP' } },
    });
    expect(body).toContain('value="Bow Jackzon"');
    expect(body).toMatch(/<option value="pvp"[^>]*selected/);
    // The region has no source in the string, so it starts on its placeholder option.
    expect(body).toMatch(/<option value=""[^>]*selected[^>]*>Region</);
  });

  it('draws the fields as fields: the name a bordered input, region and ruleset the shared Select', () => {
    const { body } = render(AddonPasteSave, { props: { signedIn: true, code: CODE } });
    expect(body).toMatch(/<input[^>]*class="[^"]*\bborder\b[^"]*"[^>]*data-testid="addon-paste-name"/);
    for (const testid of ['addon-paste-region', 'addon-paste-ruleset']) {
      expect(body).toMatch(new RegExp(`<select class="sel-native"[^>]*data-testid="${testid}"`));
    }
  });

  it('keeps the save button live with empty fields: a press names what is missing, a dead button reads as broken', () => {
    const { body } = render(AddonPasteSave, { props: { signedIn: true, code: CODE } });
    const match = /<button[^>]*data-testid="addon-paste-save-button"[^>]*>/.exec(body);
    if (match === null) throw new Error('save button not rendered');
    expect(match[0]).not.toContain('disabled');
  });
});
