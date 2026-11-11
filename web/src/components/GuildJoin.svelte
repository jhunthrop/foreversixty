<!-- web/src/components/GuildJoin.svelte -->
<!-- /guild/invite/<token> (spec 2026-11-05 §3.4; web layout spec 2026-11-05 §3A): the guild's PageHeader, then one Section
     on the form measure, no card. Which guild, which of the
     viewer's characters joins, one `Join`. The token is looked up (GET /v1/guild-invites/{token},
     public: it names the guild without joining). Signed out, the card carries the sign-in button;
     an unknown or rotated token says so plainly with no button; a member of the guild already is
     told so. Joining calls POST /v1/guilds/invite/{token}/accept and shows the joined state: the
     guild's plan line, then `Open the guild`. The inviter is not named until the link records who
     made it (§4.1, held). -->
<script lang="ts">
  import { battlenetStartUrl, fetchMeOnce, forgetSession, type Me } from '../lib/account/api';
  import { guildHref, guildInviteHref, rulesetLabel } from '../lib/characters';
  import {
    GuildApiError,
    acceptInvite,
    peekInvite,
    type InviteAcceptResult,
    type InvitePeekResult,
  } from '../lib/guild/api';
  import { guildJoinCopy } from '../lib/guild/copy';
  import { joinedPlanLineFor } from '../lib/guild/join-plan';
  import { holdsErrorRoom, joiningRow, previewAfterJoin, showsInvitedLine } from '../lib/guild/join-view';
  import { GUILD_LOADING } from '../lib/guild/layout';
  import { guildMarkSrc } from '../lib/guild/mark';
  import { classTone } from '../lib/ui/primitives';
  import Crest from './ui/Crest.svelte';
  import FactionCrest from './character/FactionCrest.svelte';
  import GuildJoinedBlock from './guild/GuildJoinedBlock.svelte';
  import GuildJoinInvited from './guild/GuildJoinInvited.svelte';
  import GuildStatus from './GuildStatus.svelte';
  import EmptyState from './ui/EmptyState.svelte';
  import Button from './ui/Button.svelte';
  import Page from './ui/Page.svelte';
  import PageHeader from './ui/PageHeader.svelte';
  import Row from './ui/Row.svelte';
  import Section from './ui/Section.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  let { token }: { token: string } = $props();

  let me = $state<Me | null>(null);
  let peek = $state<InvitePeekResult | null>(null);
  let status = $state<'loading' | 'ready' | 'invalid' | 'failed'>('loading');
  let joined = $state<InviteAcceptResult | null>(null);
  let error = $state('');
  let busy = $state(false);
  let planLine = $state<string | null>(null);

  async function load(): Promise<void> {
    status = 'loading';
    try {
      // The lookup answers a 404 for a bad token: that is a state, not a failure.
      const [session, result] = await Promise.all([
        fetchMeOnce(),
        peekInvite(token).catch((thrown: unknown) => {
          if (thrown instanceof GuildApiError && thrown.status === 404) return null;
          throw thrown;
        }),
      ]);
      me = session;
      peek = result;
      status = result === null ? 'invalid' : 'ready';
    } catch {
      status = 'failed';
    }
  }

  $effect(() => {
    void load();
  });

  const guild = $derived(peek?.guild ?? null);
  const signedIn = $derived(me !== null);
  const alreadyIn = $derived(guild !== null && (me?.guilds.some((g) => g.id === guild.id) ?? false));
  const realm = $derived(
    guild === null ? '' : `${rulesetLabel(guild.ruleset)} ${guild.region.toUpperCase()}`,
  );
  const joining = $derived(guild === null ? null : joiningRow(peek?.joins_as, guild.name));
  /** A member already sees row 3 only for a character the guild's roster names. */
  const showJoining = $derived(
    joined === null && signedIn && (!alreadyIn || peek?.joins_as?.on_roster === true),
  );

  /** The guild's plan sentence from a fresh `/v1/me`, for an officer or the guild master; null for a
   *  member or when that read fails (the line is left out). */
  async function readPlanLine(guildId: number, guildName: string): Promise<string | null> {
    forgetSession();
    try {
      const fresh = await fetchMeOnce();
      const entry = fresh?.guilds.find((g) => g.id === guildId);
      return entry === undefined ? null : joinedPlanLineFor(guildName, entry);
    } catch {
      return null;
    }
  }

  async function onJoin(): Promise<void> {
    busy = true;
    error = '';
    try {
      joined = await acceptInvite(token);
      busy = false;
      const before = peek;
      const [line, after] = await Promise.all([
        readPlanLine(joined.guild.id, joined.guild.name),
        before === null ? Promise.resolve(null) : previewAfterJoin(before, () => peekInvite(token)),
      ]);
      planLine = line;
      if (after !== null) peek = after;
    } catch (thrown) {
      error = thrown instanceof GuildApiError ? thrown.message : guildJoinCopy.failed;
      busy = false;
    }
  }
</script>

<div class="reveal" data-testid="guild-join">
  {#if status === 'loading'}
    <div class="min-h-[380px] md:min-h-[320px]" data-testid="guild-join-loading">
      <Skeleton lines={6} testid="guild-join-skeleton" />
    </div>
  {:else if status === 'failed'}
    <GuildStatus
      status="failed"
      error={guildJoinCopy.failed}
      onRetry={() => void load()}
      lines={GUILD_LOADING.join.lines}
      minHeight={GUILD_LOADING.join.minHeight}
      testid="guild-join"
    />
  {:else if status === 'invalid' || guild === null}
    <EmptyState message={guildJoinCopy.invalid} testid="guild-join-invalid" />
  {:else}
    {#snippet header()}
      <PageHeader title={guild.name} tone="strong" testid="guild-join-head">
        {#snippet crest()}
          <FactionCrest faction={guild.faction} src={guildMarkSrc(guild)} testid="guild-join-crest" />
        {/snippet}
        {#snippet meta()}
          <span data-testid="guild-join-realm">{guildJoinCopy.guildLine(realm, peek?.members)}</span>
        {/snippet}
      </PageHeader>
    {/snippet}
    <Page layout="form" {header} testid="guild-join-card">
      <Section title={guildJoinCopy.sectionHeading} panel frame="framed" testid="guild-join-section">
        <div class="flex flex-col items-start gap-4">
          {#if showsInvitedLine(alreadyIn, joined !== null)}
            <GuildJoinInvited guildName={guild.name} inviter={peek?.inviter ?? null} />
          {/if}
          {#if showJoining && joining !== null}
            <div class="w-full" data-testid="guild-join-as">
              <Row
                title={joining.label}
                titleTone={joining.labelClass === null ? undefined : classTone(joining.labelClass)}
                sub={joining.sub}
                open
              >
                {#snippet lead()}<Crest cls={joining.crestClass} size={32} />{/snippet}
              </Row>
            </div>
          {/if}
          {#if joined !== null}
            <GuildJoinedBlock
              {planLine}
              openHref={guildHref(joined.guild.region, joined.guild.ruleset, joined.guild.name)}
            />
          {:else if !signedIn}
            <p class="text-text text-[14px]" data-testid="guild-join-signin-line">
              {guildJoinCopy.signInLine}
            </p>
            <Button
              tone="primary"
              href={battlenetStartUrl(guildInviteHref(token))}
              testid="guild-join-signin"
            >
              {guildJoinCopy.signInButton}
            </Button>
          {:else if alreadyIn}
            <p class="text-text text-[14px]" data-testid="guild-join-already">
              {guildJoinCopy.alreadyIn(guild.name)}
            </p>
            <Button
              tone="primary"
              href={guildHref(guild.region, guild.ruleset, guild.name)}
              testid="guild-join-open"
            >
              {guildJoinCopy.openGuild}
            </Button>
          {:else}
            <div class="flex flex-col gap-1">
              <p class="text-muted text-[13px]" data-testid="guild-join-verified">
                {guildJoinCopy.verifiedAtOnce}
              </p>
              <p class="text-muted text-[13px]" data-testid="guild-join-consequence">
                {guildJoinCopy.consequence}
              </p>
            </div>
            <div class="flex w-full flex-col gap-3 md:flex-row md:items-center">
              <Button
                tone="primary"
                class="w-full md:w-fit"
                onclick={() => void onJoin()}
                {busy}
                busyLabel={guildJoinCopy.joining}
                testid="guild-join-button"
              >
                {guildJoinCopy.joinButton(guild.name)}
              </Button>
              <Button
                tone="ghost"
                href={guildHref(guild.region, guild.ruleset, guild.name)}
                testid="guild-join-not-now"
              >
                {guildJoinCopy.notNow}
              </Button>
            </div>
            <!-- The error line's room is held only while a join is in flight (or failed), so the
                 section at rest ends at its buttons (guild round 2). -->
            {#if holdsErrorRoom(busy, error)}
              <div class="min-h-[21px]" data-testid="guild-join-error-slot">
                {#if error !== ''}<p
                    class="text-wipe text-[14px]"
                    role="alert"
                    data-testid="guild-join-action-error"
                  >
                    {error}
                  </p>{/if}
              </div>
            {/if}
          {/if}
        </div>
      </Section>
    </Page>
  {/if}
</div>
