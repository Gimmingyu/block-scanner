package private

import (
	"encoding/asn1"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type asn1EcSig struct {
	R asn1.RawValue
	S asn1.RawValue
}

var secp256k1N = crypto.S256().Params().N
var secp256k1HalfN = new(big.Int).Div(secp256k1N, big.NewInt(2))

func New(sess *session.Session) *kms.KMS {
	return kms.New(sess)
}

func encrypt(svc *kms.KMS, keyID, plaintext []byte) (*kms.EncryptOutput, error) {
	return svc.Encrypt(&kms.EncryptInput{
		KeyId:     aws.String(string(keyID)),
		Plaintext: plaintext,
	})
}

func Encrypt(svc *kms.KMS, keyID string, value any) ([]byte, error) {
	bytes, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal value: %w", err)
	}

	output, err := encrypt(svc, []byte(keyID), bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt value: %w", err)
	}

	return output.CiphertextBlob, nil
}

func decrypt(svc *kms.KMS, keyID string, ciphertext []byte) (*kms.DecryptOutput, error) {
	return svc.Decrypt(&kms.DecryptInput{
		CiphertextBlob: ciphertext,
		KeyId:          aws.String(keyID),
	})
}

func Decrypt(svc *kms.KMS, keyID string, ciphertext []byte) ([]byte, error) {
	output, err := decrypt(svc, keyID, ciphertext)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt value: %w", err)
	}

	return output.Plaintext, nil
}

func getPubKey(svc *kms.KMS, keyID string) (*kms.GetPublicKeyOutput, error) {
	return svc.GetPublicKey(&kms.GetPublicKeyInput{KeyId: aws.String(keyID)})
}

func GetPublicKey(svc *kms.KMS, keyID string) ([]byte, error) {
	output, err := getPubKey(svc, keyID)
	if err != nil {
		return nil, fmt.Errorf("failed to get public key: %w", err)
	}

	return output.PublicKey, nil
}

func sign(svc *kms.KMS, keyID string, message []byte) (*kms.SignOutput, error) {
	return svc.Sign(&kms.SignInput{
		KeyId:            aws.String(keyID),
		Message:          message,
		MessageType:      aws.String("DIGEST"),
		SigningAlgorithm: aws.String("ECDSA_SHA_256"),
	})
}

func Sign(svc *kms.KMS, keyID string, chainID *big.Int, tx *types.Transaction) ([]byte, error) {
	signer := types.LatestSignerForChainID(chainID)
	txHashBytes := signer.Hash(tx).Bytes()
	output, err := sign(svc, keyID, txHashBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %w", err)
	}

	return output.Signature, nil
}
