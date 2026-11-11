import Foundation
import Models

/// An officer's answers from a push without opening the app (guild-officer N1, N2): Accept or
/// Decline an applicant, Mention the short role in Discord. Accept and Decline return true when the
/// API took it; Mention returns what the ping did, nil when it failed.
struct OfficerPushActions: Sendable {
    var accept: @Sendable (_ guild: Int64, _ application: Int64) async -> Bool = { _, _ in false }
    var decline: @Sendable (_ guild: Int64, _ application: Int64) async -> Bool = { _, _ in false }
    var mention: @Sendable (_ guild: Int64, _ raid: Int64, _ role: String) async -> RaidPingResult? = { _, _, _ in nil }
}

/// What replaces an officer push once its action ran (round 5, row 3; Android's
/// RaidShortPushAnswerer and ApplicantPushAnswerer): `Mentioned 3 healers in Discord.` over the
/// raid, `Application accepted` over the applicant; a failure says what failed and how to retry
/// (round 7, row 2): `Could not mention the healers in Discord` over `Onyxia's Lair. Open the raid to
/// try again.`, `Could not accept Mirelle` over `Mirelle, Holy Priest 60. Open Forever Sixty to try again.`
@MainActor
enum OfficerPushResult {
    /// `title` is the delivered push's title (the raid, the applicant).
    static func mention(_ result: RaidPingResult?, request: RaidMentionRequest, title: String) -> RaidPushOutcome.Confirmation {
        guard let result else { return mentionFailed(who: request.who, raid: title) }
        return .init(title: NotAnsweredModel.status(result, role: request.role), body: title)
    }

    static func applicant(done: Bool, accepted: Bool, applicant: ApplicantPushName) -> RaidPushOutcome.Confirmation {
        guard done else { return applicantFailed(accepted: accepted, applicant: applicant) }
        return .init(title: accepted ? Copy.pushApplicantAccepted : Copy.pushApplicantDeclined, body: applicant.pushTitle)
    }

    /// An action whose payload lacks what it needs (the server sends no Mention inside the cooldown).
    static func unreadable(action: String, title: String, applicant: ApplicantPushName) -> RaidPushOutcome.Confirmation {
        switch action {
        case RaidPushCategory.actionMention: mentionFailed(who: nil, raid: title)
        case RaidPushCategory.actionDecline: applicantFailed(accepted: false, applicant: applicant)
        default: applicantFailed(accepted: true, applicant: applicant)
        }
    }

    /// `who` is the short role as the Mention button names it (`healers`, the push's `mention_who`); a
    /// Mention without it keeps the plain failure title.
    private static func mentionFailed(who: String?, raid: String) -> RaidPushOutcome.Confirmation {
        let title = who.map(Copy.pushShortFailed(who:)) ?? Copy.commonFailed
        return .init(title: title, body: Copy.pushShortRetry(raid: raid))
    }

    private static func applicantFailed(accepted: Bool, applicant: ApplicantPushName) -> RaidPushOutcome.Confirmation {
        let title = applicant.name.map { accepted ? Copy.pushApplicantAcceptFailed(name: $0) : Copy.pushApplicantDeclineFailed(name: $0) }
        return .init(title: title ?? Copy.commonFailed, body: Copy.pushApplicantFailed(applicant: applicant.line))
    }
}

/// Who an applicant push is about, from its data (`applicant_name` `Mirelle`, `applicant`
/// `Mirelle, Holy Priest 60`); a push without them falls back to its own title.
struct ApplicantPushName: Equatable, Sendable {
    let name: String?
    let line: String
    /// The delivered push's title, under a done answer.
    let pushTitle: String

    static let nameKey = "applicant_name"
    static let lineKey = "applicant"

    init(data: [String: String], pushTitle: String) {
        name = data[Self.nameKey].flatMap { $0.isEmpty ? nil : $0 }
        line = data[Self.lineKey].flatMap { $0.isEmpty ? nil : $0 } ?? pushTitle
        self.pushTitle = pushTitle
    }
}

/// An officer push's action, read from the notification before it crosses to the main actor.
enum OfficerPushAnswer: Equatable, Sendable {
    case accept(guild: Int64, application: Int64)
    case decline(guild: Int64, application: Int64)
    case mention(RaidMentionRequest)
    /// An officer action whose payload lacks what it needs.
    case unreadable

    /// Nil for an action that is not an officer's.
    init?(action: String, userInfo: [AnyHashable: Any]) {
        switch action {
        case RaidPushCategory.actionAccept, RaidPushCategory.actionDecline:
            guard let found = PushLanding.applicant(data: PushLanding.data(userInfo: userInfo)) else { self = .unreadable; return }
            self = action == RaidPushCategory.actionAccept
                ? .accept(guild: found.guild, application: found.application)
                : .decline(guild: found.guild, application: found.application)
        case RaidPushCategory.actionMention:
            self = RaidMentionRequest(userInfo: userInfo).map(OfficerPushAnswer.mention) ?? .unreadable
        default:
            return nil
        }
    }
}

/// What `raid.short`'s Mention action posts: the raid and the short role its `actions` entry names.
struct RaidMentionRequest: Equatable, Sendable {
    let guild: Int64
    let raid: Int64
    let role: String
    /// The role as the Mention button names it (`healers`), from `data.mention_who`.
    let who: String?

    static let actionsKey = "actions"
    static let whoKey = "mention_who"
    static let idKey = "id"
    static let bodyKey = "body"
    static let roleKey = "role"

    /// Read from the APNs userInfo: `data.guild_id`, `data.raid_id`, and the `raid_mention_short`
    /// action's `body.role`; nil when any is missing (the server sends no action inside the cooldown).
    init?(userInfo: [AnyHashable: Any]) {
        let data = PushLanding.data(userInfo: userInfo)
        let actions = userInfo[Self.actionsKey] as? [[String: Any]] ?? []
        let action = actions.first { $0[Self.idKey] as? String == RaidPushCategory.actionMention }
        guard let guild = data["guild_id"].flatMap(Int64.init), let raid = data["raid_id"].flatMap(Int64.init),
              let body = action?[Self.bodyKey] as? [String: Any], let role = body[Self.roleKey] as? String, !role.isEmpty
        else { return nil }
        self.guild = guild
        self.raid = raid
        self.role = role
        who = data[Self.whoKey].flatMap { $0.isEmpty ? nil : $0 }
    }
}
