<!-- web/src/components/planner/PlannerHeader.svelte -->
<!-- Rebuild spec §4.A, fix round 1 (the approved mock, `day3/shots/boards/Planner.png`):
     the planner's own class-page header -- eyebrow, 64px crest, h1 with the always-present
     level/race suffix (plus a spec clause once one tree leads), a one-line description, the
     class/race selects, and the facts rail (Points left/Spent/Level/DPS) as the header's own
     bottom row. Lives inside the Planner island (not a `.astro` sibling like the BiS/home
     rebuild's `ClassHeader.astro`) because every one of these facts -- the class, the race,
     the level, the leading spec, the live build state -- is the island's own reactive store,
     not build-time frontmatter; `ClassHeader.astro`'s own markup is server-rendered once per
     faction and has no runtime state to read at all (see that component's header comment).

     Fix round 1, item 2.b: the mock's background is the dominant tree's own client art
     (TreeGrid.svelte's identical `dataUrl(store.treeVersion, 'trees/{background}.webp')`
     call, same "before one tree leads, the first tree's art" inference the h1 suffix already
     uses) behind a night-to-bg gradient at ~50% -- the coordinator's own ruling that the
     mock is the confirmation spec §4.A/§7's "ship flat until confirmed" clause asked for, so
     this no longer ships the flat `ArtPanel` fallback. The crest still inlines `ClassCrest.
     astro`'s own recipe (a Svelte file cannot import an `.astro` one at all, the same
     constraint `CharacterCard.svelte` already documents). -->
<script lang="ts">
  import { classCrestSrc } from '../../lib/class-crest';
  import { classColorVar } from '../../lib/report/format';
  import { plannerHeaderCopy } from '../../lib/planner/copy';
  import { dataUrl } from '../../lib/planner/load';
  import type { LiveDps } from '../../lib/planner/live-dps.svelte';
  import type { LiveGate } from '../../lib/planner/live-gate';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { MAX_POINTS } from '../../lib/planner/types';
  import { specOf } from '../../lib/sim/character';
  import { simCopy } from '../../lib/sim/copy';
  import { specLabel } from '../../lib/sim/spec-label';
  import PlannerDps from './PlannerDps.svelte';

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

  /** The dominant tree's own index -- the same "most points, ties to the first tree"
   *  inference `specKeyFor`/the h1 suffix already use, read here for the background art
   *  instead of the spec label. */
  const dominantTreeIndex = $derived.by(() => {
    const split = store.split;
    let best = 0;
    for (let i = 1; i < split.length; i += 1) {
      if (split[i] > split[best]) best = i;
    }
    return best;
  });

  const headerArtSrc = $derived.by(() => {
    const tree = store.talentIndex?.trees[dominantTreeIndex];
    return tree ? dataUrl(store.treeVersion, `trees/${tree.background}.webp`) : null;
  });
  // The bar's one refusal line already carries the planner's own refusals (an illegal move,
  // a read-only build); a failed live estimate is the same kind of fact, so it goes in the
  // same line rather than a second one (SummaryBar.svelte's own reasoning, moved here with
  // the rest of the facts rail).
  const statusLine = $derived(store.refusal ?? (live.state === 'error' ? live.message : null) ?? '');
</script>

<div class="planner-header" data-testid="planner-header">
  <!-- A CSS background-image, not an <img>: a nicer fit for a decorative, `aria-hidden`
       backdrop (no onerror fallback needed -- `.planner-header`'s own flat `--color-raised`
       background already shows through a missing or still-loading image, the same "ships
       flat" floor §4.A/§7 always guaranteed) -- but, contrary to this comment's own earlier
       round-1 claim, NOT a way to dodge LCP candidacy: Chrome counts a CSS background the
       moment it is visibly painted, `<img>`/`<video>`/text or not.

       Fix round 2, item 3 (ux-designer review): the header renders flat -- this div painted
       at `opacity:0`, nothing drawn -- until `store.talentIndex` resolves and names the
       dominant tree's own background, several round trips after first paint the crest/h1
       already won on (known from the URL alone, no fetch). An animated fade-in was tried
       here first and measured worse, not better (see `.planner-header-art`'s own style
       comment below for the numbers): Chrome still names this div the LCP element either
       way, so the fade bought nothing and cost real milliseconds. The swap is instant once
       the data is in hand instead. -->
  <div
    class="planner-header-art"
    class:planner-header-art-visible={headerArtSrc !== null}
    style={headerArtSrc !== null ? `background-image:url(${headerArtSrc})` : undefined}
    aria-hidden="true"
  ></div>
  <div class="planner-header-gradient" aria-hidden="true"></div>
  <div class="planner-header-content">
    <p class="label planner-header-eyebrow">
      <i class="planner-header-eyebrow-rule" aria-hidden="true"></i>
      {plannerHeaderCopy.eyebrow}
    </p>
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
      <div class="planner-header-titles">
        <h1 class="planner-header-h1" style={`color:${color}`} data-testid="planner-header-h1">
          {plannerHeaderCopy.h1(store.classRow?.name ?? store.classSlug)}
          <span class="planner-header-suffix">
            · {store.raceRow?.name ?? ''} · Level {store.level}{specSuffix}
          </span>
        </h1>
        <p class="planner-header-description">{plannerHeaderCopy.description}</p>
      </div>
    </div>

    <div class="planner-header-controls">
      <label class="planner-header-select-label">
        <span class="label text-muted">Class</span>
        <!-- A function binding rather than a plain `value` attribute: the value has to be
             applied after the options exist, and Svelte only does that for bindings. -->
        <select
          class="planner-header-select"
          bind:value={() => store.classSlug, (slug) => store.selectClass(slug)}
          disabled={store.readOnly}
        >
          {#each store.classes as row (row.slug)}
            <option value={row.slug}>{row.name}</option>
          {/each}
        </select>
      </label>
      <label class="planner-header-select-label">
        <span class="label text-muted">Race</span>
        <select
          class="planner-header-select"
          bind:value={() => store.raceSlug, (slug) => store.selectRace(slug)}
          disabled={store.readOnly}
          data-testid="planner-race-select"
        >
          {#each store.legalRaces as row (row.slug)}
            <option value={row.slug}>{row.name}</option>
          {/each}
        </select>
      </label>
      <span class="planner-header-faction-note">{plannerHeaderCopy.factionFollowsRace}</span>
    </div>

    <!-- The facts rail (spec §4.C, moved here by fix round 1 item 2.a): Points left, Spent,
       Level and DPS now sit inside the header's own panel as its bottom row, matching the
       mock exactly -- the same four facts `SummaryBar.svelte` has always computed, just
       relocated. "Sim this build" is no longer bundled with the DPS figure here; it is its
       own outlined gold button in the rail, above the Share panel. -->
    <div
      class="planner-header-facts grid grid-cols-2 items-start gap-x-5 gap-y-3 sm:grid-cols-4"
      data-testid="planner-facts"
    >
      <div class="flex flex-col gap-1">
        <span class="label text-muted">Points left</span>
        <span class="tabular text-gold font-mono text-[20px] leading-11" data-testid="planner-remaining">
          {MAX_POINTS - store.spent}
        </span>
        <span
          class="tabular text-muted block h-[18px] font-mono text-[12px] leading-[18px]"
          data-testid="planner-remaining-note"
        >
          {gate === 'unfinished' ? simCopy.plannerDpsPointsToGo(MAX_POINTS - store.spent) : ''}
        </span>
      </div>
      <div class="flex flex-col gap-1">
        <span class="label text-muted">Spent</span>
        <span class="tabular text-strong font-mono text-[20px] leading-11" data-testid="planner-spent">
          {store.spent}/{MAX_POINTS}
        </span>
      </div>
      <div class="flex flex-col gap-1">
        <span class="label text-muted">Level</span>
        <span class="tabular text-strong font-mono text-[20px] leading-11" data-testid="planner-level">
          {store.level}
        </span>
      </div>
      <PlannerDps
        {live}
        {gate}
        pointsLeft={MAX_POINTS - store.spent}
        onshow={onshowdps}
        showSimLink={false}
        unfinishedNote="static"
      />
      <p
        role="status"
        aria-live="polite"
        class="text-muted col-span-2 min-h-[20px] w-full text-[13px] leading-tight sm:col-span-4"
        data-testid="planner-refusal"
      >
        {statusLine}
      </p>
    </div>
  </div>
</div>

<style>
  .planner-header {
    position: relative;
    display: flex;
    flex-direction: column;
    width: 100%;
    overflow: hidden;
    border-radius: var(--radius-panel);
    background: var(--color-raised);
  }
  .planner-header-art {
    position: absolute;
    inset: 0;
    width: 100%;
    height: 100%;
    background-size: cover;
    background-position: center;
    /* Flat until the data resolves (fix round 2, item 3): opacity 0 so the element paints
       nothing at first render, whatever `.planner-header`'s own flat background already
       shows through is the only "floor" on first paint.

       No `transition` here deliberately -- an animated fade was tried first and measured
       with Lighthouse (mobile-throttled, real data): it moved the LCP timestamp *later*
       (3660ms vs 3273ms; performance 0.89 vs 0.92, under this repo's own 0.90 floor for
       this page) for no candidacy benefit at all, since Chrome still names this element
       the LCP candidate the instant it is first painted at non-zero opacity, transition or
       not (`largest-contentful-paint-element` names `.planner-header-art` in both a
       transitioned and an un-transitioned build -- a CSS background-image is always LCP-
       eligible once visible; nothing about *how* it becomes visible changes that). An
       instant swap costs nothing extra once the image has already decoded off the critical
       path, which an animated one does not get back. See the build report's "Fix round 2 /
       item 3" section for the measurements and the real, still-open fix this would take (a
       build-time class-to-background map, or shrinking the art below the h1's own
       footprint) -- neither fits this round's scope. */
    opacity: 0;
  }
  .planner-header-art-visible {
    opacity: 1;
  }
  .planner-header-gradient {
    position: absolute;
    inset: 0;
    /* Night-to-bg gradient at ~50% (fix round 1, item 2.b): the art reads behind it at full
       strength, tinted top to bottom, rather than washed out by an opaque overlay. */
    background: linear-gradient(180deg, var(--color-night) 0%, var(--color-bg) 100%);
    opacity: 0.5;
  }
  .planner-header-content {
    position: relative;
    display: flex;
    flex-direction: column;
    gap: 16px;
    padding: 22px 0 18px 0;
  }
  @media (min-width: 1024px) {
    .planner-header-content {
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
  .planner-header-identity {
    display: flex;
    align-items: center;
    gap: 18px;
    min-width: 0;
  }
  .planner-header-titles {
    display: flex;
    flex-direction: column;
    gap: 6px;
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
  .planner-header-description {
    margin: 0;
    font-size: 14px;
    color: var(--color-nav);
    max-width: 62ch;
  }
  .planner-header-controls {
    display: flex;
    align-items: center;
    gap: 14px;
    flex-wrap: wrap;
  }
  .planner-header-select-label {
    display: flex;
    flex-direction: column;
    gap: 4px;
  }
  .planner-header-select {
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
  .planner-header-faction-note {
    font-size: 13px;
    color: var(--color-muted);
  }
  .planner-header-facts {
    /* Layout itself is the Tailwind grid classes on the element (fix round 1: a fixed
       column count, 2 base / 4 from sm, so the row never organically reflows the way a
       flex-wrap row does as the DPS figure's own width changes -- that reflow was this
       panel's own CLS regression, caught by planner-dps.spec.ts's height-stability tests). */
    border: 1px solid var(--color-line);
    border-radius: var(--radius-panel);
    background: color-mix(in srgb, var(--color-raised) 70%, transparent);
    padding: 16px 20px;
  }
</style>
