package gg.foreversixty.app.network

import android.content.res.AssetManager
import gg.foreversixty.app.model.ApiJson
import io.ktor.client.engine.HttpClientEngine
import io.ktor.client.engine.mock.MockEngine
import io.ktor.client.engine.mock.respond
import io.ktor.http.ContentType
import io.ktor.http.HttpHeaders
import io.ktor.http.HttpMethod
import io.ktor.http.HttpStatusCode
import io.ktor.http.headersOf
import java.io.FileNotFoundException

/** Reads one contract fixture by file name, or null when there is none. */
fun interface FixtureSource {
    fun read(name: String): String?
}

/** Fixtures from the debug APK's assets (copied there by the contract-fixtures Gradle task). */
class AssetFixtureSource(private val assets: AssetManager) : FixtureSource {
    override fun read(name: String): String? = try {
        assets.open("$FIXTURE_ASSET_DIR/$name").bufferedReader().use { it.readText() }
    } catch (_: FileNotFoundException) {
        null
    }
}

/** Fixtures from the JVM test classpath (the same files, copied by the same task). */
class ClasspathFixtureSource(private val loader: ClassLoader) : FixtureSource {
    override fun read(name: String): String? =
        loader.getResource("$FIXTURE_ASSET_DIR/$name")?.readText()
}

/** Where fixtures sit inside assets and the test classpath. */
const val FIXTURE_ASSET_DIR = "fixtures/app"

/**
 * An engine that answers every request from [FixtureSource], with no network. [variants] picks
 * a state per route stem (`mapOf("me" to FixtureVariant.Named("empty"))`); every other route
 * answers its success fixture. Fixtures are served exactly as lane C wrote them: envelopes for
 * the API, bare JSON for `cdn.*`.
 */
object FixtureEngine {
    fun create(source: FixtureSource, variants: Map<String, FixtureVariant> = emptyMap()): HttpClientEngine =
        MockEngine { request ->
            val path = request.url.encodedPath
            val jsonHeaders = headersOf(HttpHeaders.ContentType, ContentType.Application.Json.toString())
            val (route, _) = FixtureRoutes.resolve(request.method, path)
                ?: return@MockEngine respond(notFoundBody(path), HttpStatusCode.NotFound, jsonHeaders)
            val status = route.fixture?.let { variants[it]?.status } ?: route.status
            val file = FixtureRoutes.fileName(request.method, path, variants)
                ?: return@MockEngine respond("", status)
            val body = source.read(file)
                ?: return@MockEngine respond(notFoundBody("$path -> $file"), HttpStatusCode.NotFound, jsonHeaders)
            respond(body, status, jsonHeaders)
        }

    private fun notFoundBody(what: String): String = ApiJson.encodeToString(
        ApiEnvelope.serializer(),
        ApiEnvelope(
            ok = false,
            error = ApiErrorBody(code = "not_found", message = "no contract fixture for $what"),
            requestId = "fixture-missing",
        ),
    )
}
