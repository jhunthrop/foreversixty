package gg.foreversixty.app.home.flows

import gg.foreversixty.app.flows.ClockPart
import gg.foreversixty.app.flows.TallyItem
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.model.ActivityRow
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.logic.StatusLine

/** Home's head (spec 2026-11-08 4.1): the class crest, the name in the class colour, the fact line. */
data class HomeHead(val name: String, val className: String?, val facts: List<String>)

/** One Next upgrades row: the item, where it comes from and its gain (`+14.5` `DPS`, or the server's words). */
data class UpgradeLine(
    val look: ItemLook,
    val source: String,
    val gain: String?,
    val unit: String?,
    val slot: String,
    val wornId: Int?,
    val sourceLine: String,
)

/** The next raid as Home and the Raid tab draw it, every word prepared. */
data class RaidCardModel(
    val raidId: Long,
    val icon: String?,
    /** `Olympus XXVII · Onyxia's Lair tonight`. */
    val eyebrow: String,
    val clock: List<ClockPart>,
    val clockLabel: String,
    /** `Wed 6 Jan, 8:00 pm CST · answer by 6:00 pm CST`. */
    val aside: String,
    /** The line style's title, `Onyxia's Lair · Wed 6 Jan, 8:00 pm CST`. */
    val title: String,
    /** The line style's sub line: the counts and who has not answered; once answered the counts and who answered. */
    val sub: String?,
    /**
     * `You have not answered, as Thandrel, Damage` (the answer alone with one character in the guild, E7);
     * once answered only `As Thandrel, Damage`, null with one character: the segmented control shows it (AS5).
     */
    val me: String?,
    val answer: SignupState?,
    val lateOpen: Boolean,
    val tally: List<TallyItem>,
    val ready: ReadyLine?,
    val officer: OfficerItems?,
    /** The card's title (round 5): the zone, `Onyxia's Lair`. */
    val zone: String = "",
    /** The line under it: `Olympus XXVII · Wed 6 Jan, 8:00 pm CST`. */
    val place: String = "",
    /** The Raid tab's one-line head (spec 4.2): `8:00 pm CST · pull in 2 h 48 min`, or the date and time a day or more out. */
    val headLine: String = "",
    /** `Invites 7:45 pm · 40` under the head line; null without an invite time or size. */
    val facts: String? = null,
    /** `As Thandrel, Damage` under the head line, only with more than one character in the guild (E7). */
    val who: String? = null,
    /** How late a Late answer said it would be (15, 30, 60); the selected chip. */
    val lateMinutes: Int? = null,
    /** The answers' live state from the raid's controller: enabled, the status line, Try again. */
    val control: AnswerControl = AnswerControl(),
)

/**
 * The answers' live state under a raid card (spec F4.3, R16): whether they take a press (offline,
 * locked or a character still waiting do not), the status line under them (`Saving…`, `Saved.`, an
 * error, the offline or locked sentence) and Try again after a failed write.
 */
data class AnswerControl(val enabled: Boolean = true, val status: StatusLine? = null, val canRetry: Boolean = false)

/** `2 things before you pull` and what they are, with the export's age. */
data class ReadyLine(val title: String, val sub: String)

/** An officer's open items on the raid card (F3). */
data class OfficerItems(val unanswered: String?, val short: String?, val applicants: String?, val applicantLine: String?)

/** The raid live card. */
data class RaidLiveModel(val raidId: Long, val eyebrow: String, val line: String)

/** An Activity row as Home draws it: its line with when (`+21 vs your last · 3 h ago`); [age] the when alone, after a styled body. */
data class HomeActivityRow(val row: ActivityRow, val line: String, val age: String? = null)

/** One block of the signed-in Home, in the order `/v1/home` `sections` gives. */
sealed interface HomeBlock {
    /** Next upgrades: the rows, Best run as a row, the meta (`13 upgrades · +157.6 DPS`) and the synced stamp. */
    data class Upgrades(
        val rows: List<UpgradeLine>,
        val bestRun: UpgradeLine?,
        val build: String,
        val meta: String? = null,
        val synced: String? = null,
        /** The sync is over a week old: the stamp in the warning tone (row 12). */
        val syncedStale: Boolean = false,
    ) : HomeBlock
    data class NextRaidCard(val raid: RaidCardModel) : HomeBlock
    data class NextRaidLine(val raid: RaidCardModel) : HomeBlock
    data class RaidLive(val live: RaidLiveModel) : HomeBlock
    data class LastNight(val rows: List<HomeActivityRow>, val simAgain: Boolean) : HomeBlock
    data class Activity(val rows: List<HomeActivityRow>, val unread: Int) : HomeBlock
    data class FindGuild(val line: String) : HomeBlock
    data class Notice(val text: String, val href: String?) : HomeBlock
}

/** What Home shows. */
sealed interface HomeView {
    data class Loading(val head: HomeHead) : HomeView
    data class Failed(val head: HomeHead) : HomeView

    /** Signed in, Battle.net lists no character yet (F8). */
    data object NoCharacter : HomeView
    data class Blocks(val head: HomeHead, val blocks: List<HomeBlock>, val iconBuild: String?) : HomeView
}

/** What a tap on Home does; the route binds each to its screen. */
data class HomeActions(
    val onGear: () -> Unit = {},
    val onItem: (UpgradeLine) -> Unit = {},
    val onAnswer: (Long, SignupState) -> Unit = { _, _ -> },
    val onLate: (Long, Int) -> Unit = { _, _ -> },
    /** Try again under a failed write on the raid card. */
    val onRetryAnswer: (Long) -> Unit = {},
    val onRaid: (Long) -> Unit = {},
    val onRaidNight: (Long) -> Unit = {},
    val onReady: (Long) -> Unit = {},
    val onPreviewMention: (Long) -> Unit = {},
    val onApplicants: () -> Unit = {},
    val onActivityRow: (ActivityRow) -> Unit = {},
    val onAllActivity: () -> Unit = {},
    val onSimAgain: () -> Unit = {},
    val onFindGuild: () -> Unit = {},
    val onNotice: (String) -> Unit = {},
    val onRetry: () -> Unit = {},
    val onCheckAgain: () -> Unit = {},
)
