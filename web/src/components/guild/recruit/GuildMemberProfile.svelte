<!-- web/src/components/guild/recruit/GuildMemberProfile.svelte -->
<!-- /guild/{region}/{ruleset}/{slug}/members/{account} (web layout spec 2026-11-05 §7.2 Member
     profile, "What do I know about this raider?"), an officer's read of anyone or a member's of
     themselves; mock web-guild-profile. The head names the raider in class colour with their figures
     (no guild tabs: the page is about one raider); Sections hold the
     gear gap in the BiS page's unit, the last 4 raids as a Grid and the 30-raid figure, and loot
     this tier. The rail, officers only, is Officer notes: a member's own read carries no notes and
     draws no trace of them. GuildShell loads this file as its own chunk. -->
<script lang="ts">
  import activeBuild from '../../../data/active-build.json';
  import { characterHref, guildHref, type CharacterPath } from '../../../lib/characters';
  import { liveProfileDeps } from '../../../lib/guild/guild-view-deps';
  import { ProfilePage } from '../../../lib/guild/profile-page.svelte';
  import {
    attendanceView,
    gearGapLines,
    lootRows,
    profileFigures,
    profileHead,
  } from '../../../lib/guild/profile-view';
  import { viewerLocale } from '../../../lib/guild/raid-time';
  import { profileCopy as copy } from '../../../lib/guild/recruit-copy';
  import { classTone } from '../../../lib/ui/primitives';
  import { untrack } from 'svelte';
  import ClassCrestRing from '../../character/ClassCrestRing.svelte';
  import SignInPrompt from '../../SignInPrompt.svelte';
  import Figure from '../../ui/Figure.svelte';
  import Grid from '../../ui/Grid.svelte';
  import ItemRow from '../../ui/ItemRow.svelte';
  import LoadError from '../../ui/LoadError.svelte';
  import Notice from '../../ui/Notice.svelte';
  import Page from '../../ui/Page.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';
  import Skeleton from '../../ui/Skeleton.svelte';
  import GuildViewHeader from './GuildViewHeader.svelte';
  import OfficerNotes from './OfficerNotes.svelte';

  let {
    path,
    account,
    page: given,
    viewerTimezone = new Intl.DateTimeFormat().resolvedOptions().timeZone,
    locale = viewerLocale(),
  }: {
    path: CharacterPath;
    /** `me` or the roster's account key (`u:42`). */
    account: string;
    page?: ProfilePage;
    viewerTimezone?: string;
    locale?: string;
  } = $props();

  // svelte-ignore state_referenced_locally
  const p = given ?? new ProfilePage(liveProfileDeps(), path, account);
  // svelte-ignore state_referenced_locally
  const own = given === undefined;
  $effect(() => (own ? void untrack(() => p.start()) : undefined));

  const build = activeBuild.build;
  const guildName = $derived(p.context?.guild.name ?? '');
  const guildPath = $derived(guildHref(path.region, path.ruleset, guildName || path.slug));
  const next = typeof location === 'undefined' ? '/account?signed_in=1' : location.pathname;
  const stateLine: Partial<Record<string, string>> = {
    forbidden: copy.forbidden,
    'not-found': copy.notInGuild,
  };
</script>

<!-- One root: GuildShell lays its views out in a gapped flex column, and the head and the Page
     below it are one block with the Page's own gap. -->
<div>
  {#if p.load === 'failed'}
    <LoadError message={copy.loadError} onRetry={() => void p.start()} testid="guild-profile-error" />
  {:else if p.load === 'loading' || p.context === null}
    <div class="min-h-[640px]" data-testid="guild-profile-loading">
      <Skeleton lines={6} label={copy.loading} testid="guild-profile-skeleton" />
    </div>
  {:else if p.load !== 'ready' || p.profile === null}
    {@const context = p.context}
    {#snippet plainHeader()}
      <GuildViewHeader guild={context.guild} {guildPath} testid="guild-profile-head" />
    {/snippet}
    <Page layout="split" header={plainHeader} testid="guild-profile-page">
      {#if p.load === 'signed-out'}
        <SignInPrompt line={copy.forbidden} {next} testid="guild-profile-sign-in" />
      {:else}
        <Notice text={stateLine[p.load] ?? copy.loadError} tone="neutral" testid="guild-profile-refused" />
      {/if}
    </Page>
  {:else}
    {@const context = p.context}
    {@const profile = p.profile}
    {@const head = profileHead(profile)}
    {@const gap = gearGapLines(profile)}
    {@const attendance = attendanceView(profile, viewerTimezone, locale)}
    {@const loot = lootRows(profile, viewerTimezone, locale)}
    {@const member = profile.member}
    {#snippet crest()}<ClassCrestRing characterClass={head.classSlug} size={56} />{/snippet}
    {#snippet figures()}
      {@const list = profileFigures(profile)}
      <div
        class="grid gap-x-6 gap-y-5"
        style:grid-template-columns={`repeat(${list.length}, minmax(0, 1fr))`}
        data-testid="guild-profile-figures"
      >
        {#each list as figure (figure.key)}
          <Figure value={figure.value} label={figure.label} testid={`guild-profile-figure-${figure.key}`} />
        {/each}
      </div>
    {/snippet}
    {#snippet header()}
      <GuildViewHeader
        guild={context.guild}
        {guildPath}
        title={head.name}
        titleTone={classTone(head.classSlug)}
        lead={crest}
        facts={head.facts}
        answer={figures}
        testid="guild-profile-head"
      />
    {/snippet}
    {#snippet rail()}
      {#if p.notes !== null}
        <OfficerNotes
          notes={p.notes}
          {viewerTimezone}
          {locale}
          saving={p.saving}
          deletingId={p.deletingId}
          error={p.noteError}
          onAdd={(body) => p.addNote(body)}
          onDelete={(id) => void p.deleteNote(id)}
        />
      {/if}
    {/snippet}
    <Page
      layout="split"
      panels
      {header}
      rail={p.notes === null ? undefined : rail}
      testid="guild-profile-page"
    >
      {#if member === null}
        <Notice text={copy.noCharacter} tone="neutral" testid="guild-profile-no-character" />
      {/if}
      <Section
        panel
        title={copy.gearGap}
        meta={copy.gearGapMeta}
        action={member === null
          ? undefined
          : { label: copy.openCharacter, href: characterHref(member.region, member.ruleset, member.name) }}
        testid="guild-profile-gear-gap"
      >
        {#if gap === null}
          <p class="text-muted text-[14px]">{copy.noGearShared}</p>
        {:else}
          {#each gap as line (line.key)}
            <Row title={line.text} open density="slim" testid={`guild-profile-gap-${line.key}`} />
          {/each}
        {/if}
      </Section>
      <Section title={copy.attendance} panel testid="guild-profile-attendance">
        {#if attendance === null}
          <p class="text-muted text-[14px]">{copy.noNights}</p>
        {:else}
          {#if attendance.rows.length > 0}
            <Grid
              caption={copy.lastNights(attendance.columns.length)}
              columns={attendance.columns}
              rows={attendance.rows}
              testid="guild-profile-attendance-grid"
            />
          {/if}
          {#if attendance.line !== ''}
            <p class="text-text pt-3 text-[14px]" data-testid="guild-profile-attendance-line">
              {attendance.line}
            </p>
          {/if}
        {/if}
      </Section>
      <Section title={copy.loot} panel testid="guild-profile-loot">
        {#if loot.length === 0}
          <p class="text-muted text-[14px]">{copy.noLoot}</p>
        {:else}
          <ul>
            {#each loot as item (item.key)}
              <li>
                <ItemRow
                  item={{ id: item.id, name: item.name, quality: item.quality, icon: item.icon }}
                  {build}
                  iconSize={32}
                  source={item.source}
                  testid="guild-profile-loot-row"
                />
              </li>
            {/each}
          </ul>
        {/if}
      </Section>
    </Page>
  {/if}
</div>
