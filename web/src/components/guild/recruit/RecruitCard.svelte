<!-- web/src/components/guild/recruit/RecruitCard.svelte -->
<!-- The public recruiting card, the page's one Plate (web layout spec 2026-11-05 §7.2 Recruitment;
     mock web-guild-recruit): the guild with its mark, realm, faction and size and whether they
     recruit; the raid nights in the officers' words and the next planned nights in the viewer's
     clock; progress this tier; and one Looking for line (each class or spec with its note). -->
<script lang="ts">
  import type { RecruitCard } from '../../../lib/guild/recruit-api';
  import { recruitCopy as copy } from '../../../lib/guild/recruit-copy';
  import { cardFacts, progressLine, upcomingNights, wantedLine } from '../../../lib/guild/recruit-view';
  import FactionCrest from '../../character/FactionCrest.svelte';
  import Plate from '../../ui/Plate.svelte';
  import Row from '../../ui/Row.svelte';
  import Tag from '../../ui/Tag.svelte';

  let {
    card,
    mark,
    viewerTimezone,
    locale,
  }: { card: RecruitCard; mark: string | null; viewerTimezone: string; locale: string } = $props();

  const upcoming = $derived(upcomingNights(card, viewerTimezone, locale));
  const progress = $derived(progressLine(card));
</script>

<Plate label={copy.card.label} testid="guild-recruit-card">
  <Row title={card.guild.name} sub={cardFacts(card).join(' · ')} testid="guild-recruit-status">
    {#snippet lead()}<FactionCrest faction={card.guild.faction} src={mark} small />{/snippet}
    {#snippet trail()}
      <Tag tone={card.open ? 'success' : 'muted'} dot testid="guild-recruit-open"
        >{card.open ? copy.card.recruiting : copy.card.closed}</Tag
      >
    {/snippet}
  </Row>
  <Row
    title={copy.card.raids}
    sub={card.raid_nights.text ?? copy.card.noNights}
    testid="guild-recruit-nights"
  />
  {#if upcoming.length > 0}
    <Row title={copy.card.upcoming} sub={upcoming.join(', ')} testid="guild-recruit-upcoming" />
  {/if}
  {#if progress !== ''}
    <Row title={copy.card.progress} sub={progress} testid="guild-recruit-progress" />
  {/if}
  {#if card.open}
    <Row title={copy.card.lookingFor} sub={wantedLine(card.wanted)} testid="guild-recruit-wanted" />
  {/if}
</Plate>
