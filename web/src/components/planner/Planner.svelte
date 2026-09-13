<!-- web/src/components/planner/Planner.svelte -->
<!-- The planner island. Mounted two ways: by src/pages/planner.astro with client:load, and
     by src/planner-island.ts on the API-rendered /b/:id page, where `record` is supplied and
     the build starts read-only. Desktop lays the trees out side by side; phone shows one
     tree at a time behind a tab switcher (Task 9). -->
<script lang="ts">
  import { untrack } from 'svelte';
  import { DEFAULT_CLASS_SLUG, ERA_DATA_NOTICE } from '../../lib/planner/config';
  import {
    DATA_LOAD_FAILED,
    DataLoadError,
    loadItems,
    loadReference,
    loadSets,
    loadTalents,
  } from '../../lib/planner/load';
  import { createPlannerStore } from '../../lib/planner/store.svelte';
  import { SECONDARY_BUTTON } from '../../lib/planner/styles';
  import type { BuildRecord } from '../../lib/planner/types';
  import GearPanel from './GearPanel.svelte';
  import OrderStrip from './OrderStrip.svelte';
  import SharePanel from './SharePanel.svelte';
  import SummaryBar from './SummaryBar.svelte';
  import TreeGrid from './TreeGrid.svelte';

  let {
    treeVersion,
    classSlug = DEFAULT_CLASS_SLUG,
    raceSlug,
    record = null,
  }: {
    treeVersion: string;
    classSlug?: string;
    raceSlug?: string;
    record?: BuildRecord | null;
  } = $props();

  // The page is static, so ?class= and ?race= can only be read in the browser. A record
  // (from /b/:id) wins over the query string: that build already names its class and race.
  function fromQuery(name: string): string | undefined {
    if (typeof window === 'undefined') return undefined;
    return new URLSearchParams(window.location.search).get(name) ?? undefined;
  }

  // The store is seeded once, from the props as they arrive. `untrack` says so: without it
  // the compiler warns that these reads only capture the initial value, which is the point --
  // after mount the store owns the class, the race and the order, not the props.
  const store = untrack(() =>
    createPlannerStore({
      treeVersion: record?.tree_version ?? treeVersion,
      classSlug: record ? classSlug : (fromQuery('class') ?? classSlug),
      raceSlug: record ? (raceSlug ?? '') : (fromQuery('race') ?? raceSlug ?? ''),
      order: record?.point_order,
      gear: record?.gear,
      title: record?.title,
      sourceId: record?.id ?? null,
      readOnly: record !== null,
    }),
  );

  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let attempt = $state(0);
  // Reset asks in the toolbar rather than through window.confirm: a browser dialog cannot be
  // styled, cannot say what it is about to clear, and reads badly on phone.
  let confirmingReset = $state(false);
  // Which tree the phone shows. Desktop ignores it and lays every tree out side by side.
  let activeTree = $state(0);

  // The roving tabindex keeps the unselected tabs out of the tab order, so arrow keys are the
  // only way to reach them: without this a keyboard could never open the second tree. Moving
  // selects, which is the automatic-activation half of the ARIA tabs pattern -- switching tree
  // costs nothing, so there is no reason to make it a second keypress. The tabs come off the
  // event rather than a `bind:this`: the listener is on the tablist itself, so currentTarget
  // is already the element, and there is no reference to go stale when the branch unmounts.
  function onTabKeys(event: KeyboardEvent & { currentTarget: HTMLDivElement }): void {
    const step = event.key === 'ArrowRight' ? 1 : event.key === 'ArrowLeft' ? -1 : 0;
    if (step === 0) return;
    event.preventDefault();
    const tabs = event.currentTarget.querySelectorAll<HTMLButtonElement>('[role="tab"]');
    activeTree = (activeTree + step + tabs.length) % tabs.length;
    tabs[activeTree].focus();
  }

  async function load(): Promise<void> {
    status = 'loading';
    try {
      const reference = await loadReference(store.treeVersion);
      store.setReference(reference);
      store.setTalents(await loadTalents(store.treeVersion, store.classSlug));
      store.setSets(await loadSets(store.treeVersion));
      // Gear is optional in the same way sets are: a build whose item table did not
      // normalize ships no items/<class>.json, and the planner is complete without a gear
      // panel. Only a 404 means that. A 5xx, an unreachable network or a malformed file is a
      // broken build rather than an absent one, so it is rethrown into the failure state --
      // the same line loadSets draws, for the same reason.
      try {
        store.setItems(await loadItems(store.treeVersion, store.classSlug));
      } catch (error) {
        if (!(error instanceof DataLoadError) || error.status !== 404) throw error;
        store.setItems({ build: store.treeVersion, class_slug: store.classSlug, items: [] });
      }
      status = 'ready';
    } catch {
      status = 'failed';
    }
  }

  // Re-runs whenever the class changes (selectClass drops the loaded trees) or Retry bumps
  // `attempt`. Reading both synchronously here is what registers them as dependencies; the
  // writes `load` performs happen after the tracking window, so this cannot loop.
  $effect(() => {
    void store.classSlug;
    void attempt;
    void load();
  });

  // Switching class empties the build on its own, and the class selector stays reachable while
  // the confirm is open. Without this the planner comes back asking whether to clear a build the
  // switch already cleared. The phone tab goes back to the first tree for the same reason: the
  // open tree belonged to the class that just left, and a class with fewer trees than the last
  // one would leave no tab selected and no panel shown at all. Its own effect rather than a line
  // in the one above: that effect documents a careful no-loop invariant, and this has nothing to
  // do with loading.
  $effect(() => {
    void store.classSlug;
    confirmingReset = false;
    activeTree = 0;
  });
</script>

<div class="flex flex-col gap-[22px] md:gap-8" data-testid="planner">
  <SummaryBar {store} />

  <p class="text-muted px-[18px] text-[13px] md:px-0">{ERA_DATA_NOTICE}</p>

  <!-- The three states below swap in place once the talent data arrives over the network, and
       they are wildly different heights: one line of status text against a planner fourteen
       hundred pixels tall on a phone. Whatever is under the planner -- the footer, mainly --
       moves by that difference, which Lighthouse measured as 0.185 of /planner.html's 0.186
       CLS against the 0.05 lighthouserc.json budget. This min-height reserves the room up
       front so the swap moves nothing below it.

       What it does not do is make the three states identical, and it does not abolish the
       swap. The reserve is one number; the ready height is not. These two come from the
       loaded layout of the default class, so a class whose trees run to more tiers grows past
       them and still moves the footer -- by the difference rather than by the whole planner.
       A class the build ships no items for is the same story in the other direction: no gear
       panel, so its ready state comes in some 680px under this and leaves that much dead
       space. Both are bounded by the reserve; neither is the whole-planner jump it replaces.
       Re-derive them by loading /planner, setting this element's min-height to 0, and reading
       its `getBoundingClientRect().height` below and above the md breakpoint. They measure
       1412.5 at 360px and 1038.5 from md up, now that Task 16's gear panel -- seventeen slot
       buttons two to a row, plus the totals and sets columns -- has joined the column. Each
       value here is set a hair under what was measured, because under costs a pixel of
       movement and over leaves dead space below the ready planner for good.

       Task 16 asked whether this still earns its keep, now that the ready planner is tall
       enough that the footer is below the fold in both states. It does, and the numbers are
       here so the question does not have to be re-opened blind: removing it entirely takes
       /planner.html's CLS from 0.002 to a median 0.176 and its performance score from 0.98 to
       0.91, because without it the *loading* state is short enough to leave the footer on
       screen, and the swap then hauls it 1400px down from inside the viewport. Nor is the
       answer a smaller number tuned to the audit: a 700px reserve scores an identical 0.0009
       CLS purely because it clears Lighthouse's emulated 640px fold, while still moving the
       footer 712px on the 800px-tall phone tests/e2e/planner-phone.spec.ts drives -- which is
       what those two footer assertions are for, and they fail it. The reserve has to cover the
       ready height, not the audit's viewport.

       The phone figure is deliberately the one measured at 360px, the narrowest width the
       site designs for and the width Lighthouse emulates (lighthouserc.json). It is the
       tallest: the toolbar row wraps one button further at 360 than it does from 390px up,
       which is 57px, so above 360 the ready planner comes in under this reserve and the
       reserve is what the region measures in all three states. That is dead space rather
       than movement, and it is the safe direction to err.

       It wraps the swapping branches only, not the planner as a whole, and that is what lets
       one number hold: the summary bar and the notice above are in all three states and
       reflow with the viewport width, so keeping them outside the reserve takes their
       wrapping out of the figure. Inside it every part is a fixed height -- the tab strip,
       the toolbar (including the always-visible title field and Share button), the order
       strip's reserved row, a tree grid sized by tier count rather than by width, and a gear
       panel whose slot grid is a fixed count of fixed-height rows and whose totals and sets
       columns start on their one-line empty state. The md value is the smaller one because
       desktop drops the tab strip, lays the trees out side by side and puts the slots four to
       a row, so the tallest tree sets the height, not their sum.

       Fork replaces Reset and drops the SharePanel section, but only on the read-only mount --
       the editable toolbar this measures is untouched. The read-only mount is the shorter one,
       1331 and 962, so it sits about 80px under the reserve and leaves that much space above
       the footer on the API's /b/:id. Reserving the taller figure in both is deliberate:
       Fork grows the toolbar back to the editable height, and a reserve that
       tracked `readOnly` would spend that growth shoving the footer down the moment it is
       pressed. /b/:id carries no CLS budget of its own -- it is server-rendered, so the
       island's whole planner arrives after first paint regardless of what this reserves. -->
  <div class="flex min-h-[1411px] flex-col gap-[22px] md:min-h-[1037px] md:gap-8">
    {#if status === 'loading'}
      <!-- The planner's own panel chrome rather than a bare line on a blank reserve: a
           viewport of empty space reads as a broken page, and the frame reads as the planner
           arriving. It cannot show the trees themselves -- their names, tiers and columns are
           the very thing still loading -- so it grows to fill the reserve and says so. -->
      <div class="border-line bg-raised rounded-panel mx-[18px] flex grow flex-col gap-3 border p-4 md:mx-0">
        <p class="text-muted text-[14px]" role="status">Loading talent data</p>
      </div>
    {:else if status === 'failed'}
      <div class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-3 border p-5 md:mx-0">
        <p class="text-strong text-[15px] font-semibold">{DATA_LOAD_FAILED}</p>
        <!-- Honest for every branch of `load`: the reference files fail into this state too,
             not just talents/<class>.json. -->
        <p class="text-muted text-[13px]">
          Build {store.treeVersion} did not return the files the planner needs.
        </p>
        <button
          type="button"
          class="{SECONDARY_BUTTON} border-line-warm-strong text-text w-fit px-4"
          onclick={() => (attempt += 1)}
        >
          Retry
        </button>
      </div>
    {:else if store.talentIndex}
      <!-- One tree at a time on a phone: three trees side by side do not fit 360px, and
           stacking them puts the third one two screens down. Desktop keeps the columns and
           hides this. The roving tabindex lives on the tabs, as it does on TreeGrid's cells.
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
        {#each store.talentIndex.trees as tree, i (tree.id)}
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
      </div>

      <div class="grid grid-cols-1 gap-4 px-[18px] md:grid-cols-3 md:px-0" data-testid="tree-columns">
        {#each store.talentIndex.trees as tree, i (tree.id)}
          <!-- The inactive trees are hidden with a class, not the `hidden` attribute: the
               attribute would hide them on desktop too, where `md:flex` cannot override it. -->
          <div
            id={`tree-panel-${tree.id}`}
            role="tabpanel"
            aria-labelledby={`tree-tab-${tree.id}`}
            class="border-line bg-raised rounded-panel flex-col gap-3 border p-4 md:flex {activeTree === i
              ? 'flex'
              : 'hidden'}"
          >
            <header class="flex items-baseline justify-between">
              <h2 class="section-title text-[15px]">{tree.name}</h2>
              <span class="tabular text-gold font-mono text-[15px]">{store.split[i] ?? 0}</span>
            </header>
            <TreeGrid {store} {tree} />
          </div>
        {/each}
      </div>

      <div class="flex flex-wrap items-center gap-3 px-[18px] md:px-0" data-testid="planner-toolbar">
        {#if store.readOnly}
          <!-- A build opened from a share link. Every edit is refused until Fork, so the
               toolbar says so up front rather than leaving the refusal message to explain it
               after the first click. Reset and Share are gone with it: there is nothing of
               one's own to clear, and re-sharing someone else's build under a new id is the
               one thing Fork is for. -->
          <p class="text-muted text-[13px]">
            This build was shared as a link. Fork it to spend points of your own.
          </p>
          <button
            type="button"
            class="{SECONDARY_BUTTON} border-line-warm-strong text-gold px-4"
            onclick={() => store.fork()}
          >
            Fork
          </button>
        {:else}
          <!-- Nested rather than a third arm of the branch above, so `readOnly` is asked once:
               Reset and Share belong to the same half of that decision, and SharePanel has to
               sit outside the confirm to survive it -- it holds the title being typed and the
               link of the last save, and re-mounting it when the confirm opens would throw
               both away. -->
          {#if confirmingReset}
            <span class="text-muted text-[13px]">Clear every point in this build?</span>
            <button
              type="button"
              class="{SECONDARY_BUTTON} border-line-warm-strong text-gold px-4"
              onclick={() => {
                store.reset();
                confirmingReset = false;
              }}
            >
              Clear all points
            </button>
            <!-- Reset leaves the DOM the moment it is pressed, so the keyboard lands on the
                 question it just asked rather than back at the top of the document. It lands
                 on the safe answer: a second Enter pressed out of habit keeps the build rather
                 than clearing it, which is the only reason the second step exists. -->
            <button
              type="button"
              class="{SECONDARY_BUTTON} border-line-warm text-text px-4"
              {@attach (node) => node.focus()}
              onclick={() => (confirmingReset = false)}
            >
              Keep the build
            </button>
          {:else}
            <button
              type="button"
              class="{SECONDARY_BUTTON} border-line-warm text-text px-4"
              onclick={() => (confirmingReset = true)}
            >
              Reset
            </button>
          {/if}

          <SharePanel {store} />
        {/if}
      </div>

      <OrderStrip {store} />

      <!-- Gear is the optional half of a build. A build with no item file has an empty index
           and no panel at all, rather than seventeen slots nothing can ever fill. -->
      {#if store.itemIndex.size > 0}
        <GearPanel {store} />
      {/if}
    {/if}
  </div>
</div>
