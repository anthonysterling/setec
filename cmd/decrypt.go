package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"setec/internal"

	"github.com/bitnami-labs/sealed-secrets/pkg/crypto"
	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "",
	Long:  "",
	RunE:  decrypt,
}

func init() {
	decryptCmd.Flags().StringP("private-key-path", "", "", "the path to the private key used by the Sealed Secrets Controller")
	_ = decryptCmd.MarkFlagRequired("private-key-path")
	decryptCmd.Flags().StringP("namespace", "", "", "the namespace used by the Sealed Secrets Controller")
	decryptCmd.Flags().StringP("name", "", "", "the name used by the Sealed Secrets Controller")
	decryptCmd.Flags().BoolP("base64-decode", "", true, "should the input be base64 decoded")
}

func decrypt(cmd *cobra.Command, args []string) error {

	path, err := cmd.Flags().GetString("private-key-path")
	if err != nil {
		return err
	}

	key, err := internal.NewPrivateKeyFromFile(path)
	if err != nil {
		return err
	}

	namespace, err := cmd.Flags().GetString("namespace")
	if err != nil {
		return err
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	var label []byte

	if namespace != "" && name != "" {
		label = internal.NewLabel(namespace, name)
	}

	keys := map[string]*rsa.PrivateKey{
		"default": key,
	}

	cipher, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	shouldDecode, err := cmd.Flags().GetBool("base64-decode")
	if err != nil {
		return err
	}

	if shouldDecode {
		data, err := base64.StdEncoding.DecodeString(string(cipher))
		if err != nil {
			return err
		}
		cipher = data
	}

	plain, err := crypto.HybridDecrypt(rand.Reader, keys, cipher, label)
	if err != nil {
		return err
	}

	fmt.Fprint(os.Stdout, string(plain))
	return nil
}
