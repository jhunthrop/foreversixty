package gg.foreversixty.app.raids

import android.content.res.Resources
import androidx.test.core.app.ApplicationProvider
import android.content.Context
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.GuildRaids
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.RaidNightSummary
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.model.SignupAnswer
import gg.foreversixty.app.model.SignupRequest
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiErrorBody
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.logic.Cached
import gg.foreversixty.app.raids.logic.NotificationPermission
import gg.foreversixty.app.raids.logic.RaidCache
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidHaptics
import gg.foreversixty.app.raids.logic.RaidService
import gg.foreversixty.app.raids.logic.ReminderAskPolicy
import gg.foreversixty.app.raids.logic.ReminderAskStore
import gg.foreversixty.app.raids.logic.SignupController
import java.time.Instant
import java.time.ZoneId
import kotlin.time.Duration.Companion.seconds
import kotlinx.serialization.KSerializer
import kotlinx.serialization.json.JsonElement
import kotlinx.serialization.json.JsonObject
import kotlinx.serialization.json.jsonObject

/**
 * The contract fixtures on the test classpath (forever.contract-fixtures). Each helper throws an
 * ordinary exception, so a missing fixture fails its own test and nothing else.
 */
object RaidFixture {
    fun text(name: String): String {
        val resource = requireNotNull(javaClass.classLoader?.getResource("fixtures/app/$name.json")) { "no fixture $name" }
        return resource.readText()
    }

    private fun data(name: String): JsonElement {
        val root = ApiJson.parseToJsonElement(text(name)).jsonObject
        return requireNotNull(root["data"]) { "fixture $name has no data" }
    }

    fun <T> load(serializer: KSerializer<T>, name: String): T = ApiJson.decodeFromJsonElement(serializer, data(name))

    fun raid(name: String = "guild.raid"): RaidView = load(RaidView.serializer(), name)

    fun error(name: String, status: Int): ApiError {
        val root = ApiJson.parseToJsonElement(text(name)).jsonObject
        val body = ApiJson.decodeFromJsonElement(ApiErrorBody.serializer(), requireNotNull(root["error"]) { "fixture $name has no error" })
        return ApiError.Server(status, body.code, body.message, body.fields, "test")
    }

    fun push(name: String): JsonObject = ApiJson.parseToJsonElement(text(name)).jsonObject
}

/** Fixed instants: two days before the fixture raid (invites 2026-12-20 00:45Z, pull 01:00Z). */
object RaidTime {
    val chicago: ZoneId = ZoneId.of("America/Chicago")
    val now: Instant = Instant.parse("2026-12-17T22:30:00Z")
}

fun resources(): Resources = ApplicationProvider.getApplicationContext<Context>().resources

fun clock(): RaidClock = RaidClock(resources(), RaidTime.chicago)

/** A scripted raid service: each call answers from its lambda; every sign-up is recorded. */
class FakeRaidService : RaidService {
    var meAnswer: () -> ApiResult<Me> = { ok(RaidFixture.load(Me.serializer(), "me")) }
    var raidsAnswer: () -> ApiResult<GuildRaids> = { ok(RaidFixture.load(GuildRaids.serializer(), "guild.raids")) }
    var raidAnswer: () -> ApiResult<RaidView> = { ok(RaidFixture.raid()) }
    var signupAnswer: (SignupRequest) -> ApiResult<SignupAnswer> = { ok(RaidFixture.load(SignupAnswer.serializer(), "guild.signup")) }
    var lockAnswer: () -> ApiResult<RaidView> = { ok(RaidFixture.raid("guild.raid.locked")) }
    var nightAnswer: () -> ApiResult<RaidNightSummary> = { ok(RaidFixture.load(RaidNightSummary.serializer(), "me.raid-night")) }
    val signups = mutableListOf<Pair<String, SignupRequest>>()
    val locks = mutableListOf<String>()

    override suspend fun me() = meAnswer()
    override suspend fun raids(guild: Long) = raidsAnswer()
    override suspend fun raid(guild: Long, raid: Long) = raidAnswer()
    override suspend fun signup(guild: Long, raid: Long, characterKey: String, request: SignupRequest): ApiResult<SignupAnswer> {
        signups += characterKey to request
        return signupAnswer(request)
    }
    override suspend fun lock(guild: Long, raid: Long): ApiResult<RaidView> {
        locks += "lock"
        return lockAnswer()
    }
    override suspend fun unlock(guild: Long, raid: Long): ApiResult<RaidView> {
        locks += "unlock"
        return ok(RaidFixture.raid("guild.raid.officer"))
    }
    override suspend fun night(raid: Long) = nightAnswer()
}

fun <T> ok(data: T): ApiResult<T> = ApiResult.Success(data, "test")
fun fail(error: ApiError): ApiResult<Nothing> = ApiResult.Failure(error)

class RecordingHaptics : RaidHaptics {
    val calls = mutableListOf<String>()
    override fun selection() { calls += "selection" }
    override fun success() { calls += "success" }
    override fun error() { calls += "error" }
}

class MemoryAskStore : ReminderAskStore {
    var shown = 0
    var last: Instant? = null
    override fun policy() = ReminderAskPolicy(shown, last)
    override fun recordShown(at: Instant) {
        shown += 1
        last = at
    }
}

class MemoryRaidCache : RaidCache {
    private val store = mutableMapOf<String, Pair<String, Instant>>()

    override fun <T> read(key: String, serializer: KSerializer<T>): Cached<T>? =
        store[key]?.let { (json, at) -> Cached(ApiJson.decodeFromString(serializer, json), at) }

    override fun <T> write(key: String, serializer: KSerializer<T>, value: T, at: Instant) {
        store[key] = ApiJson.encodeToString(serializer, value) to at
    }
}

fun makeController(
    raid: RaidView = RaidFixture.raid(),
    service: FakeRaidService = FakeRaidService(),
    haptics: RecordingHaptics = RecordingHaptics(),
    undetermined: Boolean = false,
    askStore: MemoryAskStore = MemoryAskStore(),
    now: Instant = RaidTime.now,
): SignupController = SignupController(
    2,
    raid,
    SignupController.Dependencies(
        service = service,
        res = resources(),
        haptics = haptics,
        permission = NotificationPermission { undetermined },
        askStore = askStore,
        now = { now },
        scope = null,
        hold = 4.seconds,
    ),
)
