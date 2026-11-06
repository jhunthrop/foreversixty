<!-- web/src/components/guild/overview/BuffCover.svelte -->
<!-- Buff and debuff cover (web layout spec 2026-11-05 §7.1, owner decision D15): what the raid needs
     against the raiders signed In or Late who can give it, each row with the spell's own icon, red
     at none. Buffs, then debuffs, in the catalogue's order. -->
<script lang="ts">
  import activeBuild from '../../../data/active-build.json';
  import type { GuildRaidCoverRow } from '../../../lib/guild/api';
  import { coverIconUrl, coverSections } from '../../../lib/guild/cover-view';
  import { coverGridRows } from '../../../lib/guild/overview-role-view';
  import SpellIcon from '../../coach/SpellIcon.svelte';
  import Grid from '../../ui/Grid.svelte';
  import Section from '../../ui/Section.svelte';

  let { cover }: { cover: readonly GuildRaidCoverRow[] } = $props();

  const SECTION_WORD = { buff: 'Buffs', debuff: 'Debuffs' } as const;
  const sections = $derived(coverSections(cover));
</script>

<Section title="Buff and debuff cover" meta="Signed-up raiders" testid="guild-overview-cover">
  {#each sections as section (section.kind)}
    <Grid
      caption={`${SECTION_WORD[section.kind]}: need against signed-up raiders`}
      labelHeader={SECTION_WORD[section.kind]}
      columns={['Need', 'In']}
      rows={coverGridRows(section.rows)}
      testid={`guild-cover-${section.kind}`}
    >
      {#snippet lead(row)}
        <SpellIcon url={coverIconUrl(activeBuild.build, row.row)} size={24} name={row.row.name} />
      {/snippet}
    </Grid>
  {/each}
</Section>
