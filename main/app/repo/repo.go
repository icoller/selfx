/*
 * @Author: coller
 * @Date: 2023-12-20 21:46:14
 * @LastEditors: coller
 * @LastEditTime: 2023-12-27 23:03:49
 * @Desc:
 */
package repo

import (
	"selfx/init/log"
)

func init() {
	MigrateTable()
}

func MigrateTable() {

	if err := Article.MigrateTable(); err != nil {
		log.Error("migrate article table error", log.Err(err))
	}

	if err := Category.MigrateTable(); err != nil {
		log.Error("migrate category table error", log.Err(err))
	}

	if err := Tag.MigrateTable(); err != nil {
		log.Error("migrate tag table error", log.Err(err))
	}

	if err := Mapping.MigrateTable(); err != nil {
		log.Error("migrate mapping table error", log.Err(err))
	}

	if err := Link.MigrateTable(); err != nil {
		log.Error("migrate link table error", log.Err(err))
	}

	if err := Crawl.MigrateTable(); err != nil {
		log.Error("migrate crawl table error", log.Err(err))
	}

	if err := Verify.MigrateTable(); err != nil {
		log.Error("migrate crawl table error", log.Err(err))
	}
}
