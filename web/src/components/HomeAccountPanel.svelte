<!-- web/src/components/HomeAccountPanel.svelte -->
<!-- The home page's signed-in panel (home signed-in panel spec 2026-10-10, merged into the shared
     character header by design/specs/2026-10-11-character-header.md): the shared character header
     (who, "Upgrades in your band" as two figures, and the one button down to "Your upgrades"),
     then the Talents and Simulator cards as one two-card row under it. The header selector is the
     one way to change character; the header follows it in place.

     While the session is loading the header draws its skeleton (only for a session-hinted visit)
     and a signed-out visitor sees nothing here, so the sky band's own server-rendered signed-out
     hero (index.astro), which keeps its own sign-in, stays exactly where it is -- the same "server
     shell first, island swaps in place" trick SessionNav.svelte's header link uses, so the hero
     never shows two competing versions. Both this island's root and index.astro's signed-out
     block share `[grid-area:1/1]` in a shared grid wrapper, so once this mounts signed-in it
     visually occludes the signed-out block (a solid background over the same cell) instead of the
     two stacking and reflowing the page underneath. -->
<script lang="ts">
  import { mount, unmount, type Component } from 'svelte';
  import CharacterHeader from './character/CharacterHeader.svelte';
  import { createCharacterCardState } from '../lib/bis/character-card-state.svelte';
  import { homeAnswerSpec, homeAction } from '../lib/home/header-answer';
  import { HOME_SIGNED_OUT_ATTR, HOME_SIGNED_OUT_ID } from '../lib/home-panel-copy';

  // One session read, shared with the header: the island that decides which character the page
  // speaks to is the same one the header draws.
  const session = createCharacterCardState();
  const me = $derived(session.me);
  const ready = $derived(session.status === 'ready' || session.status === 'failed');
  const hero = $derived(session.character);
  const signedIn = $derived(ready && me !== null && hero !== null);

  /**
   * The grid-overlay CLS trick (this component's root and index.astro's signed-out block
   * share the same `[grid-area:1/1]` cell) only ever covers the signed-out row visually --
   * it stays mounted underneath, so without this a signed-in keyboard/screen-reader user
   * could still tab to, or hear, a duplicate "Sign in with Battle.net" link sitting behind
   * the visible strip. When the session hint was wrong (nothing signed in), the hint-only
   * CSS that collapses the signed-out hero's cells is switched off through
   * `HOME_SIGNED_OUT_ATTR` on <html>, restoring the two-cell signed-out hero.
   */
  $effect(() => {
    const signedOut = document.getElementById(HOME_SIGNED_OUT_ID);
    if (signedOut === null) return;
    if (signedIn) {
      signedOut.setAttribute('inert', '');
      signedOut.setAttribute('aria-hidden', 'true');
    } else {
      signedOut.removeAttribute('inert');
      signedOut.removeAttribute('aria-hidden');
      if (ready) signedOut.removeAttribute('data-session-hide');
    }
    document.documentElement.toggleAttribute(HOME_SIGNED_OUT_ATTR, ready && !signedIn);
  });

  /**
   * Mounts one of the home page's lazily imported islands into its slot once signed in, and
   * unmounts it when the character changes or the session ends. A dynamic import, so a
   * signed-out visitor's browser never downloads the module or its fetches (review round 3,
   * "islands" item).
   */
  function mountInSlot<Props extends Record<string, unknown>>(
    slotTestid: string,
    load: () => Promise<{ default: Component<Props> }>,
    props: Props,
  ): (() => void) | undefined {
    const slot = document.querySelector<HTMLElement>(`[data-testid="${slotTestid}"]`);
    if (slot === null) return undefined;
    let island: Record<string, unknown> | null = null;
    let cancelled = false;
    void load().then(({ default: component }) => {
      if (cancelled) return;
      island = mount(component, { target: slot, props });
    });
    return () => {
      cancelled = true;
      if (island !== null) void unmount(island);
    };
  }

  /** The Talents and Simulator cards, mounted into this panel's own slot below the header. */
  $effect(() => {
    if (!signedIn || hero === null) return;
    return mountInSlot('home-hero-cards-slot', () => import('./character/HomeHeroCards.svelte'), { hero });
  });

  /** §3.B.3's "Your upgrades" table, mounted into `index.astro`'s own `home-upgrades-slot`. */
  $effect(() => {
    if (!signedIn || hero === null) return;
    return mountInSlot('home-upgrades-slot', () => import('./character/HomeUpgradesPanel.svelte'), { hero });
  });
</script>

<!-- The root always renders, even empty: the island hydrates with client:idle-after-load, and an
     observer needs a box to see. Empty and pointer-events-none, it occludes nothing until signed in. -->
<div
  class="relative flex flex-col gap-4 [grid-area:1/1] md:gap-[22px]"
  class:pointer-events-none={!signedIn}
  data-testid={signedIn ? 'home-account-panel' : undefined}
>
  <CharacterHeader
    source={session}
    action={homeAction}
    answer={homeAnswerSpec}
    whenSignedOut="none"
    nameHeading
  />
  <!-- The Talents and Simulator cards render from HomeHeroCards.svelte, mounted above once
       signed in -- not this island's own static markup, so a signed-out visitor's browser never
       downloads the Simulator card's fetch or its dependencies. The slot reserves the card
       row's height (a session-hinted visit only) so the mount does not shift the table beneath.
       Re-measured for stage 2, where both became Cards with a 48 px compact head and a 56 px
       answer box in every state (HomeHeroCards.svelte): the row is 402 px stacked below 640 and
       195 px side by side from 640 up. The old 236/112 let the row grow on mount, and a header
       switch remounting it moved the upgrades table 120 px (home-panel). -->
  <div
    class={['ch-hint-only min-w-0', (!ready || signedIn) && 'min-h-[402px] sm:min-h-[195px]']}
    data-testid="home-hero-cards-slot"
  ></div>
</div>
