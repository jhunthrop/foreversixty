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
  import { loadActionNames, type ActionNames } from '../../lib/sim/action-names';
  import { fetchSpecs } from '../../lib/sim/api';
  import { requestKind, type BulkResult, type WeightsResult } from '../../lib/sim/bulk-types';
  import { SIM_LEVEL, plannerHrefForSpec, type SimCharacter } from '../../lib/sim/character';
  import { bulkCopy, simCopy } from '../../lib/sim/copy';
  import { encounterLabel } from '../../lib/sim/encounter';
  import { confidenceBand, formatMargin } from '../../lib/sim/estimate';
  import { specLabel } from '../../lib/sim/spec-label';
  import { mergeSpecRows } from '../../lib/sim/spec-state';
  import type { SimResult, SpecFidelity } from '../../lib/sim/types';
  import { engineLabel, isStale } from '../../lib/sim/version';
  import CharacterStrip from './CharacterStrip.svelte';
  import DetailsCard from './DetailsCard.svelte';
  import RotationCard from './RotationCard.svelte';

  let { result, onrerun }: { result: SimResult; onrerun: () => void } = $props();

  let copied = $state(false);

  const stale = $derived(isStale(result.engine_version));
  const hasFigure = $derived(result.dps.mean > 0);
  const figure = $derived(hasFigure ? Math.round(result.dps.mean).toLocaleString('en-US') : '—');
  const band = $derived(
    hasFigure && result.dps.error > 0 ? `± ${formatMargin(confidenceBand(result.dps))}` : '',
  );
  // The stored request carries no settings preset, only the raw encounter and buff list --
  // the same shape `og-meta.ts`'s unfurl reads, and the same derivation: a raid-buffed run
  // logs at least one BUFF-type aura, a solo one logs none.
  const buffed = $derived(result.summary.auras.some((track) => track.type === 'BUFF'));
  const encounterText = $derived(encounterLabel(result.request.encounter, buffed));
  // The request's own capture time is the only timestamp a SimResult carries; an absent one
  // (an empty string) is not turned into a fabricated date.
  const savedDate = $derived(result.request.source.captured_at.slice(0, 10));

  // Kind is derived from the stored request, never sent and never stored twice (contract
  // 1.1) -- this is what lets a saved page pick its own results view without a second field
  // a client could set to anything.
  const kind = $derived(requestKind(result.request));

  const KIND_TITLES: Record<'gear' | 'talents' | 'drops' | 'weights', string> = {
    gear: bulkCopy.gearTitle,
    talents: bulkCopy.talentsTitle,
    drops: bulkCopy.dropsTitle,
    weights: bulkCopy.weightsTitle,
  };

  // Defect fix: GET /v1/sims/{id} used to drop the name a member gave this sim entirely
  // (api/internal/sims/handler.go's GetOutput carries it now, as `result.title`) -- the
  // heading fell back to the composed spec/kind line unconditionally, so a reviewer's own
  // "Thoradin - Fury Warrior, raid-buffed BWL night" never appeared anywhere on the page
  // they had just typed it into. The composed line remains the fallback for the common
  // case (naming a sim is optional): every kind title, spec label and DPS figure this page
  // already knows how to say, unchanged.
  const fallbackHeading = $derived(
    kind === 'run'
      ? specLabel(result.request.spec)
      : `${KIND_TITLES[kind]} · ${specLabel(result.request.spec)}`,
  );
  // Svelte's own text interpolation escapes this on render (SaveSimForm.svelte's input is
  // free text, kept exactly as typed): no `{@html}` here or anywhere else this reaches.
  const heading = $derived(
    result.title !== undefined && result.title !== '' ? result.title : fallbackHeading,
  );

  const gearKnown = $derived(result.request.character.gear.length > 0);

  /** The stored request back into the strip's own shape. Every field is already there. */
  const character = $derived<SimCharacter>({
    name: result.request.character.name,
    spec: result.request.spec,
    class_slug: result.request.character.class,
    race_slug: result.request.character.race,
    talent_level: SIM_LEVEL,
    tree_version: activeBuild.build,
    // Genuinely unknowable here, same as a combat-log character (sources.ts): a saved sim's
    // stored request carries the engine's final talent STRING, not the click order that
    // produced it, and there is no honest way back from one to the other. This is why the
    // strip's own default "Open in planner" derivation (plannerHrefFor, which needs
    // point_order to rebuild a talents string) is NOT what this page uses -- see
    // `plannerHref` below, which reaches the stored talents string directly instead.
    point_order: [],
    gear: Object.fromEntries(
      result.request.character.gear.map((slot): [string, number] => [slot.slot, slot.item_id]),
    ),
    // The richest form available on this page: the stored request's own gear list,
    // enchants and suffixes included, rather than a re-derivation from the lossy id map.
    gear_slots: result.request.character.gear,
    professions: [...(result.request.character.professions ?? [])],
    bags: [],
    bank: [],
    sets: [],
    loadouts: [],
    buffs: [...result.request.character.buffs],
    consumables: [...result.request.character.consumes],
    source: result.request.source,
  });

  /**
   * "Open in planner", built straight from the stored request's own `CharacterSpec` --
   * `result.request.character.talents` is the engine's final talents string, already
   * correct, with no `point_order`/`TalentIndex` reconstruction needed at all (the same
   * shortcut `combos.ts`'s `planItHref` takes for a bulk row's own "Plan it" link). Defect
   * fix: this page used to hand `CharacterStrip` a `character` with `point_order: []` and
   * let it fall through to its own `plannerHrefFor`, which -- having no order to work with
   * -- always encoded zeroed talents through `toCharacterSpec`.
   */
  const plannerHref = $derived(plannerHrefForSpec(result.request.character, activeBuild.build));

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

  // D48 (dps-minmaxer review round 2): a saved run reloaded cold used to hardcode
  // `actionNames={null}` below, so every ability, aura and cast name on this page read as
  // the engine's own raw key forever -- live and saved disagreed, which is the defect's
  // own name. The build's name table is loaded exactly like the item file just above, the
  // same loader `store.svelte.ts`'s own `ensureActionNames` calls for a live run
  // (`loadActionNames`, action-names.ts) -- one loader, reused, not a second path.
  let actionNames = $state<ActionNames | null>(null);
  void (async () => {
    try {
      actionNames = await loadActionNames(character.tree_version, character.class_slug);
    } catch {
      // resolveActionName's own humanised fallback (action-names.ts) covers a null table:
      // the page still reads as English, never as a raw key.
      actionNames = null;
    }
  })();

  // One load, for this component's one unchanging result -- not an `$effect`, the same
  // reason the item-file load above is not one. A failed fetch leaves the card without a
  // fidelity note, which is the honest rendering of "we do not know yet".
  let fidelity = $state<SpecFidelity | null>(null);
  void (async () => {
    try {
      const rows = mergeSpecRows(await fetchSpecs());
      fidelity = rows.find((row) => row.spec === result.request.spec) ?? null;
    } catch {
      fidelity = null;
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
  //
  // Three separate chunks, one per kind's own results view, so a saved weights page
  // downloads neither the damage tables (SimResults) nor the combination table
  // (SavedCombos), and a saved bulk page downloads neither the damage tables nor the
  // weights table -- only the one this particular result's kind actually needs.
  const simResultsLazy = createLazyComponent(() => import('./SimResults.svelte'));
  const savedCombosLazy = createLazyComponent(() => import('./SavedCombos.svelte'));
  const savedWeightsLazy = createLazyComponent(() => import('./SavedWeights.svelte'));

  if (kind === 'weights') savedWeightsLazy.load();
  else if (kind !== 'run') savedCombosLazy.load();
  else simResultsLazy.load();
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

<!-- A weights result has no gear story (design 7): nothing was tried in any slot, so the
     strip shows no grid rather than the equipped set a weights run never touched for its
     own sake. -->
<CharacterStrip
  {character}
  {items}
  gearKnown={kind !== 'weights' && gearKnown}
  readonly
  {plannerHref}
  onchange={() => {}}
/>

{#if kind === 'weights'}
  {#if savedWeightsLazy.current}
    <savedWeightsLazy.current result={result as WeightsResult} />
  {:else}
    {@render lazyFallback(savedWeightsLazy)}
  {/if}
{:else if kind !== 'run'}
  {#if savedCombosLazy.current}
    <savedCombosLazy.current result={result as BulkResult} treeVersion={character.tree_version} />
  {:else}
    {@render lazyFallback(savedCombosLazy)}
  {/if}
{:else if simResultsLazy.current}
  <simResultsLazy.current
    summary={result.summary}
    estimate={result.dps}
    iterationsRun={result.iterations_run}
    {actionNames}
    sample={result.sample}
  />
{:else}
  {@render lazyFallback(simResultsLazy)}
{/if}

<DetailsCard {result} />

<RotationCard spec={result.request.spec} {fidelity} />

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
