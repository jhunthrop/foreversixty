// web/src/components/sim/SimRunBlock.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../../lib/account/api';
import { landingCopy } from '../../lib/sim/landing-copy';
import SimRunBlock from './SimRunBlock.svelte';

const WITH_BUILD: MeCharacter = {
  key: 'us/pvp/reloadd',
  region: 'us',
  ruleset: 'pvp',
  name: 'Reloadd',
  class: 'Hunter',
  build: { source: 'blizzard', captured_at: '2026-09-24T00:00:00Z' },
};

const NO_BUILD: MeCharacter = {
  key: 'us/pvp/victorie',
  region: 'us',
  ruleset: 'pvp',
  name: 'Victorie',
  class: 'Warrior',
};

const PATH = { region: 'us', ruleset: 'pvp', slug: 'reloadd' } as const;

describe('SimRunBlock', () => {
  it('renders a Run sim button for a character with a build', () => {
    const { body } = render(SimRunBlock, {
      props: { character: WITH_BUILD, path: PATH, busy: false, onrun: () => {} },
    });
    expect(body).toContain('data-testid="sim-run-block-action"');
    expect(body).toContain(landingCopy.runSim);
    expect(body).not.toContain('data-testid="sim-run-block-paste"');
  });

  it('disables the button while busy', () => {
    const { body } = render(SimRunBlock, {
      props: { character: WITH_BUILD, path: PATH, busy: true, onrun: () => {} },
    });
    const match = /<button[^>]*data-testid="sim-run-block-action"[^>]*>/.exec(body);
    if (match === null) throw new Error('run button not found');
    expect(match[0]).toContain('disabled=""');
  });

  it('renders the no-build line and a Paste export link, no button, for a character with none', () => {
    const { body } = render(SimRunBlock, {
      props: { character: NO_BUILD, path: null, busy: false, onrun: () => {} },
    });
    expect(body).not.toContain('data-testid="sim-run-block-action"');
    expect(body).toContain(landingCopy.runBlockNoBuild('Victorie'));
    expect(body).toContain('data-testid="sim-run-block-paste"');
    expect(body).toContain(`href="${landingCopy.pasteExportHref}"`);
    expect(body).toContain(landingCopy.pasteExport);
  });
});
