<!-- web/src/components/planner/TreeTabs.svelte -->
<!-- The phone tab strip (one tree at a time) and the tree row it switches between -- split
     out of Planner.svelte (design loop, planner round) so that file stays under the
     project's file-size guideline as the responsive rail layout grew it. From md up the tab
     strip is `md:hidden` and every tree panel is `md:flex` regardless of `activeTree`, so
     this renders the same desktop columns Planner.svelte always has; only the phone,
     one-panel-at-a-time behaviour depends on the state this owns.

     Rebuild spec §4.G/§11 (owner ruling): the Gear tab is withdrawn with GearPanel itself --
     a planner build spans only its talent trees now, so this strip switches between trees
     alone. Rebuild spec §4.E.1: each tree's header (and, below md, its own tab) now carries
     the spec's own 28px icon from `data/curated/specs.json` -- a label, not a second
     navigation action; the phone tab strip's `role="tab"` buttons are still the only thing
     that switches the visible panel. -->
<script lang="ts">
  import type { BandDiffView } from '../../lib/planner/band-compare';
  import { dataUrl } from '../../lib/planner/load';
  import { SECONDARY_BUTTON } from '../../lib/planner/styles';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import type { TalentIndex } from '../../lib/planner/rules';
  import { SPECS } from '../../lib/sim/specs';
  import TreeGrid from './TreeGrid.svelte';

  let {
    store,
    talentIndex,
    activeTree = $bindable(),
    treeColumnsClass,
    bandDiff = null,
  }: {
    store: PlannerStore;
    talentIndex: TalentIndex;
    /** Which tree tab is open. Ignored from md up, where every tree panel shows regardless
     *  (Planner.svelte's own `md:flex`). */
    activeTree: number;
    /** One literal Tailwind class per tree count (styles.ts's `treeRowColumnsClass`), so the
     *  row never reserves a column no tree fills (build review round 1, finding 2). */
    treeColumnsClass: string;
    /** `BandCompare`'s own diff, once a band has loaded (spec §4.D/§4.E.4); `null` before
     *  then. Shared across every tree -- each `TreeGrid` only ever reads the ids its own
     *  tree's talents carry. */
    bandDiff?: BandDiffView | null;
  } = $props();

  /** The spec's own talent-tab icon for one tree, by its position in `talentIndex.trees`
   *  (that position IS `data/curated/specs.json`'s own `tree_index`, the same inference
   *  `specKeyFor`/`specOf` already key off of). `undefined` renders no icon rather than a
   *  guess (tenet 8). */
  function specIconFor(treeIndex: number): string | undefined {
    return SPECS.find((row) => row.class_slug === store.classSlug && row.tree_index === treeIndex)?.icon;
  }

  // The roving tabindex keeps the unselected tabs out of the tab order, so arrow keys are the
  // only way to reach them: without this a keyboard could never open the second tree. Moving
  // selects, which is the automatic-activation half of the ARIA tabs pattern -- switching
  // panel costs nothing, so there is no reason to make it a second keypress. The tabs come
  // off the event rather than a `bind:this`: the listener is on the tablist itself, so
  // currentTarget is already the element, and there is no reference to go stale when the
  // branch unmounts.
  function onTabKeys(event: KeyboardEvent & { currentTarget: HTMLDivElement }): void {
    const step = event.key === 'ArrowRight' ? 1 : event.key === 'ArrowLeft' ? -1 : 0;
    if (step === 0) return;
    event.preventDefault();
    const tabs = event.currentTarget.querySelectorAll<HTMLButtonElement>('[role="tab"]');
    activeTree = (activeTree + step + tabs.length) % tabs.length;
    tabs[activeTree].focus();
  }
</script>

<!-- One panel at a time on a phone: three trees side by side do not fit 360px. Desktop keeps
   the columns and hides this. The roving tabindex lives on the tabs, as it does on
   TreeGrid's cells.
   The -1 on the container changes nothing about the keyboard order -- a bare div was
   never a tab stop -- and is there to satisfy the compiler's a11y rule that an element
   carrying an interactive role and a key handler declare a tabindex; -1 declares one
   without adding a stop, and it matches what TreeGrid's `role="grid"` does. -->
<div
  role="tablist"
  tabindex={-1}
  aria-label="Talent trees"
  class="border-line-soft mx-[18px] flex gap-2 border-b pb-2 md:hidden"
  onkeydown={onTabKeys}
>
  {#each talentIndex.trees as tree, i (tree.id)}
    {@const icon = specIconFor(i)}
    <button
      type="button"
      role="tab"
      id={`tree-tab-${tree.id}`}
      aria-selected={activeTree === i}
      aria-controls={`tree-panel-${tree.id}`}
      tabindex={activeTree === i ? 0 : -1}
      class="{SECONDARY_BUTTON} flex-1 justify-center gap-2 {activeTree === i
        ? 'border-gold text-gold'
        : 'border-line text-nav'}"
      onclick={() => (activeTree = i)}
    >
      {#if icon !== undefined}
        <img
          src={dataUrl(store.treeVersion, `icons/${icon}.webp`)}
          alt=""
          width="28"
          height="28"
          loading="lazy"
          decoding="async"
          class="rounded-control border-line h-7 w-7 border"
        />
      {/if}
      <span>{tree.name}</span>
      <span class="tabular text-muted ml-2 font-mono">{store.split[i] ?? 0}</span>
    </button>
  {/each}
</div>

<div class="grid grid-cols-1 gap-4 px-[18px] md:px-0 {treeColumnsClass}" data-testid="tree-columns">
  {#each talentIndex.trees as tree, i (tree.id)}
    {@const icon = specIconFor(i)}
    <!-- The inactive trees are hidden with a class, not the `hidden` attribute: the
       attribute would hide them on desktop too, where `md:flex` cannot override it. -->
    <div
      id={`tree-panel-${tree.id}`}
      role="tabpanel"
      aria-labelledby={`tree-tab-${tree.id}`}
      data-testid={`tree-panel-${tree.id}`}
      class="border-line-warm bg-raised rounded-panel flex-col gap-3 border p-4 md:flex {activeTree === i
        ? 'flex'
        : 'hidden'}"
    >
      <!-- The game's tree header: spec icon, name, points in the tree, a rule under both --
         the same "icon - name - point count" order the BiS header's own SpecTabs use (spec
         §4.E.1). Not a button: a planner build spans every tree in one page, so this is a
         label, never a second navigation control beside the phone strip above. -->
      <header class="border-line-soft flex items-center justify-between gap-2 border-b pb-2">
        <span class="flex min-w-0 items-center gap-2">
          {#if icon !== undefined}
            <img
              src={dataUrl(store.treeVersion, `icons/${icon}.webp`)}
              alt=""
              width="28"
              height="28"
              loading="lazy"
              decoding="async"
              class="rounded-control border-line h-7 w-7 shrink-0 border"
            />
          {/if}
          <h2 class="section-title truncate text-[15px]">{tree.name}</h2>
        </span>
        <span class="tabular text-gold font-mono text-[15px]" data-testid={`tree-points-${tree.id}`}>
          {store.split[i] ?? 0}
        </span>
      </header>
      <TreeGrid {store} {tree} {bandDiff} />
    </div>
  {/each}
</div>
