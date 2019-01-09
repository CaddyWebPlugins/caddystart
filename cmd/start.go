// Copyright © 2018 LEGION MARKET <software@legionmarket.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"fmt"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Runs Caddy Server",
	Long: `If env TESTING has text will use testing to get cert`,
	Run: func(cmd *cobra.Command, args []string) {
		tet := os.Getenv("TESTING")
		switch tet {
		case "true", "True", "TRUE", "test", "Test", "TEST":
			if os.Getenv("DEBUG") == "true"{
				cmd := exec.Command("caddy", "-agree=true" , os.Getenv("EMAIL"), "-conf=/opt/caddy/Caddyfile",
					"-ca=https://acme-staging-v02.api.letsencrypt.org/directory", "-disabled-metrics")
				fmt.Println(cmd)
				fmt.Println("test is true")
			} else {
				cmd := exec.Command("caddy", "-agree=true" , os.Getenv("EMAIL"), "-conf=/opt/caddy/Caddyfile",
					"-ca=https://acme-staging-v02.api.letsencrypt.org/directory", "-disabled-metrics")
				cmd.Run()
			}

		default:
			if os.Getenv("DEBUG") == "true"{
				cmd := exec.Command("caddy", "-agree=true" , os.Getenv("EMAIL"), "-conf=/opt/caddy/Caddyfile", "-disabled-metrics")
				fmt.Println(cmd)
				fmt.Println("test is false")
			} else {
				cmd := exec.Command("caddy", "-agree=true" , os.Getenv("EMAIL"), "-conf=/opt/caddy/Caddyfile", "-disabled-metrics")
				cmd.Run()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

}

/*
    echo "-agree=true -email=ssl@example.com -conf=/opt/caddy/Caddyfile -ca=https://acme-staging-v02.api.letsencrypt.org/directory  -ca=https://acme-staging-v02.api.letsencrypt.org/directory -log=/var/log/caddy/server.log"
    caddy  -agree=true -email=ssl@example.com -conf=/opt/caddy/Caddyfile -quiet=true -ca=https://acme-staging-v02.api.letsencrypt.org/directory -log=/var/log/caddy/server.log
    else
    echo "-agree=true -email=ssl@example.com -conf=/opt/caddy/Caddyfile -quiet=true -log=/var/log/caddy/server.log"
    caddy  -agree=true -email=ssl@example.com -conf=/opt/caddy/Caddyfile -quiet=true -log=/var/log/caddy/server.log

*/