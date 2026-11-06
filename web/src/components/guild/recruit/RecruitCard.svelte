<!-- web/src/components/guild/recruit/RecruitCard.svelte -->
<!-- The public recruiting card, the page's one Plate (web layout spec 2026-11-05 §7.2 Recruitment):
     whether they recruit, the raid nights in the officers' words and the next planned nights in the
     viewer's clock, progress this tier, and what they look for (class crest, spec, the note). -->
<script lang="ts">
  import type { RecruitCard } from '../../../lib/guild/recruit-api';
  import { recruitCopy as copy } from '../../../lib/guild/recruit-copy';
  import { progressLine, upcomingNights, wantedRows } from '../../../lib/guild/recruit-view';
  import { classTone } from '../../../lib/ui/primitives';
  import ClassCrestRing from '../../character/ClassCrestRing.svelte';
  import Plate from '../../ui/Plate.svelte';
  import Row from '../../ui/Row.svelte';
  import Tag from '../../ui/Tag.svelte';

  let { card, viewerTimezone, locale }: { card: RecruitCard; viewerTimezone: string; locale: string } =
    $props();

  const upcoming = $derived(upcomingNights(card, viewerTimezone, locale));
  const progress = $derived(progressLine(card));
  const wanted = $derived(wantedRows(card.wanted));
</script>

<Plate label={copy.card.label} testid="guild-recruit-card">
  <Row title={card.guild.name} testid="guild-recruit-status">
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
    {#if wanted.length === 0}
      <Row title={copy.card.lookingFor} sub={copy.card.nobodyListed} testid="guild-recruit-wanted-any" />
    {:else}
      {#each wanted as line (line.key)}
        <Row
          title={line.title}
          titleTone={classTone(line.classSlug)}
          sub={line.note === '' ? copy.card.lookingFor : line.note}
          testid="guild-recruit-wanted"
        >
          {#snippet lead()}<ClassCrestRing characterClass={line.classSlug} size={32} />{/snippet}
        </Row>
      {/each}
    {/if}
  {/if}
</Plate>
