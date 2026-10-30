<!-- web/src/components/planner/PlannerHeadControls.svelte -->
<!-- The planner head's `below` row (site alignment spec 2026-10-29 §5.5): the Class and Race
     selects, the facts rail (Points left, Spent, Level and the role's score) and the one refusal
     line. The title, crest and character card are PlannerHead's own PageHead slots. -->
<script lang="ts">
  import { plannerHeaderCopy } from '../../lib/planner/copy';
  import type { LiveDps } from '../../lib/planner/live-dps.svelte';
  import type { LiveGate } from '../../lib/planner/live-gate';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { FACT_FIGURE_CLASS } from '../../lib/planner/styles';
  import PlannerDps from './PlannerDps.svelte';
  import StatTile from '../ui/StatTile.svelte';
  import Select from '../ui/Select.svelte';
  import { scoreUnitForBuild } from '../../lib/planner/score-unit';

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

  // The bar's one refusal line already carries the planner's own refusals (an illegal move,
  // a read-only build); a failed live estimate is the same kind of fact, so it goes in the
  // same line rather than a second one (SummaryBar.svelte's own reasoning, moved here with
  // the rest of the facts rail).
  const statusLine = $derived(store.refusal ?? (live.state === 'error' ? live.message : null) ?? '');
  const pointsLeft = $derived(Math.max(0, store.pointsAvailable - store.spent));
  const scoreUnit = $derived(scoreUnitForBuild(store.classSlug, store.split));
  // `Select` binds a value of its own. The store is the truth: the pickers mirror it, and only a
  // choice that differs from it (and names something) reaches `selectClass` / `selectRace` -- a
  // bound select that wrote its initial, option-less '' back would reset the build on mount.
  let classValue = $derived(store.classSlug);
  let raceValue = $derived(store.raceSlug);
  $effect(() => {
    if (classValue !== '' && classValue !== store.classSlug) store.selectClass(classValue);
  });
  $effect(() => {
    if (raceValue !== '' && raceValue !== store.raceSlug) store.selectRace(raceValue);
  });
  const classOptions = $derived(store.classes.map((row) => ({ value: row.slug, label: row.name })));
  const raceOptions = $derived(store.legalRaces.map((row) => ({ value: row.slug, label: row.name })));
</script>

<div class="flex w-full min-w-0 flex-col gap-4" data-testid="planner-header">
  <div class="planner-header-controls">
    <div class="planner-header-select-field">
      <label class="quiet-label" for="planner-class-select">Class</label>

      <Select
        id="planner-class-select"
        compact
        options={classOptions}
        bind:value={classValue}
        disabled={store.readOnly}
      />
    </div>
    <div class="planner-header-select-field">
      <label class="quiet-label" for="planner-race-select">Race</label>
      <Select
        id="planner-race-select"
        compact
        options={raceOptions}
        bind:value={raceValue}
        disabled={store.readOnly}
        testid="planner-race-select"
      />
    </div>
    <span class="planner-header-faction-note">{plannerHeaderCopy.factionFollowsRace}</span>
  </div>

  <!-- The facts rail (spec 4.C, polish spec 5.1): four tiles -- Points left, Spent, Level and
       the role's score -- in a 2 x 2 grid on a phone, one wrapping row from 1100px. Each grid
       child is `min-w-0` (StatTile) and the grid is the column's full width below 640px,
       so nothing runs past the viewport; the DPS note wraps under its figure. The tiles
       stretch to the tallest one, so the rail reads as one row of equal tiles. "Sim this
       build" is its own button in the rail, above the Share panel. -->
  <div class="planner-header-facts" data-testid="planner-facts">
    <StatTile
      label="Points left"
      value={String(pointsLeft)}
      figureTestid="planner-remaining"
      figureClass={FACT_FIGURE_CLASS}
      tone="gold"
    />
    <StatTile
      label="Spent"
      value={`${store.spent}/${store.pointsAvailable}`}
      figureTestid="planner-spent"
      figureClass={FACT_FIGURE_CLASS}
    />
    <StatTile
      label="Level"
      value={String(store.level)}
      figureTestid="planner-level"
      figureClass={FACT_FIGURE_CLASS}
    />
    <PlannerDps
      unit={scoreUnit}
      {live}
      {gate}
      {pointsLeft}
      onshow={onshowdps}
      showSimLink={false}
      unfinishedNote="static"
      tile
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

<style>
  .planner-header-controls {
    display: flex;
    align-items: center;
    gap: 14px;
    flex-wrap: wrap;
  }
  /* The server-rendered selects hold no options until the talent data lands, so a select that
     sizes to its widest option is only as wide as its padding and arrow and then jumps to its
     hydrated width. That widened pair once pushed the faction note onto its own row below 560px
     and shifted everything under the controls by 33px (0.086 of layout shift on /planner at
     360px). A fixed field width, wider than any class or race name at 15px, holds the row the
     same before and after the data arrives. */
  .planner-header-select-field {
    display: flex;
    flex-direction: column;
    gap: 4px;
    width: 152px;
  }
  .planner-header-faction-note {
    font-size: 13px;
    color: var(--color-muted);
  }
  .planner-header-facts {
    /* Polish spec 5.1: the rail is its tiles -- no panel around them. Below 640px a 2 x 2 grid
       of equal columns that is the column's full width (`minmax(0, 1fr)` lets a tile shrink
       below its content, which is what the old `auto` columns could not: the DPS note forced
       the grid past the viewport). From 640px the grid hugs its tiles, two by two; from 1100px
       one wrapping row. */
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    align-items: stretch;
    width: 100%;
    gap: 12px;
  }
  @media (min-width: 640px) {
    .planner-header-facts {
      grid-template-columns: repeat(2, auto);
      width: fit-content;
      max-width: 100%;
      justify-content: start;
    }
  }
  @media (min-width: 1100px) {
    .planner-header-facts {
      display: flex;
      flex-wrap: wrap;
      align-items: stretch;
    }
  }
</style>
