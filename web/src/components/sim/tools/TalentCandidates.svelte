<!-- web/src/components/sim/tools/TalentCandidates.svelte -->
<!-- Design 3.1.6: the character's own build, every planner build the signed-in player has
     saved for this class, the in-game loadouts the addon export carried, and "add a build"
     which mounts the planner inline and reads its code back. The planner is a large
     component and /sim/weights must never download it, so it ships as its own lazy chunk,
     opened only when the player actually asks for it -- the same idiom SimView.svelte and
     ReportView.svelte already use for their own non-default views/modes. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import activeBuild from '../../../data/active-build.json';
  import { fetchBisFile } from '../../../lib/bis/hover';
  import { decodeFS1, orderFromRanks } from '../../../lib/planner/fs1';
  import { loadTalents } from '../../../lib/planner/load';
  import { indexTalents } from '../../../lib/planner/rules';
  import type { BuildRecord, TalentFile } from '../../../lib/planner/types';
  import { createLazyComponent, type LazyLoadState } from '../../../lib/report/lazy-component.svelte';
  import { fetchMyBuilds, SimApiError } from '../../../lib/sim/api';
  import type { TalentLoadout } from '../../../lib/sim/bulk-types';
  import { plannerGearFor, talentsString, type SimCharacter } from '../../../lib/sim/character';
  import { bulkCopy, simCopy } from '../../../lib/sim/copy';
  import { dedupeByName } from '../../../lib/sim/dedupe';
  import { poolQualityCopy } from '../../../lib/sim/pool-quality-copy';
  import {
    customLoadouts,
    guideLoadout,
    hasUnaddedBuild,
    loadoutDiff,
    loadoutTrees,
    savedBuildsMessage,
  } from '../../../lib/sim/talent-loadouts';
  import { talentSplit } from '../../../lib/sim/result-band';
  import { ROW_ITEM } from '../../../lib/ui/primitives';
  import Button from '../../ui/Button.svelte';
  import TalentTrees from '../../ui/TalentTrees.svelte';
  import ToolSection from './ToolSection.svelte';

  let {
    character,
    picked,
    ontoggle,
    onpendingcustom,
    plate = false,
  }: {
    character: SimCharacter;
    picked: readonly TalentLoadout[];
    ontoggle: (loadout: TalentLoadout, on: boolean) => void;
    /**
     * newcomer round 4 (review.md:83-110): fires whenever a valid pasted build starts or
     * stops sitting unadded in the editor, so BulkRunBar (a sibling, not a child) can block
     * RUN on it -- the one remaining way the two-buttons-same-label trap could still bite
     * once the buttons themselves are told apart.
     */
    onpendingcustom?: (pending: boolean) => void;
    /** The builds sit on the page's one Plate (/sim/talents). */
    plate?: boolean;
  } = $props();

  let talents = $state<TalentFile | null>(null);
  let saved = $state<BuildRecord[] | null>(null);
  let savedFailed = $state(false);
  /** dps D31: a 401/403 (signed out) reads as guidance, not an error -- `savedBuildsMessage`. */
  let savedSignedOut = $state(false);
  /** The spec's guide build at the character's band, once its published file loads. */
  let guide = $state<TalentLoadout | null>(null);
  let plannerOpen = $state(false);
  let customCode = $state('');

  const plannerLazy = createLazyComponent(() => import('../../planner/Planner.svelte'));

  onMount(() => {
    fetchBisFile(activeBuild.build, character.spec)
      .then(
        (file) => (guide = file === null ? null : guideLoadout(file, character.level, character.race_slug)),
      )
      .catch((error: unknown) => console.error('sim talents: the guide build did not load', error));
    void loadTalents(character.tree_version, character.class_slug)
      .then((file) => (talents = file))
      .catch(() => (talents = null));
    // Contract 10.6's own ruling: every failure -- a 404 from a deployment older than the
    // migration, or anything else -- reads as "no saved builds", one line, no error banner,
    // because the rest of the page is correct without this list. dps D31: a 401/403 is not
    // "anything else" to the player -- it is the expected, permanent state for a visitor
    // with no account, so it gets its own, non-alarming wording (`savedBuildsMessage`).
    void fetchMyBuilds()
      .then((page) => {
        saved = page.rows.filter((row) => row.tree_version === character.tree_version);
      })
      .catch((error: unknown) => {
        saved = [];
        savedFailed = true;
        savedSignedOut = error instanceof SimApiError && (error.status === 401 || error.status === 403);
      });
  });

  const index = $derived(talents === null ? null : indexTalents(talents));

  /** The character's own build, always the first row, available once talents load. */
  const own = $derived<TalentLoadout | null>(
    index === null
      ? null
      : { name: bulkCopy.talentsOwn, talents: talentsString(index, character.point_order) },
  );

  /** A saved build's point order as the engine's talents string, through the one converter. */
  function loadoutFor(record: BuildRecord): TalentLoadout | null {
    if (index === null) return null;
    return {
      name: record.title === undefined || record.title === '' ? record.id : record.title,
      talents: talentsString(index, record.point_order),
    };
  }

  /**
   * Deduplicated by name: a saved planner build titled the same as an in-game loadout the
   * addon export carries would otherwise render two checkboxes under the identical
   * `sim-loadout-<name>` test id -- not just confusing, a Playwright strict-mode failure the
   * moment a test looks for that id. Saved builds are listed first on screen, so they claim
   * a name first; an in-game loadout sharing one is the entry dropped, below.
   */
  const savedLoadouts = $derived(
    dedupeByName(
      new Set<string>(),
      (saved ?? []).map(loadoutFor).filter((entry): entry is TalentLoadout => entry !== null),
    ),
  );

  /**
   * The addon export's in-game loadouts (part A's FS1 v2 decoder). `SimCharacter.loadouts`
   * carries the decoder's own shape -- one rank array per tree -- not the engine's talents
   * string, so each one goes through the same rank-to-order-to-string pipeline
   * `characterFromFs1` already uses for the character itself. Named entries already claimed
   * by `savedLoadouts` (above) are dropped rather than rendered a second time.
   */
  const exported = $derived<TalentLoadout[]>(
    dedupeByName(
      new Set(savedLoadouts.map((entry) => entry.name)),
      index === null
        ? []
        : character.loadouts.map((loadout) => {
            const { order } = orderFromRanks(index, loadout.treeRanks);
            return { name: loadout.name, talents: talentsString(index, order) };
          }),
    ),
  );

  /** dps D31/E2: a picked loadout none of the three named lists above claims -- ADD A
   *  BUILD's own paste or hand-built tree, otherwise invisible once accepted. */
  const custom = $derived(
    customLoadouts(picked, own, [savedLoadouts, exported, guide === null ? [] : [guide]]),
  );

  /** newcomer round 4: true whenever `addCustom` could accept `customCode` right now but
   *  has not -- the trap's own precondition, reported to the RUN button's own gate. */
  const pending = $derived(hasUnaddedBuild(customCode, index));
  $effect(() => {
    onpendingcustom?.(pending);
    return () => onpendingcustom?.(false);
  });

  /** What a build's diff reads against: the character's own build, or for that one the guide build. */
  function againstOf(loadout: TalentLoadout): TalentLoadout | null {
    const other = own !== null && loadout.name !== own.name ? own : guide;
    return other !== null && other.name !== loadout.name ? other : null;
  }

  function isPicked(loadout: TalentLoadout): boolean {
    return picked.some((entry) => entry.name === loadout.name);
  }

  /** The next "Build N" name not already in `picked` -- `store.addLoadout` itself dedupes
   *  by name and silently no-ops on a repeat, so a name this component hands it is always
   *  one that will actually be added. */
  function nextCustomName(): string {
    let n = picked.length + 1;
    while (picked.some((entry) => entry.name === `Build ${n}`)) n += 1;
    return `Build ${n}`;
  }

  /** The inline planner's code, turned into a loadout the moment the player accepts it. */
  function addCustom(): void {
    if (index === null || customCode === '') return;
    const decoded = decodeFS1(customCode);
    if (!decoded.ok) return;
    const { order } = orderFromRanks(index, decoded.build.treeRanks);
    ontoggle({ name: nextCustomName(), talents: talentsString(index, order) }, true);
    plannerOpen = false;
    customCode = '';
  }
</script>

{#snippet lazyFallback(lazy: LazyLoadState)}
  {#if lazy.error !== ''}
    <p class="text-muted text-[13px]" role="alert">
      {lazy.error}
      <button
        type="button"
        class="text-strong ml-1 inline-flex min-h-11 items-center underline"
        onclick={() => lazy.load()}>{simCopy.tryAgain}</button
      >
    </p>
  {/if}
{/snippet}

{#snippet loadoutCheck(loadout: TalentLoadout, testid: string)}
  <label class="text-text flex min-h-11 items-center gap-3 text-[14px]">
    <input
      type="checkbox"
      class="accent-gold h-5 w-5"
      data-testid={testid}
      checked={isPicked(loadout)}
      onchange={(event) => ontoggle(loadout, event.currentTarget.checked)}
    />
    <span class="min-w-0 flex-1">{loadout.name}</span>
    {#if plate && !loadout.name.includes(talentSplit(loadout.talents))}
      <span class="text-muted text-[13px] tabular-nums">{talentSplit(loadout.talents)}</span>
    {/if}
  </label>
{/snippet}

{#snippet loadoutRow(loadout: TalentLoadout, testid: string)}
  {#if plate && index !== null}
    <!-- /sim/talents: the build's three trees beside its tick, so the page shows what it compares. -->
    <div class="{ROW_ITEM} grid gap-x-6 gap-y-1 py-3 md:grid-cols-[260px_minmax(0,1fr)] md:items-start">
      {@render loadoutCheck(loadout, testid)}
      <!-- The trees size to their track (TalentTrees reads its box's width): 600 px with the diff
           beside them from 1024, the full width with the diff under them below; cells up to 44 px. -->
      <div class="grid min-w-0 items-start gap-x-8 gap-y-3 lg:grid-cols-[600px_minmax(0,1fr)]">
        <div class="sim-loadout-trees min-w-0">
          <TalentTrees trees={loadoutTrees(index, loadout.talents)} build={character.tree_version} />
        </div>
        <!-- Each build against the other side: a build against the character's own, the character's
             own against the guide build (h-audit 2026-11-09: its column stood empty). -->
        {#if againstOf(loadout) !== null}
          {@const changes = loadoutDiff(index, loadout.talents, againstOf(loadout)?.talents ?? '')}
          <!-- Against the character's own build: what this one takes more of (Gains), then less (Drops). -->
          <div
            class="flex flex-wrap gap-x-8 gap-y-2 text-[13px] leading-5 tabular-nums"
            data-testid={`${testid}-diff`}
          >
            {#each [{ title: bulkCopy.talentsGains, rows: changes.filter((change) => change.delta > 0), tone: 'text-kill' }, { title: bulkCopy.talentsDrops, rows: changes.filter((change) => change.delta < 0), tone: 'text-wipe' }] as group (group.title)}
              {#if group.rows.length > 0}
                <div class="flex flex-col">
                  <h4 class="text-label text-muted m-0 pb-1 font-bold">{group.title}</h4>
                  <ul class="m-0 flex list-none flex-col p-0">
                    {#each group.rows as change (change.name)}
                      <li class={group.tone}>
                        {change.delta > 0 ? '+' : '−'}{Math.abs(change.delta)}
                        <span class="text-text">{change.name}</span>
                      </li>
                    {/each}
                  </ul>
                </div>
              {/if}
            {/each}
          </div>
        {/if}
      </div>
    </div>
  {:else}
    {@render loadoutCheck(loadout, testid)}
  {/if}
{/snippet}

<ToolSection title={bulkCopy.talentBuildsTitle} testid="sim-talent-candidates">
  <div class="flex flex-col">
    {#if own !== null}
      {@render loadoutRow(own, 'sim-loadout-current')}
    {/if}
    {#if guide !== null}
      {@render loadoutRow(guide, 'sim-loadout-guide')}
    {/if}

    <h3 class="text-strong pt-3 text-[15px] leading-5 font-bold">{bulkCopy.talentsSaved}</h3>
    {#each savedLoadouts as loadout (loadout.name)}
      {@render loadoutRow(loadout, `sim-loadout-${loadout.name}`)}
    {/each}

    {#if savedFailed}
      <p class="text-muted pt-1 text-[13px]" data-testid="sim-loadouts-unavailable">
        {savedBuildsMessage(savedSignedOut)}
      </p>
    {:else if savedLoadouts.length === 0 && saved !== null}
      <p class="text-muted pt-1 text-[13px]">{bulkCopy.talentsNoSaved}</p>
    {/if}

    {#if exported.length > 0}
      <h3 class="text-strong pt-3 text-[15px] leading-5 font-bold">{bulkCopy.talentsLoadouts}</h3>
      {#each exported as loadout (loadout.name)}
        {@render loadoutRow(loadout, `sim-loadout-${loadout.name}`)}
      {/each}
    {/if}

    {#each custom as loadout (loadout.name)}
      {@render loadoutRow(loadout, `sim-loadout-${loadout.name}`)}
    {/each}

    <div class="pt-3">
      <Button
        tone="ghost"
        size="sm"
        testid="sim-loadout-add"
        onclick={() => {
          plannerOpen = !plannerOpen;
          if (plannerOpen) plannerLazy.load();
        }}>{bulkCopy.talentsAddCustom}</Button
      >
    </div>

    {#if plannerOpen}
      <div class="pt-3" data-testid="sim-inline-planner">
        {#if plannerLazy.current}
          <!-- dps D39: without a seeded `gear`, this inline card's live DPS estimate ran a
               gearless sim while the comparison table below simmed the loaded character's real
               gear, so the same build read two different DPS numbers on one screen. Passing
               the loaded character's gear here (through plannerGearFor) is what keeps the
               inline card and the ranking table honest with each other. -->
          <plannerLazy.current
            treeVersion={character.tree_version}
            classSlug={character.class_slug}
            raceSlug={character.race_slug}
            gear={plannerGearFor(character)}
            oncode={(code: string) => (customCode = code)}
            standalone={false}
          />
          <div class="pt-2">
            <Button size="sm" testid="sim-loadout-accept" disabled={customCode === ''} onclick={addCustom}
              >{poolQualityCopy.talentsAddCustomConfirm}</Button
            >
          </div>
        {:else}
          {@render lazyFallback(plannerLazy)}
        {/if}
      </div>
    {/if}
  </div>
</ToolSection>

<style>
  /* A compare page reads the trees, so they take more of the plate than the guide's compact strip. */
  .sim-loadout-trees :global(.ui-trees-box) {
    --cell-max: 44px;
  }
  /* Below 768 one tree per row at a readable cell, never three trees squeezed to 100 px. */
  @media (max-width: 767px) {
    .sim-loadout-trees :global(.ui-trees-box .ui-trees) {
      /* A tree fills the plate's width: four cells and three gaps, up to 56 px a cell. */
      --cell: min(56px, calc((100cqw - 3 * var(--cell-gap)) / 4));
      grid-template-columns: calc(4 * var(--cell) + 3 * var(--cell-gap));
    }
  }
</style>
