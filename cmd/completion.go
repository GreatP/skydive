/*
 * Copyright (C) 2017 Red Hat, Inc.
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// BashCompletion skydive root command
var BashCompletion = &cobra.Command{
	Use:          "bash-completion",
	Short:        "generate bash completion helper",
	Long:         "generate bash completion helper (skydive-bash-completion.sh)",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		RootCmd.GenBashCompletionFile("skydive-bash-completion.sh")
		fmt.Println("skydive-bash-completion.sh has been generated")
	},
}
