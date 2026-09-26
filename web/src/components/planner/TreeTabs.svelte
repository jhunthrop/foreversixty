<!-- web/src/components/planner/TreeTabs.svelte -->
<!-- The phone tab strip (one tree at a time, Gear last) and the tree row it switches
     between -- split out of Planner.svelte (design loop, planner round) so that file stays
     under the project's file-size guideline as the responsive rail layout grew it. From md
     up the tab strip is `md:hidden` and every tree panel is `md:flex` regardless of
     `activeTree`, so this renders the same desktop columns Planner.svelte always has; only
     the phone, one-panel-at-a-time behaviour depends on the state this owns. -->
<script lang="ts">
  import { SECONDARY_BUTTON } from '../../lib/planner/styles';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import type { TalentIndex } from '../../lib/planner/rules';
  import TreeGrid from './TreeGrid.svelte';

  let {
    store,
    talentIndex,
    activeTree = $bindable(),
    hasGear,
    gearTabIndex,
    treeColumnsClass,
  }: {
    store: PlannerStore;
    talentIndex: TalentIndex;
    /** Which tab is open: a tree by index, or `gearTabIndex` for Gear. Ignored from md up,
     *  where every tree panel shows regardless (Planner.svelte's own `md:flex`). */
    activeTree: number;
    hasGear: boolean;
    gearTabIndex: number;
    /** One literal Tailwind class per tree count (styles.ts's `treeRowColumnsClass`), so the
     *  row never reserves a column no tree fills (build review round 1, finding 2). */
    treeColumnsClass: string;
  } = $props();

  // The roving tabindex keeps the unselected tabs out of the tab order, so arrow keys are the
  // only way to reach them: without this a keyboard could never open the second tree, or the
  // gear panel behind the last tab. Moving selects, which is the automatic-activation half of
  // the ARIA tabs pattern -- switching panel costs nothing, so there is no reason to make it
  // a second keypress. The tabs come off the event rather than a `bind:this`: the listener is
  // on the tablist itself, so currentTarget is already the element, and there is no reference
  // to go stale when the branch unmounts.
  function onTabKeys(event: KeyboardEvent & { currentTarget: HTMLDivElement }): void {
    const step = event.key === 'ArrowRight' ? 1 : event.key === 'ArrowLeft' ? -1 : 0;
    if (step === 0) return;
    event.preventDefault();
    const tabs = event.currentTarget.querySelectorAll<HTMLButtonElement>('[role="tab"]');
    activeTree = (activeTree + step + tabs.length) % tabs.length;
    tabs[activeTree].focus();
  }
</script>

<!-- One panel at a time on a phone: three trees side by side do not fit 360px, and
   stacking them -- with the gear panel's seventeen slots under them -- puts the last
   one several screens down. Desktop keeps the columns and hides this. The roving
   tabindex lives on the tabs, as it does on TreeGrid's cells.
   The -1 on the container changes nothing about the keyboard order -- a bare div was
   never a tab stop -- and is there to satisfy the compiler's a11y rule that an element
   carrying an interactive role and a key handler declare a tabindex; -1 declares one
   without adding a stop, and it matches what TreeGrid's `role="grid"` does. -->
<div
  role="tablist"
  tabindex={-1}
  aria-label="Planner sections"
  class="border-line-soft mx-[18px] flex gap-2 border-b pb-2 md:hidden"
  onkeydown={onTabKeys}
>
  {#each talentIndex.trees as tree, i (tree.id)}
    <button
      type="button"
      role="tab"
      id={`tree-tab-${tree.id}`}
      aria-selected={activeTree === i}
      aria-controls={`tree-panel-${tree.id}`}
      tabindex={activeTree === i ? 0 : -1}
      class="{SECONDARY_BUTTON} flex-1 justify-center {activeTree === i
        ? 'border-gold text-gold'
        : 'border-line text-nav'}"
      onclick={() => (activeTree = i)}
    >
      <span>{tree.name}</span>
      <span class="tabular text-muted ml-2 font-mono">{store.split[i] ?? 0}</span>
    </button>
  {/each}
  <!-- Last, so `gearTabIndex` is the tree count and onTabKeys picks it up from the
     tablist's own DOM order without knowing gear exists. No count beside the name:
     the trees show the points spent in them because that number is otherwise only on
     the panel behind the tab, and the gear panel's own totals are not one number. -->
  {#if hasGear}
    <button
      type="button"
      role="tab"
      id="gear-tab"
      aria-selected={activeTree === gearTabIndex}
      aria-controls="gear-tabpanel"
      tabindex={activeTree === gearTabIndex ? 0 : -1}
      class="{SECONDARY_BUTTON} flex-1 justify-center {activeTree === gearTabIndex
        ? 'border-gold text-gold'
        : 'border-line text-nav'}"
      onclick={() => (activeTree = gearTabIndex)}
    >
      Gear
    </button>
  {/if}
</div>

<div
  class="order-1 grid grid-cols-1 gap-4 px-[18px] md:order-1 md:col-span-2 md:px-0 lg:order-1 lg:col-span-8 {treeColumnsClass}"
  data-testid="tree-columns"
>
  {#each talentIndex.trees as tree, i (tree.id)}
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
      <!-- The game's tree header: name on the left, points in the tree on the
         right, a rule under both. Warm border and gold number are the
         design system's; the proportions are the client's. -->
      <header class="border-line-soft flex items-baseline justify-between border-b pb-2">
        <h2 class="section-title text-[15px]">{tree.name}</h2>
        <span class="tabular text-gold font-mono text-[15px]" data-testid={`tree-points-${tree.id}`}>
          {store.split[i] ?? 0}
        </span>
      </header>
      <TreeGrid {store} {tree} />
    </div>
  {/each}
</div>
