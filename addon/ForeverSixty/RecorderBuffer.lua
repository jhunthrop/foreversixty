-- addon/ForeverSixty/RecorderBuffer.lua
-- The measurement recorder's ring buffer: a fixed number of events, the
-- oldest overwritten first. It lives in its own file so the storage rule
-- (what the SavedVariables table looks like) is one place, and the pipeline
-- reader (data/pipeline/recorder.py) has one contract to mirror.
--
-- The store is a plain table, so it serialises as SavedVariables as is:
--   { schema = n, cap = n, start = n, count = n, events = { ... } }
-- `events` is the physical array, never longer than `cap`; `start` is the
-- index of the OLDEST event, so chronological order is start..count, then
-- 1..start-1. Until the buffer first fills, `start` is 1.
local RecorderBuffer = {}

--- A fresh, empty store.
function RecorderBuffer.new(schema, capacity)
	assert(type(capacity) == "number" and capacity >= 1, "capacity must be a positive number")
	return { schema = schema, cap = capacity, start = 1, count = 0, events = {} }
end

--- Append one event, overwriting the oldest when the store is full.
function RecorderBuffer.push(store, event)
	if store.count < store.cap then
		store.count = store.count + 1
		store.events[store.count] = event
		return
	end
	store.events[store.start] = event
	store.start = store.start % store.cap + 1
end

--- The events oldest first, as a new array.
function RecorderBuffer.ordered(store)
	local ordered = {}
	for offset = 0, store.count - 1 do
		ordered[offset + 1] = store.events[(store.start - 1 + offset) % store.cap + 1]
	end
	return ordered
end

--- Empty the store in place, keeping its schema and capacity.
function RecorderBuffer.clear(store)
	store.start, store.count, store.events = 1, 0, {}
end

return RecorderBuffer
