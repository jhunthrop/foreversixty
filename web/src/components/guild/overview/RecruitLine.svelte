<!-- web/src/components/guild/overview/RecruitLine.svelte -->
<!-- Recruitment on a visitor's Overview (web layout spec 2026-11-05 §7.2): a recruiting guild gets one
     Notice to its public card. It reads the card itself and draws nothing until it answers, if it
     fails, or while the guild is not recruiting. Officers reach applications from Roster. -->
<script lang="ts">
  import { guildRecruitHref } from '../../../lib/characters';
  import { fetchRecruitCard, type RecruitCard } from '../../../lib/guild/recruit-api';
  import Notice from '../../ui/Notice.svelte';

  let { guildId, guildPath }: { guildId: number; guildPath: string } = $props();

  let card = $state<RecruitCard | null>(null);

  $effect(() => {
    void fetchRecruitCard(guildId)
      .then((answer) => (card = answer))
      .catch(() => (card = null));
  });
</script>

{#if card?.open}
  <Notice
    text={`${card.guild.name} is recruiting.`}
    action={{ label: 'See what they look for', href: guildRecruitHref(guildPath) }}
    testid="guild-overview-recruit"
  />
{/if}
