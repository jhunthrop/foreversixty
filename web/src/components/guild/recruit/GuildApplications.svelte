<!-- web/src/components/guild/recruit/GuildApplications.svelte -->
<!-- /guild/{region}/{ruleset}/{slug}/applications (web layout spec 2026-11-05 §7.2 Applications, "Do
     we take them?"; mock web-guild-applications), officers only. The one Plate holds the open
     applications on one figure grid (ApplicantRow) under a head line, each with the seat its spec
     fills for the next raid; a decision made here says what happened on Discord in a Notice above
     them. The rail is the next raid's open seats with who is applying for each, the public card (its
     fields, each opening the edit) and the record of recent decisions.
     GuildShell loads this file as its own chunk; lib/guild/applications-page.svelte.ts owns the reads and writes. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import { guildHref, guildRecruitHref, rulesetLabel, type CharacterPath } from '../../../lib/characters';
  import { bandFacts } from '../../../lib/guild/fact-line';
  import { decisionTitle, discordLines, splitApplications } from '../../../lib/guild/applications-view';
  import { ApplicationsPage } from '../../../lib/guild/applications-page.svelte';
  import { liveApplicationsDeps } from '../../../lib/guild/guild-view-deps';
  import { fetchGuildHome } from '../../../lib/guild/api';
  import { fetchGuildRaid } from '../../../lib/guild/raids-api';
  import { compGaps, type CompGapRow } from '../../../lib/guild/dense/week-board';
  import { GROUP_WORD } from '../../../lib/guild/applicant-figures';
  import { medianItemLevel } from '../../../lib/guild/planner-facts';
  import { raidNightDay } from '../../../lib/guild/overview-view';
  import Row from '../../ui/Row.svelte';
  import { APPLICANT_ACTIONS, APPLICANT_COLUMNS, APPLICANT_HEADS } from './applicant-columns';
  import { applicationsCopy as copy } from '../../../lib/guild/recruit-copy';
  import LoadError from '../../ui/LoadError.svelte';
  import Notice from '../../ui/Notice.svelte';
  import Page from '../../ui/Page.svelte';
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

  /** The next raid's seats against its comp target, and its start; null until read (or none planned). */
  let gaps = $state<readonly CompGapRow[] | null>(null);
  let nextAt = $state<string | null>(null);
  let median = $state<number | null>(null);
  $effect(() => {
    const guildId = p.context?.guild.id;
    if (guildId === undefined || p.load !== 'ready') return;
    void untrack(async () => {
      try {
        const home = await fetchGuildHome(guildId);
        median = medianItemLevel(home.roster);
        const next = home.next_raid;
        if (next == null) return;
        const raid = await fetchGuildRaid(guildId, next.id);
        gaps = compGaps(raid);
        nextAt = raid.starts_at;
      } catch {
        gaps = null;
      }
    });
  });
  const shortRows = $derived((gaps ?? []).filter((row) => row.gap > 0));
  const applying = (group: string): number =>
    (split?.pending ?? []).filter((app) => app.character.group === group).length;
</script>

<!-- One root: GuildShell lays its views out in a gapped flex column, and the head and the Page
     below it are one block with the Page's own gap. -->
<div>
  {#if p.load === 'failed'}
    <LoadError
      message={copy.page.loadError}
      onRetry={() => void p.start()}
      testid="guild-applications-error"
    />
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
        facts={bandFacts(rulesetLabel(path.ruleset), path.region)}
        tabRole="officer"
        views="applications"
        testid="guild-applications-head"
      />
    {/snippet}
    {#snippet rail()}
      {#if shortRows.length > 0 && nextAt !== null}
        <Section title={`Short for ${raidNightDay(nextAt)}`} panel testid="guild-applications-short">
          <ul>
            {#each shortRows as row (row.group)}
              <li>
                <Row
                  title={GROUP_WORD[row.group]}
                  sub={`${row.in + row.late} of ${row.need} · ${applying(row.group)} applying`}
                  density="slim"
                  open
                >
                  {#snippet trail()}<span class="text-wipe text-[14px] font-bold tabular-nums"
                      >{row.gap} short</span
                    >{/snippet}
                </Row>
              </li>
            {/each}
          </ul>
        </Section>
      {/if}
      <RecruitEdit
        {card}
        roleId={list.discord_role_id}
        recruitHref={guildRecruitHref(guildPath)}
        saving={p.saving}
        line={p.saveLine}
        onSave={(draft) => void p.saveCard(draft)}
      />
      <Section title={copy.recent} panel testid="guild-applications-recent">
        {#if split.decided.length === 0}
          <p class="text-muted text-[14px]">{copy.noRecent}</p>
        {:else}
          {#each split.decided as application (application.id)}
            <DecisionRow {application} />
          {/each}
        {/if}
      </Section>
    {/snippet}
    <Page layout="split" panels {header} {rail} testid="guild-applications-page">
      {#if decided !== null}
        <Notice
          text={[decisionTitle(decided), ...discordLines(decided)].join('. ') + '.'}
          tone="neutral"
          testid="guild-applications-outcome"
        />
      {/if}
      <Section title={copy.open} panel testid="guild-applications-open">
        {#if split.pending.length === 0}
          <p class="text-muted text-[14px]" data-testid="guild-applications-none">{copy.none}</p>
        {:else}
          <div role="region" aria-label={copy.open} data-testid="guild-applications-plate">
            <div class="hidden justify-end gap-x-3 pb-1 xl:flex" aria-hidden="true">
              <span class="{APPLICANT_COLUMNS} text-label text-muted grid font-bold">
                {#each APPLICANT_HEADS as head (head)}<span>{head}</span>{/each}
              </span>
              <span class={APPLICANT_ACTIONS}></span>
            </div>
            {#each split.pending as application (application.id)}
              <ApplicantRow
                {gaps}
                {median}
                {application}
                busy={p.busyId === application.id}
                error={p.rowError?.id === application.id ? p.rowError.text : ''}
                onVote={(vote) => void p.vote(application.id, vote)}
                onDecide={(decision) => void p.decide(application.id, decision)}
              />
            {/each}
          </div>
        {/if}
      </Section>
    </Page>
  {/if}
</div>
