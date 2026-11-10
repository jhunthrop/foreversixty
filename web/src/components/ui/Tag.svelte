<!-- web/src/components/ui/Tag.svelte -->
<!-- A state tag: 22 px, 12 px 600 in sentence case, 1 px border on the control radius, five tones,
     an optional dot. For a state (Founder, Connected, Current, Main); source pills stay `pill`. -->
<script lang="ts" module>
  export type TagTone = 'gold' | 'success' | 'muted' | 'danger' | 'warn' | 'premium' | 'yours';
</script>

<script lang="ts">
  import type { Snippet } from 'svelte';

  const TONES: Readonly<Record<TagTone, string>> = {
    gold: 'text-gold border-gold/40 bg-gold/10',
    success: 'text-gain border-gain/50 bg-gain/12',
    muted: 'text-muted border-white/20 bg-white/[.05]',
    danger: 'text-wipe border-wipe/40 bg-wipe/10',
    // The site has no separate warning hue; ember is its amber (a state to look at, not a failure).
    warn: 'text-ember border-ember/40 bg-ember/12',
    // Premium: the teal edge flat with white ink (design-system.md: "Premium is the edge colour flat").
    premium: 'text-strong border-teal-edge bg-teal-edge',
    // The viewer's own row ("Yours", "You"): the selected state's accent edge and tint, white ink.
    yours: 'text-strong border-selected-mark bg-selected',
  };

  let {
    tone = 'gold',
    dot = false,
    title,
    testid,
    children,
  }: {
    tone?: TagTone;
    dot?: boolean;
    title?: string;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

<span
  class="rounded-pill inline-flex h-[22px] items-center gap-[6px] border px-2 text-[12px] font-semibold tracking-[.01em] whitespace-nowrap {TONES[
    tone
  ]}"
  {title}
  data-testid={testid}
  data-tone={tone}
>
  {#if dot}<span class="h-[7px] w-[7px] rounded-full bg-current" aria-hidden="true" data-testid="tag-dot"
    ></span>{/if}{@render children()}
</span>
