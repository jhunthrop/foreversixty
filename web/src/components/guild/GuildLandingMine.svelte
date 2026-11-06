<!-- web/src/components/guild/GuildLandingMine.svelte -->
<!-- The guild landing's `Your guilds` card (guild round 3): every guild any of the account's
     characters is in, above the lookup, so a player on an alt with no guild still lands one click
     from their raid guild. Rows follow HomeTopGuilds' list recipe: the name as the row link, where
     it lives muted on the right. Not drawn signed out or with no guild. -->
<script lang="ts">
  import { guildLandingCopy as copy } from '../../lib/guild/copy';
  import type { MyGuildRow } from '../../lib/guild/landing-mine';
  import { rowLink } from '../../lib/report/format';
  import Card from '../ui/Card.svelte';

  let { rows }: { rows: readonly MyGuildRow[] } = $props();
</script>

{#if rows.length > 0}
  <Card glyph="guild" title={copy.mineHeading} body="none" testid="guild-landing-mine">
    <ul class="border-line-soft flex flex-col border-t px-5 py-2">
      {#each rows as row (row.key)}
        <li
          class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-2 border-b py-1 text-[14px] last:border-b-0"
        >
          <a class="{rowLink} truncate font-semibold" href={row.href}>{row.name}</a>
          <span class="text-muted text-[13px]">{row.where}</span>
        </li>
      {/each}
    </ul>
  </Card>
{/if}
