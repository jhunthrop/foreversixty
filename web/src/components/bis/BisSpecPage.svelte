<!-- web/src/components/bis/BisSpecPage.svelte -->
<!-- The Best in slot page (web layout spec 2026-11-05 §5): the page head with the set's figure,
     the one control bar, the summary line with the By slot | Where to go toggle, the paperdoll on
     its plate, What to stack, and the guide line. One island owns the state (band, faction,
     preset, race, view, sources) from the URL; it is server-rendered at the default state so first
     paint needs no script, and a deep link's non-default state is applied on mount behind
     html[data-bis-pending] (geometry kept, content hidden) so nothing visibly swaps. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import ControlBar from '../ui/ControlBar.svelte';
  import Segmented from '../ui/Segmented.svelte';
  import Plate from '../ui/Plate.svelte';
  import Paperdoll from '../ui/Paperdoll.svelte';
  import BisAnswer from './BisAnswer.svelte';
  import BisSlotDetail from './BisSlotDetail.svelte';
  import BisSources from './BisSources.svelte';
  import BisWeights from './BisWeights.svelte';
  import { readBisPageData } from '../../lib/bis/page-registry';
  import { bandModelKey } from '../../lib/bis/model-key';
  import {
    defaultSources,
    isDefaultState,
    parseState,
    stateSearch,
    withChange,
    type BisPageState,
  } from '../../lib/bis/page-state';
  import { choiceWords, defaultRuleLine, setFigureFor, sourceChoiceFor } from '../../lib/bis/source-choice';
  import {
    changedCount,
    dollCellsFor,
    dollSummary,
    filteredPicks,
    sourceCounts,
    whereGroups,
    type DollCharacter,
  } from '../../lib/bis/doll-view';
  import { BIS_SHOW_EVENT, takePendingShow, type BisShowDetail } from '../../lib/bis/show-event';
  import type { CompareSession } from '../../lib/bis/compare-session';
  import { readStoredSources, storeSources } from '../../lib/bis/source-store';
  import type { BisPresetId, Faction, SourceFilterKind } from '../../lib/bis/types';
  import type { ControlGroup, ControlOption } from '../../lib/ui/primitives';
  import { SPECS } from '../../lib/sim/specs';
  import { REGION } from '../../lib/ui/regions';
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
    crestSrc,
    crestColor,
    specOptions,
    raceOptions,
    guideHref,
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
    crestSrc: string;
    crestColor: string;
    specOptions: readonly ControlOption[];
    raceOptions: Readonly<Record<Faction, readonly { value: string; label: string }[]>>;
    guideHref: string;
    /** The page's full data (every band); the HTML carries only the default state's slice. */
    dataHref: string;
  } = $props();

  const inline = readBisPageData(spec);
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
    data = (await response.json()) as typeof inline;
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
  const groups = $derived(state.view === 'where' ? whereGroups(model, cells) : undefined);
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
  const rowsBySlot = $derived(new Map(model.rows.map((row) => [row.slot, row])));
  const ruleLine = $derived(defaultRuleLine(state.sources, state.band));
  const raceName = (faction: Faction, slug: string) =>
    raceOptions[faction].find((option) => option.value === slug)?.label ?? slug;
  const raceNote = $derived(
    state.race === model.race ? undefined : `Picks simmed as ${raceName(state.faction, model.race)}`,
  );
  const nextBand = $derived(data.bands[data.bands.indexOf(state.band) + 1]);

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

  /** The tallest answer variant measured at 1024 to 2000 px: a healer's figure with its metrics row
   *  (183 px; its trust line wraps on the real run notes). Every spec of a class keeps its title and
   *  control bar at the same y, so a class with a healer spec takes the healer floor; a class
   *  without one keeps the mocks' 165 px. */
  const HEAD_ROW_MIN_HEIGHT_DEFAULT = '165px';
  const HEALER_HEAD_ROW_MIN_HEIGHT = '183px';
  const HEAD_ROW_MIN_HEIGHT = SPECS.some((s) => s.class_slug === classSlug && s.role === 'healer')
    ? HEALER_HEAD_ROW_MIN_HEIGHT
    : HEAD_ROW_MIN_HEIGHT_DEFAULT;
</script>

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
    {#snippet header()}
      <PageHeader
        title={bisSeoCopy.heading(seoSubject)}
        tone={{ classSlug } as never}
        minRowHeight={HEAD_ROW_MIN_HEIGHT}
        testid="bis-page-head"
      >
        {#snippet crest()}
          <img
            src={crestSrc}
            alt=""
            width="56"
            height="56"
            class="bis-crest"
            style={`--c: ${crestColor}`}
            data-testid={`class-crest-${classSlug}`}
          />
        {/snippet}
        {#snippet meta()}<span data-testid="bis-page-sub">{bisSeoCopy.meta(seoSubject, state.band)}</span
          >{/snippet}
        {#snippet answer()}
          <div class="bis-pending-hide">
            <BisAnswer answer={model.answer} {figure} unit={model.unit} {raceNote} />
          </div>
        {/snippet}
      </PageHeader>
    {/snippet}
    {#snippet bar()}
      <ControlBar
        groups={groupsBar}
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
      <div class="bis-summary" data-region={REGION.summary}>
        <p class="bis-summary-line" data-testid="bis-summary">
          {dollSummary(cells, character, model.unit)}
          {#if character === undefined}<a href="/login" class="bis-gold">Sign in</a> to see them against your gear.{/if}
          {#if nextBand !== undefined}
            <a
              class="bis-link"
              href={bandHref(nextBand)}
              onclick={(event) => {
                if (event.metaKey || event.ctrlKey || event.shiftKey || event.altKey) return;
                event.preventDefault();
                onControl('band', String(nextBand));
              }}
              >Changes from {model.label}: see {nextBand >= 60 ? '60' : `${nextBand} to ${nextBand + 9}`} ›</a
            >
          {:else if preRaidToggle}
            {@const target = {
              band: state.band,
              sources: preRaidView ? defaultSources(state.band) : [...PRE_RAID_KINDS],
            }}
            <a
              class="bis-link"
              href={stateHref(target)}
              data-testid="bis-pre-raid-toggle"
              onclick={(event) => {
                if (event.metaKey || event.ctrlKey || event.shiftKey || event.altKey) return;
                event.preventDefault();
                whenReady(() => apply(withChange(state, { sources: target.sources }, context)));
              }}>{preRaidView ? preRaidCopy.turnOff : preRaidCopy.turnOn} ›</a
            >
          {/if}
        </p>
        <Segmented
          label="View"
          options={[
            { value: 'slot', label: 'By slot' },
            { value: 'where', label: 'Where to go' },
          ]}
          value={state.view}
          onChange={(value) => apply({ ...state, view: value === 'where' ? 'where' : 'slot' })}
          testid="control-view"
        />
      </div>
      {#if ruleLine || model.weights.hitCap || model.weaponNote}
        <div data-region={REGION.notes}>
          {#if ruleLine}<p class="bis-note" data-testid="bis-default-rule">{ruleLine}</p>{/if}
          {#if model.weights.hitCap}
            <p class="bis-note" title={model.weights.hitCap.title} data-testid="bis-hit-cap">
              {model.weights.hitCap.text}
            </p>
          {/if}
          {#if model.weaponNote}<p class="bis-note" data-testid="bis-weapon-note">{model.weaponNote}</p>{/if}
        </div>
      {/if}

      <Plate label="Gear" testid="bis-plate">
        <Paperdoll
          {cells}
          mode={character ? 'compare' : 'pick'}
          build={data.build}
          bind:selected
          {groups}
          label={`${fullName} best in slot, ${model.label}`}
          testid="bis-doll"
        >
          {#snippet detail(cell, placement)}
            <BisSlotDetail
              {cell}
              row={rowsBySlot.get(cell.slot)}
              items={{ ...data.items, ...(character?.items ?? {}) }}
              unit={model.unit}
              build={data.build}
              simHref={model.answer.simHref}
              {placement}
            />
          {/snippet}
        </Paperdoll>
      </Plate>

      <div class="bis-stack">
        <BisWeights weights={model.weights} testid={`bis-weights-${state.faction}-${state.band}`} />
      </div>
    </div>

    <p class="bis-guide-line" data-region={REGION.footerNote} data-testid="bis-guide-line">
      Talents, rotation and stat priority for this spec: <a href={guideHref}>{fullName} guide ›</a>
    </p>
  </Page>
</div>

<style>
  .bis-crest {
    flex-shrink: 0;
    border-radius: 999px;
    box-shadow: 0 0 0 2px color-mix(in srgb, var(--c) 55%, transparent);
    background: var(--color-bg);
  }
  .bis-summary {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    justify-content: space-between;
    gap: 10px 16px;
    padding: 10px 0;
  }
  .bis-summary-line {
    font-size: 15px;
    line-height: 24px;
    min-width: 0;
  }
  .bis-gold,
  .bis-link,
  .bis-guide-line a {
    color: var(--color-gold);
    font-weight: 600;
  }
  .bis-link {
    margin-left: 6px;
  }
  .bis-link:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .bis-note {
    color: var(--color-muted);
    font-size: 13px;
    line-height: 20px;
    margin-bottom: 8px;
  }
  .bis-stack {
    margin-top: 32px;
  }
  /* The Page's section gap already sits above it (it is the ledger's last flex item). */
  .bis-guide-line {
    color: var(--color-muted);
    font-size: 13px;
    line-height: 20px;
    margin: 0 0 8px;
  }
  :global(html[data-bis-pending]) .bis-pending-hide {
    visibility: hidden;
  }
</style>
