package gg.foreversixty.app

import android.content.Context
import gg.foreversixty.app.network.ApiConfig
import gg.foreversixty.app.network.ApiMode
import gg.foreversixty.app.telemetry.CrashReportingPolicy
import gg.foreversixty.app.telemetry.Scrubber
import io.sentry.Breadcrumb
import io.sentry.SentryEvent
import io.sentry.android.core.SentryAndroid
import io.sentry.android.core.SentryAndroidOptions
import io.sentry.protocol.Message
import io.sentry.protocol.User

/**
 * Crash and hang (ANR) reporting through Sentry, privacy-first (apps/shared/privacy.md): no user
 * beyond Sentry's random install id, no IP, no screenshots or view hierarchy, no performance
 * tracing, and every URL and message scrubbed before it leaves. Off unless [CrashReportingPolicy]
 * allows it (release build, live API, a DSN). Auto-init is off in the manifest; this is the
 * only way Sentry starts.
 */
object CrashReporting {
    fun start(context: Context, config: ApiConfig) {
        val policy = CrashReportingPolicy.of(
            dsn = BuildConfig.SENTRY_DSN,
            isDebugBuild = BuildConfig.DEBUG,
            isLiveApi = config.mode == ApiMode.Live,
            version = BuildConfig.VERSION_NAME,
            build = BuildConfig.VERSION_CODE,
        ) ?: return
        SentryAndroid.init(context) { options -> configure(options, policy) }
    }

    private fun configure(options: SentryAndroidOptions, policy: CrashReportingPolicy) {
        options.dsn = policy.dsn
        options.release = policy.release
        options.environment = policy.environment
        // sendDefaultPii stays at its default, off (deprecated in favour of dataCollection, whose
        // default collects no PII): no IP, no user, no device name.
        options.isAttachScreenshot = false
        options.isAttachViewHierarchy = false
        options.tracesSampleRate = null
        options.isEnableUserInteractionBreadcrumbs = false
        options.isAnrEnabled = true
        options.anrTimeoutIntervalMillis = CrashReportingPolicy.HANG_THRESHOLD_MILLIS
        options.isEnableAutoSessionTracking = true
        options.setBeforeBreadcrumb { crumb, _ -> scrub(crumb) }
        options.setBeforeSend { event, _ -> scrub(event) }
    }

    private fun scrub(crumb: Breadcrumb): Breadcrumb {
        crumb.message = crumb.message?.let(Scrubber::text)
        crumb.data.entries.forEach { entry ->
            val value = entry.value
            if (value is String) entry.setValue(Scrubber.value(value))
        }
        return crumb
    }

    private fun scrub(event: SentryEvent): SentryEvent {
        // Sentry's random install id stays (crash-free counts need it); nothing else about the player.
        event.user = event.user?.id?.let { id -> User().apply { this.id = id } }
        event.serverName = null
        event.request = null
        event.message = event.message?.formatted?.let { text -> Message().apply { formatted = Scrubber.text(text) } }
        event.exceptions?.forEach { exception -> exception.value = exception.value?.let(Scrubber::text) }
        event.breadcrumbs?.forEach { scrub(it) }
        event.extras?.keys?.toList()?.forEach { key ->
            val value = event.getExtra(key)
            if (value is String) event.setExtra(key, Scrubber.value(value))
        }
        return event
    }
}
