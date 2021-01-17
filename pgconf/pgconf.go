package pgconf

import cnf "github.com/Terry-Mao/goconf"

type PgConfig struct {
	Token     string `goconf:"main:token"`
	TargetUid string `goconf:"main:targetuid"`
}

func GetConf() *PgConfig {
	c := cnf.New()
	c.Spliter = "="
	if err := c.Parse("/etc/pgn/pgn.conf"); err != nil {
		panic(err)
	}
	pc := &PgConfig{}
	if err := c.Unmarshal(pc); err != nil {
		panic(err)
	}
	return pc
}
