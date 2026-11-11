<!-- web/src/components/ui/CharacterName.svelte -->
<!-- A character named the way the game names one (component audit 2026-11-10 cluster 18; spec
     2026-11-05 §3.3): an optional ringed crest (ui/Crest), then the name in its class colour at 600,
     a link when `href` is set. Every name beside a crest on the web is this component; the caller
     sets only layout and size through `class`. crestSize 0 (the default) draws the name alone. -->
<script lang="ts">
  import { classColorVar } from '../../lib/report/format';
  import Crest from './Crest.svelte';

  let {
    name,
    cls,
    crestSize = 0,
    href,
    class: extra = '',
    testid,
    crestTestid,
  }: {
    name: string;
    /** A class name or slug; unknown or absent, the body colour and the LX crest. */
    cls: string | null | undefined;
    crestSize?: number;
    href?: string;
    /** The name's own size, weight override or truncation. */
    class?: string;
    testid?: string;
    crestTestid?: string;
  } = $props();

  const colour = $derived(classColorVar(cls ?? undefined));
</script>

{#if crestSize > 0}<Crest {cls} size={crestSize} testid={crestTestid} />{/if}{#if href !== undefined}<a
    {href}
    class="ui-cname font-semibold hover:underline {extra}"
    style:color={colour}
    data-testid={testid}>{name}</a
  >{:else}<span class="ui-cname font-semibold {extra}" style:color={colour} data-testid={testid}>{name}</span
  >{/if}
