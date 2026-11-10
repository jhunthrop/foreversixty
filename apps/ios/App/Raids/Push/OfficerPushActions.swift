import Foundation

/// An officer's answers from a push without opening the app (guild-officer N1, N2): Accept or
/// Decline an applicant, Mention the short role in Discord. Each returns true when the API took it;
/// a refusal or failure opens the push's landing instead.
struct OfficerPushActions: Sendable {
    var accept: @Sendable (_ guild: Int64, _ application: Int64) async -> Bool = { _, _ in false }
    var decline: @Sendable (_ guild: Int64, _ application: Int64) async -> Bool = { _, _ in false }
    var mention: @Sendable (_ guild: Int64, _ raid: Int64, _ role: String) async -> Bool = { _, _, _ in false }
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

    static let actionsKey = "actions"
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
    }
}
