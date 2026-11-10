<!-- web/src/components/guides/GuideBisDoll.svelte -->
<!-- The guide's Best in slot doll (owner 2026-11-09): the level 60 picks on the shared paperdoll,
     the frame /bis draws, in the section's reading column beside the lane's slot list. It opens on
     the guide's default faction and follows the header character's faction once the answer
     settles. Hover, focus or a tap puts a slot's client tooltip in the frame's centre; the full
     list link and the band links stay in the lane. The lane already lists every pick, so under the
     frame's minimum width (the phone's list form) the doll steps aside instead of listing them twice. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import PaperdollFrame from '../ui/PaperdollFrame.svelte';
  import ItemTooltip from '../ItemTooltip.svelte';
  import { bisCopy } from '../../lib/bis/copy';
  import { dollForFaction, type GuideDollView } from '../../lib/guides/guide-bis-doll';
  import type { ItemTooltipModel } from '../../lib/items/tooltip';
  import { defaultSelection, type DollCell } from '../../lib/ui/paperdoll';

  let {
    views,
    tooltips,
    build,
    title,
  }: {
    /** Each faction's level 60 doll, the guide's default first (lib/guides/guide-bis-doll.ts). */
    views: readonly GuideDollView[];
    tooltips: Readonly<Record<number, ItemTooltipModel>>;
    build: string;
    /** The spec's full name, the frame's title bar. */
    title: string;
  } = $props();

  let faction = $state<string | undefined>(undefined);
  let selected = $state<string | null>(null);

  const view = $derived(dollForFaction(views, faction));
  const shownSlot = $derived(
    selected !== null && view?.cells.some((cell) => cell.slot === selected)
      ? selected
      : defaultSelection(view?.cells ?? []),
  );

  onMount(() => {
    let stop: (() => void) | undefined;
    void import('../../lib/guides/guide-character').then(({ onGuideCharacter }) => {
      stop = onGuideCharacter((character) => {
        faction = character?.faction ?? undefined;
      });
    });
    return () => stop?.();
  });
</script>

{#snippet centre(cell: DollCell | undefined)}
  {@const shown = cell?.item ?? cell?.coveredBy}
  {@const tip = shown ? tooltips[shown.id] : undefined}
  {#if tip}<ItemTooltip model={tip} {build} floating={false} />{/if}
{/snippet}

{#if view !== undefined}
  <div class="guide-doll" data-testid="guide-bis-doll" data-faction={view.faction}>
    <div class="guide-doll-frame">
      <PaperdollFrame
        cells={view.cells}
        {build}
        {title}
        meta={bisCopy.frameMeta('60', view.raceName, view.preRaid)}
        selected={shownSlot}
        onselect={(slot) => (selected = slot)}
        {centre}
        label={`${title} best in slot, level 60`}
        testid="guide-bis-frame"
      />
    </div>
  </div>
{/if}

<style>
  .guide-doll {
    container-type: inline-size;
  }
  .guide-doll-frame {
    max-width: 640px;
    margin-top: 24px;
  }
  /* A section with no prose of its own opens on the doll. */
  :global(.ui-prose:empty + astro-island) .guide-doll-frame {
    margin-top: 0;
  }
  /* Under the frame's minimum (lib/ui/paperdoll-frame.ts MIN_FRAME_PX) the lane's list says it all. */
  @container (max-width: 422.98px) {
    .guide-doll-frame {
      display: none;
    }
  }
</style>
