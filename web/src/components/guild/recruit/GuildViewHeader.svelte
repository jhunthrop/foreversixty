<!-- web/src/components/guild/recruit/GuildViewHeader.svelte -->
<!-- The head of a view under a guild page (recruit, applications, a member profile): the back
     link to the guild, then the one PageHeader with the guild's mark and faction-colour title, its
     fact line and, for a member or officer, the guild's tab strip linking back to the page. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { GuildViewGuild } from '../../../lib/guild/guild-view-load';
  import type { TitleTone } from '../../../lib/ui/primitives';
  import FactionCrest from '../../character/FactionCrest.svelte';
  import FactLine from '../FactLine.svelte';
  import GuildTabs, { type GuildTabId } from '../GuildTabs.svelte';
  import PageHeader from '../../ui/PageHeader.svelte';

  let {
    guild,
    guildPath,
    title,
    facts = [],
    tabRole,
    activeTab = 'roster',
    answer,
    titleTone,
    lead,
    testid,
  }: {
    guild: GuildViewGuild;
    guildPath: string;
    /** The page's own title; the guild's name when absent. */
    title?: string;
    facts?: readonly string[];
    /** Member or officer: the guild's tabs as links. Absent for a visitor. */
    tabRole?: 'member' | 'officer';
    activeTab?: GuildTabId;
    answer?: Snippet;
    /** The title's colour; the guild's faction when absent. */
    titleTone?: TitleTone;
    /** The mark beside the title; the guild's crest when absent. */
    lead?: Snippet;
    testid: string;
  } = $props();

  const tone = $derived<TitleTone>(
    titleTone ?? (guild.faction === null ? 'strong' : { faction: guild.faction }),
  );
</script>

{#snippet crest()}{#if lead}{@render lead()}{:else}<FactionCrest
      faction={guild.faction}
      src={guild.mark}
    />{/if}{/snippet}
{#snippet meta()}<FactLine facts={facts.map((text) => ({ text }))} />{/snippet}
{#snippet tabs()}<GuildTabs active={activeTab} role={tabRole ?? 'member'} hrefBase={guildPath} />{/snippet}

<div>
  <a
    href={guildPath}
    class="text-muted inline-flex min-h-11 items-center text-[13px] underline"
    data-testid={`${testid}-back`}>← {guild.name}</a
  >
  <PageHeader
    title={title ?? guild.name}
    {tone}
    {crest}
    meta={facts.length > 0 ? meta : undefined}
    {answer}
    tabs={tabRole === undefined ? undefined : tabs}
    {testid}
  />
</div>
