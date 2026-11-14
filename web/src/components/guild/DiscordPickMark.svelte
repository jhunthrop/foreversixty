<!-- web/src/components/guild/DiscordPickMark.svelte -->
<!-- An officer's mark on a character picked on Discord and not yet confirmed by a Battle.net
     sign-in: the `Picked on Discord` tag and the picker's Discord name, so touch reads it as well
     as a mouse; the whole sentence is its accessible name and its title. In a narrow cell the name
     wraps under the tag rather than being cut: who picked is what an officer needs. The one mark
     for the raid roster, the Roster board and the account table. -->
<script lang="ts">
  import { raidCopy } from '../../lib/guild/raid-copy';
  import Tag from '../ui/Tag.svelte';

  let { who, testid }: { who: string; testid?: string } = $props();
</script>

<span
  class="flex max-w-full min-w-0 flex-wrap items-center gap-x-1.5 gap-y-0.5 text-[12px]"
  title={raidCopy.discordClaim(who)}
  data-testid={testid}
>
  <span class="sr-only">{raidCopy.discordClaim(who)}</span>
  <span class="contents" aria-hidden="true"
    ><Tag tone="warn">{raidCopy.discordPicked}</Tag><span class="text-muted min-w-0 break-all"
      >{raidCopy.discordPickedBy(who)}</span
    ></span
  >
</span>
