// web/src/components/sim/ExampleResultCard.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { landingCopy } from '../../lib/sim/landing-copy';
import ExampleResultCard from './ExampleResultCard.svelte';

describe('ExampleResultCard', () => {
  it('renders the static example, captioned Example, with no fetch of its own', () => {
    const { body } = render(ExampleResultCard, { props: {} });
    expect(body).toContain('data-testid="sim-example-result"');
    expect(body).toContain(landingCopy.exampleCaption);
    expect(body).toContain(landingCopy.exampleSpec);
    expect(body).toContain(landingCopy.exampleDps);
    expect(body).toContain(landingCopy.exampleHint);
  });
});
