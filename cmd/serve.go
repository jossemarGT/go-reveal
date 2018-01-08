// Copyright © 2017 Jonnatan Jossemar Cordero Ramirez
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
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/jossemargt/go-reveal/reveal"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		r := http.NewServeMux()
		t, _ := reveal.NewTokenizer("", "", "")

		// ah := reveal.NewAssetsHandler()
		// r.Handle("/static/", http.StripPrefix("/static/", ah))

		basepath, err := os.Getwd()
		if err != nil {
			return err
		}

		if len(args) > 0 {
			if filepath.Ext(args[0]) != "" {
				basepath = filepath.Dir(filepath.Join(basepath, args[0]))
			} else {
				basepath = filepath.Join(basepath, args[0])
			}
		}

		fmt.Fprintf(cmd.OutOrStdout(), "[go-reveal] Using %s as basepath\n", basepath)

		f := reveal.NewRevealHandler(r, t, basepath)

		srv := &http.Server{
			Handler:      f,
			Addr:         "127.0.0.1:8080",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		fmt.Fprintln(cmd.OutOrStdout(), "[go-reveal] Open your browser at http://127.0.0.1:8080")
		return srv.ListenAndServe()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
