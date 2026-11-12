<!-- web/src/components/guild/report/ReportDiscord.svelte -->
<!-- The Discord summary, officers only (web layout spec 2026-11-05 §7.2 "Raid report"): what the
     post carries (the embed's fields: kills, wipes, time, came, the death count, the best damage
     and the loot; it names nobody for dying), then `Post to Discord`, or after a post its time and
     `Post again`. Without a connected channel the line points at Settings. `Copy for Discord`
     stays for a guild that pastes by hand. Officer-read-only keeps Copy live and disables Post. -->
<script lang="ts">
  import type { RaidNight, RaidReport } from '../../../lib/guild/attendance-contract';
  import { attendanceCopy, postedLine, postedTime } from '../../../lib/guild/attendance-copy';
  import { selectOnMount } from '../../../lib/guild/focus-now';
  import { raidCopy } from '../../../lib/guild/raid-copy';
  import { splitAtLink } from '../../../lib/guild/raid-form';
  import { discordPreview } from '../../../lib/guild/report-night-view';
  import { reportNightCopy as copy } from '../../../lib/guild/report-night-copy';
  import { CopyState } from '../../../lib/ui/copy-state.svelte';
  import Button from '../../ui/Button.svelte';
  import Section from '../../ui/Section.svelte';

  const COPY_FALLBACK_ROWS = 8;
  const connectParts = splitAtLink(raidCopy.form.connectDiscord, raidCopy.form.connectDiscordLink);

  let {
    title,
    report,
    night,
    connected,
    settingsHref,
    officerTools,
    posting,
    postLine,
    viewerTimezone,
    locale,
    copyText,
    onPost,
    copy: clip = new CopyState(),
  }: {
    title: string;
    report: RaidReport;
    night: RaidNight;
    /** The guild has a raid channel the bot can post to. */
    connected: boolean;
    settingsHref: string;
    /** False in officer-read-only: Post is disabled (the page shows the gate sentence). */
    officerTools: boolean;
    posting: boolean;
    postLine: string | null;
    viewerTimezone: string;
    locale: string;
    /** The plain text `Copy for Discord` puts on the clipboard. */
    copyText: () => string;
    onPost: (repost: boolean) => void;
    /** The clipboard state; a test seeds it. */
    copy?: CopyState;
  } = $props();

  $effect(() => () => clip.dispose());

  const lines = $derived(discordPreview(title, report, night));
  const posted = $derived(report.posted_at);
</script>

<Section title={copy.discord.head} anchor="raid-discord" testid="raid-discord">
  <div class="text-muted flex flex-col gap-0.5 text-[13px] leading-[18px]" data-testid="raid-discord-preview">
    {#each lines as line, index (index)}
      <p class={index === 0 ? 'text-strong font-bold' : ''}>{line}</p>
    {/each}
  </div>
  <div class="mt-3 flex flex-wrap items-center gap-3">
    {#if connected}
      <Button
        tone={posted === null ? 'primary' : 'secondary'}
        size="sm"
        disabled={!officerTools || posting}
        onclick={() => onPost(posted !== null)}
        testid="raid-report-post"
        >{posting
          ? copy.discord.posting
          : posted === null
            ? copy.discord.post
            : copy.discord.postAgain}</Button
      >
    {/if}
    <Button size="sm" onclick={() => void clip.copy(copyText())} testid="raid-report-copy"
      >{clip.state === 'copied' ? attendanceCopy.report.copied : attendanceCopy.report.copy}</Button
    >
    <span class="sr-only" role="status" aria-live="polite" data-testid="raid-report-copied"
      >{clip.state === 'copied' ? attendanceCopy.report.copied : ''}</span
    >
  </div>
  {#if posted !== null}
    <p class="text-muted mt-2 text-[13px]" data-testid="raid-report-posted">
      {postedLine(postedTime(posted, viewerTimezone, locale))}
    </p>
  {:else if !connected}
    <p class="text-muted mt-2 text-[13px]" data-testid="raid-report-connect">
      {connectParts.before}<a href={settingsHref} class="text-text underline"
        >{raidCopy.form.connectDiscordLink}</a
      >{connectParts.after}
    </p>
  {/if}
  {#if postLine !== null}
    <p class="text-wipe mt-2 text-[13px]" role="alert" data-testid="raid-report-post-error">{postLine}</p>
  {/if}
  {#if clip.state === 'failed'}
    <div class="mt-2 flex flex-col gap-2" data-testid="raid-report-copy-failed">
      <p class="text-wipe text-[13px]" role="alert">{attendanceCopy.errors.copyFailed}</p>
      <textarea
        class="rounded-field text-text bg-bg w-full border p-3 font-mono text-[12px]"
        rows={COPY_FALLBACK_ROWS}
        readonly
        aria-label={attendanceCopy.report.copy}
        use:selectOnMount>{clip.text}</textarea
      >
    </div>
  {/if}
</Section>
