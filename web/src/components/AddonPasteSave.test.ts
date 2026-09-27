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
    // The region has no source in the string, so the save still waits for it.
    const match = /<button[^>]*data-testid="addon-paste-save-button"[^>]*>/.exec(body);
    if (match === null) throw new Error('save button not rendered');
    expect(match[0]).toContain('disabled');
  });

  it('draws the fields as fields: each has a border', () => {
    const { body } = render(AddonPasteSave, { props: { signedIn: true, code: CODE } });
    for (const testid of ['addon-paste-name', 'addon-paste-region', 'addon-paste-ruleset']) {
      const tag = new RegExp(`<(?:input|select)[^>]*data-testid="${testid}"[^>]*>`).exec(body);
      if (tag === null) throw new Error(`${testid} not rendered`);
      expect(tag[0]).toMatch(/class="[^"]*\bborder\b/);
    }
  });

  it('disables the save button until every field is filled', () => {
    const { body } = render(AddonPasteSave, { props: { signedIn: true, code: CODE } });
    const match = /<button[^>]*data-testid="addon-paste-save-button"[^>]*>/.exec(body);
    if (match === null) throw new Error('save button not rendered');
    expect(match[0]).toContain('disabled');
  });
});
