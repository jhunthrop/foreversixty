<!-- web/src/components/guild/recruit/GuildViewHeader.svelte -->
<!-- The head of a view under a guild page (recruit, applications, a member profile): the one
     PageHeader with the guild's mark and faction-colour title, its fact line and, for a member or
     officer, the guild's tab strip linking back to the page (the mocks carry no separate back link;
     the site header's Guild door and the tabs lead back). `views` adds the Roster's sub-views
     (Members, Applications) under the tabs, the current one marked (batch 3). -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { GuildViewGuild } from '../../../lib/guild/guild-view-load';
  import type { TitleTone } from '../../../lib/ui/primitives';
  import FactionCrest from '../../character/FactionCrest.svelte';
  import FactLine from '../FactLine.svelte';
  import { guildApplicationsHref } from '../../../lib/characters';
  import GuildTabs, { type GuildTabId } from '../GuildTabs.svelte';
  import TabStrip from '../../ui/TabStrip.svelte';
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
    views,
    meta: metaContent,
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
    /** The Roster's sub-view row, its current view. */
    views?: 'members' | 'applications';
    /** The meta line as markup (a Tag), in place of the facts. */
    meta?: Snippet;
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
{#snippet tabs()}
  <GuildTabs active={activeTab} role={tabRole ?? 'member'} hrefBase={guildPath} />
  {#if views !== undefined}
    <TabStrip
      tabs={[
        { id: 'members', label: 'Members', href: `${guildPath}#roster` },
        { id: 'applications', label: 'Applications', href: guildApplicationsHref(guildPath) },
      ]}
      active={views}
      label="Roster views"
      testid="guild-roster-views"
    />
  {/if}
{/snippet}

<PageHeader
  title={title ?? guild.name}
  {tone}
  {crest}
  meta={metaContent ?? (facts.length > 0 ? meta : undefined)}
  {answer}
  tabs={tabRole === undefined ? undefined : tabs}
  {testid}
/>
