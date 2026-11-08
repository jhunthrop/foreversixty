package gg.foreversixty.app.home.flows

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.activity.Activity
import gg.foreversixty.app.activity.ActivityDay
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.bis.BisRange
import gg.foreversixty.app.gear.RunPick
import gg.foreversixty.app.home.HomeCharacter
import gg.foreversixty.app.home.HomePresenter
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemNumbers
import gg.foreversixty.app.model.ActivityRow
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.HomeSection
import gg.foreversixty.app.phrase
import java.math.BigDecimal
import java.math.RoundingMode
import java.time.Instant
import java.time.ZoneId

/** Everything Home's blocks are made of, gathered by the route. Null is "not loaded" or "none". */
data class HomeInputs(
    val sections: List<HomeSection>,
    val upgrades: CharacterHome?,
    val bestRun: RunPick?,
    val raidCard: RaidCardModel?,
    val raidLive: RaidLiveModel?,
    val activity: List<ActivityRow>,
    val unread: Int,
    /** The first raids' date for Find a guild (`9 Dec`), until it passes. */
    val firstRaids: String?,
)

/**
 * Composes Home from `/v1/home` `sections` (spec 2026-11-08 4.1): one block per section the app knows,
 * in the server's order; a section with nothing true to say is left out; an unknown type is skipped.
 */
class HomeComposer(private val res: Resources, private val now: Instant, private val zone: ZoneId) {
    fun head(character: HomeCharacter): HomeHead {
        val who = listOfNotNull(character.specName, character.className).joinToString(" ") { HomePresenter.words(it) }
        val level = character.level
        val second = when {
            level != null && level < BisRange.LAST_BAND -> toBand(level)
            else -> character.guild
        }
        return HomeHead(
            character.name,
            character.className,
            listOfNotNull(level?.let { res.getString(R.string.home_head_level, it, who) } ?: who.ifEmpty { null }, second),
        )
    }

    fun blocks(inputs: HomeInputs): List<HomeBlock> {
        val sections = inputs.sections.ifEmpty { DEFAULT }
        val hasCard = sections.any { it.type == NEXT_RAID && it.style == CARD } && inputs.raidCard != null
        return sections.mapNotNull { section ->
            when (section.type) {
                UPGRADES -> upgrades(inputs, section.limit ?: LIMIT, bestRun = !hasCard)
                NEXT_RAID -> inputs.raidCard?.let { if (section.style == CARD) HomeBlock.NextRaidCard(it) else HomeBlock.NextRaidLine(it) }
                RAID_LIVE -> inputs.raidLive?.let { HomeBlock.RaidLive(it) }
                RAID_SUMMARY -> lastNight(inputs, sections)
                ACTIVITY -> activity(inputs, section)
                GUILD_CTA -> HomeBlock.FindGuild(
                    inputs.firstRaids?.let { res.getString(R.string.home_findGuild_firstRaids, it) } ?: res.getString(R.string.home_findGuild_line),
                )
                NOTICE -> section.text?.let { HomeBlock.Notice(it, section.href) }
                else -> null
            }
        }
    }

    private fun upgrades(inputs: HomeInputs, limit: Int, bestRun: Boolean): HomeBlock? {
        val home = inputs.upgrades ?: return null
        val unit = home.unitWord.orEmpty()
        val rows = HomePresenter.presentationOrder(home.upgrades).take(limit).map { upgrade ->
            val pick = upgrade.pick
            val gain = upgrade.gain?.takeIf { upgrade.modelled }
            UpgradeLine(
                look = ItemLook(pick.id, pick.name, pick.quality, pick.icon, pick.itemLevel),
                source = "${upgrade.slotLabel} · ${pick.sourceLine}",
                gain = gain?.let(ItemNumbers::signedTenths) ?: upgrade.gainWords,
                unit = if (gain != null) unit else null,
                slot = upgrade.slot,
                wornId = upgrade.worn?.id,
                sourceLine = pick.sourceLine,
            )
        }
        if (rows.isEmpty()) return null
        val line = inputs.bestRun?.takeIf { bestRun }?.let { bestRunLine(it, unit) }
        return HomeBlock.Upgrades(rows, line, home.dataBuild)
    }

    /** `Best run: Blackrock Spire, Blackcrow +1.9 DPS, 8%`. */
    fun bestRunLine(pick: RunPick, unit: String): String {
        val gain = "${ItemNumbers.signedTenths(pick.gain)} $unit".trim()
        val chance = pick.source.chance?.let { BigDecimal.valueOf(it).setScale(0, RoundingMode.HALF_UP).toPlainString() + "%" }
        return if (chance != null) {
            res.getString(R.string.home_bestRun_chance, pick.source.name, pick.itemName, gain, chance)
        } else {
            res.getString(R.string.home_bestRun, pick.source.name, pick.itemName, gain)
        }
    }

    /** Last night: the rows the day-after preview leaves out (your loot, the report), from the last day. */
    private fun lastNight(inputs: HomeInputs, sections: List<HomeSection>): HomeBlock? {
        val kinds = sections.firstOrNull { it.type == ACTIVITY }?.excludeKinds.orEmpty().ifEmpty { LAST_NIGHT_KINDS }
        val since = now.minusSeconds(LAST_NIGHT_SECONDS)
        val rows = Activity.items(inputs.activity.filter { it.kind in kinds }).filter { it.at.isAfter(since) }
        if (rows.isEmpty()) return null
        val loot = rows.any { it.row.kind == LOOT }
        return HomeBlock.LastNight(rows.map { HomeActivityRow(it.row, it.row.body) }, simAgain = loot)
    }

    /** New for you: the newest rows of the section's group, past the kinds a block above shows. */
    private fun activity(inputs: HomeInputs, section: HomeSection): HomeBlock {
        val excluded = section.excludeKinds.orEmpty().toSet()
        val rows = Activity.items(inputs.activity)
            .filter { (section.group == null || it.row.group == section.group) && it.row.kind !in excluded }
            .take(section.limit ?: LIMIT)
            .map { HomeActivityRow(it.row, line(it.row.body, it.at)) }
        return HomeBlock.Activity(rows, inputs.unread)
    }

    /** `+21 vs your last · 3 h ago`; yesterday and earlier read as a day. */
    fun line(body: String, at: Instant): String {
        val today = now.atZone(zone).toLocalDate()
        val age = when (Activity.dayOf(at.atZone(zone).toLocalDate(), today)) {
            ActivityDay.Today -> RelativeAge.of(at, now).phrase(res)
            ActivityDay.Yesterday -> res.getString(R.string.home_yesterday)
            else -> RelativeAge.of(at, now).phrase(res)
        }
        return listOf(body, age).filter { it.isNotBlank() }.joinToString(" · ")
    }

    private fun toBand(level: Int): String {
        val next = (level / BisRange.BAND_STEP + 1) * BisRange.BAND_STEP
        val left = next - level
        return if (left == 1) res.getString(R.string.home_head_toBandOne, next) else res.getString(R.string.home_head_toBand, left, next)
    }

    companion object {
        const val UPGRADES = "upgrades"
        const val NEXT_RAID = "next_raid"
        const val RAID_LIVE = "raid_live"
        const val RAID_SUMMARY = "raid_summary"
        const val ACTIVITY = "activity"
        const val GUILD_CTA = "guild_cta"
        const val NOTICE = "notice"
        const val CARD = "card"
        private const val LOOT = "loot.awarded"
        private const val LIMIT = 3
        private const val LAST_NIGHT_SECONDS = 36L * 60 * 60

        /** What Last night shows when the section names no kinds (spec 4.1). */
        val LAST_NIGHT_KINDS = listOf(LOOT, "report.ready", "coach.ready", "ranking.best")

        /** Home before `/v1/home` answers: the any-other-day order. */
        val DEFAULT = listOf(HomeSection("head"), HomeSection(UPGRADES, limit = LIMIT), HomeSection(NEXT_RAID, style = "line"), HomeSection(ACTIVITY, limit = LIMIT))
    }
}
