package model

type User struct {
	Alias  string `json:"alias,omitempty"`
	Server string `json:"server,omitempty"`
	Image  string `json:"image,omitempty"`
	Level  int    `json:"level,omitempty"`
	Rank   struct {
		Solo struct {
			Rank    string `json:"rank,omitempty"`
			Image   string `json:"image,omitempty"`
			Lp      int    `json:"lp,omitempty"`
			Win     int    `json:"win,omitempty"`
			Lose    int    `json:"lose,omitempty"`
			Winrate int    `json:"winrate,omitempty"`
			Promos  string `json:"promos,omitempty"`
		} `json:"solo,omitempty"`
		Flex struct {
			Rank    string `json:"rank,omitempty"`
			Image   string `json:"image,omitempty"`
			Lp      int    `json:"lp,omitempty"`
			Win     int    `json:"win,omitempty"`
			Lose    int    `json:"lose,omitempty"`
			Winrate int    `json:"winrate,omitempty"`
			Promos  string `json:"promos,omitempty"`
		} `json:"flex,omitempty"`
		Arena struct {
			Rank    string `json:"rank,omitempty"`
			Image   string `json:"image,omitempty"`
			Lp      int    `json:"lp,omitempty"`
			Win     int    `json:"win,omitempty"`
			Lose    int    `json:"lose,omitempty"`
			Winrate int    `json:"winrate,omitempty"`
			Promos  string `json:"promos,omitempty"`
		} `json:"arena,omitempty"`
	} `json:"rank,omitempty"`
}

type Stats struct {
	GamesUsed []string `json:"gamesUsed,omitempty"`
	Friends   []struct {
		Name  string `json:"name,omitempty"`
		Games int    `json:"games,omitempty"`
		Wins  int    `json:"wins,omitempty"`
	} `json:"friends,omitempty"`
	StatsByChamp []struct {
		ChampionName      string  `json:"championName,omitempty"`
		Games             int     `json:"games,omitempty"`
		Wins              int     `json:"wins,omitempty"`
		Kda               float64 `json:"kda,omitempty"`
		GoldMin           float64 `json:"goldMin,omitempty"`
		CsMin             float64 `json:"csMin,omitempty"`
		VisionMin         float64 `json:"visionMin,omitempty"`
		KillParticipation float64 `json:"killParticipation,omitempty"`
		DamageDealt       float64 `json:"damageDealt,omitempty"`
		DamageTaken       float64 `json:"damageTaken,omitempty"`
	} `json:"statsByChamp,omitempty"`
	StatsByPosition []struct {
		Games    int    `json:"games,omitempty"`
		Wins     int    `json:"wins,omitempty"`
		Position string `json:"position,omitempty"`
	} `json:"statsByPosition,omitempty"`
	Records struct {
		Kda struct {
			Value        float64 `json:"value,omitempty"`
			MatchID      string  `json:"matchId,omitempty"`
			ChampionName string  `json:"championName,omitempty"`
			GameMode     string  `json:"gameMode,omitempty"`
			GameCreation int64   `json:"gameCreation,omitempty"`
			GameDuration int     `json:"gameDuration,omitempty"`
		} `json:"kda,omitempty"`
		Kills struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"kills,omitempty"`
		Deaths struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"deaths,omitempty"`
		Assists struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"assists,omitempty"`
		Gold struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"gold,omitempty"`
		GoldPerMin struct {
			Value        float64 `json:"value,omitempty"`
			MatchID      string  `json:"matchId,omitempty"`
			ChampionName string  `json:"championName,omitempty"`
			GameMode     string  `json:"gameMode,omitempty"`
			GameCreation int64   `json:"gameCreation,omitempty"`
			GameDuration int     `json:"gameDuration,omitempty"`
		} `json:"goldPerMin,omitempty"`
		Cs struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"cs,omitempty"`
		CsPerMin struct {
			Value        float64 `json:"value,omitempty"`
			MatchID      string  `json:"matchId,omitempty"`
			ChampionName string  `json:"championName,omitempty"`
			GameMode     string  `json:"gameMode,omitempty"`
			GameCreation int64   `json:"gameCreation,omitempty"`
			GameDuration int     `json:"gameDuration,omitempty"`
		} `json:"csPerMin,omitempty"`
		Vision struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int    `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"vision,omitempty"`
		VisionPerMin struct {
			Value        float64 `json:"value,omitempty"`
			MatchID      string  `json:"matchId,omitempty"`
			ChampionName string  `json:"championName,omitempty"`
			GameMode     string  `json:"gameMode,omitempty"`
			GameCreation int     `json:"gameCreation,omitempty"`
			GameDuration int     `json:"gameDuration,omitempty"`
		} `json:"visionPerMin,omitempty"`
		GameDuration struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"gameDuration,omitempty"`
		DoubleKills struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"doubleKills,omitempty"`
		TripleKills struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"tripleKills,omitempty"`
		QuadraKills struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int    `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"quadraKills,omitempty"`
		PentaKills struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int    `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"pentaKills,omitempty"`
	} `json:"records,omitempty"`
	LowRecords struct {
		Kda struct {
			Value        float64 `json:"value,omitempty"`
			MatchID      string  `json:"matchId,omitempty"`
			ChampionName string  `json:"championName,omitempty"`
			GameMode     string  `json:"gameMode,omitempty"`
			GameCreation int64   `json:"gameCreation,omitempty"`
			GameDuration int     `json:"gameDuration,omitempty"`
		} `json:"kda,omitempty"`
		Kills struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"kills,omitempty"`
		Deaths struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"deaths,omitempty"`
		Assists struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"assists,omitempty"`
		Gold struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"gold,omitempty"`
		GoldPerMin struct {
			Value        float64 `json:"value,omitempty"`
			MatchID      string  `json:"matchId,omitempty"`
			ChampionName string  `json:"championName,omitempty"`
			GameMode     string  `json:"gameMode,omitempty"`
			GameCreation int64   `json:"gameCreation,omitempty"`
			GameDuration int     `json:"gameDuration,omitempty"`
		} `json:"goldPerMin,omitempty"`
		Cs struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"cs,omitempty"`
		CsPerMin struct {
			Value        float64 `json:"value,omitempty"`
			MatchID      string  `json:"matchId,omitempty"`
			ChampionName string  `json:"championName,omitempty"`
			GameMode     string  `json:"gameMode,omitempty"`
			GameCreation int64   `json:"gameCreation,omitempty"`
			GameDuration int     `json:"gameDuration,omitempty"`
		} `json:"csPerMin,omitempty"`
		Vision struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"vision,omitempty"`
		VisionPerMin struct {
			Value        float64 `json:"value,omitempty"`
			MatchID      string  `json:"matchId,omitempty"`
			ChampionName string  `json:"championName,omitempty"`
			GameMode     string  `json:"gameMode,omitempty"`
			GameCreation int64   `json:"gameCreation,omitempty"`
			GameDuration int     `json:"gameDuration,omitempty"`
		} `json:"visionPerMin,omitempty"`
		GameDuration struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"gameDuration,omitempty"`
		DoubleKills struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"doubleKills,omitempty"`
		TripleKills struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"tripleKills,omitempty"`
		QuadraKills struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"quadraKills,omitempty"`
		PentaKills struct {
			Value        int    `json:"value,omitempty"`
			MatchID      string `json:"matchId,omitempty"`
			ChampionName string `json:"championName,omitempty"`
			GameMode     string `json:"gameMode,omitempty"`
			GameCreation int64  `json:"gameCreation,omitempty"`
			GameDuration int    `json:"gameDuration,omitempty"`
		} `json:"pentaKills,omitempty"`
	} `json:"lowRecords,omitempty"`
}
