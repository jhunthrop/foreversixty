<!-- web/src/components/setup/SetupStepList.svelte -->
<!-- /setup's three steps (setup spec 2026-11-07 sections 4 and 5) drawn from lib/setup/steps.ts's
     rows: each row's mark, its lines (the done line in --color-text, a warning in gold, a failure in
     the wipe colour with role="alert", a pairing code as Code, a reserved Skeleton line while the
     reads are in flight) and its one action. A link action is the system Button with its href, a
     button action calls back (Try again, Refresh from Battle.net), a text link is gold with `›`. -->
<script lang="ts">
  import { characterListCopy } from '../../lib/account/character-list-copy';
  import { stepHref } from '../../lib/setup/links';
  import type { StepAction, StepLineOut, StepRow } from '../../lib/setup/steps';
  import { LINK } from '../../lib/ui/primitives';
  import Button from '../ui/Button.svelte';
  import Code from '../ui/Code.svelte';
  import Steps from '../ui/Steps.svelte';

  const SKELETON_WIDTH_PX = 240;

  let {
    rows,
    companionAsset,
    refreshing = false,
    onAction,
    testid,
  }: {
    rows: readonly StepRow[];
    /** companion.json's first download: the file the Download step links to. */
    companionAsset: string;
    /** Refresh from Battle.net has started: its button reads `refreshing` and is disabled. */
    refreshing?: boolean;
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
      <span class="flex h-5 items-center"
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
    {@const href = stepHref(act.id, companionAsset)}
    {#if act.kind === 'link'}
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
      <Button
        tone={act.primary ? 'primary' : 'secondary'}
        disabled={busy}
        onclick={() => onAction?.(act)}
        testid="setup-action-{act.id}">{busy ? characterListCopy.refreshing : act.label}</Button
      >
    {/if}
  {/if}
{/snippet}

<Steps {items} {sub} {action} {testid} />
