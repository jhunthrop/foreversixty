package gg.foreversixty.app.guides

import gg.foreversixty.app.bis.BisTestData
import gg.foreversixty.app.guides.logic.GuideFiles
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.Changelog
import gg.foreversixty.app.model.GuideFile
import gg.foreversixty.app.model.GuideIndex
import gg.foreversixty.app.model.KeyDate
import gg.foreversixty.app.model.TiersFile
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.CachedFile
import gg.foreversixty.app.network.ClasspathFixtureSource
import gg.foreversixty.app.network.FixtureRoutes
import gg.foreversixty.app.network.InMemoryStaticFileCache
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.StaticEndpoints
import gg.foreversixty.app.network.StaticResource
import io.ktor.http.HttpMethod
import java.io.IOException
import java.time.Instant
import java.time.ZoneId
import kotlinx.serialization.builtins.ListSerializer

/**
 * F9's test data: the W3-C1 contract fixtures, read as the app reads them. The story matches the
 * boards: Lightningbolltz (Elemental Shaman) for the guide, Emberquill (Fire Mage) for the tier list.
 */
object GuidesTestData {
    /** The fixtures' guide build. */
    const val BUILD = "1.60.1.70707"
    const val ELEMENTAL = "shaman-elemental"
    const val FIRE = "mage-fire"
    const val PROTECTION = "warrior-protection"

    /** Sunday 1 November 2026, 15:00 UTC: Name reservation is running, First raids ahead. */
    val now: Instant = Instant.parse("2026-11-01T15:00:00Z")
    val zone: ZoneId = ZoneId.of("UTC")

    fun index(): GuideIndex = BisTestData.fixture("cdn.guides.index", GuideIndex.serializer())
    fun elemental(): GuideFile = guide(ELEMENTAL)
    fun fire(): GuideFile = guide(FIRE)
    fun protection(): GuideFile = guide(PROTECTION)
    fun guide(spec: String): GuideFile = BisTestData.fixture("cdn.guides.$spec", GuideFile.serializer())
    fun tiers(): TiersFile = BisTestData.fixture("cdn.tiers", TiersFile.serializer())
    fun changelog(): Changelog = BisTestData.fixture("cdn.changelog", Changelog.serializer())
    fun dates(): List<KeyDate> = BisTestData.fixture("cdn.dates", ListSerializer(KeyDate.serializer()))

    val lightningbolltz = character("Lightningbolltz", "shaman", "Elemental")
    val emberquill = character("Emberquill", "mage", "Fire")
    val kraggor = character("Kraggor", "warrior", "Protection")

    private fun character(name: String, classSlug: String, spec: String) = AccountCharacter(
        key = "us/pvp/${name.lowercase()}",
        region = "us",
        ruleset = "pvp",
        name = name,
        source = "bnet",
        className = classSlug,
        level = 60,
        faction = "horde",
        spec = spec,
    )
}

/**
 * The static files from the fixtures through an in-memory offline cache, the way the live
 * [gg.foreversixty.app.guides.logic.LiveGuideFiles] answers: a good read is saved; offline answers the
 * saved copy or a network failure. [activeBuild] replaces the active-build fixture's answer.
 */
class FakeGuideFiles(val cache: InMemoryStaticFileCache = InMemoryStaticFileCache()) : GuideFiles {
    var offline = false
    var activeBuild: String? = null
    var clock = SAVED_AT
    val loads = mutableListOf<String>()
    private val source = ClasspathFixtureSource(requireNotNull(javaClass.classLoader))

    override suspend fun <T> load(resource: StaticResource<T>): ApiResult<CachedFile<T>> {
        loads += resource.path
        if (offline) return cached(resource)?.let { ApiResult.Success(it, "") } ?: ApiResult.Failure(ApiError.Network(IOException("offline")))
        val body = body(resource.path) ?: return ApiResult.Failure(ApiError.Status(NOT_FOUND))
        cache.write(body, resource.path, clock)
        return when (val decoded = StaticDataClient.decode(resource, body)) {
            is ApiResult.Success -> ApiResult.Success(CachedFile(decoded.data, clock, isFromCache = false), "")
            is ApiResult.Failure -> decoded
        }
    }

    override fun <T> cached(resource: StaticResource<T>): CachedFile<T>? {
        val entry = cache.read(resource.path) ?: return null
        val decoded = StaticDataClient.decode(resource, entry.body) as? ApiResult.Success ?: return null
        return CachedFile(decoded.data, entry.savedAtMillis, isFromCache = true)
    }

    override fun savedBytes(resource: StaticResource<*>): Long? = cache.read(resource.path)?.body?.toByteArray()?.size?.toLong()

    private fun body(path: String): String? {
        if (path == StaticEndpoints.getActiveBuild().path) activeBuild?.let { return "{\"build\":\"$it\"}" }
        val name = FixtureRoutes.fileName(HttpMethod.Get, path) ?: return null
        return source.read(name)
    }

    companion object {
        /** Thursday 29 October 2026, 21:41 Chicago: when the saved copies were fetched. */
        val SAVED_AT: Long = Instant.parse("2026-10-30T02:41:00Z").toEpochMilli()
        private const val NOT_FOUND = 404
    }
}
