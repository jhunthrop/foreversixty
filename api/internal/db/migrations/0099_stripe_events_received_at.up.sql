-- stripe-reconcile deletes stripe_events rows past the retention window every night (issue #3
-- M8): a raw payload carries the buyer's name, email and address, and nothing reads it once
-- Stripe can no longer resend the event. This migration deletes nothing itself.
create index stripe_events_received_at_idx on stripe_events (received_at);
