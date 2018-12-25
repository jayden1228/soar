/*
 * Copyright 2018 Xiaomi, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"fmt"
	"testing"

	"github.com/XiaoMi/soar/common"
)

var update = flag.Bool("update", false, "update .golden files")

func TestMain(m *testing.M) {
	// 初始化 init
	common.BaseDir = common.DevPath
	err := common.ParseConfig("")
	common.LogIfError(err, "init ParseConfig")
	common.Log.Debug("mysql_test init")
	_ = update // check if var success init

	// 分割线
	flag.Parse()
	m.Run()

	// 环境清理
	//
}

func Test_Main(_ *testing.T) {
	common.Log.Debug("Entering function: %s", common.GetFunctionName())
	common.Config.OnlineDSN.Disable = true
	common.Config.LogLevel = 0
	common.Config.Query = "select * from film;alter table city add index idx_country_id(country_id);"
	main()
	common.Log.Debug("Exiting function: %s", common.GetFunctionName())
}

func Test_Main_More(_ *testing.T) {
	common.Log.Debug("Entering function: %s", common.GetFunctionName())
	common.Config.LogLevel = 0
	common.Config.Profiling = true
	common.Config.Explain = true
	common.Config.Query = "select * from film where country_id = 1;use sakila;alter table city add index idx_country_id(country_id);"
	orgRerportType := common.Config.ReportType
	for _, typ := range []string{
		"json", "html", "markdown", "fingerprint", "compress", "pretty", "rewrite",
		"ast", "tiast", "ast-json", "tiast-json", "tokenize",
	} {
		common.Config.ReportType = typ
		main()
	}
	common.Config.ReportType = orgRerportType
	common.Log.Debug("Exiting function: %s", common.GetFunctionName())
}

func Test_Main_checkConfig(t *testing.T) {
	common.Log.Debug("Entering function: %s", common.GetFunctionName())
	if checkConfig() != 0 {
		t.Error("checkConfig error")
	}
	common.Log.Debug("Exiting function: %s", common.GetFunctionName())
}

func Test_Main_initQuery(t *testing.T) {
	common.Log.Debug("Entering function: %s", common.GetFunctionName())
	// direct query
	query := initQuery("select 1")
	if query != "select 1" {
		t.Errorf("want 'select 1', got %s", query)
	}

	// read from file
	initQuery(common.DevPath + "/README.md")

	// TODO: read from stdin
	// initQuery("")
	common.Log.Debug("Exiting function: %s", common.GetFunctionName())
}

func Test_Main_reportTool(t *testing.T) {
	common.Log.Debug("Entering function: %s", common.GetFunctionName())
	orgRerportType := common.Config.ReportType
	types := []string{"html", "md2html", "explain-digest", "chardet", "remove-comment"}
	for _, tp := range types {
		common.Config.ReportType = tp
		fmt.Println(reportTool(tp, []byte{}))
	}
	common.Config.ReportType = orgRerportType
	common.Log.Debug("Exiting function: %s", common.GetFunctionName())
}
