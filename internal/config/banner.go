package config

import (
	"fmt"
	"os"
)

func Banner() {
	fmt.Fprintf(os.Stdout, `
                                
   ___  ___ ________ ___ _ __ __
  / _ \/ _ '/ __/ _ '/  ' \ \ /
 / .__/\_,_/_/  \_,_/_/_/_/_\_\  @CyInnove
/_/                             				
`)
}
