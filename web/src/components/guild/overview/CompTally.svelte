<!-- web/src/components/guild/overview/CompTally.svelte -->
<!-- Comp against the target (batch 3 mock, owner decision D13): one row per role, the seats the target
     wants, who is In, and how many short, Late named under the role; the target is the zone's default
     until an officer sets one on the raid planner. On its panel, the grid edge to edge. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { GuildRaidView } from '../../../lib/guild/api';
  import type { GuildCompTarget } from '../../../lib/guild/overview-api';
  import { compGridRows } from '../../../lib/guild/overview-role-view';
  import Grid from '../../ui/Grid.svelte';
  import Section from '../../ui/Section.svelte';

  let {
    raid,
    target,
    title = 'Comp against the target',
    editHref,
    edit,
  }: {
    raid: GuildRaidView;
    target: GuildCompTarget;
    title?: string;
    editHref?: string;
    /** The planner's officer edit, under the grid. */
    edit?: Snippet;
  } = $props();
</script>

<Section
  {title}
  panel
  flush
  action={editHref === undefined ? undefined : { label: 'Edit target', href: editHref }}
  testid="guild-overview-comp"
>
  <Grid
    caption="Comp against the target"
    columns={['Need', 'In', 'Short']}
    rows={compGridRows(raid.counts, target)}
    edge
    testid="guild-comp-grid"
  />
  {#if edit}<div class="px-5 pt-4 pb-2">{@render edit()}</div>{/if}
</Section>
