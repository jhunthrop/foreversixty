package gg.foreversixty.app.model

import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable

/**
 * GET /v1/me: who you are, what you own and what you are entitled to.
 * Mirrors api/internal/auth/me_view.go `Me`. Fields the apps do not read yet (seat views,
 * Discord channel detail) are left out; [ApiJson] ignores them.
 */
@Serializable
data class Me(
    val user: User,
    val characters: List<Character> = emptyList(),
    val guilds: List<Guild> = emptyList(),
    val entitlements: Entitlements,
    @SerialName("bnet_imported_at") val bnetImportedAt: String? = null,
    @SerialName("main_character_key") val mainCharacterKey: String? = null,
    @SerialName("finished_sims") val finishedSims: List<FinishedSim> = emptyList(),
    @SerialName("server_sims") val serverSims: ServerSims? = null,
    val notifications: Notifications? = null,
    @SerialName("discord_link_available") val discordLinkAvailable: Boolean = false,
    @SerialName("discord_link") val discordLink: DiscordLink? = null,
    @SerialName("discord_link_suggested") val discordLinkSuggested: String? = null,
) {
    /** The account's chosen main, else its first character; null for an account with none. */
    val mainCharacter: Character?
        get() = characters.firstOrNull { it.key == mainCharacterKey } ?: characters.firstOrNull()
}

/** api/internal/auth/store.go `User`. An email-only account has no battletag. */
@Serializable
data class User(
    val id: Long,
    val battletag: String? = null,
    val email: String? = null,
    val role: String,
    val anonymize: Boolean = false,
)

/** api/internal/auth/handler.go `FinishedSim`: a server run that finished while you were away. */
@Serializable
data class FinishedSim(
    @SerialName("sim_id") val simId: String,
    val spec: String,
    val title: String,
    val dps: Double,
    @SerialName("finished_at") val finishedAt: String,
)

/** api/internal/auth/handler.go `ServerSimsView`. */
@Serializable
data class ServerSims(
    @SerialName("runs_this_month") val runsThisMonth: Int,
)

/** api/internal/auth/me_view.go `NotificationsBlock`, keyed by channel. */
@Serializable
data class Notifications(
    val connect: Boolean = false,
)

/** api/internal/auth/me_view.go `DiscordLinkBlock`. */
@Serializable
data class DiscordLink(
    @SerialName("display_name") val displayName: String,
)
