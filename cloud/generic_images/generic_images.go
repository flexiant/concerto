/*
	A generic image is represents an abstract operating system image that can be 'used' to deploy servers across multiple and heterogeneous clouds.
	Thus, it allows the user not to worry about which image to use depending on the provider where the server is deployed.

	The available commands are:
		list	generic images information


	Use "cloud generic_images --help" on the commandline interface for more information about the available subcommands.

	Generic images list

	This action lists the available generic images.

	Usage:

		generic_images list

*/
package generic_images

// import (
// 	"encoding/json"
// 	"fmt"
// 	"github.com/codegangsta/cli"
// 	"github.com/flexiant/concerto/utils"
// 	"github.com/flexiant/concerto/webservice"
// 	"os"
// 	"text/tabwriter"
// )

// func cmdList(c *cli.Context) {
// 	var images []GenericImage

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get("/v1/cloud/generic_images")
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &images)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\r")

// 	for _, image := range images {
// 		fmt.Fprintf(w, "%s\t%s\n", image.Id, image.Name)
// 	}

// 	w.Flush()
// }
