<!-- web/src/components/sim/RunReport.svelte -->
<!-- The body of one `run` result (sim spec 2026-11-06 §4.5): everything under the page head and the
     control bar. Worth knowing as Notice lines, then the split page: Damage by ability on the
     page's framed panel, Timeline, Gear, Talents, the resource and Buffs and debuffs in main; the DPS
     spread, then whatever the page adds (`railTail`: Your sims), then This run and Rotation in the rail.
     Every Section is a panel (round 5), its heading inside. Each one under the fold is its own chunk, mounted as it nears the
     viewport with its head and a skeleton holding its ready height, so nothing shifts. Every file
     it reads is optional: a row without its icon or name still reads.

     The one body for a saved run (/sim/<id>, SavedSim.svelte) and a finished run on /sim
     (SimView.svelte, `live`). The head and its answer are the page's own. -->
<script lang="ts">
  import { onMount, type Snippet } from 'svelte';
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
  import { loadBuffNames, type BuffNames } from '../../lib/sim/buff-names';
  import { plannerHrefForSpec } from '../../lib/sim/character';
  import { extraAttackLines } from '../../lib/sim/extra-attacks';
  import {
    aurasCardCopy,
    gearCardCopy,
    rotationCardCopy,
    talentsCardCopy,
  } from '../../lib/sim/input-cards-copy';
  import { withoutAbsentPets } from '../../lib/sim/pet-rows';
  import { damageCopy, spreadCopy, warningsCopy } from '../../lib/sim/result-copy';
  import { resourceLabel } from '../../lib/sim/result-sections';
  import { resultTip } from '../../lib/sim/result-tips';
  import { sectionReserves } from '../../lib/sim/section-reserve';
  import { summaryActionKeys } from '../../lib/sim/sentence';
  import { settingsCopy } from '../../lib/sim/settings-copy';
  import { NO_SPELL_INFO, loadSpellInfo, type SpellInfo } from '../../lib/sim/spell-info';
  import { timelineCopy } from '../../lib/sim/timeline-copy';
  import { benchmarkFor, type BisBenchmarks } from '../../lib/sim/bis-benchmark';
  import type { SimResult } from '../../lib/sim/types';
  import type { PrerenderedNote } from '../../lib/sim/prerender-notes';
  import { simWarnings } from '../../lib/sim/warnings';
  import { REGION } from '../../lib/ui/regions';
  import { LINK } from '../../lib/ui/primitives';
  import { yieldToMain } from '../../lib/yield';
  import AbilityTable, { type AbilityDecoration, type AbilityWeapon } from '../report/AbilityTable.svelte';
  import Notice from '../ui/Notice.svelte';
  import Page from '../ui/Page.svelte';
  import Section from '../ui/Section.svelte';
  import DpsDistribution from './DpsDistribution.svelte';
  import LazySection from './LazySection.svelte';

  let {
    result,
    live = false,
    railTail,
    onrerun,
    simId = '',
    initialNotes,
    bisBenchmarks = {},
  }: {
    result: SimResult;
    /** A finished run on /sim rather than a saved one. */
    live?: boolean;
    /** The rail's last sections, the page's own (Your sims on /sim). */
    railTail?: Snippet;
    /** Saved: a Worth knowing line's "Open in the sim". Absent on /sim, where the bar runs again. */
    onrerun?: () => void;
    /** The saved sim's id, for the downloaded file's name. */
    simId?: string;
    /** Worth knowing computed at build time for a prerendered run (lib/sim/prerender-notes.ts):
     *  drawn until the names and spell facts load here, so the lines never arrive late. */
    initialNotes?: readonly PrerenderedNote[];
    /** The published sets a run is read against (lib/sim/bis-benchmark.ts). */
    bisBenchmarks?: BisBenchmarks;
  } = $props();

  const build = activeBuild.build;
  const WARNING_ICON = 20;
  /** The spell names Heroic Strike's ranks carry: a main-hand swing it replaced is counted on that row. */
  const HEROIC_STRIKE_RANKS: readonly number[] = [78, 284, 285, 1608, 11564, 11565, 11566, 11567, 25286];
  /** A tooltip that hits several enemies: its row says the hits a cast (spec 4.7). */
  const MULTI_TARGET = /up to \d+ enemies/i;

  // svelte-ignore state_referenced_locally
  const classSlug = result.request.character.class;

  let names = $state<ActionNames | null>(null);
  let spells = $state<SpellInfo>(NO_SPELL_INFO);
  let spellsLoaded = $state(false);
  let namesSettled = $state(false);
  let items = $state<ReadonlyMap<number, Item>>(new Map());
  let buffNames = $state<BuffNames | null>(null);

  const summary = $derived(withoutAbsentPets(result.summary));
  const reserves = $derived(sectionReserves(result));
  const request = $derived(result.request);
  const character = $derived(request.character);
  const powerTitle = $derived(resourceLabel(result));

  /** Runs one optional load; a failure is logged and the page keeps its fallback. */
  function loadQuietly(what: string, load: () => Promise<void>): void {
    load().catch((error: unknown) => console.error(`sim result: ${what} did not load`, error));
  }

  onMount(() => {
    loadQuietly('the action names', async () => {
      try {
        const loaded = await loadActionNames(build, classSlug);
        names = loaded;
        names = await withShardNames(loaded, summaryActionKeys(summary));
      } finally {
        namesSettled = true;
      }
    });
    loadQuietly('the spell data', async () => {
      spells = await loadSpellInfo(build, classSlug);
      spellsLoaded = true;
    });
    loadQuietly('the item file', async () => {
      const file = await loadItems(build, classSlug);
      items = new Map(file.items.map((item) => [item.id, item]));
    });
    loadQuietly('the buff table', async () => {
      buffNames = await loadBuffNames(build);
    });
  });

  const computedNotes = $derived(
    simWarnings({ result, names, fact: spells.fact }).map((note) => ({
      ...note,
      iconUrl: note.actionKey === undefined ? null : spells.iconUrl(note.actionKey),
    })),
  );
  const notes = $derived(
    initialNotes !== undefined && (!namesSettled || !spellsLoaded) ? initialNotes : computedNotes,
  );
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
      targets:
        fact !== undefined && MULTI_TARGET.test(fact.tip.description) ? request.encounter.targets : null,
    };
  }

  // Below the fold each section is its own chunk; the shared tooltip host loads after first paint.
  const timelineLazy = createLazyComponent(() => import('./SimTimeline.svelte'));
  const gearLazy = createLazyComponent(() => import('./SimGearSection.svelte'));
  const talentsLazy = createLazyComponent(() => import('./SimTalentsSection.svelte'));
  const resourceLazy = createLazyComponent(() => import('./SimResourceSection.svelte'));
  const buffsLazy = createLazyComponent(() => import('./SimBuffsSection.svelte'));
  const rotationLazy = createLazyComponent(() => import('./SimRotationSection.svelte'));
  const thisRunLazy = createLazyComponent(() => import('./SimThisRunSection.svelte'));
  const tooltipHostLazy = createLazyComponent(() => import('../BisTooltipHost.svelte'));

  // The page draws in two tasks: the table and the spread first, then everything under them. One
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

{#snippet rail()}
  <Section panel title={spreadCopy.title} testid="sim-spread-card">
    <DpsDistribution
      estimate={result.dps}
      iterationsRun={result.iterations_run}
      histogram={result.dps_histogram}
      spec={request.spec}
      benchmark={benchmarkFor(result, bisBenchmarks)}
    />
  </Section>
  <!-- Your sims under the spread, so a reader compares the number before the inputs and the
       rotation (ux-designer, round 5). -->
  {@render railTail?.()}
  {#if drawnRest}
    <LazySection
      id="settings"
      title={settingsCopy.title}
      lazy={thisRunLazy}
      ready={thisRunLazy.current !== null}
      reserve={reserves.settings}
    >
      {#if thisRunLazy.current}<thisRunLazy.current {result} {simId} />{/if}
    </LazySection>
    <LazySection
      id="rotation"
      title={rotationCardCopy.title}
      lazy={rotationLazy}
      ready={rotationLazy.current !== null}
      reserve={reserves.rotation}
    >
      {#if rotationLazy.current}<rotationLazy.current {request} {result} {names} />{/if}
    </LazySection>
  {/if}
{/snippet}

<div class="sim-body" data-testid={live ? 'sim-live-report' : 'sim-saved-run'}>
  {#if notes.length > 0}
    <div class="px-[var(--page-inset)]" data-region={REGION.summary} data-testid="sim-worth-knowing">
      {#each notes as note (`${note.id}:${note.actionKey ?? ''}`)}
        {@const icon = note.iconUrl}
        <Notice
          text={note.text}
          tone="neutral"
          action={note.action === 'planner'
            ? { label: warningsCopy.openInPlanner, href: plannerHref }
            : undefined}
          testid={`sim-note-${note.id}`}
        >
          {#snippet lead()}
            <span class="block h-5 w-5 overflow-hidden rounded-[4px]"
              >{#if icon !== null}<img
                  src={icon}
                  alt=""
                  width={WARNING_ICON}
                  height={WARNING_ICON}
                />{/if}</span
            >
          {/snippet}
          {#snippet tail()}{#if note.action === 'sim' && onrerun !== undefined}<button
                type="button"
                class="{LINK} whitespace-nowrap"
                onclick={onrerun}>{warningsCopy.openInSim} ›</button
              >{/if}{/snippet}
        </Notice>
      {/each}
    </div>
  {/if}

  <Page layout="split" panels {rail}>
    <Section title={damageCopy.title} panel frame="framed" flush testid="sim-damage-card">
      <AbilityTable {table} {decorate} />
    </Section>
    {#if drawnRest}
      <LazySection
        id="timeline"
        title={timelineCopy.title}
        lazy={timelineLazy}
        ready={timelineLazy.current !== null}
        reserve={reserves.timeline}
      >
        {#if timelineLazy.current}<timelineLazy.current {result} {names} {live} />{/if}
      </LazySection>
      <LazySection
        id="gear"
        title={gearCardCopy.title}
        lazy={gearLazy}
        ready={gearLazy.current !== null}
        reserve={reserves.gear}
      >
        {#if gearLazy.current}<gearLazy.current {request} {result} />{/if}
      </LazySection>
      <LazySection
        id="talents"
        title={talentsCardCopy.title}
        lazy={talentsLazy}
        ready={talentsLazy.current !== null}
        reserve={reserves.talents}
      >
        {#if talentsLazy.current}<talentsLazy.current {request} />{/if}
      </LazySection>
      {#if powerTitle !== null}
        <LazySection
          id="resource"
          title={powerTitle}
          lazy={resourceLazy}
          ready={resourceLazy.current !== null}
          reserve={reserves.resource}
        >
          {#if resourceLazy.current}<resourceLazy.current {result} {names} />{/if}
        </LazySection>
      {/if}
      <LazySection
        id="buffs"
        title={aurasCardCopy.title}
        lazy={buffsLazy}
        ready={buffsLazy.current !== null}
        reserve={reserves.buffs}
      >
        {#if buffsLazy.current}<buffsLazy.current {result} {names} />{/if}
      </LazySection>
    {/if}
  </Page>
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
  /* Phone and tablet (spec 8): the rail interleaves, so the spread follows the damage table and
     This run and Rotation close the page. Main and rail give their sections to one column. */
  @media (max-width: 1023px) {
    .sim-body :global(.ui-page) {
      display: flex;
      flex-direction: column;
      align-items: stretch;
      gap: var(--section-gap);
    }
    .sim-body :global(.ui-page-main),
    .sim-body :global(.ui-page-rail) {
      display: contents;
    }
    .sim-body :global([data-testid='sim-spread-card']) {
      order: 1;
    }
    .sim-body :global(.ui-page-main > *:not(:first-child)) {
      order: 2;
    }
    .sim-body :global(.ui-page-rail > *:not(:first-child)) {
      order: 3;
    }
  }
</style>
