import { describe, expect, it, vi } from 'vitest';
import {
  EVENTS_TABLE,
  FIGHT_MS,
  NOT_A_READ,
  QUERY_ABANDONED,
  QUERY_TEMPLATES,
  QUERY_TOO_SLOW,
  createQueryLayer,
  isReadOnlySql,
  normalizeCell,
  withWindow,
  type QueryEngine,
  type QueryResult,
} from './query';

/** Stands in for DuckDB: records what it was asked, answers a fixed shape. */
function fakeEngine(): QueryEngine & { opened: string[]; asked: string[]; closed: boolean } {
  const engine = {
    opened: [] as string[],
    asked: [] as string[],
    closed: false,
    async open(name: string): Promise<void> {
      engine.opened.push(name);
    },
    async query(sql: string): Promise<QueryResult> {
      engine.asked.push(sql);
      return { columns: ['event', 'n'], rows: [['SPELL_DAMAGE', 4]], elapsedMs: 3 };
    },
    async close(): Promise<void> {
      engine.closed = true;
    },
  };
  return engine;
}

/** A promise plus the handle to settle it, so a test can hold one call open. */
function held<T>(): { promise: Promise<T>; resolve: (value: T) => void } {
  let resolve: (value: T) => void = () => {};
  const promise = new Promise<T>((settle) => {
    resolve = settle;
  });
  return { promise, resolve };
}

describe('the SQL surface', () => {
  it('names the events table the templates read', () => {
    expect(EVENTS_TABLE).toBe("read_parquet('events.parquet')");
  });

  it('ships templates that all read the events table and all parse as reads', () => {
    expect(QUERY_TEMPLATES.length).toBeGreaterThanOrEqual(5);
    for (const template of QUERY_TEMPLATES) {
      expect(template.sql).toContain(EVENTS_TABLE);
      expect(isReadOnlySql(template.sql)).toBe(true);
      expect(template.label.length).toBeGreaterThan(0);
    }
  });

  it('refuses anything that is not a SELECT or a WITH', () => {
    expect(isReadOnlySql('SELECT 1')).toBe(true);
    expect(isReadOnlySql('  with x as (select 1) select * from x ')).toBe(true);
    expect(isReadOnlySql("COPY (SELECT 1) TO 'x.csv'")).toBe(false);
    expect(isReadOnlySql('DROP TABLE events')).toBe(false);
    expect(isReadOnlySql('INSTALL httpfs; SELECT 1')).toBe(false);
    expect(isReadOnlySql('SELECT 1; DROP TABLE events')).toBe(false);
  });

  it('scopes a template to the window by substituting the two bounds', () => {
    expect(withWindow('SELECT * FROM t WHERE :start <= 1 AND 1 < :end', { startMs: 1000, endMs: 5000 })).toBe(
      'SELECT * FROM t WHERE 1000 <= 1 AND 1 < 5000',
    );
  });

  // The Parquet's time column is absolute Unix nanoseconds (logs/engine/parquet/schema.go
  // `time_unix_nano`), while every millisecond the report shows -- the brush, the presets,
  // ?start= and ?end= -- counts from the fight's own start. A template that compared the
  // two directly would silently answer nothing, so every windowed template converts first.
  it('converts the event clock to milliseconds from the fight start before comparing', () => {
    expect(FIGHT_MS).toContain('time_unix_nano');
    expect(FIGHT_MS).toContain(`min(time_unix_nano) FROM ${EVENTS_TABLE}`);
    for (const template of QUERY_TEMPLATES) {
      if (!template.sql.includes(':start')) continue;
      expect(template.sql).toContain(FIGHT_MS);
      expect(template.sql).not.toMatch(/time_unix_nano\s+BETWEEN\s+:start/);
    }
  });
});

describe('normalizeCell', () => {
  it('passes primitives through and makes null of nothing', () => {
    expect(normalizeCell('SPELL_DAMAGE')).toBe('SPELL_DAMAGE');
    expect(normalizeCell(42)).toBe(42);
    expect(normalizeCell(true)).toBe(true);
    expect(normalizeCell(null)).toBeNull();
    expect(normalizeCell(undefined)).toBeNull();
  });

  it('turns a 64-bit column into a number, or a string when it will not fit', () => {
    expect(normalizeCell(5300n)).toBe(5300);
    expect(normalizeCell(12345678901234567890n)).toBe('12345678901234567890');
  });

  // sum() over a BIGINT is HUGEINT, which Arrow hands back as a four-word view. String()
  // on the raw view prints "1,2,0,0"; the view's own toString prints the decimal.
  it('reads Arrow’s 128-bit view as its decimal value', () => {
    const bigNum = Object.assign(Object.create({ [Symbol.for('isArrowBigNum')]: true }), {
      toString: () => '170141183460469231731687303715884105727',
    });
    expect(normalizeCell(bigNum)).toBe('170141183460469231731687303715884105727');
  });
});

describe('createQueryLayer', () => {
  const bytes = new Uint8Array([1, 2, 3]);

  it('does nothing at all until the first query', async () => {
    const load = vi.fn(async () => fakeEngine());
    const fetchBytes = vi.fn(async () => bytes);
    const layer = createQueryLayer({ load, fetchBytes });

    expect(layer.ready).toBe(false);
    expect(load).not.toHaveBeenCalled();
    expect(fetchBytes).not.toHaveBeenCalled();
  });

  it('loads the engine and the events file once, then reuses both', async () => {
    const engine = fakeEngine();
    const load = vi.fn(async () => engine);
    const fetchBytes = vi.fn(async () => bytes);
    const layer = createQueryLayer({ load, fetchBytes });

    await layer.run('/logs-data/reports/x/fights/3/events.parquet', 'SELECT 1');
    await layer.run('/logs-data/reports/x/fights/3/events.parquet', 'SELECT 2');

    expect(load).toHaveBeenCalledTimes(1);
    expect(fetchBytes).toHaveBeenCalledTimes(1);
    expect(engine.opened).toEqual(['events.parquet']);
    expect(engine.asked).toEqual(['SELECT 1', 'SELECT 2']);
    expect(layer.ready).toBe(true);
  });

  it('re-opens when the fight changes', async () => {
    const engine = fakeEngine();
    const layer = createQueryLayer({ load: async () => engine, fetchBytes: async () => bytes });

    await layer.run('/x/fights/3/events.parquet', 'SELECT 1');
    await layer.run('/x/fights/2/events.parquet', 'SELECT 1');

    expect(engine.opened).toEqual(['events.parquet', 'events.parquet']);
  });

  it('refuses a statement that is not a read, without loading anything', async () => {
    const load = vi.fn(async () => fakeEngine());
    const layer = createQueryLayer({ load, fetchBytes: async () => bytes });

    await expect(layer.run('/x/fights/3/events.parquet', 'DROP TABLE events')).rejects.toThrow(NOT_A_READ);
    expect(load).not.toHaveBeenCalled();
  });

  it('reports an empty answer as an empty answer, not as a failure', async () => {
    const empty: QueryEngine = {
      open: async () => undefined,
      query: async () => ({ columns: ['n'], rows: [], elapsedMs: 1 }),
      close: async () => undefined,
    };
    const layer = createQueryLayer({ load: async () => empty, fetchBytes: async () => bytes });

    await expect(layer.run('/x/fights/3/events.parquet', 'SELECT 1 WHERE false')).resolves.toEqual({
      columns: ['n'],
      rows: [],
      elapsedMs: 1,
    });
  });

  // A SELECT that parses here and then fails inside DuckDB -- a typo in a column, a table
  // that does not exist -- has to reach the visitor as DuckDB's own sentence, and must not
  // leave the layer wedged: the next query runs on the same engine.
  it('lets the engine’s own complaint through, and stays usable afterwards', async () => {
    const engine = fakeEngine();
    const fussy: QueryEngine = {
      open: (name, parquet) => engine.open(name, parquet),
      query: async (sql) => {
        if (sql.includes('nope')) throw new Error('Catalog Error: Table with name nope does not exist!');
        return engine.query(sql);
      },
      close: () => engine.close(),
    };
    const layer = createQueryLayer({ load: async () => fussy, fetchBytes: async () => bytes });

    await expect(layer.run('/x/fights/3/events.parquet', 'SELECT * FROM nope')).rejects.toThrow(
      'Catalog Error',
    );
    await expect(layer.run('/x/fights/3/events.parquet', 'SELECT 1')).resolves.toMatchObject({
      columns: ['event', 'n'],
    });
    expect(engine.closed).toBe(false);
  });
});

// Every step here is awaited and slow -- building the engine, downloading two to ten
// megabytes, running the query -- so two of them overlap the moment anyone clicks twice or
// changes fight mid-query. These hold one call open while another settles, because
// sequential awaits are the case that cannot race.
describe('createQueryLayer under overlap', () => {
  const bytes = new Uint8Array([1, 2, 3]);

  it('builds one engine for two queries that start together', async () => {
    const gate = held<QueryEngine>();
    const engine = fakeEngine();
    const load = vi.fn(() => gate.promise);
    const layer = createQueryLayer({ load, fetchBytes: async () => bytes });

    const first = layer.run('/x/fights/3/events.parquet', 'SELECT 1');
    const second = layer.run('/x/fights/3/events.parquet', 'SELECT 2');
    gate.resolve(engine);
    await Promise.all([first, second]);

    // One engine, one download, one registration: a second instance here would be a second
    // live DuckDB worker nothing ever terminates.
    expect(load).toHaveBeenCalledTimes(1);
    expect(engine.opened).toEqual(['events.parquet']);
    expect(engine.asked).toEqual(['SELECT 1', 'SELECT 2']);
  });

  // Leaving the Queries view is the ordinary way to reach this: the click lands between
  // pressing Run and the first line of work, so the cheapest right answer is to build
  // nothing at all rather than to build an engine and then terminate it.
  it('builds nothing when the layer closes before the run gets going', async () => {
    const load = vi.fn(async () => fakeEngine());
    const layer = createQueryLayer({ load, fetchBytes: async () => bytes });

    const running = layer.run('/x/fights/3/events.parquet', 'SELECT 1');
    const closing = layer.close();

    await expect(running).rejects.toThrow(QUERY_ABANDONED);
    await closing;
    expect(load).not.toHaveBeenCalled();
    expect(layer.ready).toBe(false);
  });

  it('closes an engine that finishes loading after the layer was already closed', async () => {
    const building = held<void>();
    const gate = held<QueryEngine>();
    const engine = fakeEngine();
    const layer = createQueryLayer({
      load: () => {
        building.resolve(undefined);
        return gate.promise;
      },
      fetchBytes: async () => bytes,
    });

    const running = layer.run('/x/fights/3/events.parquet', 'SELECT 1');
    await building.promise;
    const closing = layer.close();
    gate.resolve(engine);

    await expect(running).rejects.toThrow(QUERY_ABANDONED);
    await closing;
    // The engine arrives after nobody wants it and is terminated anyway: the alternative
    // is a live DuckDB worker with no reference left to stop it.
    expect(engine.closed).toBe(true);
    expect(engine.asked).toEqual([]);
    expect(layer.ready).toBe(false);
  });

  it('abandons a query whose events file is still downloading when the layer closes', async () => {
    const started = held<void>();
    const gate = held<Uint8Array>();
    const engine = fakeEngine();
    const layer = createQueryLayer({
      load: async () => engine,
      fetchBytes: () => {
        started.resolve(undefined);
        return gate.promise;
      },
    });

    const running = layer.run('/x/fights/3/events.parquet', 'SELECT 1');
    await started.promise;
    const closing = layer.close();
    gate.resolve(bytes);

    await expect(running).rejects.toThrow(QUERY_ABANDONED);
    await closing;
    expect(engine.opened).toEqual([]);
    expect(engine.closed).toBe(true);
  });

  it('stops a query that never answers, tears the engine down and works again after', async () => {
    const stuck = held<QueryResult>();
    const first = fakeEngine();
    const second = fakeEngine();
    const engines = [{ ...first, query: () => stuck.promise, close: first.close } as QueryEngine, second];
    const layer = createQueryLayer({
      load: async () => engines.shift() as QueryEngine,
      fetchBytes: async () => bytes,
      timeoutMs: 5,
    });

    await expect(layer.run('/x/fights/3/events.parquet', 'SELECT 1')).rejects.toThrow(QUERY_TOO_SLOW);
    expect(first.closed).toBe(true);
    expect(layer.ready).toBe(false);

    await expect(layer.run('/x/fights/3/events.parquet', 'SELECT 2')).resolves.toMatchObject({
      columns: ['event', 'n'],
    });
    expect(second.asked).toEqual(['SELECT 2']);
  });
});
