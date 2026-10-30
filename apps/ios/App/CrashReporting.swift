import Foundation
import Networking
import Sentry
import Telemetry

/// Crash and hang reporting through Sentry, privacy-first (apps/shared/privacy.md):
/// no user beyond Sentry's random install id, no IP inference, no screenshots or view
/// hierarchy, no performance tracing, and every URL and message scrubbed before it leaves.
/// Off unless CrashReportingPolicy allows it (release build, live API, a DSN).
enum CrashReporting {
    static func start(configuration: APIConfiguration, bundle: Bundle = .main) {
        let info = bundle.infoDictionary ?? [:]
        guard let policy = CrashReportingPolicy(
            dsn: info[CrashReportingPolicy.dsnKey] as? String,
            isDebugBuild: isDebugBuild,
            isLiveAPI: configuration.mode == .live,
            version: info["CFBundleShortVersionString"] as? String ?? "0.0.0",
            build: info["CFBundleVersion"] as? String ?? "0"
        ) else { return }
        SentrySDK.start { options in configure(options, policy: policy) }
    }

    private static var isDebugBuild: Bool {
        #if DEBUG
            true
        #else
            false
        #endif
    }

    private static func configure(_ options: Options, policy: CrashReportingPolicy) {
        options.dsn = policy.dsn
        options.releaseName = policy.release
        options.environment = policy.environment
        options.sendDefaultPii = false
        options.attachScreenshot = false
        options.attachViewHierarchy = false
        options.tracesSampleRate = 0
        options.enableAutoPerformanceTracing = false
        options.enableNetworkTracking = false
        options.enableCaptureFailedRequests = false
        options.enableUserInteractionTracing = false
        options.enableAppHangTracking = true
        options.appHangTimeoutInterval = CrashReportingPolicy.hangThreshold
        options.enableWatchdogTerminationTracking = true
        options.enableAutoSessionTracking = true
        options.beforeBreadcrumb = { crumb in scrub(crumb) }
        options.beforeSend = { event in scrub(event) }
    }

    private static func scrub(_ crumb: Breadcrumb) -> Breadcrumb {
        crumb.message = crumb.message.map(Scrubber.text)
        for (key, value) in crumb.data ?? [:] {
            crumb.setData(value: scrub(value), key: key)
        }
        return crumb
    }

    private static func scrub(_ event: Event) -> Event {
        // Sentry's random install id stays (crash-free counts need it); nothing else about the player.
        let installID = event.user?.userId
        event.user = installID.map { User(userId: $0) }
        event.serverName = nil
        event.request = nil
        event.message = event.message.map { SentryMessage(formatted: Scrubber.text($0.formatted)) }
        event.exceptions?.forEach { $0.value = $0.value.map(Scrubber.text) }
        event.breadcrumbs = event.breadcrumbs?.map(scrub)
        event.extra = event.extra?.mapValues(scrub)
        return event
    }

    /// A breadcrumb or extra value: URLs as URLs, other strings as text, the rest as is.
    private static func scrub(_ value: Any) -> Any {
        guard let text = value as? String else { return value }
        return text.contains("://") ? Scrubber.url(text) : Scrubber.text(text)
    }
}
