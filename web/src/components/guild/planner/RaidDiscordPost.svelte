<!-- web/src/components/guild/planner/RaidDiscordPost.svelte -->
<!-- Officers: the raid's Discord post. Post now for a raid that has not posted (saved with the box
     off, or a send that did not land), Repost for one that has (a buried or deleted message: the bot
     deletes the old one). A save already posts at once (owner ruling 2026-11-09), so this is the
     manual lever. Shown only while the guild has a Discord channel and the raid is still open. -->
<script lang="ts">
  import type { GuildRaidView } from '../../../lib/guild/api';
  import { channelLabel } from '../../../lib/guild/raid-form';
  import { raidPostCopy } from '../../../lib/guild/raid-copy';
  import Button from '../../ui/Button.svelte';

  let {
    raid,
    when,
    officerTools,
    posting = false,
    answer = null,
    onPost,
  }: {
    raid: GuildRaidView;
    /** A stamp in words (`Mon 9 Nov, 10:10 pm`). */
    when: (at: Date) => string;
    officerTools: boolean;
    posting?: boolean;
    answer?: string | null;
    onPost: () => void;
  } = $props();

  const postedAt = $derived(raid.discord_post?.posted_at ?? null);
  const state = $derived(
    postedAt === null
      ? raidPostCopy.notPosted
      : raidPostCopy.postedTo(channelLabel(raid.discord.channel_name), when(new Date(postedAt))),
  );
</script>

<div class="flex flex-wrap items-center gap-x-3 gap-y-1 text-[13px]" data-testid="raid-discord-post">
  <span class="text-muted" data-testid="raid-discord-post-state">{state}</span>
  <Button size="sm" disabled={!officerTools || posting} onclick={onPost} testid="raid-discord-post-button"
    >{postedAt === null ? raidPostCopy.postNow : raidPostCopy.repost}</Button
  >
  {#if answer !== null}
    <span role="status" data-testid="raid-discord-post-answer">{answer}</span>
  {/if}
</div>
