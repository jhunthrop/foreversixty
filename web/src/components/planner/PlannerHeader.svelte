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

     Fix round 5: this component is now the band's own LEFT COLUMN content only -- no
     background, no border, no radius, no art. `PlannerHeaderBand.svelte` (its only caller)
     owns all of that at the band level, full-bleed across both grid columns (this header's
     own content and `PlannerCharacterCard`'s), which a box scoped to this component's own
     column could never reach. The crest still inlines `ClassCrest.astro`'s own recipe (a
     Svelte file cannot import an `.astro` one at all, the same constraint `CharacterCard.
     svelte` already documents). -->
<script lang="ts">
  import { classCrestSrc } from '../../lib/class-crest';
  import { classColorVar } from '../../lib/report/format';
  import { plannerHeaderCopy } from '../../lib/planner/copy';
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

  // The bar's one refusal line already carries the planner's own refusals (an illegal move,
  // a read-only build); a failed live estimate is the same kind of fact, so it goes in the
  // same line rather than a second one (SummaryBar.svelte's own reasoning, moved here with
  // the rest of the facts rail).
  const statusLine = $derived(store.refusal ?? (live.state === 'error' ? live.message : null) ?? '');
</script>

<div class="planner-header" data-testid="planner-header">
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

    <!-- The facts rail (spec §4.C, moved here by fix round 1 item 2.a; fix round 5: a
       compact row, not a stretched grid -- the mock's own panel wraps tight around its
       four facts, 36px apart, not a 4-column grid stretched to the header box's own full
       width with acres of empty space past DPS). Points left, Spent, Level and DPS sit
       inside the header's own panel as its bottom row -- the same four facts
       `SummaryBar.svelte` has always computed, just relocated. "Sim this build" is no
       longer bundled with the DPS figure here; it is its own outlined gold button in the
       rail, above the Share panel. -->
    <div class="planner-header-facts items-start" data-testid="planner-facts">
      <div class="flex flex-col gap-1">
        <span class="label text-muted">Points left</span>
        <span class="tabular text-gold font-mono text-[20px] leading-11" data-testid="planner-remaining">
          {MAX_POINTS - store.spent}
        </span>
        <!-- Fix round 6 (e2e: the facts rail's own height-stability tests, caught on
             mobile): `repeat(2, auto)` sizes each grid COLUMN from the widest content any
             row puts in it -- DPS's own long "live when the build reaches 51 points" note
             (column 2) forces column 1 down to barely wider than "42" while that note is
             showing, and without `whitespace-nowrap` here too, THIS note wrapped to a
             second line in that narrow column, inflating the whole panel's height exactly
             until DPS's own note goes short (`± N`) or empty and column 1 gets its width
             back. `whitespace-nowrap` (matching `PlannerDps.svelte`'s own identical fix)
             means this note overflows its cell rather than wraps it -- the cell's own
             height never moves, whatever column 1 is given. -->
        <span
          class="tabular text-muted block h-[14px] font-mono text-[11px] leading-[14px] whitespace-nowrap"
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
    </div>
    <p
      role="status"
      aria-live="polite"
      class="text-muted min-h-[20px] w-full text-[13px] leading-tight"
      data-testid="planner-refusal"
    >
      {statusLine}
    </p>
  </div>
</div>

<style>
  /* Fix round 5: a plain content column, no background/border/radius/overflow of its own --
     `PlannerHeaderBand.svelte` (this component's only caller) owns the band's background,
     full-bleed across both of its grid columns, which a box scoped to this one could never
     reach. */
  .planner-header {
    display: flex;
    flex-direction: column;
    width: 100%;
    min-width: 0;
  }
  .planner-header-content {
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
    /* Fix round 5: fit-content, not a full-width stretch -- the mock's own compact panel,
       wrapping tight around its facts rather than a grid stretched to the header's full
       column width with empty space past DPS.

       Fix round 6, item 4 (player review): a real 2x2 grid below 1100px, not organic
       flex-wrap -- DPS alone wrapping to its own orphan row (three facts fit one row, the
       fourth didn't) read as a layout bug because it was one: `flex-wrap` breaks wherever
       the row runs out of space, not at a deliberate point. `repeat(2, auto)` always pairs
       Points left/Spent and Level/DPS, 36px apart either way; `width: fit-content` only
       matters at the single-row tier, so it is scoped there too. */
    display: grid;
    grid-template-columns: repeat(2, auto);
    /* The 2x2 tier hugs its facts too: signed out at 1024 the header's left column is the
       whole content width, and a grid left to stretch spread two facts across 870px. */
    width: fit-content;
    max-width: 100%;
    justify-content: start;
    gap: 36px;
    border: 1px solid var(--color-line);
    border-radius: var(--radius-panel);
    background: color-mix(in srgb, var(--color-raised) 90%, transparent);
    padding: 14px 20px;
  }
  @media (min-width: 1100px) {
    .planner-header-facts {
      display: flex;
      flex-wrap: wrap;
      width: fit-content;
    }
  }
</style>
