package gg.foreversixty.app.telemetry

/**
 * Whether this build reports crashes and hangs, and with what. Reporting runs only in a release
 * build talking to the live API with a DSN configured: never in debug, never in fixture mode,
 * never with the placeholder empty DSN.
 */
data class CrashReportingPolicy(val dsn: String, val release: String, val environment: String) {
    companion object {
        /** A main-thread stall this long is reported as a hang (an ANR on Android). */
        const val HANG_THRESHOLD_MILLIS = 2_000L
        private const val ENVIRONMENT = "production"

        /** Null when reporting must stay off for this build. */
        fun of(dsn: String?, isDebugBuild: Boolean, isLiveApi: Boolean, version: String, build: Int): CrashReportingPolicy? {
            if (isDebugBuild || !isLiveApi || dsn.isNullOrEmpty() || !dsn.startsWith("https://")) return null
            return CrashReportingPolicy(dsn, release = "gg.foreversixty.app@$version+$build", environment = ENVIRONMENT)
        }
    }
}
