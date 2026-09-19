<!-- web/src/components/sim/SavedSim.svelte -->
<!-- /sim/<sim_id>, read-only. A saved result is not the sim page with a result in it: no
     run control, no switcher, no settings bar -- one result header in their place, the
     strip built from the stored request rather than a live store, and the same results
     below. Task 17.

     The character strip's own item file is fetched once, the same `loadItems` call
     `store.svelte.ts`'s `adopt()` makes: a saved sim's class is learned only from `result`,
     so until it resolves the grid draws slot names and ids without icons or rarities,
     exactly as it does on /sim before its own fetch lands. -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import { loadItems } from '../../lib/planner/load';
  import type { Item } from '../../lib/planner/types';
  import { createLazyComponent, type LazyLoadState } from '../../lib/report/lazy-component.svelte';
  import { SIM_LEVEL, type SimCharacter } from '../../lib/sim/character';
  import { simCopy } from '../../lib/sim/copy';
  import { encounterLabel } from '../../lib/sim/encounter';
  import { confidenceBand } from '../../lib/sim/estimate';
  import { specLabel } from '../../lib/sim/spec-label';
  import type { SimResult } from '../../lib/sim/types';
  import { engineLabel, isStale } from '../../lib/sim/version';
  import CharacterStrip from './CharacterStrip.svelte';
  import DetailsCard from './DetailsCard.svelte';

  let { result, onrerun }: { result: SimResult; onrerun: () => void } = $props();

  let copied = $state(false);

  const stale = $derived(isStale(result.engine_version));
  const hasFigure = $derived(result.dps.mean > 0);
  const figure = $derived(hasFigure ? Math.round(result.dps.mean).toLocaleString('en-US') : '—');
  const band = $derived(
    hasFigure && result.dps.error > 0
      ? `± ${Math.round(confidenceBand(result.dps)).toLocaleString('en-US')}`
      : '',
  );
  // The stored request carries no settings preset, only the raw encounter and buff list --
  // the same shape `og-meta.ts`'s unfurl reads, and the same derivation: a raid-buffed run
  // logs at least one BUFF-type aura, a solo one logs none.
  const buffed = $derived(result.summary.auras.some((track) => track.type === 'BUFF'));
  const encounterText = $derived(encounterLabel(result.request.encounter, buffed));
  // The request's own capture time is the only timestamp a SimResult carries; an absent one
  // (an empty string) is not turned into a fabricated date.
  const savedDate = $derived(result.request.source.captured_at.slice(0, 10));

  // No title travels with a fetched SimResult -- the contract's `sims.title` column has no
  // mirror on the Go `SimResult` struct, only on the `GET /v1/sims?mine=1` row -- so the
  // heading always falls back to the spec, the one branch this shape can ever reach today.
  const heading = $derived(specLabel(result.request.spec));

  const gearKnown = $derived(result.request.character.gear.length > 0);

  /** The stored request back into the strip's own shape. Every field is already there. */
  const character = $derived<SimCharacter>({
    name: result.request.character.name,
    spec: result.request.spec,
    class_slug: result.request.character.class,
    race_slug: result.request.character.race,
    talent_level: SIM_LEVEL,
    tree_version: activeBuild.build,
    // The stored gear is the engine's list, not the planner's map; the strip only needs
    // item ids per slot and that is exactly what a GearSlot carries.
    point_order: [],
    gear: Object.fromEntries(
      result.request.character.gear.map((slot): [string, number] => [slot.slot, slot.item_id]),
    ),
    buffs: [...result.request.character.buffs],
    consumables: [...result.request.character.consumes],
    source: result.request.source,
  });

  let items = $state<Map<number, Item>>(new Map());

  // One load, for this component's one (unchanging) character -- not an `$effect`, the
  // same reason `enterCompare` in `SimView.svelte` is a plain call rather than one: this
  // runs once, for the one result this instance was mounted with.
  void (async () => {
    try {
      const file = await loadItems(character.tree_version, character.class_slug);
      items = new Map(file.items.map((item) => [item.id, item]));
    } catch {
      // The strip renders slot names and "Empty" without the item file; a failed fetch
      // must not stop the rest of the page from rendering.
      items = new Map();
    }
  })();

  async function copyLink(): Promise<void> {
    try {
      await navigator.clipboard.writeText(window.location.href);
      copied = true;
    } catch {
      copied = false;
    }
  }

  // SimResults pulls in five report components; a saved sim needs it for the very first
  // paint (it is the page's own content), but it still ships as its own chunk rather than
  // an eager import -- the same split /sim's own results use -- so the load starts the
  // moment this component mounts instead of adding its weight to the shared island bundle.
  const simResultsLazy = createLazyComponent(() => import('./SimResults.svelte'));
  simResultsLazy.load();
</script>

{#snippet lazyFallback(lazy: LazyLoadState)}
  {#if lazy.error !== ''}
    <p class="text-muted px-[18px] text-[13px] md:px-0" role="alert" data-testid="sim-results-error">
      {lazy.error}
      <button
        type="button"
        class="text-strong ml-1 inline-flex min-h-11 items-center underline"
        onclick={() => lazy.load()}>{simCopy.tryAgain}</button
      >
    </p>
  {/if}
{/snippet}

<div class="flex flex-wrap items-baseline gap-x-4 gap-y-1 px-[18px] md:px-0">
  <h1 class="section-title" data-testid="sim-saved-title">{heading}</h1>
</div>

<section
  class="bg-raised border-line rounded-panel mx-[18px] flex flex-wrap items-baseline gap-x-6 gap-y-2 border p-4 md:mx-0"
  data-testid="sim-saved-header"
>
  <div class="flex flex-wrap items-baseline gap-2">
    <span
      class={`tabular font-mono text-[32px] leading-none ${hasFigure ? 'text-gold' : 'text-muted'}`}
      data-testid="sim-dps"
    >
      {figure}
    </span>
    <span class="tabular text-muted font-mono text-[14px]" data-testid="sim-error">{band}</span>
  </div>
  <span class="text-muted text-[13px]" data-testid="sim-saved-encounter">{encounterText}</span>
  {#if stale}
    <span class="pill pill-sample" data-testid="sim-stale-pill">{engineLabel(result.engine_version)}</span>
  {:else}
    <span class="tabular text-muted font-mono text-[13px]" data-testid="sim-engine">
      {engineLabel(result.engine_version)}
    </span>
  {/if}
  {#if result.request.source.captured_at !== ''}
    <span class="tabular text-muted font-mono text-[13px]" data-testid="sim-saved-date">{savedDate}</span>
  {/if}
</section>

{#if stale}
  <p class="mx-[18px] text-[13px] md:mx-0" data-testid="sim-stale">
    <span class="text-strong">{simCopy.staleEngine}</span>
  </p>
{/if}

<CharacterStrip {character} {items} {gearKnown} readonly onchange={() => {}} />

{#if simResultsLazy.current}
  <simResultsLazy.current
    summary={result.summary}
    estimate={result.dps}
    iterationsRun={result.iterations_run}
    actionNames={null}
    sample={result.sample}
  />
{:else}
  {@render lazyFallback(simResultsLazy)}
{/if}

<DetailsCard {result} />

<div class="mx-[18px] flex flex-wrap items-center gap-3 md:mx-0">
  <button
    type="button"
    class="border-line-warm rounded-control text-nav label min-h-11 border px-4"
    onclick={onrerun}
    data-testid="sim-run-yourself"
  >
    {simCopy.runThisYourself}
  </button>
  <button
    type="button"
    class="border-line-warm rounded-control text-nav label min-h-11 border px-4"
    onclick={() => void copyLink()}
    data-testid="sim-saved-copy"
  >
    {copied ? simCopy.copied : simCopy.copyLink}
  </button>
</div>
