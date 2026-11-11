<!-- web/src/components/guild/recruit/GuildMemberProfile.svelte -->
<!-- /guild/{region}/{ruleset}/{slug}/members/{account} (web layout spec 2026-11-05 §7.2 Member
     profile, "What do I know about this raider?"), an officer's read of anyone or a member's of
     themselves; mock web-guild-profile. The head names the raider in class colour with their figures
     over the guild's tabs (Roster current). Main: Raid ready, Next upgrades (ProfileUpgrades), (the gear gap in the BiS page's unit,
     enchants, bags, talents, item level against the guild, the top fight, the last export), the last 4
     raids as a Grid with the 30-raid figure, and loot this tier. The rail: their characters, then,
     officers only, Officer notes: a member's own read carries no notes and draws no trace of them.
     GuildShell loads this file as its own chunk. -->
<script lang="ts">
  import activeBuild from '../../../data/active-build.json';
  import { characterHref, guildHref, type CharacterPath } from '../../../lib/characters';
  import { liveProfileDeps } from '../../../lib/guild/guild-view-deps';
  import { ProfilePage } from '../../../lib/guild/profile-page.svelte';
  import {
    attendanceView,
    characterRows,
    lootFairness,
    lootRows,
    profileFigures,
    profileHead,
    readyFacts,
    upgradeRows,
    readyTag,
  } from '../../../lib/guild/profile-view';

  import { profileCopy as copy } from '../../../lib/guild/recruit-copy';
  import { classTone } from '../../../lib/ui/primitives';
  import { untrack } from 'svelte';
  import { viewerZone } from '../../../lib/dates';
  import Crest from '../../ui/Crest.svelte';
  import SignInPrompt from '../../SignInPrompt.svelte';
  import Facts from '../../ui/Facts.svelte';
  import Figure from '../../ui/Figure.svelte';
  import Grid from '../../ui/Grid.svelte';
  import ItemRow from '../../ui/ItemRow.svelte';
  import LoadError from '../../ui/LoadError.svelte';
  import Notice from '../../ui/Notice.svelte';
  import Page from '../../ui/Page.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';
  import Skeleton from '../../ui/Skeleton.svelte';
  import Tag from '../../ui/Tag.svelte';
  import GuildViewHeader from './GuildViewHeader.svelte';
  import OfficerNotes from './OfficerNotes.svelte';
  import ProfileUpgrades from './ProfileUpgrades.svelte';

  let {
    path,
    account,
    page: given,
    viewerTimezone = viewerZone(),
  }: {
    path: CharacterPath;
    /** `me` or the roster's account key (`u:42`). */
    account: string;
    page?: ProfilePage;
    viewerTimezone?: string;
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
    {@const ready = readyFacts(profile, viewerTimezone)}
    {@const readyState = readyTag(profile)}
    {@const characters = characterRows(profile)}
    {@const attendance = attendanceView(profile, viewerTimezone)}
    {@const loot = lootRows(profile, viewerTimezone)}
    {@const fairness = lootFairness(profile)}
    {@const upgrades = upgradeRows(profile)}
    {@const member = profile.member}
    {#snippet crest()}<Crest cls={head.classSlug} size={56} />{/snippet}
    {#snippet figures()}
      {@const list = profileFigures(profile)}
      <!-- Two by two on a phone, one row from 768, so no label wraps under its figure. -->
      <div
        class="grid grid-cols-2 gap-x-6 gap-y-5 md:grid-cols-[repeat(var(--figures),minmax(0,1fr))]"
        style:--figures={list.length}
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
        tabRole={p.notes === null ? 'member' : 'officer'}
        testid="guild-profile-head"
      />
    {/snippet}
    {#snippet rail()}
      {#if characters.length > 0}
        <Section title={copy.characters} panel testid="guild-profile-characters">
          <ul>
            {#each characters as character (character.key)}
              <li>
                <Row
                  title={character.name}
                  titleTone={classTone(character.classSlug)}
                  href={character.href ?? undefined}
                  density="slim"
                  open
                  testid="guild-profile-character"
                >
                  {#snippet lead()}<Crest cls={character.classSlug} size={28} />{/snippet}
                  {#snippet trail()}{#if character.main}<Tag tone="muted">{copy.main}</Tag>{/if}{/snippet}
                </Row>
              </li>
            {/each}
          </ul>
        </Section>
      {/if}
      {#if p.notes !== null}
        <OfficerNotes
          notes={p.notes}
          {viewerTimezone}
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
      rail={p.notes === null && characters.length === 0 ? undefined : rail}
      testid="guild-profile-page"
    >
      {#if member === null}
        <Notice text={copy.noCharacter} tone="neutral" testid="guild-profile-no-character" />
      {/if}
      <Section
        panel
        title={copy.ready.head}
        action={member === null
          ? undefined
          : { label: copy.openCharacter, href: characterHref(member.region, member.ruleset, member.name) }}
        testid="guild-profile-ready"
      >
        {#snippet head()}{#if readyState !== null}<Tag tone={readyState.tone} testid="guild-profile-ready-tag"
              >{readyState.text}</Tag
            >{/if}{/snippet}
        {#if ready === null}
          <p class="text-muted text-[14px]">{copy.ready.noData}</p>
        {:else}
          <Facts items={ready} testid="guild-profile-ready-facts" />
        {/if}
      </Section>
      {#if upgrades.length > 0}
        <ProfileUpgrades rows={upgrades} classSlug={head.classSlug} />
      {/if}
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
      <Section
        title={copy.loot}
        meta={fairness.median === '' ? undefined : fairness.median}
        panel
        testid="guild-profile-loot"
      >
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
        {#if fairness.since !== null}
          <p class="text-text pt-3 text-[14px]" data-testid="guild-profile-loot-since">{fairness.since}</p>
        {/if}
      </Section>
    </Page>
  {/if}
</div>
