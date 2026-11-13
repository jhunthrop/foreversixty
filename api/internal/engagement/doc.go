// Package engagement keeps the daily active record (user_active_days): which signed-in accounts
// touched the product on which UTC day, and through which surfaces, so the owner can see who
// comes back. The auth middleware reports every authenticated request's actor to the Recorder,
// which writes at most once per account, day and surface per instance. The Store reads the
// owner's figures for /admin/activity: daily and weekly actives, new against returning, the
// per-surface split and week-over-week retention by signup week. The owner's account and the
// accounts holding a granted Premium are internal and stay out of every figure.
package engagement
