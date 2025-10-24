package nex

import (
	"fmt"
	"os"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/super-mario-maker-secure/globals"
)

func StartSecureServer() {
	globals.SecureServer = nex.NewServer()
	globals.SecureServer.SetPRUDPVersion(1)
	globals.SecureServer.SetPRUDPProtocolMinorVersion(2)
	globals.SecureServer.SetDefaultNEXVersion(nex.NewPatchedNEXVersion(3, 5, 5, "AX5J"))
	globals.SecureServer.SetKerberosPassword(globals.KerberosPassword)
	globals.SecureServer.SetAccessKey("4d324052")

	globals.SecureServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("==XCX1 - Secure==")
		fmt.Printf("Protocol ID: %d\n", request.ProtocolID())
		fmt.Printf("Method ID: %d\n", request.MethodID())
		fmt.Println("===============")
	})

	// * Register the common handlers first so that they can be overridden if needed
	registerCommonSecureProtocols()
	registerNEXProtocols()

	globals.SecureServer.Listen(fmt.Sprintf(":%s", os.Getenv("PN_XCX_SECURE_SERVER_PORT")))
}
