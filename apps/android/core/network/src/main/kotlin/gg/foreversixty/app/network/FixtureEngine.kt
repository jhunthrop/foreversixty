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
import kotlinx.serialization.json.JsonObject
import kotlinx.serialization.json.jsonObject
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
 * Maps a request to its fixture file. `GET /v1/me` is `me.json`; `/v1/auth/app/exchange` is
 * `auth-app-exchange.json`. A method-specific file (`me.patch.json`) wins over the plain one,
 * so a path with both a GET and a write can carry both. This is the single place to change if
 * lane C's contract names fixtures differently.
 */
object FixtureRoutes {
    private const val API_PREFIX = "/v1/"

    fun candidates(method: HttpMethod, path: String): List<String> {
        val slug = path.removePrefix(API_PREFIX).trim('/').replace('/', '-')
        val specific = "$slug.${method.value.lowercase()}.json"
        return if (method == HttpMethod.Get) listOf("$slug.json", specific) else listOf(specific, "$slug.json")
    }
}

/** An engine that answers every request from [FixtureSource], with no network. */
object FixtureEngine {
    fun create(source: FixtureSource): HttpClientEngine = MockEngine { request ->
        val name = FixtureRoutes.candidates(request.method, request.url.encodedPath)
            .firstNotNullOfOrNull { candidate -> source.read(candidate)?.let { candidate to it } }
        val jsonHeaders = headersOf(HttpHeaders.ContentType, ContentType.Application.Json.toString())
        if (name == null) {
            respond(notFoundBody(request.url.encodedPath), HttpStatusCode.NotFound, jsonHeaders)
        } else {
            respond(asEnvelope(name.first, name.second), HttpStatusCode.OK, jsonHeaders)
        }
    }

    /** A fixture may be the whole envelope or the data alone; the client always sees an envelope. */
    internal fun asEnvelope(name: String, raw: String): String {
        val root = ApiJson.parseToJsonElement(raw).jsonObject
        if ("ok" in root) return raw
        val envelope = ApiEnvelope(ok = true, data = JsonObject(root), requestId = "fixture-$name")
        return ApiJson.encodeToString(ApiEnvelope.serializer(), envelope)
    }

    private fun notFoundBody(path: String): String = ApiJson.encodeToString(
        ApiEnvelope.serializer(),
        ApiEnvelope(
            ok = false,
            error = ApiErrorBody(code = "not_found", message = "no contract fixture for $path"),
            requestId = "fixture-missing",
        ),
    )
}
