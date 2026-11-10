<!-- web/src/components/guild/setup/SetupParts.svelte -->
<!-- A step's parts in one wrapping line (guild onboarding spec §3): a small ring, ticked teal when
     the part is set, the part's name, and a plan Tag on a part a plan carries. -->
<script lang="ts">
  import type { SetupPart } from '../../../lib/guild/setup-model';
  import Tag from '../../ui/Tag.svelte';

  let { parts, testid }: { parts: readonly SetupPart[]; testid: string } = $props();
</script>

<ul class="flex flex-wrap gap-x-4 gap-y-1 text-[13px]" data-testid={testid}>
  {#each parts as part (part.key)}
    <li class="inline-flex items-center gap-1.5" data-done={part.done}>
      <span
        class="inline-flex size-[14px] flex-none items-center justify-center rounded-full text-[10px] leading-none font-bold {part.done
          ? 'bg-kill text-bg'
          : 'border-gold border'}"
        role="img"
        aria-label={part.done ? 'Set' : 'Not set'}>{part.done ? '✓' : ''}</span
      >
      <span class={part.done ? 'text-text' : 'text-muted'}>{part.label}</span>
      {#if part.plan !== undefined}<Tag tone="muted">{part.plan}</Tag>{/if}
    </li>
  {/each}
</ul>
