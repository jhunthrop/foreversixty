<!-- web/src/components/guild/GuildJoinInvited.svelte -->
<!-- The join card's invite line (spec 2026-11-05 §3.4 row 2): the officer who made the link, by
     their crest and character in class colour, `invited you to join.`; with no inviter known (a
     link made before 0070, or its maker off the roster), `You are invited to join {Guild}.` -->
<script lang="ts">
  import type { ClaimHolder } from '../../lib/guild/api';
  import { guildJoinCopy } from '../../lib/guild/copy';
  import CharacterName from '../ui/CharacterName.svelte';

  let { guildName, inviter = null }: { guildName: string; inviter?: ClaimHolder | null } = $props();
</script>

{#if inviter !== null}
  <p
    class="text-text inline-flex flex-wrap items-center gap-1.5 text-[15px]"
    data-testid="guild-join-invited"
  >
    <CharacterName name={inviter.name} cls={inviter.class} crestSize={24} />
    {guildJoinCopy.invitedBy}
  </p>
{:else}
  <p class="text-text text-[15px]" data-testid="guild-join-invited">{guildJoinCopy.invited(guildName)}</p>
{/if}
