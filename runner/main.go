package main

import (
	"database/sql"
	"fmt"
	"github.com/CovenantSQL/CovenantSQL/client"
	"github.com/CovenantSQL/CovenantSQL/utils/log"
	"math/rand"
	"strings"
)

func main() {
	log.SetLevel(log.InfoLevel)
	var config = "./test/service/node_c/config.yaml"
	var dsn = "covenantsql://0a10b74439f2376d828c9a70fd538dac4b69e0f4065424feebc0f5dbc8b34872"
	var password = ""

	// 使用节点配置文件初始化 Golang SDK
	err := client.Init(config, []byte(password))
	if err != nil {
		log.Fatal(err)
	}

	// 连接数据库实例
	db, err := sql.Open("covenantsql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	//Q := `DROP TABLE IF EXISTS cityPop;
	//  CREATE TABLE cityPop (
	//      ID INT,
	//      Name VARCHAR,
	//      CountryCode VARCHAR,
	//      District VARCHAR,
	//      Population INT
	//  );
	//  CREATE INDEX cityCountryCodeIndex ON cityPop ( CountryCode );
	//
	//  DROP TABLE IF EXISTS countryGDP;
	//  CREATE TABLE countryGDP (
	//       ID integer PRIMARY KEY,
	//       CountryCode string NOT NULL,
	//       GDP integer
	//  );
	//  CREATE INDEX countryCountryCodeIndex ON countryGDP ( CountryCode );`
	//
	//// 写入数据
	//_, err = db.Exec(Q)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("\nExec:\n %s\n", Q)
	//Q = `INSERT INTO countryGDP VALUES
	//       (0, "ZWE", 99999),(1, "CHN", 3000000000000),
	//       (2, "PSE", 321312313),(3, "JPN", 300000000);
	//   INSERT INTO cityPop VALUES (707,'Shenzhen','CHN','Guangzhou',99442);
	//   INSERT INTO cityPop VALUES (1074,'Shenzhen','CHN','Guangzhou',353632);
	//   INSERT INTO cityPop VALUES (1591,'Toyama','JPN','Toyama',325790);
	//   INSERT INTO cityPop VALUES (1649,'Takaoka','JPN','Toyama',174380);
	//   INSERT INTO cityPop VALUES (1762,'Takasago','JPN','Hyogo',97632);
	//   INSERT INTO cityPop VALUES (1763,'Fujimi','JPN','Saitama',96972);
	//   INSERT INTO cityPop VALUES (1764,'Urasoe','JPN','Okinawa',96002);
	//   INSERT INTO cityPop VALUES (1765,'Yonezawa','JPN','Yamagata',95592);
	//   INSERT INTO cityPop VALUES (1766,'Konan','JPN','Aichi',95521);
	//   INSERT INTO cityPop VALUES (1767,'Yamatokoriyama','JPN','Nara',95165);
	//   INSERT INTO cityPop VALUES (1768,'Maizuru','JPN','Kyoto',94784);
	//   INSERT INTO cityPop VALUES (1769,'Onomichi','JPN','Hiroshima',93756);
	//   INSERT INTO cityPop VALUES (1770,'Higashimatsuyama','JPN','Saitama',93342);
	//   INSERT INTO cityPop VALUES (2707,'Xai-Xai','MOZ','Gaza',99442);
	//   INSERT INTO cityPop VALUES (4074,'Gaza','PSE','Gaza',353632);
	//   INSERT INTO cityPop VALUES (4077,'Jabaliya','PSE','North Gaza',113901);`
	//_, err = db.Exec(Q)
	//if err != nil {
	//	log.Warn(err)
	//}
	//log.Printf("\nExec:\n %s\n", Q)

	for i := 0; i < 20; i++ {
		var array = make([]string, 0)
		for i := 0; i < 100000; i++ {
			var qq = fmt.Sprintf("(%d,%s,%s,%s,%d)", rand.Int(), fmt.Sprintf("%f", rand.Float64()),
				fmt.Sprintf("%f", rand.Float64()), fmt.Sprintf("%f", rand.Float64()), rand.Int())
			array = append(array, qq)
		}
		var query = strings.Join(array, ",")
		query = "INSERT INTO cityPop VALUES " + query + ";"
		_, err = db.Exec(query)
		if err != nil {
			panic(err)
		}
	}
}
