<!-- web/src/components/sim/tools/TalentCandidates.svelte -->
<!-- Design 3.1.6: the character's own build, every planner build the signed-in player has
     saved for this class, the in-game loadouts the addon export carried, and "add a build"
     which mounts the planner inline and reads its code back. The planner is a large
     component and /sim/weights must never download it, so it ships as its own lazy chunk,
     opened only when the player actually asks for it -- the same idiom SimView.svelte and
     ReportView.svelte already use for their own non-default views/modes. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { decodeFS1, orderFromRanks } from '../../../lib/planner/fs1';
  import { loadTalents } from '../../../lib/planner/load';
  import { indexTalents } from '../../../lib/planner/rules';
  import { SECONDARY_BUTTON } from '../../../lib/planner/styles';
  import type { BuildRecord, TalentFile } from '../../../lib/planner/types';
  import { createLazyComponent, type LazyLoadState } from '../../../lib/report/lazy-component.svelte';
  import { fetchMyBuilds, SimApiError } from '../../../lib/sim/api';
  import type { TalentLoadout } from '../../../lib/sim/bulk-types';
  import { plannerGearFor, talentsString, type SimCharacter } from '../../../lib/sim/character';
  import { bulkCopy, simCopy } from '../../../lib/sim/copy';
  import { dedupeByName } from '../../../lib/sim/dedupe';
  import { customLoadouts, savedBuildsMessage } from '../../../lib/sim/talent-loadouts';

  let {
    character,
    picked,
    ontoggle,
  }: {
    character: SimCharacter;
    picked: readonly TalentLoadout[];
    ontoggle: (loadout: TalentLoadout, on: boolean) => void;
  } = $props();

  let talents = $state<TalentFile | null>(null);
  let saved = $state<BuildRecord[] | null>(null);
  let savedFailed = $state(false);
  /** dps D31: a 401/403 (signed out) reads as guidance, not an error -- `savedBuildsMessage`. */
  let savedSignedOut = $state(false);
  let plannerOpen = $state(false);
  let customCode = $state('');

  const plannerLazy = createLazyComponent(() => import('../../planner/Planner.svelte'));

  onMount(() => {
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
  const custom = $derived(customLoadouts(picked, own, [savedLoadouts, exported]));

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

<section
  class="border-line rounded-panel mx-[18px] flex flex-col gap-2 border p-3 md:mx-0"
  data-testid="sim-talent-candidates"
>
  <h3 class="section-title text-[14px]">{bulkCopy.talentsSaved}</h3>

  {#if own !== null}
    <label class="text-text flex min-h-11 items-center gap-2 text-[13px]">
      <input
        type="checkbox"
        class="h-5 w-5"
        data-testid="sim-loadout-current"
        checked={isPicked(own)}
        onchange={(event) => ontoggle(own, event.currentTarget.checked)}
      />
      {bulkCopy.talentsOwn}
    </label>
  {/if}

  {#each savedLoadouts as loadout (loadout.name)}
    <label class="text-text flex min-h-11 items-center gap-2 text-[13px]">
      <input
        type="checkbox"
        class="h-5 w-5"
        data-testid={`sim-loadout-${loadout.name}`}
        checked={isPicked(loadout)}
        onchange={(event) => ontoggle(loadout, event.currentTarget.checked)}
      />
      {loadout.name}
    </label>
  {/each}

  {#if savedFailed}
    <p class="text-muted text-[12px]" data-testid="sim-loadouts-unavailable">
      {savedBuildsMessage(savedSignedOut)}
    </p>
  {:else if savedLoadouts.length === 0 && saved !== null}
    <p class="text-muted text-[12px]">{bulkCopy.talentsNoSaved}</p>
  {/if}

  {#if exported.length > 0}
    <h4 class="text-muted text-[12px]">{bulkCopy.talentsLoadouts}</h4>
    {#each exported as loadout (loadout.name)}
      <label class="text-text flex min-h-11 items-center gap-2 text-[13px]">
        <input
          type="checkbox"
          class="h-5 w-5"
          data-testid={`sim-loadout-${loadout.name}`}
          checked={isPicked(loadout)}
          onchange={(event) => ontoggle(loadout, event.currentTarget.checked)}
        />
        {loadout.name}
      </label>
    {/each}
  {/if}

  {#each custom as loadout (loadout.name)}
    <label class="text-text flex min-h-11 items-center gap-2 text-[13px]">
      <input
        type="checkbox"
        class="h-5 w-5"
        data-testid={`sim-loadout-${loadout.name}`}
        checked={isPicked(loadout)}
        onchange={(event) => ontoggle(loadout, event.currentTarget.checked)}
      />
      {loadout.name}
    </label>
  {/each}

  <button
    type="button"
    class="{SECONDARY_BUTTON} border-line-warm text-nav w-fit px-3"
    data-testid="sim-loadout-add"
    onclick={() => {
      plannerOpen = !plannerOpen;
      if (plannerOpen) plannerLazy.load();
    }}>{bulkCopy.talentsAddCustom}</button
  >

  {#if plannerOpen}
    <div class="border-line rounded-panel border p-2" data-testid="sim-inline-planner">
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
        <button
          type="button"
          class="{SECONDARY_BUTTON} border-line-warm text-nav mt-2 w-fit px-3 disabled:opacity-50"
          data-testid="sim-loadout-accept"
          disabled={customCode === ''}
          onclick={addCustom}>{bulkCopy.talentsAddCustom}</button
        >
      {:else}
        {@render lazyFallback(plannerLazy)}
      {/if}
    </div>
  {/if}
</section>
