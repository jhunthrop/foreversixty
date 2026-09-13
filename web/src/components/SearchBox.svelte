<!-- web/src/components/SearchBox.svelte -->
<script lang="ts">
  type Result = { url: string; title: string; excerpt: string };
  type Pagefind = { init: () => Promise<void>; search: (q: string) => Promise<{ results: Array<{ data: () => Promise<{ url: string; meta: { title: string }; excerpt: string }> }> }> };

  let { initial = '' }: { initial?: string } = $props();
  let query = $state(initial);
  let results = $state<Result[]>([]);
  let active = $state(-1);
  let open = $state(false);
  let input: HTMLInputElement;
  let pagefind: Pagefind | null = null;

  async function load() {
    if (pagefind) return;
    const pagefindModuleUrl = '/pagefind/pagefind.js';
    pagefind = (await import(/* @vite-ignore */ pagefindModuleUrl)) as Pagefind;
    await pagefind.init();
  }

  let timer: ReturnType<typeof setTimeout>;
  async function onInput() {
    clearTimeout(timer);
    timer = setTimeout(run, 120);
  }

  let seq = 0;
  async function run() {
    if (query.trim().length < 2) { results = []; open = false; return; }
    const mine = ++seq;
    await load();
    const res = await pagefind!.search(query);
    const top = await Promise.all(res.results.slice(0, 8).map((r) => r.data()));
    if (mine !== seq) return;
    results = top.map((d) => ({ url: d.url.replace(/\.html$/, ''), title: d.meta.title, excerpt: d.excerpt }));
    active = results.length ? 0 : -1;
    open = results.length > 0;
  }

  function onKey(e: KeyboardEvent) {
    if (e.key === 'ArrowDown') { e.preventDefault(); active = Math.min(active + 1, results.length - 1); }
    else if (e.key === 'ArrowUp') { e.preventDefault(); active = Math.max(active - 1, 0); }
    else if (e.key === 'Enter' && active >= 0 && open) { e.preventDefault(); location.href = results[active].url; }
    else if (e.key === 'Escape') { open = false; }
  }

  function globalKey(e: KeyboardEvent) {
    if (e.key === '/' && document.activeElement !== input && !(document.activeElement instanceof HTMLInputElement)) {
      e.preventDefault();
      input.focus();
    }
  }

  $effect(() => {
    window.addEventListener('keydown', globalKey);
    if (initial) run();
    return () => window.removeEventListener('keydown', globalKey);
  });
</script>

<form role="search" action="/search" class="relative max-w-[760px]" onsubmit={(e) => { if (open && active >= 0) { e.preventDefault(); location.href = results[active].url; } }}>
  <div class="flex items-center gap-3 h-14 px-[18px] border border-line-warm-strong rounded-panel bg-bg/75 shadow-[0_0_0_1px_rgba(229,185,85,.15),0_12px_30px_rgba(0,0,0,.45)]">
    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" class="text-gold" stroke-width="2" stroke-linecap="round" aria-hidden="true"><circle cx="11" cy="11" r="7"/><path d="M20 20l-3.5-3.5"/></svg>
    <input
      bind:this={input}
      bind:value={query}
      type="search"
      name="q"
      autocomplete="off"
      placeholder="Search quests, items, dungeons, zones, talents…"
      aria-label="Search the site"
      aria-expanded={open}
      aria-controls="search-results"
      aria-activedescendant={open && active >= 0 ? `search-opt-${active}` : undefined}
      class="flex-1 bg-transparent text-[17px] text-text placeholder:text-muted outline-none"
      oninput={onInput}
      onkeydown={onKey}
      onfocus={load}
    />
    <kbd class="font-mono text-[12px] text-muted px-[7px] py-[3px] border border-line-warm rounded-control">/</kbd>
  </div>
  {#if open}
    <ul id="search-results" role="listbox" class="absolute left-0 right-0 top-full mt-2 bg-raised border border-line rounded-panel overflow-hidden z-10">
      {#each results as r, i}
        <li id={`search-opt-${i}`} role="option" aria-selected={i === active}>
          <a href={r.url} class={`flex flex-col gap-1 px-4 py-3 border-b border-line-soft last:border-b-0 ${i === active ? 'bg-card-top' : ''}`} onmouseenter={() => (active = i)} aria-labelledby={`search-title-${i}`} aria-describedby={`search-excerpt-${i}`}>
            <span id={`search-title-${i}`} class="text-[15px] font-semibold text-strong">{r.title}</span>
            <span id={`search-excerpt-${i}`} class="text-[13px] text-muted">{@html r.excerpt}</span>
          </a>
        </li>
      {/each}
    </ul>
  {/if}
</form>
