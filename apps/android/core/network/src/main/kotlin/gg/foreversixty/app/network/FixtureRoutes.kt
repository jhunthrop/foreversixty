package gg.foreversixty.app.network

import io.ktor.http.HttpMethod
import io.ktor.http.HttpStatusCode

/**
 * One route answered from a contract fixture. [pattern] is the path with `{name}` segments
 * (`{spec}.json` is a placeholder plus a literal suffix); [fixture] is lane C's file stem
 * (`<endpoint>[.<variant>]`, docs/contracts/app-v1.md) and may name a path segment
 * (`cdn.bis.{spec}`). A null [fixture] answers an empty body (a 204).
 */
data class FixtureRoute(
    val method: HttpMethod,
    val pattern: String,
    val fixture: String?,
    val status: HttpStatusCode = HttpStatusCode.OK,
) {
    /** The path's `{name}` values when this route serves the request, else null. */
    fun match(method: HttpMethod, path: String): Map<String, String>? {
        if (method != this.method) return null
        val want = segments(pattern)
        val got = segments(path)
        if (want.size != got.size) return null
        val params = mutableMapOf<String, String>()
        for ((expected, actual) in want.zip(got)) {
            val placeholder = PLACEHOLDER.matchEntire(expected)
            when {
                placeholder != null -> {
                    val (name, suffix) = placeholder.destructured
                    if (!actual.endsWith(suffix) || actual.length == suffix.length) return null
                    params[name] = actual.removeSuffix(suffix)
                }
                expected != actual -> return null
            }
        }
        return params
    }

    private fun segments(path: String) = path.split('/').filter { it.isNotEmpty() }

    private companion object {
        val PLACEHOLDER = Regex("""\{([a-z_]+)\}(.*)""")
    }
}

/** Which state of an endpoint fixture mode answers: lane C's variants and error variants. */
sealed interface FixtureVariant {
    /** The file stem for a route's own stem. */
    fun file(stem: String): String

    /** The status to answer with; null keeps the route's success status. */
    val status: HttpStatusCode?

    /** A success variant: `me` with `Named("empty")` is `me.empty.json`. */
    data class Named(val variant: String) : FixtureVariant {
        override fun file(stem: String) = "$stem.$variant"
        override val status: HttpStatusCode? = null
    }

    /** An error variant: `me.patch` with `Error("invalid", 400)` is `me.patch.error-invalid.json`. */
    data class Error(val code: String, override val status: HttpStatusCode) : FixtureVariant {
        override fun file(stem: String) = "$stem.error-$code"
    }

    /** The shared `error-unauthorized.json`: a missing, revoked or expired token on any route. */
    data object Unauthorized : FixtureVariant {
        override fun file(stem: String) = "error-unauthorized"
        override val status: HttpStatusCode = HttpStatusCode.Unauthorized
    }
}

/**
 * The route-to-fixture table for docs/contracts/fixtures/app (lane C owns the files;
 * docs/contracts/app-v1.md names them). API fixtures are whole envelopes; `cdn.*` fixtures are
 * the static files as served, with no envelope. A route joins by being added to the spec.
 */
object FixtureRoutes {

    /** Generated from api/openapi.yaml: a route's default fixture is its first `x-fixtures`, or `x-fixture-route`. */
    val api: List<FixtureRoute> = ContractRoutes.api

    /** The site's static data (`https://foreversixty.gg/data/...`); bare JSON, no envelope. */
    val cdn: List<FixtureRoute> = ContractRoutes.staticData

    val all: List<FixtureRoute> = api + cdn

    /** The route serving [method] [path] and its `{name}` values, or null. */
    fun resolve(method: HttpMethod, path: String): Pair<FixtureRoute, Map<String, String>>? =
        all.firstNotNullOfOrNull { route -> route.match(method, path)?.let { route to it } }

    /** The fixture file a request reads under [variants], or null for a 204 or an unknown route. */
    fun fileName(method: HttpMethod, path: String, variants: Map<String, FixtureVariant> = emptyMap()): String? {
        val (route, params) = resolve(method, path) ?: return null
        val stem = route.fixture ?: return null
        val named = params.entries.fold(stem) { acc, (key, value) -> acc.replace("{$key}", value) }
        return "${variants[stem]?.file(named) ?: named}.json"
    }
}
