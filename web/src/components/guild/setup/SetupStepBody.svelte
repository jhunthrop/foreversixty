<!-- web/src/components/guild/setup/SetupStepBody.svelte -->
<!-- One setup step's lines under its title (guild onboarding spec §3): its live facts, its parts,
     and, while it is to do, the controls that finish it in place. A done or skipped step shows
     its facts only. Every write goes through `run`, which the page holds (busy, the error line). -->
<script lang="ts">
  import type { GuildSettingsData, GuildVisibility } from '../../../lib/guild/api';
  import { updateGuildSettings } from '../../../lib/guild/api';
  import { guildBotInstallUrl, saveBotChannel } from '../../../lib/guild/discord-bot-api';
  import { inviteUrl } from '../../../lib/guild/invite-link';
  import { VISIBILITY_OPTIONS } from '../../../lib/guild/settings-view';
  import { saveRecruitChannel, type GuildSetup, type SetupStep } from '../../../lib/guild/setup-api';
  import { setupCopy } from '../../../lib/guild/setup-copy';
  import {
    discordLine,
    inviteLine,
    officerRank,
    recruitingLine,
    recruitingParts,
    scheduleLine,
    scheduleParts,
    syncLine,
    syncShare,
    visibilityLine,
  } from '../../../lib/guild/setup-model';
  import { SITE_URL } from '../../../lib/seo';
  import Button from '../../ui/Button.svelte';
  import Meter from '../../ui/Meter.svelte';
  import Segmented from '../../ui/Segmented.svelte';
  import GuildInviteLink from '../GuildInviteLink.svelte';
  import SetupChannelPick from './SetupChannelPick.svelte';
  import SetupParts from './SetupParts.svelte';
  import SetupPaste from './SetupPaste.svelte';

  let {
    step,
    setup,
    settings,
    guildId,
    guildName,
    guildPath,
    primary,
    busy,
    when,
    run,
    onConfirm,
    onSettings,
  }: {
    step: SetupStep;
    setup: GuildSetup;
    settings: GuildSettingsData;
    guildId: number;
    guildName: string;
    guildPath: string;
    /** This step is the first to do: its main control is the page's one primary. */
    primary: boolean;
    busy: boolean;
    /** The next raid's start in words (`Tue 8 Dec, 8:00 pm`). */
    when: (iso: string) => string;
    /** Runs a write and reads the checklist again. */
    run: (action: () => Promise<void>) => void;
    /** Marks a confirmable step done. */
    onConfirm: () => void;
    onSettings: (fresh: GuildSettingsData) => void;
  } = $props();

  const todo = $derived(step.state === 'todo');
  const tone = $derived(primary ? 'primary' : 'secondary');
  const testid = $derived(`setup-${step.key.replaceAll('_', '-')}`);

  let rankDraft = $state<number | null>(null);
  const rankValue = $derived(rankDraft ?? setup.ranks.officer_max_rank_index);
  const rankChanged = $derived(rankValue !== setup.ranks.officer_max_rank_index && rankValue >= 0);

  const invitePath = $derived(settings.invite.url);
  const inviteText = $derived(
    invitePath === null ? null : setupCopy.invite.paste(guildName, inviteUrl(invitePath)),
  );
  const visibility = $derived(settings.default_visibility);
  const visibilityLabel = $derived(VISIBILITY_OPTIONS.find((o) => o.value === visibility)?.label ?? '');

  const saveSettings = (patch: Parameters<typeof updateGuildSettings>[1]): void =>
    run(async () => onSettings(await updateGuildSettings(guildId, patch)));
</script>

<div class="flex flex-col gap-2 pt-1" data-testid={`${testid}-body`}>
  {#if step.key === 'officer_ranks'}
    {#if setup.ranks.ranks.length > 0}
      <p class="flex flex-wrap gap-x-3 text-[13px]" data-testid="setup-ranks">
        {#each setup.ranks.ranks as line (line.rank_index)}
          <span
            class={officerRank(line.rank_index, setup.ranks.officer_max_rank_index)
              ? 'text-strong font-semibold'
              : 'text-muted'}
            title={line.names.join(', ')}>{setupCopy.ranks.line(line.rank_index, line.members)}</span
          >
        {/each}
      </p>
    {:else}
      <p class="text-muted text-[13px]">{setupCopy.ranks.noRanks}</p>
    {/if}
    <p class="text-muted text-[13px]">
      {setupCopy.ranks.officers(setup.ranks.officer_max_rank_index, setup.ranks.officers)}
    </p>
    {#if todo}
      <div class="flex flex-wrap items-center gap-3 pt-1">
        <label class="flex items-center gap-2 text-[14px]">
          {setupCopy.ranks.inputLabel}
          <input
            type="number"
            min="0"
            max="9"
            class="rounded-control text-strong h-11 w-[64px] border border-white/15 bg-white/[.045] px-3 text-[15px] font-semibold"
            value={rankValue}
            oninput={(e) => (rankDraft = Number((e.currentTarget as HTMLInputElement).value))}
            disabled={busy}
            data-testid="setup-ranks-input"
          />
        </label>
        <Button
          size="sm"
          tone={rankChanged ? tone : 'secondary'}
          disabled={busy || !rankChanged}
          onclick={() => saveSettings({ officer_max_rank_index: rankValue })}
          testid="setup-ranks-save">{setupCopy.ranks.save}</Button
        >
        <Button size="sm" tone="ghost" disabled={busy} onclick={onConfirm} testid="setup-ranks-keep"
          >{setupCopy.ranks.keep(setup.ranks.officer_max_rank_index)}</Button
        >
      </div>
    {/if}
  {:else if step.key === 'discord_bot'}
    <p class="text-muted text-[13px]">{discordLine(setup.discord)}</p>
    {#if todo && setup.discord.installed && setup.discord.channel_name === null}
      <SetupChannelPick
        {guildId}
        label={setupCopy.discord.channelLabel}
        {busy}
        onSave={(id) => run(async () => void (await saveBotChannel(guildId, id)))}
        testid="setup-discord-channel"
      />
    {:else if todo}
      <div class="flex flex-col items-start gap-1.5 pt-1">
        <Button
          {tone}
          href={`${guildBotInstallUrl(guildId)}?return=setup`}
          rel="noopener"
          testid="setup-discord-add">{setupCopy.discord.add}</Button
        >
        <span class="text-muted text-[13px]">{setupCopy.discord.free}</span>
      </div>
    {/if}
  {:else if step.key === 'recruiting'}
    <p class="text-muted text-[13px]">{recruitingLine(setup.recruiting)}</p>
    <SetupParts parts={recruitingParts(setup)} testid="setup-recruiting-parts" />
    {#if todo}
      <div class="flex flex-wrap items-center gap-3 pt-1">
        {#if setup.discord.installed}
          <SetupChannelPick
            {guildId}
            label={setupCopy.recruiting.channelLabel}
            current={setup.discord.recruit_channel_name}
            {busy}
            onSave={(id) => run(async () => void (await saveRecruitChannel(guildId, id)))}
            testid="setup-recruit-channel"
          />
        {:else}
          <span class="text-muted text-[13px]">{setupCopy.recruiting.botFirst}</span>
        {/if}
        <a class="text-gold hover:text-gold-hover text-[14px] font-semibold" href={`${guildPath}/recruit`}
          >{setupCopy.recruiting.editPitch} ›</a
        >
      </div>
    {:else if setup.discord.recruit_channel_name !== null}
      <p class="text-muted text-[13px]">
        {setupCopy.recruiting.channelIn(setup.discord.recruit_channel_name)}
      </p>
    {/if}
  {:else if step.key === 'raid_schedule'}
    <p class="text-muted text-[13px]">{scheduleLine(setup.schedule, when)}</p>
    <SetupParts parts={scheduleParts(setup)} testid="setup-schedule-parts" />
    {#if todo}
      <div class="pt-1">
        <Button size="sm" {tone} href={`${guildPath}#raids`} testid="setup-schedule-plan"
          >{setupCopy.schedule.plan}</Button
        >
      </div>
    {/if}
  {:else if step.key === 'invite_link'}
    <p class="text-muted text-[13px]">{inviteLine(setup.invite)}</p>
    {#if todo}
      <GuildInviteLink
        {guildId}
        rotatedAt={settings.invite.rotated_at}
        keptUrl={settings.invite.url}
        madeBy={settings.invite.made_by}
        disabled={busy}
        onRotated={(invite) => onSettings({ ...settings, invite })}
      />
      {#if inviteText !== null}
        <SetupPaste label={setupCopy.invite.pasteLabel} text={inviteText} testid="setup-invite-paste" />
        <div>
          <Button size="sm" {tone} disabled={busy} onclick={onConfirm} testid="setup-invite-posted"
            >{setupCopy.invite.posted}</Button
          >
        </div>
      {/if}
    {/if}
  {:else if step.key === 'addon_sync'}
    <p class="text-muted text-[13px]">{syncLine(setup.sync)}</p>
    <div class="max-w-[320px]">
      <Meter value={syncShare(setup.sync)} tone="gold" testid="setup-sync-meter" />
    </div>
    {#if todo}
      <SetupPaste
        label={setupCopy.sync.pasteLabel}
        text={setupCopy.sync.paste(`${SITE_URL}/setup`)}
        testid="setup-sync-paste"
      />
    {/if}
  {:else if step.key === 'report_visibility'}
    <p class="text-muted text-[13px]">{visibilityLine(setup.visibility)}</p>
    {#if todo}
      <div class="flex flex-wrap items-center gap-3 pt-1">
        <Segmented
          options={VISIBILITY_OPTIONS}
          value={visibility}
          label={setupCopy.visibility.label}
          size="sm"
          disabled={busy}
          onChange={(value: GuildVisibility) => saveSettings({ default_visibility: value })}
          testid="setup-visibility"
        />
        <Button size="sm" tone="ghost" disabled={busy} onclick={onConfirm} testid="setup-visibility-keep"
          >{setupCopy.visibility.keep(visibilityLabel)}</Button
        >
      </div>
    {/if}
  {:else if step.key === 'linked_role'}
    <p class="text-muted text-[13px]">{setupCopy.linkedRole.line}</p>
    {#if todo}
      <div class="pt-1">
        <Button size="sm" {tone} disabled={busy} onclick={onConfirm} testid="setup-linked-role-done"
          >{setupCopy.linkedRole.done}</Button
        >
      </div>
    {/if}
  {/if}
</div>
