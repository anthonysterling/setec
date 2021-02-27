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

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "",
	Long:  "",
	RunE:  encrypt,
}

func init() {
	encryptCmd.Flags().StringP("public-key-path", "", "", "the path to the public key used by the Sealed Secrets Controller")
	encryptCmd.Flags().StringP("public-key-url", "", "", "the url to the public key used by the Sealed Secrets Controller")
	encryptCmd.Flags().StringP("namespace", "", "", "the namespace used by the Sealed Secrets Controller")
	encryptCmd.Flags().StringP("name", "", "", "the name used by the Sealed Secrets Controller")
	encryptCmd.Flags().BoolP("base64-encode", "", true, "should the output be base64 encoded")
}

func encrypt(cmd *cobra.Command, args []string) error {

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

	key, err := getPublicKey(cmd)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	cipher, err := crypto.HybridEncrypt(rand.Reader, key, data, label)
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

func getPublicKey(cmd *cobra.Command) (*rsa.PublicKey, error) {
	path, err := cmd.Flags().GetString("public-key-path")
	if err != nil {
		return nil, err
	}

	if path != "" {
		return internal.NewPublicKeyFromFile(path)
	}

	url, err := cmd.Flags().GetString("public-key-url")
	if err != nil {
		return nil, err
	}

	if url != "" {
		return internal.NewPublicKeyFromURL(url)
	}

	return nil, fmt.Errorf("either public-key-path or public-key-url must be set and be non-empty")
}
