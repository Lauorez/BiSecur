package bisecur

import (
	"bisecur/cli"
	"bisecur/sdk"
)

func Logout(localMac [6]byte, mac [6]byte, host string, port int, token uint32) error {
	if token == 0 {
		return nil //fmt.Errorf("invalid token value: 0x%X. Logout request ignored", token)
	}

	client := sdk.NewClient(cli.Log, localMac, mac, host, port, token)
	defer func() {
		err2 := client.Close()
		if err2 != nil {
			cli.Log.Errorf("%v", err2)
		}
	}()
	
	err := client.Open()
	if err != nil {
		return err
	}

	client.SetToken(0) // clear token in local client instance

	err = client.Logout()
	if err != nil {
		return err
	}

	return nil
}
