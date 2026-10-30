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
 * the static files as served, with no envelope. Feature lanes add a row per route they call.
 */
object FixtureRoutes {
    private val Get = HttpMethod.Get
    private val Post = HttpMethod.Post
    private val Put = HttpMethod.Put
    private val Patch = HttpMethod.Patch
    private val Delete = HttpMethod.Delete

    val api: List<FixtureRoute> = listOf(
        FixtureRoute(Get, "/v1/me", "me"),
        FixtureRoute(Patch, "/v1/me", "me.patch"),
        FixtureRoute(Get, "/v1/characters/search", "character.search"),
        FixtureRoute(Get, "/v1/characters/{region}/{ruleset}/{name}", "character.card"),
        FixtureRoute(Get, "/v1/characters/{region}/{ruleset}/{name}/sim-input", "character.sim-input.addon"),
        FixtureRoute(Get, "/v1/sims", "sims.list"),
        FixtureRoute(Get, "/v1/sims/{id}", "sims.result"),
        FixtureRoute(Get, "/v1/sims/{id}/progress", "sims.progress"),
        FixtureRoute(Post, "/v1/sims/{id}/rerun", "sims.rerun", HttpStatusCode.Accepted),
        FixtureRoute(Post, "/v1/sims/{id}/seen", null, HttpStatusCode.NoContent),
        FixtureRoute(Get, "/v1/guilds/{id}/home", "guild.home"),
        FixtureRoute(Get, "/v1/guilds/{id}/raids", "guild.raids"),
        FixtureRoute(Get, "/v1/guilds/{id}/raids/{raid}", "guild.raid"),
        FixtureRoute(Patch, "/v1/guilds/{id}/raids/{raid}", "guild.raid.patch"),
        FixtureRoute(Post, "/v1/guilds/{id}/raids/{raid}/lock", "guild.raid.patch"),
        FixtureRoute(Post, "/v1/guilds/{id}/raids/{raid}/unlock", "guild.raid.patch"),
        FixtureRoute(Post, "/v1/guilds/{id}/raids/{raid}/cancel", "guild.raid.patch"),
        FixtureRoute(Put, "/v1/guilds/{id}/raids/{raid}/signups/{region}/{ruleset}/{name}", "guild.signup"),
        FixtureRoute(Get, "/v1/guilds/{id}/loot", "guild.loot"),
        FixtureRoute(Get, "/v1/reports", "reports.list"),
        FixtureRoute(Get, "/v1/reports/{id}", "reports.get"),
        FixtureRoute(Get, "/v1/builds", "builds.list"),
        FixtureRoute(Get, "/v1/builds/{id}", "builds.get"),
        FixtureRoute(Post, "/v1/builds", "builds.save", HttpStatusCode.Created),
        FixtureRoute(Post, "/v1/auth/app/exchange", "app.signin.exchange", HttpStatusCode.Created),
        FixtureRoute(Post, "/v1/auth/app/logout", "app.logout"),
        FixtureRoute(Get, "/v1/me/notifications", "app.notifications"),
        FixtureRoute(Patch, "/v1/me/notifications", "app.notifications"),
        FixtureRoute(Post, "/v1/me/push-tokens", "app.push-tokens.register", HttpStatusCode.Created),
        FixtureRoute(Delete, "/v1/me/push-tokens/{token}", null, HttpStatusCode.NoContent),
        FixtureRoute(Get, "/v1/app/config", "app.config"),
    )

    /** The site's static data (`https://foreversixty.gg/data/...`); bare JSON, no envelope. */
    val cdn: List<FixtureRoute> = listOf(
        FixtureRoute(Get, "/data/active-build.json", "cdn.active-build"),
        FixtureRoute(Get, "/data/dates.json", "cdn.dates"),
        FixtureRoute(Get, "/data/{build}/classes.json", "cdn.classes"),
        FixtureRoute(Get, "/data/{build}/bis/{spec}.json", "cdn.bis.{spec}"),
        FixtureRoute(Get, "/data/{build}/items/{class}.json", "cdn.items.{class}"),
        FixtureRoute(Get, "/data/{build}/talents/{class}.json", "cdn.talents.{class}"),
    )

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
