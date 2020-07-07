package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"setec/internal"

	"github.com/bitnami-labs/sealed-secrets/pkg/crypto"
	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "",
	Long:  "",
	RunE:  encrypt,
}

func init() {
	encryptCmd.Flags().StringP("private-key-path", "", "", "the path to the private key used by the Sealed Secrets Controller")
	_ = encryptCmd.MarkFlagRequired("private-key-path")
	encryptCmd.Flags().StringP("namespace", "", "", "the namespace used by the Sealed Secrets Controller")
	encryptCmd.Flags().StringP("name", "", "", "the name used by the Sealed Secrets Controller")
	encryptCmd.Flags().BoolP("base64-encode", "", true, "should the output be base64 encoded")
}

func encrypt(cmd *cobra.Command, args []string) error {

	path, err := cmd.Flags().GetString("private-key-path")
	if err != nil {
		return err
	}

	key, err := internal.NewPrivateKeyFromFile(path)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(os.Stdin)
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

	cipher, err := crypto.HybridEncrypt(rand.Reader, &key.PublicKey, data, label)
	if err != nil {
		return err
	}

	shouldEncode, err := cmd.Flags().GetBool("base64-encode")
	if err != nil {
		return err
	}

	if shouldEncode {
		str := base64.StdEncoding.EncodeToString(cipher)
		fmt.Fprint(os.Stdout, str)
		return nil
	}

	fmt.Fprint(os.Stdout, string(cipher))
	return nil
}
