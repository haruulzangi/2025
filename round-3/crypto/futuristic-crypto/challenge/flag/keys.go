package flag

import (
	"crypto/ecdh"
	"crypto/mlkem"
	"crypto/rand"
	"encoding/base64"
	"os"
	"path"
)

func GetDataPath() string {
	path := os.Getenv("DATA_PATH")
	if path == "" {
		return "/data"
	}
	return path
}

const (
	MLKEMKeyFile = "mlkem"
	AliceKeyFile = "x25519-alice"
	BobKeyFile   = "x25519-bob"
)

func generateKeys() error {
	dk, err := mlkem.GenerateKey1024()
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(GetDataPath(), MLKEMKeyFile), []byte(base64.StdEncoding.EncodeToString(dk.Bytes())), 0600)
	if err != nil {
		return err
	}

	curve := ecdh.X25519()
	alicePriv, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(GetDataPath(), AliceKeyFile), []byte(base64.StdEncoding.EncodeToString(alicePriv.Bytes())), 0600)
	if err != nil {
		return err
	}

	bobPriv, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}
	err = os.WriteFile(path.Join(GetDataPath(), BobKeyFile), []byte(base64.StdEncoding.EncodeToString(bobPriv.Bytes())), 0600)
	if err != nil {
		return err
	}
	return nil
}

var fileList = []string{
	MLKEMKeyFile,
	AliceKeyFile,
	BobKeyFile,
}

func EnsureKeysExist() error {
	for _, file := range fileList {
		if _, err := os.Stat(path.Join(GetDataPath(), file)); os.IsNotExist(err) {
			return generateKeys()
		}
	}
	return nil
}

var mlkemDecKey *mlkem.DecapsulationKey1024

func GetMLKEMDecapsulationKey() (*mlkem.DecapsulationKey1024, error) {
	if mlkemDecKey != nil {
		return mlkemDecKey, nil
	}

	dkBytes, err := os.ReadFile(path.Join(GetDataPath(), MLKEMKeyFile))
	if err != nil {
		return nil, err
	}
	decodedKey, err := base64.StdEncoding.DecodeString(string(dkBytes))
	if err != nil {
		return nil, err
	}

	dk, err := mlkem.NewDecapsulationKey1024(decodedKey)
	if err != nil {
		return nil, err
	}
	mlkemDecKey = dk
	return dk, nil
}

func ensureX25519Key(path string) (*ecdh.PrivateKey, error) {
	key, err := os.ReadFile(path)
	if err == nil {
		curve := ecdh.X25519()
		decodedKey, err := base64.StdEncoding.DecodeString(string(key))
		if err != nil {
			return nil, err
		}
		priv, err := curve.NewPrivateKey(decodedKey)
		if err != nil {
			return nil, err
		}
		return priv, nil
	}

	curve := ecdh.X25519()
	priv, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(path, priv.Bytes(), 0600)
	if err != nil {
		return nil, err
	}
	return priv, nil
}

var alicePrivKey *ecdh.PrivateKey

func GetAlicePrivateKey() (*ecdh.PrivateKey, error) {
	if alicePrivKey != nil {
		return alicePrivKey, nil
	}
	priv, err := ensureX25519Key(path.Join(GetDataPath(), AliceKeyFile))
	if err != nil {
		return nil, err
	}
	alicePrivKey = priv
	return priv, nil
}

var bobPrivKey *ecdh.PrivateKey

func GetBobPrivateKey() (*ecdh.PrivateKey, error) {
	if bobPrivKey != nil {
		return bobPrivKey, nil
	}
	priv, err := ensureX25519Key(path.Join(GetDataPath(), BobKeyFile))
	if err != nil {
		return nil, err
	}
	bobPrivKey = priv
	return priv, nil
}
