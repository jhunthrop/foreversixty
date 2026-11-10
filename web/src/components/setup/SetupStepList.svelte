<!-- web/src/components/setup/SetupStepList.svelte -->
<!-- /setup's three steps (setup spec 2026-11-07 sections 4 and 5) drawn from lib/setup/steps.ts's
     rows: each row's mark, its lines (the done line in --color-text, a warning in gold, a failure in
     the wipe colour with role="alert", a pairing code as Code, a reserved Skeleton line while the
     reads are in flight, and on a phone the second line an answer wraps to there) and its one action. A link action is the system Button with its href, a
     button action calls back (Try again, Refresh from Battle.net, Copy link), a text link is gold
     with `›`, the Download step is the visitor's own build (CompanionDownload), and the addon's
     install is CurseForge then Wago as one group. -->
<script lang="ts">
  import { characterListCopy } from '../../lib/account/character-list-copy';
  import { addonCopy } from '../../lib/addon/copy';
  import type { CompanionRelease } from '../../lib/setup/release';
  import type { VisitorPlatform } from '../../lib/setup/visitor-platform';
  import { CURSEFORGE_URL, WAGO_URL, stepHref } from '../../lib/setup/links';
  import type { StepAction, StepLineOut, StepRow } from '../../lib/setup/steps';
  import { LINK } from '../../lib/ui/primitives';
  import Button from '../ui/Button.svelte';
  import CompanionDownload from './CompanionDownload.svelte';
  import Code from '../ui/Code.svelte';
  import Steps from '../ui/Steps.svelte';

  const SKELETON_WIDTH_PX = 240;

  let {
    rows,
    platform,
    release = null,
    refreshing = false,
    copied = false,
    onAction,
    testid,
  }: {
    rows: readonly StepRow[];
    /** The visitor's system: which build the Download step offers (lib/setup/downloads.ts). */
    platform: VisitorPlatform;
    /** The feed's newest companion release, for the download's link, version and size. */
    release?: CompanionRelease | null;
    /** Refresh from Battle.net has started: its button reads `refreshing` and is disabled. */
    refreshing?: boolean;
    /** Copy link has just copied: its button says so. */
    copied?: boolean;
    /** A button action was pressed (retry, refresh). */
    onAction?: (action: StepAction) => void;
    testid?: string;
  } = $props();

  const items = $derived(
    rows.map((row) => ({ key: `step-${row.n}`, n: row.n, mark: row.mark, title: row.title })),
  );
  const byKey = $derived(new Map(rows.map((row) => [`step-${row.n}`, row])));

  const TONE: Record<StepLineOut['tone'], string> = {
    muted: 'text-muted',
    text: 'text-text',
    ok: 'text-text',
    warn: 'text-gold',
    miss: 'text-wipe',
  };
</script>

{#snippet sub(item: { key: string })}
  {#each byKey.get(item.key)?.lines ?? [] as line, index (index)}
    {#if line.skeleton === true}
      <span class={['flex h-5 items-center', line.phoneOnly === true && 'md:hidden']}
        ><span
          class="skeleton-block h-[14px] rounded-[3px]"
          style:width="{SKELETON_WIDTH_PX}px"
          aria-hidden="true"
        ></span></span
      >
    {:else}
      <span class="block">
        <span class={TONE[line.tone]} role={line.alert === true ? 'alert' : undefined}>{line.text}</span>
        {#if line.code !== undefined}<Code text={line.code} testid="setup-pairing-code" />{/if}
      </span>
    {/if}
  {/each}
{/snippet}

{#snippet action(item: { key: string })}
  {@const step = byKey.get(item.key)}
  {#if step?.action}
    {@const act = step.action}
    {@const href = stepHref(act.id)}
    {#if act.id === 'download'}
      <CompanionDownload {platform} {release} primary={act.primary} />
    {:else if act.id === 'install'}
      <div class="flex flex-wrap items-center gap-3" role="group" aria-label={act.label}>
        <Button
          tone={act.primary ? 'primary' : 'secondary'}
          href={CURSEFORGE_URL}
          rel="noopener"
          testid="setup-action-curseforge">{addonCopy.installCurseForge}</Button
        >
        <Button href={WAGO_URL} rel="noopener" testid="setup-action-wago">{addonCopy.installWago}</Button>
      </div>
    {:else if act.id === 'update'}
      <!-- A ticked step's update: quiet gold links, never a button that competes with the next step. -->
      <div class="flex flex-wrap items-center gap-x-5 gap-y-1" role="group" aria-label={act.label}>
        <a
          class="{LINK} inline-flex min-h-11 items-center md:min-h-0"
          href={CURSEFORGE_URL}
          rel="noopener"
          data-testid="setup-action-curseforge">{addonCopy.setupUpdateCurseForge} ›</a
        >
        <a
          class="{LINK} inline-flex min-h-11 items-center md:min-h-0"
          href={WAGO_URL}
          rel="noopener"
          data-testid="setup-action-wago">{addonCopy.setupUpdateWago} ›</a
        >
      </div>
    {:else if act.kind === 'link'}
      <a
        class="{LINK} inline-flex min-h-11 items-center md:min-h-0"
        href={href ?? '#'}
        data-testid="setup-action-{act.id}">{act.label} ›</a
      >
    {:else if act.kind === 'link-button' && href !== null}
      <Button
        tone={act.primary ? 'primary' : 'secondary'}
        {href}
        rel="noopener"
        testid="setup-action-{act.id}">{act.label}</Button
      >
    {:else}
      {@const busy = act.id === 'refresh' && refreshing}
      {@const done = act.id === 'copylink' && copied}
      <Button
        tone={act.primary ? 'primary' : 'secondary'}
        disabled={busy}
        onclick={() => onAction?.(act)}
        testid="setup-action-{act.id}"
        >{busy ? characterListCopy.refreshing : done ? addonCopy.setupLinkCopied : act.label}</Button
      >
    {/if}
  {/if}
{/snippet}

<Steps {items} {sub} {action} {testid} />
