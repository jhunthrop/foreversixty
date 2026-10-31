package gg.foreversixty.app.raids.logic

import android.content.res.Resources
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.R
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.AccountGuild
import gg.foreversixty.app.model.CharacterGuild
import gg.foreversixty.app.model.GuildRaids
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.RaidNightRow
import gg.foreversixty.app.model.RaidNightSummary
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.model.mainCharacter
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import java.time.Instant
import java.util.Locale
import kotlinx.coroutines.delay
import kotlin.time.Duration.Companion.seconds

/** The guild head (spec F4.1 item 1). */
data class GuildHead(
    val id: Long,
    val name: String,
    val crestUrl: String?,
    val faction: String?,
    /** `Horde · US Normal · 38 on the roster`, or null when the account's guild row is missing. */
    val line: String?,
    /** The guild's page on the site, when the account row has its slug. */
    val webUrl: String?,
) {
    companion object {
        fun of(character: CharacterGuild, guild: AccountGuild?, res: Resources) = GuildHead(
            id = character.id,
            name = guild?.name ?: character.name,
            crestUrl = guild?.crestUrl,
            faction = guild?.faction,
            line = guild?.let {
                res.getString(
                    R.string.raids_guildLine,
                    it.faction?.replaceFirstChar { c -> c.titlecase(Locale.UK) }.orEmpty(),
                    it.region.uppercase(Locale.UK),
                    it.rulesetLabel ?: it.ruleset.replaceFirstChar { c -> c.titlecase(Locale.UK) },
                    it.rosterCount ?: it.verifiedCount,
                )
            },
            webUrl = guild?.slug?.let { "${RaidLinks.SITE}/guild/${guild.region}/${guild.ruleset}/$it" },
        )
    }
}

/** What the Raids tab draws (spec F4.1 and F4.13 to F4.17). */
sealed interface RaidsTabState {
    data object Loading : RaidsTabState
    data object SignedOut : RaidsTabState
    data class NoGuild(val name: String) : RaidsTabState
    data class Waiting(val head: GuildHead, val character: AccountCharacter) : RaidsTabState
    data class Loaded(val content: RaidsTabContent) : RaidsTabState
    data object Failed : RaidsTabState
}

data class RaidsTabContent(
    val head: GuildHead,
    val character: AccountCharacter,
    val planned: List<RaidView>,
    val nights: List<RaidNightRow>,
    /** The officer's waiting count; null for anyone who cannot approve. */
    val waiting: Int?,
    /** Set while the cached answer shows because a refresh could not connect. */
    val offlineSince: Instant?,
) {
    fun replacing(raid: RaidView) = copy(planned = planned.map { if (it.id == raid.id) raid else it })
}

/** Loads `/v1/me` for the current character's guild, then the guild's raids; the cache draws first. */
class RaidsTabModel(
    private val service: RaidService,
    private val cache: RaidCache,
    private val res: Resources,
    private val now: () -> Instant,
    initial: RaidsTabState = RaidsTabState.Loading,
) {
    var state: RaidsTabState by mutableStateOf(initial)
        private set

    /** The character the tab is for; null follows the account's main (lane F1's selector sets it). */
    var currentKey: String? = null

    suspend fun load() {
        val me = when (val result = service.me()) {
            is ApiResult.Success -> result.data.also { cache.write(ME, Me.serializer(), it, now()) }
            is ApiResult.Failure -> {
                val error = result.error
                if (error is ApiError.Server && error.isUnauthorized) {
                    state = RaidsTabState.SignedOut
                    return
                }
                cache.read(ME, Me.serializer())?.value ?: run {
                    state = RaidsTabState.Failed
                    return
                }
            }
        }
        load(me)
    }

    private suspend fun load(me: Me) {
        val character = me.characters.firstOrNull { it.key == currentKey } ?: me.mainCharacter
        if (character == null) {
            state = RaidsTabState.NoGuild("")
            return
        }
        val membership = character.guild ?: run {
            state = RaidsTabState.NoGuild(character.name)
            return
        }
        val head = GuildHead.of(membership, me.guilds.firstOrNull { it.id == membership.id }, res)
        if (!membership.verified) {
            state = RaidsTabState.Waiting(head, character)
            return
        }
        val key = "guild-${membership.id}-raids"
        if (state == RaidsTabState.Loading) {
            cache.read(key, GuildRaids.serializer())?.let { state = RaidsTabState.Loaded(content(head, character, it.value, null)) }
        }
        state = when (val result = service.raids(membership.id)) {
            is ApiResult.Success -> {
                cache.write(key, GuildRaids.serializer(), result.data, now())
                RaidsTabState.Loaded(content(head, character, result.data, null))
            }
            is ApiResult.Failure -> {
                val cached = cache.read(key, GuildRaids.serializer())
                if (SignupRefusal.isOffline(result.error) && cached != null) {
                    RaidsTabState.Loaded(content(head, character, cached.value, cached.savedAt))
                } else {
                    RaidsTabState.Failed
                }
            }
        }
    }

    private fun content(head: GuildHead, character: AccountCharacter, raids: GuildRaids, offlineSince: Instant?) =
        RaidsTabContent(
            head = head,
            character = character,
            planned = raids.planned.orEmpty(),
            nights = raids.rows.filter { it.raidId != null }.take(NIGHT_ROWS),
            waiting = raids.waitingForApproval?.takeIf { it > 0 },
            offlineSince = offlineSince,
        )

    /** A sign-up on the card or the raid page replaced this raid. */
    fun replace(raid: RaidView) {
        (state as? RaidsTabState.Loaded)?.let { state = RaidsTabState.Loaded(it.content.replacing(raid)) }
    }

    companion object {
        /** The raid nights row count (spec F4.1 item 4). */
        const val NIGHT_ROWS = 3
        private const val ME = "me"
    }
}

/** What a deep link asks the raid page to do on open (links.json: `spec=ask`, `state=in`). */
data class RaidOpenAction(val state: SignupState)

/** Loads one raid for the raid page (cache first), and keeps it live in its last 24 h. */
class RaidPageModel(
    val guild: Long,
    val raidId: Long,
    private val deps: SignupController.Dependencies,
    private val cache: RaidCache,
    private var openAction: RaidOpenAction? = null,
    private val onChange: (RaidView) -> Unit = {},
) {
    sealed interface State {
        data object Loading : State
        data class Loaded(val controller: SignupController) : State
        data object Failed : State
        data object NotAvailable : State
    }

    var state: State by mutableStateOf(State.Loading)
        private set
    var offlineSince: Instant? by mutableStateOf(null)
        private set

    private val key get() = "guild-$guild-raid-$raidId"

    /** A page built straight from a controller (previews and screenshot tests). */
    constructor(controller: SignupController, deps: SignupController.Dependencies, offlineSince: Instant? = null) :
        this(controller.guild, controller.raid.id, deps, NoRaidCache) {
        state = State.Loaded(controller)
        this.offlineSince = offlineSince
        controller.offline = offlineSince != null
    }

    suspend fun load() {
        val loaded = state as? State.Loaded
        if (loaded != null) {
            refresh(loaded.controller)
            return
        }
        cache.read(key, RaidView.serializer())?.let { show(it.value) }
        when (val result = deps.service.raid(guild, raidId)) {
            is ApiResult.Success -> {
                cache.write(key, RaidView.serializer(), result.data, deps.now())
                offlineSince = null
                show(result.data)
                runOpenAction()
            }
            is ApiResult.Failure -> fail(result.error)
        }
    }

    private suspend fun refresh(controller: SignupController) {
        controller.refresh()
        if (controller.offline) {
            offlineSince = offlineSince ?: cache.read(key, RaidView.serializer())?.savedAt ?: deps.now()
        } else {
            offlineSince = null
            cache.write(key, RaidView.serializer(), controller.raid, deps.now())
        }
    }

    private fun show(view: RaidView) {
        (state as? State.Loaded)?.let {
            it.controller.replace(view)
            return
        }
        val controller = SignupController(guild, view, deps)
        controller.onRaidChanged = { changed ->
            cache.write(key, RaidView.serializer(), changed, deps.now())
            onChange(changed)
        }
        state = State.Loaded(controller)
    }

    private fun fail(error: ApiError) {
        val loaded = state as? State.Loaded
        if (loaded != null) {
            if (SignupRefusal.isOffline(error)) {
                loaded.controller.offline = true
                offlineSince = cache.read(key, RaidView.serializer())?.savedAt ?: deps.now()
            }
            return
        }
        state = if (error is ApiError.Server && error.status == HTTP_NOT_FOUND) State.NotAvailable else State.Failed
    }

    private suspend fun runOpenAction() {
        val action = openAction ?: return
        val controller = (state as? State.Loaded)?.controller ?: return
        val character = controller.raid.viewer.characters.firstOrNull { it.verified } ?: return
        openAction = null
        controller.press(character.characterKey, action.state)
    }

    /** Refetches every 30 s while the raid starts within 24 h; the caller's scope cancels it. */
    suspend fun stayLive() {
        while (true) {
            delay(LIVE_PERIOD)
            val controller = (state as? State.Loaded)?.controller ?: continue
            val facts = controller.facts
            val untilStart = java.time.Duration.between(deps.now(), facts.startsAt)
            if (untilStart < LIVE_WINDOW && facts.isSignable) refresh(controller)
        }
    }

    private companion object {
        const val HTTP_NOT_FOUND = 404
        val LIVE_PERIOD = 30.seconds
        val LIVE_WINDOW: java.time.Duration = java.time.Duration.ofHours(24)
    }
}

/** Loads `GET /v1/me/raid-nights/{raid}` for Your raid night (spec F4.22). */
class RaidNightModel(
    val raidId: Long,
    private val service: RaidService,
    private val cache: RaidCache,
    private val now: () -> Instant,
    initial: State = State.Loading,
) {
    sealed interface State {
        data object Loading : State
        data class Loaded(val night: RaidNightSummary) : State
        data object Failed : State
    }

    var state: State by mutableStateOf(initial)
        private set
    var offlineSince: Instant? by mutableStateOf(null)
        private set

    private val key get() = "night-$raidId"

    suspend fun load() {
        if (state == State.Loading) cache.read(key, RaidNightSummary.serializer())?.let { state = State.Loaded(it.value) }
        when (val result = service.night(raidId)) {
            is ApiResult.Success -> {
                cache.write(key, RaidNightSummary.serializer(), result.data, now())
                offlineSince = null
                state = State.Loaded(result.data)
            }
            is ApiResult.Failure -> {
                val cached = cache.read(key, RaidNightSummary.serializer())
                when {
                    SignupRefusal.isOffline(result.error) && cached != null -> {
                        offlineSince = cached.savedAt
                        state = State.Loaded(cached.value)
                    }
                    state is State.Loaded -> Unit
                    else -> state = State.Failed
                }
            }
        }
    }
}

/** The links the raid screens open and share. */
object RaidLinks {
    const val SITE = "https://foreversixty.gg"

    /** The raid's https app link (links.json); the site redirects it without the app. */
    fun share(guild: Long, raid: Long): String = gg.foreversixty.app.navigation.DeepLink.Raid(guild, raid).url()

    fun report(id: String): String = "$SITE/reports/$id"
}
