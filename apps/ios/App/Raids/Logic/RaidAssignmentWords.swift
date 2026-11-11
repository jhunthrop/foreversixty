import Foundation
import Models

/// One job the viewer holds (spec 2026-11-10-guild-tab.md 2.3a): `Whelps, west cave with Kaelra and
/// Thalrin`, then `Onyxia · <the slot's purpose>`.
struct RaidJobLine: Equatable, Identifiable {
    var id: String { encounter + "/" + slot }
    let encounter: String
    let slot: String
    let title: String
    let line: String
}

/// One encounter's filled slots, read only, for an officer.
struct AssignmentEncounterLine: Equatable, Identifiable {
    struct Slot: Equatable, Identifiable {
        var id: String { key }
        let key: String
        let label: String
        let names: [RaidAssignedName]
        /// `2 of 3`.
        let filled: String
    }

    var id: Int64 { encounterID }
    let encounterID: Int64
    let title: String
    let slots: [Slot]
}

/// The raid's assignments (`RaidView.assignments`) in the web planner's and the addon's words.
enum RaidAssignmentWords {
    /// The viewer's jobs in the raid: every slot one of their characters holds, the partners named
    /// as the addon's `%s with %s: %s` names them. Empty when they hold none.
    static func jobs(_ raid: RaidView?) -> [RaidJobLine] {
        guard let encounters = raid?.assignments?.encounters else { return [] }
        return encounters.flatMap { encounter in
            encounter.slots.filter { $0.names.contains(where: \.own) }.map { slot in
                let partners = slot.names.filter { !$0.own }.map(\.name)
                let title = partners.isEmpty
                    ? slot.label
                    : Copy.sharedAssignJobWith(label: slot.label, names: RaidFlowPresenter.list(partners))
                return RaidJobLine(
                    encounter: encounter.name, slot: slot.key, title: title,
                    line: Copy.guildTabJobLine(boss: encounter.name, purpose: slot.purpose)
                )
            }
        }
    }

    /// Every encounter with a pick, every slot of its template in order (an empty one reads `0 of 2`), as the
    /// web planner lists them; the encounter titled by its boss.
    static func encounters(_ raid: RaidView?) -> [AssignmentEncounterLine] {
        guard let encounters = raid?.assignments?.encounters else { return [] }
        return encounters.compactMap { encounter in
            guard encounter.slots.contains(where: { !$0.names.isEmpty }) else { return nil }
            let slots = encounter.slots.map { slot in
                AssignmentEncounterLine.Slot(
                    key: slot.key, label: slot.label, names: slot.names,
                    filled: Copy.sharedAssignFilled(n: slot.names.count, count: slot.count)
                )
            }
            return AssignmentEncounterLine(encounterID: encounter.encounterId, title: encounter.name, slots: slots)
        }
    }
}
