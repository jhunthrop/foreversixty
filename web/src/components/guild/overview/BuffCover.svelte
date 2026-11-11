<!-- web/src/components/guild/overview/BuffCover.svelte -->
<!-- Cover, thinnest first (batch 3 mock, owner decision D15): what the raid needs against the raiders
     signed In or Late who can give it, each row with the spell's own icon, red at none. The gaps first;
     the first six show, the head's "All" opens the rest in place. -->
<script lang="ts">
  import activeBuild from '../../../data/active-build.json';
  import type { GuildRaidCoverRow } from '../../../lib/guild/api';
  import { coverIconUrl, coverThinnestFirst } from '../../../lib/guild/cover-view';
  import { coverGridRows } from '../../../lib/guild/overview-role-view';
  import SpellIcon from '../../ui/SpellIcon.svelte';
  import Button from '../../ui/Button.svelte';
  import Grid from '../../ui/Grid.svelte';
  import Section from '../../ui/Section.svelte';

  /** The rows shown before "All": the gaps and the thinnest lines. */
  const SHOWN = 6;

  let { cover }: { cover: readonly GuildRaidCoverRow[] } = $props();

  let all = $state(false);
  const sorted = $derived(coverThinnestFirst(cover));
  const rows = $derived(coverGridRows(all ? sorted : sorted.slice(0, SHOWN)));
</script>

<Section title="Cover, thinnest first" panel flush testid="guild-overview-cover">
  {#snippet head()}
    {#if sorted.length > SHOWN}
      <Button tone="ghost" size="sm" onclick={() => (all = !all)} testid="guild-cover-all"
        >{all ? 'Fewer' : `All ${sorted.length} ›`}</Button
      >
    {/if}
  {/snippet}
  <Grid
    caption="Buff and debuff cover against signed-up raiders"
    columns={['Need', 'In']}
    {rows}
    edge
    testid="guild-cover"
  >
    {#snippet lead(row)}
      <SpellIcon url={coverIconUrl(activeBuild.build, row.row)} size={24} />
    {/snippet}
  </Grid>
</Section>
