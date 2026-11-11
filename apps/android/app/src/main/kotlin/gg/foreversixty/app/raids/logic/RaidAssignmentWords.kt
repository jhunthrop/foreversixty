package gg.foreversixty.app.raids.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.model.RaidAssignedName
import gg.foreversixty.app.model.RaidView

/** One job the viewer holds (spec 2026-11-10-guild-tab.md 2.3a): the slot with its partners, then the boss and the purpose. */
data class RaidJobLine(val encounter: String, val slot: String, val title: String, val line: String)

/** One encounter's filled slots, read only, for an officer. */
data class AssignmentEncounterLine(val encounterId: Long, val title: String, val slots: List<Slot>) {
    data class Slot(val key: String, val label: String, val names: List<RaidAssignedName>, val filled: String)
}

/** The raid's assignments (`RaidView.assignments`) in the web planner's and the addon's words. */
class RaidAssignmentWords(private val res: Resources) {
    /** The viewer's jobs: every slot one of their characters holds, the partners named as the addon names them. */
    fun jobs(raid: RaidView?): List<RaidJobLine> =
        raid?.assignments?.encounters.orEmpty().flatMap { encounter ->
            encounter.slots.filter { slot -> slot.names.any { it.own } }.map { slot ->
                val partners = slot.names.filterNot { it.own }.map { it.name }
                val title = if (partners.isEmpty()) slot.label else res.getString(R.string.shared_assign_jobWith, slot.label, list(partners))
                RaidJobLine(encounter.name, slot.key, title, res.getString(R.string.guild_tab_jobLine, encounter.name, slot.purpose))
            }
        }

    /** Every encounter with a pick, its filled slots in the template's order. */
    fun encounters(raid: RaidView?): List<AssignmentEncounterLine> =
        raid?.assignments?.encounters.orEmpty().mapNotNull { encounter ->
            val slots = encounter.slots.filter { it.names.isNotEmpty() }.map {
                AssignmentEncounterLine.Slot(it.key, it.label, it.names, res.getString(R.string.shared_assign_filled, it.names.size, it.count))
            }
            slots.takeIf { it.isNotEmpty() }?.let {
                AssignmentEncounterLine(encounter.encounterId, res.getString(R.string.shared_assign_encounterTitle, encounter.name), it)
            }
        }

    private fun list(names: List<String>): String =
        if (names.size < 2) names.firstOrNull().orEmpty() else res.getString(R.string.raid_list_and, names.dropLast(1).joinToString(", "), names.last())
}
