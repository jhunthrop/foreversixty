// web/src/components/CharacterCard.test.ts
// The card's session-resolved states (signed-in, synced/not-synced) only ever render after
// `$effect` runs, which neither `svelte/server`'s `render()` nor this project's own
// `mount()` (hands Svelte's server entry, same constraint HelpNote.test.ts documents) ever
// fires -- so there is no harness in this repo that mounts the component past its initial,
// pre-session SSR shape. The two real, signed-in states (ux-designer/wow-player review
// round 1: the `Sample` pill must be absent in BOTH, not just one) are instead covered by
// Playwright screenshots (scratchpad/day3/shots/bis-build/03-*). This is the regression
// guard available here: the pill's markup and copy are gone from the component entirely,
// not conditioned on `synced`, so there is no code path left that could ever reintroduce it
// for a real character.
import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import CharacterCard from './CharacterCard.svelte';

describe('CharacterCard', () => {
  it('renders the loading Skeleton before the session resolves (SSR/pre-hydration)', () => {
    const { body } = render(CharacterCard, { props: { nextPath: '/bis/hunter/marksmanship' } });
    expect(body).toContain('data-testid="bis-character-card-loading"');
    expect(body).not.toContain('data-testid="bis-character-card"');
  });

  it('never references the Sample pill for a real character -- the mock’s own illustrative label, never a runtime state for either signed-in state', () => {
    const source = readFileSync(fileURLToPath(new URL('./CharacterCard.svelte', import.meta.url)), 'utf8');
    expect(source).not.toContain('pill-sample');
    expect(source).not.toContain('samplePillLabel');
  });

  it('has no Switch control: the header selector is the one way to switch character', () => {
    const source = readFileSync(fileURLToPath(new URL('./CharacterCard.svelte', import.meta.url)), 'utf8');
    expect(source).not.toContain('switchHref');
    expect(source).not.toContain('switchLabel');
    expect(source).not.toContain('character-card-switch');
  });
});
