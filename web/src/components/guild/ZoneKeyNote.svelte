<!-- web/src/components/guild/ZoneKeyNote.svelte -->
<!-- `Needs a key: {Item}.` (guild events spec 3.3): under the form's What select and in the event
     page's band under the when line. The name is the shared item link, in its quality colour, with
     a tooltip of the facts the API carries: a key sits in no class file, so nothing is fetched. -->
<script lang="ts">
  import ItemName from '../ui/ItemName.svelte';
  import activeBuild from '../../data/active-build.json';
  import type { GuildKeyItem } from '../../lib/guild/api';
  import { raidCopy } from '../../lib/guild/raid-copy';
  import { keyItemTooltipModel } from '../../lib/guild/loot-icon';
  import ItemHover from '../ItemHover.svelte';

  let { item, testid = 'zone-key-note' }: { item: GuildKeyItem; testid?: string } = $props();

  /** A key opens a door for every class, so the tooltip reads no class's file. */
  const NO_CLASS = '';
</script>

{#snippet keyLink()}<ItemHover
    itemId={item.id}
    classSlug={NO_CLASS}
    build={activeBuild.build}
    model={keyItemTooltipModel(item)}><ItemName name={item.name} quality={item.quality} /></ItemHover
  >{/snippet}

<!-- A key that opens one of two ways in (Stratholme: Undead's service entrance; groups also enter
     through the main gate) is named for the door it opens, never as a key the run needs. -->
<p class="text-muted text-[13px]" data-testid={testid}>
  {#if item.door === undefined}{raidCopy.form
      .needsKey}{@render keyLink()}.{:else}{@render keyLink()}{raidCopy.form.keyOpens(item.door)}{/if}
</p>
