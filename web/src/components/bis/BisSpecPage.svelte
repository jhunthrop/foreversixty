<!-- web/src/components/bis/BisSpecPage.svelte -->
<!-- The Best in slot page on the in-game paperdoll (design/specs/2026-11-09-directions.md,
     Paperdoll and Best in slot): the band holds the nav and the one control bar; the frame's title
     bar is the page's h1 (the spec, the band and faction, the set's figure), its centre the
     selected slot's client tooltip, what it is worth and Open in the sim. Signed in, the
     character's Battle.net render stands in the centre and those lines go in a bar under the frame.
     Beside the frame: the set's figures (a tank, a healer), What to stack and Where to go; below
     it the itinerary. Under 423 px the frame is its strip and two-column list. One island owns the
     state (band, faction, preset, race, sources) from the URL; it is server-rendered at the
     default state so first paint needs no script, and a deep link's non-default state is applied
     on mount behind html[data-bis-pending] (geometry kept, content hidden) so nothing swaps. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import Page from '../ui/Page.svelte';
  import ControlBar from '../ui/ControlBar.svelte';
  import PaperdollFrame from '../ui/PaperdollFrame.svelte';
  import ItemTooltip from '../ItemTooltip.svelte';
  import Plate from '../ui/Plate.svelte';
  import BisSources from './BisSources.svelte';
  import BisSlotLines from './BisSlotLines.svelte';
  import BisSide from './BisSide.svelte';
  import BisItinerary from './BisItinerary.svelte';
  import { addTooltipModels, readBisPageData, readTooltipModels } from '../../lib/bis/page-registry';
  import { bandModelKey } from '../../lib/bis/model-key';
  import type { BisFullData } from '../../lib/bis/page-model';
  import type { ItemTooltipModel } from '../../lib/items/tooltip';
  import {
    defaultSources,
    isDefaultState,
    parseState,
    stateSearch,
    withChange,
    type BisPageState,
  } from '../../lib/bis/page-state';
  import { choiceWords, setFigureFor, sourceChoiceFor } from '../../lib/bis/source-choice';
  import {
    changedCount,
    dollCellsFor,
    dollSummary,
    filteredPicks,
    sourceCounts,
    type DollCharacter,
  } from '../../lib/bis/doll-view';
  import { itinerary, simWeightsHref, stackView } from '../../lib/bis/frame-view';
  import { bisCopy } from '../../lib/bis/copy';
  import { defaultSelection, type DollCell } from '../../lib/ui/paperdoll';
  import { BIS_SHOW_EVENT, takePendingShow, type BisShowDetail } from '../../lib/bis/show-event';
  import type { CompareSession } from '../../lib/bis/compare-session';
  import { readStoredSources, storeSources } from '../../lib/bis/source-store';
  import type { BisPresetId, Faction, SourceFilterKind } from '../../lib/bis/types';
  import { controlSummary, type ControlGroup, type ControlOption } from '../../lib/ui/primitives';
  import { bisContext, presetLabel } from '../../lib/bis/page-context';
  import { bandSlug, bisBandPath } from '../../lib/bis/band-path';
  import { bisSeoCopy } from '../../lib/bis/seo-copy';
  import {
    PRE_RAID_KINDS,
    PRE_RAID_SLUG,
    isPreRaidView,
    preRaidCopy,
    preRaidOffered,
  } from '../../lib/bis/pre-raid';
  import { specLabel } from '../../lib/sim/spec-label';

  let {
    spec,
    classSlug,
    specSlug,
    pathBand,
    pathPreRaid = false,
    preRaid,
    crestColor,
    specOptions,
    raceOptions,
    dataHref,
  }: {
    spec: string;
    classSlug: string;
    specSlug: string;
    /** The band this page's URL names (lib/bis/band-path.ts); the server renders it. */
    pathBand: number;
    /** The URL is the band's pre-raid list, /bis/<class>/<spec>/60/pre-raid (lib/bis/pre-raid.ts). */
    pathPreRaid?: boolean;
    /** The level 60 list holds no raid drop: its h1 says Pre-Raid (lib/bis/seo-copy.ts). */
    preRaid: boolean;
    /** The class colour, for the signed-in character's name in the centre. */
    crestColor: string;
    specOptions: readonly ControlOption[];
    raceOptions: Readonly<Record<Faction, readonly { value: string; label: string }[]>>;
    /** The page's full data (every band); the HTML carries only the default state's slice. */
    dataHref: string;
  } = $props();

  const inline = readBisPageData(spec);
  const inlineTooltips = readTooltipModels(spec);
  let lateTooltips = $state.raw<ReadonlyMap<number, ItemTooltipModel>>(new Map());
  let data = $state.raw(inline);
  const context = $derived(bisContext(data, raceOptions, pathBand, pathPreRaid));
  /** A band's own page: the spec's path for its first band, else /bis/<class>/<spec>/<band>. */
  const bandHref = (band: number): string => bisBandPath(classSlug, specSlug, band, data.bands[0] ?? band);
  /** A state's page: its band's, and /pre-raid under 60 for the pre-raid list. */
  const stateHref = (next: Pick<BisPageState, 'band' | 'sources'>): string =>
    `${bandHref(next.band)}${isPreRaidView(next.band, next.sources) ? `/${PRE_RAID_SLUG}` : ''}`;
  /** Settles once the full data is in (at once when the page inlined all of it). */
  let ready: Promise<void> = Promise.resolve();

  async function loadFull(): Promise<void> {
    if (!inline.partial) return;
    const response = await fetch(dataHref);
    if (!response.ok) throw new Error(`${dataHref} answered ${response.status}`);
    const { tooltips: more, ...full } = (await response.json()) as BisFullData;
    lateTooltips = new Map(addTooltipModels(more));
    data = full;
  }

  /** Runs a state change once the full data is in; a failed load keeps the default state shown. */
  function whenReady(change: () => void): void {
    void ready.then(change, () => {});
  }

  let state = $state<BisPageState>(
    parseState('', '', bisContext(inline, raceOptions, pathBand, pathPreRaid)),
  );
  let selected = $state<string | null>(null);
  let character = $state<DollCharacter | undefined>(undefined);
  let root = $state<HTMLElement | undefined>(undefined);
  let compare: CompareSession | undefined;

  function shownBand() {
    return { band: state.band, faction: state.faction, preset: state.preset };
  }

  $effect(() => {
    compare?.show(shownBand());
  });

  const model = $derived(data.models[bandModelKey(state.faction, state.band, state.preset)]!);
  const choice = $derived(sourceChoiceFor(state.sources, state.band));
  const picks = $derived(filteredPicks(model, choice));
  const changed = $derived(changedCount(model, picks));
  const deps = $derived({ items: data.items, unit: model.unit, character });
  const cells = $derived(dollCellsFor(model, deps, picks, choiceWords(state.sources, state.band)));
  const stops = $derived(itinerary(model, cells));
  const stack = $derived(stackView(model));
  const figure = $derived(setFigureFor(model, state.sources, changed));
  const counts = $derived(sourceCounts(model));
  const preRaidView = $derived(isPreRaidView(state.band, state.sources));
  /** The Pre-raid toggle shows by itself once a level 60 band holds raid or PvP gear. */
  const preRaidToggle = $derived(
    preRaidOffered(
      state.band,
      model.rows.map((row) => row.filter),
    ),
  );
  // The spec's full name, the catalogue's own ("Feral Druid (Bear)"): the h1, the list's label, the guide link.
  const fullName = $derived(specLabel(spec));
  const seoSubject = $derived({ fullName, preRaid: preRaid || preRaidView });
  const notes = $derived(
    [
      model.weights.hitCap
        ? { text: model.weights.hitCap.text, title: model.weights.hitCap.title, testid: 'bis-hit-cap' }
        : undefined,
      model.weights.expertiseCap
        ? {
            text: model.weights.expertiseCap.text,
            title: model.weights.expertiseCap.title,
            testid: 'bis-expertise-cap',
          }
        : undefined,
    ].filter((note) => note !== undefined),
  );
  /** The selected slot, else the doll's default (the first upgrade, else the first pick). */
  const shownSlot = $derived(
    selected !== null && cells.some((cell) => cell.slot === selected) ? selected : defaultSelection(cells),
  );
  const preRaidAction = $derived.by(() => {
    if (!preRaidToggle) return undefined;
    const target = {
      band: state.band,
      sources: preRaidView ? defaultSources(state.band) : [...PRE_RAID_KINDS],
    };
    return {
      label: preRaidView ? preRaidCopy.turnOff : preRaidCopy.turnOn,
      href: stateHref(target),
      onclick: (event: MouseEvent) => {
        if (event.metaKey || event.ctrlKey || event.shiftKey || event.altKey) return;
        event.preventDefault();
        whenReady(() => apply(withChange(state, { sources: target.sources }, context)));
      },
    };
  });

  /** Fits "Night Elf"; the rare long race name ("High Order Skyborne") truncates in the field only. */
  const RACE_FIELD_PX = 112;

  const groupsBar = $derived.by((): ControlGroup[] => {
    const presets = context.presetsFor(state.faction, state.band);
    const list: ControlGroup[] = [
      { label: 'Spec', param: 'spec', kind: 'segmented', options: specOptions, value: spec },
      {
        label: 'Level',
        param: 'band',
        kind: 'segmented',
        options: data.bands.map((band) => ({
          value: String(band),
          label: bandSlug(band),
          href: bandHref(band),
          inPlace: true,
        })),
        value: String(state.band),
      },
      {
        label: 'Faction',
        param: 'faction',
        kind: 'segmented',
        options: [
          { value: 'alliance', label: 'Alliance' },
          { value: 'horde', label: 'Horde' },
        ],
        value: state.faction,
      },
    ];
    if (presets.length > 1) {
      list.push({
        label: 'Profile',
        param: 'preset',
        kind: 'segmented',
        options: presets.map((preset) => ({ value: preset, label: presetLabel(preset) })),
        value: state.preset,
      });
    }
    list.push({
      label: 'Race',
      param: 'race',
      kind: 'select',
      options: raceOptions[state.faction],
      value: state.race,
      fieldWidth: RACE_FIELD_PX,
    });
    return list;
  });
  /** The phone line: range, faction and profile. The spec is the doll's title and the race sits in
   *  the open stack, so the line fits 390 without cutting a word. */
  const barSummary = $derived(
    controlSummary(groupsBar.filter((group) => group.param !== 'spec' && group.param !== 'race')),
  );

  function apply(next: BisPageState, push = true): void {
    if (data.models[bandModelKey(next.faction, next.band, next.preset)] === undefined) return;
    if (next.band !== state.band || next.faction !== state.faction) selected = null;
    if (next.band !== state.band || next.faction !== state.faction || next.preset !== state.preset)
      character = undefined;
    state = next;
    if (!push) return;
    const url = `${stateHref(next)}${stateSearch(next, context)}`;
    window.history.replaceState(null, '', url);
  }

  function onControl(param: string, value: string): void {
    whenReady(() => control(param, value));
  }

  function control(param: string, value: string): void {
    if (param === 'band') apply(withChange(state, { band: Number(value) }, context));
    else if (param === 'faction') apply(withChange(state, { faction: value as Faction }, context));
    else if (param === 'preset') apply(withChange(state, { preset: value as BisPresetId }, context));
    else if (param === 'race') apply(withChange(state, { race: value }, context));
  }

  function onSources(next: SourceFilterKind[]): void {
    storeSources(state.band, next);
    whenReady(() => apply(withChange(state, { sources: next }, context)));
  }

  function applyUrlState(): void {
    const params = new URLSearchParams(window.location.search);
    let initial = parseState(window.location.search, window.location.hash, context);
    if (!params.has('sources') && !pathPreRaid) {
      const stored = readStoredSources(initial.band);
      if (stored !== undefined) initial = { ...initial, sources: stored };
    }
    // A legacy #band-<faction>-<band> link keeps its URL: the hash is read, not rewritten.
    if (!isDefaultState(initial, context)) apply(initial, false);
  }

  onMount(() => {
    ready = loadFull();
    void ready
      .then(applyUrlState, (error: unknown) => console.error('BiS page data failed to load', error))
      .finally(() => document.documentElement.removeAttribute('data-bis-pending'));
    const show = (detail: BisShowDetail) => {
      if (root !== undefined) delete root.dataset.bisShow;
      if (!data.bands.includes(detail.band)) return;
      whenReady(() => apply(withChange(state, { faction: detail.faction, band: detail.band }, context)));
    };
    const onShow = (event: Event) => show((event as CustomEvent<BisShowDetail>).detail);
    root?.addEventListener(BIS_SHOW_EVENT, onShow);
    const pending = root === undefined ? undefined : takePendingShow(root);
    if (pending !== undefined) show(pending);
    void import('../../lib/bis/compare-session').then(({ watchCompare }) => {
      compare = watchCompare({ build: data.build, spec, classSlug }, shownBand(), (next, race) => {
        character = next;
        if (race !== undefined && race !== state.race && context.racesFor(state.faction).includes(race)) {
          whenReady(() => apply(withChange(state, { race }, context)));
        }
      });
    });
    return () => {
      root?.removeEventListener(BIS_SHOW_EVENT, onShow);
      compare?.stop();
    };
  });
</script>

{#snippet simLink()}
  {#if model.answer.simHref}<a class="bis-sim-link" href={model.answer.simHref} data-testid="bis-sim-link"
      >{bisCopy.openInSim} ›</a
    >{/if}
{/snippet}

{#snippet centre(cell: DollCell | undefined, form: 'frame' | 'list')}
  {@const shown = cell?.item ?? cell?.coveredBy}
  {@const tip = shown ? (inlineTooltips.get(shown.id) ?? lateTooltips.get(shown.id)) : undefined}
  {#if form === 'frame' && character?.render}
    <p class="bis-who" style:--c={crestColor} data-testid="bis-who">
      <b>{character.name}</b>
      <small>{bisCopy.whoLine(character.level, character.raceClass)}</small>
      <small class="bis-who-as">{bisCopy.asLastLoggedOut}</small>
    </p>
  {/if}
  {#if tip}
    <div class="bis-tip" data-form={form} data-picked={selected !== null ? '' : undefined}>
      <ItemTooltip model={tip} build={data.build} floating={false}>
        {#snippet tail()}<BisSlotLines {cell} />{/snippet}
      </ItemTooltip>
    </div>
  {/if}
  {#if character === undefined}{@render simLink()}{/if}
{/snippet}

<div
  class="bis-root"
  bind:this={root}
  data-bis-spec={spec}
  data-bis-faction={state.faction}
  data-bis-band={state.band}
  data-bis-preset={state.preset}
  data-bis-path-band={pathBand !== data.bands[0] ? pathBand : undefined}
  data-testid="bis-page"
>
  <Page layout="ledger" testid="bis-ledger">
    {#snippet bar()}
      <ControlBar
        groups={groupsBar}
        summary={barSummary}
        label="Best in slot controls"
        labelsFrom={1920}
        onchange={onControl}
        testid="bis-controls"
      >
        {#snippet tail()}
          <BisSources on={state.sources} {counts} band={state.band} onchange={onSources} />
        {/snippet}
      </ControlBar>
    {/snippet}

    <div class="bis-pending-hide" data-testid={`bis-band-${state.faction}-${state.band}`}>
      <div class="bis-top">
        <div class="bis-doll">
          <PaperdollFrame
            {cells}
            build={data.build}
            title={bisSeoCopy.heading(seoSubject)}
            shortTitle={fullName}
            meta={bisCopy.frameMeta(model.label, state.band >= 60 && seoSubject.preRaid)}
            figure={figure ? { value: figure.value, unit: model.unit } : undefined}
            avatar={character?.avatar ?? null}
            render={character?.render ?? null}
            heading="h1"
            selected={shownSlot}
            onselect={(slot) => (selected = slot)}
            {centre}
            label={`${fullName} best in slot, ${model.label}`}
            testid="bis-doll"
          />
          {#if character !== undefined}
            <div class="bis-after" data-testid="bis-after">
              <Plate flush={false}>
                <div class="bis-after-row">
                  <p data-testid="bis-summary">{dollSummary(cells, character, model.unit)}</p>
                  {@render simLink()}
                </div>
              </Plate>
            </div>
          {/if}
        </div>
        <BisSide
          {model}
          {stack}
          {stops}
          {notes}
          weightsHref={simWeightsHref(model.answer.simHref)}
          whereAction={preRaidAction}
        />
      </div>
      <div class="bis-route">
        <BisItinerary {stops} build={data.build} label={`${fullName}, where to go`} />
      </div>
    </div>
  </Page>
</div>

<style>
  /* The phone list shows a slot's tooltip once the player taps it, never one left open by default. */
  @media (max-width: 767px) {
    .bis-tip[data-form='list']:not([data-picked]) {
      display: none;
    }
  }
  .bis-top {
    display: grid;
    grid-template-columns: minmax(0, 1fr);
    gap: 28px;
    align-items: start;
  }
  .bis-doll {
    width: 100%;
    max-width: 640px;
    justify-self: center;
  }
  @media (min-width: 1100px) {
    .bis-top {
      grid-template-columns: minmax(0, 640px) minmax(0, 1fr);
      gap: 36px;
    }
    .bis-doll {
      justify-self: start;
    }
  }
  /* The wide frame (from 1760): the doll grows to 760 and the side takes the rest of the frame, so
     the pair sits on the frame's edges as the bar above and the itinerary below do. */
  @media (min-width: 1760px) {
    .bis-top {
      grid-template-columns: minmax(0, 760px) minmax(0, 1fr);
      column-gap: 64px;
    }
    .bis-doll {
      max-width: 760px;
    }
  }
  /* The name stands at the top of the centre, the tooltip at its foot (the centre is a column
     packed to the end), so the model's face shows between them. */
  .bis-who {
    margin-bottom: auto;
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
    color: var(--c);
    font-size: 15px;
    line-height: 20px;
    text-shadow:
      0 1px 2px black,
      0 0 6px black;
  }
  .bis-who small {
    color: var(--color-text);
    font-size: 13px;
  }
  .bis-who .bis-who-as {
    font-style: italic;
    color: var(--color-muted-card);
  }
  .bis-after {
    margin-top: 12px;
  }
  .bis-after-row {
    display: flex;
    flex-wrap: wrap;
    align-items: baseline;
    justify-content: space-between;
    gap: 8px 16px;
    font-size: 14px;
    line-height: 20px;
  }
  .bis-sim-link {
    color: var(--color-gold);
    font-size: 14px;
    font-weight: 600;
    white-space: nowrap;
    text-shadow: 0 1px 2px black;
  }
  .bis-route {
    margin-top: var(--section-gap);
  }
  :global(html[data-bis-pending]) .bis-pending-hide {
    visibility: hidden;
  }
</style>
