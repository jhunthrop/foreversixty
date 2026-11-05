<!-- web/src/components/sim/SavedRun.svelte -->
<!-- A saved `run` result on one scrolling page (spec 2026-11-04-sim-results): the band and its
     answer, the sticky section nav, Worth knowing, Damage by ability beside the DPS spread, then
     each deeper card in its own chunk, mounted as it nears the viewport with a skeleton holding
     its ready height. Every file it reads is optional: a row without its icon or name still
     reads, nothing here can stop the figure from drawing. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import activeBuild from '../../data/active-build.json';
  import { dataUrl, loadItems } from '../../lib/planner/load';
  import type { Item } from '../../lib/planner/types';
  import { abilityTable, type AbilityRow } from '../../lib/report/ability-table';
  import { createLazyComponent } from '../../lib/report/lazy-component.svelte';
  import {
    loadActionNames,
    parseActionKey,
    resolveActionName,
    withShardNames,
    type ActionNames,
  } from '../../lib/sim/action-names';
  import { fetchSpecs } from '../../lib/sim/api';
  import { loadBuffNames, type BuffNames } from '../../lib/sim/buff-names';
  import { plannerHrefForSpec } from '../../lib/sim/character';
  import { extraAttackLines } from '../../lib/sim/extra-attacks';
  import { withoutAbsentPets } from '../../lib/sim/pet-rows';
  import { loadPresetSets, raidSetAt } from '../../lib/sim/preset-sets';
  import { clock, damageCopy, grouped, oneDecimal, sectionHeadCopy, spreadCopy } from '../../lib/sim/result-copy';
  import { SECTION_RESERVE, resultSections, timelineSkeletonHeight } from '../../lib/sim/result-sections';
  import { resultTip } from '../../lib/sim/result-tips';
  import { summaryActionKeys } from '../../lib/sim/sentence';
  import { specDisplayName } from '../../lib/sim/spec-label';
  import { mergeSpecRows } from '../../lib/sim/spec-state';
  import { NO_SPELL_INFO, loadSpellInfo, type SpellInfo } from '../../lib/sim/spell-info';
  import type { SimResult, SpecFidelity } from '../../lib/sim/types';
  import { simWarnings } from '../../lib/sim/warnings';
  import { yieldToMain } from '../../lib/yield';
  import AbilityTable, { type AbilityDecoration, type AbilityWeapon } from '../report/AbilityTable.svelte';
  import Card from '../ui/Card.svelte';
  import DpsDistribution from './DpsDistribution.svelte';
  import LazySection from './LazySection.svelte';
  import ResultBand from './ResultBand.svelte';
  import SectionNav from './SectionNav.svelte';
  import WorthKnowing from './WorthKnowing.svelte';

  let { result, onrerun, simId = '' }: { result: SimResult; onrerun: () => void; simId?: string } = $props();

  const build = activeBuild.build;
  // svelte-ignore state_referenced_locally
  const classSlug = result.request.character.class;

  let names = $state<ActionNames | null>(null);
  let spells = $state<SpellInfo>(NO_SPELL_INFO);
  let items = $state<ReadonlyMap<number, Item>>(new Map());
  let buffNames = $state<BuffNames | null>(null);
  let raidBuffs = $state<readonly string[] | null>(null);
  let fidelity = $state<SpecFidelity | null>(null);

  const summary = $derived(withoutAbsentPets(result.summary));
  const sections = $derived(resultSections(result));
  const resourceTitle = $derived(sections.find((section) => section.id === 'resource')?.label ?? '');
  const hasSample = $derived((result.sample?.length ?? 0) > 0);
  const hasGear = $derived(result.request.character.gear.length > 0);

  // One load each, for this component's one unchanging result. Each failure leaves the page
  // standing on its fallback: unnamed rows, empty icon boxes, no weapon names.
  /** Runs one optional load; a failure is logged and the page keeps its fallback. */
  function loadQuietly(what: string, load: () => Promise<void>): void {
    load().catch((error: unknown) => console.error(`sim result: ${what} did not load`, error));
  }

  // svelte-ignore state_referenced_locally
  const { spec, character } = result.request;

  onMount(() => {
    loadQuietly('the action names', async () => {
      const loaded = await loadActionNames(build, classSlug);
      names = loaded;
      names = await withShardNames(loaded, summaryActionKeys(summary));
    });
    loadQuietly('the spell data', async () => {
      spells = await loadSpellInfo(build, classSlug);
    });
    loadQuietly('the item file', async () => {
      const file = await loadItems(build, classSlug);
      items = new Map(file.items.map((item) => [item.id, item]));
    });
    loadQuietly('the buff table', async () => {
      buffNames = await loadBuffNames(build);
    });
    loadQuietly('the preset table', async () => {
      await loadPresetSets();
      raidBuffs = raidSetAt(spec, character.level)?.buffs ?? [];
    });
    loadQuietly('the fidelity row', async () => {
      fidelity = mergeSpecRows(await fetchSpecs()).find((row) => row.spec === spec) ?? null;
    });
  });

  const notes = $derived(simWarnings({ result, names, fact: spells.fact }));
  const plannerHref = $derived(plannerHrefForSpec(character, build));
  const table = $derived(
    abilityTable({
      actors: summary.damage_done,
      casts: summary.casts,
      auras: summary.auras,
      durationMs: summary.duration_ms,
      names,
    }),
  );
  const damageMeta = $derived(damageCopy.meta(oneDecimal(result.dps.mean), clock(summary.duration_ms / 1000)));
  const figure = $derived(oneDecimal(result.dps.mean));

  const HEROIC_STRIKE_RANKS: readonly number[] = [78, 284, 285, 1608, 11564, 11565, 11566, 11567, 25286];
  const swingsBecame = $derived(
    summary.casts
      .filter((row) => {
        const parsed = parseActionKey(row.spell_name);
        return parsed?.kind === 'spell' && parsed.tag === 0 && HEROIC_STRIKE_RANKS.includes(parsed.id);
      })
      .reduce((sum, row) => sum + row.succeeded, 0),
  );

  /** The weapon a swing row's hand holds; Extra attacks swing the main hand. */
  function weaponFor(row: AbilityRow): AbilityWeapon | null {
    if (row.hand === null) return null;
    const slot = row.hand === 'off' ? 'off_hand' : 'main_hand';
    const id = character.gear.find((entry) => entry.slot === slot)?.item_id;
    const item = id === undefined ? undefined : items.get(id);
    if (item === undefined) return null;
    return {
      id: item.id,
      name: item.name,
      quality: item.quality,
      iconUrl: item.icon === '' ? null : dataUrl(build, `icons/${item.icon}.webp`),
    };
  }

  const extraSources = $derived(
    extraAttackLines(
      result,
      buffNames,
      (id) => items.get(id)?.name,
      (key) => resolveActionName(key, names),
    ),
  );

  function decorate(row: AbilityRow): AbilityDecoration {
    const fact = row.spellId === 0 ? undefined : spells.fact(row.spellId);
    return {
      iconUrl: spells.iconUrl(row.key),
      rank: fact?.rank ?? 0,
      weapon: weaponFor(row),
      sources: row.hand === 'extra' ? extraSources : [],
      swingsBecame: row.hand === 'main' ? swingsBecame : 0,
      tip: fact !== undefined,
    };
  }

  // The band's figure scrolling out brings the nav's compact figure in.
  let figureElement = $state<HTMLElement | undefined>(undefined);
  let figureGone = $state(false);
  $effect(() => {
    if (figureElement === undefined || typeof IntersectionObserver === 'undefined') return;
    const observer = new IntersectionObserver((entries) => {
      figureGone = entries.every((entry) => !entry.isIntersecting);
    });
    observer.observe(figureElement);
    return () => observer.disconnect();
  });

  // Below the fold each card is its own chunk; the shared tooltip host loads after first paint.
  const timelineLazy = createLazyComponent(() => import('./SimTimeline.svelte'));
  const gearLazy = createLazyComponent(() => import('./SimGearCard.svelte'));
  const talentsLazy = createLazyComponent(() => import('./SimTalentsCard.svelte'));
  const resourceLazy = createLazyComponent(() => import('./SimResourceCard.svelte'));
  const aurasLazy = createLazyComponent(() => import('./SimAurasCard.svelte'));
  const rotationLazy = createLazyComponent(() => import('./RotationCard.svelte'));
  const settingsLazy = createLazyComponent(() => import('./SimSettingsCard.svelte'));
  const tooltipHostLazy = createLazyComponent(() => import('../BisTooltipHost.svelte'));

  // The page draws in two tasks: the band and the answer first, then everything under them. One
  // task here is most of /sim/<id>'s Total Blocking Time under Lighthouse's 4x CPU profile.
  let drawnRest = $state(typeof window === 'undefined');
  onMount(() => {
    let cancelled = false;
    void yieldToMain().then(() => {
      if (cancelled) return;
      drawnRest = true;
      tooltipHostLazy.load();
    });
    return () => {
      cancelled = true;
    };
  });

  const resolveTip = (key: string) => Promise.resolve(resultTip(key, spells));
  const resolveItem = async (id: number) => {
    const { fetchItemTooltipModel } = await import('../../lib/items/lookup');
    return fetchItemTooltipModel(build, classSlug, id);
  };
</script>

<div class="sim-result flex flex-col" data-testid="sim-saved-run">
  <ResultBand {result} {onrerun} {raidBuffs} bind:figureElement />
  <SectionNav {sections} {figure} figureVisible={figureGone} drawn={drawnRest} />

  <div class="flex flex-col gap-[22px] pt-[22px] md:gap-8 md:pt-8">
    {#if notes.length > 0}
      <WorthKnowing {notes} {spells} {plannerHref} {onrerun} />
    {/if}

    <div class="grid grid-cols-1 items-start gap-4 xl:grid-cols-[minmax(0,8fr)_minmax(0,4fr)]">
      <Card
        id="damage"
        glyph="chart"
        label={damageCopy.label}
        title={damageCopy.title}
        meta={damageMeta}
        body="none"
        class="mx-[18px] scroll-mt-16 md:mx-0"
        testid="sim-damage-card"
      >
        <AbilityTable {table} {decorate} />
      </Card>
      <Card
        glyph="spark"
        label={spreadCopy.label(grouped(result.iterations_run))}
        title={spreadCopy.title}
        body="none"
        class="mx-[18px] md:mx-0"
        testid="sim-spread-card"
      >
        <div class="border-line-soft border-t">
          <DpsDistribution
            estimate={result.dps}
            iterationsRun={result.iterations_run}
            histogram={result.dps_histogram}
          />
        </div>
      </Card>
    </div>

    {#if drawnRest}
      <!-- A run with no sample or no gear keeps the card (its head and one line); only the nav drops
           the link, so these two carry no reservation then. -->
      <LazySection
        id="timeline"
        head={{ glyph: 'chart', ...sectionHeadCopy.timeline }}
        lazy={timelineLazy}
        ready={timelineLazy.current !== null}
        style={hasSample ? `min-height: ${timelineSkeletonHeight(result)}px` : undefined}
      >
        {#if timelineLazy.current}<timelineLazy.current {result} {names} />{/if}
      </LazySection>
      <LazySection
        id="gear"
        head={{ glyph: 'guild', ...sectionHeadCopy.gear }}
        lazy={gearLazy}
        ready={gearLazy.current !== null}
        reserve={hasGear ? SECTION_RESERVE.gear : ''}
      >
        {#if gearLazy.current}<gearLazy.current {result} {names} />{/if}
      </LazySection>
      <LazySection
        id="talents"
        head={{ glyph: 'tree', ...sectionHeadCopy.talents }}
        lazy={talentsLazy}
        ready={talentsLazy.current !== null}
        reserve={SECTION_RESERVE.talents}
      >
        {#if talentsLazy.current}<talentsLazy.current {result} {names} />{/if}
      </LazySection>
      {#if sections.some((section) => section.id === 'resource')}
        <LazySection
          id="resource"
        head={{ glyph: 'bolt', ...sectionHeadCopy.resource, title: resourceTitle }}
          lazy={resourceLazy}
          ready={resourceLazy.current !== null}
          reserve={SECTION_RESERVE.resource}
        >
          {#if resourceLazy.current}<resourceLazy.current {result} {names} />{/if}
        </LazySection>
      {/if}
      <LazySection id="buffs"
        head={{ glyph: 'spark', ...sectionHeadCopy.buffs }} lazy={aurasLazy} ready={aurasLazy.current !== null} reserve={SECTION_RESERVE.buffs}>
        {#if aurasLazy.current}<aurasLazy.current {result} {names} />{/if}
      </LazySection>
      <LazySection
        id="rotation"
        head={{ glyph: 'list', title: sectionHeadCopy.rotation.title, label: sectionHeadCopy.rotation.label(specDisplayName(spec), character.level) }}
        lazy={rotationLazy}
        ready={rotationLazy.current !== null}
        reserve={SECTION_RESERVE.rotation}
      >
        {#if rotationLazy.current}<rotationLazy.current {spec} {fidelity} />{/if}
      </LazySection>
      <LazySection
        id="settings"
        head={{ glyph: 'gear', ...sectionHeadCopy.settings }}
        lazy={settingsLazy}
        ready={settingsLazy.current !== null}
        reserve={SECTION_RESERVE.settings}
      >
        {#if settingsLazy.current}<settingsLazy.current {result} {simId} />{/if}
      </LazySection>
    {/if}
  </div>
</div>

{#if tooltipHostLazy.current}
  <tooltipHostLazy.current
    {build}
    modelsElementId="sim-result-tooltip-models"
    resolve={resolveItem}
    {resolveTip}
  />
{/if}

<style>
  /* The result page's card heads (spec 3.2): the gold micro label over the title in caps. */
  .sim-result :global([data-card-label]) {
    color: var(--color-gold);
  }
  .sim-result :global(.heading-plan) {
    text-transform: uppercase;
  }
</style>
