// web/src/components/character/FactionCrest.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import FactionCrest from './FactionCrest.svelte';

describe('FactionCrest', () => {
  it('draws the flat Horde logo, ringed in the Horde bar colour', () => {
    const { body } = render(FactionCrest, { props: { faction: 'horde', testid: 't' } });
    expect(body).toContain('/icons/hd/faction/horde-logo-512.webp');
    expect(body).toContain('box-shadow: 0 0 0 2px #c0392b');
    expect(body).toContain('data-testid="t"');
  });

  it('draws the flat Alliance logo, ringed in the Alliance bar colour', () => {
    const { body } = render(FactionCrest, { props: { faction: 'alliance', testid: 't' } });
    expect(body).toContain('/icons/hd/faction/alliance-logo-512.webp');
    expect(body).toContain('box-shadow: 0 0 0 2px #2f6fd6');
  });

  it('renders nothing for a null faction -- never a neutral disc', () => {
    // Svelte's server renderer leaves its own empty {#if}-block comment markers in the
    // output even when the block is false -- asserting no <img> (the only element this
    // component ever emits) is the real "renders nothing" check, not an exact-empty body.
    const { body } = render(FactionCrest, { props: { faction: null } });
    expect(body).not.toContain('<img');
  });

  it('renders nothing for an undefined faction', () => {
    const { body } = render(FactionCrest, { props: { faction: undefined } });
    expect(body).not.toContain('<img');
  });

  it('is a pure decorative mark -- alt="" and aria-hidden, never a link', () => {
    const { body } = render(FactionCrest, { props: { faction: 'horde' } });
    expect(body).toContain('alt=""');
    expect(body).toContain('aria-hidden="true"');
    expect(body).not.toContain('<a ');
  });
});
