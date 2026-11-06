<!-- web/src/components/guild/recruit/GuildApplications.svelte -->
<!-- /guild/{region}/{ruleset}/{slug}/applications (web layout spec 2026-11-05 §7.2 Applications, "Do
     we take them?"), officers only. The one Plate holds the open applications; a decision made here
     says what happened on Discord in a Notice above them; the record of recent decisions follows.
     The rail is the public card's edit. GuildShell loads this file as its own chunk;
     lib/guild/applications-page.svelte.ts owns the reads and writes. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import { guildHref, guildRecruitHref, type CharacterPath } from '../../../lib/characters';
  import { decisionTitle, discordLines, splitApplications } from '../../../lib/guild/applications-view';
  import { ApplicationsPage } from '../../../lib/guild/applications-page.svelte';
  import { liveApplicationsDeps } from '../../../lib/guild/guild-view-deps';
  import { applicationsCopy as copy } from '../../../lib/guild/recruit-copy';
  import LoadError from '../../ui/LoadError.svelte';
  import Notice from '../../ui/Notice.svelte';
  import Page from '../../ui/Page.svelte';
  import Plate from '../../ui/Plate.svelte';
  import Section from '../../ui/Section.svelte';
  import Skeleton from '../../ui/Skeleton.svelte';
  import ApplicantRow from './ApplicantRow.svelte';
  import DecisionRow from './DecisionRow.svelte';
  import GuildViewHeader from './GuildViewHeader.svelte';
  import RecruitEdit from './RecruitEdit.svelte';

  let { path, page: given }: { path: CharacterPath; page?: ApplicationsPage } = $props();

  // svelte-ignore state_referenced_locally
  const p = given ?? new ApplicationsPage(liveApplicationsDeps(), path);
  // svelte-ignore state_referenced_locally
  const own = given === undefined;
  $effect(() => (own ? void untrack(() => p.start()) : undefined));

  const guildName = $derived(p.context?.guild.name ?? '');
  const guildPath = $derived(guildHref(path.region, path.ruleset, guildName || path.slug));
  const split = $derived(p.list === null ? null : splitApplications(p.list));
  const decided = $derived(p.lastDecision);
</script>

{#if p.load === 'failed'}
  <LoadError message={copy.page.loadError} onRetry={() => void p.start()} testid="guild-applications-error" />
{:else if p.load === 'forbidden' && p.context !== null}
  {@const context = p.context}
  {#snippet forbiddenHeader()}
    <GuildViewHeader
      guild={context.guild}
      {guildPath}
      title={copy.page.title}
      testid="guild-applications-head"
    />
  {/snippet}
  <Page layout="split" header={forbiddenHeader} testid="guild-applications-page">
    <Notice text={copy.page.officersOnly} tone="neutral" testid="guild-applications-forbidden" />
  </Page>
{:else if p.load !== 'ready' || p.context === null || p.card === null || p.list === null || split === null}
  <div class="min-h-[640px]" data-testid="guild-applications-loading">
    <Skeleton lines={6} label={copy.page.loading} testid="guild-applications-skeleton" />
  </div>
{:else}
  {@const context = p.context}
  {@const card = p.card}
  {@const list = p.list}
  {#snippet header()}
    <GuildViewHeader
      guild={context.guild}
      {guildPath}
      title={copy.page.title}
      facts={[copy.openMeta(list.pending), copy.roleNote]}
      tabRole="officer"
      testid="guild-applications-head"
    />
  {/snippet}
  {#snippet rail()}
    <RecruitEdit
      {card}
      roleId={list.discord_role_id}
      recruitHref={guildRecruitHref(guildPath)}
      saving={p.saving}
      line={p.saveLine}
      onSave={(draft) => void p.saveCard(draft)}
    />
  {/snippet}
  <Page layout="split" {header} {rail} testid="guild-applications-page">
    {#if decided !== null}
      <Notice
        text={[decisionTitle(decided), ...discordLines(decided)].join('. ') + '.'}
        tone="neutral"
        testid="guild-applications-outcome"
      />
    {/if}
    <Section title={copy.open} meta={copy.openMeta(split.pending.length)} testid="guild-applications-open">
      {#if split.pending.length === 0}
        <p class="text-muted text-[14px]" data-testid="guild-applications-none">{copy.none}</p>
      {:else}
        <Plate label={copy.open} testid="guild-applications-plate">
          {#each split.pending as application (application.id)}
            <ApplicantRow
              {application}
              busy={p.busyId === application.id}
              error={p.rowError?.id === application.id ? p.rowError.text : ''}
              onVote={(vote) => void p.vote(application.id, vote)}
              onDecide={(decision) => void p.decide(application.id, decision)}
            />
          {/each}
        </Plate>
      {/if}
    </Section>
    <Section title={copy.recent} testid="guild-applications-recent">
      {#if split.decided.length === 0}
        <p class="text-muted text-[14px]">{copy.noRecent}</p>
      {:else}
        {#each split.decided as application (application.id)}
          <DecisionRow {application} />
        {/each}
      {/if}
    </Section>
  </Page>
{/if}
