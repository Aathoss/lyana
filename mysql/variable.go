package mysql

var (
	member   sMember
	msgcount sMessageCount
	sanction sSanction
	rule     sRule
	count    int
)

// Member stock les informations de la base de données
type sMember struct {
	id             int
	uid_discord    string
	player_mc      string
	date_whitelist int64
	inactif        int64
	notif          int
}

type sMessageCount struct {
	id          int
	uid_discord string
	count_msg   int
}

type sSanction struct {
	id           int
	uid          string
	id_message   string
	id_msg_notif string
}

type sRule struct {
	id        int
	uid       string
	timestamp int64
}
