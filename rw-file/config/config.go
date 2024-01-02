package config

import (
	types "rw-file/types"
)

var Info = []types.Website{
	{
		"Golang", "http://c.biancheng.net/golang/",
		[]string{
			"http://c.biancheng.net/cplus/",
			"http://c.biancheng.net/linux_tutorial/",
		},
	}, {
		"Java",
		"http://c.biancheng.net/java/",
		[]string{
			"http://c.biancheng.net/socket/",
			"http://c.biancheng.net/python/",
		},
	},
}
