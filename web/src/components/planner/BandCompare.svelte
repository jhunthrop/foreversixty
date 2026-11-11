<!-- web/src/components/planner/BandCompare.svelte -->
<!-- The Level build section (spec 2026-11-08 section 4.6, rulings Q1 and MUST 3): the build
     against its level band's published build. One line: `Matches` in green; the difference named
     (`Differs by 2 points: this build takes Rapid Killing 2/2; the level 60 build takes ...`);
     the band points still to take; or, for a character with more points than the band's build
     covers, how far it reaches and which build the rest follow -- never `Matches` then. Load (with
     its confirm over spent points), See what differs, and the band's gear on the BiS page. The
     band's own set figure retired (decision P1): the head's live score now runs on this band's
     gear, which this section hands up through `onband`. The band file is `bis/<spec>.json`,
     fetched once per spec (`fetchBisFile`); there is no band picker here (owner ruling 12). -->
<script lang="ts">
  import EmptyState from '../ui/EmptyState.svelte';
  import { bandForLevel, bandEntryFor, bisPageHref, fetchBisFile } from '../../lib/bis/hover';
  import type { BisBand, BisFile, Faction } from '../../lib/bis/types';
  import {
    bandCoverage,
    bandDiffNames,
    bandDiffStatus,
    bandForPanel,
    bandPointTotal,
    bandTitleForDraft,
    diffAgainstBand,
    loadFromBandAtLevel,
    type BandDiffView,
  } from '../../lib/planner/band-compare';
  import { bandRangeLabel } from '../../lib/planner/band-title';
  import { bandCompareCopy } from '../../lib/planner/copy';
  import { levelReached, pointsAvailableAt } from '../../lib/planner/derive';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { siblingSpecNames, specDisplayName } from '../../lib/sim/spec-label';
  import { specKeyFor } from '../../lib/addon/score';
  import { LINK } from '../../lib/ui/primitives';
  import type { LadderOrderFor } from '../../lib/planner/ladder-load';
  import Button from '../ui/Button.svelte';
  import Section from '../ui/Section.svelte';
  import Tag from '../ui/Tag.svelte';

  let {
    store,
    specOverride,
    characterLevel = null,
    onbanddiff,
    onquiet,
    onloaded,
    onband,
    heading,
    quietMatch = false,
    panel = false,
  }: {
    /** The heading inside its panel (the standalone planner's rail and column). */
    panel?: boolean;
    store: PlannerStore;
    /** `?spec=` (spec §4.I): names which spec's `bis/<spec>.json` this panel reads,
     *  overriding the points-based inference for the life of this mount. */
    specOverride?: string;
    /** The selected character's level; the load cuts the band's point order at it. */
    characterLevel?: number | null;
    /** Fires with the panel's own diff view (or `null` once there is none) so the tree row
     *  can draw the per-cell "differs from this band" marker (spec §4.E.4) from the same
     *  computation, rather than a second fetch/derivation of its own. */
    onbanddiff?: (diff: BandDiffView | null) => void;
    /** Whether the card draws nothing (quietMatch on a matching build): the planner lays out the rail by it. */
    onquiet?: (quiet: boolean) => void;
    /** Fires once a load actually applies, with the band's own label -- fix round 2
     *  (ux-designer review finding 2): a band load is never a character/addon import and
     *  must never claim to be one in Planner.svelte's own reconstruction note. */
    onloaded?: (bandLabel: string) => void;
    /** Fires with the band entry the section compares against (or null) and its label: the
     *  head's live score sims this band's gear (decision P1). */
    onband?: (band: BisBand | null, bandLabel: string) => void;
    /** The section's title in place of `{band} build`: a shared build's `Guide build` (ruling SHOULD). */
    heading?: string;
    /** The standalone planner: a build that matches its level build draws no card (h-audit 2026-11-09:
     *  a plate holding one Matches chip and a button); the guide row's Loaded tag says it. */
    quietMatch?: boolean;
  } = $props();

  let file = $state<BisFile | null | undefined>(undefined);
  const faction = $derived<Faction>(store.raceRow?.faction === 'horde' ? 'horde' : 'alliance');
  /** The band of the points spent (what a bare build is compared with). */
  const reachedBand = $derived(bandForLevel(levelReached(store.order)));
  /** The band of the build's level: a level-60 character with a build short of 51 points is still
   *  a level 60, compared with the level-60 build rather than the band its points happen to reach.
   *  `store.level` is the points' level for a bare build, so nothing else moves. */
  const levelBand = $derived(bandForLevel(store.level));
  const wantedBand = $derived(bandForPanel(store.order.length, characterLevel, levelBand));
  /** The visitor's level-60 default only while the spec's file has that band: a spec whose list
   *  stops short of 60 compares at the band the build has reached, as before. */
  const band = $derived(
    wantedBand !== reachedBand &&
      file !== undefined &&
      file !== null &&
      bandEntryFor(file, wantedBand, faction) === undefined
      ? reachedBand
      : wantedBand,
  );
  const specKey = $derived(
    specOverride ?? (store.talentIndex ? specKeyFor(store.classSlug, store.split) : null),
  );
  const bandLabel = $derived(bandRangeLabel(band));
  // The spec's ladder order (/data/leveling-plans.json), so "Load" spends the band's points the
  // way the sim and the guide's Leveling strip do; {} until read, and on a failed read.
  let ladderOrderFor = $state<LadderOrderFor>(() => undefined);
  $effect(() => {
    void import('../../lib/planner/ladder-load')
      .then((m) => m.loadLadderOrderFor())
      .then((fn) => (ladderOrderFor = fn))
      .catch(() => {});
  });

  // undefined: still loading (or not started). null: fetched, and there is no file for this
  // spec yet -- the honest empty state, never a guess (BisSlotPopover.svelte's own pattern).
  let loadedFor = $state('');

  $effect(() => {
    const build = store.treeVersion;
    const spec = specKey;
    if (spec === null || spec === '' || build === '') {
      file = null;
      loadedFor = '';
      return;
    }
    const key = `${build}::${spec}`;
    if (loadedFor === key) return;
    loadedFor = key;
    file = undefined;
    void fetchBisFile(build, spec).then((result) => {
      if (loadedFor === key) file = result;
    });
  });

  const bandEntry = $derived(
    file !== null && file !== undefined ? bandEntryFor(file, band, faction) : undefined,
  );

  const diff = $derived.by(() => {
    if (store.talentIndex === null || bandEntry === undefined) return null;
    const result = diffAgainstBand(store.talentIndex, store.ranks, bandEntry.talents);
    return { ...result, bandLabel } as BandDiffView;
  });

  /** Every band point taken reads as a match; points past the band's total are a quiet note. */
  const status = $derived(diff === null ? null : bandDiffStatus(diff));
  /** A character with more points than the band's build covers never reads `Matches` (MUST 3). */
  const coverage = $derived(
    status === null || bandEntry === undefined || store.talentIndex === null
      ? null
      : bandCoverage(
          status,
          bandPointTotal(store.talentIndex, bandEntry.talents),
          characterLevel === null ? null : pointsAvailableAt(characterLevel),
        ),
  );
  const beyond = $derived(
    status === null || status.kind === 'toTake' || coverage !== null ? 0 : status.beyond,
  );
  /** The difference named, not counted (ruling Q1). */
  const namedLine = $derived.by(() => {
    if (status?.kind !== 'differs' || store.talentIndex === null || bandEntry === undefined) return '';
    const names = bandDiffNames(store.talentIndex, store.ranks, bandEntry.talents);
    return bandCompareCopy.differsNamed(status.count, names.mine, names.theirs, bandLabel);
  });

  $effect(() => {
    onbanddiff?.(diff);
  });
  $effect(() => {
    onquiet?.(quiet);
  });
  // Only once the band file has answered: the estimate waits for the gear it runs on.
  $effect(() => {
    if (file !== undefined) onband?.(bandEntry ?? null, bandLabel);
  });

  const gearHref = $derived(
    specKey !== null && specKey !== '' ? bisPageHref(specKey, faction, band) : '/bis',
  );

  let confirming = $state(false);
  /** What the last load did, shown until the build changes from what it loaded. */
  let loaded = $state({ key: '', line: '' });
  const loadedLine = $derived(loaded.key === store.order.join(',') ? loaded.line : '');
  const quiet = $derived(
    quietMatch && status?.kind === 'matches' && coverage === null && !confirming && loadedLine === '',
  );

  /** The band's talents, applied in tree/tier/column order -- the same order the
   *  reconstruction itself walks -- so "see what differs" always jumps to a real,
   *  deterministic first cell rather than whichever `Set` iteration happens to yield. */
  function firstDiffingTalentId(): number | null {
    if (store.talentIndex === null || diff === null) return null;
    for (const tree of store.talentIndex.trees) {
      const byTier = [...tree.talents].sort((a, b) => a.tier - b.tier || a.column - b.column);
      for (const talent of byTier) {
        if (diff.diffTalentIds.has(talent.id)) return talent.id;
      }
    }
    return null;
  }

  function scrollToFirstDiff(): void {
    const id = firstDiffingTalentId();
    if (id === null) return;
    const cell = document.querySelector<HTMLElement>(`[data-testid="talent-${id}"]`);
    cell?.scrollIntoView({ behavior: 'smooth', block: 'center' });
    cell?.focus({ preventScroll: true });
  }

  function applyLoad(): void {
    if (store.talentIndex === null || bandEntry === undefined) return;
    const ordered =
      specKey === null ? undefined : ladderOrderFor(store.talentIndex, bandEntry.talents, specKey);
    const result = loadFromBandAtLevel(store.talentIndex, bandEntry.talents, characterLevel, ordered);
    store.applyOrder(result.order, {});
    const points = result.order.length;
    loaded = {
      key: store.order.join(','),
      line: bandCompareCopy.loadedLine(points + 9, points, result.fullPoints),
    };
    store.nameBandBuild(bandTitleForDraft(bandLabel, store.classSlug, store.split));
    onloaded?.(bandLabel);
  }

  function requestLoad(): void {
    if (store.order.length === 0) {
      applyLoad();
      return;
    }
    confirming = true;
  }

  function confirmLoad(): void {
    confirming = false;
    applyLoad();
  }
</script>

{#if !quiet}<Section title={heading ?? bandCompareCopy.heading(bandLabel)} {panel} testid="band-compare">
    {#if file === undefined}
      <div class="pl-sec" data-testid="band-compare-skeleton" aria-busy="true">
        <span class="skeleton-block block max-w-full" style="width:260px;height:20px"></span>
        <span class="skeleton-block block max-w-full" style="width:200px;height:44px"></span>
      </div>
    {:else if file === null || bandEntry === undefined}
      <div class="pl-pad">
        <EmptyState
          message={bandCompareCopy.noLevelBuild(
            store.level,
            specKey !== null ? specDisplayName(specKey) : '',
            specKey !== null ? siblingSpecNames(specKey) : [],
          )}
          testid="band-compare-empty"
        />
      </div>
    {:else if diff !== null}
      <div class="pl-sec">
        <p class="pl-line" data-testid="band-compare-differs">
          {#if coverage !== null}
            {bandCompareCopy.coversFewer(coverage.covered, coverage.total)}
          {:else if status?.kind === 'matches'}
            <Tag tone="success"
              >{bandCompareCopy.matches(specKey !== null ? specDisplayName(specKey) : '')}</Tag
            >
          {:else if status?.kind === 'differs'}
            {namedLine}
          {:else if status?.kind === 'toTake'}
            <b class="text-gold tabular-nums">{status.count}</b>
            {bandCompareCopy.toTakeSuffix}
          {/if}
        </p>
        {#if beyond > 0}
          <p class="text-small text-muted" data-testid="band-compare-beyond">
            {bandCompareCopy.beyondGuide(beyond)}
          </p>
        {/if}
        {#if confirming}
          <p class="text-small text-muted" data-testid="band-compare-load-confirm">
            {bandCompareCopy.loadConfirmMessage}
          </p>
          <div class="pl-acts">
            <Button onclick={confirmLoad} testid="band-compare-load-confirm-proceed"
              >{bandCompareCopy.loadConfirmProceed}</Button
            >
            <Button onclick={() => (confirming = false)} testid="band-compare-load-confirm-cancel"
              >{bandCompareCopy.loadConfirmCancel}</Button
            >
          </div>
        {:else}
          <div class="pl-acts">
            {#if !store.readOnly}
              <Button onclick={requestLoad} testid="band-compare-load"
                >{bandCompareCopy.loadButton(bandLabel)}</Button
              >
            {/if}
            {#if status?.kind === 'differs'}
              <Button onclick={scrollToFirstDiff} testid="band-compare-see-differs"
                >{bandCompareCopy.seeWhatDiffers}</Button
              >
            {/if}
          </div>
        {/if}
        {#if loadedLine !== ''}
          <p class="text-small text-muted" data-testid="band-compare-loaded">{loadedLine}</p>
        {/if}
        <p class="pl-gear">
          <a class={LINK} href={gearHref} data-testid="band-compare-gear-link"
            >{bandCompareCopy.bestInSlotList} ›</a
          >
        </p>
      </div>
    {/if}
  </Section>{/if}
