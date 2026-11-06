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
  import { bandModelKey } from '../../lib/bis/page-model';
  import { isDefaultState, parseState, stateSearch, withChange, type BisPageState } from '../../lib/bis/page-state';
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
  import { BIS_SHOW_EVENT, type BisShowDetail } from '../../lib/bis/show-event';
  import type { CompareSession } from '../../lib/bis/compare-session';
  import { readStoredSources, storeSources } from '../../lib/bis/source-store';
  import type { BisPresetId, Faction, SourceFilterKind } from '../../lib/bis/types';
  import type { ControlGroup, ControlOption } from '../../lib/ui/primitives';
  import { bisContext, presetLabel } from '../../lib/bis/page-context';

  let {
    spec,
    classSlug,
    className,
    specName,
    crestSrc,
    crestColor,
    specOptions,
    raceOptions,
    guideHref,
  }: {
    spec: string;
    classSlug: string;
    className: string;
    specName: string;
    crestSrc: string;
    crestColor: string;
    specOptions: readonly ControlOption[];
    raceOptions: Readonly<Record<Faction, readonly { value: string; label: string }[]>>;
    guideHref: string;
  } = $props();

  const data = readBisPageData(spec);
  const context = bisContext(data, raceOptions);

  let state = $state<BisPageState>(parseState('', '', context));
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
  const rowsBySlot = $derived(new Map(model.rows.map((row) => [row.slot, row])));
  const ruleLine = $derived(defaultRuleLine(state.sources, state.band));
  const raceName = (faction: Faction, slug: string) =>
    raceOptions[faction].find((option) => option.value === slug)?.label ?? slug;
  const raceNote = $derived(
    state.race === model.race ? undefined : `Picks simmed as ${raceName(state.faction, model.race)}`,
  );
  const nextBand = $derived(data.bands[data.bands.indexOf(state.band) + 1]);

  const groupsBar = $derived.by((): ControlGroup[] => {
    const presets = context.presetsFor(state.faction, state.band);
    const list: ControlGroup[] = [
      { label: 'Spec', param: 'spec', kind: 'segmented', options: specOptions, value: spec },
      {
        label: 'Level',
        param: 'band',
        kind: 'segmented',
        options: data.bands.map((band) => ({ value: String(band), label: band >= 60 ? '60' : `${band}-${band + 9}` })),
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
    list.push({ label: 'Race', param: 'race', kind: 'select', options: raceOptions[state.faction], value: state.race });
    return list;
  });

  function apply(next: BisPageState, push = true): void {
    if (next.band !== state.band || next.faction !== state.faction) selected = null;
    if (next.band !== state.band || next.faction !== state.faction || next.preset !== state.preset) character = undefined;
    state = next;
    if (!push) return;
    const url = `${window.location.pathname}${stateSearch(next, context)}`;
    window.history.replaceState(null, '', url);
  }

  function onControl(param: string, value: string): void {
    if (param === 'band') apply(withChange(state, { band: Number(value) }, context));
    else if (param === 'faction') apply(withChange(state, { faction: value as Faction }, context));
    else if (param === 'preset') apply(withChange(state, { preset: value as BisPresetId }, context));
    else if (param === 'race') apply(withChange(state, { race: value }, context));
  }

  function onSources(next: SourceFilterKind[]): void {
    storeSources(state.band, next);
    apply(withChange(state, { sources: next }, context));
  }

  onMount(() => {
    const params = new URLSearchParams(window.location.search);
    let initial = parseState(window.location.search, window.location.hash, context);
    if (!params.has('sources')) {
      const stored = readStoredSources(initial.band);
      if (stored !== undefined) initial = { ...initial, sources: stored };
    }
    const hadHash = /^#band-/.test(window.location.hash);
    if (!isDefaultState(initial, context) || hadHash) apply(initial, hadHash);
    document.documentElement.removeAttribute('data-bis-pending');
    const onShow = (event: Event) => {
      const detail = (event as CustomEvent<BisShowDetail>).detail;
      apply(withChange(state, { faction: detail.faction, band: detail.band }, context));
    };
    root?.addEventListener(BIS_SHOW_EVENT, onShow);
    void import('../../lib/bis/compare-session').then(({ watchCompare }) => {
      compare = watchCompare({ build: data.build, spec, classSlug }, shownBand(), (next, race) => {
        character = next;
        if (race !== undefined && race !== state.race && context.racesFor(state.faction).includes(race)) {
          apply(withChange(state, { race }, context));
        }
      });
    });
    return () => {
      root?.removeEventListener(BIS_SHOW_EVENT, onShow);
      compare?.stop();
    };
  });
</script>

<div
  class="bis-root"
  bind:this={root}
  data-bis-spec={spec}
  data-bis-faction={state.faction}
  data-bis-band={state.band}
  data-bis-preset={state.preset}
  data-testid="bis-page"
>
  <Page layout="ledger" testid="bis-ledger">
    {#snippet header()}
      <PageHeader title={`${specName} ${className}`} tone={{ classSlug } as never} testid="bis-page-head">
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
        {#snippet meta()}<span data-testid="bis-page-sub">Best in slot</span>{/snippet}
        {#snippet answer()}
          <div class="bis-pending-hide"><BisAnswer answer={model.answer} {figure} unit={model.unit} {raceNote} /></div>
        {/snippet}
      </PageHeader>
    {/snippet}
    {#snippet bar()}
      <ControlBar groups={groupsBar} label="Best in slot controls" onchange={onControl} testid="bis-controls">
        {#snippet tail()}
          <BisSources on={state.sources} {counts} band={state.band} onchange={onSources} />
        {/snippet}
      </ControlBar>
    {/snippet}

    <div class="bis-pending-hide" data-testid={`bis-band-${state.faction}-${state.band}`}>
      <div class="bis-summary">
        <p class="bis-summary-line" data-testid="bis-summary">
          {dollSummary(cells, character, model.unit)}
          {#if character === undefined}<a href="/login" class="bis-gold">Sign in</a> to see them against your gear.{/if}
          {#if nextBand !== undefined}
            <button type="button" class="bis-link" onclick={() => onControl('band', String(nextBand))}
              >Changes from {model.label}: see {nextBand >= 60 ? '60' : `${nextBand} to ${nextBand + 9}`} ›</button
            >
          {/if}
        </p>
        <Segmented
          param="view"
          label="View"
          options={[
            { value: 'slot', label: 'By slot' },
            { value: 'where', label: 'Where to go' },
          ]}
          value={state.view}
          onchange={(_, value) => apply({ ...state, view: value === 'where' ? 'where' : 'slot' })}
          testid="bis-view"
        />
      </div>
      {#if ruleLine}<p class="bis-note" data-testid="bis-default-rule">{ruleLine}</p>{/if}
      {#if model.weights.hitCap}
        <p class="bis-note" title={model.weights.hitCap.title} data-testid="bis-hit-cap">{model.weights.hitCap.text}</p>
      {/if}
      {#if model.weaponNote}<p class="bis-note" data-testid="bis-weapon-note">{model.weaponNote}</p>{/if}

      <Plate label="Gear" testid="bis-plate">
        <Paperdoll
          {cells}
          mode={character ? 'compare' : 'pick'}
          build={data.build}
          bind:selected
          groups={groups}
          label={`${specName} ${className} best in slot, ${model.label}`}
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

      <BisWeights weights={model.weights} testid={`bis-weights-${state.faction}-${state.band}`} />
    </div>

    <p class="bis-guide-line" data-testid="bis-guide-line">
      Talents, rotation and stat priority for this spec: <a href={guideHref}>{specName} {className} guide ›</a>
    </p>
  </Page>
</div>

<style>
  .bis-crest {
    flex-shrink: 0;
    border-radius: 999px;
    box-shadow: 0 0 0 2px color-mix(in srgb, var(--c) 55%, transparent);
    background: var(--color-raised, #0d111a);
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
  .bis-guide-line {
    color: var(--color-muted);
    font-size: 13px;
    margin: 24px 0 8px;
  }
  :global(html[data-bis-pending]) .bis-pending-hide {
    visibility: hidden;
  }
</style>
