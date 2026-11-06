<!-- web/src/components/guild/overview/RecruitLine.svelte -->
<!-- Recruitment on the Overview (web layout spec 2026-11-05 §7.2): a visitor of a recruiting guild
     gets one Notice to its card; an officer gets one Row to the applications and the card's edit,
     open or closed. It reads the public card itself and draws nothing until it answers, or if it
     fails. Members get neither. -->
<script lang="ts">
  import { guildApplicationsHref, guildRecruitHref } from '../../../lib/characters';
  import { fetchRecruitCard, type RecruitCard } from '../../../lib/guild/recruit-api';
  import Notice from '../../ui/Notice.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';

  let { guildId, guildPath, officer }: { guildId: number; guildPath: string; officer: boolean } = $props();

  let card = $state<RecruitCard | null>(null);

  $effect(() => {
    void fetchRecruitCard(guildId)
      .then((answer) => (card = answer))
      .catch(() => (card = null));
  });
</script>

{#if card !== null}
  {#if officer}
    <Section title="Recruitment" testid="guild-overview-recruit">
      <Row
        title={card.open ? 'Taking applications' : 'Not taking applications'}
        sub="Applications, votes and your public card"
        href={guildApplicationsHref(guildPath)}
        open
        testid="guild-overview-recruit-link"
      />
    </Section>
  {:else if card.open}
    <Notice
      text={`${card.guild.name} is recruiting.`}
      action={{ label: 'See what they look for', href: guildRecruitHref(guildPath) }}
      testid="guild-overview-recruit"
    />
  {/if}
{/if}
