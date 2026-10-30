<!-- web/src/components/planner/SummaryBar.svelte -->
<!-- design/DESIGN-SYSTEM.md "State panel": label rows on a raised panel, mono numbers,
     gold accent. Pinned to the top of the viewport on phone (design/Mobile.dc.html keeps
     the current state above the fold), static on desktop. -->
<script lang="ts">
  import type { LiveDps } from '../../lib/planner/live-dps.svelte';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { MAX_POINTS } from '../../lib/planner/types';
  import PlannerDps from './PlannerDps.svelte';
  import { scoreUnitForBuild } from '../../lib/planner/score-unit';
  import type { LiveGate } from '../../lib/planner/live-gate';
  import { classOption } from '../../lib/ui/select-options';
  import Select from '../ui/Select.svelte';

  let {
    store,
    live,
    simHref,
    gate,
    standalone,
    onshowdps,
  }: {
    store: PlannerStore;
    live: LiveDps;
    simHref: string;
    gate: LiveGate;
    /**
     * newcomer round 4 (review.md:83-110): the one embedder that mounts this NOT standalone
     * -- TalentCandidates.svelte's inline ADD A BUILD editor -- can hold a pasted build
     * whose gear differs from the player's own current character, which the ranked
     * comparison beside it locks to. `PlannerDps` needs this to say so beside its own
     * figure; the real /planner page always shows the player's own gear, where the caveat
     * would be noise.
     */
    standalone: boolean;
    onshowdps: () => void;
  } = $props();

  // The bar's one refusal line already carries the planner's own refusals (an illegal move,
  // a read-only build); a failed live estimate is the same kind of fact -- something the
  // player did produced no result -- so it goes in the same line rather than a second one.
  // The planner's own refusal wins when both are set: it is about the edit that just
  // happened, and the DPS failure describes the build as a whole.
  const statusLine = $derived(store.refusal ?? (live.state === 'error' ? live.message : null) ?? '');

  const scoreUnit = $derived(scoreUnitForBuild(store.classSlug, store.split));
</script>

<div
  class="border-line bg-raised md:rounded-panel sticky top-0 z-20 flex flex-wrap items-start gap-3 border-b px-[18px] py-3 md:static md:gap-5 md:border md:px-5"
>
  <div class="flex flex-col gap-1">
    <label class="quiet-label" for="planner-summary-class">Class</label>
    <!-- A function binding: the store is the truth, and a pick reaches selectClass. -->
    <Select
      id="planner-summary-class"
      class="w-[180px]"
      size="compact"
      bind:value={() => store.classSlug, (slug) => store.selectClass(slug)}
      options={store.classes.map((row) =>
        classOption({ value: row.slug, label: row.name, className: row.slug }),
      )}
      disabled={store.readOnly}
    />
  </div>

  <!-- Rebuild spec §4.A: the race select moves to the header, restyled onto ClassHeader's
       own token -- this bar keeps only the class select, which the header has no room for
       (the header's h1 already names the class). -->

  <!-- Rebuild spec §4.C, review finding 2: Level is now always rendered -- it needs no
       character, only `levelReached(order)`, which is computable from the moment the store
       exists. Hiding it behind `hasCharacter && store.spent > 0` was exactly the "level gate
       is an afterthought" complaint the review named. -->
  <div class="flex flex-col gap-1">
    <span class="quiet-label">Level</span>
    <span class="text-strong numeric text-[20px] leading-11" data-testid="planner-level">
      {store.level}
    </span>
  </div>

  <div class="flex flex-col gap-1">
    <span class="quiet-label">Points left</span>
    <span class="text-gold numeric text-[20px] leading-11" data-testid="planner-remaining">
      {MAX_POINTS - store.spent}
    </span>
  </div>

  <div class="flex flex-col gap-1">
    <span class="quiet-label">Spent</span>
    <span class="text-strong numeric text-[20px] leading-11" data-testid="planner-spent">
      {store.spent}/{MAX_POINTS}
    </span>
  </div>

  <PlannerDps
    unit={scoreUnit}
    {live}
    href={simHref}
    {gate}
    pointsLeft={MAX_POINTS - store.spent}
    onshow={onshowdps}
    ownGearCaveat={!standalone}
  />

  <p
    role="status"
    aria-live="polite"
    class="text-muted order-last min-h-[20px] w-full text-[13px] leading-tight"
    data-testid="planner-refusal"
  >
    {statusLine}
  </p>
</div>
