package gg.foreversixty.app.network

import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.AppExchangeRequest
import gg.foreversixty.app.model.AppExchangeResponse
import gg.foreversixty.app.model.ClassInfo
import gg.foreversixty.app.model.Me
import io.ktor.http.HttpMethod
import io.ktor.http.HttpStatusCode
import kotlinx.coroutines.test.runTest
import kotlinx.serialization.builtins.ListSerializer
import kotlinx.serialization.builtins.serializer
import kotlinx.serialization.json.JsonObject
import org.junit.Assert.assertEquals
import org.junit.Assert.assertFalse
import org.junit.Assert.assertNotNull
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test

class FixtureEngineTest {
    private val source = ClasspathFixtureSource(requireNotNull(javaClass.classLoader))
    private val config = ApiConfig(baseUrl = "https://api.example.test", mode = ApiMode.Fixture)

    private fun client(variants: Map<String, FixtureVariant> = emptyMap()) =
        ApiClient(ForeverHttpClient.create(config, NoToken, FixtureEngine.create(source, variants)))

    /** A request path for a route pattern; the CDN fixtures are for one spec and one class. */
    private fun samplePath(pattern: String): String =
        mapOf("spec" to "warrior-fury", "class" to "warrior", "build" to "1.60.1.70658")
            .entries.fold(pattern) { acc, (key, value) -> acc.replace("{$key}", value) }
            .replace(Regex("""\{[a-z_]+\}"""), "1")

    @Test
    fun `routes use lane C's endpoint names`() {
        assertEquals("me.json", FixtureRoutes.fileName(HttpMethod.Get, "/v1/me"))
        assertEquals("me.patch.json", FixtureRoutes.fileName(HttpMethod.Patch, "/v1/me"))
        assertEquals("character.card.json", FixtureRoutes.fileName(HttpMethod.Get, "/v1/characters/us/pvp/obnoxious-yell"))
        assertEquals("app.signin.exchange.json", FixtureRoutes.fileName(HttpMethod.Post, "/v1/auth/app/exchange"))
        assertEquals("cdn.bis.warrior-fury.json", FixtureRoutes.fileName(HttpMethod.Get, "/data/1.60.1.70658/bis/warrior-fury.json"))
        assertNull(FixtureRoutes.fileName(HttpMethod.Post, "/v1/sims/7/seen"))
    }

    @Test
    fun `variants pick lane C's variant and error files`() {
        val path = "/v1/me"
        assertEquals("me.empty.json", FixtureRoutes.fileName(HttpMethod.Get, path, mapOf("me" to FixtureVariant.Named("empty"))))
        assertEquals(
            "me.patch.error-invalid.json",
            FixtureRoutes.fileName(HttpMethod.Patch, path, mapOf("me.patch" to FixtureVariant.Error("invalid", HttpStatusCode.BadRequest))),
        )
        assertEquals("error-unauthorized.json", FixtureRoutes.fileName(HttpMethod.Get, path, mapOf("me" to FixtureVariant.Unauthorized)))
    }

    @Test
    fun `every catalogued route reads a fixture lane C wrote`() {
        for (route in FixtureRoutes.all) {
            val file = FixtureRoutes.fileName(route.method, samplePath(route.pattern))
            if (route.fixture == null) {
                assertNull(route.pattern, file)
                continue
            }
            assertNotNull("${route.pattern} unmatched", file)
            assertNotNull("missing $file", source.read(requireNotNull(file)))
        }
    }

    @Test
    fun `API fixtures are envelopes and CDN fixtures are bare`() {
        for (route in FixtureRoutes.all.filter { it.fixture != null }) {
            val file = requireNotNull(FixtureRoutes.fileName(route.method, samplePath(route.pattern)))
            val root = ApiJson.parseToJsonElement(requireNotNull(source.read(file)))
            val enveloped = (root as? JsonObject)?.containsKey("ok") == true
            assertEquals(file, route in FixtureRoutes.api, enveloped)
        }
    }

    @Test
    fun `me is served from the contract fixture`() = runTest {
        val result = AccountApi(client()).me()

        assertTrue("expected success, got $result", result is ApiResult.Success)
    }

    @Test
    fun `a named variant serves the empty account`() = runTest {
        val result = AccountApi(client(mapOf("me" to FixtureVariant.Named("empty")))).me()

        assertTrue((result as ApiResult.Success<Me>).data.characters.isEmpty())
    }

    @Test
    fun `the shared unauthorized fixture answers 401`() = runTest {
        val result = AccountApi(client(mapOf("me" to FixtureVariant.Unauthorized))).me()

        assertTrue(((result as ApiResult.Failure).error as ApiError.Server).isUnauthorized)
    }

    @Test
    fun `the exchange is served from the contract fixture`() = runTest {
        val result = client().post(
            "/v1/auth/app/exchange",
            AppExchangeRequest("code", "verifier", "Pixel"),
            AppExchangeRequest.serializer(),
            AppExchangeResponse.serializer(),
        )

        val answer = (result as ApiResult.Success).data
        assertTrue(answer.token.startsWith("fsd_"))
        assertTrue(answer.me.user.id > 0)
    }

    @Test
    fun `a 204 route answers no body`() = runTest {
        val result = client().post("/v1/sims/7/seen", Unit, Unit.serializer(), Unit.serializer())

        assertTrue("expected success, got $result", result is ApiResult.Success)
    }

    @Test
    fun `CDN fixtures are read bare through the static client`() = runTest {
        val static = StaticDataClient(FixtureEngine.create(source))

        val build = (static.activeBuild() as ApiResult.Success).data.build
        val classes = static.get("/data/$build/classes.json", ListSerializer(ClassInfo.serializer()))

        assertFalse(build.isEmpty())
        assertTrue((classes as ApiResult.Success).data.all { it.slug.isNotEmpty() })
    }

    @Test
    fun `a path with no fixture answers a not found envelope`() = runTest {
        val result = client().get("/v1/nothing/here", AppExchangeResponse.serializer())

        val error = (result as ApiResult.Failure).error as ApiError.Server
        assertEquals(404, error.status)
        assertEquals("not_found", error.code)
    }
}
