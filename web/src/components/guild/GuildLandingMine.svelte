<!-- web/src/components/guild/GuildLandingMine.svelte -->
<!-- The guild landing's `Your guilds` (guild round 3; web layout spec 2026-11-05 §6.1 rail): every
     guild any of the account's characters is in, so a player on an alt with no guild still lands one
     click from their raid guild. Open link Rows: the name, where it lives, the account's rank as a
     Tag. Not drawn signed out or with no guild. -->
<script lang="ts">
  import { guildLandingCopy as copy } from '../../lib/guild/copy';
  import type { MyGuildRow } from '../../lib/guild/landing-mine';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import Tag from '../ui/Tag.svelte';

  let { rows }: { rows: readonly MyGuildRow[] } = $props();
</script>

{#if rows.length > 0}
  <Section title={copy.mineHeading} testid="guild-landing-mine">
    {#each rows as row (row.key)}
      <Row title={row.name} sub={row.where} href={row.href} open>
        {#snippet trail()}
          {#if row.rank !== null}<Tag tone="muted">{row.rank}</Tag>{/if}
        {/snippet}
      </Row>
    {/each}
  </Section>
{/if}
