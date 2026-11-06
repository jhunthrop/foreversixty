<!-- web/src/components/guild/overview/CompTally.svelte -->
<!-- Comp against the target (web layout spec 2026-11-05 §7.1, owner decision D13): In per role
     against the raid's seats, Late named under the count; the target is the zone's default until an
     officer sets one on the raid planner. -->
<script lang="ts">
  import type { GuildRaidView } from '../../../lib/guild/api';
  import type { GuildCompTarget } from '../../../lib/guild/overview-api';
  import { compTallyRows, compTargetLine } from '../../../lib/guild/overview-role-view';
  import Section from '../../ui/Section.svelte';
  import Tally from '../../ui/Tally.svelte';

  let {
    raid,
    target,
    title = 'Comp against the target',
    editHref,
  }: { raid: GuildRaidView; target: GuildCompTarget; title?: string; editHref?: string } = $props();

  const items = $derived(
    compTallyRows(raid.counts, target).map((row) => ({
      key: row.role,
      label: row.label,
      have: row.have,
      want: row.want,
      note: row.late > 0 ? `${row.late} late` : undefined,
    })),
  );
</script>

<Section
  {title}
  meta={compTargetLine(raid.zone_name, target)}
  action={editHref === undefined ? undefined : { label: 'Edit target', href: editHref }}
  testid="guild-overview-comp"
>
  <Tally {items} testid="guild-comp-tally" />
</Section>
