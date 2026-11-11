<!-- web/src/components/sim/SavedKind.svelte -->
<!-- /sim/<sim_id> for a bulk or weights result, read-only (the `run` kind is RunReport.svelte).
     A saved result is not the sim page with a result in it: no run control, no switcher, no
     settings bar -- one result header in their place, the strip built from the stored request
     rather than a live store, and the same results below. Task 17.

     The character strip's own item file is fetched once, the same `loadItems` call
     `store.svelte.ts`'s `adopt()` makes: a saved sim's class is learned only from `result`,
     so until it resolves the grid draws slot names and ids without icons or rarities,
     exactly as it does on /sim before its own fetch lands.

     Sim spec 2026-11-06 §4.5 grammar, the saved quick sim's: the character head with its answer
     (the figure, the 95% band, Open in the sim and Copy link), then a split page: the result
     table (its one Plate) and the character in the main column, This run (Facts) and Rotation
     in the rail. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import activeBuild from '../../data/active-build.json';
  import { loadItems, loadTalents } from '../../lib/planner/load';
  import { indexTalents, type TalentIndex } from '../../lib/planner/rules';
  import type { Item } from '../../lib/planner/types';
  import { createLazyComponent, type LazyLoadState } from '../../lib/report/lazy-component.svelte';
  import { fetchSpecs } from '../../lib/sim/api';
  import { requestKind, type BulkResult, type WeightsResult } from '../../lib/sim/bulk-types';
  import {
    MAX_LEVEL,
    NAMELESS_REQUEST_NAME,
    plannerHrefForSpec,
    talentPointsFromString,
    type SimCharacter,
  } from '../../lib/sim/character';
  import { bulkCopy, simCopy } from '../../lib/sim/copy';
  import { encounterLabel } from '../../lib/sim/encounter';
  import { savedDay } from '../../lib/sim/result-band';
  import type { AnswerModel } from '../../lib/sim/answer';
  import { confidenceBand, formatMargin } from '../../lib/sim/estimate';
  import { gameCodeForSimResult } from '../../lib/sim/game-code';
  import { simPageCopy } from '../../lib/sim/page-copy';
  import { grouped } from '../../lib/format/number';
  import { specLabel } from '../../lib/sim/spec-label';
  import { mergeSpecRows } from '../../lib/sim/spec-state';
  import type { ItemsStatus } from '../../lib/sim/store.svelte';
  import type { SimResult, SpecFidelity } from '../../lib/sim/types';
  import { isStale } from '../../lib/sim/version';
  import { yieldToMain } from '../../lib/yield';
  import CopyForGame from '../CopyForGame.svelte';
  import Button from '../ui/Button.svelte';
  import CopyButton from '../ui/CopyButton.svelte';
  import Notice from '../ui/Notice.svelte';
  import Page from '../ui/Page.svelte';
  import CharacterStrip from './CharacterStrip.svelte';
  import DetailsCard from './DetailsCard.svelte';
  import LiveRotationCard from './LiveRotationCard.svelte';
  import SimAnswer from './SimAnswer.svelte';
  import SimCharacterHead from './SimCharacterHead.svelte';

  let { result, onrerun }: { result: SimResult; onrerun: () => void } = $props();

  const stale = $derived(isStale(result.engine_version));
  const hasFigure = $derived(result.dps.mean > 0);
  // The stored request carries no settings preset, only the raw encounter and buff list --
  // the same shape `og-meta.ts`'s unfurl reads, and the same derivation: a raid-buffed run
  // logs at least one BUFF-type aura, a solo one logs none.
  const buffed = $derived(result.summary.auras.some((track) => track.type === 'BUFF'));
  const encounterText = $derived(encounterLabel(result.request.encounter, buffed));
  // The request's own capture time is the only timestamp a SimResult carries; an absent one
  // (an empty string) is not turned into a fabricated date.
  const savedOn = $derived(savedDay(result.request.source.captured_at));
  const headLines = $derived(
    savedOn === null ? [encounterText] : [encounterText, simPageCopy.head.saved(savedOn)],
  );

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
    kind === 'run' ? '' : `${KIND_TITLES[kind]} · ${specLabel(result.request.spec)}`,
  );
  // Svelte's own text interpolation escapes this on render (SaveSimForm.svelte's input is
  // free text, kept exactly as typed): no `{@html}` here or anywhere else this reaches.
  const heading = $derived(
    result.title !== undefined && result.title !== '' ? result.title : fallbackHeading,
  );

  /** The head's answer: the kind's name, the figure and its band, labelled once as 95%. */
  const answerModel = $derived<AnswerModel>({
    state: 'result',
    label: kind === 'run' ? '' : KIND_TITLES[kind],
    figure: hasFigure ? { value: grouped(result.dps.mean), tone: 'gold', unit: true } : undefined,
    caption:
      hasFigure && result.dps.error > 0
        ? { bold: `± ${formatMargin(confidenceBand(result.dps))} DPS` }
        : undefined,
    primary: 'run-again',
  });

  const gearKnown = $derived(result.request.character.gear.length > 0);

  /** The stored request back into the strip's own shape. Every field is already there. */
  const character = $derived<SimCharacter>({
    name: result.request.character.name === NAMELESS_REQUEST_NAME ? '' : result.request.character.name,
    spec: result.request.spec,
    class_slug: result.request.character.class,
    race_slug: result.request.character.race,
    talent_level: MAX_LEVEL,
    // The stored request's own level (design 6): a saved sim's own CharacterSpec already
    // carries whatever level it actually ran at, so the strip shows that rather than
    // asserting MAX_LEVEL for every saved sim regardless of what it was.
    level: result.request.character.level,
    // Stamped with the ACTIVE data build because the stored request carries no build of its
    // own: a sim saved against an older talent layout maps its digits onto today's layout,
    // and the addon cannot catch that (the code it receives names the active build).
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

  /**
   * The talent point count, the same digit-sum `plannerHref` above's own comment describes
   * for the talents string -- no `point_order` needed (2026-09-21 result-page review round
   * 3, newcomer's own finding: the live page shows "51 points" beside Open in planner, this
   * page showed nothing).
   */
  const talentPoints = $derived(talentPointsFromString(result.request.character.talents));

  let items = $state<Map<number, Item>>(new Map());
  let itemsStatus = $state<ItemsStatus>('loading');

  // One load, for this component's one (unchanging) character -- not an `$effect`, the
  // same reason `enterCompare` in `SimView.svelte` is a plain call rather than one: this
  // runs once, for the one result this instance was mounted with.
  void (async () => {
    try {
      const file = await loadItems(character.tree_version, character.class_slug);
      items = new Map(file.items.map((item) => [item.id, item]));
      itemsStatus = 'ready';
    } catch {
      itemsStatus = 'failed';
      // The strip renders slot names and "Empty" without the item file; a failed fetch
      // must not stop the rest of the page from rendering.
      items = new Map();
    }
  })();

  // Task L8.1, "Copy for the game": the code needs the class's tree layout to turn the stored
  // talents string's digits into cells. Nothing else on this page reads it (the strip skips
  // its own load because `plannerHref` is precomputed), so it is loaded here, once, through
  // the same cached `loadTalents` the strip and the planner use -- a static build file, not
  // an API call. A failed load leaves the button disabled rather than copying a build with
  // no talents in it.
  let talentIndex = $state<TalentIndex | null>(null);
  void (async () => {
    try {
      talentIndex = indexTalents(await loadTalents(character.tree_version, character.class_slug));
    } catch {
      talentIndex = null;
    }
  })();

  /**
   * The build this page's result is about, as the FSB1 code the addon's build dialog reads:
   * a bulk run's winning set, otherwise the stored request's own character
   * (`gameCodeForSimResult`, game-code.ts). Item stats come from the item file loaded above;
   * a slot with no record there is left out of the code.
   */
  const gameCode = $derived(gameCodeForSimResult(result, character.tree_version, talentIndex, items));

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

  // Two separate chunks, one per kind's own results view: a saved weights page downloads no
  // combination table and a saved bulk page no weights table.
  const savedCombosLazy = createLazyComponent(() => import('./SavedCombos.svelte'));
  const savedWeightsLazy = createLazyComponent(() => import('./SavedWeights.svelte'));

  if (kind === 'weights') savedWeightsLazy.load();
  else savedCombosLazy.load();

  // The page draws in three tasks, not one: the heading and the DPS header, then the character
  // strip (about 6 ms of the mount alone on a fast desktop: the paperdoll and its slots), then the
  // cards under the results (details, rotation, the actions row). As one task the mount was
  // ~9 ms there and ~186 ms of CI's throttled Lighthouse profile, most of /sim/<id>'s Total
  // Blocking Time against its 200 ms budget (2026-11-02). Each stage only adds below what is
  // drawn, and the results themselves arrive later still (their own chunk). A server render has
  // no tasks to spread, so it draws every stage at once.
  const ALL_STAGES = 2;
  let drawnStage = $state(typeof window === 'undefined' ? ALL_STAGES : 0);
  onMount(() => {
    let cancelled = false;
    void (async () => {
      while (drawnStage < ALL_STAGES) {
        await yieldToMain();
        if (cancelled) return;
        drawnStage += 1;
      }
    })();
    return () => {
      cancelled = true;
    };
  });
</script>

{#snippet lazyFallback(lazy: LazyLoadState)}
  {#if lazy.error !== ''}
    <p class="text-muted text-[14px] leading-5" role="alert" data-testid="sim-results-error">
      {lazy.error}
      <button
        type="button"
        class="text-strong ml-1 inline-flex min-h-11 items-center underline"
        onclick={() => lazy.load()}>{simCopy.tryAgain}</button
      >
    </p>
  {/if}
{/snippet}

{#snippet rail()}
  {#if drawnStage >= 2}
    <DetailsCard {result}>
      <CopyForGame code={gameCode} testid="sim-saved-copy-for-game" />
    </DetailsCard>
    <LiveRotationCard spec={result.request.spec} {fidelity} />
  {/if}
{/snippet}

<div data-testid="sim-saved-header">
  <SimCharacterHead name={heading} classSlug={result.request.character.class} lines={headLines}>
    {#snippet answer()}
      <SimAnswer model={answerModel} />
    {/snippet}
    {#snippet actions()}
      <Button tone="primary" size="sm" onclick={onrerun} testid="sim-run-yourself"
        >{simPageCopy.answer.openInSim}</Button
      >
      <CopyButton
        text={() => window.location.href}
        label={simPageCopy.answer.copyLink}
        testid="sim-saved-copy"
      />
    {/snippet}
  </SimCharacterHead>
</div>

{#if stale}
  <div class="px-[var(--page-inset)]">
    <Notice text={simCopy.staleEngine} tone="neutral" testid="sim-stale" />
  </div>
{/if}

<Page layout="split" panels {rail}>
  {#if kind === 'weights'}
    {#if savedWeightsLazy.current}
      <savedWeightsLazy.current result={result as WeightsResult} />
    {:else}
      {@render lazyFallback(savedWeightsLazy)}
    {/if}
  {:else if savedCombosLazy.current}
    <savedCombosLazy.current result={result as BulkResult} treeVersion={character.tree_version} />
  {:else}
    {@render lazyFallback(savedCombosLazy)}
  {/if}

  <!-- A weights result has no gear story (design 7): nothing was tried in any slot, so the
       strip shows no grid rather than the equipped set a weights run never touched for its
       own sake. -->
  {#if drawnStage >= 1}
    <CharacterStrip
      {character}
      {items}
      {itemsStatus}
      gearKnown={kind !== 'weights' && gearKnown}
      readonly
      plated={false}
      {plannerHref}
      {talentPoints}
      onchange={() => {}}
    />
  {/if}
</Page>
