package gg.foreversixty.app.telemetry

import java.net.URI
import java.net.URISyntaxException

/**
 * What leaves the device in a crash or hang report, made safe first. The same rules run on iOS
 * (Sources/Telemetry/Scrubber.swift); apps/shared/privacy.md is what they promise.
 *
 * - A URL keeps scheme, host and path and loses its query, fragment and user info, so the
 *   sign-in callback's `code` and `state` never travel.
 * - Free text loses device tokens (`fsd_…`), bearer credentials and any `code`, `state`,
 *   `token`, `code_verifier` or `key` parameter value.
 */
object Scrubber {
    const val REDACTED = "[redacted]"

    private val rules = listOf(
        Regex("""fsd_[A-Za-z0-9_\-]+""") to "fsd_$REDACTED",
        Regex("""(?i)\bbearer\s+[A-Za-z0-9._~+/=\-]+""") to "Bearer $REDACTED",
        Regex("""(?i)\b(code|code_verifier|state|token|access_token|refresh_token|key)=[^&\s#"']+""") to
            "$1=$REDACTED",
    )

    /** The URL without its query and fragment. Text that is not an absolute URL is scrubbed as text. */
    fun url(raw: String): String {
        val uri = try {
            URI(raw)
        } catch (_: URISyntaxException) {
            return text(raw)
        }
        if (uri.scheme == null || uri.host == null) return text(raw)
        val port = if (uri.port == -1) "" else ":${uri.port}"
        return "${uri.scheme}://${uri.host}$port${uri.rawPath.orEmpty()}"
    }

    /** Free text with every credential-shaped run replaced. */
    fun text(raw: String): String = rules.fold(raw) { result, (pattern, template) ->
        pattern.replace(result, template)
    }

    /** A breadcrumb or extra value: URLs as URLs, everything else as text. */
    fun value(raw: String): String = if ("://" in raw) url(raw) else text(raw)
}
