<!-- web/src/components/guild/recruit/GuildRecruit.svelte -->
<!-- /guild/{region}/{ruleset}/{slug}/recruit (batch 3, b3-guild-recruit; web layout spec 2026-11-05
     §7.2 Recruitment, "Should I apply?"): where a stranger meets the guild. The band reads Join and
     whether they recruit; the framed panel is who they look for, then what the viewer can do (sign
     in, apply, their own application, or the line that the guild is closed); the rail holds the
     guild's facts, the officers' pitch and the officers. GuildShell loads this file as its own
     chunk; lib/guild/recruit-page.svelte.ts owns the reads and writes. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import { guildApplicationsHref, guildHref, type CharacterPath } from '../../../lib/characters';
  import { fetchOfficerApplications } from '../../../lib/guild/recruit-api';
  import { liveRecruitDeps } from '../../../lib/guild/guild-view-deps';

  import { recruitCopy as copy } from '../../../lib/guild/recruit-copy';
  import { RecruitPage } from '../../../lib/guild/recruit-page.svelte';
  import { officerRows, realmLabel, wantedRows } from '../../../lib/guild/recruit-view';
  import { classTone } from '../../../lib/ui/primitives';
  import Crest from '../../ui/Crest.svelte';
  import { battlenetStartUrl } from '../../../lib/account/api';
  import { viewerLocale, viewerZone } from '../../../lib/dates';
  import Button from '../../ui/Button.svelte';
  import LoadError from '../../ui/LoadError.svelte';
  import Notice from '../../ui/Notice.svelte';
  import Page from '../../ui/Page.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';
  import Skeleton from '../../ui/Skeleton.svelte';
  import Tag from '../../ui/Tag.svelte';
  import ApplyForm from './ApplyForm.svelte';
  import GuildViewHeader from './GuildViewHeader.svelte';
  import MyApplication from './MyApplication.svelte';
  import RecruitCard from './RecruitCard.svelte';

  let {
    path,
    page: given,
    viewerTimezone = viewerZone(),
    locale = viewerLocale(),
  }: {
    path: CharacterPath;
    /** A test passes a seeded page; the view builds its own otherwise. */
    page?: RecruitPage;
    viewerTimezone?: string;
    locale?: string;
  } = $props();

  // svelte-ignore state_referenced_locally
  const p = given ?? new RecruitPage(liveRecruitDeps(), path);
  // svelte-ignore state_referenced_locally
  const own = given === undefined;
  $effect(() => (own ? void untrack(() => p.start()) : undefined));

  const guildName = $derived(p.context?.guild.name ?? '');
  const guildPath = $derived(guildHref(path.region, path.ruleset, guildName || path.slug));
  const role = $derived(p.context?.role);
  const tabRole = $derived(role === 'officer' ? 'officer' : role === 'member' ? 'member' : undefined);
  /** An officer reads the open applications against this card; null until read. */
  let openApplications = $state<number | null>(null);
  $effect(() => {
    const guildId = p.context?.guild.id;
    if (role !== 'officer' || guildId === undefined) return;
    void untrack(() =>
      fetchOfficerApplications(guildId)
        .then((list) => (openApplications = list.pending))
        .catch(() => (openApplications = null)),
    );
  });
  const next = typeof location === 'undefined' ? '/account?signed_in=1' : location.pathname;
</script>

<!-- One root: GuildShell lays its views out in a gapped flex column, and the head and the Page
     below it are one block with the Page's own gap. -->
<div>
  {#snippet rail()}
    {#if p.card !== null}
      <RecruitCard card={p.card} {viewerTimezone} {locale} />
      {#if p.card.pitch !== ''}
        <Section title={copy.card.why} panel testid="guild-recruit-pitch">
          <p class="text-text text-[14px] leading-[22px] whitespace-pre-line">{p.card.pitch}</p>
        </Section>
      {/if}
      {#if p.card.officers.length > 0}
        <Section title={copy.card.officers} panel testid="guild-recruit-officers">
          {#each officerRows(p.card.officers) as officer (officer.key)}
            <Row title={officer.name} titleTone={classTone(officer.classSlug)} sub={officer.sub} open>
              {#snippet lead()}<Crest cls={officer.classSlug} size={32} />{/snippet}
            </Row>
          {/each}
        </Section>
      {/if}
    {/if}
  {/snippet}

  {#if p.load === 'failed'}
    <LoadError message={copy.page.loadError} onRetry={() => void p.start()} testid="guild-recruit-error" />
  {:else if p.load === 'loading' || p.context === null || p.card === null || p.state === null}
    <div class="min-h-[640px]" data-testid="guild-recruit-loading">
      <Skeleton lines={6} label={copy.page.loading} testid="guild-recruit-skeleton" />
    </div>
  {:else}
    {@const state = p.state}
    {@const card = p.card}
    {@const context = p.context}
    {#snippet status()}
      <Tag tone={card.open ? 'success' : 'muted'} testid="guild-recruit-open"
        >{card.open ? copy.card.recruiting : copy.card.closed}</Tag
      >
    {/snippet}
    {#snippet header()}
      <GuildViewHeader
        guild={context.guild}
        {guildPath}
        title={copy.page.title(card.guild.name)}
        meta={status}
        {tabRole}
        testid="guild-recruit-head"
      />
    {/snippet}
    <Page layout="split" panels {header} {rail} testid="guild-recruit-page">
      {#if card.open}
        <Section title={copy.card.lookingFor} panel testid="guild-recruit-wanted">
          {#if card.wanted.length === 0}
            <p class="text-text text-[14px]">{copy.card.nobodyListed}</p>
          {:else}
            <ul>
              {#each wantedRows(card.wanted) as line (line.key)}
                <li>
                  <Row
                    title={line.title}
                    titleTone={classTone(line.classSlug)}
                    sub={line.note || undefined}
                    open
                  >
                    {#snippet lead()}<Crest cls={line.classSlug} size={30} />{/snippet}
                  </Row>
                </li>
              {/each}
            </ul>
          {/if}
        </Section>
      {/if}
      {#if state.kind === 'closed'}
        <Notice text={copy.closedLine(card.guild.name)} tone="neutral" testid="guild-recruit-closed" />
      {:else if state.kind === 'member' && openApplications !== null}
        <Notice
          text={copy.mine.officerOpen(openApplications)}
          tone="neutral"
          action={{ label: copy.mine.review, href: guildApplicationsHref(guildPath) }}
          testid="guild-recruit-officer"
        />
      {:else if state.kind === 'member'}
        <Notice
          text={copy.mine.member(card.guild.name)}
          tone="neutral"
          action={{ label: copy.mine.openGuild, href: guildPath }}
          testid="guild-recruit-member"
        />
      {:else if state.kind === 'application'}
        <MyApplication
          application={state.application}
          guildName={card.guild.name}
          {viewerTimezone}
          busy={p.busy}
          error={p.error}
          onWithdraw={() => void p.withdraw()}
        />
      {:else}
        <Section title={copy.apply.heading} panel testid="guild-recruit-apply">
          {#if state.kind === 'signed-out'}
            <Button tone="primary" href={battlenetStartUrl(next)} testid="guild-recruit-sign-in"
              >{copy.apply.signIn}</Button
            >
          {:else if state.kind === 'no-character'}
            <p class="text-text text-[14px]" data-testid="guild-recruit-no-character">
              {copy.apply.noCharacter(realmLabel(card))}
            </p>
          {:else}
            <ApplyForm
              characters={state.characters}
              busy={p.busy}
              error={p.error}
              onApply={(input) => void p.apply(input)}
            />
          {/if}
        </Section>
      {/if}
    </Page>
  {/if}
</div>
