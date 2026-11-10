<!-- web/src/components/guild/overview/NextRaidPlate.svelte -->
<!-- Next raid (batch 3 mock), a panel: the raid's Row (its icon, the guild's time and the viewer's own,
     the countdown), then the viewer's own Row (`you`: RaidBandLine's row, which writes the same sign-up
     the Discord buttons write). Officers also get the In, Late, Bench and Not signed figures and the
     planner link. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { GuildRaidView, NextRaidView } from '../../../lib/guild/api';
  import Figure from '../../ui/Figure.svelte';
  import Section from '../../ui/Section.svelte';
  import NextRaidRow from './NextRaidRow.svelte';

  let {
    next,
    raid,
    officer,
    planHref,
    viewerTz,
    now,
    you,
  }: {
    next: NextRaidView;
    /** The full raid read, for the officer's counts; null until it answers. */
    raid: GuildRaidView | null;
    officer: boolean;
    planHref: string;
    viewerTz: string;
    now: Date;
    you?: Snippet;
  } = $props();

</script>

<Section
  title="Next raid"
  panel
  action={officer ? { label: 'Open the planner', href: planHref } : undefined}
  testid="guild-overview-next-raid"
>
  <NextRaidRow {next} durationMinutes={raid?.duration_minutes ?? 180} {viewerTz} {now} />
  {#if you}<div data-testid="guild-next-raid-you">{@render you()}</div>{/if}
  {#if officer && raid !== null}
    <div class="flex flex-wrap gap-x-8 gap-y-3 pt-3" data-testid="guild-next-raid-figures">
      <Figure value={String(raid.counts.in)} label="In" tone="kill" />
      <Figure value={String(raid.counts.late)} label="Late" tone="gold" />
      <Figure value={String(raid.counts.bench)} label="Bench" />
      <Figure
        value={String(raid.not_signed_raiders?.length ?? raid.not_signed.length)}
        label="Not signed"
        tone="wipe"
      />
    </div>
  {/if}
</Section>
