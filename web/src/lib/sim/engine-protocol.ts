// web/src/lib/sim/engine-protocol.ts
// Adapts between the wasm engine's simSplit/simCombine shapes (sim/cmd/wasm/main.go: each
// crosses the boundary as ONE JSON string, never a JS array) and the worker pool's message
// protocol (worker.ts's ToWorker/FromWorker), which carries one JSON string per shard both
// ways. Shared by the real worker (sim.worker.ts) and the in-process fake
// (test-support/fake-worker.ts) so the two adapt identically -- the duplication the final
// review's H1 fix would otherwise leave behind.
import type { SimRequest, SimResult } from './types';

/**
 * simSplit returns ONE JSON string encoding a SimRequest[] (main.go:198-203's
 * json.Marshal(parts)); the pool's 'many' message wants one JSON string per shard, since
 * each shard is handed straight back to simRun as its requestJSON.
 */
export function shardsFromSplit(splitJSON: string): string[] {
  return (JSON.parse(splitJSON) as SimRequest[]).map((request) => JSON.stringify(request));
}

/**
 * simCombine takes ONE JSON string encoding a SimResult[] (main.go:211-213's
 * json.Unmarshal(args[0].String(), &parts)); the pool holds one JSON string per shard
 * result, gathered from each worker's 'one' message.
 */
export function combineInputFromShards(shardResultsJSON: readonly string[]): string {
  return JSON.stringify(shardResultsJSON.map((json) => JSON.parse(json) as SimResult));
}
