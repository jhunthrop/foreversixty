<!-- web/src/components/planner/PlannerHeaderBand.svelte -->
<!-- Fix round 5 (owner: the live planner did not match the mock, `day3/shots/boards/
     Planner.png`, structurally): the page's own header BAND -- full-bleed, edge to edge,
     the same shape the BiS page's `ArtPanel` + its own `.header-row` already ship (that
     page's `<main>` carries no max-width at all, so `ArtPanel` is full-bleed simply by
     being `width: 100%` inside it; `/planner`'s own `<main>` is a constrained
     `max-w-[1344px]` column instead, so this band breaks out of it with the standard
     `width: 100vw; margin-left: calc(50% - 50vw)` trick rather than restructuring
     `planner.astro`'s own page shell).

     Previously: `PlannerHeader` was its own bordered, rounded, `bg-raised` BOX sitting
     inside the page's gutters, with `PlannerCharacterCard` a second, separate box beside
     it (a flex row) -- the band's own background, the full-bleed edges and the two
     columns sharing one background never existed. This component is the fix: it owns the
     band's background (flat `--color-raised`, the night-to-`--bg` gradient over it --
     spec §13 keeps the dominant-tree art itself off, `PlannerHeader.svelte`'s own `headerArtSrc`/
     `HEADER_ART_ENABLED` constant, until it can paint from first byte without regressing
     Lighthouse's TBT again), the inner content wrapper (max-width 1344px, centred, the
     page's own 48px gutters from 1024px, 18px below it -- the exact figures the live
     captures were checked against at 1024/1280/1440/1920), and the `1fr 440px` grid (24px
     gap, bottom-aligned) that puts `PlannerCharacterCard` inside the band as its own right
     column, the same shape `ClassHeader.astro`'s own `.header-row` uses on the BiS page.

     The grid collapses to one column when there is nothing in the right column to show:
     `PlannerCharacterCard` renders nothing at all for a confirmed-signed-out visitor (its
     own spec §4.B rule, unchanged by this round), and reserving 440px of horizontal width
     for an empty column would be the same "dead gap for the common, anonymous visit" fix
     round 4 already ruled out for its own vertical reservation. Reads the same character-
     card state the card itself reads (`createCharacterCardState`, cache-scoped on the
     `/v1/me` query -- a second caller costs no second fetch, the same sharing that
     component's own header comment already documents) purely to decide the column count. -->
<script lang="ts">
  import { createCharacterCardState } from '../../lib/bis/character-card-state.svelte';
  import { dataUrl } from '../../lib/planner/load';
  import type { LiveDps } from '../../lib/planner/live-dps.svelte';
  import type { LiveGate } from '../../lib/planner/live-gate';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import PlannerCharacterCard from './PlannerCharacterCard.svelte';
  import PlannerHeader from './PlannerHeader.svelte';

  let {
    store,
    live,
    gate,
    onshowdps,
  }: {
    store: PlannerStore;
    live: LiveDps;
    gate: LiveGate;
    onshowdps: () => void;
  } = $props();

  const cardState = createCharacterCardState();
  const hasCard = $derived(
    cardState.status === 'loading' || cardState.status === 'failed' || cardState.character !== null,
  );

  /** The dominant tree's own index -- the same "most points, ties to the first tree"
   *  inference the h1 suffix (`PlannerHeader.svelte`) already uses, read here for the
   *  background art instead of the spec label. Moved here from `PlannerHeader.svelte`
   *  (fix round 5): the art has to cover the whole band, both grid columns, not just the
   *  header's own left one, so it can only ever live at this level. */
  const dominantTreeIndex = $derived.by(() => {
    const split = store.split;
    let best = 0;
    for (let i = 1; i < split.length; i += 1) {
      if (split[i] > split[best]) best = i;
    }
    return best;
  });

  /** Spec §13 (2026-10-02): the band ships flat. Art painted after talent data resolves is
   *  always the page's largest paint and lands at ~3.3 s under the audit's throttling,
   *  which put the merged build at 0.89 against the 0.90 performance floor. The art comes
   *  back when the planner can paint it from the first byte (per-class static routes or a
   *  server-rendered class); flipping this constant is the whole switch. */
  const HEADER_ART_ENABLED = false;
  const headerArtSrc = $derived.by(() => {
    if (!HEADER_ART_ENABLED) return null;
    const tree = store.talentIndex?.trees[dominantTreeIndex];
    return tree ? dataUrl(store.treeVersion, `trees/${tree.background}.webp`) : null;
  });
</script>

<div class="planner-header-band" data-testid="planner-header-band">
  <!-- A CSS background-image, not an <img>: a nicer fit for a decorative, `aria-hidden`
       backdrop over the whole band -- see `PlannerHeader.svelte`'s own, now-historical
       comment on why this is not a way to dodge LCP candidacy (it isn't; Chrome counts a
       CSS background the moment it is visibly painted) and why the swap is instant, no
       opacity transition, once the art is back in use. -->
  <div
    class="planner-header-band-art"
    class:planner-header-band-art-visible={headerArtSrc !== null}
    style={headerArtSrc !== null ? `background-image:url(${headerArtSrc})` : undefined}
    aria-hidden="true"
  ></div>
  <div class="planner-header-band-gradient" aria-hidden="true"></div>
  <div class="planner-header-band-inner">
    <div class="planner-header-band-grid" class:planner-header-band-grid-split={hasCard}>
      <PlannerHeader {store} {live} {gate} {onshowdps} />
      <PlannerCharacterCard />
    </div>
  </div>
</div>

<style>
  .planner-header-band {
    position: relative;
    width: 100vw;
    margin-left: calc(50% - 50vw);
    overflow: hidden;
    /* Flat `--color-raised` floor (spec §13): the dominant-tree art stays off until it can
       paint from first byte without becoming the page's own LCP/TBT regression again --
       see this file's own `HEADER_ART_ENABLED` comment above. */
    background: var(--color-raised);
  }
  .planner-header-band-art {
    position: absolute;
    inset: 0;
    background-size: cover;
    background-position: center;
    /* Flat until the data resolves and (if the art is ever re-enabled) names the dominant
       tree's own background: opacity 0 so the element paints nothing, the flat
       `--color-raised` band behind it the only "floor" at first paint. No `transition` --
       measured, an animated fade only moved the LCP timestamp later for no candidacy
       benefit (see git history on `PlannerHeader.svelte`'s own prior copy of this note). */
    opacity: 0;
  }
  .planner-header-band-art-visible {
    opacity: 1;
  }
  .planner-header-band-gradient {
    position: absolute;
    inset: 0;
    /* Night-to-bg gradient (unchanged from the header's own prior recipe): reads as the
       page's header band, not a card -- no border, no radius, no inset box. */
    background: linear-gradient(180deg, var(--color-night) 0%, var(--color-bg) 100%);
    opacity: 0.5;
  }
  .planner-header-band-inner {
    position: relative;
    max-width: 1344px;
    margin: 0 auto;
    padding: 0 18px;
  }
  @media (min-width: 1024px) {
    .planner-header-band-inner {
      padding: 0 48px;
    }
  }
  .planner-header-band-grid {
    display: grid;
    grid-template-columns: 1fr;
    gap: 16px;
  }
  @media (min-width: 1024px) {
    .planner-header-band-grid-split {
      grid-template-columns: 1fr 440px;
      gap: 24px;
      align-items: end;
    }
  }
</style>
