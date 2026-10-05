<!-- web/src/components/guild/GuildRaids.svelte -->
<!-- Guild control-centre spec §4.C: every raid night, one expandable in place (no
     navigation). A night in Barrow Deeps or Hyjal Summit shows "Pull N — no named
     encounter published" rather than inventing a boss name (§12.1). Shown to every role
     (public sees public reports only, per the existing report-visibility rule server
     side) -- this component renders whatever `rows` it is given. -->
<script lang="ts">
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import { classColorVar } from '../../lib/report/format';
  import type { GuildRaidRow } from '../../lib/guild/api';
  import EmptyState from '../ui/EmptyState.svelte';

  let { rows, officer }: { rows: GuildRaidRow[]; officer: boolean } = $props();

  let expanded = $state<string | null>(null);
  let defaulted = $state(false);

  /** `rows` arrives after this component's own mount (Guild.svelte only renders it once
   *  the raids fetch resolves, but a future caller could pass an empty array first) -- a
   *  plain `$state` initializer only captures its value once, so the first night's own
   *  default-open state is set here instead, exactly once. */
  $effect(() => {
    if (defaulted || rows.length === 0) return;
    expanded = rows[0].id;
    defaulted = true;
  });

  function minutes(ms: number): string {
    const total = Math.round(ms / 60000);
    return `${total} min`;
  }
  function clock(ms: number): string {
    const totalSeconds = Math.round(ms / 1000);
    const m = Math.floor(totalSeconds / 60);
    const s = totalSeconds % 60;
    return `${m}:${String(s).padStart(2, '0')}`;
  }
</script>

<section class="flex flex-col gap-2" data-testid="guild-raids-tab">
  <h2 class="section-title text-[18px]">Raids</h2>
  {#if rows.length === 0}
    <EmptyState
      message="No raid nights logged yet. Upload a raid log to get started."
      testid="guild-raids-empty"
    />
  {:else}
    <div class="flex flex-col" data-testid="guild-raids-rows">
      {#each rows as row (row.id)}
        {@const isOpen = expanded === row.id}
        <div class="border-line-soft border-b">
          <button
            class="flex w-full flex-wrap items-center gap-3 py-2 text-left"
            onclick={() => (expanded = isOpen ? null : row.id)}
            aria-expanded={isOpen}
            data-testid={`guild-raid-row-${row.id}`}
          >
            <span class="min-w-[140px] text-[14px] font-semibold">{row.zone}</span>
            <span class="text-muted tabular font-mono text-[13px]">{row.created_at.slice(0, 10)}</span>
            <span class="text-muted tabular font-mono text-[13px]">{minutes(row.duration_ms)}</span>
            <span class="text-muted tabular ml-auto font-mono text-[13px]">
              {row.kill_count} named kill{row.kill_count === 1 ? '' : 's'} · {row.wipe_count} wipe{row.wipe_count ===
              1
                ? ''
                : 's'}
            </span>
          </button>
          {#if isOpen}
            <div class="flex flex-col gap-3 pb-4" data-testid={`guild-raid-expanded-${row.id}`}>
              <div class="flex flex-wrap items-center gap-1">
                {#each row.present.slice(0, 20) as p (p.character_key)}
                  <ClassCrestRing characterClass={p.class} size={22} />
                {/each}
                <span class="text-muted ml-2 text-[12px]">{row.present.length} of {row.raiders} present</span>
              </div>
              {#if row.top_parse !== null}
                <p class="text-[13px]">
                  Top parse: <span style={`color:${classColorVar(row.top_parse.class)};font-weight:600`}
                    >{row.top_parse.name}</span
                  >
                  {row.top_parse.value.toFixed(1)}
                  {row.top_parse.metric.toUpperCase()}
                </p>
              {/if}
              <div class="flex flex-col" data-testid={`guild-raid-fights-${row.id}`}>
                {#each row.fights as fight (fight.index)}
                  <div class="border-line-soft flex items-center gap-3 border-b py-1.5">
                    {#if fight.name !== null}
                      <span class="min-w-[120px] text-[13px]">{fight.name} · pull {fight.index + 1}</span>
                      <span class="text-muted tabular font-mono text-[12px]">{clock(fight.duration_ms)}</span>
                      <span
                        class="pill"
                        class:pill-kill={fight.kill}
                        style={fight.kill
                          ? 'color:#7bff5c;background:rgba(30,255,0,.10);border-color:rgba(30,255,0,.25)'
                          : ''}
                      >
                        {fight.kill ? 'Kill' : 'Wipe'}
                      </span>
                    {:else}
                      <span class="min-w-[120px] text-[13px]">Pull {fight.index + 1}</span>
                      <span class="text-muted text-[12px]">no named encounter published</span>
                    {/if}
                  </div>
                {/each}
              </div>
              {#if officer}
                <p class="text-muted text-[11px]">
                  Attach to this guild / visibility controls are officer-only (unchanged from round 1).
                </p>
              {/if}
            </div>
          {/if}
        </div>
      {/each}
    </div>
  {/if}
</section>
