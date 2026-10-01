<!-- web/src/components/planner/PlannerHeader.svelte -->
<!-- Rebuild spec §4.A: the planner's own class-page header -- eyebrow, 64px crest, h1 with
     the always-present level/race suffix (plus a spec clause once one tree leads), and the
     race select. Lives inside the Planner island (not a `.astro` sibling like the BiS/home
     rebuild's `ClassHeader.astro`) because every one of these facts -- the class, the race,
     the level, the leading spec -- is the island's own reactive store, not build-time
     frontmatter; `ClassHeader.astro`'s own markup is server-rendered once per faction and
     has no runtime state to read at all (see that component's header comment).

     Reuses `ArtPanel`'s current look (a flat `bg-raised` panel, §4.A/§7: "ships flat until
     the art lane delivers the per-class tree art") directly as a Tailwind class on this
     component's own root, rather than importing `ArtPanel.astro` -- a Svelte component
     cannot import an `.astro` file at all (the same constraint `CharacterCard.svelte`
     already documents for `ClassCrest`/`FactionMark`), and `ArtPanel` has no behaviour
     beyond that one background rule to duplicate. Same reasoning for the crest: inlines the
     identical `<img>` `ClassCrest.astro` renders, from the same `classCrestSrc` path
     function, rather than forking a second crest component. -->
<script lang="ts">
  import { classCrestSrc } from '../../lib/class-crest';
  import { classColorVar } from '../../lib/report/format';
  import { levelReached } from '../../lib/planner/derive';
  import { plannerHeaderCopy } from '../../lib/planner/copy';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { specOf } from '../../lib/sim/character';
  import { specLabel } from '../../lib/sim/spec-label';

  let { store }: { store: PlannerStore } = $props();

  const color = $derived(classColorVar(store.classRow?.name));

  /** True once one tree holds strictly more points than each of the other two --
   *  `specKeyFor`'s own tie-break, reused unchanged. A tie, including 0/0/0, is false: the
   *  h1's third clause only ever appears once a build genuinely leans one way. */
  const hasLeadingTree = $derived.by(() => {
    const split = store.split;
    if (split.length === 0) return false;
    const max = Math.max(...split);
    return max > 0 && split.filter((points) => points === max).length === 1;
  });

  const specSuffix = $derived(
    hasLeadingTree && store.talentIndex !== null
      ? ` · ${specLabel(specOf(store.talentIndex, [...store.order]))}`
      : '',
  );
</script>

<div class="planner-header bg-raised" data-testid="planner-header">
  <p class="label planner-header-eyebrow">
    <i class="planner-header-eyebrow-rule" aria-hidden="true"></i>
    {plannerHeaderCopy.eyebrow}
  </p>
  <div class="planner-header-row">
    <div class="planner-header-identity">
      <img
        src={classCrestSrc(store.classSlug)}
        alt=""
        width="64"
        height="64"
        loading="eager"
        decoding="async"
        class="planner-header-crest"
        style={`--c:${color}`}
        data-testid="planner-header-crest"
      />
      <h1 class="planner-header-h1" style={`color:${color}`} data-testid="planner-header-h1">
        {plannerHeaderCopy.h1(store.classRow?.name ?? store.classSlug)}
        <span class="planner-header-suffix">
          · {store.raceRow?.name ?? ''} · Level {levelReached(store.order)}{specSuffix}
        </span>
      </h1>
    </div>
    <label class="planner-header-race">
      <span class="label text-muted">Race</span>
      <!-- A function binding rather than a plain `value` attribute: the value has to be
           applied after the options exist, and Svelte only does that for bindings
           (SummaryBar.svelte's own class select follows the identical pattern). -->
      <select
        class="planner-header-race-select"
        bind:value={() => store.raceSlug, (slug) => store.selectRace(slug)}
        disabled={store.readOnly}
        data-testid="planner-race-select"
      >
        {#each store.legalRaces as row (row.slug)}
          <option value={row.slug}>{row.name}</option>
        {/each}
      </select>
    </label>
  </div>
</div>

<style>
  .planner-header {
    display: flex;
    flex-direction: column;
    gap: 16px;
    width: 100%;
    /* Horizontal gutter comes from the caller (Planner.svelte's own `px-[18px] md:px-12`,
       the same convention every other panel in this page follows) -- only the vertical
       rhythm is this component's own, so it never double-pads alongside that wrapper. */
    padding: 22px 0 18px 0;
    border-radius: var(--radius-panel);
  }
  @media (min-width: 1024px) {
    .planner-header {
      padding: 28px 0 24px 0;
    }
  }
  .planner-header-eyebrow {
    display: flex;
    align-items: center;
    gap: 10px;
    color: var(--color-gold);
  }
  .planner-header-eyebrow-rule {
    width: 28px;
    height: 1px;
    background: var(--color-gold);
    display: inline-block;
  }
  .planner-header-row {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 18px;
    flex-wrap: wrap;
  }
  .planner-header-identity {
    display: flex;
    align-items: center;
    gap: 18px;
    min-width: 0;
  }
  .planner-header-crest {
    flex-shrink: 0;
    border-radius: 999px;
    box-shadow: 0 0 0 2px var(--c);
    background: var(--color-raised);
    object-fit: cover;
  }
  .planner-header-h1 {
    margin: 0;
    font-family: var(--font-display);
    font-size: 22px;
    line-height: 1.1;
    font-weight: 700;
    text-transform: none;
  }
  .planner-header-suffix {
    color: var(--color-muted);
    font-weight: 600;
    font-size: 14px;
    letter-spacing: 0.1em;
    text-transform: none;
  }
  .planner-header-race {
    display: flex;
    flex-direction: column;
    gap: 4px;
  }
  .planner-header-race-select {
    height: 44px;
    box-sizing: border-box;
    padding: 0 12px;
    border-radius: var(--radius-control);
    border: 1px solid var(--color-line);
    background: transparent;
    color: var(--color-text);
    font-size: 13px;
    font-family: var(--font-body);
  }
</style>
