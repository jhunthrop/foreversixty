import Foundation

/// Fixture mode only: a sign-up write (`PUT …/signups/{region}/{ruleset}/{name}`) answers its canned raid with the
/// tapped answer in place, so the fixture app shows the answer moving (In to Late and back), as the server does.
/// The canned file stays the contract's; only the answering character's `state` and `late_minutes` follow the body.
enum SignupEcho {
    private static let stateKey = "state"
    private static let lateKey = "late_minutes"
    private static let keyField = "character_key"

    /// The canned answer with the request's state written onto the character's sign-up and viewer row; nil when the
    /// request is not a sign-up or either side does not parse.
    static func apply(_ data: Data, request body: Data, params: [String: String]) -> Data? {
        guard let region = params["region"], let ruleset = params["ruleset"], let name = params["name"],
              let sent = try? JSONSerialization.jsonObject(with: body) as? [String: Any],
              let state = sent[stateKey] as? String,
              var root = try? JSONSerialization.jsonObject(with: data) as? [String: Any],
              var payload = root["data"] as? [String: Any],
              var raid = payload["raid"] as? [String: Any]
        else { return nil }
        let key = "\(region)/\(ruleset)/\(name)"
        let late = sent[lateKey] ?? NSNull()
        func echoed(_ rows: Any?) -> Any? {
            (rows as? [[String: Any]])?.map { row in
                guard row[keyField] as? String == key else { return row }
                var row = row
                row[stateKey] = state
                if row.keys.contains(lateKey) { row[lateKey] = late }
                return row
            }
        }
        raid["signups"] = echoed(raid["signups"]) ?? raid["signups"]
        if var viewer = raid["viewer"] as? [String: Any] {
            viewer["characters"] = echoed(viewer["characters"]) ?? viewer["characters"]
            raid["viewer"] = viewer
        }
        payload["raid"] = raid
        root["data"] = payload
        return try? JSONSerialization.data(withJSONObject: root)
    }
}
