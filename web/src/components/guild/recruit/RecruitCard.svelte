<!-- web/src/components/guild/recruit/RecruitCard.svelte -->
<!-- The guild's public facts, the recruit rail's first panel (batch 3, b3-guild-recruit; web layout
     spec 2026-11-05 §7.2 Recruitment): where it plays, its side and size, the raid nights in the
     officers' words, the next planned nights in the viewer's clock, and its progress this tier. -->
<script lang="ts">
  import type { RecruitCard } from '../../../lib/guild/recruit-api';
  import { recruitFacts } from '../../../lib/guild/recruit-view';
  import Facts from '../../ui/Facts.svelte';
  import Section from '../../ui/Section.svelte';

  let { card, viewerTimezone, locale }: { card: RecruitCard; viewerTimezone: string; locale: string } =
    $props();
</script>

<Section title={card.guild.name} panel testid="guild-recruit-card">
  <Facts items={recruitFacts(card, viewerTimezone, locale)} testid="guild-recruit-facts" />
</Section>
