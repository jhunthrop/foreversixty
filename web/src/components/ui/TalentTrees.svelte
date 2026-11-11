<script lang="ts">
  import type { Snippet } from 'svelte';
  import { radioKeyTarget } from '../../lib/ui/control-bar';
  import type { PanelFrame } from '../../lib/ui/primitives';
  import {
    MODE_CELL_SIZE,
    openTreeIndex,
    type TalentCellSize,
    type TalentTreesCopy,
    type TalentTreesMode,
    type TalentTreeView,
  } from '../../lib/ui/talent-tree';
  import Plate from './Plate.svelte';
  import Section from './Section.svelte';
  import TalentTree from './TalentTree.svelte';

  let {
    trees,
    build,
    mode,
    size,
    active = $bindable(),
    label = 'Talent trees',
    copy,
    onadd,
    onremove,
    onresettree,
    foot,
    title,
    frame = 'plain',
    tipScope = false,
    root = $bindable(),
    testid,
  }: {
    trees: readonly TalentTreeView[];
    /** The data build the icons and tooltips come from. */
    build: string;
    mode: TalentTreesMode;
    /** The cell size token; the mode's own by default (edit lg, view md, compact sm). */
    size?: TalentCellSize;
    /** The open tree below 768 (lg and md); the tree with the most points until one is picked. */
    active?: number;
    /** The trees' accessible name (the editor's Plate, the tab list). */
    label?: string;
    /** `edit` only: the per-tree Reset's words. */
    copy?: TalentTreesCopy;
    onadd?: (id: number) => void;
    onremove?: (id: number) => void;
    /** `edit` only: the per-tree Reset in each tree head (shown when the tree has points). */
    onresettree?: (treeIndex: number) => void;
    /** `edit` only: the status line and actions under the trees. */
    foot?: Snippet;
    /** `edit` only: a heading inside the panel (the planner's "Talents"), the trees in a Section. */
    title?: string;
    /** `title` only: the panel's frame. */
    frame?: PanelFrame;
    /** The trees' tooltips belong to a host of their own (TalentTreesWithTips): a page's host
     *  leaves them alone. */
    tipScope?: boolean;
    /** The trees' box, for that host's scope. */
    root?: HTMLDivElement;
    testid?: string;
  } = $props();

  const editable = $derived(mode === 'edit');
  const cellSize = $derived(size ?? MODE_CELL_SIZE[mode]);
  const tabbed = $derived(cellSize !== 'sm');
  const uid = $props.id();
  const open = $derived(active ?? openTreeIndex(trees));

  function onTabKey(event: KeyboardEvent, index: number): void {
    const target = radioKeyTarget(event.key, index, trees.length);
    if (target === null) return;
    event.preventDefault();
    active = target;
    const tablist = (event.currentTarget as HTMLElement).parentElement;
    tablist?.querySelector<HTMLElement>(`[data-testid="tree-tab-${target}"]`)?.focus();
  }
</script>

<!-- web/src/components/ui/TalentTrees.svelte -->
<!-- Every talent tree on the site (talent-tree lane 2026-11-10; planner spec 2026-11-08 section 3A,
     4.4; guides spec 2026-11-08 §3A): a build's trees in one of three modes with one look, one
     tooltip and one keyboard. `edit` (the planner) on the page's Plate or a titled Section, with
     the per-tree Reset and a foot; `view` (a read-only build: the sim, /sim/talents) bare, at the md
     cell; `compact` (a guide) three trees fitted to the box. lg and md show the tree tabs and the
     open tree below 768, one column per tree from 768; sm keeps the trees side by side and stacks
     them under the 18 px floor. The look lives in components/ui/talent-tree.css. -->

{#if editable && title}
  <Section {title} panel {frame} flush {testid}>{@render body()}</Section>
{:else if editable}
  <Plate {label} {testid}>{@render body()}</Plate>
{:else}
  <div data-testid={testid}>{@render body()}</div>
{/if}

{#snippet body()}
  <div
    class="tt-trees"
    data-size={cellSize}
    data-mode={mode}
    data-bare={editable ? undefined : ''}
    data-tip-scope={tipScope ? '' : undefined}
    bind:this={root}
  >
    {#if tabbed}
      <div class="tt-tabs" role="tablist" aria-label={label}>
        {#each trees as tree, index (tree.id)}
          <button
            type="button"
            role="tab"
            id={`${uid}-tab-${index}`}
            aria-selected={index === open}
            aria-controls={`${uid}-panel-${index}`}
            tabindex={index === open ? 0 : -1}
            data-testid={`tree-tab-${index}`}
            onclick={() => (active = index)}
            onkeydown={(event) => onTabKey(event, index)}
          >
            <span
              >{#if tree.iconSrc !== ''}<img src={tree.iconSrc} alt="" width="22" height="22" />{/if}<b
                >{tree.points}</b
              ></span
            >{tree.name}
          </button>
        {/each}
      </div>
    {/if}
    <div class="tt-cols" style:--n={trees.length}>
      {#each trees as tree, index (tree.id)}
        <div
          class="tt-col"
          class:on={index === open}
          role={tabbed ? 'tabpanel' : undefined}
          id={tabbed ? `${uid}-panel-${index}` : undefined}
          aria-label={tree.name}
          data-testid={`tree-panel-${tree.id}`}
        >
          <TalentTree
            {tree}
            {build}
            {editable}
            {copy}
            {onadd}
            {onremove}
            onreset={onresettree ? () => onresettree(index) : undefined}
          />
        </div>
      {/each}
    </div>
    {#if foot}<div class="tt-f">{@render foot()}</div>{/if}
  </div>
{/snippet}
