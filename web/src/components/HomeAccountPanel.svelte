<!-- web/src/components/HomeAccountPanel.svelte -->
<!-- The home page's signed-in swap (spec 2026-09-22 §3.2): while this is loading or the
     visitor is signed out, it renders nothing and the Astro shell's own server-rendered
     sentence + button (index.astro) stays exactly where it is -- the same "server shell
     first, island swaps in place" trick SessionNav.svelte's header link uses, so the panel
     never shows two competing versions. Both this island's root and index.astro's
     signed-out block share `[grid-area:1/1]` in a shared grid wrapper, so once this mounts
     signed-in it visually occludes the signed-out row (a solid background over the same
     cell) instead of the two stacking and reflowing the page underneath. -->
<script lang="ts">
  import { fetchMeOnce, type Me } from '../lib/account/api';
  import { readCurrent } from '../lib/current-character';
  import { classColorVar } from '../lib/report/format';
  import { homePanelCopy } from '../lib/home-panel-copy';

  let me = $state<Me | null>(null);
  let ready = $state(false);

  $effect(() => {
    void fetchMeOnce()
      .then((result) => {
        me = result;
        ready = true;
      })
      .catch(() => {
        ready = true;
      });
  });

  const pointer = $derived(readCurrentIfReady());
  function readCurrentIfReady() {
    if (!ready || me === null) return null;
    return readCurrent();
  }
  const current = $derived(
    pointer === null ? null : me?.characters.find((c) => c.key === pointer.ref || pointer.ref === ''),
  );
  const displayName = $derived(pointer?.label ?? me?.characters[0]?.name ?? '');
  const colour = $derived(classColorVar(current?.class));
</script>

{#if ready && me !== null}
  <div
    class="flex flex-wrap items-center gap-3 bg-[var(--color-bg)] [grid-area:1/1]"
    data-testid="home-account-panel"
  >
    <span class="rounded-control h-7 w-7 shrink-0" style={`background: ${colour}`} aria-hidden="true"></span>
    <span class="text-[15px] font-semibold" style={`color: ${colour}`}>{displayName}</span>
    <a class="text-nav text-[13px] font-semibold" href="/planner">{homePanelCopy.openInPlanner}</a>
    <a class="text-nav text-[13px] font-semibold" href="/sim">{homePanelCopy.openInSimulator}</a>
    <a class="text-nav text-[13px] font-semibold" href="/logs">{homePanelCopy.logs}</a>
    <a class="text-nav text-[13px] font-semibold" href="/account">{homePanelCopy.yourCharacters}</a>
  </div>
{/if}
